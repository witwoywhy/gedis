package stringer

import (
	"strings"
	"time"
	"unicode/utf8"
)

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

func (s *String) SetEx(key, value string, ttl int) {
	ttlDuration := time.Duration(ttl)
	s.Set(key, value, ttlDuration)
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
