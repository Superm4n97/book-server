package main

import (
	"context"
	"fmt"
	"github.com/Superm4n97/Book-Server/handlers"
	"github.com/Superm4n97/Book-Server/middlewares"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"os/signal"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Pong"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

var R = chi.NewRouter()

func RouterSetup() {
	//R = chi.NewRouter()

	R.Use(chimiddleware.Logger)

	//R.Use(middlewares.Authentication)
	R.Get("/ping", Pong)

	R.Route("/apis/v1/books", func(r chi.Router) {
		//R.Use(middlewares.Authentication)
		r.Use(middlewares.Authentication)

		r.Get("/", handlers.GetAllBooks)
		r.Post("/", handlers.AddNewBook)
		r.Get("/{id}", handlers.GetBookInfoWithID)
		r.Put("/{id}", handlers.UpdateBookInformation)
		r.Delete("/{id}", handlers.RemoveBookFromList)
	})
	R.Post("/apis/v1/login", handlers.Login)

	//r.Get("/apis/v1/books", handlers.GetAllBooks)
	//r.Post("/apis/v1/books", handlers.AddNewBook)
	//r.Get("/apis/v1/books/{id}", handlers.GetBookInfoWithID)
	//r.Put("/apis/v1/books/{id}", handlers.UpdateBookInformation)
	//r.Delete("/apis/v1/books/{id}", handlers.RemoveBookFromList)
}

func Init() {
	RouterSetup()
}

func main() {
	Init()

	fmt.Println("Server Running on port: 8080")
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)

	server := &http.Server{
		Addr:    ":8080",
		Handler: R,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	<-stopCh
	fmt.Println("Server is shutting down!!")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("failed to shutdown the server")
		return
	}
	fmt.Println("Server is gracefully shutdown")
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

admin:1234:Superm4n
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM

'{"id":1,"title":"llad","isbn":"4324","authors":[{"name":"abul","email":"abul@gmail.com","city":"dhaka"}]}'
*/
