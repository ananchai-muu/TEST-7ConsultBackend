package left_right

import "testing"

func TestFindBestNumberFromEncoded(t *testing.T) {
	testCases := []struct {
		name           string
		input         string
		wantDecoded   string
		wantSum       int
	}{
		{
			name:         "Test case 1: LLRR=",
			input:       "LLRR=",
			wantDecoded: "210122",
			wantSum:     8,
		},
		{
			name:         "Test case 2: ==RLL",
			input:       "==RLL",
			wantDecoded: "000210",
			wantSum:     3,
		},
		{
			name:         "Test case 3: =LLRR",
			input:       "=LLRR",
			wantDecoded: "221012",
			wantSum:     8,
		},
		{
			name:         "Test case 4: RRL=R",
			input:       "RRL=R",
			wantDecoded: "012001",
			wantSum:     4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotDecoded, gotSum := FindBestNumberFromEncoded(tc.input)
			if gotDecoded != tc.wantDecoded {
				t.Errorf("FindBestNumberFromEncoded(%q) decoded = %v, want %v", 
					tc.input, gotDecoded, tc.wantDecoded)
			}
			if gotSum != tc.wantSum {
				t.Errorf("FindBestNumberFromEncoded(%q) sum = %v, want %v", 
					tc.input, gotSum, tc.wantSum)
			}
		})
	}
}

func TestDecodeSequence(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		want     string
	}{
		{
			name:  "Test LLRR=",
			input: "LLRR=",
			want:  "210122",
		},
		{
			name:  "Test ==RLL",
			input: "==RLL",
			want:  "000210",
		},
		{
			name:  "Test =LLRR",
			input: "=LLRR",
			want:  "221012",
		},
		{
			name:  "Test RRL=R",
			input: "RRL=R",
			want:  "012001",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := decodeSequence(tc.input)
			if got != tc.want {
				t.Errorf("decodeSequence(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestSumOfDigits(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		want     int
	}{
		{
			name:  "Test 210122",
			input: "210122",
			want:  8,
		},
		{
			name:  "Test 000210",
			input: "000210",
			want:  3,
		},
		{
			name:  "Test 221012",
			input: "221012",
			want:  8,
		},
		{
			name:  "Test 012001",
			input: "012001",
			want:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumOfDigits(tc.input)
			if got != tc.want {
				t.Errorf("sumOfDigits(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
} 