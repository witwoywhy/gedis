package stringer

import (
	"context"
	"time"
)

type Stringer interface {
	Set(string, string, time.Duration)
	SetEx(string, string, int)
	MSet(map[string]string)

	Get(string) string
	MGet([]string) []string

	TTL(string) int
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

func (s *String) Get(key string) string {
	return s.storage[key].value
}

func (s *String) get(key string) repository {
	return s.storage[key]
}

func (s *String) MGet(keys []string) []string {
	var values []string = make([]string, len(keys))
	for i, key := range keys {
		v := s.Get(key)
		values[i] = v
	}
	return values
}
