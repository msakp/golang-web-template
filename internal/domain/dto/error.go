package dto

type HttpErr struct {
	HttpCode    int    `json:"-"`
	Message     string `json:"err" validate:"required" example:"Short error message : 'Not Found' | 'Internal Server Error' | etc"`
	Description string `json:"description" validate:"required" example:"verbose error description"`
}

func (err *HttpErr) Error() string {
	return err.Description
}

type OkResponse struct {
	Message string `json:"status" validate:"required" example:"ok"`
}

const (
	MsgUserNotFound      = "user with given id/email not found"
	MsgUserAlreadyExists = "user with given emain already exists"

	MsgInvalidPassword = "invalid password"
)
