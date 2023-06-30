package lstructs

type Queue[T any] struct {
	Length int
	Data   []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{0, nil}
}

func (q *Queue[T]) Enqueue(x T) {
	q.Length++
	q.Data = append(q.Data, x)
}

func (q *Queue[T]) Dequeue() T {
	var x T
	if len(q.Data) > 0 {
		q.Length--
		x, q.Data = q.Data[0], q.Data[1:]
	}
	return x
}

func (q *Queue[T]) Peek() T {
	var x T
	if len(q.Data) > 0 {
		x = q.Data[0]
	}
	return x
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Data) == 0
}

func (q *Queue[T]) Size() int {
	return q.Length
}
