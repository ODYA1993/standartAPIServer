package models

//type User struct {
//	ID         int    `json:"id"`
//	Name       string `json:"username"`
//	Surname    string `json:"email"`
//	Patronymic string `json:"password"`
//}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	}
}
