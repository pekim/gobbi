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

	void callback_childwatchfuncHandler(GObject *, GPid, gint, gpointer, gpointer);

*/
/*

	void callback_dataforeachfuncHandler(GObject *, GQuark, gpointer, gpointer, gpointer);

*/
/*

	void callback_destroynotifyHandler(GObject *, gpointer, gpointer);

*/
/*

	void callback_freefuncHandler(GObject *, gpointer, gpointer);

*/
/*

	void callback_funcHandler(GObject *, gpointer, gpointer, gpointer);

*/
/*

	gboolean callback_hookcheckfuncHandler(GObject *, gpointer, gpointer);

*/
/*

	gint callback_hookcomparefuncHandler(GObject *, GHook*, GHook*, gpointer);

*/
/*

	void callback_hookfinalizefuncHandler(GObject *, GHookList*, GHook*, gpointer);

*/
/*

	gboolean callback_hookfindfuncHandler(GObject *, GHook*, gpointer, gpointer);

*/
/*

	void callback_hookfuncHandler(GObject *, gpointer, gpointer);

*/
/*

	gboolean callback_iofuncHandler(GObject *, GIOChannel*, GIOCondition, gpointer, gpointer);

*/
/*

	void callback_logfuncHandler(GObject *, const gchar*, GLogLevelFlags, const gchar*, gpointer, gpointer);

*/
/*

	void callback_nodeforeachfuncHandler(GObject *, GNode*, gpointer, gpointer);

*/
/*

	gboolean callback_nodetraversefuncHandler(GObject *, GNode*, gpointer, gpointer);

*/
/*

	gboolean callback_optionargfuncHandler(GObject *, const gchar*, const gchar*, gpointer, gpointer);

*/
/*

	void callback_optionerrorfuncHandler(GObject *, GOptionContext*, GOptionGroup*, gpointer, gpointer);

*/
/*

	gboolean callback_optionparsefuncHandler(GObject *, GOptionContext*, GOptionGroup*, gpointer, gpointer);

*/
/*

	gint callback_pollfuncHandler(GObject *, GPollFD*, guint, gint, gpointer);

*/
/*

	void callback_printfuncHandler(GObject *, const gchar*, gpointer);

*/
/*

	void callback_scannermsgfuncHandler(GObject *, GScanner*, gchar*, gboolean, gpointer);

*/
/*

	gint callback_sequenceitercomparefuncHandler(GObject *, GSequenceIter*, GSequenceIter*, gpointer, gpointer);

*/
/*

	void callback_sourcedummymarshalHandler(GObject *, gpointer);

*/
/*

	gboolean callback_sourcefuncHandler(GObject *, gpointer, gpointer);

*/
/*

	void callback_spawnchildsetupfuncHandler(GObject *, gpointer, gpointer);

*/
/*

	const gchar* callback_translatefuncHandler(GObject *, const gchar*, gpointer, gpointer);

*/
/*

	gboolean callback_unixfdsourcefuncHandler(GObject *, gint, GIOCondition, gpointer, gpointer);

*/
/*

	void callback_voidfuncHandler(GObject *, gpointer);

*/
import "C"

var callbackChildwatchfuncId int
var callbackChildwatchfuncMap = make(map[int]ChildwatchfuncCallback)
var callbackChildwatchfuncLock sync.RWMutex

// ChildwatchfuncCallback is a callback function for a 'ChildWatchFunc' callback.
type ChildwatchfuncCallback func(pid Pid, status int32)

//export callback_childwatchfuncHandler
func callback_childwatchfuncHandler(_ *C.GObject, c_pid C.GPid, c_status C.gint, data C.gpointer) {
	callbackChildwatchfuncLock.RLock()
	defer callbackChildwatchfuncLock.RUnlock()

	pid := (c_pid)

	status := int32(c_status)

	index := int(uintptr(data))
	callback := callbackChildwatchfuncMap[index].callback
	callback(pid, status, userData)
}

// Unsupported : callback CompareDataFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

// Unsupported : callback CompareFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

var callbackDataforeachfuncId int
var callbackDataforeachfuncMap = make(map[int]DataforeachfuncCallback)
var callbackDataforeachfuncLock sync.RWMutex

// DataforeachfuncCallback is a callback function for a 'DataForeachFunc' callback.
type DataforeachfuncCallback func(keyId Quark)

//export callback_dataforeachfuncHandler
func callback_dataforeachfuncHandler(_ *C.GObject, c_key_id C.GQuark, data C.gpointer) {
	callbackDataforeachfuncLock.RLock()
	defer callbackDataforeachfuncLock.RUnlock()

	keyId := (c_key_id)

	index := int(uintptr(data))
	callback := callbackDataforeachfuncMap[index].callback
	callback(keyId, data, userData)
}

var callbackDestroynotifyId int
var callbackDestroynotifyMap = make(map[int]DestroynotifyCallback)
var callbackDestroynotifyLock sync.RWMutex

// DestroynotifyCallback is a callback function for a 'DestroyNotify' callback.
type DestroynotifyCallback func()

//export callback_destroynotifyHandler
func callback_destroynotifyHandler(_ *C.GObject, data C.gpointer) {
	callbackDestroynotifyLock.RLock()
	defer callbackDestroynotifyLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackDestroynotifyMap[index].callback
	callback(data)
}

// Unsupported : callback DuplicateFunc : no return generator

// Unsupported : callback EqualFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

var callbackFreefuncId int
var callbackFreefuncMap = make(map[int]FreefuncCallback)
var callbackFreefuncLock sync.RWMutex

// FreefuncCallback is a callback function for a 'FreeFunc' callback.
type FreefuncCallback func()

//export callback_freefuncHandler
func callback_freefuncHandler(_ *C.GObject, data C.gpointer) {
	callbackFreefuncLock.RLock()
	defer callbackFreefuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackFreefuncMap[index].callback
	callback(data)
}

var callbackFuncId int
var callbackFuncMap = make(map[int]FuncCallback)
var callbackFuncLock sync.RWMutex

// FuncCallback is a callback function for a 'Func' callback.
type FuncCallback func()

//export callback_funcHandler
func callback_funcHandler(_ *C.GObject, data C.gpointer) {
	callbackFuncLock.RLock()
	defer callbackFuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackFuncMap[index].callback
	callback(data, userData)
}

// Unsupported : callback HFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : callback HRFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : callback HashFunc : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

var callbackHookcheckfuncId int
var callbackHookcheckfuncMap = make(map[int]HookcheckfuncCallback)
var callbackHookcheckfuncLock sync.RWMutex

// HookcheckfuncCallback is a callback function for a 'HookCheckFunc' callback.
type HookcheckfuncCallback func() bool

//export callback_hookcheckfuncHandler
func callback_hookcheckfuncHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	callbackHookcheckfuncLock.RLock()
	defer callbackHookcheckfuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackHookcheckfuncMap[index].callback
	retGo := callback(data)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback HookCheckMarshaller : unsupported parameter marshal_data : no type generator for gpointer (gpointer) for param marshal_data

var callbackHookcomparefuncId int
var callbackHookcomparefuncMap = make(map[int]HookcomparefuncCallback)
var callbackHookcomparefuncLock sync.RWMutex

// HookcomparefuncCallback is a callback function for a 'HookCompareFunc' callback.
type HookcomparefuncCallback func(newHook *Hook, sibling *Hook) int32

//export callback_hookcomparefuncHandler
func callback_hookcomparefuncHandler(_ *C.GObject, c_new_hook *C.GHook, c_sibling *C.GHook, data C.gpointer) C.gint {
	callbackHookcomparefuncLock.RLock()
	defer callbackHookcomparefuncLock.RUnlock()

	newHook := HookNewFromC(unsafe.Pointer(c_new_hook))

	sibling := HookNewFromC(unsafe.Pointer(c_sibling))

	index := int(uintptr(data))
	callback := callbackHookcomparefuncMap[index].callback
	retGo := callback(newHook, sibling)
	retC :=
		(C.gint)(retGo)
	return retC
}

var callbackHookfinalizefuncId int
var callbackHookfinalizefuncMap = make(map[int]HookfinalizefuncCallback)
var callbackHookfinalizefuncLock sync.RWMutex

// HookfinalizefuncCallback is a callback function for a 'HookFinalizeFunc' callback.
type HookfinalizefuncCallback func(hookList *HookList, hook *Hook)

//export callback_hookfinalizefuncHandler
func callback_hookfinalizefuncHandler(_ *C.GObject, c_hook_list *C.GHookList, c_hook *C.GHook, data C.gpointer) {
	callbackHookfinalizefuncLock.RLock()
	defer callbackHookfinalizefuncLock.RUnlock()

	hookList := HookListNewFromC(unsafe.Pointer(c_hook_list))

	hook := HookNewFromC(unsafe.Pointer(c_hook))

	index := int(uintptr(data))
	callback := callbackHookfinalizefuncMap[index].callback
	callback(hookList, hook)
}

var callbackHookfindfuncId int
var callbackHookfindfuncMap = make(map[int]HookfindfuncCallback)
var callbackHookfindfuncLock sync.RWMutex

// HookfindfuncCallback is a callback function for a 'HookFindFunc' callback.
type HookfindfuncCallback func(hook *Hook) bool

//export callback_hookfindfuncHandler
func callback_hookfindfuncHandler(_ *C.GObject, c_hook *C.GHook, data C.gpointer) C.gboolean {
	callbackHookfindfuncLock.RLock()
	defer callbackHookfindfuncLock.RUnlock()

	hook := HookNewFromC(unsafe.Pointer(c_hook))

	index := int(uintptr(data))
	callback := callbackHookfindfuncMap[index].callback
	retGo := callback(hook, data)
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
func callback_hookfuncHandler(_ *C.GObject, data C.gpointer) {
	callbackHookfuncLock.RLock()
	defer callbackHookfuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackHookfuncMap[index].callback
	callback(data)
}

// Unsupported : callback HookMarshaller : unsupported parameter marshal_data : no type generator for gpointer (gpointer) for param marshal_data

var callbackIofuncId int
var callbackIofuncMap = make(map[int]IofuncCallback)
var callbackIofuncLock sync.RWMutex

// IofuncCallback is a callback function for a 'IOFunc' callback.
type IofuncCallback func(source *IOChannel, condition IOCondition) bool

//export callback_iofuncHandler
func callback_iofuncHandler(_ *C.GObject, c_source *C.GIOChannel, c_condition C.GIOCondition, data C.gpointer) C.gboolean {
	callbackIofuncLock.RLock()
	defer callbackIofuncLock.RUnlock()

	source := IOChannelNewFromC(unsafe.Pointer(c_source))

	condition := IOCondition(c_condition)

	index := int(uintptr(data))
	callback := callbackIofuncMap[index].callback
	retGo := callback(source, condition, data)
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
func callback_logfuncHandler(_ *C.GObject, c_log_domain *C.gchar, c_log_level C.GLogLevelFlags, c_message *C.gchar, data C.gpointer) {
	callbackLogfuncLock.RLock()
	defer callbackLogfuncLock.RUnlock()

	logDomain := C.GoString(c_log_domain)

	logLevel := LogLevelFlags(c_log_level)

	message := C.GoString(c_message)

	index := int(uintptr(data))
	callback := callbackLogfuncMap[index].callback
	callback(logDomain, logLevel, message, userData)
}

var callbackNodeforeachfuncId int
var callbackNodeforeachfuncMap = make(map[int]NodeforeachfuncCallback)
var callbackNodeforeachfuncLock sync.RWMutex

// NodeforeachfuncCallback is a callback function for a 'NodeForeachFunc' callback.
type NodeforeachfuncCallback func(node *Node)

//export callback_nodeforeachfuncHandler
func callback_nodeforeachfuncHandler(_ *C.GObject, c_node *C.GNode, data C.gpointer) {
	callbackNodeforeachfuncLock.RLock()
	defer callbackNodeforeachfuncLock.RUnlock()

	node := NodeNewFromC(unsafe.Pointer(c_node))

	index := int(uintptr(data))
	callback := callbackNodeforeachfuncMap[index].callback
	callback(node, data)
}

var callbackNodetraversefuncId int
var callbackNodetraversefuncMap = make(map[int]NodetraversefuncCallback)
var callbackNodetraversefuncLock sync.RWMutex

// NodetraversefuncCallback is a callback function for a 'NodeTraverseFunc' callback.
type NodetraversefuncCallback func(node *Node) bool

//export callback_nodetraversefuncHandler
func callback_nodetraversefuncHandler(_ *C.GObject, c_node *C.GNode, data C.gpointer) C.gboolean {
	callbackNodetraversefuncLock.RLock()
	defer callbackNodetraversefuncLock.RUnlock()

	node := NodeNewFromC(unsafe.Pointer(c_node))

	index := int(uintptr(data))
	callback := callbackNodetraversefuncMap[index].callback
	retGo := callback(node, data)
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
func callback_optionargfuncHandler(_ *C.GObject, c_option_name *C.gchar, c_value *C.gchar, data C.gpointer) C.gboolean {
	callbackOptionargfuncLock.RLock()
	defer callbackOptionargfuncLock.RUnlock()

	optionName := C.GoString(c_option_name)

	value := C.GoString(c_value)

	index := int(uintptr(data))
	callback := callbackOptionargfuncMap[index].callback
	retGo := callback(optionName, value, data)
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
func callback_optionerrorfuncHandler(_ *C.GObject, c_context *C.GOptionContext, c_group *C.GOptionGroup, data C.gpointer) {
	callbackOptionerrorfuncLock.RLock()
	defer callbackOptionerrorfuncLock.RUnlock()

	context := OptionContextNewFromC(unsafe.Pointer(c_context))

	group := OptionGroupNewFromC(unsafe.Pointer(c_group))

	index := int(uintptr(data))
	callback := callbackOptionerrorfuncMap[index].callback
	callback(context, group, data)
}

var callbackOptionparsefuncId int
var callbackOptionparsefuncMap = make(map[int]OptionparsefuncCallback)
var callbackOptionparsefuncLock sync.RWMutex

// OptionparsefuncCallback is a callback function for a 'OptionParseFunc' callback.
type OptionparsefuncCallback func(context *OptionContext, group *OptionGroup) bool

//export callback_optionparsefuncHandler
func callback_optionparsefuncHandler(_ *C.GObject, c_context *C.GOptionContext, c_group *C.GOptionGroup, data C.gpointer) C.gboolean {
	callbackOptionparsefuncLock.RLock()
	defer callbackOptionparsefuncLock.RUnlock()

	context := OptionContextNewFromC(unsafe.Pointer(c_context))

	group := OptionGroupNewFromC(unsafe.Pointer(c_group))

	index := int(uintptr(data))
	callback := callbackOptionparsefuncMap[index].callback
	retGo := callback(context, group, data)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackPollfuncId int
var callbackPollfuncMap = make(map[int]PollfuncCallback)
var callbackPollfuncLock sync.RWMutex

// PollfuncCallback is a callback function for a 'PollFunc' callback.
type PollfuncCallback func(ufds *PollFD, nfsd uint32, timeout int32) int32

//export callback_pollfuncHandler
func callback_pollfuncHandler(_ *C.GObject, c_ufds *C.GPollFD, c_nfsd C.guint, c_timeout_ C.gint, data C.gpointer) C.gint {
	callbackPollfuncLock.RLock()
	defer callbackPollfuncLock.RUnlock()

	ufds := PollFDNewFromC(unsafe.Pointer(c_ufds))

	nfsd := uint32(c_nfsd)

	timeout := int32(c_timeout_)

	index := int(uintptr(data))
	callback := callbackPollfuncMap[index].callback
	retGo := callback(ufds, nfsd, timeout)
	retC :=
		(C.gint)(retGo)
	return retC
}

var callbackPrintfuncId int
var callbackPrintfuncMap = make(map[int]PrintfuncCallback)
var callbackPrintfuncLock sync.RWMutex

// PrintfuncCallback is a callback function for a 'PrintFunc' callback.
type PrintfuncCallback func(string_ string)

//export callback_printfuncHandler
func callback_printfuncHandler(_ *C.GObject, c_string *C.gchar, data C.gpointer) {
	callbackPrintfuncLock.RLock()
	defer callbackPrintfuncLock.RUnlock()

	string_ := C.GoString(c_string)

	index := int(uintptr(data))
	callback := callbackPrintfuncMap[index].callback
	callback(string_)
}

var callbackScannermsgfuncId int
var callbackScannermsgfuncMap = make(map[int]ScannermsgfuncCallback)
var callbackScannermsgfuncLock sync.RWMutex

// ScannermsgfuncCallback is a callback function for a 'ScannerMsgFunc' callback.
type ScannermsgfuncCallback func(scanner *Scanner, message string, error bool)

//export callback_scannermsgfuncHandler
func callback_scannermsgfuncHandler(_ *C.GObject, c_scanner *C.GScanner, c_message *C.gchar, c_error C.gboolean, data C.gpointer) {
	callbackScannermsgfuncLock.RLock()
	defer callbackScannermsgfuncLock.RUnlock()

	scanner := ScannerNewFromC(unsafe.Pointer(c_scanner))

	message := C.GoString(c_message)

	error := c_error == C.TRUE

	index := int(uintptr(data))
	callback := callbackScannermsgfuncMap[index].callback
	callback(scanner, message, error)
}

var callbackSequenceitercomparefuncId int
var callbackSequenceitercomparefuncMap = make(map[int]SequenceitercomparefuncCallback)
var callbackSequenceitercomparefuncLock sync.RWMutex

// SequenceitercomparefuncCallback is a callback function for a 'SequenceIterCompareFunc' callback.
type SequenceitercomparefuncCallback func(a *SequenceIter, b *SequenceIter) int32

//export callback_sequenceitercomparefuncHandler
func callback_sequenceitercomparefuncHandler(_ *C.GObject, c_a *C.GSequenceIter, c_b *C.GSequenceIter, data C.gpointer) C.gint {
	callbackSequenceitercomparefuncLock.RLock()
	defer callbackSequenceitercomparefuncLock.RUnlock()

	a := SequenceIterNewFromC(unsafe.Pointer(c_a))

	b := SequenceIterNewFromC(unsafe.Pointer(c_b))

	index := int(uintptr(data))
	callback := callbackSequenceitercomparefuncMap[index].callback
	retGo := callback(a, b, data)
	retC :=
		(C.gint)(retGo)
	return retC
}

var callbackSourcedummymarshalId int
var callbackSourcedummymarshalMap = make(map[int]SourcedummymarshalCallback)
var callbackSourcedummymarshalLock sync.RWMutex

// SourcedummymarshalCallback is a callback function for a 'SourceDummyMarshal' callback.
type SourcedummymarshalCallback func()

//export callback_sourcedummymarshalHandler
func callback_sourcedummymarshalHandler(_ *C.GObject, data C.gpointer) {
	callbackSourcedummymarshalLock.RLock()
	defer callbackSourcedummymarshalLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackSourcedummymarshalMap[index].callback
	callback()
}

var callbackSourcefuncId int
var callbackSourcefuncMap = make(map[int]SourcefuncCallback)
var callbackSourcefuncLock sync.RWMutex

// SourcefuncCallback is a callback function for a 'SourceFunc' callback.
type SourcefuncCallback func() bool

//export callback_sourcefuncHandler
func callback_sourcefuncHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	callbackSourcefuncLock.RLock()
	defer callbackSourcefuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackSourcefuncMap[index].callback
	retGo := callback(userData)
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
func callback_spawnchildsetupfuncHandler(_ *C.GObject, data C.gpointer) {
	callbackSpawnchildsetupfuncLock.RLock()
	defer callbackSpawnchildsetupfuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackSpawnchildsetupfuncMap[index].callback
	callback(userData)
}

// Unsupported : callback ThreadFunc : no return generator

var callbackTranslatefuncId int
var callbackTranslatefuncMap = make(map[int]TranslatefuncCallback)
var callbackTranslatefuncLock sync.RWMutex

// TranslatefuncCallback is a callback function for a 'TranslateFunc' callback.
type TranslatefuncCallback func(str string) string

//export callback_translatefuncHandler
func callback_translatefuncHandler(_ *C.GObject, c_str *C.gchar, data C.gpointer) {
	callbackTranslatefuncLock.RLock()
	defer callbackTranslatefuncLock.RUnlock()

	str := C.GoString(c_str)

	index := int(uintptr(data))
	callback := callbackTranslatefuncMap[index].callback
	retGo := callback(str, data)
	retC :=
		C.CString(retGo)
	return retC
}

// Unsupported : callback TraverseFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

var callbackUnixfdsourcefuncId int
var callbackUnixfdsourcefuncMap = make(map[int]UnixfdsourcefuncCallback)
var callbackUnixfdsourcefuncLock sync.RWMutex

// UnixfdsourcefuncCallback is a callback function for a 'UnixFDSourceFunc' callback.
type UnixfdsourcefuncCallback func(fd int32, condition IOCondition) bool

//export callback_unixfdsourcefuncHandler
func callback_unixfdsourcefuncHandler(_ *C.GObject, c_fd C.gint, c_condition C.GIOCondition, data C.gpointer) C.gboolean {
	callbackUnixfdsourcefuncLock.RLock()
	defer callbackUnixfdsourcefuncLock.RUnlock()

	fd := int32(c_fd)

	condition := IOCondition(c_condition)

	index := int(uintptr(data))
	callback := callbackUnixfdsourcefuncMap[index].callback
	retGo := callback(fd, condition, userData)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackVoidfuncId int
var callbackVoidfuncMap = make(map[int]VoidfuncCallback)
var callbackVoidfuncLock sync.RWMutex

// VoidfuncCallback is a callback function for a 'VoidFunc' callback.
type VoidfuncCallback func()

//export callback_voidfuncHandler
func callback_voidfuncHandler(_ *C.GObject, data C.gpointer) {
	callbackVoidfuncLock.RLock()
	defer callbackVoidfuncLock.RUnlock()

	index := int(uintptr(data))
	callback := callbackVoidfuncMap[index].callback
	callback()
}
