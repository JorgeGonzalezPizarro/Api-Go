package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {

	router := chi.NewRouter()
	router.Get("/{todoID}", GetTodo)
	router.Delete("/{todoID}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodo)
	return router
}

func GetAllTodo(w http.ResponseWriter, request *http.Request) {

	todos := Model.todos
}
