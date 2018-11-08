// This is a generated file - DO NOT EDIT

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// AccelGroupsActivate is a wrapper around the C function gtk_accel_groups_activate.
func AccelGroupsActivate(object *gobject.Object, accelKey uint32, accelMods gdk.ModifierType) bool {
	c_object := (*C.GObject)(object.ToC())

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_accel_groups_activate(c_object, c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// AccelGroupsFromObject is a wrapper around the C function gtk_accel_groups_from_object.
func AccelGroupsFromObject(object *gobject.Object) *glib.SList {
	c_object := (*C.GObject)(object.ToC())

	retC := C.gtk_accel_groups_from_object(c_object)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AcceleratorGetDefaultModMask is a wrapper around the C function gtk_accelerator_get_default_mod_mask.
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	retC := C.gtk_accelerator_get_default_mod_mask()
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// AcceleratorName is a wrapper around the C function gtk_accelerator_name.
func AcceleratorName(acceleratorKey uint32, acceleratorMods gdk.ModifierType) string {
	c_accelerator_key := (C.guint)(acceleratorKey)

	c_accelerator_mods := (C.GdkModifierType)(acceleratorMods)

	retC := C.gtk_accelerator_name(c_accelerator_key, c_accelerator_mods)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : output array param accelerator_codes

// AcceleratorSetDefaultModMask is a wrapper around the C function gtk_accelerator_set_default_mod_mask.
func AcceleratorSetDefaultModMask(defaultModMask gdk.ModifierType) {
	c_default_mod_mask := (C.GdkModifierType)(defaultModMask)

	C.gtk_accelerator_set_default_mod_mask(c_default_mod_mask)

	return
}

// AcceleratorValid is a wrapper around the C function gtk_accelerator_valid.
func AcceleratorValid(keyval uint32, modifiers gdk.ModifierType) bool {
	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_accelerator_valid(c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// BindingEntryAddSignall is a wrapper around the C function gtk_binding_entry_add_signall.
func BindingEntryAddSignall(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType, signalName string, bindingArgs *glib.SList) {
	c_binding_set := (*C.GtkBindingSet)(bindingSet.ToC())

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_signal_name := C.CString(signalName)
	defer C.free(unsafe.Pointer(c_signal_name))

	c_binding_args := (*C.GSList)(bindingArgs.ToC())

	C.gtk_binding_entry_add_signall(c_binding_set, c_keyval, c_modifiers, c_signal_name, c_binding_args)

	return
}

// BindingEntryRemove is a wrapper around the C function gtk_binding_entry_remove.
func BindingEntryRemove(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(bindingSet.ToC())

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_remove(c_binding_set, c_keyval, c_modifiers)

	return
}

// BindingSetByClass is a wrapper around the C function gtk_binding_set_by_class.
func BindingSetByClass(objectClass uintptr) *BindingSet {
	c_object_class := (C.gpointer)(objectClass)

	retC := C.gtk_binding_set_by_class(c_object_class)
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BindingSetFind is a wrapper around the C function gtk_binding_set_find.
func BindingSetFind(setName string) *BindingSet {
	c_set_name := C.CString(setName)
	defer C.free(unsafe.Pointer(c_set_name))

	retC := C.gtk_binding_set_find(c_set_name)
	var retGo (*BindingSet)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BindingSetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// BindingSetNew is a wrapper around the C function gtk_binding_set_new.
func BindingSetNew(setName string) *BindingSet {
	c_set_name := C.CString(setName)
	defer C.free(unsafe.Pointer(c_set_name))

	retC := C.gtk_binding_set_new(c_set_name)
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BindingsActivate is a wrapper around the C function gtk_bindings_activate.
func BindingsActivate(object *gobject.Object, keyval uint32, modifiers gdk.ModifierType) bool {
	c_object := (*C.GObject)(object.ToC())

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_bindings_activate(c_object, c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// BuilderErrorQuark is a wrapper around the C function gtk_builder_error_quark.
func BuilderErrorQuark() glib.Quark {
	retC := C.gtk_builder_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// CheckVersion is a wrapper around the C function gtk_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	retC := C.gtk_check_version(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// CssProviderErrorQuark is a wrapper around the C function gtk_css_provider_error_quark.
func CssProviderErrorQuark() glib.Quark {
	retC := C.gtk_css_provider_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// DisableSetlocale is a wrapper around the C function gtk_disable_setlocale.
func DisableSetlocale() {
	C.gtk_disable_setlocale()

	return
}

// DistributeNaturalAllocation is a wrapper around the C function gtk_distribute_natural_allocation.
func DistributeNaturalAllocation(extraSpace int32, nRequestedSizes uint32, sizes *RequestedSize) int32 {
	c_extra_space := (C.gint)(extraSpace)

	c_n_requested_sizes := (C.guint)(nRequestedSizes)

	c_sizes := (*C.GtkRequestedSize)(sizes.ToC())

	retC := C.gtk_distribute_natural_allocation(c_extra_space, c_n_requested_sizes, c_sizes)
	retGo := (int32)(retC)

	return retGo
}

// DragFinish is a wrapper around the C function gtk_drag_finish.
func DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	c_context := (*C.GdkDragContext)(context.ToC())

	c_success :=
		boolToGboolean(success)

	c_del :=
		boolToGboolean(del)

	c_time_ := (C.guint32)(time)

	C.gtk_drag_finish(c_context, c_success, c_del, c_time_)

	return
}

// DragGetSourceWidget is a wrapper around the C function gtk_drag_get_source_widget.
func DragGetSourceWidget(context *gdk.DragContext) *Widget {
	c_context := (*C.GdkDragContext)(context.ToC())

	retC := C.gtk_drag_get_source_widget(c_context)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// DragSetIconDefault is a wrapper around the C function gtk_drag_set_icon_default.
func DragSetIconDefault(context *gdk.DragContext) {
	c_context := (*C.GdkDragContext)(context.ToC())

	C.gtk_drag_set_icon_default(c_context)

	return
}

// Unsupported : gtk_drag_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// DragSetIconPixbuf is a wrapper around the C function gtk_drag_set_icon_pixbuf.
func DragSetIconPixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(context.ToC())

	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_pixbuf(c_context, c_pixbuf, c_hot_x, c_hot_y)

	return
}

// DragSetIconStock is a wrapper around the C function gtk_drag_set_icon_stock.
func DragSetIconStock(context *gdk.DragContext, stockId string, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(context.ToC())

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_stock(c_context, c_stock_id, c_hot_x, c_hot_y)

	return
}

// DragSetIconSurface is a wrapper around the C function gtk_drag_set_icon_surface.
func DragSetIconSurface(context *gdk.DragContext, surface *cairo.Surface) {
	c_context := (*C.GdkDragContext)(context.ToC())

	c_surface := (*C.cairo_surface_t)(surface.ToC())

	C.gtk_drag_set_icon_surface(c_context, c_surface)

	return
}

// DragSetIconWidget is a wrapper around the C function gtk_drag_set_icon_widget.
func DragSetIconWidget(context *gdk.DragContext, widget *Widget, hotX int32, hotY int32) {
	c_context := (*C.GdkDragContext)(context.ToC())

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_hot_x := (C.gint)(hotX)

	c_hot_y := (C.gint)(hotY)

	C.gtk_drag_set_icon_widget(c_context, c_widget, c_hot_x, c_hot_y)

	return
}

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// EventsPending is a wrapper around the C function gtk_events_pending.
func EventsPending() bool {
	retC := C.gtk_events_pending()
	retGo := retC == C.TRUE

	return retGo
}

// False is a wrapper around the C function gtk_false.
func False() bool {
	retC := C.gtk_false()
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_get_current_event : no return generator

// GetCurrentEventDevice is a wrapper around the C function gtk_get_current_event_device.
func GetCurrentEventDevice() *gdk.Device {
	retC := C.gtk_get_current_event_device()
	var retGo (*gdk.Device)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.DeviceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// GetCurrentEventTime is a wrapper around the C function gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {
	retC := C.gtk_get_current_event_time()
	retGo := (uint32)(retC)

	return retGo
}

// GetDebugFlags is a wrapper around the C function gtk_get_debug_flags.
func GetDebugFlags() uint32 {
	retC := C.gtk_get_debug_flags()
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultLanguage is a wrapper around the C function gtk_get_default_language.
func GetDefaultLanguage() *pango.Language {
	retC := C.gtk_get_default_language()
	retGo := pango.LanguageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// GrabGetCurrent is a wrapper around the C function gtk_grab_get_current.
func GrabGetCurrent() *Widget {
	retC := C.gtk_grab_get_current()
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_icon_size_from_name : no return generator

// Unsupported : gtk_icon_size_get_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_lookup : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_lookup_for_settings : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_register : no return generator

// Unsupported : gtk_icon_size_register_alias : unsupported parameter target : no type generator for gint (GtkIconSize) for param target

// IconThemeErrorQuark is a wrapper around the C function gtk_icon_theme_error_quark.
func IconThemeErrorQuark() glib.Quark {
	retC := C.gtk_icon_theme_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gtk_init : unsupported parameter argv :

// Unsupported : gtk_init_check : unsupported parameter argv :

// Unsupported : gtk_init_with_args : unsupported parameter argv :

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// KeySnooperRemove is a wrapper around the C function gtk_key_snooper_remove.
func KeySnooperRemove(snooperHandlerId uint32) {
	c_snooper_handler_id := (C.guint)(snooperHandlerId)

	C.gtk_key_snooper_remove(c_snooper_handler_id)

	return
}

// Main is a wrapper around the C function gtk_main.
func Main() {
	C.gtk_main()

	return
}

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// MainIteration is a wrapper around the C function gtk_main_iteration.
func MainIteration() bool {
	retC := C.gtk_main_iteration()
	retGo := retC == C.TRUE

	return retGo
}

// MainIterationDo is a wrapper around the C function gtk_main_iteration_do.
func MainIterationDo(blocking bool) bool {
	c_blocking :=
		boolToGboolean(blocking)

	retC := C.gtk_main_iteration_do(c_blocking)
	retGo := retC == C.TRUE

	return retGo
}

// MainLevel is a wrapper around the C function gtk_main_level.
func MainLevel() uint32 {
	retC := C.gtk_main_level()
	retGo := (uint32)(retC)

	return retGo
}

// MainQuit is a wrapper around the C function gtk_main_quit.
func MainQuit() {
	C.gtk_main_quit()

	return
}

// PaintArrow is a wrapper around the C function gtk_paint_arrow.
func PaintArrow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, arrowType ArrowType, fill bool, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_arrow_type := (C.GtkArrowType)(arrowType)

	c_fill :=
		boolToGboolean(fill)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_arrow(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_arrow_type, c_fill, c_x, c_y, c_width, c_height)

	return
}

// PaintBox is a wrapper around the C function gtk_paint_box.
func PaintBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_box(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintBoxGap is a wrapper around the C function gtk_paint_box_gap.
func PaintBoxGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, gapSide PositionType, gapX int32, gapWidth int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_gap_x := (C.gint)(gapX)

	c_gap_width := (C.gint)(gapWidth)

	C.gtk_paint_box_gap(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gap_side, c_gap_x, c_gap_width)

	return
}

// PaintCheck is a wrapper around the C function gtk_paint_check.
func PaintCheck(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_check(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintDiamond is a wrapper around the C function gtk_paint_diamond.
func PaintDiamond(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_diamond(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintExpander is a wrapper around the C function gtk_paint_expander.
func PaintExpander(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x int32, y int32, expanderStyle ExpanderStyle) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_expander_style := (C.GtkExpanderStyle)(expanderStyle)

	C.gtk_paint_expander(c_style, c_cr, c_state_type, c_widget, c_detail, c_x, c_y, c_expander_style)

	return
}

// PaintExtension is a wrapper around the C function gtk_paint_extension.
func PaintExtension(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, gapSide PositionType) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	C.gtk_paint_extension(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gap_side)

	return
}

// PaintFlatBox is a wrapper around the C function gtk_paint_flat_box.
func PaintFlatBox(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_flat_box(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintFocus is a wrapper around the C function gtk_paint_focus.
func PaintFocus(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_focus(c_style, c_cr, c_state_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintHandle is a wrapper around the C function gtk_paint_handle.
func PaintHandle(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, orientation Orientation) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_paint_handle(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_orientation)

	return
}

// PaintHline is a wrapper around the C function gtk_paint_hline.
func PaintHline(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, x1 int32, x2 int32, y int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x1 := (C.gint)(x1)

	c_x2 := (C.gint)(x2)

	c_y := (C.gint)(y)

	C.gtk_paint_hline(c_style, c_cr, c_state_type, c_widget, c_detail, c_x1, c_x2, c_y)

	return
}

// PaintLayout is a wrapper around the C function gtk_paint_layout.
func PaintLayout(style *Style, cr *cairo.Context, stateType StateType, useText bool, widget *Widget, detail string, x int32, y int32, layout *pango.Layout) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_use_text :=
		boolToGboolean(useText)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_layout := (*C.PangoLayout)(layout.ToC())

	C.gtk_paint_layout(c_style, c_cr, c_state_type, c_use_text, c_widget, c_detail, c_x, c_y, c_layout)

	return
}

// PaintOption is a wrapper around the C function gtk_paint_option.
func PaintOption(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_option(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintResizeGrip is a wrapper around the C function gtk_paint_resize_grip.
func PaintResizeGrip(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, edge gdk.WindowEdge, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_edge := (C.GdkWindowEdge)(edge)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_resize_grip(c_style, c_cr, c_state_type, c_widget, c_detail, c_edge, c_x, c_y, c_width, c_height)

	return
}

// PaintShadow is a wrapper around the C function gtk_paint_shadow.
func PaintShadow(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_shadow(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintShadowGap is a wrapper around the C function gtk_paint_shadow_gap.
func PaintShadowGap(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, gapSide PositionType, gapX int32, gapWidth int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_gap_x := (C.gint)(gapX)

	c_gap_width := (C.gint)(gapWidth)

	C.gtk_paint_shadow_gap(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gap_side, c_gap_x, c_gap_width)

	return
}

// PaintSlider is a wrapper around the C function gtk_paint_slider.
func PaintSlider(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32, orientation Orientation) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_paint_slider(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height, c_orientation)

	return
}

// PaintSpinner is a wrapper around the C function gtk_paint_spinner.
func PaintSpinner(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, step uint32, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_step := (C.guint)(step)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_spinner(c_style, c_cr, c_state_type, c_widget, c_detail, c_step, c_x, c_y, c_width, c_height)

	return
}

// PaintTab is a wrapper around the C function gtk_paint_tab.
func PaintTab(style *Style, cr *cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int32, y int32, width int32, height int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_tab(c_style, c_cr, c_state_type, c_shadow_type, c_widget, c_detail, c_x, c_y, c_width, c_height)

	return
}

// PaintVline is a wrapper around the C function gtk_paint_vline.
func PaintVline(style *Style, cr *cairo.Context, stateType StateType, widget *Widget, detail string, y1 int32, y2 int32, x int32) {
	c_style := (*C.GtkStyle)(style.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(widget.ToC())

	c_detail := C.CString(detail)
	defer C.free(unsafe.Pointer(c_detail))

	c_y1_ := (C.gint)(y1)

	c_y2_ := (C.gint)(y2)

	c_x := (C.gint)(x)

	C.gtk_paint_vline(c_style, c_cr, c_state_type, c_widget, c_detail, c_y1_, c_y2_, c_x)

	return
}

// Unsupported : gtk_parse_args : unsupported parameter argv :

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc (GtkPageSetupDoneFunc) for param done_cb

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// RcAddDefaultFile is a wrapper around the C function gtk_rc_add_default_file.
func RcAddDefaultFile(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_rc_add_default_file(c_filename)

	return
}

// RcFindModuleInPath is a wrapper around the C function gtk_rc_find_module_in_path.
func RcFindModuleInPath(moduleFile string) string {
	c_module_file := C.CString(moduleFile)
	defer C.free(unsafe.Pointer(c_module_file))

	retC := C.gtk_rc_find_module_in_path(c_module_file)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RcFindPixmapInPath is a wrapper around the C function gtk_rc_find_pixmap_in_path.
func RcFindPixmapInPath(settings *Settings, scanner *glib.Scanner, pixmapFile string) string {
	c_settings := (*C.GtkSettings)(settings.ToC())

	c_scanner := (*C.GScanner)(scanner.ToC())

	c_pixmap_file := C.CString(pixmapFile)
	defer C.free(unsafe.Pointer(c_pixmap_file))

	retC := C.gtk_rc_find_pixmap_in_path(c_settings, c_scanner, c_pixmap_file)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_rc_get_default_files : no return type

// RcGetImModuleFile is a wrapper around the C function gtk_rc_get_im_module_file.
func RcGetImModuleFile() string {
	retC := C.gtk_rc_get_im_module_file()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RcGetImModulePath is a wrapper around the C function gtk_rc_get_im_module_path.
func RcGetImModulePath() string {
	retC := C.gtk_rc_get_im_module_path()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RcGetModuleDir is a wrapper around the C function gtk_rc_get_module_dir.
func RcGetModuleDir() string {
	retC := C.gtk_rc_get_module_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RcGetStyle is a wrapper around the C function gtk_rc_get_style.
func RcGetStyle(widget *Widget) *Style {
	c_widget := (*C.GtkWidget)(widget.ToC())

	retC := C.gtk_rc_get_style(c_widget)
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RcGetStyleByPaths is a wrapper around the C function gtk_rc_get_style_by_paths.
func RcGetStyleByPaths(settings *Settings, widgetPath string, classPath string, type_ gobject.Type) *Style {
	c_settings := (*C.GtkSettings)(settings.ToC())

	c_widget_path := C.CString(widgetPath)
	defer C.free(unsafe.Pointer(c_widget_path))

	c_class_path := C.CString(classPath)
	defer C.free(unsafe.Pointer(c_class_path))

	c_type := (C.GType)(type_)

	retC := C.gtk_rc_get_style_by_paths(c_settings, c_widget_path, c_class_path, c_type)
	var retGo (*Style)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// RcGetThemeDir is a wrapper around the C function gtk_rc_get_theme_dir.
func RcGetThemeDir() string {
	retC := C.gtk_rc_get_theme_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RcParse is a wrapper around the C function gtk_rc_parse.
func RcParse(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_rc_parse(c_filename)

	return
}

// RcParseColor is a wrapper around the C function gtk_rc_parse_color.
func RcParseColor(scanner *glib.Scanner) (uint32, *gdk.Color) {
	c_scanner := (*C.GScanner)(scanner.ToC())

	var c_color C.GdkColor

	retC := C.gtk_rc_parse_color(c_scanner, &c_color)
	retGo := (uint32)(retC)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// RcParseString is a wrapper around the C function gtk_rc_parse_string.
func RcParseString(rcString string) {
	c_rc_string := C.CString(rcString)
	defer C.free(unsafe.Pointer(c_rc_string))

	C.gtk_rc_parse_string(c_rc_string)

	return
}

// Unsupported : gtk_rc_property_parse_border : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_color : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_enum : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_flags : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_requisition : unsupported parameter pspec : Blacklisted record : GParamSpec

// RcReparseAll is a wrapper around the C function gtk_rc_reparse_all.
func RcReparseAll() bool {
	retC := C.gtk_rc_reparse_all()
	retGo := retC == C.TRUE

	return retGo
}

// RcReparseAllForSettings is a wrapper around the C function gtk_rc_reparse_all_for_settings.
func RcReparseAllForSettings(settings *Settings, forceLoad bool) bool {
	c_settings := (*C.GtkSettings)(settings.ToC())

	c_force_load :=
		boolToGboolean(forceLoad)

	retC := C.gtk_rc_reparse_all_for_settings(c_settings, c_force_load)
	retGo := retC == C.TRUE

	return retGo
}

// RcScannerNew is a wrapper around the C function gtk_rc_scanner_new.
func RcScannerNew() *glib.Scanner {
	retC := C.gtk_rc_scanner_new()
	retGo := glib.ScannerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames :

// RecentChooserErrorQuark is a wrapper around the C function gtk_recent_chooser_error_quark.
func RecentChooserErrorQuark() glib.Quark {
	retC := C.gtk_recent_chooser_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// RecentManagerErrorQuark is a wrapper around the C function gtk_recent_manager_error_quark.
func RecentManagerErrorQuark() glib.Quark {
	retC := C.gtk_recent_manager_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gtk_render_background_get_clip : unsupported parameter out_clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_render_icon_pixbuf : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// SelectionRemoveAll is a wrapper around the C function gtk_selection_remove_all.
func SelectionRemoveAll(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_selection_remove_all(c_widget)

	return
}

// SetDebugFlags is a wrapper around the C function gtk_set_debug_flags.
func SetDebugFlags(flags uint32) {
	c_flags := (C.guint)(flags)

	C.gtk_set_debug_flags(c_flags)

	return
}

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

// StockListIds is a wrapper around the C function gtk_stock_list_ids.
func StockListIds() *glib.SList {
	retC := C.gtk_stock_list_ids()
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StockLookup is a wrapper around the C function gtk_stock_lookup.
func StockLookup(stockId string) (bool, *StockItem) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	var c_item C.GtkStockItem

	retC := C.gtk_stock_lookup(c_stock_id, &c_item)
	retGo := retC == C.TRUE

	item := StockItemNewFromC(unsafe.Pointer(&c_item))

	return retGo, item
}

// Unsupported : gtk_stock_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GtkTranslateFunc) for param func

// Unsupported : gtk_target_table_free : unsupported parameter targets :

// Unsupported : gtk_target_table_new_from_list : no return type

// Unsupported : gtk_targets_include_image : unsupported parameter targets :

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_uri : unsupported parameter targets :

// Unsupported : gtk_test_create_widget : unsupported parameter ... : varargs

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// Unsupported : gtk_test_init : unsupported parameter argvp :

// Unsupported : gtk_test_list_all_types : no return type

// Unsupported : gtk_tree_get_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel (GtkTreeModel**) for param tree_model

// TreeRowReferenceDeleted is a wrapper around the C function gtk_tree_row_reference_deleted.
func TreeRowReferenceDeleted(proxy *gobject.Object, path *TreePath) {
	c_proxy := (*C.GObject)(proxy.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_row_reference_deleted(c_proxy, c_path)

	return
}

// TreeRowReferenceInserted is a wrapper around the C function gtk_tree_row_reference_inserted.
func TreeRowReferenceInserted(proxy *gobject.Object, path *TreePath) {
	c_proxy := (*C.GObject)(proxy.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_row_reference_inserted(c_proxy, c_path)

	return
}

// TreeRowReferenceReordered is a wrapper around the C function gtk_tree_row_reference_reordered.
func TreeRowReferenceReordered(proxy *gobject.Object, path *TreePath, iter *TreeIter, newOrder []int32) {
	c_proxy := (*C.GObject)(proxy.ToC())

	c_path := (*C.GtkTreePath)(path.ToC())

	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_new_order := &newOrder[0]

	C.gtk_tree_row_reference_reordered(c_proxy, c_path, c_iter, (*C.gint)(unsafe.Pointer(c_new_order)))

	return
}

// Unsupported : gtk_tree_set_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel (GtkTreeModel*) for param tree_model

// True is a wrapper around the C function gtk_true.
func True() bool {
	retC := C.gtk_true()
	retGo := retC == C.TRUE

	return retGo
}
