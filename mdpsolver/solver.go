package mdpsolver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/furudenipa/diceraceDP/config"
)

func Run() {
	fmt.Println("Settings")
	fmt.Println(" numSquares: ", config.NumSquares)
	fmt.Println(" maxTickets: ", config.MaxTickets)
	fmt.Println(" numSteps:  ", config.NumSteps)
	fmt.Println("----------start----------")
	var currentStateValues [config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64
	var prevStateValues [config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64
	var policy [config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]byte

	file, err := os.Create("policy.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var buffer bytes.Buffer

	for step := 0; step < config.NumSteps; step++ {
		fmt.Println("step: ", step)
		for ticketSum := 0; ticketSum <= (config.MaxTickets-1)*6; ticketSum++ {
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
				for square := 1; square < config.NumSquares; square++ {
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

		// Write policy to file
		for i := 0; i < config.NumSquares; i++ {
			for j := 0; j < config.MaxTickets; j++ {
				for k := 0; k < config.MaxTickets; k++ {
					for l := 0; l < config.MaxTickets; l++ {
						for m := 0; m < config.MaxTickets; m++ {
							err = binary.Write(&buffer, binary.LittleEndian, policy[i][j][k][l][m])
							if err != nil {
								fmt.Println("Error writing policy:", err)
								return
							}
						}
					}
				}
			}
		}
		if _, err := file.Write(buffer.Bytes()); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		buffer.Reset()

		prevStateValues = currentStateValues
	}
	fmt.Println("----------end----------")
}

// maxIndex returns the maximum value and its index from the given values.
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
