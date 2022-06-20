package main

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/handlers"
	"github.com/Superm4n97/Book-Server/middlewares"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	r := chi.NewRouter()

	//r.Use(middlewares.Authentication)
	r.Get("/ping", pong)
	r.Route("/apis/v1/books", func(r chi.Router) {
		//r.Use(middlewares.Authentication)
		r.Use(middlewares.Authentication)

		r.Get("/", handlers.GetAllBooks)
		r.Post("/", handlers.AddNewBook)
		r.Get("/{id}", handlers.GetBookInfoWithID)
		r.Put("/{id}", handlers.UpdateBookInformation)
		r.Delete("/{id}", handlers.RemoveBookFromList)
	})
	r.Post("/apis/v1/login", handlers.Login)

	//r.Get("/apis/v1/books", handlers.GetAllBooks)
	//r.Post("/apis/v1/books", handlers.AddNewBook)
	//r.Get("/apis/v1/books/{id}", handlers.GetBookInfoWithID)
	//r.Put("/apis/v1/books/{id}", handlers.UpdateBookInformation)
	//r.Delete("/apis/v1/books/{id}", handlers.RemoveBookFromList)

	fmt.Println("Server Running on port: 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Failed to start the server")
		return
	}
}

/*
JSON Format Data
====================

{
    "id" : 1,
    "title" : "lazy lad",
    "isbn" : "01623",
    "authors" : [
        {
            "name":"abc",
            "email": "abc.com",
            "city" : "dhaka"
        },
        {
            "name":"vd",
            "email": "vd.com",
            "city" : "sylhet"
        }
    ]
}

Admin JWT Token
==================
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiMTIzNCIsInVzZXJOYW1lIjoiYWRtaW4ifQ.DX81oiggc9PA0qhU-LSJflUUTmqfOU1sig4wk39DPmA
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiMTIzNCIsInVzZXJOYW1lIjoiYWRtaW4ifQ.DX81oiggc9PA0qhU-LSJflUUTmqfOU1sig4wk39DPmA
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM
*/
