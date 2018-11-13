// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Creates a new #GAppInfo from the given information.
//
// Note that for @commandline, the quoting rules of the Exec key of the
// [freedesktop.org Desktop Entry Specification](http://freedesktop.org/Standards/desktop-entry-spec)
// are applied. For example, if the @commandline contains
// percent-encoded URIs, the percent-character must be doubled in order to prevent it from
// being swallowed by Exec key unquoting. See the specification for exact quoting rules.
/*

C function

g_app_info_create_from_commandline
*/
func AppInfoCreateFromCommandline(commandline string, applicationName string, flags AppInfoCreateFlags) (*AppInfo, error) {
	c_commandline := C.CString(commandline)
	defer C.free(unsafe.Pointer(c_commandline))

	c_application_name := C.CString(applicationName)
	defer C.free(unsafe.Pointer(c_application_name))

	c_flags := (C.GAppInfoCreateFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_app_info_create_from_commandline(c_commandline, c_application_name, c_flags, &cThrowableError)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets a list of all of the applications currently registered
// on this system.
//
// For desktop files, this includes applications that have
// `NoDisplay=true` set or are excluded from display by means
// of `OnlyShowIn` or `NotShowIn`. See g_app_info_should_show().
// The returned list does not include applications which have
// the `Hidden` key set.
/*

C function

g_app_info_get_all
*/
func AppInfoGetAll() *glib.List {
	retC := C.g_app_info_get_all()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a list of all #GAppInfos for a given content type,
// including the recommended and fallback #GAppInfos. See
// g_app_info_get_recommended_for_type() and
// g_app_info_get_fallback_for_type().
/*

C function

g_app_info_get_all_for_type
*/
func AppInfoGetAllForType(contentType string) *glib.List {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.g_app_info_get_all_for_type(c_content_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the default #GAppInfo for a given content type.
/*

C function

g_app_info_get_default_for_type
*/
func AppInfoGetDefaultForType(contentType string, mustSupportUris bool) *AppInfo {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	c_must_support_uris :=
		boolToGboolean(mustSupportUris)

	retC := C.g_app_info_get_default_for_type(c_content_type, c_must_support_uris)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the default application for handling URIs with
// the given URI scheme. A URI scheme is the initial part
// of the URI, up to but not including the ':', e.g. "http",
// "ftp" or "sip".
/*

C function

g_app_info_get_default_for_uri_scheme
*/
func AppInfoGetDefaultForUriScheme(uriScheme string) *AppInfo {
	c_uri_scheme := C.CString(uriScheme)
	defer C.free(unsafe.Pointer(c_uri_scheme))

	retC := C.g_app_info_get_default_for_uri_scheme(c_uri_scheme)
	retGo := AppInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Utility function that launches the default application
// registered to handle the specified uri. Synchronous I/O
// is done on the uri to detect the type of the file if
// required.
/*

C function

g_app_info_launch_default_for_uri
*/
func AppInfoLaunchDefaultForUri(uri string, context *AppLaunchContext) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_context := (*C.GAppLaunchContext)(C.NULL)
	if context != nil {
		c_context = (*C.GAppLaunchContext)(context.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_app_info_launch_default_for_uri(c_uri, c_context, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if a content type can be executable. Note that for instance
// things like text files can be executables (i.e. scripts and batch files).
/*

C function

g_content_type_can_be_executable
*/
func ContentTypeCanBeExecutable(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_can_be_executable(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Compares two content types for equality.
/*

C function

g_content_type_equals
*/
func ContentTypeEquals(type1 string, type2 string) bool {
	c_type1 := C.CString(type1)
	defer C.free(unsafe.Pointer(c_type1))

	c_type2 := C.CString(type2)
	defer C.free(unsafe.Pointer(c_type2))

	retC := C.g_content_type_equals(c_type1, c_type2)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the human readable description of the content type.
/*

C function

g_content_type_get_description
*/
func ContentTypeGetDescription(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_description(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the icon for a content type.
/*

C function

g_content_type_get_icon
*/
func ContentTypeGetIcon(type_ string) *Icon {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_icon(c_type)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the mime type for the content type, if one is registered.
/*

C function

g_content_type_get_mime_type
*/
func ContentTypeGetMimeType(type_ string) string {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_get_mime_type(c_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Guesses the content type based on example data. If the function is
// uncertain, @result_uncertain will be set to %TRUE. Either @filename
// or @data may be %NULL, in which case the guess will be based solely
// on the other argument.
/*

C function

g_content_type_guess
*/
func ContentTypeGuess(filename string, data []uint8) (string, bool) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_data := &data[0]

	c_data_size := (C.gsize)(len(data))

	var c_result_uncertain C.gboolean

	retC := C.g_content_type_guess(c_filename, (*C.guchar)(unsafe.Pointer(c_data)), c_data_size, &c_result_uncertain)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	resultUncertain := c_result_uncertain == C.TRUE

	return retGo, resultUncertain
}

// Determines if @type is a subset of @supertype.
/*

C function

g_content_type_is_a
*/
func ContentTypeIsA(type_ string, supertype string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_supertype := C.CString(supertype)
	defer C.free(unsafe.Pointer(c_supertype))

	retC := C.g_content_type_is_a(c_type, c_supertype)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the content type is the generic "unknown" type.
// On UNIX this is the "application/octet-stream" mimetype,
// while on win32 it is "*" and on OSX it is a dynamic type
// or octet-stream.
/*

C function

g_content_type_is_unknown
*/
func ContentTypeIsUnknown(type_ string) bool {
	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	retC := C.g_content_type_is_unknown(c_type)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a list of strings containing all the registered content types
// known to the system. The list and its data should be freed using
// g_list_free_full (list, g_free).
/*

C function

g_content_types_get_registered
*/
func ContentTypesGetRegistered() *glib.List {
	retC := C.g_content_types_get_registered()
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

/*

C function

g_dbus_error_quark
*/
func DbusErrorQuark() glib.Quark {
	retC := C.g_dbus_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Creates a #GFile with the given argument from the command line.
// The value of @arg can be either a URI, an absolute path or a
// relative path resolved relative to the current working directory.
// This operation never fails, but the returned object might not
// support any I/O operation if @arg points to a malformed path.
//
// Note that on Windows, this function expects its argument to be in
// UTF-8 -- not the system code page.  This means that you
// should not use this function with string from argv as it is passed
// to main().  g_win32_get_command_line() will return a UTF-8 version of
// the commandline.  #GApplication also uses UTF-8 but
// g_application_command_line_create_file_for_arg() may be more useful
// for you there.  It is also always possible to use this function with
// #GOptionContext arguments of type %G_OPTION_ARG_FILENAME.
/*

C function

g_file_new_for_commandline_arg
*/
func FileNewForCommandlineArg(arg string) *File {
	c_arg := C.CString(arg)
	defer C.free(unsafe.Pointer(c_arg))

	retC := C.g_file_new_for_commandline_arg(c_arg)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Constructs a #GFile for a given path. This operation never
// fails, but the returned object might not support any I/O
// operation if @path is malformed.
/*

C function

g_file_new_for_path
*/
func FileNewForPath(path string) *File {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_file_new_for_path(c_path)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Constructs a #GFile for a given URI. This operation never
// fails, but the returned object might not support any I/O
// operation if @uri is malformed or if the uri type is
// not supported.
/*

C function

g_file_new_for_uri
*/
func FileNewForUri(uri string) *File {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_file_new_for_uri(c_uri)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Constructs a #GFile with the given @parse_name (i.e. something
// given by g_file_get_parse_name()). This operation never fails,
// but the returned object might not support any I/O operation if
// the @parse_name cannot be parsed.
/*

C function

g_file_parse_name
*/
func FileParseName(parseName string) *File {
	c_parse_name := C.CString(parseName)
	defer C.free(unsafe.Pointer(c_parse_name))

	retC := C.g_file_parse_name(c_parse_name)
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets a hash for an icon.
/*

C function

g_icon_hash
*/
func IconHash(icon uintptr) uint32 {
	c_icon := (C.gconstpointer)(icon)

	retC := C.g_icon_hash(c_icon)
	retGo := (uint32)(retC)

	return retGo
}

// Converts errno.h error codes into GIO error codes. The fallback
// value %G_IO_ERROR_FAILED is returned for error codes not currently
// handled (but note that future GLib releases may return a more
// specific value instead).
//
// As %errno is global and may be modified by intermediate function
// calls, you should save its value as soon as the call which sets it
/*

C function

g_io_error_from_errno
*/
func IoErrorFromErrno(errNo int32) IOErrorEnum {
	c_err_no := (C.gint)(errNo)

	retC := C.g_io_error_from_errno(c_err_no)
	retGo := (IOErrorEnum)(retC)

	return retGo
}

// Gets the GIO Error Quark.
/*

C function

g_io_error_quark
*/
func IoErrorQuark() glib.Quark {
	retC := C.g_io_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Registers @type as extension for the extension point with name
// @extension_point_name.
//
// If @type has already been registered as an extension for this
// extension point, the existing #GIOExtension object is returned.
/*

C function

g_io_extension_point_implement
*/
func IoExtensionPointImplement(extensionPointName string, type_ gobject.Type, extensionName string, priority int32) *IOExtension {
	c_extension_point_name := C.CString(extensionPointName)
	defer C.free(unsafe.Pointer(c_extension_point_name))

	c_type := (C.GType)(type_)

	c_extension_name := C.CString(extensionName)
	defer C.free(unsafe.Pointer(c_extension_name))

	c_priority := (C.gint)(priority)

	retC := C.g_io_extension_point_implement(c_extension_point_name, c_type, c_extension_name, c_priority)
	retGo := IOExtensionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Looks up an existing extension point.
/*

C function

g_io_extension_point_lookup
*/
func IoExtensionPointLookup(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_lookup(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Registers an extension point.
/*

C function

g_io_extension_point_register
*/
func IoExtensionPointRegister(name string) *IOExtensionPoint {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_io_extension_point_register(c_name)
	retGo := IOExtensionPointNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Loads all the modules in the specified directory.
//
// If don't require all modules to be initialized (and thus registering
// all gtypes) then you can use g_io_modules_scan_all_in_directory()
// which allows delayed/lazy loading of modules.
/*

C function

g_io_modules_load_all_in_directory
*/
func IoModulesLoadAllInDirectory(dirname string) *glib.List {
	c_dirname := C.CString(dirname)
	defer C.free(unsafe.Pointer(c_dirname))

	retC := C.g_io_modules_load_all_in_directory(c_dirname)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Cancels all cancellable I/O jobs.
//
// A job is cancellable if a #GCancellable was passed into
// g_io_scheduler_push_job().
/*

C function

g_io_scheduler_cancel_all_jobs
*/
func IoSchedulerCancelAllJobs() {
	C.g_io_scheduler_cancel_all_jobs()

	return
}

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc (GIOSchedulerJobFunc) for param job_func

// Blacklisted : g_keyfile_settings_backend_new

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Determines if @mount_path is considered an implementation of the
// OS. This is primarily used for hiding mountable and mounted volumes
// that only are used in the OS and has little to no relevance to the
// casual user.
/*

C function

g_unix_is_mount_path_system_internal
*/
func UnixIsMountPathSystemInternal(mountPath string) bool {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	retC := C.g_unix_is_mount_path_system_internal(c_mount_path)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a #GUnixMountEntry for a given mount path. If @time_read
// is set, it will be filled with a unix timestamp for checking
// if the mounts have changed since with g_unix_mounts_changed_since().
/*

C function

g_unix_mount_at
*/
func UnixMountAt(mountPath string) (*UnixMountEntry, uint64) {
	c_mount_path := C.CString(mountPath)
	defer C.free(unsafe.Pointer(c_mount_path))

	var c_time_read C.guint64

	retC := C.g_unix_mount_at(c_mount_path, &c_time_read)
	retGo := UnixMountEntryNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// Compares two unix mounts.
/*

C function

g_unix_mount_compare
*/
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int32 {
	c_mount1 := (*C.GUnixMountEntry)(C.NULL)
	if mount1 != nil {
		c_mount1 = (*C.GUnixMountEntry)(mount1.ToC())
	}

	c_mount2 := (*C.GUnixMountEntry)(C.NULL)
	if mount2 != nil {
		c_mount2 = (*C.GUnixMountEntry)(mount2.ToC())
	}

	retC := C.g_unix_mount_compare(c_mount1, c_mount2)
	retGo := (int32)(retC)

	return retGo
}

// Frees a unix mount.
/*

C function

g_unix_mount_free
*/
func UnixMountFree(mountEntry *UnixMountEntry) {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	C.g_unix_mount_free(c_mount_entry)

	return
}

// Gets the device path for a unix mount.
/*

C function

g_unix_mount_get_device_path
*/
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_device_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the filesystem type for the unix mount.
/*

C function

g_unix_mount_get_fs_type
*/
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_fs_type(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// Gets the mount path for a unix mount.
/*

C function

g_unix_mount_get_mount_path
*/
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_get_mount_path(c_mount_entry)
	retGo := C.GoString(retC)

	return retGo
}

// Guesses whether a Unix mount can be ejected.
/*

C function

g_unix_mount_guess_can_eject
*/
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_can_eject(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// Guesses the icon of a Unix mount.
/*

C function

g_unix_mount_guess_icon
*/
func UnixMountGuessIcon(mountEntry *UnixMountEntry) *Icon {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_icon(c_mount_entry)
	retGo := IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Guesses the name of a Unix mount.
// The result is a translated string.
/*

C function

g_unix_mount_guess_name
*/
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_name(c_mount_entry)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Guesses whether a Unix mount should be displayed in the UI.
/*

C function

g_unix_mount_guess_should_display
*/
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_guess_should_display(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a unix mount is mounted read only.
/*

C function

g_unix_mount_is_readonly
*/
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_is_readonly(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a Unix mount is a system mount. This is the Boolean OR of
// g_unix_is_system_fs_type(), g_unix_is_system_device_path() and
// g_unix_is_mount_path_system_internal() on @mount_entry’s properties.
//
// The definition of what a ‘system’ mount entry is may change over time as new
// file system types and device paths are ignored.
/*

C function

g_unix_mount_is_system_internal
*/
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	c_mount_entry := (*C.GUnixMountEntry)(C.NULL)
	if mountEntry != nil {
		c_mount_entry = (*C.GUnixMountEntry)(mountEntry.ToC())
	}

	retC := C.g_unix_mount_is_system_internal(c_mount_entry)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if the unix mount points have changed since a given unix time.
/*

C function

g_unix_mount_points_changed_since
*/
func UnixMountPointsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mount_points_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a #GList of #GUnixMountPoint containing the unix mount points.
// If @time_read is set, it will be filled with the mount timestamp,
// allowing for checking if the mounts have changed with
// g_unix_mount_points_changed_since().
/*

C function

g_unix_mount_points_get
*/
func UnixMountPointsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mount_points_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}

// Checks if the unix mounts have changed since a given unix time.
/*

C function

g_unix_mounts_changed_since
*/
func UnixMountsChangedSince(time uint64) bool {
	c_time := (C.guint64)(time)

	retC := C.g_unix_mounts_changed_since(c_time)
	retGo := retC == C.TRUE

	return retGo
}

// Gets a #GList of #GUnixMountEntry containing the unix mounts.
// If @time_read is set, it will be filled with the mount
// timestamp, allowing for checking if the mounts have changed
// with g_unix_mounts_changed_since().
/*

C function

g_unix_mounts_get
*/
func UnixMountsGet() (*glib.List, uint64) {
	var c_time_read C.guint64

	retC := C.g_unix_mounts_get(&c_time_read)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	timeRead := (uint64)(c_time_read)

	return retGo, timeRead
}
