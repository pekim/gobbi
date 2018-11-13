// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Creates a directory and any parent directories that may not
// exist similar to 'mkdir -p'. If the file system does not support
// creating directories, this function will fail, setting @error to
// %G_IO_ERROR_NOT_SUPPORTED. If the directory itself already exists,
// this function will fail setting @error to %G_IO_ERROR_EXISTS, unlike
// the similar g_mkdir_with_parents().
//
// For a local #GFile the newly created directories will have the default
// (current) ownership and permissions of the current process.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function

g_file_make_directory_with_parents
*/
func (recv *File) MakeDirectoryWithParents(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_directory_with_parents((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Obtains a file or directory monitor for the given file,
// depending on the type of the file.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function

g_file_monitor
*/
func (recv *File) Monitor(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Utility function to inspect the #GFileType of a file. This is
// implemented using g_file_query_info() and as such does blocking I/O.
//
// The primary use case of this method is to check if a file is
// a regular file, directory, or symlink.
/*

C function

g_file_query_file_type
*/
func (recv *File) QueryFileType(flags FileQueryInfoFlags, cancellable *Cancellable) FileType {
	c_flags := (C.GFileQueryInfoFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_file_query_file_type((*C.GFile)(recv.native), c_flags, c_cancellable)
	retGo := (FileType)(retC)

	return retGo
}

// Unsupported : g_mount_guess_content_type : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_mount_guess_content_type_finish : no return type

// Unsupported : g_mount_guess_content_type_sync : no return type

// Gets the activation root for a #GVolume if it is known ahead of
// mount time. Returns %NULL otherwise. If not %NULL and if @volume
// is mounted, then the result of g_mount_get_root() on the
// #GMount object obtained from g_volume_get_mount() will always
// either be equal or a prefix of what this function returns. In
// other words, in code
//
// |[<!-- language="C" -->
// GMount *mount;
// GFile *mount_root
// GFile *volume_activation_root;
//
// mount = g_volume_get_mount (volume); // mounted, so never NULL
// mount_root = g_mount_get_root (mount);
// volume_activation_root = g_volume_get_activation_root (volume); // assume not NULL
// ]|
// then the expression
// |[<!-- language="C" -->
// (g_file_has_prefix (volume_activation_root, mount_root) ||
// g_file_equal (volume_activation_root, mount_root))
// ]|
// will always be %TRUE.
//
// Activation roots are typically used in #GVolumeMonitor
// implementations to find the underlying mount to shadow, see
// g_mount_is_shadowed() for more details.
/*

C function

g_volume_get_activation_root
*/
func (recv *Volume) GetActivationRoot() *File {
	retC := C.g_volume_get_activation_root((*C.GVolume)(recv.native))
	var retGo (*File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}
