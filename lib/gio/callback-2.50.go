// This is a generated file - DO NOT EDIT
// +build gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

	GFile* callback_vfsfilelookupfuncHandler(GObject *, GVfs*, const char*, gpointer, gpointer);

*/
import "C"

var callbackVfsfilelookupfuncId int
var callbackVfsfilelookupfuncMap = make(map[int]VfsfilelookupfuncCallback)
var callbackVfsfilelookupfuncLock sync.RWMutex

// VfsfilelookupfuncCallback is a callback function for a 'VfsFileLookupFunc' callback.
type VfsfilelookupfuncCallback func(vfs *Vfs, identifier string) *File
