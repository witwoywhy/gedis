package stringer

import (
	"strconv"
	"unicode/utf8"
)

func (s *String) decrBy(key string, decr int) (int, error) {
	value := s.Get(key)
	if utf8.RuneCountInString(value) == 0 {
		s.Set(key, "-1", 0)
		return -1, nil
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	n -= decr

	str := strconv.Itoa(n)
	s.Set(key, str, 0)

	return n, nil
}

func (s *String) Decr(key string) (int, error) {
	return s.decrBy(key, 1)
}
