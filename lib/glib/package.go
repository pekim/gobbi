// This is a generated file - DO NOT EDIT

package glib

// #cgo pkg-config: glib-2.0
import "C"

// Error makes Error implement the error interface
func (e Error) Error() string {
	return e.Message
}
