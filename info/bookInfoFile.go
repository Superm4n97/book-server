package info

import (
	"encoding/json"
	"log"
	"net/http"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	City  string `json:"city"`
}
type Book struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	ISBN    string   `json:"isbn"`
	Authors []Author `json:"authors"`
}

var Books = make(map[int]Book)

//GetAllBooks gives the response about all the books in the Books
//it encodes Books in JSON format
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	//the JSON request is converted into a [string]string map
	//element can be accessed by request["json_field_name"]
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	Books[book.Id] = book

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Println("failed to write data in response body")
	}
	w.WriteHeader(http.StatusOK)
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
