package lstructs

import "testing"

func TestFastLog2(t *testing.T) {
	funcName := "fastLog2"

	if fastLog2(0) > 0 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=0\n", funcName)
	}

	if fastLog2(2) != 1 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=2\n", funcName)
	}
	if fastLog2(3) != 1 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=3\n", funcName)
	}

	if fastLog2(8) != 3 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=8\n", funcName)
	}

	if fastLog2(1024) != 10 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=1024\n", funcName)
	}
}

func TestFastPow2(t *testing.T) {
	funcName := "fastLog2"

	if fastPow2(0) != 1 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=0\n", funcName)
	}

	if fastPow2(1) != 2 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=1\n", funcName)
	}

	if fastPow2(2) != 4 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=2\n", funcName)
	}

	if fastPow2(10) != 1024 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=10\n", funcName)
	}

	if fastPow2(20) != 1048576 {
		t.Fatalf("[func: %s] is not performing as expected; failed on n=20\n", funcName)
	}
}
