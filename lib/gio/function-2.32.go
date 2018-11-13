// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Opens a file in the preferred directory for temporary files (as
// returned by g_get_tmp_dir()) and returns a #GFile and
// #GFileIOStream pointing to it.
//
// @tmpl should be a string in the GLib file name encoding
// containing a sequence of six 'X' characters, and containing no
// directory components. If it is %NULL, a default template is used.
//
// Unlike the other #GFile constructors, this will return %NULL if
// a temporary file could not be created.
/*

C function : g_file_new_tmp
*/
func FileNewTmp(tmpl string) (*File, *FileIOStream, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var c_iostream *C.GFileIOStream

	var cThrowableError *C.GError

	retC := C.g_file_new_tmp(c_tmpl, &c_iostream, &cThrowableError)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	iostream := FileIOStreamNewFromC(unsafe.Pointer(c_iostream))

	return retGo, iostream, goThrowableError
}

// Gets the default #GNetworkMonitor for the system.
/*

C function : g_network_monitor_get_default
*/
func NetworkMonitorGetDefault() *NetworkMonitor {
	retC := C.g_network_monitor_get_default()
	retGo := NetworkMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GResource Error Quark.
/*

C function : g_resource_error_quark
*/
func ResourceErrorQuark() glib.Quark {
	retC := C.g_resource_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Loads a binary resource bundle and creates a #GResource representation of it, allowing
// you to query it for data.
//
// If you want to use this resource in the global resource namespace you need
// to register it with g_resources_register().
/*

C function : g_resource_load
*/
func ResourceLoad(filename string) (*Resource, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_resource_load(c_filename, &cThrowableError)
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resources_enumerate_children : no return type

// Looks for a file at the specified @path in the set of
// globally registered resources and if found returns information about it.
//
// @lookup_flags controls the behaviour of the lookup.
/*

C function : g_resources_get_info
*/
func ResourcesGetInfo(path string, lookupFlags ResourceLookupFlags) (bool, uint64, uint32, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var c_size C.gsize

	var c_flags C.guint32

	var cThrowableError *C.GError

	retC := C.g_resources_get_info(c_path, c_lookup_flags, &c_size, &c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	size := (uint64)(c_size)

	flags := (uint32)(c_flags)

	return retGo, size, flags, goThrowableError
}

// Looks for a file at the specified @path in the set of
// globally registered resources and returns a #GBytes that
// lets you directly access the data in memory.
//
// The data is always followed by a zero byte, so you
// can safely use the data as a C string. However, that byte
// is not included in the size of the GBytes.
//
// For uncompressed resource files this is a pointer directly into
// the resource bundle, which is typically in some readonly data section
// in the program binary. For compressed files we allocate memory on
// the heap and automatically uncompress the data.
//
// @lookup_flags controls the behaviour of the lookup.
/*

C function : g_resources_lookup_data
*/
func ResourcesLookupData(path string, lookupFlags ResourceLookupFlags) (*glib.Bytes, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resources_lookup_data(c_path, c_lookup_flags, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Looks for a file at the specified @path in the set of
// globally registered resources and returns a #GInputStream
// that lets you read the data.
//
// @lookup_flags controls the behaviour of the lookup.
/*

C function : g_resources_open_stream
*/
func ResourcesOpenStream(path string, lookupFlags ResourceLookupFlags) (*InputStream, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resources_open_stream(c_path, c_lookup_flags, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Registers the resource with the process-global set of resources.
// Once a resource is registered the files in it can be accessed
// with the global resource lookup functions like g_resources_lookup_data().
/*

C function : g_resources_register
*/
func ResourcesRegister(resource *Resource) {
	c_resource := (*C.GResource)(C.NULL)
	if resource != nil {
		c_resource = (*C.GResource)(resource.ToC())
	}

	C.g_resources_register(c_resource)

	return
}

// Unregisters the resource from the process-global set of resources.
/*

C function : g_resources_unregister
*/
func ResourcesUnregister(resource *Resource) {
	c_resource := (*C.GResource)(C.NULL)
	if resource != nil {
		c_resource = (*C.GResource)(resource.ToC())
	}

	C.g_resources_unregister(c_resource)

	return
}

// Gets the default system schema source.
//
// This function is not required for normal uses of #GSettings but it
// may be useful to authors of plugin management systems or to those who
// want to introspect the content of schemas.
//
// If no schemas are installed, %NULL will be returned.
//
// The returned source may actually consist of multiple schema sources
// from different directories, depending on which directories were given
// in `XDG_DATA_DIRS` and `GSETTINGS_SCHEMA_DIR`. For this reason, all
// lookups performed against the default source should probably be done
// recursively.
/*

C function : g_settings_schema_source_get_default
*/
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	retC := C.g_settings_schema_source_get_default()
	var retGo (*SettingsSchemaSource)
	if retC == nil {
		retGo = nil
	} else {
		retGo = SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
