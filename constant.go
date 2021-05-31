package confide

import "net/http"

const (
	// Success response code
	SCodeCreated = 1
	SCodeFound   = 2
	SCodeUp      = 3

	// Fail response code
	FCodeDataNotFound = 1
	FCodeDb           = 2
	FCodeBadRequest   = 3
	FCodeUriNotFound  = 4
	FCodeSelfRefer    = 5
)

// Success text response
var sCodeText = map[int]string{
	SCodeCreated: "successfully created",
	SCodeFound:   "data found",
	SCodeUp:      "we are up and running",
}

// Fail text response
var fCodeText = map[int]string{
	FCodeDataNotFound: "data not found",
	FCodeDb:           "database error",
	FCodeBadRequest:   "bad request",
	FCodeUriNotFound:  "uri not found",
	FCodeSelfRefer:    "self-refer is not allowed",
}

// Success http status
var sHttpStatus = map[int]int{
	SCodeCreated: http.StatusCreated,
	SCodeFound:   http.StatusOK,
	SCodeUp:      http.StatusOK,
}

// Fail http status
var fHttpStatus = map[int]int{
	FCodeDataNotFound: http.StatusOK,
	FCodeDb:           http.StatusInternalServerError,
	FCodeBadRequest:   http.StatusBadRequest,
	FCodeUriNotFound:  http.StatusNotFound,
	FCodeSelfRefer:    http.StatusBadRequest,
}
