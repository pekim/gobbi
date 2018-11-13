// This is a generated file - DO NOT EDIT
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
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

// Unsupported : gtk_cell_area_class_find_cell_property : return type : Blacklisted record : GParamSpec

// Unsupported : gtk_cell_area_class_install_cell_property : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_cell_area_class_list_cell_properties : no return type

// Creates a new linear gradient along the line defined by (x0, y0) and (x1, y1). Before using the gradient
// a number of stop colors must be added through gtk_gradient_add_color_stop().
/*

C function : gtk_gradient_new_linear
*/
func GradientNewLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	retC := C.gtk_gradient_new_linear(c_x0, c_y0, c_x1, c_y1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new radial gradient along the two circles defined by (x0, y0, radius0) and
// (x1, y1, radius1). Before using the gradient a number of stop colors must be added
// through gtk_gradient_add_color_stop().
/*

C function : gtk_gradient_new_radial
*/
func GradientNewRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_radius0 := (C.gdouble)(radius0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	c_radius1 := (C.gdouble)(radius1)

	retC := C.gtk_gradient_new_radial(c_x0, c_y0, c_radius0, c_x1, c_y1, c_radius1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a stop color to @gradient.
/*

C function : gtk_gradient_add_color_stop
*/
func (recv *Gradient) AddColorStop(offset float64, color *SymbolicColor) {
	c_offset := (C.gdouble)(offset)

	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	C.gtk_gradient_add_color_stop((*C.GtkGradient)(recv.native), c_offset, c_color)

	return
}

// Increases the reference count of @gradient.
/*

C function : gtk_gradient_ref
*/
func (recv *Gradient) Ref() *Gradient {
	retC := C.gtk_gradient_ref((*C.GtkGradient)(recv.native))
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @gradient is resolvable, @resolved_gradient will be filled in
// with the resolved gradient as a cairo_pattern_t, and %TRUE will
// be returned. Generally, if @gradient can’t be resolved, it is
// due to it being defined on top of a named color that doesn't
// exist in @props.
/*

C function : gtk_gradient_resolve
*/
func (recv *Gradient) Resolve(props *StyleProperties) (bool, *cairo.Pattern) {
	c_props := (*C.GtkStyleProperties)(C.NULL)
	if props != nil {
		c_props = (*C.GtkStyleProperties)(props.ToC())
	}

	var c_resolved_gradient *C.cairo_pattern_t

	retC := C.gtk_gradient_resolve((*C.GtkGradient)(recv.native), c_props, &c_resolved_gradient)
	retGo := retC == C.TRUE

	resolvedGradient := cairo.PatternNewFromC(unsafe.Pointer(c_resolved_gradient))

	return retGo, resolvedGradient
}

// Decreases the reference count of @gradient, freeing its memory
// if the reference count reaches 0.
/*

C function : gtk_gradient_unref
*/
func (recv *Gradient) Unref() {
	C.gtk_gradient_unref((*C.GtkGradient)(recv.native))

	return
}

// Renders an icon using gtk_render_icon_pixbuf(). In most cases,
// gtk_widget_render_icon_pixbuf() is better, since it automatically provides
// most of the arguments from the current widget settings.  This
// function never returns %NULL; if the icon can’t be rendered
// (perhaps because an image file fails to load), a default "missing
// image" icon will be returned instead.
/*

C function : gtk_icon_set_render_icon_pixbuf
*/
func (recv *IconSet) RenderIconPixbuf(context *StyleContext, size IconSize) *gdkpixbuf.Pixbuf {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_icon_set_render_icon_pixbuf((*C.GtkIconSet)(recv.native), c_context, c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Allocates a new #GtkRequisition-struct and initializes its elements to zero.
/*

C function : gtk_requisition_new
*/
func RequisitionNew() *Requisition {
	retC := C.gtk_requisition_new()
	retGo := RequisitionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_selection_data_get_data_with_length : no return type

// Creates a symbolic color by modifying the relative alpha
// value of @color. A factor < 1.0 would resolve to a more
// transparent color, while > 1.0 would resolve to a more
// opaque color.
/*

C function : gtk_symbolic_color_new_alpha
*/
func SymbolicColorNewAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_alpha(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a symbolic color pointing to a literal color.
/*

C function : gtk_symbolic_color_new_literal
*/
func SymbolicColorNewLiteral(color *gdk.RGBA) *SymbolicColor {
	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	retC := C.gtk_symbolic_color_new_literal(c_color)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a symbolic color defined as a mix of another
// two colors. a mix factor of 0 would resolve to @color1,
// while a factor of 1 would resolve to @color2.
/*

C function : gtk_symbolic_color_new_mix
*/
func SymbolicColorNewMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	c_color1 := (*C.GtkSymbolicColor)(C.NULL)
	if color1 != nil {
		c_color1 = (*C.GtkSymbolicColor)(color1.ToC())
	}

	c_color2 := (*C.GtkSymbolicColor)(C.NULL)
	if color2 != nil {
		c_color2 = (*C.GtkSymbolicColor)(color2.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_mix(c_color1, c_color2, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a symbolic color pointing to an unresolved named
// color. See gtk_style_context_lookup_color() and
// gtk_style_properties_lookup_color().
/*

C function : gtk_symbolic_color_new_name
*/
func SymbolicColorNewName(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_symbolic_color_new_name(c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a symbolic color defined as a shade of
// another color. A factor > 1.0 would resolve to
// a brighter color, while < 1.0 would resolve to
// a darker color.
/*

C function : gtk_symbolic_color_new_shade
*/
func SymbolicColorNewShade(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_shade(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Increases the reference count of @color
/*

C function : gtk_symbolic_color_ref
*/
func (recv *SymbolicColor) Ref() *SymbolicColor {
	retC := C.gtk_symbolic_color_ref((*C.GtkSymbolicColor)(recv.native))
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If @color is resolvable, @resolved_color will be filled in
// with the resolved color, and %TRUE will be returned. Generally,
// if @color can’t be resolved, it is due to it being defined on
// top of a named color that doesn’t exist in @props.
//
// When @props is %NULL, resolving of named colors will fail, so if
// your @color is or references such a color, this function will
// return %FALSE.
/*

C function : gtk_symbolic_color_resolve
*/
func (recv *SymbolicColor) Resolve(props *StyleProperties) (bool, *gdk.RGBA) {
	c_props := (*C.GtkStyleProperties)(C.NULL)
	if props != nil {
		c_props = (*C.GtkStyleProperties)(props.ToC())
	}

	var c_resolved_color C.GdkRGBA

	retC := C.gtk_symbolic_color_resolve((*C.GtkSymbolicColor)(recv.native), c_props, &c_resolved_color)
	retGo := retC == C.TRUE

	resolvedColor := gdk.RGBANewFromC(unsafe.Pointer(&c_resolved_color))

	return retGo, resolvedColor
}

// Decreases the reference count of @color, freeing its memory if the
// reference count reaches 0.
/*

C function : gtk_symbolic_color_unref
*/
func (recv *SymbolicColor) Unref() {
	C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(recv.native))

	return
}

// Unsupported : gtk_tree_path_get_indices_with_depth : no return type

// Returns an empty widget path.
/*

C function : gtk_widget_path_new
*/
func WidgetPathNew() *WidgetPath {
	retC := C.gtk_widget_path_new()
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends a widget type to the widget hierarchy represented by @path.
/*

C function : gtk_widget_path_append_type
*/
func (recv *WidgetPath) AppendType(type_ gobject.Type) int32 {
	c_type := (C.GType)(type_)

	retC := C.gtk_widget_path_append_type((*C.GtkWidgetPath)(recv.native), c_type)
	retGo := (int32)(retC)

	return retGo
}

// Returns a copy of @path
/*

C function : gtk_widget_path_copy
*/
func (recv *WidgetPath) Copy() *WidgetPath {
	retC := C.gtk_widget_path_copy((*C.GtkWidgetPath)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count on @path, freeing the structure
// if the reference count reaches 0.
/*

C function : gtk_widget_path_free
*/
func (recv *WidgetPath) Free() {
	C.gtk_widget_path_free((*C.GtkWidgetPath)(recv.native))

	return
}

// Returns the topmost object type, that is, the object type this path
// is representing.
/*

C function : gtk_widget_path_get_object_type
*/
func (recv *WidgetPath) GetObjectType() gobject.Type {
	retC := C.gtk_widget_path_get_object_type((*C.GtkWidgetPath)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// Returns %TRUE if any of the parents of the widget represented
// in @path is of type @type, or any subtype of it.
/*

C function : gtk_widget_path_has_parent
*/
func (recv *WidgetPath) HasParent(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.gtk_widget_path_has_parent((*C.GtkWidgetPath)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the widget type represented by this path
// is @type, or a subtype of it.
/*

C function : gtk_widget_path_is_type
*/
func (recv *WidgetPath) IsType(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.gtk_widget_path_is_type((*C.GtkWidgetPath)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Adds the class @name to the widget at position @pos in
// the hierarchy defined in @path. See
// gtk_style_context_add_class().
/*

C function : gtk_widget_path_iter_add_class
*/
func (recv *WidgetPath) IterAddClass(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_add_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// Adds the region @name to the widget at position @pos in
// the hierarchy defined in @path. See
// gtk_style_context_add_region().
//
// Region names must only contain lowercase letters
// and “-”, starting always with a lowercase letter.
/*

C function : gtk_widget_path_iter_add_region
*/
func (recv *WidgetPath) IterAddRegion(pos int32, name string, flags RegionFlags) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GtkRegionFlags)(flags)

	C.gtk_widget_path_iter_add_region((*C.GtkWidgetPath)(recv.native), c_pos, c_name, c_flags)

	return
}

// Removes all classes from the widget at position @pos in the
// hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_clear_classes
*/
func (recv *WidgetPath) IterClearClasses(pos int32) {
	c_pos := (C.gint)(pos)

	C.gtk_widget_path_iter_clear_classes((*C.GtkWidgetPath)(recv.native), c_pos)

	return
}

// Removes all regions from the widget at position @pos in the
// hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_clear_regions
*/
func (recv *WidgetPath) IterClearRegions(pos int32) {
	c_pos := (C.gint)(pos)

	C.gtk_widget_path_iter_clear_regions((*C.GtkWidgetPath)(recv.native), c_pos)

	return
}

// Returns the object #GType that is at position @pos in the widget
// hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_get_object_type
*/
func (recv *WidgetPath) IterGetObjectType(pos int32) gobject.Type {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_object_type((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := (gobject.Type)(retC)

	return retGo
}

// Returns %TRUE if the widget at position @pos has the class @name
// defined, %FALSE otherwise.
/*

C function : gtk_widget_path_iter_has_class
*/
func (recv *WidgetPath) IterHasClass(pos int32, name string) bool {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_widget_path_iter_has_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the widget at position @pos has the name @name,
// %FALSE otherwise.
/*

C function : gtk_widget_path_iter_has_name
*/
func (recv *WidgetPath) IterHasName(pos int32, name string) bool {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_widget_path_iter_has_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)
	retGo := retC == C.TRUE

	return retGo
}

// See gtk_widget_path_iter_has_class(). This is a version that operates
// with GQuarks.
/*

C function : gtk_widget_path_iter_has_qclass
*/
func (recv *WidgetPath) IterHasQclass(pos int32, qname glib.Quark) bool {
	c_pos := (C.gint)(pos)

	c_qname := (C.GQuark)(qname)

	retC := C.gtk_widget_path_iter_has_qclass((*C.GtkWidgetPath)(recv.native), c_pos, c_qname)
	retGo := retC == C.TRUE

	return retGo
}

// See gtk_widget_path_iter_has_name(). This is a version
// that operates on #GQuarks.
/*

C function : gtk_widget_path_iter_has_qname
*/
func (recv *WidgetPath) IterHasQname(pos int32, qname glib.Quark) bool {
	c_pos := (C.gint)(pos)

	c_qname := (C.GQuark)(qname)

	retC := C.gtk_widget_path_iter_has_qname((*C.GtkWidgetPath)(recv.native), c_pos, c_qname)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_path_iter_has_qregion : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Returns a list with all the class names defined for the widget
// at position @pos in the hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_list_classes
*/
func (recv *WidgetPath) IterListClasses(pos int32) *glib.SList {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_list_classes((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns a list with all the region names defined for the widget
// at position @pos in the hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_list_regions
*/
func (recv *WidgetPath) IterListRegions(pos int32) *glib.SList {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_list_regions((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes the class @name from the widget at position @pos in
// the hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_remove_class
*/
func (recv *WidgetPath) IterRemoveClass(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_remove_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// Removes the region @name from the widget at position @pos in
// the hierarchy defined in @path.
/*

C function : gtk_widget_path_iter_remove_region
*/
func (recv *WidgetPath) IterRemoveRegion(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_remove_region((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// Sets the widget name for the widget found at position @pos
// in the widget hierarchy defined by @path.
/*

C function : gtk_widget_path_iter_set_name
*/
func (recv *WidgetPath) IterSetName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// Sets the object type for a given position in the widget hierarchy
// defined by @path.
/*

C function : gtk_widget_path_iter_set_object_type
*/
func (recv *WidgetPath) IterSetObjectType(pos int32, type_ gobject.Type) {
	c_pos := (C.gint)(pos)

	c_type := (C.GType)(type_)

	C.gtk_widget_path_iter_set_object_type((*C.GtkWidgetPath)(recv.native), c_pos, c_type)

	return
}

// Returns the number of #GtkWidget #GTypes between the represented
// widget and its topmost container.
/*

C function : gtk_widget_path_length
*/
func (recv *WidgetPath) Length() int32 {
	retC := C.gtk_widget_path_length((*C.GtkWidgetPath)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Prepends a widget type to the widget hierachy represented by @path.
/*

C function : gtk_widget_path_prepend_type
*/
func (recv *WidgetPath) PrependType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.gtk_widget_path_prepend_type((*C.GtkWidgetPath)(recv.native), c_type)

	return
}
