// This is a generated file - DO NOT EDIT

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
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

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_key : no type generator for guint, guint*

// Unsupported : gtk_accelerator_set_default_mod_mask : no return generator

// AcceleratorValid is a wrapper around the C function gtk_accelerator_valid.
func AcceleratorValid(keyval uint32, modifiers gdk.ModifierType) bool {
	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_accelerator_valid(c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_binding_entry_add_signall : no return generator

// Unsupported : gtk_binding_entry_remove : no return generator

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
	retGo := BindingSetNewFromC(unsafe.Pointer(retC))

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

// Unsupported : gtk_disable_setlocale : no return generator

// DistributeNaturalAllocation is a wrapper around the C function gtk_distribute_natural_allocation.
func DistributeNaturalAllocation(extraSpace int32, nRequestedSizes uint32, sizes *RequestedSize) int32 {
	c_extra_space := (C.gint)(extraSpace)

	c_n_requested_sizes := (C.guint)(nRequestedSizes)

	c_sizes := (*C.GtkRequestedSize)(sizes.ToC())

	retC := C.gtk_distribute_natural_allocation(c_extra_space, c_n_requested_sizes, c_sizes)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_drag_finish : no return generator

// DragGetSourceWidget is a wrapper around the C function gtk_drag_get_source_widget.
func DragGetSourceWidget(context *gdk.DragContext) *Widget {
	c_context := (*C.GdkDragContext)(context.ToC())

	retC := C.gtk_drag_get_source_widget(c_context)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_drag_set_icon_default : no return generator

// Unsupported : gtk_drag_set_icon_pixbuf : no return generator

// Unsupported : gtk_drag_set_icon_stock : no return generator

// Unsupported : gtk_drag_set_icon_surface : no return generator

// Unsupported : gtk_drag_set_icon_widget : no return generator

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
	retGo := gdk.DeviceNewFromC(unsafe.Pointer(retC))

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

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GrabGetCurrent is a wrapper around the C function gtk_grab_get_current.
func GrabGetCurrent() *Widget {
	retC := C.gtk_grab_get_current()
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_size_from_name : no return generator

// Unsupported : gtk_icon_size_get_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_lookup : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_icon_size_register : no return generator

// Unsupported : gtk_icon_size_register_alias : unsupported parameter target : no type generator for gint, GtkIconSize

// IconThemeErrorQuark is a wrapper around the C function gtk_icon_theme_error_quark.
func IconThemeErrorQuark() glib.Quark {
	retC := C.gtk_icon_theme_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gtk_init : unsupported parameter argc : no type generator for gint, int*

// Unsupported : gtk_init_check : unsupported parameter argc : no type generator for gint, int*

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc, GtkKeySnoopFunc

// Unsupported : gtk_key_snooper_remove : no return generator

// Unsupported : gtk_main : no return generator

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

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

// Unsupported : gtk_main_quit : no return generator

// Unsupported : gtk_paint_arrow : no return generator

// Unsupported : gtk_paint_box : no return generator

// Unsupported : gtk_paint_box_gap : no return generator

// Unsupported : gtk_paint_check : no return generator

// Unsupported : gtk_paint_diamond : no return generator

// Unsupported : gtk_paint_expander : no return generator

// Unsupported : gtk_paint_extension : no return generator

// Unsupported : gtk_paint_flat_box : no return generator

// Unsupported : gtk_paint_focus : no return generator

// Unsupported : gtk_paint_handle : no return generator

// Unsupported : gtk_paint_hline : no return generator

// Unsupported : gtk_paint_layout : no return generator

// Unsupported : gtk_paint_option : no return generator

// Unsupported : gtk_paint_resize_grip : no return generator

// Unsupported : gtk_paint_shadow : no return generator

// Unsupported : gtk_paint_shadow_gap : no return generator

// Unsupported : gtk_paint_slider : no return generator

// Unsupported : gtk_paint_spinner : no return generator

// Unsupported : gtk_paint_tab : no return generator

// Unsupported : gtk_paint_vline : no return generator

// Unsupported : gtk_parse_args : unsupported parameter argc : no type generator for gint, int*

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// Unsupported : gtk_rc_add_default_file : no return generator

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

// Unsupported : gtk_rc_get_style_by_paths : unsupported parameter type : no type generator for GType, GType

// RcGetThemeDir is a wrapper around the C function gtk_rc_get_theme_dir.
func RcGetThemeDir() string {
	retC := C.gtk_rc_get_theme_dir()
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_rc_parse : no return generator

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

// Unsupported : gtk_rc_parse_string : no return generator

// RcPropertyParseBorder is a wrapper around the C function gtk_rc_property_parse_border.
func RcPropertyParseBorder(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_gstring := (*C.GString)(gstring.ToC())

	c_property_value := (*C.GValue)(propertyValue.ToC())

	retC := C.gtk_rc_property_parse_border(c_pspec, c_gstring, c_property_value)
	retGo := retC == C.TRUE

	return retGo
}

// RcPropertyParseColor is a wrapper around the C function gtk_rc_property_parse_color.
func RcPropertyParseColor(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_gstring := (*C.GString)(gstring.ToC())

	c_property_value := (*C.GValue)(propertyValue.ToC())

	retC := C.gtk_rc_property_parse_color(c_pspec, c_gstring, c_property_value)
	retGo := retC == C.TRUE

	return retGo
}

// RcPropertyParseEnum is a wrapper around the C function gtk_rc_property_parse_enum.
func RcPropertyParseEnum(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_gstring := (*C.GString)(gstring.ToC())

	c_property_value := (*C.GValue)(propertyValue.ToC())

	retC := C.gtk_rc_property_parse_enum(c_pspec, c_gstring, c_property_value)
	retGo := retC == C.TRUE

	return retGo
}

// RcPropertyParseFlags is a wrapper around the C function gtk_rc_property_parse_flags.
func RcPropertyParseFlags(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_gstring := (*C.GString)(gstring.ToC())

	c_property_value := (*C.GValue)(propertyValue.ToC())

	retC := C.gtk_rc_property_parse_flags(c_pspec, c_gstring, c_property_value)
	retGo := retC == C.TRUE

	return retGo
}

// RcPropertyParseRequisition is a wrapper around the C function gtk_rc_property_parse_requisition.
func RcPropertyParseRequisition(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) bool {
	c_pspec := (*C.GParamSpec)(pspec.ToC())

	c_gstring := (*C.GString)(gstring.ToC())

	c_property_value := (*C.GValue)(propertyValue.ToC())

	retC := C.gtk_rc_property_parse_requisition(c_pspec, c_gstring, c_property_value)
	retGo := retC == C.TRUE

	return retGo
}

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

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames : no param type

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

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_remove_all : no return generator

// Unsupported : gtk_set_debug_flags : no return generator

// Unsupported : gtk_stock_add : unsupported parameter items : no param type

// Unsupported : gtk_stock_add_static : unsupported parameter items : no param type

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

// Unsupported : gtk_tree_get_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel**

// Unsupported : gtk_tree_row_reference_deleted : no return generator

// Unsupported : gtk_tree_row_reference_inserted : no return generator

// Unsupported : gtk_tree_row_reference_reordered : unsupported parameter new_order : no param type

// Unsupported : gtk_tree_set_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// True is a wrapper around the C function gtk_true.
func True() bool {
	retC := C.gtk_true()
	retGo := retC == C.TRUE

	return retGo
}
