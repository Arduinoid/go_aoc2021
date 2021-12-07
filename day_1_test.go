package main

import "testing"

func TestCountIncrease(t *testing.T) {
	input := []int{5, 3, 4, 5, 1, 9}
	expectedCount := 3

	result := countIncrease(input)

	if result != expectedCount {
		t.Logf("expected: %d, got: %d", expectedCount, result)
		t.Fail()
	}
}

func TestCountWindowedIncrease(t *testing.T) {
	input := []int{5, 3, 4, 5, 1, 9} // 12, 12, 10, 15
	expectedCount := 1

	result := countWindowedIncrease(input)

	if result != expectedCount {
		t.Logf("expected: %d, got: %d", expectedCount, result)
		t.Fail()
	}
}
