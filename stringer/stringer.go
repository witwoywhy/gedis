package stringer

import (
	"context"
	"time"
)

type Stringer interface {
	Set(string, string, time.Duration)
	SetEx(string, string, int)

	Get(string) string

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

func (s *String) expired(key string, expired time.Time) {
	ctx, cancel := context.WithDeadline(s.ctx, expired)
	defer cancel()

	for {
		<-ctx.Done()
		delete(s.storage, key)
		return
	}
}

func (s *String) Get(key string) string {
	return s.storage[key].value
}

func (s *String) get(key string) repository {
	return s.storage[key]
}

func (s *String) TTL(key string) int {
	r := s.get(key)
	t := time.Now()
	return int(r.ttl.Sub(t).Seconds())
}
