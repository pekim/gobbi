// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

/*

C type

GtkFileChooserNativeClass
*/
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

/*

C type

GtkNativeDialogClass
*/
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

// Struct defining a pad action entry.
/*

C type

GtkPadActionEntry
*/
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

/*

C type

GtkPadControllerClass
*/
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

/*

C type

GtkShortcutLabelClass
*/
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

/*

C type

GtkShortcutsGroupClass
*/
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

/*

C type

GtkShortcutsSectionClass
*/
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

/*

C type

GtkShortcutsShortcutClass
*/
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

/*

C type

GtkShortcutsWindowClass
*/
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

// Returns %TRUE if @tag is toggled on at exactly this point. If @tag
// is %NULL, returns %TRUE if any tag is toggled on at this point.
//
// Note that if gtk_text_iter_starts_tag() returns %TRUE, it means that @iter is
// at the beginning of the tagged range, and that the
// character at @iter is inside the tagged range. In other
// words, unlike gtk_text_iter_ends_tag(), if gtk_text_iter_starts_tag() returns
// %TRUE, gtk_text_iter_has_tag() will also return %TRUE for the same
// parameters.
/*

C function

gtk_text_iter_starts_tag
*/
func (recv *TextIter) StartsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_starts_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the name used by this class for matching in CSS code. See
// gtk_widget_class_set_css_name() for details.
/*

C function

gtk_widget_class_get_css_name
*/
func (recv *WidgetClass) GetCssName() string {
	retC := C.gtk_widget_class_get_css_name((*C.GtkWidgetClass)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the name to be used for CSS matching of widgets.
//
// If this function is not called for a given class, the name
// of the parent class is used.
/*

C function

gtk_widget_class_set_css_name
*/
func (recv *WidgetClass) SetCssName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_class_set_css_name((*C.GtkWidgetClass)(recv.native), c_name)

	return
}

// Returns the object name that is at position @pos in the widget
// hierarchy defined in @path.
/*

C function

gtk_widget_path_iter_get_object_name
*/
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_object_name((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := C.GoString(retC)

	return retGo
}

// Sets the object name for a given position in the widget hierarchy
// defined by @path.
//
// When set, the object name overrides the object type when matching
// CSS.
/*

C function

gtk_widget_path_iter_set_object_name
*/
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_object_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}
