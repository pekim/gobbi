// This is a generated file - DO NOT EDIT
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
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

// BuilderNewFromFile is a wrapper around the C function gtk_builder_new_from_file.
func BuilderNewFromFile(filename string) *Builder {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_builder_new_from_file(c_filename)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BuilderNewFromResource is a wrapper around the C function gtk_builder_new_from_resource.
func BuilderNewFromResource(resourcePath string) *Builder {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_builder_new_from_resource(c_resource_path)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BuilderNewFromString is a wrapper around the C function gtk_builder_new_from_string.
func BuilderNewFromString(string_ string) *Builder {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gssize)(len(string_))

	retC := C.gtk_builder_new_from_string(c_string, c_length)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ButtonNewFromIconName is a wrapper around the C function gtk_button_new_from_icon_name.
func ButtonNewFromIconName(iconName string, size IconSize) *Button {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_button_new_from_icon_name(c_icon_name, c_size)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HeaderBarNew is a wrapper around the C function gtk_header_bar_new.
func HeaderBarNew() *HeaderBar {
	retC := C.gtk_header_bar_new()
	retGo := HeaderBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromSurface is a wrapper around the C function gtk_image_new_from_surface.
func ImageNewFromSurface(surface *cairo.Surface) *Image {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	retC := C.gtk_image_new_from_surface(c_surface)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListBoxNew is a wrapper around the C function gtk_list_box_new.
func ListBoxNew() *ListBox {
	retC := C.gtk_list_box_new()
	retGo := ListBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListBoxRowNew is a wrapper around the C function gtk_list_box_row_new.
func ListBoxRowNew() *ListBoxRow {
	retC := C.gtk_list_box_row_new()
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PlacesSidebarNew is a wrapper around the C function gtk_places_sidebar_new.
func PlacesSidebarNew() *PlacesSidebar {
	retC := C.gtk_places_sidebar_new()
	retGo := PlacesSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RevealerNew is a wrapper around the C function gtk_revealer_new.
func RevealerNew() *Revealer {
	retC := C.gtk_revealer_new()
	retGo := RevealerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SearchBarNew is a wrapper around the C function gtk_search_bar_new.
func SearchBarNew() *SearchBar {
	retC := C.gtk_search_bar_new()
	retGo := SearchBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StackNew is a wrapper around the C function gtk_stack_new.
func StackNew() *Stack {
	retC := C.gtk_stack_new()
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StackSwitcherNew is a wrapper around the C function gtk_stack_switcher_new.
func StackSwitcherNew() *StackSwitcher {
	retC := C.gtk_stack_switcher_new()
	retGo := StackSwitcherNewFromC(unsafe.Pointer(retC))

	return retGo
}
