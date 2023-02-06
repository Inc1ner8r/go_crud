package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/inciner8r/go_crud_t2/models"
)

var books = []models.Book{}

func indexById(books []models.Book, id int) int {
	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			return i
		}
	}
	return -1
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		log.Fatal(err.Error())
	}

}
func PostBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatal(err.Error())
	}

	books = append(books, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	index := indexById(books, idInt)

	if index < 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	books = append(books[:index], books[index+1:]...)
	w.WriteHeader(http.StatusOK)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	index := indexById(books, idInt)

	if index < 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books[index]); err != nil {
		log.Fatal(err.Error())
	}
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		log.Fatal(err.Error())
	}

	id := mux.Vars(r)["id"]
	idInt, _ := strconv.Atoi(id)
	index := indexById(books, idInt)

	books[index] = book
	w.WriteHeader(http.StatusOK)
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api is up")
}
