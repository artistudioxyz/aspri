package library

// CustomError is a custom error type with a message.
type CustomError struct {
	message string
}

// Error returns the error message.
func (e *CustomError) Error() string {
	return e.message
}
