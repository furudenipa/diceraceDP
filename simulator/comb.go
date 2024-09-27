package main

func generateCombination(n, currentSum, targetSum int, currentTuple []int, results *[][]int) {
	if n == 6 {
		newTuple := make([]int, len(currentTuple))
		copy(newTuple, currentTuple)
		*results = append(*results, newTuple)
		return
	}

	start := max(0, targetSum-currentSum-(5-n)*(maxTickets-1))
	end := min(maxTickets-1, targetSum-currentSum) + 1
	for i := start; i < end; i++ {
		currentTuple = append(currentTuple, i)
		generateCombination(n+1, currentSum+i, targetSum, currentTuple, results)
		currentTuple = currentTuple[:len(currentTuple)-1]
	}
}
