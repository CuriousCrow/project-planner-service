package typed_error

type ErrorType byte

const (
	NotFound ErrorType = iota
	BadRequest
	ServerError
)

// New ...
func New(et ErrorType, message string) error {
	return TypedError{
		Type:    et,
		Message: message,
	}
}

// NewServerError ...
func NewServerError(err error) error {
	return TypedError{
		Type:    ServerError,
		Message: err.Error(),
	}
}

type TypedError struct {
	Type    ErrorType
	Message string
}

func (e TypedError) Error() string {
	return e.Message
}
