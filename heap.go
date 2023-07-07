package lstructs

type Heap[T comparable] struct {
	Length int
	Data   []T

	level     int
	maxLength int
	less      LessFn[T]
}

func NewHeap[T comparable](less LessFn[T]) *Heap[T] {
	return &Heap[T]{0, nil, 0, 0, less}
}

func (h *Heap[T]) Insert(x T) {
	if h.Length == h.maxLength {
		// Recreate the new data array, copy over
		tmp := h.Data
		newLen := fastPow2(h.level+1) - 1
		h.Data = make([]T, newLen)
		h.level++
		h.maxLength = newLen
		copy(h.Data, tmp)
	}

	h.Data[h.Length] = x

	// Heapify
	heapUp(h.Data, h.Length, h.less)

	h.Length++
}

// The delete function will return True on successful deletion
func (h *Heap[T]) Delete(x T) bool {
	var null T
	for i, y := range h.Data {
		if x == y {
			if i == h.Length-1 {
				h.Data[i] = null
			} else {
				h.Data[i] = h.Data[h.Length-1]
				heapUp(h.Data, i, h.less)
			}

			h.Length--

			return true
		}
	}

	return false
}

func (h *Heap[T]) Pop() T {
	if h.Length == 0 {
		var nullValue T
		return nullValue
	}

	var null T
	tmp := h.Data[0]
	h.Data[0] = h.Data[h.Length-1]
	h.Data[h.Length-1] = null
	h.Length--

	heapDown(h.Data, h.Length, 0, h.less)

	return tmp
}

func (h *Heap[T]) Peek() T {
	if len(h.Data) == 0 {
		var nullValue T
		return nullValue
	}

	return h.Data[0]
}

func (h *Heap[T]) Size() int {
	return h.Length
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Length == 0
}

func heapUp[T comparable](arr []T, i int, less LessFn[T]) {
	if i == 0 {
		return
	}

	parent := (i - 1) / 2

	if less(arr[i], arr[parent]) {
		Swap(arr, i, parent)
		heapUp(arr, parent, less)
	}

}

func heapDown[T comparable](arr []T, n int, i int, less LessFn[T]) {
	left := i*2 + 1
	right := left + 1
	min := -1

	if left < n && less(arr[left], arr[i]) {
		min = left
	} else {
		min = i
	}

	if right < n && less(arr[right], arr[min]) {
		min = right
	}

	if min != i {
		Swap(arr, i, min)
		heapDown(arr, n, min, less)
	}
}
