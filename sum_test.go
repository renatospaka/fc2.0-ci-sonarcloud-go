package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	if result != 5 {
		t.Error("The result should be 5.")
	}
}

func TestSub(t *testing.T) {
	result := Sub(3, 2)
	if result != 1 {
		t.Error("The result should be 1.")
	}
}

func TestTimes(t *testing.T) {
	result := Times(2, 3)
	if result != 6 {
		t.Error("The result should be 6.")
	}
}

func TestDiv(t *testing.T) {
	result := Div(6, 3)
	if result != 2 {
		t.Error("The result should be 2.")
	}
}

func TestSumX(t *testing.T) {
	result := SumX(2, 3)
	if result != 7 {
		t.Error("The result should be 7.")
	}
}