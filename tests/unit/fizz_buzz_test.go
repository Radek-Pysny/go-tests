package unit

import (
	"strconv"
	"testing"

	"github.com/Radek-Pysny/go-tests/put"
	"github.com/stretchr/testify/require"
)

func Test_FizzBuzz(t *testing.T) {
	testCases := []struct {
		input       int
		expected    string
		expectedErr error
	}{
		{-10, "", put.ErrNegativeArgument},
		{-2, "", put.ErrNegativeArgument},
		{-1, "", put.ErrNegativeArgument},
		{0, "", put.ErrZeroArgument},
		{1, "1", nil},
		{2, "2", nil},
		{3, "Fizz", nil},
		{4, "4", nil},
		{5, "Buzz", nil},
		{6, "Fizz", nil},
		{7, "7", nil},
		{8, "8", nil},
		{9, "Fizz", nil},
		{10, "Buzz", nil},
		{11, "11", nil},
		{12, "Fizz", nil},
		{13, "13", nil},
		{14, "14", nil},
		{15, "Fizz Buzz", nil},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			returned, err := put.FizzBuzz(tc.input)

			require.Equal(t, tc.expected, returned)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}
