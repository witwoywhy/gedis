package stringer

import (
	"context"
	"time"
	"unicode/utf8"
)

type Stringer interface {
	Set(string, string, time.Duration)
	SetRange(string, string, int) int
	SetEx(string, string, int)
	SetNx(string, string) bool
	MSet(map[string]string)
	PSetEx(string, string, int)

	Get(string) string
	GetDel(string) string
	GetRange(string, int, int) string
	GetSet(string, string) string
	GetEx(string, int) string
	MGet([]string) []string

	Append(string, string) int
	StrLen(string) int

	Incr(string) (int, error)
	IncrBy(string, int) (int, error)

	Decr(string) (int, error)
	DecrBy(string, int) (int, error)

	TTL(string) int
	Exists(string) bool
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
	s.Set(key, v, 0)
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
