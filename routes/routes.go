package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inciner8r/go_crud_t2/controllers"
)

func Routes(r *mux.Router) {
	r.Path("/getBooks").Methods(http.MethodGet).HandlerFunc(controllers.GetBooks)
	r.Path("/postBook").Methods(http.MethodPost).HandlerFunc(controllers.PostBook)
	r.Path("/deleteBook/{id}").Methods(http.MethodDelete).HandlerFunc(controllers.DeleteBook)
	r.Path("/getBook/{id}").Methods(http.MethodGet).HandlerFunc(controllers.GetBookById)
	r.Path("/updateBook/{id}").Methods(http.MethodPut).HandlerFunc(controllers.UpdateBook)
	r.Path("/").Methods(http.MethodGet).HandlerFunc(controllers.Hello)
}
