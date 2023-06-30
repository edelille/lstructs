package lstructs

import "testing"

func TestEmptySet(t *testing.T) {
	funcName := "Stack.IsEmpty()"
	x := NewStack[int]()
	if x.IsEmpty() == false {
		t.Fatalf("[func: %s] is not performing as expected\n", funcName)
	}
}

func TestAddSet(t *testing.T) {
	testAmt := 100
	funcName := "Set.Add()"
	x := NewSet[int]()
	for i := 0; i < testAmt; i++ {
		x.Add(i)
		x.Add(i)
	}

	if x.Size() != testAmt {
		t.Fatalf("[func: %s] is not performing as expected; actual: %d\n", funcName, x.Size())
	}
}

func TestRemoveSet(t *testing.T) {
	testAmt := 100
	funcName := "Set.Remove()"
	x := NewSet[int]()

	for i := 0; i <= testAmt; i++ {
		x.Add(i)
	}

	for i := testAmt; i >= 0; i-- {
		x.Remove(i)
	}

	if x.Size() != 0 {
		t.Fatalf("[func: %s] is not performing as expected; actual: %d\n", funcName, x.Size())
	}
}

func TestExistsSet(t *testing.T) {
	testElem := 5
	funcName := "Set.Exists()"
	x := NewSet[int]()
	if x.Has(testElem) {
		t.Fatalf("[func: %s] is not performing as expected; should not find\n", funcName)
	}

	x.Add(testElem)
	if !x.Has(testElem) {
		t.Fatalf("[func: %s] is not performing as expected; should find\n", funcName)
	}
}

func testClearSet(t *testing.T) {
	testAmt := 100
	funcName := "Set.Clear()"
	x := NewSet[int]()
	for i := 0; i < testAmt; i++ {
		x.Add(i)
	}

	x.Clear()
	for i := 0; i < testAmt; i++ {
		if x.Has(i) {
			t.Fatalf("[func: %s] is not performing as expected; should not find\n", funcName)
		}
	}

	if !x.IsEmpty() {
		t.Fatalf("[func: %s] is not performing as expected; should be empty\n", funcName)
	}
}
