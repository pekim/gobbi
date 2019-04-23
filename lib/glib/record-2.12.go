// This is a generated file - DO NOT EDIT
// +build glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Blacklisted : g_bookmark_file_new

// Blacklisted : g_bookmark_file_add_application

// Blacklisted : g_bookmark_file_add_group

// Blacklisted : g_bookmark_file_free

// Blacklisted : g_bookmark_file_get_added

// Blacklisted : g_bookmark_file_get_app_info

// Blacklisted : g_bookmark_file_get_applications

// Blacklisted : g_bookmark_file_get_description

// Blacklisted : g_bookmark_file_get_groups

// Blacklisted : g_bookmark_file_get_icon

// Blacklisted : g_bookmark_file_get_is_private

// Blacklisted : g_bookmark_file_get_mime_type

// Blacklisted : g_bookmark_file_get_modified

// Blacklisted : g_bookmark_file_get_size

// GetTitle is a wrapper around the C function g_bookmark_file_get_title.
func (recv *BookmarkFile) GetTitle(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_title((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Blacklisted : g_bookmark_file_get_uris

// Blacklisted : g_bookmark_file_get_visited

// Blacklisted : g_bookmark_file_has_application

// Blacklisted : g_bookmark_file_has_group

// Blacklisted : g_bookmark_file_has_item

// Blacklisted : g_bookmark_file_load_from_data

// Blacklisted : g_bookmark_file_load_from_data_dirs

// Blacklisted : g_bookmark_file_load_from_file

// Blacklisted : g_bookmark_file_move_item

// Blacklisted : g_bookmark_file_remove_application

// Blacklisted : g_bookmark_file_remove_group

// Blacklisted : g_bookmark_file_remove_item

// Blacklisted : g_bookmark_file_set_added

// Blacklisted : g_bookmark_file_set_app_info

// Blacklisted : g_bookmark_file_set_description

// Blacklisted : g_bookmark_file_set_groups

// Blacklisted : g_bookmark_file_set_icon

// Blacklisted : g_bookmark_file_set_is_private

// Blacklisted : g_bookmark_file_set_mime_type

// Blacklisted : g_bookmark_file_set_modified

// SetTitle is a wrapper around the C function g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_bookmark_file_set_title((*C.GBookmarkFile)(recv.native), c_uri, c_title)

	return
}

// Blacklisted : g_bookmark_file_set_visited

// Unsupported : g_bookmark_file_to_data : array return type :

// Blacklisted : g_bookmark_file_to_file

// Blacklisted : g_hash_table_remove_all

// Blacklisted : g_hash_table_steal_all

// Blacklisted : g_key_file_get_double

// Unsupported : g_key_file_get_double_list : array return type :

// Blacklisted : g_key_file_set_double

// Blacklisted : g_key_file_set_double_list

// Blacklisted : g_option_context_get_description

// Blacklisted : g_option_context_get_summary

// Blacklisted : g_option_context_set_description

// Blacklisted : g_option_context_set_summary

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// Blacklisted : g_option_context_set_translation_domain

// Blacklisted : g_source_is_destroyed

// Blacklisted : g_source_set_funcs

// Blacklisted : g_time_val_from_iso8601

// Blacklisted : g_time_val_to_iso8601
