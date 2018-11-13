// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_sequence_lookup : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_lookup_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp

// Adds @child_source to @source as a "polled" source; when @source is
// added to a #GMainContext, @child_source will be automatically added
// with the same priority, when @child_source is triggered, it will
// cause @source to dispatch (in addition to calling its own
// callback), and when @source is destroyed, it will destroy
// @child_source as well. (@source will also still be dispatched if
// its own prepare/check functions indicate that it is ready.)
//
// If you don't need @child_source to do anything on its own when it
// triggers, you can call g_source_set_dummy_callback() on it to set a
// callback that does nothing (except return %TRUE if appropriate).
//
// @source will hold a reference on @child_source while @child_source
// is attached to it.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
/*

C function

g_source_add_child_source
*/
func (recv *Source) AddChildSource(childSource *Source) {
	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	C.g_source_add_child_source((*C.GSource)(recv.native), c_child_source)

	return
}

// Gets the time to be used when checking this source. The advantage of
// calling this function over calling g_get_monotonic_time() directly is
// that when checking multiple sources, GLib can cache a single value
// instead of having to repeatedly get the system monotonic time.
//
// The time here is the system monotonic time, if available, or some
// other reasonable alternative otherwise.  See g_get_monotonic_time().
/*

C function

g_source_get_time
*/
func (recv *Source) GetTime() int64 {
	retC := C.g_source_get_time((*C.GSource)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Detaches @child_source from @source and destroys it.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
/*

C function

g_source_remove_child_source
*/
func (recv *Source) RemoveChildSource(childSource *Source) {
	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	C.g_source_remove_child_source((*C.GSource)(recv.native), c_child_source)

	return
}
