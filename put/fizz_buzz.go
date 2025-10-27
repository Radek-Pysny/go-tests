package put

import (
	"strconv"
)

func FizzBuzz(x int) (string, error) {
	switch {
	case x < 0:
		return "", ErrNegativeArgument

	case x == 0:
		return "", ErrZeroArgument
	}

	addFizz := x%3 == 0
	addBuzz := x%5 == 0

	switch {
	case addFizz && addBuzz:
		return "Fizz Buzz", nil

	case addFizz:
		return "Fizz", nil

	case addBuzz:
		return "Buzz", nil

	default:
		return strconv.Itoa(x), nil
	}
}
