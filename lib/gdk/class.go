// This is a generated file - DO NOT EDIT

package gdk

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkAppLaunchContext

// Blacklisted : GdkCursor

// Blacklisted : GdkDevice

// Blacklisted : GdkDeviceManager

// Blacklisted : GdkDisplay

// Blacklisted : GdkDisplayManager

// DragContext is a wrapper around the C record GdkDragContext.
type DragContext struct {
	native *C.GdkDragContext
}

func DragContextNewFromC(u unsafe.Pointer) *DragContext {
	c := (*C.GdkDragContext)(u)
	if c == nil {
		return nil
	}

	g := &DragContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DragContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DragContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DragContext with another DragContext, and returns true if they represent the same GObject.
func (recv *DragContext) Equals(other *DragContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DragContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DragContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a DragContext.
func CastToDragContext(object *gobject.Object) *DragContext {
	return DragContextNewFromC(object.ToC())
}

// Blacklisted : gdk_drag_context_get_device

// Blacklisted : gdk_drag_context_set_device

// Blacklisted : GdkFrameClock

// Blacklisted : GdkGLContext

// Blacklisted : GdkKeymap

// Blacklisted : GdkScreen

// Blacklisted : GdkVisual

// Window is a wrapper around the C record GdkWindow.
type Window struct {
	native *C.GdkWindow
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.GdkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Window) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Window with another Window, and returns true if they represent the same GObject.
func (recv *Window) Equals(other *Window) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Window) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Window.
// Exercise care, as this is a potentially dangerous function if the Object is not a Window.
func CastToWindow(object *gobject.Object) *Window {
	return WindowNewFromC(object.ToC())
}

// WindowNew is a wrapper around the C function gdk_window_new.
func WindowNew(parent *Window, attributes *WindowAttr, attributesMask WindowAttributesType) *Window {
	c_parent := (*C.GdkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GdkWindow)(parent.ToC())
	}

	c_attributes := (*C.GdkWindowAttr)(C.NULL)
	if attributes != nil {
		c_attributes = (*C.GdkWindowAttr)(attributes.ToC())
	}

	c_attributes_mask := (C.gint)(attributesMask)

	retC := C.gdk_window_new(c_parent, c_attributes, c_attributes_mask)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Blacklisted : gdk_window_at_pointer

// Blacklisted : gdk_window_constrain_size

// Blacklisted : gdk_window_process_all_updates

// Blacklisted : gdk_window_set_debug_updates

// Unsupported : gdk_window_add_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Blacklisted : gdk_window_begin_move_drag

// Blacklisted : gdk_window_begin_paint_rect

// Blacklisted : gdk_window_begin_paint_region

// Blacklisted : gdk_window_begin_resize_drag

// Blacklisted : gdk_window_deiconify

// Destroy is a wrapper around the C function gdk_window_destroy.
func (recv *Window) Destroy() {
	C.gdk_window_destroy((*C.GdkWindow)(recv.native))

	return
}

// Blacklisted : gdk_window_destroy_notify

// Blacklisted : gdk_window_end_paint

// Blacklisted : gdk_window_focus

// Blacklisted : gdk_window_freeze_toplevel_updates_libgtk_only

// Blacklisted : gdk_window_freeze_updates

// Blacklisted : gdk_window_fullscreen_on_monitor

// Blacklisted : gdk_window_get_children

// Blacklisted : gdk_window_get_clip_region

// Unsupported : gdk_window_get_decorations : unsupported parameter decorations : GdkWMDecoration* with indirection level of 1

// Blacklisted : gdk_window_get_events

// Blacklisted : gdk_window_get_frame_extents

// Blacklisted : gdk_window_get_geometry

// Blacklisted : gdk_window_get_origin

// Blacklisted : gdk_window_get_parent

// Unsupported : gdk_window_get_pointer : unsupported parameter mask : GdkModifierType* with indirection level of 1

// Blacklisted : gdk_window_get_position

// Blacklisted : gdk_window_get_root_origin

// Blacklisted : gdk_window_get_source_events

// Blacklisted : gdk_window_get_state

// Blacklisted : gdk_window_get_toplevel

// Blacklisted : gdk_window_get_update_area

// Unsupported : gdk_window_get_user_data : unsupported parameter data : no type generator for gpointer (gpointer*) for param data

// Blacklisted : gdk_window_get_visible_region

// Blacklisted : gdk_window_get_window_type

// Blacklisted : gdk_window_hide

// Blacklisted : gdk_window_iconify

// Unsupported : gdk_window_invalidate_maybe_recurse : unsupported parameter child_func : no type generator for WindowChildFunc (GdkWindowChildFunc) for param child_func

// Blacklisted : gdk_window_invalidate_rect

// Blacklisted : gdk_window_invalidate_region

// Blacklisted : gdk_window_is_viewable

// IsVisible is a wrapper around the C function gdk_window_is_visible.
func (recv *Window) IsVisible() bool {
	retC := C.gdk_window_is_visible((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : gdk_window_lower

// Blacklisted : gdk_window_maximize

// Blacklisted : gdk_window_merge_child_shapes

// Blacklisted : gdk_window_move

// Blacklisted : gdk_window_move_resize

// Blacklisted : gdk_window_peek_children

// Blacklisted : gdk_window_process_updates

// Blacklisted : gdk_window_raise

// Blacklisted : gdk_window_register_dnd

// Unsupported : gdk_window_remove_filter : unsupported parameter function : no type generator for FilterFunc (GdkFilterFunc) for param function

// Blacklisted : gdk_window_reparent

// Blacklisted : gdk_window_resize

// Blacklisted : gdk_window_scroll

// Blacklisted : gdk_window_set_background

// Blacklisted : gdk_window_set_background_pattern

// Blacklisted : gdk_window_set_background_rgba

// Blacklisted : gdk_window_set_child_shapes

// Blacklisted : gdk_window_set_cursor

// Blacklisted : gdk_window_set_decorations

// Blacklisted : gdk_window_set_events

// Blacklisted : gdk_window_set_functions

// Blacklisted : gdk_window_set_geometry_hints

// Blacklisted : gdk_window_set_group

// Blacklisted : gdk_window_set_icon_list

// Blacklisted : gdk_window_set_icon_name

// Blacklisted : gdk_window_set_modal_hint

// Blacklisted : gdk_window_set_override_redirect

// Blacklisted : gdk_window_set_role

// Blacklisted : gdk_window_set_static_gravities

// SetTitle is a wrapper around the C function gdk_window_set_title.
func (recv *Window) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gdk_window_set_title((*C.GdkWindow)(recv.native), c_title)

	return
}

// Blacklisted : gdk_window_set_transient_for

// Blacklisted : gdk_window_set_type_hint

// Blacklisted : gdk_window_set_user_data

// Blacklisted : gdk_window_shape_combine_region

// Blacklisted : gdk_window_show

// Blacklisted : gdk_window_show_unraised

// Blacklisted : gdk_window_stick

// Blacklisted : gdk_window_thaw_toplevel_updates_libgtk_only

// Blacklisted : gdk_window_thaw_updates

// Blacklisted : gdk_window_unmaximize

// Blacklisted : gdk_window_unstick

// Blacklisted : gdk_window_withdraw
