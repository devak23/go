package main

import "testing"

func TestGreet_English(t *testing.T) {
	want := "Hello, World"
	got := greet(EN)
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour, le Monde"
	got := greet(lang)
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_Spanish(t *testing.T) {
	lang := language("es")
	want := ""
	got := greet(lang)
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
