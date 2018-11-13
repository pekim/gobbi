// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Defines a part of a CSS document. Because sections are nested into
// one another, you can use gtk_css_section_get_parent() to get the
// containing region.
/*

C record/class : GtkCssSection
*/
type CssSection struct {
	native *C.GtkCssSection
}

func CssSectionNewFromC(u unsafe.Pointer) *CssSection {
	c := (*C.GtkCssSection)(u)
	if c == nil {
		return nil
	}

	g := &CssSection{native: c}

	return g
}

func (recv *CssSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the line in the CSS document where this section end.
// The line number is 0-indexed, so the first line of the document
// will return 0.
// This value may change in future invocations of this function if
// @section is not yet parsed completely. This will for example
// happen in the GtkCssProvider::parsing-error signal.
// The end position and line may be identical to the start
// position and line for sections which failed to parse anything
// successfully.
/*

C function : gtk_css_section_get_end_line
*/
func (recv *CssSection) GetEndLine() uint32 {
	retC := C.gtk_css_section_get_end_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_end_line().
// This value may change in future invocations of this function if
// @section is not yet parsed completely. This will for example
// happen in the GtkCssProvider::parsing-error signal.
// The end position and line may be identical to the start
// position and line for sections which failed to parse anything
// successfully.
/*

C function : gtk_css_section_get_end_position
*/
func (recv *CssSection) GetEndPosition() uint32 {
	retC := C.gtk_css_section_get_end_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the file that @section was parsed from. If no such file exists,
// for example because the CSS was loaded via
// @gtk_css_provider_load_from_data(), then %NULL is returned.
/*

C function : gtk_css_section_get_file
*/
func (recv *CssSection) GetFile() *gio.File {
	retC := C.gtk_css_section_get_file((*C.GtkCssSection)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the parent section for the given @section. The parent section is
// the section that contains this @section. A special case are sections of
// type #GTK_CSS_SECTION_DOCUMENT. Their parent will either be %NULL
// if they are the original CSS document that was loaded by
// gtk_css_provider_load_from_file() or a section of type
// #GTK_CSS_SECTION_IMPORT if it was loaded with an import rule from
// a different file.
/*

C function : gtk_css_section_get_parent
*/
func (recv *CssSection) GetParent() *CssSection {
	retC := C.gtk_css_section_get_parent((*C.GtkCssSection)(recv.native))
	var retGo (*CssSection)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CssSectionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the type of information that @section describes.
/*

C function : gtk_css_section_get_section_type
*/
func (recv *CssSection) GetSectionType() CssSectionType {
	retC := C.gtk_css_section_get_section_type((*C.GtkCssSection)(recv.native))
	retGo := (CssSectionType)(retC)

	return retGo
}

// Returns the line in the CSS document where this section starts.
// The line number is 0-indexed, so the first line of the document
// will return 0.
/*

C function : gtk_css_section_get_start_line
*/
func (recv *CssSection) GetStartLine() uint32 {
	retC := C.gtk_css_section_get_start_line((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_start_line().
/*

C function : gtk_css_section_get_start_position
*/
func (recv *CssSection) GetStartPosition() uint32 {
	retC := C.gtk_css_section_get_start_position((*C.GtkCssSection)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Increments the reference count on @section.
/*

C function : gtk_css_section_ref
*/
func (recv *CssSection) Ref() *CssSection {
	retC := C.gtk_css_section_ref((*C.GtkCssSection)(recv.native))
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count on @section, freeing the
// structure if the reference count reaches 0.
/*

C function : gtk_css_section_unref
*/
func (recv *CssSection) Unref() {
	C.gtk_css_section_unref((*C.GtkCssSection)(recv.native))

	return
}

// Assigns the value of @other to @iter.  This function
// is not useful in applications, because iterators can be assigned
// with `GtkTextIter i = j;`. The
// function is used by language bindings.
/*

C function : gtk_text_iter_assign
*/
func (recv *TextIter) Assign(other *TextIter) {
	c_other := (*C.GtkTextIter)(C.NULL)
	if other != nil {
		c_other = (*C.GtkTextIter)(other.ToC())
	}

	C.gtk_text_iter_assign((*C.GtkTextIter)(recv.native), c_other)

	return
}

// Sets the default #AtkRole to be set on accessibles created for
// widgets of @widget_class. Accessibles may decide to not honor this
// setting if their role reporting is more refined. Calls to
// gtk_widget_class_set_accessible_type() will reset this value.
//
// In cases where you want more fine-grained control over the role of
// accessibles created for @widget_class, you should provide your own
// accessible type and use gtk_widget_class_set_accessible_type()
// instead.
//
// If @role is #ATK_ROLE_INVALID, the default role will not be changed
// and the accessible’s default role will be used instead.
//
// This function should only be called from class init functions of widgets.
/*

C function : gtk_widget_class_set_accessible_role
*/
func (recv *WidgetClass) SetAccessibleRole(role atk.Role) {
	c_role := (C.AtkRole)(role)

	C.gtk_widget_class_set_accessible_role((*C.GtkWidgetClass)(recv.native), c_role)

	return
}

// Sets the type to be used for creating accessibles for widgets of
// @widget_class. The given @type must be a subtype of the type used for
// accessibles of the parent class.
//
// This function should only be called from class init functions of widgets.
/*

C function : gtk_widget_class_set_accessible_type
*/
func (recv *WidgetClass) SetAccessibleType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.gtk_widget_class_set_accessible_type((*C.GtkWidgetClass)(recv.native), c_type)

	return
}

// Appends the data from @widget to the widget hierarchy represented
// by @path. This function is a shortcut for adding information from
// @widget to the given @path. This includes setting the name or
// adding the style classes from @widget.
/*

C function : gtk_widget_path_append_for_widget
*/
func (recv *WidgetPath) AppendForWidget(widget *Widget) int32 {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_widget_path_append_for_widget((*C.GtkWidgetPath)(recv.native), c_widget)
	retGo := (int32)(retC)

	return retGo
}

// Appends a widget type with all its siblings to the widget hierarchy
// represented by @path. Using this function instead of
// gtk_widget_path_append_type() will allow the CSS theming to use
// sibling matches in selectors and apply :nth-child() pseudo classes.
// In turn, it requires a lot more care in widget implementations as
// widgets need to make sure to call gtk_widget_reset_style() on all
// involved widgets when the @siblings path changes.
/*

C function : gtk_widget_path_append_with_siblings
*/
func (recv *WidgetPath) AppendWithSiblings(siblings *WidgetPath, siblingIndex uint32) int32 {
	c_siblings := (*C.GtkWidgetPath)(C.NULL)
	if siblings != nil {
		c_siblings = (*C.GtkWidgetPath)(siblings.ToC())
	}

	c_sibling_index := (C.guint)(siblingIndex)

	retC := C.gtk_widget_path_append_with_siblings((*C.GtkWidgetPath)(recv.native), c_siblings, c_sibling_index)
	retGo := (int32)(retC)

	return retGo
}

// Increments the reference count on @path.
/*

C function : gtk_widget_path_ref
*/
func (recv *WidgetPath) Ref() *WidgetPath {
	retC := C.gtk_widget_path_ref((*C.GtkWidgetPath)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Dumps the widget path into a string representation. It tries to match
// the CSS style as closely as possible (Note that there might be paths
// that cannot be represented in CSS).
//
// The main use of this code is for debugging purposes, so that you can
// g_print() the path or dump it in a gdb session.
/*

C function : gtk_widget_path_to_string
*/
func (recv *WidgetPath) ToString() string {
	retC := C.gtk_widget_path_to_string((*C.GtkWidgetPath)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count on @path, freeing the structure
// if the reference count reaches 0.
/*

C function : gtk_widget_path_unref
*/
func (recv *WidgetPath) Unref() {
	C.gtk_widget_path_unref((*C.GtkWidgetPath)(recv.native))

	return
}
