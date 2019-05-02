// This is a generated file - DO NOT EDIT

package gio

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// AppLaunchContext is a wrapper around the C record GAppLaunchContext.
type AppLaunchContext struct {
	native *C.GAppLaunchContext
	// parent_instance : record
	// Private : priv
}

func AppLaunchContextNewFromC(u unsafe.Pointer) *AppLaunchContext {
	c := (*C.GAppLaunchContext)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppLaunchContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppLaunchContextNew is a wrapper around the C function g_app_launch_context_new.
func AppLaunchContextNew() *AppLaunchContext {
	retC := C.g_app_launch_context_new()
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ApplicationCommandLine is a wrapper around the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native *C.GApplicationCommandLine
	// Private : parent_instance
	// Private : priv
}

func ApplicationCommandLineNewFromC(u unsafe.Pointer) *ApplicationCommandLine {
	c := (*C.GApplicationCommandLine)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLine{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ApplicationCommandLine) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedInputStream is a wrapper around the C record GBufferedInputStream.
type BufferedInputStream struct {
	native *C.GBufferedInputStream
	// parent_instance : record
	// Private : priv
}

func BufferedInputStreamNewFromC(u unsafe.Pointer) *BufferedInputStream {
	c := (*C.GBufferedInputStream)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BufferedInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BufferedInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedInputStreamNew is a wrapper around the C function g_buffered_input_stream_new.
func BufferedInputStreamNew(baseStream *InputStream) *BufferedInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.g_buffered_input_stream_new(c_base_stream)
	retGo := BufferedInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BufferedInputStreamNewSized is a wrapper around the C function g_buffered_input_stream_new_sized.
func BufferedInputStreamNewSized(baseStream *InputStream, size uint64) *BufferedInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	c_size := (C.gsize)(size)

	retC := C.g_buffered_input_stream_new_sized(c_base_stream, c_size)
	retGo := BufferedInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BufferedOutputStream is a wrapper around the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native *C.GBufferedOutputStream
	// parent_instance : record
	// priv : record
}

func BufferedOutputStreamNewFromC(u unsafe.Pointer) *BufferedOutputStream {
	c := (*C.GBufferedOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BufferedOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BufferedOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedOutputStreamNew is a wrapper around the C function g_buffered_output_stream_new.
func BufferedOutputStreamNew(baseStream *OutputStream) *BufferedOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	retC := C.g_buffered_output_stream_new(c_base_stream)
	retGo := BufferedOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BufferedOutputStreamNewSized is a wrapper around the C function g_buffered_output_stream_new_sized.
func BufferedOutputStreamNewSized(baseStream *OutputStream, size uint64) *BufferedOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	c_size := (C.gsize)(size)

	retC := C.g_buffered_output_stream_new_sized(c_base_stream, c_size)
	retGo := BufferedOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BytesIcon is a wrapper around the C record GBytesIcon.
type BytesIcon struct {
	native *C.GBytesIcon
}

func BytesIconNewFromC(u unsafe.Pointer) *BytesIcon {
	c := (*C.GBytesIcon)(u)
	if c == nil {
		return nil
	}

	g := &BytesIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BytesIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BytesIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Cancellable is a wrapper around the C record GCancellable.
type Cancellable struct {
	native *C.GCancellable
	// parent_instance : record
	// Private : priv
}

func CancellableNewFromC(u unsafe.Pointer) *Cancellable {
	c := (*C.GCancellable)(u)
	if c == nil {
		return nil
	}

	g := &Cancellable{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Cancellable) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Cancellable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CancellableNew is a wrapper around the C function g_cancellable_new.
func CancellableNew() *Cancellable {
	retC := C.g_cancellable_new()
	retGo := CancellableNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CharsetConverter is a wrapper around the C record GCharsetConverter.
type CharsetConverter struct {
	native *C.GCharsetConverter
}

func CharsetConverterNewFromC(u unsafe.Pointer) *CharsetConverter {
	c := (*C.GCharsetConverter)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CharsetConverter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CharsetConverter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterInputStream is a wrapper around the C record GConverterInputStream.
type ConverterInputStream struct {
	native *C.GConverterInputStream
	// parent_instance : record
	// Private : priv
}

func ConverterInputStreamNewFromC(u unsafe.Pointer) *ConverterInputStream {
	c := (*C.GConverterInputStream)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ConverterInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ConverterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterInputStreamNew is a wrapper around the C function g_converter_input_stream_new.
func ConverterInputStreamNew(baseStream *InputStream, converter *Converter) *ConverterInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	c_converter := (*C.GConverter)(converter.ToC())

	retC := C.g_converter_input_stream_new(c_base_stream, c_converter)
	retGo := ConverterInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ConverterOutputStream is a wrapper around the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native *C.GConverterOutputStream
	// parent_instance : record
	// Private : priv
}

func ConverterOutputStreamNewFromC(u unsafe.Pointer) *ConverterOutputStream {
	c := (*C.GConverterOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ConverterOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ConverterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterOutputStreamNew is a wrapper around the C function g_converter_output_stream_new.
func ConverterOutputStreamNew(baseStream *OutputStream, converter *Converter) *ConverterOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	c_converter := (*C.GConverter)(converter.ToC())

	retC := C.g_converter_output_stream_new(c_base_stream, c_converter)
	retGo := ConverterOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DBusActionGroup is a wrapper around the C record GDBusActionGroup.
type DBusActionGroup struct {
	native *C.GDBusActionGroup
}

func DBusActionGroupNewFromC(u unsafe.Pointer) *DBusActionGroup {
	c := (*C.GDBusActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &DBusActionGroup{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusActionGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusMenuModel is a wrapper around the C record GDBusMenuModel.
type DBusMenuModel struct {
	native *C.GDBusMenuModel
}

func DBusMenuModelNewFromC(u unsafe.Pointer) *DBusMenuModel {
	c := (*C.GDBusMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &DBusMenuModel{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DBusMenuModel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DBusMenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataInputStream is a wrapper around the C record GDataInputStream.
type DataInputStream struct {
	native *C.GDataInputStream
	// parent_instance : record
	// Private : priv
}

func DataInputStreamNewFromC(u unsafe.Pointer) *DataInputStream {
	c := (*C.GDataInputStream)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DataInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DataInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataInputStreamNew is a wrapper around the C function g_data_input_stream_new.
func DataInputStreamNew(baseStream *InputStream) *DataInputStream {
	c_base_stream := (*C.GInputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GInputStream)(baseStream.ToC())
	}

	retC := C.g_data_input_stream_new(c_base_stream)
	retGo := DataInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DataOutputStream is a wrapper around the C record GDataOutputStream.
type DataOutputStream struct {
	native *C.GDataOutputStream
	// parent_instance : record
	// Private : priv
}

func DataOutputStreamNewFromC(u unsafe.Pointer) *DataOutputStream {
	c := (*C.GDataOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DataOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DataOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataOutputStreamNew is a wrapper around the C function g_data_output_stream_new.
func DataOutputStreamNew(baseStream *OutputStream) *DataOutputStream {
	c_base_stream := (*C.GOutputStream)(C.NULL)
	if baseStream != nil {
		c_base_stream = (*C.GOutputStream)(baseStream.ToC())
	}

	retC := C.g_data_output_stream_new(c_base_stream)
	retGo := DataOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DesktopAppInfo is a wrapper around the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native *C.GDesktopAppInfo
}

func DesktopAppInfoNewFromC(u unsafe.Pointer) *DesktopAppInfo {
	c := (*C.GDesktopAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DesktopAppInfo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DesktopAppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DesktopAppInfoNew is a wrapper around the C function g_desktop_app_info_new.
func DesktopAppInfoNew(desktopId string) *DesktopAppInfo {
	c_desktop_id := C.CString(desktopId)
	defer C.free(unsafe.Pointer(c_desktop_id))

	retC := C.g_desktop_app_info_new(c_desktop_id)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DesktopAppInfoNewFromFilename is a wrapper around the C function g_desktop_app_info_new_from_filename.
func DesktopAppInfoNewFromFilename(filename string) *DesktopAppInfo {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_desktop_app_info_new_from_filename(c_filename)
	retGo := DesktopAppInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Emblem is a wrapper around the C record GEmblem.
type Emblem struct {
	native *C.GEmblem
}

func EmblemNewFromC(u unsafe.Pointer) *Emblem {
	c := (*C.GEmblem)(u)
	if c == nil {
		return nil
	}

	g := &Emblem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Emblem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Emblem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EmblemedIcon is a wrapper around the C record GEmblemedIcon.
type EmblemedIcon struct {
	native *C.GEmblemedIcon
	// parent_instance : record
	// Private : priv
}

func EmblemedIconNewFromC(u unsafe.Pointer) *EmblemedIcon {
	c := (*C.GEmblemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EmblemedIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EmblemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileEnumerator is a wrapper around the C record GFileEnumerator.
type FileEnumerator struct {
	native *C.GFileEnumerator
	// parent_instance : record
	// Private : priv
}

func FileEnumeratorNewFromC(u unsafe.Pointer) *FileEnumerator {
	c := (*C.GFileEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumerator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileEnumerator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIOStream is a wrapper around the C record GFileIOStream.
type FileIOStream struct {
	native *C.GFileIOStream
	// parent_instance : record
	// Private : priv
}

func FileIOStreamNewFromC(u unsafe.Pointer) *FileIOStream {
	c := (*C.GFileIOStream)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileIOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIcon is a wrapper around the C record GFileIcon.
type FileIcon struct {
	native *C.GFileIcon
}

func FileIconNewFromC(u unsafe.Pointer) *FileIcon {
	c := (*C.GFileIcon)(u)
	if c == nil {
		return nil
	}

	g := &FileIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIconNew is a wrapper around the C function g_file_icon_new.
func FileIconNew(file *File) *FileIcon {
	c_file := (*C.GFile)(file.ToC())

	retC := C.g_file_icon_new(c_file)
	retGo := FileIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FileInfo is a wrapper around the C record GFileInfo.
type FileInfo struct {
	native *C.GFileInfo
}

func FileInfoNewFromC(u unsafe.Pointer) *FileInfo {
	c := (*C.GFileInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileInfo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileInfo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileInfoNew is a wrapper around the C function g_file_info_new.
func FileInfoNew() *FileInfo {
	retC := C.g_file_info_new()
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FileInputStream is a wrapper around the C record GFileInputStream.
type FileInputStream struct {
	native *C.GFileInputStream
	// parent_instance : record
	// Private : priv
}

func FileInputStreamNewFromC(u unsafe.Pointer) *FileInputStream {
	c := (*C.GFileInputStream)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileMonitor is a wrapper around the C record GFileMonitor.
type FileMonitor struct {
	native *C.GFileMonitor
	// parent_instance : record
	// Private : priv
}

func FileMonitorNewFromC(u unsafe.Pointer) *FileMonitor {
	c := (*C.GFileMonitor)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileOutputStream is a wrapper around the C record GFileOutputStream.
type FileOutputStream struct {
	native *C.GFileOutputStream
	// parent_instance : record
	// Private : priv
}

func FileOutputStreamNewFromC(u unsafe.Pointer) *FileOutputStream {
	c := (*C.GFileOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilenameCompleter is a wrapper around the C record GFilenameCompleter.
type FilenameCompleter struct {
	native *C.GFilenameCompleter
}

func FilenameCompleterNewFromC(u unsafe.Pointer) *FilenameCompleter {
	c := (*C.GFilenameCompleter)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FilenameCompleter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FilenameCompleter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilenameCompleterNew is a wrapper around the C function g_filename_completer_new.
func FilenameCompleterNew() *FilenameCompleter {
	retC := C.g_filename_completer_new()
	retGo := FilenameCompleterNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FilterInputStream is a wrapper around the C record GFilterInputStream.
type FilterInputStream struct {
	native *C.GFilterInputStream
	// parent_instance : record
	// base_stream : record
}

func FilterInputStreamNewFromC(u unsafe.Pointer) *FilterInputStream {
	c := (*C.GFilterInputStream)(u)
	if c == nil {
		return nil
	}

	g := &FilterInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FilterInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FilterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilterOutputStream is a wrapper around the C record GFilterOutputStream.
type FilterOutputStream struct {
	native *C.GFilterOutputStream
	// parent_instance : record
	// base_stream : record
}

func FilterOutputStreamNewFromC(u unsafe.Pointer) *FilterOutputStream {
	c := (*C.GFilterOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &FilterOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FilterOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FilterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOModule is a wrapper around the C record GIOModule.
type IOModule struct {
	native *C.GIOModule
}

func IOModuleNewFromC(u unsafe.Pointer) *IOModule {
	c := (*C.GIOModule)(u)
	if c == nil {
		return nil
	}

	g := &IOModule{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IOModule) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IOModule) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOModuleNew is a wrapper around the C function g_io_module_new.
func IOModuleNew(filename string) *IOModule {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_io_module_new(c_filename)
	retGo := IOModuleNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IOStream is a wrapper around the C record GIOStream.
type IOStream struct {
	native *C.GIOStream
	// parent_instance : record
	// Private : priv
}

func IOStreamNewFromC(u unsafe.Pointer) *IOStream {
	c := (*C.GIOStream)(u)
	if c == nil {
		return nil
	}

	g := &IOStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *IOStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *IOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddress is a wrapper around the C record GInetAddress.
type InetAddress struct {
	native *C.GInetAddress
	// parent_instance : record
	// Private : priv
}

func InetAddressNewFromC(u unsafe.Pointer) *InetAddress {
	c := (*C.GInetAddress)(u)
	if c == nil {
		return nil
	}

	g := &InetAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InetAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InetAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetSocketAddress is a wrapper around the C record GInetSocketAddress.
type InetSocketAddress struct {
	native *C.GInetSocketAddress
	// parent_instance : record
	// Private : priv
}

func InetSocketAddressNewFromC(u unsafe.Pointer) *InetSocketAddress {
	c := (*C.GInetSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InetSocketAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InetSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStream is a wrapper around the C record GInputStream.
type InputStream struct {
	native *C.GInputStream
	// parent_instance : record
	// Private : priv
}

func InputStreamNewFromC(u unsafe.Pointer) *InputStream {
	c := (*C.GInputStream)(u)
	if c == nil {
		return nil
	}

	g := &InputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListStore is a wrapper around the C record GListStore.
type ListStore struct {
	native *C.GListStore
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	c := (*C.GListStore)(u)
	if c == nil {
		return nil
	}

	g := &ListStore{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ListStore) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemoryInputStream is a wrapper around the C record GMemoryInputStream.
type MemoryInputStream struct {
	native *C.GMemoryInputStream
	// parent_instance : record
	// Private : priv
}

func MemoryInputStreamNewFromC(u unsafe.Pointer) *MemoryInputStream {
	c := (*C.GMemoryInputStream)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MemoryInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MemoryInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemoryInputStreamNew is a wrapper around the C function g_memory_input_stream_new.
func MemoryInputStreamNew() *MemoryInputStream {
	retC := C.g_memory_input_stream_new()
	retGo := MemoryInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// MemoryOutputStream is a wrapper around the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native *C.GMemoryOutputStream
	// parent_instance : record
	// Private : priv
}

func MemoryOutputStreamNewFromC(u unsafe.Pointer) *MemoryOutputStream {
	c := (*C.GMemoryOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MemoryOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MemoryOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_memory_output_stream_new : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// MountOperation is a wrapper around the C record GMountOperation.
type MountOperation struct {
	native *C.GMountOperation
	// parent_instance : record
	// priv : record
}

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	c := (*C.GMountOperation)(u)
	if c == nil {
		return nil
	}

	g := &MountOperation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MountOperation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperationNew is a wrapper around the C function g_mount_operation_new.
func MountOperationNew() *MountOperation {
	retC := C.g_mount_operation_new()
	retGo := MountOperationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// NativeVolumeMonitor is a wrapper around the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native *C.GNativeVolumeMonitor
	// parent_instance : record
}

func NativeVolumeMonitorNewFromC(u unsafe.Pointer) *NativeVolumeMonitor {
	c := (*C.GNativeVolumeMonitor)(u)
	if c == nil {
		return nil
	}

	g := &NativeVolumeMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NativeVolumeMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkAddress is a wrapper around the C record GNetworkAddress.
type NetworkAddress struct {
	native *C.GNetworkAddress
	// parent_instance : record
	// Private : priv
}

func NetworkAddressNewFromC(u unsafe.Pointer) *NetworkAddress {
	c := (*C.GNetworkAddress)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NetworkAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NetworkAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkService is a wrapper around the C record GNetworkService.
type NetworkService struct {
	native *C.GNetworkService
	// parent_instance : record
	// Private : priv
}

func NetworkServiceNewFromC(u unsafe.Pointer) *NetworkService {
	c := (*C.GNetworkService)(u)
	if c == nil {
		return nil
	}

	g := &NetworkService{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NetworkService) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NetworkService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStream is a wrapper around the C record GOutputStream.
type OutputStream struct {
	native *C.GOutputStream
	// parent_instance : record
	// Private : priv
}

func OutputStreamNewFromC(u unsafe.Pointer) *OutputStream {
	c := (*C.GOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &OutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *OutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *OutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Permission is a wrapper around the C record GPermission.
type Permission struct {
	native *C.GPermission
	// parent_instance : record
	// Private : priv
}

func PermissionNewFromC(u unsafe.Pointer) *Permission {
	c := (*C.GPermission)(u)
	if c == nil {
		return nil
	}

	g := &Permission{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Permission) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Permission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressEnumerator is a wrapper around the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native *C.GProxyAddressEnumerator
	// parent_instance : record
	// priv : record
}

func ProxyAddressEnumeratorNewFromC(u unsafe.Pointer) *ProxyAddressEnumerator {
	c := (*C.GProxyAddressEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumerator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ProxyAddressEnumerator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Resolver is a wrapper around the C record GResolver.
type Resolver struct {
	native *C.GResolver
	// parent_instance : record
	// priv : record
}

func ResolverNewFromC(u unsafe.Pointer) *Resolver {
	c := (*C.GResolver)(u)
	if c == nil {
		return nil
	}

	g := &Resolver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Resolver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Resolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Settings is a wrapper around the C record GSettings.
type Settings struct {
	native *C.GSettings
	// parent_instance : record
	// priv : record
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	c := (*C.GSettings)(u)
	if c == nil {
		return nil
	}

	g := &Settings{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Settings) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsBackend is a wrapper around the C record GSettingsBackend.
type SettingsBackend struct {
	native *C.GSettingsBackend
	// parent_instance : record
	// Private : priv
}

func SettingsBackendNewFromC(u unsafe.Pointer) *SettingsBackend {
	c := (*C.GSettingsBackend)(u)
	if c == nil {
		return nil
	}

	g := &SettingsBackend{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SettingsBackend) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SettingsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleAction is a wrapper around the C record GSimpleAction.
type SimpleAction struct {
	native *C.GSimpleAction
}

func SimpleActionNewFromC(u unsafe.Pointer) *SimpleAction {
	c := (*C.GSimpleAction)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAction{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleAction) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleAsyncResult is a wrapper around the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native *C.GSimpleAsyncResult
}

func SimpleAsyncResultNewFromC(u unsafe.Pointer) *SimpleAsyncResult {
	c := (*C.GSimpleAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResult{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleAsyncResult) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SimplePermission is a wrapper around the C record GSimplePermission.
type SimplePermission struct {
	native *C.GSimplePermission
}

func SimplePermissionNewFromC(u unsafe.Pointer) *SimplePermission {
	c := (*C.GSimplePermission)(u)
	if c == nil {
		return nil
	}

	g := &SimplePermission{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimplePermission) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimplePermission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleProxyResolver is a wrapper around the C record GSimpleProxyResolver.
type SimpleProxyResolver struct {
	native *C.GSimpleProxyResolver
	// parent_instance : record
	// Private : priv
}

func SimpleProxyResolverNewFromC(u unsafe.Pointer) *SimpleProxyResolver {
	c := (*C.GSimpleProxyResolver)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SimpleProxyResolver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SimpleProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddress is a wrapper around the C record GSocketAddress.
type SocketAddress struct {
	native *C.GSocketAddress
	// parent_instance : record
}

func SocketAddressNewFromC(u unsafe.Pointer) *SocketAddress {
	c := (*C.GSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddressEnumerator is a wrapper around the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native *C.GSocketAddressEnumerator
	// parent_instance : record
}

func SocketAddressEnumeratorNewFromC(u unsafe.Pointer) *SocketAddressEnumerator {
	c := (*C.GSocketAddressEnumerator)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressEnumerator{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketAddressEnumerator) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketControlMessage is a wrapper around the C record GSocketControlMessage.
type SocketControlMessage struct {
	native *C.GSocketControlMessage
	// parent_instance : record
	// priv : record
}

func SocketControlMessageNewFromC(u unsafe.Pointer) *SocketControlMessage {
	c := (*C.GSocketControlMessage)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketControlMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketControlMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Task is a wrapper around the C record GTask.
type Task struct {
	native *C.GTask
}

func TaskNewFromC(u unsafe.Pointer) *Task {
	c := (*C.GTask)(u)
	if c == nil {
		return nil
	}

	g := &Task{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Task) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Task) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TcpWrapperConnection is a wrapper around the C record GTcpWrapperConnection.
type TcpWrapperConnection struct {
	native *C.GTcpWrapperConnection
	// parent_instance : record
	// priv : record
}

func TcpWrapperConnectionNewFromC(u unsafe.Pointer) *TcpWrapperConnection {
	c := (*C.GTcpWrapperConnection)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TcpWrapperConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TcpWrapperConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemedIcon is a wrapper around the C record GThemedIcon.
type ThemedIcon struct {
	native *C.GThemedIcon
}

func ThemedIconNewFromC(u unsafe.Pointer) *ThemedIcon {
	c := (*C.GThemedIcon)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIcon{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ThemedIcon) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ThemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemedIconNew is a wrapper around the C function g_themed_icon_new.
func ThemedIconNew(iconname string) *ThemedIcon {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	retC := C.g_themed_icon_new(c_iconname)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ThemedIconNewFromNames is a wrapper around the C function g_themed_icon_new_from_names.
func ThemedIconNewFromNames(iconnames []string) *ThemedIcon {
	c_iconnames_array := make([]*C.char, len(iconnames)+1, len(iconnames)+1)
	for i, item := range iconnames {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_iconnames_array[i] = c
	}
	c_iconnames_array[len(iconnames)] = nil
	c_iconnames_arrayPtr := &c_iconnames_array[0]
	c_iconnames := (**C.char)(unsafe.Pointer(c_iconnames_arrayPtr))

	c_len := (C.int)(len(iconnames))

	retC := C.g_themed_icon_new_from_names(c_iconnames, c_len)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ThemedIconNewWithDefaultFallbacks is a wrapper around the C function g_themed_icon_new_with_default_fallbacks.
func ThemedIconNewWithDefaultFallbacks(iconname string) *ThemedIcon {
	c_iconname := C.CString(iconname)
	defer C.free(unsafe.Pointer(c_iconname))

	retC := C.g_themed_icon_new_with_default_fallbacks(c_iconname)
	retGo := ThemedIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixConnection is a wrapper around the C record GUnixConnection.
type UnixConnection struct {
	native *C.GUnixConnection
	// parent_instance : record
	// priv : record
}

func UnixConnectionNewFromC(u unsafe.Pointer) *UnixConnection {
	c := (*C.GUnixConnection)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDList is a wrapper around the C record GUnixFDList.
type UnixFDList struct {
	native *C.GUnixFDList
	// parent_instance : record
	// priv : record
}

func UnixFDListNewFromC(u unsafe.Pointer) *UnixFDList {
	c := (*C.GUnixFDList)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDList{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixFDList) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixFDList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDMessage is a wrapper around the C record GUnixFDMessage.
type UnixFDMessage struct {
	native *C.GUnixFDMessage
	// parent_instance : record
	// priv : record
}

func UnixFDMessageNewFromC(u unsafe.Pointer) *UnixFDMessage {
	c := (*C.GUnixFDMessage)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessage{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixFDMessage) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixFDMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixInputStream is a wrapper around the C record GUnixInputStream.
type UnixInputStream struct {
	native *C.GUnixInputStream
	// parent_instance : record
	// Private : priv
}

func UnixInputStreamNewFromC(u unsafe.Pointer) *UnixInputStream {
	c := (*C.GUnixInputStream)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixInputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixInputStreamNew is a wrapper around the C function g_unix_input_stream_new.
func UnixInputStreamNew(fd int32, closeFd bool) *UnixInputStream {
	c_fd := (C.gint)(fd)

	c_close_fd :=
		boolToGboolean(closeFd)

	retC := C.g_unix_input_stream_new(c_fd, c_close_fd)
	retGo := UnixInputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixMountMonitor is a wrapper around the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native *C.GUnixMountMonitor
}

func UnixMountMonitorNewFromC(u unsafe.Pointer) *UnixMountMonitor {
	c := (*C.GUnixMountMonitor)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixMountMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixMountMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixMountMonitorNew is a wrapper around the C function g_unix_mount_monitor_new.
func UnixMountMonitorNew() *UnixMountMonitor {
	retC := C.g_unix_mount_monitor_new()
	retGo := UnixMountMonitorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixOutputStream is a wrapper around the C record GUnixOutputStream.
type UnixOutputStream struct {
	native *C.GUnixOutputStream
	// parent_instance : record
	// Private : priv
}

func UnixOutputStreamNewFromC(u unsafe.Pointer) *UnixOutputStream {
	c := (*C.GUnixOutputStream)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStream{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixOutputStream) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixOutputStreamNew is a wrapper around the C function g_unix_output_stream_new.
func UnixOutputStreamNew(fd int32, closeFd bool) *UnixOutputStream {
	c_fd := (C.gint)(fd)

	c_close_fd :=
		boolToGboolean(closeFd)

	retC := C.g_unix_output_stream_new(c_fd, c_close_fd)
	retGo := UnixOutputStreamNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixSocketAddress is a wrapper around the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native *C.GUnixSocketAddress
	// parent_instance : record
	// Private : priv
}

func UnixSocketAddressNewFromC(u unsafe.Pointer) *UnixSocketAddress {
	c := (*C.GUnixSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddress{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UnixSocketAddress) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UnixSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixSocketAddressNewAbstract is a wrapper around the C function g_unix_socket_address_new_abstract.
func UnixSocketAddressNewAbstract(path []rune) *UnixSocketAddress {
	c_path_array := make([]C.gchar, len(path)+1, len(path)+1)
	for i, item := range path {
		c := (C.gchar)(item)
		c_path_array[i] = c
	}
	c_path_array[len(path)] = 0
	c_path_arrayPtr := &c_path_array[0]
	c_path := (*C.gchar)(unsafe.Pointer(c_path_arrayPtr))

	c_path_len := (C.gint)(len(path))

	retC := C.g_unix_socket_address_new_abstract(c_path, c_path_len)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Vfs is a wrapper around the C record GVfs.
type Vfs struct {
	native *C.GVfs
	// parent_instance : record
}

func VfsNewFromC(u unsafe.Pointer) *Vfs {
	c := (*C.GVfs)(u)
	if c == nil {
		return nil
	}

	g := &Vfs{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Vfs) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Vfs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VolumeMonitor is a wrapper around the C record GVolumeMonitor.
type VolumeMonitor struct {
	native *C.GVolumeMonitor
	// parent_instance : record
	// Private : priv
}

func VolumeMonitorNewFromC(u unsafe.Pointer) *VolumeMonitor {
	c := (*C.GVolumeMonitor)(u)
	if c == nil {
		return nil
	}

	g := &VolumeMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *VolumeMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *VolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ZlibCompressor is a wrapper around the C record GZlibCompressor.
type ZlibCompressor struct {
	native *C.GZlibCompressor
}

func ZlibCompressorNewFromC(u unsafe.Pointer) *ZlibCompressor {
	c := (*C.GZlibCompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ZlibCompressor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ZlibCompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ZlibDecompressor is a wrapper around the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native *C.GZlibDecompressor
}

func ZlibDecompressorNewFromC(u unsafe.Pointer) *ZlibDecompressor {
	c := (*C.GZlibDecompressor)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ZlibDecompressor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ZlibDecompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : FILE_ATTRIBUTE_FILESYSTEM_REMOTE

// Blacklisted : SETTINGS_BACKEND_EXTENSION_POINT_NAME

// Unsupported : g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon

// Unsupported : g_io_scheduler_push_job : unsupported parameter job_func : no type generator for IOSchedulerJobFunc (GIOSchedulerJobFunc) for param job_func

// Blacklisted : g_keyfile_settings_backend_new

// Unsupported : g_simple_async_report_error_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_report_gerror_in_idle : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Action is a wrapper around the C record GAction.
type Action struct {
	native *C.GAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionGroup is a wrapper around the C record GActionGroup.
type ActionGroup struct {
	native *C.GActionGroup
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionMap is a wrapper around the C record GActionMap.
type ActionMap struct {
	native *C.GActionMap
}

func ActionMapNewFromC(u unsafe.Pointer) *ActionMap {
	c := (*C.GActionMap)(u)
	if c == nil {
		return nil
	}

	g := &ActionMap{native: c}

	return g
}

func (recv *ActionMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppInfo is a wrapper around the C record GAppInfo.
type AppInfo struct {
	native *C.GAppInfo
}

func AppInfoNewFromC(u unsafe.Pointer) *AppInfo {
	c := (*C.GAppInfo)(u)
	if c == nil {
		return nil
	}

	g := &AppInfo{native: c}

	return g
}

func (recv *AppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AsyncResult is a wrapper around the C record GAsyncResult.
type AsyncResult struct {
	native *C.GAsyncResult
}

func AsyncResultNewFromC(u unsafe.Pointer) *AsyncResult {
	c := (*C.GAsyncResult)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResult{native: c}

	return g
}

func (recv *AsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObject is a wrapper around the C record GDBusObject.
type DBusObject struct {
	native *C.GDBusObject
}

func DBusObjectNewFromC(u unsafe.Pointer) *DBusObject {
	c := (*C.GDBusObject)(u)
	if c == nil {
		return nil
	}

	g := &DBusObject{native: c}

	return g
}

func (recv *DBusObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectManager is a wrapper around the C record GDBusObjectManager.
type DBusObjectManager struct {
	native *C.GDBusObjectManager
}

func DBusObjectManagerNewFromC(u unsafe.Pointer) *DBusObjectManager {
	c := (*C.GDBusObjectManager)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManager{native: c}

	return g
}

func (recv *DBusObjectManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DesktopAppInfoLookup is a wrapper around the C record GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native *C.GDesktopAppInfoLookup
}

func DesktopAppInfoLookupNewFromC(u unsafe.Pointer) *DesktopAppInfoLookup {
	c := (*C.GDesktopAppInfoLookup)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookup{native: c}

	return g
}

func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Drive is a wrapper around the C record GDrive.
type Drive struct {
	native *C.GDrive
}

func DriveNewFromC(u unsafe.Pointer) *Drive {
	c := (*C.GDrive)(u)
	if c == nil {
		return nil
	}

	g := &Drive{native: c}

	return g
}

func (recv *Drive) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// File is a wrapper around the C record GFile.
type File struct {
	native *C.GFile
}

func FileNewFromC(u unsafe.Pointer) *File {
	c := (*C.GFile)(u)
	if c == nil {
		return nil
	}

	g := &File{native: c}

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileDescriptorBased is a wrapper around the C record GFileDescriptorBased.
type FileDescriptorBased struct {
	native *C.GFileDescriptorBased
}

func FileDescriptorBasedNewFromC(u unsafe.Pointer) *FileDescriptorBased {
	c := (*C.GFileDescriptorBased)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBased{native: c}

	return g
}

func (recv *FileDescriptorBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Icon is a wrapper around the C record GIcon.
type Icon struct {
	native *C.GIcon
}

func IconNewFromC(u unsafe.Pointer) *Icon {
	c := (*C.GIcon)(u)
	if c == nil {
		return nil
	}

	g := &Icon{native: c}

	return g
}

func (recv *Icon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListModel is a wrapper around the C record GListModel.
type ListModel struct {
	native *C.GListModel
}

func ListModelNewFromC(u unsafe.Pointer) *ListModel {
	c := (*C.GListModel)(u)
	if c == nil {
		return nil
	}

	g := &ListModel{native: c}

	return g
}

func (recv *ListModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LoadableIcon is a wrapper around the C record GLoadableIcon.
type LoadableIcon struct {
	native *C.GLoadableIcon
}

func LoadableIconNewFromC(u unsafe.Pointer) *LoadableIcon {
	c := (*C.GLoadableIcon)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIcon{native: c}

	return g
}

func (recv *LoadableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Mount is a wrapper around the C record GMount.
type Mount struct {
	native *C.GMount
}

func MountNewFromC(u unsafe.Pointer) *Mount {
	c := (*C.GMount)(u)
	if c == nil {
		return nil
	}

	g := &Mount{native: c}

	return g
}

func (recv *Mount) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RemoteActionGroup is a wrapper around the C record GRemoteActionGroup.
type RemoteActionGroup struct {
	native *C.GRemoteActionGroup
}

func RemoteActionGroupNewFromC(u unsafe.Pointer) *RemoteActionGroup {
	c := (*C.GRemoteActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &RemoteActionGroup{native: c}

	return g
}

func (recv *RemoteActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Seekable is a wrapper around the C record GSeekable.
type Seekable struct {
	native *C.GSeekable
}

func SeekableNewFromC(u unsafe.Pointer) *Seekable {
	c := (*C.GSeekable)(u)
	if c == nil {
		return nil
	}

	g := &Seekable{native: c}

	return g
}

func (recv *Seekable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketConnectable is a wrapper around the C record GSocketConnectable.
type SocketConnectable struct {
	native *C.GSocketConnectable
}

func SocketConnectableNewFromC(u unsafe.Pointer) *SocketConnectable {
	c := (*C.GSocketConnectable)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectable{native: c}

	return g
}

func (recv *SocketConnectable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Volume is a wrapper around the C record GVolume.
type Volume struct {
	native *C.GVolume
}

func VolumeNewFromC(u unsafe.Pointer) *Volume {
	c := (*C.GVolume)(u)
	if c == nil {
		return nil
	}

	g := &Volume{native: c}

	return g
}

func (recv *Volume) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionEntry is a wrapper around the C record GActionEntry.
type ActionEntry struct {
	native *C.GActionEntry
	Name   string
	// no type for activate
	ParameterType string
	State         string
	// no type for change_state
	// Private : padding
}

func ActionEntryNewFromC(u unsafe.Pointer) *ActionEntry {
	c := (*C.GActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &ActionEntry{
		Name:          C.GoString(c.name),
		ParameterType: C.GoString(c.parameter_type),
		State:         C.GoString(c.state),
		native:        c,
	}

	return g
}

func (recv *ActionEntry) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.parameter_type =
		C.CString(recv.ParameterType)
	recv.native.state =
		C.CString(recv.State)

	return (unsafe.Pointer)(recv.native)
}

// AppInfoIface is a wrapper around the C record GAppInfoIface.
type AppInfoIface struct {
	native *C.GAppInfoIface
	// g_iface : record
	// no type for dup
	// no type for equal
	// no type for get_id
	// no type for get_name
	// no type for get_description
	// no type for get_executable
	// no type for get_icon
	// no type for launch
	// no type for supports_uris
	// no type for supports_files
	// no type for launch_uris
	// no type for should_show
	// no type for set_as_default_for_type
	// no type for set_as_default_for_extension
	// no type for add_supports_type
	// no type for can_remove_supports_type
	// no type for remove_supports_type
	// no type for can_delete
	// no type for do_delete
	// no type for get_commandline
	// no type for get_display_name
	// no type for set_as_last_used_for_type
	// no type for get_supported_types
}

func AppInfoIfaceNewFromC(u unsafe.Pointer) *AppInfoIface {
	c := (*C.GAppInfoIface)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoIface{native: c}

	return g
}

func (recv *AppInfoIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppLaunchContextClass is a wrapper around the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native *C.GAppLaunchContextClass
	// parent_class : record
	// no type for get_display
	// no type for get_startup_notify_id
	// no type for launch_failed
	// no type for launched
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

func AppLaunchContextClassNewFromC(u unsafe.Pointer) *AppLaunchContextClass {
	c := (*C.GAppLaunchContextClass)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContextClass{native: c}

	return g
}

func (recv *AppLaunchContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppLaunchContextPrivate is a wrapper around the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native *C.GAppLaunchContextPrivate
}

func AppLaunchContextPrivateNewFromC(u unsafe.Pointer) *AppLaunchContextPrivate {
	c := (*C.GAppLaunchContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AppLaunchContextPrivate{native: c}

	return g
}

func (recv *AppLaunchContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationCommandLinePrivate is a wrapper around the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native *C.GApplicationCommandLinePrivate
}

func ApplicationCommandLinePrivateNewFromC(u unsafe.Pointer) *ApplicationCommandLinePrivate {
	c := (*C.GApplicationCommandLinePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationCommandLinePrivate{native: c}

	return g
}

func (recv *ApplicationCommandLinePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ApplicationPrivate is a wrapper around the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native *C.GApplicationPrivate
}

func ApplicationPrivateNewFromC(u unsafe.Pointer) *ApplicationPrivate {
	c := (*C.GApplicationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationPrivate{native: c}

	return g
}

func (recv *ApplicationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AsyncResultIface is a wrapper around the C record GAsyncResultIface.
type AsyncResultIface struct {
	native *C.GAsyncResultIface
	// g_iface : record
	// no type for get_user_data
	// no type for get_source_object
	// no type for is_tagged
}

func AsyncResultIfaceNewFromC(u unsafe.Pointer) *AsyncResultIface {
	c := (*C.GAsyncResultIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncResultIface{native: c}

	return g
}

func (recv *AsyncResultIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedInputStreamClass is a wrapper around the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native *C.GBufferedInputStreamClass
	// parent_class : record
	// no type for fill
	// no type for fill_async
	// no type for fill_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func BufferedInputStreamClassNewFromC(u unsafe.Pointer) *BufferedInputStreamClass {
	c := (*C.GBufferedInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStreamClass{native: c}

	return g
}

func (recv *BufferedInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedInputStreamPrivate is a wrapper around the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native *C.GBufferedInputStreamPrivate
}

func BufferedInputStreamPrivateNewFromC(u unsafe.Pointer) *BufferedInputStreamPrivate {
	c := (*C.GBufferedInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferedInputStreamPrivate{native: c}

	return g
}

func (recv *BufferedInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedOutputStreamClass is a wrapper around the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native *C.GBufferedOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func BufferedOutputStreamClassNewFromC(u unsafe.Pointer) *BufferedOutputStreamClass {
	c := (*C.GBufferedOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStreamClass{native: c}

	return g
}

func (recv *BufferedOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedOutputStreamPrivate is a wrapper around the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native *C.GBufferedOutputStreamPrivate
}

func BufferedOutputStreamPrivateNewFromC(u unsafe.Pointer) *BufferedOutputStreamPrivate {
	c := (*C.GBufferedOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferedOutputStreamPrivate{native: c}

	return g
}

func (recv *BufferedOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CancellableClass is a wrapper around the C record GCancellableClass.
type CancellableClass struct {
	native *C.GCancellableClass
	// parent_class : record
	// no type for cancelled
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func CancellableClassNewFromC(u unsafe.Pointer) *CancellableClass {
	c := (*C.GCancellableClass)(u)
	if c == nil {
		return nil
	}

	g := &CancellableClass{native: c}

	return g
}

func (recv *CancellableClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CancellablePrivate is a wrapper around the C record GCancellablePrivate.
type CancellablePrivate struct {
	native *C.GCancellablePrivate
}

func CancellablePrivateNewFromC(u unsafe.Pointer) *CancellablePrivate {
	c := (*C.GCancellablePrivate)(u)
	if c == nil {
		return nil
	}

	g := &CancellablePrivate{native: c}

	return g
}

func (recv *CancellablePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CharsetConverterClass is a wrapper around the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native *C.GCharsetConverterClass
	// parent_class : record
}

func CharsetConverterClassNewFromC(u unsafe.Pointer) *CharsetConverterClass {
	c := (*C.GCharsetConverterClass)(u)
	if c == nil {
		return nil
	}

	g := &CharsetConverterClass{native: c}

	return g
}

func (recv *CharsetConverterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterInputStreamClass is a wrapper around the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native *C.GConverterInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ConverterInputStreamClassNewFromC(u unsafe.Pointer) *ConverterInputStreamClass {
	c := (*C.GConverterInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStreamClass{native: c}

	return g
}

func (recv *ConverterInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterInputStreamPrivate is a wrapper around the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native *C.GConverterInputStreamPrivate
}

func ConverterInputStreamPrivateNewFromC(u unsafe.Pointer) *ConverterInputStreamPrivate {
	c := (*C.GConverterInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ConverterInputStreamPrivate{native: c}

	return g
}

func (recv *ConverterInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterOutputStreamClass is a wrapper around the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native *C.GConverterOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ConverterOutputStreamClassNewFromC(u unsafe.Pointer) *ConverterOutputStreamClass {
	c := (*C.GConverterOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStreamClass{native: c}

	return g
}

func (recv *ConverterOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ConverterOutputStreamPrivate is a wrapper around the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native *C.GConverterOutputStreamPrivate
}

func ConverterOutputStreamPrivateNewFromC(u unsafe.Pointer) *ConverterOutputStreamPrivate {
	c := (*C.GConverterOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ConverterOutputStreamPrivate{native: c}

	return g
}

func (recv *ConverterOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusInterfaceSkeletonPrivate is a wrapper around the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native *C.GDBusInterfaceSkeletonPrivate
}

func DBusInterfaceSkeletonPrivateNewFromC(u unsafe.Pointer) *DBusInterfaceSkeletonPrivate {
	c := (*C.GDBusInterfaceSkeletonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusInterfaceSkeletonPrivate{native: c}

	return g
}

func (recv *DBusInterfaceSkeletonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectManagerClientPrivate is a wrapper around the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native *C.GDBusObjectManagerClientPrivate
}

func DBusObjectManagerClientPrivateNewFromC(u unsafe.Pointer) *DBusObjectManagerClientPrivate {
	c := (*C.GDBusObjectManagerClientPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerClientPrivate{native: c}

	return g
}

func (recv *DBusObjectManagerClientPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectManagerServerPrivate is a wrapper around the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native *C.GDBusObjectManagerServerPrivate
}

func DBusObjectManagerServerPrivateNewFromC(u unsafe.Pointer) *DBusObjectManagerServerPrivate {
	c := (*C.GDBusObjectManagerServerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectManagerServerPrivate{native: c}

	return g
}

func (recv *DBusObjectManagerServerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectProxyPrivate is a wrapper around the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native *C.GDBusObjectProxyPrivate
}

func DBusObjectProxyPrivateNewFromC(u unsafe.Pointer) *DBusObjectProxyPrivate {
	c := (*C.GDBusObjectProxyPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectProxyPrivate{native: c}

	return g
}

func (recv *DBusObjectProxyPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectSkeletonPrivate is a wrapper around the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native *C.GDBusObjectSkeletonPrivate
}

func DBusObjectSkeletonPrivateNewFromC(u unsafe.Pointer) *DBusObjectSkeletonPrivate {
	c := (*C.GDBusObjectSkeletonPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusObjectSkeletonPrivate{native: c}

	return g
}

func (recv *DBusObjectSkeletonPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusProxyPrivate is a wrapper around the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native *C.GDBusProxyPrivate
}

func DBusProxyPrivateNewFromC(u unsafe.Pointer) *DBusProxyPrivate {
	c := (*C.GDBusProxyPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DBusProxyPrivate{native: c}

	return g
}

func (recv *DBusProxyPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataInputStreamClass is a wrapper around the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native *C.GDataInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func DataInputStreamClassNewFromC(u unsafe.Pointer) *DataInputStreamClass {
	c := (*C.GDataInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStreamClass{native: c}

	return g
}

func (recv *DataInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataInputStreamPrivate is a wrapper around the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native *C.GDataInputStreamPrivate
}

func DataInputStreamPrivateNewFromC(u unsafe.Pointer) *DataInputStreamPrivate {
	c := (*C.GDataInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DataInputStreamPrivate{native: c}

	return g
}

func (recv *DataInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataOutputStreamClass is a wrapper around the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native *C.GDataOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func DataOutputStreamClassNewFromC(u unsafe.Pointer) *DataOutputStreamClass {
	c := (*C.GDataOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStreamClass{native: c}

	return g
}

func (recv *DataOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataOutputStreamPrivate is a wrapper around the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native *C.GDataOutputStreamPrivate
}

func DataOutputStreamPrivateNewFromC(u unsafe.Pointer) *DataOutputStreamPrivate {
	c := (*C.GDataOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DataOutputStreamPrivate{native: c}

	return g
}

func (recv *DataOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DesktopAppInfoClass is a wrapper around the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native *C.GDesktopAppInfoClass
	// parent_class : record
}

func DesktopAppInfoClassNewFromC(u unsafe.Pointer) *DesktopAppInfoClass {
	c := (*C.GDesktopAppInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoClass{native: c}

	return g
}

func (recv *DesktopAppInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DesktopAppInfoLookupIface is a wrapper around the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native *C.GDesktopAppInfoLookupIface
	// g_iface : record
	// no type for get_default_for_uri_scheme
}

func DesktopAppInfoLookupIfaceNewFromC(u unsafe.Pointer) *DesktopAppInfoLookupIface {
	c := (*C.GDesktopAppInfoLookupIface)(u)
	if c == nil {
		return nil
	}

	g := &DesktopAppInfoLookupIface{native: c}

	return g
}

func (recv *DesktopAppInfoLookupIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DriveIface is a wrapper around the C record GDriveIface.
type DriveIface struct {
	native *C.GDriveIface
	// g_iface : record
	// no type for changed
	// no type for disconnected
	// no type for eject_button
	// no type for get_name
	// no type for get_icon
	// no type for has_volumes
	// no type for get_volumes
	// no type for is_media_removable
	// no type for has_media
	// no type for is_media_check_automatic
	// no type for can_eject
	// no type for can_poll_for_media
	// no type for eject
	// no type for eject_finish
	// no type for poll_for_media
	// no type for poll_for_media_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for get_start_stop_type
	// no type for can_start
	// no type for can_start_degraded
	// no type for start
	// no type for start_finish
	// no type for can_stop
	// no type for stop
	// no type for stop_finish
	// no type for stop_button
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
	// no type for is_removable
}

func DriveIfaceNewFromC(u unsafe.Pointer) *DriveIface {
	c := (*C.GDriveIface)(u)
	if c == nil {
		return nil
	}

	g := &DriveIface{native: c}

	return g
}

func (recv *DriveIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EmblemClass is a wrapper around the C record GEmblemClass.
type EmblemClass struct {
	native *C.GEmblemClass
}

func EmblemClassNewFromC(u unsafe.Pointer) *EmblemClass {
	c := (*C.GEmblemClass)(u)
	if c == nil {
		return nil
	}

	g := &EmblemClass{native: c}

	return g
}

func (recv *EmblemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EmblemedIconClass is a wrapper around the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native *C.GEmblemedIconClass
	// parent_class : record
}

func EmblemedIconClassNewFromC(u unsafe.Pointer) *EmblemedIconClass {
	c := (*C.GEmblemedIconClass)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIconClass{native: c}

	return g
}

func (recv *EmblemedIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EmblemedIconPrivate is a wrapper around the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native *C.GEmblemedIconPrivate
}

func EmblemedIconPrivateNewFromC(u unsafe.Pointer) *EmblemedIconPrivate {
	c := (*C.GEmblemedIconPrivate)(u)
	if c == nil {
		return nil
	}

	g := &EmblemedIconPrivate{native: c}

	return g
}

func (recv *EmblemedIconPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileAttributeInfo is a wrapper around the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native *C.GFileAttributeInfo
	Name   string
	Type   FileAttributeType
	Flags  FileAttributeInfoFlags
}

func FileAttributeInfoNewFromC(u unsafe.Pointer) *FileAttributeInfo {
	c := (*C.GFileAttributeInfo)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeInfo{
		Flags:  (FileAttributeInfoFlags)(c.flags),
		Name:   C.GoString(c.name),
		Type:   (FileAttributeType)(c._type),
		native: c,
	}

	return g
}

func (recv *FileAttributeInfo) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native._type =
		(C.GFileAttributeType)(recv.Type)
	recv.native.flags =
		(C.GFileAttributeInfoFlags)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// FileAttributeInfoList is a wrapper around the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native *C.GFileAttributeInfoList
	// infos : record
	NInfos int32
}

func FileAttributeInfoListNewFromC(u unsafe.Pointer) *FileAttributeInfoList {
	c := (*C.GFileAttributeInfoList)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeInfoList{
		NInfos: (int32)(c.n_infos),
		native: c,
	}

	return g
}

func (recv *FileAttributeInfoList) ToC() unsafe.Pointer {
	recv.native.n_infos =
		(C.int)(recv.NInfos)

	return (unsafe.Pointer)(recv.native)
}

// FileAttributeInfoListNew is a wrapper around the C function g_file_attribute_info_list_new.
func FileAttributeInfoListNew() *FileAttributeInfoList {
	retC := C.g_file_attribute_info_list_new()
	retGo := FileAttributeInfoListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileAttributeMatcher is a wrapper around the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native *C.GFileAttributeMatcher
}

func FileAttributeMatcherNewFromC(u unsafe.Pointer) *FileAttributeMatcher {
	c := (*C.GFileAttributeMatcher)(u)
	if c == nil {
		return nil
	}

	g := &FileAttributeMatcher{native: c}

	return g
}

func (recv *FileAttributeMatcher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileAttributeMatcherNew is a wrapper around the C function g_file_attribute_matcher_new.
func FileAttributeMatcherNew(attributes string) *FileAttributeMatcher {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	retC := C.g_file_attribute_matcher_new(c_attributes)
	retGo := FileAttributeMatcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileDescriptorBasedIface is a wrapper around the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native *C.GFileDescriptorBasedIface
	// g_iface : record
	// no type for get_fd
}

func FileDescriptorBasedIfaceNewFromC(u unsafe.Pointer) *FileDescriptorBasedIface {
	c := (*C.GFileDescriptorBasedIface)(u)
	if c == nil {
		return nil
	}

	g := &FileDescriptorBasedIface{native: c}

	return g
}

func (recv *FileDescriptorBasedIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileEnumeratorClass is a wrapper around the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native *C.GFileEnumeratorClass
	// parent_class : record
	// no type for next_file
	// no type for close_fn
	// no type for next_files_async
	// no type for next_files_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

func FileEnumeratorClassNewFromC(u unsafe.Pointer) *FileEnumeratorClass {
	c := (*C.GFileEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumeratorClass{native: c}

	return g
}

func (recv *FileEnumeratorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileEnumeratorPrivate is a wrapper around the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native *C.GFileEnumeratorPrivate
}

func FileEnumeratorPrivateNewFromC(u unsafe.Pointer) *FileEnumeratorPrivate {
	c := (*C.GFileEnumeratorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileEnumeratorPrivate{native: c}

	return g
}

func (recv *FileEnumeratorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIOStreamClass is a wrapper around the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native *C.GFileIOStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileIOStreamClassNewFromC(u unsafe.Pointer) *FileIOStreamClass {
	c := (*C.GFileIOStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStreamClass{native: c}

	return g
}

func (recv *FileIOStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIOStreamPrivate is a wrapper around the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native *C.GFileIOStreamPrivate
}

func FileIOStreamPrivateNewFromC(u unsafe.Pointer) *FileIOStreamPrivate {
	c := (*C.GFileIOStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileIOStreamPrivate{native: c}

	return g
}

func (recv *FileIOStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIconClass is a wrapper around the C record GFileIconClass.
type FileIconClass struct {
	native *C.GFileIconClass
}

func FileIconClassNewFromC(u unsafe.Pointer) *FileIconClass {
	c := (*C.GFileIconClass)(u)
	if c == nil {
		return nil
	}

	g := &FileIconClass{native: c}

	return g
}

func (recv *FileIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIface is a wrapper around the C record GFileIface.
type FileIface struct {
	native *C.GFileIface
	// g_iface : record
	// no type for dup
	// no type for hash
	// no type for equal
	// no type for is_native
	// no type for has_uri_scheme
	// no type for get_uri_scheme
	// no type for get_basename
	// no type for get_path
	// no type for get_uri
	// no type for get_parse_name
	// no type for get_parent
	// no type for prefix_matches
	// no type for get_relative_path
	// no type for resolve_relative_path
	// no type for get_child_for_display_name
	// no type for enumerate_children
	// no type for enumerate_children_async
	// no type for enumerate_children_finish
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for query_filesystem_info
	// no type for query_filesystem_info_async
	// no type for query_filesystem_info_finish
	// no type for find_enclosing_mount
	// no type for find_enclosing_mount_async
	// no type for find_enclosing_mount_finish
	// no type for set_display_name
	// no type for set_display_name_async
	// no type for set_display_name_finish
	// no type for query_settable_attributes
	// no type for _query_settable_attributes_async
	// no type for _query_settable_attributes_finish
	// no type for query_writable_namespaces
	// no type for _query_writable_namespaces_async
	// no type for _query_writable_namespaces_finish
	// no type for set_attribute
	// no type for set_attributes_from_info
	// no type for set_attributes_async
	// no type for set_attributes_finish
	// no type for read_fn
	// no type for read_async
	// no type for read_finish
	// no type for append_to
	// no type for append_to_async
	// no type for append_to_finish
	// no type for create
	// no type for create_async
	// no type for create_finish
	// no type for replace
	// no type for replace_async
	// no type for replace_finish
	// no type for delete_file
	// no type for delete_file_async
	// no type for delete_file_finish
	// no type for trash
	// no type for trash_async
	// no type for trash_finish
	// no type for make_directory
	// no type for make_directory_async
	// no type for make_directory_finish
	// no type for make_symbolic_link
	// no type for _make_symbolic_link_async
	// no type for _make_symbolic_link_finish
	// no type for copy
	// no type for copy_async
	// no type for copy_finish
	// no type for move
	// no type for _move_async
	// no type for _move_finish
	// no type for mount_mountable
	// no type for mount_mountable_finish
	// no type for unmount_mountable
	// no type for unmount_mountable_finish
	// no type for eject_mountable
	// no type for eject_mountable_finish
	// no type for mount_enclosing_volume
	// no type for mount_enclosing_volume_finish
	// no type for monitor_dir
	// no type for monitor_file
	// no type for open_readwrite
	// no type for open_readwrite_async
	// no type for open_readwrite_finish
	// no type for create_readwrite
	// no type for create_readwrite_async
	// no type for create_readwrite_finish
	// no type for replace_readwrite
	// no type for replace_readwrite_async
	// no type for replace_readwrite_finish
	// no type for start_mountable
	// no type for start_mountable_finish
	// no type for stop_mountable
	// no type for stop_mountable_finish
	SupportsThreadContexts bool
	// no type for unmount_mountable_with_operation
	// no type for unmount_mountable_with_operation_finish
	// no type for eject_mountable_with_operation
	// no type for eject_mountable_with_operation_finish
	// no type for poll_mountable
	// no type for poll_mountable_finish
	// no type for measure_disk_usage
	// no type for measure_disk_usage_async
	// no type for measure_disk_usage_finish
}

func FileIfaceNewFromC(u unsafe.Pointer) *FileIface {
	c := (*C.GFileIface)(u)
	if c == nil {
		return nil
	}

	g := &FileIface{
		SupportsThreadContexts: c.supports_thread_contexts == C.TRUE,
		native:                 c,
	}

	return g
}

func (recv *FileIface) ToC() unsafe.Pointer {
	recv.native.supports_thread_contexts =
		boolToGboolean(recv.SupportsThreadContexts)

	return (unsafe.Pointer)(recv.native)
}

// FileInfoClass is a wrapper around the C record GFileInfoClass.
type FileInfoClass struct {
	native *C.GFileInfoClass
}

func FileInfoClassNewFromC(u unsafe.Pointer) *FileInfoClass {
	c := (*C.GFileInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &FileInfoClass{native: c}

	return g
}

func (recv *FileInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileInputStreamClass is a wrapper around the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native *C.GFileInputStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileInputStreamClassNewFromC(u unsafe.Pointer) *FileInputStreamClass {
	c := (*C.GFileInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStreamClass{native: c}

	return g
}

func (recv *FileInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileInputStreamPrivate is a wrapper around the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native *C.GFileInputStreamPrivate
}

func FileInputStreamPrivateNewFromC(u unsafe.Pointer) *FileInputStreamPrivate {
	c := (*C.GFileInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileInputStreamPrivate{native: c}

	return g
}

func (recv *FileInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileMonitorClass is a wrapper around the C record GFileMonitorClass.
type FileMonitorClass struct {
	native *C.GFileMonitorClass
	// parent_class : record
	// no type for changed
	// no type for cancel
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileMonitorClassNewFromC(u unsafe.Pointer) *FileMonitorClass {
	c := (*C.GFileMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitorClass{native: c}

	return g
}

func (recv *FileMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileMonitorPrivate is a wrapper around the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native *C.GFileMonitorPrivate
}

func FileMonitorPrivateNewFromC(u unsafe.Pointer) *FileMonitorPrivate {
	c := (*C.GFileMonitorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileMonitorPrivate{native: c}

	return g
}

func (recv *FileMonitorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileOutputStreamClass is a wrapper around the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native *C.GFileOutputStreamClass
	// parent_class : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
	// no type for query_info
	// no type for query_info_async
	// no type for query_info_finish
	// no type for get_etag
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func FileOutputStreamClassNewFromC(u unsafe.Pointer) *FileOutputStreamClass {
	c := (*C.GFileOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStreamClass{native: c}

	return g
}

func (recv *FileOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileOutputStreamPrivate is a wrapper around the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native *C.GFileOutputStreamPrivate
}

func FileOutputStreamPrivateNewFromC(u unsafe.Pointer) *FileOutputStreamPrivate {
	c := (*C.GFileOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileOutputStreamPrivate{native: c}

	return g
}

func (recv *FileOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilenameCompleterClass is a wrapper around the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native *C.GFilenameCompleterClass
	// parent_class : record
	// no type for got_completion_data
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilenameCompleterClassNewFromC(u unsafe.Pointer) *FilenameCompleterClass {
	c := (*C.GFilenameCompleterClass)(u)
	if c == nil {
		return nil
	}

	g := &FilenameCompleterClass{native: c}

	return g
}

func (recv *FilenameCompleterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilterInputStreamClass is a wrapper around the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native *C.GFilterInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilterInputStreamClassNewFromC(u unsafe.Pointer) *FilterInputStreamClass {
	c := (*C.GFilterInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FilterInputStreamClass{native: c}

	return g
}

func (recv *FilterInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilterOutputStreamClass is a wrapper around the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native *C.GFilterOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
}

func FilterOutputStreamClassNewFromC(u unsafe.Pointer) *FilterOutputStreamClass {
	c := (*C.GFilterOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &FilterOutputStreamClass{native: c}

	return g
}

func (recv *FilterOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOExtension is a wrapper around the C record GIOExtension.
type IOExtension struct {
	native *C.GIOExtension
}

func IOExtensionNewFromC(u unsafe.Pointer) *IOExtension {
	c := (*C.GIOExtension)(u)
	if c == nil {
		return nil
	}

	g := &IOExtension{native: c}

	return g
}

func (recv *IOExtension) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOExtensionPoint is a wrapper around the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native *C.GIOExtensionPoint
}

func IOExtensionPointNewFromC(u unsafe.Pointer) *IOExtensionPoint {
	c := (*C.GIOExtensionPoint)(u)
	if c == nil {
		return nil
	}

	g := &IOExtensionPoint{native: c}

	return g
}

func (recv *IOExtensionPoint) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOModuleClass is a wrapper around the C record GIOModuleClass.
type IOModuleClass struct {
	native *C.GIOModuleClass
}

func IOModuleClassNewFromC(u unsafe.Pointer) *IOModuleClass {
	c := (*C.GIOModuleClass)(u)
	if c == nil {
		return nil
	}

	g := &IOModuleClass{native: c}

	return g
}

func (recv *IOModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOSchedulerJob is a wrapper around the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native *C.GIOSchedulerJob
}

func IOSchedulerJobNewFromC(u unsafe.Pointer) *IOSchedulerJob {
	c := (*C.GIOSchedulerJob)(u)
	if c == nil {
		return nil
	}

	g := &IOSchedulerJob{native: c}

	return g
}

func (recv *IOSchedulerJob) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStreamAdapter is a wrapper around the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native *C.GIOStreamAdapter
}

func IOStreamAdapterNewFromC(u unsafe.Pointer) *IOStreamAdapter {
	c := (*C.GIOStreamAdapter)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamAdapter{native: c}

	return g
}

func (recv *IOStreamAdapter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStreamClass is a wrapper around the C record GIOStreamClass.
type IOStreamClass struct {
	native *C.GIOStreamClass
	// parent_class : record
	// no type for get_input_stream
	// no type for get_output_stream
	// no type for close_fn
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

func IOStreamClassNewFromC(u unsafe.Pointer) *IOStreamClass {
	c := (*C.GIOStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamClass{native: c}

	return g
}

func (recv *IOStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStreamPrivate is a wrapper around the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native *C.GIOStreamPrivate
}

func IOStreamPrivateNewFromC(u unsafe.Pointer) *IOStreamPrivate {
	c := (*C.GIOStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &IOStreamPrivate{native: c}

	return g
}

func (recv *IOStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IconIface is a wrapper around the C record GIconIface.
type IconIface struct {
	native *C.GIconIface
	// g_iface : record
	// no type for hash
	// no type for equal
	// no type for to_tokens
	// no type for from_tokens
	// no type for serialize
}

func IconIfaceNewFromC(u unsafe.Pointer) *IconIface {
	c := (*C.GIconIface)(u)
	if c == nil {
		return nil
	}

	g := &IconIface{native: c}

	return g
}

func (recv *IconIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddressClass is a wrapper around the C record GInetAddressClass.
type InetAddressClass struct {
	native *C.GInetAddressClass
	// parent_class : record
	// no type for to_string
	// no type for to_bytes
}

func InetAddressClassNewFromC(u unsafe.Pointer) *InetAddressClass {
	c := (*C.GInetAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressClass{native: c}

	return g
}

func (recv *InetAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddressMaskClass is a wrapper around the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native *C.GInetAddressMaskClass
	// parent_class : record
}

func InetAddressMaskClassNewFromC(u unsafe.Pointer) *InetAddressMaskClass {
	c := (*C.GInetAddressMaskClass)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMaskClass{native: c}

	return g
}

func (recv *InetAddressMaskClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddressMaskPrivate is a wrapper around the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native *C.GInetAddressMaskPrivate
}

func InetAddressMaskPrivateNewFromC(u unsafe.Pointer) *InetAddressMaskPrivate {
	c := (*C.GInetAddressMaskPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMaskPrivate{native: c}

	return g
}

func (recv *InetAddressMaskPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddressPrivate is a wrapper around the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native *C.GInetAddressPrivate
}

func InetAddressPrivateNewFromC(u unsafe.Pointer) *InetAddressPrivate {
	c := (*C.GInetAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressPrivate{native: c}

	return g
}

func (recv *InetAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetSocketAddressClass is a wrapper around the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native *C.GInetSocketAddressClass
	// parent_class : record
}

func InetSocketAddressClassNewFromC(u unsafe.Pointer) *InetSocketAddressClass {
	c := (*C.GInetSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddressClass{native: c}

	return g
}

func (recv *InetSocketAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetSocketAddressPrivate is a wrapper around the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native *C.GInetSocketAddressPrivate
}

func InetSocketAddressPrivateNewFromC(u unsafe.Pointer) *InetSocketAddressPrivate {
	c := (*C.GInetSocketAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InetSocketAddressPrivate{native: c}

	return g
}

func (recv *InetSocketAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStreamClass is a wrapper around the C record GInputStreamClass.
type InputStreamClass struct {
	native *C.GInputStreamClass
	// parent_class : record
	// no type for read_fn
	// no type for skip
	// no type for close_fn
	// no type for read_async
	// no type for read_finish
	// no type for skip_async
	// no type for skip_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func InputStreamClassNewFromC(u unsafe.Pointer) *InputStreamClass {
	c := (*C.GInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &InputStreamClass{native: c}

	return g
}

func (recv *InputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputStreamPrivate is a wrapper around the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native *C.GInputStreamPrivate
}

func InputStreamPrivateNewFromC(u unsafe.Pointer) *InputStreamPrivate {
	c := (*C.GInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InputStreamPrivate{native: c}

	return g
}

func (recv *InputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListStoreClass is a wrapper around the C record GListStoreClass.
type ListStoreClass struct {
	native *C.GListStoreClass
	// parent_class : record
}

func ListStoreClassNewFromC(u unsafe.Pointer) *ListStoreClass {
	c := (*C.GListStoreClass)(u)
	if c == nil {
		return nil
	}

	g := &ListStoreClass{native: c}

	return g
}

func (recv *ListStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LoadableIconIface is a wrapper around the C record GLoadableIconIface.
type LoadableIconIface struct {
	native *C.GLoadableIconIface
	// g_iface : record
	// no type for load
	// no type for load_async
	// no type for load_finish
}

func LoadableIconIfaceNewFromC(u unsafe.Pointer) *LoadableIconIface {
	c := (*C.GLoadableIconIface)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIconIface{native: c}

	return g
}

func (recv *LoadableIconIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemoryInputStreamClass is a wrapper around the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native *C.GMemoryInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func MemoryInputStreamClassNewFromC(u unsafe.Pointer) *MemoryInputStreamClass {
	c := (*C.GMemoryInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStreamClass{native: c}

	return g
}

func (recv *MemoryInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemoryInputStreamPrivate is a wrapper around the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native *C.GMemoryInputStreamPrivate
}

func MemoryInputStreamPrivateNewFromC(u unsafe.Pointer) *MemoryInputStreamPrivate {
	c := (*C.GMemoryInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MemoryInputStreamPrivate{native: c}

	return g
}

func (recv *MemoryInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemoryOutputStreamClass is a wrapper around the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native *C.GMemoryOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func MemoryOutputStreamClassNewFromC(u unsafe.Pointer) *MemoryOutputStreamClass {
	c := (*C.GMemoryOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStreamClass{native: c}

	return g
}

func (recv *MemoryOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemoryOutputStreamPrivate is a wrapper around the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native *C.GMemoryOutputStreamPrivate
}

func MemoryOutputStreamPrivateNewFromC(u unsafe.Pointer) *MemoryOutputStreamPrivate {
	c := (*C.GMemoryOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MemoryOutputStreamPrivate{native: c}

	return g
}

func (recv *MemoryOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuAttributeIterClass is a wrapper around the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native *C.GMenuAttributeIterClass
	// parent_class : record
	// no type for get_next
}

func MenuAttributeIterClassNewFromC(u unsafe.Pointer) *MenuAttributeIterClass {
	c := (*C.GMenuAttributeIterClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIterClass{native: c}

	return g
}

func (recv *MenuAttributeIterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuAttributeIterPrivate is a wrapper around the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native *C.GMenuAttributeIterPrivate
}

func MenuAttributeIterPrivateNewFromC(u unsafe.Pointer) *MenuAttributeIterPrivate {
	c := (*C.GMenuAttributeIterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIterPrivate{native: c}

	return g
}

func (recv *MenuAttributeIterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuLinkIterClass is a wrapper around the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native *C.GMenuLinkIterClass
	// parent_class : record
	// no type for get_next
}

func MenuLinkIterClassNewFromC(u unsafe.Pointer) *MenuLinkIterClass {
	c := (*C.GMenuLinkIterClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIterClass{native: c}

	return g
}

func (recv *MenuLinkIterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuLinkIterPrivate is a wrapper around the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native *C.GMenuLinkIterPrivate
}

func MenuLinkIterPrivateNewFromC(u unsafe.Pointer) *MenuLinkIterPrivate {
	c := (*C.GMenuLinkIterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIterPrivate{native: c}

	return g
}

func (recv *MenuLinkIterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuModelClass is a wrapper around the C record GMenuModelClass.
type MenuModelClass struct {
	native *C.GMenuModelClass
	// parent_class : record
	// no type for is_mutable
	// no type for get_n_items
	// no type for get_item_attributes
	// no type for iterate_item_attributes
	// no type for get_item_attribute_value
	// no type for get_item_links
	// no type for iterate_item_links
	// no type for get_item_link
}

func MenuModelClassNewFromC(u unsafe.Pointer) *MenuModelClass {
	c := (*C.GMenuModelClass)(u)
	if c == nil {
		return nil
	}

	g := &MenuModelClass{native: c}

	return g
}

func (recv *MenuModelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuModelPrivate is a wrapper around the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native *C.GMenuModelPrivate
}

func MenuModelPrivateNewFromC(u unsafe.Pointer) *MenuModelPrivate {
	c := (*C.GMenuModelPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MenuModelPrivate{native: c}

	return g
}

func (recv *MenuModelPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountIface is a wrapper around the C record GMountIface.
type MountIface struct {
	native *C.GMountIface
	// g_iface : record
	// no type for changed
	// no type for unmounted
	// no type for get_root
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_volume
	// no type for get_drive
	// no type for can_unmount
	// no type for can_eject
	// no type for unmount
	// no type for unmount_finish
	// no type for eject
	// no type for eject_finish
	// no type for remount
	// no type for remount_finish
	// no type for guess_content_type
	// no type for guess_content_type_finish
	// no type for guess_content_type_sync
	// no type for pre_unmount
	// no type for unmount_with_operation
	// no type for unmount_with_operation_finish
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_default_location
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

func MountIfaceNewFromC(u unsafe.Pointer) *MountIface {
	c := (*C.GMountIface)(u)
	if c == nil {
		return nil
	}

	g := &MountIface{native: c}

	return g
}

func (recv *MountIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperationClass is a wrapper around the C record GMountOperationClass.
type MountOperationClass struct {
	native *C.GMountOperationClass
	// parent_class : record
	// no type for ask_password
	// no type for ask_question
	// no type for reply
	// no type for aborted
	// no type for show_processes
	// no type for show_unmount_progress
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
}

func MountOperationClassNewFromC(u unsafe.Pointer) *MountOperationClass {
	c := (*C.GMountOperationClass)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationClass{native: c}

	return g
}

func (recv *MountOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperationPrivate is a wrapper around the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native *C.GMountOperationPrivate
}

func MountOperationPrivateNewFromC(u unsafe.Pointer) *MountOperationPrivate {
	c := (*C.GMountOperationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MountOperationPrivate{native: c}

	return g
}

func (recv *MountOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NativeSocketAddress is a wrapper around the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native *C.GNativeSocketAddress
}

func NativeSocketAddressNewFromC(u unsafe.Pointer) *NativeSocketAddress {
	c := (*C.GNativeSocketAddress)(u)
	if c == nil {
		return nil
	}

	g := &NativeSocketAddress{native: c}

	return g
}

func (recv *NativeSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NativeVolumeMonitorClass is a wrapper around the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native *C.GNativeVolumeMonitorClass
	// parent_class : record
	// no type for get_mount_for_mount_path
}

func NativeVolumeMonitorClassNewFromC(u unsafe.Pointer) *NativeVolumeMonitorClass {
	c := (*C.GNativeVolumeMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &NativeVolumeMonitorClass{native: c}

	return g
}

func (recv *NativeVolumeMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkAddressClass is a wrapper around the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native *C.GNetworkAddressClass
	// parent_class : record
}

func NetworkAddressClassNewFromC(u unsafe.Pointer) *NetworkAddressClass {
	c := (*C.GNetworkAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddressClass{native: c}

	return g
}

func (recv *NetworkAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkAddressPrivate is a wrapper around the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native *C.GNetworkAddressPrivate
}

func NetworkAddressPrivateNewFromC(u unsafe.Pointer) *NetworkAddressPrivate {
	c := (*C.GNetworkAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NetworkAddressPrivate{native: c}

	return g
}

func (recv *NetworkAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkServiceClass is a wrapper around the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native *C.GNetworkServiceClass
	// parent_class : record
}

func NetworkServiceClassNewFromC(u unsafe.Pointer) *NetworkServiceClass {
	c := (*C.GNetworkServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &NetworkServiceClass{native: c}

	return g
}

func (recv *NetworkServiceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkServicePrivate is a wrapper around the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native *C.GNetworkServicePrivate
}

func NetworkServicePrivateNewFromC(u unsafe.Pointer) *NetworkServicePrivate {
	c := (*C.GNetworkServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &NetworkServicePrivate{native: c}

	return g
}

func (recv *NetworkServicePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStreamClass is a wrapper around the C record GOutputStreamClass.
type OutputStreamClass struct {
	native *C.GOutputStreamClass
	// parent_class : record
	// no type for write_fn
	// no type for splice
	// no type for flush
	// no type for close_fn
	// no type for write_async
	// no type for write_finish
	// no type for splice_async
	// no type for splice_finish
	// no type for flush_async
	// no type for flush_finish
	// no type for close_async
	// no type for close_finish
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
}

func OutputStreamClassNewFromC(u unsafe.Pointer) *OutputStreamClass {
	c := (*C.GOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &OutputStreamClass{native: c}

	return g
}

func (recv *OutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OutputStreamPrivate is a wrapper around the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native *C.GOutputStreamPrivate
}

func OutputStreamPrivateNewFromC(u unsafe.Pointer) *OutputStreamPrivate {
	c := (*C.GOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &OutputStreamPrivate{native: c}

	return g
}

func (recv *OutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PermissionClass is a wrapper around the C record GPermissionClass.
type PermissionClass struct {
	native *C.GPermissionClass
	// parent_class : record
	// no type for acquire
	// no type for acquire_async
	// no type for acquire_finish
	// no type for release
	// no type for release_async
	// no type for release_finish
	// no type for reserved
}

func PermissionClassNewFromC(u unsafe.Pointer) *PermissionClass {
	c := (*C.GPermissionClass)(u)
	if c == nil {
		return nil
	}

	g := &PermissionClass{native: c}

	return g
}

func (recv *PermissionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PermissionPrivate is a wrapper around the C record GPermissionPrivate.
type PermissionPrivate struct {
	native *C.GPermissionPrivate
}

func PermissionPrivateNewFromC(u unsafe.Pointer) *PermissionPrivate {
	c := (*C.GPermissionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PermissionPrivate{native: c}

	return g
}

func (recv *PermissionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressEnumeratorClass is a wrapper around the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native *C.GProxyAddressEnumeratorClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
}

func ProxyAddressEnumeratorClassNewFromC(u unsafe.Pointer) *ProxyAddressEnumeratorClass {
	c := (*C.GProxyAddressEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumeratorClass{native: c}

	return g
}

func (recv *ProxyAddressEnumeratorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressEnumeratorPrivate is a wrapper around the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native *C.GProxyAddressEnumeratorPrivate
}

func ProxyAddressEnumeratorPrivateNewFromC(u unsafe.Pointer) *ProxyAddressEnumeratorPrivate {
	c := (*C.GProxyAddressEnumeratorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressEnumeratorPrivate{native: c}

	return g
}

func (recv *ProxyAddressEnumeratorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressPrivate is a wrapper around the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native *C.GProxyAddressPrivate
}

func ProxyAddressPrivateNewFromC(u unsafe.Pointer) *ProxyAddressPrivate {
	c := (*C.GProxyAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ProxyAddressPrivate{native: c}

	return g
}

func (recv *ProxyAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyResolverInterface is a wrapper around the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native *C.GProxyResolverInterface
	// g_iface : record
	// no type for is_supported
	// no type for lookup
	// no type for lookup_async
	// no type for lookup_finish
}

func ProxyResolverInterfaceNewFromC(u unsafe.Pointer) *ProxyResolverInterface {
	c := (*C.GProxyResolverInterface)(u)
	if c == nil {
		return nil
	}

	g := &ProxyResolverInterface{native: c}

	return g
}

func (recv *ProxyResolverInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ResolverClass is a wrapper around the C record GResolverClass.
type ResolverClass struct {
	native *C.GResolverClass
	// parent_class : record
	// no type for reload
	// no type for lookup_by_name
	// no type for lookup_by_name_async
	// no type for lookup_by_name_finish
	// no type for lookup_by_address
	// no type for lookup_by_address_async
	// no type for lookup_by_address_finish
	// no type for lookup_service
	// no type for lookup_service_async
	// no type for lookup_service_finish
	// no type for lookup_records
	// no type for lookup_records_async
	// no type for lookup_records_finish
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func ResolverClassNewFromC(u unsafe.Pointer) *ResolverClass {
	c := (*C.GResolverClass)(u)
	if c == nil {
		return nil
	}

	g := &ResolverClass{native: c}

	return g
}

func (recv *ResolverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ResolverPrivate is a wrapper around the C record GResolverPrivate.
type ResolverPrivate struct {
	native *C.GResolverPrivate
}

func ResolverPrivateNewFromC(u unsafe.Pointer) *ResolverPrivate {
	c := (*C.GResolverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ResolverPrivate{native: c}

	return g
}

func (recv *ResolverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SeekableIface is a wrapper around the C record GSeekableIface.
type SeekableIface struct {
	native *C.GSeekableIface
	// g_iface : record
	// no type for tell
	// no type for can_seek
	// no type for seek
	// no type for can_truncate
	// no type for truncate_fn
}

func SeekableIfaceNewFromC(u unsafe.Pointer) *SeekableIface {
	c := (*C.GSeekableIface)(u)
	if c == nil {
		return nil
	}

	g := &SeekableIface{native: c}

	return g
}

func (recv *SeekableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GSettingsBackendClass

// Blacklisted : GSettingsBackendPrivate

// SettingsClass is a wrapper around the C record GSettingsClass.
type SettingsClass struct {
	native *C.GSettingsClass
	// parent_class : record
	// no type for writable_changed
	// no type for changed
	// no type for writable_change_event
	// no type for change_event
	// no type for padding
}

func SettingsClassNewFromC(u unsafe.Pointer) *SettingsClass {
	c := (*C.GSettingsClass)(u)
	if c == nil {
		return nil
	}

	g := &SettingsClass{native: c}

	return g
}

func (recv *SettingsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsPrivate is a wrapper around the C record GSettingsPrivate.
type SettingsPrivate struct {
	native *C.GSettingsPrivate
}

func SettingsPrivateNewFromC(u unsafe.Pointer) *SettingsPrivate {
	c := (*C.GSettingsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SettingsPrivate{native: c}

	return g
}

func (recv *SettingsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsSchemaKey is a wrapper around the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native *C.GSettingsSchemaKey
}

func SettingsSchemaKeyNewFromC(u unsafe.Pointer) *SettingsSchemaKey {
	c := (*C.GSettingsSchemaKey)(u)
	if c == nil {
		return nil
	}

	g := &SettingsSchemaKey{native: c}

	return g
}

func (recv *SettingsSchemaKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleActionGroupClass is a wrapper around the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native *C.GSimpleActionGroupClass
	// Private : parent_class
	// Private : padding
}

func SimpleActionGroupClassNewFromC(u unsafe.Pointer) *SimpleActionGroupClass {
	c := (*C.GSimpleActionGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroupClass{native: c}

	return g
}

func (recv *SimpleActionGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleActionGroupPrivate is a wrapper around the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native *C.GSimpleActionGroupPrivate
}

func SimpleActionGroupPrivateNewFromC(u unsafe.Pointer) *SimpleActionGroupPrivate {
	c := (*C.GSimpleActionGroupPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SimpleActionGroupPrivate{native: c}

	return g
}

func (recv *SimpleActionGroupPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleAsyncResultClass is a wrapper around the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native *C.GSimpleAsyncResultClass
}

func SimpleAsyncResultClassNewFromC(u unsafe.Pointer) *SimpleAsyncResultClass {
	c := (*C.GSimpleAsyncResultClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleAsyncResultClass{native: c}

	return g
}

func (recv *SimpleAsyncResultClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleProxyResolverClass is a wrapper around the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native *C.GSimpleProxyResolverClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func SimpleProxyResolverClassNewFromC(u unsafe.Pointer) *SimpleProxyResolverClass {
	c := (*C.GSimpleProxyResolverClass)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolverClass{native: c}

	return g
}

func (recv *SimpleProxyResolverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleProxyResolverPrivate is a wrapper around the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native *C.GSimpleProxyResolverPrivate
}

func SimpleProxyResolverPrivateNewFromC(u unsafe.Pointer) *SimpleProxyResolverPrivate {
	c := (*C.GSimpleProxyResolverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SimpleProxyResolverPrivate{native: c}

	return g
}

func (recv *SimpleProxyResolverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddressClass is a wrapper around the C record GSocketAddressClass.
type SocketAddressClass struct {
	native *C.GSocketAddressClass
	// parent_class : record
	// no type for get_family
	// no type for get_native_size
	// no type for to_native
}

func SocketAddressClassNewFromC(u unsafe.Pointer) *SocketAddressClass {
	c := (*C.GSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressClass{native: c}

	return g
}

func (recv *SocketAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketAddressEnumeratorClass is a wrapper around the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native *C.GSocketAddressEnumeratorClass
	// parent_class : record
	// no type for next
	// no type for next_async
	// no type for next_finish
}

func SocketAddressEnumeratorClassNewFromC(u unsafe.Pointer) *SocketAddressEnumeratorClass {
	c := (*C.GSocketAddressEnumeratorClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketAddressEnumeratorClass{native: c}

	return g
}

func (recv *SocketAddressEnumeratorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketClass is a wrapper around the C record GSocketClass.
type SocketClass struct {
	native *C.GSocketClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
	// no type for _g_reserved7
	// no type for _g_reserved8
	// no type for _g_reserved9
	// no type for _g_reserved10
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.GSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketClientClass is a wrapper around the C record GSocketClientClass.
type SocketClientClass struct {
	native *C.GSocketClientClass
	// parent_class : record
	// no type for event
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
}

func SocketClientClassNewFromC(u unsafe.Pointer) *SocketClientClass {
	c := (*C.GSocketClientClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClientClass{native: c}

	return g
}

func (recv *SocketClientClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketClientPrivate is a wrapper around the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native *C.GSocketClientPrivate
}

func SocketClientPrivateNewFromC(u unsafe.Pointer) *SocketClientPrivate {
	c := (*C.GSocketClientPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketClientPrivate{native: c}

	return g
}

func (recv *SocketClientPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketConnectableIface is a wrapper around the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native *C.GSocketConnectableIface
	// g_iface : record
	// no type for enumerate
	// no type for proxy_enumerate
	// no type for to_string
}

func SocketConnectableIfaceNewFromC(u unsafe.Pointer) *SocketConnectableIface {
	c := (*C.GSocketConnectableIface)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectableIface{native: c}

	return g
}

func (recv *SocketConnectableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketConnectionClass is a wrapper around the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native *C.GSocketConnectionClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketConnectionClassNewFromC(u unsafe.Pointer) *SocketConnectionClass {
	c := (*C.GSocketConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectionClass{native: c}

	return g
}

func (recv *SocketConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketConnectionPrivate is a wrapper around the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native *C.GSocketConnectionPrivate
}

func SocketConnectionPrivateNewFromC(u unsafe.Pointer) *SocketConnectionPrivate {
	c := (*C.GSocketConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnectionPrivate{native: c}

	return g
}

func (recv *SocketConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketControlMessageClass is a wrapper around the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native *C.GSocketControlMessageClass
	// parent_class : record
	// no type for get_size
	// no type for get_level
	// no type for get_type
	// no type for serialize
	// no type for deserialize
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func SocketControlMessageClassNewFromC(u unsafe.Pointer) *SocketControlMessageClass {
	c := (*C.GSocketControlMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessageClass{native: c}

	return g
}

func (recv *SocketControlMessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketControlMessagePrivate is a wrapper around the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native *C.GSocketControlMessagePrivate
}

func SocketControlMessagePrivateNewFromC(u unsafe.Pointer) *SocketControlMessagePrivate {
	c := (*C.GSocketControlMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketControlMessagePrivate{native: c}

	return g
}

func (recv *SocketControlMessagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketListenerClass is a wrapper around the C record GSocketListenerClass.
type SocketListenerClass struct {
	native *C.GSocketListenerClass
	// parent_class : record
	// no type for changed
	// no type for event
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketListenerClassNewFromC(u unsafe.Pointer) *SocketListenerClass {
	c := (*C.GSocketListenerClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketListenerClass{native: c}

	return g
}

func (recv *SocketListenerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketListenerPrivate is a wrapper around the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native *C.GSocketListenerPrivate
}

func SocketListenerPrivateNewFromC(u unsafe.Pointer) *SocketListenerPrivate {
	c := (*C.GSocketListenerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketListenerPrivate{native: c}

	return g
}

func (recv *SocketListenerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketPrivate is a wrapper around the C record GSocketPrivate.
type SocketPrivate struct {
	native *C.GSocketPrivate
}

func SocketPrivateNewFromC(u unsafe.Pointer) *SocketPrivate {
	c := (*C.GSocketPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketPrivate{native: c}

	return g
}

func (recv *SocketPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketServiceClass is a wrapper around the C record GSocketServiceClass.
type SocketServiceClass struct {
	native *C.GSocketServiceClass
	// parent_class : record
	// no type for incoming
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func SocketServiceClassNewFromC(u unsafe.Pointer) *SocketServiceClass {
	c := (*C.GSocketServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketServiceClass{native: c}

	return g
}

func (recv *SocketServiceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketServicePrivate is a wrapper around the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native *C.GSocketServicePrivate
}

func SocketServicePrivateNewFromC(u unsafe.Pointer) *SocketServicePrivate {
	c := (*C.GSocketServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SocketServicePrivate{native: c}

	return g
}

func (recv *SocketServicePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SrvTarget is a wrapper around the C record GSrvTarget.
type SrvTarget struct {
	native *C.GSrvTarget
}

func SrvTargetNewFromC(u unsafe.Pointer) *SrvTarget {
	c := (*C.GSrvTarget)(u)
	if c == nil {
		return nil
	}

	g := &SrvTarget{native: c}

	return g
}

func (recv *SrvTarget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StaticResource is a wrapper around the C record GStaticResource.
type StaticResource struct {
	native *C.GStaticResource
	// Private : data
	// Private : data_len
	// Private : resource
	// Private : next
	// Private : padding
}

func StaticResourceNewFromC(u unsafe.Pointer) *StaticResource {
	c := (*C.GStaticResource)(u)
	if c == nil {
		return nil
	}

	g := &StaticResource{native: c}

	return g
}

func (recv *StaticResource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TaskClass is a wrapper around the C record GTaskClass.
type TaskClass struct {
	native *C.GTaskClass
}

func TaskClassNewFromC(u unsafe.Pointer) *TaskClass {
	c := (*C.GTaskClass)(u)
	if c == nil {
		return nil
	}

	g := &TaskClass{native: c}

	return g
}

func (recv *TaskClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TcpConnectionClass is a wrapper around the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native *C.GTcpConnectionClass
	// parent_class : record
}

func TcpConnectionClassNewFromC(u unsafe.Pointer) *TcpConnectionClass {
	c := (*C.GTcpConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnectionClass{native: c}

	return g
}

func (recv *TcpConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TcpConnectionPrivate is a wrapper around the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native *C.GTcpConnectionPrivate
}

func TcpConnectionPrivateNewFromC(u unsafe.Pointer) *TcpConnectionPrivate {
	c := (*C.GTcpConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnectionPrivate{native: c}

	return g
}

func (recv *TcpConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TcpWrapperConnectionClass is a wrapper around the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native *C.GTcpWrapperConnectionClass
	// parent_class : record
}

func TcpWrapperConnectionClassNewFromC(u unsafe.Pointer) *TcpWrapperConnectionClass {
	c := (*C.GTcpWrapperConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnectionClass{native: c}

	return g
}

func (recv *TcpWrapperConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TcpWrapperConnectionPrivate is a wrapper around the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native *C.GTcpWrapperConnectionPrivate
}

func TcpWrapperConnectionPrivateNewFromC(u unsafe.Pointer) *TcpWrapperConnectionPrivate {
	c := (*C.GTcpWrapperConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TcpWrapperConnectionPrivate{native: c}

	return g
}

func (recv *TcpWrapperConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemedIconClass is a wrapper around the C record GThemedIconClass.
type ThemedIconClass struct {
	native *C.GThemedIconClass
}

func ThemedIconClassNewFromC(u unsafe.Pointer) *ThemedIconClass {
	c := (*C.GThemedIconClass)(u)
	if c == nil {
		return nil
	}

	g := &ThemedIconClass{native: c}

	return g
}

func (recv *ThemedIconClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThreadedSocketServiceClass is a wrapper around the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native *C.GThreadedSocketServiceClass
	// parent_class : record
	// no type for run
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func ThreadedSocketServiceClassNewFromC(u unsafe.Pointer) *ThreadedSocketServiceClass {
	c := (*C.GThreadedSocketServiceClass)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketServiceClass{native: c}

	return g
}

func (recv *ThreadedSocketServiceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThreadedSocketServicePrivate is a wrapper around the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native *C.GThreadedSocketServicePrivate
}

func ThreadedSocketServicePrivateNewFromC(u unsafe.Pointer) *ThreadedSocketServicePrivate {
	c := (*C.GThreadedSocketServicePrivate)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketServicePrivate{native: c}

	return g
}

func (recv *ThreadedSocketServicePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsCertificateClass is a wrapper around the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native *C.GTlsCertificateClass
	// parent_class : record
	// no type for verify
	// Private : padding
}

func TlsCertificateClassNewFromC(u unsafe.Pointer) *TlsCertificateClass {
	c := (*C.GTlsCertificateClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificateClass{native: c}

	return g
}

func (recv *TlsCertificateClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsCertificatePrivate is a wrapper around the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native *C.GTlsCertificatePrivate
}

func TlsCertificatePrivateNewFromC(u unsafe.Pointer) *TlsCertificatePrivate {
	c := (*C.GTlsCertificatePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsCertificatePrivate{native: c}

	return g
}

func (recv *TlsCertificatePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsConnectionClass is a wrapper around the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native *C.GTlsConnectionClass
	// parent_class : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// Private : padding
}

func TlsConnectionClassNewFromC(u unsafe.Pointer) *TlsConnectionClass {
	c := (*C.GTlsConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnectionClass{native: c}

	return g
}

func (recv *TlsConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsConnectionPrivate is a wrapper around the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native *C.GTlsConnectionPrivate
}

func TlsConnectionPrivateNewFromC(u unsafe.Pointer) *TlsConnectionPrivate {
	c := (*C.GTlsConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsConnectionPrivate{native: c}

	return g
}

func (recv *TlsConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsDatabasePrivate is a wrapper around the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native *C.GTlsDatabasePrivate
}

func TlsDatabasePrivateNewFromC(u unsafe.Pointer) *TlsDatabasePrivate {
	c := (*C.GTlsDatabasePrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsDatabasePrivate{native: c}

	return g
}

func (recv *TlsDatabasePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsFileDatabaseInterface is a wrapper around the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native *C.GTlsFileDatabaseInterface
	// g_iface : record
	// Private : padding
}

func TlsFileDatabaseInterfaceNewFromC(u unsafe.Pointer) *TlsFileDatabaseInterface {
	c := (*C.GTlsFileDatabaseInterface)(u)
	if c == nil {
		return nil
	}

	g := &TlsFileDatabaseInterface{native: c}

	return g
}

func (recv *TlsFileDatabaseInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsInteractionPrivate is a wrapper around the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native *C.GTlsInteractionPrivate
}

func TlsInteractionPrivateNewFromC(u unsafe.Pointer) *TlsInteractionPrivate {
	c := (*C.GTlsInteractionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsInteractionPrivate{native: c}

	return g
}

func (recv *TlsInteractionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsPasswordClass is a wrapper around the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native *C.GTlsPasswordClass
	// parent_class : record
	// no type for get_value
	// no type for set_value
	// no type for get_default_warning
	// Private : padding
}

func TlsPasswordClassNewFromC(u unsafe.Pointer) *TlsPasswordClass {
	c := (*C.GTlsPasswordClass)(u)
	if c == nil {
		return nil
	}

	g := &TlsPasswordClass{native: c}

	return g
}

func (recv *TlsPasswordClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TlsPasswordPrivate is a wrapper around the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native *C.GTlsPasswordPrivate
}

func TlsPasswordPrivateNewFromC(u unsafe.Pointer) *TlsPasswordPrivate {
	c := (*C.GTlsPasswordPrivate)(u)
	if c == nil {
		return nil
	}

	g := &TlsPasswordPrivate{native: c}

	return g
}

func (recv *TlsPasswordPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixConnectionClass is a wrapper around the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native *C.GUnixConnectionClass
	// parent_class : record
}

func UnixConnectionClassNewFromC(u unsafe.Pointer) *UnixConnectionClass {
	c := (*C.GUnixConnectionClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnectionClass{native: c}

	return g
}

func (recv *UnixConnectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixConnectionPrivate is a wrapper around the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native *C.GUnixConnectionPrivate
}

func UnixConnectionPrivateNewFromC(u unsafe.Pointer) *UnixConnectionPrivate {
	c := (*C.GUnixConnectionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixConnectionPrivate{native: c}

	return g
}

func (recv *UnixConnectionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixCredentialsMessagePrivate is a wrapper around the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native *C.GUnixCredentialsMessagePrivate
}

func UnixCredentialsMessagePrivateNewFromC(u unsafe.Pointer) *UnixCredentialsMessagePrivate {
	c := (*C.GUnixCredentialsMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixCredentialsMessagePrivate{native: c}

	return g
}

func (recv *UnixCredentialsMessagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDListClass is a wrapper around the C record GUnixFDListClass.
type UnixFDListClass struct {
	native *C.GUnixFDListClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixFDListClassNewFromC(u unsafe.Pointer) *UnixFDListClass {
	c := (*C.GUnixFDListClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDListClass{native: c}

	return g
}

func (recv *UnixFDListClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDListPrivate is a wrapper around the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native *C.GUnixFDListPrivate
}

func UnixFDListPrivateNewFromC(u unsafe.Pointer) *UnixFDListPrivate {
	c := (*C.GUnixFDListPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDListPrivate{native: c}

	return g
}

func (recv *UnixFDListPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDMessageClass is a wrapper around the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native *C.GUnixFDMessageClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
}

func UnixFDMessageClassNewFromC(u unsafe.Pointer) *UnixFDMessageClass {
	c := (*C.GUnixFDMessageClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessageClass{native: c}

	return g
}

func (recv *UnixFDMessageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDMessagePrivate is a wrapper around the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native *C.GUnixFDMessagePrivate
}

func UnixFDMessagePrivateNewFromC(u unsafe.Pointer) *UnixFDMessagePrivate {
	c := (*C.GUnixFDMessagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixFDMessagePrivate{native: c}

	return g
}

func (recv *UnixFDMessagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixInputStreamClass is a wrapper around the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native *C.GUnixInputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixInputStreamClassNewFromC(u unsafe.Pointer) *UnixInputStreamClass {
	c := (*C.GUnixInputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStreamClass{native: c}

	return g
}

func (recv *UnixInputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixInputStreamPrivate is a wrapper around the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native *C.GUnixInputStreamPrivate
}

func UnixInputStreamPrivateNewFromC(u unsafe.Pointer) *UnixInputStreamPrivate {
	c := (*C.GUnixInputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixInputStreamPrivate{native: c}

	return g
}

func (recv *UnixInputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixMountEntry is a wrapper around the C record GUnixMountEntry.
type UnixMountEntry struct {
	native *C.GUnixMountEntry
}

func UnixMountEntryNewFromC(u unsafe.Pointer) *UnixMountEntry {
	c := (*C.GUnixMountEntry)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountEntry{native: c}

	return g
}

func (recv *UnixMountEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixMountMonitorClass is a wrapper around the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native *C.GUnixMountMonitorClass
}

func UnixMountMonitorClassNewFromC(u unsafe.Pointer) *UnixMountMonitorClass {
	c := (*C.GUnixMountMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountMonitorClass{native: c}

	return g
}

func (recv *UnixMountMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixMountPoint is a wrapper around the C record GUnixMountPoint.
type UnixMountPoint struct {
	native *C.GUnixMountPoint
}

func UnixMountPointNewFromC(u unsafe.Pointer) *UnixMountPoint {
	c := (*C.GUnixMountPoint)(u)
	if c == nil {
		return nil
	}

	g := &UnixMountPoint{native: c}

	return g
}

func (recv *UnixMountPoint) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixOutputStreamClass is a wrapper around the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native *C.GUnixOutputStreamClass
	// parent_class : record
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
}

func UnixOutputStreamClassNewFromC(u unsafe.Pointer) *UnixOutputStreamClass {
	c := (*C.GUnixOutputStreamClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStreamClass{native: c}

	return g
}

func (recv *UnixOutputStreamClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixOutputStreamPrivate is a wrapper around the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native *C.GUnixOutputStreamPrivate
}

func UnixOutputStreamPrivateNewFromC(u unsafe.Pointer) *UnixOutputStreamPrivate {
	c := (*C.GUnixOutputStreamPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixOutputStreamPrivate{native: c}

	return g
}

func (recv *UnixOutputStreamPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixSocketAddressClass is a wrapper around the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native *C.GUnixSocketAddressClass
	// parent_class : record
}

func UnixSocketAddressClassNewFromC(u unsafe.Pointer) *UnixSocketAddressClass {
	c := (*C.GUnixSocketAddressClass)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddressClass{native: c}

	return g
}

func (recv *UnixSocketAddressClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixSocketAddressPrivate is a wrapper around the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native *C.GUnixSocketAddressPrivate
}

func UnixSocketAddressPrivateNewFromC(u unsafe.Pointer) *UnixSocketAddressPrivate {
	c := (*C.GUnixSocketAddressPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UnixSocketAddressPrivate{native: c}

	return g
}

func (recv *UnixSocketAddressPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VfsClass is a wrapper around the C record GVfsClass.
type VfsClass struct {
	native *C.GVfsClass
	// parent_class : record
	// no type for is_active
	// no type for get_file_for_path
	// no type for get_file_for_uri
	// no type for get_supported_uri_schemes
	// no type for parse_name
	// no type for local_file_add_info
	// no type for add_writable_namespaces
	// no type for local_file_set_attributes
	// no type for local_file_removed
	// no type for local_file_moved
	// no type for deserialize_icon
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func VfsClassNewFromC(u unsafe.Pointer) *VfsClass {
	c := (*C.GVfsClass)(u)
	if c == nil {
		return nil
	}

	g := &VfsClass{native: c}

	return g
}

func (recv *VfsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VolumeIface is a wrapper around the C record GVolumeIface.
type VolumeIface struct {
	native *C.GVolumeIface
	// g_iface : record
	// no type for changed
	// no type for removed
	// no type for get_name
	// no type for get_icon
	// no type for get_uuid
	// no type for get_drive
	// no type for get_mount
	// no type for can_mount
	// no type for can_eject
	// no type for mount_fn
	// no type for mount_finish
	// no type for eject
	// no type for eject_finish
	// no type for get_identifier
	// no type for enumerate_identifiers
	// no type for should_automount
	// no type for get_activation_root
	// no type for eject_with_operation
	// no type for eject_with_operation_finish
	// no type for get_sort_key
	// no type for get_symbolic_icon
}

func VolumeIfaceNewFromC(u unsafe.Pointer) *VolumeIface {
	c := (*C.GVolumeIface)(u)
	if c == nil {
		return nil
	}

	g := &VolumeIface{native: c}

	return g
}

func (recv *VolumeIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VolumeMonitorClass is a wrapper around the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native *C.GVolumeMonitorClass
	// parent_class : record
	// no type for volume_added
	// no type for volume_removed
	// no type for volume_changed
	// no type for mount_added
	// no type for mount_removed
	// no type for mount_pre_unmount
	// no type for mount_changed
	// no type for drive_connected
	// no type for drive_disconnected
	// no type for drive_changed
	// no type for is_supported
	// no type for get_connected_drives
	// no type for get_volumes
	// no type for get_mounts
	// no type for get_volume_for_uuid
	// no type for get_mount_for_uuid
	// no type for adopt_orphan_mount
	// no type for drive_eject_button
	// no type for drive_stop_button
	// no type for _g_reserved1
	// no type for _g_reserved2
	// no type for _g_reserved3
	// no type for _g_reserved4
	// no type for _g_reserved5
	// no type for _g_reserved6
}

func VolumeMonitorClassNewFromC(u unsafe.Pointer) *VolumeMonitorClass {
	c := (*C.GVolumeMonitorClass)(u)
	if c == nil {
		return nil
	}

	g := &VolumeMonitorClass{native: c}

	return g
}

func (recv *VolumeMonitorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ZlibCompressorClass is a wrapper around the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native *C.GZlibCompressorClass
	// parent_class : record
}

func ZlibCompressorClassNewFromC(u unsafe.Pointer) *ZlibCompressorClass {
	c := (*C.GZlibCompressorClass)(u)
	if c == nil {
		return nil
	}

	g := &ZlibCompressorClass{native: c}

	return g
}

func (recv *ZlibCompressorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ZlibDecompressorClass is a wrapper around the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native *C.GZlibDecompressorClass
	// parent_class : record
}

func ZlibDecompressorClassNewFromC(u unsafe.Pointer) *ZlibDecompressorClass {
	c := (*C.GZlibDecompressorClass)(u)
	if c == nil {
		return nil
	}

	g := &ZlibDecompressorClass{native: c}

	return g
}

func (recv *ZlibDecompressorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
