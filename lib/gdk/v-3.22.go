// Code generated - DO NOT EDIT.
// +build gdk_3.22 gdk_3.24

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// UNSUPPORTED : XEvent : blacklisted

// AnchorHints is a representation of the C bitfield GdkAnchorHints.
type AnchorHints int

// AnchorHints_flip_x is a representation of the C bitfield member GDK_ANCHOR_FLIP_X.
const AnchorHints_flip_x = AnchorHints(1)

// AnchorHints_flip_y is a representation of the C bitfield member GDK_ANCHOR_FLIP_Y.
const AnchorHints_flip_y = AnchorHints(2)

// AnchorHints_slide_x is a representation of the C bitfield member GDK_ANCHOR_SLIDE_X.
const AnchorHints_slide_x = AnchorHints(4)

// AnchorHints_slide_y is a representation of the C bitfield member GDK_ANCHOR_SLIDE_Y.
const AnchorHints_slide_y = AnchorHints(8)

// AnchorHints_resize_x is a representation of the C bitfield member GDK_ANCHOR_RESIZE_X.
const AnchorHints_resize_x = AnchorHints(16)

// AnchorHints_resize_y is a representation of the C bitfield member GDK_ANCHOR_RESIZE_Y.
const AnchorHints_resize_y = AnchorHints(32)

// AnchorHints_flip is a representation of the C bitfield member GDK_ANCHOR_FLIP.
const AnchorHints_flip = AnchorHints(3)

// AnchorHints_slide is a representation of the C bitfield member GDK_ANCHOR_SLIDE.
const AnchorHints_slide = AnchorHints(12)

// AnchorHints_resize is a representation of the C bitfield member GDK_ANCHOR_RESIZE.
const AnchorHints_resize = AnchorHints(48)

// AxisFlags is a representation of the C bitfield GdkAxisFlags.
type AxisFlags int

// AxisFlags_x is a representation of the C bitfield member GDK_AXIS_FLAG_X.
const AxisFlags_x = AxisFlags(2)

// AxisFlags_y is a representation of the C bitfield member GDK_AXIS_FLAG_Y.
const AxisFlags_y = AxisFlags(4)

// AxisFlags_pressure is a representation of the C bitfield member GDK_AXIS_FLAG_PRESSURE.
const AxisFlags_pressure = AxisFlags(8)

// AxisFlags_xtilt is a representation of the C bitfield member GDK_AXIS_FLAG_XTILT.
const AxisFlags_xtilt = AxisFlags(16)

// AxisFlags_ytilt is a representation of the C bitfield member GDK_AXIS_FLAG_YTILT.
const AxisFlags_ytilt = AxisFlags(32)

// AxisFlags_wheel is a representation of the C bitfield member GDK_AXIS_FLAG_WHEEL.
const AxisFlags_wheel = AxisFlags(64)

// AxisFlags_distance is a representation of the C bitfield member GDK_AXIS_FLAG_DISTANCE.
const AxisFlags_distance = AxisFlags(128)

// AxisFlags_rotation is a representation of the C bitfield member GDK_AXIS_FLAG_ROTATION.
const AxisFlags_rotation = AxisFlags(256)

// AxisFlags_slider is a representation of the C bitfield member GDK_AXIS_FLAG_SLIDER.
const AxisFlags_slider = AxisFlags(512)

// DeviceToolType is a representation of the C enumeration GdkDeviceToolType.
type DeviceToolType int

// DeviceToolType_unknown is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_UNKNOWN.
const DeviceToolType_unknown = DeviceToolType(0)

// DeviceToolType_pen is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_PEN.
const DeviceToolType_pen = DeviceToolType(1)

// DeviceToolType_eraser is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_ERASER.
const DeviceToolType_eraser = DeviceToolType(2)

// DeviceToolType_brush is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_BRUSH.
const DeviceToolType_brush = DeviceToolType(3)

// DeviceToolType_pencil is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_PENCIL.
const DeviceToolType_pencil = DeviceToolType(4)

// DeviceToolType_airbrush is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_AIRBRUSH.
const DeviceToolType_airbrush = DeviceToolType(5)

// DeviceToolType_mouse is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_MOUSE.
const DeviceToolType_mouse = DeviceToolType(6)

// DeviceToolType_lens is a representation of the C enumeration member GDK_DEVICE_TOOL_TYPE_LENS.
const DeviceToolType_lens = DeviceToolType(7)

// SubpixelLayout is a representation of the C enumeration GdkSubpixelLayout.
type SubpixelLayout int

// SubpixelLayout_unknown is a representation of the C enumeration member GDK_SUBPIXEL_LAYOUT_UNKNOWN.
const SubpixelLayout_unknown = SubpixelLayout(0)

// SubpixelLayout_none is a representation of the C enumeration member GDK_SUBPIXEL_LAYOUT_NONE.
const SubpixelLayout_none = SubpixelLayout(1)

// SubpixelLayout_horizontal_rgb is a representation of the C enumeration member GDK_SUBPIXEL_LAYOUT_HORIZONTAL_RGB.
const SubpixelLayout_horizontal_rgb = SubpixelLayout(2)

// SubpixelLayout_horizontal_bgr is a representation of the C enumeration member GDK_SUBPIXEL_LAYOUT_HORIZONTAL_BGR.
const SubpixelLayout_horizontal_bgr = SubpixelLayout(3)

// SubpixelLayout_vertical_rgb is a representation of the C enumeration member GDK_SUBPIXEL_LAYOUT_VERTICAL_RGB.
const SubpixelLayout_vertical_rgb = SubpixelLayout(4)

// SubpixelLayout_vertical_bgr is a representation of the C enumeration member GDK_SUBPIXEL_LAYOUT_VERTICAL_BGR.
const SubpixelLayout_vertical_bgr = SubpixelLayout(5)

// UNSUPPORTED : gdk_cairo_get_clip_rectangle : has [in]out param, rect

// CairoGetDrawingContext wraps the C function gdk_cairo_get_drawing_context.
//
// since 3.22
func CairoGetDrawingContext(cr *cairo.Context) *DrawingContext {
	sys_cr := cr.ToC()
	retSys := gdk.Fn_gdk_cairo_get_drawing_context(sys_cr)
	ret := DrawingContextNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gdk_color_parse : has [in]out param, color

// UNSUPPORTED : gdk_drag_find_window_for_screen : has [in]out param, dest_window

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_events_get_angle : has [in]out param, angle

// UNSUPPORTED : gdk_events_get_center : has [in]out param, x

// UNSUPPORTED : gdk_events_get_distance : has [in]out param, distance

// UNSUPPORTED : gdk_init : has array param, argv

// UNSUPPORTED : gdk_init_check : has array param, argv

// UNSUPPORTED : gdk_keyval_convert_case : has [in]out param, lower

// PangoContextGetForDisplay wraps the C function gdk_pango_context_get_for_display.
//
// since 3.22
func PangoContextGetForDisplay(display *Display) *pango.Context {
	sys_display := display.ToC()
	retSys := gdk.Fn_gdk_pango_context_get_for_display(sys_display)
	ret := pango.ContextNewFromC(retSys)

	return ret
}

// UNSUPPORTED : gdk_pango_layout_line_get_clip_region : parameter 'index_ranges' is array parameter without length parameter

// UNSUPPORTED : gdk_parse_args : has array param, argv

// UNSUPPORTED : gdk_property_get : has [in]out param, actual_property_type

// UNSUPPORTED : gdk_query_depths : has array param, depths

// UNSUPPORTED : gdk_query_visual_types : has array param, visual_types

// UNSUPPORTED : gdk_synthesize_window_state : blacklisted

// UNSUPPORTED : gdk_text_property_to_utf8_list_for_display : parameter 'list' is array parameter without length parameter

// UNSUPPORTED : gdk_threads_add_idle : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_idle_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_add_timeout_seconds_full : parameter 'function' is callback

// UNSUPPORTED : gdk_threads_set_lock_functions : parameter 'enter_fn' is callback

// EventPadAxis is a representation of the C record GdkEventPadAxis.
//
// since 3.22
type EventPadAxis struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventPadAxis that represents the EventPadAxis.
func (recv *EventPadAxis) ToC() unsafe.Pointer {
	return recv.native
}

// EventPadAxisNewFromC creates a new EventPadAxis from a pointer to the C GdkEventPadAxis that represents the EventPadAxis.
func EventPadAxisNewFromC(native unsafe.Pointer) *EventPadAxis {
	return &EventPadAxis{native: native}
}

// EventPadButton is a representation of the C record GdkEventPadButton.
//
// since 3.22
type EventPadButton struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventPadButton that represents the EventPadButton.
func (recv *EventPadButton) ToC() unsafe.Pointer {
	return recv.native
}

// EventPadButtonNewFromC creates a new EventPadButton from a pointer to the C GdkEventPadButton that represents the EventPadButton.
func EventPadButtonNewFromC(native unsafe.Pointer) *EventPadButton {
	return &EventPadButton{native: native}
}

// EventPadGroupMode is a representation of the C record GdkEventPadGroupMode.
//
// since 3.22
type EventPadGroupMode struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GdkEventPadGroupMode that represents the EventPadGroupMode.
func (recv *EventPadGroupMode) ToC() unsafe.Pointer {
	return recv.native
}

// EventPadGroupModeNewFromC creates a new EventPadGroupMode from a pointer to the C GdkEventPadGroupMode that represents the EventPadGroupMode.
func EventPadGroupModeNewFromC(native unsafe.Pointer) *EventPadGroupMode {
	return &EventPadGroupMode{native: native}
}
