package net

import (
	"fmt"
)

const (
	// response code
	ERROR_CODE_OK      = 0
	ERROR_SIGN_INVAILD = 1000 + iota
	ERROR_TOKEN_INVALID
	ERROR_AUTH_FAILED
	ERROR_HTTP_INNER_ERROR
	ERROR_HTTP_PARAM_INVALID
	ERROR_HTTP_RESOURCE_EXISTS
	ERROR_HTTP_RESOURCE_NOT_FOUND

	// HTTP Method
	HTTP_METHOD_GET    string = "GET"
	HTTP_METHOD_POST   string = "POST"
	HTTP_METHOD_PUT    string = "PUT"
	HTTP_METHOD_PATCH  string = "PATCH"
	HTTP_METHOD_DELETE string = "DELETE"
	HTTP_METHOD_HEAD   string = "HEAD"

	// Detail Error
	NEW_REQUEST_ERROR     string = "New http request err: %v, url: %v"
	DO_REQUEST_ERROR      string = "Do http request err: %v, url: %v"
	DO_GET_REQUEST_ERROR  string = "Do get http request err: %v, url: %v"
	DO_POST_REQUEST_ERROR string = "Do post http request err: %v, url: %v"
	READ_RESP_BODY_ERROR  string = "Read resp body err: %v, url: %v"
	MARSHAL_JSON_ERROR    string = "Marshal json  err: %v"
	UNMARSHAL_JSON_ERROR  string = "Unmarshal json  err: %v"
)

var (
	respCodeInfo = map[int]string{
		ERROR_CODE_OK:                 "Success",
		ERROR_TOKEN_INVALID:           "Token invalid.",
		ERROR_HTTP_INNER_ERROR:        "Http inner error err: %v",
		ERROR_HTTP_PARAM_INVALID:      "Http param invalid err: %v",
		ERROR_SIGN_INVAILD:            "Sign invaild.",
		ERROR_HTTP_RESOURCE_EXISTS:    "Http resource already exists. err: %v",
		ERROR_HTTP_RESOURCE_NOT_FOUND: "Http resource not found. err: %v",
		ERROR_AUTH_FAILED:             "Authentication failed",

		// ERROR_REDIS:              "Redis error err: %v",
		// ERROR_MONGO:              "Mongo error err: %v",
		// ERROR_UNMARSHAL_JSON:     "Unmarshal json  err: %v",
		// ERROR_MARSHAL_JSON:       "Marshal json  err: %v",
	}

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

func ErrorCodeToString(code int) string {
	return respCodeInfo[code]
}

type BaseRepsonse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewBaseResponse(code int, data interface{}) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    code,
		Data:    data,
		Message: ErrorCodeToString(code),
	}
}

func NewHttpInnerErrorResp(err error) *BaseRepsonse {
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

func NewResourceExistsResp(err error) *BaseRepsonse {
	return &BaseRepsonse{
		Code:    ERROR_HTTP_RESOURCE_EXISTS,
		Data:    nil,
		Message: fmt.Sprintf(ErrorCodeToString(ERROR_HTTP_RESOURCE_EXISTS), err),
	}
}
