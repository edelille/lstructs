package lstructs

type Set[T comparable] struct {
	Length int
	Data   map[T]interface{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{0, map[T]interface{}{}}
}

func (s *Set[T]) Add(x T) {
	if _, f := s.Data[x]; !f {
		s.Length++
		s.Data[x] = struct{}{}
	}
}

func (s *Set[T]) Remove(x T) {
	if _, f := s.Data[x]; f {
		s.Length--
		delete(s.Data, x)
	}
}

func (s *Set[T]) Clear() {
	for k, _ := range s.Data {
		s.Length--
		delete(s.Data, k)
	}
}

func (s *Set[T]) Has(x T) bool {
	_, f := s.Data[x]
	return f
}

func (s *Set[T]) IsEmpty() bool {
	return s.Length == 0
}

func (s *Set[T]) Size() int {
	return s.Length
}
