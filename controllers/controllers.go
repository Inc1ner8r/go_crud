package controllers

import (
	"fmt"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all books")
}
