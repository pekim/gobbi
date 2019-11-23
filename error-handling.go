package gobbi

import "github.com/pekim/gobbi/internal/gi"

/*
	SetErrorHandler sets a handler function that will be called
	if an error occurs dynamically referencing an entity while
	calling a function.

	Such errors may occur for a number of reasons.
	A library is not installed.
	A function is not available in the installed version of a library.
	A record or class is not available in the installed version of a library.
*/
func SetErrorHandler(handler gi.ErrorHandler) {
	gi.SetErrorHandler(handler)
}
