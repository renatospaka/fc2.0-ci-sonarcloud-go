package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	if result != 5 {
		t.Error("The result should be 5.")
	}
}
func TestSub(t *testing.T) {
	result := Sub(2, 3)
	if result != -1 {
		t.Error("The result should be -1.")
	}
}

func TestTimes(t *testing.T) {
	result := Times(2, 3)
	if result != 6 {
		t.Error("The result should be 6.")
	}
}
