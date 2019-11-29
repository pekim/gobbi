// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'gtk_accel_groups_activate' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_accel_groups_from_object' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_default_mod_mask' : return type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label_with_keycode' : parameter 'display' of type 'Gdk.Display' not supported

// UNSUPPORTED : C value 'gtk_accelerator_name' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_name_with_keycode' : parameter 'display' of type 'Gdk.Display' not supported

// UNSUPPORTED : C value 'gtk_accelerator_parse' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_parse_with_keycode' : parameter 'accelerator_codes' has no type

// UNSUPPORTED : C value 'gtk_accelerator_set_default_mod_mask' : parameter 'default_mod_mask' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_valid' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_alternative_dialog_button_order' : parameter 'screen' of type 'Gdk.Screen' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_add_signal_from_string' : return type 'GLib.TokenType' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_add_signall' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_remove' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_binding_entry_skip' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_binding_set_by_class' : parameter 'object_class' of type 'gpointer' not supported

var bindingSetFindFunction *gi.Function
var bindingSetFindFunction_Once sync.Once

func bindingSetFindFunction_Set() error {
	var err error
	bindingSetFindFunction_Once.Do(func() {
		bindingSetFindFunction, err = gi.FunctionInvokerNew("Gtk", "binding_set_find")
	})
	return err
}

// BindingSetFind is a representation of the C type gtk_binding_set_find.
func BindingSetFind(setName string) *BindingSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(setName)

	var ret gi.Argument

	err := bindingSetFindFunction_Set()
	if err == nil {
		ret = bindingSetFindFunction.Invoke(inArgs[:], nil)
	}

	retGo := &BindingSet{native: ret.Pointer()}

	return retGo
}

var bindingSetNewFunction *gi.Function
var bindingSetNewFunction_Once sync.Once

func bindingSetNewFunction_Set() error {
	var err error
	bindingSetNewFunction_Once.Do(func() {
		bindingSetNewFunction, err = gi.FunctionInvokerNew("Gtk", "binding_set_new")
	})
	return err
}

// BindingSetNew is a representation of the C type gtk_binding_set_new.
func BindingSetNew(setName string) *BindingSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(setName)

	var ret gi.Argument

	err := bindingSetNewFunction_Set()
	if err == nil {
		ret = bindingSetNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &BindingSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_bindings_activate' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_bindings_activate_event' : parameter 'object' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_builder_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_cairo_should_draw_window' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_cairo_transform_to_window' : parameter 'cr' of type 'cairo.Context' not supported

var checkVersionFunction *gi.Function
var checkVersionFunction_Once sync.Once

func checkVersionFunction_Set() error {
	var err error
	checkVersionFunction_Once.Do(func() {
		checkVersionFunction, err = gi.FunctionInvokerNew("Gtk", "check_version")
	})
	return err
}

// CheckVersion is a representation of the C type gtk_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(requiredMajor)
	inArgs[1].SetUint32(requiredMinor)
	inArgs[2].SetUint32(requiredMicro)

	var ret gi.Argument

	err := checkVersionFunction_Set()
	if err == nil {
		ret = checkVersionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_provider_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_device_grab_add' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_device_grab_remove' : parameter 'widget' of type 'Widget' not supported

var disableSetlocaleFunction *gi.Function
var disableSetlocaleFunction_Once sync.Once

func disableSetlocaleFunction_Set() error {
	var err error
	disableSetlocaleFunction_Once.Do(func() {
		disableSetlocaleFunction, err = gi.FunctionInvokerNew("Gtk", "disable_setlocale")
	})
	return err
}

// DisableSetlocale is a representation of the C type gtk_disable_setlocale.
func DisableSetlocale() {

	err := disableSetlocaleFunction_Set()
	if err == nil {
		disableSetlocaleFunction.Invoke(nil, nil)
	}

	return
}

var distributeNaturalAllocationFunction *gi.Function
var distributeNaturalAllocationFunction_Once sync.Once

func distributeNaturalAllocationFunction_Set() error {
	var err error
	distributeNaturalAllocationFunction_Once.Do(func() {
		distributeNaturalAllocationFunction, err = gi.FunctionInvokerNew("Gtk", "distribute_natural_allocation")
	})
	return err
}

// DistributeNaturalAllocation is a representation of the C type gtk_distribute_natural_allocation.
func DistributeNaturalAllocation(extraSpace int32, nRequestedSizes uint32, sizes *RequestedSize) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt32(extraSpace)
	inArgs[1].SetUint32(nRequestedSizes)
	inArgs[2].SetPointer(sizes.native)

	var ret gi.Argument

	err := distributeNaturalAllocationFunction_Set()
	if err == nil {
		ret = distributeNaturalAllocationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_drag_cancel' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_finish' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_get_source_widget' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_default' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_gicon' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_name' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_pixbuf' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_stock' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_surface' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_drag_set_icon_widget' : parameter 'context' of type 'Gdk.DragContext' not supported

// UNSUPPORTED : C value 'gtk_draw_insertion_cursor' : parameter 'widget' of type 'Widget' not supported

var eventsPendingFunction *gi.Function
var eventsPendingFunction_Once sync.Once

func eventsPendingFunction_Set() error {
	var err error
	eventsPendingFunction_Once.Do(func() {
		eventsPendingFunction, err = gi.FunctionInvokerNew("Gtk", "events_pending")
	})
	return err
}

// EventsPending is a representation of the C type gtk_events_pending.
func EventsPending() bool {

	var ret gi.Argument

	err := eventsPendingFunction_Set()
	if err == nil {
		ret = eventsPendingFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var falseFunction *gi.Function
var falseFunction_Once sync.Once

func falseFunction_Set() error {
	var err error
	falseFunction_Once.Do(func() {
		falseFunction, err = gi.FunctionInvokerNew("Gtk", "false")
	})
	return err
}

// False is a representation of the C type gtk_false.
func False() bool {

	var ret gi.Argument

	err := falseFunction_Set()
	if err == nil {
		ret = falseFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_file_chooser_error_quark' : return type 'GLib.Quark' not supported

var getBinaryAgeFunction *gi.Function
var getBinaryAgeFunction_Once sync.Once

func getBinaryAgeFunction_Set() error {
	var err error
	getBinaryAgeFunction_Once.Do(func() {
		getBinaryAgeFunction, err = gi.FunctionInvokerNew("Gtk", "get_binary_age")
	})
	return err
}

// GetBinaryAge is a representation of the C type gtk_get_binary_age.
func GetBinaryAge() uint32 {

	var ret gi.Argument

	err := getBinaryAgeFunction_Set()
	if err == nil {
		ret = getBinaryAgeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_current_event' : return type 'Gdk.Event' not supported

// UNSUPPORTED : C value 'gtk_get_current_event_device' : return type 'Gdk.Device' not supported

// UNSUPPORTED : C value 'gtk_get_current_event_state' : parameter 'state' of type 'Gdk.ModifierType' not supported

var getCurrentEventTimeFunction *gi.Function
var getCurrentEventTimeFunction_Once sync.Once

func getCurrentEventTimeFunction_Set() error {
	var err error
	getCurrentEventTimeFunction_Once.Do(func() {
		getCurrentEventTimeFunction, err = gi.FunctionInvokerNew("Gtk", "get_current_event_time")
	})
	return err
}

// GetCurrentEventTime is a representation of the C type gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {

	var ret gi.Argument

	err := getCurrentEventTimeFunction_Set()
	if err == nil {
		ret = getCurrentEventTimeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getDebugFlagsFunction *gi.Function
var getDebugFlagsFunction_Once sync.Once

func getDebugFlagsFunction_Set() error {
	var err error
	getDebugFlagsFunction_Once.Do(func() {
		getDebugFlagsFunction, err = gi.FunctionInvokerNew("Gtk", "get_debug_flags")
	})
	return err
}

// GetDebugFlags is a representation of the C type gtk_get_debug_flags.
func GetDebugFlags() uint32 {

	var ret gi.Argument

	err := getDebugFlagsFunction_Set()
	if err == nil {
		ret = getDebugFlagsFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_default_language' : return type 'Pango.Language' not supported

// UNSUPPORTED : C value 'gtk_get_event_widget' : parameter 'event' of type 'Gdk.Event' not supported

var getInterfaceAgeFunction *gi.Function
var getInterfaceAgeFunction_Once sync.Once

func getInterfaceAgeFunction_Set() error {
	var err error
	getInterfaceAgeFunction_Once.Do(func() {
		getInterfaceAgeFunction, err = gi.FunctionInvokerNew("Gtk", "get_interface_age")
	})
	return err
}

// GetInterfaceAge is a representation of the C type gtk_get_interface_age.
func GetInterfaceAge() uint32 {

	var ret gi.Argument

	err := getInterfaceAgeFunction_Set()
	if err == nil {
		ret = getInterfaceAgeFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_locale_direction' : return type 'TextDirection' not supported

var getMajorVersionFunction *gi.Function
var getMajorVersionFunction_Once sync.Once

func getMajorVersionFunction_Set() error {
	var err error
	getMajorVersionFunction_Once.Do(func() {
		getMajorVersionFunction, err = gi.FunctionInvokerNew("Gtk", "get_major_version")
	})
	return err
}

// GetMajorVersion is a representation of the C type gtk_get_major_version.
func GetMajorVersion() uint32 {

	var ret gi.Argument

	err := getMajorVersionFunction_Set()
	if err == nil {
		ret = getMajorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMicroVersionFunction *gi.Function
var getMicroVersionFunction_Once sync.Once

func getMicroVersionFunction_Set() error {
	var err error
	getMicroVersionFunction_Once.Do(func() {
		getMicroVersionFunction, err = gi.FunctionInvokerNew("Gtk", "get_micro_version")
	})
	return err
}

// GetMicroVersion is a representation of the C type gtk_get_micro_version.
func GetMicroVersion() uint32 {

	var ret gi.Argument

	err := getMicroVersionFunction_Set()
	if err == nil {
		ret = getMicroVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var getMinorVersionFunction *gi.Function
var getMinorVersionFunction_Once sync.Once

func getMinorVersionFunction_Set() error {
	var err error
	getMinorVersionFunction_Once.Do(func() {
		getMinorVersionFunction, err = gi.FunctionInvokerNew("Gtk", "get_minor_version")
	})
	return err
}

// GetMinorVersion is a representation of the C type gtk_get_minor_version.
func GetMinorVersion() uint32 {

	var ret gi.Argument

	err := getMinorVersionFunction_Set()
	if err == nil {
		ret = getMinorVersionFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_get_option_group' : return type 'GLib.OptionGroup' not supported

// UNSUPPORTED : C value 'gtk_grab_get_current' : return type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_icon_size_from_name' : return type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_get_name' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_lookup' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_lookup_for_settings' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'gtk_icon_size_register' : return type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_size_register_alias' : parameter 'target' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_theme_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_init' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_init_check' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_init_with_args' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_key_snooper_install' : parameter 'snooper' of type 'KeySnoopFunc' not supported

var keySnooperRemoveFunction *gi.Function
var keySnooperRemoveFunction_Once sync.Once

func keySnooperRemoveFunction_Set() error {
	var err error
	keySnooperRemoveFunction_Once.Do(func() {
		keySnooperRemoveFunction, err = gi.FunctionInvokerNew("Gtk", "key_snooper_remove")
	})
	return err
}

// KeySnooperRemove is a representation of the C type gtk_key_snooper_remove.
func KeySnooperRemove(snooperHandlerId uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(snooperHandlerId)

	err := keySnooperRemoveFunction_Set()
	if err == nil {
		keySnooperRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mainFunction *gi.Function
var mainFunction_Once sync.Once

func mainFunction_Set() error {
	var err error
	mainFunction_Once.Do(func() {
		mainFunction, err = gi.FunctionInvokerNew("Gtk", "main")
	})
	return err
}

// Main is a representation of the C type gtk_main.
func Main() {

	err := mainFunction_Set()
	if err == nil {
		mainFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_main_do_event' : parameter 'event' of type 'Gdk.Event' not supported

var mainIterationFunction *gi.Function
var mainIterationFunction_Once sync.Once

func mainIterationFunction_Set() error {
	var err error
	mainIterationFunction_Once.Do(func() {
		mainIterationFunction, err = gi.FunctionInvokerNew("Gtk", "main_iteration")
	})
	return err
}

// MainIteration is a representation of the C type gtk_main_iteration.
func MainIteration() bool {

	var ret gi.Argument

	err := mainIterationFunction_Set()
	if err == nil {
		ret = mainIterationFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var mainIterationDoFunction *gi.Function
var mainIterationDoFunction_Once sync.Once

func mainIterationDoFunction_Set() error {
	var err error
	mainIterationDoFunction_Once.Do(func() {
		mainIterationDoFunction, err = gi.FunctionInvokerNew("Gtk", "main_iteration_do")
	})
	return err
}

// MainIterationDo is a representation of the C type gtk_main_iteration_do.
func MainIterationDo(blocking bool) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(blocking)

	var ret gi.Argument

	err := mainIterationDoFunction_Set()
	if err == nil {
		ret = mainIterationDoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var mainLevelFunction *gi.Function
var mainLevelFunction_Once sync.Once

func mainLevelFunction_Set() error {
	var err error
	mainLevelFunction_Once.Do(func() {
		mainLevelFunction, err = gi.FunctionInvokerNew("Gtk", "main_level")
	})
	return err
}

// MainLevel is a representation of the C type gtk_main_level.
func MainLevel() uint32 {

	var ret gi.Argument

	err := mainLevelFunction_Set()
	if err == nil {
		ret = mainLevelFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var mainQuitFunction *gi.Function
var mainQuitFunction_Once sync.Once

func mainQuitFunction_Set() error {
	var err error
	mainQuitFunction_Once.Do(func() {
		mainQuitFunction, err = gi.FunctionInvokerNew("Gtk", "main_quit")
	})
	return err
}

// MainQuit is a representation of the C type gtk_main_quit.
func MainQuit() {

	err := mainQuitFunction_Set()
	if err == nil {
		mainQuitFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_paint_arrow' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_box' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_box_gap' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_check' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_diamond' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_expander' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_extension' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_flat_box' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_focus' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_handle' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_hline' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_layout' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_option' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_resize_grip' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_shadow' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_shadow_gap' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_slider' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_spinner' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_tab' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_paint_vline' : parameter 'style' of type 'Style' not supported

var paperSizeGetDefaultFunction *gi.Function
var paperSizeGetDefaultFunction_Once sync.Once

func paperSizeGetDefaultFunction_Set() error {
	var err error
	paperSizeGetDefaultFunction_Once.Do(func() {
		paperSizeGetDefaultFunction, err = gi.FunctionInvokerNew("Gtk", "paper_size_get_default")
	})
	return err
}

// PaperSizeGetDefault is a representation of the C type gtk_paper_size_get_default.
func PaperSizeGetDefault() string {

	var ret gi.Argument

	err := paperSizeGetDefaultFunction_Set()
	if err == nil {
		ret = paperSizeGetDefaultFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_paper_sizes' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'gtk_parse_args' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gtk_print_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog_async' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_propagate_event' : parameter 'widget' of type 'Widget' not supported

var rcAddDefaultFileFunction *gi.Function
var rcAddDefaultFileFunction_Once sync.Once

func rcAddDefaultFileFunction_Set() error {
	var err error
	rcAddDefaultFileFunction_Once.Do(func() {
		rcAddDefaultFileFunction, err = gi.FunctionInvokerNew("Gtk", "rc_add_default_file")
	})
	return err
}

// RcAddDefaultFile is a representation of the C type gtk_rc_add_default_file.
func RcAddDefaultFile(filename string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	err := rcAddDefaultFileFunction_Set()
	if err == nil {
		rcAddDefaultFileFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rcFindModuleInPathFunction *gi.Function
var rcFindModuleInPathFunction_Once sync.Once

func rcFindModuleInPathFunction_Set() error {
	var err error
	rcFindModuleInPathFunction_Once.Do(func() {
		rcFindModuleInPathFunction, err = gi.FunctionInvokerNew("Gtk", "rc_find_module_in_path")
	})
	return err
}

// RcFindModuleInPath is a representation of the C type gtk_rc_find_module_in_path.
func RcFindModuleInPath(moduleFile string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(moduleFile)

	var ret gi.Argument

	err := rcFindModuleInPathFunction_Set()
	if err == nil {
		ret = rcFindModuleInPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_rc_find_pixmap_in_path' : parameter 'settings' of type 'Settings' not supported

var rcGetDefaultFilesFunction *gi.Function
var rcGetDefaultFilesFunction_Once sync.Once

func rcGetDefaultFilesFunction_Set() error {
	var err error
	rcGetDefaultFilesFunction_Once.Do(func() {
		rcGetDefaultFilesFunction, err = gi.FunctionInvokerNew("Gtk", "rc_get_default_files")
	})
	return err
}

// RcGetDefaultFiles is a representation of the C type gtk_rc_get_default_files.
func RcGetDefaultFiles() {

	err := rcGetDefaultFilesFunction_Set()
	if err == nil {
		rcGetDefaultFilesFunction.Invoke(nil, nil)
	}

	return
}

var rcGetImModuleFileFunction *gi.Function
var rcGetImModuleFileFunction_Once sync.Once

func rcGetImModuleFileFunction_Set() error {
	var err error
	rcGetImModuleFileFunction_Once.Do(func() {
		rcGetImModuleFileFunction, err = gi.FunctionInvokerNew("Gtk", "rc_get_im_module_file")
	})
	return err
}

// RcGetImModuleFile is a representation of the C type gtk_rc_get_im_module_file.
func RcGetImModuleFile() string {

	var ret gi.Argument

	err := rcGetImModuleFileFunction_Set()
	if err == nil {
		ret = rcGetImModuleFileFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var rcGetImModulePathFunction *gi.Function
var rcGetImModulePathFunction_Once sync.Once

func rcGetImModulePathFunction_Set() error {
	var err error
	rcGetImModulePathFunction_Once.Do(func() {
		rcGetImModulePathFunction, err = gi.FunctionInvokerNew("Gtk", "rc_get_im_module_path")
	})
	return err
}

// RcGetImModulePath is a representation of the C type gtk_rc_get_im_module_path.
func RcGetImModulePath() string {

	var ret gi.Argument

	err := rcGetImModulePathFunction_Set()
	if err == nil {
		ret = rcGetImModulePathFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var rcGetModuleDirFunction *gi.Function
var rcGetModuleDirFunction_Once sync.Once

func rcGetModuleDirFunction_Set() error {
	var err error
	rcGetModuleDirFunction_Once.Do(func() {
		rcGetModuleDirFunction, err = gi.FunctionInvokerNew("Gtk", "rc_get_module_dir")
	})
	return err
}

// RcGetModuleDir is a representation of the C type gtk_rc_get_module_dir.
func RcGetModuleDir() string {

	var ret gi.Argument

	err := rcGetModuleDirFunction_Set()
	if err == nil {
		ret = rcGetModuleDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_rc_get_style' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_rc_get_style_by_paths' : parameter 'settings' of type 'Settings' not supported

var rcGetThemeDirFunction *gi.Function
var rcGetThemeDirFunction_Once sync.Once

func rcGetThemeDirFunction_Set() error {
	var err error
	rcGetThemeDirFunction_Once.Do(func() {
		rcGetThemeDirFunction, err = gi.FunctionInvokerNew("Gtk", "rc_get_theme_dir")
	})
	return err
}

// RcGetThemeDir is a representation of the C type gtk_rc_get_theme_dir.
func RcGetThemeDir() string {

	var ret gi.Argument

	err := rcGetThemeDirFunction_Set()
	if err == nil {
		ret = rcGetThemeDirFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var rcParseFunction *gi.Function
var rcParseFunction_Once sync.Once

func rcParseFunction_Set() error {
	var err error
	rcParseFunction_Once.Do(func() {
		rcParseFunction, err = gi.FunctionInvokerNew("Gtk", "rc_parse")
	})
	return err
}

// RcParse is a representation of the C type gtk_rc_parse.
func RcParse(filename string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	err := rcParseFunction_Set()
	if err == nil {
		rcParseFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_rc_parse_color' : parameter 'scanner' of type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_parse_color_full' : parameter 'scanner' of type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_parse_priority' : parameter 'scanner' of type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_parse_state' : parameter 'scanner' of type 'GLib.Scanner' not supported

var rcParseStringFunction *gi.Function
var rcParseStringFunction_Once sync.Once

func rcParseStringFunction_Set() error {
	var err error
	rcParseStringFunction_Once.Do(func() {
		rcParseStringFunction, err = gi.FunctionInvokerNew("Gtk", "rc_parse_string")
	})
	return err
}

// RcParseString is a representation of the C type gtk_rc_parse_string.
func RcParseString(rcString string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(rcString)

	err := rcParseStringFunction_Set()
	if err == nil {
		rcParseStringFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_rc_property_parse_border' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_color' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_enum' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_flags' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_requisition' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var rcReparseAllFunction *gi.Function
var rcReparseAllFunction_Once sync.Once

func rcReparseAllFunction_Set() error {
	var err error
	rcReparseAllFunction_Once.Do(func() {
		rcReparseAllFunction, err = gi.FunctionInvokerNew("Gtk", "rc_reparse_all")
	})
	return err
}

// RcReparseAll is a representation of the C type gtk_rc_reparse_all.
func RcReparseAll() bool {

	var ret gi.Argument

	err := rcReparseAllFunction_Set()
	if err == nil {
		ret = rcReparseAllFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_rc_reparse_all_for_settings' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'gtk_rc_reset_styles' : parameter 'settings' of type 'Settings' not supported

// UNSUPPORTED : C value 'gtk_rc_scanner_new' : return type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_set_default_files' : parameter 'filenames' has no type

// UNSUPPORTED : C value 'gtk_recent_chooser_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_recent_manager_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_render_activity' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_arrow' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_background' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_background_get_clip' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_check' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_expander' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_extension' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_focus' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_frame' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_frame_gap' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_handle' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_icon' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_icon_pixbuf' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_icon_surface' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_insertion_cursor' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_layout' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_line' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_option' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_render_slider' : parameter 'context' of type 'StyleContext' not supported

var rgbToHsvFunction *gi.Function
var rgbToHsvFunction_Once sync.Once

func rgbToHsvFunction_Set() error {
	var err error
	rgbToHsvFunction_Once.Do(func() {
		rgbToHsvFunction, err = gi.FunctionInvokerNew("Gtk", "rgb_to_hsv")
	})
	return err
}

// RgbToHsv is a representation of the C type gtk_rgb_to_hsv.
func RgbToHsv(r float64, g float64, b float64) (float64, float64, float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetFloat64(r)
	inArgs[1].SetFloat64(g)
	inArgs[2].SetFloat64(b)

	var outArgs [3]gi.Argument

	err := rgbToHsvFunction_Set()
	if err == nil {
		rgbToHsvFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()
	out2 := outArgs[2].Float64()

	return out0, out1, out2
}

// UNSUPPORTED : C value 'gtk_selection_add_target' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_add_targets' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_clear_targets' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_convert' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_owner_set' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_selection_owner_set_for_display' : parameter 'display' of type 'Gdk.Display' not supported

// UNSUPPORTED : C value 'gtk_selection_remove_all' : parameter 'widget' of type 'Widget' not supported

var setDebugFlagsFunction *gi.Function
var setDebugFlagsFunction_Once sync.Once

func setDebugFlagsFunction_Set() error {
	var err error
	setDebugFlagsFunction_Once.Do(func() {
		setDebugFlagsFunction, err = gi.FunctionInvokerNew("Gtk", "set_debug_flags")
	})
	return err
}

// SetDebugFlags is a representation of the C type gtk_set_debug_flags.
func SetDebugFlags(flags uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(flags)

	err := setDebugFlagsFunction_Set()
	if err == nil {
		setDebugFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_show_about_dialog' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_show_uri' : parameter 'screen' of type 'Gdk.Screen' not supported

// UNSUPPORTED : C value 'gtk_show_uri_on_window' : parameter 'parent' of type 'Window' not supported

// UNSUPPORTED : C value 'gtk_stock_add' : parameter 'items' has no type

// UNSUPPORTED : C value 'gtk_stock_add_static' : parameter 'items' has no type

// UNSUPPORTED : C value 'gtk_stock_list_ids' : return type 'GLib.SList' not supported

var stockLookupFunction *gi.Function
var stockLookupFunction_Once sync.Once

func stockLookupFunction_Set() error {
	var err error
	stockLookupFunction_Once.Do(func() {
		stockLookupFunction, err = gi.FunctionInvokerNew("Gtk", "stock_lookup")
	})
	return err
}

// StockLookup is a representation of the C type gtk_stock_lookup.
func StockLookup(stockId string) (bool, *StockItem) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(stockId)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := stockLookupFunction_Set()
	if err == nil {
		ret = stockLookupFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := &StockItem{native: outArgs[0].Pointer()}

	return retGo, out0
}

// UNSUPPORTED : C value 'gtk_stock_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

// UNSUPPORTED : C value 'gtk_target_table_free' : parameter 'targets' has no type

var targetTableNewFromListFunction *gi.Function
var targetTableNewFromListFunction_Once sync.Once

func targetTableNewFromListFunction_Set() error {
	var err error
	targetTableNewFromListFunction_Once.Do(func() {
		targetTableNewFromListFunction, err = gi.FunctionInvokerNew("Gtk", "target_table_new_from_list")
	})
	return err
}

// TargetTableNewFromList is a representation of the C type gtk_target_table_new_from_list.
func TargetTableNewFromList(list *TargetList) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(list.native)

	var outArgs [1]gi.Argument

	err := targetTableNewFromListFunction_Set()
	if err == nil {
		targetTableNewFromListFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_targets_include_image' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_targets_include_rich_text' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_targets_include_text' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_targets_include_uri' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_test_create_simple_window' : return type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_create_widget' : parameter 'widget_type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_test_display_button_window' : parameter '...' has no type

// UNSUPPORTED : C value 'gtk_test_find_label' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_find_sibling' : parameter 'base_widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_find_widget' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_init' : parameter 'argvp' has no type

var testListAllTypesFunction *gi.Function
var testListAllTypesFunction_Once sync.Once

func testListAllTypesFunction_Set() error {
	var err error
	testListAllTypesFunction_Once.Do(func() {
		testListAllTypesFunction, err = gi.FunctionInvokerNew("Gtk", "test_list_all_types")
	})
	return err
}

// TestListAllTypes is a representation of the C type gtk_test_list_all_types.
func TestListAllTypes() uint32 {

	var outArgs [1]gi.Argument

	err := testListAllTypesFunction_Set()
	if err == nil {
		testListAllTypesFunction.Invoke(nil, outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var testRegisterAllTypesFunction *gi.Function
var testRegisterAllTypesFunction_Once sync.Once

func testRegisterAllTypesFunction_Set() error {
	var err error
	testRegisterAllTypesFunction_Once.Do(func() {
		testRegisterAllTypesFunction, err = gi.FunctionInvokerNew("Gtk", "test_register_all_types")
	})
	return err
}

// TestRegisterAllTypes is a representation of the C type gtk_test_register_all_types.
func TestRegisterAllTypes() {

	err := testRegisterAllTypesFunction_Set()
	if err == nil {
		testRegisterAllTypesFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_test_slider_get_value' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_slider_set_perc' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_spin_button_click' : parameter 'spinner' of type 'SpinButton' not supported

// UNSUPPORTED : C value 'gtk_test_text_get' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_text_set' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_widget_click' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_widget_send_key' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_test_widget_wait_for_draw' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_tree_get_row_drag_data' : parameter 'tree_model' of type 'TreeModel' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_deleted' : parameter 'proxy' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_inserted' : parameter 'proxy' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_reordered' : parameter 'proxy' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gtk_tree_set_row_drag_data' : parameter 'tree_model' of type 'TreeModel' not supported

var trueFunction *gi.Function
var trueFunction_Once sync.Once

func trueFunction_Set() error {
	var err error
	trueFunction_Once.Do(func() {
		trueFunction, err = gi.FunctionInvokerNew("Gtk", "true")
	})
	return err
}

// True is a representation of the C type gtk_true.
func True() bool {

	var ret gi.Argument

	err := trueFunction_Set()
	if err == nil {
		ret = trueFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}
