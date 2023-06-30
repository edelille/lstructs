package lstructs

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	funcName := "Queue.IsEmpty()"
	x := NewQueue[int]()
	if x.IsEmpty == false {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
}

func TestEnqueueDequeue(t *testing.T) {
	funcName := "Queue.Enqueue() or Queue.Dequeue()"
	x := NewQueue[int]()
	if x.IsEmpty == false {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
}
