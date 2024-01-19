package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Artur")
	want := "Hello, Artur"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
