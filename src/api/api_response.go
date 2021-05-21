// Package api
// @Title  api_response.go
// @Description  record some defined response.
// @Author  peanut996
// @Update  peanut996  2021/5/22 0:22
package api

import (
	"fmt"
	"framework/api/model"
)

type BaseRepsonse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
}

var (
	SignInvaildResp = &BaseRepsonse{
		Code:    ERROR_SIGN_INVAILD,
		Message: fmt.Sprint(ErrorCodeToString(ERROR_SIGN_INVAILD)),
		Data:    nil,
	}

	ResourceExistsResp = &BaseRepsonse{
		Code:    ERROR_HTTP_RESOURCE_EXISTS,
		Message: fmt.Sprintf(ErrorCodeToString(ERROR_HTTP_RESOURCE_EXISTS), nil),
		Data:    nil,
	}

	ResourceNotFoundResp = &BaseRepsonse{
		Code:    ERROR_HTTP_RESOURCE_NOT_FOUND,
		Message: fmt.Sprintf(ErrorCodeToString(ERROR_HTTP_RESOURCE_NOT_FOUND), nil),
		Data:    nil,
	}

	AuthFaildResp = &BaseRepsonse{
		Code:    ERROR_AUTH_FAILED,
		Message: fmt.Sprintf(ErrorCodeToString(ERROR_AUTH_FAILED)),
		Data:    nil,
	}

	TokenInvaildResp = &BaseRepsonse{
		Code:    ERROR_TOKEN_INVALID,
		Message: fmt.Sprint(ErrorCodeToString(ERROR_TOKEN_INVALID)),
		Data:    nil,
	}
)

func NewBaseResponse(code int, data interface{}) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    code,
		Data:    data,
		Message: ErrorCodeToString(code),
	}
}

func NewHttpInnerErrorResponse(err error) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    ERROR_HTTP_INNER_ERROR,
		Message: fmt.Sprintf(ErrorCodeToString(ERROR_HTTP_INNER_ERROR), err),
		Data:    nil,
	}
}

func NewSuccessResponse(data interface{}) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    ERROR_CODE_OK,
		Data:    data,
		Message: ErrorCodeToString(ERROR_CODE_OK),
	}
}

func NewResourceExistsResponse(err error) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    ERROR_HTTP_RESOURCE_EXISTS,
		Data:    nil,
		Message: fmt.Sprintf(ErrorCodeToString(ERROR_HTTP_RESOURCE_EXISTS), err),
	}
}
