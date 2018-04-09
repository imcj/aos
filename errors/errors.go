package errors

type Error interface {
	error
	GetCode() int
}

type StatusError struct {
	Code    int
	Message string
	Payload interface{}
}

func (e StatusError) Error() string {
	return e.Message
}

func (e StatusError) GetCode() int {
	return e.Code
}

func New(code int, message string) error {
	return StatusError{
		code,
		message,
		nil,
	}
}
