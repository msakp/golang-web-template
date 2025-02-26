package contracts

import "github.com/msakp/golang-web-template/internal/domain/dto"

type ValidatorService interface {
	ValidateRequestData(dataModel any) *dto.HttpErr
	ValidateRequestSlice(dataModels any) *dto.HttpErr
}
