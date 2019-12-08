// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
)

// UNSUPPORTED : C value 'gtk_accel_groups_activate' : parameter 'accel_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accel_groups_from_object' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_default_mod_mask' : return type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_get_label_with_keycode' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_name' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_name_with_keycode' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_parse' : parameter 'accelerator_mods' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_parse_with_keycode' : parameter 'accelerator_codes' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_accelerator_set_default_mod_mask' : parameter 'default_mod_mask' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_accelerator_valid' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

var alternativeDialogButtonOrderFunction *gi.Function
var alternativeDialogButtonOrderFunction_Once sync.Once

func alternativeDialogButtonOrderFunction_Set() error {
	var err error
	alternativeDialogButtonOrderFunction_Once.Do(func() {
		alternativeDialogButtonOrderFunction, err = gi.FunctionInvokerNew("Gtk", "alternative_dialog_button_order")
	})
	return err
}

// AlternativeDialogButtonOrder is a representation of the C type gtk_alternative_dialog_button_order.
func AlternativeDialogButtonOrder(screen *gdk.Screen) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(screen.Native())

	var ret gi.Argument

	err := alternativeDialogButtonOrderFunction_Set()
	if err == nil {
		ret = alternativeDialogButtonOrderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

	retGo := BindingSetNewFromNative(ret.Pointer())

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

	retGo := BindingSetNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_bindings_activate' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_bindings_activate_event' : parameter 'event' of type 'Gdk.EventKey' not supported

var builderErrorQuarkFunction *gi.Function
var builderErrorQuarkFunction_Once sync.Once

func builderErrorQuarkFunction_Set() error {
	var err error
	builderErrorQuarkFunction_Once.Do(func() {
		builderErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "builder_error_quark")
	})
	return err
}

// BuilderErrorQuark is a representation of the C type gtk_builder_error_quark.
func BuilderErrorQuark() glib.Quark {

	var ret gi.Argument

	err := builderErrorQuarkFunction_Set()
	if err == nil {
		ret = builderErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

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

var cssProviderErrorQuarkFunction *gi.Function
var cssProviderErrorQuarkFunction_Once sync.Once

func cssProviderErrorQuarkFunction_Set() error {
	var err error
	cssProviderErrorQuarkFunction_Once.Do(func() {
		cssProviderErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "css_provider_error_quark")
	})
	return err
}

// CssProviderErrorQuark is a representation of the C type gtk_css_provider_error_quark.
func CssProviderErrorQuark() glib.Quark {

	var ret gi.Argument

	err := cssProviderErrorQuarkFunction_Set()
	if err == nil {
		ret = cssProviderErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var deviceGrabAddFunction *gi.Function
var deviceGrabAddFunction_Once sync.Once

func deviceGrabAddFunction_Set() error {
	var err error
	deviceGrabAddFunction_Once.Do(func() {
		deviceGrabAddFunction, err = gi.FunctionInvokerNew("Gtk", "device_grab_add")
	})
	return err
}

// DeviceGrabAdd is a representation of the C type gtk_device_grab_add.
func DeviceGrabAdd(widget *Widget, device *gdk.Device, blockOthers bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(widget.Native())
	inArgs[1].SetPointer(device.Native())
	inArgs[2].SetBoolean(blockOthers)

	err := deviceGrabAddFunction_Set()
	if err == nil {
		deviceGrabAddFunction.Invoke(inArgs[:], nil)
	}

	return
}

var deviceGrabRemoveFunction *gi.Function
var deviceGrabRemoveFunction_Once sync.Once

func deviceGrabRemoveFunction_Set() error {
	var err error
	deviceGrabRemoveFunction_Once.Do(func() {
		deviceGrabRemoveFunction, err = gi.FunctionInvokerNew("Gtk", "device_grab_remove")
	})
	return err
}

// DeviceGrabRemove is a representation of the C type gtk_device_grab_remove.
func DeviceGrabRemove(widget *Widget, device *gdk.Device) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(widget.Native())
	inArgs[1].SetPointer(device.Native())

	err := deviceGrabRemoveFunction_Set()
	if err == nil {
		deviceGrabRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

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
	inArgs[2].SetPointer(sizes.Native())

	var ret gi.Argument

	err := distributeNaturalAllocationFunction_Set()
	if err == nil {
		ret = distributeNaturalAllocationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var dragCancelFunction *gi.Function
var dragCancelFunction_Once sync.Once

func dragCancelFunction_Set() error {
	var err error
	dragCancelFunction_Once.Do(func() {
		dragCancelFunction, err = gi.FunctionInvokerNew("Gtk", "drag_cancel")
	})
	return err
}

// DragCancel is a representation of the C type gtk_drag_cancel.
func DragCancel(context *gdk.DragContext) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	err := dragCancelFunction_Set()
	if err == nil {
		dragCancelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragFinishFunction *gi.Function
var dragFinishFunction_Once sync.Once

func dragFinishFunction_Set() error {
	var err error
	dragFinishFunction_Once.Do(func() {
		dragFinishFunction, err = gi.FunctionInvokerNew("Gtk", "drag_finish")
	})
	return err
}

// DragFinish is a representation of the C type gtk_drag_finish.
func DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetBoolean(success)
	inArgs[2].SetBoolean(del)
	inArgs[3].SetUint32(time)

	err := dragFinishFunction_Set()
	if err == nil {
		dragFinishFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragGetSourceWidgetFunction *gi.Function
var dragGetSourceWidgetFunction_Once sync.Once

func dragGetSourceWidgetFunction_Set() error {
	var err error
	dragGetSourceWidgetFunction_Once.Do(func() {
		dragGetSourceWidgetFunction, err = gi.FunctionInvokerNew("Gtk", "drag_get_source_widget")
	})
	return err
}

// DragGetSourceWidget is a representation of the C type gtk_drag_get_source_widget.
func DragGetSourceWidget(context *gdk.DragContext) *Widget {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	var ret gi.Argument

	err := dragGetSourceWidgetFunction_Set()
	if err == nil {
		ret = dragGetSourceWidgetFunction.Invoke(inArgs[:], nil)
	}

	retGo := WidgetNewFromNative(ret.Pointer())

	return retGo
}

var dragSetIconDefaultFunction *gi.Function
var dragSetIconDefaultFunction_Once sync.Once

func dragSetIconDefaultFunction_Set() error {
	var err error
	dragSetIconDefaultFunction_Once.Do(func() {
		dragSetIconDefaultFunction, err = gi.FunctionInvokerNew("Gtk", "drag_set_icon_default")
	})
	return err
}

// DragSetIconDefault is a representation of the C type gtk_drag_set_icon_default.
func DragSetIconDefault(context *gdk.DragContext) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native())

	err := dragSetIconDefaultFunction_Set()
	if err == nil {
		dragSetIconDefaultFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_drag_set_icon_gicon' : parameter 'icon' of type 'Gio.Icon' not supported

var dragSetIconNameFunction *gi.Function
var dragSetIconNameFunction_Once sync.Once

func dragSetIconNameFunction_Set() error {
	var err error
	dragSetIconNameFunction_Once.Do(func() {
		dragSetIconNameFunction, err = gi.FunctionInvokerNew("Gtk", "drag_set_icon_name")
	})
	return err
}

// DragSetIconName is a representation of the C type gtk_drag_set_icon_name.
func DragSetIconName(context *gdk.DragContext, iconName string, hotX int32, hotY int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetString(iconName)
	inArgs[2].SetInt32(hotX)
	inArgs[3].SetInt32(hotY)

	err := dragSetIconNameFunction_Set()
	if err == nil {
		dragSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragSetIconPixbufFunction *gi.Function
var dragSetIconPixbufFunction_Once sync.Once

func dragSetIconPixbufFunction_Set() error {
	var err error
	dragSetIconPixbufFunction_Once.Do(func() {
		dragSetIconPixbufFunction, err = gi.FunctionInvokerNew("Gtk", "drag_set_icon_pixbuf")
	})
	return err
}

// DragSetIconPixbuf is a representation of the C type gtk_drag_set_icon_pixbuf.
func DragSetIconPixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int32, hotY int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetPointer(pixbuf.Native())
	inArgs[2].SetInt32(hotX)
	inArgs[3].SetInt32(hotY)

	err := dragSetIconPixbufFunction_Set()
	if err == nil {
		dragSetIconPixbufFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragSetIconStockFunction *gi.Function
var dragSetIconStockFunction_Once sync.Once

func dragSetIconStockFunction_Set() error {
	var err error
	dragSetIconStockFunction_Once.Do(func() {
		dragSetIconStockFunction, err = gi.FunctionInvokerNew("Gtk", "drag_set_icon_stock")
	})
	return err
}

// DragSetIconStock is a representation of the C type gtk_drag_set_icon_stock.
func DragSetIconStock(context *gdk.DragContext, stockId string, hotX int32, hotY int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetString(stockId)
	inArgs[2].SetInt32(hotX)
	inArgs[3].SetInt32(hotY)

	err := dragSetIconStockFunction_Set()
	if err == nil {
		dragSetIconStockFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_drag_set_icon_surface' : parameter 'surface' of type 'cairo.Surface' not supported

var dragSetIconWidgetFunction *gi.Function
var dragSetIconWidgetFunction_Once sync.Once

func dragSetIconWidgetFunction_Set() error {
	var err error
	dragSetIconWidgetFunction_Once.Do(func() {
		dragSetIconWidgetFunction, err = gi.FunctionInvokerNew("Gtk", "drag_set_icon_widget")
	})
	return err
}

// DragSetIconWidget is a representation of the C type gtk_drag_set_icon_widget.
func DragSetIconWidget(context *gdk.DragContext, widget *Widget, hotX int32, hotY int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetPointer(widget.Native())
	inArgs[2].SetInt32(hotX)
	inArgs[3].SetInt32(hotY)

	err := dragSetIconWidgetFunction_Set()
	if err == nil {
		dragSetIconWidgetFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_draw_insertion_cursor' : parameter 'cr' of type 'cairo.Context' not supported

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

var fileChooserErrorQuarkFunction *gi.Function
var fileChooserErrorQuarkFunction_Once sync.Once

func fileChooserErrorQuarkFunction_Set() error {
	var err error
	fileChooserErrorQuarkFunction_Once.Do(func() {
		fileChooserErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "file_chooser_error_quark")
	})
	return err
}

// FileChooserErrorQuark is a representation of the C type gtk_file_chooser_error_quark.
func FileChooserErrorQuark() glib.Quark {

	var ret gi.Argument

	err := fileChooserErrorQuarkFunction_Set()
	if err == nil {
		ret = fileChooserErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

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

var getCurrentEventDeviceFunction *gi.Function
var getCurrentEventDeviceFunction_Once sync.Once

func getCurrentEventDeviceFunction_Set() error {
	var err error
	getCurrentEventDeviceFunction_Once.Do(func() {
		getCurrentEventDeviceFunction, err = gi.FunctionInvokerNew("Gtk", "get_current_event_device")
	})
	return err
}

// GetCurrentEventDevice is a representation of the C type gtk_get_current_event_device.
func GetCurrentEventDevice() *gdk.Device {

	var ret gi.Argument

	err := getCurrentEventDeviceFunction_Set()
	if err == nil {
		ret = getCurrentEventDeviceFunction.Invoke(nil, nil)
	}

	retGo := gdk.DeviceNewFromNative(ret.Pointer())

	return retGo
}

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

var getLocaleDirectionFunction *gi.Function
var getLocaleDirectionFunction_Once sync.Once

func getLocaleDirectionFunction_Set() error {
	var err error
	getLocaleDirectionFunction_Once.Do(func() {
		getLocaleDirectionFunction, err = gi.FunctionInvokerNew("Gtk", "get_locale_direction")
	})
	return err
}

// GetLocaleDirection is a representation of the C type gtk_get_locale_direction.
func GetLocaleDirection() TextDirection {

	var ret gi.Argument

	err := getLocaleDirectionFunction_Set()
	if err == nil {
		ret = getLocaleDirectionFunction.Invoke(nil, nil)
	}

	retGo := TextDirection(ret.Int32())

	return retGo
}

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

var grabGetCurrentFunction *gi.Function
var grabGetCurrentFunction_Once sync.Once

func grabGetCurrentFunction_Set() error {
	var err error
	grabGetCurrentFunction_Once.Do(func() {
		grabGetCurrentFunction, err = gi.FunctionInvokerNew("Gtk", "grab_get_current")
	})
	return err
}

// GrabGetCurrent is a representation of the C type gtk_grab_get_current.
func GrabGetCurrent() *Widget {

	var ret gi.Argument

	err := grabGetCurrentFunction_Set()
	if err == nil {
		ret = grabGetCurrentFunction.Invoke(nil, nil)
	}

	retGo := WidgetNewFromNative(ret.Pointer())

	return retGo
}

var iconSizeFromNameFunction *gi.Function
var iconSizeFromNameFunction_Once sync.Once

func iconSizeFromNameFunction_Set() error {
	var err error
	iconSizeFromNameFunction_Once.Do(func() {
		iconSizeFromNameFunction, err = gi.FunctionInvokerNew("Gtk", "icon_size_from_name")
	})
	return err
}

// IconSizeFromName is a representation of the C type gtk_icon_size_from_name.
func IconSizeFromName(name string) IconSize {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := iconSizeFromNameFunction_Set()
	if err == nil {
		ret = iconSizeFromNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconSize(ret.Int32())

	return retGo
}

var iconSizeGetNameFunction *gi.Function
var iconSizeGetNameFunction_Once sync.Once

func iconSizeGetNameFunction_Set() error {
	var err error
	iconSizeGetNameFunction_Once.Do(func() {
		iconSizeGetNameFunction, err = gi.FunctionInvokerNew("Gtk", "icon_size_get_name")
	})
	return err
}

// IconSizeGetName is a representation of the C type gtk_icon_size_get_name.
func IconSizeGetName(size IconSize) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(size))

	var ret gi.Argument

	err := iconSizeGetNameFunction_Set()
	if err == nil {
		ret = iconSizeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var iconSizeLookupFunction *gi.Function
var iconSizeLookupFunction_Once sync.Once

func iconSizeLookupFunction_Set() error {
	var err error
	iconSizeLookupFunction_Once.Do(func() {
		iconSizeLookupFunction, err = gi.FunctionInvokerNew("Gtk", "icon_size_lookup")
	})
	return err
}

// IconSizeLookup is a representation of the C type gtk_icon_size_lookup.
func IconSizeLookup(size IconSize) (bool, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(size))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := iconSizeLookupFunction_Set()
	if err == nil {
		ret = iconSizeLookupFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var iconSizeLookupForSettingsFunction *gi.Function
var iconSizeLookupForSettingsFunction_Once sync.Once

func iconSizeLookupForSettingsFunction_Set() error {
	var err error
	iconSizeLookupForSettingsFunction_Once.Do(func() {
		iconSizeLookupForSettingsFunction, err = gi.FunctionInvokerNew("Gtk", "icon_size_lookup_for_settings")
	})
	return err
}

// IconSizeLookupForSettings is a representation of the C type gtk_icon_size_lookup_for_settings.
func IconSizeLookupForSettings(settings *Settings, size IconSize) (bool, int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(settings.Native())
	inArgs[1].SetInt32(int32(size))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := iconSizeLookupForSettingsFunction_Set()
	if err == nil {
		ret = iconSizeLookupForSettingsFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var iconSizeRegisterFunction *gi.Function
var iconSizeRegisterFunction_Once sync.Once

func iconSizeRegisterFunction_Set() error {
	var err error
	iconSizeRegisterFunction_Once.Do(func() {
		iconSizeRegisterFunction, err = gi.FunctionInvokerNew("Gtk", "icon_size_register")
	})
	return err
}

// IconSizeRegister is a representation of the C type gtk_icon_size_register.
func IconSizeRegister(name string, width int32, height int32) IconSize {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)

	var ret gi.Argument

	err := iconSizeRegisterFunction_Set()
	if err == nil {
		ret = iconSizeRegisterFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconSize(ret.Int32())

	return retGo
}

var iconSizeRegisterAliasFunction *gi.Function
var iconSizeRegisterAliasFunction_Once sync.Once

func iconSizeRegisterAliasFunction_Set() error {
	var err error
	iconSizeRegisterAliasFunction_Once.Do(func() {
		iconSizeRegisterAliasFunction, err = gi.FunctionInvokerNew("Gtk", "icon_size_register_alias")
	})
	return err
}

// IconSizeRegisterAlias is a representation of the C type gtk_icon_size_register_alias.
func IconSizeRegisterAlias(alias string, target IconSize) {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(alias)
	inArgs[1].SetInt32(int32(target))

	err := iconSizeRegisterAliasFunction_Set()
	if err == nil {
		iconSizeRegisterAliasFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iconThemeErrorQuarkFunction *gi.Function
var iconThemeErrorQuarkFunction_Once sync.Once

func iconThemeErrorQuarkFunction_Set() error {
	var err error
	iconThemeErrorQuarkFunction_Once.Do(func() {
		iconThemeErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "icon_theme_error_quark")
	})
	return err
}

// IconThemeErrorQuark is a representation of the C type gtk_icon_theme_error_quark.
func IconThemeErrorQuark() glib.Quark {

	var ret gi.Argument

	err := iconThemeErrorQuarkFunction_Set()
	if err == nil {
		ret = iconThemeErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_init' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_init_check' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_init_with_args' : parameter 'argv' of type 'nil' not supported

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

// UNSUPPORTED : C value 'gtk_paint_arrow' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_box' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_box_gap' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_check' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_diamond' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_expander' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_extension' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_flat_box' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_focus' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_handle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_hline' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_layout' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_option' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_resize_grip' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_shadow' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_shadow_gap' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_slider' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_spinner' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_tab' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_paint_vline' : parameter 'cr' of type 'cairo.Context' not supported

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

// UNSUPPORTED : C value 'gtk_parse_args' : parameter 'argv' of type 'nil' not supported

var printErrorQuarkFunction *gi.Function
var printErrorQuarkFunction_Once sync.Once

func printErrorQuarkFunction_Set() error {
	var err error
	printErrorQuarkFunction_Once.Do(func() {
		printErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "print_error_quark")
	})
	return err
}

// PrintErrorQuark is a representation of the C type gtk_print_error_quark.
func PrintErrorQuark() glib.Quark {

	var ret gi.Argument

	err := printErrorQuarkFunction_Set()
	if err == nil {
		ret = printErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var printRunPageSetupDialogFunction *gi.Function
var printRunPageSetupDialogFunction_Once sync.Once

func printRunPageSetupDialogFunction_Set() error {
	var err error
	printRunPageSetupDialogFunction_Once.Do(func() {
		printRunPageSetupDialogFunction, err = gi.FunctionInvokerNew("Gtk", "print_run_page_setup_dialog")
	})
	return err
}

// PrintRunPageSetupDialog is a representation of the C type gtk_print_run_page_setup_dialog.
func PrintRunPageSetupDialog(parent *Window, pageSetup *PageSetup, settings *PrintSettings) *PageSetup {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(parent.Native())
	inArgs[1].SetPointer(pageSetup.Native())
	inArgs[2].SetPointer(settings.Native())

	var ret gi.Argument

	err := printRunPageSetupDialogFunction_Set()
	if err == nil {
		ret = printRunPageSetupDialogFunction.Invoke(inArgs[:], nil)
	}

	retGo := PageSetupNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_print_run_page_setup_dialog_async' : parameter 'done_cb' of type 'PageSetupDoneFunc' not supported

// UNSUPPORTED : C value 'gtk_propagate_event' : parameter 'event' of type 'Gdk.Event' not supported

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

// UNSUPPORTED : C value 'gtk_rc_find_pixmap_in_path' : parameter 'scanner' of type 'GLib.Scanner' not supported

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

var rcGetStyleFunction *gi.Function
var rcGetStyleFunction_Once sync.Once

func rcGetStyleFunction_Set() error {
	var err error
	rcGetStyleFunction_Once.Do(func() {
		rcGetStyleFunction, err = gi.FunctionInvokerNew("Gtk", "rc_get_style")
	})
	return err
}

// RcGetStyle is a representation of the C type gtk_rc_get_style.
func RcGetStyle(widget *Widget) *Style {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(widget.Native())

	var ret gi.Argument

	err := rcGetStyleFunction_Set()
	if err == nil {
		ret = rcGetStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := StyleNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_rc_get_style_by_paths' : parameter 'type' of type 'GType' not supported

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

// UNSUPPORTED : C value 'gtk_rc_property_parse_border' : parameter 'gstring' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_color' : parameter 'gstring' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_enum' : parameter 'gstring' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_flags' : parameter 'gstring' of type 'GLib.String' not supported

// UNSUPPORTED : C value 'gtk_rc_property_parse_requisition' : parameter 'gstring' of type 'GLib.String' not supported

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

var rcReparseAllForSettingsFunction *gi.Function
var rcReparseAllForSettingsFunction_Once sync.Once

func rcReparseAllForSettingsFunction_Set() error {
	var err error
	rcReparseAllForSettingsFunction_Once.Do(func() {
		rcReparseAllForSettingsFunction, err = gi.FunctionInvokerNew("Gtk", "rc_reparse_all_for_settings")
	})
	return err
}

// RcReparseAllForSettings is a representation of the C type gtk_rc_reparse_all_for_settings.
func RcReparseAllForSettings(settings *Settings, forceLoad bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(settings.Native())
	inArgs[1].SetBoolean(forceLoad)

	var ret gi.Argument

	err := rcReparseAllForSettingsFunction_Set()
	if err == nil {
		ret = rcReparseAllForSettingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var rcResetStylesFunction *gi.Function
var rcResetStylesFunction_Once sync.Once

func rcResetStylesFunction_Set() error {
	var err error
	rcResetStylesFunction_Once.Do(func() {
		rcResetStylesFunction, err = gi.FunctionInvokerNew("Gtk", "rc_reset_styles")
	})
	return err
}

// RcResetStyles is a representation of the C type gtk_rc_reset_styles.
func RcResetStyles(settings *Settings) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(settings.Native())

	err := rcResetStylesFunction_Set()
	if err == nil {
		rcResetStylesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_rc_scanner_new' : return type 'GLib.Scanner' not supported

// UNSUPPORTED : C value 'gtk_rc_set_default_files' : parameter 'filenames' of type 'nil' not supported

var recentChooserErrorQuarkFunction *gi.Function
var recentChooserErrorQuarkFunction_Once sync.Once

func recentChooserErrorQuarkFunction_Set() error {
	var err error
	recentChooserErrorQuarkFunction_Once.Do(func() {
		recentChooserErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "recent_chooser_error_quark")
	})
	return err
}

// RecentChooserErrorQuark is a representation of the C type gtk_recent_chooser_error_quark.
func RecentChooserErrorQuark() glib.Quark {

	var ret gi.Argument

	err := recentChooserErrorQuarkFunction_Set()
	if err == nil {
		ret = recentChooserErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var recentManagerErrorQuarkFunction *gi.Function
var recentManagerErrorQuarkFunction_Once sync.Once

func recentManagerErrorQuarkFunction_Set() error {
	var err error
	recentManagerErrorQuarkFunction_Once.Do(func() {
		recentManagerErrorQuarkFunction, err = gi.FunctionInvokerNew("Gtk", "recent_manager_error_quark")
	})
	return err
}

// RecentManagerErrorQuark is a representation of the C type gtk_recent_manager_error_quark.
func RecentManagerErrorQuark() glib.Quark {

	var ret gi.Argument

	err := recentManagerErrorQuarkFunction_Set()
	if err == nil {
		ret = recentManagerErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'gtk_render_activity' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_arrow' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_background' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_background_get_clip' : parameter 'out_clip' of type 'Gdk.Rectangle' not supported

// UNSUPPORTED : C value 'gtk_render_check' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_expander' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_extension' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_focus' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_frame' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_frame_gap' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_handle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_icon' : parameter 'cr' of type 'cairo.Context' not supported

var renderIconPixbufFunction *gi.Function
var renderIconPixbufFunction_Once sync.Once

func renderIconPixbufFunction_Set() error {
	var err error
	renderIconPixbufFunction_Once.Do(func() {
		renderIconPixbufFunction, err = gi.FunctionInvokerNew("Gtk", "render_icon_pixbuf")
	})
	return err
}

// RenderIconPixbuf is a representation of the C type gtk_render_icon_pixbuf.
func RenderIconPixbuf(context *StyleContext, source *IconSource, size IconSize) *gdkpixbuf.Pixbuf {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(context.Native())
	inArgs[1].SetPointer(source.Native())
	inArgs[2].SetInt32(int32(size))

	var ret gi.Argument

	err := renderIconPixbufFunction_Set()
	if err == nil {
		ret = renderIconPixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := gdkpixbuf.PixbufNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_render_icon_surface' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_insertion_cursor' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_layout' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_line' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_option' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gtk_render_slider' : parameter 'cr' of type 'cairo.Context' not supported

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

// UNSUPPORTED : C value 'gtk_selection_add_target' : parameter 'selection' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_add_targets' : parameter 'selection' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_clear_targets' : parameter 'selection' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_convert' : parameter 'selection' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_owner_set' : parameter 'selection' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_owner_set_for_display' : parameter 'selection' of type 'Gdk.Atom' not supported

var selectionRemoveAllFunction *gi.Function
var selectionRemoveAllFunction_Once sync.Once

func selectionRemoveAllFunction_Set() error {
	var err error
	selectionRemoveAllFunction_Once.Do(func() {
		selectionRemoveAllFunction, err = gi.FunctionInvokerNew("Gtk", "selection_remove_all")
	})
	return err
}

// SelectionRemoveAll is a representation of the C type gtk_selection_remove_all.
func SelectionRemoveAll(widget *Widget) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(widget.Native())

	err := selectionRemoveAllFunction_Set()
	if err == nil {
		selectionRemoveAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

// UNSUPPORTED : C value 'gtk_show_about_dialog' : parameter '...' of type 'nil' not supported

var showUriFunction *gi.Function
var showUriFunction_Once sync.Once

func showUriFunction_Set() error {
	var err error
	showUriFunction_Once.Do(func() {
		showUriFunction, err = gi.FunctionInvokerNew("Gtk", "show_uri")
	})
	return err
}

// ShowUri is a representation of the C type gtk_show_uri.
func ShowUri(screen *gdk.Screen, uri string, timestamp uint32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(screen.Native())
	inArgs[1].SetString(uri)
	inArgs[2].SetUint32(timestamp)

	var ret gi.Argument

	err := showUriFunction_Set()
	if err == nil {
		ret = showUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var showUriOnWindowFunction *gi.Function
var showUriOnWindowFunction_Once sync.Once

func showUriOnWindowFunction_Set() error {
	var err error
	showUriOnWindowFunction_Once.Do(func() {
		showUriOnWindowFunction, err = gi.FunctionInvokerNew("Gtk", "show_uri_on_window")
	})
	return err
}

// ShowUriOnWindow is a representation of the C type gtk_show_uri_on_window.
func ShowUriOnWindow(parent *Window, uri string, timestamp uint32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(parent.Native())
	inArgs[1].SetString(uri)
	inArgs[2].SetUint32(timestamp)

	var ret gi.Argument

	err := showUriOnWindowFunction_Set()
	if err == nil {
		ret = showUriOnWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_stock_add' : parameter 'items' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_stock_add_static' : parameter 'items' of type 'nil' not supported

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
	out0 := StockItemNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

// UNSUPPORTED : C value 'gtk_stock_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

// UNSUPPORTED : C value 'gtk_target_table_free' : parameter 'targets' of type 'nil' not supported

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
	inArgs[0].SetPointer(list.Native())

	var outArgs [1]gi.Argument

	err := targetTableNewFromListFunction_Set()
	if err == nil {
		targetTableNewFromListFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_targets_include_image' : parameter 'targets' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_targets_include_rich_text' : parameter 'targets' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_targets_include_text' : parameter 'targets' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_targets_include_uri' : parameter 'targets' of type 'nil' not supported

var testCreateSimpleWindowFunction *gi.Function
var testCreateSimpleWindowFunction_Once sync.Once

func testCreateSimpleWindowFunction_Set() error {
	var err error
	testCreateSimpleWindowFunction_Once.Do(func() {
		testCreateSimpleWindowFunction, err = gi.FunctionInvokerNew("Gtk", "test_create_simple_window")
	})
	return err
}

// TestCreateSimpleWindow is a representation of the C type gtk_test_create_simple_window.
func TestCreateSimpleWindow(windowTitle string, dialogText string) *Widget {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(windowTitle)
	inArgs[1].SetString(dialogText)

	var ret gi.Argument

	err := testCreateSimpleWindowFunction_Set()
	if err == nil {
		ret = testCreateSimpleWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := WidgetNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_test_create_widget' : parameter 'widget_type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_test_display_button_window' : parameter '...' of type 'nil' not supported

var testFindLabelFunction *gi.Function
var testFindLabelFunction_Once sync.Once

func testFindLabelFunction_Set() error {
	var err error
	testFindLabelFunction_Once.Do(func() {
		testFindLabelFunction, err = gi.FunctionInvokerNew("Gtk", "test_find_label")
	})
	return err
}

// TestFindLabel is a representation of the C type gtk_test_find_label.
func TestFindLabel(widget *Widget, labelPattern string) *Widget {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(widget.Native())
	inArgs[1].SetString(labelPattern)

	var ret gi.Argument

	err := testFindLabelFunction_Set()
	if err == nil {
		ret = testFindLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := WidgetNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_test_find_sibling' : parameter 'widget_type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_test_find_widget' : parameter 'widget_type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_test_init' : parameter 'argvp' of type 'nil' not supported

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

var testSliderGetValueFunction *gi.Function
var testSliderGetValueFunction_Once sync.Once

func testSliderGetValueFunction_Set() error {
	var err error
	testSliderGetValueFunction_Once.Do(func() {
		testSliderGetValueFunction, err = gi.FunctionInvokerNew("Gtk", "test_slider_get_value")
	})
	return err
}

// TestSliderGetValue is a representation of the C type gtk_test_slider_get_value.
func TestSliderGetValue(widget *Widget) float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(widget.Native())

	var ret gi.Argument

	err := testSliderGetValueFunction_Set()
	if err == nil {
		ret = testSliderGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var testSliderSetPercFunction *gi.Function
var testSliderSetPercFunction_Once sync.Once

func testSliderSetPercFunction_Set() error {
	var err error
	testSliderSetPercFunction_Once.Do(func() {
		testSliderSetPercFunction, err = gi.FunctionInvokerNew("Gtk", "test_slider_set_perc")
	})
	return err
}

// TestSliderSetPerc is a representation of the C type gtk_test_slider_set_perc.
func TestSliderSetPerc(widget *Widget, percentage float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(widget.Native())
	inArgs[1].SetFloat64(percentage)

	err := testSliderSetPercFunction_Set()
	if err == nil {
		testSliderSetPercFunction.Invoke(inArgs[:], nil)
	}

	return
}

var testSpinButtonClickFunction *gi.Function
var testSpinButtonClickFunction_Once sync.Once

func testSpinButtonClickFunction_Set() error {
	var err error
	testSpinButtonClickFunction_Once.Do(func() {
		testSpinButtonClickFunction, err = gi.FunctionInvokerNew("Gtk", "test_spin_button_click")
	})
	return err
}

// TestSpinButtonClick is a representation of the C type gtk_test_spin_button_click.
func TestSpinButtonClick(spinner *SpinButton, button uint32, upwards bool) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(spinner.Native())
	inArgs[1].SetUint32(button)
	inArgs[2].SetBoolean(upwards)

	var ret gi.Argument

	err := testSpinButtonClickFunction_Set()
	if err == nil {
		ret = testSpinButtonClickFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var testTextGetFunction *gi.Function
var testTextGetFunction_Once sync.Once

func testTextGetFunction_Set() error {
	var err error
	testTextGetFunction_Once.Do(func() {
		testTextGetFunction, err = gi.FunctionInvokerNew("Gtk", "test_text_get")
	})
	return err
}

// TestTextGet is a representation of the C type gtk_test_text_get.
func TestTextGet(widget *Widget) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(widget.Native())

	var ret gi.Argument

	err := testTextGetFunction_Set()
	if err == nil {
		ret = testTextGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var testTextSetFunction *gi.Function
var testTextSetFunction_Once sync.Once

func testTextSetFunction_Set() error {
	var err error
	testTextSetFunction_Once.Do(func() {
		testTextSetFunction, err = gi.FunctionInvokerNew("Gtk", "test_text_set")
	})
	return err
}

// TestTextSet is a representation of the C type gtk_test_text_set.
func TestTextSet(widget *Widget, string_ string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(widget.Native())
	inArgs[1].SetString(string_)

	err := testTextSetFunction_Set()
	if err == nil {
		testTextSetFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_test_widget_click' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_test_widget_send_key' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

var testWidgetWaitForDrawFunction *gi.Function
var testWidgetWaitForDrawFunction_Once sync.Once

func testWidgetWaitForDrawFunction_Set() error {
	var err error
	testWidgetWaitForDrawFunction_Once.Do(func() {
		testWidgetWaitForDrawFunction, err = gi.FunctionInvokerNew("Gtk", "test_widget_wait_for_draw")
	})
	return err
}

// TestWidgetWaitForDraw is a representation of the C type gtk_test_widget_wait_for_draw.
func TestWidgetWaitForDraw(widget *Widget) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(widget.Native())

	err := testWidgetWaitForDrawFunction_Set()
	if err == nil {
		testWidgetWaitForDrawFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_tree_get_row_drag_data' : parameter 'tree_model' of type 'TreeModel' not supported

var treeRowReferenceDeletedFunction *gi.Function
var treeRowReferenceDeletedFunction_Once sync.Once

func treeRowReferenceDeletedFunction_Set() error {
	var err error
	treeRowReferenceDeletedFunction_Once.Do(func() {
		treeRowReferenceDeletedFunction, err = gi.FunctionInvokerNew("Gtk", "tree_row_reference_deleted")
	})
	return err
}

// TreeRowReferenceDeleted is a representation of the C type gtk_tree_row_reference_deleted.
func TreeRowReferenceDeleted(proxy *gobject.Object, path *TreePath) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(proxy.Native())
	inArgs[1].SetPointer(path.Native())

	err := treeRowReferenceDeletedFunction_Set()
	if err == nil {
		treeRowReferenceDeletedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeRowReferenceInsertedFunction *gi.Function
var treeRowReferenceInsertedFunction_Once sync.Once

func treeRowReferenceInsertedFunction_Set() error {
	var err error
	treeRowReferenceInsertedFunction_Once.Do(func() {
		treeRowReferenceInsertedFunction, err = gi.FunctionInvokerNew("Gtk", "tree_row_reference_inserted")
	})
	return err
}

// TreeRowReferenceInserted is a representation of the C type gtk_tree_row_reference_inserted.
func TreeRowReferenceInserted(proxy *gobject.Object, path *TreePath) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(proxy.Native())
	inArgs[1].SetPointer(path.Native())

	err := treeRowReferenceInsertedFunction_Set()
	if err == nil {
		treeRowReferenceInsertedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_reordered' : parameter 'new_order' of type 'nil' not supported

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
