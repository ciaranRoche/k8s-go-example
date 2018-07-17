package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "A pain in the axe", Author: "C. Roche", ISBN: "12345trewq"}
	json := book.toJSON()

	assert.Equal(t, `{"title":"A pain in the axe","author":"C. Roche","isbn":"12345trewq"}`,
		string(json), "Book JSON-marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"A pain in the axe","author":"C. Roche","isbn":"12345trewq"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "A pain in the axe", Author: "C. Roche", ISBN: "12345trewq"},
		book, "Book JSON unmarshalling wrong.")
}
