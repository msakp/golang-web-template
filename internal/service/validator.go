package service

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/msakp/golang-web-template/internal/domain/contracts"
)

var _ contracts.ValidatorService = (*validatorService)(nil)

type validatorService struct{
	validator *validator.Validate

}

func NewValidatorService() *validatorService{
	var service validatorService
	service.validator = validator.New(validator.WithRequiredStructEnabled())
	service.initTags()
	return &service
}


func (s *validatorService) initTags(){
	s.validator.RegisterValidation("password", func(fl validator.FieldLevel) bool{
		value := fl.Field().String()
		length := fl.Field().Len()
		requiredLength := 8 <= length && length <= 64

		hasLower := value != strings.ToUpper(value)
		hasUpper := value != strings.ToLower(value)
		hasDigit := strings.IndexAny(value, "0123456789") != -1
		hasSpecial := strings.IndexAny(value, "!@#$%^") != -1
		return hasLower && hasUpper && hasDigit && hasSpecial && requiredLength
	})

	s.validator.RegisterValidation("username", func(fl validator.FieldLevel) bool{
		length := fl.Field().Len()
		requiredLength := 3 <= length && length <= 16
		return requiredLength
	})
}

func (s *validatorService) ValidateRequestData(d interface{}) bool{
	err := s.validator.Struct(d)
	log.Println("validated")
	if err != nil{
		log.Printf("ERR: %s , %T", err, err)
		return false
	}

	return true
} 
