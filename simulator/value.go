package simulator

import "github.com/furudenipa/diceraceDP/config"

func calcRollValue(
	step, square int, state []int,
	prevStateValues *[config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64) float64 {

	var v float64

	if step == 0 {
		return v
	}

	var sv float64
	for i := 1; i <= 6; i++ {
		nextSquare := (square + i) % config.NumSquares
		//fmt.Println(calcStateValue(step-1, nextSquare, state, stateValues))
		sv += calcStateValue(nextSquare, state, prevStateValues)
	}
	v += sv
	v += config.DiceRewardsMap[square]

	return v / 6
}

func calcTicketValue(
	n, square int, state []int,
	currentStateValues *[config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64) float64 {

	state[n]--
	nextSquare := (square + n + 1) % config.NumSquares
	var v float64
	v += calcStateValue(nextSquare, state, currentStateValues)
	v += config.RewardsMap[nextSquare]
	state[n]++

	return v
}

func calcStateValue(nextSquare int, state []int, stateValues *[config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64) float64 {

	var sv float64
	if nextSquare != config.AdvanceSixSquare && nextSquare != config.TicketSquare {
		sv += stateValues[nextSquare][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]]
	} else {
		nextSquare = config.TicketSquare
		var tmp float64
		for ticketTpye := 0; ticketTpye < 6; ticketTpye++ {
			canIncrement := state[ticketTpye] < config.MaxTickets-1
			if canIncrement {
				state[ticketTpye]++
			}
			tmp += stateValues[nextSquare][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]]
			if canIncrement {
				state[ticketTpye]--
			}
		}
		sv += tmp / 6
	}
	return sv
}
