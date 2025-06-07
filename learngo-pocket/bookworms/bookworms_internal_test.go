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

func TestBookCount(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{
					handmaidsTale, theBellJar,
				}},
				{Name: "Peggy", Books: []Book{
					orxyAndCrake, handmaidsTale, janeEyre,
				}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				orxyAndCrake:  1,
				janeEyre:      1,
			},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
			},
			want: map[Book]uint{},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{
					theBellJar, theBellJar,
				}},
			},
			want: map[Book]uint{
				theBellJar: 2,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("Different result! got %v, want %v", got, tc.want)
			}
		})
	}
}

// ///////// Helper methods ///////////////
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

// we are comparing two maps and checking if their sizes are equal, and the books at the given index are the same
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}
