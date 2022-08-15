package model

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

//read form environment variable
//os.Getenv("UNAME")
//os.Getenv("UPASS")
//var ServerSecretKey = os.Getenv("SSKEY")

//testing purpose
var UNAME = "admin"
var UPASS = "1234"
var ServerSecretKey = "Superm4n"

//store in userinfo database
var UserInfo = map[string]string{
	UNAME: UPASS,
}
