package left_right

import (
	"backendtest/left_right"
	"testing"
)

// Test function to check the test cases
func TestDecodePattern(t *testing.T) {
	tests := []struct {
		pattern string
		expected string
	}{
		{"LLRR=", "210122"},
		{"==RLL", "000210"},
		{"=LLRR", "221012"},
		{"RRL=R", "012001"},
	}

	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			result := left_right.DecodePattern(test.pattern)
			if result != test.expected {
				t.Errorf("For pattern %s, expected %s, but got %s", test.pattern, test.expected, result)
			}
		})
	}
}