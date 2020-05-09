package main

import "testing"

func Fact(fact int64) (res int64) {
	res = 1
	var i int64
	for i = 1; i <= fact; i++ {
		res *= i
	}

	return
}

func TestFact(t *testing.T) {
	if Fact(0) != 1 {
		t.Error("Test fail", 0)
	}

	if Fact(1) != 1 {
		t.Error("Test fail", 1)
	}

	if Fact(5) != 120 {
		t.Error("Test fail", 5)
	}

	if Fact(6) != 720 {
		t.Error("Test fail", 6)
	}

}
