package reader

import "github.com/furudenipa/diceraceDP/config"

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
func GetFlatIndex(step, square int, remainingTickets []int, strides []int) int {
	return step*strides[0] + square*strides[1] +
		remainingTickets[0]*strides[2] + remainingTickets[1]*strides[3] + remainingTickets[2]*strides[4] +
		remainingTickets[3]*strides[5] + remainingTickets[4]*strides[6] + remainingTickets[5]*strides[7]
}
