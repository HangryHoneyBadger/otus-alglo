package main

import (
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
			result := lackyTicketCount(tc.In)
			require.Equal(t, tc.Out, result)
		})
	}
}
