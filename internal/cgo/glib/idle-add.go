package glib

import (
	"sync"
)

// #cgo pkg-config: glib-2.0
// #include <glib.h>
/*
	gboolean idleAddHandler (gpointer user_data);

	static void idle_add(gpointer data) {
		g_idle_add(idleAddHandler, data);
	}
*/
import "C"

type IdleAddCallback func() bool

var idleAddId int
var idleAddMap = make(map[int]IdleAddCallback)
var idleAddLock sync.Mutex

func IdleAdd(callback IdleAddCallback) {
	idleAddLock.Lock()
	defer idleAddLock.Unlock()

	idleAddId++
	idleAddMap[idleAddId] = callback

	C.idle_add(C.gpointer(uintptr(idleAddId)))
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
