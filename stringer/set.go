package stringer

import (
	"strings"
	"time"
	"unicode/utf8"
)

func (s *String) Set(key, value string, ttl time.Duration) {
	var r repository
	if repo, ok := s.get(key); ok {
		if !repo.ttl.IsZero() {
			repo.ch <- struct{}{}
		}

		r = repo
	}

	r.value = value

	if ttl != 0 {
		expired := time.Now()
		expired = expired.Add(ttl)

		r.ttl = expired
		r.ch = make(chan struct{})

		go s.expire(key, &r)
	}

	s.storage[key] = r
}

func (s *String) SetEx(key, value string, ttl int) {
	s.Set(key, value, s.secondDuration(ttl))
}

func (s *String) MSet(values map[string]string) {
	for k, v := range values {
		s.Set(k, v, 0)
	}
}

func (s *String) SetRange(key, value string, index int) int {
	v := s.Get(key)
	l := utf8.RuneCountInString(v)
	if l == 0 {
		v = strings.Repeat(" ", index)
		v += value
	} else if index > l {
		diff := index - l
		v += strings.Repeat(" ", diff)
		v += value
	} else {
		v = v[:index] + value
	}

	s.Set(key, v, 0)
	return utf8.RuneCountInString(v)
}

func (s *String) PSetEx(key, value string, ttl int) {
	s.Set(key, value, s.milliSecondDuration(ttl))
}

func (s *String) SetNx(key string, value string) bool {
	_, ok := s.get(key)
	if ok {
		return false
	}

	s.Set(key, value, 0)
	return true
}
