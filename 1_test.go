package main

import "testing"

func TestMatchNumbers1(t *testing.T) {
	result := matchNumbers("1212")
	if result != 6 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 6)
	}
}

func TestMatchNumbers2(t *testing.T) {
	result := matchNumbers("1221")
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestMatchNumbers3(t *testing.T) {
	result := matchNumbers("123425")
	if result != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 4)
	}
}

func TestMatchNumbers4(t *testing.T) {
	result := matchNumbers("123123")
	if result != 12 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 12)
	}
}

func TestMatchNumbers5(t *testing.T) {
	result := matchNumbers("12131415")
	if result != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 4)
	}
}
