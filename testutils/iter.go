package testutils

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/require"
)

func AssertSeq2[K comparable, V comparable](t *testing.T, expected iter.Seq2[K, V], returned iter.Seq2[K, V]) {
	t.Helper()

	var (
		expectedKeys   []K
		expectedValues []V
		returnedKeys   []K
		returnedValues []V
	)

	for k, v := range expected {
		expectedKeys = append(expectedKeys, k)
		expectedValues = append(expectedValues, v)
	}

	for k, v := range returned {
		returnedKeys = append(returnedKeys, k)
		returnedValues = append(returnedValues, v)
	}

	require.Equal(t, expectedKeys, returnedKeys)
	require.Equal(t, expectedValues, returnedValues)
}
