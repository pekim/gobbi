// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// RecentManager is a wrapper around the C record GtkRecentManager.
type RecentManager struct {
	native *C.GtkRecentManager
	// Private : parent_instance
	// Private : priv
}

func RecentManagerNewFromC(u unsafe.Pointer) *RecentManager {
	c := (*C.GtkRecentManager)(u)
	if c == nil {
		return nil
	}

	g := &RecentManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RecentManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RecentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : STOCK_ORIENTATION_LANDSCAPE

// Blacklisted : STOCK_ORIENTATION_PORTRAIT

// Blacklisted : STOCK_ORIENTATION_REVERSE_LANDSCAPE

// Blacklisted : STOCK_ORIENTATION_REVERSE_PORTRAIT

// Blacklisted : STOCK_SELECT_ALL

// Unsupported : gtk_paper_size_get_default : return type :

// Unsupported : gtk_print_error_quark : return type :

// Unsupported : gtk_print_run_page_setup_dialog : return type :

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc (GtkPageSetupDoneFunc) for param done_cb

// Unsupported : gtk_target_table_free : unsupported parameter targets :

// Unsupported : gtk_target_table_new_from_list : array return type :

// Unsupported : gtk_targets_include_image : unsupported parameter targets :

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_uri : unsupported parameter targets :

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
