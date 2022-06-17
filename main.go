package main

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/actions"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", pong)
	r.Get("/apis/v1/books", actions.GetAllBooks)
	r.Post("/apis/v1/books", actions.AddNewBook)
	r.Get("/apis/v1/book/{id}", actions.GetBookInfoWithID)
	r.Put("/apis/v1/book/{id}", actions.UpdateBookInformation)
	r.Delete("/apis/v1/book/{id}", actions.RemoveBookFromList)

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
*/
