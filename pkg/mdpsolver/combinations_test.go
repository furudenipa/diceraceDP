package mdpsolver

import (
	"reflect"
	"testing"

	"github.com/furudenipa/diceraceDP/config"
)

func TestGenerateCombination(t *testing.T) {
	config.SetTestConfig()
	tests := []struct {
		name         string
		n            int
		currentSum   int
		targetSum    int
		currentTuple []int
		expected     [][]int
	}{
		{
			name:         "初期状態からの組み合わせ生成",
			n:            0,
			currentSum:   0,
			targetSum:    3,
			currentTuple: []int{},
			expected: [][]int{
				{0, 0, 0, 0, 0, 3},
				{0, 0, 0, 0, 1, 2},
				{0, 0, 0, 0, 2, 1},
				{0, 0, 0, 0, 3, 0},
				{0, 0, 0, 1, 0, 2},
				{0, 0, 0, 1, 1, 1},
				{0, 0, 0, 1, 2, 0},
				{0, 0, 0, 2, 0, 1},
				{0, 0, 0, 2, 1, 0},
				{0, 0, 0, 3, 0, 0},
				{0, 0, 1, 0, 0, 2},
				{0, 0, 1, 0, 1, 1},
				{0, 0, 1, 0, 2, 0},
				{0, 0, 1, 1, 0, 1},
				{0, 0, 1, 1, 1, 0},
				{0, 0, 1, 2, 0, 0},
				{0, 0, 2, 0, 0, 1},
				{0, 0, 2, 0, 1, 0},
				{0, 0, 2, 1, 0, 0},
				{0, 0, 3, 0, 0, 0},
				{0, 1, 0, 0, 0, 2},
				{0, 1, 0, 0, 1, 1},
				{0, 1, 0, 0, 2, 0},
				{0, 1, 0, 1, 0, 1},
				{0, 1, 0, 1, 1, 0},
				{0, 1, 0, 2, 0, 0},
				{0, 1, 1, 0, 0, 1},
				{0, 1, 1, 0, 1, 0},
				{0, 1, 1, 1, 0, 0},
				{0, 1, 2, 0, 0, 0},
				{0, 2, 0, 0, 0, 1},
				{0, 2, 0, 0, 1, 0},
				{0, 2, 0, 1, 0, 0},
				{0, 2, 1, 0, 0, 0},
				{0, 3, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 2},
				{1, 0, 0, 0, 1, 1},
				{1, 0, 0, 0, 2, 0},
				{1, 0, 0, 1, 0, 1},
				{1, 0, 0, 1, 1, 0},
				{1, 0, 0, 2, 0, 0},
				{1, 0, 1, 0, 0, 1},
				{1, 0, 1, 0, 1, 0},
				{1, 0, 1, 1, 0, 0},
				{1, 0, 2, 0, 0, 0},
				{1, 1, 0, 0, 0, 1},
				{1, 1, 0, 0, 1, 0},
				{1, 1, 0, 1, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{1, 2, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 1},
				{2, 0, 0, 0, 1, 0},
				{2, 0, 0, 1, 0, 0},
				{2, 0, 1, 0, 0, 0},
				{2, 1, 0, 0, 0, 0},
				{3, 0, 0, 0, 0, 0},
			},
		},
		{
			name:         "部分的な組み合わせ生成",
			n:            2,
			currentSum:   1,
			targetSum:    4,
			currentTuple: []int{1, 0},
			expected: [][]int{
				{1, 0, 0, 0, 0, 3},
				{1, 0, 0, 0, 1, 2},
				{1, 0, 0, 0, 2, 1},
				{1, 0, 0, 0, 3, 0},
				{1, 0, 0, 1, 0, 2},
				{1, 0, 0, 1, 1, 1},
				{1, 0, 0, 1, 2, 0},
				{1, 0, 0, 2, 0, 1},
				{1, 0, 0, 2, 1, 0},
				{1, 0, 0, 3, 0, 0},
				{1, 0, 1, 0, 0, 2},
				{1, 0, 1, 0, 1, 1},
				{1, 0, 1, 0, 2, 0},
				{1, 0, 1, 1, 0, 1},
				{1, 0, 1, 1, 1, 0},
				{1, 0, 1, 2, 0, 0},
				{1, 0, 2, 0, 0, 1},
				{1, 0, 2, 0, 1, 0},
				{1, 0, 2, 1, 0, 0},
				{1, 0, 3, 0, 0, 0},
			},
		},
		{
			name:         "目標合計に達する組み合わせ",
			n:            5,
			currentSum:   0,
			targetSum:    2,
			currentTuple: []int{0, 0, 0, 0, 0},
			expected: [][]int{
				{0, 0, 0, 0, 0, 2},
			},
		},
		{
			name:         "目標合計を超える組み合わせは生成されない",
			n:            4,
			currentSum:   4,
			targetSum:    3,
			currentTuple: []int{1, 1, 1, 1},
			expected:     [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var results [][]int
			generateCombination(tt.n, tt.currentSum, tt.targetSum, tt.currentTuple, &results)

			// 結果の比較
			if !compareSlices(results, tt.expected) {
				t.Errorf("Test %s failed.\nExpected: %v\nGot: %v", tt.name, tt.expected, results)
			}
		})
	}
}

// compareSlices は、二次元スライスが同じ内容を持つかを比較します。
func compareSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}
