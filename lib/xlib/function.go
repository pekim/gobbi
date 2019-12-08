// Code generated - DO NOT EDIT.

package xlib

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	"sync"
)

var openDisplayFunction *gi.Function
var openDisplayFunction_Once sync.Once

func openDisplayFunction_Set() error {
	var err error
	openDisplayFunction_Once.Do(func() {
		openDisplayFunction, err = gi.FunctionInvokerNew("xlib", "open_display")
	})
	return err
}

// OpenDisplay is a representation of the C type XOpenDisplay.
func OpenDisplay() {

	err := openDisplayFunction_Set()
	if err == nil {
		openDisplayFunction.Invoke(nil, nil)
	}

	return
}
