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

// The virtual function table for #GActionMap.
/*

C type

GActionMapInterface
*/
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

C function

g_file_attribute_matcher_to_string
*/
func (recv *FileAttributeMatcher) ToString() string {
	retC := C.g_file_attribute_matcher_to_string((*C.GFileAttributeMatcher)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// The virtual function table for #GNetworkMonitor.
/*

C type

GNetworkMonitorInterface
*/
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

// The virtual function table for #GRemoteActionGroup.
/*

C type

GRemoteActionGroupInterface
*/
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

// Applications and libraries often contain binary or textual data that is
// really part of the application, rather than user data. For instance
// #GtkBuilder .ui files, splashscreen images, GMenu markup XML, CSS files,
// icons, etc. These are often shipped as files in `$datadir/appname`, or
// manually included as literal strings in the code.
//
// The #GResource API and the [glib-compile-resources][glib-compile-resources] program
// provide a convenient and efficient alternative to this which has some nice properties. You
// maintain the files as normal files, so its easy to edit them, but during the build the files
// are combined into a binary bundle that is linked into the executable. This means that loading
// the resource files are efficient (as they are already in memory, shared with other instances) and
// simple (no need to check for things like I/O errors or locate the files in the filesystem). It
// also makes it easier to create relocatable applications.
//
// Resource files can also be marked as compressed. Such files will be included in the resource bundle
// in a compressed form, but will be automatically uncompressed when the resource is used. This
// is very useful e.g. for larger text files that are parsed once (or rarely) and then thrown away.
//
// Resource files can also be marked to be preprocessed, by setting the value of the
// `preprocess` attribute to a comma-separated list of preprocessing options.
// The only options currently supported are:
//
// `xml-stripblanks` which will use the xmllint command
// to strip ignorable whitespace from the XML file. For this to work,
// the `XMLLINT` environment variable must be set to the full path to
// the xmllint executable, or xmllint must be in the `PATH`; otherwise
// the preprocessing step is skipped.
//
// `to-pixdata` which will use the gdk-pixbuf-pixdata command to convert
// images to the GdkPixdata format, which allows you to create pixbufs directly using the data inside
// the resource file, rather than an (uncompressed) copy if it. For this, the gdk-pixbuf-pixdata
// program must be in the PATH, or the `GDK_PIXBUF_PIXDATA` environment variable must be
// set to the full path to the gdk-pixbuf-pixdata executable; otherwise the resource compiler will
// abort.
//
// Resource files will be exported in the GResource namespace using the
// combination of the given `prefix` and the filename from the `file` element.
// The `alias` attribute can be used to alter the filename to expose them at a
// different location in the resource namespace. Typically, this is used to
// include files from a different source directory without exposing the source
// directory in the resource namespace, as in the example below.
//
// Resource bundles are created by the [glib-compile-resources][glib-compile-resources] program
// which takes an XML file that describes the bundle, and a set of files that the XML references. These
// are combined into a binary resource bundle.
//
// An example resource description:
// |[
// <?xml version="1.0" encoding="UTF-8"?>
// <gresources>
// <gresource prefix="/org/gtk/Example">
// <file>data/splashscreen.png</file>
// <file compressed="true">dialog.ui</file>
// <file preprocess="xml-stripblanks">menumarkup.xml</file>
// <file alias="example.css">data/example.css</file>
// </gresource>
// </gresources>
// ]|
//
// This will create a resource bundle with the following files:
// |[
// /org/gtk/Example/data/splashscreen.png
// /org/gtk/Example/dialog.ui
// /org/gtk/Example/menumarkup.xml
// /org/gtk/Example/example.css
// ]|
//
// Note that all resources in the process share the same namespace, so use Java-style
// path prefixes (like in the above example) to avoid conflicts.
//
// You can then use [glib-compile-resources][glib-compile-resources] to compile the XML to a
// binary bundle that you can load with g_resource_load(). However, its more common to use the --generate-source and
// --generate-header arguments to create a source file and header to link directly into your application.
// This will generate `get_resource()`, `register_resource()` and
// `unregister_resource()` functions, prefixed by the `--c-name` argument passed
// to [glib-compile-resources][glib-compile-resources]. `get_resource()` returns
// the generated #GResource object. The register and unregister functions
// register the resource so its files can be accessed using
// g_resources_lookup_data().
//
// Once a #GResource has been created and registered all the data in it can be accessed globally in the process by
// using API calls like g_resources_open_stream() to stream the data or g_resources_lookup_data() to get a direct pointer
// to the data. You can also use URIs like "resource:///org/gtk/Example/data/splashscreen.png" with #GFile to access
// the resource data.
//
// Some higher-level APIs, such as #GtkApplication, will automatically load
// resources from certain well-known paths in the resource namespace as a
// convenience. See the documentation for those APIs for details.
//
// There are two forms of the generated source, the default version uses the compiler support for constructor
// and destructor functions (where available) to automatically create and register the #GResource on startup
// or library load time. If you pass `--manual-register`, two functions to register/unregister the resource are created
// instead. This requires an explicit initialization call in your application/library, but it works on all platforms,
// even on the minor ones where constructors are not supported. (Constructor support is available for at least Win32, Mac OS and Linux.)
//
// Note that resource data can point directly into the data segment of e.g. a library, so if you are unloading libraries
// during runtime you need to be very careful with keeping around pointers to data from a resource, as this goes away
// when the library is unloaded. However, in practice this is not generally a problem, since most resource accesses
// are for your own resources, and resource data is often used once, during parsing, and then released.
//
// When debugging a program or testing a change to an installed version, it is often useful to be able to
// replace resources in the program or library, without recompiling, for debugging or quick hacking and testing
// purposes. Since GLib 2.50, it is possible to use the `G_RESOURCE_OVERLAYS` environment variable to selectively overlay
// resources with replacements from the filesystem.  It is a colon-separated list of substitutions to perform
// during resource lookups.
//
// A substitution has the form
//
// |[
// /org/gtk/libgtk=/home/desrt/gtk-overlay
// ]|
//
// The part before the `=` is the resource subpath for which the overlay applies.  The part after is a
// filesystem path which contains files and subdirectories as you would like to be loaded as resources with the
// equivalent names.
//
// In the example above, if an application tried to load a resource with the resource path
// `/org/gtk/libgtk/ui/gtkdialog.ui` then GResource would check the filesystem path
// `/home/desrt/gtk-overlay/ui/gtkdialog.ui`.  If a file was found there, it would be used instead.  This is an
// overlay, not an outright replacement, which means that if a file is not found at that path, the built-in
// version will be used instead.  Whiteouts are not currently supported.
//
// Substitutions must start with a slash, and must not contain a trailing slash before the '='.  The path after
// the slash should ideally be absolute, but this is not strictly required.  It is possible to overlay the
// location of a single resource with an individual file.
/*

C type

GResource
*/
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

C function

g_resource_new_from_data
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

C function

g_resources_register
*/
func (recv *Resource) Register() {
	C.g_resources_register((*C.GResource)(recv.native))

	return
}

// Unregisters the resource from the process-global set of resources.
/*

C function

g_resources_unregister
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

C function

g_resource_get_info
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

C function

g_resource_lookup_data
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

C function

g_resource_open_stream
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

C function

g_resource_ref
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

C function

g_resource_unref
*/
func (recv *Resource) Unref() {
	C.g_resource_unref((*C.GResource)(recv.native))

	return
}

// The #GSettingsSchemaSource and #GSettingsSchema APIs provide a
// mechanism for advanced control over the loading of schemas and a
// mechanism for introspecting their content.
//
// Plugin loading systems that wish to provide plugins a way to access
// settings face the problem of how to make the schemas for these
// settings visible to GSettings.  Typically, a plugin will want to ship
// the schema along with itself and it won't be installed into the
// standard system directories for schemas.
//
// #GSettingsSchemaSource provides a mechanism for dealing with this by
// allowing the creation of a new 'schema source' from which schemas can
// be acquired.  This schema source can then become part of the metadata
// associated with the plugin and queried whenever the plugin requires
// access to some settings.
//
// Consider the following example:
//
// |[<!-- language="C" -->
// typedef struct
// {
// ...
// GSettingsSchemaSource *schema_source;
// ...
// } Plugin;
//
// Plugin *
// initialise_plugin (const gchar *dir)
// {
// Plugin *plugin;
//
// ...
//
// plugin->schema_source =
// g_settings_schema_source_new_from_directory (dir,
// g_settings_schema_source_get_default (), FALSE, NULL);
//
// ...
//
// return plugin;
// }
//
// ...
//
// GSettings *
// plugin_get_settings (Plugin      *plugin,
// const gchar *schema_id)
// {
// GSettingsSchema *schema;
//
// if (schema_id == NULL)
// schema_id = plugin->identifier;
//
// schema = g_settings_schema_source_lookup (plugin->schema_source,
// schema_id, FALSE);
//
// if (schema == NULL)
// {
// ... disable the plugin or abort, etc ...
// }
//
// return g_settings_new_full (schema, NULL, NULL);
// }
// ]|
//
// The code above shows how hooks should be added to the code that
// initialises (or enables) the plugin to create the schema source and
// how an API can be added to the plugin system to provide a convenient
// way for the plugin to access its settings, using the schemas that it
// ships.
//
// From the standpoint of the plugin, it would need to ensure that it
// ships a gschemas.compiled file as part of itself, and then simply do
// the following:
//
// |[<!-- language="C" -->
// {
// GSettings *settings;
// gint some_value;
//
// settings = plugin_get_settings (self, NULL);
// some_value = g_settings_get_int (settings, "some-value");
// ...
// }
// ]|
//
// It's also possible that the plugin system expects the schema source
// files (ie: .gschema.xml files) instead of a gschemas.compiled file.
// In that case, the plugin loading system must compile the schemas for
// itself before attempting to create the settings source.
/*

C type

GSettingsSchema
*/
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

C function

g_settings_schema_get_id
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

C function

g_settings_schema_get_path
*/
func (recv *SettingsSchema) GetPath() string {
	retC := C.g_settings_schema_get_path((*C.GSettingsSchema)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Increase the reference count of @schema, returning a new reference.
/*

C function

g_settings_schema_ref
*/
func (recv *SettingsSchema) Ref() *SettingsSchema {
	retC := C.g_settings_schema_ref((*C.GSettingsSchema)(recv.native))
	retGo := SettingsSchemaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrease the reference count of @schema, possibly freeing it.
/*

C function

g_settings_schema_unref
*/
func (recv *SettingsSchema) Unref() {
	C.g_settings_schema_unref((*C.GSettingsSchema)(recv.native))

	return
}

// This is an opaque structure type.  You may not access it directly.
/*

C type

GSettingsSchemaSource
*/
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

C function

g_settings_schema_source_new_from_directory
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

C function

g_settings_schema_source_lookup
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

C function

g_settings_schema_source_ref
*/
func (recv *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	retC := C.g_settings_schema_source_ref((*C.GSettingsSchemaSource)(recv.native))
	retGo := SettingsSchemaSourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrease the reference count of @source, possibly freeing it.
/*

C function

g_settings_schema_source_unref
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

C function

g_static_resource_fini
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

C function

g_static_resource_get_resource
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

C function

g_static_resource_init
*/
func (recv *StaticResource) Init() {
	C.g_static_resource_init((*C.GStaticResource)(recv.native))

	return
}

// Gets the options for the mount point.
/*

C function

g_unix_mount_point_get_options
*/
func (recv *UnixMountPoint) GetOptions() string {
	retC := C.g_unix_mount_point_get_options((*C.GUnixMountPoint)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}
