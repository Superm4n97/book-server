package actions

import (
	"encoding/json"
	"github.com/Superm4n97/Book-Server/info"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

var Books = make(map[int]info.Book)

func GetBookInfoWithID(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")
	bid, err := strconv.Atoi(bookId)

	if err != nil || Books[bid].Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book id"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books[bid])
}

//GetAllBooks gives the response about all the books in the Books
//it encodes Books in JSON format
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var booksSlice []info.Book
	for _, val := range Books {
		booksSlice = append(booksSlice, val)
	}

	json.NewEncoder(w).Encode(booksSlice)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book info.Book

	//the JSON request is converted into a [string]string map
	//element can be accessed by request["json_field_name"]
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	Books[book.Id] = book

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Println("failed to write data in response body")
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateBookInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stBookId := chi.URLParam(r, "id")
	bookId, err := strconv.Atoi(stBookId)

	if err != nil || Books[bookId].Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Inappropriate update operation."))

		return
	}

	var book info.Book
	e := json.NewDecoder(r.Body).Decode(&book)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON format"))

		return
	}

	if book.Id != bookId {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book Id mismatched"))

		return
	}

	Books[bookId] = book

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func RemoveBookFromList(w http.ResponseWriter, r *http.Request) {
	strBookId := chi.URLParam(r, "id")
	bookId, err := strconv.Atoi(strBookId)

	if err != nil || Books[bookId].Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book id for deletion!!!"))

		return
	}

	delete(Books, bookId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted!!!"))
}
