package handler

import "encoding/json"

type JsonError struct {
	Message string `json:"message"`
}

func jsonError(message string) []byte {
	myError := &JsonError{
		Message: message,
	}

	r, err := json.Marshal(myError)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
