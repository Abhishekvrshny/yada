package yadaerror

import "encoding/json"

type Error struct {
	InternalCode int    `json:"internal_code"`
	Message      string `json:"message"`
}

func (h Error) Error() string {
	return h.Message
}

func (h Error) ToJSONBytes() []byte {
	b, _ := json.Marshal(h)
	return b
}

func New(msg string, internalCode int) Error {
	return Error{internalCode, msg}
}
