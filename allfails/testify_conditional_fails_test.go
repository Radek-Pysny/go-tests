package allfails

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_testify_assert_False(t *testing.T) {
	assert.False(t, true) // false != true

	t.Log("it will print")
}

func Test_testify_assert_True(t *testing.T) {
	assert.True(t, false) // true != false

	t.Log("it will print")
}

func Test_testify_assert_Zero(t *testing.T) {
	assert.Zero(t, 1) // 0 != 1

	t.Log("it will print")
}

func Test_testify_assert_Equal(t *testing.T) {
	assert.Equal(t, false, true)        // false != true
	assert.Equal(t, 1, "1")             // different types
	assert.Equal(t, int32(1), int64(1)) // yet another different types

	t.Log("it will print")
}

func Test_testify_assert_Same(t *testing.T) {
	x, y := 42, 42
	ptrX, ptrY := &x, &y

	require.Equal(t, x, y)     // no fail yet
	assert.Same(t, ptrX, ptrY) // pointers points to a different addresses

	t.Log("it will print")
}

func Test_testify_assert_InDelta(t *testing.T) {
	x := 1.0
	y := 0.0
	for range 10 {
		y += 0.1
	}

	assert.InDelta(t, x, y, 1e-20) // too small delta requested
}
