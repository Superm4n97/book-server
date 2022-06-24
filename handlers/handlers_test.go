package handlers

import (
	"bytes"
	"encoding/base64"
	"github.com/Superm4n97/Book-Server/middlewares"
	"github.com/Superm4n97/Book-Server/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

//var R = chi.NewRouter()

//func init() {
//	initialize()
//}

type Test struct {
	method         string
	rurl           string
	body           io.Reader
	expectedStatus int
	authen         string
}

func addSomeBooks() {
	for i := 1; i <= 5; i++ {
		Books[i] = model.Book{
			Id:    i,
			ISBN:  "7890",
			Title: "hello world",
			Authors: []model.Author{
				{
					Name:  "Alexa D.",
					Email: "alexa@gmail.com",
					City:  "Boston",
				},
				{
					Name:  "Robert P.",
					Email: "probert@gmail.com",
					City:  "London",
				},
			},
		}
	}
}

func initialize() *chi.Mux {
	R := chi.NewRouter()
	R.Use(middleware.Logger)
	R.Route("/apis/v1/books", func(r chi.Router) {
		//R.Use(middlewares.Authentication)
		r.Use(middlewares.Authentication)

		r.Get("/", GetAllBooks)
		r.Post("/", AddNewBook)
		r.Get("/{id}", GetBookInfoWithID)
		r.Put("/{id}", UpdateBookInformation)
		r.Delete("/{id}", RemoveBookFromList)
	})
	R.Post("/apis/v1/login", Login)

	addSomeBooks()
	return R
}

/*
func TestAddNewBook(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	httptest.NewRecorder()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Add new book test with sample book data",
			args: args{
				w: httptest.NewRecorder(),
				r: getRequestForAddNewBook(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddNewBook(tt.args.w, tt.args.r)
			gotResp := tt.args.w.(*httptest.ResponseRecorder)
			if gotResp.Result().StatusCode != http.StatusOK {
				t.Fail()
			}
		})
	}
}

*/

///*
func TestLogin(t *testing.T) {
	R := initialize()

	tests := []Test{
		{
			method:         "POST",
			rurl:           "/apis/v1/login",
			body:           nil,
			expectedStatus: 200,
			authen:         "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:1234")),
		},
		{
			method:         "POST",
			rurl:           "/apis/v1/login",
			body:           nil,
			expectedStatus: 401, //wrong username or password
			authen:         "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:12345")),
		},
		{
			method:         "POST",
			rurl:           "/apis/v1/login",
			body:           nil,
			expectedStatus: 401,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.rurl, test.body)

		req.Header.Set("Authorization", test.authen)

		res := httptest.NewRecorder()

		R.ServeHTTP(res, req)
		assert.Equal(t, res.Code, test.expectedStatus)
	}
}

//*/

func TestGetAllBooks(t *testing.T) {
	R := initialize()

	tests := []Test{
		{
			method:         "GET",
			rurl:           "/apis/v1/books",
			body:           nil,
			expectedStatus: 200,
			authen:         "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:1234")),
			//authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",

		},
	}
	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.rurl, test.body)
		req.Header.Set("Authorization", test.authen)
		res := httptest.NewRecorder()

		R.ServeHTTP(res, req)
		assert.Equal(t, res.Code, test.expectedStatus)
	}
}

func TestGetBookInfoWithID(t *testing.T) {
	R := initialize()

	tests := []Test{
		{
			method:         "GET",
			rurl:           "/apis/v1/books/1",
			body:           nil,
			expectedStatus: 200,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
		{
			method:         "GET",
			rurl:           "/apis/v1/books/7",
			expectedStatus: 404,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.rurl, test.body)
		req.Header.Set("Authorization", test.authen)
		res := httptest.NewRecorder()

		R.ServeHTTP(res, req)

		assert.Equal(t, res.Code, test.expectedStatus)
	}
}

func TestAddNewBook(t *testing.T) {
	R := initialize()

	tests := []Test{
		{
			method: "POST",
			rurl:   "/apis/v1/books",
			body: bytes.NewReader([]byte(`{
    "id" : 10,
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
}`)),
			expectedStatus: 200,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
		{
			method: "POST",
			rurl:   "/apis/v1/books",
			body: bytes.NewReader([]byte(`{
    "id" : 10,
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
}`)),
			expectedStatus: 400,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.rurl, test.body)
		req.Header.Set("Authorization", test.authen)

		res := httptest.NewRecorder()

		R.ServeHTTP(res, req)
		assert.Equal(t, res.Code, test.expectedStatus)
	}
}

func TestUpdateBookInformation(t *testing.T) {
	R := initialize()

	tests := []Test{
		{
			method: "PUT",
			rurl:   "/apis/v1/books/1",
			body: bytes.NewReader([]byte(`{
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
}`)),
			expectedStatus: 200,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
		{
			method: "PUT",
			rurl:   "/apis/v1/books/15",
			body: bytes.NewReader([]byte(`{
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
}`)),
			expectedStatus: 404,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
		{
			method: "PUT",
			rurl:   "/apis/v1/books/2",
			body: bytes.NewReader([]byte(`{
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
}`)),
			expectedStatus: 400,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.rurl, test.body)
		req.Header.Set("Authorization", test.authen)

		res := httptest.NewRecorder()

		R.ServeHTTP(res, req)
		assert.Equal(t, res.Code, test.expectedStatus)
	}
}

func TestRemoveBookFromList(t *testing.T) {
	R := initialize()

	tests := []Test{
		{
			method:         "DELETE",
			rurl:           "/apis/v1/books/5",
			body:           nil,
			expectedStatus: 200,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
		{
			method:         "DELETE",
			rurl:           "/apis/v1/books/5",
			body:           nil,
			expectedStatus: 404,
			authen:         "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6IjEyMzQifQ.-i0If6mLRGHQeXDkK_NQbqxjfJbvKXcVU6GF6e55FuM",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.rurl, test.body)
		req.Header.Set("Authorization", test.authen)
		res := httptest.NewRecorder()

		R.ServeHTTP(res, req)

		assert.Equal(t, res.Code, test.expectedStatus)
	}
}
