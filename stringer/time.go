package stringer

import (
	"context"
	"time"
)

func (s *String) expired(key string, expired time.Time) {
	ctx, cancel := context.WithDeadline(s.ctx, expired)
	defer cancel()

	for {
		<-ctx.Done()
		delete(s.storage, key)
		return
	}
}

func (s *String) TTL(key string) int {
	r := s.get(key)
	t := time.Now()
	return int(r.ttl.Sub(t).Seconds())
}
