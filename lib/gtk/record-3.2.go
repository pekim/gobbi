// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

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

// CssSection is a wrapper around the C record GtkCssSection.
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

// Equals compares this CssSection with another CssSection, and returns true if they represent the same GObject.
func (recv *CssSection) Equals(other *CssSection) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : gtk_css_section_get_end_line

// Blacklisted : gtk_css_section_get_end_position

// Blacklisted : gtk_css_section_get_file

// Blacklisted : gtk_css_section_get_parent

// Blacklisted : gtk_css_section_get_section_type

// Blacklisted : gtk_css_section_get_start_line

// Blacklisted : gtk_css_section_get_start_position

// Blacklisted : gtk_css_section_ref

// Blacklisted : gtk_css_section_unref

// Blacklisted : gtk_text_iter_assign

// Blacklisted : gtk_widget_class_set_accessible_role

// Blacklisted : gtk_widget_class_set_accessible_type

// Blacklisted : gtk_widget_path_append_for_widget

// Blacklisted : gtk_widget_path_append_with_siblings

// Blacklisted : gtk_widget_path_ref

// Blacklisted : gtk_widget_path_to_string

// Blacklisted : gtk_widget_path_unref
