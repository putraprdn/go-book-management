package routes

import (
	"github.com/gorilla/mux"
	"github.com/putraprdn/go-book-management/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
