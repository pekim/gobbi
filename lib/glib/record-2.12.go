// This is a generated file - DO NOT EDIT
// +build glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_async_queue_push_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_push_sorted_unlocked : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_sort : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_async_queue_sort_unlocked : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// AddApplication is a wrapper around the C function g_bookmark_file_add_application.
func (recv *BookmarkFile) AddApplication(uri string, name string, exec string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_exec := C.CString(exec)
	defer C.free(unsafe.Pointer(c_exec))

	C.g_bookmark_file_add_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, c_exec)

	return
}

// AddGroup is a wrapper around the C function g_bookmark_file_add_group.
func (recv *BookmarkFile) AddGroup(uri string, group string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	C.g_bookmark_file_add_group((*C.GBookmarkFile)(recv.native), c_uri, c_group)

	return
}

// Free is a wrapper around the C function g_bookmark_file_free.
func (recv *BookmarkFile) Free() {
	C.g_bookmark_file_free((*C.GBookmarkFile)(recv.native))

	return
}

// Unsupported : g_bookmark_file_get_added : no return generator

// Unsupported : g_bookmark_file_get_app_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : g_bookmark_file_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// GetDescription is a wrapper around the C function g_bookmark_file_get_description.
func (recv *BookmarkFile) GetDescription(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_description((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bookmark_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// GetIcon is a wrapper around the C function g_bookmark_file_get_icon.
func (recv *BookmarkFile) GetIcon(uri string) (bool, string, string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_href *C.gchar

	var c_mime_type *C.gchar

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_icon((*C.GBookmarkFile)(recv.native), c_uri, &c_href, &c_mime_type, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	href := C.GoString(c_href)
	defer C.free(unsafe.Pointer(c_href))

	mimeType := C.GoString(c_mime_type)
	defer C.free(unsafe.Pointer(c_mime_type))

	return retGo, href, mimeType, goThrowableError
}

// GetIsPrivate is a wrapper around the C function g_bookmark_file_get_is_private.
func (recv *BookmarkFile) GetIsPrivate(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_is_private((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// GetMimeType is a wrapper around the C function g_bookmark_file_get_mime_type.
func (recv *BookmarkFile) GetMimeType(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_mime_type((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bookmark_file_get_modified : no return generator

// GetSize is a wrapper around the C function g_bookmark_file_get_size.
func (recv *BookmarkFile) GetSize() int32 {
	retC := C.g_bookmark_file_get_size((*C.GBookmarkFile)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTitle is a wrapper around the C function g_bookmark_file_get_title.
func (recv *BookmarkFile) GetTitle(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_title((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bookmark_file_get_uris : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_visited : no return generator

// HasApplication is a wrapper around the C function g_bookmark_file_has_application.
func (recv *BookmarkFile) HasApplication(uri string, name string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_has_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// HasGroup is a wrapper around the C function g_bookmark_file_has_group.
func (recv *BookmarkFile) HasGroup(uri string, group string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_has_group((*C.GBookmarkFile)(recv.native), c_uri, c_group, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// HasItem is a wrapper around the C function g_bookmark_file_has_item.
func (recv *BookmarkFile) HasItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_bookmark_file_has_item((*C.GBookmarkFile)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_bookmark_file_load_from_data : unsupported parameter data : no param type

// LoadFromDataDirs is a wrapper around the C function g_bookmark_file_load_from_data_dirs.
func (recv *BookmarkFile) LoadFromDataDirs(file string) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var c_full_path *C.gchar

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_data_dirs((*C.GBookmarkFile)(recv.native), c_file, &c_full_path, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goThrowableError
}

// LoadFromFile is a wrapper around the C function g_bookmark_file_load_from_file.
func (recv *BookmarkFile) LoadFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_file((*C.GBookmarkFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// MoveItem is a wrapper around the C function g_bookmark_file_move_item.
func (recv *BookmarkFile) MoveItem(oldUri string, newUri string) (bool, error) {
	c_old_uri := C.CString(oldUri)
	defer C.free(unsafe.Pointer(c_old_uri))

	c_new_uri := C.CString(newUri)
	defer C.free(unsafe.Pointer(c_new_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_move_item((*C.GBookmarkFile)(recv.native), c_old_uri, c_new_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveApplication is a wrapper around the C function g_bookmark_file_remove_application.
func (recv *BookmarkFile) RemoveApplication(uri string, name string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveGroup is a wrapper around the C function g_bookmark_file_remove_group.
func (recv *BookmarkFile) RemoveGroup(uri string, group string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_group((*C.GBookmarkFile)(recv.native), c_uri, c_group, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveItem is a wrapper around the C function g_bookmark_file_remove_item.
func (recv *BookmarkFile) RemoveItem(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_item((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bookmark_file_set_added : unsupported parameter added : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_app_info : unsupported parameter stamp : no type generator for glong, time_t

// SetDescription is a wrapper around the C function g_bookmark_file_set_description.
func (recv *BookmarkFile) SetDescription(uri string, description string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_bookmark_file_set_description((*C.GBookmarkFile)(recv.native), c_uri, c_description)

	return
}

// Unsupported : g_bookmark_file_set_groups : unsupported parameter groups : no param type

// SetIcon is a wrapper around the C function g_bookmark_file_set_icon.
func (recv *BookmarkFile) SetIcon(uri string, href string, mimeType string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_href := C.CString(href)
	defer C.free(unsafe.Pointer(c_href))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.g_bookmark_file_set_icon((*C.GBookmarkFile)(recv.native), c_uri, c_href, c_mime_type)

	return
}

// SetIsPrivate is a wrapper around the C function g_bookmark_file_set_is_private.
func (recv *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_is_private :=
		boolToGboolean(isPrivate)

	C.g_bookmark_file_set_is_private((*C.GBookmarkFile)(recv.native), c_uri, c_is_private)

	return
}

// SetMimeType is a wrapper around the C function g_bookmark_file_set_mime_type.
func (recv *BookmarkFile) SetMimeType(uri string, mimeType string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.g_bookmark_file_set_mime_type((*C.GBookmarkFile)(recv.native), c_uri, c_mime_type)

	return
}

// Unsupported : g_bookmark_file_set_modified : unsupported parameter modified : no type generator for glong, time_t

// SetTitle is a wrapper around the C function g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_bookmark_file_set_title((*C.GBookmarkFile)(recv.native), c_uri, c_title)

	return
}

// Unsupported : g_bookmark_file_set_visited : unsupported parameter visited : no type generator for glong, time_t

// Unsupported : g_bookmark_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// ToFile is a wrapper around the C function g_bookmark_file_to_file.
func (recv *BookmarkFile) ToFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_to_file((*C.GBookmarkFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

// Unsupported : g_bytes_get_data : unsupported parameter size : no type generator for gsize, gsize*

// Unsupported : g_bytes_unref_to_array : no return type

// Unsupported : g_bytes_unref_to_data : unsupported parameter size : no type generator for gsize, gsize*

// Unsupported : g_checksum_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// Unsupported : g_checksum_update : unsupported parameter data : no param type

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_cond_wait_until : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_date_set_time_t : unsupported parameter timet : no type generator for glong, time_t

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer, tm*

// Unsupported : g_date_time_get_ymd : unsupported parameter year : no type generator for gint, gint*

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer, gpointer*

// Unsupported : g_hmac_get_digest : unsupported parameter buffer : no type generator for guint8, guint8*

// Unsupported : g_hmac_update : unsupported parameter data : no param type

// Unsupported : g_hook_list_marshal : unsupported parameter marshaller : no type generator for HookMarshaller, GHookMarshaller

// Unsupported : g_hook_list_marshal_check : unsupported parameter marshaller : no type generator for HookCheckMarshaller, GHookCheckMarshaller

// Blacklisted : GIConv

// Blacklisted : GIOChannel

// Unsupported : g_key_file_get_boolean_list : unsupported parameter length : no type generator for gsize, gsize*

// GetDouble is a wrapper around the C function g_key_file_get_double.
func (recv *KeyFile) GetDouble(groupName string, key string) (float64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_double((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (float64)(retC)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_get_double_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_integer_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_keys : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_locale_string_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_string_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_load_from_dirs : unsupported parameter search_dirs : no param type

// Unsupported : g_key_file_set_boolean_list : unsupported parameter list : no param type

// SetDouble is a wrapper around the C function g_key_file_set_double.
func (recv *KeyFile) SetDouble(groupName string, key string, value float64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	C.g_key_file_set_double((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Unsupported : g_key_file_set_double_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_integer_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_locale_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_set_string_list : unsupported parameter list : no param type

// Unsupported : g_key_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_main_context_check : unsupported parameter fds : no param type

// Unsupported : g_main_context_get_poll_func : no return generator

// Unsupported : g_main_context_invoke : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_invoke_full : unsupported parameter function : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_main_context_prepare : unsupported parameter priority : no type generator for gint, gint*

// Unsupported : g_main_context_query : unsupported parameter timeout_ : no type generator for gint, gint*

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc, GPollFunc

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex, GMutex*

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_markup_parse_context_get_position : unsupported parameter line_number : no type generator for gint, gint*

// Unsupported : g_match_info_fetch_all : no return type

// Unsupported : g_match_info_fetch_named_pos : unsupported parameter start_pos : no type generator for gint, gint*

// Unsupported : g_match_info_fetch_pos : unsupported parameter start_pos : no type generator for gint, gint*

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc, GNodeForeachFunc

// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc, GCopyFunc

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc, GNodeTraverseFunc

// Unsupported : g_once_impl : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// GetDescription is a wrapper around the C function g_option_context_get_description.
func (recv *OptionContext) GetDescription() string {
	retC := C.g_option_context_get_description((*C.GOptionContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSummary is a wrapper around the C function g_option_context_get_summary.
func (recv *OptionContext) GetSummary() string {
	retC := C.g_option_context_get_summary((*C.GOptionContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_option_context_parse : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : g_option_context_parse_strv : unsupported parameter arguments : no param type

// SetDescription is a wrapper around the C function g_option_context_set_description.
func (recv *OptionContext) SetDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_option_context_set_description((*C.GOptionContext)(recv.native), c_description)

	return
}

// SetSummary is a wrapper around the C function g_option_context_set_summary.
func (recv *OptionContext) SetSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_option_context_set_summary((*C.GOptionContext)(recv.native), c_summary)

	return
}

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// SetTranslationDomain is a wrapper around the C function g_option_context_set_translation_domain.
func (recv *OptionContext) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_context_set_translation_domain((*C.GOptionContext)(recv.native), c_domain)

	return
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc, GOptionErrorFunc

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc, GOptionParseFunc

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// Unsupported : g_queue_find_custom : unsupported parameter func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_queue_free_full : unsupported parameter free_func : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_queue_insert_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_rand_set_seed_array : unsupported parameter seed : no type generator for guint32, const guint32*

// Unsupported : g_regex_match : unsupported parameter match_info : record with indirection level of 2

// Unsupported : g_regex_match_all : unsupported parameter match_info : record with indirection level of 2

// Unsupported : g_regex_match_all_full : unsupported parameter string : no param type

// Unsupported : g_regex_match_full : unsupported parameter string : no param type

// Unsupported : g_regex_replace : unsupported parameter string : no param type

// Unsupported : g_regex_replace_eval : unsupported parameter string : no param type

// Unsupported : g_regex_replace_literal : unsupported parameter string : no param type

// Unsupported : g_regex_split : no return type

// Unsupported : g_regex_split_full : unsupported parameter string : no param type

// Unsupported : g_scanner_cur_value : no return generator

// Unsupported : g_scanner_error : unsupported parameter ... : varargs

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc, GHFunc

// Unsupported : g_scanner_warn : unsupported parameter ... : varargs

// Unsupported : g_sequence_foreach : unsupported parameter func : no type generator for Func, GFunc

// Unsupported : g_sequence_insert_sorted : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_insert_sorted_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_lookup : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_lookup_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_search : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_search_iter : unsupported parameter iter_cmp : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// Unsupported : g_sequence_sort : unsupported parameter cmp_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unsupported : g_sequence_sort_iter : unsupported parameter cmp_func : no type generator for SequenceIterCompareFunc, GSequenceIterCompareFunc

// IsDestroyed is a wrapper around the C function g_source_is_destroyed.
func (recv *Source) IsDestroyed() bool {
	retC := C.g_source_is_destroyed((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc, GSourceFunc

// SetFuncs is a wrapper around the C function g_source_set_funcs.
func (recv *Source) SetFuncs(funcs *SourceFuncs) {
	c_funcs := (*C.GSourceFuncs)(funcs.ToC())

	C.g_source_set_funcs((*C.GSource)(recv.native), c_funcs)

	return
}

// Unsupported : g_string_append_printf : unsupported parameter ... : varargs

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_string_printf : unsupported parameter ... : varargs

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// ToIso8601 is a wrapper around the C function g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() string {
	retC := C.g_time_val_to_iso8601((*C.GTimeVal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_time_zone_adjust_time : unsupported parameter time_ : no type generator for gint64, gint64*

// Unsupported : g_timer_elapsed : unsupported parameter microseconds : no type generator for gulong, gulong*

// Unsupported : g_tree_foreach : unsupported parameter func : no type generator for TraverseFunc, GTraverseFunc

// Unsupported : g_tree_lookup_extended : unsupported parameter orig_key : no type generator for gpointer, gpointer*

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc, GTraverseFunc

// Blacklisted : GVariant

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_add : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_parsed : unsupported parameter ... : varargs

// Unsupported : g_variant_builder_add_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_builder_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_builder_init : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_builder_open : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_end : return type : Blacklisted record : GVariant

// Unsupported : g_variant_dict_init : unsupported parameter from_asv : Blacklisted record : GVariant

// Unsupported : g_variant_dict_insert : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_insert_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_dict_lookup : unsupported parameter ... : varargs

// Unsupported : g_variant_dict_lookup_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Unsupported : g_variant_iter_init : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_variant_iter_loop : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next : unsupported parameter ... : varargs

// Unsupported : g_variant_iter_next_value : return type : Blacklisted record : GVariant

// Blacklisted : GVariantType
