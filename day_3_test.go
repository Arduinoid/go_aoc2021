package main

import "testing"

func TestPowerConsumption(t *testing.T) {
	// Arrange
	expected := 90
	input := []byte{
		0b0100,
		0b1101,
		0b0101,
		0b0010,
		0b1100,
	} //  0100

	// Act
	p := powerConsumption(input)

	// Assert
	if p != expected {
		t.Logf("expected: %d, got: %d", expected, p)
		t.Fail()
	}
}
