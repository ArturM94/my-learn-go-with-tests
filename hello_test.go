package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Artur", "")
		want := "Hello, Artur"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Ukrainian", func(t *testing.T) {
		got := Hello("Артур", "Ukrainian")
		want := "Привіт, Артур"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Artur", "French")
		want := "Bonjour, Artur"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() define assertCorrectMessage function as a helper
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
