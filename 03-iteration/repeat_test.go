package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeat(t, repeated, expected)
	})
	t.Run("repeat 'x' 12 times", func(t *testing.T) {
		repeated := Repeat("x", 12)
		expected := "xxxxxxxxxxxx"
		assertRepeat(t, repeated, expected)
	})
}

func assertRepeat(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	b.Run("repeat 'x' 5 times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("x", 5)
		}
	})
	b.Run("repeat 'a' 10 times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 10)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: "aaa"
}
