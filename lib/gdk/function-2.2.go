// This is a generated file - DO NOT EDIT
// +build gdk_2.2 gdk_2.4 gdk_2.6 gdk_2.8 gdk_2.10 gdk_2.12 gdk_2.14 gdk_2.16 gdk_2.18 gdk_2.20 gdk_2.22 gdk_2.24 gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	"C"
	pango "github.com/pekim/gobbi/lib/pango"
	"unsafe"
)

// Unsupported : gdk_drag_find_window_for_screen : unsupported parameter protocol : GdkDragProtocol* with indirection level of 1

// GetDisplayArgName is a wrapper around the C function gdk_get_display_arg_name.
func GetDisplayArgName() string {
	retC := C.gdk_get_display_arg_name()
	retGo := C.GoString(retC)

	return retGo
}

// NotifyStartupComplete is a wrapper around the C function gdk_notify_startup_complete.
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()

	return
}

// PangoContextGetForScreen is a wrapper around the C function gdk_pango_context_get_for_screen.
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gdk_pango_context_get_for_screen(c_screen)
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ParseArgs is a wrapper around the C function gdk_parse_args.
func ParseArgs(args []string) []string {
	cArgc, cArgv := argsIn(args)

	C.gdk_parse_args(&cArgc, &cArgv)

	args = argsOut(cArgc, cArgv)

	return args
}

// SelectionOwnerGetForDisplay is a wrapper around the C function gdk_selection_owner_get_for_display.
func SelectionOwnerGetForDisplay(display *Display, selection *Atom) *Window {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	retC := C.gdk_selection_owner_get_for_display(c_display, c_selection)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SelectionOwnerSetForDisplay is a wrapper around the C function gdk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_owner := (*C.GdkWindow)(C.NULL)
	if owner != nil {
		c_owner = (*C.GdkWindow)(owner.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_time_ := (C.guint32)(time)

	c_send_event :=
		boolToGboolean(sendEvent)

	retC := C.gdk_selection_owner_set_for_display(c_display, c_owner, c_selection, c_time_, c_send_event)
	retGo := retC == C.TRUE

	return retGo
}

// SelectionSendNotifyForDisplay is a wrapper around the C function gdk_selection_send_notify_for_display.
func SelectionSendNotifyForDisplay(display *Display, requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	c_requestor := (*C.GdkWindow)(C.NULL)
	if requestor != nil {
		c_requestor = (*C.GdkWindow)(requestor.ToC())
	}

	c_selection := (C.GdkAtom)(C.NULL)
	if selection != nil {
		c_selection = (C.GdkAtom)(selection.ToC())
	}

	c_target := (C.GdkAtom)(C.NULL)
	if target != nil {
		c_target = (C.GdkAtom)(target.ToC())
	}

	c_property := (C.GdkAtom)(C.NULL)
	if property != nil {
		c_property = (C.GdkAtom)(property.ToC())
	}

	c_time_ := (C.guint32)(time)

	C.gdk_selection_send_notify_for_display(c_display, c_requestor, c_selection, c_target, c_property, c_time_)

	return
}

// Unsupported : gdk_text_property_to_utf8_list_for_display : unsupported parameter list : output array param list
