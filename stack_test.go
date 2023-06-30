package lstructs

import (
	"testing"
)

func TestEmptyStack(t *testing.T) {
	funcName := "Stack.IsEmpty()"
	x := NewStack[int]()
	if x.IsEmpty() == false {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
}

func TestPushPop(t *testing.T) {
	testAmt := 100
	funcName := "Stack.Push() or Stack.Pop()"
	x := NewStack[int]()
	for i := 0; i <= testAmt; i++ {
		x.Push(i)
	}

	for i := testAmt; i >= 0; i-- {
		if x.Pop() != i {
			t.Fatalf("[func: %s] is not performing as expected; failed on %d\n", funcName, i)
		}
	}
}

func TestPeek(t *testing.T) {
	funcName := "Stack.Peek()"
	x := NewStack[int]()
	if x.Peek() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on empty\n", funcName)
	}
	x.Push(5)
	if x.Peek() != 5 {
		t.Fatalf("[func: %s] is not performing as expected; failed after pushing\n", funcName)
	}
}

func TestLength(t *testing.T) {
	testAmt := 100
	funcName := "Stack.Size()"
	x := NewStack[int]()

	if x.Size() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on empty\n", funcName)
	}

	for i := 1; i < testAmt; i++ {
		x.Push(i)
		if x.Size() != i {
			t.Fatalf("[func: %s] is not performing as expected; failed size: %d\n", funcName, i)
		}
	}
}
