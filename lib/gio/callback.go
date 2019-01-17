// This is a generated file - DO NOT EDIT

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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

var callbackDesktopapplaunchcallbackId int
var callbackDesktopapplaunchcallbackMap = make(map[int]DesktopapplaunchcallbackCallback)
var callbackDesktopapplaunchcallbackLock sync.RWMutex

// DesktopapplaunchcallbackCallback is a callback function for a 'DesktopAppLaunchCallback' callback.
type DesktopapplaunchcallbackCallback func(appinfo *DesktopAppInfo, pid glib.Pid)

var callbackFileprogresscallbackId int
var callbackFileprogresscallbackMap = make(map[int]FileprogresscallbackCallback)
var callbackFileprogresscallbackLock sync.RWMutex

// FileprogresscallbackCallback is a callback function for a 'FileProgressCallback' callback.
type FileprogresscallbackCallback func(currentNumBytes int64, totalNumBytes int64)

// Unsupported : callback FileReadMoreCallback : unsupported parameter callback_data : no type generator for gpointer (gpointer) for param callback_data

var callbackIoschedulerjobfuncId int
var callbackIoschedulerjobfuncMap = make(map[int]IoschedulerjobfuncCallback)
var callbackIoschedulerjobfuncLock sync.RWMutex

// IoschedulerjobfuncCallback is a callback function for a 'IOSchedulerJobFunc' callback.
type IoschedulerjobfuncCallback func(job *IOSchedulerJob, cancellable *Cancellable) bool

// Unsupported : callback ReallocFunc : no return generator

var callbackSettingsbindgetmappingId int
var callbackSettingsbindgetmappingMap = make(map[int]SettingsbindgetmappingCallback)
var callbackSettingsbindgetmappingLock sync.RWMutex

// SettingsbindgetmappingCallback is a callback function for a 'SettingsBindGetMapping' callback.
type SettingsbindgetmappingCallback func(value *gobject.Value, variant *glib.Variant) bool

var callbackSettingsbindsetmappingId int
var callbackSettingsbindsetmappingMap = make(map[int]SettingsbindsetmappingCallback)
var callbackSettingsbindsetmappingLock sync.RWMutex

// SettingsbindsetmappingCallback is a callback function for a 'SettingsBindSetMapping' callback.
type SettingsbindsetmappingCallback func(value *gobject.Value, expectedType *glib.VariantType) *glib.Variant

// Unsupported : callback SettingsGetMapping : unsupported parameter result : no type generator for gpointer (gpointer*) for param result

var callbackSimpleasyncthreadfuncId int
var callbackSimpleasyncthreadfuncMap = make(map[int]SimpleasyncthreadfuncCallback)
var callbackSimpleasyncthreadfuncLock sync.RWMutex

// SimpleasyncthreadfuncCallback is a callback function for a 'SimpleAsyncThreadFunc' callback.
type SimpleasyncthreadfuncCallback func(res *SimpleAsyncResult, object *gobject.Object, cancellable *Cancellable)
