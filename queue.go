package lstructs

type Queue[T any] struct {
	length int
	data   []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{0, nil}
}

func (q *Queue[T]) Enqueue(x T) {
	q.length++
	q.data = append(q.data, x)
}

func (q *Queue[T]) Dequeue() T {
	var x T
	if len(q.data) > 0 {
		q.length--
		x, q.data = q.data[0], q.data[1:]
	}
	return x
}

func (q *Queue[T]) Peek() T {
	var x T
	if len(q.data) > 0 {
		x = q.data[0]
	}
	return x
}

func (q *Queue[T]) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return q.length
}
