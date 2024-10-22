package mdpsolver

import "github.com/furudenipa/diceraceDP/config"

type stateValues struct {
	stateValues [config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64
}

func (s *stateValues) GetValue(square int, remainingTickets []int) float64 {
	return s.stateValues[square][remainingTickets[0]][remainingTickets[1]][remainingTickets[2]][remainingTickets[3]][remainingTickets[4]][remainingTickets[5]]
}

func (s *stateValues) SetValue(square int, remainingTickets []int, value float64) {
	s.stateValues[square][remainingTickets[0]][remainingTickets[1]][remainingTickets[2]][remainingTickets[3]][remainingTickets[4]][remainingTickets[5]] = value
}

// Calculate the value when the action diceroll is selected
func calcRollValue(
	square, remainingRolls int, remainingTickets []int,
	prevStateValues *stateValues) float64 {

	var v float64

	if remainingRolls == 0 {
		return float64(-1)
	}

	var sv float64
	for i := 1; i <= 6; i++ {
		nextSquare := (square + i) % config.NumSquares
		sv += calcStateValue(nextSquare, remainingTickets, prevStateValues)
	}
	v += sv
	v += config.ExpRewards[square]

	return v / 6
}

// Calculate the value when the action ticket_n is selected
// if use ticket_1, then n = 0
func calcTicketValue(
	n, square int, remainingTickets []int,
	currentStateValues *stateValues) float64 {

	remainingTickets[n]--
	nextSquare := (square + n + 1) % config.NumSquares
	var v float64
	v += calcStateValue(nextSquare, remainingTickets, currentStateValues)
	v += config.Rewards[nextSquare]
	remainingTickets[n]++

	return v
}

// Calculate the state value
func calcStateValue(nextSquare int, remainingTickets []int, stateValues *stateValues) float64 {

	var sv float64
	if nextSquare != config.AdvanceSixSquare && nextSquare != config.TicketSquare {
		sv += stateValues.GetValue(nextSquare, remainingTickets)
	} else {
		nextSquare = config.TicketSquare
		var tmp float64
		for ticketTpye := 0; ticketTpye < 6; ticketTpye++ {
			canIncrement := remainingTickets[ticketTpye] < config.MaxTickets-1
			if canIncrement {
				remainingTickets[ticketTpye]++
			}
			tmp += stateValues.GetValue(nextSquare, remainingTickets)
			if canIncrement {
				remainingTickets[ticketTpye]--
			}
		}
		sv += tmp / 6
	}
	return sv
}
