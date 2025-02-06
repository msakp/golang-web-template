package dto

type HttpErr struct {
	Message string `json:"err" validate:"required" example:"some error description"`
}
