package stringer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetandGet(t *testing.T) {
	type testCase struct {
		key   string
		value string
		want  string
	}
	test := testCase{key: "mykey", value: "Hello", want: "Hello"}

	var str Stringer = NewString()

	t.Run("can set", func(t *testing.T) {
		str.Set(test.key, test.value, 0)
	})

	t.Run("can get", func(t *testing.T) {
		got := str.Get(test.key)
		assert.Equal(t, test.want, got)
	})

	t.Run("get exist key", func(t *testing.T) {
		test := testCase{key: "anotherkey", want: ""}

		got := str.Get(test.key)
		assert.Equal(t, test.want, got)
	})
}

func TestExists(t *testing.T) {
	var str Stringer = NewString()

	t.Run("exists key", func(t *testing.T) {
		type testCase struct {
			key   string
			value string
		}
		test := testCase{key: "mykey", value: "Hello"}

		str.Set(test.key, test.value, 0)

		got := str.Exists(test.key)
		assert.Equal(t, true, got)
	})

	t.Run("not exists key", func(t *testing.T) {
		key := "anotherkey"
		want := false

		got := str.Exists(key)
		assert.Equal(t, want, got)
	})
}
