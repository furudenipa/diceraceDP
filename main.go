package main

import (
	"fmt"
)

//var check float64 = 0

func main() {
	// show settings
	var stateValues [numSteps][numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]float64
	var policy [numSteps][numSquares][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets][maxTickets]byte
	for step := 0; step < numSteps; step++ {
		for ticketSum := 0; ticketSum <= (maxTickets-1)*6; ticketSum++ {
			fmt.Println("step:", step, "ticketSum:", ticketSum)
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

				newValue, action := maxIndex(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])
				stateValues[step][0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = newValue
				policy[step][0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = byte(action)

				/* debug
				if check != stateValues[0][0][1][1][1][1][1][1] {
					check = stateValues[0][0][1][1][1][1][1][1]
					fmt.Println("check:", check, "square:", square, "state:", state, "rollValue:", rollValue, "ticketValues:", ticketValues, "action:", action)
				}

				if ticketSum > 2 && ticketSum != 6 {
					break
				}
				if step == 0 || step == 1 {
					fmt.Println("step: ", step, "state:", state, "rollValue:", rollValue, "ticketValues:", ticketValues, "action:", action)
					fmt.Println("value: ", stateValues[step][0][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]])
				}

				if step > 1 {
					fmt.Println("step: ", step, "state:", state, "rollValue:", rollValue, "ticketValues:", ticketValues, "action:", action)
				}

				if ticketSum == 54 {
					fmt.Println(rollValue, ticketValues)
					fmt.Println(newValue, action)
				}
				*/
			}
			for _, state := range combinations {

				for square := 1; square < numSquares; square++ {
					rollValue := calcRollValue(step, square, state, &stateValues)
					var ticketValues [6]float64
					for n := 0; n < 6; n++ {
						if state[n] > 0 {
							ticketValues[n] = calcTicketValue(n, step, square, state, &stateValues)
						}
					}
					newValue, action := maxIndex(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])
					stateValues[step][square][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = newValue

					policy[step][square][state[0]][state[1]][state[2]][state[3]][state[4]][state[5]] = byte(action)
				}
			}
		}
	}

	/*miniWriter()
	fmt.Println("done")

	for i := 0; i < numSquares; i++ {
		fmt.Printf("policy[0][%d][1][1][1][1][1][1] = %f\n", i, float64(policy[0][i][1][1][1][1][1][1]))
		fmt.Printf("stateValues[0][%d][1][1][1][1][1][1] = %f\n", i, stateValues[0][i][1][1][1][1][1][1])
	}
	*/
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
