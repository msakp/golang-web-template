package dto

type UserRegister struct {
	Name             string `json:"name" validate:"required" minLenth:"3" maxLength:"16" example:"msa"`
	Email            string `json:"email" format:"email" example:"yoyoyo@femail.ru"`
	PasswordUnhashed string `json:"password" example:"qwerty123_AOISROKT(:#*L(*))"`
	PasswordHashed   string `json:"-"`
}

type UserView struct {
	Name  string
	Email string
}

type UserRegisterResponse struct {
	Token string `json:"token"`
}
