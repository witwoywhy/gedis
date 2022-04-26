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
		str.Set(test.key, test.value)
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
