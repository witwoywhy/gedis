package stringer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGetDel(t *testing.T) {
	key := "mykey"
	hello := "Hello"

	var str Stringer = NewString()
	str.Set(key, hello, 0)

	got := str.GetDel(key)
	assert.Equal(t, hello, got)

	got = str.Get(key)
	assert.Zero(t, got)
}
