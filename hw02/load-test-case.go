package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
