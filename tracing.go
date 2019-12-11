package gobbi

import (
	"github.com/pekim/gobbi/internal/cgo"
)

/*
	SetTraceHandler sets a handler function that will be called
	with messages detailing function calls made through
	gobject introspection.

	Pass nil to remove a previous set handler.
*/
func SetTraceHandler(handler cgo.TraceHandler) {
	cgo.SetTraceHandler(handler)
}
