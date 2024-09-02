package main

func calcRollValue(
	step, square int, state []int,
	prevStateValues *[numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64) float64 {

	var v float64

	if step == 0 {
		return v
	}

	var sv float64
	for i := 1; i <= 6; i++ {
		nextSquare := (square + i) % numSquares
		//fmt.Println(calcStateValue(step-1, nextSquare, state, stateValues))
		sv += calcStateValue(nextSquare, state, prevStateValues)
	}
	v += sv
	v += diceRewardsMap[square]

	return v / 6
}

func calcTicketValue(
	n, square int, state []int,
	currentStateValues *[numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64) float64 {

	state[n]--
	nextSquare := (square + n + 1) % numSquares
	var v float64
	v += calcStateValue(nextSquare, state, currentStateValues)
	v += rewardsMap[nextSquare]
	state[n]++

	return v
}

func calcStateValue(nextSquare int, state []int, stateValues *[numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64) float64 {

	var sv float64
	if nextSquare != advanceSixSquare && nextSquare != ticketSquare {
		sv += stateValues[nextSquare][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]]
	} else {
		nextSquare = ticketSquare
		var tmp float64
		for ticketTpye := 0; ticketTpye < 6; ticketTpye++ {
			canIncrement := state[ticketTpye] < maxTickets-1
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
