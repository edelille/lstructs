package lstructs

type Stack[T any] struct {
	Length int
	Data   []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{0, nil}
}

func (s *Stack[T]) Push(x T) {
	s.Length++
	s.Data = append(s.Data, x)
}

func (s *Stack[T]) Pop() T {
	var x T
	if len(s.Data) > 0 {
		s.Length--
		x, s.Data = s.Data[len(s.Data)-1], s.Data[:len(s.Data)-1]
	}
	return x
}

func (s *Stack[T]) Peek() T {
	var x T
	if len(s.Data) > 0 {
		x = s.Data[len(s.Data)-1]
	}
	return x
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length == 0
}

func (s *Stack[T]) Size() int {
	return s.Length
}
