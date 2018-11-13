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

// Adds the application with @name and @exec to the list of
// applications that have registered a bookmark for @uri into
// @bookmark.
//
// Every bookmark inside a #GBookmarkFile must have at least an
// application registered.  Each application must provide a name, a
// command line useful for launching the bookmark, the number of times
// the bookmark has been registered by the application and the last
// time the application registered this bookmark.
//
// If @name is %NULL, the name of the application will be the
// same returned by g_get_application_name(); if @exec is %NULL, the
// command line will be a composition of the program name as
// returned by g_get_prgname() and the "\%u" modifier, which will be
// expanded to the bookmark's URI.
//
// This function will automatically take care of updating the
// registrations count and timestamping in case an application
// with the same @name had already registered a bookmark for
// @uri inside @bookmark.
//
// If no bookmark for @uri is found, one is created.
/*

C function

g_bookmark_file_add_application
*/
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

// Adds @group to the list of groups to which the bookmark for @uri
// belongs to.
//
// If no bookmark for @uri is found then it is created.
/*

C function

g_bookmark_file_add_group
*/
func (recv *BookmarkFile) AddGroup(uri string, group string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_group := C.CString(group)
	defer C.free(unsafe.Pointer(c_group))

	C.g_bookmark_file_add_group((*C.GBookmarkFile)(recv.native), c_uri, c_group)

	return
}

// Frees a #GBookmarkFile.
/*

C function

g_bookmark_file_free
*/
func (recv *BookmarkFile) Free() {
	C.g_bookmark_file_free((*C.GBookmarkFile)(recv.native))

	return
}

// Unsupported : g_bookmark_file_get_added : no return generator

// Unsupported : g_bookmark_file_get_app_info : unsupported parameter stamp : no type generator for glong (time_t*) for param stamp

// Unsupported : g_bookmark_file_get_applications : no return type

// Retrieves the description of the bookmark for @uri.
//
// In the event the URI cannot be found, %NULL is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
/*

C function

g_bookmark_file_get_description
*/
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

// Gets the icon of the bookmark for @uri.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
/*

C function

g_bookmark_file_get_icon
*/
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

// Gets whether the private flag of the bookmark for @uri is set.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.  In the
// event that the private flag cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_INVALID_VALUE.
/*

C function

g_bookmark_file_get_is_private
*/
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

// Retrieves the MIME type of the resource pointed by @uri.
//
// In the event the URI cannot be found, %NULL is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.  In the
// event that the MIME type cannot be found, %NULL is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_INVALID_VALUE.
/*

C function

g_bookmark_file_get_mime_type
*/
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

// Gets the number of bookmarks inside @bookmark.
/*

C function

g_bookmark_file_get_size
*/
func (recv *BookmarkFile) GetSize() int32 {
	retC := C.g_bookmark_file_get_size((*C.GBookmarkFile)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the title of the bookmark for @uri.
//
// If @uri is %NULL, the title of @bookmark is returned.
//
// In the event the URI cannot be found, %NULL is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
/*

C function

g_bookmark_file_get_title
*/
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

// Checks whether the bookmark for @uri inside @bookmark has been
// registered by application @name.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
/*

C function

g_bookmark_file_has_application
*/
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

// Checks whether @group appears in the list of groups to which
// the bookmark for @uri belongs to.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
/*

C function

g_bookmark_file_has_group
*/
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

// Looks whether the desktop bookmark has an item with its URI set to @uri.
/*

C function

g_bookmark_file_has_item
*/
func (recv *BookmarkFile) HasItem(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_bookmark_file_has_item((*C.GBookmarkFile)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// Loads a bookmark file from memory into an empty #GBookmarkFile
// structure.  If the object cannot be created then @error is set to a
// #GBookmarkFileError.
/*

C function

g_bookmark_file_load_from_data
*/
func (recv *BookmarkFile) LoadFromData(data []uint8) (bool, error) {
	c_data := &data[0]

	c_length := (C.gsize)(len(data))

	var cThrowableError *C.GError

	retC := C.g_bookmark_file_load_from_data((*C.GBookmarkFile)(recv.native), (*C.gchar)(unsafe.Pointer(c_data)), c_length, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This function looks for a desktop bookmark file named @file in the
// paths returned from g_get_user_data_dir() and g_get_system_data_dirs(),
// loads the file into @bookmark and returns the file's full path in
// @full_path.  If the file could not be loaded then an %error is
// set to either a #GFileError or #GBookmarkFileError.
/*

C function

g_bookmark_file_load_from_data_dirs
*/
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

// Loads a desktop bookmark file into an empty #GBookmarkFile structure.
// If the file could not be loaded then @error is set to either a #GFileError
// or #GBookmarkFileError.
/*

C function

g_bookmark_file_load_from_file
*/
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

// Changes the URI of a bookmark item from @old_uri to @new_uri.  Any
// existing bookmark for @new_uri will be overwritten.  If @new_uri is
// %NULL, then the bookmark is removed.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
/*

C function

g_bookmark_file_move_item
*/
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

// Removes application registered with @name from the list of applications
// that have registered a bookmark for @uri inside @bookmark.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
// In the event that no application with name @app_name has registered
// a bookmark for @uri,  %FALSE is returned and error is set to
// #G_BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED.
/*

C function

g_bookmark_file_remove_application
*/
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

// Removes @group from the list of groups to which the bookmark
// for @uri belongs to.
//
// In the event the URI cannot be found, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
// In the event no group was defined, %FALSE is returned and
// @error is set to #G_BOOKMARK_FILE_ERROR_INVALID_VALUE.
/*

C function

g_bookmark_file_remove_group
*/
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

// Removes the bookmark for @uri from the bookmark file @bookmark.
/*

C function

g_bookmark_file_remove_item
*/
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

// Unsupported : g_bookmark_file_set_added : unsupported parameter added : no type generator for glong (time_t) for param added

// Unsupported : g_bookmark_file_set_app_info : unsupported parameter stamp : no type generator for glong (time_t) for param stamp

// Sets @description as the description of the bookmark for @uri.
//
// If @uri is %NULL, the description of @bookmark is set.
//
// If a bookmark for @uri cannot be found then it is created.
/*

C function

g_bookmark_file_set_description
*/
func (recv *BookmarkFile) SetDescription(uri string, description string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_bookmark_file_set_description((*C.GBookmarkFile)(recv.native), c_uri, c_description)

	return
}

// Unsupported : g_bookmark_file_set_groups : unsupported parameter groups :

// Sets the icon for the bookmark for @uri. If @href is %NULL, unsets
// the currently set icon. @href can either be a full URL for the icon
// file or the icon name following the Icon Naming specification.
//
// If no bookmark for @uri is found one is created.
/*

C function

g_bookmark_file_set_icon
*/
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

// Sets the private flag of the bookmark for @uri.
//
// If a bookmark for @uri cannot be found then it is created.
/*

C function

g_bookmark_file_set_is_private
*/
func (recv *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_is_private :=
		boolToGboolean(isPrivate)

	C.g_bookmark_file_set_is_private((*C.GBookmarkFile)(recv.native), c_uri, c_is_private)

	return
}

// Sets @mime_type as the MIME type of the bookmark for @uri.
//
// If a bookmark for @uri cannot be found then it is created.
/*

C function

g_bookmark_file_set_mime_type
*/
func (recv *BookmarkFile) SetMimeType(uri string, mimeType string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.g_bookmark_file_set_mime_type((*C.GBookmarkFile)(recv.native), c_uri, c_mime_type)

	return
}

// Unsupported : g_bookmark_file_set_modified : unsupported parameter modified : no type generator for glong (time_t) for param modified

// Sets @title as the title of the bookmark for @uri inside the
// bookmark file @bookmark.
//
// If @uri is %NULL, the title of @bookmark is set.
//
// If a bookmark for @uri cannot be found then it is created.
/*

C function

g_bookmark_file_set_title
*/
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_bookmark_file_set_title((*C.GBookmarkFile)(recv.native), c_uri, c_title)

	return
}

// Unsupported : g_bookmark_file_set_visited : unsupported parameter visited : no type generator for glong (time_t) for param visited

// Unsupported : g_bookmark_file_to_data : no return type

// This function outputs @bookmark into a file.  The write process is
// guaranteed to be atomic by using g_file_set_contents() internally.
/*

C function

g_bookmark_file_to_file
*/
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

// Returns the value associated with @key under @group_name as a
// double. If @group_name is %NULL, the start_group is used.
//
// If @key cannot be found then 0.0 is returned and @error is set to
// #G_KEY_FILE_ERROR_KEY_NOT_FOUND. Likewise, if the value associated
// with @key cannot be interpreted as a double then 0.0 is returned
// and @error is set to #G_KEY_FILE_ERROR_INVALID_VALUE.
/*

C function

g_key_file_get_double
*/
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

// Associates a new double value with @key under @group_name.
// If @key cannot be found then it is created.
/*

C function

g_key_file_set_double
*/
func (recv *KeyFile) SetDouble(groupName string, key string, value float64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gdouble)(value)

	C.g_key_file_set_double((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Associates a list of double values with @key under
// @group_name.  If @key cannot be found then it is created.
/*

C function

g_key_file_set_double_list
*/
func (recv *KeyFile) SetDoubleList(groupName string, key string, list []float64) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list := &list[0]

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_double_list((*C.GKeyFile)(recv.native), c_group_name, c_key, (*C.gdouble)(unsafe.Pointer(c_list)), c_length)

	return
}

// Returns the description. See g_option_context_set_description().
/*

C function

g_option_context_get_description
*/
func (recv *OptionContext) GetDescription() string {
	retC := C.g_option_context_get_description((*C.GOptionContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns the summary. See g_option_context_set_summary().
/*

C function

g_option_context_get_summary
*/
func (recv *OptionContext) GetSummary() string {
	retC := C.g_option_context_get_summary((*C.GOptionContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Adds a string to be displayed in `--help` output after the list
// of options. This text often includes a bug reporting address.
//
// Note that the summary is translated (see
// g_option_context_set_translate_func()).
/*

C function

g_option_context_set_description
*/
func (recv *OptionContext) SetDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_option_context_set_description((*C.GOptionContext)(recv.native), c_description)

	return
}

// Adds a string to be displayed in `--help` output before the list
// of options. This is typically a summary of the program functionality.
//
// Note that the summary is translated (see
// g_option_context_set_translate_func() and
// g_option_context_set_translation_domain()).
/*

C function

g_option_context_set_summary
*/
func (recv *OptionContext) SetSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_option_context_set_summary((*C.GOptionContext)(recv.native), c_summary)

	return
}

// Unsupported : g_option_context_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// A convenience function to use gettext() for translating
// user-visible strings.
/*

C function

g_option_context_set_translation_domain
*/
func (recv *OptionContext) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_context_set_translation_domain((*C.GOptionContext)(recv.native), c_domain)

	return
}

// Returns whether @source has been destroyed.
//
// This is important when you operate upon your objects
// from within idle handlers, but may have freed the object
// before the dispatch of your idle handler.
//
// |[<!-- language="C" -->
// static gboolean
// idle_callback (gpointer data)
// {
// SomeWidget *self = data;
//
// GDK_THREADS_ENTER ();
// do stuff with self
// self->idle_id = 0;
// GDK_THREADS_LEAVE ();
//
// return G_SOURCE_REMOVE;
// }
//
// static void
// some_widget_do_stuff_later (SomeWidget *self)
// {
// self->idle_id = g_idle_add (idle_callback, self);
// }
//
// static void
// some_widget_finalize (GObject *object)
// {
// SomeWidget *self = SOME_WIDGET (object);
//
// if (self->idle_id)
// g_source_remove (self->idle_id);
//
// G_OBJECT_CLASS (parent_class)->finalize (object);
// }
// ]|
//
// This will fail in a multi-threaded application if the
// widget is destroyed before the idle handler fires due
// to the use after free in the callback. A solution, to
// this particular problem, is to check to if the source
// has already been destroy within the callback.
//
// |[<!-- language="C" -->
// static gboolean
// idle_callback (gpointer data)
// {
// SomeWidget *self = data;
//
// GDK_THREADS_ENTER ();
// if (!g_source_is_destroyed (g_main_current_source ()))
// {
// do stuff with self
// }
// GDK_THREADS_LEAVE ();
//
// return FALSE;
// }
// ]|
//
// Calls to this function from a thread other than the one acquired by the
// #GMainContext the #GSource is attached to are typically redundant, as the
// source could be destroyed immediately after this function returns. However,
// once a source is destroyed it cannot be un-destroyed, so this function can be
// used for opportunistic checks from any thread.
/*

C function

g_source_is_destroyed
*/
func (recv *Source) IsDestroyed() bool {
	retC := C.g_source_is_destroyed((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the source functions (can be used to override
// default implementations) of an unattached source.
/*

C function

g_source_set_funcs
*/
func (recv *Source) SetFuncs(funcs *SourceFuncs) {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	C.g_source_set_funcs((*C.GSource)(recv.native), c_funcs)

	return
}

// Converts @time_ into an RFC 3339 encoded string, relative to the
// Coordinated Universal Time (UTC). This is one of the many formats
// allowed by ISO 8601.
//
// ISO 8601 allows a large number of date/time formats, with or without
// punctuation and optional elements. The format returned by this function
// is a complete date and time, with optional punctuation included, the
// UTC time zone represented as "Z", and the @tv_usec part included if
// and only if it is nonzero, i.e. either
// "YYYY-MM-DDTHH:MM:SSZ" or "YYYY-MM-DDTHH:MM:SS.fffffZ".
//
// This corresponds to the Internet date/time format defined by
// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt),
// and to either of the two most-precise formats defined by
// the W3C Note
// [Date and Time Formats](http://www.w3.org/TR/NOTE-datetime-19980827).
// Both of these documents are profiles of ISO 8601.
//
// Use g_date_time_format() or g_strdup_printf() if a different
// variation of ISO 8601 format is required.
//
// If @time_ represents a date which is too large to fit into a `struct tm`,
// %NULL will be returned. This is platform dependent, but it is safe to assume
// years up to 3000 are supported. The return value of g_time_val_to_iso8601()
// has been nullable since GLib 2.54; before then, GLib would crash under the
// same conditions.
/*

C function

g_time_val_to_iso8601
*/
func (recv *TimeVal) ToIso8601() string {
	retC := C.g_time_val_to_iso8601((*C.GTimeVal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
