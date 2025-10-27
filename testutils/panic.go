package testutils

import (
	"testing"
)

func ExpectPanic(t *testing.T, f func()) {
	t.Helper()

	defer func() {
		_ = recover()
	}()

	// the given function should panic...
	f()

	t.Errorf("Not caught an expected panic!")
}
