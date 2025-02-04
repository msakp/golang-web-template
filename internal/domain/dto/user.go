package dto

type UserRegister struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	PasswordUnhashed string `json:"password"`
}

type UserView struct {
	Name  string
	Email string
}
