package main

import "testing"

func TestHello(t *testing.T) {
	// Testing For names
	t.Run("Say Hello To Someone", func(t *testing.T) {
		got := Hello("Potato")
		want := "Hello, Potato"

		if got != want {
			t.Errorf("got : %q want : %q", got, want)
		}
	})

	// Testing For Default Values
	t.Run("Testing For Default Values", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got : %q want : %q", got, want)
		}
	})
}
