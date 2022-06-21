package handlers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Superm4n97/Book-Server/middlewares"
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
		w.Write([]byte("Invalid book id"))
		return
	}

	if _, ok := Books[bid]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(Books[bid])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}

	if _, ok := Books[book.Id]; ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Books[book.Id] = book

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Println("failed to write data in response body")
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
		return
	}

	if _, ok := Books[bookId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var book model.Book
	e := json.NewDecoder(r.Body).Decode(&book)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if book.Id != bookId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Books[bookId] = book

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}

	if _, ok := Books[bookId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(Books, bookId)

	w.WriteHeader(http.StatusOK)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	headerString := r.Header.Get("Authorization")

	if headerString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	infoStr := strings.Split(headerString, " ")

	if len(infoStr) != 2 || infoStr[0] != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	temp, err := base64.StdEncoding.DecodeString(infoStr[1])

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	infoStr = strings.Split(string(temp), ":")

	if model.UserInfo[infoStr[0]] != infoStr[1] {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong username or password"))
		return
	}

	//============GENERATE TOKEN===================
	st := middlewares.CreateJwtToken(infoStr[0], infoStr[1])
	_, err = w.Write([]byte(st))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
