package main

func luckyTicketCount(halfNumberCount int) int {
	maxSum := 9 * halfNumberCount
	stepCurrent := make([]int, maxSum+1)
	stepPrevious := make([]int, maxSum+1)
	stepPrevious[0] = 1

	for i := 1; i <= halfNumberCount; i++ {
		for j := 0; j <= maxSum; j++ {
			stepCurrent[j] = 0
			for k := 0; k <= 9; k++ {
				if j-k >= 0 {
					stepCurrent[j] += stepPrevious[j-k]
				}
			}
		}
		stepPrevious, stepCurrent = stepCurrent, stepPrevious
	}

	count := 0
	for _, v := range stepPrevious {
		count += v * v
	}
	return count
}
