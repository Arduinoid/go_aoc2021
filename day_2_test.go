package main

import "testing"

const (
	input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`
)

/*
aim   = 10
depth = 60
horz  = 15
*/

func TestPositionCalc(t *testing.T) {
	exPosCalc := 150
	pos := positionCalc(input)

	if pos != exPosCalc {
		t.Logf("expected: %d, got: %d", exPosCalc, pos)
		t.Fail()
	}
}

func TestPositionAdvCalc(t *testing.T) {
	exPosCalc := 900
	pos := positionAdvCalc(input)

	if pos != exPosCalc {
		t.Logf("expected: %d, got: %d", exPosCalc, pos)
		t.Fail()
	}
}
