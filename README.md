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

### Setup:
#### First set all the environment variables.

For username and password
>export UNAME=<i>user_name</i><br>
>export UPASS=<i>user_password</i><br>
 
Then set the secret key of the server by: <br>

>export SSKEY=<i>server_secret_key </i><br>

To run the program use<br>

>go run main.go

### How to use:
#### POST
> /apis/v1/books<br>

The body of this post request contains the information of the book 
in json format. If the request is successful then the server add a new book
consisting the given information. The body of the response contains the 
book that just added.
#### GET
>/apis/v1/books<br>

it returns the list of information all books.

> apis/v1/books/{id}<br>

it returns the information of a specific book, having the book id {id}
#### PUT
>apis/v1/books/{id}<br>



#### DELETE
> apis/v1/books/{id}<br>


