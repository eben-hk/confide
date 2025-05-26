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
	SCodeSynced    = 9
	SCodeLogin     = 10
	SCodeLogout    = 11

	// Fail response code
	FCodeDataNotFound         = 1
	FCodeDb                   = 2
	FCodeBadRequest           = 3
	FCodeUriNotFound          = 4
	FCodeSelfRefer            = 5
	FCodeMethodNotAllowed     = 6
	FCodeDuplicateEntry       = 7
	FCodeMaximumLimit         = 8
	FCodeUnauthorized         = 9
	FCodeOneDevicePolicy      = 10
	FCodeInvalidLogin         = 11
	FCodeOTP                  = 12
	FCodeGeneral              = 13
	FCodeInternalServerError  = 14
	FCodeIncompleteData       = 15
	FCodeTokenInvalid         = 16
	FCodeTokenExpired         = 17
	FCodeInvalidFileExtension = 18
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
	SCodeSynced:    "successfully synced",
	SCodeLogin:     "successfully login",
	SCodeLogout:    "successfully logout",
}

// Fail message
var fCodeText = map[int]string{
	FCodeDataNotFound:         "data not found",
	FCodeDb:                   "data source error",
	FCodeBadRequest:           "bad request",
	FCodeUriNotFound:          "uri not found",
	FCodeSelfRefer:            "self-refer is not allowed",
	FCodeMethodNotAllowed:     "method not allowed",
	FCodeDuplicateEntry:       "duplicate entry",
	FCodeMaximumLimit:         "maximum limit reached",
	FCodeUnauthorized:         "unauthorized",
	FCodeOneDevicePolicy:      "multiple accounts in one device is not allowed",
	FCodeInvalidLogin:         "invalid login",
	FCodeOTP:                  "otp code: invalid / expired / reach max attempt",
	FCodeGeneral:              "general error",
	FCodeInternalServerError:  "internal server error",
	FCodeIncompleteData:       "incomplete data",
	FCodeTokenInvalid:         "invalid token",
	FCodeTokenExpired:         "expired token",
	FCodeInvalidFileExtension: "invalid file extension",
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
	SCodeSynced:    http.StatusOK,
	SCodeLogin:     http.StatusOK,
	SCodeLogout:    http.StatusOK,
}

// Fail http status
var fHttpStatus = map[int]int{
	FCodeDataNotFound:         http.StatusOK,
	FCodeDb:                   http.StatusInternalServerError,
	FCodeBadRequest:           http.StatusBadRequest,
	FCodeUriNotFound:          http.StatusNotFound,
	FCodeSelfRefer:            http.StatusOK,
	FCodeMethodNotAllowed:     http.StatusMethodNotAllowed,
	FCodeDuplicateEntry:       http.StatusOK,
	FCodeMaximumLimit:         http.StatusOK,
	FCodeUnauthorized:         http.StatusUnauthorized,
	FCodeOneDevicePolicy:      http.StatusOK,
	FCodeInvalidLogin:         http.StatusOK,
	FCodeOTP:                  http.StatusOK,
	FCodeGeneral:              http.StatusOK,
	FCodeInternalServerError:  http.StatusInternalServerError,
	FCodeIncompleteData:       http.StatusOK,
	FCodeTokenInvalid:         http.StatusOK,
	FCodeTokenExpired:         http.StatusOK,
	FCodeInvalidFileExtension: http.StatusUnprocessableEntity,
}
