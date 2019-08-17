package yadaerror

import "encoding/json"

type HTTPError struct {
	Message string
}

func NewHTTPError(httpErr HTTPError) []byte{
	b, _ :=json.Marshal(httpErr)
	return b
}