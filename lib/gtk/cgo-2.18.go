// This is a generated file - DO NOT EDIT
// +build gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// EntryNewWithBuffer is a wrapper around the C function gtk_entry_new_with_buffer.
func EntryNewWithBuffer(buffer *EntryBuffer) *Entry {
	c_buffer := (*C.GtkEntryBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkEntryBuffer)(buffer.ToC())
	}

	retC := C.gtk_entry_new_with_buffer(c_buffer)
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EntryBufferNew is a wrapper around the C function gtk_entry_buffer_new.
func EntryBufferNew(initialChars string, nInitialChars int32) *EntryBuffer {
	c_initial_chars := C.CString(initialChars)
	defer C.free(unsafe.Pointer(c_initial_chars))

	c_n_initial_chars := (C.gint)(nInitialChars)

	retC := C.gtk_entry_buffer_new(c_initial_chars, c_n_initial_chars)
	retGo := EntryBufferNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InfoBarNew is a wrapper around the C function gtk_info_bar_new.
func InfoBarNew() *InfoBar {
	retC := C.gtk_info_bar_new()
	retGo := InfoBarNewFromC(unsafe.Pointer(retC))

	return retGo
}
