// Code generated - DO NOT EDIT.
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static void c_gtk_cell_area_add_with_properties(GtkCellArea* area, GtkCellRenderer* renderer, const gchar* first_prop_name) {
    return gtk_cell_area_add_with_properties(area, renderer, first_prop_name, NULL);
}
*/
/*

static void c_gtk_cell_area_cell_get(GtkCellArea* area, GtkCellRenderer* renderer, const gchar* first_prop_name) {
    return gtk_cell_area_cell_get(area, renderer, first_prop_name, NULL);
}
*/
/*

static void c_gtk_cell_area_cell_get_valist(GtkCellArea* area, GtkCellRenderer* renderer, const gchar* first_property_name) {
    return gtk_cell_area_cell_get_valist(area, renderer, first_property_name, NULL);
}
*/
/*

static void c_gtk_cell_area_cell_set(GtkCellArea* area, GtkCellRenderer* renderer, const gchar* first_prop_name) {
    return gtk_cell_area_cell_set(area, renderer, first_prop_name, NULL);
}
*/
/*

static void c_gtk_cell_area_cell_set_valist(GtkCellArea* area, GtkCellRenderer* renderer, const gchar* first_property_name) {
    return gtk_cell_area_cell_set_valist(area, renderer, first_property_name, NULL);
}
*/
/*

static void c_gtk_style_context_get(GtkStyleContext* context, GtkStateFlags state) {
    return gtk_style_context_get(context, state, NULL);
}
*/
/*

static void c_gtk_style_context_get_style(GtkStyleContext* context) {
    return gtk_style_context_get_style(context, NULL);
}
*/
/*

static void c_gtk_style_context_get_style_valist(GtkStyleContext* context) {
    return gtk_style_context_get_style_valist(context, NULL);
}
*/
/*

static void c_gtk_style_context_get_valist(GtkStyleContext* context, GtkStateFlags state) {
    return gtk_style_context_get_valist(context, state, NULL);
}
*/
/*

static void c_gtk_style_properties_get(GtkStyleProperties* props, GtkStateFlags state) {
    return gtk_style_properties_get(props, state, NULL);
}
*/
/*

static void c_gtk_style_properties_get_valist(GtkStyleProperties* props, GtkStateFlags state) {
    return gtk_style_properties_get_valist(props, state, NULL);
}
*/
/*

static void c_gtk_style_properties_set(GtkStyleProperties* props, GtkStateFlags state) {
    return gtk_style_properties_set(props, state, NULL);
}
*/
/*

static void c_gtk_style_properties_set_valist(GtkStyleProperties* props, GtkStateFlags state) {
    return gtk_style_properties_set_valist(props, state, NULL);
}
*/
/*

static void c_gtk_theming_engine_get(GtkThemingEngine* engine, GtkStateFlags state) {
    return gtk_theming_engine_get(engine, state, NULL);
}
*/
/*

static void c_gtk_theming_engine_get_style(GtkThemingEngine* engine) {
    return gtk_theming_engine_get_style(engine, NULL);
}
*/
/*

static void c_gtk_theming_engine_get_style_valist(GtkThemingEngine* engine) {
    return gtk_theming_engine_get_style_valist(engine, NULL);
}
*/
/*

static void c_gtk_theming_engine_get_valist(GtkThemingEngine* engine, GtkStateFlags state) {
    return gtk_theming_engine_get_valist(engine, state, NULL);
}
*/
import "C"

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_binding_entry_add_signal_from_string(param0 unsafe.Pointer, param1 string) int {
	cValue0 := (*C.GtkBindingSet)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_binding_entry_add_signal_from_string(cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_cell_area_class_find_cell_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkCellAreaClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_cell_area_class_find_cell_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_class_install_cell_property(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAreaClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	C.gtk_cell_area_class_install_cell_property(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_class_list_cell_properties(paramInstance unsafe.Pointer, param0 *uint) []unsafe.Pointer {
	cValueInstance := (*C.GtkCellAreaClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_class_list_cell_properties(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]unsafe.Pointer, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_gradient_new_linear(param0 float64, param1 float64, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gtk_gradient_new_linear(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_new_radial(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	ret := C.gtk_gradient_new_radial(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_add_color_stop(paramInstance unsafe.Pointer, param0 float64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (*C.GtkSymbolicColor)(unsafe.Pointer(param1))

	C.gtk_gradient_add_color_stop(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_gradient_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gradient_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_resolve(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProperties)(unsafe.Pointer(param0))

	cValue1 := (**C.cairo_pattern_t)(unsafe.Pointer(param1))

	ret := C.gtk_gradient_resolve(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gradient_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	C.gtk_gradient_unref(cValueInstance)
}

func Fn_gtk_icon_set_render_icon_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_icon_set_render_icon_pixbuf(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_requisition_new() unsafe.Pointer {
	ret := C.gtk_requisition_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

func Fn_gtk_selection_data_get_data_with_length(paramInstance unsafe.Pointer, param0 *int) []uint8 {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	ret := C.gtk_selection_data_get_data_with_length(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

func Fn_gtk_symbolic_color_new_alpha(param0 unsafe.Pointer, param1 float64) unsafe.Pointer {
	cValue0 := (*C.GtkSymbolicColor)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_symbolic_color_new_alpha(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_literal(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	ret := C.gtk_symbolic_color_new_literal(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_mix(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64) unsafe.Pointer {
	cValue0 := (*C.GtkSymbolicColor)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkSymbolicColor)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	ret := C.gtk_symbolic_color_new_mix(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_name(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_symbolic_color_new_name(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_shade(param0 unsafe.Pointer, param1 float64) unsafe.Pointer {
	cValue0 := (*C.GtkSymbolicColor)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_symbolic_color_new_shade(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	ret := C.gtk_symbolic_color_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_resolve(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProperties)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	ret := C.gtk_symbolic_color_resolve(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_symbolic_color_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	C.gtk_symbolic_color_unref(cValueInstance)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

func Fn_gtk_tree_path_get_indices_with_depth(paramInstance unsafe.Pointer, param0 *int) []int {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	ret := C.gtk_tree_path_get_indices_with_depth(cValueInstance, cValue0)

	retLen := int(*cValue0)
	retGo := make([]int, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](int))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_widget_path_new() unsafe.Pointer {
	ret := C.gtk_widget_path_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_append_type(paramInstance unsafe.Pointer, param0 uint64) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_path_append_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_path_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	C.gtk_widget_path_free(cValueInstance)
}

func Fn_gtk_widget_path_get_object_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_get_object_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_gtk_widget_path_has_parent(paramInstance unsafe.Pointer, param0 uint64) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_path_has_parent(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_is_type(paramInstance unsafe.Pointer, param0 uint64) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_path_is_type(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_add_class(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_add_class(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_add_region(paramInstance unsafe.Pointer, param0 int, param1 string, param2 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GtkRegionFlags)(param2)

	C.gtk_widget_path_iter_add_region(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_path_iter_clear_classes(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_path_iter_clear_classes(cValueInstance, cValue0)
}

func Fn_gtk_widget_path_iter_clear_regions(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_path_iter_clear_regions(cValueInstance, cValue0)
}

func Fn_gtk_widget_path_iter_get_object_type(paramInstance unsafe.Pointer, param0 int) uint64 {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_object_type(cValueInstance, cValue0)

	return (uint64)(ret)
}

func Fn_gtk_widget_path_iter_has_class(paramInstance unsafe.Pointer, param0 int, param1 string) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_widget_path_iter_has_class(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_name(paramInstance unsafe.Pointer, param0 int, param1 string) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_widget_path_iter_has_name(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_qclass(paramInstance unsafe.Pointer, param0 int, param1 uint32) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GQuark)(param1)

	ret := C.gtk_widget_path_iter_has_qclass(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_qname(paramInstance unsafe.Pointer, param0 int, param1 uint32) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GQuark)(param1)

	ret := C.gtk_widget_path_iter_has_qname(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_qregion(paramInstance unsafe.Pointer, param0 int, param1 uint32, param2 *int) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GQuark)(param1)

	cValue2 := (*C.GtkRegionFlags)(unsafe.Pointer(param2))

	ret := C.gtk_widget_path_iter_has_qregion(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_region(paramInstance unsafe.Pointer, param0 int, param1 string, param2 *int) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GtkRegionFlags)(unsafe.Pointer(param2))

	ret := C.gtk_widget_path_iter_has_region(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_list_classes(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_list_classes(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_iter_list_regions(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_list_regions(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_iter_remove_class(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_remove_class(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_remove_region(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_remove_region(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_set_name(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_set_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_set_object_type(paramInstance unsafe.Pointer, param0 int, param1 uint64) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GType)(param1)

	C.gtk_widget_path_iter_set_object_type(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_path_prepend_type(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.gtk_widget_path_prepend_type(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

func Fn_gtk_cairo_should_draw_window(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	ret := C.gtk_cairo_should_draw_window(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cairo_transform_to_window(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkWindow)(unsafe.Pointer(param2))

	C.gtk_cairo_transform_to_window(cValue0, cValue1, cValue2)
}

func Fn_gtk_device_grab_add(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_device_grab_add(cValue0, cValue1, cValue2)
}

func Fn_gtk_device_grab_remove(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))

	C.gtk_device_grab_remove(cValue0, cValue1)
}

func Fn_gtk_draw_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool, param4 int, param5 bool) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	cValue4 := (C.GtkTextDirection)(param4)

	cValue5 := toCBool(param5)

	C.gtk_draw_insertion_cursor(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_get_binary_age() uint {
	ret := C.gtk_get_binary_age()

	return (uint)(ret)
}

func Fn_gtk_get_interface_age() uint {
	ret := C.gtk_get_interface_age()

	return (uint)(ret)
}

func Fn_gtk_get_major_version() uint {
	ret := C.gtk_get_major_version()

	return (uint)(ret)
}

func Fn_gtk_get_micro_version() uint {
	ret := C.gtk_get_micro_version()

	return (uint)(ret)
}

func Fn_gtk_get_minor_version() uint {
	ret := C.gtk_get_minor_version()

	return (uint)(ret)
}

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

func Fn_gtk_render_activity(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_activity(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_arrow(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_background(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_background(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_check(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_expander(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GtkPositionType)(param6)

	C.gtk_render_extension(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gtk_render_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_focus(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_frame(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_frame(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_frame_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int, param7 float64, param8 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GtkPositionType)(param6)

	cValue7 := (C.gdouble)(param7)

	cValue8 := (C.gdouble)(param8)

	C.gtk_render_frame_gap(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8)
}

func Fn_gtk_render_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_handle(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) unsafe.Pointer {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconSource)(unsafe.Pointer(param1))

	cValue2 := (C.GtkIconSize)(param2)

	ret := C.gtk_render_icon_pixbuf(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_render_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (*C.PangoLayout)(unsafe.Pointer(param4))

	C.gtk_render_layout(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_render_line(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_line(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_option(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GtkOrientation)(param6)

	C.gtk_render_slider(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_add_credit_section : parameter 'people' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

func Fn_gtk_about_dialog_get_license_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_license_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

func Fn_gtk_about_dialog_set_license_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkLicense)(param0)

	C.gtk_about_dialog_set_license_type(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

func Fn_gtk_app_chooser_button_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_app_chooser_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_button_append_custom_item(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GIcon)(unsafe.Pointer(param2))

	C.gtk_app_chooser_button_append_custom_item(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_app_chooser_button_append_separator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	C.gtk_app_chooser_button_append_separator(cValueInstance)
}

func Fn_gtk_app_chooser_button_get_show_dialog_item(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_button_get_show_dialog_item(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_button_set_active_custom_item(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_button_set_active_custom_item(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_button_set_show_dialog_item(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_button_set_show_dialog_item(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_dialog_new(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDialogFlags)(param1)

	cValue2 := (*C.GFile)(unsafe.Pointer(param2))

	ret := C.gtk_app_chooser_dialog_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_dialog_new_for_content_type(param0 unsafe.Pointer, param1 int, param2 string) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDialogFlags)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.gtk_app_chooser_dialog_new_for_content_type(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_dialog_get_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAppChooserDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_dialog_get_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_widget_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_app_chooser_widget_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_widget_get_default_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_default_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_widget_get_show_all(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_all(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_fallback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_fallback(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_other(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_other(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_recommended(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_recommended(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_set_show_all(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_all(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_default(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_fallback(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_fallback(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_other(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_other(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_recommended(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_recommended(cValueInstance, cValue0)
}

func Fn_gtk_application_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GApplicationFlags)(param1)

	ret := C.gtk_application_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_add_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_application_add_window(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

func Fn_gtk_application_get_windows(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_get_windows(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

func Fn_gtk_application_remove_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_application_remove_window(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_application_set_accels_for_action : parameter 'accels' is array parameter without length parameter

func Fn_gtk_assistant_next_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_next_page(cValueInstance)
}

func Fn_gtk_assistant_previous_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_previous_page(cValueInstance)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

func Fn_gtk_box_new(param0 int, param1 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_box_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_button_box_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_button_box_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_calendar_get_day_is_marked(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_calendar_get_day_is_marked(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

func Fn_gtk_cell_area_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 bool) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (C.GtkCellRendererState)(param3)

	cValue4 := toCBool(param4)

	ret := C.gtk_cell_area_activate(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_activate_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkEvent)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (C.GtkCellRendererState)(param4)

	ret := C.gtk_cell_area_activate_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_area_add(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_add_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	C.gtk_cell_area_add_focus_sibling(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_add_with_properties(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_cell_area_add_with_properties(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_apply_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_cell_area_apply_attributes(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_attribute_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_cell_area_attribute_connect(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_attribute_disconnect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_cell_area_attribute_disconnect(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_cell_get(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_cell_area_cell_get(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_cell_get_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_cell_area_cell_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_cell_get_valist(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_cell_area_cell_get_valist(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_cell_set(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_cell_area_cell_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_cell_set_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_cell_area_cell_set_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_cell_set_valist(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.c_gtk_cell_area_cell_set_valist(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_copy_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_copy_context(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_create_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_create_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) int {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkEvent)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (C.GtkCellRendererState)(param4)

	ret := C.gtk_cell_area_event(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return (int)(ret)
}

func Fn_gtk_cell_area_focus(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	ret := C.gtk_cell_area_focus(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

func Fn_gtk_cell_area_get_cell_allocation(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkCellRenderer)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	C.gtk_cell_area_get_cell_allocation(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_area_get_cell_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (*C.GdkRectangle)(unsafe.Pointer(param5))

	ret := C.gtk_cell_area_get_cell_at_position(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_current_path_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_current_path_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_cell_area_get_edit_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_edit_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_edited_cell(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_edited_cell(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_focus_cell(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_focus_cell(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_focus_from_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_get_focus_from_sibling(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_focus_siblings(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_get_focus_siblings(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_preferred_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_area_get_preferred_height(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_cell_area_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_area_get_preferred_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_area_get_preferred_width(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_cell_area_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_area_get_request_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_request_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_cell_area_has_renderer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_has_renderer(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_inner_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_cell_area_inner_cell_area(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_is_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_is_activatable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_is_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	ret := C.gtk_cell_area_is_focus_sibling(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_area_remove(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_remove_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	C.gtk_cell_area_remove_focus_sibling(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_render(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int, param6 bool) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.cairo_t)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	cValue5 := (C.GtkCellRendererState)(param5)

	cValue6 := toCBool(param6)

	C.gtk_cell_area_render(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gtk_cell_area_request_renderer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 int, param4 *int, param5 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (C.GtkOrientation)(param1)

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	C.gtk_cell_area_request_renderer(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_cell_area_set_focus_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_area_set_focus_cell(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_stop_editing(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_area_stop_editing(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_box_new() unsafe.Pointer {
	ret := C.gtk_cell_area_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_box_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_box_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_cell_area_box_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_cell_area_box_pack_end(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_box_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_cell_area_box_pack_start(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_box_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_cell_area_box_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_context_get_allocation(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_area_context_get_allocation(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_context_get_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_context_get_preferred_height(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_area_context_get_preferred_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_area_context_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_context_get_preferred_width(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_area_context_get_preferred_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_area_context_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_context_push_preferred_height(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_area_context_push_preferred_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_push_preferred_width(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_area_context_push_preferred_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_get_aligned_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkCellRendererState)(param1)

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	C.gtk_cell_renderer_get_aligned_area(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_renderer_get_preferred_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_renderer_get_preferred_height(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_renderer_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_renderer_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_renderer_get_preferred_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkRequisition)(unsafe.Pointer(param2))

	C.gtk_cell_renderer_get_preferred_size(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_renderer_get_preferred_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_renderer_get_preferred_width(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_renderer_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_renderer_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_renderer_get_request_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_get_request_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_cell_renderer_get_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) int {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkCellRendererState)(param1)

	ret := C.gtk_cell_renderer_get_state(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_cell_renderer_is_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_is_activatable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_get_draw_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_draw_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_get_fit_model(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_fit_model(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_set_background_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_background_rgba(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_draw_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_view_set_draw_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_fit_model(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_view_set_fit_model(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_color_button_new_with_rgba(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	ret := C.gtk_color_button_new_with_rgba(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_get_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_button_get_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_button_set_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_get_current_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_current_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_get_previous_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_previous_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_current_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_set_current_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_previous_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_set_previous_rgba(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

func Fn_gtk_combo_box_get_active_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_active_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_combo_box_get_id_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_id_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_popup_fixed_width(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_popup_fixed_width(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

func Fn_gtk_combo_box_popup_for_device(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	C.gtk_combo_box_popup_for_device(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_active_id(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_combo_box_set_active_id(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_set_id_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_id_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_popup_fixed_width(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_popup_fixed_width(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_combo_box_text_insert(paramInstance unsafe.Pointer, param0 int, param1 string, param2 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_combo_box_text_insert(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_combo_box_text_remove_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_text_remove_all(cValueInstance)
}

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_entry_get_icon_area(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gtk_entry_get_icon_area(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_get_text_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_entry_get_text_area(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_entry_completion_new_with_area(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

func Fn_gtk_icon_info_load_symbolic(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *bool, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRGBA)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRGBA)(unsafe.Pointer(param3))

	cValue4 := (*C.gboolean)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_symbolic(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

func Fn_gtk_icon_info_load_symbolic_for_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.gboolean)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_symbolic_for_context(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

func Fn_gtk_icon_info_load_symbolic_for_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *bool, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (C.GtkStateType)(param1)

	cValue2 := (*C.gboolean)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_symbolic_for_style(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

func Fn_gtk_icon_view_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_new_with_area(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_item_get_reserve_indicator(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_reserve_indicator(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_item_set_reserve_indicator(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_item_set_reserve_indicator(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_get_parent_shell(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_shell_get_parent_shell(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_shell_get_selected_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_shell_get_selected_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_get_background_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_background_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_get_background_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_background_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_numerable_icon_get_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_count(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_numerable_icon_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_numerable_icon_get_style_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_style_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_set_background_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_numerable_icon_set_background_gicon(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_background_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_numerable_icon_set_background_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_count(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_numerable_icon_set_count(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_numerable_icon_set_label(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_style_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	C.gtk_numerable_icon_set_style_context(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.gtk_numerable_icon_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_new_with_style_context(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStyleContext)(unsafe.Pointer(param1))

	ret := C.gtk_numerable_icon_new_with_style_context(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paned_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_paned_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_progress_bar_get_show_text(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_show_text(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_progress_bar_set_show_text(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_progress_bar_set_show_text(cValueInstance, cValue0)
}

func Fn_gtk_radio_action_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRadioAction)(unsafe.Pointer(param0))

	C.gtk_radio_action_join_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_button_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	C.gtk_radio_button_join_group(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_scale_new(param0 int, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_scale_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_new_with_range(param0 int, param1 float64, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gtk_scale_new_with_range(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

func Fn_gtk_scrollbar_new(param0 int, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_scrollbar_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_get_min_content_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_min_content_height(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_min_content_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_min_content_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_set_min_content_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scrolled_window_set_min_content_height(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_min_content_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scrolled_window_set_min_content_width(cValueInstance, cValue0)
}

func Fn_gtk_separator_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_separator_new(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_style_has_context(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_has_context(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_style_context_add_class(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_add_class(cValueInstance, cValue0)
}

func Fn_gtk_style_context_add_provider(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProvider)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	C.gtk_style_context_add_provider(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_add_region(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkRegionFlags)(param1)

	C.gtk_style_context_add_region(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_cancel_animations(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_style_context_cancel_animations(cValueInstance, cValue0)
}

func Fn_gtk_style_context_get(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_style_context_get(cValueInstance, cValue0)
}

func Fn_gtk_style_context_get_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_get_background_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_border(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_style_context_get_border(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_border_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_get_border_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_get_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_style_context_get_font(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	ret := C.gtk_style_context_get_font(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_junction_sides(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_junction_sides(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_style_context_get_margin(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_style_context_get_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_padding(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_style_context_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_path(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_path(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_context_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_style_context_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_state(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_style_context_get_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.c_gtk_style_context_get_style(cValueInstance)
}

func Fn_gtk_style_context_get_style_valist(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.c_gtk_style_context_get_style_valist(cValueInstance)
}

func Fn_gtk_style_context_get_valist(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_style_context_get_valist(cValueInstance, cValue0)
}

func Fn_gtk_style_context_has_class(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_context_has_class(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_style_context_has_region(paramInstance unsafe.Pointer, param0 string, param1 *int) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRegionFlags)(unsafe.Pointer(param1))

	ret := C.gtk_style_context_has_region(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_style_context_invalidate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_invalidate(cValueInstance)
}

func Fn_gtk_style_context_list_classes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_list_classes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_list_regions(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_list_regions(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_notify_state_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 bool) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := toCBool(param3)

	C.gtk_style_context_notify_state_change(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_style_context_pop_animatable_region(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_pop_animatable_region(cValueInstance)
}

func Fn_gtk_style_context_push_animatable_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_style_context_push_animatable_region(cValueInstance, cValue0)
}

func Fn_gtk_style_context_remove_class(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_remove_class(cValueInstance, cValue0)
}

func Fn_gtk_style_context_remove_provider(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProvider)(unsafe.Pointer(param0))

	C.gtk_style_context_remove_provider(cValueInstance, cValue0)
}

func Fn_gtk_style_context_remove_region(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_remove_region(cValueInstance, cValue0)
}

func Fn_gtk_style_context_restore(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_restore(cValueInstance)
}

func Fn_gtk_style_context_save(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_save(cValueInstance)
}

func Fn_gtk_style_context_scroll_animations(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_style_context_scroll_animations(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_style_context_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_style_context_set_background(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextDirection)(param0)

	C.gtk_style_context_set_direction(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_junction_sides(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkJunctionSides)(param0)

	C.gtk_style_context_set_junction_sides(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	C.gtk_style_context_set_path(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_style_context_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_state(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.gtk_style_context_set_state(cValueInstance, cValue0)
}

func Fn_gtk_style_context_state_is_running(paramInstance unsafe.Pointer, param0 int, param1 *float64) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_style_context_state_is_running(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_style_context_add_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStyleProvider)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	C.gtk_style_context_add_provider_for_screen(cValue0, cValue1, cValue2)
}

func Fn_gtk_style_context_remove_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStyleProvider)(unsafe.Pointer(param1))

	C.gtk_style_context_remove_provider_for_screen(cValue0, cValue1)
}

func Fn_gtk_style_context_reset_widgets(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_style_context_reset_widgets(cValue0)
}

func Fn_gtk_style_properties_get(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_style_properties_get(cValueInstance, cValue0)
}

func Fn_gtk_style_properties_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_style_properties_get_property(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_style_properties_get_valist(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_style_properties_get_valist(cValueInstance, cValue0)
}

func Fn_gtk_style_properties_lookup_color(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_properties_lookup_color(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_properties_map_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkSymbolicColor)(unsafe.Pointer(param1))

	C.gtk_style_properties_map_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_properties_merge(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProperties)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_style_properties_merge(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_properties_set(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_style_properties_set(cValueInstance, cValue0)
}

func Fn_gtk_style_properties_set_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_properties_set_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_style_properties_set_valist(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_style_properties_set_valist(cValueInstance, cValue0)
}

func Fn_gtk_style_properties_unset_property(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	C.gtk_style_properties_unset_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_switch_new() unsafe.Pointer {
	ret := C.gtk_switch_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_switch_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSwitch)(unsafe.Pointer(paramInstance))

	ret := C.gtk_switch_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_switch_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSwitch)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_switch_set_active(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_get_cursor_locations(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_text_view_get_cursor_locations(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_theming_engine_get(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_theming_engine_get(cValueInstance, cValue0)
}

func Fn_gtk_theming_engine_get_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_background_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_border(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_border(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_border_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_border_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_theming_engine_get_font(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	ret := C.gtk_theming_engine_get_font(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_theming_engine_get_junction_sides(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_junction_sides(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_theming_engine_get_margin(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_padding(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_path(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_path(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_theming_engine_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_theming_engine_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_theming_engine_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_state(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_theming_engine_get_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	C.c_gtk_theming_engine_get_style(cValueInstance)
}

func Fn_gtk_theming_engine_get_style_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_style_property(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_style_valist(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	C.c_gtk_theming_engine_get_style_valist(cValueInstance)
}

func Fn_gtk_theming_engine_get_valist(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.c_gtk_theming_engine_get_valist(cValueInstance, cValue0)
}

func Fn_gtk_theming_engine_has_class(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_theming_engine_has_class(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_has_region(paramInstance unsafe.Pointer, param0 string, param1 *int) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRegionFlags)(unsafe.Pointer(param1))

	ret := C.gtk_theming_engine_has_region(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	ret := C.gtk_theming_engine_lookup_color(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_state_is_running(paramInstance unsafe.Pointer, param0 int, param1 *float64) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_theming_engine_state_is_running(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

func Fn_gtk_tree_view_is_blank_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkTreeViewColumn)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	ret := C.gtk_tree_view_is_blank_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

func Fn_gtk_tree_view_column_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_column_new_with_area(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_get_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_button(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_widget_add_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (C.GdkEventMask)(param1)

	C.gtk_widget_add_device_events(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_device_is_shadowed(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_widget_device_is_shadowed(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_draw(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	C.gtk_widget_draw(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_device_enabled(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_widget_get_device_enabled(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_widget_get_device_events(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_bottom(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_bottom(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_left(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_left(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_right(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_right(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_top(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_top(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_preferred_height(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_preferred_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_widget_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_get_preferred_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	C.gtk_widget_get_preferred_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_width(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_preferred_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_widget_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_get_request_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_request_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_state_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_state_flags(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_input_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_input_shape_combine_region(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_override_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_background_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_override_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_override_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_override_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.gtk_widget_override_font(cValueInstance, cValue0)
}

func Fn_gtk_widget_override_symbolic_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_symbolic_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_queue_draw_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_queue_draw_region(cValueInstance, cValue0)
}

func Fn_gtk_widget_render_icon_pixbuf(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_widget_render_icon_pixbuf(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_reset_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_reset_style(cValueInstance)
}

func Fn_gtk_widget_set_device_enabled(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_widget_set_device_enabled(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (C.GdkEventMask)(param1)

	C.gtk_widget_set_device_events(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_margin_bottom(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_bottom(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_left(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_left(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_right(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_right(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_top(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_top(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_state_flags(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := toCBool(param1)

	C.gtk_widget_set_state_flags(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_support_multidevice(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_support_multidevice(cValueInstance, cValue0)
}

func Fn_gtk_widget_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_shape_combine_region(cValueInstance, cValue0)
}

func Fn_gtk_widget_unset_state_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.gtk_widget_unset_state_flags(cValueInstance, cValue0)
}

func Fn_gtk_window_get_application(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_application(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_has_resize_grip(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_has_resize_grip(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_resize_grip_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_window_get_resize_grip_area(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_window_resize_grip_is_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_resize_grip_is_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_resize_to_geometry(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_resize_to_geometry(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_set_application(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkApplication)(unsafe.Pointer(param0))

	C.gtk_window_set_application(cValueInstance, cValue0)
}

func Fn_gtk_window_set_default_geometry(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_set_default_geometry(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_set_has_resize_grip(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_has_resize_grip(cValueInstance, cValue0)
}

func Fn_gtk_window_set_has_user_ref_count(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_has_user_ref_count(cValueInstance, cValue0)
}

func Fn_gtk_window_group_get_current_device_grab(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_window_group_get_current_device_grab(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_get_app_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAppChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_get_app_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_get_content_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAppChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_get_content_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_refresh(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooser)(unsafe.Pointer(paramInstance))

	C.gtk_app_chooser_refresh(cValueInstance)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

func Fn_gtk_cell_layout_get_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_layout_get_area(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_scrollable_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrollable_get_hscroll_policy(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_hscroll_policy(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrollable_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrollable_get_vscroll_policy(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_vscroll_policy(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrollable_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrollable_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_set_hscroll_policy(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkScrollablePolicy)(param0)

	C.gtk_scrollable_set_hscroll_policy(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrollable_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_set_vscroll_policy(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkScrollablePolicy)(param0)

	C.gtk_scrollable_set_vscroll_policy(cValueInstance, cValue0)
}

func Fn_gtk_style_provider_get_icon_factory(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	ret := C.gtk_style_provider_get_icon_factory(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_provider_get_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	ret := C.gtk_style_provider_get_style(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_provider_get_style_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyleProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GParamSpec)(unsafe.Pointer(param2))

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	ret := C.gtk_style_provider_get_style_property(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

func Fn_gtk_tree_model_iter_previous(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_iter_previous(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback
