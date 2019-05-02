// This is a generated file - DO NOT EDIT
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

// AppInfoMonitor is a wrapper around the C record GAppInfoMonitor.
type AppInfoMonitor struct {
	native *C.GAppInfoMonitor
}

func AppInfoMonitorNewFromC(u unsafe.Pointer) *AppInfoMonitor {
	c := (*C.GAppInfoMonitor)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppInfoMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppInfoMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetSocketAddressNewFromString is a wrapper around the C function g_inet_socket_address_new_from_string.
func InetSocketAddressNewFromString(address string, port uint32) *InetSocketAddress {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_port := (C.guint)(port)

	retC := C.g_inet_socket_address_new_from_string(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Notification is a wrapper around the C record GNotification.
type Notification struct {
	native *C.GNotification
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	c := (*C.GNotification)(u)
	if c == nil {
		return nil
	}

	g := &Notification{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Notification) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NotificationNew is a wrapper around the C function g_notification_new.
func NotificationNew(title string) *Notification {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.g_notification_new(c_title)
	retGo := NotificationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Subprocess is a wrapper around the C record GSubprocess.
type Subprocess struct {
	native *C.GSubprocess
}

func SubprocessNewFromC(u unsafe.Pointer) *Subprocess {
	c := (*C.GSubprocess)(u)
	if c == nil {
		return nil
	}

	g := &Subprocess{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Subprocess) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Subprocess) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_subprocess_new : unsupported parameter ... : varargs

// SubprocessNewv is a wrapper around the C function g_subprocess_newv.
func SubprocessNewv(argv []string, flags SubprocessFlags) (*Subprocess, error) {
	c_argv_array := make([]*C.gchar, len(argv)+1, len(argv)+1)
	for i, item := range argv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_argv_array[i] = c
	}
	c_argv_array[len(argv)] = nil
	c_argv_arrayPtr := &c_argv_array[0]
	c_argv := (**C.gchar)(unsafe.Pointer(c_argv_arrayPtr))

	c_flags := (C.GSubprocessFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_subprocess_newv(c_argv, c_flags, &cThrowableError)
	retGo := SubprocessNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SubprocessLauncher is a wrapper around the C record GSubprocessLauncher.
type SubprocessLauncher struct {
	native *C.GSubprocessLauncher
}

func SubprocessLauncherNewFromC(u unsafe.Pointer) *SubprocessLauncher {
	c := (*C.GSubprocessLauncher)(u)
	if c == nil {
		return nil
	}

	g := &SubprocessLauncher{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SubprocessLauncher) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SubprocessLauncher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SubprocessLauncherNew is a wrapper around the C function g_subprocess_launcher_new.
func SubprocessLauncherNew(flags SubprocessFlags) *SubprocessLauncher {
	c_flags := (C.GSubprocessFlags)(flags)

	retC := C.g_subprocess_launcher_new(c_flags)
	retGo := SubprocessLauncherNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}
