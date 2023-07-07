package lstructs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeapSize(t *testing.T) {
	funcName := "Heap.Size()"

	cases := []struct {
		name     string
		less     LessFn[int]
		input    []int
		expected int
	}{
		{
			name:     "empty",
			less:     func(x, y int) bool { return x < y },
			input:    []int{},
			expected: 0,
		},
		{
			name:     "sizeOne",
			less:     func(x, y int) bool { return x < y },
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "sizeTen",
			less:     func(x, y int) bool { return x < y },
			input:    []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: 10,
		},
	}

	for _, c := range cases {
		h := NewHeap(c.less)

		for _, x := range c.input {
			h.Insert(x)
		}

		if h.Size() != c.expected {
			t.Fatalf("[func: %s] is not performing as expected; failed on %s\n", funcName, c.name)
		}
	}
}

func TestHeapPeek(t *testing.T) {
	funcName := "Heap.Peek()"

	h := NewHeap(func(x, y int) bool { return x < y })

	if h.Peek() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on empty\n", funcName)
	}

	h.Insert(1)
	if h.Peek() != 1 {
		t.Fatalf("[func: %s] is not performing as expected; failed on 1\n", funcName)
	}
}

func TestHeapSort(t *testing.T) {
	funcName := "Heap.Pop()"

	cases := []struct {
		name     string
		less     LessFn[int]
		input    []int
		expected []int
	}{
		{
			name:     "minHeap",
			less:     func(x, y int) bool { return x < y },
			input:    []int{1, 3, 4, 2, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "maxHeap",
			less:     func(x, y int) bool { return x > y },
			input:    []int{1, 3, 4, 2, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, c := range cases {
		h := NewHeap(c.less)

		for _, x := range c.input {
			h.Insert(x)
		}

		test := []int{}
		for !h.IsEmpty() {
			test = append(test, h.Pop())
		}

		if !reflect.DeepEqual(test, c.expected) {
			fmt.Println(test, c.expected)
			t.Fatalf("[func: %s] is not performing as expected; failed on %s\n", funcName, c.name)
		}
	}
}

func TestHeapDelete(t *testing.T) {
	funcName := "Heap.Delete()"

	cases := []struct {
		name     string
		less     LessFn[int]
		input    []int
		toDelete []int
		expected []int
	}{
		{
			name:     "deleteNothing",
			less:     func(x, y int) bool { return x < y },
			input:    []int{1, 3, 4, 2, 5},
			toDelete: []int{6, 7},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "deleteSomething",
			less:     func(x, y int) bool { return x > y },
			input:    []int{1, 3, 4, 2, 5},
			toDelete: []int{2, 2},
			expected: []int{5, 4, 3, 1},
		},
		{
			name:     "deleteAll",
			less:     func(x, y int) bool { return x > y },
			input:    []int{1, 3, 4, 2, 5},
			toDelete: []int{2, 2, 5, 4, 3, 1},
			expected: []int{},
		},
	}

	for _, c := range cases {
		h := NewHeap(c.less)

		for _, x := range c.input {
			h.Insert(x)
		}

		for _, x := range c.toDelete {
			h.Delete(x)
		}

		test := []int{}
		for !h.IsEmpty() {
			test = append(test, h.Pop())
		}

		if !reflect.DeepEqual(test, c.expected) {
			fmt.Println(test, c.expected)
			t.Fatalf("[func: %s] is not performing as expected; failed on %s\n", funcName, c.name)
		}
	}
}
