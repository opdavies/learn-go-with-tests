package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Oliver")
	want := "Hello, failure"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
