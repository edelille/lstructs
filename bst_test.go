package lstructs

import (
	"reflect"
	"testing"
)

func TestBSTSize(t *testing.T) {
	funcName := "BST.Size()"
	b := NewBST[int](func(x, y int) bool { return x < y })

	if b.Size() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on empty\n", funcName)
	}

	b.Insert(1)
	if b.Size() != 1 {
		t.Fatalf("[func: %s] is not performing as expected; failed on 1\n", funcName)
	}

}

func TestBSTInorderTraversal(t *testing.T) {
	funcName := "BST.InorderTraversal()"

	cases := []struct {
		name     string
		input    []int
		expected []int
		less     func(x, y int) bool
	}{
		{
			name:     "empty",
			input:    []int{},
			expected: []int{},
			less:     func(x, y int) bool { return x < y },
		},
		{
			name:     "ascending",
			input:    []int{2, 3, 1, 5, 4},
			expected: []int{1, 2, 3, 4, 5},
			less:     func(x, y int) bool { return x < y },
		},
		{
			name:     "descending",
			input:    []int{2, 3, 1, 5, 4},
			expected: []int{5, 4, 3, 2, 1},
			less:     func(x, y int) bool { return x > y },
		},
	}

	for _, c := range cases {
		b := NewBST[int](c.less)
		for _, x := range c.input {
			b.Insert(x)
		}
		output := b.InorderTraversal()

		if !reflect.DeepEqual(output, c.expected) {
			t.Fatalf("[func: %s] is not performing as expected; failed on %s\n", funcName, c.name)
		}
	}

}

func TestBSTRemove(t *testing.T) {
	funcName := "BST.Remove()"
	cases := []struct {
		name     string
		less     LessFn[int]
		input    []int
		toRemove []int
		expected []int
	}{
		{
			name:     "basic",
			less:     func(x, y int) bool { return x < y },
			input:    []int{1, 2, 3, 4, 5},
			toRemove: []int{2, 4},
			expected: []int{1, 3, 5},
		},
		{
			name:     "large",
			less:     func(x, y int) bool { return x < y },
			input:    []int{5, 4, 7, 4, 2, 1, 2, 3, 6},
			toRemove: []int{4, 2, 1},
			expected: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name:     "noDelete",
			less:     func(x, y int) bool { return x < y },
			input:    []int{1, 2, 3, 4, 5},
			toRemove: []int{6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, c := range cases {
		b := NewBST(c.less)
		for _, x := range c.input {
			b.Insert(x)
		}

		for _, x := range c.toRemove {
			b.Remove(x)
		}

		if !reflect.DeepEqual(b.InorderTraversal(), c.expected) {
			t.Fatalf("[func: %s] is not performing as expected; failed on %s\n", funcName, c.name)
		}
	}
}

func TestBSTRebalance(t *testing.T) {
	funcName := "BST.Rebalance()"

	cases := []struct {
		name  string
		less  LessFn[int]
		input []int
	}{
		{
			name:  "noChange",
			less:  func(x, y int) bool { return x < y },
			input: []int{1, 2, 3, 4, 5},
		},
	}

	for _, c := range cases {
		b := NewBST(c.less)

		for _, x := range c.input {
			b.Insert(x)
		}

		b.Rebalance()

		if !reflect.DeepEqual(b.InorderTraversal(), c.input) {
			t.Fatalf("[func: %s] is not performing as expected; failed on %s\n", funcName, c.name)
		}
	}
}
