package stringer

import (
	"strconv"
	"unicode/utf8"
)

func (s *String) incrBy(key string, incr int) (int, error) {
	value := s.Get(key)
	if utf8.RuneCountInString(value) == 0 {
		s.Set(key, strconv.Itoa(incr), 0)
		return incr, nil
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	n += incr

	str := strconv.Itoa(n)
	s.Set(key, str, 0)

	return n, nil
}

func (s *String) Incr(key string) (int, error) {
	return s.incrBy(key, 1)
}
