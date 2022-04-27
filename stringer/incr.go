package stringer

import (
	"strconv"
)

func (s *String) incrBy(key string, incr int) (int, error) {
	var num int
	r, ok := s.get(key)
	if ok {
		n, err := strconv.Atoi(r.value)
		if err != nil {
			return 0, err
		}
		num = n
	}

	num += incr

	str := strconv.Itoa(num)
	s.Set(key, str)

	return num, nil
}

func (s *String) Incr(key string) (int, error) {
	return s.incrBy(key, 1)
}

func (s *String) IncrBy(key string, incr int) (int, error) {
	return s.incrBy(key, incr)
}

func (s *String) IncrByFloat(key string, incr float64) (string, error) {
	var value string
	var num float64

	r, ok := s.get(key)
	if ok {
		value = r.value
		n, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return "", err
		}
		num = n
	}

	num += incr

	str := strconv.FormatFloat(num, 'f', -1, 64)
	s.Set(key, str)
	return str, nil
}
