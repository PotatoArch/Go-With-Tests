package main

import "testing"

func TestHello(t *testing.T) {
	// Testing For names
	t.Run("Say Hello To Someone", func(t *testing.T) {
		got := Hello("Potato")
		want := "Hello, Potato"

		isCorrect(t, got, want)
	})

	// Testing For Default Values
	t.Run("Testing For Default Values", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		isCorrect(t, got, want)
	})
}

func isCorrect(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got : %q want : %q", got, want)
	}
}
