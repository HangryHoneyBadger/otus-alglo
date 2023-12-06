package main

import (
	"fmt"
	"sync"
)

func main() {
	tests, err := LoadTestCases("./test_case")
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(tests))
	for name, tc := range tests {
		go func(name string, t *TestCase) {
			defer wg.Done()
			result := lackyTicketCount(t.In)
			fmt.Println(name, result == t.Out)
		}(name, tc)
	}
	wg.Wait()
}
