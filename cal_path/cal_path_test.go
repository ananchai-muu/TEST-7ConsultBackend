package cal_path

import (
	"testing"
)

func TestCalculateMaxPath(t *testing.T) {
	tests := []struct {
		name     string
		data     Data
		expected int
	}{
		{
			name: "simple triangle",
			data: Data{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			expected: 237,
		},
		{
			name:     "empty triangle",
			data:     Data{},
			expected: 0,
		},
		{
			name: "triangle with one row",
			data: Data{
				{10},
			},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalMaxPath(tt.data)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}
