// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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
/*

	void callback_asyncreadycallbackHandler(GObject *, GObject*, GAsyncResult*, gpointer, gpointer);

*/
/*

	void callback_desktopapplaunchcallbackHandler(GObject *, GDesktopAppInfo*, GPid, gpointer, gpointer);

*/
/*

	void callback_fileprogresscallbackHandler(GObject *, goffset, goffset, gpointer, gpointer);

*/
/*

	gboolean callback_ioschedulerjobfuncHandler(GObject *, GIOSchedulerJob*, GCancellable*, gpointer, gpointer);

*/
/*

	gboolean callback_settingsbindgetmappingHandler(GObject *, GValue*, GVariant*, gpointer, gpointer);

*/
/*

	GVariant* callback_settingsbindsetmappingHandler(GObject *, const GValue*, const GVariantType*, gpointer, gpointer);

*/
/*

	void callback_simpleasyncthreadfuncHandler(GObject *, GSimpleAsyncResult*, GObject*, GCancellable*, gpointer);

*/
import "C"

var callbackAsyncreadycallbackId int
var callbackAsyncreadycallbackMap = make(map[int]AsyncreadycallbackCallback)
var callbackAsyncreadycallbackLock sync.RWMutex

// AsyncreadycallbackCallback is a callback function for a 'AsyncReadyCallback' callback.
type AsyncreadycallbackCallback func(sourceObject *gobject.Object, res *AsyncResult)

//export callback_asyncreadycallbackHandler
func callback_asyncreadycallbackHandler(_ *C.GObject, c_source_object *C.GObject, c_res *C.GAsyncResult, data C.gpointer) {
	callbackAsyncreadycallbackLock.RLock()
	defer callbackAsyncreadycallbackLock.RUnlock()

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	res := AsyncResultNewFromC(unsafe.Pointer(c_res))

	index := int(uintptr(data))
	callback := callbackAsyncreadycallbackMap[index].callback
	callback(sourceObject, res, userData)
}

var callbackDesktopapplaunchcallbackId int
var callbackDesktopapplaunchcallbackMap = make(map[int]DesktopapplaunchcallbackCallback)
var callbackDesktopapplaunchcallbackLock sync.RWMutex

// DesktopapplaunchcallbackCallback is a callback function for a 'DesktopAppLaunchCallback' callback.
type DesktopapplaunchcallbackCallback func(appinfo *DesktopAppInfo, pid glib.Pid)

//export callback_desktopapplaunchcallbackHandler
func callback_desktopapplaunchcallbackHandler(_ *C.GObject, c_appinfo *C.GDesktopAppInfo, c_pid C.GPid, data C.gpointer) {
	callbackDesktopapplaunchcallbackLock.RLock()
	defer callbackDesktopapplaunchcallbackLock.RUnlock()

	appinfo := DesktopAppInfoNewFromC(unsafe.Pointer(c_appinfo))

	pid := (c_pid)

	index := int(uintptr(data))
	callback := callbackDesktopapplaunchcallbackMap[index].callback
	callback(appinfo, pid, userData)
}

var callbackFileprogresscallbackId int
var callbackFileprogresscallbackMap = make(map[int]FileprogresscallbackCallback)
var callbackFileprogresscallbackLock sync.RWMutex

// FileprogresscallbackCallback is a callback function for a 'FileProgressCallback' callback.
type FileprogresscallbackCallback func(currentNumBytes int64, totalNumBytes int64)

//export callback_fileprogresscallbackHandler
func callback_fileprogresscallbackHandler(_ *C.GObject, c_current_num_bytes C.goffset, c_total_num_bytes C.goffset, data C.gpointer) {
	callbackFileprogresscallbackLock.RLock()
	defer callbackFileprogresscallbackLock.RUnlock()

	currentNumBytes := uint64(c_current_num_bytes)

	totalNumBytes := uint64(c_total_num_bytes)

	index := int(uintptr(data))
	callback := callbackFileprogresscallbackMap[index].callback
	callback(currentNumBytes, totalNumBytes, userData)
}

// Unsupported : callback FileReadMoreCallback : unsupported parameter callback_data : no type generator for gpointer (gpointer) for param callback_data

var callbackIoschedulerjobfuncId int
var callbackIoschedulerjobfuncMap = make(map[int]IoschedulerjobfuncCallback)
var callbackIoschedulerjobfuncLock sync.RWMutex

// IoschedulerjobfuncCallback is a callback function for a 'IOSchedulerJobFunc' callback.
type IoschedulerjobfuncCallback func(job *IOSchedulerJob, cancellable *Cancellable) bool

//export callback_ioschedulerjobfuncHandler
func callback_ioschedulerjobfuncHandler(_ *C.GObject, c_job *C.GIOSchedulerJob, c_cancellable *C.GCancellable, data C.gpointer) C.gboolean {
	callbackIoschedulerjobfuncLock.RLock()
	defer callbackIoschedulerjobfuncLock.RUnlock()

	job := IOSchedulerJobNewFromC(unsafe.Pointer(c_job))

	cancellable := CancellableNewFromC(unsafe.Pointer(c_cancellable))

	index := int(uintptr(data))
	callback := callbackIoschedulerjobfuncMap[index].callback
	retGo := callback(job, cancellable, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback ReallocFunc : no return generator

var callbackSettingsbindgetmappingId int
var callbackSettingsbindgetmappingMap = make(map[int]SettingsbindgetmappingCallback)
var callbackSettingsbindgetmappingLock sync.RWMutex

// SettingsbindgetmappingCallback is a callback function for a 'SettingsBindGetMapping' callback.
type SettingsbindgetmappingCallback func(value *gobject.Value, variant *glib.Variant) bool

//export callback_settingsbindgetmappingHandler
func callback_settingsbindgetmappingHandler(_ *C.GObject, c_value *C.GValue, c_variant *C.GVariant, data C.gpointer) C.gboolean {
	callbackSettingsbindgetmappingLock.RLock()
	defer callbackSettingsbindgetmappingLock.RUnlock()

	value := gobject.ValueNewFromC(unsafe.Pointer(c_value))

	variant := glib.VariantNewFromC(unsafe.Pointer(c_variant))

	index := int(uintptr(data))
	callback := callbackSettingsbindgetmappingMap[index].callback
	retGo := callback(value, variant, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackSettingsbindsetmappingId int
var callbackSettingsbindsetmappingMap = make(map[int]SettingsbindsetmappingCallback)
var callbackSettingsbindsetmappingLock sync.RWMutex

// SettingsbindsetmappingCallback is a callback function for a 'SettingsBindSetMapping' callback.
type SettingsbindsetmappingCallback func(value *gobject.Value, expectedType *glib.VariantType) *glib.Variant

//export callback_settingsbindsetmappingHandler
func callback_settingsbindsetmappingHandler(_ *C.GObject, c_value *C.GValue, c_expected_type *C.GVariantType, data C.gpointer) **C.GVariant {
	callbackSettingsbindsetmappingLock.RLock()
	defer callbackSettingsbindsetmappingLock.RUnlock()

	value := gobject.ValueNewFromC(unsafe.Pointer(c_value))

	expectedType := glib.VariantTypeNewFromC(unsafe.Pointer(c_expected_type))

	index := int(uintptr(data))
	callback := callbackSettingsbindsetmappingMap[index].callback
	retGo := callback(value, expectedType, userData)
	retC :=
		(*C.GVariant)(retGo.ToC())
	return retC
}

// Unsupported : callback SettingsGetMapping : unsupported parameter result : no type generator for gpointer (gpointer*) for param result

var callbackSimpleasyncthreadfuncId int
var callbackSimpleasyncthreadfuncMap = make(map[int]SimpleasyncthreadfuncCallback)
var callbackSimpleasyncthreadfuncLock sync.RWMutex

// SimpleasyncthreadfuncCallback is a callback function for a 'SimpleAsyncThreadFunc' callback.
type SimpleasyncthreadfuncCallback func(res *SimpleAsyncResult, object *gobject.Object, cancellable *Cancellable)

//export callback_simpleasyncthreadfuncHandler
func callback_simpleasyncthreadfuncHandler(_ *C.GObject, c_res *C.GSimpleAsyncResult, c_object *C.GObject, c_cancellable *C.GCancellable, data C.gpointer) {
	callbackSimpleasyncthreadfuncLock.RLock()
	defer callbackSimpleasyncthreadfuncLock.RUnlock()

	res := SimpleAsyncResultNewFromC(unsafe.Pointer(c_res))

	object := gobject.ObjectNewFromC(unsafe.Pointer(c_object))

	cancellable := CancellableNewFromC(unsafe.Pointer(c_cancellable))

	index := int(uintptr(data))
	callback := callbackSimpleasyncthreadfuncMap[index].callback
	callback(res, object, cancellable)
}
