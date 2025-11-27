package unit

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_std_assertions presents how assertions are supposed to be done within Go standard testing package.
// One has to write all test conditions using basic branching control structures (aka if or switch) and
// mark the test as failed explicitly.
func Test_std_assertions(t *testing.T) {
	if 1 == 2 {
		t.Error("1 == 2 ?!")
	}
}

func Test_testify_assertions(t *testing.T) {
	t.Run("generic equality assertion (beware of different data types; even int × int32 etc.)", func(t *testing.T) {
		assert.Equal(t, 1, 1)
		assert.NotEqual(t, 1, int32(1))

		type indication bool
		assert.NotEqual(t, indication(false), false) // defined type differs from its "underlying type"
	})

	t.Run("bool equality assertion", func(t *testing.T) {
		assert.True(t, true)
		assert.False(t, false)
	})

	t.Run("nil equality assertions", func(t *testing.T) {
		assert.Nil(t, nil)
		assert.NotNil(t, 1)
	})

	t.Run("order assertions", func(t *testing.T) {
		assert.Greater(t, 42.1, 42.099)
		assert.GreaterOrEqual(t, 42.1, 42.099)
		assert.GreaterOrEqual(t, 42.1, 42.1)
		assert.LessOrEqual(t, -1, -1)
		assert.LessOrEqual(t, -1, 32987)
		assert.Less(t, -1, 23)

		assert.IsIncreasing(t, []int{})   // empty is both increasing and decreasing
		assert.IsIncreasing(t, []int{42}) // single-element is both increasing and decreasing
		assert.IsIncreasing(t, []int{1, 2, 3})
		assert.IsIncreasing(t, []string{"a", "alma", "alma mater", "axis", "base"})

		assert.IsDecreasing(t, []int{})
		assert.IsDecreasing(t, []int{127})
		assert.IsDecreasing(t, []int{1, 0, -1})
		assert.IsDecreasing(t, []string{"zzz", "z", "xyz", "xab"})

		assert.IsNonIncreasing(t, []int{3, 3, 3, 2, 2, 2, 1, 1, 1})
		assert.IsNonDecreasing(t, []int{1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4})
	})

	t.Run("zero value equality assertions", func(t *testing.T) {
		assert.Zero(t, false)
		assert.NotZero(t, true)

		assert.Zero(t, 0)
		assert.NotZero(t, 1)

		assert.Zero(t, 0.0)
		assert.NotZero(t, 1.85)

		assert.Zero(t, "")
		assert.NotZero(t, "non-empty string")

		var value = 42
		assert.Zero(t, nil) // zero value of the given value
		assert.NotZero(t, &value)
		assert.NotZero(t, []int{})
		assert.NotZero(t, map[string]int{})
		assert.NotZero(t, make(chan int))
		assert.NotZero(t, func() {})
		assert.NotZero(t, any(1))
	})

	t.Run("emptiness assertions", func(t *testing.T) {
		var (
			emptyArray             = [10]int{}
			emptySlice             = []int{}
			emptyMap               = map[string]int{}
			emptyUnbufferedChannel = make(chan int)
			emptyBufferedChannel   = make(chan int, 1)

			number                  = 42
			nonEmptyArray           = [10]int{1, 2, 3, 9: 42}
			nonEmptySlice           = []int{1, 2, 4}
			nonEmptyMap             = map[string]int{"one": 1, "two": 2}
			nonEmptyBufferedChannel = make(chan int, 1)
		)

		nonEmptyBufferedChannel <- 0

		// all zero values are considered empty
		assert.Empty(t, 0)
		assert.Empty(t, 0.0)
		assert.Empty(t, false)
		assert.Empty(t, "")
		assert.Empty(t, nil)
		// empty "containers"
		assert.Empty(t, emptyArray)
		assert.Empty(t, emptySlice)
		assert.Empty(t, emptyMap)
		assert.Empty(t, emptyUnbufferedChannel)
		assert.Empty(t, emptyBufferedChannel)
		// every pointer to empty value is also considered empty
		assert.Empty(t, &emptyArray)
		assert.Empty(t, &emptySlice)
		assert.Empty(t, &emptyMap)
		assert.Empty(t, &emptyUnbufferedChannel)
		assert.Empty(t, &emptyBufferedChannel)

		assert.NotEmpty(t, 1)
		assert.NotEmpty(t, 6.9)
		assert.NotEmpty(t, true)
		assert.NotEmpty(t, "text")
		assert.NotEmpty(t, &number)
		assert.NotEmpty(t, nonEmptyArray)
		assert.NotEmpty(t, nonEmptySlice)
		assert.NotEmpty(t, nonEmptyMap)
		assert.NotEmpty(t, nonEmptyBufferedChannel)
		assert.NotEmpty(t, &nonEmptyArray)
		assert.NotEmpty(t, &nonEmptySlice)
		assert.NotEmpty(t, &nonEmptyMap)
		assert.NotEmpty(t, &nonEmptyBufferedChannel)
	})

	t.Run("error assertions", func(t *testing.T) {
		assert.NoError(t, nil)

		err := fmt.Errorf("wrapped %w", context.Canceled)
		assert.ErrorIs(t, err, context.Canceled) // similar to errors.Is

		_, err = os.Open("non-existing-file!@%")
		var pathError *os.PathError
		assert.ErrorAs(t, err, &pathError) // similar to errors.As

		err = fmt.Errorf("extra context for error: %w", err)
		assert.ErrorContains(t, err, "extra context for error") // checking substring of Error() output

		err = fmt.Errorf("wrapped: %w", context.Canceled)
		assert.EqualError(t, err, "wrapped: context canceled")
	})

	t.Run("floating-point inexact equality assertions", func(t *testing.T) {
		assert.InDelta(t, 1.0, 0.999999, 0.00001)
		assert.InDeltaSlice(t, []float64{0.25, 0.50, 0.75}, []float64{0.24999, 0.49999, 0.74999}, 0.00002)
		assert.InDeltaMapValues(t, map[bool]float64{true: 42, false: 69}, map[bool]float64{true: 42, false: 69}, 0.0)
	})

	t.Run("containment assertions", func(t *testing.T) {
		assert.Contains(t, "Hello, world!", "Hello")
		assert.NotContains(t, "Hello, world!", "hi")

		assert.Contains(t, []int{1, 2, 4, 42}, 42)
		assert.NotContains(t, []int{1, 2, 4}, 42)
		assert.NotContains(t, []float64{0.0, 1.0}, 0.9999)

		assert.Contains(t, map[string]struct{}{"abc": {}}, "abc")
		assert.NotContains(t, map[string]struct{}{"abc": {}}, "xyz")
	})

	t.Run("length assertion", func(t *testing.T) {
		assert.Len(t, []int{}, 0)
		assert.Len(t, []int(nil), 0)
		assert.Len(t, []int{1, 2, 3}, 3)

		assert.Len(t, "", 0)
		assert.Len(t, "xyz", 3)

		assert.Len(t, map[bool]int{}, 0)
		assert.Len(t, map[bool]int{true: 1, false: 0}, 2)
	})

	t.Run("unordered sequence assertion", func(t *testing.T) {
		assert.ElementsMatch(t, [...]int{1, 2, 2, 3, 3, 3}, [...]int{3, 2, 1, 2, 3, 3})
		assert.ElementsMatch(t, [...]int{1, 2, 2, 3, 3, 3}, []int{3, 2, 1, 2, 3, 3})
		assert.ElementsMatch(t, []int{1, 2, 2, 3, 3, 3}, [...]int{3, 2, 1, 2, 3, 3})
		assert.ElementsMatch(t, []int{1, 2, 2, 3, 3, 3}, []int{3, 2, 1, 2, 3, 3})

		assert.NotElementsMatch(t, []int{1, 2, 3}, []int{3, 1, 2, 0})
		assert.NotElementsMatch(t, []int{1, 2, 3, -1}, []int{3, 1, 2, 0})
		assert.NotElementsMatch(t, []int{1, 2, 3, -1}, []int{3, 1, 2})
		assert.NotElementsMatch(t, []int{1, 2, 2, 3, 3, 3}, []int{3, 1, 2})
	})

	t.Run("pointer address assertions", func(t *testing.T) {
		var (
			x = 42
			y = 42
		)

		assert.Same(t, (*int)(nil), (*int)(nil))
		assert.Same(t, &x, &x)

		assert.NotSame(t, &x, (*int)(nil))
		assert.NotSame(t, &x, &y) // pointing to the same value, but each one on different address of the memory
	})

	t.Run("subset assertions", func(t *testing.T) {
		assert.Subset(t, []int{1, 2, 3}, []int{1, 2, 3})
		assert.Subset(t, []int{1, 2, 3}, []int{1, 2})
		assert.Subset(t, []int{1, 2, 3}, []int{})
		assert.Subset(t, map[string]int{"one": 1, "two": 2}, map[string]int{"one": 1, "two": 2})
		assert.Subset(t, map[string]int{"one": 1, "two": 2}, map[string]int{"one": 1})
		assert.Subset(t, map[string]int{"one": 1, "two": 2}, map[string]int{"two": 2})
		assert.Subset(t, map[string]int{"one": 1, "two": 2}, map[string]int{})

		assert.NotSubset(t, map[string]int{}, map[string]int{"one": 1, "two": 2})
	})

	t.Run("positive/negative assertions", func(t *testing.T) {
		assert.Positive(t, 1)
		assert.Negative(t, -1)
	})

	t.Run("RegEx assertions", func(t *testing.T) {
		assert.Regexp(t, regexp.MustCompile("^( [[:punct:]][A-Z]l{2,2}o)+.$"), " 'Allo 'Allo!")
	})
}

type DAO interface {
	Fetch(id string) (any, error)
	Store(id string, value any) error
}

type DaoMock struct {
	DAO
	ReturnValue any
	ReturnError error
}

func (d *DaoMock) Fetch(_ string) (any, error) {
	return d.ReturnValue, d.ReturnError
}

func NewFetchDaoMock(x any, err error) DaoMock {
	return DaoMock{nil, x, err}
}
