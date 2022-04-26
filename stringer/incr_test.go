package stringer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncr(t *testing.T) {
	type testCase struct {
		key     string
		value   string
		wantStr string
		wantInt int
	}

	var str Stringer = NewString()

	t.Run("exists key", func(t *testing.T) {
		test := testCase{
			key:     "mykey",
			value:   "10",
			wantStr: "11",
			wantInt: 11,
		}

		str.Set(test.key, test.value, 0)

		got1, err := str.Incr(test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("not exists key", func(t *testing.T) {
		test := testCase{
			key:     "anotherkey",
			wantStr: "1",
			wantInt: 1,
		}

		got1, err := str.Incr(test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})
}
