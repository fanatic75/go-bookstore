package routes

import (
	"net/http"

	"github.com/fanatic75/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	//router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods(http.MethodDelete)
}
