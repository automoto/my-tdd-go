package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to a person by name", func(t *testing.T) {
		got := Hello("john", "english")
		want:= "Hello, john!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say Hello, Person when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, person!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("john", "spanish")
		want := "Hola, john!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("john", "french")
		want := "Bonjour, john!"
		assertCorrectMessage(t, got, want)
	})
}
