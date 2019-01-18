// This is a generated file - DO NOT EDIT

package glib

import "sync"

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

// Unsupported : callback CompareDataFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

// Unsupported : callback CompareFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

var callbackDataforeachfuncId int
var callbackDataforeachfuncMap = make(map[int]DataforeachfuncCallback)
var callbackDataforeachfuncLock sync.RWMutex

// DataforeachfuncCallback is a callback function for a 'DataForeachFunc' callback.
type DataforeachfuncCallback func(keyId Quark)

var callbackDestroynotifyId int
var callbackDestroynotifyMap = make(map[int]DestroynotifyCallback)
var callbackDestroynotifyLock sync.RWMutex

// DestroynotifyCallback is a callback function for a 'DestroyNotify' callback.
type DestroynotifyCallback func()

// Unsupported : callback DuplicateFunc : no return generator

// Unsupported : callback EqualFunc : unsupported parameter a : no type generator for gpointer (gconstpointer) for param a

var callbackFreefuncId int
var callbackFreefuncMap = make(map[int]FreefuncCallback)
var callbackFreefuncLock sync.RWMutex

// FreefuncCallback is a callback function for a 'FreeFunc' callback.
type FreefuncCallback func()

var callbackFuncId int
var callbackFuncMap = make(map[int]FuncCallback)
var callbackFuncLock sync.RWMutex

// FuncCallback is a callback function for a 'Func' callback.
type FuncCallback func()

// Unsupported : callback HFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : callback HRFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

// Unsupported : callback HashFunc : unsupported parameter key : no type generator for gpointer (gconstpointer) for param key

var callbackHookcheckfuncId int
var callbackHookcheckfuncMap = make(map[int]HookcheckfuncCallback)
var callbackHookcheckfuncLock sync.RWMutex

// HookcheckfuncCallback is a callback function for a 'HookCheckFunc' callback.
type HookcheckfuncCallback func() bool

// Unsupported : callback HookCheckMarshaller : unsupported parameter marshal_data : no type generator for gpointer (gpointer) for param marshal_data

var callbackHookcomparefuncId int
var callbackHookcomparefuncMap = make(map[int]HookcomparefuncCallback)
var callbackHookcomparefuncLock sync.RWMutex

// HookcomparefuncCallback is a callback function for a 'HookCompareFunc' callback.
type HookcomparefuncCallback func(newHook *Hook, sibling *Hook) int32

var callbackHookfinalizefuncId int
var callbackHookfinalizefuncMap = make(map[int]HookfinalizefuncCallback)
var callbackHookfinalizefuncLock sync.RWMutex

// HookfinalizefuncCallback is a callback function for a 'HookFinalizeFunc' callback.
type HookfinalizefuncCallback func(hookList *HookList, hook *Hook)

var callbackHookfindfuncId int
var callbackHookfindfuncMap = make(map[int]HookfindfuncCallback)
var callbackHookfindfuncLock sync.RWMutex

// HookfindfuncCallback is a callback function for a 'HookFindFunc' callback.
type HookfindfuncCallback func(hook *Hook) bool

var callbackHookfuncId int
var callbackHookfuncMap = make(map[int]HookfuncCallback)
var callbackHookfuncLock sync.RWMutex

// HookfuncCallback is a callback function for a 'HookFunc' callback.
type HookfuncCallback func()

// Unsupported : callback HookMarshaller : unsupported parameter marshal_data : no type generator for gpointer (gpointer) for param marshal_data

var callbackIofuncId int
var callbackIofuncMap = make(map[int]IofuncCallback)
var callbackIofuncLock sync.RWMutex

// IofuncCallback is a callback function for a 'IOFunc' callback.
type IofuncCallback func(source *IOChannel, condition IOCondition) bool

var callbackLogfuncId int
var callbackLogfuncMap = make(map[int]LogfuncCallback)
var callbackLogfuncLock sync.RWMutex

// LogfuncCallback is a callback function for a 'LogFunc' callback.
type LogfuncCallback func(logDomain string, logLevel LogLevelFlags, message string)

var callbackNodeforeachfuncId int
var callbackNodeforeachfuncMap = make(map[int]NodeforeachfuncCallback)
var callbackNodeforeachfuncLock sync.RWMutex

// NodeforeachfuncCallback is a callback function for a 'NodeForeachFunc' callback.
type NodeforeachfuncCallback func(node *Node)

var callbackNodetraversefuncId int
var callbackNodetraversefuncMap = make(map[int]NodetraversefuncCallback)
var callbackNodetraversefuncLock sync.RWMutex

// NodetraversefuncCallback is a callback function for a 'NodeTraverseFunc' callback.
type NodetraversefuncCallback func(node *Node) bool

var callbackOptionargfuncId int
var callbackOptionargfuncMap = make(map[int]OptionargfuncCallback)
var callbackOptionargfuncLock sync.RWMutex

// OptionargfuncCallback is a callback function for a 'OptionArgFunc' callback.
type OptionargfuncCallback func(optionName string, value string) bool

var callbackOptionerrorfuncId int
var callbackOptionerrorfuncMap = make(map[int]OptionerrorfuncCallback)
var callbackOptionerrorfuncLock sync.RWMutex

// OptionerrorfuncCallback is a callback function for a 'OptionErrorFunc' callback.
type OptionerrorfuncCallback func(context *OptionContext, group *OptionGroup)

var callbackOptionparsefuncId int
var callbackOptionparsefuncMap = make(map[int]OptionparsefuncCallback)
var callbackOptionparsefuncLock sync.RWMutex

// OptionparsefuncCallback is a callback function for a 'OptionParseFunc' callback.
type OptionparsefuncCallback func(context *OptionContext, group *OptionGroup) bool

var callbackPollfuncId int
var callbackPollfuncMap = make(map[int]PollfuncCallback)
var callbackPollfuncLock sync.RWMutex

// PollfuncCallback is a callback function for a 'PollFunc' callback.
type PollfuncCallback func(ufds *PollFD, nfsd uint32, timeout int32) int32

var callbackPrintfuncId int
var callbackPrintfuncMap = make(map[int]PrintfuncCallback)
var callbackPrintfuncLock sync.RWMutex

// PrintfuncCallback is a callback function for a 'PrintFunc' callback.
type PrintfuncCallback func(string_ string)

var callbackScannermsgfuncId int
var callbackScannermsgfuncMap = make(map[int]ScannermsgfuncCallback)
var callbackScannermsgfuncLock sync.RWMutex

// ScannermsgfuncCallback is a callback function for a 'ScannerMsgFunc' callback.
type ScannermsgfuncCallback func(scanner *Scanner, message string, error bool)

var callbackSequenceitercomparefuncId int
var callbackSequenceitercomparefuncMap = make(map[int]SequenceitercomparefuncCallback)
var callbackSequenceitercomparefuncLock sync.RWMutex

// SequenceitercomparefuncCallback is a callback function for a 'SequenceIterCompareFunc' callback.
type SequenceitercomparefuncCallback func(a *SequenceIter, b *SequenceIter) int32

var callbackSourcedummymarshalId int
var callbackSourcedummymarshalMap = make(map[int]SourcedummymarshalCallback)
var callbackSourcedummymarshalLock sync.RWMutex

// SourcedummymarshalCallback is a callback function for a 'SourceDummyMarshal' callback.
type SourcedummymarshalCallback func()

var callbackSourcefuncId int
var callbackSourcefuncMap = make(map[int]SourcefuncCallback)
var callbackSourcefuncLock sync.RWMutex

// SourcefuncCallback is a callback function for a 'SourceFunc' callback.
type SourcefuncCallback func() bool

var callbackSpawnchildsetupfuncId int
var callbackSpawnchildsetupfuncMap = make(map[int]SpawnchildsetupfuncCallback)
var callbackSpawnchildsetupfuncLock sync.RWMutex

// SpawnchildsetupfuncCallback is a callback function for a 'SpawnChildSetupFunc' callback.
type SpawnchildsetupfuncCallback func()

// Unsupported : callback ThreadFunc : no return generator

var callbackTranslatefuncId int
var callbackTranslatefuncMap = make(map[int]TranslatefuncCallback)
var callbackTranslatefuncLock sync.RWMutex

// TranslatefuncCallback is a callback function for a 'TranslateFunc' callback.
type TranslatefuncCallback func(str string) string

// Unsupported : callback TraverseFunc : unsupported parameter key : no type generator for gpointer (gpointer) for param key

var callbackUnixfdsourcefuncId int
var callbackUnixfdsourcefuncMap = make(map[int]UnixfdsourcefuncCallback)
var callbackUnixfdsourcefuncLock sync.RWMutex

// UnixfdsourcefuncCallback is a callback function for a 'UnixFDSourceFunc' callback.
type UnixfdsourcefuncCallback func(fd int32, condition IOCondition) bool

var callbackVoidfuncId int
var callbackVoidfuncMap = make(map[int]VoidfuncCallback)
var callbackVoidfuncLock sync.RWMutex

// VoidfuncCallback is a callback function for a 'VoidFunc' callback.
type VoidfuncCallback func()
