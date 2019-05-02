// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// Equals compares this FileChooserNativeClass with another FileChooserNativeClass, and returns true if they represent the same GObject.
func (recv *FileChooserNativeClass) Equals(other *FileChooserNativeClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this NativeDialogClass with another NativeDialogClass, and returns true if they represent the same GObject.
func (recv *NativeDialogClass) Equals(other *NativeDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PadActionEntry with another PadActionEntry, and returns true if they represent the same GObject.
func (recv *PadActionEntry) Equals(other *PadActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this PadControllerClass with another PadControllerClass, and returns true if they represent the same GObject.
func (recv *PadControllerClass) Equals(other *PadControllerClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ShortcutLabelClass with another ShortcutLabelClass, and returns true if they represent the same GObject.
func (recv *ShortcutLabelClass) Equals(other *ShortcutLabelClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ShortcutsGroupClass with another ShortcutsGroupClass, and returns true if they represent the same GObject.
func (recv *ShortcutsGroupClass) Equals(other *ShortcutsGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ShortcutsSectionClass with another ShortcutsSectionClass, and returns true if they represent the same GObject.
func (recv *ShortcutsSectionClass) Equals(other *ShortcutsSectionClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ShortcutsShortcutClass with another ShortcutsShortcutClass, and returns true if they represent the same GObject.
func (recv *ShortcutsShortcutClass) Equals(other *ShortcutsShortcutClass) bool {
	return other.ToC() == recv.ToC()
}

// Equals compares this ShortcutsWindowClass with another ShortcutsWindowClass, and returns true if they represent the same GObject.
func (recv *ShortcutsWindowClass) Equals(other *ShortcutsWindowClass) bool {
	return other.ToC() == recv.ToC()
}

// StartsTag is a wrapper around the C function gtk_text_iter_starts_tag.
func (recv *TextIter) StartsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_starts_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// GetCssName is a wrapper around the C function gtk_widget_class_get_css_name.
func (recv *WidgetClass) GetCssName() string {
	retC := C.gtk_widget_class_get_css_name((*C.GtkWidgetClass)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetCssName is a wrapper around the C function gtk_widget_class_set_css_name.
func (recv *WidgetClass) SetCssName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_class_set_css_name((*C.GtkWidgetClass)(recv.native), c_name)

	return
}

// IterGetObjectName is a wrapper around the C function gtk_widget_path_iter_get_object_name.
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_object_name((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := C.GoString(retC)

	return retGo
}

// IterSetObjectName is a wrapper around the C function gtk_widget_path_iter_set_object_name.
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_object_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}
