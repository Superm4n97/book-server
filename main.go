package main

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/actions"
	"github.com/Superm4n97/Book-Server/middlewares"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	r := chi.NewRouter()

	//r.Use(middlewares.BasicAuthentication)
	r.Get("/ping", pong)
	r.Route("/apis/v1/books", func(r chi.Router) {
		//r.Use(middlewares.BasicAuthentication)
		r.Use(middlewares.BearerAuthentication)

		r.Get("/", actions.GetAllBooks)
		r.Post("/", actions.AddNewBook)
		r.Get("/{id}", actions.GetBookInfoWithID)
		r.Put("/{id}", actions.UpdateBookInformation)
		r.Delete("/{id}", actions.RemoveBookFromList)
	})
	r.Post("/apis/v1/login", actions.Login)

	//r.Get("/apis/v1/books", actions.GetAllBooks)
	//r.Post("/apis/v1/books", actions.AddNewBook)
	//r.Get("/apis/v1/books/{id}", actions.GetBookInfoWithID)
	//r.Put("/apis/v1/books/{id}", actions.UpdateBookInformation)
	//r.Delete("/apis/v1/books/{id}", actions.RemoveBookFromList)

	fmt.Println("Server Running on port: 8080")
	http.ListenAndServe(":8080", r)

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

*/
