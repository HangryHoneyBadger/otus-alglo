package main

func luckyTicketCount(halfNumberCount int) int {
	maxSum := 9 * halfNumberCount
	dpCurrent := make([]int, maxSum+1)
	dpPrevious := make([]int, maxSum+1)
	dpPrevious[0] = 1

	for i := 1; i <= halfNumberCount; i++ {
		for j := 0; j <= maxSum; j++ {
			dpCurrent[j] = 0
			for k := 0; k <= 9; k++ {
				if j-k >= 0 {
					dpCurrent[j] += dpPrevious[j-k]
				}
			}
		}
		dpPrevious, dpCurrent = dpCurrent, dpPrevious
	}

	count := 0
	for _, v := range dpPrevious {
		count += v * v
	}
	return count
}
