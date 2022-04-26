package stringer

import (
	"context"
	"time"
	"unicode/utf8"
)

type Stringer interface {
	Set(string, string, time.Duration)
	SetEx(string, string, int)
	MSet(map[string]string)

	Get(string) string
	MGet([]string) []string

	Append(string, string) int

	TTL(string) int
	Exists(string) bool
}

type repository struct {
	value string
	ttl   time.Time
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
	if _, ok := s.storage[key]; ok {
		return true
	}

	return false
}

func (s *String) Append(key, value string) int {
	v := s.Get(key)
	v += value
	s.Set(key, v, 0)
	return utf8.RuneCountInString(v)
}