package glib

// Error makes Error implement the error interface
func (e *Error) Error() string {
	return e.Message
}
