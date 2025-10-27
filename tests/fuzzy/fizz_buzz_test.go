package fuzzy

import (
	"strconv"
	"testing"

	"github.com/Radek-Pysny/go-tests/put"
	"github.com/stretchr/testify/require"
)

func Fuzz_FizzBuzz(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, x int) {
		returned, err := put.FizzBuzz(x)

		switch {
		case x < 0:
			require.ErrorIs(t, err, put.ErrNegativeArgument)
			require.Equal(t, "", returned)

		case x == 0:
			require.ErrorIs(t, err, put.ErrZeroArgument)
			require.Equal(t, "", returned)

		case x%15 == 0:
			require.NoError(t, err)
			require.Equal(t, "Fizz Buzz", returned)

		case x%3 == 0:
			require.NoError(t, err)
			require.Equal(t, "Fizz", returned)

		case x%5 == 0:
			require.NoError(t, err)
			require.Equal(t, "Buzz", returned)

		default:
			require.NoError(t, err)

			y, err := strconv.Atoi(returned)
			require.NoError(t, err)
			require.Equal(t, x, y)
		}

	})
}
