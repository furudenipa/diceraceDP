package mdpsolver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/furudenipa/diceraceDP/config"
)

func Run(filename string) {
	fmt.Println("Settings")
	fmt.Println(" numSquares: ", config.NumSquares)
	fmt.Println(" maxTickets: ", config.MaxTickets)
	fmt.Println(" numSteps:  ", config.MaxRolls)
	fmt.Println("----------start----------")
	var currentStateValues stateValues
	var prevStateValues stateValues
	var policy policy

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var buffer bytes.Buffer

	for remainingRolls := 0; remainingRolls < config.MaxRolls; remainingRolls++ {
		fmt.Println("step: ", remainingRolls)
		for ticketSum := 0; ticketSum <= (config.MaxTickets-1)*6; ticketSum++ {
			combinations := [][]int{}
			generateCombination(0, 0, ticketSum, []int{}, &combinations)

			for _, remainingTickets := range combinations {
				square := 0
				rollValue := calcRollValue(square, remainingRolls, remainingTickets, &prevStateValues)
				ticketValues := [6]float64{-1, -1, -1, -1, -1, -1}
				for n := 0; n < 6; n++ {
					if remainingTickets[n] > 0 {
						ticketValues[n] = calcTicketValue(n, square, remainingTickets, &currentStateValues)
					}
				}
				newValue, action := maxIndex(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])

				// If the value is -1, there are no tickets and rolls left.
				if newValue < 0 {
					newValue = 0
					action = 7
				}

				currentStateValues.SetValue(0, remainingTickets, newValue)
				policy.SetAction(0, remainingTickets, byte(action))
			}

			for _, remainingTickets := range combinations {
				for square := 1; square < config.NumSquares; square++ {
					rollValue := calcRollValue(square, remainingRolls, remainingTickets, &prevStateValues)
					ticketValues := [6]float64{-1, -1, -1, -1, -1, -1}
					for n := 0; n < 6; n++ {
						if remainingTickets[n] > 0 {
							ticketValues[n] = calcTicketValue(n, square, remainingTickets, &currentStateValues)
						}
					}
					newValue, action := maxIndex(rollValue, ticketValues[0], ticketValues[1], ticketValues[2], ticketValues[3], ticketValues[4], ticketValues[5])

					// If the value is -1, there are no tickets and rolls left.
					if newValue < 0 {
						newValue = 0
						action = 7
					}

					currentStateValues.SetValue(square, remainingTickets, newValue)
					policy.SetAction(square, remainingTickets, byte(action))
				}
			}
		}

		fmt.Println(" currentStateValues[0][1][0][0][0][0][0]:", currentStateValues.GetValue(0, []int{1, 0, 0, 0, 0, 0}))
		fmt.Println(" currentStateValues[0][0][1][0][0][0][0]:", currentStateValues.GetValue(0, []int{0, 1, 0, 0, 0, 0}))
		fmt.Println(" currentStateValues[0][0][0][1][0][0][0]:", currentStateValues.GetValue(0, []int{0, 0, 1, 0, 0, 0}))
		fmt.Println(" currentStateValues[0][0][0][0][1][0][0]:", currentStateValues.GetValue(0, []int{0, 0, 0, 1, 0, 0}))
		fmt.Println(" currentStateValues[0][0][0][0][0][1][0]:", currentStateValues.GetValue(0, []int{0, 0, 0, 0, 1, 0}))
		fmt.Println(" currentStateValues[0][0][0][0][0][0][1]:", currentStateValues.GetValue(0, []int{0, 0, 0, 0, 0, 1}))
		fmt.Println(" !currentStateValues[0][1][1][1][1][1][1]:", currentStateValues.GetValue(0, []int{1, 1, 1, 1, 1, 1}))
		fmt.Println(" !prevStateValues[0][1][1][1][1][1][1]:", prevStateValues.GetValue(0, []int{1, 1, 1, 1, 1, 1}))
		fmt.Println(" currentStateValues[0][1][1][1][0][0][0]:", currentStateValues.GetValue(0, []int{1, 1, 1, 0, 0, 0}))
		prevStateValues = currentStateValues // stateValuesのコピー, 配列型は値渡し

		// Write policy to file
		// TODO: rewrite this parts
		for i := 0; i < config.NumSquares; i++ {
			for j := 0; j < config.MaxTickets; j++ {
				for k := 0; k < config.MaxTickets; k++ {
					for l := 0; l < config.MaxTickets; l++ {
						for m := 0; m < config.MaxTickets; m++ {
							err = binary.Write(&buffer, binary.LittleEndian, policy.GetPolicyTwoDimensionalSlice(i, j, k, l, m))
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
