package utils

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
