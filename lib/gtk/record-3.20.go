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

// FileChooserNativeClass is a wrapper around the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native *C.GtkFileChooserNativeClass
	// parent_class : record
}

func FileChooserNativeClassNewFromC(u unsafe.Pointer) *FileChooserNativeClass {
	c := (*C.GtkFileChooserNativeClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserNativeClass{native: c}

	return g
}

func (recv *FileChooserNativeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileChooserNativeClass with another FileChooserNativeClass, and returns true if they represent the same GObject.
func (recv *FileChooserNativeClass) Equals(other *FileChooserNativeClass) bool {
	return other.ToC() == recv.ToC()
}

// NativeDialogClass is a wrapper around the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native *C.GtkNativeDialogClass
	// parent_class : record
	// no type for response
	// no type for show
	// no type for hide
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func NativeDialogClassNewFromC(u unsafe.Pointer) *NativeDialogClass {
	c := (*C.GtkNativeDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &NativeDialogClass{native: c}

	return g
}

func (recv *NativeDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NativeDialogClass with another NativeDialogClass, and returns true if they represent the same GObject.
func (recv *NativeDialogClass) Equals(other *NativeDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// PadActionEntry is a wrapper around the C record GtkPadActionEntry.
type PadActionEntry struct {
	native     *C.GtkPadActionEntry
	Type       PadActionType
	Index      int32
	Mode       int32
	Label      string
	ActionName string
}

func PadActionEntryNewFromC(u unsafe.Pointer) *PadActionEntry {
	c := (*C.GtkPadActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &PadActionEntry{
		ActionName: C.GoString(c.action_name),
		Index:      (int32)(c.index),
		Label:      C.GoString(c.label),
		Mode:       (int32)(c.mode),
		Type:       (PadActionType)(c._type),
		native:     c,
	}

	return g
}

func (recv *PadActionEntry) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GtkPadActionType)(recv.Type)
	recv.native.index =
		(C.gint)(recv.Index)
	recv.native.mode =
		(C.gint)(recv.Mode)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.action_name =
		C.CString(recv.ActionName)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PadActionEntry with another PadActionEntry, and returns true if they represent the same GObject.
func (recv *PadActionEntry) Equals(other *PadActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// PadControllerClass is a wrapper around the C record GtkPadControllerClass.
type PadControllerClass struct {
	native *C.GtkPadControllerClass
}

func PadControllerClassNewFromC(u unsafe.Pointer) *PadControllerClass {
	c := (*C.GtkPadControllerClass)(u)
	if c == nil {
		return nil
	}

	g := &PadControllerClass{native: c}

	return g
}

func (recv *PadControllerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PadControllerClass with another PadControllerClass, and returns true if they represent the same GObject.
func (recv *PadControllerClass) Equals(other *PadControllerClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutLabelClass is a wrapper around the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native *C.GtkShortcutLabelClass
}

func ShortcutLabelClassNewFromC(u unsafe.Pointer) *ShortcutLabelClass {
	c := (*C.GtkShortcutLabelClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutLabelClass{native: c}

	return g
}

func (recv *ShortcutLabelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutLabelClass with another ShortcutLabelClass, and returns true if they represent the same GObject.
func (recv *ShortcutLabelClass) Equals(other *ShortcutLabelClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsGroupClass is a wrapper around the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native *C.GtkShortcutsGroupClass
}

func ShortcutsGroupClassNewFromC(u unsafe.Pointer) *ShortcutsGroupClass {
	c := (*C.GtkShortcutsGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsGroupClass{native: c}

	return g
}

func (recv *ShortcutsGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsGroupClass with another ShortcutsGroupClass, and returns true if they represent the same GObject.
func (recv *ShortcutsGroupClass) Equals(other *ShortcutsGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsSectionClass is a wrapper around the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native *C.GtkShortcutsSectionClass
}

func ShortcutsSectionClassNewFromC(u unsafe.Pointer) *ShortcutsSectionClass {
	c := (*C.GtkShortcutsSectionClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsSectionClass{native: c}

	return g
}

func (recv *ShortcutsSectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsSectionClass with another ShortcutsSectionClass, and returns true if they represent the same GObject.
func (recv *ShortcutsSectionClass) Equals(other *ShortcutsSectionClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsShortcutClass is a wrapper around the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native *C.GtkShortcutsShortcutClass
}

func ShortcutsShortcutClassNewFromC(u unsafe.Pointer) *ShortcutsShortcutClass {
	c := (*C.GtkShortcutsShortcutClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsShortcutClass{native: c}

	return g
}

func (recv *ShortcutsShortcutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsShortcutClass with another ShortcutsShortcutClass, and returns true if they represent the same GObject.
func (recv *ShortcutsShortcutClass) Equals(other *ShortcutsShortcutClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsWindowClass is a wrapper around the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native *C.GtkShortcutsWindowClass
	// parent_class : record
	// no type for close
	// no type for search
}

func ShortcutsWindowClassNewFromC(u unsafe.Pointer) *ShortcutsWindowClass {
	c := (*C.GtkShortcutsWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsWindowClass{native: c}

	return g
}

func (recv *ShortcutsWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsWindowClass with another ShortcutsWindowClass, and returns true if they represent the same GObject.
func (recv *ShortcutsWindowClass) Equals(other *ShortcutsWindowClass) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : gtk_text_iter_starts_tag

// Blacklisted : gtk_widget_class_get_css_name

// Blacklisted : gtk_widget_class_set_css_name

// Blacklisted : gtk_widget_path_iter_get_object_name

// Blacklisted : gtk_widget_path_iter_set_object_name
