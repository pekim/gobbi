// This is a generated file - DO NOT EDIT

package atk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void callback_eventlistenerHandler(GObject *, AtkObject*, gpointer);

*/
/*

	void callback_eventlistenerinitHandler(GObject *, gpointer);

*/
/*

	void callback_focushandlerHandler(GObject *, AtkObject*, gboolean, gpointer);

*/
/*

	gboolean callback_functionHandler(GObject *, gpointer, gpointer);

*/
/*

	gint callback_keysnoopfuncHandler(GObject *, AtkKeyEventStruct*, gpointer, gpointer);

*/
/*

	void callback_propertychangehandlerHandler(GObject *, AtkObject*, AtkPropertyValues*, gpointer);

*/
import "C"

var callbackEventlistenerId int
var callbackEventlistenerMap = make(map[int]EventlistenerCallback)
var callbackEventlistenerLock sync.RWMutex

// EventlistenerCallback is a callback function for a 'EventListener' callback.
type EventlistenerCallback func(obj *Object)

var callbackEventlistenerinitId int
var callbackEventlistenerinitMap = make(map[int]EventlistenerinitCallback)
var callbackEventlistenerinitLock sync.RWMutex

// EventlistenerinitCallback is a callback function for a 'EventListenerInit' callback.
type EventlistenerinitCallback func()

var callbackFocushandlerId int
var callbackFocushandlerMap = make(map[int]FocushandlerCallback)
var callbackFocushandlerLock sync.RWMutex

// FocushandlerCallback is a callback function for a 'FocusHandler' callback.
type FocushandlerCallback func(object *Object, focusIn bool)

var callbackFunctionId int
var callbackFunctionMap = make(map[int]FunctionCallback)
var callbackFunctionLock sync.RWMutex

// FunctionCallback is a callback function for a 'Function' callback.
type FunctionCallback func() bool

var callbackKeysnoopfuncId int
var callbackKeysnoopfuncMap = make(map[int]KeysnoopfuncCallback)
var callbackKeysnoopfuncLock sync.RWMutex

// KeysnoopfuncCallback is a callback function for a 'KeySnoopFunc' callback.
type KeysnoopfuncCallback func(event *KeyEventStruct) int32

var callbackPropertychangehandlerId int
var callbackPropertychangehandlerMap = make(map[int]PropertychangehandlerCallback)
var callbackPropertychangehandlerLock sync.RWMutex

// PropertychangehandlerCallback is a callback function for a 'PropertyChangeHandler' callback.
type PropertychangehandlerCallback func(obj *Object, vals *PropertyValues)
