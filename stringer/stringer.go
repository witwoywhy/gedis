package stringer

import (
	"context"
	"time"
	"unicode/utf8"
)

type Stringer interface {
	Set(key string, value string)
	SetRange(key string, value string, index int) int
	SetEx(key string, value string, ttl int)
	SetNx(key string, value string) bool
	MSet(values map[string]string)
	MSetNx(values map[string]string) bool
	PSetEx(key string, value string, ttl int)

	Get(key string) string
	GetDel(key string) string
	GetRange(key string, start int, end int) string
	GetSet(key string, value string) string
	GetEx(key string, ttl int) string
	MGet(key []string) []string

	Append(key string, value string) int
	StrLen(key string) int

	Incr(key string) (int, error)
	IncrBy(key string, incr int) (int, error)
	IncrByFloat(key string, incr float64) (string, error)

	Decr(key string) (int, error)
	DecrBy(key string, decr int) (int, error)

	TTL(key string) int
	Exists(key string) bool
}

type repository struct {
	value string

	ttl time.Time
	ch  chan struct{}
}

type String struct {
	storage map[string]repository

	ctx context.Context
}

func NewString() Stringer {
	return &String{
		storage: make(map[string]repository),
		ctx:     context.Background(),
	}
}

func (s *String) Exists(key string) bool {
	_, ok := s.storage[key]
	return ok
}

func (s *String) Append(key, value string) int {
	v := s.Get(key)
	v += value
	s.Set(key, v)
	return utf8.RuneCountInString(v)
}

func (s *String) StrLen(key string) int {
	value := s.Get(key)
	return utf8.RuneCountInString(value)
}

func (s *String) secondDuration(second int) time.Duration {
	return time.Duration(second) * time.Second
}
func (s *String) milliSecondDuration(milliSecond int) time.Duration {
	return time.Duration(milliSecond) * time.Millisecond
}
