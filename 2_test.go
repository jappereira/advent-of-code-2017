package main

import "testing"

func TestComputeEvenlyDivisorsSum(t *testing.T) {
	result := computeEvenlyDivisorsSum("2_test.in")
	if result != 9 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 9)
	}
}