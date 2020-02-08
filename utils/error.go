package utils

import "net/http"

type ApiError struct {
	Code   int
	Reason string
	error  error
}

func (e ApiError) Error() string {
	if e.error != nil {
		return e.error.Error()
	}
	return e.Reason
}

func CheckInternalError(w http.ResponseWriter, err error) bool {
	if err != nil {
		JsonError(err.Error(), w, 500)
		return true
	}
	return false
}
