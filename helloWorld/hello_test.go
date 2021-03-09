package helloWorld

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "")
		want := englishHelloPrefix + "Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when empty string is provided", func(t *testing.T) {
		want := "Hello, world"
		got := hello("", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("In spanish", func(t *testing.T) {
		got := hello("melodie", "spanish")
		want := "Hola, melodie"
		assertCorrectMessage(t, got, want)
	})


	t.Run("In french", func(t *testing.T) {
		got := hello("melodie", "french")
		want := "Bonjour, melodie"
		assertCorrectMessage(t, got, want)
	})


	t.Run("In unsupported language", func(t *testing.T) {
		got := hello("melodie", "indian")
		want := "Hello, melodie"
		assertCorrectMessage(t, got, want)
	})
}
