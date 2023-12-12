package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLuckyTickit(t *testing.T) {
	tests, err := LoadTestCases("./test_case")
	if err != nil {
		panic(err)
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			result := luckyTicketCount(tc.In)
			require.Equal(t, tc.Out, result)
		})
	}
}

func BenchmarkLackyTicketCount(b *testing.B) {
	for i := 3; i <= 10; i++ {
		b.Run(fmt.Sprintf("halfNumberCount %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				luckyTicketCount(i)
			}
		})
	}
}
