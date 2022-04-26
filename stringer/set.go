package stringer

import "time"

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
