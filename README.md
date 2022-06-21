# Book-Server
REST API Server using Go and Chi .It is a book server, that a user can add, update and remove any book
or a user can check the list of all books. A book object consist of,
<br>
<i>
> Book id<br>
> Book title<br>
> ISBN number<br>
> List of authors<br>
</i>

An author object contains,<br>
<i>
> Author Name<br>
> Email address<br>
> Living city<br>
</i>

<br>It contains both basic
and bearer authentication. <br>

<br>

### Setup (Server side):
#### First set all the environment variables.

For username and password
>$ export UNAME=<i>user_name</i><br>
>$ export UPASS=<i>user_password</i><br>
 
Then set the secret key of the server by: <br>

>$ export SSKEY=<i>server_secret_key </i><br>

To start the server use<br>

>$ go run main.go

### How to use:(Client side requests)
#### Login
>$ curl -X POST --user username:password localhost:8080/apis/v1/login
 

#### POST
>$ curl -X POST --user username:password -d <'json_book_object'> localhost:8080/apis/v1/books<br>

The body of this post request contains the information of the book 
in json format. If the request is successful then the server add a new book
consisting the given information. The body of the response contains the 
book that just added.
#### GET
Basic,
>$ curl --user <'username:password'> localhost:8080/apis/v1/books<br>

or, Bearer,

>$ curl -H "Authorization: Bearer <'token'>" localhost:8080/apis/v1/books


it returns the list of information all books.

>$ curl -H "Authorization: Bearer <'token''>" localhost:8080/apis/v1/books/{id}<br>

it returns the information of a specific book, having the book id {id}
#### PUT
>$ curl -X PUT --user <'username:password'> -d <'new josn data'> localhost:8080/apis/v1/books/{id}<br>

Updates a book information with newer one. The book id from the url and the given one should be mached.

#### DELETE
>$ curl -X DELETE --user <'username:password'> localhost/apis/v1/books/{id}<br>

Deletes a book information respect with book id = id

