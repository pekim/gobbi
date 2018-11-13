// This is a generated file - DO NOT EDIT
// +build glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Resets the state of the @checksum back to its initial state.
/*

C function : g_checksum_reset
*/
func (recv *Checksum) Reset() {
	C.g_checksum_reset((*C.GChecksum)(recv.native))

	return
}

// Returns the user_data associated with @context.
//
// This will either be the user_data that was provided to
// g_markup_parse_context_new() or to the most recent call
// of g_markup_parse_context_push().
/*

C function : g_markup_parse_context_get_user_data
*/
func (recv *MarkupParseContext) GetUserData() uintptr {
	retC := C.g_markup_parse_context_get_user_data((*C.GMarkupParseContext)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Completes the process of a temporary sub-parser redirection.
//
// This function exists to collect the user_data allocated by a
// matching call to g_markup_parse_context_push(). It must be called
// in the end_element handler corresponding to the start_element
// handler during which g_markup_parse_context_push() was called.
// You must not call this function from the error callback -- the
// @user_data is provided directly to the callback in that case.
//
// This function is not intended to be directly called by users
// interested in invoking subparsers. Instead, it is intended to
// be used by the subparsers themselves to implement a higher-level
// interface.
/*

C function : g_markup_parse_context_pop
*/
func (recv *MarkupParseContext) Pop() uintptr {
	retC := C.g_markup_parse_context_pop((*C.GMarkupParseContext)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Temporarily redirects markup data to a sub-parser.
//
// This function may only be called from the start_element handler of
// a #GMarkupParser. It must be matched with a corresponding call to
// g_markup_parse_context_pop() in the matching end_element handler
// (except in the case that the parser aborts due to an error).
//
// All tags, text and other data between the matching tags is
// redirected to the subparser given by @parser. @user_data is used
// as the user_data for that parser. @user_data is also passed to the
// error callback in the event that an error occurs. This includes
// errors that occur in subparsers of the subparser.
//
// The end tag matching the start tag for which this call was made is
// handled by the previous parser (which is given its own user_data)
// which is why g_markup_parse_context_pop() is provided to allow "one
// last access" to the @user_data provided to this function. In the
// case of error, the @user_data provided here is passed directly to
// the error callback of the subparser and g_markup_parse_context_pop()
// should not be called. In either case, if @user_data was allocated
// then it ought to be freed from both of these locations.
//
// This function is not intended to be directly called by users
// interested in invoking subparsers. Instead, it is intended to be
// used by the subparsers themselves to implement a higher-level
// interface.
//
// As an example, see the following implementation of a simple
// parser that counts the number of tags encountered.
//
// |[<!-- language="C" -->
// typedef struct
// {
// gint tag_count;
// } CounterData;
//
// static void
// counter_start_element (GMarkupParseContext  *context,
// const gchar          *element_name,
// const gchar         **attribute_names,
// const gchar         **attribute_values,
// gpointer              user_data,
// GError              **error)
// {
// CounterData *data = user_data;
//
// data->tag_count++;
// }
//
// static void
// counter_error (GMarkupParseContext *context,
// GError              *error,
// gpointer             user_data)
// {
// CounterData *data = user_data;
//
// g_slice_free (CounterData, data);
// }
//
// static GMarkupParser counter_subparser =
// {
// counter_start_element,
// NULL,
// NULL,
// NULL,
// counter_error
// };
// ]|
//
// In order to allow this parser to be easily used as a subparser, the
// following interface is provided:
//
// |[<!-- language="C" -->
// void
// start_counting (GMarkupParseContext *context)
// {
// CounterData *data = g_slice_new (CounterData);
//
// data->tag_count = 0;
// g_markup_parse_context_push (context, &counter_subparser, data);
// }
//
// gint
// end_counting (GMarkupParseContext *context)
// {
// CounterData *data = g_markup_parse_context_pop (context);
// int result;
//
// result = data->tag_count;
// g_slice_free (CounterData, data);
//
// return result;
// }
// ]|
//
// The subparser would then be used as follows:
//
// |[<!-- language="C" -->
// static void start_element (context, element_name, ...)
// {
// if (strcmp (element_name, "count-these") == 0)
// start_counting (context);
//
// else, handle other tags...
// }
//
// static void end_element (context, element_name, ...)
// {
// if (strcmp (element_name, "count-these") == 0)
// g_print ("Counted %d tags\n", end_counting (context));
//
// else, handle other tags...
// }
// ]|
/*

C function : g_markup_parse_context_push
*/
func (recv *MarkupParseContext) Push(parser *MarkupParser, userData uintptr) {
	c_parser := (*C.GMarkupParser)(C.NULL)
	if parser != nil {
		c_parser = (*C.GMarkupParser)(parser.ToC())
	}

	c_user_data := (C.gpointer)(userData)

	C.g_markup_parse_context_push((*C.GMarkupParseContext)(recv.native), c_parser, c_user_data)

	return
}
