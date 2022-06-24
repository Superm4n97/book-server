package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Superm4n97/Book-Server/auth"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var Books = make(map[int]model.Book)

func GetBookInfoWithID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookId := chi.URLParam(r, "id")
	bid, err := strconv.Atoi(bookId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := Books[bid]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("no book with id %d", bid)))
		return
	}

	err = json.NewEncoder(w).Encode(Books[bid])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

//GetAllBooks gives the response about all the books in the Books
//it encodes Books in JSON format
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var booksSlice []model.Book
	for _, val := range Books {
		booksSlice = append(booksSlice, val)
	}

	err := json.NewEncoder(w).Encode(booksSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book

	//the JSON request is converted into a [string]string map
	//element can be accessed by request["json_field_name"]
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := Books[book.Id]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("book with the id already exists"))
		return
	}

	w.Write([]byte("book successfully added. "))
	Books[book.Id] = book

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Println("failed to write data in response body")
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	return
}

func UpdateBookInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stBookId := chi.URLParam(r, "id")
	bookId, err := strconv.Atoi(stBookId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := Books[bookId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("no book with id %d", bookId)))
		return
	}

	var book model.Book
	e := json.NewDecoder(r.Body).Decode(&book)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e.Error()))
		return
	}

	if book.Id != bookId {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("book id do not match with url id"))
		return
	}

	Books[bookId] = book

	w.Write([]byte("book successfully updated"))

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func RemoveBookFromList(w http.ResponseWriter, r *http.Request) {
	strBookId := chi.URLParam(r, "id")
	bookId, err := strconv.Atoi(strBookId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := Books[bookId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no book for deletion"))
		return
	}

	delete(Books, bookId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book successfully deleted"))
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	headerString := r.Header.Get("Authorization")

	if headerString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("login information missing"))
		return
	}

	infoStr := strings.Split(headerString, " ")

	if len(infoStr) != 2 || infoStr[0] != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid login information"))
		return
	}

	temp, err := base64.StdEncoding.DecodeString(infoStr[1])

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	infoStr = strings.Split(string(temp), ":")

	if model.UserInfo[infoStr[0]] != infoStr[1] {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong username or password"))
		return
	}

	w.Write([]byte("successfully logged in. bearer token -> "))

	//============GENERATE TOKEN===================
	st := auth.CreateJwtToken(infoStr[0], infoStr[1])
	_, err = w.Write([]byte(st))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
