package contracts

import "solution/internal/domain/dto"

type ValidatorService interface {
	ValidateRequestData(dataModel any) *dto.HttpErr
	ValidateRequestSlice(dataModels any) *dto.HttpErr
}
