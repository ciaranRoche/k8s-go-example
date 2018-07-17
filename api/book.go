package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) toJSON() []byte {
	toJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return toJSON
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	Book{Title: "Attach Handle for Free Parasol", Author: "C. Roche", ISBN: "0987654321"},
	Book{Title: "Only 988 Years Until We Live Underwater", Author: "C. Roche", ISBN: "1357908642"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
