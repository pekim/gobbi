// Code generated - DO NOT EDIT.
// +build gtk_3.0

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gtk "github.com/pekim/gobbi/lib/internal/c/gtk"
	pango "github.com/pekim/gobbi/lib/pango"
)

// License is a representation of the C enumeration GtkLicense.
type License int

// License_unknown is a representation of the C enumeration member GTK_LICENSE_UNKNOWN.
const License_unknown = License(0)

// License_custom is a representation of the C enumeration member GTK_LICENSE_CUSTOM.
const License_custom = License(1)

// License_gpl_2_0 is a representation of the C enumeration member GTK_LICENSE_GPL_2_0.
const License_gpl_2_0 = License(2)

// License_gpl_3_0 is a representation of the C enumeration member GTK_LICENSE_GPL_3_0.
const License_gpl_3_0 = License(3)

// License_lgpl_2_1 is a representation of the C enumeration member GTK_LICENSE_LGPL_2_1.
const License_lgpl_2_1 = License(4)

// License_lgpl_3_0 is a representation of the C enumeration member GTK_LICENSE_LGPL_3_0.
const License_lgpl_3_0 = License(5)

// License_bsd is a representation of the C enumeration member GTK_LICENSE_BSD.
const License_bsd = License(6)

// License_mit_x11 is a representation of the C enumeration member GTK_LICENSE_MIT_X11.
const License_mit_x11 = License(7)

// License_artistic is a representation of the C enumeration member GTK_LICENSE_ARTISTIC.
const License_artistic = License(8)

// License_gpl_2_0_only is a representation of the C enumeration member GTK_LICENSE_GPL_2_0_ONLY.
const License_gpl_2_0_only = License(9)

// License_gpl_3_0_only is a representation of the C enumeration member GTK_LICENSE_GPL_3_0_ONLY.
const License_gpl_3_0_only = License(10)

// License_lgpl_2_1_only is a representation of the C enumeration member GTK_LICENSE_LGPL_2_1_ONLY.
const License_lgpl_2_1_only = License(11)

// License_lgpl_3_0_only is a representation of the C enumeration member GTK_LICENSE_LGPL_3_0_ONLY.
const License_lgpl_3_0_only = License(12)

// License_agpl_3_0 is a representation of the C enumeration member GTK_LICENSE_AGPL_3_0.
const License_agpl_3_0 = License(13)

// License_agpl_3_0_only is a representation of the C enumeration member GTK_LICENSE_AGPL_3_0_ONLY.
const License_agpl_3_0_only = License(14)

// UNSUPPORTED : gtk_accelerator_parse : has [in]out param, accelerator_key

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// CairoShouldDrawWindow wraps the C function gtk_cairo_should_draw_window.
//
// since 3.0
func CairoShouldDrawWindow(cr *cairo.Context, window *gdk.Window) bool {
	sys_cr := cr.ToC()
	sys_window := window.ToC()
	retSys := gtk.Fn_gtk_cairo_should_draw_window(sys_cr, sys_window)
	ret := retSys

	return ret
}

// CairoTransformToWindow wraps the C function gtk_cairo_transform_to_window.
//
// since 3.0
func CairoTransformToWindow(cr *cairo.Context, widget *Widget, window *gdk.Window) {
	sys_cr := cr.ToC()
	sys_widget := widget.ToC()
	sys_window := window.ToC()
	gtk.Fn_gtk_cairo_transform_to_window(sys_cr, sys_widget, sys_window)
}

// DeviceGrabAdd wraps the C function gtk_device_grab_add.
//
// since 3.0
func DeviceGrabAdd(widget *Widget, device *gdk.Device, blockOthers bool) {
	sys_widget := widget.ToC()
	sys_device := device.ToC()
	sys_blockOthers := blockOthers
	gtk.Fn_gtk_device_grab_add(sys_widget, sys_device, sys_blockOthers)
}

// DeviceGrabRemove wraps the C function gtk_device_grab_remove.
//
// since 3.0
func DeviceGrabRemove(widget *Widget, device *gdk.Device) {
	sys_widget := widget.ToC()
	sys_device := device.ToC()
	gtk.Fn_gtk_device_grab_remove(sys_widget, sys_device)
}

// DrawInsertionCursor wraps the C function gtk_draw_insertion_cursor.
//
// since 3.0
func DrawInsertionCursor(widget *Widget, cr *cairo.Context, location *gdk.Rectangle, isPrimary bool, direction TextDirection, drawArrow bool) {
	sys_widget := widget.ToC()
	sys_cr := cr.ToC()
	sys_location := location.ToC()
	sys_isPrimary := isPrimary
	sys_direction := (int)(direction)
	sys_drawArrow := drawArrow
	gtk.Fn_gtk_draw_insertion_cursor(sys_widget, sys_cr, sys_location, sys_isPrimary, sys_direction, sys_drawArrow)
}

// GetBinaryAge wraps the C function gtk_get_binary_age.
//
// since 3.0
func GetBinaryAge() uint {
	retSys := gtk.Fn_gtk_get_binary_age()
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_get_current_event_state : has [in]out param, state

// GetInterfaceAge wraps the C function gtk_get_interface_age.
//
// since 3.0
func GetInterfaceAge() uint {
	retSys := gtk.Fn_gtk_get_interface_age()
	ret := retSys

	return ret
}

// GetMajorVersion wraps the C function gtk_get_major_version.
//
// since 3.0
func GetMajorVersion() uint {
	retSys := gtk.Fn_gtk_get_major_version()
	ret := retSys

	return ret
}

// GetMicroVersion wraps the C function gtk_get_micro_version.
//
// since 3.0
func GetMicroVersion() uint {
	retSys := gtk.Fn_gtk_get_micro_version()
	ret := retSys

	return ret
}

// GetMinorVersion wraps the C function gtk_get_minor_version.
//
// since 3.0
func GetMinorVersion() uint {
	retSys := gtk.Fn_gtk_get_minor_version()
	ret := retSys

	return ret
}

// UNSUPPORTED : gtk_icon_size_lookup : has [in]out param, width

// UNSUPPORTED : gtk_icon_size_lookup_for_settings : has [in]out param, width

// UNSUPPORTED : gtk_init : has array param, argv

// UNSUPPORTED : gtk_init_check : has array param, argv

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_parse_args : has array param, argv

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_parse_color : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_color_full : has [in]out param, color

// UNSUPPORTED : gtk_rc_parse_state : has [in]out param, state

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// RenderActivity wraps the C function gtk_render_activity.
//
// since 3.0
func RenderActivity(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_activity(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderArrow wraps the C function gtk_render_arrow.
//
// since 3.0
func RenderArrow(context *StyleContext, cr *cairo.Context, angle float64, x float64, y float64, size float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_angle := angle
	sys_x := x
	sys_y := y
	sys_size := size
	gtk.Fn_gtk_render_arrow(sys_context, sys_cr, sys_angle, sys_x, sys_y, sys_size)
}

// RenderBackground wraps the C function gtk_render_background.
//
// since 3.0.
func RenderBackground(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_background(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// UNSUPPORTED : gtk_render_background_get_clip : has [in]out param, out_clip

// RenderCheck wraps the C function gtk_render_check.
//
// since 3.0
func RenderCheck(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_check(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderExpander wraps the C function gtk_render_expander.
//
// since 3.0
func RenderExpander(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_expander(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderExtension wraps the C function gtk_render_extension.
//
// since 3.0
func RenderExtension(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := (int)(gapSide)
	gtk.Fn_gtk_render_extension(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height, sys_gapSide)
}

// RenderFocus wraps the C function gtk_render_focus.
//
// since 3.0
func RenderFocus(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_focus(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderFrame wraps the C function gtk_render_frame.
//
// since 3.0
func RenderFrame(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_frame(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderFrameGap wraps the C function gtk_render_frame_gap.
//
// since 3.0
func RenderFrameGap(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType, xy0Gap float64, xy1Gap float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_gapSide := (int)(gapSide)
	sys_xy0Gap := xy0Gap
	sys_xy1Gap := xy1Gap
	gtk.Fn_gtk_render_frame_gap(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height, sys_gapSide, sys_xy0Gap, sys_xy1Gap)
}

// RenderHandle wraps the C function gtk_render_handle.
//
// since 3.0
func RenderHandle(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_handle(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderIconPixbuf wraps the C function gtk_render_icon_pixbuf.
//
// since 3.0
func RenderIconPixbuf(context *StyleContext, source *IconSource, size IconSize) *gdkpixbuf.Pixbuf {
	sys_context := context.ToC()
	sys_source := source.ToC()
	sys_size := (int)(size)
	retSys := gtk.Fn_gtk_render_icon_pixbuf(sys_context, sys_source, sys_size)
	ret := gdkpixbuf.PixbufNewFromC(retSys)

	return ret
}

// RenderLayout wraps the C function gtk_render_layout.
//
// since 3.0
func RenderLayout(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_layout := layout.ToC()
	gtk.Fn_gtk_render_layout(sys_context, sys_cr, sys_x, sys_y, sys_layout)
}

// RenderLine wraps the C function gtk_render_line.
//
// since 3.0
func RenderLine(context *StyleContext, cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x0 := x0
	sys_y0 := y0
	sys_x1 := x1
	sys_y1 := y1
	gtk.Fn_gtk_render_line(sys_context, sys_cr, sys_x0, sys_y0, sys_x1, sys_y1)
}

// RenderOption wraps the C function gtk_render_option.
//
// since 3.0
func RenderOption(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	gtk.Fn_gtk_render_option(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height)
}

// RenderSlider wraps the C function gtk_render_slider.
//
// since 3.0
func RenderSlider(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, orientation Orientation) {
	sys_context := context.ToC()
	sys_cr := cr.ToC()
	sys_x := x
	sys_y := y
	sys_width := width
	sys_height := height
	sys_orientation := (int)(orientation)
	gtk.Fn_gtk_render_slider(sys_context, sys_cr, sys_x, sys_y, sys_width, sys_height, sys_orientation)
}

// UNSUPPORTED : gtk_rgb_to_hsv : has [in]out param, h

// UNSUPPORTED : gtk_selection_add_targets : has array param, targets

// UNSUPPORTED : gtk_show_uri : throws

// UNSUPPORTED : gtk_show_uri_on_window : throws

// UNSUPPORTED : gtk_stock_add : has array param, items

// UNSUPPORTED : gtk_stock_add_static : has array param, items

// UNSUPPORTED : gtk_stock_lookup : has [in]out param, item

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_target_table_free : has array param, targets

// UNSUPPORTED : gtk_target_table_new_from_list : has [in]out param, n_targets

// UNSUPPORTED : gtk_targets_include_image : has array param, targets

// UNSUPPORTED : gtk_targets_include_rich_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_text : has array param, targets

// UNSUPPORTED : gtk_targets_include_uri : has array param, targets

// UNSUPPORTED : gtk_test_init : has array param, argvp

// UNSUPPORTED : gtk_test_list_all_types : has [in]out param, n_types

// UNSUPPORTED : gtk_tree_get_row_drag_data : has [in]out param, tree_model

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : EventControllerMotionClass : blacklisted

// UNSUPPORTED : EventControllerScrollClass : blacklisted

// UNSUPPORTED : GestureStylusClass : blacklisted

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted

// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted

// UNSUPPORTED : StackAccessibleClass : blacklisted

// UNSUPPORTED : EntryIconAccessible : blacklisted

// UNSUPPORTED : EventControllerKey : blacklisted

// UNSUPPORTED : EventControllerMotion : blacklisted

// UNSUPPORTED : EventControllerScroll : blacklisted

// UNSUPPORTED : GestureStylus : blacklisted

// UNSUPPORTED : HeaderBarAccessible : blacklisted

// UNSUPPORTED : StackAccessible : blacklisted
