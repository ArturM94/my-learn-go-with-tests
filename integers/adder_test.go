package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("both positive", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertSum(t, sum, expected)
	})
	t.Run("one positive, one negative", func(t *testing.T) {
		sum := Add(2, -3)
		expected := -1
		assertSum(t, sum, expected)
	})
	t.Run("both negative", func(t *testing.T) {
		sum := Add(-1, -5)
		expected := -6
		assertSum(t, sum, expected)
	})
}

func assertSum(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
