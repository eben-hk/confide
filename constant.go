package confide

import "net/http"

const (
	// Success response code
	SCodeCreated   = 1
	SCodeFound     = 2
	SCodeUp        = 3
	SCodeGenerated = 4
	SCodeDeleted   = 5
	SCodeUpdated   = 6
	SCodeValidated = 7
	SCodeUploaded  = 8

	// Fail response code
	FCodeDataNotFound     = 1
	FCodeDb               = 2
	FCodeBadRequest       = 3
	FCodeUriNotFound      = 4
	FCodeSelfRefer        = 5
	FCodeMethodNotAllowed = 6
	FCodeDuplicateEntry   = 7
	FCodeMaximumLimit     = 8
	FCodeUnauthorized     = 9
	FCodeOneDevicePolicy  = 10
)

// Success message
var sCodeText = map[int]string{
	SCodeCreated:   "successfully created",
	SCodeFound:     "data found",
	SCodeUp:        "service is up and running",
	SCodeGenerated: "successfully generated",
	SCodeDeleted:   "successfully deleted",
	SCodeUpdated:   "successfully updated",
	SCodeValidated: "successfully validated",
	SCodeUploaded:  "successfully uploaded",
}

// Fail message
var fCodeText = map[int]string{
	FCodeDataNotFound:     "data not found",
	FCodeDb:               "data source error",
	FCodeBadRequest:       "bad request",
	FCodeUriNotFound:      "uri not found",
	FCodeSelfRefer:        "self-refer is not allowed",
	FCodeMethodNotAllowed: "method not allowed",
	FCodeDuplicateEntry:   "duplicate entry",
	FCodeMaximumLimit:     "maximum limit reached",
	FCodeUnauthorized:     "unauthorized",
	FCodeOneDevicePolicy:  "multiple accounts in one device is not allowed",
}

// Success http status
var sHttpStatus = map[int]int{
	SCodeCreated:   http.StatusCreated,
	SCodeFound:     http.StatusOK,
	SCodeUp:        http.StatusOK,
	SCodeGenerated: http.StatusOK,
	SCodeDeleted:   http.StatusOK,
	SCodeUpdated:   http.StatusOK,
	SCodeValidated: http.StatusOK,
	SCodeUploaded:  http.StatusOK,
}

// Fail http status
var fHttpStatus = map[int]int{
	FCodeDataNotFound:     http.StatusOK,
	FCodeDb:               http.StatusInternalServerError,
	FCodeBadRequest:       http.StatusBadRequest,
	FCodeUriNotFound:      http.StatusNotFound,
	FCodeSelfRefer:        http.StatusOK,
	FCodeMethodNotAllowed: http.StatusMethodNotAllowed,
	FCodeDuplicateEntry:   http.StatusConflict,
	FCodeMaximumLimit:     http.StatusOK,
	FCodeUnauthorized:     http.StatusUnauthorized,
	FCodeOneDevicePolicy:  http.StatusOK,
}
