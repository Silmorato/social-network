package errors

type APIError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Cause   error  `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(status int, message string, cause error) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
		Cause:   cause,
	}
}
