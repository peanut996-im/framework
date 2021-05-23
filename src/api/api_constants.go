package api

import "errors"

const (
	EventAuth         = "auth"
	EventLoad         = "load"
	EventAddFriend    = "addFriend"
	EventDeleteFriend = "deleteFriend"
	EventCreateGroup  = "createGroup"
	EventJoinGroup    = "joinGroup"
	EventLeaveGroup   = "leaveGroup"
	EventChat         = "chat"
)

const (
	ErrorCodeOK      = 0
	ErrorSignInvalid = 1000 + iota
	ErrorTokenInvalid
	ErrorAuthFailed
	ErrorHttpInnerError
	ErrorHttpParamInvalid
	ErrorHttpResourceExists
	ErrorHttpResourceNotFound

	HTTPMethodGet    string = "GET"
	HTTPMethodPost   string = "POST"
	HTTPMethodPut    string = "PUT"
	HTTPMethodPatch  string = "PATCH"
	HTTPMethodDelete string = "DELETE"
	HTTPMethodHead   string = "HEAD"
)

const (
	NewRequestError    string = "New http request err: %v, url: %v"
	DoRequestError     string = "Do http request err: %v, url: %v"
	DoGetRequestError  string = "Do get http request err: %v, url: %v"
	DoPostRequestError string = "Do post http request err: %v, url: %v"
	ReadRespBodyError  string = "Read resp body err: %v, url: %v"
	MarshalJsonError   string = "Marshal json  err: %v"
	UnmarshalJsonError string = "Unmarshal json  err: %v"
)

var (
	respCodeInfo = map[int]string{
		ErrorCodeOK:               "Success",
		ErrorTokenInvalid:         "Token invalid.",
		ErrorHttpInnerError:       "Http inner error err: %v",
		ErrorHttpParamInvalid:     "Http param invalid err: %v",
		ErrorSignInvalid:          "Sign invalid.",
		ErrorHttpResourceExists:   "Http resource already exists. err: %v",
		ErrorHttpResourceNotFound: "Http resource not found. err: %v",
		ErrorAuthFailed:           "Authentication failed",

		// ERROR_REDIS:              "Redis error err: %v",
		// ERROR_MONGO:              "Mongo error err: %v",
		// ERROR_UNMARSHAL_JSON:     "Unmarshal json  err: %v",
		// ERROR_MARSHAL_JSON:       "Marshal json  err: %v",
	}
)

func ErrorCodeToString(code int) string {
	return respCodeInfo[code]
}

func ErrorStringToError(err string) error {
	return errors.New(err)
}
