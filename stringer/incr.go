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

func (s *String) IncrBy(key string, incr int) (int, error) {
	return s.incrBy(key, incr)
}

func (s *String) IncrByFloat(key string, incr float64) (string, error) {
	var value string
	var float float64

	r, ok := s.get(key)
	if ok {
		value = r.value
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return "", err
		}
		float = n
	}

	float += incr

	str := strconv.FormatFloat(float, 'f', -1, 64)
	s.Set(key, str, 0)
	return str, nil
}
