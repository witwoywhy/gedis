package stringer

import (
	"context"
	"time"
)

type Stringer interface {
	Set(string, string, time.Duration)
	SetEx(string, string, int)

	Get(string) string
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

func (s *String) Set(key, value string, ttl time.Duration) {
	var r repository
	r.value = value

	if ttl != 0 {
		expired := time.Now()
		expired = expired.Add(time.Second * ttl)
		r.ttl = expired

		go s.expired(key, expired)
	}

	s.storage[key] = r
}

func (s *String) Get(key string) string {
	return s.storage[key].value
}

func (s *String) SetEx(key, value string, ttl int) {
	ttlDuration := time.Duration(ttl)
	s.Set(key, value, ttlDuration)
}
