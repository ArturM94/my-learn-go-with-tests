package generics_test

import (
	generics "learn-go-with-tests/19-generics"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(generics.Stack[int])

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		stackOfInts.Push(456)
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(1)
		stackOfInts.Push(2)
		firstNum, _ := stackOfInts.Pop()
		secondNum, _ := stackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
