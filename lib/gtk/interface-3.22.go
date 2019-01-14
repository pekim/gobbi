// This is a generated file - DO NOT EDIT
// +build gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// AddChoice is a wrapper around the C function gtk_file_chooser_add_choice.
func (recv *FileChooser) AddChoice(id string, label string, options []string, optionLabels []string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_options_array := make([]*C.char, len(options), len(options))
	for i, item := range options {
		g := options[i]
		c := C.CString(g)
		c_options_array[i] = c
	}
	c_options_arrayPtr := &c_options_array[0]
	c_options := (**C.char)(unsafe.Pointer(c_options_arrayPtr))

	c_option_labels_array := make([]*C.char, len(optionLabels), len(optionLabels))
	for i, item := range optionLabels {
		g := optionLabels[i]
		c := C.CString(g)
		c_option_labels_array[i] = c
	}
	c_option_labels_arrayPtr := &c_option_labels_array[0]
	c_option_labels := (**C.char)(unsafe.Pointer(c_option_labels_arrayPtr))

	C.gtk_file_chooser_add_choice((*C.GtkFileChooser)(recv.native), c_id, c_label, c_options, c_option_labels)

	return
}

// GetChoice is a wrapper around the C function gtk_file_chooser_get_choice.
func (recv *FileChooser) GetChoice(id string) string {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	retC := C.gtk_file_chooser_get_choice((*C.GtkFileChooser)(recv.native), c_id)
	retGo := C.GoString(retC)

	return retGo
}

// RemoveChoice is a wrapper around the C function gtk_file_chooser_remove_choice.
func (recv *FileChooser) RemoveChoice(id string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	C.gtk_file_chooser_remove_choice((*C.GtkFileChooser)(recv.native), c_id)

	return
}

// SetChoice is a wrapper around the C function gtk_file_chooser_set_choice.
func (recv *FileChooser) SetChoice(id string, option string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_option := C.CString(option)
	defer C.free(unsafe.Pointer(c_option))

	C.gtk_file_chooser_set_choice((*C.GtkFileChooser)(recv.native), c_id, c_option)

	return
}
