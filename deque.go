package lstructs

type Deque[T any] struct {
	Length int
	Data   []T
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{0, nil}
}

func (d *Deque[T]) PushLeft(x T) {
	d.Length++
	d.Data = append([]T{x}, d.Data...)
}

func (d *Deque[T]) PushRight(x T) {
	d.Length++
	d.Data = append(d.Data, x)
}

func (d *Deque[T]) PopLeft() T {
	var x T
	if d.Length > 0 {
		d.Length--
		x, d.Data = d.Data[0], d.Data[1:]
	}
	return x
}

func (d *Deque[T]) PopRight() T {
	var x T
	if d.Length > 0 {
		d.Length--
		x, d.Data = d.Data[len(d.Data)-1], d.Data[:len(d.Data)-1]
	}
	return x
}

func (d *Deque[T]) PeekLeft() T {
	var x T
	if d.Length > 0 {
		x = d.Data[0]
	}
	return x
}

func (d *Deque[T]) PeekRight() T {
	var x T
	if d.Length > 0 {
		x = d.Data[0]
	}
	return x
}

func (d *Deque[T]) IsEmpty() bool {
	return d.Length == 0
}

func (d *Deque[T]) Size() int {
	return d.Length
}
