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

// ActionMapInterface is a wrapper around the C record GActionMapInterface.
type ActionMapInterface struct {
	native *C.GActionMapInterface
	// g_iface : record
	// no type for lookup_action
	// no type for add_action
	// no type for remove_action
}

func ActionMapInterfaceNewFromC(u unsafe.Pointer) *ActionMapInterface {
	c := (*C.GActionMapInterface)(u)
	if c == nil {
		return nil
	}

	g := &ActionMapInterface{native: c}

	return g
}

func (recv *ActionMapInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Prints what the matcher is matching against. The format will be
// equal to the format passed to g_file_attribute_matcher_new().
// The output however, might not be identical, as the matcher may
// decide to use a different order or omit needless parts.
/*

C function : g_file_attribute_matcher_to_string
*/
func (recv *FileAttributeMatcher) ToString() string {
	retC := C.g_file_attribute_matcher_to_string((*C.GFileAttributeMatcher)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// NetworkMonitorInterface is a wrapper around the C record GNetworkMonitorInterface.
type NetworkMonitorInterface struct {
	native *C.GNetworkMonitorInterface
	// g_iface : record
	// no type for network_changed
	// no type for can_reach
	// no type for can_reach_async
	// no type for can_reach_finish
}

func NetworkMonitorInterfaceNewFromC(u unsafe.Pointer) *NetworkMonitorInterface {
	c := (*C.GNetworkMonitorInterface)(u)
	if c == nil {
		return nil
	}

	g := &NetworkMonitorInterface{native: c}

	return g
}

func (recv *NetworkMonitorInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RemoteActionGroupInterface is a wrapper around the C record GRemoteActionGroupInterface.
type RemoteActionGroupInterface struct {
	native *C.GRemoteActionGroupInterface
	// g_iface : record
	// no type for activate_action_full
	// no type for change_action_state_full
}

func RemoteActionGroupInterfaceNewFromC(u unsafe.Pointer) *RemoteActionGroupInterface {
	c := (*C.GRemoteActionGroupInterface)(u)
	if c == nil {
		return nil
	}

	g := &RemoteActionGroupInterface{native: c}

	return g
}

func (recv *RemoteActionGroupInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Resource is a wrapper around the C record GResource.
type Resource struct {
	native *C.GResource
}

func ResourceNewFromC(u unsafe.Pointer) *Resource {
	c := (*C.GResource)(u)
	if c == nil {
		return nil
	}

	g := &Resource{native: c}

	return g
}

func (recv *Resource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a GResource from a reference to the binary resource bundle.
// This will keep a reference to @data while the resource lives, so
// the data should not be modified or freed.
//
// If you want to use this resource in the global resource namespace you need
// to register it with g_resources_register().
//
// Note: @data must be backed by memory that is at least pointer aligned.
// Otherwise this function will internally create a copy of the memory since
// GLib 2.56, or in older versions fail and exit the process.
/*

C function : g_resource_new_from_data
*/
func ResourceNewFromData(data *glib.Bytes) (*Resource, error) {
	c_data := (*C.GBytes)(C.NULL)
	if data != nil {
		c_data = (*C.GBytes)(data.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resource_new_from_data(c_data, &cThrowableError)
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

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
func (recv *Resource) Register() {
	C.g_resources_register((*C.GResource)(recv.native))

	return
}

// Unregisters the resource from the process-global set of resources.
/*

C function : g_resources_unregister
*/
func (recv *Resource) Unregister() {
	C.g_resources_unregister((*C.GResource)(recv.native))

	return
}

// Unsupported : g_resource_enumerate_children : no return type

// Looks for a file at the specified @path in the resource and
// if found returns information about it.
//
// @lookup_flags controls the behaviour of the lookup.
/*

C function : g_resource_get_info
*/
func (recv *Resource) GetInfo(path string, lookupFlags ResourceLookupFlags) (bool, uint64, uint32, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var c_size C.gsize

	var c_flags C.guint32

	var cThrowableError *C.GError

	retC := C.g_resource_get_info((*C.GResource)(recv.native), c_path, c_lookup_flags, &c_size, &c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	size := (uint64)(c_size)

	flags := (uint32)(c_flags)

	return retGo, size, flags, goThrowableError
}

// Looks for a file at the specified @path in the resource and
// returns a #GBytes that lets you directly access the data in
// memory.
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

C function : g_resource_lookup_data
*/
func (recv *Resource) LookupData(path string, lookupFlags ResourceLookupFlags) (*glib.Bytes, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resource_lookup_data((*C.GResource)(recv.native), c_path, c_lookup_flags, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Looks for a file at the specified @path in the resource and
// returns a #GInputStream that lets you read the data.
//
// @lookup_flags controls the behaviour of the lookup.
/*

C function : g_resource_open_stream
*/
func (recv *Resource) OpenStream(path string, lookupFlags ResourceLookupFlags) (*InputStream, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_lookup_flags := (C.GResourceLookupFlags)(lookupFlags)

	var cThrowableError *C.GError

	retC := C.g_resource_open_stream((*C.GResource)(recv.native), c_path, c_lookup_flags, &cThrowableError)
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Atomically increments the reference count of @resource by one. This
// function is MT-safe and may be called from any thread.
/*

C function : g_resource_ref
*/
func (recv *Resource) Ref() *Resource {
	retC := C.g_resource_ref((*C.GResource)(recv.native))
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Atomically decrements the reference count of @resource by one. If the
// reference count drops to 0, all memory allocated by the resource is
// released. This function is MT-safe and may be called from any
// thread.
/*

C function : g_resource_unref
*/
func (recv *Resource) Unref() {
	C.g_resource_unref((*C.GResource)(recv.native))

	return
}

// SettingsSchema is a wrapper around the C record GSettingsSchema.
type SettingsSchema struct {
	native *C.GSettingsSchema
}

func SettingsSchemaNewFromC(u unsafe.Pointer) *SettingsSchema {
	c := (*C.GSettingsSchema)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchema{native: c}

	return g
}

func (recv *SettingsSchema) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Get the ID of @schema.
/*

C function : g_settings_schema_get_id
*/
func (recv *SettingsSchema) GetId() string {
	retC := C.g_settings_schema_get_id((*C.GSettingsSchema)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the path associated with @schema, or %NULL.
//
// Schemas may be single-instance or relocatable.  Single-instance
// schemas correspond to exactly one set of keys in the backend
// database: those located at the path returned by this function.
//
// Relocatable schemas can be referenced by other schemas and can
// threfore describe multiple sets of keys at different locations.  For
// relocatable schemas, this function will return %NULL.
/*

C function : g_settings_schema_get_path
*/
func (recv *SettingsSchema) GetPath() string {
	retC := C.g_settings_schema_get_path((*C.GSettingsSchema)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Increase the reference count of @schema, returning a new reference.
/*

C function : g_settings_schema_ref
*/
func (recv *SettingsSchema) Ref() *SettingsSchema {
	retC := C.g_settings_schema_ref((*C.GSettingsSchema)(recv.native))
	retGo := SettingsSchemaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrease the reference count of @schema, possibly freeing it.
/*

C function : g_settings_schema_unref
*/
func (recv *SettingsSchema) Unref() {
	C.g_settings_schema_unref((*C.GSettingsSchema)(recv.native))

	return
}

// SettingsSchemaSource is a wrapper around the C record GSettingsSchemaSource.
type SettingsSchemaSource struct {
	native *C.GSettingsSchemaSource
}

func SettingsSchemaSourceNewFromC(u unsafe.Pointer) *SettingsSchemaSource {
	c := (*C.GSettingsSchemaSource)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchemaSource{native: c}

	return g
}

func (recv *SettingsSchemaSource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Attempts to create a new schema source corresponding to the contents
// of the given directory.
//
// This function is not required for normal uses of #GSettings but it
// may be useful to authors of plugin management systems.
//
// The directory should contain a file called `gschemas.compiled` as
// produced by the [glib-compile-schemas][glib-compile-schemas] tool.
//
// If @trusted is %TRUE then `gschemas.compiled` is trusted not to be
// corrupted. This assumption has a performance advantage, but can result
// in crashes or inconsistent behaviour in the case of a corrupted file.
// Generally, you should set @trusted to %TRUE for files installed by the
// system and to %FALSE for files in the home directory.
//
// If @parent is non-%NULL then there are two effects.
//
// First, if g_settings_schema_source_lookup() is called with the
// @recursive flag set to %TRUE and the schema can not be found in the
// source, the lookup will recurse to the parent.
//
// Second, any references to other schemas specified within this
// source (ie: `child` or `extends`) references may be resolved
// from the @parent.
//
// For this second reason, except in very unusual situations, the
// @parent should probably be given as the default schema source, as
// returned by g_settings_schema_source_get_default().
/*

C function : g_settings_schema_source_new_from_directory
*/
func SettingsSchemaSourceNewFromDirectory(directory string, parent *SettingsSchemaSource, trusted bool) (*SettingsSchemaSource, error) {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	c_parent := (*C.GSettingsSchemaSource)(C.NULL)
	if parent != nil {
		c_parent = (*C.GSettingsSchemaSource)(parent.ToC())
	}

	c_trusted :=
		boolToGboolean(trusted)

	var cThrowableError *C.GError

	retC := C.g_settings_schema_source_new_from_directory(c_directory, c_parent, c_trusted, &cThrowableError)
	retGo := SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Looks up a schema with the identifier @schema_id in @source.
//
// This function is not required for normal uses of #GSettings but it
// may be useful to authors of plugin management systems or to those who
// want to introspect the content of schemas.
//
// If the schema isn't found directly in @source and @recursive is %TRUE
// then the parent sources will also be checked.
//
// If the schema isn't found, %NULL is returned.
/*

C function : g_settings_schema_source_lookup
*/
func (recv *SettingsSchemaSource) Lookup(schemaId string, recursive bool) *SettingsSchema {
	c_schema_id := C.CString(schemaId)
	defer C.free(unsafe.Pointer(c_schema_id))

	c_recursive :=
		boolToGboolean(recursive)

	retC := C.g_settings_schema_source_lookup((*C.GSettingsSchemaSource)(recv.native), c_schema_id, c_recursive)
	var retGo (*SettingsSchema)
	if retC == nil {
		retGo = nil
	} else {
		retGo = SettingsSchemaNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Increase the reference count of @source, returning a new reference.
/*

C function : g_settings_schema_source_ref
*/
func (recv *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	retC := C.g_settings_schema_source_ref((*C.GSettingsSchemaSource)(recv.native))
	retGo := SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrease the reference count of @source, possibly freeing it.
/*

C function : g_settings_schema_source_unref
*/
func (recv *SettingsSchemaSource) Unref() {
	C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(recv.native))

	return
}

// Finalized a GResource initialized by g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources]
// and is not typically used by other code.
/*

C function : g_static_resource_fini
*/
func (recv *StaticResource) Fini() {
	C.g_static_resource_fini((*C.GStaticResource)(recv.native))

	return
}

// Gets the GResource that was registered by a call to g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources]
// and is not typically used by other code.
/*

C function : g_static_resource_get_resource
*/
func (recv *StaticResource) GetResource() *Resource {
	retC := C.g_static_resource_get_resource((*C.GStaticResource)(recv.native))
	retGo := ResourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Initializes a GResource from static data using a
// GStaticResource.
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources]
// and is not typically used by other code.
/*

C function : g_static_resource_init
*/
func (recv *StaticResource) Init() {
	C.g_static_resource_init((*C.GStaticResource)(recv.native))

	return
}

// Gets the options for the mount point.
/*

C function : g_unix_mount_point_get_options
*/
func (recv *UnixMountPoint) GetOptions() string {
	retC := C.g_unix_mount_point_get_options((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
