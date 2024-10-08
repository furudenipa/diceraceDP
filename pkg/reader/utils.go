package reader

import (
	"fmt"

	"github.com/furudenipa/diceraceDP/config"
)

// Powはbaseのexp乗を計算します。
func Pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

// 1次元配列policyのインデックス計算に必要なstridesを計算します。
func ComputeStrides() []int {
	strides := make([]int, config.NumDimensions)
	strides[config.NumDimensions-1] = 1
	for d := config.NumDimensions - 2; d >= 1; d-- {
		strides[d] = strides[d+1] * config.MaxTickets
	}
	strides[0] = config.NumSquares * strides[1]
	return strides
}

// getFlatIndexは多次元インデックスをフラットインデックスに変換します。
func GetFlatIndex(remainingRolls, square int, remainingTickets []int, strides []int) (int, error) {
	// Error handling
	// 重いのでdebug時のみ
	/*
		if remainingRolls < 0 || remainingRolls >= config.MaxRolls {
			return 0, fmt.Errorf("invalid remainingRolls: %d", remainingRolls)
		} else if square < 0 || square >= config.NumSquares {
			return 0, fmt.Errorf("invalid square: %d", square)
		}
		if err := checkRemainingTickets(remainingTickets); err != nil {
			return 0, err
		}*/

	index := remainingRolls*strides[0] + square*strides[1] +
		remainingTickets[0]*strides[2] + remainingTickets[1]*strides[3] + remainingTickets[2]*strides[4] +
		remainingTickets[3]*strides[5] + remainingTickets[4]*strides[6] + remainingTickets[5]*strides[7]

	return index, nil
}

func checkRemainingTickets(remainingTickets []int) error {
	for _, ticket := range remainingTickets {
		if ticket < 0 || ticket >= config.MaxTickets {
			return fmt.Errorf("invalid remainingTickets: %v", remainingTickets)
		}
	}
	return nil
}
