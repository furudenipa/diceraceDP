package mdpsolver

import (
	"testing"

	"github.com/furudenipa/diceraceDP/config"
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
	tests := []struct {
		name             string
		s                *stateValues
		square           int
		remainingRolls   int
		remainingTickets []int
		expectedValue    float64
	}{
		{
			name:             "基本的なテスト",
			s:                &stateValues{},
			square:           0,
			remainingRolls:   1,
			remainingTickets: []int{1, 1, 1, 1, 1, 1},
			expectedValue: func() float64 {
				var sum float64
				for i := 1; i <= 6; i++ {
					sum += config.Rewards[0+i]
				}
				return sum / 6
			}(),
		},
		{
			name:             "squareが0をまたぐ場合",
			s:                &stateValues{},
			square:           16,
			remainingRolls:   1,
			remainingTickets: []int{1, 1, 1, 1, 1, 1},
			expectedValue: func() float64 {
				var sum float64
				for i := 1; i <= 6; i++ {
					sum += config.Rewards[(16+i)%config.NumSquares]
				}
				return sum / 6
			}(),
		},
		{
			name:             "remainingRollが0のとき",
			s:                &stateValues{},
			square:           0,
			remainingRolls:   0,
			remainingTickets: []int{1, 1, 1, 1, 1, 1},
			expectedValue:    -1,
		},
		{
			name: "prevStateがあるとき",
			s: func() *stateValues {
				s := stateValues{}
				for i := 0; i < config.NumSquares; i++ {
					s.SetValue(i, []int{1, 1, 1, 1, 1, 1}, float64(i*i))
				}
				return &s
			}(),
			square:           0,
			remainingRolls:   1,
			remainingTickets: []int{1, 1, 1, 1, 1, 1},
			expectedValue: func() float64 {
				var sum float64
				for i := 1; i <= 6; i++ {
					sum += config.Rewards[0+i] + float64((0+i)*(0+i))
				}
				return sum / 6
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := calcRollValue(tt.square, tt.remainingRolls, tt.remainingTickets, tt.s)
			if value != tt.expectedValue {
				t.Errorf("expected %v, got %v", tt.expectedValue, value)
			} else {
				t.Logf("Test %s passed. Got expected value: %v", tt.name, value)
			}
		})
	}
}

func TestCalcTicketValue(t *testing.T) {
	tests := []struct {
		name             string
		s                *stateValues
		n                int
		square           int
		remainingTickets []int
		expectedValue    float64
	}{
		{
			name:             "基本的なテスト",
			s:                &stateValues{},
			n:                0,
			square:           0,
			remainingTickets: []int{1, 2, 3, 4, 5, 6},
			expectedValue:    config.Rewards[0+1],
		},
		{
			name:             "squareが16の場合",
			s:                &stateValues{},
			n:                3,
			square:           16,
			remainingTickets: []int{1, 2, 3, 4, 5, 6},
			expectedValue:    config.Rewards[(16+4)%config.NumSquares],
		},
		{
			name: "prevStateがあるとき",
			s: func() *stateValues {
				s := stateValues{}
				for i := 0; i < config.NumSquares; i++ {
					s.SetValue(i, []int{1, 2, 3, 4, 5, 6 - 1}, float64(i*i))
				}
				return &s
			}(),
			n:                5,
			square:           0,
			remainingTickets: []int{1, 2, 3, 4, 5, 6},
			expectedValue:    (0+6)*(0+6) + config.Rewards[0+6],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := calcTicketValue(tt.n, tt.square, tt.remainingTickets, tt.s)
			if value != tt.expectedValue {
				t.Errorf("expected %v, got %v", tt.expectedValue, value)
			} else {
				t.Logf("Test %s passed. Got expected value: %v", tt.name, value)
			}
		})
	}
}

func TestCalcStateValue(t *testing.T) {
	tests := []struct {
		name             string
		s                *stateValues
		nextSquare       int
		remainingTickets []int
		expectedValue    float64
	}{
		{
			name:             "基本的なテスト",
			s:                &stateValues{},
			nextSquare:       0,
			remainingTickets: []int{1, 2, 3, 4, 5, 6},
			expectedValue:    0.0, // Adjust this based on your config
		},
		{
			name:             "nextSquareが0をまたぐ場合",
			s:                &stateValues{},
			nextSquare:       16,
			remainingTickets: []int{1, 2, 3, 4, 5, 6},
			expectedValue:    0.0, // Adjust this based on your config
		},
		{
			name: "6進むマスの場合",
			s: func() *stateValues {
				s := stateValues{}
				for i := 0; i < config.NumSquares; i++ {
					s.SetValue(i, []int{1, 2, 3, 4, 5, 6}, float64((i * i)))

					s.SetValue(i, []int{2, 2, 3, 4, 5, 6}, float64(100-(i*i*i)))
					s.SetValue(i, []int{1, 3, 3, 4, 5, 6}, float64(100-(i*i*i)))
					s.SetValue(i, []int{1, 2, 4, 4, 5, 6}, float64(100-(i*i*i)))
					s.SetValue(i, []int{1, 2, 3, 5, 5, 6}, float64(100-(i*i*i)))
					s.SetValue(i, []int{1, 2, 3, 4, 6, 6}, float64(100-(i*i*i)))
					s.SetValue(i, []int{1, 2, 3, 4, 5, 7}, float64(100-(i*i*i)))

				}
				return &s
			}(),
			nextSquare:       12,
			remainingTickets: []int{1, 2, 3, 4, 5, 6},
			expectedValue:    100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := calcStateValue(tt.nextSquare, tt.remainingTickets, tt.s)
			if value != tt.expectedValue {
				t.Errorf("expected %v, got %v", tt.expectedValue, value)
			} else {
				t.Logf("Test %s passed. Got expected value: %v", tt.name, value)
			}
		})
	}
}
