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

func TestMGet(t *testing.T) {
	type testCase struct {
		key   string
		value string
	}
	key1 := testCase{key: "key1", value: "Hello"}
	key2 := testCase{key: "key2", value: "World"}

	var str Stringer = NewString()

	str.Set(key1.key, key1.value, 0)
	str.Set(key2.key, key2.value, 0)

	t.Run("can mget", func(t *testing.T) {
		keys := []string{key1.key, "anotherkey", key2.key}
		want := []string{key1.value, "", key2.value}

		got := str.MGet(keys)
		assert.Equal(t, want, got)
	})
}
