package domain

type ErrorWithStatus struct {
	Status  int
	Message string
}

func (e *ErrorWithStatus) Error() string {
	return e.Message
}
