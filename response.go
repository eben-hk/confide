package confide

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data" default:""`
	HttpStatus int         `json:"-"`
	IsSuccess  bool        `json:"is_success"`
	Message    string      `json:"message" default:""`
}

// JSON send payload to consumer in JSON format
func JSON(w http.ResponseWriter, p Payload) {
	httpStatus := p.HttpStatus

	if p.IsSuccess {
		if p.Message == "" {
			p.Message = sCodeText[p.Code]
		}
		if httpStatus == 0 {
			httpStatus = sHttpStatus[p.Code]
		}
	} else if !p.IsSuccess {
		if p.Message == "" {
			p.Message = fCodeText[p.Code]
		}
		if httpStatus == 0 {
			httpStatus = fHttpStatus[p.Code]
		}
	}

	if p.Data == nil {
		p.Data = map[string]interface{}{}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
