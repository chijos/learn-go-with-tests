package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Chirag", "")
    want := "Hello, Chirag"
    assertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
    got := Hello("", "English")
    want := "Hello, world"
    assertCorrectMessage(t, got, want)
  })

  t.Run("greet in Spanish", func(t *testing.T) {
    got := Hello("Chirag", "Spanish")
    want := "Hola, Chirag"
    assertCorrectMessage(t, got, want)
  })

  t.Run("greet in French", func(t *testing.T) {
    got := Hello("Chirag", "French")
    want := "Bonjour, Chirag"
    assertCorrectMessage(t, got, want)
  })

  t.Run("greet in Hindi", func(t *testing.T) {
    got := Hello("Chirag", "Hindi")
    want := "Namaste, Chirag"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
      t.Errorf("got %q want %q", got, want)
  }
}
