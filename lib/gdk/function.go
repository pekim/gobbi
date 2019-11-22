// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'gdk_add_option_entries_libgtk_only' : parameter 'group' of type 'GLib.OptionGroup' not supported

// UNSUPPORTED : C value 'gdk_atom_intern' : parameter 'only_if_exists' of type 'gboolean' not supported

var atomInternStaticStringFunction *gi.Function
var atomInternStaticStringFunction_Once sync.Once

func atomInternStaticStringFunction_Set() error {
	var err error
	atomInternStaticStringFunction_Once.Do(func() {
		atomInternStaticStringFunction, err = gi.FunctionInvokerNew("Gdk", "atom_intern_static_string")
	})
	return err
}

// AtomInternStaticString is a representation of the C type gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) *Atom {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(atomName)

	var ret gi.Argument

	err := atomInternStaticStringFunction_Set()
	if err == nil {
		ret = atomInternStaticStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Atom{native: ret.Pointer()}

	return retGo
}

var beepFunction *gi.Function
var beepFunction_Once sync.Once

func beepFunction_Set() error {
	var err error
	beepFunction_Once.Do(func() {
		beepFunction, err = gi.FunctionInvokerNew("Gdk", "beep")
	})
	return err
}

// Beep is a representation of the C type gdk_beep.
func Beep() {

	err := beepFunction_Set()
	if err == nil {
		beepFunction.Invoke(nil, nil)
	}

}

// UNSUPPORTED : C value 'gdk_cairo_create' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_cairo_draw_from_gl' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_get_clip_rectangle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_get_drawing_context' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_rectangle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_region' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_region_create_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_color' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_pixbuf' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_rgba' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_window' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_surface_create_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_color_parse' : parameter 'color' of type 'Color' not supported

var disableMultideviceFunction *gi.Function
var disableMultideviceFunction_Once sync.Once

func disableMultideviceFunction_Set() error {
	var err error
	disableMultideviceFunction_Once.Do(func() {
		disableMultideviceFunction, err = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	})
	return err
}

// DisableMultidevice is a representation of the C type gdk_disable_multidevice.
func DisableMultidevice() {

	err := disableMultideviceFunction_Set()
	if err == nil {
		disableMultideviceFunction.Invoke(nil, nil)
	}

}

// UNSUPPORTED : C value 'gdk_drag_abort' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_begin' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_drag_begin_for_device' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_drag_begin_from_point' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_drag_drop' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_drop_done' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_drop_succeeded' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_find_window_for_screen' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_get_selection' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_motion' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drag_status' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drop_finish' : parameter 'context' of type 'DragContext' not supported

// UNSUPPORTED : C value 'gdk_drop_reply' : parameter 'context' of type 'DragContext' not supported

var errorTrapPopFunction *gi.Function
var errorTrapPopFunction_Once sync.Once

func errorTrapPopFunction_Set() error {
	var err error
	errorTrapPopFunction_Once.Do(func() {
		errorTrapPopFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	})
	return err
}

// ErrorTrapPop is a representation of the C type gdk_error_trap_pop.
func ErrorTrapPop() int32 {

	var ret gi.Argument

	err := errorTrapPopFunction_Set()
	if err == nil {
		ret = errorTrapPopFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
}

var errorTrapPopIgnoredFunction *gi.Function
var errorTrapPopIgnoredFunction_Once sync.Once

func errorTrapPopIgnoredFunction_Set() error {
	var err error
	errorTrapPopIgnoredFunction_Once.Do(func() {
		errorTrapPopIgnoredFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	})
	return err
}

// ErrorTrapPopIgnored is a representation of the C type gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {

	err := errorTrapPopIgnoredFunction_Set()
	if err == nil {
		errorTrapPopIgnoredFunction.Invoke(nil, nil)
	}

}

var errorTrapPushFunction *gi.Function
var errorTrapPushFunction_Once sync.Once

func errorTrapPushFunction_Set() error {
	var err error
	errorTrapPushFunction_Once.Do(func() {
		errorTrapPushFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	})
	return err
}

// ErrorTrapPush is a representation of the C type gdk_error_trap_push.
func ErrorTrapPush() {

	err := errorTrapPushFunction_Set()
	if err == nil {
		errorTrapPushFunction.Invoke(nil, nil)
	}

}

// UNSUPPORTED : C value 'gdk_event_get' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_handler_set' : parameter 'func' of type 'EventFunc' not supported

// UNSUPPORTED : C value 'gdk_event_peek' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_request_motions' : parameter 'event' of type 'EventMotion' not supported

// UNSUPPORTED : C value 'gdk_events_get_angle' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_center' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_distance' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_pending' : return type 'gboolean' not supported

var flushFunction *gi.Function
var flushFunction_Once sync.Once

func flushFunction_Set() error {
	var err error
	flushFunction_Once.Do(func() {
		flushFunction, err = gi.FunctionInvokerNew("Gdk", "flush")
	})
	return err
}

// Flush is a representation of the C type gdk_flush.
func Flush() {

	err := flushFunction_Set()
	if err == nil {
		flushFunction.Invoke(nil, nil)
	}

}

// UNSUPPORTED : C value 'gdk_get_default_root_window' : return type 'Window' not supported

var getDisplayFunction *gi.Function
var getDisplayFunction_Once sync.Once

func getDisplayFunction_Set() error {
	var err error
	getDisplayFunction_Once.Do(func() {
		getDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_get_display.
func GetDisplay() string {

	var ret gi.Argument

	err := getDisplayFunction_Set()
	if err == nil {
		ret = getDisplayFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var getDisplayArgNameFunction *gi.Function
var getDisplayArgNameFunction_Once sync.Once

func getDisplayArgNameFunction_Set() error {
	var err error
	getDisplayArgNameFunction_Once.Do(func() {
		getDisplayArgNameFunction, err = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	})
	return err
}

// GetDisplayArgName is a representation of the C type gdk_get_display_arg_name.
func GetDisplayArgName() string {

	var ret gi.Argument

	err := getDisplayArgNameFunction_Set()
	if err == nil {
		ret = getDisplayArgNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getProgramClassFunction *gi.Function
var getProgramClassFunction_Once sync.Once

func getProgramClassFunction_Set() error {
	var err error
	getProgramClassFunction_Once.Do(func() {
		getProgramClassFunction, err = gi.FunctionInvokerNew("Gdk", "get_program_class")
	})
	return err
}

// GetProgramClass is a representation of the C type gdk_get_program_class.
func GetProgramClass() string {

	var ret gi.Argument

	err := getProgramClassFunction_Set()
	if err == nil {
		ret = getProgramClassFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gdk_get_show_events' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_gl_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gdk_init' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gdk_init_check' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gdk_keyboard_grab' : parameter 'window' of type 'Window' not supported

var keyboardUngrabFunction *gi.Function
var keyboardUngrabFunction_Once sync.Once

func keyboardUngrabFunction_Set() error {
	var err error
	keyboardUngrabFunction_Once.Do(func() {
		keyboardUngrabFunction, err = gi.FunctionInvokerNew("Gdk", "keyboard_ungrab")
	})
	return err
}

// KeyboardUngrab is a representation of the C type gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	err := keyboardUngrabFunction_Set()
	if err == nil {
		keyboardUngrabFunction.Invoke(inArgs[:], nil)
	}

}

var keyvalConvertCaseFunction *gi.Function
var keyvalConvertCaseFunction_Once sync.Once

func keyvalConvertCaseFunction_Set() error {
	var err error
	keyvalConvertCaseFunction_Once.Do(func() {
		keyvalConvertCaseFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_convert_case")
	})
	return err
}

// KeyvalConvertCase is a representation of the C type gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(symbol)

	var outArgs [2]gi.Argument

	err := keyvalConvertCaseFunction_Set()
	if err == nil {
		keyvalConvertCaseFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Uint32()

	return out0, out1
}

var keyvalFromNameFunction *gi.Function
var keyvalFromNameFunction_Once sync.Once

func keyvalFromNameFunction_Set() error {
	var err error
	keyvalFromNameFunction_Once.Do(func() {
		keyvalFromNameFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_from_name")
	})
	return err
}

// KeyvalFromName is a representation of the C type gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(keyvalName)

	var ret gi.Argument

	err := keyvalFromNameFunction_Set()
	if err == nil {
		ret = keyvalFromNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_keyval_is_lower' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_keyval_is_upper' : return type 'gboolean' not supported

var keyvalNameFunction *gi.Function
var keyvalNameFunction_Once sync.Once

func keyvalNameFunction_Set() error {
	var err error
	keyvalNameFunction_Once.Do(func() {
		keyvalNameFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_name")
	})
	return err
}

// KeyvalName is a representation of the C type gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalNameFunction_Set()
	if err == nil {
		ret = keyvalNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var keyvalToLowerFunction *gi.Function
var keyvalToLowerFunction_Once sync.Once

func keyvalToLowerFunction_Set() error {
	var err error
	keyvalToLowerFunction_Once.Do(func() {
		keyvalToLowerFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_lower")
	})
	return err
}

// KeyvalToLower is a representation of the C type gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToLowerFunction_Set()
	if err == nil {
		ret = keyvalToLowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalToUnicodeFunction *gi.Function
var keyvalToUnicodeFunction_Once sync.Once

func keyvalToUnicodeFunction_Set() error {
	var err error
	keyvalToUnicodeFunction_Once.Do(func() {
		keyvalToUnicodeFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_unicode")
	})
	return err
}

// KeyvalToUnicode is a representation of the C type gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToUnicodeFunction_Set()
	if err == nil {
		ret = keyvalToUnicodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalToUpperFunction *gi.Function
var keyvalToUpperFunction_Once sync.Once

func keyvalToUpperFunction_Set() error {
	var err error
	keyvalToUpperFunction_Once.Do(func() {
		keyvalToUpperFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_upper")
	})
	return err
}

// KeyvalToUpper is a representation of the C type gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToUpperFunction_Set()
	if err == nil {
		ret = keyvalToUpperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_list_visuals' : return type 'GLib.List' not supported

var notifyStartupCompleteFunction *gi.Function
var notifyStartupCompleteFunction_Once sync.Once

func notifyStartupCompleteFunction_Set() error {
	var err error
	notifyStartupCompleteFunction_Once.Do(func() {
		notifyStartupCompleteFunction, err = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	})
	return err
}

// NotifyStartupComplete is a representation of the C type gdk_notify_startup_complete.
func NotifyStartupComplete() {

	err := notifyStartupCompleteFunction_Set()
	if err == nil {
		notifyStartupCompleteFunction.Invoke(nil, nil)
	}

}

var notifyStartupCompleteWithIdFunction *gi.Function
var notifyStartupCompleteWithIdFunction_Once sync.Once

func notifyStartupCompleteWithIdFunction_Set() error {
	var err error
	notifyStartupCompleteWithIdFunction_Once.Do(func() {
		notifyStartupCompleteWithIdFunction, err = gi.FunctionInvokerNew("Gdk", "notify_startup_complete_with_id")
	})
	return err
}

// NotifyStartupCompleteWithId is a representation of the C type gdk_notify_startup_complete_with_id.
func NotifyStartupCompleteWithId(startupId string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(startupId)

	err := notifyStartupCompleteWithIdFunction_Set()
	if err == nil {
		notifyStartupCompleteWithIdFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'gdk_offscreen_window_get_embedder' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_offscreen_window_get_surface' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_offscreen_window_set_embedder' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_screen' : parameter 'screen' of type 'Screen' not supported

// UNSUPPORTED : C value 'gdk_pango_layout_get_clip_region' : parameter 'layout' of type 'Pango.Layout' not supported

// UNSUPPORTED : C value 'gdk_pango_layout_line_get_clip_region' : parameter 'line' of type 'Pango.LayoutLine' not supported

// UNSUPPORTED : C value 'gdk_parse_args' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_window' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_pointer_grab' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_pointer_is_grabbed' : return type 'gboolean' not supported

var pointerUngrabFunction *gi.Function
var pointerUngrabFunction_Once sync.Once

func pointerUngrabFunction_Set() error {
	var err error
	pointerUngrabFunction_Once.Do(func() {
		pointerUngrabFunction, err = gi.FunctionInvokerNew("Gdk", "pointer_ungrab")
	})
	return err
}

// PointerUngrab is a representation of the C type gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	err := pointerUngrabFunction_Set()
	if err == nil {
		pointerUngrabFunction.Invoke(inArgs[:], nil)
	}

}

var preParseLibgtkOnlyFunction *gi.Function
var preParseLibgtkOnlyFunction_Once sync.Once

func preParseLibgtkOnlyFunction_Set() error {
	var err error
	preParseLibgtkOnlyFunction_Once.Do(func() {
		preParseLibgtkOnlyFunction, err = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	})
	return err
}

// PreParseLibgtkOnly is a representation of the C type gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {

	err := preParseLibgtkOnlyFunction_Set()
	if err == nil {
		preParseLibgtkOnlyFunction.Invoke(nil, nil)
	}

}

// UNSUPPORTED : C value 'gdk_property_change' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_property_delete' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_property_get' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_query_depths' : parameter 'depths' has no type

// UNSUPPORTED : C value 'gdk_query_visual_types' : parameter 'visual_types' has no type

// UNSUPPORTED : C value 'gdk_selection_convert' : parameter 'requestor' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_get' : parameter 'selection' of type 'Atom' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_get_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_set' : parameter 'owner' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_owner_set_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_selection_property_get' : parameter 'requestor' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_send_notify' : parameter 'requestor' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_selection_send_notify_for_display' : parameter 'display' of type 'Display' not supported

var setAllowedBackendsFunction *gi.Function
var setAllowedBackendsFunction_Once sync.Once

func setAllowedBackendsFunction_Set() error {
	var err error
	setAllowedBackendsFunction_Once.Do(func() {
		setAllowedBackendsFunction, err = gi.FunctionInvokerNew("Gdk", "set_allowed_backends")
	})
	return err
}

// SetAllowedBackends is a representation of the C type gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(backends)

	err := setAllowedBackendsFunction_Set()
	if err == nil {
		setAllowedBackendsFunction.Invoke(inArgs[:], nil)
	}

}

var setDoubleClickTimeFunction *gi.Function
var setDoubleClickTimeFunction_Once sync.Once

func setDoubleClickTimeFunction_Set() error {
	var err error
	setDoubleClickTimeFunction_Once.Do(func() {
		setDoubleClickTimeFunction, err = gi.FunctionInvokerNew("Gdk", "set_double_click_time")
	})
	return err
}

// SetDoubleClickTime is a representation of the C type gdk_set_double_click_time.
func SetDoubleClickTime(msec uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(msec)

	err := setDoubleClickTimeFunction_Set()
	if err == nil {
		setDoubleClickTimeFunction.Invoke(inArgs[:], nil)
	}

}

var setProgramClassFunction *gi.Function
var setProgramClassFunction_Once sync.Once

func setProgramClassFunction_Set() error {
	var err error
	setProgramClassFunction_Once.Do(func() {
		setProgramClassFunction, err = gi.FunctionInvokerNew("Gdk", "set_program_class")
	})
	return err
}

// SetProgramClass is a representation of the C type gdk_set_program_class.
func SetProgramClass(programClass string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(programClass)

	err := setProgramClassFunction_Set()
	if err == nil {
		setProgramClassFunction.Invoke(inArgs[:], nil)
	}

}

// UNSUPPORTED : C value 'gdk_set_show_events' : parameter 'show_events' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_setting_get' : parameter 'value' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'gdk_synthesize_window_state' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_test_render_sync' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_test_simulate_button' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_test_simulate_key' : parameter 'window' of type 'Window' not supported

// UNSUPPORTED : C value 'gdk_text_property_to_utf8_list_for_display' : parameter 'display' of type 'Display' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

var threadsEnterFunction *gi.Function
var threadsEnterFunction_Once sync.Once

func threadsEnterFunction_Set() error {
	var err error
	threadsEnterFunction_Once.Do(func() {
		threadsEnterFunction, err = gi.FunctionInvokerNew("Gdk", "threads_enter")
	})
	return err
}

// ThreadsEnter is a representation of the C type gdk_threads_enter.
func ThreadsEnter() {

	err := threadsEnterFunction_Set()
	if err == nil {
		threadsEnterFunction.Invoke(nil, nil)
	}

}

var threadsInitFunction *gi.Function
var threadsInitFunction_Once sync.Once

func threadsInitFunction_Set() error {
	var err error
	threadsInitFunction_Once.Do(func() {
		threadsInitFunction, err = gi.FunctionInvokerNew("Gdk", "threads_init")
	})
	return err
}

// ThreadsInit is a representation of the C type gdk_threads_init.
func ThreadsInit() {

	err := threadsInitFunction_Set()
	if err == nil {
		threadsInitFunction.Invoke(nil, nil)
	}

}

var threadsLeaveFunction *gi.Function
var threadsLeaveFunction_Once sync.Once

func threadsLeaveFunction_Set() error {
	var err error
	threadsLeaveFunction_Once.Do(func() {
		threadsLeaveFunction, err = gi.FunctionInvokerNew("Gdk", "threads_leave")
	})
	return err
}

// ThreadsLeave is a representation of the C type gdk_threads_leave.
func ThreadsLeave() {

	err := threadsLeaveFunction_Set()
	if err == nil {
		threadsLeaveFunction.Invoke(nil, nil)
	}

}

// UNSUPPORTED : C value 'gdk_threads_set_lock_functions' : parameter 'enter_fn' of type 'GObject.Callback' not supported

var unicodeToKeyvalFunction *gi.Function
var unicodeToKeyvalFunction_Once sync.Once

func unicodeToKeyvalFunction_Set() error {
	var err error
	unicodeToKeyvalFunction_Once.Do(func() {
		unicodeToKeyvalFunction, err = gi.FunctionInvokerNew("Gdk", "unicode_to_keyval")
	})
	return err
}

// UnicodeToKeyval is a representation of the C type gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(wc)

	var ret gi.Argument

	err := unicodeToKeyvalFunction_Set()
	if err == nil {
		ret = unicodeToKeyvalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var utf8ToStringTargetFunction *gi.Function
var utf8ToStringTargetFunction_Once sync.Once

func utf8ToStringTargetFunction_Set() error {
	var err error
	utf8ToStringTargetFunction_Once.Do(func() {
		utf8ToStringTargetFunction, err = gi.FunctionInvokerNew("Gdk", "utf8_to_string_target")
	})
	return err
}

// Utf8ToStringTarget is a representation of the C type gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := utf8ToStringTargetFunction_Set()
	if err == nil {
		ret = utf8ToStringTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}
