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
