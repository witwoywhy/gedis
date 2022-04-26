package stringer

import (
	"fmt"
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

func TestGetRange(t *testing.T) {
	type testCase struct {
		start int
		end   int
		want  string
	}
	tests := []testCase{
		{start: 0, end: 3, want: "This"},
		{start: -3, end: -1, want: "ing"},
		{start: 0, end: -1, want: "This is a string"},
		{start: 10, end: 100, want: "string"},
	}
	key := "mykey"
	value := "This is a string"

	var str Stringer = NewString()
	str.Set(key, value, 0)

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d:%d", test.start, test.end), func(t *testing.T) {
			got := str.GetRange(key, test.start, test.end)
			assert.Equal(t, test.want, got)
		})
	}

	t.Run("not exists key", func(t *testing.T) {
		got := str.GetRange("anotherkey", 100, -100)
		assert.Zero(t, got)
	})
}
