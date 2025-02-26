package service

import (
	"reflect"

	"github.com/msakp/golang-web-template/internal/domain/contracts"
	"github.com/msakp/golang-web-template/internal/domain/dto"
	"github.com/msakp/golang-web-template/internal/wrapper"

	"github.com/go-playground/validator/v10"
)

var _ contracts.ValidatorService = (*validatorService)(nil)

type validatorService struct {
	validator *validator.Validate
}

func NewValidatorService() *validatorService {
	service := validatorService{}
	service.validator = validator.New(validator.WithRequiredStructEnabled())
	service.initTags()
	return &service
}

func (s *validatorService) initTags() {
	// s.validator.RegisterValidation("some-tag", validationFunc)
}

// dto must be a pointer
func (s *validatorService) ValidateRequestData(dto any) *dto.HttpErr {
	err := s.validator.Struct(dto)
	if err != nil {
		return wrapper.ValidationErr(err.Error())
	}
	return nil
}

func (s *validatorService) ValidateRequestSlice(dataModels any) *dto.HttpErr {
	// no additional validation required (type assertion of dataModel being a slice is already executed in appropriate handler with fiber.BodyParser())
	// usage without dataModels type assertion being a slice is prohibited!
	v := reflect.ValueOf(dataModels)
	for i := range v.Len() {
		err := s.validator.Struct(v.Index(i).Interface())
		if err != nil {
			return wrapper.ValidationErr(err.Error())
		}
	}
	return nil
}
