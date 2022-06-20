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

var UserInfo = map[string]string{
	"admin": "1234",
}

var ServerSecretKey = "Superm4n"
