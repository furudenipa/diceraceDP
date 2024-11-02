package mdpsolver

import (
	"testing"
)

func TestStateValues_GetValue(t *testing.T) {
	s := &stateValues{}
	square := 0
	remainingTickets := []int{1, 2, 3, 4, 5, 6}
	expectedValue := 1.23
	s.SetValue(square, remainingTickets, expectedValue)

	value := s.GetValue(square, remainingTickets)
	if value != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, value)
	}
}

func TestCalcRollValue(t *testing.T) {
	s := &stateValues{}
	square := 0
	remainingRolls := 1
	remainingTickets := []int{1, 2, 3, 4, 5, 6}
	expectedValue := float64(-1)

	value := calcRollValue(square, remainingRolls, remainingTickets, s)
	if value != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, value)
	}
}

func TestCalcTicketValue(t *testing.T) {
	s := &stateValues{}
	square := 0
	n := 0
	remainingTickets := []int{1, 2, 3, 4, 5, 6}
	expectedValue := 0.0 // Adjust this based on your config

	value := calcTicketValue(n, square, remainingTickets, s)
	if value != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, value)
	}
}

func TestCalcStateValue(t *testing.T) {
	s := &stateValues{}
	nextSquare := 0
	remainingTickets := []int{1, 2, 3, 4, 5, 6}
	expectedValue := 0.0 // Adjust this based on your config

	value := calcStateValue(nextSquare, remainingTickets, s)
	if value != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, value)
	}
}
