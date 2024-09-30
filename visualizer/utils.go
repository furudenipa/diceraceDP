package visualizer

import "github.com/furudenipa/diceraceDP/config"

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func computeStrides() []int {
	strides := make([]int, config.NumDimensions)
	strides[config.NumDimensions-1] = 1
	for d := config.NumDimensions - 2; d >= 1; d-- {
		strides[d] = strides[d+1] * config.MaxTickets
	}
	strides[0] = config.NumSquares * strides[1]
	return strides
}

// getFlatIndexは多次元インデックスをフラットインデックスに変換します。
func getFlatIndex(step, square, t1, t2, t3, t4, t5, t6 int, strides []int) int {
	return step*strides[0] + square*strides[1] +
		t1*strides[2] + t2*strides[3] + t3*strides[4] +
		t4*strides[5] + t5*strides[6] + t6*strides[7]
}
