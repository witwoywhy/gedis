package stringer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetEx(t *testing.T) {
	type testCase struct {
		key   string
		value string
		want  string
		ttl   int
	}
	test := testCase{key: "mykey", value: "Hello", ttl: 3, want: "Hello"}

	var str Stringer = NewString()

	t.Run("can set with ttl", func(t *testing.T) {
		str.SetEx(test.key, test.value, test.ttl)

		got := str.Get(test.key)
		assert.Equal(t, test.want, got)
	})

	t.Run("can expired", func(t *testing.T) {
		time.Sleep(4 * time.Second)

		got := str.Get(test.key)
		assert.Equal(t, "", got)
	})

	t.Run("can change ttl", func(t *testing.T) {
		test := testCase{key: "change", value: "Hello", ttl: 3, want: "Hello"}
		str.SetEx(test.key, "", test.ttl)

		test.ttl = 10
		str.SetEx(test.key, test.value, test.ttl)

		time.Sleep(3 * time.Second)
		got := str.Get(test.key)
		assert.Equal(t, test.want, got)

		time.Sleep(8 * time.Second)
		got = str.Get(test.key)
		assert.Equal(t, "", got)
	})
}

func TestMSet(t *testing.T) {
	type testCase struct {
		key   string
		value string
	}
	key1 := testCase{key: "key1", value: "Hello"}
	key2 := testCase{key: "key2", value: "World"}
	tests := []testCase{
		key1,
		key2,
	}

	var str Stringer = NewString()

	t.Run("can mset", func(t *testing.T) {
		m := make(map[string]string)
		for _, test := range tests {
			m[test.key] = test.value
		}

		str.MSet(m)
	})

	t.Run("can get", func(t *testing.T) {
		got := str.Get(key1.key)
		assert.Equal(t, key1.value, got)

		got = str.Get(key2.key)
		assert.Equal(t, key2.value, got)
	})
}

func TestSetRange(t *testing.T) {
	type testCase struct {
		key   string
		index int
		want  int
	}

	var str Stringer = NewString()

	t.Run("basic", func(t *testing.T) {
		test := testCase{key: "key1", index: 6, want: 11}
		str.Set(test.key, "Hello World", 0)

		got := str.SetRange(test.key, "Redis", test.index)
		assert.Equal(t, test.want, got)
	})

	t.Run("zero padding", func(t *testing.T) {
		test := testCase{key: "key2", index: 6, want: 11}
		got := str.SetRange(test.key, "Redis", test.index)
		assert.Equal(t, test.want, got)
	})

	t.Run("index over exists", func(t *testing.T) {
		test := testCase{key: "key1", index: 20, want: 25}
		str.Set(test.key, "Hello World", 0)

		got := str.SetRange(test.key, "Redis", test.index)
		assert.Equal(t, test.want, got)
	})
}
