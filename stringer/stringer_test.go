package stringer

import (
	"testing"
	"time"

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
