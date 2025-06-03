package main

import (
	"slices"
	"testing"
)

type testCase struct {
	bookwormsFile string
	want          []Bookworm
	wantErr       bool
}

var (
	handmaidsTale            = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	orxyAndCrake             = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar               = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	nineteenEightyFour       = Book{Author: "George Orwell", Title: "1984"}
	animalFarm               = Book{Author: "George Orwell", Title: "Animal Farm"}
	braveNewWorld            = Book{Author: "Aldous Huxley", Title: "Brave New World"}
	theGreatGatsby           = Book{Author: "J.K.Rowling", Title: "The Great Gatsby"}
	janeEyre                 = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	theAdventuresOfTomSawyer = Book{Author: "Jane Austen", Title: "The Adventures of Tom Sawyer"}
)

func TestLoadBookworms(t *testing.T) {
	t.Parallel()

	testCases := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{orxyAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file does not exist": {
			bookwormsFile: "testdata/does-not-exist.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid Json": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.bookwormsFile, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)
			if err != nil && !tc.wantErr {
				t.Fatalf("unexpected error: %v", err.Error())
			}

			if err == nil && tc.wantErr {
				t.Fatalf("expected error %s, got none", err.Error())
			}

			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("Different result! got %v, want %v", got, tc.want)
			}
		})
	}
}

func equalBookworms(t *testing.T, got, want []Bookworm) bool {
	t.Helper()
	if len(got) != len(want) {
		return false
	}
	for i, bookworm := range got {
		if bookworm.Name != want[i].Name {
			return false
		}

		if !slices.Equal(bookworm.Books, want[i].Books) {
			return false
		}
	}
	return true
}

func equalBooks(t *testing.T, got, want []Book) bool {
	t.Helper()
	if len(got) != len(want) {
		return false
	}
	for i, book := range got {
		if book.Author != want[i].Author || book.Title != want[i].Title {
			return false
		}
	}

	return true
}
