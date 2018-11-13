// This is a generated file - DO NOT EDIT
// +build gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// Gets the contents of the selection data as a #GdkPixbuf.
/*

C function

gtk_selection_data_get_pixbuf
*/
func (recv *SelectionData) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_selection_data_get_pixbuf((*C.GtkSelectionData)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_selection_data_get_uris : no return type

// Sets the contents of the selection from a #GdkPixbuf
// The pixbuf is converted to the form determined by
// @selection_data->target.
/*

C function

gtk_selection_data_set_pixbuf
*/
func (recv *SelectionData) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) bool {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_selection_data_set_pixbuf((*C.GtkSelectionData)(recv.native), c_pixbuf)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_selection_data_set_uris : unsupported parameter uris :

// Given a #GtkSelectionData object holding a list of targets,
// determines if any of the targets in @targets can be used to
// provide a #GdkPixbuf.
/*

C function

gtk_selection_data_targets_include_image
*/
func (recv *SelectionData) TargetsIncludeImage(writable bool) bool {
	c_writable :=
		boolToGboolean(writable)

	retC := C.gtk_selection_data_targets_include_image((*C.GtkSelectionData)(recv.native), c_writable)
	retGo := retC == C.TRUE

	return retGo
}

// Appends the image targets supported by #GtkSelectionData to
// the target list. All targets are added with the same @info.
/*

C function

gtk_target_list_add_image_targets
*/
func (recv *TargetList) AddImageTargets(info uint32, writable bool) {
	c_info := (C.guint)(info)

	c_writable :=
		boolToGboolean(writable)

	C.gtk_target_list_add_image_targets((*C.GtkTargetList)(recv.native), c_info, c_writable)

	return
}

// Appends the text targets supported by #GtkSelectionData to
// the target list. All targets are added with the same @info.
/*

C function

gtk_target_list_add_text_targets
*/
func (recv *TargetList) AddTextTargets(info uint32) {
	c_info := (C.guint)(info)

	C.gtk_target_list_add_text_targets((*C.GtkTargetList)(recv.native), c_info)

	return
}

// Appends the URI targets supported by #GtkSelectionData to
// the target list. All targets are added with the same @info.
/*

C function

gtk_target_list_add_uri_targets
*/
func (recv *TargetList) AddUriTargets(info uint32) {
	c_info := (C.guint)(info)

	C.gtk_target_list_add_uri_targets((*C.GtkTargetList)(recv.native), c_info)

	return
}
