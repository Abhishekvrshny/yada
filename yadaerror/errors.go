package yadaerror

import "encoding/json"

// Error is custom error type that implements
// error. It additionally holds internal codes
type Error struct {
	InternalCode int    `json:"internal_code"`
	Message      string `json:"message"`
}

// Error is the overriden function
func (h Error) Error() string {
	return h.Message
}

// ToJSONBytes marshals Error to json
func (h Error) ToJSONBytes() []byte {
	b, _ := json.Marshal(h)
	return b
}

// New returns a new Error
func New(msg string, internalCode int) Error {
	return Error{internalCode, msg}
}
