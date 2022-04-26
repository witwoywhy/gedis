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
