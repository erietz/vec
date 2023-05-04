package vec

import (
	"strconv"
	"testing"
)

func TestFilterInts(t *testing.T) {
	testCases := []struct {
		name   string
		input  Vec[int]
		output Vec[int]
		fn     func(i int, index int) bool
	}{
		{
			name:   "is even",
			input:  Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			output: Vec[int]{0, 2, 4, 6, 8, 10, 12, 14},
			fn: func(i int, _ int) bool {
				return i%2 == 0
			},
		},
		{
			name:   "is odd",
			input:  Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			output: Vec[int]{1, 3, 5, 7, 9, 11, 13, 15},
			fn: func(i int, _ int) bool {
				return i%2 != 0
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[int]{},
			output: Vec[int]{},
			fn: func(i int, _ int) bool {
				return i%2 != 0
			},
		},
		{
			name:   "index sanity check",
			input:  Vec[int]{321, 1, 2, 3, 4, 5, 6, 123},
			output: Vec[int]{321, 123},
			fn: func(i int, index int) bool {
				return index == 0 || index == 7
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
		fn     func(i string, index int) bool
	}{
		{
			name:   "is even",
			input:  Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7"},
			output: Vec[string]{"0", "2", "4", "6"},
			fn: func(i string, _ int) bool {
				return i == "0" || i == "2" || i == "4" || i == "6"
			},
		},
		{
			name:   "is odd",
			input:  Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7"},
			output: Vec[string]{"1", "3", "5", "7"},
			fn: func(i string, _ int) bool {
				return i == "1" || i == "3" || i == "5" || i == "7"
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[string]{},
			output: Vec[string]{},
			fn: func(i string, _ int) bool {
				return i == "potato"
			},
		},
		{
			name:   "index sanity check",
			input:  Vec[string]{"first", "1", "2", "3", "4", "5", "6", "last"},
			output: Vec[string]{"first", "last"},
			fn: func(i string, index int) bool {
				return index == 0 || index == 7
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
		fn     func(i int, index int) int
	}{
		{
			name:   "square",
			input:  Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			output: Vec[int]{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100},
			fn: func(i int, _ int) int {
				return i * i
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[int]{},
			output: Vec[int]{},
			fn: func(i int, _ int) int {
				return i * i
			},
		},
		{
			name:   "index sanity check",
			input:  Vec[int]{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			output: Vec[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			fn: func(i int, index int) int {
				return index
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
		fn     func(i string, index int) string
	}{
		{
			name:   "Make cooler",
			input:  Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			output: Vec[string]{"0ice", "1ice", "2ice", "3ice", "4ice", "5ice", "6ice", "7ice", "8ice", "9ice", "10ice"},
			fn: func(s string, _ int) string {
				return s + "ice"
			},
		},
		{
			name:   "empty input, empty output",
			input:  Vec[string]{},
			output: Vec[string]{},
			fn: func(s string, _ int) string {
				return s + "ice"
			},
		},
		{
			name:   "index sanity check",
			input:  Vec[string]{"0ice", "1ice", "2ice", "3ice", "4ice", "5ice", "6ice", "7ice", "8ice", "9ice", "10ice"},
			output: Vec[string]{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			fn: func(s string, index int) string {
				return strconv.Itoa(index)
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
