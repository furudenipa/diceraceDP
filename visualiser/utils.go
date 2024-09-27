package main

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

const (
	numDimensions    = 8
	numSquares       = 18
	maxTickets       = 10
	numSteps         = 100 //100
	ticketSquare     = 0
	advanceSixSquare = 12
	STONE            = 4.3  // yellow
	CREDIT           = 13.5 // 100K
	REPORT           = 4    // yellow
	MIYU             = 0
	ELEPH            = 50
)

func computeStrides() []int {
	strides := make([]int, numDimensions)
	strides[numDimensions-1] = 1
	for d := numDimensions - 2; d >= 1; d-- {
		strides[d] = strides[d+1] * maxTickets
	}
	strides[0] = numSquares * strides[1]
	return strides
}

// getFlatIndexは多次元インデックスをフラットインデックスに変換します。
func getFlatIndex(step, square, t1, t2, t3, t4, t5, t6 int, strides []int) int {
	return step*strides[0] + square*strides[1] +
		t1*strides[2] + t2*strides[3] + t3*strides[4] +
		t4*strides[5] + t5*strides[6] + t6*strides[7]
}

/*
// getUserIndicesはユーザーから6つのインデックスを取得します。
func getUserIndices(reader *bufio.Reader) ([]int, error) {
	var indices []int
	prompts := []string{"t1: ", "t2: ", "t3: ", "t4: ", "t5: ", "t6: "}
	for _, prompt := range prompts {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			return nil, fmt.Errorf("有効な整数を入力してください: %v", err)
		}
		indices = append(indices, num)
	}
	return indices, nil
}

// validateIndicesはインデックスが有効な範囲内にあるかをチェックします。
func validateIndices(t1, t2, t3, t4, t5, t6 int) bool {
	return allInRange(0, maxTickets-1, t1, t2, t3, t4, t5, t6)
}

// allInRangeは全ての値が指定された範囲内にあるかをチェックします。
func allInRange(min, max int, values ...int) bool {
	for _, v := range values {
		if v < min || v > max {
			return false
		}
	}
	return true
}*/
