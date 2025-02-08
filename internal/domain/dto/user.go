package dto

import "github.com/google/uuid"

type UserRegister struct {
	Name             string `json:"name" validate:"required,username" minLength:"3" maxLength:"16" example:"msa"`
	Email            string `json:"email" validate:"required,email" example:"yoyoyo@femail.ru"`
	PasswordUnhashed string `json:"password" validate:"required,password" example:"qwerty123_AOISROKT(:#*L(*))"`
	PasswordHashed   string `json:"-"`
}

type UserLogin struct {
	Email            string `json:"email" validate:"required,required" example:"yoyoyo@femail.ru"`
	PasswordUnHashed string `json:"password" validate:"required,password" example:"qwerty234sraiekvaroisehw{}$"`
}

type UserView struct {
	Name  string `json:"name" example:"vanya228"`
	Email string `json:"email" example:"me@femail.ru"`
	PasswordHashed string `json:"password" example:"SOME_HASHED_PASSWD"`
}

type UserAuthResponse struct {
	Token string `json:"token" validate:"required"`
	Id uuid.UUID `json:"id" validate:"required"`
}
