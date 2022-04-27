package stringer

import (
	"context"
	"time"
)

func (s *String) expire(key string, r *repository) {
	ctx, cancel := context.WithDeadline(s.ctx, r.ttl)
	defer cancel()

	for {
		select {
		case <-r.ch:
			r.ttl = time.Time{}
			return
		case <-ctx.Done():
			delete(s.storage, key)
			return
		}
	}
}

func (s *String) TTL(key string) int {
	r, ok := s.get(key)
	if ok {
		if r.ttl.IsZero() {
			return -1
		}

		t := time.Now()
		return int(r.ttl.Sub(t).Seconds())
	}
	return -2
}
