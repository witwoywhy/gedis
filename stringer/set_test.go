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
