// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// RecentInfo is a wrapper around the C record GtkRecentInfo.
type RecentInfo struct {
	native *C.GtkRecentInfo
}

func RecentInfoNewFromC(u unsafe.Pointer) *RecentInfo {
	c := (*C.GtkRecentInfo)(u)
	if c == nil {
		return nil
	}

	g := &RecentInfo{native: c}

	return g
}

func (recv *RecentInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_recent_info_create_app_info : no return generator

// Exists is a wrapper around the C function gtk_recent_info_exists.
func (recv *RecentInfo) Exists() bool {
	retC := C.gtk_recent_info_exists((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_recent_info_get_added : no return generator

// GetAge is a wrapper around the C function gtk_recent_info_get_age.
func (recv *RecentInfo) GetAge() int32 {
	retC := C.gtk_recent_info_get_age((*C.GtkRecentInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_application_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : gtk_recent_info_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// GetDescription is a wrapper around the C function gtk_recent_info_get_description.
func (recv *RecentInfo) GetDescription() string {
	retC := C.gtk_recent_info_get_description((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_recent_info_get_display_name.
func (recv *RecentInfo) GetDisplayName() string {
	retC := C.gtk_recent_info_get_display_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_gicon : no return generator

// Unsupported : gtk_recent_info_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// GetIcon is a wrapper around the C function gtk_recent_info_get_icon.
func (recv *RecentInfo) GetIcon(size int32) *gdkpixbuf.Pixbuf {
	c_size := (C.gint)(size)

	retC := C.gtk_recent_info_get_icon((*C.GtkRecentInfo)(recv.native), c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMimeType is a wrapper around the C function gtk_recent_info_get_mime_type.
func (recv *RecentInfo) GetMimeType() string {
	retC := C.gtk_recent_info_get_mime_type((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_modified : no return generator

// GetPrivateHint is a wrapper around the C function gtk_recent_info_get_private_hint.
func (recv *RecentInfo) GetPrivateHint() bool {
	retC := C.gtk_recent_info_get_private_hint((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShortName is a wrapper around the C function gtk_recent_info_get_short_name.
func (recv *RecentInfo) GetShortName() string {
	retC := C.gtk_recent_info_get_short_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function gtk_recent_info_get_uri.
func (recv *RecentInfo) GetUri() string {
	retC := C.gtk_recent_info_get_uri((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUriDisplay is a wrapper around the C function gtk_recent_info_get_uri_display.
func (recv *RecentInfo) GetUriDisplay() string {
	retC := C.gtk_recent_info_get_uri_display((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_recent_info_get_visited : no return generator

// HasApplication is a wrapper around the C function gtk_recent_info_has_application.
func (recv *RecentInfo) HasApplication(appName string) bool {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	retC := C.gtk_recent_info_has_application((*C.GtkRecentInfo)(recv.native), c_app_name)
	retGo := retC == C.TRUE

	return retGo
}

// HasGroup is a wrapper around the C function gtk_recent_info_has_group.
func (recv *RecentInfo) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.gtk_recent_info_has_group((*C.GtkRecentInfo)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// IsLocal is a wrapper around the C function gtk_recent_info_is_local.
func (recv *RecentInfo) IsLocal() bool {
	retC := C.gtk_recent_info_is_local((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// LastApplication is a wrapper around the C function gtk_recent_info_last_application.
func (recv *RecentInfo) LastApplication() string {
	retC := C.gtk_recent_info_last_application((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Match is a wrapper around the C function gtk_recent_info_match.
func (recv *RecentInfo) Match(infoB *RecentInfo) bool {
	c_info_b := (*C.GtkRecentInfo)(infoB.ToC())

	retC := C.gtk_recent_info_match((*C.GtkRecentInfo)(recv.native), c_info_b)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function gtk_recent_info_ref.
func (recv *RecentInfo) Ref() *RecentInfo {
	retC := C.gtk_recent_info_ref((*C.GtkRecentInfo)(recv.native))
	retGo := RecentInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_recent_info_unref : no return generator

// RecentManagerClass is a wrapper around the C record GtkRecentManagerClass.
type RecentManagerClass struct {
	native *C.GtkRecentManagerClass
	// Private : parent_class
	// no type for changed
	// no type for _gtk_recent1
	// no type for _gtk_recent2
	// no type for _gtk_recent3
	// no type for _gtk_recent4
}

func RecentManagerClassNewFromC(u unsafe.Pointer) *RecentManagerClass {
	c := (*C.GtkRecentManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentManagerClass{native: c}

	return g
}

func (recv *RecentManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
