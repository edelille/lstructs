package lstructs

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() T {
	var x T
	if len(s.data) > 0 {
		x, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	}
	return x
}

func (s *Stack[T]) Peek() T {
	var x T
	if len(s.data) > 0 {
		x = s.data[len(s.data)-1]
	}
	return x
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
