package confide

import "net/http"

const (
	// Success response code
	SCodeCreated   = 1
	SCodeFound     = 2
	SCodeUp        = 3
	SCodeGenerated = 4
	SCodeDeleted   = 5

	// Fail response code
	FCodeDataNotFound     = 1
	FCodeDb               = 2
	FCodeBadRequest       = 3
	FCodeUriNotFound      = 4
	FCodeSelfRefer        = 5
	FCodeMethodNotAllowed = 6
)

// Success message
var sCodeText = map[int]string{
	SCodeCreated:   "successfully created",
	SCodeFound:     "data found",
	SCodeUp:        "we are up and running",
	SCodeGenerated: "successfully generated",
	SCodeDeleted:   "successfully deleted",
}

// Fail message
var fCodeText = map[int]string{
	FCodeDataNotFound:     "data not found",
	FCodeDb:               "database error",
	FCodeBadRequest:       "bad request",
	FCodeUriNotFound:      "uri not found",
	FCodeSelfRefer:        "self-refer is not allowed",
	FCodeMethodNotAllowed: "method not allowed",
}

// Success http status
var sHttpStatus = map[int]int{
	SCodeCreated:   http.StatusCreated,
	SCodeFound:     http.StatusOK,
	SCodeUp:        http.StatusOK,
	SCodeGenerated: http.StatusOK,
	SCodeDeleted:   http.StatusOK,
}

// Fail http status
var fHttpStatus = map[int]int{
	FCodeDataNotFound:     http.StatusOK,
	FCodeDb:               http.StatusInternalServerError,
	FCodeBadRequest:       http.StatusBadRequest,
	FCodeUriNotFound:      http.StatusNotFound,
	FCodeSelfRefer:        http.StatusBadRequest,
	FCodeMethodNotAllowed: http.StatusMethodNotAllowed,
}
