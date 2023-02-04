package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/inciner8r/go_crud_t2/models"
)

var books = []models.Book{}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatal("get books failed")
	}

}
func PostBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatal("error in decoding")
	}

	books = append(books, book)
}
