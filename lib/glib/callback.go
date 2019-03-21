// This is a generated file - DO NOT EDIT

package glib

import (
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <glib-object.h>
// #include <stdlib.h>
/*

	void callback_childwatchfuncHandler(GObject *, GPid, gint, gpointer);

*/
/*

	void callback_destroynotifyHandler(GObject *, gpointer);

*/
/*

	void callback_freefuncHandler(GObject *, gpointer);

*/
/*

	gboolean callback_hookcheckfuncHandler(GObject *, gpointer);

*/
/*

	gboolean callback_hookfindfuncHandler(GObject *, GHook*, gpointer);

*/
/*

	void callback_hookfuncHandler(GObject *, gpointer);

*/
/*

	gboolean callback_iofuncHandler(GObject *, GIOChannel*, GIOCondition, gpointer);

*/
/*

	void callback_logfuncHandler(GObject *, gchar*, GLogLevelFlags, gchar*, gpointer);

*/
/*

	void callback_nodeforeachfuncHandler(GObject *, GNode*, gpointer);

*/
/*

	gboolean callback_nodetraversefuncHandler(GObject *, GNode*, gpointer);

*/
/*

	gboolean callback_optionargfuncHandler(GObject *, gchar*, gchar*, gpointer);

*/
/*

	void callback_optionerrorfuncHandler(GObject *, GOptionContext*, GOptionGroup*, gpointer);

*/
/*

	gboolean callback_optionparsefuncHandler(GObject *, GOptionContext*, GOptionGroup*, gpointer);

*/
/*

	gint callback_sequenceitercomparefuncHandler(GObject *, GSequenceIter*, GSequenceIter*, gpointer);

*/
/*

	gboolean callback_sourcefuncHandler(GObject *, gpointer);

*/
/*

	void callback_spawnchildsetupfuncHandler(GObject *, gpointer);

*/
/*

//	const gchar* callback_translatefuncHandler(GObject *, gchar*, gpointer);

*/
/*

	gboolean callback_unixfdsourcefuncHandler(GObject *, gint, GIOCondition, gpointer);

*/
import "C"

var callbackChildwatchfuncId int
var callbackChildwatchfuncMap = make(map[int]ChildwatchfuncCallback)
var callbackChildwatchfuncLock sync.RWMutex

// ChildwatchfuncCallback is a callback function for a 'ChildWatchFunc' callback.
type ChildwatchfuncCallback func(pid Pid, status int32)

//export callback_childwatchfuncHandler
func callback_childwatchfuncHandler(_ *C.GObject, c_pid C.GPid, c_status C.gint, c_user_data C.gpointer) {
	callbackChildwatchfuncLock.RLock()
	defer callbackChildwatchfuncLock.RUnlock()

	pid := Pid(c_pid)

	status := int32(c_status)

	index := int(uintptr(c_user_data))
	callback := callbackChildwatchfuncMap[index]
	callback(pid, status)
}

// Unsupported : callback CompareDataFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

// Unsupported : callback CompareFunc : no [user_]data param

// Unsupported : callback DataForeachFunc : unsupported parameter data : no type generator for gpointer (gpointer) for param data

var callbackDestroynotifyId int
var callbackDestroynotifyMap = make(map[int]DestroynotifyCallback)
var callbackDestroynotifyLock sync.RWMutex

// DestroynotifyCallback is a callback function for a 'DestroyNotify' callback.
type DestroynotifyCallback func()

//export callback_destroynotifyHandler
func callback_destroynotifyHandler(_ *C.GObject, c_data C.gpointer) {
	callbackDestroynotifyLock.RLock()
	defer callbackDestroynotifyLock.RUnlock()

	index := int(uintptr(c_data))
	callback := callbackDestroynotifyMap[index]
	callback()
}

// Unsupported : callback DuplicateFunc : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : callback EqualFunc : no [user_]data param

var callbackFreefuncId int
var callbackFreefuncMap = make(map[int]FreefuncCallback)
var callbackFreefuncLock sync.RWMutex

// FreefuncCallback is a callback function for a 'FreeFunc' callback.
type FreefuncCallback func()

//export callback_freefuncHandler
func callback_freefuncHandler(_ *C.GObject, c_data C.gpointer) {
	callbackFreefuncLock.RLock()
	defer callbackFreefuncLock.RUnlock()

	index := int(uintptr(c_data))
	callback := callbackFreefuncMap[index]
	callback()
}

// Unsupported : callback Func : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : callback HFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : callback HRFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : callback HashFunc : no [user_]data param

var callbackHookcheckfuncId int
var callbackHookcheckfuncMap = make(map[int]HookcheckfuncCallback)
var callbackHookcheckfuncLock sync.RWMutex

// HookcheckfuncCallback is a callback function for a 'HookCheckFunc' callback.
type HookcheckfuncCallback func() bool

//export callback_hookcheckfuncHandler
func callback_hookcheckfuncHandler(_ *C.GObject, c_data C.gpointer) C.gboolean {
	callbackHookcheckfuncLock.RLock()
	defer callbackHookcheckfuncLock.RUnlock()

	index := int(uintptr(c_data))
	callback := callbackHookcheckfuncMap[index]
	retGo := callback()
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback HookCheckMarshaller : no [user_]data param

// Unsupported : callback HookCompareFunc : no [user_]data param

// Unsupported : callback HookFinalizeFunc : no [user_]data param

var callbackHookfindfuncId int
var callbackHookfindfuncMap = make(map[int]HookfindfuncCallback)
var callbackHookfindfuncLock sync.RWMutex

// HookfindfuncCallback is a callback function for a 'HookFindFunc' callback.
type HookfindfuncCallback func(hook *Hook) bool

//export callback_hookfindfuncHandler
func callback_hookfindfuncHandler(_ *C.GObject, c_hook *C.GHook, c_data C.gpointer) C.gboolean {
	callbackHookfindfuncLock.RLock()
	defer callbackHookfindfuncLock.RUnlock()

	hook := HookNewFromC(unsafe.Pointer(c_hook))

	index := int(uintptr(c_data))
	callback := callbackHookfindfuncMap[index]
	retGo := callback(hook)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackHookfuncId int
var callbackHookfuncMap = make(map[int]HookfuncCallback)
var callbackHookfuncLock sync.RWMutex

// HookfuncCallback is a callback function for a 'HookFunc' callback.
type HookfuncCallback func()

//export callback_hookfuncHandler
func callback_hookfuncHandler(_ *C.GObject, c_data C.gpointer) {
	callbackHookfuncLock.RLock()
	defer callbackHookfuncLock.RUnlock()

	index := int(uintptr(c_data))
	callback := callbackHookfuncMap[index]
	callback()
}

// Unsupported : callback HookMarshaller : no [user_]data param

var callbackIofuncId int
var callbackIofuncMap = make(map[int]IofuncCallback)
var callbackIofuncLock sync.RWMutex

// IofuncCallback is a callback function for a 'IOFunc' callback.
type IofuncCallback func(source *IOChannel, condition IOCondition) bool

//export callback_iofuncHandler
func callback_iofuncHandler(_ *C.GObject, c_source *C.GIOChannel, c_condition C.GIOCondition, c_data C.gpointer) C.gboolean {
	callbackIofuncLock.RLock()
	defer callbackIofuncLock.RUnlock()

	source := IOChannelNewFromC(unsafe.Pointer(c_source))

	condition := IOCondition(c_condition)

	index := int(uintptr(c_data))
	callback := callbackIofuncMap[index]
	retGo := callback(source, condition)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackLogfuncId int
var callbackLogfuncMap = make(map[int]LogfuncCallback)
var callbackLogfuncLock sync.RWMutex

// LogfuncCallback is a callback function for a 'LogFunc' callback.
type LogfuncCallback func(logDomain string, logLevel LogLevelFlags, message string)

//export callback_logfuncHandler
func callback_logfuncHandler(_ *C.GObject, c_log_domain *C.gchar, c_log_level C.GLogLevelFlags, c_message *C.gchar, c_user_data C.gpointer) {
	callbackLogfuncLock.RLock()
	defer callbackLogfuncLock.RUnlock()

	logDomain := C.GoString(c_log_domain)

	logLevel := LogLevelFlags(c_log_level)

	message := C.GoString(c_message)

	index := int(uintptr(c_user_data))
	callback := callbackLogfuncMap[index]
	callback(logDomain, logLevel, message)
}

var callbackNodeforeachfuncId int
var callbackNodeforeachfuncMap = make(map[int]NodeforeachfuncCallback)
var callbackNodeforeachfuncLock sync.RWMutex

// NodeforeachfuncCallback is a callback function for a 'NodeForeachFunc' callback.
type NodeforeachfuncCallback func(node *Node)

//export callback_nodeforeachfuncHandler
func callback_nodeforeachfuncHandler(_ *C.GObject, c_node *C.GNode, c_data C.gpointer) {
	callbackNodeforeachfuncLock.RLock()
	defer callbackNodeforeachfuncLock.RUnlock()

	node := NodeNewFromC(unsafe.Pointer(c_node))

	index := int(uintptr(c_data))
	callback := callbackNodeforeachfuncMap[index]
	callback(node)
}

var callbackNodetraversefuncId int
var callbackNodetraversefuncMap = make(map[int]NodetraversefuncCallback)
var callbackNodetraversefuncLock sync.RWMutex

// NodetraversefuncCallback is a callback function for a 'NodeTraverseFunc' callback.
type NodetraversefuncCallback func(node *Node) bool

//export callback_nodetraversefuncHandler
func callback_nodetraversefuncHandler(_ *C.GObject, c_node *C.GNode, c_data C.gpointer) C.gboolean {
	callbackNodetraversefuncLock.RLock()
	defer callbackNodetraversefuncLock.RUnlock()

	node := NodeNewFromC(unsafe.Pointer(c_node))

	index := int(uintptr(c_data))
	callback := callbackNodetraversefuncMap[index]
	retGo := callback(node)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackOptionargfuncId int
var callbackOptionargfuncMap = make(map[int]OptionargfuncCallback)
var callbackOptionargfuncLock sync.RWMutex

// OptionargfuncCallback is a callback function for a 'OptionArgFunc' callback.
type OptionargfuncCallback func(optionName string, value string) bool

//export callback_optionargfuncHandler
func callback_optionargfuncHandler(_ *C.GObject, c_option_name *C.gchar, c_value *C.gchar, c_data C.gpointer) C.gboolean {
	callbackOptionargfuncLock.RLock()
	defer callbackOptionargfuncLock.RUnlock()

	optionName := C.GoString(c_option_name)

	value := C.GoString(c_value)

	index := int(uintptr(c_data))
	callback := callbackOptionargfuncMap[index]
	retGo := callback(optionName, value)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackOptionerrorfuncId int
var callbackOptionerrorfuncMap = make(map[int]OptionerrorfuncCallback)
var callbackOptionerrorfuncLock sync.RWMutex

// OptionerrorfuncCallback is a callback function for a 'OptionErrorFunc' callback.
type OptionerrorfuncCallback func(context *OptionContext, group *OptionGroup)

//export callback_optionerrorfuncHandler
func callback_optionerrorfuncHandler(_ *C.GObject, c_context *C.GOptionContext, c_group *C.GOptionGroup, c_data C.gpointer) {
	callbackOptionerrorfuncLock.RLock()
	defer callbackOptionerrorfuncLock.RUnlock()

	context := OptionContextNewFromC(unsafe.Pointer(c_context))

	group := OptionGroupNewFromC(unsafe.Pointer(c_group))

	index := int(uintptr(c_data))
	callback := callbackOptionerrorfuncMap[index]
	callback(context, group)
}

var callbackOptionparsefuncId int
var callbackOptionparsefuncMap = make(map[int]OptionparsefuncCallback)
var callbackOptionparsefuncLock sync.RWMutex

// OptionparsefuncCallback is a callback function for a 'OptionParseFunc' callback.
type OptionparsefuncCallback func(context *OptionContext, group *OptionGroup) bool

//export callback_optionparsefuncHandler
func callback_optionparsefuncHandler(_ *C.GObject, c_context *C.GOptionContext, c_group *C.GOptionGroup, c_data C.gpointer) C.gboolean {
	callbackOptionparsefuncLock.RLock()
	defer callbackOptionparsefuncLock.RUnlock()

	context := OptionContextNewFromC(unsafe.Pointer(c_context))

	group := OptionGroupNewFromC(unsafe.Pointer(c_group))

	index := int(uintptr(c_data))
	callback := callbackOptionparsefuncMap[index]
	retGo := callback(context, group)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback PollFunc : no [user_]data param

// Unsupported : callback PrintFunc : no [user_]data param

// Unsupported : callback ScannerMsgFunc : no [user_]data param

var callbackSequenceitercomparefuncId int
var callbackSequenceitercomparefuncMap = make(map[int]SequenceitercomparefuncCallback)
var callbackSequenceitercomparefuncLock sync.RWMutex

// SequenceitercomparefuncCallback is a callback function for a 'SequenceIterCompareFunc' callback.
type SequenceitercomparefuncCallback func(a *SequenceIter, b *SequenceIter) int32

//export callback_sequenceitercomparefuncHandler
func callback_sequenceitercomparefuncHandler(_ *C.GObject, c_a *C.GSequenceIter, c_b *C.GSequenceIter, c_data C.gpointer) C.gint {
	callbackSequenceitercomparefuncLock.RLock()
	defer callbackSequenceitercomparefuncLock.RUnlock()

	a := SequenceIterNewFromC(unsafe.Pointer(c_a))

	b := SequenceIterNewFromC(unsafe.Pointer(c_b))

	index := int(uintptr(c_data))
	callback := callbackSequenceitercomparefuncMap[index]
	retGo := callback(a, b)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported : callback SourceDummyMarshal : no [user_]data param

var callbackSourcefuncId int
var callbackSourcefuncMap = make(map[int]SourcefuncCallback)
var callbackSourcefuncLock sync.RWMutex

// SourcefuncCallback is a callback function for a 'SourceFunc' callback.
type SourcefuncCallback func() bool

//export callback_sourcefuncHandler
func callback_sourcefuncHandler(_ *C.GObject, c_user_data C.gpointer) C.gboolean {
	callbackSourcefuncLock.RLock()
	defer callbackSourcefuncLock.RUnlock()

	index := int(uintptr(c_user_data))
	callback := callbackSourcefuncMap[index]
	retGo := callback()
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackSpawnchildsetupfuncId int
var callbackSpawnchildsetupfuncMap = make(map[int]SpawnchildsetupfuncCallback)
var callbackSpawnchildsetupfuncLock sync.RWMutex

// SpawnchildsetupfuncCallback is a callback function for a 'SpawnChildSetupFunc' callback.
type SpawnchildsetupfuncCallback func()

//export callback_spawnchildsetupfuncHandler
func callback_spawnchildsetupfuncHandler(_ *C.GObject, c_user_data C.gpointer) {
	callbackSpawnchildsetupfuncLock.RLock()
	defer callbackSpawnchildsetupfuncLock.RUnlock()

	index := int(uintptr(c_user_data))
	callback := callbackSpawnchildsetupfuncMap[index]
	callback()
}

// Unsupported : callback ThreadFunc : no return generator

var callbackTranslatefuncId int
var callbackTranslatefuncMap = make(map[int]TranslatefuncCallback)
var callbackTranslatefuncLock sync.RWMutex

// TranslatefuncCallback is a callback function for a 'TranslateFunc' callback.
type TranslatefuncCallback func(str string) string

////export callback_translatefuncHandler
//func callback_translatefuncHandler(_ *C.GObject, c_str *C.gchar, c_data C.gpointer) *C.gchar {
//	callbackTranslatefuncLock.RLock()
//	defer callbackTranslatefuncLock.RUnlock()
//
//	str := C.GoString(c_str)
//
//	index := int(uintptr(c_data))
//	callback := callbackTranslatefuncMap[index]
//	retGo := callback(str)
//	retC :=
//		C.CString(retGo)
//	return retC
//}

// Unsupported : callback TraverseFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

var callbackUnixfdsourcefuncId int
var callbackUnixfdsourcefuncMap = make(map[int]UnixfdsourcefuncCallback)
var callbackUnixfdsourcefuncLock sync.RWMutex

// UnixfdsourcefuncCallback is a callback function for a 'UnixFDSourceFunc' callback.
type UnixfdsourcefuncCallback func(fd int32, condition IOCondition) bool

//export callback_unixfdsourcefuncHandler
func callback_unixfdsourcefuncHandler(_ *C.GObject, c_fd C.gint, c_condition C.GIOCondition, c_user_data C.gpointer) C.gboolean {
	callbackUnixfdsourcefuncLock.RLock()
	defer callbackUnixfdsourcefuncLock.RUnlock()

	fd := int32(c_fd)

	condition := IOCondition(c_condition)

	index := int(uintptr(c_user_data))
	callback := callbackUnixfdsourcefuncMap[index]
	retGo := callback(fd, condition)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback VoidFunc : no [user_]data param
