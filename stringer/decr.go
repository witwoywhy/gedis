package stringer

import (
	"strconv"
)

func (s *String) decrBy(key string, decr int) (int, error) {
	var num int
	r, ok := s.get(key)
	if ok {
		n, err := strconv.Atoi(r.value)
		if err != nil {
			return 0, err
		}
		num = n
	}

	num -= decr
	
	str := strconv.Itoa(num)
	s.Set(key, str)
	return num, nil
}

func (s *String) Decr(key string) (int, error) {
	return s.decrBy(key, 1)
}

func (s *String) DecrBy(key string, decr int) (int, error) {
	return s.decrBy(key, decr)
}
