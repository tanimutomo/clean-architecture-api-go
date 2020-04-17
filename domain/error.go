package domain

type ErrorResponse struct {
	Status  int
	Message string
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
