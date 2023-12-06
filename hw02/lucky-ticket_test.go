package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

type TestCase struct {
	In  int
	Out int
}

func LoadTestCases(dir string) (map[string]*TestCase, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	testCases := make(map[string]*TestCase, len(files)/2)
	for _, file := range files {
		name := file.Name()
		base := strings.TrimSuffix(name, filepath.Ext(name))
		content, err := os.ReadFile(filepath.Join(dir, name))
		contentStr := strings.TrimRight(string(content), "\r\n")
		val, err := strconv.Atoi(contentStr)
		if err != nil {
			return nil, err
		}

		if _, ok := testCases[base]; !ok {
			testCases[base] = new(TestCase)
		}

		if strings.HasSuffix(name, ".in") {
			testCases[base].In = val
		} else if strings.HasSuffix(name, ".out") {
			testCases[base].Out = val
		}
	}

	return testCases, nil
}

func TestLuckyTickit(t *testing.T) {
	tests, err := LoadTestCases("./test_case")
	if err != nil {
		panic(err)
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			result := lackyTicket(tc.In)
			require.Equal(t, tc.Out, result)
		})
	}
}
