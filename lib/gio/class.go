// This is a generated file - DO NOT EDIT

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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
import "C"

// Blacklisted : GAppLaunchContext

// ApplicationCommandLine is a wrapper around the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native *C.GApplicationCommandLine
	// Private : parent_instance
	// Private : priv
}

func ApplicationCommandLineNewFromC(u unsafe.Pointer) *ApplicationCommandLine {
	c := (*C.GApplicationCommandLine)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLine{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ApplicationCommandLine) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationCommandLine with another ApplicationCommandLine, and returns true if they represent the same GObject.
func (recv *ApplicationCommandLine) Equals(other *ApplicationCommandLine) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ApplicationCommandLine) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ApplicationCommandLine.
// Exercise care, as this is a potentially dangerous function if the Object is not a ApplicationCommandLine.
func CastToApplicationCommandLine(object *gobject.Object) *ApplicationCommandLine {
	return ApplicationCommandLineNewFromC(object.ToC())
}

// Blacklisted : GBufferedInputStream

// Blacklisted : GBufferedOutputStream

// Blacklisted : GBytesIcon

// Blacklisted : GCancellable

// Blacklisted : GCharsetConverter

// Blacklisted : GConverterInputStream

// Blacklisted : GConverterOutputStream

// Blacklisted : GDBusActionGroup

// Blacklisted : GDBusMenuModel

// Blacklisted : GDataInputStream

// Blacklisted : GDataOutputStream

// Blacklisted : GDesktopAppInfo

// Blacklisted : GEmblem

// Blacklisted : GEmblemedIcon

// Blacklisted : GFileEnumerator

// Blacklisted : GFileIOStream

// Blacklisted : GFileIcon

// Blacklisted : GFileInfo

// Blacklisted : GFileInputStream

// Blacklisted : GFileMonitor

// Blacklisted : GFileOutputStream

// Blacklisted : GFilenameCompleter

// Blacklisted : GFilterInputStream

// Blacklisted : GFilterOutputStream

// Blacklisted : GIOModule

// Blacklisted : GIOStream

// Blacklisted : GInetAddress

// Blacklisted : GInetSocketAddress

// Blacklisted : GInputStream

// Blacklisted : GListStore

// Blacklisted : GMemoryInputStream

// Blacklisted : GMemoryOutputStream

// Blacklisted : GMountOperation

// Blacklisted : GNativeVolumeMonitor

// Blacklisted : GNetworkAddress

// Blacklisted : GNetworkService

// Blacklisted : GOutputStream

// Blacklisted : GPermission

// Blacklisted : GProxyAddressEnumerator

// Blacklisted : GResolver

// Blacklisted : GSettings

// Blacklisted : GSettingsBackend

// Blacklisted : GSimpleAction

// Blacklisted : GSimpleAsyncResult

// Blacklisted : GSimplePermission

// Blacklisted : GSimpleProxyResolver

// Blacklisted : GSocketAddress

// Blacklisted : GSocketAddressEnumerator

// Blacklisted : GSocketControlMessage

// Blacklisted : GTask

// Blacklisted : GTcpWrapperConnection

// Blacklisted : GThemedIcon

// Blacklisted : GUnixConnection

// Blacklisted : GUnixFDList

// Blacklisted : GUnixFDMessage

// Blacklisted : GUnixInputStream

// Blacklisted : GUnixMountMonitor

// Blacklisted : GUnixOutputStream

// Blacklisted : GUnixSocketAddress

// Blacklisted : GVfs

// Blacklisted : GVolumeMonitor

// Blacklisted : GZlibCompressor

// Blacklisted : GZlibDecompressor
