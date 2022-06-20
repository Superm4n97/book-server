package actions

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Superm4n97/Book-Server/info"
	"github.com/Superm4n97/Book-Server/middlewares"
	"github.com/Superm4n97/Book-Server/testdir"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	//json.NewEncoder(w).Encode(booksSlice)

	json.NewEncoder(w).Encode(testdir.GetTestStruct())
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

func Login(w http.ResponseWriter, r *http.Request) {
	headerString := r.Header.Get("Authorization")

	if headerString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	infoStr := strings.Split(headerString, " ")

	if len(infoStr) != 2 || infoStr[0] != "Basic" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	temp, _ := base64.StdEncoding.DecodeString(infoStr[1])
	infoStr = strings.Split(string(temp), ":")

	if info.UserInfo[infoStr[0]] != infoStr[1] {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Username/Password!!!"))

		return
	}

	w.Write([]byte("Successfully Loged In........"))

	//
	//============GENERATE TOKEN===================
	st := middlewares.CreateJwtToken(infoStr[0], infoStr[1])
	w.Write([]byte(st))
}
