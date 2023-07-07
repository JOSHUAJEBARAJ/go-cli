package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{name: "No filter", root: "testdata", cfg: config{}, expected: "testdata/dir.log\ntestdata/dir2/empty.sh\n"},
		{name: "Filter by extension", root: "testdata", cfg: config{ext: ".log"}, expected: "testdata/dir.log\n"},
		{name: "Filter by non-existent extension", root: "testdata", cfg: config{ext: "txt"}, expected: ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Exepected %q, got %q instead", tc.expected, res)
			}
		})

	}
}
