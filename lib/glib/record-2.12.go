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

// BookmarkFileNew is a wrapper around the C function g_bookmark_file_new.
func BookmarkFileNew() *BookmarkFile {
	retC := C.g_bookmark_file_new()
	retGo := BookmarkFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetAdded is a wrapper around the C function g_bookmark_file_get_added.
func (recv *BookmarkFile) GetAdded(uri string) (int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_added((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAppInfo is a wrapper around the C function g_bookmark_file_get_app_info.
func (recv *BookmarkFile) GetAppInfo(uri string, name string) (bool, string, uint32, int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var c_exec *C.gchar

	var c_count C.guint

	var c_stamp C.time_t

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_app_info((*C.GBookmarkFile)(recv.native), c_uri, c_name, &c_exec, &c_count, &c_stamp, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	exec := C.GoString(c_exec)
	defer C.free(unsafe.Pointer(c_exec))

	count := (uint32)(c_count)

	stamp := (int64)(c_stamp)

	return retGo, exec, count, stamp, goError
}

// GetApplications is a wrapper around the C function g_bookmark_file_get_applications.
func (recv *BookmarkFile) GetApplications(uri string) ([]string, uint64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_applications((*C.GBookmarkFile)(recv.native), c_uri, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetDescription is a wrapper around the C function g_bookmark_file_get_description.
func (recv *BookmarkFile) GetDescription(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_description((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
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

// GetGroups is a wrapper around the C function g_bookmark_file_get_groups.
func (recv *BookmarkFile) GetGroups(uri string) ([]string, uint64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_groups((*C.GBookmarkFile)(recv.native), c_uri, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetIcon is a wrapper around the C function g_bookmark_file_get_icon.
func (recv *BookmarkFile) GetIcon(uri string) (bool, string, string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var c_href *C.gchar

	var c_mime_type *C.gchar

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_icon((*C.GBookmarkFile)(recv.native), c_uri, &c_href, &c_mime_type, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	href := C.GoString(c_href)
	defer C.free(unsafe.Pointer(c_href))

	mimeType := C.GoString(c_mime_type)
	defer C.free(unsafe.Pointer(c_mime_type))

	return retGo, href, mimeType, goError
}

// GetIsPrivate is a wrapper around the C function g_bookmark_file_get_is_private.
func (recv *BookmarkFile) GetIsPrivate(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_is_private((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetMimeType is a wrapper around the C function g_bookmark_file_get_mime_type.
func (recv *BookmarkFile) GetMimeType(uri string) (string, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_mime_type((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
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

// GetModified is a wrapper around the C function g_bookmark_file_get_modified.
func (recv *BookmarkFile) GetModified(uri string) (int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_modified((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetUris is a wrapper around the C function g_bookmark_file_get_uris.
func (recv *BookmarkFile) GetUris() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_bookmark_file_get_uris((*C.GBookmarkFile)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetVisited is a wrapper around the C function g_bookmark_file_get_visited.
func (recv *BookmarkFile) GetVisited(uri string) (int64, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_get_visited((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasApplication is a wrapper around the C function g_bookmark_file_has_application.
func (recv *BookmarkFile) HasApplication(uri string, name string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_has_application((*C.GBookmarkFile)(recv.native), c_uri, c_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasItem is a wrapper around the C function g_bookmark_file_has_item.
func (recv *BookmarkFile) HasItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_bookmark_file_has_item((*C.GBookmarkFile)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// LoadFromData is a wrapper around the C function g_bookmark_file_load_from_data.
func (recv *BookmarkFile) LoadFromData(data []uint8) (bool, error) {
	c_data_array := make([]C.guint8, len(data), len(data))
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.gchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_data((*C.GBookmarkFile)(recv.native), c_data, c_length, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadFromDataDirs is a wrapper around the C function g_bookmark_file_load_from_data_dirs.
func (recv *BookmarkFile) LoadFromDataDirs(file string) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var c_full_path *C.gchar

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_data_dirs((*C.GBookmarkFile)(recv.native), c_file, &c_full_path, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goError
}

// LoadFromFile is a wrapper around the C function g_bookmark_file_load_from_file.
func (recv *BookmarkFile) LoadFromFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_file((*C.GBookmarkFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveItem is a wrapper around the C function g_bookmark_file_remove_item.
func (recv *BookmarkFile) RemoveItem(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_remove_item((*C.GBookmarkFile)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAdded is a wrapper around the C function g_bookmark_file_set_added.
func (recv *BookmarkFile) SetAdded(uri string, added int64) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_added := (C.time_t)(added)

	C.g_bookmark_file_set_added((*C.GBookmarkFile)(recv.native), c_uri, c_added)

	return
}

// SetAppInfo is a wrapper around the C function g_bookmark_file_set_app_info.
func (recv *BookmarkFile) SetAppInfo(uri string, name string, exec string, count int32, stamp int64) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_exec := C.CString(exec)
	defer C.free(unsafe.Pointer(c_exec))

	c_count := (C.gint)(count)

	c_stamp := (C.time_t)(stamp)

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_set_app_info((*C.GBookmarkFile)(recv.native), c_uri, c_name, c_exec, c_count, c_stamp, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetDescription is a wrapper around the C function g_bookmark_file_set_description.
func (recv *BookmarkFile) SetDescription(uri string, description string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_bookmark_file_set_description((*C.GBookmarkFile)(recv.native), c_uri, c_description)

	return
}

// SetGroups is a wrapper around the C function g_bookmark_file_set_groups.
func (recv *BookmarkFile) SetGroups(uri string, groups []string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_groups_array := make([]*C.gchar, len(groups), len(groups))
	for i, item := range groups {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_groups_array[i] = c
	}
	c_groups_arrayPtr := &c_groups_array[0]
	c_groups := (**C.gchar)(unsafe.Pointer(c_groups_arrayPtr))

	c_length := (C.gsize)(len(groups))

	C.g_bookmark_file_set_groups((*C.GBookmarkFile)(recv.native), c_uri, c_groups, c_length)

	return
}

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

// SetModified is a wrapper around the C function g_bookmark_file_set_modified.
func (recv *BookmarkFile) SetModified(uri string, modified int64) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_modified := (C.time_t)(modified)

	C.g_bookmark_file_set_modified((*C.GBookmarkFile)(recv.native), c_uri, c_modified)

	return
}

// SetTitle is a wrapper around the C function g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_bookmark_file_set_title((*C.GBookmarkFile)(recv.native), c_uri, c_title)

	return
}

// SetVisited is a wrapper around the C function g_bookmark_file_set_visited.
func (recv *BookmarkFile) SetVisited(uri string, visited int64) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_visited := (C.time_t)(visited)

	C.g_bookmark_file_set_visited((*C.GBookmarkFile)(recv.native), c_uri, c_visited)

	return
}

// Unsupported : g_bookmark_file_to_data : array return type :

// ToFile is a wrapper around the C function g_bookmark_file_to_file.
func (recv *BookmarkFile) ToFile(filename string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_to_file((*C.GBookmarkFile)(recv.native), c_filename, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HashTableRemoveAll is a wrapper around the C function g_hash_table_remove_all.
func HashTableRemoveAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_remove_all(c_hash_table)

	return
}

// HashTableStealAll is a wrapper around the C function g_hash_table_steal_all.
func HashTableStealAll(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_steal_all(c_hash_table)

	return
}

// GetDouble is a wrapper around the C function g_key_file_get_double.
func (recv *KeyFile) GetDouble(groupName string, key string) (float64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_double((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (float64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_key_file_get_double_list : array return type :

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

// SetDoubleList is a wrapper around the C function g_key_file_set_double_list.
func (recv *KeyFile) SetDoubleList(groupName string, key string, list []float64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list_array := make([]C.gdouble, len(list), len(list))
	for i, item := range list {
		c := (C.gdouble)(item)
		c_list_array[i] = c
	}
	c_list_arrayPtr := &c_list_array[0]
	c_list := (*C.gdouble)(unsafe.Pointer(c_list_arrayPtr))

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_double_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_list, c_length)

	return
}

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

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// SetTranslationDomain is a wrapper around the C function g_option_context_set_translation_domain.
func (recv *OptionContext) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_context_set_translation_domain((*C.GOptionContext)(recv.native), c_domain)

	return
}

// IsDestroyed is a wrapper around the C function g_source_is_destroyed.
func (recv *Source) IsDestroyed() bool {
	retC := C.g_source_is_destroyed((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFuncs is a wrapper around the C function g_source_set_funcs.
func (recv *Source) SetFuncs(funcs *SourceFuncs) {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	C.g_source_set_funcs((*C.GSource)(recv.native), c_funcs)

	return
}

// TimeValFromIso8601 is a wrapper around the C function g_time_val_from_iso8601.
func TimeValFromIso8601(isoDate string) (bool, *TimeVal) {
	c_iso_date := C.CString(isoDate)
	defer C.free(unsafe.Pointer(c_iso_date))

	var c_time_ C.GTimeVal

	retC := C.g_time_val_from_iso8601(c_iso_date, &c_time_)
	retGo := retC == C.TRUE

	time := TimeValNewFromC(unsafe.Pointer(&c_time_))

	return retGo, time
}

// ToIso8601 is a wrapper around the C function g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() string {
	retC := C.g_time_val_to_iso8601((*C.GTimeVal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
