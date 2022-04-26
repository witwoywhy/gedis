package stringer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecr(t *testing.T) {
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
			wantStr: "9",
			wantInt: 9,
		}

		str.Set(test.key, test.value, 0)

		got1, err := str.Decr(test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("not exists key", func(t *testing.T) {
		test := testCase{
			key:     "anotherkey",
			wantStr: "-1",
			wantInt: -1,
		}

		got1, err := str.Decr(test.key)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("key is not integer", func(t *testing.T) {
		key := "string"

		str.Set(key, "This is a string", 0)

		_, err := str.Decr(key)
		assert.NotNil(t, err)
	})
}

func TestDecrBy(t *testing.T) {
	type testCase struct {
		key     string
		value   string
		wantStr string
		decr    int
		wantInt int
	}

	var str Stringer = NewString()

	t.Run("exists key", func(t *testing.T) {
		test := testCase{
			key:     "mykey",
			value:   "10",
			wantStr: "5",
			wantInt: 5,
			decr:    5,
		}

		str.Set(test.key, test.value, 0)

		got1, err := str.DecrBy(test.key, test.decr)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("not exists key", func(t *testing.T) {
		test := testCase{
			key:     "anotherkey",
			wantStr: "5",
			wantInt: 5,
			decr:    -5,
		}

		got1, err := str.DecrBy(test.key, test.decr)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("not exists key and decr by zero", func(t *testing.T) {
		test := testCase{
			key:     "byzero",
			wantStr: "0",
			wantInt: 0,
			decr:    0,
		}

		got1, err := str.DecrBy(test.key, test.decr)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})
}
