package mdpsolver

import "github.com/furudenipa/diceraceDP/config"

// generateCombination は、指定された条件に基づいて数値の組み合わせを生成します。
// n: 現在の深さ（現在のタプルの長さ）
// currentSum: 現在のタプルの合計
// targetSum: 目標とする合計値
// currentTuple: 現在のタプルのスライス
// results: 結果を格納する二次元スライスへのポインタ
func generateCombination(n, currentSum, targetSum int, currentTuple []int, results *[][]int) {
	if n == 6 {
		newTuple := make([]int, len(currentTuple))
		copy(newTuple, currentTuple)
		*results = append(*results, newTuple)
		return
	}

	start := max(0, targetSum-currentSum-(5-n)*(config.MaxTickets-1))
	end := min(config.MaxTickets-1, targetSum-currentSum) + 1
	for i := start; i < end; i++ {
		currentTuple = append(currentTuple, i)
		generateCombination(n+1, currentSum+i, targetSum, currentTuple, results)
		currentTuple = currentTuple[:len(currentTuple)-1]
	}
}
