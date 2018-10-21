package glib

import "sync"

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

var idleAddId int
var idleAddMap = make(map[int]IdleAddCallback)
var idleAddLock sync.Mutex

/*
IdleAdd adds a function to be called whenever there are no higher priority events
pending to the default main loop.
If the function returns falseit is automatically removed from the list of event
sources and will not be called again.
*/
func IdleAdd(callback IdleAddCallback) {
	idleAddLock.Lock()
	defer idleAddLock.Unlock()

	idleAddId++
	idleAddMap[idleAddId] = callback
	C.idle_add(C.gpointer(uintptr(idleAddId)))
}

//export idleAddHandler
func idleAddHandler(data C.gpointer) C.gboolean {
	index := int(uintptr(data))
	callback := idleAddMap[index]
	return boolToGboolean(callback())
}