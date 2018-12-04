// This is a generated file - DO NOT EDIT
// +build glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// g_list_free_full : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func
// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc (GSourceFunc) for param function

// g_slist_free_full : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func
// Unsupported : g_sequence_lookup : unsupported parameter cmp_func : no type generator for CompareDataFunc (GCompareDataFunc) for param cmp_func

// Unsupported : g_sequence_lookup_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc (GSequenceIterCompareFunc) for param iter_cmp

// AddChildSource is a wrapper around the C function g_source_add_child_source.
func (recv *Source) AddChildSource(childSource *Source) {
	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	C.g_source_add_child_source((*C.GSource)(recv.native), c_child_source)

	return
}

// GetTime is a wrapper around the C function g_source_get_time.
func (recv *Source) GetTime() int64 {
	retC := C.g_source_get_time((*C.GSource)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// RemoveChildSource is a wrapper around the C function g_source_remove_child_source.
func (recv *Source) RemoveChildSource(childSource *Source) {
	c_child_source := (*C.GSource)(C.NULL)
	if childSource != nil {
		c_child_source = (*C.GSource)(childSource.ToC())
	}

	C.g_source_remove_child_source((*C.GSource)(recv.native), c_child_source)

	return
}
