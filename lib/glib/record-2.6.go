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

// Returns the week of the year, where weeks are interpreted according
// to ISO 8601.
/*

C function

g_date_get_iso8601_week_of_year
*/
func (recv *Date) GetIso8601WeekOfYear() uint32 {
	retC := C.g_date_get_iso8601_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Creates a new empty #GKeyFile object. Use
// g_key_file_load_from_file(), g_key_file_load_from_data(),
// g_key_file_load_from_dirs() or g_key_file_load_from_data_dirs() to
// read an existing key file.
/*

C function

g_key_file_new
*/
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Clears all keys and groups from @key_file, and decreases the
// reference count by 1. If the reference count reaches zero,
// frees the key file and all its allocated memory.
/*

C function

g_key_file_free
*/
func (recv *KeyFile) Free() {
	C.g_key_file_free((*C.GKeyFile)(recv.native))

	return
}

// Returns the value associated with @key under @group_name as a
// boolean.
//
// If @key cannot be found then %FALSE is returned and @error is set
// to #G_KEY_FILE_ERROR_KEY_NOT_FOUND. Likewise, if the value
// associated with @key cannot be interpreted as a boolean then %FALSE
// is returned and @error is set to #G_KEY_FILE_ERROR_INVALID_VALUE.
/*

C function

g_key_file_get_boolean
*/
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

// Unsupported : g_key_file_get_boolean_list : no return type

// Retrieves a comment above @key from @group_name.
// If @key is %NULL then @comment will be read from above
// @group_name. If both @key and @group_name are %NULL, then
// @comment will be read from above the first group in the file.
//
// Note that the returned string includes the '#' comment markers.
/*

C function

g_key_file_get_comment
*/
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

// Unsupported : g_key_file_get_groups : no return type

// Returns the value associated with @key under @group_name as an
// integer.
//
// If @key cannot be found then 0 is returned and @error is set to
// #G_KEY_FILE_ERROR_KEY_NOT_FOUND. Likewise, if the value associated
// with @key cannot be interpreted as an integer, or is out of range
// for a #gint, then 0 is returned
// and @error is set to #G_KEY_FILE_ERROR_INVALID_VALUE.
/*

C function

g_key_file_get_integer
*/
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

// Unsupported : g_key_file_get_integer_list : no return type

// Unsupported : g_key_file_get_keys : no return type

// Returns the value associated with @key under @group_name
// translated in the given @locale if available.  If @locale is
// %NULL then the current locale is assumed.
//
// If @locale is to be non-%NULL, or if the current locale will change over
// the lifetime of the #GKeyFile, it must be loaded with
// %G_KEY_FILE_KEEP_TRANSLATIONS in order to load strings for all locales.
//
// If @key cannot be found then %NULL is returned and @error is set
// to #G_KEY_FILE_ERROR_KEY_NOT_FOUND. If the value associated
// with @key cannot be interpreted or no suitable translation can
// be found then the untranslated value is returned.
/*

C function

g_key_file_get_locale_string
*/
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

// Unsupported : g_key_file_get_locale_string_list : no return type

// Returns the name of the start group of the file.
/*

C function

g_key_file_get_start_group
*/
func (recv *KeyFile) GetStartGroup() string {
	retC := C.g_key_file_get_start_group((*C.GKeyFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns the string value associated with @key under @group_name.
// Unlike g_key_file_get_value(), this function handles escape sequences
// like \s.
//
// In the event the key cannot be found, %NULL is returned and
// @error is set to #G_KEY_FILE_ERROR_KEY_NOT_FOUND.  In the
// event that the @group_name cannot be found, %NULL is returned
// and @error is set to #G_KEY_FILE_ERROR_GROUP_NOT_FOUND.
/*

C function

g_key_file_get_string
*/
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

// Unsupported : g_key_file_get_string_list : no return type

// Returns the raw value associated with @key under @group_name.
// Use g_key_file_get_string() to retrieve an unescaped UTF-8 string.
//
// In the event the key cannot be found, %NULL is returned and
// @error is set to #G_KEY_FILE_ERROR_KEY_NOT_FOUND.  In the
// event that the @group_name cannot be found, %NULL is returned
// and @error is set to #G_KEY_FILE_ERROR_GROUP_NOT_FOUND.
/*

C function

g_key_file_get_value
*/
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

// Looks whether the key file has the group @group_name.
/*

C function

g_key_file_has_group
*/
func (recv *KeyFile) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.g_key_file_has_group((*C.GKeyFile)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// Looks whether the key file has the key @key in the group
// @group_name.
//
// Note that this function does not follow the rules for #GError strictly;
// the return value both carries meaning and signals an error.  To use
// this function, you must pass a #GError pointer in @error, and check
// whether it is not %NULL to see if an error occurred.
//
// Language bindings should use g_key_file_get_value() to test whether
// or not a key exists.
/*

C function

g_key_file_has_key
*/
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

// Loads a key file from memory into an empty #GKeyFile structure.
// If the object cannot be created then %error is set to a #GKeyFileError.
/*

C function

g_key_file_load_from_data
*/
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

// This function looks for a key file named @file in the paths
// returned from g_get_user_data_dir() and g_get_system_data_dirs(),
// loads the file into @key_file and returns the file's full path in
// @full_path.  If the file could not be loaded then an %error is
// set to either a #GFileError or #GKeyFileError.
/*

C function

g_key_file_load_from_data_dirs
*/
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

// Loads a key file into an empty #GKeyFile structure.
//
// If the OS returns an error when opening or reading the file, a
// %G_FILE_ERROR is returned. If there is a problem parsing the file, a
// %G_KEY_FILE_ERROR is returned.
//
// This function will never return a %G_KEY_FILE_ERROR_NOT_FOUND error. If the
// @file is not found, %G_FILE_ERROR_NOENT is returned.
/*

C function

g_key_file_load_from_file
*/
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

// Removes a comment above @key from @group_name.
// If @key is %NULL then @comment will be removed above @group_name.
// If both @key and @group_name are %NULL, then @comment will
// be removed above the first group in the file.
/*

C function

g_key_file_remove_comment
*/
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

// Removes the specified group, @group_name,
// from the key file.
/*

C function

g_key_file_remove_group
*/
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

// Removes @key in @group_name from the key file.
/*

C function

g_key_file_remove_key
*/
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

// Associates a new boolean value with @key under @group_name.
// If @key cannot be found then it is created.
/*

C function

g_key_file_set_boolean
*/
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

// Associates a list of boolean values with @key under @group_name.
// If @key cannot be found then it is created.
// If @group_name is %NULL, the start_group is used.
/*

C function

g_key_file_set_boolean_list
*/
func (recv *KeyFile) SetBooleanList(groupName string, key string, list []bool) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list := &list[0]

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_boolean_list((*C.GKeyFile)(recv.native), c_group_name, c_key, (*C.gboolean)(unsafe.Pointer(c_list)), c_length)

	return
}

// Places a comment above @key from @group_name.
//
// If @key is %NULL then @comment will be written above @group_name.
// If both @key and @group_name  are %NULL, then @comment will be
// written above the first group in the file.
//
// Note that this function prepends a '#' comment marker to
// each line of @comment.
/*

C function

g_key_file_set_comment
*/
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

// Associates a new integer value with @key under @group_name.
// If @key cannot be found then it is created.
/*

C function

g_key_file_set_integer
*/
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	C.g_key_file_set_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// Associates a list of integer values with @key under @group_name.
// If @key cannot be found then it is created.
/*

C function

g_key_file_set_integer_list
*/
func (recv *KeyFile) SetIntegerList(groupName string, key string, list []int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list := &list[0]

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_integer_list((*C.GKeyFile)(recv.native), c_group_name, c_key, (*C.gint)(unsafe.Pointer(c_list)), c_length)

	return
}

// Sets the character which is used to separate
// values in lists. Typically ';' or ',' are used
// as separators. The default list separator is ';'.
/*

C function

g_key_file_set_list_separator
*/
func (recv *KeyFile) SetListSeparator(separator rune) {
	c_separator := (C.gchar)(separator)

	C.g_key_file_set_list_separator((*C.GKeyFile)(recv.native), c_separator)

	return
}

// Associates a string value for @key and @locale under @group_name.
// If the translation for @key cannot be found then it is created.
/*

C function

g_key_file_set_locale_string
*/
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

// Unsupported : g_key_file_set_locale_string_list : unsupported parameter list :

// Associates a new string value with @key under @group_name.
// If @key cannot be found then it is created.
// If @group_name cannot be found then it is created.
// Unlike g_key_file_set_value(), this function handles characters
// that need escaping, such as newlines.
/*

C function

g_key_file_set_string
*/
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

// Unsupported : g_key_file_set_string_list : unsupported parameter list :

// Associates a new value with @key under @group_name.
//
// If @key cannot be found then it is created. If @group_name cannot
// be found then it is created. To set an UTF-8 string which may contain
// characters that need escaping (such as newlines or spaces), use
// g_key_file_set_string().
/*

C function

g_key_file_set_value
*/
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

// This function outputs @key_file as a string.
//
// Note that this function never reports an error,
// so it is safe to pass %NULL as @error.
/*

C function

g_key_file_to_data
*/
func (recv *KeyFile) ToData() (string, uint64, error) {
	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_to_data((*C.GKeyFile)(recv.native), &c_length, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goThrowableError
}

// Adds a #GOptionGroup to the @context, so that parsing with @context
// will recognize the options in the group. Note that this will take
// ownership of the @group and thus the @group should not be freed.
/*

C function

g_option_context_add_group
*/
func (recv *OptionContext) AddGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_option_context_add_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// A convenience function which creates a main group if it doesn't
// exist, adds the @entries to it and sets the translation domain.
/*

C function

g_option_context_add_main_entries
*/
func (recv *OptionContext) AddMainEntries(entries *OptionEntry, translationDomain string) {
	c_entries := (*C.GOptionEntry)(C.NULL)
	if entries != nil {
		c_entries = (*C.GOptionEntry)(entries.ToC())
	}

	c_translation_domain := C.CString(translationDomain)
	defer C.free(unsafe.Pointer(c_translation_domain))

	C.g_option_context_add_main_entries((*C.GOptionContext)(recv.native), c_entries, c_translation_domain)

	return
}

// Frees context and all the groups which have been
// added to it.
//
// Please note that parsed arguments need to be freed separately (see
// #GOptionEntry).
/*

C function

g_option_context_free
*/
func (recv *OptionContext) Free() {
	C.g_option_context_free((*C.GOptionContext)(recv.native))

	return
}

// Returns whether automatic `--help` generation
// is turned on for @context. See g_option_context_set_help_enabled().
/*

C function

g_option_context_get_help_enabled
*/
func (recv *OptionContext) GetHelpEnabled() bool {
	retC := C.g_option_context_get_help_enabled((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns whether unknown options are ignored or not. See
// g_option_context_set_ignore_unknown_options().
/*

C function

g_option_context_get_ignore_unknown_options
*/
func (recv *OptionContext) GetIgnoreUnknownOptions() bool {
	retC := C.g_option_context_get_ignore_unknown_options((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns a pointer to the main group of @context.
/*

C function

g_option_context_get_main_group
*/
func (recv *OptionContext) GetMainGroup() *OptionGroup {
	retC := C.g_option_context_get_main_group((*C.GOptionContext)(recv.native))
	retGo := OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parses the command line arguments, recognizing options
// which have been added to @context. A side-effect of
// calling this function is that g_set_prgname() will be
// called.
//
// If the parsing is successful, any parsed arguments are
// removed from the array and @argc and @argv are updated
// accordingly. A '--' option is stripped from @argv
// unless there are unparsed options before and after it,
// or some of the options after it start with '-'. In case
// of an error, @argc and @argv are left unmodified.
//
// If automatic `--help` support is enabled
// (see g_option_context_set_help_enabled()), and the
// @argv array contains one of the recognized help options,
// this function will produce help output to stdout and
// call `exit (0)`.
//
// Note that function depends on the [current locale][setlocale] for
// automatic character set conversion of string and filename
// arguments.
/*

C function

g_option_context_parse
*/
func (recv *OptionContext) Parse(args []string) (bool, []string, error) {
	cArgc, cArgv := argsIn(args)

	var cThrowableError *C.GError

	retC := C.g_option_context_parse((*C.GOptionContext)(recv.native), &cArgc, &cArgv, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	args = argsOut(cArgc, cArgv)

	return retGo, args, goThrowableError
}

// Enables or disables automatic generation of `--help` output.
// By default, g_option_context_parse() recognizes `--help`, `-h`,
// `-?`, `--help-all` and `--help-groupname` and creates suitable
// output to stdout.
/*

C function

g_option_context_set_help_enabled
*/
func (recv *OptionContext) SetHelpEnabled(helpEnabled bool) {
	c_help_enabled :=
		boolToGboolean(helpEnabled)

	C.g_option_context_set_help_enabled((*C.GOptionContext)(recv.native), c_help_enabled)

	return
}

// Sets whether to ignore unknown options or not. If an argument is
// ignored, it is left in the @argv array after parsing. By default,
// g_option_context_parse() treats unknown options as error.
//
// This setting does not affect non-option arguments (i.e. arguments
// which don't start with a dash). But note that GOption cannot reliably
// determine whether a non-option belongs to a preceding unknown option.
/*

C function

g_option_context_set_ignore_unknown_options
*/
func (recv *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	c_ignore_unknown :=
		boolToGboolean(ignoreUnknown)

	C.g_option_context_set_ignore_unknown_options((*C.GOptionContext)(recv.native), c_ignore_unknown)

	return
}

// Sets a #GOptionGroup as main group of the @context.
// This has the same effect as calling g_option_context_add_group(),
// the only difference is that the options in the main group are
// treated differently when generating `--help` output.
/*

C function

g_option_context_set_main_group
*/
func (recv *OptionContext) SetMainGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_option_context_set_main_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Adds the options specified in @entries to @group.
/*

C function

g_option_group_add_entries
*/
func (recv *OptionGroup) AddEntries(entries *OptionEntry) {
	c_entries := (*C.GOptionEntry)(C.NULL)
	if entries != nil {
		c_entries = (*C.GOptionEntry)(entries.ToC())
	}

	C.g_option_group_add_entries((*C.GOptionGroup)(recv.native), c_entries)

	return
}

// Frees a #GOptionGroup. Note that you must not free groups
// which have been added to a #GOptionContext.
/*

C function

g_option_group_free
*/
func (recv *OptionGroup) Free() {
	C.g_option_group_free((*C.GOptionGroup)(recv.native))

	return
}

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc (GOptionErrorFunc) for param error_func

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc (GOptionParseFunc) for param pre_parse_func

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// A convenience function to use gettext() for translating
// user-visible strings.
/*

C function

g_option_group_set_translation_domain
*/
func (recv *OptionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_group_set_translation_domain((*C.GOptionGroup)(recv.native), c_domain)

	return
}
