package allfails

import (
	"testing"
)

// t.Fail, t.Error, and t.Errorf mark test as failed, but they continue the test

func Test_std_Fail(t *testing.T) {
	t.Fail()
	t.Log("it will print")
}

func Test_std_Error(t *testing.T) {
	t.Error("failing with custom error message")
	t.Log("it will print")
}

func Test_std_Errorf(t *testing.T) {
	t.Errorf("failing with custom error message and %s", "variable")
	t.Log("it will print")
}

// t.FailNow, t.Fatal, adn t.Fatalf will mark test as failed and stop suddenly

func Test_std_FailNow(t *testing.T) {
	t.FailNow()
	t.Log("it will NOT print")
}

func Test_std_Fatal(t *testing.T) {
	t.Fatal("failing with custom error message")
	t.Log("it will NOT print")
}

func Test_std_Fatalf(t *testing.T) {
	t.Fatalf("failing with custom error message and %s", "variable")
	t.Log("it will NOT print")
}

// Failed can be used to detect if the test is already failed or not

func Test_std_Failed(t *testing.T) {
	if t.Failed() {
		panic("it shall never happen!!!")
	}

	t.Fail()

	if !t.Failed() {
		panic("it has to be failed already!!!")
	}
}

// Test_Failed_subtest shows that subtests does not share fail mark, so t.Failed can see only
// the current subtest.
func Test_std_Failed_subtest(t *testing.T) {
	t.Run("phase 1", func(t *testing.T) {
		if t.Failed() {
			panic("it shall never happen!!!")
		}

		t.Fail()
	})

	t.Run("phase 2", func(t *testing.T) {
		if t.Failed() {
			panic("it shall never happen!!!")
		}

		t.Fail()
	})
}
