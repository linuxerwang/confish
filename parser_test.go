package confish

import (
	"reflect"
	"strings"
	"testing"
)

type Book struct {
	Name   string          `cfg-attr:"name"`
	ISBN   string          `cfg-attr:"isbn"`
	Author []string        `cfg-attr:"author"`
	Year   int             `cfg-attr:"year"`
	Prices map[int]float32 `cfg-attr:"prices"`
}

type BookShelf struct {
	Category string  `cfg-attr:"category"`
	Books    []*Book `cfg-attr:"book"`
}

func TestParse(t *testing.T) {
	booksConf := strings.NewReader(`
	bookshelf {
		category: "Computer Technology"

		book {
			name:   "The Go Programming Language"
			isbn:   "978-0134190440"
			author: [
				"Alan Donovan",
				"Brian W. Kernighan",
			]
			year:   2015
			prices: {
				10:  30.49,
				50:  28.49,
				200: 25.49,
			}
		}

		book {
			name:   "Go in Action"
			isbn:   "978-1617291784"
			author: [
				"William Kennedy",
				"Brian Ketelsen",
				"Erik St. Martin",
			]
			year:   2015
			prices: {
				10:   26.99,
				50:   25.89,
				500:  23.29,
				1000: 20.29,
			}
		}

		book {
			name:   "Conf-ish in Go"
			author: []
			prices: {}
		}
	}
	`)

	bs := &BookShelf{}
	err := Parse(booksConf, bs)
	if err != nil {
		t.Fatalf("failed to parse confish file")
	}

	if bs.Category != "Philosophy" {
		t.Fatalf("got %s, want Philosophy", bs.Category)
	}

	if len(bs.Books) != 3 {
		t.Fatalf("got %d books, want 3", len(bs.Books))
	}

	var b *Book

	b = &Book{
		"The Go Programming Language",
		"978-0134190440",
		[]string{"Alan Donovan", "Brian W. Kernighan"},
		2015,
		map[int]float32{10: 30.49, 50: 28.49, 200: 25.49},
	}
	if !reflect.DeepEqual(bs.Books[0], b) {
		t.Fatalf("got book %+v, want %+v", bs.Books[0], b)
	}

	b = &Book{
		"Go in Action",
		"978-1617291784",
		[]string{"William Kennedy", "Brian Ketelsen", "Erik St. Martin"},
		2015,
		map[int]float32{10: 26.99, 50: 25.89, 500: 23.29, 1000: 20.29},
	}
	if !reflect.DeepEqual(bs.Books[1], b) {
		t.Fatalf("got book %+v, want %+v", bs.Books[1], b)
	}

	b = &Book{
		"Conf-ish in Go",
		"",
		[]string{},
		0,
		map[int]float32{},
	}
	if !reflect.DeepEqual(bs.Books[2], b) {
		t.Fatalf("got book %+v, want %+v", bs.Books[2], b)
	}
}
