package dto

type HttpErr struct {
	HttpCode int `json:"-"`
	Data interface{} `json:"data" validate:"required" example:"some error description"`
}

