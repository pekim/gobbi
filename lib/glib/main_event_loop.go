package glib

import (
	"sync"
)

// #include <glib.h>
/*
	gboolean idleAddHandler (gpointer user_data);

	static void idle_add(gpointer data) {
		g_idle_add(idleAddHandler, data);
	}
*/
import "C"

// IdleAddCallback is a function that can be passed to IdleAdd, that will
// be called from glib's main event loop.
type IdleAddCallback func() bool

// IdleAddOnceCallback is a function that can be passed to IdleAddOnce, that will
// be called once from glib's main event loop.
type IdleAddOnceCallback func()

var idleAddId int
var idleAddMap = make(map[int]IdleAddCallback)
var idleAddLock sync.Mutex

/*
IdleAdd adds a function to be called whenever there are no higher priority events
pending to the default main loop.
If the function returns false it is automatically removed from the list of event
sources and will not be called again.
*/
func IdleAdd(callback IdleAddCallback) {
	idleAddLock.Lock()
	defer idleAddLock.Unlock()

	idleAddId++
	idleAddMap[idleAddId] = callback
	C.idle_add(C.gpointer(uintptr(idleAddId)))
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
		return SOURCE_REMOVE
	})
}

//export idleAddHandler
func idleAddHandler(data C.gpointer) C.gboolean {
	idleAddLock.Lock()
	defer idleAddLock.Unlock()

	index := int(uintptr(data))
	callback := idleAddMap[index]

	if !callback() {
		delete(idleAddMap, index)
		return C.FALSE
	}
	return C.TRUE
}
