package dto

type UserRegister struct {
	Name             string `json:"name" validate:"required" minLenth:"3" maxLength:"16" example:"msa"`
	Email            string `json:"email" validate:"required" example:"yoyoyo@femail.ru"`
	PasswordUnhashed string `json:"password" validate:"required" example:"qwerty123_AOISROKT(:#*L(*))"`
	PasswordHashed   string `json:"-"`
}

type UserLogin struct {
	Email            string `json:"email" validate:"required" example:"yoyoyo@femail.ru"`
	PasswordUnHashed string `json:"password" validate:"required" example:"qwerty234sraiekvaroisehw{}$"`
}

type UserView struct {
	Name  string
	Email string
}

type UserAuthResponse struct {
	Token string `json:"token" validate:"required"`
}
