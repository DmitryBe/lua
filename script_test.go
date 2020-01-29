// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package lua

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newScript(file string) (*Script, error) {
	f, _ := os.Open(file)
	return FromReader("test.lua", f, testModule())
}

type Person struct {
	Name string
	Age  int
}

// Benchmark_Serial/fib-8         	 5870025	       203 ns/op	      16 B/op	       2 allocs/op
// Benchmark_Serial/empty-8       	 8592448	       137 ns/op	       0 B/op	       0 allocs/op
// Benchmark_Serial/update-8      	 1000000	      1069 ns/op	     224 B/op	      14 allocs/op
// Benchmark_Serial/module-8      	 1900801	       629 ns/op	     160 B/op	       8 allocs/op
func Benchmark_Serial(b *testing.B) {
	b.Run("fib", func(b *testing.B) {
		s, _ := newScript("fixtures/fib.lua")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Run(context.Background(), 1)
		}
	})

	b.Run("empty", func(b *testing.B) {
		s, _ := newScript("fixtures/empty.lua")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Run(context.Background())
		}
	})

	b.Run("update", func(b *testing.B) {
		s, _ := newScript("fixtures/update.lua")
		input := &Person{Name: "Roman"}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Run(context.Background(), input)
		}
	})

}

// Benchmark_Module/echo-8         	 3467438	       346 ns/op	      48 B/op	       3 allocs/op
// Benchmark_Module/hash-8         	 3323775	       359 ns/op	      48 B/op	       5 allocs/op
func Benchmark_Module(b *testing.B) {
	b.Run("echo", func(b *testing.B) {
		s, _ := newScript("fixtures/echo.lua")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Run(context.Background(), "abc")
		}
	})

	b.Run("hash", func(b *testing.B) {
		s, _ := newScript("fixtures/hash.lua")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Run(context.Background(), "abc")
		}
	})
}

// Benchmark_Fib_Parallel-8   	 4951480	       242 ns/op	      16 B/op	       2 allocs/op
func Benchmark_Fib_Parallel(b *testing.B) {
	s, _ := newScript("fixtures/fib.lua")
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Run(context.Background(), 1)
		}
	})
}

func Test_Update(t *testing.T) {
	s, err := newScript("fixtures/update.lua")
	assert.NoError(t, err)
	assert.Equal(t, "test.lua", s.Name())

	input := &Person{
		Name: "Roman",
	}
	out, err := s.Run(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, TypeString, out.Type())
	assert.Equal(t, "Updated", input.Name)
	assert.Equal(t, "Updated", out.String())
}

func Test_Fib(t *testing.T) {

	s, err := newScript("fixtures/fib.lua")
	assert.NoError(t, err)

	out, err := s.Run(context.Background(), 10)
	assert.NoError(t, err)
	assert.Equal(t, TypeNumber, out.Type())
	assert.Equal(t, Number(89), out.(Number))
	assert.Equal(t, "89", out.String())

	err = s.Close()
	assert.NoError(t, err)
}

func Test_Empty(t *testing.T) {
	s, err := newScript("fixtures/empty.lua")
	assert.NoError(t, err)

	out, err := s.Run(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, TypeNil, out.Type())
}

func Test_Print(t *testing.T) {
	s, err := newScript("fixtures/print.lua")
	assert.NoError(t, err)

	out, err := s.Run(context.Background(), &Person{
		Name: "Roman",
	})
	assert.NoError(t, err)
	assert.Equal(t, TypeString, out.Type())
	assert.Equal(t, "Hello, Roman!", out.String())
}

func Test_InvalidScript(t *testing.T) {
	_, err := FromString("", `
	xxx main()
		local x = 1
	end`)
	assert.Error(t, err)
}

func Test_NoMain(t *testing.T) {
	{
		s := new(Script)
		_, err := s.Run(context.Background())
		assert.Error(t, err)
	}

	{
		_, err := FromString("", `main = 1`)
		assert.Error(t, err)
	}

	{
		_, err := FromString("", `
		function notmain()
			local x = 1
		end`)
		assert.Error(t, err)
	}

	{
		_, err := FromString("", `
		function xxx()
			local x = 1
		end`)
		assert.Error(t, err)
	}
}

func Test_Error(t *testing.T) {
	{
		_, err := FromString("", `
		error() 
		function main()
			local x = 1
		end`)
		assert.Error(t, err)
	}

	{
		s, err := FromString("", `
		function main()
			error() 
		end`)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.Error(t, err)
	}
}

func Test_JSON(t *testing.T) {
	input := map[string]interface{}{
		"a": 123,
		"b": "hello",
		"c": 10.15,
		"d": true,
		"e": &Person{Name: "Roman", Age: 15},
	}

	s, err := newScript("fixtures/json.lua")
	assert.NoError(t, err)

	out, err := s.Run(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, TypeString, out.Type())
	assert.Equal(t, `{"a":123,"b":"hello","c":10.15,"d":true,"e":{"Name":"Roman","Age":15}}`,
		out.String())
}
