package contracts

type ValidatorService interface{
	ValidateRequestData(d interface{}) bool
}
