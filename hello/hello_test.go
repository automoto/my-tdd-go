package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("john")
	want:= "Hello, john!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
