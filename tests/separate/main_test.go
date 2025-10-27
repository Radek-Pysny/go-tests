//go:build separate && !unit && !integration

package separate

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting test suite with all failing tests")

	m.Run()
}
