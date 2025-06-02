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

// Structures are Go’s way of aggregating data types together in a meaningful entity. In our case, because the structure
// represents a test case, we’ll name it testCase. Our structure needs only to be accessible in the TestGreet function
// (and nowhere else), so let’s define it there:
type testCase struct {
	lang language
	want string
}

func TestGreet(t *testing.T) {
	t.Parallel()

	var tests = map[string]testCase{
		"English":                 {EN, "Hello, World"},
		"French":                  {FR, "Bonjour, le Monde"},
		"Akkadian, not supported": {language("akk"), "unsupported language \"akk\""},
		"Greek":                   {EL, "γεια σου κόσμε"},
		"Urdu":                    {UR, "ہیلو دنیا"},
		"Hebrew":                  {HE, "שלום עולם"},
		"German":                  {DE, "Hallo Welt"},
		"Vietnamese":              {VI, "xin chào thế giới"},
		"Empty":                   {language(""), "unsupported language \"\""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := greetWithPhrasebook(test.lang)
			if got != test.want {
				t.Errorf("expected: %q, got: %q", test.want, got)
			}
		})
	}
}
