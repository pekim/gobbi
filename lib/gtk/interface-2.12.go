// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AddChild is a wrapper around the C function gtk_buildable_add_child.
func (recv *Buildable) AddChild(builder *Builder, child *gobject.Object, type_ string) {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_child := (*C.GObject)(child.ToC())

	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	C.gtk_buildable_add_child((*C.GtkBuildable)(recv.native), c_builder, c_child, c_type)

	return
}

// ConstructChild is a wrapper around the C function gtk_buildable_construct_child.
func (recv *Buildable) ConstructChild(builder *Builder, name string) *gobject.Object {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_buildable_construct_child((*C.GtkBuildable)(recv.native), c_builder, c_name)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CustomFinished is a wrapper around the C function gtk_buildable_custom_finished.
func (recv *Buildable) CustomFinished(builder *Builder, child *gobject.Object, tagname string, data uintptr) {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_child := (*C.GObject)(child.ToC())

	c_tagname := C.CString(tagname)
	defer C.free(unsafe.Pointer(c_tagname))

	c_data := (C.gpointer)(data)

	C.gtk_buildable_custom_finished((*C.GtkBuildable)(recv.native), c_builder, c_child, c_tagname, c_data)

	return
}

// CustomTagEnd is a wrapper around the C function gtk_buildable_custom_tag_end.
func (recv *Buildable) CustomTagEnd(builder *Builder, child *gobject.Object, tagname string, data uintptr) {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_child := (*C.GObject)(child.ToC())

	c_tagname := C.CString(tagname)
	defer C.free(unsafe.Pointer(c_tagname))

	c_data := (C.gpointer)(data)

	C.gtk_buildable_custom_tag_end((*C.GtkBuildable)(recv.native), c_builder, c_child, c_tagname, &c_data)

	return
}

// CustomTagStart is a wrapper around the C function gtk_buildable_custom_tag_start.
func (recv *Buildable) CustomTagStart(builder *Builder, child *gobject.Object, tagname string) (bool, *glib.MarkupParser, uintptr) {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_child := (*C.GObject)(child.ToC())

	c_tagname := C.CString(tagname)
	defer C.free(unsafe.Pointer(c_tagname))

	var c_parser C.GMarkupParser

	var c_data C.gpointer

	retC := C.gtk_buildable_custom_tag_start((*C.GtkBuildable)(recv.native), c_builder, c_child, c_tagname, &c_parser, &c_data)
	retGo := retC == C.TRUE

	parser := glib.MarkupParserNewFromC(unsafe.Pointer(&c_parser))

	data := (uintptr)(unsafe.Pointer(&c_data))

	return retGo, parser, data
}

// GetInternalChild is a wrapper around the C function gtk_buildable_get_internal_child.
func (recv *Buildable) GetInternalChild(builder *Builder, childname string) *gobject.Object {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_childname := C.CString(childname)
	defer C.free(unsafe.Pointer(c_childname))

	retC := C.gtk_buildable_get_internal_child((*C.GtkBuildable)(recv.native), c_builder, c_childname)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function gtk_buildable_get_name.
func (recv *Buildable) GetName() string {
	retC := C.gtk_buildable_get_name((*C.GtkBuildable)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// ParserFinished is a wrapper around the C function gtk_buildable_parser_finished.
func (recv *Buildable) ParserFinished(builder *Builder) {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	C.gtk_buildable_parser_finished((*C.GtkBuildable)(recv.native), c_builder)

	return
}

// SetBuildableProperty is a wrapper around the C function gtk_buildable_set_buildable_property.
func (recv *Buildable) SetBuildableProperty(builder *Builder, name string, value *gobject.Value) {
	c_builder := (*C.GtkBuilder)(builder.ToC())

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_buildable_set_buildable_property((*C.GtkBuildable)(recv.native), c_builder, c_name, c_value)

	return
}

// SetName is a wrapper around the C function gtk_buildable_set_name.
func (recv *Buildable) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_buildable_set_name((*C.GtkBuildable)(recv.native), c_name)

	return
}

// GetCells is a wrapper around the C function gtk_cell_layout_get_cells.
func (recv *CellLayout) GetCells() *glib.List {
	retC := C.gtk_cell_layout_get_cells((*C.GtkCellLayout)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}
