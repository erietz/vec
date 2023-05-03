package vec

import "testing"

func TestFilterInts(t *testing.T) {
	testCases := []struct {
		name   string
		input  Vec[int]
		output Vec[int]
		fn     func(i int) bool
	}{
		{
			name:   "is even",
			input:  Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			output: Vec[int]{0, 2, 4, 6, 8, 10, 12, 14},
			fn: func(i int) bool {
				return i%2 == 0
			},
		},
		{
			name:   "is odd",
			input:  Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			output: Vec[int]{1, 3, 5, 7, 9, 11, 13, 15},
			fn: func(i int) bool {
				return i%2 != 0
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[int]{},
			output: Vec[int]{},
			fn: func(i int) bool {
				return i%2 != 0
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.input.Filter(tc.fn)

			if len(out) != len(tc.output) {
				t.Errorf("got %v, want %v", out, tc.output)
			}

			for i := range out {
				if out[i] != tc.output[i] {
					t.Errorf("got %v, want %v", out[i], tc.output[i])
				}
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	testCases := []struct {
		name   string
		input  Vec[string]
		output Vec[string]
		fn     func(i string) bool
	}{
		{
			name:   "is even",
			input:  Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7"},
			output: Vec[string]{"0", "2", "4", "6"},
			fn: func(i string) bool {
				return i == "0" || i == "2" || i == "4" || i == "6"
			},
		},
		{
			name:   "is odd",
			input:  Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7"},
			output: Vec[string]{"1", "3", "5", "7"},
			fn: func(i string) bool {
				return i == "1" || i == "3" || i == "5" || i == "7"
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[string]{},
			output: Vec[string]{},
			fn: func(i string) bool {
				return i == "potato"
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.input.Filter(tc.fn)

			if len(out) != len(tc.output) {
				t.Errorf("got %v, want %v", out, tc.output)
			}

			for i := range out {
				if out[i] != tc.output[i] {
					t.Errorf("got %v, want %v", out[i], tc.output[i])
				}
			}
		})
	}
}

func TestMapInts(t *testing.T) {
	testCases := []struct {
		name   string
		input  Vec[int]
		output Vec[int]
		fn     func(i int) int
	}{
		{
			name:   "square",
			input:  Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			output: Vec[int]{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100},
			fn: func(i int) int {
				return i * i
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[int]{},
			output: Vec[int]{},
			fn: func(i int) int {
				return i * i
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.input.Map(tc.fn)

			if len(out) != len(tc.output) {
				t.Errorf("got %v, want %v", out, tc.output)
			}

			for i := range out {
				if out[i] != tc.output[i] {
					t.Errorf("got %v, want %v", out[i], tc.output[i])
				}
			}
		})
	}
}

func TestMapString(t *testing.T) {
	testCases := []struct {
		name   string
		input  Vec[string]
		output Vec[string]
		fn     func(i string) string
	}{
		{
			name:   "Make cooler",
			input:  Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			output: Vec[string]{"0ice", "1ice", "2ice", "3ice", "4ice", "5ice", "6ice", "7ice", "8ice", "9ice", "10ice"},
			fn: func(s string) string {
				return s + "ice"
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[string]{},
			output: Vec[string]{},
			fn: func(s string) string {
				return s + "ice"
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.input.Map(tc.fn)

			if len(out) != len(tc.output) {
				t.Errorf("got %v, want %v", out, tc.output)
			}

			for i := range out {
				if out[i] != tc.output[i] {
					t.Errorf("got %v, want %v", out[i], tc.output[i])
				}
			}
		})
	}
}
