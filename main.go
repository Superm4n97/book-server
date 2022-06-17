package main

import (
	"fmt"
	"github.com/Superm4n97/Book-Server/info"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/ping", pong)
	r.Get("/apis/v1/books", info.GetAllBooks)
	r.Post("/apis/v1/books", info.AddNewBook)

	fmt.Println("Server Running on port: 8080")
	http.ListenAndServe(":8080", r)

}
