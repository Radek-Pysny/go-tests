package unit

import (
	"testing"
)

func TestVerbose(t *testing.T) {
	t.Run("flat-part", func(t *testing.T) {
		t.Log("starting the complex test")

		// ...
	})

	t.Run("loop-part", func(t *testing.T) {
		for i := range 10 {
			t.Logf("will run %dth step", i)

			// ...
		}
	})
}
