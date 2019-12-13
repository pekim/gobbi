package glib

import (
	"github.com/pekim/gobbi/internal/cgo/glib"
)

// IdleAddCallback is a function that can be passed to IdleAdd, that will
// be called from glib's main event loop.
type IdleAddCallback glib.IdleAddCallback

// IdleAddOnceCallback is a function that can be passed to IdleAddOnce, that will
// be called once from glib's main event loop.
type IdleAddOnceCallback func()

/*
IdleAdd adds a function to be called whenever there are no higher priority events
pending to the default main loop.

If the function returns false it is automatically removed from the list of event
sources and will not be called again.
The constants SOURCE_REMOVE and SOURCE_CONTINUE may be used for the return value.
*/
func IdleAdd(callback IdleAddCallback) {
	glib.IdleAdd(glib.IdleAddCallback(callback))
}

/*
IdleAddOnce adds a function to be called whenever there are no higher priority
events pending to the default main loop.

It is a convenience function that works the same way as IdleAdd, other than
the callback has no return value, and it will be removed from the list event
sources after one invocation. That is, the callback will only be called once.
*/
func IdleAddOnce(callback IdleAddOnceCallback) {
	IdleAdd(func() bool {
		callback()
		return false
	})
}
