// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

// UNSUPPORTED : C value 'gdk_add_option_entries_libgtk_only' : parameter 'group' of type 'GLib.OptionGroup' not supported

// UNSUPPORTED : C value 'gdk_atom_intern' : parameter 'only_if_exists' of type 'gboolean' not supported

var AtomInternStaticStringFunction *gi.Function
var AtomInternStaticStringFunctionOnce sync.Once

func AtomInternStaticStringFunctionSet() {
	AtomInternStaticStringFunctionOnce.Do(func() {
		AtomInternStaticStringFunction = gi.FunctionInvokerNew("Gdk", "atom_intern_static_string")
	})
}

var atomInternStaticStringInvoker *gi.Function

// AtomInternStaticString is a representation of the C type gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) *Atom {
	if atomInternStaticStringInvoker == nil {
		atomInternStaticStringInvoker = gi.FunctionInvokerNew("Gdk", "atom_intern_static_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(atomName)

	ret := atomInternStaticStringInvoker.Invoke(inArgs[:], nil)

	retGo := &Atom{native: ret.Pointer()}

	return retGo
}

var BeepFunction *gi.Function
var BeepFunctionOnce sync.Once

func BeepFunctionSet() {
	BeepFunctionOnce.Do(func() {
		BeepFunction = gi.FunctionInvokerNew("Gdk", "beep")
	})
}

var beepInvoker *gi.Function

// Beep is a representation of the C type gdk_beep.
func Beep() {
	if beepInvoker == nil {
		beepInvoker = gi.FunctionInvokerNew("Gdk", "beep")
	}

	beepInvoker.Invoke(nil, nil)

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

var DisableMultideviceFunction *gi.Function
var DisableMultideviceFunctionOnce sync.Once

func DisableMultideviceFunctionSet() {
	DisableMultideviceFunctionOnce.Do(func() {
		DisableMultideviceFunction = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	})
}

var disableMultideviceInvoker *gi.Function

// DisableMultidevice is a representation of the C type gdk_disable_multidevice.
func DisableMultidevice() {
	if disableMultideviceInvoker == nil {
		disableMultideviceInvoker = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	}

	disableMultideviceInvoker.Invoke(nil, nil)

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

var ErrorTrapPopFunction *gi.Function
var ErrorTrapPopFunctionOnce sync.Once

func ErrorTrapPopFunctionSet() {
	ErrorTrapPopFunctionOnce.Do(func() {
		ErrorTrapPopFunction = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	})
}

var errorTrapPopInvoker *gi.Function

// ErrorTrapPop is a representation of the C type gdk_error_trap_pop.
func ErrorTrapPop() int32 {
	if errorTrapPopInvoker == nil {
		errorTrapPopInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	}

	ret := errorTrapPopInvoker.Invoke(nil, nil)

	retGo := ret.Int32()

	return retGo
}

var ErrorTrapPopIgnoredFunction *gi.Function
var ErrorTrapPopIgnoredFunctionOnce sync.Once

func ErrorTrapPopIgnoredFunctionSet() {
	ErrorTrapPopIgnoredFunctionOnce.Do(func() {
		ErrorTrapPopIgnoredFunction = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	})
}

var errorTrapPopIgnoredInvoker *gi.Function

// ErrorTrapPopIgnored is a representation of the C type gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {
	if errorTrapPopIgnoredInvoker == nil {
		errorTrapPopIgnoredInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	}

	errorTrapPopIgnoredInvoker.Invoke(nil, nil)

}

var ErrorTrapPushFunction *gi.Function
var ErrorTrapPushFunctionOnce sync.Once

func ErrorTrapPushFunctionSet() {
	ErrorTrapPushFunctionOnce.Do(func() {
		ErrorTrapPushFunction = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	})
}

var errorTrapPushInvoker *gi.Function

// ErrorTrapPush is a representation of the C type gdk_error_trap_push.
func ErrorTrapPush() {
	if errorTrapPushInvoker == nil {
		errorTrapPushInvoker = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	}

	errorTrapPushInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gdk_event_get' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_handler_set' : parameter 'func' of type 'EventFunc' not supported

// UNSUPPORTED : C value 'gdk_event_peek' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_request_motions' : parameter 'event' of type 'EventMotion' not supported

// UNSUPPORTED : C value 'gdk_events_get_angle' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_center' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_distance' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_pending' : return type 'gboolean' not supported

var FlushFunction *gi.Function
var FlushFunctionOnce sync.Once

func FlushFunctionSet() {
	FlushFunctionOnce.Do(func() {
		FlushFunction = gi.FunctionInvokerNew("Gdk", "flush")
	})
}

var flushInvoker *gi.Function

// Flush is a representation of the C type gdk_flush.
func Flush() {
	if flushInvoker == nil {
		flushInvoker = gi.FunctionInvokerNew("Gdk", "flush")
	}

	flushInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gdk_get_default_root_window' : return type 'Window' not supported

var GetDisplayFunction *gi.Function
var GetDisplayFunctionOnce sync.Once

func GetDisplayFunctionSet() {
	GetDisplayFunctionOnce.Do(func() {
		GetDisplayFunction = gi.FunctionInvokerNew("Gdk", "get_display")
	})
}

var getDisplayInvoker *gi.Function

// GetDisplay is a representation of the C type gdk_get_display.
func GetDisplay() string {
	if getDisplayInvoker == nil {
		getDisplayInvoker = gi.FunctionInvokerNew("Gdk", "get_display")
	}

	ret := getDisplayInvoker.Invoke(nil, nil)

	retGo := ret.String(true)

	return retGo
}

var GetDisplayArgNameFunction *gi.Function
var GetDisplayArgNameFunctionOnce sync.Once

func GetDisplayArgNameFunctionSet() {
	GetDisplayArgNameFunctionOnce.Do(func() {
		GetDisplayArgNameFunction = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	})
}

var getDisplayArgNameInvoker *gi.Function

// GetDisplayArgName is a representation of the C type gdk_get_display_arg_name.
func GetDisplayArgName() string {
	if getDisplayArgNameInvoker == nil {
		getDisplayArgNameInvoker = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	}

	ret := getDisplayArgNameInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

var GetProgramClassFunction *gi.Function
var GetProgramClassFunctionOnce sync.Once

func GetProgramClassFunctionSet() {
	GetProgramClassFunctionOnce.Do(func() {
		GetProgramClassFunction = gi.FunctionInvokerNew("Gdk", "get_program_class")
	})
}

var getProgramClassInvoker *gi.Function

// GetProgramClass is a representation of the C type gdk_get_program_class.
func GetProgramClass() string {
	if getProgramClassInvoker == nil {
		getProgramClassInvoker = gi.FunctionInvokerNew("Gdk", "get_program_class")
	}

	ret := getProgramClassInvoker.Invoke(nil, nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gdk_get_show_events' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_gl_error_quark' : return type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gdk_init' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gdk_init_check' : parameter 'argv' has no type

// UNSUPPORTED : C value 'gdk_keyboard_grab' : parameter 'window' of type 'Window' not supported

var KeyboardUngrabFunction *gi.Function
var KeyboardUngrabFunctionOnce sync.Once

func KeyboardUngrabFunctionSet() {
	KeyboardUngrabFunctionOnce.Do(func() {
		KeyboardUngrabFunction = gi.FunctionInvokerNew("Gdk", "keyboard_ungrab")
	})
}

var keyboardUngrabInvoker *gi.Function

// KeyboardUngrab is a representation of the C type gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	if keyboardUngrabInvoker == nil {
		keyboardUngrabInvoker = gi.FunctionInvokerNew("Gdk", "keyboard_ungrab")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	keyboardUngrabInvoker.Invoke(inArgs[:], nil)

}

var KeyvalConvertCaseFunction *gi.Function
var KeyvalConvertCaseFunctionOnce sync.Once

func KeyvalConvertCaseFunctionSet() {
	KeyvalConvertCaseFunctionOnce.Do(func() {
		KeyvalConvertCaseFunction = gi.FunctionInvokerNew("Gdk", "keyval_convert_case")
	})
}

var keyvalConvertCaseInvoker *gi.Function

// KeyvalConvertCase is a representation of the C type gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	if keyvalConvertCaseInvoker == nil {
		keyvalConvertCaseInvoker = gi.FunctionInvokerNew("Gdk", "keyval_convert_case")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(symbol)

	var outArgs [2]gi.Argument

	keyvalConvertCaseInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Uint32()

	return out0, out1
}

var KeyvalFromNameFunction *gi.Function
var KeyvalFromNameFunctionOnce sync.Once

func KeyvalFromNameFunctionSet() {
	KeyvalFromNameFunctionOnce.Do(func() {
		KeyvalFromNameFunction = gi.FunctionInvokerNew("Gdk", "keyval_from_name")
	})
}

var keyvalFromNameInvoker *gi.Function

// KeyvalFromName is a representation of the C type gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	if keyvalFromNameInvoker == nil {
		keyvalFromNameInvoker = gi.FunctionInvokerNew("Gdk", "keyval_from_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(keyvalName)

	ret := keyvalFromNameInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_keyval_is_lower' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_keyval_is_upper' : return type 'gboolean' not supported

var KeyvalNameFunction *gi.Function
var KeyvalNameFunctionOnce sync.Once

func KeyvalNameFunctionSet() {
	KeyvalNameFunctionOnce.Do(func() {
		KeyvalNameFunction = gi.FunctionInvokerNew("Gdk", "keyval_name")
	})
}

var keyvalNameInvoker *gi.Function

// KeyvalName is a representation of the C type gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	if keyvalNameInvoker == nil {
		keyvalNameInvoker = gi.FunctionInvokerNew("Gdk", "keyval_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	ret := keyvalNameInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var KeyvalToLowerFunction *gi.Function
var KeyvalToLowerFunctionOnce sync.Once

func KeyvalToLowerFunctionSet() {
	KeyvalToLowerFunctionOnce.Do(func() {
		KeyvalToLowerFunction = gi.FunctionInvokerNew("Gdk", "keyval_to_lower")
	})
}

var keyvalToLowerInvoker *gi.Function

// KeyvalToLower is a representation of the C type gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	if keyvalToLowerInvoker == nil {
		keyvalToLowerInvoker = gi.FunctionInvokerNew("Gdk", "keyval_to_lower")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	ret := keyvalToLowerInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var KeyvalToUnicodeFunction *gi.Function
var KeyvalToUnicodeFunctionOnce sync.Once

func KeyvalToUnicodeFunctionSet() {
	KeyvalToUnicodeFunctionOnce.Do(func() {
		KeyvalToUnicodeFunction = gi.FunctionInvokerNew("Gdk", "keyval_to_unicode")
	})
}

var keyvalToUnicodeInvoker *gi.Function

// KeyvalToUnicode is a representation of the C type gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	if keyvalToUnicodeInvoker == nil {
		keyvalToUnicodeInvoker = gi.FunctionInvokerNew("Gdk", "keyval_to_unicode")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	ret := keyvalToUnicodeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var KeyvalToUpperFunction *gi.Function
var KeyvalToUpperFunctionOnce sync.Once

func KeyvalToUpperFunctionSet() {
	KeyvalToUpperFunctionOnce.Do(func() {
		KeyvalToUpperFunction = gi.FunctionInvokerNew("Gdk", "keyval_to_upper")
	})
}

var keyvalToUpperInvoker *gi.Function

// KeyvalToUpper is a representation of the C type gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	if keyvalToUpperInvoker == nil {
		keyvalToUpperInvoker = gi.FunctionInvokerNew("Gdk", "keyval_to_upper")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	ret := keyvalToUpperInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_list_visuals' : return type 'GLib.List' not supported

var NotifyStartupCompleteFunction *gi.Function
var NotifyStartupCompleteFunctionOnce sync.Once

func NotifyStartupCompleteFunctionSet() {
	NotifyStartupCompleteFunctionOnce.Do(func() {
		NotifyStartupCompleteFunction = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	})
}

var notifyStartupCompleteInvoker *gi.Function

// NotifyStartupComplete is a representation of the C type gdk_notify_startup_complete.
func NotifyStartupComplete() {
	if notifyStartupCompleteInvoker == nil {
		notifyStartupCompleteInvoker = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	}

	notifyStartupCompleteInvoker.Invoke(nil, nil)

}

var NotifyStartupCompleteWithIdFunction *gi.Function
var NotifyStartupCompleteWithIdFunctionOnce sync.Once

func NotifyStartupCompleteWithIdFunctionSet() {
	NotifyStartupCompleteWithIdFunctionOnce.Do(func() {
		NotifyStartupCompleteWithIdFunction = gi.FunctionInvokerNew("Gdk", "notify_startup_complete_with_id")
	})
}

var notifyStartupCompleteWithIdInvoker *gi.Function

// NotifyStartupCompleteWithId is a representation of the C type gdk_notify_startup_complete_with_id.
func NotifyStartupCompleteWithId(startupId string) {
	if notifyStartupCompleteWithIdInvoker == nil {
		notifyStartupCompleteWithIdInvoker = gi.FunctionInvokerNew("Gdk", "notify_startup_complete_with_id")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(startupId)

	notifyStartupCompleteWithIdInvoker.Invoke(inArgs[:], nil)

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

var PointerUngrabFunction *gi.Function
var PointerUngrabFunctionOnce sync.Once

func PointerUngrabFunctionSet() {
	PointerUngrabFunctionOnce.Do(func() {
		PointerUngrabFunction = gi.FunctionInvokerNew("Gdk", "pointer_ungrab")
	})
}

var pointerUngrabInvoker *gi.Function

// PointerUngrab is a representation of the C type gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	if pointerUngrabInvoker == nil {
		pointerUngrabInvoker = gi.FunctionInvokerNew("Gdk", "pointer_ungrab")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	pointerUngrabInvoker.Invoke(inArgs[:], nil)

}

var PreParseLibgtkOnlyFunction *gi.Function
var PreParseLibgtkOnlyFunctionOnce sync.Once

func PreParseLibgtkOnlyFunctionSet() {
	PreParseLibgtkOnlyFunctionOnce.Do(func() {
		PreParseLibgtkOnlyFunction = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	})
}

var preParseLibgtkOnlyInvoker *gi.Function

// PreParseLibgtkOnly is a representation of the C type gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {
	if preParseLibgtkOnlyInvoker == nil {
		preParseLibgtkOnlyInvoker = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	}

	preParseLibgtkOnlyInvoker.Invoke(nil, nil)

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

var SetAllowedBackendsFunction *gi.Function
var SetAllowedBackendsFunctionOnce sync.Once

func SetAllowedBackendsFunctionSet() {
	SetAllowedBackendsFunctionOnce.Do(func() {
		SetAllowedBackendsFunction = gi.FunctionInvokerNew("Gdk", "set_allowed_backends")
	})
}

var setAllowedBackendsInvoker *gi.Function

// SetAllowedBackends is a representation of the C type gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	if setAllowedBackendsInvoker == nil {
		setAllowedBackendsInvoker = gi.FunctionInvokerNew("Gdk", "set_allowed_backends")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(backends)

	setAllowedBackendsInvoker.Invoke(inArgs[:], nil)

}

var SetDoubleClickTimeFunction *gi.Function
var SetDoubleClickTimeFunctionOnce sync.Once

func SetDoubleClickTimeFunctionSet() {
	SetDoubleClickTimeFunctionOnce.Do(func() {
		SetDoubleClickTimeFunction = gi.FunctionInvokerNew("Gdk", "set_double_click_time")
	})
}

var setDoubleClickTimeInvoker *gi.Function

// SetDoubleClickTime is a representation of the C type gdk_set_double_click_time.
func SetDoubleClickTime(msec uint32) {
	if setDoubleClickTimeInvoker == nil {
		setDoubleClickTimeInvoker = gi.FunctionInvokerNew("Gdk", "set_double_click_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(msec)

	setDoubleClickTimeInvoker.Invoke(inArgs[:], nil)

}

var SetProgramClassFunction *gi.Function
var SetProgramClassFunctionOnce sync.Once

func SetProgramClassFunctionSet() {
	SetProgramClassFunctionOnce.Do(func() {
		SetProgramClassFunction = gi.FunctionInvokerNew("Gdk", "set_program_class")
	})
}

var setProgramClassInvoker *gi.Function

// SetProgramClass is a representation of the C type gdk_set_program_class.
func SetProgramClass(programClass string) {
	if setProgramClassInvoker == nil {
		setProgramClassInvoker = gi.FunctionInvokerNew("Gdk", "set_program_class")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(programClass)

	setProgramClassInvoker.Invoke(inArgs[:], nil)

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

var ThreadsEnterFunction *gi.Function
var ThreadsEnterFunctionOnce sync.Once

func ThreadsEnterFunctionSet() {
	ThreadsEnterFunctionOnce.Do(func() {
		ThreadsEnterFunction = gi.FunctionInvokerNew("Gdk", "threads_enter")
	})
}

var threadsEnterInvoker *gi.Function

// ThreadsEnter is a representation of the C type gdk_threads_enter.
func ThreadsEnter() {
	if threadsEnterInvoker == nil {
		threadsEnterInvoker = gi.FunctionInvokerNew("Gdk", "threads_enter")
	}

	threadsEnterInvoker.Invoke(nil, nil)

}

var ThreadsInitFunction *gi.Function
var ThreadsInitFunctionOnce sync.Once

func ThreadsInitFunctionSet() {
	ThreadsInitFunctionOnce.Do(func() {
		ThreadsInitFunction = gi.FunctionInvokerNew("Gdk", "threads_init")
	})
}

var threadsInitInvoker *gi.Function

// ThreadsInit is a representation of the C type gdk_threads_init.
func ThreadsInit() {
	if threadsInitInvoker == nil {
		threadsInitInvoker = gi.FunctionInvokerNew("Gdk", "threads_init")
	}

	threadsInitInvoker.Invoke(nil, nil)

}

var ThreadsLeaveFunction *gi.Function
var ThreadsLeaveFunctionOnce sync.Once

func ThreadsLeaveFunctionSet() {
	ThreadsLeaveFunctionOnce.Do(func() {
		ThreadsLeaveFunction = gi.FunctionInvokerNew("Gdk", "threads_leave")
	})
}

var threadsLeaveInvoker *gi.Function

// ThreadsLeave is a representation of the C type gdk_threads_leave.
func ThreadsLeave() {
	if threadsLeaveInvoker == nil {
		threadsLeaveInvoker = gi.FunctionInvokerNew("Gdk", "threads_leave")
	}

	threadsLeaveInvoker.Invoke(nil, nil)

}

// UNSUPPORTED : C value 'gdk_threads_set_lock_functions' : parameter 'enter_fn' of type 'GObject.Callback' not supported

var UnicodeToKeyvalFunction *gi.Function
var UnicodeToKeyvalFunctionOnce sync.Once

func UnicodeToKeyvalFunctionSet() {
	UnicodeToKeyvalFunctionOnce.Do(func() {
		UnicodeToKeyvalFunction = gi.FunctionInvokerNew("Gdk", "unicode_to_keyval")
	})
}

var unicodeToKeyvalInvoker *gi.Function

// UnicodeToKeyval is a representation of the C type gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	if unicodeToKeyvalInvoker == nil {
		unicodeToKeyvalInvoker = gi.FunctionInvokerNew("Gdk", "unicode_to_keyval")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(wc)

	ret := unicodeToKeyvalInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var Utf8ToStringTargetFunction *gi.Function
var Utf8ToStringTargetFunctionOnce sync.Once

func Utf8ToStringTargetFunctionSet() {
	Utf8ToStringTargetFunctionOnce.Do(func() {
		Utf8ToStringTargetFunction = gi.FunctionInvokerNew("Gdk", "utf8_to_string_target")
	})
}

var utf8ToStringTargetInvoker *gi.Function

// Utf8ToStringTarget is a representation of the C type gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	if utf8ToStringTargetInvoker == nil {
		utf8ToStringTargetInvoker = gi.FunctionInvokerNew("Gdk", "utf8_to_string_target")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	ret := utf8ToStringTargetInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}
