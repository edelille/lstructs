package lstructs

import (
	"testing"
)

func TestEmptyQueue(t *testing.T) {
	funcName := "Queue.IsEmpty()"
	x := NewQueue[int]()
	if x.IsEmpty() == false {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
}

func TestEnqueueDequeueQueue(t *testing.T) {
	testAmt := 100
	funcName := "Queue.Enqueue() or Queue.Dequeue()"
	x := NewQueue[int]()

	for i := 0; i < testAmt; i++ {
		x.Enqueue(i)
	}

	for i := 0; i < testAmt; i++ {
		if i != x.Dequeue() {
			t.Fatalf("[func: %s] is not performing as expected; failed on %d\n", funcName, i)
		}
	}
}

func TestPeekQueue(t *testing.T) {
	funcName := "Queue.Peek()"
	x := NewQueue[int]()
	if x.Peek() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on empty\n", funcName)
	}

	x.Enqueue(2)
	if x.Peek() != 2 {
		t.Fatalf("[func: %s] is not performing as expected; failed on peek\n", funcName)
	}
}
