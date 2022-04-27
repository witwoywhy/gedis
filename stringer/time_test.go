package stringer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

	t.Run("key not expired", func(t *testing.T) {
		key := "anotherkey1"
		want := -1
		str.Set(key, "", 0)

		got := str.TTL(key)
		assert.Equal(t, want, got)
	})

	t.Run("not exists key", func(t *testing.T) {
		key := "anotherkey2"
		want := -2

		got := str.TTL(key)
		assert.Equal(t, want, got)
	})
}
