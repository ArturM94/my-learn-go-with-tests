package main

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	assertSum(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	assertSum(t, got, want)
}

func assertSum(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
