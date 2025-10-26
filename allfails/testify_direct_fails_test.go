package allfails

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// assert.Fail and assert.Failf mark test as failed, but they continue the tests

func Test_testify_assert_Fail(t *testing.T) {
	assert.Fail(t, "expected failure")

	t.Log("it will print")
}

func Test_testify_assert_Failf(t *testing.T) {
	assert.Failf(t, "expected failure", "custom message with %s", "variable")

	t.Log("it will print")
}

// assert.FailNow, t.FailNowf, and all direct failing function from package require
// (namely require.Fail, require.Failf, require.FailNow, require.FailNowf) will
// mark test as failed and stop suddenly

func Test_testify_assert_FailNow(t *testing.T) {
	assert.FailNow(t, "expected failure")

	t.Log("it will NOT print")
}

func Test_testify_assert_FailNowf(t *testing.T) {
	assert.FailNowf(t, "expected failure", "custom message with %s", "variable")

	t.Log("it will NOT print")
}
func Test_testify_require_Fail(t *testing.T) {
	require.Fail(t, "expected failure")

	t.Log("it will NOT print")
}

func Test_testify_require_Failf(t *testing.T) {
	require.Failf(t, "expected failure", "custom message with %s", "variable")

	t.Log("it will NOT print")
}

func Test_testify_require_FailNow(t *testing.T) {
	require.FailNow(t, "expected failure")

	t.Log("it will NOT print")
}

func Test_testify_require_FailNowf(t *testing.T) {
	require.FailNowf(t, "expected failure", "custom message with %s", "variable")

	t.Log("it will NOT print")
}
