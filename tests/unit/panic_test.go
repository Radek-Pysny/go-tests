package unit

import (
	"testing"

	"github.com/Radek-Pysny/go-tests/testutils"
	"github.com/stretchr/testify/assert"
)

func functionThatPanics() {
	panic("kernel panic")
}

func Test_testutils_ExpectPanic_anyPanic(t *testing.T) {
	testutils.ExpectPanic(t, functionThatPanics)
}

func Test_assert_Panic_anyPanic(t *testing.T) {
	assert.Panics(t, functionThatPanics)
}
