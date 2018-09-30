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

// Unsupported : g_bookmark_file_get_app_info : unsupported parameter stamp : no type generator for glong, time_t*

// Unsupported : g_bookmark_file_get_applications : no return type

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

// Unsupported : g_bookmark_file_get_groups : no return type

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

// Unsupported : g_bookmark_file_get_uris : no return type

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

// Unsupported : g_bookmark_file_to_data : no return type

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

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

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

// Unsupported : g_key_file_get_double_list : no return type

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

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

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

// IsDestroyed is a wrapper around the C function g_source_is_destroyed.
func (recv *Source) IsDestroyed() bool {
	retC := C.g_source_is_destroyed((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFuncs is a wrapper around the C function g_source_set_funcs.
func (recv *Source) SetFuncs(funcs *SourceFuncs) {
	c_funcs := (*C.GSourceFuncs)(funcs.ToC())

	C.g_source_set_funcs((*C.GSource)(recv.native), c_funcs)

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// ToIso8601 is a wrapper around the C function g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() string {
	retC := C.g_time_val_to_iso8601((*C.GTimeVal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant
