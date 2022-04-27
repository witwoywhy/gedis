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

	t.Run("key is not integer", func(t *testing.T) {
		key := "string"

		str.Set(key, "This is a string", 0)

		_, err := str.Incr(key)
		assert.NotNil(t, err)
	})
}

func TestIncrBy(t *testing.T) {
	type testCase struct {
		key     string
		value   string
		wantStr string
		incr    int
		wantInt int
	}

	var str Stringer = NewString()

	t.Run("exists key", func(t *testing.T) {
		test := testCase{
			key:     "mykey",
			value:   "10",
			wantStr: "15",
			wantInt: 15,
			incr:    5,
		}

		str.Set(test.key, test.value, 0)

		got1, err := str.IncrBy(test.key, test.incr)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("not exists key", func(t *testing.T) {
		test := testCase{
			key:     "anotherkey",
			wantStr: "-5",
			wantInt: -5,
			incr:    -5,
		}

		got1, err := str.IncrBy(test.key, test.incr)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})

	t.Run("not exists key and incr by zero", func(t *testing.T) {
		test := testCase{
			key:     "byzero",
			wantStr: "0",
			wantInt: 0,
			incr:    0,
		}

		got1, err := str.DecrBy(test.key, test.incr)
		assert.Nil(t, err)
		assert.Equal(t, test.wantInt, got1)

		got2 := str.Get(test.key)
		assert.Equal(t, test.wantStr, got2)
	})
}

func TestIncrByFloat(t *testing.T) {
	var str Stringer = NewString()

	t.Run("exists key", func(t *testing.T) {
		key := "key1"
		str.Set(key, "10.50", 0)

		got, err := str.IncrByFloat(key, 0.1)
		assert.Nil(t, err)
		assert.Equal(t, "10.6", got)

		got, err = str.IncrByFloat(key, -5)
		assert.Nil(t, err)
		assert.Equal(t, "5.6", got)
	})

	t.Run("not exists key", func(t *testing.T) {
		got, err := str.IncrByFloat("key2", -1)
		assert.Nil(t, err)
		assert.Equal(t, "-1", got)
	})

	t.Run("value is string", func(t *testing.T) {
		key := "string"
		str.Set(key, "hello", 0)

		_, err := str.IncrByFloat(key, 1.55)
		assert.NotNil(t, err)
	})
}
