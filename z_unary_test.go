// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package lua

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_In_String(t *testing.T) {
	m := &Module{Name: "test"}
	m.Register("test1", func(v String) error {
		return nil
	})
	m.Register("test2", func(v String) error {
		return errors.New("boom")
	})

	{ // Happy path
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test1(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background(), newTestValue(TypeString).(String))
		assert.NoError(t, err)
	}

	{ // Invalid argument
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test2(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background(), newTestValue(TypeString).(String))
		assert.Error(t, err)
	}
}

func Test_Out_String(t *testing.T) {
	m := &Module{Name: "test"}
	m.Register("test1", func() (String, error) {
		return newTestValue(TypeString).(String), nil
	})
	m.Register("test2", func() (String, error) {
		return newTestValue(TypeString).(String), errors.New("boom")
	})

	{ // Happy path
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test1(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.NoError(t, err)
	}

	{ // Invalid argument
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test2(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.Error(t, err)
	}
}

func Test_In_Number(t *testing.T) {
	m := &Module{Name: "test"}
	m.Register("test1", func(v Number) error {
		return nil
	})
	m.Register("test2", func(v Number) error {
		return errors.New("boom")
	})

	{ // Happy path
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test1(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background(), newTestValue(TypeNumber).(Number))
		assert.NoError(t, err)
	}

	{ // Invalid argument
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test2(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background(), newTestValue(TypeNumber).(Number))
		assert.Error(t, err)
	}
}

func Test_Out_Number(t *testing.T) {
	m := &Module{Name: "test"}
	m.Register("test1", func() (Number, error) {
		return newTestValue(TypeNumber).(Number), nil
	})
	m.Register("test2", func() (Number, error) {
		return newTestValue(TypeNumber).(Number), errors.New("boom")
	})

	{ // Happy path
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test1(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.NoError(t, err)
	}

	{ // Invalid argument
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test2(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.Error(t, err)
	}
}

func Test_In_Bool(t *testing.T) {
	m := &Module{Name: "test"}
	m.Register("test1", func(v Bool) error {
		return nil
	})
	m.Register("test2", func(v Bool) error {
		return errors.New("boom")
	})

	{ // Happy path
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test1(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background(), newTestValue(TypeBool).(Bool))
		assert.NoError(t, err)
	}

	{ // Invalid argument
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test2(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background(), newTestValue(TypeBool).(Bool))
		assert.Error(t, err)
	}
}

func Test_Out_Bool(t *testing.T) {
	m := &Module{Name: "test"}
	m.Register("test1", func() (Bool, error) {
		return newTestValue(TypeBool).(Bool), nil
	})
	m.Register("test2", func() (Bool, error) {
		return newTestValue(TypeBool).(Bool), errors.New("boom")
	})

	{ // Happy path
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test1(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.NoError(t, err)
	}

	{ // Invalid argument
		s, err := FromString("", `
		local api = require("test")
		function main(input)
			return api.test2(input)
		end`, m)
		assert.NotNil(t, s)
		assert.NoError(t, err)
		_, err = s.Run(context.Background())
		assert.Error(t, err)
	}
}
