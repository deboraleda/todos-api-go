package main

import (
	"fmt"
	"net/http"
	"todos-api/dbconfig"
	"todos-api/internal/auth"
	"todos-api/internal/handlers"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {

	dbconfig.ConfigDB()

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/login", handlers.Login)
		r.Post("/user", handlers.CreateUser)
	})

	r.Group(func(r chi.Router) {
		r.Use(auth.MyMiddleware)
		r.Post("/", handlers.Create)
		r.Put("/{id}", handlers.Update)
		r.Delete("/{id}", handlers.Delete)
		r.Get("/", handlers.List)
		r.Get("/{id}", handlers.Get)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r)

	defer dbconfig.DB.Close()

}
