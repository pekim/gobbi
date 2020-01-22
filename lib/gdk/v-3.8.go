// Code generated - DO NOT EDIT.
// +build gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22 gdk_3.24

package gdk

// UNSUPPORTED : XEvent : blacklisted

// FrameClockPhase is a representation of the C bitfield GdkFrameClockPhase.
type FrameClockPhase int

// FrameClockPhase_none is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_NONE.
const FrameClockPhase_none = FrameClockPhase(0)

// FrameClockPhase_flush_events is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS.
const FrameClockPhase_flush_events = FrameClockPhase(1)

// FrameClockPhase_before_paint is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT.
const FrameClockPhase_before_paint = FrameClockPhase(2)

// FrameClockPhase_update is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_UPDATE.
const FrameClockPhase_update = FrameClockPhase(4)

// FrameClockPhase_layout is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_LAYOUT.
const FrameClockPhase_layout = FrameClockPhase(8)

// FrameClockPhase_paint is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_PAINT.
const FrameClockPhase_paint = FrameClockPhase(16)

// FrameClockPhase_resume_events is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS.
const FrameClockPhase_resume_events = FrameClockPhase(32)

// FrameClockPhase_after_paint is a representation of the C bitfield member GDK_FRAME_CLOCK_PHASE_AFTER_PAINT.
const FrameClockPhase_after_paint = FrameClockPhase(64)

// FullscreenMode is a representation of the C enumeration GdkFullscreenMode.
type FullscreenMode int

// FullscreenMode_current_monitor is a representation of the C enumeration member GDK_FULLSCREEN_ON_CURRENT_MONITOR.
const FullscreenMode_current_monitor = FullscreenMode(0)

// FullscreenMode_all_monitors is a representation of the C enumeration member GDK_FULLSCREEN_ON_ALL_MONITORS.
const FullscreenMode_all_monitors = FullscreenMode(1)

// UNSUPPORTED : gdk_cairo_get_clip_rectangle : has [in]out param, rect

// UNSUPPORTED : gdk_color_parse : has [in]out param, color

// UNSUPPORTED : gdk_drag_find_window_for_screen : has [in]out param, dest_window

// UNSUPPORTED : gdk_event_handler_set : parameter 'func' is callback

// UNSUPPORTED : gdk_events_get_angle : has [in]out param, angle

// UNSUPPORTED : gdk_events_get_center : has [in]out param, x

// UNSUPPORTED : gdk_events_get_distance : has [in]out param, distance

// UNSUPPORTED : gdk_init : has array param, argv

// UNSUPPORTED : gdk_init_check : has array param, argv

// UNSUPPORTED : gdk_keyval_convert_case : has [in]out param, lower

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
