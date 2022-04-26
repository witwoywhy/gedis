package stringer

type Stringer interface {
	Set(string, string)

	Get(string) string
}

type String struct {
	storage map[string]string
}

func NewString() Stringer {
	return &String{
		storage: make(map[string]string),
	}
}

func (s *String) Set(key, value string) {
	s.storage[key] = value
}

func (s *String) Get(key string) string {
	return s.storage[key]
}
