package stringer

import "unicode/utf8"

func (s *String) Get(key string) string {
	return s.storage[key].value
}

func (s *String) get(key string) repository {
	return s.storage[key]
}

func (s *String) MGet(keys []string) []string {
	var values []string = make([]string, len(keys))
	for i, key := range keys {
		v := s.Get(key)
		values[i] = v
	}
	return values
}

func (s *String) GetDel(key string) string {
	value := s.Get(key)
	delete(s.storage, key)
	return value
}

func (s *String) GetRange(key string, start, end int) string {
	value := s.Get(key)
	l := utf8.RuneCountInString(value)

	switch {
	case l == 0:
		return ""
	case start < 0:
		start = l + start
		fallthrough
	case end < 0:
		end = l + end + 1
	case end > l:
		end = l
	case end > 0:
		end++
	}

	return value[start:end]
}
