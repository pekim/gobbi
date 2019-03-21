// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
/*

	gboolean callback_cancellablesourcefuncHandler(GObject *, GCancellable*, gpointer);

*/
/*

	gboolean callback_pollablesourcefuncHandler(GObject *, GObject*, gpointer);

*/
import "C"

var callbackCancellablesourcefuncId int
var callbackCancellablesourcefuncMap = make(map[int]CancellablesourcefuncCallback)
var callbackCancellablesourcefuncLock sync.RWMutex

// CancellablesourcefuncCallback is a callback function for a 'CancellableSourceFunc' callback.
type CancellablesourcefuncCallback func(cancellable *Cancellable) bool

//export callback_cancellablesourcefuncHandler
func callback_cancellablesourcefuncHandler(_ *C.GObject, c_cancellable *C.GCancellable, c_user_data C.gpointer) C.gboolean {
	callbackCancellablesourcefuncLock.RLock()
	defer callbackCancellablesourcefuncLock.RUnlock()

	cancellable := CancellableNewFromC(unsafe.Pointer(c_cancellable))

	index := int(uintptr(c_user_data))
	callback := callbackCancellablesourcefuncMap[index]
	retGo := callback(cancellable)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackPollablesourcefuncId int
var callbackPollablesourcefuncMap = make(map[int]PollablesourcefuncCallback)
var callbackPollablesourcefuncLock sync.RWMutex

// PollablesourcefuncCallback is a callback function for a 'PollableSourceFunc' callback.
type PollablesourcefuncCallback func(pollableStream *gobject.Object) bool

//export callback_pollablesourcefuncHandler
func callback_pollablesourcefuncHandler(_ *C.GObject, c_pollable_stream *C.GObject, c_user_data C.gpointer) C.gboolean {
	callbackPollablesourcefuncLock.RLock()
	defer callbackPollablesourcefuncLock.RUnlock()

	pollableStream := gobject.ObjectNewFromC(unsafe.Pointer(c_pollable_stream))

	index := int(uintptr(c_user_data))
	callback := callbackPollablesourcefuncMap[index]
	retGo := callback(pollableStream)
	retC :=
		boolToGboolean(retGo)
	return retC
}
