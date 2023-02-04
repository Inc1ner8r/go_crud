package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inciner8r/go_crud_t2/controllers"
)

func Routes(r *mux.Router) {
	r.Methods(http.MethodGet).HandlerFunc(controllers.GetBooks)
}
