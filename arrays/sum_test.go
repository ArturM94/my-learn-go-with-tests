package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nubmers := []int{1, 2, 3, 4, 5}
		got := Sum(nubmers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, nubmers)
		}
	})
}
