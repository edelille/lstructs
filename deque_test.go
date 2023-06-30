package lstructs

import "testing"

func TestEmptyDeque(t *testing.T) {
	funcName := "Deque.IsEmpty()"
	x := NewDeque[int]()
	if x.IsEmpty() == false {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
	x.PushLeft(1)
	if x.IsEmpty() == true {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
}

func TestStackDeque(t *testing.T) {
	testAmt := 100
	funcName := "Deque.QueueFns()"
	x := NewDeque[int]()

	for i := 0; i <= testAmt; i++ {
		x.PushLeft(i)
	}

	for i := testAmt; i >= 0; i-- {
		if i != x.PopLeft() {
			t.Fatalf("[func: %s] is not performing as expected; failed on stack\n", funcName)
		}
	}
}

func TestPushRightPopRightDeque(t *testing.T) {
	testAmt := 100
	funcName := "Deque.StackFns()"
	x := NewDeque[int]()

	for i := 0; i < testAmt; i++ {
		x.PushLeft(i)
	}

	for i := 0; i < testAmt; i++ {
		if i != x.PopRight() {
			t.Fatalf("[func: %s] is not performing as expected; failed on queue\n", funcName)
		}
	}
}

func TestPeekLeftDeque(t *testing.T) {
	funcName := "Deque.Peek()"
	x := NewDeque[int]()
	if x.PeekLeft() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on empty\n", funcName)
	}

	x.PushLeft(2)
	if x.PeekLeft() != 2 || x.PeekLeft() != x.PeekRight() {
		t.Fatalf("[func: %s] is not performing as expected; failed on peek\n", funcName)
	}
}
