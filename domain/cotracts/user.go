package contracts

import "github.com/msakp/golang-web-template/domain/dto"

type UserRepository interface{
	Create(u *dto.UserRegister) error
}



type UserService interface{
	Create(u *dto.UserRegister) error
}
