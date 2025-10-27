package unit

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"testing"

	"github.com/Radek-Pysny/go-tests/put"
	"github.com/Radek-Pysny/go-tests/testutils"
	"github.com/stretchr/testify/require"
)

func Test_custom_EverySecond(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{nil, nil},
		{[]int{0}, nil},
		{[]int{1}, nil},
		{[]int{0, 1, 2, 3}, []int{1, 3}},
		{[]int{0, 1, 2, 3, 4}, []int{1, 3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v -> %v", tc.input, tc.expected), func(t *testing.T) {
			output := slices.Collect(put.EverySecond(tc.input))

			require.Len(t, output, len(tc.expected))
			require.Equal(t, tc.expected, output)
		})
	}
}

func Test_custom_Enumerate(t *testing.T) {
	testCases := []struct {
		title    string
		input    iter.Seq[int]
		expected []int
	}{
		{
			title:    "slice-values",
			input:    slices.Values([]int{0, 1, 2, 3, 4}),
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			title:    "every-second-iterator",
			input:    put.EverySecond([]int{0, 1, 2, 3, 4}),
			expected: []int{1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			outputIt := put.Enumerate(tc.input)

			testutils.AssertSeq2(
				t,
				slices.All(tc.expected),
				outputIt,
			)
		})
	}
}

func Test_custom_EnumerateFrom(t *testing.T) {
	testCases := []struct {
		title    string
		from     int
		input    iter.Seq[int]
		expected []int
	}{
		{
			title:    "slice-values-from-0",
			input:    slices.Values([]int{0, 1, 2, 3, 4}),
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			title:    "slice-values-from-1",
			input:    slices.Values([]int{0, 1, 2, 3, 4}),
			from:     1,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			title:    "slice-values-from-42",
			input:    slices.Values([]int{0, 1, 2, 3, 4}),
			from:     42,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			title:    "slice-values-from-minus-1",
			input:    slices.Values([]int{0, 1, 2, 3, 4}),
			from:     -1,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			title:    "slice-values-from-min-int",
			input:    slices.Values([]int{0, 1, 2, 3, 4}),
			from:     math.MinInt,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			title:    "every-second-iterator-from-0",
			input:    put.EverySecond([]int{0, 1, 2, 3, 4}),
			expected: []int{1, 3},
		},
		{
			title:    "every-second-iterator-from-1",
			input:    put.EverySecond([]int{0, 1, 2, 3, 4}),
			from:     1,
			expected: []int{1, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {

			i := 0
			for returnedIndex, returnedValue := range put.EnumerateFrom(tc.input, tc.from) {
				expectedValue := tc.expected[i]
				expectedIndex := tc.from + i

				require.Equal(t, expectedIndex, returnedIndex, "index[%d]", i)
				require.Equal(t, expectedValue, returnedValue, "value[%d]", i)

				i++
			}
		})
	}
}

func Test_custom_Pairs(t *testing.T) {
	type pair struct {
		fst int
		snd int
	}

	testCases := []struct {
		title    string
		input    []int
		expected []pair
	}{
		{
			title:    "empty",
			input:    nil,
			expected: []pair{}, // cannot be nil due to unit test implementation
		},
		{
			title:    "[1]",
			input:    []int{1},
			expected: []pair{{1, 0}},
		},
		{
			title:    "[1,2]",
			input:    []int{1, 2},
			expected: []pair{{1, 2}},
		},
		{
			title:    "[1,2,3]",
			input:    []int{1, 2, 3},
			expected: []pair{{1, 2}, {3, 0}},
		},
		{
			title:    "[1,2,3,4]",
			input:    []int{1, 2, 3, 4},
			expected: []pair{{1, 2}, {3, 4}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			returnedIt := put.Pairs(slices.Values(tc.input))

			returned := make([]pair, 0, len(tc.expected))
			for fst, snd := range returnedIt {
				returned = append(returned, pair{fst: fst, snd: snd})
			}

			require.Len(t, returned, len(tc.expected))
			require.Equal(t, tc.expected, returned)
		})
	}
}
