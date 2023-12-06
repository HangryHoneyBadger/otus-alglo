package main

func lackyTicket(halfNumberCount int) int {
	wayvs := make(map[int]int, 9*halfNumberCount)
	lackyStep(0, 0, halfNumberCount, wayvs)
	count := 0
	for _, v := range wayvs {
		count += v * v
	}
	return count
}

func lackyStep(step, sum, halfNumberCount int, value map[int]int) {
	for i := 0; i < 10; i++ {
		if step+1 == halfNumberCount {
			value[sum+i]++
		} else {
			lackyStep(step+1, sum+i, halfNumberCount, value)
		}
	}
}
