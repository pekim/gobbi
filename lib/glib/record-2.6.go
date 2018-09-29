// This is a generated file - DO NOT EDIT
// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// Unsupported : g_bookmark_file_get_added : no return generator

// Unsupported : g_bookmark_file_get_app_info : unsupported parameter count : no type generator for guint, guint*

// Unsupported : g_bookmark_file_get_applications : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_modified : no return generator

// Unsupported : g_bookmark_file_get_uris : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_bookmark_file_get_visited : no return generator

// Unsupported : g_bookmark_file_load_from_data : unsupported parameter data : no param type

// Unsupported : g_bookmark_file_set_added : unsupported parameter added : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_app_info : unsupported parameter stamp : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_groups : unsupported parameter groups : no param type

// Unsupported : g_bookmark_file_set_modified : unsupported parameter modified : no type generator for glong, time_t

// Unsupported : g_bookmark_file_set_visited : unsupported parameter visited : no type generator for glong, time_t

// Unsupported : g_bookmark_file_to_data : unsupported parameter length : no type generator for gsize, gsize*

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

// GetIso8601WeekOfYear is a wrapper around the C function g_date_get_iso8601_week_of_year.
func (recv *Date) GetIso8601WeekOfYear() uint32 {
	retC := C.g_date_get_iso8601_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

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

// KeyFileNew is a wrapper around the C function g_key_file_new.
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_key_file_free.
func (recv *KeyFile) Free() {
	C.g_key_file_free((*C.GKeyFile)(recv.native))

	return
}

// GetBoolean is a wrapper around the C function g_key_file_get_boolean.
func (recv *KeyFile) GetBoolean(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_boolean((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_get_boolean_list : unsupported parameter length : no type generator for gsize, gsize*

// GetComment is a wrapper around the C function g_key_file_get_comment.
func (recv *KeyFile) GetComment(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_get_double_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_groups : unsupported parameter length : no type generator for gsize, gsize*

// GetInteger is a wrapper around the C function g_key_file_get_integer.
func (recv *KeyFile) GetInteger(groupName string, key string) (int32, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_get_integer_list : unsupported parameter length : no type generator for gsize, gsize*

// Unsupported : g_key_file_get_keys : unsupported parameter length : no type generator for gsize, gsize*

// GetLocaleString is a wrapper around the C function g_key_file_get_locale_string.
func (recv *KeyFile) GetLocaleString(groupName string, key string, locale string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_locale_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_get_locale_string_list : unsupported parameter length : no type generator for gsize, gsize*

// GetStartGroup is a wrapper around the C function g_key_file_get_start_group.
func (recv *KeyFile) GetStartGroup() string {
	retC := C.g_key_file_get_start_group((*C.GKeyFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetString is a wrapper around the C function g_key_file_get_string.
func (recv *KeyFile) GetString(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_string((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_get_string_list : unsupported parameter length : no type generator for gsize, gsize*

// GetValue is a wrapper around the C function g_key_file_get_value.
func (recv *KeyFile) GetValue(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_value((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// HasGroup is a wrapper around the C function g_key_file_has_group.
func (recv *KeyFile) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.g_key_file_has_group((*C.GKeyFile)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// HasKey is a wrapper around the C function g_key_file_has_key.
func (recv *KeyFile) HasKey(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_has_key((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LoadFromData is a wrapper around the C function g_key_file_load_from_data.
func (recv *KeyFile) LoadFromData(data string, length uint64, flags KeyFileFlags) (bool, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gsize)(length)

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_data((*C.GKeyFile)(recv.native), c_data, c_length, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// LoadFromDataDirs is a wrapper around the C function g_key_file_load_from_data_dirs.
func (recv *KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var c_full_path *C.gchar

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_data_dirs((*C.GKeyFile)(recv.native), c_file, &c_full_path, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goThrowableError
}

// Unsupported : g_key_file_load_from_dirs : unsupported parameter search_dirs : no param type

// LoadFromFile is a wrapper around the C function g_key_file_load_from_file.
func (recv *KeyFile) LoadFromFile(file string, flags KeyFileFlags) (bool, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_file((*C.GKeyFile)(recv.native), c_file, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveComment is a wrapper around the C function g_key_file_remove_comment.
func (recv *KeyFile) RemoveComment(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveGroup is a wrapper around the C function g_key_file_remove_group.
func (recv *KeyFile) RemoveGroup(groupName string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_group((*C.GKeyFile)(recv.native), c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// RemoveKey is a wrapper around the C function g_key_file_remove_key.
func (recv *KeyFile) RemoveKey(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_key((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SetBoolean is a wrapper around the C function g_key_file_set_boolean.
func (recv *KeyFile) SetBoolean(groupName string, key string, value bool) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	C.g_key_file_set_boolean((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Unsupported : g_key_file_set_boolean_list : unsupported parameter list : no param type

// SetComment is a wrapper around the C function g_key_file_set_comment.
func (recv *KeyFile) SetComment(groupName string, key string, comment string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_comment := C.CString(comment)
	defer C.free(unsafe.Pointer(c_comment))

	var cThrowableError *C.GError

	retC := C.g_key_file_set_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, c_comment, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_key_file_set_double_list : unsupported parameter list : no param type

// SetInteger is a wrapper around the C function g_key_file_set_integer.
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	C.g_key_file_set_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Unsupported : g_key_file_set_integer_list : unsupported parameter list : no param type

// SetListSeparator is a wrapper around the C function g_key_file_set_list_separator.
func (recv *KeyFile) SetListSeparator(separator rune) {
	c_separator := (C.gchar)(separator)

	C.g_key_file_set_list_separator((*C.GKeyFile)(recv.native), c_separator)

	return
}

// SetLocaleString is a wrapper around the C function g_key_file_set_locale_string.
func (recv *KeyFile) SetLocaleString(groupName string, key string, locale string, string string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_key_file_set_locale_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, c_string)

	return
}

// Unsupported : g_key_file_set_locale_string_list : unsupported parameter list : no param type

// SetString is a wrapper around the C function g_key_file_set_string.
func (recv *KeyFile) SetString(groupName string, key string, string string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	C.g_key_file_set_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_string)

	return
}

// Unsupported : g_key_file_set_string_list : unsupported parameter list : no param type

// SetValue is a wrapper around the C function g_key_file_set_value.
func (recv *KeyFile) SetValue(groupName string, key string, value string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_key_file_set_value((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

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

// AddGroup is a wrapper around the C function g_option_context_add_group.
func (recv *OptionContext) AddGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(group.ToC())

	C.g_option_context_add_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// AddMainEntries is a wrapper around the C function g_option_context_add_main_entries.
func (recv *OptionContext) AddMainEntries(entries *OptionEntry, translationDomain string) {
	c_entries := (*C.GOptionEntry)(entries.ToC())

	c_translation_domain := C.CString(translationDomain)
	defer C.free(unsafe.Pointer(c_translation_domain))

	C.g_option_context_add_main_entries((*C.GOptionContext)(recv.native), c_entries, c_translation_domain)

	return
}

// Free is a wrapper around the C function g_option_context_free.
func (recv *OptionContext) Free() {
	C.g_option_context_free((*C.GOptionContext)(recv.native))

	return
}

// GetHelpEnabled is a wrapper around the C function g_option_context_get_help_enabled.
func (recv *OptionContext) GetHelpEnabled() bool {
	retC := C.g_option_context_get_help_enabled((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIgnoreUnknownOptions is a wrapper around the C function g_option_context_get_ignore_unknown_options.
func (recv *OptionContext) GetIgnoreUnknownOptions() bool {
	retC := C.g_option_context_get_ignore_unknown_options((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMainGroup is a wrapper around the C function g_option_context_get_main_group.
func (recv *OptionContext) GetMainGroup() *OptionGroup {
	retC := C.g_option_context_get_main_group((*C.GOptionContext)(recv.native))
	retGo := OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_option_context_parse : unsupported parameter argc : no type generator for gint, gint*

// Unsupported : g_option_context_parse_strv : unsupported parameter arguments : no param type

// SetHelpEnabled is a wrapper around the C function g_option_context_set_help_enabled.
func (recv *OptionContext) SetHelpEnabled(helpEnabled bool) {
	c_help_enabled :=
		boolToGboolean(helpEnabled)

	C.g_option_context_set_help_enabled((*C.GOptionContext)(recv.native), c_help_enabled)

	return
}

// SetIgnoreUnknownOptions is a wrapper around the C function g_option_context_set_ignore_unknown_options.
func (recv *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	c_ignore_unknown :=
		boolToGboolean(ignoreUnknown)

	C.g_option_context_set_ignore_unknown_options((*C.GOptionContext)(recv.native), c_ignore_unknown)

	return
}

// SetMainGroup is a wrapper around the C function g_option_context_set_main_group.
func (recv *OptionContext) SetMainGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(group.ToC())

	C.g_option_context_set_main_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// AddEntries is a wrapper around the C function g_option_group_add_entries.
func (recv *OptionGroup) AddEntries(entries *OptionEntry) {
	c_entries := (*C.GOptionEntry)(entries.ToC())

	C.g_option_group_add_entries((*C.GOptionGroup)(recv.native), c_entries)

	return
}

// Free is a wrapper around the C function g_option_group_free.
func (recv *OptionGroup) Free() {
	C.g_option_group_free((*C.GOptionGroup)(recv.native))

	return
}

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc, GOptionErrorFunc

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc, GOptionParseFunc

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc, GTranslateFunc

// SetTranslationDomain is a wrapper around the C function g_option_group_set_translation_domain.
func (recv *OptionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_group_set_translation_domain((*C.GOptionGroup)(recv.native), c_domain)

	return
}

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

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc, GSourceFunc

// Unsupported : g_string_append_printf : unsupported parameter ... : varargs

// Unsupported : g_string_append_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_string_printf : unsupported parameter ... : varargs

// Unsupported : g_string_vprintf : unsupported parameter args : no type generator for va_list, va_list

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

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
