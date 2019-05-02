// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// MakeDirectoryWithParents is a wrapper around the C function g_file_make_directory_with_parents.
func (recv *File) MakeDirectoryWithParents(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_make_directory_with_parents((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Monitor is a wrapper around the C function g_file_monitor.
func (recv *File) Monitor(flags FileMonitorFlags, cancellable *Cancellable) (*FileMonitor, error) {
	c_flags := (C.GFileMonitorFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_monitor((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileMonitorNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// QueryFileType is a wrapper around the C function g_file_query_file_type.
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

// GuessContentTypeFinish is a wrapper around the C function g_mount_guess_content_type_finish.
func (recv *Mount) GuessContentTypeFinish(result *AsyncResult) ([]string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_guess_content_type_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GuessContentTypeSync is a wrapper around the C function g_mount_guess_content_type_sync.
func (recv *Mount) GuessContentTypeSync(forceRescan bool, cancellable *Cancellable) ([]string, error) {
	c_force_rescan :=
		boolToGboolean(forceRescan)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_mount_guess_content_type_sync((*C.GMount)(recv.native), c_force_rescan, c_cancellable, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetActivationRoot is a wrapper around the C function g_volume_get_activation_root.
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
