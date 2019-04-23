// This is a generated file - DO NOT EDIT
// +build gdk_3.20 gdk_3.22

package gdk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	void dragcontext_actionChangedHandler(GObject *, GdkDragAction, gpointer);

	static gulong DragContext_signal_connect_action_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-changed", G_CALLBACK(dragcontext_actionChangedHandler), data);
	}

*/
/*

	void dragcontext_cancelHandler(GObject *, GdkDragCancelReason, gpointer);

	static gulong DragContext_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", G_CALLBACK(dragcontext_cancelHandler), data);
	}

*/
/*

	void dragcontext_dndFinishedHandler(GObject *, gpointer);

	static gulong DragContext_signal_connect_dnd_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "dnd-finished", G_CALLBACK(dragcontext_dndFinishedHandler), data);
	}

*/
/*

	void dragcontext_dropPerformedHandler(GObject *, gint, gpointer);

	static gulong DragContext_signal_connect_drop_performed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drop-performed", G_CALLBACK(dragcontext_dropPerformedHandler), data);
	}

*/
import "C"

type signalDragContextActionChangedDetail struct {
	callback  DragContextSignalActionChangedCallback
	handlerID C.gulong
}

var signalDragContextActionChangedId int
var signalDragContextActionChangedMap = make(map[int]signalDragContextActionChangedDetail)
var signalDragContextActionChangedLock sync.RWMutex

// DragContextSignalActionChangedCallback is a callback function for a 'action-changed' signal emitted from a DragContext.
type DragContextSignalActionChangedCallback func(action DragAction)

/*
ConnectActionChanged connects the callback to the 'action-changed' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectActionChanged to remove it.
*/
func (recv *DragContext) ConnectActionChanged(callback DragContextSignalActionChangedCallback) int {
	signalDragContextActionChangedLock.Lock()
	defer signalDragContextActionChangedLock.Unlock()

	signalDragContextActionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_action_changed(instance, C.gpointer(uintptr(signalDragContextActionChangedId)))

	detail := signalDragContextActionChangedDetail{callback, handlerID}
	signalDragContextActionChangedMap[signalDragContextActionChangedId] = detail

	return signalDragContextActionChangedId
}

/*
DisconnectActionChanged disconnects a callback from the 'action-changed' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectActionChanged.
*/
func (recv *DragContext) DisconnectActionChanged(connectionID int) {
	signalDragContextActionChangedLock.Lock()
	defer signalDragContextActionChangedLock.Unlock()

	detail, exists := signalDragContextActionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextActionChangedMap, connectionID)
}

//export dragcontext_actionChangedHandler
func dragcontext_actionChangedHandler(_ *C.GObject, c_action C.GdkDragAction, data C.gpointer) {
	signalDragContextActionChangedLock.RLock()
	defer signalDragContextActionChangedLock.RUnlock()

	action := DragAction(c_action)

	index := int(uintptr(data))
	callback := signalDragContextActionChangedMap[index].callback
	callback(action)
}

type signalDragContextCancelDetail struct {
	callback  DragContextSignalCancelCallback
	handlerID C.gulong
}

var signalDragContextCancelId int
var signalDragContextCancelMap = make(map[int]signalDragContextCancelDetail)
var signalDragContextCancelLock sync.RWMutex

// DragContextSignalCancelCallback is a callback function for a 'cancel' signal emitted from a DragContext.
type DragContextSignalCancelCallback func(reason DragCancelReason)

/*
ConnectCancel connects the callback to the 'cancel' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectCancel to remove it.
*/
func (recv *DragContext) ConnectCancel(callback DragContextSignalCancelCallback) int {
	signalDragContextCancelLock.Lock()
	defer signalDragContextCancelLock.Unlock()

	signalDragContextCancelId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_cancel(instance, C.gpointer(uintptr(signalDragContextCancelId)))

	detail := signalDragContextCancelDetail{callback, handlerID}
	signalDragContextCancelMap[signalDragContextCancelId] = detail

	return signalDragContextCancelId
}

/*
DisconnectCancel disconnects a callback from the 'cancel' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectCancel.
*/
func (recv *DragContext) DisconnectCancel(connectionID int) {
	signalDragContextCancelLock.Lock()
	defer signalDragContextCancelLock.Unlock()

	detail, exists := signalDragContextCancelMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextCancelMap, connectionID)
}

//export dragcontext_cancelHandler
func dragcontext_cancelHandler(_ *C.GObject, c_reason C.GdkDragCancelReason, data C.gpointer) {
	signalDragContextCancelLock.RLock()
	defer signalDragContextCancelLock.RUnlock()

	reason := DragCancelReason(c_reason)

	index := int(uintptr(data))
	callback := signalDragContextCancelMap[index].callback
	callback(reason)
}

type signalDragContextDndFinishedDetail struct {
	callback  DragContextSignalDndFinishedCallback
	handlerID C.gulong
}

var signalDragContextDndFinishedId int
var signalDragContextDndFinishedMap = make(map[int]signalDragContextDndFinishedDetail)
var signalDragContextDndFinishedLock sync.RWMutex

// DragContextSignalDndFinishedCallback is a callback function for a 'dnd-finished' signal emitted from a DragContext.
type DragContextSignalDndFinishedCallback func()

/*
ConnectDndFinished connects the callback to the 'dnd-finished' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectDndFinished to remove it.
*/
func (recv *DragContext) ConnectDndFinished(callback DragContextSignalDndFinishedCallback) int {
	signalDragContextDndFinishedLock.Lock()
	defer signalDragContextDndFinishedLock.Unlock()

	signalDragContextDndFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_dnd_finished(instance, C.gpointer(uintptr(signalDragContextDndFinishedId)))

	detail := signalDragContextDndFinishedDetail{callback, handlerID}
	signalDragContextDndFinishedMap[signalDragContextDndFinishedId] = detail

	return signalDragContextDndFinishedId
}

/*
DisconnectDndFinished disconnects a callback from the 'dnd-finished' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectDndFinished.
*/
func (recv *DragContext) DisconnectDndFinished(connectionID int) {
	signalDragContextDndFinishedLock.Lock()
	defer signalDragContextDndFinishedLock.Unlock()

	detail, exists := signalDragContextDndFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextDndFinishedMap, connectionID)
}

//export dragcontext_dndFinishedHandler
func dragcontext_dndFinishedHandler(_ *C.GObject, data C.gpointer) {
	signalDragContextDndFinishedLock.RLock()
	defer signalDragContextDndFinishedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDragContextDndFinishedMap[index].callback
	callback()
}

type signalDragContextDropPerformedDetail struct {
	callback  DragContextSignalDropPerformedCallback
	handlerID C.gulong
}

var signalDragContextDropPerformedId int
var signalDragContextDropPerformedMap = make(map[int]signalDragContextDropPerformedDetail)
var signalDragContextDropPerformedLock sync.RWMutex

// DragContextSignalDropPerformedCallback is a callback function for a 'drop-performed' signal emitted from a DragContext.
type DragContextSignalDropPerformedCallback func(time int32)

/*
ConnectDropPerformed connects the callback to the 'drop-performed' signal for the DragContext.

The returned value represents the connection, and may be passed to DisconnectDropPerformed to remove it.
*/
func (recv *DragContext) ConnectDropPerformed(callback DragContextSignalDropPerformedCallback) int {
	signalDragContextDropPerformedLock.Lock()
	defer signalDragContextDropPerformedLock.Unlock()

	signalDragContextDropPerformedId++
	instance := C.gpointer(recv.native)
	handlerID := C.DragContext_signal_connect_drop_performed(instance, C.gpointer(uintptr(signalDragContextDropPerformedId)))

	detail := signalDragContextDropPerformedDetail{callback, handlerID}
	signalDragContextDropPerformedMap[signalDragContextDropPerformedId] = detail

	return signalDragContextDropPerformedId
}

/*
DisconnectDropPerformed disconnects a callback from the 'drop-performed' signal for the DragContext.

The connectionID should be a value returned from a call to ConnectDropPerformed.
*/
func (recv *DragContext) DisconnectDropPerformed(connectionID int) {
	signalDragContextDropPerformedLock.Lock()
	defer signalDragContextDropPerformedLock.Unlock()

	detail, exists := signalDragContextDropPerformedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDragContextDropPerformedMap, connectionID)
}

//export dragcontext_dropPerformedHandler
func dragcontext_dropPerformedHandler(_ *C.GObject, c_time C.gint, data C.gpointer) {
	signalDragContextDropPerformedLock.RLock()
	defer signalDragContextDropPerformedLock.RUnlock()

	time := int32(c_time)

	index := int(uintptr(data))
	callback := signalDragContextDropPerformedMap[index].callback
	callback(time)
}

// Blacklisted : gdk_drag_context_get_drag_window

// Blacklisted : gdk_drag_context_manage_dnd

// Blacklisted : gdk_drag_context_set_hotspot
