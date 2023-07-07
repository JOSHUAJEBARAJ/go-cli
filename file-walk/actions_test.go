package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtenstion", "testdata/dir.log", "", 0, false},
		{"FilterExtenstionMatch", "testdata/dir.log", ".log", 0, false},
		{"FilterExtenstionNoMatch", "testdata/dir.log", ".sh", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatal(err)
			}
			f := filterOut(tc.file, tc.ext, tc.minSize, info)
			if f != tc.expected {
				t.Errorf("Expected %t and got %t", tc.expected, f)
			}

		})
	}
}
