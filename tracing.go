package gobbi

import (
	"github.com/pekim/gobbi/internal/cgo"
)

/*
	SetTraceHandler sets a handler function that will be called
	with messages detailing function calls and signal callbacks
	made through gobject introspection.

	Pass nil to remove a previously set handler.

	Alternatively an environment variable GOBBI_TRACE with a value
	of "1", "t", "T", "TRUE", "true", or "True"
	will cause tracing messages to be written to stdout.
*/
func SetTraceHandler(handler cgo.TraceHandler) {
	cgo.SetTraceHandler(handler)
}
