// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "sync"

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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	void callback_filemeasureprogresscallbackHandler(GObject *, gboolean, guint64, guint64, guint64, gpointer, gpointer);

*/
import "C"

var callbackFilemeasureprogresscallbackId int
var callbackFilemeasureprogresscallbackMap = make(map[int]FilemeasureprogresscallbackCallback)
var callbackFilemeasureprogresscallbackLock sync.RWMutex

// FilemeasureprogresscallbackCallback is a callback function for a 'FileMeasureProgressCallback' callback.
type FilemeasureprogresscallbackCallback func(reporting bool, currentSize uint64, numDirs uint64, numFiles uint64)

//export callback_filemeasureprogresscallbackHandler
func callback_filemeasureprogresscallbackHandler(_ *C.GObject, c_reporting C.gboolean, c_current_size C.guint64, c_num_dirs C.guint64, c_num_files C.guint64, data C.gpointer) {
	callbackFilemeasureprogresscallbackLock.RLock()
	defer callbackFilemeasureprogresscallbackLock.RUnlock()

	reporting := c_reporting == C.TRUE

	currentSize := uint64(c_current_size)

	numDirs := uint64(c_num_dirs)

	numFiles := uint64(c_num_files)

	index := int(uintptr(data))
	callback := callbackFilemeasureprogresscallbackMap[index].callback
	callback(reporting, currentSize, numDirs, numFiles, userData)
}
