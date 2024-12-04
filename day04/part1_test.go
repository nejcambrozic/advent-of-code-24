package day04

import (
	"testing"
)

// run with
// `go test day04/part1_test.go day04/part1.go`
func TestIsMatch(t *testing.T) {
	testParams := []struct {
		in  string
		out bool
	}{
		// valid
		{"X", true},
		{"XM", true},
		{"XMA", true},
		{"XMAS", true},
		{"S", true},
		{"SA", true},
		{"SAM", true},
		{"SAMX", true},
		// negative cases
		{"A", false},
	}

	for _, tt := range testParams {
		t.Run(tt.in, func(t *testing.T) {
			result := IsMatch(tt.in)
			if result != tt.out {
				t.Errorf("Error `%s`: expected `%t` got %t", tt.in, tt.out, result)
			}
		})
	}

}
