// Package api
// @Title  api_response.go
// @Description  record some defined response.
// @Author  peanut996
// @Update  peanut996  2021/5/22 0:22
package api

import (
	"fmt"
)

type BaseRepsonse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	SignInvaildResp = &BaseRepsonse{
		Code:    ErrorSignInvalid,
		Message: fmt.Sprint(ErrorCodeToString(ErrorSignInvalid)),
		Data:    nil,
	}

	ResourceExistsResp = &BaseRepsonse{
		Code:    ErrorHttpResourceExists,
		Message: fmt.Sprintf(ErrorCodeToString(ErrorHttpResourceExists), nil),
		Data:    nil,
	}

	ResourceNotFoundResp = &BaseRepsonse{
		Code:    ErrorHttpResourceNotFound,
		Message: fmt.Sprintf(ErrorCodeToString(ErrorHttpResourceNotFound), nil),
		Data:    nil,
	}

	AuthFaildResp = &BaseRepsonse{
		Code:    ErrorAuthFailed,
		Message: fmt.Sprintf(ErrorCodeToString(ErrorAuthFailed)),
		Data:    nil,
	}

	TokenInvaildResp = &BaseRepsonse{
		Code:    ErrorTokenInvalid,
		Message: fmt.Sprint(ErrorCodeToString(ErrorTokenInvalid)),
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
		Code:    ErrorHttpInnerError,
		Message: fmt.Sprintf(ErrorCodeToString(ErrorHttpInnerError), err),
		Data:    nil,
	}
}

func NewSuccessResponse(data interface{}) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    ErrorCodeOK,
		Data:    data,
		Message: ErrorCodeToString(ErrorCodeOK),
	}
}

func NewResourceExistsResponse(err error) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    ErrorHttpResourceExists,
		Data:    nil,
		Message: fmt.Sprintf(ErrorCodeToString(ErrorHttpResourceExists), err),
	}
}
