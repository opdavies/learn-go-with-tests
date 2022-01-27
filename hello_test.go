package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("Oliver", "")
		want := "Hello, Oliver"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func (t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Welsh", func (t *testing.T) {
		got := Hello("Oliver", "Welsh")
		want := "Helo, Oliver"

		assertCorrectMessage(t, got, want)
	})
}
