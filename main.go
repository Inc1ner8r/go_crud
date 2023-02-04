package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inciner8r/go_crud_t2/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Routes(r)
	http.ListenAndServe(":8080", r)
}
