package gobbi

import "github.com/pekim/gobbi/internal/cgo/gi"

func init() {
	SetErrorHandler(func(err error) {
		panic(err)
	})
}

/*
	SetErrorHandler sets a handler function that will be called
	if an error occurs while calling a function.

	Such errors may occur for a number of reasons.
	A library is not installed.
	A function is not available in the installed version of a library.
	A record or class is not available in the installed version of a library.

	Pass nil to remove a previous set handler.

	If no error handler is set,
	the default handler will panic when an error occurs.
*/
func SetErrorHandler(handler gi.ErrorHandler) {
	gi.SetErrorHandler(handler)
}
