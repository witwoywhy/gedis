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

func TestTTL(t *testing.T) {
	type testCase struct {
		key   string
		value string
		ttl   int
		want  int
	}
	test := testCase{key: "mykey", value: "Hello", ttl: 10}

	var str Stringer = NewString()

	t.Run("set value", func(t *testing.T) {
		str.SetEx(test.key, test.value, test.ttl)

		test.want = 9
		got := str.TTL(test.key)
		assert.Equal(t, test.want, got)
	})

	t.Run("after 5 sec", func(t *testing.T) {
		time.Sleep(time.Second * 5)

		test.want = 4
		got := str.TTL(test.key)
		assert.Equal(t, test.want, got)
	})
}
