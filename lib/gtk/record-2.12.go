// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"C"
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// BindingEntrySkip is a wrapper around the C function gtk_binding_entry_skip.
func BindingEntrySkip(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_skip(c_binding_set, c_keyval, c_modifiers)

	return
}

// PaperSizeGetPaperSizes is a wrapper around the C function gtk_paper_size_get_paper_sizes.
func PaperSizeGetPaperSizes(includeCustom bool) *glib.List {
	c_include_custom :=
		boolToGboolean(includeCustom)

	retC := C.gtk_paper_size_get_paper_sizes(c_include_custom)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToKeyFile is a wrapper around the C function gtk_paper_size_to_key_file.
func (recv *PaperSize) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_paper_size_to_key_file((*C.GtkPaperSize)(recv.native), c_key_file, c_group_name)

	return
}
