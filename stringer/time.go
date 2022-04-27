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
	r, ok := s.get(key)
	if ok {
		if r.ttl.IsZero() {
			return -2
		}
		
		t := time.Now()
		return int(r.ttl.Sub(t).Seconds())
	}
	return -2
}
