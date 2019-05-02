// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

// DesktopAppInfoNewFromKeyfile is a wrapper around the C function g_desktop_app_info_new_from_keyfile.
func DesktopAppInfoNewFromKeyfile(keyFile *glib.KeyFile) *DesktopAppInfo {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	retC := C.g_desktop_app_info_new_from_keyfile(c_key_file)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// EmblemNew is a wrapper around the C function g_emblem_new.
func EmblemNew(icon *Icon) *Emblem {
	c_icon := (*C.GIcon)(icon.ToC())

	retC := C.g_emblem_new(c_icon)
	retGo := EmblemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// EmblemNewWithOrigin is a wrapper around the C function g_emblem_new_with_origin.
func EmblemNewWithOrigin(icon *Icon, origin EmblemOrigin) *Emblem {
	c_icon := (*C.GIcon)(icon.ToC())

	c_origin := (C.GEmblemOrigin)(origin)

	retC := C.g_emblem_new_with_origin(c_icon, c_origin)
	retGo := EmblemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// EmblemedIconNew is a wrapper around the C function g_emblemed_icon_new.
func EmblemedIconNew(icon *Icon, emblem *Emblem) *EmblemedIcon {
	c_icon := (*C.GIcon)(icon.ToC())

	c_emblem := (*C.GEmblem)(C.NULL)
	if emblem != nil {
		c_emblem = (*C.GEmblem)(emblem.ToC())
	}

	retC := C.g_emblemed_icon_new(c_icon, c_emblem)
	retGo := EmblemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}
