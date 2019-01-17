// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	gboolean callback_datagrambasedsourcefuncHandler(GObject *, GDatagramBased*, GIOCondition, gpointer, gpointer);

*/
import "C"

var callbackDatagrambasedsourcefuncId int
var callbackDatagrambasedsourcefuncMap = make(map[int]DatagrambasedsourcefuncCallback)
var callbackDatagrambasedsourcefuncLock sync.RWMutex

// DatagrambasedsourcefuncCallback is a callback function for a 'DatagramBasedSourceFunc' callback.
type DatagrambasedsourcefuncCallback func(datagramBased *DatagramBased, condition glib.IOCondition) bool
