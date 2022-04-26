package stringer

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
