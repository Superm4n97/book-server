package testdir

type FullName struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last_name"`
}

type TestInfo struct {
	Name  FullName `json:"name"`
	Email string   `json:"email"`
	Age   string   `json:"age"`
}

func GetTestStruct() TestInfo {
	return TestInfo{
		Name: FullName{
			FirstName: "Rasel ",
			LastName:  "Hossain",
		},
		Email: "rasehossain@gmail.com",
		Age:   "25",
	}
}
