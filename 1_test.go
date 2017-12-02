package main

import "testing"

func Test1(t *testing.T) {
	result := compute("1212")
	if result != 6 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 6)
	}
}

func Test2(t *testing.T) {
	result := compute("1221")
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func Test3(t *testing.T) {
	result := compute("123425")
	if result != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 4)
	}
}

func Test4(t *testing.T) {
	result := compute("123123")
	if result != 12 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 12)
	}
}

func Test5(t *testing.T) {
	result := compute("12131415")
	if result != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 4)
	}
}
