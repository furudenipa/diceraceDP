package main

import (
	"fmt"
)

func main() {
	fmt.Println("Settings")
	fmt.Println(" numSquares: ", numSquares)
	fmt.Println(" maxTickets: ", maxTickets)
	fmt.Println(" numSteps:  ", numSteps)
	fmt.Println("----------start----------")
	var currentStateValues [numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64
	var prevStateValues [numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64
	var policy [numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]byte

	for step := 0; step < numSteps; step++ {
		fmt.Println("step: ", step)

		for ticketSum := 0; ticketSum <= (maxTickets-1)*6; ticketSum++ {
			combinations := [][]int{}
			generateCombination(0, 0, ticketSum, []int{}, &combinations)
			for _, state := range combinations {
				square := 0
				rollValue := calcRollValue(step, square, state, &prevStateValues)
				var ticketValues [6]float64
				for n := 0; n < 6; n++ {
					if state[n] > 0 {
						ticketValues[n] = calcTicketValue(n, square, state, &currentStateValues)
					}
				}
				newValue, action := maxIndex(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])
				currentStateValues[0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = newValue
				policy[0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = byte(action)
			}

			for _, state := range combinations {
				for square := 1; square < numSquares; square++ {
					rollValue := calcRollValue(step, square, state, &prevStateValues)
					var ticketValues [6]float64
					for n := 0; n < 6; n++ {
						if state[n] > 0 {
							ticketValues[n] = calcTicketValue(n, square, state, &currentStateValues)
						}
					}
					newValue, action := maxIndex(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])
					currentStateValues[square][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = newValue
					policy[square][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = byte(action)
				}
			}
		}
		for i := 0; i < 18; i++ {
			fmt.Printf("currentStateValues[%d][0][0][1][0][0][0]: %f\n", i, currentStateValues[i][0][0][1][0][0][0])
		}
		fmt.Println("currentStateValues[0][1][0][1][0][1][0]: ", currentStateValues[0][1][0][1][0][1][0])
		fmt.Println("currentStateValues[0][1][1][1][1][1][1]: ", currentStateValues[0][1][1][1][1][1][1])

		prevStateValues = currentStateValues
	}
	fmt.Println("----------end----------")
}

func maxIndex(values ...float64) (float64, int) {
	if len(values) == 0 {
		return 0, -1
	}
	maxValue := values[0]
	maxIndex := 0
	for i, value := range values {
		if value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}

	return maxValue, maxIndex
}
