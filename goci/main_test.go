package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {

	var testCases = []struct {
		name   string
		proj   string
		out    string
		expErr error
	}{
		{
			name:   "Success",
			proj:   "./testdata/tool",
			out:    "Build was success\n\n",
			expErr: nil,
		},
		{
			name:   "Failure",
			proj:   "./testdata/toolErr",
			out:    "",
			expErr: &stepErr{step: "go build"},
		},
	}
	// iteration over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var out bytes.Buffer
			err := run(tc.proj, &out)
			// checking the error

			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error %q and got Nil instead", tc.expErr)
					return
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error %q and got %q", tc.expErr, err)
					return
				}
			}

			//
			if err != nil {
				t.Errorf("Unexptected error :%q", err)
			}
			if out.String() != tc.out {
				t.Errorf("Exptected output not matched %q %q", out.String(), tc.out)
			}
		})
	}
}
