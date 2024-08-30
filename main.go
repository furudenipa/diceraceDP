package main

import (
	"fmt"
)

func main() {
	var stateValues [numSteps][numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64
	for step := 0; step < numSteps; step++ {
		for ticketSum := 0; ticketSum < (maxTickets-1)*6; ticketSum++ {
			combinations := [][]int{}
			generateCombination(0, 0, ticketSum, []int{}, &combinations)
			for _, state := range combinations {
				square := 0
				rollValue := calcRollValue(step, square, state, &stateValues)
				var ticketValues [6]float64
				for n := 0; n < 6; n++ {
					if state[n] > 0 {
						ticketValues[n] = calcTicketValue(n, step, square, state, &stateValues)
					}
				}
				newValue := max(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])
				stateValues[step][0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = newValue
			}

			for _, state := range combinations{
				for square:= 1; square <= numSquares; square++{

					square := 0
					rollValue := calcRollValue(step, square, state, &stateValues)
					var ticketValues [6]float64
					for n := 0; n < 6; n++ {
						if state[n] > 0 {
							ticketValues[n] = calcTicketValue(n, step, square, state, &stateValues)
						}
					}
					newValue := max(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])
					stateValues[step][0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = newValue
	
				}
			}
		}
		fmt.Println(step)
	}
}
