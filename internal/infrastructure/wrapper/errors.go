package wrapper

import "github.com/msakp/golang-web-template/internal/domain/dto"



func ServiceErr(httpCode int, message string) *dto.HttpErr{
	return &dto.HttpErr{HttpCode: httpCode, Message: message}
}
