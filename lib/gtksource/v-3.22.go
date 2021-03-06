// Code generated - DO NOT EDIT.
// +build gtksource_3.22

package gtksource

import (
	atk "github.com/pekim/gobbi/lib/atk"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gtk "github.com/pekim/gobbi/lib/gtk"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtksourceview/gtksource.h>
// #include <gtksourceview/completion-providers/words/gtksourcecompletionwords.h>
// #include <../gdk/gdk_event.h>
// #include <stdlib.h>
/*

	void buffer_bracketMatchedHandler(GObject *, GtkTextIter *, GtkSourceBracketMatchType, gpointer);

	static gulong Buffer_signal_connect_bracket_matched(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "bracket-matched", G_CALLBACK(buffer_bracketMatchedHandler), data);
	}

*/
/*

	void buffer_highlightUpdatedHandler(GObject *, GtkTextIter *, GtkTextIter *, gpointer);

	static gulong Buffer_signal_connect_highlight_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "highlight-updated", G_CALLBACK(buffer_highlightUpdatedHandler), data);
	}

*/
/*

	void buffer_redoHandler(GObject *, gpointer);

	static gulong Buffer_signal_connect_redo(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "redo", G_CALLBACK(buffer_redoHandler), data);
	}

*/
/*

	void buffer_sourceMarkUpdatedHandler(GObject *, GtkTextMark *, gpointer);

	static gulong Buffer_signal_connect_source_mark_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "source-mark-updated", G_CALLBACK(buffer_sourceMarkUpdatedHandler), data);
	}

*/
/*

	void buffer_undoHandler(GObject *, gpointer);

	static gulong Buffer_signal_connect_undo(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "undo", G_CALLBACK(buffer_undoHandler), data);
	}

*/
/*

	void completion_activateProposalHandler(GObject *, gpointer);

	static gulong Completion_signal_connect_activate_proposal(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-proposal", G_CALLBACK(completion_activateProposalHandler), data);
	}

*/
/*

	void completion_hideHandler(GObject *, gpointer);

	static gulong Completion_signal_connect_hide(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "hide", G_CALLBACK(completion_hideHandler), data);
	}

*/
/*

	void completion_moveCursorHandler(GObject *, GtkScrollStep, gint, gpointer);

	static gulong Completion_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", G_CALLBACK(completion_moveCursorHandler), data);
	}

*/
/*

	void completion_movePageHandler(GObject *, GtkScrollStep, gint, gpointer);

	static gulong Completion_signal_connect_move_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-page", G_CALLBACK(completion_movePageHandler), data);
	}

*/
/*

	void completion_populateContextHandler(GObject *, GtkSourceCompletionContext *, gpointer);

	static gulong Completion_signal_connect_populate_context(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-context", G_CALLBACK(completion_populateContextHandler), data);
	}

*/
/*

	void completion_showHandler(GObject *, gpointer);

	static gulong Completion_signal_connect_show(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show", G_CALLBACK(completion_showHandler), data);
	}

*/
/*

	void completioncontext_cancelledHandler(GObject *, gpointer);

	static gulong CompletionContext_signal_connect_cancelled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancelled", G_CALLBACK(completioncontext_cancelledHandler), data);
	}

*/
/*

	void completioninfo_beforeShowHandler(GObject *, gpointer);

	static gulong CompletionInfo_signal_connect_before_show(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "before-show", G_CALLBACK(completioninfo_beforeShowHandler), data);
	}

*/
/*

	void gutterrenderer_activateHandler(GObject *, GtkTextIter *, GdkRectangle *, GdkEvent_ *, gpointer);

	static gulong GutterRenderer_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(gutterrenderer_activateHandler), data);
	}

*/
/*

	gboolean gutterrenderer_queryActivatableHandler(GObject *, GtkTextIter *, GdkRectangle *, GdkEvent_ *, gpointer);

	static gulong GutterRenderer_signal_connect_query_activatable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-activatable", G_CALLBACK(gutterrenderer_queryActivatableHandler), data);
	}

*/
/*

	void gutterrenderer_queryDataHandler(GObject *, GtkTextIter *, GtkTextIter *, GtkSourceGutterRendererState, gpointer);

	static gulong GutterRenderer_signal_connect_query_data(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-data", G_CALLBACK(gutterrenderer_queryDataHandler), data);
	}

*/
/*

	gboolean gutterrenderer_queryTooltipHandler(GObject *, GtkTextIter *, GdkRectangle *, gint, gint, GtkTooltip *, gpointer);

	static gulong GutterRenderer_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", G_CALLBACK(gutterrenderer_queryTooltipHandler), data);
	}

*/
/*

	void gutterrenderer_queueDrawHandler(GObject *, gpointer);

	static gulong GutterRenderer_signal_connect_queue_draw(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "queue-draw", G_CALLBACK(gutterrenderer_queueDrawHandler), data);
	}

*/
/*

	void view_changeCaseHandler(GObject *, GtkSourceChangeCaseType, gpointer);

	static gulong View_signal_connect_change_case(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-case", G_CALLBACK(view_changeCaseHandler), data);
	}

*/
/*

	void view_changeNumberHandler(GObject *, gint, gpointer);

	static gulong View_signal_connect_change_number(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-number", G_CALLBACK(view_changeNumberHandler), data);
	}

*/
/*

	void view_joinLinesHandler(GObject *, gpointer);

	static gulong View_signal_connect_join_lines(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "join-lines", G_CALLBACK(view_joinLinesHandler), data);
	}

*/
/*

	void view_lineMarkActivatedHandler(GObject *, GtkTextIter *, GdkEvent_ *, gpointer);

	static gulong View_signal_connect_line_mark_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "line-mark-activated", G_CALLBACK(view_lineMarkActivatedHandler), data);
	}

*/
/*

	void view_moveLinesHandler(GObject *, gboolean, gint, gpointer);

	static gulong View_signal_connect_move_lines(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-lines", G_CALLBACK(view_moveLinesHandler), data);
	}

*/
/*

	void view_moveToMatchingBracketHandler(GObject *, gboolean, gpointer);

	static gulong View_signal_connect_move_to_matching_bracket(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-to-matching-bracket", G_CALLBACK(view_moveToMatchingBracketHandler), data);
	}

*/
/*

	void view_moveWordsHandler(GObject *, gint, gpointer);

	static gulong View_signal_connect_move_words(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-words", G_CALLBACK(view_moveWordsHandler), data);
	}

*/
/*

	void view_redoHandler(GObject *, gpointer);

	static gulong View_signal_connect_redo(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "redo", G_CALLBACK(view_redoHandler), data);
	}

*/
/*

	void view_showCompletionHandler(GObject *, gpointer);

	static gulong View_signal_connect_show_completion(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-completion", G_CALLBACK(view_showCompletionHandler), data);
	}

*/
/*

	void view_smartHomeEndHandler(GObject *, GtkTextIter *, gint, gpointer);

	static gulong View_signal_connect_smart_home_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "smart-home-end", G_CALLBACK(view_smartHomeEndHandler), data);
	}

*/
/*

	void view_undoHandler(GObject *, gpointer);

	static gulong View_signal_connect_undo(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "undo", G_CALLBACK(view_undoHandler), data);
	}

*/
/*

	void completionproposal_changedHandler(GObject *, gpointer);

	static gulong CompletionProposal_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(completionproposal_changedHandler), data);
	}

*/
/*

	void undomanager_canRedoChangedHandler(GObject *, gpointer);

	static gulong UndoManager_signal_connect_can_redo_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "can-redo-changed", G_CALLBACK(undomanager_canRedoChangedHandler), data);
	}

*/
/*

	void undomanager_canUndoChangedHandler(GObject *, gpointer);

	static gulong UndoManager_signal_connect_can_undo_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "can-undo-changed", G_CALLBACK(undomanager_canUndoChangedHandler), data);
	}

*/
import "C"

type CompletionActivation C.GtkSourceCompletionActivation

const (
	GTK_SOURCE_COMPLETION_ACTIVATION_NONE           CompletionActivation = 0
	GTK_SOURCE_COMPLETION_ACTIVATION_INTERACTIVE    CompletionActivation = 1
	GTK_SOURCE_COMPLETION_ACTIVATION_USER_REQUESTED CompletionActivation = 2
)

type DrawSpacesFlags C.GtkSourceDrawSpacesFlags

const (
	GTK_SOURCE_DRAW_SPACES_SPACE    DrawSpacesFlags = 1
	GTK_SOURCE_DRAW_SPACES_TAB      DrawSpacesFlags = 2
	GTK_SOURCE_DRAW_SPACES_NEWLINE  DrawSpacesFlags = 4
	GTK_SOURCE_DRAW_SPACES_NBSP     DrawSpacesFlags = 8
	GTK_SOURCE_DRAW_SPACES_LEADING  DrawSpacesFlags = 16
	GTK_SOURCE_DRAW_SPACES_TEXT     DrawSpacesFlags = 32
	GTK_SOURCE_DRAW_SPACES_TRAILING DrawSpacesFlags = 64
	GTK_SOURCE_DRAW_SPACES_ALL      DrawSpacesFlags = 127
)

type FileSaverFlags C.GtkSourceFileSaverFlags

const (
	GTK_SOURCE_FILE_SAVER_FLAGS_NONE                     FileSaverFlags = 0
	GTK_SOURCE_FILE_SAVER_FLAGS_IGNORE_INVALID_CHARS     FileSaverFlags = 1
	GTK_SOURCE_FILE_SAVER_FLAGS_IGNORE_MODIFICATION_TIME FileSaverFlags = 2
	GTK_SOURCE_FILE_SAVER_FLAGS_CREATE_BACKUP            FileSaverFlags = 4
)

type GutterRendererState C.GtkSourceGutterRendererState

const (
	GTK_SOURCE_GUTTER_RENDERER_STATE_NORMAL   GutterRendererState = 0
	GTK_SOURCE_GUTTER_RENDERER_STATE_CURSOR   GutterRendererState = 1
	GTK_SOURCE_GUTTER_RENDERER_STATE_PRELIT   GutterRendererState = 2
	GTK_SOURCE_GUTTER_RENDERER_STATE_SELECTED GutterRendererState = 4
)

type SortFlags C.GtkSourceSortFlags

const (
	GTK_SOURCE_SORT_FLAGS_NONE              SortFlags = 0
	GTK_SOURCE_SORT_FLAGS_CASE_SENSITIVE    SortFlags = 1
	GTK_SOURCE_SORT_FLAGS_REVERSE_ORDER     SortFlags = 2
	GTK_SOURCE_SORT_FLAGS_REMOVE_DUPLICATES SortFlags = 4
)

// Buffer is a wrapper around the C record GtkSourceBuffer.
type Buffer struct {
	native *C.GtkSourceBuffer
	// parent_instance : record
	// priv : record
}

func BufferNewFromC(u unsafe.Pointer) *Buffer {
	c := (*C.GtkSourceBuffer)(u)
	if c == nil {
		return nil
	}

	g := &Buffer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Buffer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Buffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Buffer with another Buffer, and returns true if they represent the same GObject.
func (recv *Buffer) Equals(other *Buffer) bool {
	return other.ToC() == recv.ToC()
}

// TextBuffer upcasts to *TextBuffer
func (recv *Buffer) TextBuffer() *gtk.TextBuffer {
	return gtk.TextBufferNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Buffer) Object() *gobject.Object {
	return recv.TextBuffer().Object()
}

// CastToWidget down casts any arbitrary Object to Buffer.
// Exercise care, as this is a potentially dangerous function if the Object is not a Buffer.
func CastToBuffer(object *gobject.Object) *Buffer {
	return BufferNewFromC(object.ToC())
}

type signalBufferBracketMatchedDetail struct {
	callback  BufferSignalBracketMatchedCallback
	handlerID C.gulong
}

var signalBufferBracketMatchedId int
var signalBufferBracketMatchedMap = make(map[int]signalBufferBracketMatchedDetail)
var signalBufferBracketMatchedLock sync.RWMutex

// BufferSignalBracketMatchedCallback is a callback function for a 'bracket-matched' signal emitted from a Buffer.
type BufferSignalBracketMatchedCallback func(targetObject *Buffer, iter *gtk.TextIter, state BracketMatchType)

/*
ConnectBracketMatched connects the callback to the 'bracket-matched' signal for the Buffer.

The returned value represents the connection, and may be passed to DisconnectBracketMatched to remove it.
*/
func (recv *Buffer) ConnectBracketMatched(callback BufferSignalBracketMatchedCallback) int {
	signalBufferBracketMatchedLock.Lock()
	defer signalBufferBracketMatchedLock.Unlock()

	signalBufferBracketMatchedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Buffer_signal_connect_bracket_matched(instance, C.gpointer(uintptr(signalBufferBracketMatchedId)))

	detail := signalBufferBracketMatchedDetail{callback, handlerID}
	signalBufferBracketMatchedMap[signalBufferBracketMatchedId] = detail

	return signalBufferBracketMatchedId
}

/*
DisconnectBracketMatched disconnects a callback from the 'bracket-matched' signal for the Buffer.

The connectionID should be a value returned from a call to ConnectBracketMatched.
*/
func (recv *Buffer) DisconnectBracketMatched(connectionID int) {
	signalBufferBracketMatchedLock.Lock()
	defer signalBufferBracketMatchedLock.Unlock()

	detail, exists := signalBufferBracketMatchedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalBufferBracketMatchedMap, connectionID)
}

//export buffer_bracketMatchedHandler
func buffer_bracketMatchedHandler(c_targetObject *C.GObject, c_iter *C.GtkTextIter, c_state C.GtkSourceBracketMatchType, data C.gpointer) {
	signalBufferBracketMatchedLock.RLock()
	defer signalBufferBracketMatchedLock.RUnlock()

	iter := gtk.TextIterNewFromC(unsafe.Pointer(c_iter))

	state := BracketMatchType(c_state)

	targetObject := BufferNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalBufferBracketMatchedMap[index].callback
	callback(targetObject, iter, state)
}

type signalBufferHighlightUpdatedDetail struct {
	callback  BufferSignalHighlightUpdatedCallback
	handlerID C.gulong
}

var signalBufferHighlightUpdatedId int
var signalBufferHighlightUpdatedMap = make(map[int]signalBufferHighlightUpdatedDetail)
var signalBufferHighlightUpdatedLock sync.RWMutex

// BufferSignalHighlightUpdatedCallback is a callback function for a 'highlight-updated' signal emitted from a Buffer.
type BufferSignalHighlightUpdatedCallback func(targetObject *Buffer, start *gtk.TextIter, end *gtk.TextIter)

/*
ConnectHighlightUpdated connects the callback to the 'highlight-updated' signal for the Buffer.

The returned value represents the connection, and may be passed to DisconnectHighlightUpdated to remove it.
*/
func (recv *Buffer) ConnectHighlightUpdated(callback BufferSignalHighlightUpdatedCallback) int {
	signalBufferHighlightUpdatedLock.Lock()
	defer signalBufferHighlightUpdatedLock.Unlock()

	signalBufferHighlightUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Buffer_signal_connect_highlight_updated(instance, C.gpointer(uintptr(signalBufferHighlightUpdatedId)))

	detail := signalBufferHighlightUpdatedDetail{callback, handlerID}
	signalBufferHighlightUpdatedMap[signalBufferHighlightUpdatedId] = detail

	return signalBufferHighlightUpdatedId
}

/*
DisconnectHighlightUpdated disconnects a callback from the 'highlight-updated' signal for the Buffer.

The connectionID should be a value returned from a call to ConnectHighlightUpdated.
*/
func (recv *Buffer) DisconnectHighlightUpdated(connectionID int) {
	signalBufferHighlightUpdatedLock.Lock()
	defer signalBufferHighlightUpdatedLock.Unlock()

	detail, exists := signalBufferHighlightUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalBufferHighlightUpdatedMap, connectionID)
}

//export buffer_highlightUpdatedHandler
func buffer_highlightUpdatedHandler(c_targetObject *C.GObject, c_start *C.GtkTextIter, c_end *C.GtkTextIter, data C.gpointer) {
	signalBufferHighlightUpdatedLock.RLock()
	defer signalBufferHighlightUpdatedLock.RUnlock()

	start := gtk.TextIterNewFromC(unsafe.Pointer(c_start))

	end := gtk.TextIterNewFromC(unsafe.Pointer(c_end))

	targetObject := BufferNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalBufferHighlightUpdatedMap[index].callback
	callback(targetObject, start, end)
}

type signalBufferRedoDetail struct {
	callback  BufferSignalRedoCallback
	handlerID C.gulong
}

var signalBufferRedoId int
var signalBufferRedoMap = make(map[int]signalBufferRedoDetail)
var signalBufferRedoLock sync.RWMutex

// BufferSignalRedoCallback is a callback function for a 'redo' signal emitted from a Buffer.
type BufferSignalRedoCallback func(targetObject *Buffer)

/*
ConnectRedo connects the callback to the 'redo' signal for the Buffer.

The returned value represents the connection, and may be passed to DisconnectRedo to remove it.
*/
func (recv *Buffer) ConnectRedo(callback BufferSignalRedoCallback) int {
	signalBufferRedoLock.Lock()
	defer signalBufferRedoLock.Unlock()

	signalBufferRedoId++
	instance := C.gpointer(recv.native)
	handlerID := C.Buffer_signal_connect_redo(instance, C.gpointer(uintptr(signalBufferRedoId)))

	detail := signalBufferRedoDetail{callback, handlerID}
	signalBufferRedoMap[signalBufferRedoId] = detail

	return signalBufferRedoId
}

/*
DisconnectRedo disconnects a callback from the 'redo' signal for the Buffer.

The connectionID should be a value returned from a call to ConnectRedo.
*/
func (recv *Buffer) DisconnectRedo(connectionID int) {
	signalBufferRedoLock.Lock()
	defer signalBufferRedoLock.Unlock()

	detail, exists := signalBufferRedoMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalBufferRedoMap, connectionID)
}

//export buffer_redoHandler
func buffer_redoHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalBufferRedoLock.RLock()
	defer signalBufferRedoLock.RUnlock()

	targetObject := BufferNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalBufferRedoMap[index].callback
	callback(targetObject)
}

type signalBufferSourceMarkUpdatedDetail struct {
	callback  BufferSignalSourceMarkUpdatedCallback
	handlerID C.gulong
}

var signalBufferSourceMarkUpdatedId int
var signalBufferSourceMarkUpdatedMap = make(map[int]signalBufferSourceMarkUpdatedDetail)
var signalBufferSourceMarkUpdatedLock sync.RWMutex

// BufferSignalSourceMarkUpdatedCallback is a callback function for a 'source-mark-updated' signal emitted from a Buffer.
type BufferSignalSourceMarkUpdatedCallback func(targetObject *Buffer, mark *gtk.TextMark)

/*
ConnectSourceMarkUpdated connects the callback to the 'source-mark-updated' signal for the Buffer.

The returned value represents the connection, and may be passed to DisconnectSourceMarkUpdated to remove it.
*/
func (recv *Buffer) ConnectSourceMarkUpdated(callback BufferSignalSourceMarkUpdatedCallback) int {
	signalBufferSourceMarkUpdatedLock.Lock()
	defer signalBufferSourceMarkUpdatedLock.Unlock()

	signalBufferSourceMarkUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Buffer_signal_connect_source_mark_updated(instance, C.gpointer(uintptr(signalBufferSourceMarkUpdatedId)))

	detail := signalBufferSourceMarkUpdatedDetail{callback, handlerID}
	signalBufferSourceMarkUpdatedMap[signalBufferSourceMarkUpdatedId] = detail

	return signalBufferSourceMarkUpdatedId
}

/*
DisconnectSourceMarkUpdated disconnects a callback from the 'source-mark-updated' signal for the Buffer.

The connectionID should be a value returned from a call to ConnectSourceMarkUpdated.
*/
func (recv *Buffer) DisconnectSourceMarkUpdated(connectionID int) {
	signalBufferSourceMarkUpdatedLock.Lock()
	defer signalBufferSourceMarkUpdatedLock.Unlock()

	detail, exists := signalBufferSourceMarkUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalBufferSourceMarkUpdatedMap, connectionID)
}

//export buffer_sourceMarkUpdatedHandler
func buffer_sourceMarkUpdatedHandler(c_targetObject *C.GObject, c_mark *C.GtkTextMark, data C.gpointer) {
	signalBufferSourceMarkUpdatedLock.RLock()
	defer signalBufferSourceMarkUpdatedLock.RUnlock()

	mark := gtk.TextMarkNewFromC(unsafe.Pointer(c_mark))

	targetObject := BufferNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalBufferSourceMarkUpdatedMap[index].callback
	callback(targetObject, mark)
}

type signalBufferUndoDetail struct {
	callback  BufferSignalUndoCallback
	handlerID C.gulong
}

var signalBufferUndoId int
var signalBufferUndoMap = make(map[int]signalBufferUndoDetail)
var signalBufferUndoLock sync.RWMutex

// BufferSignalUndoCallback is a callback function for a 'undo' signal emitted from a Buffer.
type BufferSignalUndoCallback func(targetObject *Buffer)

/*
ConnectUndo connects the callback to the 'undo' signal for the Buffer.

The returned value represents the connection, and may be passed to DisconnectUndo to remove it.
*/
func (recv *Buffer) ConnectUndo(callback BufferSignalUndoCallback) int {
	signalBufferUndoLock.Lock()
	defer signalBufferUndoLock.Unlock()

	signalBufferUndoId++
	instance := C.gpointer(recv.native)
	handlerID := C.Buffer_signal_connect_undo(instance, C.gpointer(uintptr(signalBufferUndoId)))

	detail := signalBufferUndoDetail{callback, handlerID}
	signalBufferUndoMap[signalBufferUndoId] = detail

	return signalBufferUndoId
}

/*
DisconnectUndo disconnects a callback from the 'undo' signal for the Buffer.

The connectionID should be a value returned from a call to ConnectUndo.
*/
func (recv *Buffer) DisconnectUndo(connectionID int) {
	signalBufferUndoLock.Lock()
	defer signalBufferUndoLock.Unlock()

	detail, exists := signalBufferUndoMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalBufferUndoMap, connectionID)
}

//export buffer_undoHandler
func buffer_undoHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalBufferUndoLock.RLock()
	defer signalBufferUndoLock.RUnlock()

	targetObject := BufferNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalBufferUndoMap[index].callback
	callback(targetObject)
}

// BufferNew is a wrapper around the C function gtk_source_buffer_new.
func BufferNew(table *gtk.TextTagTable) *Buffer {
	c_table := (*C.GtkTextTagTable)(C.NULL)
	if table != nil {
		c_table = (*C.GtkTextTagTable)(table.ToC())
	}

	retC := C.gtk_source_buffer_new(c_table)
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BufferNewWithLanguage is a wrapper around the C function gtk_source_buffer_new_with_language.
func BufferNewWithLanguage(language *Language) *Buffer {
	c_language := (*C.GtkSourceLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.GtkSourceLanguage)(language.ToC())
	}

	retC := C.gtk_source_buffer_new_with_language(c_language)
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BackwardIterToSourceMark is a wrapper around the C function gtk_source_buffer_backward_iter_to_source_mark.
func (recv *Buffer) BackwardIterToSourceMark(iter *gtk.TextIter, category string) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_buffer_backward_iter_to_source_mark((*C.GtkSourceBuffer)(recv.native), c_iter, c_category)
	retGo := retC == C.TRUE

	return retGo
}

// BeginNotUndoableAction is a wrapper around the C function gtk_source_buffer_begin_not_undoable_action.
func (recv *Buffer) BeginNotUndoableAction() {
	C.gtk_source_buffer_begin_not_undoable_action((*C.GtkSourceBuffer)(recv.native))

	return
}

// CanRedo is a wrapper around the C function gtk_source_buffer_can_redo.
func (recv *Buffer) CanRedo() bool {
	retC := C.gtk_source_buffer_can_redo((*C.GtkSourceBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanUndo is a wrapper around the C function gtk_source_buffer_can_undo.
func (recv *Buffer) CanUndo() bool {
	retC := C.gtk_source_buffer_can_undo((*C.GtkSourceBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ChangeCase is a wrapper around the C function gtk_source_buffer_change_case.
func (recv *Buffer) ChangeCase(caseType ChangeCaseType, start *gtk.TextIter, end *gtk.TextIter) {
	c_case_type := (C.GtkSourceChangeCaseType)(caseType)

	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	C.gtk_source_buffer_change_case((*C.GtkSourceBuffer)(recv.native), c_case_type, c_start, c_end)

	return
}

// CreateSourceMark is a wrapper around the C function gtk_source_buffer_create_source_mark.
func (recv *Buffer) CreateSourceMark(name string, category string, where *gtk.TextIter) *Mark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	c_where := (*C.GtkTextIter)(C.NULL)
	if where != nil {
		c_where = (*C.GtkTextIter)(where.ToC())
	}

	retC := C.gtk_source_buffer_create_source_mark((*C.GtkSourceBuffer)(recv.native), c_name, c_category, c_where)
	retGo := MarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_source_buffer_create_source_tag : unsupported parameter ... : varargs

// EndNotUndoableAction is a wrapper around the C function gtk_source_buffer_end_not_undoable_action.
func (recv *Buffer) EndNotUndoableAction() {
	C.gtk_source_buffer_end_not_undoable_action((*C.GtkSourceBuffer)(recv.native))

	return
}

// EnsureHighlight is a wrapper around the C function gtk_source_buffer_ensure_highlight.
func (recv *Buffer) EnsureHighlight(start *gtk.TextIter, end *gtk.TextIter) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	C.gtk_source_buffer_ensure_highlight((*C.GtkSourceBuffer)(recv.native), c_start, c_end)

	return
}

// ForwardIterToSourceMark is a wrapper around the C function gtk_source_buffer_forward_iter_to_source_mark.
func (recv *Buffer) ForwardIterToSourceMark(iter *gtk.TextIter, category string) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_buffer_forward_iter_to_source_mark((*C.GtkSourceBuffer)(recv.native), c_iter, c_category)
	retGo := retC == C.TRUE

	return retGo
}

// GetContextClassesAtIter is a wrapper around the C function gtk_source_buffer_get_context_classes_at_iter.
func (recv *Buffer) GetContextClassesAtIter(iter *gtk.TextIter) []string {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	retC := C.gtk_source_buffer_get_context_classes_at_iter((*C.GtkSourceBuffer)(recv.native), c_iter)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetHighlightMatchingBrackets is a wrapper around the C function gtk_source_buffer_get_highlight_matching_brackets.
func (recv *Buffer) GetHighlightMatchingBrackets() bool {
	retC := C.gtk_source_buffer_get_highlight_matching_brackets((*C.GtkSourceBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHighlightSyntax is a wrapper around the C function gtk_source_buffer_get_highlight_syntax.
func (recv *Buffer) GetHighlightSyntax() bool {
	retC := C.gtk_source_buffer_get_highlight_syntax((*C.GtkSourceBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetImplicitTrailingNewline is a wrapper around the C function gtk_source_buffer_get_implicit_trailing_newline.
func (recv *Buffer) GetImplicitTrailingNewline() bool {
	retC := C.gtk_source_buffer_get_implicit_trailing_newline((*C.GtkSourceBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLanguage is a wrapper around the C function gtk_source_buffer_get_language.
func (recv *Buffer) GetLanguage() *Language {
	retC := C.gtk_source_buffer_get_language((*C.GtkSourceBuffer)(recv.native))
	var retGo (*Language)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LanguageNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetMaxUndoLevels is a wrapper around the C function gtk_source_buffer_get_max_undo_levels.
func (recv *Buffer) GetMaxUndoLevels() int32 {
	retC := C.gtk_source_buffer_get_max_undo_levels((*C.GtkSourceBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSourceMarksAtIter is a wrapper around the C function gtk_source_buffer_get_source_marks_at_iter.
func (recv *Buffer) GetSourceMarksAtIter(iter *gtk.TextIter, category string) *glib.SList {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_buffer_get_source_marks_at_iter((*C.GtkSourceBuffer)(recv.native), c_iter, c_category)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSourceMarksAtLine is a wrapper around the C function gtk_source_buffer_get_source_marks_at_line.
func (recv *Buffer) GetSourceMarksAtLine(line int32, category string) *glib.SList {
	c_line := (C.gint)(line)

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_buffer_get_source_marks_at_line((*C.GtkSourceBuffer)(recv.native), c_line, c_category)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStyleScheme is a wrapper around the C function gtk_source_buffer_get_style_scheme.
func (recv *Buffer) GetStyleScheme() *StyleScheme {
	retC := C.gtk_source_buffer_get_style_scheme((*C.GtkSourceBuffer)(recv.native))
	var retGo (*StyleScheme)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleSchemeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetUndoManager is a wrapper around the C function gtk_source_buffer_get_undo_manager.
func (recv *Buffer) GetUndoManager() *UndoManager {
	retC := C.gtk_source_buffer_get_undo_manager((*C.GtkSourceBuffer)(recv.native))
	var retGo (*UndoManager)
	if retC == nil {
		retGo = nil
	} else {
		retGo = UndoManagerNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// IterBackwardToContextClassToggle is a wrapper around the C function gtk_source_buffer_iter_backward_to_context_class_toggle.
func (recv *Buffer) IterBackwardToContextClassToggle(iter *gtk.TextIter, contextClass string) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_context_class := C.CString(contextClass)
	defer C.free(unsafe.Pointer(c_context_class))

	retC := C.gtk_source_buffer_iter_backward_to_context_class_toggle((*C.GtkSourceBuffer)(recv.native), c_iter, c_context_class)
	retGo := retC == C.TRUE

	return retGo
}

// IterForwardToContextClassToggle is a wrapper around the C function gtk_source_buffer_iter_forward_to_context_class_toggle.
func (recv *Buffer) IterForwardToContextClassToggle(iter *gtk.TextIter, contextClass string) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_context_class := C.CString(contextClass)
	defer C.free(unsafe.Pointer(c_context_class))

	retC := C.gtk_source_buffer_iter_forward_to_context_class_toggle((*C.GtkSourceBuffer)(recv.native), c_iter, c_context_class)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasContextClass is a wrapper around the C function gtk_source_buffer_iter_has_context_class.
func (recv *Buffer) IterHasContextClass(iter *gtk.TextIter, contextClass string) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_context_class := C.CString(contextClass)
	defer C.free(unsafe.Pointer(c_context_class))

	retC := C.gtk_source_buffer_iter_has_context_class((*C.GtkSourceBuffer)(recv.native), c_iter, c_context_class)
	retGo := retC == C.TRUE

	return retGo
}

// JoinLines is a wrapper around the C function gtk_source_buffer_join_lines.
func (recv *Buffer) JoinLines(start *gtk.TextIter, end *gtk.TextIter) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	C.gtk_source_buffer_join_lines((*C.GtkSourceBuffer)(recv.native), c_start, c_end)

	return
}

// Redo is a wrapper around the C function gtk_source_buffer_redo.
func (recv *Buffer) Redo() {
	C.gtk_source_buffer_redo((*C.GtkSourceBuffer)(recv.native))

	return
}

// RemoveSourceMarks is a wrapper around the C function gtk_source_buffer_remove_source_marks.
func (recv *Buffer) RemoveSourceMarks(start *gtk.TextIter, end *gtk.TextIter, category string) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	C.gtk_source_buffer_remove_source_marks((*C.GtkSourceBuffer)(recv.native), c_start, c_end, c_category)

	return
}

// SetHighlightMatchingBrackets is a wrapper around the C function gtk_source_buffer_set_highlight_matching_brackets.
func (recv *Buffer) SetHighlightMatchingBrackets(highlight bool) {
	c_highlight :=
		boolToGboolean(highlight)

	C.gtk_source_buffer_set_highlight_matching_brackets((*C.GtkSourceBuffer)(recv.native), c_highlight)

	return
}

// SetHighlightSyntax is a wrapper around the C function gtk_source_buffer_set_highlight_syntax.
func (recv *Buffer) SetHighlightSyntax(highlight bool) {
	c_highlight :=
		boolToGboolean(highlight)

	C.gtk_source_buffer_set_highlight_syntax((*C.GtkSourceBuffer)(recv.native), c_highlight)

	return
}

// SetImplicitTrailingNewline is a wrapper around the C function gtk_source_buffer_set_implicit_trailing_newline.
func (recv *Buffer) SetImplicitTrailingNewline(implicitTrailingNewline bool) {
	c_implicit_trailing_newline :=
		boolToGboolean(implicitTrailingNewline)

	C.gtk_source_buffer_set_implicit_trailing_newline((*C.GtkSourceBuffer)(recv.native), c_implicit_trailing_newline)

	return
}

// SetLanguage is a wrapper around the C function gtk_source_buffer_set_language.
func (recv *Buffer) SetLanguage(language *Language) {
	c_language := (*C.GtkSourceLanguage)(C.NULL)
	if language != nil {
		c_language = (*C.GtkSourceLanguage)(language.ToC())
	}

	C.gtk_source_buffer_set_language((*C.GtkSourceBuffer)(recv.native), c_language)

	return
}

// SetMaxUndoLevels is a wrapper around the C function gtk_source_buffer_set_max_undo_levels.
func (recv *Buffer) SetMaxUndoLevels(maxUndoLevels int32) {
	c_max_undo_levels := (C.gint)(maxUndoLevels)

	C.gtk_source_buffer_set_max_undo_levels((*C.GtkSourceBuffer)(recv.native), c_max_undo_levels)

	return
}

// SetStyleScheme is a wrapper around the C function gtk_source_buffer_set_style_scheme.
func (recv *Buffer) SetStyleScheme(scheme *StyleScheme) {
	c_scheme := (*C.GtkSourceStyleScheme)(C.NULL)
	if scheme != nil {
		c_scheme = (*C.GtkSourceStyleScheme)(scheme.ToC())
	}

	C.gtk_source_buffer_set_style_scheme((*C.GtkSourceBuffer)(recv.native), c_scheme)

	return
}

// SetUndoManager is a wrapper around the C function gtk_source_buffer_set_undo_manager.
func (recv *Buffer) SetUndoManager(manager *UndoManager) {
	c_manager := (*C.GtkSourceUndoManager)(manager.ToC())

	C.gtk_source_buffer_set_undo_manager((*C.GtkSourceBuffer)(recv.native), c_manager)

	return
}

// SortLines is a wrapper around the C function gtk_source_buffer_sort_lines.
func (recv *Buffer) SortLines(start *gtk.TextIter, end *gtk.TextIter, flags SortFlags, column int32) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	c_flags := (C.GtkSourceSortFlags)(flags)

	c_column := (C.gint)(column)

	C.gtk_source_buffer_sort_lines((*C.GtkSourceBuffer)(recv.native), c_start, c_end, c_flags, c_column)

	return
}

// Undo is a wrapper around the C function gtk_source_buffer_undo.
func (recv *Buffer) Undo() {
	C.gtk_source_buffer_undo((*C.GtkSourceBuffer)(recv.native))

	return
}

// Completion is a wrapper around the C record GtkSourceCompletion.
type Completion struct {
	native *C.GtkSourceCompletion
	// parent_instance : record
	// priv : record
}

func CompletionNewFromC(u unsafe.Pointer) *Completion {
	c := (*C.GtkSourceCompletion)(u)
	if c == nil {
		return nil
	}

	g := &Completion{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Completion) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Completion) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Completion with another Completion, and returns true if they represent the same GObject.
func (recv *Completion) Equals(other *Completion) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Completion) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Completion.
// Exercise care, as this is a potentially dangerous function if the Object is not a Completion.
func CastToCompletion(object *gobject.Object) *Completion {
	return CompletionNewFromC(object.ToC())
}

type signalCompletionActivateProposalDetail struct {
	callback  CompletionSignalActivateProposalCallback
	handlerID C.gulong
}

var signalCompletionActivateProposalId int
var signalCompletionActivateProposalMap = make(map[int]signalCompletionActivateProposalDetail)
var signalCompletionActivateProposalLock sync.RWMutex

// CompletionSignalActivateProposalCallback is a callback function for a 'activate-proposal' signal emitted from a Completion.
type CompletionSignalActivateProposalCallback func(targetObject *Completion)

/*
ConnectActivateProposal connects the callback to the 'activate-proposal' signal for the Completion.

The returned value represents the connection, and may be passed to DisconnectActivateProposal to remove it.
*/
func (recv *Completion) ConnectActivateProposal(callback CompletionSignalActivateProposalCallback) int {
	signalCompletionActivateProposalLock.Lock()
	defer signalCompletionActivateProposalLock.Unlock()

	signalCompletionActivateProposalId++
	instance := C.gpointer(recv.native)
	handlerID := C.Completion_signal_connect_activate_proposal(instance, C.gpointer(uintptr(signalCompletionActivateProposalId)))

	detail := signalCompletionActivateProposalDetail{callback, handlerID}
	signalCompletionActivateProposalMap[signalCompletionActivateProposalId] = detail

	return signalCompletionActivateProposalId
}

/*
DisconnectActivateProposal disconnects a callback from the 'activate-proposal' signal for the Completion.

The connectionID should be a value returned from a call to ConnectActivateProposal.
*/
func (recv *Completion) DisconnectActivateProposal(connectionID int) {
	signalCompletionActivateProposalLock.Lock()
	defer signalCompletionActivateProposalLock.Unlock()

	detail, exists := signalCompletionActivateProposalMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionActivateProposalMap, connectionID)
}

//export completion_activateProposalHandler
func completion_activateProposalHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCompletionActivateProposalLock.RLock()
	defer signalCompletionActivateProposalLock.RUnlock()

	targetObject := CompletionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionActivateProposalMap[index].callback
	callback(targetObject)
}

type signalCompletionHideDetail struct {
	callback  CompletionSignalHideCallback
	handlerID C.gulong
}

var signalCompletionHideId int
var signalCompletionHideMap = make(map[int]signalCompletionHideDetail)
var signalCompletionHideLock sync.RWMutex

// CompletionSignalHideCallback is a callback function for a 'hide' signal emitted from a Completion.
type CompletionSignalHideCallback func(targetObject *Completion)

/*
ConnectHide connects the callback to the 'hide' signal for the Completion.

The returned value represents the connection, and may be passed to DisconnectHide to remove it.
*/
func (recv *Completion) ConnectHide(callback CompletionSignalHideCallback) int {
	signalCompletionHideLock.Lock()
	defer signalCompletionHideLock.Unlock()

	signalCompletionHideId++
	instance := C.gpointer(recv.native)
	handlerID := C.Completion_signal_connect_hide(instance, C.gpointer(uintptr(signalCompletionHideId)))

	detail := signalCompletionHideDetail{callback, handlerID}
	signalCompletionHideMap[signalCompletionHideId] = detail

	return signalCompletionHideId
}

/*
DisconnectHide disconnects a callback from the 'hide' signal for the Completion.

The connectionID should be a value returned from a call to ConnectHide.
*/
func (recv *Completion) DisconnectHide(connectionID int) {
	signalCompletionHideLock.Lock()
	defer signalCompletionHideLock.Unlock()

	detail, exists := signalCompletionHideMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionHideMap, connectionID)
}

//export completion_hideHandler
func completion_hideHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCompletionHideLock.RLock()
	defer signalCompletionHideLock.RUnlock()

	targetObject := CompletionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionHideMap[index].callback
	callback(targetObject)
}

type signalCompletionMoveCursorDetail struct {
	callback  CompletionSignalMoveCursorCallback
	handlerID C.gulong
}

var signalCompletionMoveCursorId int
var signalCompletionMoveCursorMap = make(map[int]signalCompletionMoveCursorDetail)
var signalCompletionMoveCursorLock sync.RWMutex

// CompletionSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a Completion.
type CompletionSignalMoveCursorCallback func(targetObject *Completion, step gtk.ScrollStep, num int32)

/*
ConnectMoveCursor connects the callback to the 'move-cursor' signal for the Completion.

The returned value represents the connection, and may be passed to DisconnectMoveCursor to remove it.
*/
func (recv *Completion) ConnectMoveCursor(callback CompletionSignalMoveCursorCallback) int {
	signalCompletionMoveCursorLock.Lock()
	defer signalCompletionMoveCursorLock.Unlock()

	signalCompletionMoveCursorId++
	instance := C.gpointer(recv.native)
	handlerID := C.Completion_signal_connect_move_cursor(instance, C.gpointer(uintptr(signalCompletionMoveCursorId)))

	detail := signalCompletionMoveCursorDetail{callback, handlerID}
	signalCompletionMoveCursorMap[signalCompletionMoveCursorId] = detail

	return signalCompletionMoveCursorId
}

/*
DisconnectMoveCursor disconnects a callback from the 'move-cursor' signal for the Completion.

The connectionID should be a value returned from a call to ConnectMoveCursor.
*/
func (recv *Completion) DisconnectMoveCursor(connectionID int) {
	signalCompletionMoveCursorLock.Lock()
	defer signalCompletionMoveCursorLock.Unlock()

	detail, exists := signalCompletionMoveCursorMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionMoveCursorMap, connectionID)
}

//export completion_moveCursorHandler
func completion_moveCursorHandler(c_targetObject *C.GObject, c_step C.GtkScrollStep, c_num C.gint, data C.gpointer) {
	signalCompletionMoveCursorLock.RLock()
	defer signalCompletionMoveCursorLock.RUnlock()

	step := gtk.ScrollStep(c_step)

	num := int32(c_num)

	targetObject := CompletionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionMoveCursorMap[index].callback
	callback(targetObject, step, num)
}

type signalCompletionMovePageDetail struct {
	callback  CompletionSignalMovePageCallback
	handlerID C.gulong
}

var signalCompletionMovePageId int
var signalCompletionMovePageMap = make(map[int]signalCompletionMovePageDetail)
var signalCompletionMovePageLock sync.RWMutex

// CompletionSignalMovePageCallback is a callback function for a 'move-page' signal emitted from a Completion.
type CompletionSignalMovePageCallback func(targetObject *Completion, step gtk.ScrollStep, num int32)

/*
ConnectMovePage connects the callback to the 'move-page' signal for the Completion.

The returned value represents the connection, and may be passed to DisconnectMovePage to remove it.
*/
func (recv *Completion) ConnectMovePage(callback CompletionSignalMovePageCallback) int {
	signalCompletionMovePageLock.Lock()
	defer signalCompletionMovePageLock.Unlock()

	signalCompletionMovePageId++
	instance := C.gpointer(recv.native)
	handlerID := C.Completion_signal_connect_move_page(instance, C.gpointer(uintptr(signalCompletionMovePageId)))

	detail := signalCompletionMovePageDetail{callback, handlerID}
	signalCompletionMovePageMap[signalCompletionMovePageId] = detail

	return signalCompletionMovePageId
}

/*
DisconnectMovePage disconnects a callback from the 'move-page' signal for the Completion.

The connectionID should be a value returned from a call to ConnectMovePage.
*/
func (recv *Completion) DisconnectMovePage(connectionID int) {
	signalCompletionMovePageLock.Lock()
	defer signalCompletionMovePageLock.Unlock()

	detail, exists := signalCompletionMovePageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionMovePageMap, connectionID)
}

//export completion_movePageHandler
func completion_movePageHandler(c_targetObject *C.GObject, c_step C.GtkScrollStep, c_num C.gint, data C.gpointer) {
	signalCompletionMovePageLock.RLock()
	defer signalCompletionMovePageLock.RUnlock()

	step := gtk.ScrollStep(c_step)

	num := int32(c_num)

	targetObject := CompletionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionMovePageMap[index].callback
	callback(targetObject, step, num)
}

type signalCompletionPopulateContextDetail struct {
	callback  CompletionSignalPopulateContextCallback
	handlerID C.gulong
}

var signalCompletionPopulateContextId int
var signalCompletionPopulateContextMap = make(map[int]signalCompletionPopulateContextDetail)
var signalCompletionPopulateContextLock sync.RWMutex

// CompletionSignalPopulateContextCallback is a callback function for a 'populate-context' signal emitted from a Completion.
type CompletionSignalPopulateContextCallback func(targetObject *Completion, context *CompletionContext)

/*
ConnectPopulateContext connects the callback to the 'populate-context' signal for the Completion.

The returned value represents the connection, and may be passed to DisconnectPopulateContext to remove it.
*/
func (recv *Completion) ConnectPopulateContext(callback CompletionSignalPopulateContextCallback) int {
	signalCompletionPopulateContextLock.Lock()
	defer signalCompletionPopulateContextLock.Unlock()

	signalCompletionPopulateContextId++
	instance := C.gpointer(recv.native)
	handlerID := C.Completion_signal_connect_populate_context(instance, C.gpointer(uintptr(signalCompletionPopulateContextId)))

	detail := signalCompletionPopulateContextDetail{callback, handlerID}
	signalCompletionPopulateContextMap[signalCompletionPopulateContextId] = detail

	return signalCompletionPopulateContextId
}

/*
DisconnectPopulateContext disconnects a callback from the 'populate-context' signal for the Completion.

The connectionID should be a value returned from a call to ConnectPopulateContext.
*/
func (recv *Completion) DisconnectPopulateContext(connectionID int) {
	signalCompletionPopulateContextLock.Lock()
	defer signalCompletionPopulateContextLock.Unlock()

	detail, exists := signalCompletionPopulateContextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionPopulateContextMap, connectionID)
}

//export completion_populateContextHandler
func completion_populateContextHandler(c_targetObject *C.GObject, c_context *C.GtkSourceCompletionContext, data C.gpointer) {
	signalCompletionPopulateContextLock.RLock()
	defer signalCompletionPopulateContextLock.RUnlock()

	context := CompletionContextNewFromC(unsafe.Pointer(c_context))

	targetObject := CompletionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionPopulateContextMap[index].callback
	callback(targetObject, context)
}

type signalCompletionShowDetail struct {
	callback  CompletionSignalShowCallback
	handlerID C.gulong
}

var signalCompletionShowId int
var signalCompletionShowMap = make(map[int]signalCompletionShowDetail)
var signalCompletionShowLock sync.RWMutex

// CompletionSignalShowCallback is a callback function for a 'show' signal emitted from a Completion.
type CompletionSignalShowCallback func(targetObject *Completion)

/*
ConnectShow connects the callback to the 'show' signal for the Completion.

The returned value represents the connection, and may be passed to DisconnectShow to remove it.
*/
func (recv *Completion) ConnectShow(callback CompletionSignalShowCallback) int {
	signalCompletionShowLock.Lock()
	defer signalCompletionShowLock.Unlock()

	signalCompletionShowId++
	instance := C.gpointer(recv.native)
	handlerID := C.Completion_signal_connect_show(instance, C.gpointer(uintptr(signalCompletionShowId)))

	detail := signalCompletionShowDetail{callback, handlerID}
	signalCompletionShowMap[signalCompletionShowId] = detail

	return signalCompletionShowId
}

/*
DisconnectShow disconnects a callback from the 'show' signal for the Completion.

The connectionID should be a value returned from a call to ConnectShow.
*/
func (recv *Completion) DisconnectShow(connectionID int) {
	signalCompletionShowLock.Lock()
	defer signalCompletionShowLock.Unlock()

	detail, exists := signalCompletionShowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionShowMap, connectionID)
}

//export completion_showHandler
func completion_showHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCompletionShowLock.RLock()
	defer signalCompletionShowLock.RUnlock()

	targetObject := CompletionNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionShowMap[index].callback
	callback(targetObject)
}

// AddProvider is a wrapper around the C function gtk_source_completion_add_provider.
func (recv *Completion) AddProvider(provider *CompletionProvider) (bool, error) {
	c_provider := (*C.GtkSourceCompletionProvider)(provider.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_source_completion_add_provider((*C.GtkSourceCompletion)(recv.native), c_provider, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// BlockInteractive is a wrapper around the C function gtk_source_completion_block_interactive.
func (recv *Completion) BlockInteractive() {
	C.gtk_source_completion_block_interactive((*C.GtkSourceCompletion)(recv.native))

	return
}

// CreateContext is a wrapper around the C function gtk_source_completion_create_context.
func (recv *Completion) CreateContext(position *gtk.TextIter) *CompletionContext {
	c_position := (*C.GtkTextIter)(C.NULL)
	if position != nil {
		c_position = (*C.GtkTextIter)(position.ToC())
	}

	retC := C.gtk_source_completion_create_context((*C.GtkSourceCompletion)(recv.native), c_position)
	retGo := CompletionContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInfoWindow is a wrapper around the C function gtk_source_completion_get_info_window.
func (recv *Completion) GetInfoWindow() *CompletionInfo {
	retC := C.gtk_source_completion_get_info_window((*C.GtkSourceCompletion)(recv.native))
	retGo := CompletionInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProviders is a wrapper around the C function gtk_source_completion_get_providers.
func (recv *Completion) GetProviders() *glib.List {
	retC := C.gtk_source_completion_get_providers((*C.GtkSourceCompletion)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetView is a wrapper around the C function gtk_source_completion_get_view.
func (recv *Completion) GetView() *View {
	retC := C.gtk_source_completion_get_view((*C.GtkSourceCompletion)(recv.native))
	var retGo (*View)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ViewNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Hide is a wrapper around the C function gtk_source_completion_hide.
func (recv *Completion) Hide() {
	C.gtk_source_completion_hide((*C.GtkSourceCompletion)(recv.native))

	return
}

// MoveWindow is a wrapper around the C function gtk_source_completion_move_window.
func (recv *Completion) MoveWindow(iter *gtk.TextIter) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	C.gtk_source_completion_move_window((*C.GtkSourceCompletion)(recv.native), c_iter)

	return
}

// RemoveProvider is a wrapper around the C function gtk_source_completion_remove_provider.
func (recv *Completion) RemoveProvider(provider *CompletionProvider) (bool, error) {
	c_provider := (*C.GtkSourceCompletionProvider)(provider.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_source_completion_remove_provider((*C.GtkSourceCompletion)(recv.native), c_provider, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Show is a wrapper around the C function gtk_source_completion_show.
func (recv *Completion) Show(providers *glib.List, context *CompletionContext) bool {
	c_providers := (*C.GList)(C.NULL)
	if providers != nil {
		c_providers = (*C.GList)(providers.ToC())
	}

	c_context := (*C.GtkSourceCompletionContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkSourceCompletionContext)(context.ToC())
	}

	retC := C.gtk_source_completion_show((*C.GtkSourceCompletion)(recv.native), c_providers, c_context)
	retGo := retC == C.TRUE

	return retGo
}

// UnblockInteractive is a wrapper around the C function gtk_source_completion_unblock_interactive.
func (recv *Completion) UnblockInteractive() {
	C.gtk_source_completion_unblock_interactive((*C.GtkSourceCompletion)(recv.native))

	return
}

// Buildable returns the Buildable interface implemented by Completion
func (recv *Completion) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// CompletionContext is a wrapper around the C record GtkSourceCompletionContext.
type CompletionContext struct {
	native *C.GtkSourceCompletionContext
	// parent : record
	// priv : record
}

func CompletionContextNewFromC(u unsafe.Pointer) *CompletionContext {
	c := (*C.GtkSourceCompletionContext)(u)
	if c == nil {
		return nil
	}

	g := &CompletionContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CompletionContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CompletionContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionContext with another CompletionContext, and returns true if they represent the same GObject.
func (recv *CompletionContext) Equals(other *CompletionContext) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CompletionContext) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *CompletionContext) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitrary Object to CompletionContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a CompletionContext.
func CastToCompletionContext(object *gobject.Object) *CompletionContext {
	return CompletionContextNewFromC(object.ToC())
}

type signalCompletionContextCancelledDetail struct {
	callback  CompletionContextSignalCancelledCallback
	handlerID C.gulong
}

var signalCompletionContextCancelledId int
var signalCompletionContextCancelledMap = make(map[int]signalCompletionContextCancelledDetail)
var signalCompletionContextCancelledLock sync.RWMutex

// CompletionContextSignalCancelledCallback is a callback function for a 'cancelled' signal emitted from a CompletionContext.
type CompletionContextSignalCancelledCallback func(targetObject *CompletionContext)

/*
ConnectCancelled connects the callback to the 'cancelled' signal for the CompletionContext.

The returned value represents the connection, and may be passed to DisconnectCancelled to remove it.
*/
func (recv *CompletionContext) ConnectCancelled(callback CompletionContextSignalCancelledCallback) int {
	signalCompletionContextCancelledLock.Lock()
	defer signalCompletionContextCancelledLock.Unlock()

	signalCompletionContextCancelledId++
	instance := C.gpointer(recv.native)
	handlerID := C.CompletionContext_signal_connect_cancelled(instance, C.gpointer(uintptr(signalCompletionContextCancelledId)))

	detail := signalCompletionContextCancelledDetail{callback, handlerID}
	signalCompletionContextCancelledMap[signalCompletionContextCancelledId] = detail

	return signalCompletionContextCancelledId
}

/*
DisconnectCancelled disconnects a callback from the 'cancelled' signal for the CompletionContext.

The connectionID should be a value returned from a call to ConnectCancelled.
*/
func (recv *CompletionContext) DisconnectCancelled(connectionID int) {
	signalCompletionContextCancelledLock.Lock()
	defer signalCompletionContextCancelledLock.Unlock()

	detail, exists := signalCompletionContextCancelledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionContextCancelledMap, connectionID)
}

//export completioncontext_cancelledHandler
func completioncontext_cancelledHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCompletionContextCancelledLock.RLock()
	defer signalCompletionContextCancelledLock.RUnlock()

	targetObject := CompletionContextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionContextCancelledMap[index].callback
	callback(targetObject)
}

// AddProposals is a wrapper around the C function gtk_source_completion_context_add_proposals.
func (recv *CompletionContext) AddProposals(provider *CompletionProvider, proposals *glib.List, finished bool) {
	c_provider := (*C.GtkSourceCompletionProvider)(provider.ToC())

	c_proposals := (*C.GList)(C.NULL)
	if proposals != nil {
		c_proposals = (*C.GList)(proposals.ToC())
	}

	c_finished :=
		boolToGboolean(finished)

	C.gtk_source_completion_context_add_proposals((*C.GtkSourceCompletionContext)(recv.native), c_provider, c_proposals, c_finished)

	return
}

// GetActivation is a wrapper around the C function gtk_source_completion_context_get_activation.
func (recv *CompletionContext) GetActivation() CompletionActivation {
	retC := C.gtk_source_completion_context_get_activation((*C.GtkSourceCompletionContext)(recv.native))
	retGo := (CompletionActivation)(retC)

	return retGo
}

// GetIter is a wrapper around the C function gtk_source_completion_context_get_iter.
func (recv *CompletionContext) GetIter() (bool, *gtk.TextIter) {
	var c_iter C.GtkTextIter

	retC := C.gtk_source_completion_context_get_iter((*C.GtkSourceCompletionContext)(recv.native), &c_iter)
	retGo := retC == C.TRUE

	iter := gtk.TextIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// CompletionInfo is a wrapper around the C record GtkSourceCompletionInfo.
type CompletionInfo struct {
	native *C.GtkSourceCompletionInfo
	// parent : record
	// priv : record
}

func CompletionInfoNewFromC(u unsafe.Pointer) *CompletionInfo {
	c := (*C.GtkSourceCompletionInfo)(u)
	if c == nil {
		return nil
	}

	g := &CompletionInfo{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CompletionInfo) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CompletionInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionInfo with another CompletionInfo, and returns true if they represent the same GObject.
func (recv *CompletionInfo) Equals(other *CompletionInfo) bool {
	return other.ToC() == recv.ToC()
}

// Window upcasts to *Window
func (recv *CompletionInfo) Window() *gtk.Window {
	return gtk.WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *CompletionInfo) Bin() *gtk.Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *CompletionInfo) Container() *gtk.Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *CompletionInfo) Widget() *gtk.Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CompletionInfo) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CompletionInfo) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitrary Object to CompletionInfo.
// Exercise care, as this is a potentially dangerous function if the Object is not a CompletionInfo.
func CastToCompletionInfo(object *gobject.Object) *CompletionInfo {
	return CompletionInfoNewFromC(object.ToC())
}

type signalCompletionInfoBeforeShowDetail struct {
	callback  CompletionInfoSignalBeforeShowCallback
	handlerID C.gulong
}

var signalCompletionInfoBeforeShowId int
var signalCompletionInfoBeforeShowMap = make(map[int]signalCompletionInfoBeforeShowDetail)
var signalCompletionInfoBeforeShowLock sync.RWMutex

// CompletionInfoSignalBeforeShowCallback is a callback function for a 'before-show' signal emitted from a CompletionInfo.
type CompletionInfoSignalBeforeShowCallback func(targetObject *CompletionInfo)

/*
ConnectBeforeShow connects the callback to the 'before-show' signal for the CompletionInfo.

The returned value represents the connection, and may be passed to DisconnectBeforeShow to remove it.
*/
func (recv *CompletionInfo) ConnectBeforeShow(callback CompletionInfoSignalBeforeShowCallback) int {
	signalCompletionInfoBeforeShowLock.Lock()
	defer signalCompletionInfoBeforeShowLock.Unlock()

	signalCompletionInfoBeforeShowId++
	instance := C.gpointer(recv.native)
	handlerID := C.CompletionInfo_signal_connect_before_show(instance, C.gpointer(uintptr(signalCompletionInfoBeforeShowId)))

	detail := signalCompletionInfoBeforeShowDetail{callback, handlerID}
	signalCompletionInfoBeforeShowMap[signalCompletionInfoBeforeShowId] = detail

	return signalCompletionInfoBeforeShowId
}

/*
DisconnectBeforeShow disconnects a callback from the 'before-show' signal for the CompletionInfo.

The connectionID should be a value returned from a call to ConnectBeforeShow.
*/
func (recv *CompletionInfo) DisconnectBeforeShow(connectionID int) {
	signalCompletionInfoBeforeShowLock.Lock()
	defer signalCompletionInfoBeforeShowLock.Unlock()

	detail, exists := signalCompletionInfoBeforeShowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionInfoBeforeShowMap, connectionID)
}

//export completioninfo_beforeShowHandler
func completioninfo_beforeShowHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCompletionInfoBeforeShowLock.RLock()
	defer signalCompletionInfoBeforeShowLock.RUnlock()

	targetObject := CompletionInfoNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionInfoBeforeShowMap[index].callback
	callback(targetObject)
}

// CompletionInfoNew is a wrapper around the C function gtk_source_completion_info_new.
func CompletionInfoNew() *CompletionInfo {
	retC := C.gtk_source_completion_info_new()
	retGo := CompletionInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidget is a wrapper around the C function gtk_source_completion_info_get_widget.
func (recv *CompletionInfo) GetWidget() *gtk.Widget {
	retC := C.gtk_source_completion_info_get_widget((*C.GtkSourceCompletionInfo)(recv.native))
	retGo := gtk.WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MoveToIter is a wrapper around the C function gtk_source_completion_info_move_to_iter.
func (recv *CompletionInfo) MoveToIter(view *gtk.TextView, iter *gtk.TextIter) {
	c_view := (*C.GtkTextView)(C.NULL)
	if view != nil {
		c_view = (*C.GtkTextView)(view.ToC())
	}

	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	C.gtk_source_completion_info_move_to_iter((*C.GtkSourceCompletionInfo)(recv.native), c_view, c_iter)

	return
}

// SetWidget is a wrapper around the C function gtk_source_completion_info_set_widget.
func (recv *CompletionInfo) SetWidget(widget *gtk.Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_source_completion_info_set_widget((*C.GtkSourceCompletionInfo)(recv.native), c_widget)

	return
}

// ImplementorIface returns the ImplementorIface interface implemented by CompletionInfo
func (recv *CompletionInfo) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CompletionInfo
func (recv *CompletionInfo) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// CompletionItem is a wrapper around the C record GtkSourceCompletionItem.
type CompletionItem struct {
	native *C.GtkSourceCompletionItem
	// parent : record
	// priv : record
}

func CompletionItemNewFromC(u unsafe.Pointer) *CompletionItem {
	c := (*C.GtkSourceCompletionItem)(u)
	if c == nil {
		return nil
	}

	g := &CompletionItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CompletionItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CompletionItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionItem with another CompletionItem, and returns true if they represent the same GObject.
func (recv *CompletionItem) Equals(other *CompletionItem) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *CompletionItem) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to CompletionItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a CompletionItem.
func CastToCompletionItem(object *gobject.Object) *CompletionItem {
	return CompletionItemNewFromC(object.ToC())
}

// CompletionItemNew is a wrapper around the C function gtk_source_completion_item_new.
func CompletionItemNew(label string, text string, icon *gdkpixbuf.Pixbuf, info string) *CompletionItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_icon := (*C.GdkPixbuf)(C.NULL)
	if icon != nil {
		c_icon = (*C.GdkPixbuf)(icon.ToC())
	}

	c_info := C.CString(info)
	defer C.free(unsafe.Pointer(c_info))

	retC := C.gtk_source_completion_item_new(c_label, c_text, c_icon, c_info)
	retGo := CompletionItemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CompletionItemNewFromStock is a wrapper around the C function gtk_source_completion_item_new_from_stock.
func CompletionItemNewFromStock(label string, text string, stock string, info string) *CompletionItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_stock := C.CString(stock)
	defer C.free(unsafe.Pointer(c_stock))

	c_info := C.CString(info)
	defer C.free(unsafe.Pointer(c_info))

	retC := C.gtk_source_completion_item_new_from_stock(c_label, c_text, c_stock, c_info)
	retGo := CompletionItemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CompletionItemNewWithMarkup is a wrapper around the C function gtk_source_completion_item_new_with_markup.
func CompletionItemNewWithMarkup(markup string, text string, icon *gdkpixbuf.Pixbuf, info string) *CompletionItem {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_icon := (*C.GdkPixbuf)(C.NULL)
	if icon != nil {
		c_icon = (*C.GdkPixbuf)(icon.ToC())
	}

	c_info := C.CString(info)
	defer C.free(unsafe.Pointer(c_info))

	retC := C.gtk_source_completion_item_new_with_markup(c_markup, c_text, c_icon, c_info)
	retGo := CompletionItemNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CompletionProposal returns the CompletionProposal interface implemented by CompletionItem
func (recv *CompletionItem) CompletionProposal() *CompletionProposal {
	return CompletionProposalNewFromC(recv.ToC())
}

// CompletionWords is a wrapper around the C record GtkSourceCompletionWords.
type CompletionWords struct {
	native *C.GtkSourceCompletionWords
	// parent : record
	// priv : record
}

func CompletionWordsNewFromC(u unsafe.Pointer) *CompletionWords {
	c := (*C.GtkSourceCompletionWords)(u)
	if c == nil {
		return nil
	}

	g := &CompletionWords{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CompletionWords) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CompletionWords) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionWords with another CompletionWords, and returns true if they represent the same GObject.
func (recv *CompletionWords) Equals(other *CompletionWords) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *CompletionWords) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to CompletionWords.
// Exercise care, as this is a potentially dangerous function if the Object is not a CompletionWords.
func CastToCompletionWords(object *gobject.Object) *CompletionWords {
	return CompletionWordsNewFromC(object.ToC())
}

// CompletionWordsNew is a wrapper around the C function gtk_source_completion_words_new.
func CompletionWordsNew(name string, icon *gdkpixbuf.Pixbuf) *CompletionWords {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_icon := (*C.GdkPixbuf)(C.NULL)
	if icon != nil {
		c_icon = (*C.GdkPixbuf)(icon.ToC())
	}

	retC := C.gtk_source_completion_words_new(c_name, c_icon)
	retGo := CompletionWordsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Register is a wrapper around the C function gtk_source_completion_words_register.
func (recv *CompletionWords) Register(buffer *gtk.TextBuffer) {
	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	C.gtk_source_completion_words_register((*C.GtkSourceCompletionWords)(recv.native), c_buffer)

	return
}

// Unregister is a wrapper around the C function gtk_source_completion_words_unregister.
func (recv *CompletionWords) Unregister(buffer *gtk.TextBuffer) {
	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	C.gtk_source_completion_words_unregister((*C.GtkSourceCompletionWords)(recv.native), c_buffer)

	return
}

// CompletionProvider returns the CompletionProvider interface implemented by CompletionWords
func (recv *CompletionWords) CompletionProvider() *CompletionProvider {
	return CompletionProviderNewFromC(recv.ToC())
}

// File is a wrapper around the C record GtkSourceFile.
type File struct {
	native *C.GtkSourceFile
	// parent : record
	// priv : record
}

func FileNewFromC(u unsafe.Pointer) *File {
	c := (*C.GtkSourceFile)(u)
	if c == nil {
		return nil
	}

	g := &File{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *File) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this File with another File, and returns true if they represent the same GObject.
func (recv *File) Equals(other *File) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *File) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to File.
// Exercise care, as this is a potentially dangerous function if the Object is not a File.
func CastToFile(object *gobject.Object) *File {
	return FileNewFromC(object.ToC())
}

// FileNew is a wrapper around the C function gtk_source_file_new.
func FileNew() *File {
	retC := C.gtk_source_file_new()
	retGo := FileNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// CheckFileOnDisk is a wrapper around the C function gtk_source_file_check_file_on_disk.
func (recv *File) CheckFileOnDisk() {
	C.gtk_source_file_check_file_on_disk((*C.GtkSourceFile)(recv.native))

	return
}

// GetCompressionType is a wrapper around the C function gtk_source_file_get_compression_type.
func (recv *File) GetCompressionType() CompressionType {
	retC := C.gtk_source_file_get_compression_type((*C.GtkSourceFile)(recv.native))
	retGo := (CompressionType)(retC)

	return retGo
}

// GetEncoding is a wrapper around the C function gtk_source_file_get_encoding.
func (recv *File) GetEncoding() *Encoding {
	retC := C.gtk_source_file_get_encoding((*C.GtkSourceFile)(recv.native))
	retGo := EncodingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLocation is a wrapper around the C function gtk_source_file_get_location.
func (recv *File) GetLocation() *gio.File {
	retC := C.gtk_source_file_get_location((*C.GtkSourceFile)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNewlineType is a wrapper around the C function gtk_source_file_get_newline_type.
func (recv *File) GetNewlineType() NewlineType {
	retC := C.gtk_source_file_get_newline_type((*C.GtkSourceFile)(recv.native))
	retGo := (NewlineType)(retC)

	return retGo
}

// IsDeleted is a wrapper around the C function gtk_source_file_is_deleted.
func (recv *File) IsDeleted() bool {
	retC := C.gtk_source_file_is_deleted((*C.GtkSourceFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsExternallyModified is a wrapper around the C function gtk_source_file_is_externally_modified.
func (recv *File) IsExternallyModified() bool {
	retC := C.gtk_source_file_is_externally_modified((*C.GtkSourceFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsLocal is a wrapper around the C function gtk_source_file_is_local.
func (recv *File) IsLocal() bool {
	retC := C.gtk_source_file_is_local((*C.GtkSourceFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsReadonly is a wrapper around the C function gtk_source_file_is_readonly.
func (recv *File) IsReadonly() bool {
	retC := C.gtk_source_file_is_readonly((*C.GtkSourceFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLocation is a wrapper around the C function gtk_source_file_set_location.
func (recv *File) SetLocation(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_source_file_set_location((*C.GtkSourceFile)(recv.native), c_location)

	return
}

// Unsupported : gtk_source_file_set_mount_operation_factory : unsupported parameter callback : no type generator for MountOperationFactory (GtkSourceMountOperationFactory) for param callback

// FileLoader is a wrapper around the C record GtkSourceFileLoader.
type FileLoader struct {
	native *C.GtkSourceFileLoader
	// parent : record
	// priv : record
}

func FileLoaderNewFromC(u unsafe.Pointer) *FileLoader {
	c := (*C.GtkSourceFileLoader)(u)
	if c == nil {
		return nil
	}

	g := &FileLoader{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileLoader) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileLoader) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileLoader with another FileLoader, and returns true if they represent the same GObject.
func (recv *FileLoader) Equals(other *FileLoader) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileLoader) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileLoader.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileLoader.
func CastToFileLoader(object *gobject.Object) *FileLoader {
	return FileLoaderNewFromC(object.ToC())
}

// FileLoaderNew is a wrapper around the C function gtk_source_file_loader_new.
func FileLoaderNew(buffer *Buffer, file *File) *FileLoader {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	c_file := (*C.GtkSourceFile)(C.NULL)
	if file != nil {
		c_file = (*C.GtkSourceFile)(file.ToC())
	}

	retC := C.gtk_source_file_loader_new(c_buffer, c_file)
	retGo := FileLoaderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FileLoaderNewFromStream is a wrapper around the C function gtk_source_file_loader_new_from_stream.
func FileLoaderNewFromStream(buffer *Buffer, file *File, stream *gio.InputStream) *FileLoader {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	c_file := (*C.GtkSourceFile)(C.NULL)
	if file != nil {
		c_file = (*C.GtkSourceFile)(file.ToC())
	}

	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	retC := C.gtk_source_file_loader_new_from_stream(c_buffer, c_file, c_stream)
	retGo := FileLoaderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_source_file_loader_get_buffer.
func (recv *FileLoader) GetBuffer() *Buffer {
	retC := C.gtk_source_file_loader_get_buffer((*C.GtkSourceFileLoader)(recv.native))
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCompressionType is a wrapper around the C function gtk_source_file_loader_get_compression_type.
func (recv *FileLoader) GetCompressionType() CompressionType {
	retC := C.gtk_source_file_loader_get_compression_type((*C.GtkSourceFileLoader)(recv.native))
	retGo := (CompressionType)(retC)

	return retGo
}

// GetEncoding is a wrapper around the C function gtk_source_file_loader_get_encoding.
func (recv *FileLoader) GetEncoding() *Encoding {
	retC := C.gtk_source_file_loader_get_encoding((*C.GtkSourceFileLoader)(recv.native))
	retGo := EncodingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFile is a wrapper around the C function gtk_source_file_loader_get_file.
func (recv *FileLoader) GetFile() *File {
	retC := C.gtk_source_file_loader_get_file((*C.GtkSourceFileLoader)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInputStream is a wrapper around the C function gtk_source_file_loader_get_input_stream.
func (recv *FileLoader) GetInputStream() *gio.InputStream {
	retC := C.gtk_source_file_loader_get_input_stream((*C.GtkSourceFileLoader)(recv.native))
	var retGo (*gio.InputStream)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.InputStreamNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLocation is a wrapper around the C function gtk_source_file_loader_get_location.
func (recv *FileLoader) GetLocation() *gio.File {
	retC := C.gtk_source_file_loader_get_location((*C.GtkSourceFileLoader)(recv.native))
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetNewlineType is a wrapper around the C function gtk_source_file_loader_get_newline_type.
func (recv *FileLoader) GetNewlineType() NewlineType {
	retC := C.gtk_source_file_loader_get_newline_type((*C.GtkSourceFileLoader)(recv.native))
	retGo := (NewlineType)(retC)

	return retGo
}

// Unsupported : gtk_source_file_loader_load_async : unsupported parameter progress_callback : no type generator for Gio.FileProgressCallback (GFileProgressCallback) for param progress_callback

// LoadFinish is a wrapper around the C function gtk_source_file_loader_load_finish.
func (recv *FileLoader) LoadFinish(result *gio.AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_source_file_loader_load_finish((*C.GtkSourceFileLoader)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetCandidateEncodings is a wrapper around the C function gtk_source_file_loader_set_candidate_encodings.
func (recv *FileLoader) SetCandidateEncodings(candidateEncodings *glib.SList) {
	c_candidate_encodings := (*C.GSList)(C.NULL)
	if candidateEncodings != nil {
		c_candidate_encodings = (*C.GSList)(candidateEncodings.ToC())
	}

	C.gtk_source_file_loader_set_candidate_encodings((*C.GtkSourceFileLoader)(recv.native), c_candidate_encodings)

	return
}

// FileSaver is a wrapper around the C record GtkSourceFileSaver.
type FileSaver struct {
	native *C.GtkSourceFileSaver
	// object : record
	// priv : record
}

func FileSaverNewFromC(u unsafe.Pointer) *FileSaver {
	c := (*C.GtkSourceFileSaver)(u)
	if c == nil {
		return nil
	}

	g := &FileSaver{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileSaver) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileSaver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileSaver with another FileSaver, and returns true if they represent the same GObject.
func (recv *FileSaver) Equals(other *FileSaver) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileSaver) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileSaver.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileSaver.
func CastToFileSaver(object *gobject.Object) *FileSaver {
	return FileSaverNewFromC(object.ToC())
}

// FileSaverNew is a wrapper around the C function gtk_source_file_saver_new.
func FileSaverNew(buffer *Buffer, file *File) *FileSaver {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	c_file := (*C.GtkSourceFile)(C.NULL)
	if file != nil {
		c_file = (*C.GtkSourceFile)(file.ToC())
	}

	retC := C.gtk_source_file_saver_new(c_buffer, c_file)
	retGo := FileSaverNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// FileSaverNewWithTarget is a wrapper around the C function gtk_source_file_saver_new_with_target.
func FileSaverNewWithTarget(buffer *Buffer, file *File, targetLocation *gio.File) *FileSaver {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	c_file := (*C.GtkSourceFile)(C.NULL)
	if file != nil {
		c_file = (*C.GtkSourceFile)(file.ToC())
	}

	c_target_location := (*C.GFile)(targetLocation.ToC())

	retC := C.gtk_source_file_saver_new_with_target(c_buffer, c_file, c_target_location)
	retGo := FileSaverNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_source_file_saver_get_buffer.
func (recv *FileSaver) GetBuffer() *Buffer {
	retC := C.gtk_source_file_saver_get_buffer((*C.GtkSourceFileSaver)(recv.native))
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCompressionType is a wrapper around the C function gtk_source_file_saver_get_compression_type.
func (recv *FileSaver) GetCompressionType() CompressionType {
	retC := C.gtk_source_file_saver_get_compression_type((*C.GtkSourceFileSaver)(recv.native))
	retGo := (CompressionType)(retC)

	return retGo
}

// GetEncoding is a wrapper around the C function gtk_source_file_saver_get_encoding.
func (recv *FileSaver) GetEncoding() *Encoding {
	retC := C.gtk_source_file_saver_get_encoding((*C.GtkSourceFileSaver)(recv.native))
	retGo := EncodingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFile is a wrapper around the C function gtk_source_file_saver_get_file.
func (recv *FileSaver) GetFile() *File {
	retC := C.gtk_source_file_saver_get_file((*C.GtkSourceFileSaver)(recv.native))
	retGo := FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFlags is a wrapper around the C function gtk_source_file_saver_get_flags.
func (recv *FileSaver) GetFlags() FileSaverFlags {
	retC := C.gtk_source_file_saver_get_flags((*C.GtkSourceFileSaver)(recv.native))
	retGo := (FileSaverFlags)(retC)

	return retGo
}

// GetLocation is a wrapper around the C function gtk_source_file_saver_get_location.
func (recv *FileSaver) GetLocation() *gio.File {
	retC := C.gtk_source_file_saver_get_location((*C.GtkSourceFileSaver)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNewlineType is a wrapper around the C function gtk_source_file_saver_get_newline_type.
func (recv *FileSaver) GetNewlineType() NewlineType {
	retC := C.gtk_source_file_saver_get_newline_type((*C.GtkSourceFileSaver)(recv.native))
	retGo := (NewlineType)(retC)

	return retGo
}

// Unsupported : gtk_source_file_saver_save_async : unsupported parameter progress_callback : no type generator for Gio.FileProgressCallback (GFileProgressCallback) for param progress_callback

// SaveFinish is a wrapper around the C function gtk_source_file_saver_save_finish.
func (recv *FileSaver) SaveFinish(result *gio.AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_source_file_saver_save_finish((*C.GtkSourceFileSaver)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetCompressionType is a wrapper around the C function gtk_source_file_saver_set_compression_type.
func (recv *FileSaver) SetCompressionType(compressionType CompressionType) {
	c_compression_type := (C.GtkSourceCompressionType)(compressionType)

	C.gtk_source_file_saver_set_compression_type((*C.GtkSourceFileSaver)(recv.native), c_compression_type)

	return
}

// SetEncoding is a wrapper around the C function gtk_source_file_saver_set_encoding.
func (recv *FileSaver) SetEncoding(encoding *Encoding) {
	c_encoding := (*C.GtkSourceEncoding)(C.NULL)
	if encoding != nil {
		c_encoding = (*C.GtkSourceEncoding)(encoding.ToC())
	}

	C.gtk_source_file_saver_set_encoding((*C.GtkSourceFileSaver)(recv.native), c_encoding)

	return
}

// SetFlags is a wrapper around the C function gtk_source_file_saver_set_flags.
func (recv *FileSaver) SetFlags(flags FileSaverFlags) {
	c_flags := (C.GtkSourceFileSaverFlags)(flags)

	C.gtk_source_file_saver_set_flags((*C.GtkSourceFileSaver)(recv.native), c_flags)

	return
}

// SetNewlineType is a wrapper around the C function gtk_source_file_saver_set_newline_type.
func (recv *FileSaver) SetNewlineType(newlineType NewlineType) {
	c_newline_type := (C.GtkSourceNewlineType)(newlineType)

	C.gtk_source_file_saver_set_newline_type((*C.GtkSourceFileSaver)(recv.native), c_newline_type)

	return
}

// Gutter is a wrapper around the C record GtkSourceGutter.
type Gutter struct {
	native *C.GtkSourceGutter
	// parent : record
	// priv : record
}

func GutterNewFromC(u unsafe.Pointer) *Gutter {
	c := (*C.GtkSourceGutter)(u)
	if c == nil {
		return nil
	}

	g := &Gutter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Gutter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Gutter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Gutter with another Gutter, and returns true if they represent the same GObject.
func (recv *Gutter) Equals(other *Gutter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Gutter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Gutter.
// Exercise care, as this is a potentially dangerous function if the Object is not a Gutter.
func CastToGutter(object *gobject.Object) *Gutter {
	return GutterNewFromC(object.ToC())
}

// GetPadding is a wrapper around the C function gtk_source_gutter_get_padding.
func (recv *Gutter) GetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_source_gutter_get_padding((*C.GtkSourceGutter)(recv.native), &c_xpad, &c_ypad)

	return
}

// GetRendererAtPos is a wrapper around the C function gtk_source_gutter_get_renderer_at_pos.
func (recv *Gutter) GetRendererAtPos(x int32, y int32) *GutterRenderer {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_source_gutter_get_renderer_at_pos((*C.GtkSourceGutter)(recv.native), c_x, c_y)
	var retGo (*GutterRenderer)
	if retC == nil {
		retGo = nil
	} else {
		retGo = GutterRendererNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetWindow is a wrapper around the C function gtk_source_gutter_get_window.
func (recv *Gutter) GetWindow() *gdk.Window {
	retC := C.gtk_source_gutter_get_window((*C.GtkSourceGutter)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function gtk_source_gutter_insert.
func (recv *Gutter) Insert(renderer *GutterRenderer, position int32) bool {
	c_renderer := (*C.GtkSourceGutterRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkSourceGutterRenderer)(renderer.ToC())
	}

	c_position := (C.gint)(position)

	retC := C.gtk_source_gutter_insert((*C.GtkSourceGutter)(recv.native), c_renderer, c_position)
	retGo := retC == C.TRUE

	return retGo
}

// QueueDraw is a wrapper around the C function gtk_source_gutter_queue_draw.
func (recv *Gutter) QueueDraw() {
	C.gtk_source_gutter_queue_draw((*C.GtkSourceGutter)(recv.native))

	return
}

// Remove is a wrapper around the C function gtk_source_gutter_remove.
func (recv *Gutter) Remove(renderer *GutterRenderer) {
	c_renderer := (*C.GtkSourceGutterRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkSourceGutterRenderer)(renderer.ToC())
	}

	C.gtk_source_gutter_remove((*C.GtkSourceGutter)(recv.native), c_renderer)

	return
}

// Reorder is a wrapper around the C function gtk_source_gutter_reorder.
func (recv *Gutter) Reorder(renderer *GutterRenderer, position int32) {
	c_renderer := (*C.GtkSourceGutterRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkSourceGutterRenderer)(renderer.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_source_gutter_reorder((*C.GtkSourceGutter)(recv.native), c_renderer, c_position)

	return
}

// SetPadding is a wrapper around the C function gtk_source_gutter_set_padding.
func (recv *Gutter) SetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_source_gutter_set_padding((*C.GtkSourceGutter)(recv.native), c_xpad, c_ypad)

	return
}

// GutterRenderer is a wrapper around the C record GtkSourceGutterRenderer.
type GutterRenderer struct {
	native *C.GtkSourceGutterRenderer
	// parent : record
	// Private : priv
}

func GutterRendererNewFromC(u unsafe.Pointer) *GutterRenderer {
	c := (*C.GtkSourceGutterRenderer)(u)
	if c == nil {
		return nil
	}

	g := &GutterRenderer{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GutterRenderer) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GutterRenderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRenderer with another GutterRenderer, and returns true if they represent the same GObject.
func (recv *GutterRenderer) Equals(other *GutterRenderer) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GutterRenderer) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *GutterRenderer) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitrary Object to GutterRenderer.
// Exercise care, as this is a potentially dangerous function if the Object is not a GutterRenderer.
func CastToGutterRenderer(object *gobject.Object) *GutterRenderer {
	return GutterRendererNewFromC(object.ToC())
}

type signalGutterRendererActivateDetail struct {
	callback  GutterRendererSignalActivateCallback
	handlerID C.gulong
}

var signalGutterRendererActivateId int
var signalGutterRendererActivateMap = make(map[int]signalGutterRendererActivateDetail)
var signalGutterRendererActivateLock sync.RWMutex

// GutterRendererSignalActivateCallback is a callback function for a 'activate' signal emitted from a GutterRenderer.
type GutterRendererSignalActivateCallback func(targetObject *GutterRenderer, iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event)

/*
ConnectActivate connects the callback to the 'activate' signal for the GutterRenderer.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *GutterRenderer) ConnectActivate(callback GutterRendererSignalActivateCallback) int {
	signalGutterRendererActivateLock.Lock()
	defer signalGutterRendererActivateLock.Unlock()

	signalGutterRendererActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.GutterRenderer_signal_connect_activate(instance, C.gpointer(uintptr(signalGutterRendererActivateId)))

	detail := signalGutterRendererActivateDetail{callback, handlerID}
	signalGutterRendererActivateMap[signalGutterRendererActivateId] = detail

	return signalGutterRendererActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the GutterRenderer.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *GutterRenderer) DisconnectActivate(connectionID int) {
	signalGutterRendererActivateLock.Lock()
	defer signalGutterRendererActivateLock.Unlock()

	detail, exists := signalGutterRendererActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGutterRendererActivateMap, connectionID)
}

//export gutterrenderer_activateHandler
func gutterrenderer_activateHandler(c_targetObject *C.GObject, c_iter *C.GtkTextIter, c_area *C.GdkRectangle, c_event *C.GdkEvent_, data C.gpointer) {
	signalGutterRendererActivateLock.RLock()
	defer signalGutterRendererActivateLock.RUnlock()

	iter := gtk.TextIterNewFromC(unsafe.Pointer(c_iter))

	area := gdk.RectangleNewFromC(unsafe.Pointer(c_area))

	event := gdk.EventNewFromC(unsafe.Pointer(c_event))

	targetObject := GutterRendererNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalGutterRendererActivateMap[index].callback
	callback(targetObject, iter, area, event)
}

type signalGutterRendererQueryActivatableDetail struct {
	callback  GutterRendererSignalQueryActivatableCallback
	handlerID C.gulong
}

var signalGutterRendererQueryActivatableId int
var signalGutterRendererQueryActivatableMap = make(map[int]signalGutterRendererQueryActivatableDetail)
var signalGutterRendererQueryActivatableLock sync.RWMutex

// GutterRendererSignalQueryActivatableCallback is a callback function for a 'query-activatable' signal emitted from a GutterRenderer.
type GutterRendererSignalQueryActivatableCallback func(targetObject *GutterRenderer, iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) bool

/*
ConnectQueryActivatable connects the callback to the 'query-activatable' signal for the GutterRenderer.

The returned value represents the connection, and may be passed to DisconnectQueryActivatable to remove it.
*/
func (recv *GutterRenderer) ConnectQueryActivatable(callback GutterRendererSignalQueryActivatableCallback) int {
	signalGutterRendererQueryActivatableLock.Lock()
	defer signalGutterRendererQueryActivatableLock.Unlock()

	signalGutterRendererQueryActivatableId++
	instance := C.gpointer(recv.native)
	handlerID := C.GutterRenderer_signal_connect_query_activatable(instance, C.gpointer(uintptr(signalGutterRendererQueryActivatableId)))

	detail := signalGutterRendererQueryActivatableDetail{callback, handlerID}
	signalGutterRendererQueryActivatableMap[signalGutterRendererQueryActivatableId] = detail

	return signalGutterRendererQueryActivatableId
}

/*
DisconnectQueryActivatable disconnects a callback from the 'query-activatable' signal for the GutterRenderer.

The connectionID should be a value returned from a call to ConnectQueryActivatable.
*/
func (recv *GutterRenderer) DisconnectQueryActivatable(connectionID int) {
	signalGutterRendererQueryActivatableLock.Lock()
	defer signalGutterRendererQueryActivatableLock.Unlock()

	detail, exists := signalGutterRendererQueryActivatableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGutterRendererQueryActivatableMap, connectionID)
}

//export gutterrenderer_queryActivatableHandler
func gutterrenderer_queryActivatableHandler(c_targetObject *C.GObject, c_iter *C.GtkTextIter, c_area *C.GdkRectangle, c_event *C.GdkEvent_, data C.gpointer) C.gboolean {
	signalGutterRendererQueryActivatableLock.RLock()
	defer signalGutterRendererQueryActivatableLock.RUnlock()

	iter := gtk.TextIterNewFromC(unsafe.Pointer(c_iter))

	area := gdk.RectangleNewFromC(unsafe.Pointer(c_area))

	event := gdk.EventNewFromC(unsafe.Pointer(c_event))

	targetObject := GutterRendererNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalGutterRendererQueryActivatableMap[index].callback
	retGo := callback(targetObject, iter, area, event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalGutterRendererQueryDataDetail struct {
	callback  GutterRendererSignalQueryDataCallback
	handlerID C.gulong
}

var signalGutterRendererQueryDataId int
var signalGutterRendererQueryDataMap = make(map[int]signalGutterRendererQueryDataDetail)
var signalGutterRendererQueryDataLock sync.RWMutex

// GutterRendererSignalQueryDataCallback is a callback function for a 'query-data' signal emitted from a GutterRenderer.
type GutterRendererSignalQueryDataCallback func(targetObject *GutterRenderer, start *gtk.TextIter, end *gtk.TextIter, state GutterRendererState)

/*
ConnectQueryData connects the callback to the 'query-data' signal for the GutterRenderer.

The returned value represents the connection, and may be passed to DisconnectQueryData to remove it.
*/
func (recv *GutterRenderer) ConnectQueryData(callback GutterRendererSignalQueryDataCallback) int {
	signalGutterRendererQueryDataLock.Lock()
	defer signalGutterRendererQueryDataLock.Unlock()

	signalGutterRendererQueryDataId++
	instance := C.gpointer(recv.native)
	handlerID := C.GutterRenderer_signal_connect_query_data(instance, C.gpointer(uintptr(signalGutterRendererQueryDataId)))

	detail := signalGutterRendererQueryDataDetail{callback, handlerID}
	signalGutterRendererQueryDataMap[signalGutterRendererQueryDataId] = detail

	return signalGutterRendererQueryDataId
}

/*
DisconnectQueryData disconnects a callback from the 'query-data' signal for the GutterRenderer.

The connectionID should be a value returned from a call to ConnectQueryData.
*/
func (recv *GutterRenderer) DisconnectQueryData(connectionID int) {
	signalGutterRendererQueryDataLock.Lock()
	defer signalGutterRendererQueryDataLock.Unlock()

	detail, exists := signalGutterRendererQueryDataMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGutterRendererQueryDataMap, connectionID)
}

//export gutterrenderer_queryDataHandler
func gutterrenderer_queryDataHandler(c_targetObject *C.GObject, c_start *C.GtkTextIter, c_end *C.GtkTextIter, c_state C.GtkSourceGutterRendererState, data C.gpointer) {
	signalGutterRendererQueryDataLock.RLock()
	defer signalGutterRendererQueryDataLock.RUnlock()

	start := gtk.TextIterNewFromC(unsafe.Pointer(c_start))

	end := gtk.TextIterNewFromC(unsafe.Pointer(c_end))

	state := GutterRendererState(c_state)

	targetObject := GutterRendererNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalGutterRendererQueryDataMap[index].callback
	callback(targetObject, start, end, state)
}

type signalGutterRendererQueryTooltipDetail struct {
	callback  GutterRendererSignalQueryTooltipCallback
	handlerID C.gulong
}

var signalGutterRendererQueryTooltipId int
var signalGutterRendererQueryTooltipMap = make(map[int]signalGutterRendererQueryTooltipDetail)
var signalGutterRendererQueryTooltipLock sync.RWMutex

// GutterRendererSignalQueryTooltipCallback is a callback function for a 'query-tooltip' signal emitted from a GutterRenderer.
type GutterRendererSignalQueryTooltipCallback func(targetObject *GutterRenderer, iter *gtk.TextIter, area *gdk.Rectangle, x int32, y int32, tooltip *gtk.Tooltip) bool

/*
ConnectQueryTooltip connects the callback to the 'query-tooltip' signal for the GutterRenderer.

The returned value represents the connection, and may be passed to DisconnectQueryTooltip to remove it.
*/
func (recv *GutterRenderer) ConnectQueryTooltip(callback GutterRendererSignalQueryTooltipCallback) int {
	signalGutterRendererQueryTooltipLock.Lock()
	defer signalGutterRendererQueryTooltipLock.Unlock()

	signalGutterRendererQueryTooltipId++
	instance := C.gpointer(recv.native)
	handlerID := C.GutterRenderer_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalGutterRendererQueryTooltipId)))

	detail := signalGutterRendererQueryTooltipDetail{callback, handlerID}
	signalGutterRendererQueryTooltipMap[signalGutterRendererQueryTooltipId] = detail

	return signalGutterRendererQueryTooltipId
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the GutterRenderer.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *GutterRenderer) DisconnectQueryTooltip(connectionID int) {
	signalGutterRendererQueryTooltipLock.Lock()
	defer signalGutterRendererQueryTooltipLock.Unlock()

	detail, exists := signalGutterRendererQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGutterRendererQueryTooltipMap, connectionID)
}

//export gutterrenderer_queryTooltipHandler
func gutterrenderer_queryTooltipHandler(c_targetObject *C.GObject, c_iter *C.GtkTextIter, c_area *C.GdkRectangle, c_x C.gint, c_y C.gint, c_tooltip *C.GtkTooltip, data C.gpointer) C.gboolean {
	signalGutterRendererQueryTooltipLock.RLock()
	defer signalGutterRendererQueryTooltipLock.RUnlock()

	iter := gtk.TextIterNewFromC(unsafe.Pointer(c_iter))

	area := gdk.RectangleNewFromC(unsafe.Pointer(c_area))

	x := int32(c_x)

	y := int32(c_y)

	tooltip := gtk.TooltipNewFromC(unsafe.Pointer(c_tooltip))

	targetObject := GutterRendererNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalGutterRendererQueryTooltipMap[index].callback
	retGo := callback(targetObject, iter, area, x, y, tooltip)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalGutterRendererQueueDrawDetail struct {
	callback  GutterRendererSignalQueueDrawCallback
	handlerID C.gulong
}

var signalGutterRendererQueueDrawId int
var signalGutterRendererQueueDrawMap = make(map[int]signalGutterRendererQueueDrawDetail)
var signalGutterRendererQueueDrawLock sync.RWMutex

// GutterRendererSignalQueueDrawCallback is a callback function for a 'queue-draw' signal emitted from a GutterRenderer.
type GutterRendererSignalQueueDrawCallback func(targetObject *GutterRenderer)

/*
ConnectQueueDraw connects the callback to the 'queue-draw' signal for the GutterRenderer.

The returned value represents the connection, and may be passed to DisconnectQueueDraw to remove it.
*/
func (recv *GutterRenderer) ConnectQueueDraw(callback GutterRendererSignalQueueDrawCallback) int {
	signalGutterRendererQueueDrawLock.Lock()
	defer signalGutterRendererQueueDrawLock.Unlock()

	signalGutterRendererQueueDrawId++
	instance := C.gpointer(recv.native)
	handlerID := C.GutterRenderer_signal_connect_queue_draw(instance, C.gpointer(uintptr(signalGutterRendererQueueDrawId)))

	detail := signalGutterRendererQueueDrawDetail{callback, handlerID}
	signalGutterRendererQueueDrawMap[signalGutterRendererQueueDrawId] = detail

	return signalGutterRendererQueueDrawId
}

/*
DisconnectQueueDraw disconnects a callback from the 'queue-draw' signal for the GutterRenderer.

The connectionID should be a value returned from a call to ConnectQueueDraw.
*/
func (recv *GutterRenderer) DisconnectQueueDraw(connectionID int) {
	signalGutterRendererQueueDrawLock.Lock()
	defer signalGutterRendererQueueDrawLock.Unlock()

	detail, exists := signalGutterRendererQueueDrawMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalGutterRendererQueueDrawMap, connectionID)
}

//export gutterrenderer_queueDrawHandler
func gutterrenderer_queueDrawHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalGutterRendererQueueDrawLock.RLock()
	defer signalGutterRendererQueueDrawLock.RUnlock()

	targetObject := GutterRendererNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalGutterRendererQueueDrawMap[index].callback
	callback(targetObject)
}

// Activate is a wrapper around the C function gtk_source_gutter_renderer_activate.
func (recv *GutterRenderer) Activate(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_area := (*C.GdkRectangle)(C.NULL)
	if area != nil {
		c_area = (*C.GdkRectangle)(area.ToC())
	}

	c_event := (*C.GdkEvent)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEvent)(event.ToC())
	}

	C.gtk_source_gutter_renderer_activate((*C.GtkSourceGutterRenderer)(recv.native), c_iter, c_area, c_event)

	return
}

// Begin is a wrapper around the C function gtk_source_gutter_renderer_begin.
func (recv *GutterRenderer) Begin(cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, start *gtk.TextIter, end *gtk.TextIter) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_background_area := (*C.GdkRectangle)(C.NULL)
	if backgroundArea != nil {
		c_background_area = (*C.GdkRectangle)(backgroundArea.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	C.gtk_source_gutter_renderer_begin((*C.GtkSourceGutterRenderer)(recv.native), c_cr, c_background_area, c_cell_area, c_start, c_end)

	return
}

// Draw is a wrapper around the C function gtk_source_gutter_renderer_draw.
func (recv *GutterRenderer) Draw(cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, start *gtk.TextIter, end *gtk.TextIter, state GutterRendererState) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_background_area := (*C.GdkRectangle)(C.NULL)
	if backgroundArea != nil {
		c_background_area = (*C.GdkRectangle)(backgroundArea.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	c_state := (C.GtkSourceGutterRendererState)(state)

	C.gtk_source_gutter_renderer_draw((*C.GtkSourceGutterRenderer)(recv.native), c_cr, c_background_area, c_cell_area, c_start, c_end, c_state)

	return
}

// End is a wrapper around the C function gtk_source_gutter_renderer_end.
func (recv *GutterRenderer) End() {
	C.gtk_source_gutter_renderer_end((*C.GtkSourceGutterRenderer)(recv.native))

	return
}

// GetAlignment is a wrapper around the C function gtk_source_gutter_renderer_get_alignment.
func (recv *GutterRenderer) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_source_gutter_renderer_get_alignment((*C.GtkSourceGutterRenderer)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// GetAlignmentMode is a wrapper around the C function gtk_source_gutter_renderer_get_alignment_mode.
func (recv *GutterRenderer) GetAlignmentMode() GutterRendererAlignmentMode {
	retC := C.gtk_source_gutter_renderer_get_alignment_mode((*C.GtkSourceGutterRenderer)(recv.native))
	retGo := (GutterRendererAlignmentMode)(retC)

	return retGo
}

// GetBackground is a wrapper around the C function gtk_source_gutter_renderer_get_background.
func (recv *GutterRenderer) GetBackground() (bool, *gdk.RGBA) {
	var c_color C.GdkRGBA

	retC := C.gtk_source_gutter_renderer_get_background((*C.GtkSourceGutterRenderer)(recv.native), &c_color)
	retGo := retC == C.TRUE

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// GetPadding is a wrapper around the C function gtk_source_gutter_renderer_get_padding.
func (recv *GutterRenderer) GetPadding() (int32, int32) {
	var c_xpad C.gint

	var c_ypad C.gint

	C.gtk_source_gutter_renderer_get_padding((*C.GtkSourceGutterRenderer)(recv.native), &c_xpad, &c_ypad)

	xpad := (int32)(c_xpad)

	ypad := (int32)(c_ypad)

	return xpad, ypad
}

// GetSize is a wrapper around the C function gtk_source_gutter_renderer_get_size.
func (recv *GutterRenderer) GetSize() int32 {
	retC := C.gtk_source_gutter_renderer_get_size((*C.GtkSourceGutterRenderer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetView is a wrapper around the C function gtk_source_gutter_renderer_get_view.
func (recv *GutterRenderer) GetView() *gtk.TextView {
	retC := C.gtk_source_gutter_renderer_get_view((*C.GtkSourceGutterRenderer)(recv.native))
	retGo := gtk.TextViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisible is a wrapper around the C function gtk_source_gutter_renderer_get_visible.
func (recv *GutterRenderer) GetVisible() bool {
	retC := C.gtk_source_gutter_renderer_get_visible((*C.GtkSourceGutterRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWindowType is a wrapper around the C function gtk_source_gutter_renderer_get_window_type.
func (recv *GutterRenderer) GetWindowType() gtk.TextWindowType {
	retC := C.gtk_source_gutter_renderer_get_window_type((*C.GtkSourceGutterRenderer)(recv.native))
	retGo := (gtk.TextWindowType)(retC)

	return retGo
}

// QueryActivatable is a wrapper around the C function gtk_source_gutter_renderer_query_activatable.
func (recv *GutterRenderer) QueryActivatable(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_area := (*C.GdkRectangle)(C.NULL)
	if area != nil {
		c_area = (*C.GdkRectangle)(area.ToC())
	}

	c_event := (*C.GdkEvent)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEvent)(event.ToC())
	}

	retC := C.gtk_source_gutter_renderer_query_activatable((*C.GtkSourceGutterRenderer)(recv.native), c_iter, c_area, c_event)
	retGo := retC == C.TRUE

	return retGo
}

// QueryData is a wrapper around the C function gtk_source_gutter_renderer_query_data.
func (recv *GutterRenderer) QueryData(start *gtk.TextIter, end *gtk.TextIter, state GutterRendererState) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	c_state := (C.GtkSourceGutterRendererState)(state)

	C.gtk_source_gutter_renderer_query_data((*C.GtkSourceGutterRenderer)(recv.native), c_start, c_end, c_state)

	return
}

// QueryTooltip is a wrapper around the C function gtk_source_gutter_renderer_query_tooltip.
func (recv *GutterRenderer) QueryTooltip(iter *gtk.TextIter, area *gdk.Rectangle, x int32, y int32, tooltip *gtk.Tooltip) bool {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	c_area := (*C.GdkRectangle)(C.NULL)
	if area != nil {
		c_area = (*C.GdkRectangle)(area.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	retC := C.gtk_source_gutter_renderer_query_tooltip((*C.GtkSourceGutterRenderer)(recv.native), c_iter, c_area, c_x, c_y, c_tooltip)
	retGo := retC == C.TRUE

	return retGo
}

// QueueDraw is a wrapper around the C function gtk_source_gutter_renderer_queue_draw.
func (recv *GutterRenderer) QueueDraw() {
	C.gtk_source_gutter_renderer_queue_draw((*C.GtkSourceGutterRenderer)(recv.native))

	return
}

// SetAlignment is a wrapper around the C function gtk_source_gutter_renderer_set_alignment.
func (recv *GutterRenderer) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_source_gutter_renderer_set_alignment((*C.GtkSourceGutterRenderer)(recv.native), c_xalign, c_yalign)

	return
}

// SetAlignmentMode is a wrapper around the C function gtk_source_gutter_renderer_set_alignment_mode.
func (recv *GutterRenderer) SetAlignmentMode(mode GutterRendererAlignmentMode) {
	c_mode := (C.GtkSourceGutterRendererAlignmentMode)(mode)

	C.gtk_source_gutter_renderer_set_alignment_mode((*C.GtkSourceGutterRenderer)(recv.native), c_mode)

	return
}

// SetBackground is a wrapper around the C function gtk_source_gutter_renderer_set_background.
func (recv *GutterRenderer) SetBackground(color *gdk.RGBA) {
	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_source_gutter_renderer_set_background((*C.GtkSourceGutterRenderer)(recv.native), c_color)

	return
}

// SetPadding is a wrapper around the C function gtk_source_gutter_renderer_set_padding.
func (recv *GutterRenderer) SetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_source_gutter_renderer_set_padding((*C.GtkSourceGutterRenderer)(recv.native), c_xpad, c_ypad)

	return
}

// SetSize is a wrapper around the C function gtk_source_gutter_renderer_set_size.
func (recv *GutterRenderer) SetSize(size int32) {
	c_size := (C.gint)(size)

	C.gtk_source_gutter_renderer_set_size((*C.GtkSourceGutterRenderer)(recv.native), c_size)

	return
}

// SetVisible is a wrapper around the C function gtk_source_gutter_renderer_set_visible.
func (recv *GutterRenderer) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_source_gutter_renderer_set_visible((*C.GtkSourceGutterRenderer)(recv.native), c_visible)

	return
}

// GutterRendererPixbuf is a wrapper around the C record GtkSourceGutterRendererPixbuf.
type GutterRendererPixbuf struct {
	native *C.GtkSourceGutterRendererPixbuf
	// Private : parent
	// Private : priv
}

func GutterRendererPixbufNewFromC(u unsafe.Pointer) *GutterRendererPixbuf {
	c := (*C.GtkSourceGutterRendererPixbuf)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererPixbuf{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GutterRendererPixbuf) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GutterRendererPixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererPixbuf with another GutterRendererPixbuf, and returns true if they represent the same GObject.
func (recv *GutterRendererPixbuf) Equals(other *GutterRendererPixbuf) bool {
	return other.ToC() == recv.ToC()
}

// GutterRenderer upcasts to *GutterRenderer
func (recv *GutterRendererPixbuf) GutterRenderer() *GutterRenderer {
	return GutterRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GutterRendererPixbuf) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.GutterRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *GutterRendererPixbuf) Object() *gobject.Object {
	return recv.GutterRenderer().Object()
}

// CastToWidget down casts any arbitrary Object to GutterRendererPixbuf.
// Exercise care, as this is a potentially dangerous function if the Object is not a GutterRendererPixbuf.
func CastToGutterRendererPixbuf(object *gobject.Object) *GutterRendererPixbuf {
	return GutterRendererPixbufNewFromC(object.ToC())
}

// GutterRendererPixbufNew is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_new.
func GutterRendererPixbufNew() *GutterRendererPixbuf {
	retC := C.gtk_source_gutter_renderer_pixbuf_new()
	retGo := GutterRendererPixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGicon is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_get_gicon.
func (recv *GutterRendererPixbuf) GetGicon() *gio.Icon {
	retC := C.gtk_source_gutter_renderer_pixbuf_get_gicon((*C.GtkSourceGutterRendererPixbuf)(recv.native))
	retGo := gio.IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconName is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_get_icon_name.
func (recv *GutterRendererPixbuf) GetIconName() string {
	retC := C.gtk_source_gutter_renderer_pixbuf_get_icon_name((*C.GtkSourceGutterRendererPixbuf)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPixbuf is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_get_pixbuf.
func (recv *GutterRendererPixbuf) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_source_gutter_renderer_pixbuf_get_pixbuf((*C.GtkSourceGutterRendererPixbuf)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStockId is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_get_stock_id.
func (recv *GutterRendererPixbuf) GetStockId() string {
	retC := C.gtk_source_gutter_renderer_pixbuf_get_stock_id((*C.GtkSourceGutterRendererPixbuf)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetGicon is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_set_gicon.
func (recv *GutterRendererPixbuf) SetGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_source_gutter_renderer_pixbuf_set_gicon((*C.GtkSourceGutterRendererPixbuf)(recv.native), c_icon)

	return
}

// SetIconName is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_set_icon_name.
func (recv *GutterRendererPixbuf) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_source_gutter_renderer_pixbuf_set_icon_name((*C.GtkSourceGutterRendererPixbuf)(recv.native), c_icon_name)

	return
}

// SetPixbuf is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_set_pixbuf.
func (recv *GutterRendererPixbuf) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_source_gutter_renderer_pixbuf_set_pixbuf((*C.GtkSourceGutterRendererPixbuf)(recv.native), c_pixbuf)

	return
}

// SetStockId is a wrapper around the C function gtk_source_gutter_renderer_pixbuf_set_stock_id.
func (recv *GutterRendererPixbuf) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_source_gutter_renderer_pixbuf_set_stock_id((*C.GtkSourceGutterRendererPixbuf)(recv.native), c_stock_id)

	return
}

// GutterRendererText is a wrapper around the C record GtkSourceGutterRendererText.
type GutterRendererText struct {
	native *C.GtkSourceGutterRendererText
	// Private : parent
	// Private : priv
}

func GutterRendererTextNewFromC(u unsafe.Pointer) *GutterRendererText {
	c := (*C.GtkSourceGutterRendererText)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererText{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GutterRendererText) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GutterRendererText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererText with another GutterRendererText, and returns true if they represent the same GObject.
func (recv *GutterRendererText) Equals(other *GutterRendererText) bool {
	return other.ToC() == recv.ToC()
}

// GutterRenderer upcasts to *GutterRenderer
func (recv *GutterRendererText) GutterRenderer() *GutterRenderer {
	return GutterRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *GutterRendererText) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.GutterRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *GutterRendererText) Object() *gobject.Object {
	return recv.GutterRenderer().Object()
}

// CastToWidget down casts any arbitrary Object to GutterRendererText.
// Exercise care, as this is a potentially dangerous function if the Object is not a GutterRendererText.
func CastToGutterRendererText(object *gobject.Object) *GutterRendererText {
	return GutterRendererTextNewFromC(object.ToC())
}

// GutterRendererTextNew is a wrapper around the C function gtk_source_gutter_renderer_text_new.
func GutterRendererTextNew() *GutterRendererText {
	retC := C.gtk_source_gutter_renderer_text_new()
	retGo := GutterRendererTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Measure is a wrapper around the C function gtk_source_gutter_renderer_text_measure.
func (recv *GutterRendererText) Measure(text string) (int32, int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	var c_width C.gint

	var c_height C.gint

	C.gtk_source_gutter_renderer_text_measure((*C.GtkSourceGutterRendererText)(recv.native), c_text, &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// MeasureMarkup is a wrapper around the C function gtk_source_gutter_renderer_text_measure_markup.
func (recv *GutterRendererText) MeasureMarkup(markup string) (int32, int32) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	var c_width C.gint

	var c_height C.gint

	C.gtk_source_gutter_renderer_text_measure_markup((*C.GtkSourceGutterRendererText)(recv.native), c_markup, &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// SetMarkup is a wrapper around the C function gtk_source_gutter_renderer_text_set_markup.
func (recv *GutterRendererText) SetMarkup(markup string, length int32) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	c_length := (C.gint)(length)

	C.gtk_source_gutter_renderer_text_set_markup((*C.GtkSourceGutterRendererText)(recv.native), c_markup, c_length)

	return
}

// SetText is a wrapper around the C function gtk_source_gutter_renderer_text_set_text.
func (recv *GutterRendererText) SetText(text string, length int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_length := (C.gint)(length)

	C.gtk_source_gutter_renderer_text_set_text((*C.GtkSourceGutterRendererText)(recv.native), c_text, c_length)

	return
}

// Language is a wrapper around the C record GtkSourceLanguage.
type Language struct {
	native *C.GtkSourceLanguage
	// parent_instance : record
	// priv : record
}

func LanguageNewFromC(u unsafe.Pointer) *Language {
	c := (*C.GtkSourceLanguage)(u)
	if c == nil {
		return nil
	}

	g := &Language{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Language) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Language) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Language with another Language, and returns true if they represent the same GObject.
func (recv *Language) Equals(other *Language) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Language) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Language.
// Exercise care, as this is a potentially dangerous function if the Object is not a Language.
func CastToLanguage(object *gobject.Object) *Language {
	return LanguageNewFromC(object.ToC())
}

// GetGlobs is a wrapper around the C function gtk_source_language_get_globs.
func (recv *Language) GetGlobs() []string {
	retC := C.gtk_source_language_get_globs((*C.GtkSourceLanguage)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetHidden is a wrapper around the C function gtk_source_language_get_hidden.
func (recv *Language) GetHidden() bool {
	retC := C.gtk_source_language_get_hidden((*C.GtkSourceLanguage)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetId is a wrapper around the C function gtk_source_language_get_id.
func (recv *Language) GetId() string {
	retC := C.gtk_source_language_get_id((*C.GtkSourceLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMetadata is a wrapper around the C function gtk_source_language_get_metadata.
func (recv *Language) GetMetadata(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_source_language_get_metadata((*C.GtkSourceLanguage)(recv.native), c_name)
	retGo := C.GoString(retC)

	return retGo
}

// GetMimeTypes is a wrapper around the C function gtk_source_language_get_mime_types.
func (recv *Language) GetMimeTypes() []string {
	retC := C.gtk_source_language_get_mime_types((*C.GtkSourceLanguage)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetName is a wrapper around the C function gtk_source_language_get_name.
func (recv *Language) GetName() string {
	retC := C.gtk_source_language_get_name((*C.GtkSourceLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSection is a wrapper around the C function gtk_source_language_get_section.
func (recv *Language) GetSection() string {
	retC := C.gtk_source_language_get_section((*C.GtkSourceLanguage)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStyleFallback is a wrapper around the C function gtk_source_language_get_style_fallback.
func (recv *Language) GetStyleFallback(styleId string) string {
	c_style_id := C.CString(styleId)
	defer C.free(unsafe.Pointer(c_style_id))

	retC := C.gtk_source_language_get_style_fallback((*C.GtkSourceLanguage)(recv.native), c_style_id)
	retGo := C.GoString(retC)

	return retGo
}

// GetStyleIds is a wrapper around the C function gtk_source_language_get_style_ids.
func (recv *Language) GetStyleIds() []string {
	retC := C.gtk_source_language_get_style_ids((*C.GtkSourceLanguage)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetStyleName is a wrapper around the C function gtk_source_language_get_style_name.
func (recv *Language) GetStyleName(styleId string) string {
	c_style_id := C.CString(styleId)
	defer C.free(unsafe.Pointer(c_style_id))

	retC := C.gtk_source_language_get_style_name((*C.GtkSourceLanguage)(recv.native), c_style_id)
	retGo := C.GoString(retC)

	return retGo
}

// LanguageManager is a wrapper around the C record GtkSourceLanguageManager.
type LanguageManager struct {
	native *C.GtkSourceLanguageManager
	// parent_instance : record
	// priv : record
}

func LanguageManagerNewFromC(u unsafe.Pointer) *LanguageManager {
	c := (*C.GtkSourceLanguageManager)(u)
	if c == nil {
		return nil
	}

	g := &LanguageManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *LanguageManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *LanguageManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LanguageManager with another LanguageManager, and returns true if they represent the same GObject.
func (recv *LanguageManager) Equals(other *LanguageManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *LanguageManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to LanguageManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a LanguageManager.
func CastToLanguageManager(object *gobject.Object) *LanguageManager {
	return LanguageManagerNewFromC(object.ToC())
}

// LanguageManagerNew is a wrapper around the C function gtk_source_language_manager_new.
func LanguageManagerNew() *LanguageManager {
	retC := C.gtk_source_language_manager_new()
	retGo := LanguageManagerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// LanguageManagerGetDefault is a wrapper around the C function gtk_source_language_manager_get_default.
func LanguageManagerGetDefault() *LanguageManager {
	retC := C.gtk_source_language_manager_get_default()
	retGo := LanguageManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLanguage is a wrapper around the C function gtk_source_language_manager_get_language.
func (recv *LanguageManager) GetLanguage(id string) *Language {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	retC := C.gtk_source_language_manager_get_language((*C.GtkSourceLanguageManager)(recv.native), c_id)
	var retGo (*Language)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LanguageNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLanguageIds is a wrapper around the C function gtk_source_language_manager_get_language_ids.
func (recv *LanguageManager) GetLanguageIds() []string {
	retC := C.gtk_source_language_manager_get_language_ids((*C.GtkSourceLanguageManager)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetSearchPath is a wrapper around the C function gtk_source_language_manager_get_search_path.
func (recv *LanguageManager) GetSearchPath() []string {
	retC := C.gtk_source_language_manager_get_search_path((*C.GtkSourceLanguageManager)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GuessLanguage is a wrapper around the C function gtk_source_language_manager_guess_language.
func (recv *LanguageManager) GuessLanguage(filename string, contentType string) *Language {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_source_language_manager_guess_language((*C.GtkSourceLanguageManager)(recv.native), c_filename, c_content_type)
	var retGo (*Language)
	if retC == nil {
		retGo = nil
	} else {
		retGo = LanguageNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetSearchPath is a wrapper around the C function gtk_source_language_manager_set_search_path.
func (recv *LanguageManager) SetSearchPath(dirs []string) {
	c_dirs_array := make([]*C.gchar, len(dirs)+1, len(dirs)+1)
	for i, item := range dirs {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_dirs_array[i] = c
	}
	c_dirs_array[len(dirs)] = nil
	c_dirs_arrayPtr := &c_dirs_array[0]
	c_dirs := (**C.gchar)(unsafe.Pointer(c_dirs_arrayPtr))

	C.gtk_source_language_manager_set_search_path((*C.GtkSourceLanguageManager)(recv.native), c_dirs)

	return
}

// Map is a wrapper around the C record GtkSourceMap.
type Map struct {
	native *C.GtkSourceMap
	// parent_instance : record
}

func MapNewFromC(u unsafe.Pointer) *Map {
	c := (*C.GtkSourceMap)(u)
	if c == nil {
		return nil
	}

	g := &Map{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Map) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Map) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Map with another Map, and returns true if they represent the same GObject.
func (recv *Map) Equals(other *Map) bool {
	return other.ToC() == recv.ToC()
}

// View upcasts to *View
func (recv *Map) View() *View {
	return ViewNewFromC(unsafe.Pointer(recv.native))
}

// TextView upcasts to *TextView
func (recv *Map) TextView() *gtk.TextView {
	return recv.View().TextView()
}

// Container upcasts to *Container
func (recv *Map) Container() *gtk.Container {
	return recv.View().Container()
}

// Widget upcasts to *Widget
func (recv *Map) Widget() *gtk.Widget {
	return recv.View().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Map) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.View().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Map) Object() *gobject.Object {
	return recv.View().Object()
}

// CastToWidget down casts any arbitrary Object to Map.
// Exercise care, as this is a potentially dangerous function if the Object is not a Map.
func CastToMap(object *gobject.Object) *Map {
	return MapNewFromC(object.ToC())
}

// MapNew is a wrapper around the C function gtk_source_map_new.
func MapNew() *Map {
	retC := C.gtk_source_map_new()
	retGo := MapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetView is a wrapper around the C function gtk_source_map_get_view.
func (recv *Map) GetView() *View {
	retC := C.gtk_source_map_get_view((*C.GtkSourceMap)(recv.native))
	var retGo (*View)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ViewNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetView is a wrapper around the C function gtk_source_map_set_view.
func (recv *Map) SetView(view *View) {
	c_view := (*C.GtkSourceView)(C.NULL)
	if view != nil {
		c_view = (*C.GtkSourceView)(view.ToC())
	}

	C.gtk_source_map_set_view((*C.GtkSourceMap)(recv.native), c_view)

	return
}

// ImplementorIface returns the ImplementorIface interface implemented by Map
func (recv *Map) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Map
func (recv *Map) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by Map
func (recv *Map) Scrollable() *gtk.Scrollable {
	return gtk.ScrollableNewFromC(recv.ToC())
}

// Mark is a wrapper around the C record GtkSourceMark.
type Mark struct {
	native *C.GtkSourceMark
	// parent_instance : record
	// priv : record
}

func MarkNewFromC(u unsafe.Pointer) *Mark {
	c := (*C.GtkSourceMark)(u)
	if c == nil {
		return nil
	}

	g := &Mark{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Mark) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Mark) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Mark with another Mark, and returns true if they represent the same GObject.
func (recv *Mark) Equals(other *Mark) bool {
	return other.ToC() == recv.ToC()
}

// TextMark upcasts to *TextMark
func (recv *Mark) TextMark() *gtk.TextMark {
	return gtk.TextMarkNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Mark) Object() *gobject.Object {
	return recv.TextMark().Object()
}

// CastToWidget down casts any arbitrary Object to Mark.
// Exercise care, as this is a potentially dangerous function if the Object is not a Mark.
func CastToMark(object *gobject.Object) *Mark {
	return MarkNewFromC(object.ToC())
}

// MarkNew is a wrapper around the C function gtk_source_mark_new.
func MarkNew(name string, category string) *Mark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_mark_new(c_name, c_category)
	retGo := MarkNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetCategory is a wrapper around the C function gtk_source_mark_get_category.
func (recv *Mark) GetCategory() string {
	retC := C.gtk_source_mark_get_category((*C.GtkSourceMark)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Next is a wrapper around the C function gtk_source_mark_next.
func (recv *Mark) Next(category string) *Mark {
	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_mark_next((*C.GtkSourceMark)(recv.native), c_category)
	var retGo (*Mark)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MarkNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Prev is a wrapper around the C function gtk_source_mark_prev.
func (recv *Mark) Prev(category string) *Mark {
	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	retC := C.gtk_source_mark_prev((*C.GtkSourceMark)(recv.native), c_category)
	var retGo (*Mark)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MarkNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// MarkAttributes is a wrapper around the C record GtkSourceMarkAttributes.
type MarkAttributes struct {
	native *C.GtkSourceMarkAttributes
	// Private : parent
	// Private : priv
}

func MarkAttributesNewFromC(u unsafe.Pointer) *MarkAttributes {
	c := (*C.GtkSourceMarkAttributes)(u)
	if c == nil {
		return nil
	}

	g := &MarkAttributes{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MarkAttributes) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MarkAttributes) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkAttributes with another MarkAttributes, and returns true if they represent the same GObject.
func (recv *MarkAttributes) Equals(other *MarkAttributes) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *MarkAttributes) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to MarkAttributes.
// Exercise care, as this is a potentially dangerous function if the Object is not a MarkAttributes.
func CastToMarkAttributes(object *gobject.Object) *MarkAttributes {
	return MarkAttributesNewFromC(object.ToC())
}

// Unsupported signal 'query-tooltip-markup' for MarkAttributes : return value utf8 :

// Unsupported signal 'query-tooltip-text' for MarkAttributes : return value utf8 :

// MarkAttributesNew is a wrapper around the C function gtk_source_mark_attributes_new.
func MarkAttributesNew() *MarkAttributes {
	retC := C.gtk_source_mark_attributes_new()
	retGo := MarkAttributesNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetBackground is a wrapper around the C function gtk_source_mark_attributes_get_background.
func (recv *MarkAttributes) GetBackground() (bool, *gdk.RGBA) {
	var c_background C.GdkRGBA

	retC := C.gtk_source_mark_attributes_get_background((*C.GtkSourceMarkAttributes)(recv.native), &c_background)
	retGo := retC == C.TRUE

	background := gdk.RGBANewFromC(unsafe.Pointer(&c_background))

	return retGo, background
}

// GetGicon is a wrapper around the C function gtk_source_mark_attributes_get_gicon.
func (recv *MarkAttributes) GetGicon() *gio.Icon {
	retC := C.gtk_source_mark_attributes_get_gicon((*C.GtkSourceMarkAttributes)(recv.native))
	retGo := gio.IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconName is a wrapper around the C function gtk_source_mark_attributes_get_icon_name.
func (recv *MarkAttributes) GetIconName() string {
	retC := C.gtk_source_mark_attributes_get_icon_name((*C.GtkSourceMarkAttributes)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPixbuf is a wrapper around the C function gtk_source_mark_attributes_get_pixbuf.
func (recv *MarkAttributes) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_source_mark_attributes_get_pixbuf((*C.GtkSourceMarkAttributes)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStockId is a wrapper around the C function gtk_source_mark_attributes_get_stock_id.
func (recv *MarkAttributes) GetStockId() string {
	retC := C.gtk_source_mark_attributes_get_stock_id((*C.GtkSourceMarkAttributes)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTooltipMarkup is a wrapper around the C function gtk_source_mark_attributes_get_tooltip_markup.
func (recv *MarkAttributes) GetTooltipMarkup(mark *Mark) string {
	c_mark := (*C.GtkSourceMark)(C.NULL)
	if mark != nil {
		c_mark = (*C.GtkSourceMark)(mark.ToC())
	}

	retC := C.gtk_source_mark_attributes_get_tooltip_markup((*C.GtkSourceMarkAttributes)(recv.native), c_mark)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTooltipText is a wrapper around the C function gtk_source_mark_attributes_get_tooltip_text.
func (recv *MarkAttributes) GetTooltipText(mark *Mark) string {
	c_mark := (*C.GtkSourceMark)(C.NULL)
	if mark != nil {
		c_mark = (*C.GtkSourceMark)(mark.ToC())
	}

	retC := C.gtk_source_mark_attributes_get_tooltip_text((*C.GtkSourceMarkAttributes)(recv.native), c_mark)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RenderIcon is a wrapper around the C function gtk_source_mark_attributes_render_icon.
func (recv *MarkAttributes) RenderIcon(widget *gtk.Widget, size int32) *gdkpixbuf.Pixbuf {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_size := (C.gint)(size)

	retC := C.gtk_source_mark_attributes_render_icon((*C.GtkSourceMarkAttributes)(recv.native), c_widget, c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetBackground is a wrapper around the C function gtk_source_mark_attributes_set_background.
func (recv *MarkAttributes) SetBackground(background *gdk.RGBA) {
	c_background := (*C.GdkRGBA)(C.NULL)
	if background != nil {
		c_background = (*C.GdkRGBA)(background.ToC())
	}

	C.gtk_source_mark_attributes_set_background((*C.GtkSourceMarkAttributes)(recv.native), c_background)

	return
}

// SetGicon is a wrapper around the C function gtk_source_mark_attributes_set_gicon.
func (recv *MarkAttributes) SetGicon(gicon *gio.Icon) {
	c_gicon := (*C.GIcon)(gicon.ToC())

	C.gtk_source_mark_attributes_set_gicon((*C.GtkSourceMarkAttributes)(recv.native), c_gicon)

	return
}

// SetIconName is a wrapper around the C function gtk_source_mark_attributes_set_icon_name.
func (recv *MarkAttributes) SetIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_source_mark_attributes_set_icon_name((*C.GtkSourceMarkAttributes)(recv.native), c_icon_name)

	return
}

// SetPixbuf is a wrapper around the C function gtk_source_mark_attributes_set_pixbuf.
func (recv *MarkAttributes) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_source_mark_attributes_set_pixbuf((*C.GtkSourceMarkAttributes)(recv.native), c_pixbuf)

	return
}

// SetStockId is a wrapper around the C function gtk_source_mark_attributes_set_stock_id.
func (recv *MarkAttributes) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_source_mark_attributes_set_stock_id((*C.GtkSourceMarkAttributes)(recv.native), c_stock_id)

	return
}

// PrintCompositor is a wrapper around the C record GtkSourcePrintCompositor.
type PrintCompositor struct {
	native *C.GtkSourcePrintCompositor
	// parent_instance : record
	// priv : record
}

func PrintCompositorNewFromC(u unsafe.Pointer) *PrintCompositor {
	c := (*C.GtkSourcePrintCompositor)(u)
	if c == nil {
		return nil
	}

	g := &PrintCompositor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PrintCompositor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PrintCompositor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintCompositor with another PrintCompositor, and returns true if they represent the same GObject.
func (recv *PrintCompositor) Equals(other *PrintCompositor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PrintCompositor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PrintCompositor.
// Exercise care, as this is a potentially dangerous function if the Object is not a PrintCompositor.
func CastToPrintCompositor(object *gobject.Object) *PrintCompositor {
	return PrintCompositorNewFromC(object.ToC())
}

// PrintCompositorNew is a wrapper around the C function gtk_source_print_compositor_new.
func PrintCompositorNew(buffer *Buffer) *PrintCompositor {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	retC := C.gtk_source_print_compositor_new(c_buffer)
	retGo := PrintCompositorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PrintCompositorNewFromView is a wrapper around the C function gtk_source_print_compositor_new_from_view.
func PrintCompositorNewFromView(view *View) *PrintCompositor {
	c_view := (*C.GtkSourceView)(C.NULL)
	if view != nil {
		c_view = (*C.GtkSourceView)(view.ToC())
	}

	retC := C.gtk_source_print_compositor_new_from_view(c_view)
	retGo := PrintCompositorNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// DrawPage is a wrapper around the C function gtk_source_print_compositor_draw_page.
func (recv *PrintCompositor) DrawPage(context *gtk.PrintContext, pageNr int32) {
	c_context := (*C.GtkPrintContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkPrintContext)(context.ToC())
	}

	c_page_nr := (C.gint)(pageNr)

	C.gtk_source_print_compositor_draw_page((*C.GtkSourcePrintCompositor)(recv.native), c_context, c_page_nr)

	return
}

// GetBodyFontName is a wrapper around the C function gtk_source_print_compositor_get_body_font_name.
func (recv *PrintCompositor) GetBodyFontName() string {
	retC := C.gtk_source_print_compositor_get_body_font_name((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetBottomMargin is a wrapper around the C function gtk_source_print_compositor_get_bottom_margin.
func (recv *PrintCompositor) GetBottomMargin(unit gtk.Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_source_print_compositor_get_bottom_margin((*C.GtkSourcePrintCompositor)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_source_print_compositor_get_buffer.
func (recv *PrintCompositor) GetBuffer() *Buffer {
	retC := C.gtk_source_print_compositor_get_buffer((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFooterFontName is a wrapper around the C function gtk_source_print_compositor_get_footer_font_name.
func (recv *PrintCompositor) GetFooterFontName() string {
	retC := C.gtk_source_print_compositor_get_footer_font_name((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetHeaderFontName is a wrapper around the C function gtk_source_print_compositor_get_header_font_name.
func (recv *PrintCompositor) GetHeaderFontName() string {
	retC := C.gtk_source_print_compositor_get_header_font_name((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetHighlightSyntax is a wrapper around the C function gtk_source_print_compositor_get_highlight_syntax.
func (recv *PrintCompositor) GetHighlightSyntax() bool {
	retC := C.gtk_source_print_compositor_get_highlight_syntax((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLeftMargin is a wrapper around the C function gtk_source_print_compositor_get_left_margin.
func (recv *PrintCompositor) GetLeftMargin(unit gtk.Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_source_print_compositor_get_left_margin((*C.GtkSourcePrintCompositor)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetLineNumbersFontName is a wrapper around the C function gtk_source_print_compositor_get_line_numbers_font_name.
func (recv *PrintCompositor) GetLineNumbersFontName() string {
	retC := C.gtk_source_print_compositor_get_line_numbers_font_name((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetNPages is a wrapper around the C function gtk_source_print_compositor_get_n_pages.
func (recv *PrintCompositor) GetNPages() int32 {
	retC := C.gtk_source_print_compositor_get_n_pages((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPaginationProgress is a wrapper around the C function gtk_source_print_compositor_get_pagination_progress.
func (recv *PrintCompositor) GetPaginationProgress() float64 {
	retC := C.gtk_source_print_compositor_get_pagination_progress((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetPrintFooter is a wrapper around the C function gtk_source_print_compositor_get_print_footer.
func (recv *PrintCompositor) GetPrintFooter() bool {
	retC := C.gtk_source_print_compositor_get_print_footer((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPrintHeader is a wrapper around the C function gtk_source_print_compositor_get_print_header.
func (recv *PrintCompositor) GetPrintHeader() bool {
	retC := C.gtk_source_print_compositor_get_print_header((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPrintLineNumbers is a wrapper around the C function gtk_source_print_compositor_get_print_line_numbers.
func (recv *PrintCompositor) GetPrintLineNumbers() uint32 {
	retC := C.gtk_source_print_compositor_get_print_line_numbers((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetRightMargin is a wrapper around the C function gtk_source_print_compositor_get_right_margin.
func (recv *PrintCompositor) GetRightMargin(unit gtk.Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_source_print_compositor_get_right_margin((*C.GtkSourcePrintCompositor)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetTabWidth is a wrapper around the C function gtk_source_print_compositor_get_tab_width.
func (recv *PrintCompositor) GetTabWidth() uint32 {
	retC := C.gtk_source_print_compositor_get_tab_width((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTopMargin is a wrapper around the C function gtk_source_print_compositor_get_top_margin.
func (recv *PrintCompositor) GetTopMargin(unit gtk.Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_source_print_compositor_get_top_margin((*C.GtkSourcePrintCompositor)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// GetWrapMode is a wrapper around the C function gtk_source_print_compositor_get_wrap_mode.
func (recv *PrintCompositor) GetWrapMode() gtk.WrapMode {
	retC := C.gtk_source_print_compositor_get_wrap_mode((*C.GtkSourcePrintCompositor)(recv.native))
	retGo := (gtk.WrapMode)(retC)

	return retGo
}

// Paginate is a wrapper around the C function gtk_source_print_compositor_paginate.
func (recv *PrintCompositor) Paginate(context *gtk.PrintContext) bool {
	c_context := (*C.GtkPrintContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkPrintContext)(context.ToC())
	}

	retC := C.gtk_source_print_compositor_paginate((*C.GtkSourcePrintCompositor)(recv.native), c_context)
	retGo := retC == C.TRUE

	return retGo
}

// SetBodyFontName is a wrapper around the C function gtk_source_print_compositor_set_body_font_name.
func (recv *PrintCompositor) SetBodyFontName(fontName string) {
	c_font_name := C.CString(fontName)
	defer C.free(unsafe.Pointer(c_font_name))

	C.gtk_source_print_compositor_set_body_font_name((*C.GtkSourcePrintCompositor)(recv.native), c_font_name)

	return
}

// SetBottomMargin is a wrapper around the C function gtk_source_print_compositor_set_bottom_margin.
func (recv *PrintCompositor) SetBottomMargin(margin float64, unit gtk.Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_source_print_compositor_set_bottom_margin((*C.GtkSourcePrintCompositor)(recv.native), c_margin, c_unit)

	return
}

// SetFooterFontName is a wrapper around the C function gtk_source_print_compositor_set_footer_font_name.
func (recv *PrintCompositor) SetFooterFontName(fontName string) {
	c_font_name := C.CString(fontName)
	defer C.free(unsafe.Pointer(c_font_name))

	C.gtk_source_print_compositor_set_footer_font_name((*C.GtkSourcePrintCompositor)(recv.native), c_font_name)

	return
}

// SetFooterFormat is a wrapper around the C function gtk_source_print_compositor_set_footer_format.
func (recv *PrintCompositor) SetFooterFormat(separator bool, left string, center string, right string) {
	c_separator :=
		boolToGboolean(separator)

	c_left := C.CString(left)
	defer C.free(unsafe.Pointer(c_left))

	c_center := C.CString(center)
	defer C.free(unsafe.Pointer(c_center))

	c_right := C.CString(right)
	defer C.free(unsafe.Pointer(c_right))

	C.gtk_source_print_compositor_set_footer_format((*C.GtkSourcePrintCompositor)(recv.native), c_separator, c_left, c_center, c_right)

	return
}

// SetHeaderFontName is a wrapper around the C function gtk_source_print_compositor_set_header_font_name.
func (recv *PrintCompositor) SetHeaderFontName(fontName string) {
	c_font_name := C.CString(fontName)
	defer C.free(unsafe.Pointer(c_font_name))

	C.gtk_source_print_compositor_set_header_font_name((*C.GtkSourcePrintCompositor)(recv.native), c_font_name)

	return
}

// SetHeaderFormat is a wrapper around the C function gtk_source_print_compositor_set_header_format.
func (recv *PrintCompositor) SetHeaderFormat(separator bool, left string, center string, right string) {
	c_separator :=
		boolToGboolean(separator)

	c_left := C.CString(left)
	defer C.free(unsafe.Pointer(c_left))

	c_center := C.CString(center)
	defer C.free(unsafe.Pointer(c_center))

	c_right := C.CString(right)
	defer C.free(unsafe.Pointer(c_right))

	C.gtk_source_print_compositor_set_header_format((*C.GtkSourcePrintCompositor)(recv.native), c_separator, c_left, c_center, c_right)

	return
}

// SetHighlightSyntax is a wrapper around the C function gtk_source_print_compositor_set_highlight_syntax.
func (recv *PrintCompositor) SetHighlightSyntax(highlight bool) {
	c_highlight :=
		boolToGboolean(highlight)

	C.gtk_source_print_compositor_set_highlight_syntax((*C.GtkSourcePrintCompositor)(recv.native), c_highlight)

	return
}

// SetLeftMargin is a wrapper around the C function gtk_source_print_compositor_set_left_margin.
func (recv *PrintCompositor) SetLeftMargin(margin float64, unit gtk.Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_source_print_compositor_set_left_margin((*C.GtkSourcePrintCompositor)(recv.native), c_margin, c_unit)

	return
}

// SetLineNumbersFontName is a wrapper around the C function gtk_source_print_compositor_set_line_numbers_font_name.
func (recv *PrintCompositor) SetLineNumbersFontName(fontName string) {
	c_font_name := C.CString(fontName)
	defer C.free(unsafe.Pointer(c_font_name))

	C.gtk_source_print_compositor_set_line_numbers_font_name((*C.GtkSourcePrintCompositor)(recv.native), c_font_name)

	return
}

// SetPrintFooter is a wrapper around the C function gtk_source_print_compositor_set_print_footer.
func (recv *PrintCompositor) SetPrintFooter(print bool) {
	c_print :=
		boolToGboolean(print)

	C.gtk_source_print_compositor_set_print_footer((*C.GtkSourcePrintCompositor)(recv.native), c_print)

	return
}

// SetPrintHeader is a wrapper around the C function gtk_source_print_compositor_set_print_header.
func (recv *PrintCompositor) SetPrintHeader(print bool) {
	c_print :=
		boolToGboolean(print)

	C.gtk_source_print_compositor_set_print_header((*C.GtkSourcePrintCompositor)(recv.native), c_print)

	return
}

// SetPrintLineNumbers is a wrapper around the C function gtk_source_print_compositor_set_print_line_numbers.
func (recv *PrintCompositor) SetPrintLineNumbers(interval uint32) {
	c_interval := (C.guint)(interval)

	C.gtk_source_print_compositor_set_print_line_numbers((*C.GtkSourcePrintCompositor)(recv.native), c_interval)

	return
}

// SetRightMargin is a wrapper around the C function gtk_source_print_compositor_set_right_margin.
func (recv *PrintCompositor) SetRightMargin(margin float64, unit gtk.Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_source_print_compositor_set_right_margin((*C.GtkSourcePrintCompositor)(recv.native), c_margin, c_unit)

	return
}

// SetTabWidth is a wrapper around the C function gtk_source_print_compositor_set_tab_width.
func (recv *PrintCompositor) SetTabWidth(width uint32) {
	c_width := (C.guint)(width)

	C.gtk_source_print_compositor_set_tab_width((*C.GtkSourcePrintCompositor)(recv.native), c_width)

	return
}

// SetTopMargin is a wrapper around the C function gtk_source_print_compositor_set_top_margin.
func (recv *PrintCompositor) SetTopMargin(margin float64, unit gtk.Unit) {
	c_margin := (C.gdouble)(margin)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_source_print_compositor_set_top_margin((*C.GtkSourcePrintCompositor)(recv.native), c_margin, c_unit)

	return
}

// SetWrapMode is a wrapper around the C function gtk_source_print_compositor_set_wrap_mode.
func (recv *PrintCompositor) SetWrapMode(wrapMode gtk.WrapMode) {
	c_wrap_mode := (C.GtkWrapMode)(wrapMode)

	C.gtk_source_print_compositor_set_wrap_mode((*C.GtkSourcePrintCompositor)(recv.native), c_wrap_mode)

	return
}

// Region is a wrapper around the C record GtkSourceRegion.
type Region struct {
	native *C.GtkSourceRegion
	// parent_instance : record
}

func RegionNewFromC(u unsafe.Pointer) *Region {
	c := (*C.GtkSourceRegion)(u)
	if c == nil {
		return nil
	}

	g := &Region{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Region) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Region) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Region with another Region, and returns true if they represent the same GObject.
func (recv *Region) Equals(other *Region) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Region) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Region.
// Exercise care, as this is a potentially dangerous function if the Object is not a Region.
func CastToRegion(object *gobject.Object) *Region {
	return RegionNewFromC(object.ToC())
}

// RegionNew is a wrapper around the C function gtk_source_region_new.
func RegionNew(buffer *gtk.TextBuffer) *Region {
	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	retC := C.gtk_source_region_new(c_buffer)
	retGo := RegionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddRegion is a wrapper around the C function gtk_source_region_add_region.
func (recv *Region) AddRegion(regionToAdd *Region) {
	c_region_to_add := (*C.GtkSourceRegion)(C.NULL)
	if regionToAdd != nil {
		c_region_to_add = (*C.GtkSourceRegion)(regionToAdd.ToC())
	}

	C.gtk_source_region_add_region((*C.GtkSourceRegion)(recv.native), c_region_to_add)

	return
}

// AddSubregion is a wrapper around the C function gtk_source_region_add_subregion.
func (recv *Region) AddSubregion(Start *gtk.TextIter, End *gtk.TextIter) {
	c__start := (*C.GtkTextIter)(C.NULL)
	if Start != nil {
		c__start = (*C.GtkTextIter)(Start.ToC())
	}

	c__end := (*C.GtkTextIter)(C.NULL)
	if End != nil {
		c__end = (*C.GtkTextIter)(End.ToC())
	}

	C.gtk_source_region_add_subregion((*C.GtkSourceRegion)(recv.native), c__start, c__end)

	return
}

// GetBounds is a wrapper around the C function gtk_source_region_get_bounds.
func (recv *Region) GetBounds() (bool, *gtk.TextIter, *gtk.TextIter) {
	var c_start C.GtkTextIter

	var c_end C.GtkTextIter

	retC := C.gtk_source_region_get_bounds((*C.GtkSourceRegion)(recv.native), &c_start, &c_end)
	retGo := retC == C.TRUE

	start := gtk.TextIterNewFromC(unsafe.Pointer(&c_start))

	end := gtk.TextIterNewFromC(unsafe.Pointer(&c_end))

	return retGo, start, end
}

// GetBuffer is a wrapper around the C function gtk_source_region_get_buffer.
func (recv *Region) GetBuffer() *gtk.TextBuffer {
	retC := C.gtk_source_region_get_buffer((*C.GtkSourceRegion)(recv.native))
	var retGo (*gtk.TextBuffer)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gtk.TextBufferNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStartRegionIter is a wrapper around the C function gtk_source_region_get_start_region_iter.
func (recv *Region) GetStartRegionIter() *RegionIter {
	var c_iter C.GtkSourceRegionIter

	C.gtk_source_region_get_start_region_iter((*C.GtkSourceRegion)(recv.native), &c_iter)

	iter := RegionIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// IntersectRegion is a wrapper around the C function gtk_source_region_intersect_region.
func (recv *Region) IntersectRegion(region2 *Region) *Region {
	c_region2 := (*C.GtkSourceRegion)(C.NULL)
	if region2 != nil {
		c_region2 = (*C.GtkSourceRegion)(region2.ToC())
	}

	retC := C.gtk_source_region_intersect_region((*C.GtkSourceRegion)(recv.native), c_region2)
	var retGo (*Region)
	if retC == nil {
		retGo = nil
	} else {
		retGo = RegionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// IntersectSubregion is a wrapper around the C function gtk_source_region_intersect_subregion.
func (recv *Region) IntersectSubregion(Start *gtk.TextIter, End *gtk.TextIter) *Region {
	c__start := (*C.GtkTextIter)(C.NULL)
	if Start != nil {
		c__start = (*C.GtkTextIter)(Start.ToC())
	}

	c__end := (*C.GtkTextIter)(C.NULL)
	if End != nil {
		c__end = (*C.GtkTextIter)(End.ToC())
	}

	retC := C.gtk_source_region_intersect_subregion((*C.GtkSourceRegion)(recv.native), c__start, c__end)
	var retGo (*Region)
	if retC == nil {
		retGo = nil
	} else {
		retGo = RegionNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// IsEmpty is a wrapper around the C function gtk_source_region_is_empty.
func (recv *Region) IsEmpty() bool {
	retC := C.gtk_source_region_is_empty((*C.GtkSourceRegion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SubtractRegion is a wrapper around the C function gtk_source_region_subtract_region.
func (recv *Region) SubtractRegion(regionToSubtract *Region) {
	c_region_to_subtract := (*C.GtkSourceRegion)(C.NULL)
	if regionToSubtract != nil {
		c_region_to_subtract = (*C.GtkSourceRegion)(regionToSubtract.ToC())
	}

	C.gtk_source_region_subtract_region((*C.GtkSourceRegion)(recv.native), c_region_to_subtract)

	return
}

// SubtractSubregion is a wrapper around the C function gtk_source_region_subtract_subregion.
func (recv *Region) SubtractSubregion(Start *gtk.TextIter, End *gtk.TextIter) {
	c__start := (*C.GtkTextIter)(C.NULL)
	if Start != nil {
		c__start = (*C.GtkTextIter)(Start.ToC())
	}

	c__end := (*C.GtkTextIter)(C.NULL)
	if End != nil {
		c__end = (*C.GtkTextIter)(End.ToC())
	}

	C.gtk_source_region_subtract_subregion((*C.GtkSourceRegion)(recv.native), c__start, c__end)

	return
}

// ToString is a wrapper around the C function gtk_source_region_to_string.
func (recv *Region) ToString() string {
	retC := C.gtk_source_region_to_string((*C.GtkSourceRegion)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// SearchContext is a wrapper around the C record GtkSourceSearchContext.
type SearchContext struct {
	native *C.GtkSourceSearchContext
	// parent : record
	// priv : record
}

func SearchContextNewFromC(u unsafe.Pointer) *SearchContext {
	c := (*C.GtkSourceSearchContext)(u)
	if c == nil {
		return nil
	}

	g := &SearchContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SearchContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SearchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SearchContext with another SearchContext, and returns true if they represent the same GObject.
func (recv *SearchContext) Equals(other *SearchContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SearchContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SearchContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a SearchContext.
func CastToSearchContext(object *gobject.Object) *SearchContext {
	return SearchContextNewFromC(object.ToC())
}

// SearchContextNew is a wrapper around the C function gtk_source_search_context_new.
func SearchContextNew(buffer *Buffer, settings *SearchSettings) *SearchContext {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	c_settings := (*C.GtkSourceSearchSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSourceSearchSettings)(settings.ToC())
	}

	retC := C.gtk_source_search_context_new(c_buffer, c_settings)
	retGo := SearchContextNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Backward is a wrapper around the C function gtk_source_search_context_backward.
func (recv *SearchContext) Backward(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	retC := C.gtk_source_search_context_backward((*C.GtkSourceSearchContext)(recv.native), c_iter, &c_match_start, &c_match_end)
	retGo := retC == C.TRUE

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	return retGo, matchStart, matchEnd
}

// Backward2 is a wrapper around the C function gtk_source_search_context_backward2.
func (recv *SearchContext) Backward2(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter, bool) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	var c_has_wrapped_around C.gboolean

	retC := C.gtk_source_search_context_backward2((*C.GtkSourceSearchContext)(recv.native), c_iter, &c_match_start, &c_match_end, &c_has_wrapped_around)
	retGo := retC == C.TRUE

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	hasWrappedAround := c_has_wrapped_around == C.TRUE

	return retGo, matchStart, matchEnd, hasWrappedAround
}

// Unsupported : gtk_source_search_context_backward_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// BackwardFinish is a wrapper around the C function gtk_source_search_context_backward_finish.
func (recv *SearchContext) BackwardFinish(result *gio.AsyncResult) (bool, *gtk.TextIter, *gtk.TextIter, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_backward_finish((*C.GtkSourceSearchContext)(recv.native), c_result, &c_match_start, &c_match_end, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	return retGo, matchStart, matchEnd, goError
}

// BackwardFinish2 is a wrapper around the C function gtk_source_search_context_backward_finish2.
func (recv *SearchContext) BackwardFinish2(result *gio.AsyncResult) (bool, *gtk.TextIter, *gtk.TextIter, bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	var c_has_wrapped_around C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_backward_finish2((*C.GtkSourceSearchContext)(recv.native), c_result, &c_match_start, &c_match_end, &c_has_wrapped_around, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	hasWrappedAround := c_has_wrapped_around == C.TRUE

	return retGo, matchStart, matchEnd, hasWrappedAround, goError
}

// Forward is a wrapper around the C function gtk_source_search_context_forward.
func (recv *SearchContext) Forward(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	retC := C.gtk_source_search_context_forward((*C.GtkSourceSearchContext)(recv.native), c_iter, &c_match_start, &c_match_end)
	retGo := retC == C.TRUE

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	return retGo, matchStart, matchEnd
}

// Forward2 is a wrapper around the C function gtk_source_search_context_forward2.
func (recv *SearchContext) Forward2(iter *gtk.TextIter) (bool, *gtk.TextIter, *gtk.TextIter, bool) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	var c_has_wrapped_around C.gboolean

	retC := C.gtk_source_search_context_forward2((*C.GtkSourceSearchContext)(recv.native), c_iter, &c_match_start, &c_match_end, &c_has_wrapped_around)
	retGo := retC == C.TRUE

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	hasWrappedAround := c_has_wrapped_around == C.TRUE

	return retGo, matchStart, matchEnd, hasWrappedAround
}

// Unsupported : gtk_source_search_context_forward_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ForwardFinish is a wrapper around the C function gtk_source_search_context_forward_finish.
func (recv *SearchContext) ForwardFinish(result *gio.AsyncResult) (bool, *gtk.TextIter, *gtk.TextIter, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_forward_finish((*C.GtkSourceSearchContext)(recv.native), c_result, &c_match_start, &c_match_end, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	return retGo, matchStart, matchEnd, goError
}

// ForwardFinish2 is a wrapper around the C function gtk_source_search_context_forward_finish2.
func (recv *SearchContext) ForwardFinish2(result *gio.AsyncResult) (bool, *gtk.TextIter, *gtk.TextIter, bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_match_start C.GtkTextIter

	var c_match_end C.GtkTextIter

	var c_has_wrapped_around C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_forward_finish2((*C.GtkSourceSearchContext)(recv.native), c_result, &c_match_start, &c_match_end, &c_has_wrapped_around, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	matchStart := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_start))

	matchEnd := gtk.TextIterNewFromC(unsafe.Pointer(&c_match_end))

	hasWrappedAround := c_has_wrapped_around == C.TRUE

	return retGo, matchStart, matchEnd, hasWrappedAround, goError
}

// GetBuffer is a wrapper around the C function gtk_source_search_context_get_buffer.
func (recv *SearchContext) GetBuffer() *Buffer {
	retC := C.gtk_source_search_context_get_buffer((*C.GtkSourceSearchContext)(recv.native))
	retGo := BufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHighlight is a wrapper around the C function gtk_source_search_context_get_highlight.
func (recv *SearchContext) GetHighlight() bool {
	retC := C.gtk_source_search_context_get_highlight((*C.GtkSourceSearchContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMatchStyle is a wrapper around the C function gtk_source_search_context_get_match_style.
func (recv *SearchContext) GetMatchStyle() *Style {
	retC := C.gtk_source_search_context_get_match_style((*C.GtkSourceSearchContext)(recv.native))
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOccurrencePosition is a wrapper around the C function gtk_source_search_context_get_occurrence_position.
func (recv *SearchContext) GetOccurrencePosition(matchStart *gtk.TextIter, matchEnd *gtk.TextIter) int32 {
	c_match_start := (*C.GtkTextIter)(C.NULL)
	if matchStart != nil {
		c_match_start = (*C.GtkTextIter)(matchStart.ToC())
	}

	c_match_end := (*C.GtkTextIter)(C.NULL)
	if matchEnd != nil {
		c_match_end = (*C.GtkTextIter)(matchEnd.ToC())
	}

	retC := C.gtk_source_search_context_get_occurrence_position((*C.GtkSourceSearchContext)(recv.native), c_match_start, c_match_end)
	retGo := (int32)(retC)

	return retGo
}

// GetOccurrencesCount is a wrapper around the C function gtk_source_search_context_get_occurrences_count.
func (recv *SearchContext) GetOccurrencesCount() int32 {
	retC := C.gtk_source_search_context_get_occurrences_count((*C.GtkSourceSearchContext)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRegexError is a wrapper around the C function gtk_source_search_context_get_regex_error.
func (recv *SearchContext) GetRegexError() *glib.Error {
	retC := C.gtk_source_search_context_get_regex_error((*C.GtkSourceSearchContext)(recv.native))
	var retGo (*glib.Error)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.ErrorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSettings is a wrapper around the C function gtk_source_search_context_get_settings.
func (recv *SearchContext) GetSettings() *SearchSettings {
	retC := C.gtk_source_search_context_get_settings((*C.GtkSourceSearchContext)(recv.native))
	retGo := SearchSettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Replace is a wrapper around the C function gtk_source_search_context_replace.
func (recv *SearchContext) Replace(matchStart *gtk.TextIter, matchEnd *gtk.TextIter, replace string, replaceLength int32) (bool, error) {
	c_match_start := (*C.GtkTextIter)(C.NULL)
	if matchStart != nil {
		c_match_start = (*C.GtkTextIter)(matchStart.ToC())
	}

	c_match_end := (*C.GtkTextIter)(C.NULL)
	if matchEnd != nil {
		c_match_end = (*C.GtkTextIter)(matchEnd.ToC())
	}

	c_replace := C.CString(replace)
	defer C.free(unsafe.Pointer(c_replace))

	c_replace_length := (C.gint)(replaceLength)

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_replace((*C.GtkSourceSearchContext)(recv.native), c_match_start, c_match_end, c_replace, c_replace_length, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Replace2 is a wrapper around the C function gtk_source_search_context_replace2.
func (recv *SearchContext) Replace2(matchStart *gtk.TextIter, matchEnd *gtk.TextIter, replace string, replaceLength int32) (bool, error) {
	c_match_start := (*C.GtkTextIter)(C.NULL)
	if matchStart != nil {
		c_match_start = (*C.GtkTextIter)(matchStart.ToC())
	}

	c_match_end := (*C.GtkTextIter)(C.NULL)
	if matchEnd != nil {
		c_match_end = (*C.GtkTextIter)(matchEnd.ToC())
	}

	c_replace := C.CString(replace)
	defer C.free(unsafe.Pointer(c_replace))

	c_replace_length := (C.gint)(replaceLength)

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_replace2((*C.GtkSourceSearchContext)(recv.native), c_match_start, c_match_end, c_replace, c_replace_length, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReplaceAll is a wrapper around the C function gtk_source_search_context_replace_all.
func (recv *SearchContext) ReplaceAll(replace string, replaceLength int32) (uint32, error) {
	c_replace := C.CString(replace)
	defer C.free(unsafe.Pointer(c_replace))

	c_replace_length := (C.gint)(replaceLength)

	var cThrowableError *C.GError

	retC := C.gtk_source_search_context_replace_all((*C.GtkSourceSearchContext)(recv.native), c_replace, c_replace_length, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetHighlight is a wrapper around the C function gtk_source_search_context_set_highlight.
func (recv *SearchContext) SetHighlight(highlight bool) {
	c_highlight :=
		boolToGboolean(highlight)

	C.gtk_source_search_context_set_highlight((*C.GtkSourceSearchContext)(recv.native), c_highlight)

	return
}

// SetMatchStyle is a wrapper around the C function gtk_source_search_context_set_match_style.
func (recv *SearchContext) SetMatchStyle(matchStyle *Style) {
	c_match_style := (*C.GtkSourceStyle)(C.NULL)
	if matchStyle != nil {
		c_match_style = (*C.GtkSourceStyle)(matchStyle.ToC())
	}

	C.gtk_source_search_context_set_match_style((*C.GtkSourceSearchContext)(recv.native), c_match_style)

	return
}

// SetSettings is a wrapper around the C function gtk_source_search_context_set_settings.
func (recv *SearchContext) SetSettings(settings *SearchSettings) {
	c_settings := (*C.GtkSourceSearchSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSourceSearchSettings)(settings.ToC())
	}

	C.gtk_source_search_context_set_settings((*C.GtkSourceSearchContext)(recv.native), c_settings)

	return
}

// SearchSettings is a wrapper around the C record GtkSourceSearchSettings.
type SearchSettings struct {
	native *C.GtkSourceSearchSettings
	// parent : record
	// priv : record
}

func SearchSettingsNewFromC(u unsafe.Pointer) *SearchSettings {
	c := (*C.GtkSourceSearchSettings)(u)
	if c == nil {
		return nil
	}

	g := &SearchSettings{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SearchSettings) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SearchSettings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SearchSettings with another SearchSettings, and returns true if they represent the same GObject.
func (recv *SearchSettings) Equals(other *SearchSettings) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SearchSettings) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SearchSettings.
// Exercise care, as this is a potentially dangerous function if the Object is not a SearchSettings.
func CastToSearchSettings(object *gobject.Object) *SearchSettings {
	return SearchSettingsNewFromC(object.ToC())
}

// SearchSettingsNew is a wrapper around the C function gtk_source_search_settings_new.
func SearchSettingsNew() *SearchSettings {
	retC := C.gtk_source_search_settings_new()
	retGo := SearchSettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetAtWordBoundaries is a wrapper around the C function gtk_source_search_settings_get_at_word_boundaries.
func (recv *SearchSettings) GetAtWordBoundaries() bool {
	retC := C.gtk_source_search_settings_get_at_word_boundaries((*C.GtkSourceSearchSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCaseSensitive is a wrapper around the C function gtk_source_search_settings_get_case_sensitive.
func (recv *SearchSettings) GetCaseSensitive() bool {
	retC := C.gtk_source_search_settings_get_case_sensitive((*C.GtkSourceSearchSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRegexEnabled is a wrapper around the C function gtk_source_search_settings_get_regex_enabled.
func (recv *SearchSettings) GetRegexEnabled() bool {
	retC := C.gtk_source_search_settings_get_regex_enabled((*C.GtkSourceSearchSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSearchText is a wrapper around the C function gtk_source_search_settings_get_search_text.
func (recv *SearchSettings) GetSearchText() string {
	retC := C.gtk_source_search_settings_get_search_text((*C.GtkSourceSearchSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWrapAround is a wrapper around the C function gtk_source_search_settings_get_wrap_around.
func (recv *SearchSettings) GetWrapAround() bool {
	retC := C.gtk_source_search_settings_get_wrap_around((*C.GtkSourceSearchSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAtWordBoundaries is a wrapper around the C function gtk_source_search_settings_set_at_word_boundaries.
func (recv *SearchSettings) SetAtWordBoundaries(atWordBoundaries bool) {
	c_at_word_boundaries :=
		boolToGboolean(atWordBoundaries)

	C.gtk_source_search_settings_set_at_word_boundaries((*C.GtkSourceSearchSettings)(recv.native), c_at_word_boundaries)

	return
}

// SetCaseSensitive is a wrapper around the C function gtk_source_search_settings_set_case_sensitive.
func (recv *SearchSettings) SetCaseSensitive(caseSensitive bool) {
	c_case_sensitive :=
		boolToGboolean(caseSensitive)

	C.gtk_source_search_settings_set_case_sensitive((*C.GtkSourceSearchSettings)(recv.native), c_case_sensitive)

	return
}

// SetRegexEnabled is a wrapper around the C function gtk_source_search_settings_set_regex_enabled.
func (recv *SearchSettings) SetRegexEnabled(regexEnabled bool) {
	c_regex_enabled :=
		boolToGboolean(regexEnabled)

	C.gtk_source_search_settings_set_regex_enabled((*C.GtkSourceSearchSettings)(recv.native), c_regex_enabled)

	return
}

// SetSearchText is a wrapper around the C function gtk_source_search_settings_set_search_text.
func (recv *SearchSettings) SetSearchText(searchText string) {
	c_search_text := C.CString(searchText)
	defer C.free(unsafe.Pointer(c_search_text))

	C.gtk_source_search_settings_set_search_text((*C.GtkSourceSearchSettings)(recv.native), c_search_text)

	return
}

// SetWrapAround is a wrapper around the C function gtk_source_search_settings_set_wrap_around.
func (recv *SearchSettings) SetWrapAround(wrapAround bool) {
	c_wrap_around :=
		boolToGboolean(wrapAround)

	C.gtk_source_search_settings_set_wrap_around((*C.GtkSourceSearchSettings)(recv.native), c_wrap_around)

	return
}

// Style is a wrapper around the C record GtkSourceStyle.
type Style struct {
	native *C.GtkSourceStyle
}

func StyleNewFromC(u unsafe.Pointer) *Style {
	c := (*C.GtkSourceStyle)(u)
	if c == nil {
		return nil
	}

	g := &Style{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Style) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Style) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Style with another Style, and returns true if they represent the same GObject.
func (recv *Style) Equals(other *Style) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Style) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Style.
// Exercise care, as this is a potentially dangerous function if the Object is not a Style.
func CastToStyle(object *gobject.Object) *Style {
	return StyleNewFromC(object.ToC())
}

// Apply is a wrapper around the C function gtk_source_style_apply.
func (recv *Style) Apply(tag *gtk.TextTag) {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	C.gtk_source_style_apply((*C.GtkSourceStyle)(recv.native), c_tag)

	return
}

// Copy is a wrapper around the C function gtk_source_style_copy.
func (recv *Style) Copy() *Style {
	retC := C.gtk_source_style_copy((*C.GtkSourceStyle)(recv.native))
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StyleScheme is a wrapper around the C record GtkSourceStyleScheme.
type StyleScheme struct {
	native *C.GtkSourceStyleScheme
	// base : record
	// priv : record
}

func StyleSchemeNewFromC(u unsafe.Pointer) *StyleScheme {
	c := (*C.GtkSourceStyleScheme)(u)
	if c == nil {
		return nil
	}

	g := &StyleScheme{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StyleScheme) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StyleScheme) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleScheme with another StyleScheme, and returns true if they represent the same GObject.
func (recv *StyleScheme) Equals(other *StyleScheme) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *StyleScheme) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to StyleScheme.
// Exercise care, as this is a potentially dangerous function if the Object is not a StyleScheme.
func CastToStyleScheme(object *gobject.Object) *StyleScheme {
	return StyleSchemeNewFromC(object.ToC())
}

// GetAuthors is a wrapper around the C function gtk_source_style_scheme_get_authors.
func (recv *StyleScheme) GetAuthors() []string {
	retC := C.gtk_source_style_scheme_get_authors((*C.GtkSourceStyleScheme)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetDescription is a wrapper around the C function gtk_source_style_scheme_get_description.
func (recv *StyleScheme) GetDescription() string {
	retC := C.gtk_source_style_scheme_get_description((*C.GtkSourceStyleScheme)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetFilename is a wrapper around the C function gtk_source_style_scheme_get_filename.
func (recv *StyleScheme) GetFilename() string {
	retC := C.gtk_source_style_scheme_get_filename((*C.GtkSourceStyleScheme)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetId is a wrapper around the C function gtk_source_style_scheme_get_id.
func (recv *StyleScheme) GetId() string {
	retC := C.gtk_source_style_scheme_get_id((*C.GtkSourceStyleScheme)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function gtk_source_style_scheme_get_name.
func (recv *StyleScheme) GetName() string {
	retC := C.gtk_source_style_scheme_get_name((*C.GtkSourceStyleScheme)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStyle is a wrapper around the C function gtk_source_style_scheme_get_style.
func (recv *StyleScheme) GetStyle(styleId string) *Style {
	c_style_id := C.CString(styleId)
	defer C.free(unsafe.Pointer(c_style_id))

	retC := C.gtk_source_style_scheme_get_style((*C.GtkSourceStyleScheme)(recv.native), c_style_id)
	var retGo (*Style)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// StyleSchemeChooserButton is a wrapper around the C record GtkSourceStyleSchemeChooserButton.
type StyleSchemeChooserButton struct {
	native *C.GtkSourceStyleSchemeChooserButton
	// parent : record
}

func StyleSchemeChooserButtonNewFromC(u unsafe.Pointer) *StyleSchemeChooserButton {
	c := (*C.GtkSourceStyleSchemeChooserButton)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeChooserButton{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StyleSchemeChooserButton) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StyleSchemeChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeChooserButton with another StyleSchemeChooserButton, and returns true if they represent the same GObject.
func (recv *StyleSchemeChooserButton) Equals(other *StyleSchemeChooserButton) bool {
	return other.ToC() == recv.ToC()
}

// Button upcasts to *Button
func (recv *StyleSchemeChooserButton) Button() *gtk.Button {
	return gtk.ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *StyleSchemeChooserButton) Bin() *gtk.Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *StyleSchemeChooserButton) Container() *gtk.Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *StyleSchemeChooserButton) Widget() *gtk.Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *StyleSchemeChooserButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *StyleSchemeChooserButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitrary Object to StyleSchemeChooserButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a StyleSchemeChooserButton.
func CastToStyleSchemeChooserButton(object *gobject.Object) *StyleSchemeChooserButton {
	return StyleSchemeChooserButtonNewFromC(object.ToC())
}

// StyleSchemeChooserButtonNew is a wrapper around the C function gtk_source_style_scheme_chooser_button_new.
func StyleSchemeChooserButtonNew() *StyleSchemeChooserButton {
	retC := C.gtk_source_style_scheme_chooser_button_new()
	retGo := StyleSchemeChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImplementorIface returns the ImplementorIface interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) Actionable() *gtk.Actionable {
	return gtk.ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) Activatable() *gtk.Activatable {
	return gtk.ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// StyleSchemeChooser returns the StyleSchemeChooser interface implemented by StyleSchemeChooserButton
func (recv *StyleSchemeChooserButton) StyleSchemeChooser() *StyleSchemeChooser {
	return StyleSchemeChooserNewFromC(recv.ToC())
}

// StyleSchemeChooserWidget is a wrapper around the C record GtkSourceStyleSchemeChooserWidget.
type StyleSchemeChooserWidget struct {
	native *C.GtkSourceStyleSchemeChooserWidget
	// parent : record
}

func StyleSchemeChooserWidgetNewFromC(u unsafe.Pointer) *StyleSchemeChooserWidget {
	c := (*C.GtkSourceStyleSchemeChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeChooserWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StyleSchemeChooserWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StyleSchemeChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeChooserWidget with another StyleSchemeChooserWidget, and returns true if they represent the same GObject.
func (recv *StyleSchemeChooserWidget) Equals(other *StyleSchemeChooserWidget) bool {
	return other.ToC() == recv.ToC()
}

// Bin upcasts to *Bin
func (recv *StyleSchemeChooserWidget) Bin() *gtk.Bin {
	return gtk.BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *StyleSchemeChooserWidget) Container() *gtk.Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *StyleSchemeChooserWidget) Widget() *gtk.Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *StyleSchemeChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *StyleSchemeChooserWidget) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitrary Object to StyleSchemeChooserWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a StyleSchemeChooserWidget.
func CastToStyleSchemeChooserWidget(object *gobject.Object) *StyleSchemeChooserWidget {
	return StyleSchemeChooserWidgetNewFromC(object.ToC())
}

// StyleSchemeChooserWidgetNew is a wrapper around the C function gtk_source_style_scheme_chooser_widget_new.
func StyleSchemeChooserWidgetNew() *StyleSchemeChooserWidget {
	retC := C.gtk_source_style_scheme_chooser_widget_new()
	retGo := StyleSchemeChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImplementorIface returns the ImplementorIface interface implemented by StyleSchemeChooserWidget
func (recv *StyleSchemeChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by StyleSchemeChooserWidget
func (recv *StyleSchemeChooserWidget) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// StyleSchemeChooser returns the StyleSchemeChooser interface implemented by StyleSchemeChooserWidget
func (recv *StyleSchemeChooserWidget) StyleSchemeChooser() *StyleSchemeChooser {
	return StyleSchemeChooserNewFromC(recv.ToC())
}

// StyleSchemeManager is a wrapper around the C record GtkSourceStyleSchemeManager.
type StyleSchemeManager struct {
	native *C.GtkSourceStyleSchemeManager
	// parent : record
	// priv : record
}

func StyleSchemeManagerNewFromC(u unsafe.Pointer) *StyleSchemeManager {
	c := (*C.GtkSourceStyleSchemeManager)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StyleSchemeManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StyleSchemeManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeManager with another StyleSchemeManager, and returns true if they represent the same GObject.
func (recv *StyleSchemeManager) Equals(other *StyleSchemeManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *StyleSchemeManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to StyleSchemeManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a StyleSchemeManager.
func CastToStyleSchemeManager(object *gobject.Object) *StyleSchemeManager {
	return StyleSchemeManagerNewFromC(object.ToC())
}

// StyleSchemeManagerNew is a wrapper around the C function gtk_source_style_scheme_manager_new.
func StyleSchemeManagerNew() *StyleSchemeManager {
	retC := C.gtk_source_style_scheme_manager_new()
	retGo := StyleSchemeManagerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// StyleSchemeManagerGetDefault is a wrapper around the C function gtk_source_style_scheme_manager_get_default.
func StyleSchemeManagerGetDefault() *StyleSchemeManager {
	retC := C.gtk_source_style_scheme_manager_get_default()
	retGo := StyleSchemeManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendSearchPath is a wrapper around the C function gtk_source_style_scheme_manager_append_search_path.
func (recv *StyleSchemeManager) AppendSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_source_style_scheme_manager_append_search_path((*C.GtkSourceStyleSchemeManager)(recv.native), c_path)

	return
}

// ForceRescan is a wrapper around the C function gtk_source_style_scheme_manager_force_rescan.
func (recv *StyleSchemeManager) ForceRescan() {
	C.gtk_source_style_scheme_manager_force_rescan((*C.GtkSourceStyleSchemeManager)(recv.native))

	return
}

// GetScheme is a wrapper around the C function gtk_source_style_scheme_manager_get_scheme.
func (recv *StyleSchemeManager) GetScheme(schemeId string) *StyleScheme {
	c_scheme_id := C.CString(schemeId)
	defer C.free(unsafe.Pointer(c_scheme_id))

	retC := C.gtk_source_style_scheme_manager_get_scheme((*C.GtkSourceStyleSchemeManager)(recv.native), c_scheme_id)
	retGo := StyleSchemeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSchemeIds is a wrapper around the C function gtk_source_style_scheme_manager_get_scheme_ids.
func (recv *StyleSchemeManager) GetSchemeIds() []string {
	retC := C.gtk_source_style_scheme_manager_get_scheme_ids((*C.GtkSourceStyleSchemeManager)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetSearchPath is a wrapper around the C function gtk_source_style_scheme_manager_get_search_path.
func (recv *StyleSchemeManager) GetSearchPath() []string {
	retC := C.gtk_source_style_scheme_manager_get_search_path((*C.GtkSourceStyleSchemeManager)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// PrependSearchPath is a wrapper around the C function gtk_source_style_scheme_manager_prepend_search_path.
func (recv *StyleSchemeManager) PrependSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_source_style_scheme_manager_prepend_search_path((*C.GtkSourceStyleSchemeManager)(recv.native), c_path)

	return
}

// SetSearchPath is a wrapper around the C function gtk_source_style_scheme_manager_set_search_path.
func (recv *StyleSchemeManager) SetSearchPath(path []string) {
	c_path_array := make([]*C.gchar, len(path)+1, len(path)+1)
	for i, item := range path {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_path_array[i] = c
	}
	c_path_array[len(path)] = nil
	c_path_arrayPtr := &c_path_array[0]
	c_path := (**C.gchar)(unsafe.Pointer(c_path_arrayPtr))

	C.gtk_source_style_scheme_manager_set_search_path((*C.GtkSourceStyleSchemeManager)(recv.native), c_path)

	return
}

// Tag is a wrapper around the C record GtkSourceTag.
type Tag struct {
	native *C.GtkSourceTag
	// parent_instance : record
}

func TagNewFromC(u unsafe.Pointer) *Tag {
	c := (*C.GtkSourceTag)(u)
	if c == nil {
		return nil
	}

	g := &Tag{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Tag) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Tag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Tag with another Tag, and returns true if they represent the same GObject.
func (recv *Tag) Equals(other *Tag) bool {
	return other.ToC() == recv.ToC()
}

// TextTag upcasts to *TextTag
func (recv *Tag) TextTag() *gtk.TextTag {
	return gtk.TextTagNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Tag) Object() *gobject.Object {
	return recv.TextTag().Object()
}

// CastToWidget down casts any arbitrary Object to Tag.
// Exercise care, as this is a potentially dangerous function if the Object is not a Tag.
func CastToTag(object *gobject.Object) *Tag {
	return TagNewFromC(object.ToC())
}

// TagNew is a wrapper around the C function gtk_source_tag_new.
func TagNew(name string) *Tag {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_source_tag_new(c_name)
	retGo := TagNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// View is a wrapper around the C record GtkSourceView.
type View struct {
	native *C.GtkSourceView
	// parent : record
	// priv : record
}

func ViewNewFromC(u unsafe.Pointer) *View {
	c := (*C.GtkSourceView)(u)
	if c == nil {
		return nil
	}

	g := &View{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *View) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *View) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this View with another View, and returns true if they represent the same GObject.
func (recv *View) Equals(other *View) bool {
	return other.ToC() == recv.ToC()
}

// TextView upcasts to *TextView
func (recv *View) TextView() *gtk.TextView {
	return gtk.TextViewNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *View) Container() *gtk.Container {
	return recv.TextView().Container()
}

// Widget upcasts to *Widget
func (recv *View) Widget() *gtk.Widget {
	return recv.TextView().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *View) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.TextView().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *View) Object() *gobject.Object {
	return recv.TextView().Object()
}

// CastToWidget down casts any arbitrary Object to View.
// Exercise care, as this is a potentially dangerous function if the Object is not a View.
func CastToView(object *gobject.Object) *View {
	return ViewNewFromC(object.ToC())
}

type signalViewChangeCaseDetail struct {
	callback  ViewSignalChangeCaseCallback
	handlerID C.gulong
}

var signalViewChangeCaseId int
var signalViewChangeCaseMap = make(map[int]signalViewChangeCaseDetail)
var signalViewChangeCaseLock sync.RWMutex

// ViewSignalChangeCaseCallback is a callback function for a 'change-case' signal emitted from a View.
type ViewSignalChangeCaseCallback func(targetObject *View, caseType ChangeCaseType)

/*
ConnectChangeCase connects the callback to the 'change-case' signal for the View.

The returned value represents the connection, and may be passed to DisconnectChangeCase to remove it.
*/
func (recv *View) ConnectChangeCase(callback ViewSignalChangeCaseCallback) int {
	signalViewChangeCaseLock.Lock()
	defer signalViewChangeCaseLock.Unlock()

	signalViewChangeCaseId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_change_case(instance, C.gpointer(uintptr(signalViewChangeCaseId)))

	detail := signalViewChangeCaseDetail{callback, handlerID}
	signalViewChangeCaseMap[signalViewChangeCaseId] = detail

	return signalViewChangeCaseId
}

/*
DisconnectChangeCase disconnects a callback from the 'change-case' signal for the View.

The connectionID should be a value returned from a call to ConnectChangeCase.
*/
func (recv *View) DisconnectChangeCase(connectionID int) {
	signalViewChangeCaseLock.Lock()
	defer signalViewChangeCaseLock.Unlock()

	detail, exists := signalViewChangeCaseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewChangeCaseMap, connectionID)
}

//export view_changeCaseHandler
func view_changeCaseHandler(c_targetObject *C.GObject, c_case_type C.GtkSourceChangeCaseType, data C.gpointer) {
	signalViewChangeCaseLock.RLock()
	defer signalViewChangeCaseLock.RUnlock()

	caseType := ChangeCaseType(c_case_type)

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewChangeCaseMap[index].callback
	callback(targetObject, caseType)
}

type signalViewChangeNumberDetail struct {
	callback  ViewSignalChangeNumberCallback
	handlerID C.gulong
}

var signalViewChangeNumberId int
var signalViewChangeNumberMap = make(map[int]signalViewChangeNumberDetail)
var signalViewChangeNumberLock sync.RWMutex

// ViewSignalChangeNumberCallback is a callback function for a 'change-number' signal emitted from a View.
type ViewSignalChangeNumberCallback func(targetObject *View, count int32)

/*
ConnectChangeNumber connects the callback to the 'change-number' signal for the View.

The returned value represents the connection, and may be passed to DisconnectChangeNumber to remove it.
*/
func (recv *View) ConnectChangeNumber(callback ViewSignalChangeNumberCallback) int {
	signalViewChangeNumberLock.Lock()
	defer signalViewChangeNumberLock.Unlock()

	signalViewChangeNumberId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_change_number(instance, C.gpointer(uintptr(signalViewChangeNumberId)))

	detail := signalViewChangeNumberDetail{callback, handlerID}
	signalViewChangeNumberMap[signalViewChangeNumberId] = detail

	return signalViewChangeNumberId
}

/*
DisconnectChangeNumber disconnects a callback from the 'change-number' signal for the View.

The connectionID should be a value returned from a call to ConnectChangeNumber.
*/
func (recv *View) DisconnectChangeNumber(connectionID int) {
	signalViewChangeNumberLock.Lock()
	defer signalViewChangeNumberLock.Unlock()

	detail, exists := signalViewChangeNumberMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewChangeNumberMap, connectionID)
}

//export view_changeNumberHandler
func view_changeNumberHandler(c_targetObject *C.GObject, c_count C.gint, data C.gpointer) {
	signalViewChangeNumberLock.RLock()
	defer signalViewChangeNumberLock.RUnlock()

	count := int32(c_count)

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewChangeNumberMap[index].callback
	callback(targetObject, count)
}

type signalViewJoinLinesDetail struct {
	callback  ViewSignalJoinLinesCallback
	handlerID C.gulong
}

var signalViewJoinLinesId int
var signalViewJoinLinesMap = make(map[int]signalViewJoinLinesDetail)
var signalViewJoinLinesLock sync.RWMutex

// ViewSignalJoinLinesCallback is a callback function for a 'join-lines' signal emitted from a View.
type ViewSignalJoinLinesCallback func(targetObject *View)

/*
ConnectJoinLines connects the callback to the 'join-lines' signal for the View.

The returned value represents the connection, and may be passed to DisconnectJoinLines to remove it.
*/
func (recv *View) ConnectJoinLines(callback ViewSignalJoinLinesCallback) int {
	signalViewJoinLinesLock.Lock()
	defer signalViewJoinLinesLock.Unlock()

	signalViewJoinLinesId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_join_lines(instance, C.gpointer(uintptr(signalViewJoinLinesId)))

	detail := signalViewJoinLinesDetail{callback, handlerID}
	signalViewJoinLinesMap[signalViewJoinLinesId] = detail

	return signalViewJoinLinesId
}

/*
DisconnectJoinLines disconnects a callback from the 'join-lines' signal for the View.

The connectionID should be a value returned from a call to ConnectJoinLines.
*/
func (recv *View) DisconnectJoinLines(connectionID int) {
	signalViewJoinLinesLock.Lock()
	defer signalViewJoinLinesLock.Unlock()

	detail, exists := signalViewJoinLinesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewJoinLinesMap, connectionID)
}

//export view_joinLinesHandler
func view_joinLinesHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalViewJoinLinesLock.RLock()
	defer signalViewJoinLinesLock.RUnlock()

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewJoinLinesMap[index].callback
	callback(targetObject)
}

type signalViewLineMarkActivatedDetail struct {
	callback  ViewSignalLineMarkActivatedCallback
	handlerID C.gulong
}

var signalViewLineMarkActivatedId int
var signalViewLineMarkActivatedMap = make(map[int]signalViewLineMarkActivatedDetail)
var signalViewLineMarkActivatedLock sync.RWMutex

// ViewSignalLineMarkActivatedCallback is a callback function for a 'line-mark-activated' signal emitted from a View.
type ViewSignalLineMarkActivatedCallback func(targetObject *View, iter *gtk.TextIter, event *gdk.Event)

/*
ConnectLineMarkActivated connects the callback to the 'line-mark-activated' signal for the View.

The returned value represents the connection, and may be passed to DisconnectLineMarkActivated to remove it.
*/
func (recv *View) ConnectLineMarkActivated(callback ViewSignalLineMarkActivatedCallback) int {
	signalViewLineMarkActivatedLock.Lock()
	defer signalViewLineMarkActivatedLock.Unlock()

	signalViewLineMarkActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_line_mark_activated(instance, C.gpointer(uintptr(signalViewLineMarkActivatedId)))

	detail := signalViewLineMarkActivatedDetail{callback, handlerID}
	signalViewLineMarkActivatedMap[signalViewLineMarkActivatedId] = detail

	return signalViewLineMarkActivatedId
}

/*
DisconnectLineMarkActivated disconnects a callback from the 'line-mark-activated' signal for the View.

The connectionID should be a value returned from a call to ConnectLineMarkActivated.
*/
func (recv *View) DisconnectLineMarkActivated(connectionID int) {
	signalViewLineMarkActivatedLock.Lock()
	defer signalViewLineMarkActivatedLock.Unlock()

	detail, exists := signalViewLineMarkActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewLineMarkActivatedMap, connectionID)
}

//export view_lineMarkActivatedHandler
func view_lineMarkActivatedHandler(c_targetObject *C.GObject, c_iter *C.GtkTextIter, c_event *C.GdkEvent_, data C.gpointer) {
	signalViewLineMarkActivatedLock.RLock()
	defer signalViewLineMarkActivatedLock.RUnlock()

	iter := gtk.TextIterNewFromC(unsafe.Pointer(c_iter))

	event := gdk.EventNewFromC(unsafe.Pointer(c_event))

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewLineMarkActivatedMap[index].callback
	callback(targetObject, iter, event)
}

type signalViewMoveLinesDetail struct {
	callback  ViewSignalMoveLinesCallback
	handlerID C.gulong
}

var signalViewMoveLinesId int
var signalViewMoveLinesMap = make(map[int]signalViewMoveLinesDetail)
var signalViewMoveLinesLock sync.RWMutex

// ViewSignalMoveLinesCallback is a callback function for a 'move-lines' signal emitted from a View.
type ViewSignalMoveLinesCallback func(targetObject *View, copy bool, count int32)

/*
ConnectMoveLines connects the callback to the 'move-lines' signal for the View.

The returned value represents the connection, and may be passed to DisconnectMoveLines to remove it.
*/
func (recv *View) ConnectMoveLines(callback ViewSignalMoveLinesCallback) int {
	signalViewMoveLinesLock.Lock()
	defer signalViewMoveLinesLock.Unlock()

	signalViewMoveLinesId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_move_lines(instance, C.gpointer(uintptr(signalViewMoveLinesId)))

	detail := signalViewMoveLinesDetail{callback, handlerID}
	signalViewMoveLinesMap[signalViewMoveLinesId] = detail

	return signalViewMoveLinesId
}

/*
DisconnectMoveLines disconnects a callback from the 'move-lines' signal for the View.

The connectionID should be a value returned from a call to ConnectMoveLines.
*/
func (recv *View) DisconnectMoveLines(connectionID int) {
	signalViewMoveLinesLock.Lock()
	defer signalViewMoveLinesLock.Unlock()

	detail, exists := signalViewMoveLinesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewMoveLinesMap, connectionID)
}

//export view_moveLinesHandler
func view_moveLinesHandler(c_targetObject *C.GObject, c_copy C.gboolean, c_count C.gint, data C.gpointer) {
	signalViewMoveLinesLock.RLock()
	defer signalViewMoveLinesLock.RUnlock()

	copy := c_copy == C.TRUE

	count := int32(c_count)

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewMoveLinesMap[index].callback
	callback(targetObject, copy, count)
}

type signalViewMoveToMatchingBracketDetail struct {
	callback  ViewSignalMoveToMatchingBracketCallback
	handlerID C.gulong
}

var signalViewMoveToMatchingBracketId int
var signalViewMoveToMatchingBracketMap = make(map[int]signalViewMoveToMatchingBracketDetail)
var signalViewMoveToMatchingBracketLock sync.RWMutex

// ViewSignalMoveToMatchingBracketCallback is a callback function for a 'move-to-matching-bracket' signal emitted from a View.
type ViewSignalMoveToMatchingBracketCallback func(targetObject *View, extendSelection bool)

/*
ConnectMoveToMatchingBracket connects the callback to the 'move-to-matching-bracket' signal for the View.

The returned value represents the connection, and may be passed to DisconnectMoveToMatchingBracket to remove it.
*/
func (recv *View) ConnectMoveToMatchingBracket(callback ViewSignalMoveToMatchingBracketCallback) int {
	signalViewMoveToMatchingBracketLock.Lock()
	defer signalViewMoveToMatchingBracketLock.Unlock()

	signalViewMoveToMatchingBracketId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_move_to_matching_bracket(instance, C.gpointer(uintptr(signalViewMoveToMatchingBracketId)))

	detail := signalViewMoveToMatchingBracketDetail{callback, handlerID}
	signalViewMoveToMatchingBracketMap[signalViewMoveToMatchingBracketId] = detail

	return signalViewMoveToMatchingBracketId
}

/*
DisconnectMoveToMatchingBracket disconnects a callback from the 'move-to-matching-bracket' signal for the View.

The connectionID should be a value returned from a call to ConnectMoveToMatchingBracket.
*/
func (recv *View) DisconnectMoveToMatchingBracket(connectionID int) {
	signalViewMoveToMatchingBracketLock.Lock()
	defer signalViewMoveToMatchingBracketLock.Unlock()

	detail, exists := signalViewMoveToMatchingBracketMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewMoveToMatchingBracketMap, connectionID)
}

//export view_moveToMatchingBracketHandler
func view_moveToMatchingBracketHandler(c_targetObject *C.GObject, c_extend_selection C.gboolean, data C.gpointer) {
	signalViewMoveToMatchingBracketLock.RLock()
	defer signalViewMoveToMatchingBracketLock.RUnlock()

	extendSelection := c_extend_selection == C.TRUE

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewMoveToMatchingBracketMap[index].callback
	callback(targetObject, extendSelection)
}

type signalViewMoveWordsDetail struct {
	callback  ViewSignalMoveWordsCallback
	handlerID C.gulong
}

var signalViewMoveWordsId int
var signalViewMoveWordsMap = make(map[int]signalViewMoveWordsDetail)
var signalViewMoveWordsLock sync.RWMutex

// ViewSignalMoveWordsCallback is a callback function for a 'move-words' signal emitted from a View.
type ViewSignalMoveWordsCallback func(targetObject *View, count int32)

/*
ConnectMoveWords connects the callback to the 'move-words' signal for the View.

The returned value represents the connection, and may be passed to DisconnectMoveWords to remove it.
*/
func (recv *View) ConnectMoveWords(callback ViewSignalMoveWordsCallback) int {
	signalViewMoveWordsLock.Lock()
	defer signalViewMoveWordsLock.Unlock()

	signalViewMoveWordsId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_move_words(instance, C.gpointer(uintptr(signalViewMoveWordsId)))

	detail := signalViewMoveWordsDetail{callback, handlerID}
	signalViewMoveWordsMap[signalViewMoveWordsId] = detail

	return signalViewMoveWordsId
}

/*
DisconnectMoveWords disconnects a callback from the 'move-words' signal for the View.

The connectionID should be a value returned from a call to ConnectMoveWords.
*/
func (recv *View) DisconnectMoveWords(connectionID int) {
	signalViewMoveWordsLock.Lock()
	defer signalViewMoveWordsLock.Unlock()

	detail, exists := signalViewMoveWordsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewMoveWordsMap, connectionID)
}

//export view_moveWordsHandler
func view_moveWordsHandler(c_targetObject *C.GObject, c_count C.gint, data C.gpointer) {
	signalViewMoveWordsLock.RLock()
	defer signalViewMoveWordsLock.RUnlock()

	count := int32(c_count)

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewMoveWordsMap[index].callback
	callback(targetObject, count)
}

type signalViewRedoDetail struct {
	callback  ViewSignalRedoCallback
	handlerID C.gulong
}

var signalViewRedoId int
var signalViewRedoMap = make(map[int]signalViewRedoDetail)
var signalViewRedoLock sync.RWMutex

// ViewSignalRedoCallback is a callback function for a 'redo' signal emitted from a View.
type ViewSignalRedoCallback func(targetObject *View)

/*
ConnectRedo connects the callback to the 'redo' signal for the View.

The returned value represents the connection, and may be passed to DisconnectRedo to remove it.
*/
func (recv *View) ConnectRedo(callback ViewSignalRedoCallback) int {
	signalViewRedoLock.Lock()
	defer signalViewRedoLock.Unlock()

	signalViewRedoId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_redo(instance, C.gpointer(uintptr(signalViewRedoId)))

	detail := signalViewRedoDetail{callback, handlerID}
	signalViewRedoMap[signalViewRedoId] = detail

	return signalViewRedoId
}

/*
DisconnectRedo disconnects a callback from the 'redo' signal for the View.

The connectionID should be a value returned from a call to ConnectRedo.
*/
func (recv *View) DisconnectRedo(connectionID int) {
	signalViewRedoLock.Lock()
	defer signalViewRedoLock.Unlock()

	detail, exists := signalViewRedoMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewRedoMap, connectionID)
}

//export view_redoHandler
func view_redoHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalViewRedoLock.RLock()
	defer signalViewRedoLock.RUnlock()

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewRedoMap[index].callback
	callback(targetObject)
}

type signalViewShowCompletionDetail struct {
	callback  ViewSignalShowCompletionCallback
	handlerID C.gulong
}

var signalViewShowCompletionId int
var signalViewShowCompletionMap = make(map[int]signalViewShowCompletionDetail)
var signalViewShowCompletionLock sync.RWMutex

// ViewSignalShowCompletionCallback is a callback function for a 'show-completion' signal emitted from a View.
type ViewSignalShowCompletionCallback func(targetObject *View)

/*
ConnectShowCompletion connects the callback to the 'show-completion' signal for the View.

The returned value represents the connection, and may be passed to DisconnectShowCompletion to remove it.
*/
func (recv *View) ConnectShowCompletion(callback ViewSignalShowCompletionCallback) int {
	signalViewShowCompletionLock.Lock()
	defer signalViewShowCompletionLock.Unlock()

	signalViewShowCompletionId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_show_completion(instance, C.gpointer(uintptr(signalViewShowCompletionId)))

	detail := signalViewShowCompletionDetail{callback, handlerID}
	signalViewShowCompletionMap[signalViewShowCompletionId] = detail

	return signalViewShowCompletionId
}

/*
DisconnectShowCompletion disconnects a callback from the 'show-completion' signal for the View.

The connectionID should be a value returned from a call to ConnectShowCompletion.
*/
func (recv *View) DisconnectShowCompletion(connectionID int) {
	signalViewShowCompletionLock.Lock()
	defer signalViewShowCompletionLock.Unlock()

	detail, exists := signalViewShowCompletionMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewShowCompletionMap, connectionID)
}

//export view_showCompletionHandler
func view_showCompletionHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalViewShowCompletionLock.RLock()
	defer signalViewShowCompletionLock.RUnlock()

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewShowCompletionMap[index].callback
	callback(targetObject)
}

type signalViewSmartHomeEndDetail struct {
	callback  ViewSignalSmartHomeEndCallback
	handlerID C.gulong
}

var signalViewSmartHomeEndId int
var signalViewSmartHomeEndMap = make(map[int]signalViewSmartHomeEndDetail)
var signalViewSmartHomeEndLock sync.RWMutex

// ViewSignalSmartHomeEndCallback is a callback function for a 'smart-home-end' signal emitted from a View.
type ViewSignalSmartHomeEndCallback func(targetObject *View, iter *gtk.TextIter, count int32)

/*
ConnectSmartHomeEnd connects the callback to the 'smart-home-end' signal for the View.

The returned value represents the connection, and may be passed to DisconnectSmartHomeEnd to remove it.
*/
func (recv *View) ConnectSmartHomeEnd(callback ViewSignalSmartHomeEndCallback) int {
	signalViewSmartHomeEndLock.Lock()
	defer signalViewSmartHomeEndLock.Unlock()

	signalViewSmartHomeEndId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_smart_home_end(instance, C.gpointer(uintptr(signalViewSmartHomeEndId)))

	detail := signalViewSmartHomeEndDetail{callback, handlerID}
	signalViewSmartHomeEndMap[signalViewSmartHomeEndId] = detail

	return signalViewSmartHomeEndId
}

/*
DisconnectSmartHomeEnd disconnects a callback from the 'smart-home-end' signal for the View.

The connectionID should be a value returned from a call to ConnectSmartHomeEnd.
*/
func (recv *View) DisconnectSmartHomeEnd(connectionID int) {
	signalViewSmartHomeEndLock.Lock()
	defer signalViewSmartHomeEndLock.Unlock()

	detail, exists := signalViewSmartHomeEndMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewSmartHomeEndMap, connectionID)
}

//export view_smartHomeEndHandler
func view_smartHomeEndHandler(c_targetObject *C.GObject, c_iter *C.GtkTextIter, c_count C.gint, data C.gpointer) {
	signalViewSmartHomeEndLock.RLock()
	defer signalViewSmartHomeEndLock.RUnlock()

	iter := gtk.TextIterNewFromC(unsafe.Pointer(c_iter))

	count := int32(c_count)

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewSmartHomeEndMap[index].callback
	callback(targetObject, iter, count)
}

type signalViewUndoDetail struct {
	callback  ViewSignalUndoCallback
	handlerID C.gulong
}

var signalViewUndoId int
var signalViewUndoMap = make(map[int]signalViewUndoDetail)
var signalViewUndoLock sync.RWMutex

// ViewSignalUndoCallback is a callback function for a 'undo' signal emitted from a View.
type ViewSignalUndoCallback func(targetObject *View)

/*
ConnectUndo connects the callback to the 'undo' signal for the View.

The returned value represents the connection, and may be passed to DisconnectUndo to remove it.
*/
func (recv *View) ConnectUndo(callback ViewSignalUndoCallback) int {
	signalViewUndoLock.Lock()
	defer signalViewUndoLock.Unlock()

	signalViewUndoId++
	instance := C.gpointer(recv.native)
	handlerID := C.View_signal_connect_undo(instance, C.gpointer(uintptr(signalViewUndoId)))

	detail := signalViewUndoDetail{callback, handlerID}
	signalViewUndoMap[signalViewUndoId] = detail

	return signalViewUndoId
}

/*
DisconnectUndo disconnects a callback from the 'undo' signal for the View.

The connectionID should be a value returned from a call to ConnectUndo.
*/
func (recv *View) DisconnectUndo(connectionID int) {
	signalViewUndoLock.Lock()
	defer signalViewUndoLock.Unlock()

	detail, exists := signalViewUndoMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalViewUndoMap, connectionID)
}

//export view_undoHandler
func view_undoHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalViewUndoLock.RLock()
	defer signalViewUndoLock.RUnlock()

	targetObject := ViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalViewUndoMap[index].callback
	callback(targetObject)
}

// ViewNew is a wrapper around the C function gtk_source_view_new.
func ViewNew() *View {
	retC := C.gtk_source_view_new()
	retGo := ViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ViewNewWithBuffer is a wrapper around the C function gtk_source_view_new_with_buffer.
func ViewNewWithBuffer(buffer *Buffer) *View {
	c_buffer := (*C.GtkSourceBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkSourceBuffer)(buffer.ToC())
	}

	retC := C.gtk_source_view_new_with_buffer(c_buffer)
	retGo := ViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAutoIndent is a wrapper around the C function gtk_source_view_get_auto_indent.
func (recv *View) GetAutoIndent() bool {
	retC := C.gtk_source_view_get_auto_indent((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetBackgroundPattern is a wrapper around the C function gtk_source_view_get_background_pattern.
func (recv *View) GetBackgroundPattern() BackgroundPatternType {
	retC := C.gtk_source_view_get_background_pattern((*C.GtkSourceView)(recv.native))
	retGo := (BackgroundPatternType)(retC)

	return retGo
}

// GetCompletion is a wrapper around the C function gtk_source_view_get_completion.
func (recv *View) GetCompletion() *Completion {
	retC := C.gtk_source_view_get_completion((*C.GtkSourceView)(recv.native))
	retGo := CompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDrawSpaces is a wrapper around the C function gtk_source_view_get_draw_spaces.
func (recv *View) GetDrawSpaces() DrawSpacesFlags {
	retC := C.gtk_source_view_get_draw_spaces((*C.GtkSourceView)(recv.native))
	retGo := (DrawSpacesFlags)(retC)

	return retGo
}

// GetGutter is a wrapper around the C function gtk_source_view_get_gutter.
func (recv *View) GetGutter(windowType gtk.TextWindowType) *Gutter {
	c_window_type := (C.GtkTextWindowType)(windowType)

	retC := C.gtk_source_view_get_gutter((*C.GtkSourceView)(recv.native), c_window_type)
	retGo := GutterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHighlightCurrentLine is a wrapper around the C function gtk_source_view_get_highlight_current_line.
func (recv *View) GetHighlightCurrentLine() bool {
	retC := C.gtk_source_view_get_highlight_current_line((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIndentOnTab is a wrapper around the C function gtk_source_view_get_indent_on_tab.
func (recv *View) GetIndentOnTab() bool {
	retC := C.gtk_source_view_get_indent_on_tab((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIndentWidth is a wrapper around the C function gtk_source_view_get_indent_width.
func (recv *View) GetIndentWidth() int32 {
	retC := C.gtk_source_view_get_indent_width((*C.GtkSourceView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetInsertSpacesInsteadOfTabs is a wrapper around the C function gtk_source_view_get_insert_spaces_instead_of_tabs.
func (recv *View) GetInsertSpacesInsteadOfTabs() bool {
	retC := C.gtk_source_view_get_insert_spaces_instead_of_tabs((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMarkAttributes is a wrapper around the C function gtk_source_view_get_mark_attributes.
func (recv *View) GetMarkAttributes(category string, priority int32) *MarkAttributes {
	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	c_priority := (C.gint)(priority)

	retC := C.gtk_source_view_get_mark_attributes((*C.GtkSourceView)(recv.native), c_category, &c_priority)
	retGo := MarkAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRightMarginPosition is a wrapper around the C function gtk_source_view_get_right_margin_position.
func (recv *View) GetRightMarginPosition() uint32 {
	retC := C.gtk_source_view_get_right_margin_position((*C.GtkSourceView)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetShowLineMarks is a wrapper around the C function gtk_source_view_get_show_line_marks.
func (recv *View) GetShowLineMarks() bool {
	retC := C.gtk_source_view_get_show_line_marks((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowLineNumbers is a wrapper around the C function gtk_source_view_get_show_line_numbers.
func (recv *View) GetShowLineNumbers() bool {
	retC := C.gtk_source_view_get_show_line_numbers((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowRightMargin is a wrapper around the C function gtk_source_view_get_show_right_margin.
func (recv *View) GetShowRightMargin() bool {
	retC := C.gtk_source_view_get_show_right_margin((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSmartBackspace is a wrapper around the C function gtk_source_view_get_smart_backspace.
func (recv *View) GetSmartBackspace() bool {
	retC := C.gtk_source_view_get_smart_backspace((*C.GtkSourceView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSmartHomeEnd is a wrapper around the C function gtk_source_view_get_smart_home_end.
func (recv *View) GetSmartHomeEnd() SmartHomeEndType {
	retC := C.gtk_source_view_get_smart_home_end((*C.GtkSourceView)(recv.native))
	retGo := (SmartHomeEndType)(retC)

	return retGo
}

// GetTabWidth is a wrapper around the C function gtk_source_view_get_tab_width.
func (recv *View) GetTabWidth() uint32 {
	retC := C.gtk_source_view_get_tab_width((*C.GtkSourceView)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetVisualColumn is a wrapper around the C function gtk_source_view_get_visual_column.
func (recv *View) GetVisualColumn(iter *gtk.TextIter) uint32 {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	retC := C.gtk_source_view_get_visual_column((*C.GtkSourceView)(recv.native), c_iter)
	retGo := (uint32)(retC)

	return retGo
}

// IndentLines is a wrapper around the C function gtk_source_view_indent_lines.
func (recv *View) IndentLines(start *gtk.TextIter, end *gtk.TextIter) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	C.gtk_source_view_indent_lines((*C.GtkSourceView)(recv.native), c_start, c_end)

	return
}

// SetAutoIndent is a wrapper around the C function gtk_source_view_set_auto_indent.
func (recv *View) SetAutoIndent(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.gtk_source_view_set_auto_indent((*C.GtkSourceView)(recv.native), c_enable)

	return
}

// SetBackgroundPattern is a wrapper around the C function gtk_source_view_set_background_pattern.
func (recv *View) SetBackgroundPattern(backgroundPattern BackgroundPatternType) {
	c_background_pattern := (C.GtkSourceBackgroundPatternType)(backgroundPattern)

	C.gtk_source_view_set_background_pattern((*C.GtkSourceView)(recv.native), c_background_pattern)

	return
}

// SetDrawSpaces is a wrapper around the C function gtk_source_view_set_draw_spaces.
func (recv *View) SetDrawSpaces(flags DrawSpacesFlags) {
	c_flags := (C.GtkSourceDrawSpacesFlags)(flags)

	C.gtk_source_view_set_draw_spaces((*C.GtkSourceView)(recv.native), c_flags)

	return
}

// SetHighlightCurrentLine is a wrapper around the C function gtk_source_view_set_highlight_current_line.
func (recv *View) SetHighlightCurrentLine(highlight bool) {
	c_highlight :=
		boolToGboolean(highlight)

	C.gtk_source_view_set_highlight_current_line((*C.GtkSourceView)(recv.native), c_highlight)

	return
}

// SetIndentOnTab is a wrapper around the C function gtk_source_view_set_indent_on_tab.
func (recv *View) SetIndentOnTab(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.gtk_source_view_set_indent_on_tab((*C.GtkSourceView)(recv.native), c_enable)

	return
}

// SetIndentWidth is a wrapper around the C function gtk_source_view_set_indent_width.
func (recv *View) SetIndentWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_source_view_set_indent_width((*C.GtkSourceView)(recv.native), c_width)

	return
}

// SetInsertSpacesInsteadOfTabs is a wrapper around the C function gtk_source_view_set_insert_spaces_instead_of_tabs.
func (recv *View) SetInsertSpacesInsteadOfTabs(enable bool) {
	c_enable :=
		boolToGboolean(enable)

	C.gtk_source_view_set_insert_spaces_instead_of_tabs((*C.GtkSourceView)(recv.native), c_enable)

	return
}

// SetMarkAttributes is a wrapper around the C function gtk_source_view_set_mark_attributes.
func (recv *View) SetMarkAttributes(category string, attributes *MarkAttributes, priority int32) {
	c_category := C.CString(category)
	defer C.free(unsafe.Pointer(c_category))

	c_attributes := (*C.GtkSourceMarkAttributes)(C.NULL)
	if attributes != nil {
		c_attributes = (*C.GtkSourceMarkAttributes)(attributes.ToC())
	}

	c_priority := (C.gint)(priority)

	C.gtk_source_view_set_mark_attributes((*C.GtkSourceView)(recv.native), c_category, c_attributes, c_priority)

	return
}

// SetRightMarginPosition is a wrapper around the C function gtk_source_view_set_right_margin_position.
func (recv *View) SetRightMarginPosition(pos uint32) {
	c_pos := (C.guint)(pos)

	C.gtk_source_view_set_right_margin_position((*C.GtkSourceView)(recv.native), c_pos)

	return
}

// SetShowLineMarks is a wrapper around the C function gtk_source_view_set_show_line_marks.
func (recv *View) SetShowLineMarks(show bool) {
	c_show :=
		boolToGboolean(show)

	C.gtk_source_view_set_show_line_marks((*C.GtkSourceView)(recv.native), c_show)

	return
}

// SetShowLineNumbers is a wrapper around the C function gtk_source_view_set_show_line_numbers.
func (recv *View) SetShowLineNumbers(show bool) {
	c_show :=
		boolToGboolean(show)

	C.gtk_source_view_set_show_line_numbers((*C.GtkSourceView)(recv.native), c_show)

	return
}

// SetShowRightMargin is a wrapper around the C function gtk_source_view_set_show_right_margin.
func (recv *View) SetShowRightMargin(show bool) {
	c_show :=
		boolToGboolean(show)

	C.gtk_source_view_set_show_right_margin((*C.GtkSourceView)(recv.native), c_show)

	return
}

// SetSmartBackspace is a wrapper around the C function gtk_source_view_set_smart_backspace.
func (recv *View) SetSmartBackspace(smartBackspace bool) {
	c_smart_backspace :=
		boolToGboolean(smartBackspace)

	C.gtk_source_view_set_smart_backspace((*C.GtkSourceView)(recv.native), c_smart_backspace)

	return
}

// SetSmartHomeEnd is a wrapper around the C function gtk_source_view_set_smart_home_end.
func (recv *View) SetSmartHomeEnd(smartHomeEnd SmartHomeEndType) {
	c_smart_home_end := (C.GtkSourceSmartHomeEndType)(smartHomeEnd)

	C.gtk_source_view_set_smart_home_end((*C.GtkSourceView)(recv.native), c_smart_home_end)

	return
}

// SetTabWidth is a wrapper around the C function gtk_source_view_set_tab_width.
func (recv *View) SetTabWidth(width uint32) {
	c_width := (C.guint)(width)

	C.gtk_source_view_set_tab_width((*C.GtkSourceView)(recv.native), c_width)

	return
}

// UnindentLines is a wrapper around the C function gtk_source_view_unindent_lines.
func (recv *View) UnindentLines(start *gtk.TextIter, end *gtk.TextIter) {
	c_start := (*C.GtkTextIter)(C.NULL)
	if start != nil {
		c_start = (*C.GtkTextIter)(start.ToC())
	}

	c_end := (*C.GtkTextIter)(C.NULL)
	if end != nil {
		c_end = (*C.GtkTextIter)(end.ToC())
	}

	C.gtk_source_view_unindent_lines((*C.GtkSourceView)(recv.native), c_start, c_end)

	return
}

// ImplementorIface returns the ImplementorIface interface implemented by View
func (recv *View) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by View
func (recv *View) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by View
func (recv *View) Scrollable() *gtk.Scrollable {
	return gtk.ScrollableNewFromC(recv.ToC())
}

type BackgroundPatternType C.GtkSourceBackgroundPatternType

const (
	GTK_SOURCE_BACKGROUND_PATTERN_TYPE_NONE BackgroundPatternType = 0
	GTK_SOURCE_BACKGROUND_PATTERN_TYPE_GRID BackgroundPatternType = 1
)

type BracketMatchType C.GtkSourceBracketMatchType

const (
	GTK_SOURCE_BRACKET_MATCH_NONE         BracketMatchType = 0
	GTK_SOURCE_BRACKET_MATCH_OUT_OF_RANGE BracketMatchType = 1
	GTK_SOURCE_BRACKET_MATCH_NOT_FOUND    BracketMatchType = 2
	GTK_SOURCE_BRACKET_MATCH_FOUND        BracketMatchType = 3
)

type ChangeCaseType C.GtkSourceChangeCaseType

const (
	GTK_SOURCE_CHANGE_CASE_LOWER  ChangeCaseType = 0
	GTK_SOURCE_CHANGE_CASE_UPPER  ChangeCaseType = 1
	GTK_SOURCE_CHANGE_CASE_TOGGLE ChangeCaseType = 2
	GTK_SOURCE_CHANGE_CASE_TITLE  ChangeCaseType = 3
)

type CompletionError C.GtkSourceCompletionError

const (
	GTK_SOURCE_COMPLETION_ERROR_ALREADY_BOUND CompletionError = 0
	GTK_SOURCE_COMPLETION_ERROR_NOT_BOUND     CompletionError = 1
)

// CompletionErrorQuark is a wrapper around the C function gtk_source_completion_error_quark.
func CompletionErrorQuark() glib.Quark {
	retC := C.gtk_source_completion_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type CompressionType C.GtkSourceCompressionType

const (
	GTK_SOURCE_COMPRESSION_TYPE_NONE CompressionType = 0
	GTK_SOURCE_COMPRESSION_TYPE_GZIP CompressionType = 1
)

type FileLoaderError C.GtkSourceFileLoaderError

const (
	GTK_SOURCE_FILE_LOADER_ERROR_TOO_BIG                        FileLoaderError = 0
	GTK_SOURCE_FILE_LOADER_ERROR_ENCODING_AUTO_DETECTION_FAILED FileLoaderError = 1
	GTK_SOURCE_FILE_LOADER_ERROR_CONVERSION_FALLBACK            FileLoaderError = 2
)

// FileLoaderErrorQuark is a wrapper around the C function gtk_source_file_loader_error_quark.
func FileLoaderErrorQuark() glib.Quark {
	retC := C.gtk_source_file_loader_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type FileSaverError C.GtkSourceFileSaverError

const (
	GTK_SOURCE_FILE_SAVER_ERROR_INVALID_CHARS       FileSaverError = 0
	GTK_SOURCE_FILE_SAVER_ERROR_EXTERNALLY_MODIFIED FileSaverError = 1
)

// FileSaverErrorQuark is a wrapper around the C function gtk_source_file_saver_error_quark.
func FileSaverErrorQuark() glib.Quark {
	retC := C.gtk_source_file_saver_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type GutterRendererAlignmentMode C.GtkSourceGutterRendererAlignmentMode

const (
	GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_CELL  GutterRendererAlignmentMode = 0
	GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_FIRST GutterRendererAlignmentMode = 1
	GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_LAST  GutterRendererAlignmentMode = 2
)

type NewlineType C.GtkSourceNewlineType

const (
	GTK_SOURCE_NEWLINE_TYPE_LF    NewlineType = 0
	GTK_SOURCE_NEWLINE_TYPE_CR    NewlineType = 1
	GTK_SOURCE_NEWLINE_TYPE_CR_LF NewlineType = 2
)

type SmartHomeEndType C.GtkSourceSmartHomeEndType

const (
	GTK_SOURCE_SMART_HOME_END_DISABLED SmartHomeEndType = 0
	GTK_SOURCE_SMART_HOME_END_BEFORE   SmartHomeEndType = 1
	GTK_SOURCE_SMART_HOME_END_AFTER    SmartHomeEndType = 2
	GTK_SOURCE_SMART_HOME_END_ALWAYS   SmartHomeEndType = 3
)

type ViewGutterPosition C.GtkSourceViewGutterPosition

const (
	GTK_SOURCE_VIEW_GUTTER_POSITION_LINES ViewGutterPosition = -30
	GTK_SOURCE_VIEW_GUTTER_POSITION_MARKS ViewGutterPosition = -20
)

// UtilsEscapeSearchText is a wrapper around the C function gtk_source_utils_escape_search_text.
func UtilsEscapeSearchText(text string) string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_source_utils_escape_search_text(c_text)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UtilsUnescapeSearchText is a wrapper around the C function gtk_source_utils_unescape_search_text.
func UtilsUnescapeSearchText(text string) string {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_source_utils_unescape_search_text(c_text)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// CompletionProposal is a wrapper around the C record GtkSourceCompletionProposal.
type CompletionProposal struct {
	native *C.GtkSourceCompletionProposal
}

func CompletionProposalNewFromC(u unsafe.Pointer) *CompletionProposal {
	c := (*C.GtkSourceCompletionProposal)(u)
	if c == nil {
		return nil
	}

	g := &CompletionProposal{native: c}

	return g
}

func (recv *CompletionProposal) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionProposal with another CompletionProposal, and returns true if they represent the same GObject.
func (recv *CompletionProposal) Equals(other *CompletionProposal) bool {
	return other.ToC() == recv.ToC()
}

type signalCompletionProposalChangedDetail struct {
	callback  CompletionProposalSignalChangedCallback
	handlerID C.gulong
}

var signalCompletionProposalChangedId int
var signalCompletionProposalChangedMap = make(map[int]signalCompletionProposalChangedDetail)
var signalCompletionProposalChangedLock sync.RWMutex

// CompletionProposalSignalChangedCallback is a callback function for a 'changed' signal emitted from a CompletionProposal.
type CompletionProposalSignalChangedCallback func(targetObject *CompletionProposal)

/*
ConnectChanged connects the callback to the 'changed' signal for the CompletionProposal.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *CompletionProposal) ConnectChanged(callback CompletionProposalSignalChangedCallback) int {
	signalCompletionProposalChangedLock.Lock()
	defer signalCompletionProposalChangedLock.Unlock()

	signalCompletionProposalChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CompletionProposal_signal_connect_changed(instance, C.gpointer(uintptr(signalCompletionProposalChangedId)))

	detail := signalCompletionProposalChangedDetail{callback, handlerID}
	signalCompletionProposalChangedMap[signalCompletionProposalChangedId] = detail

	return signalCompletionProposalChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the CompletionProposal.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *CompletionProposal) DisconnectChanged(connectionID int) {
	signalCompletionProposalChangedLock.Lock()
	defer signalCompletionProposalChangedLock.Unlock()

	detail, exists := signalCompletionProposalChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCompletionProposalChangedMap, connectionID)
}

//export completionproposal_changedHandler
func completionproposal_changedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCompletionProposalChangedLock.RLock()
	defer signalCompletionProposalChangedLock.RUnlock()

	targetObject := CompletionProposalNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCompletionProposalChangedMap[index].callback
	callback(targetObject)
}

// Changed is a wrapper around the C function gtk_source_completion_proposal_changed.
func (recv *CompletionProposal) Changed() {
	C.gtk_source_completion_proposal_changed((*C.GtkSourceCompletionProposal)(recv.native))

	return
}

// Equal is a wrapper around the C function gtk_source_completion_proposal_equal.
func (recv *CompletionProposal) Equal(other *CompletionProposal) bool {
	c_other := (*C.GtkSourceCompletionProposal)(other.ToC())

	retC := C.gtk_source_completion_proposal_equal((*C.GtkSourceCompletionProposal)(recv.native), c_other)
	retGo := retC == C.TRUE

	return retGo
}

// GetGicon is a wrapper around the C function gtk_source_completion_proposal_get_gicon.
func (recv *CompletionProposal) GetGicon() *gio.Icon {
	retC := C.gtk_source_completion_proposal_get_gicon((*C.GtkSourceCompletionProposal)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIcon is a wrapper around the C function gtk_source_completion_proposal_get_icon.
func (recv *CompletionProposal) GetIcon() *gdkpixbuf.Pixbuf {
	retC := C.gtk_source_completion_proposal_get_icon((*C.GtkSourceCompletionProposal)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIconName is a wrapper around the C function gtk_source_completion_proposal_get_icon_name.
func (recv *CompletionProposal) GetIconName() string {
	retC := C.gtk_source_completion_proposal_get_icon_name((*C.GtkSourceCompletionProposal)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetInfo is a wrapper around the C function gtk_source_completion_proposal_get_info.
func (recv *CompletionProposal) GetInfo() string {
	retC := C.gtk_source_completion_proposal_get_info((*C.GtkSourceCompletionProposal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetLabel is a wrapper around the C function gtk_source_completion_proposal_get_label.
func (recv *CompletionProposal) GetLabel() string {
	retC := C.gtk_source_completion_proposal_get_label((*C.GtkSourceCompletionProposal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetMarkup is a wrapper around the C function gtk_source_completion_proposal_get_markup.
func (recv *CompletionProposal) GetMarkup() string {
	retC := C.gtk_source_completion_proposal_get_markup((*C.GtkSourceCompletionProposal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetText is a wrapper around the C function gtk_source_completion_proposal_get_text.
func (recv *CompletionProposal) GetText() string {
	retC := C.gtk_source_completion_proposal_get_text((*C.GtkSourceCompletionProposal)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Hash is a wrapper around the C function gtk_source_completion_proposal_hash.
func (recv *CompletionProposal) Hash() uint32 {
	retC := C.gtk_source_completion_proposal_hash((*C.GtkSourceCompletionProposal)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// CompletionProvider is a wrapper around the C record GtkSourceCompletionProvider.
type CompletionProvider struct {
	native *C.GtkSourceCompletionProvider
}

func CompletionProviderNewFromC(u unsafe.Pointer) *CompletionProvider {
	c := (*C.GtkSourceCompletionProvider)(u)
	if c == nil {
		return nil
	}

	g := &CompletionProvider{native: c}

	return g
}

func (recv *CompletionProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionProvider with another CompletionProvider, and returns true if they represent the same GObject.
func (recv *CompletionProvider) Equals(other *CompletionProvider) bool {
	return other.ToC() == recv.ToC()
}

// ActivateProposal is a wrapper around the C function gtk_source_completion_provider_activate_proposal.
func (recv *CompletionProvider) ActivateProposal(proposal *CompletionProposal, iter *gtk.TextIter) bool {
	c_proposal := (*C.GtkSourceCompletionProposal)(proposal.ToC())

	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	retC := C.gtk_source_completion_provider_activate_proposal((*C.GtkSourceCompletionProvider)(recv.native), c_proposal, c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// GetActivation is a wrapper around the C function gtk_source_completion_provider_get_activation.
func (recv *CompletionProvider) GetActivation() CompletionActivation {
	retC := C.gtk_source_completion_provider_get_activation((*C.GtkSourceCompletionProvider)(recv.native))
	retGo := (CompletionActivation)(retC)

	return retGo
}

// GetGicon is a wrapper around the C function gtk_source_completion_provider_get_gicon.
func (recv *CompletionProvider) GetGicon() *gio.Icon {
	retC := C.gtk_source_completion_provider_get_gicon((*C.GtkSourceCompletionProvider)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIcon is a wrapper around the C function gtk_source_completion_provider_get_icon.
func (recv *CompletionProvider) GetIcon() *gdkpixbuf.Pixbuf {
	retC := C.gtk_source_completion_provider_get_icon((*C.GtkSourceCompletionProvider)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIconName is a wrapper around the C function gtk_source_completion_provider_get_icon_name.
func (recv *CompletionProvider) GetIconName() string {
	retC := C.gtk_source_completion_provider_get_icon_name((*C.GtkSourceCompletionProvider)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetInfoWidget is a wrapper around the C function gtk_source_completion_provider_get_info_widget.
func (recv *CompletionProvider) GetInfoWidget(proposal *CompletionProposal) *gtk.Widget {
	c_proposal := (*C.GtkSourceCompletionProposal)(proposal.ToC())

	retC := C.gtk_source_completion_provider_get_info_widget((*C.GtkSourceCompletionProvider)(recv.native), c_proposal)
	var retGo (*gtk.Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gtk.WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetInteractiveDelay is a wrapper around the C function gtk_source_completion_provider_get_interactive_delay.
func (recv *CompletionProvider) GetInteractiveDelay() int32 {
	retC := C.gtk_source_completion_provider_get_interactive_delay((*C.GtkSourceCompletionProvider)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetName is a wrapper around the C function gtk_source_completion_provider_get_name.
func (recv *CompletionProvider) GetName() string {
	retC := C.gtk_source_completion_provider_get_name((*C.GtkSourceCompletionProvider)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPriority is a wrapper around the C function gtk_source_completion_provider_get_priority.
func (recv *CompletionProvider) GetPriority() int32 {
	retC := C.gtk_source_completion_provider_get_priority((*C.GtkSourceCompletionProvider)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStartIter is a wrapper around the C function gtk_source_completion_provider_get_start_iter.
func (recv *CompletionProvider) GetStartIter(context *CompletionContext, proposal *CompletionProposal) (bool, *gtk.TextIter) {
	c_context := (*C.GtkSourceCompletionContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkSourceCompletionContext)(context.ToC())
	}

	c_proposal := (*C.GtkSourceCompletionProposal)(proposal.ToC())

	var c_iter C.GtkTextIter

	retC := C.gtk_source_completion_provider_get_start_iter((*C.GtkSourceCompletionProvider)(recv.native), c_context, c_proposal, &c_iter)
	retGo := retC == C.TRUE

	iter := gtk.TextIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// Match is a wrapper around the C function gtk_source_completion_provider_match.
func (recv *CompletionProvider) Match(context *CompletionContext) bool {
	c_context := (*C.GtkSourceCompletionContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkSourceCompletionContext)(context.ToC())
	}

	retC := C.gtk_source_completion_provider_match((*C.GtkSourceCompletionProvider)(recv.native), c_context)
	retGo := retC == C.TRUE

	return retGo
}

// Populate is a wrapper around the C function gtk_source_completion_provider_populate.
func (recv *CompletionProvider) Populate(context *CompletionContext) {
	c_context := (*C.GtkSourceCompletionContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkSourceCompletionContext)(context.ToC())
	}

	C.gtk_source_completion_provider_populate((*C.GtkSourceCompletionProvider)(recv.native), c_context)

	return
}

// UpdateInfo is a wrapper around the C function gtk_source_completion_provider_update_info.
func (recv *CompletionProvider) UpdateInfo(proposal *CompletionProposal, info *CompletionInfo) {
	c_proposal := (*C.GtkSourceCompletionProposal)(proposal.ToC())

	c_info := (*C.GtkSourceCompletionInfo)(C.NULL)
	if info != nil {
		c_info = (*C.GtkSourceCompletionInfo)(info.ToC())
	}

	C.gtk_source_completion_provider_update_info((*C.GtkSourceCompletionProvider)(recv.native), c_proposal, c_info)

	return
}

// StyleSchemeChooser is a wrapper around the C record GtkSourceStyleSchemeChooser.
type StyleSchemeChooser struct {
	native *C.GtkSourceStyleSchemeChooser
}

func StyleSchemeChooserNewFromC(u unsafe.Pointer) *StyleSchemeChooser {
	c := (*C.GtkSourceStyleSchemeChooser)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeChooser{native: c}

	return g
}

func (recv *StyleSchemeChooser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeChooser with another StyleSchemeChooser, and returns true if they represent the same GObject.
func (recv *StyleSchemeChooser) Equals(other *StyleSchemeChooser) bool {
	return other.ToC() == recv.ToC()
}

// GetStyleScheme is a wrapper around the C function gtk_source_style_scheme_chooser_get_style_scheme.
func (recv *StyleSchemeChooser) GetStyleScheme() *StyleScheme {
	retC := C.gtk_source_style_scheme_chooser_get_style_scheme((*C.GtkSourceStyleSchemeChooser)(recv.native))
	retGo := StyleSchemeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetStyleScheme is a wrapper around the C function gtk_source_style_scheme_chooser_set_style_scheme.
func (recv *StyleSchemeChooser) SetStyleScheme(scheme *StyleScheme) {
	c_scheme := (*C.GtkSourceStyleScheme)(C.NULL)
	if scheme != nil {
		c_scheme = (*C.GtkSourceStyleScheme)(scheme.ToC())
	}

	C.gtk_source_style_scheme_chooser_set_style_scheme((*C.GtkSourceStyleSchemeChooser)(recv.native), c_scheme)

	return
}

// UndoManager is a wrapper around the C record GtkSourceUndoManager.
type UndoManager struct {
	native *C.GtkSourceUndoManager
}

func UndoManagerNewFromC(u unsafe.Pointer) *UndoManager {
	c := (*C.GtkSourceUndoManager)(u)
	if c == nil {
		return nil
	}

	g := &UndoManager{native: c}

	return g
}

func (recv *UndoManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UndoManager with another UndoManager, and returns true if they represent the same GObject.
func (recv *UndoManager) Equals(other *UndoManager) bool {
	return other.ToC() == recv.ToC()
}

type signalUndoManagerCanRedoChangedDetail struct {
	callback  UndoManagerSignalCanRedoChangedCallback
	handlerID C.gulong
}

var signalUndoManagerCanRedoChangedId int
var signalUndoManagerCanRedoChangedMap = make(map[int]signalUndoManagerCanRedoChangedDetail)
var signalUndoManagerCanRedoChangedLock sync.RWMutex

// UndoManagerSignalCanRedoChangedCallback is a callback function for a 'can-redo-changed' signal emitted from a UndoManager.
type UndoManagerSignalCanRedoChangedCallback func(targetObject *UndoManager)

/*
ConnectCanRedoChanged connects the callback to the 'can-redo-changed' signal for the UndoManager.

The returned value represents the connection, and may be passed to DisconnectCanRedoChanged to remove it.
*/
func (recv *UndoManager) ConnectCanRedoChanged(callback UndoManagerSignalCanRedoChangedCallback) int {
	signalUndoManagerCanRedoChangedLock.Lock()
	defer signalUndoManagerCanRedoChangedLock.Unlock()

	signalUndoManagerCanRedoChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.UndoManager_signal_connect_can_redo_changed(instance, C.gpointer(uintptr(signalUndoManagerCanRedoChangedId)))

	detail := signalUndoManagerCanRedoChangedDetail{callback, handlerID}
	signalUndoManagerCanRedoChangedMap[signalUndoManagerCanRedoChangedId] = detail

	return signalUndoManagerCanRedoChangedId
}

/*
DisconnectCanRedoChanged disconnects a callback from the 'can-redo-changed' signal for the UndoManager.

The connectionID should be a value returned from a call to ConnectCanRedoChanged.
*/
func (recv *UndoManager) DisconnectCanRedoChanged(connectionID int) {
	signalUndoManagerCanRedoChangedLock.Lock()
	defer signalUndoManagerCanRedoChangedLock.Unlock()

	detail, exists := signalUndoManagerCanRedoChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUndoManagerCanRedoChangedMap, connectionID)
}

//export undomanager_canRedoChangedHandler
func undomanager_canRedoChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalUndoManagerCanRedoChangedLock.RLock()
	defer signalUndoManagerCanRedoChangedLock.RUnlock()

	targetObject := UndoManagerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalUndoManagerCanRedoChangedMap[index].callback
	callback(targetObject)
}

type signalUndoManagerCanUndoChangedDetail struct {
	callback  UndoManagerSignalCanUndoChangedCallback
	handlerID C.gulong
}

var signalUndoManagerCanUndoChangedId int
var signalUndoManagerCanUndoChangedMap = make(map[int]signalUndoManagerCanUndoChangedDetail)
var signalUndoManagerCanUndoChangedLock sync.RWMutex

// UndoManagerSignalCanUndoChangedCallback is a callback function for a 'can-undo-changed' signal emitted from a UndoManager.
type UndoManagerSignalCanUndoChangedCallback func(targetObject *UndoManager)

/*
ConnectCanUndoChanged connects the callback to the 'can-undo-changed' signal for the UndoManager.

The returned value represents the connection, and may be passed to DisconnectCanUndoChanged to remove it.
*/
func (recv *UndoManager) ConnectCanUndoChanged(callback UndoManagerSignalCanUndoChangedCallback) int {
	signalUndoManagerCanUndoChangedLock.Lock()
	defer signalUndoManagerCanUndoChangedLock.Unlock()

	signalUndoManagerCanUndoChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.UndoManager_signal_connect_can_undo_changed(instance, C.gpointer(uintptr(signalUndoManagerCanUndoChangedId)))

	detail := signalUndoManagerCanUndoChangedDetail{callback, handlerID}
	signalUndoManagerCanUndoChangedMap[signalUndoManagerCanUndoChangedId] = detail

	return signalUndoManagerCanUndoChangedId
}

/*
DisconnectCanUndoChanged disconnects a callback from the 'can-undo-changed' signal for the UndoManager.

The connectionID should be a value returned from a call to ConnectCanUndoChanged.
*/
func (recv *UndoManager) DisconnectCanUndoChanged(connectionID int) {
	signalUndoManagerCanUndoChangedLock.Lock()
	defer signalUndoManagerCanUndoChangedLock.Unlock()

	detail, exists := signalUndoManagerCanUndoChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUndoManagerCanUndoChangedMap, connectionID)
}

//export undomanager_canUndoChangedHandler
func undomanager_canUndoChangedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalUndoManagerCanUndoChangedLock.RLock()
	defer signalUndoManagerCanUndoChangedLock.RUnlock()

	targetObject := UndoManagerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalUndoManagerCanUndoChangedMap[index].callback
	callback(targetObject)
}

// BeginNotUndoableAction is a wrapper around the C function gtk_source_undo_manager_begin_not_undoable_action.
func (recv *UndoManager) BeginNotUndoableAction() {
	C.gtk_source_undo_manager_begin_not_undoable_action((*C.GtkSourceUndoManager)(recv.native))

	return
}

// CanRedo is a wrapper around the C function gtk_source_undo_manager_can_redo.
func (recv *UndoManager) CanRedo() bool {
	retC := C.gtk_source_undo_manager_can_redo((*C.GtkSourceUndoManager)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanRedoChanged is a wrapper around the C function gtk_source_undo_manager_can_redo_changed.
func (recv *UndoManager) CanRedoChanged() {
	C.gtk_source_undo_manager_can_redo_changed((*C.GtkSourceUndoManager)(recv.native))

	return
}

// CanUndo is a wrapper around the C function gtk_source_undo_manager_can_undo.
func (recv *UndoManager) CanUndo() bool {
	retC := C.gtk_source_undo_manager_can_undo((*C.GtkSourceUndoManager)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanUndoChanged is a wrapper around the C function gtk_source_undo_manager_can_undo_changed.
func (recv *UndoManager) CanUndoChanged() {
	C.gtk_source_undo_manager_can_undo_changed((*C.GtkSourceUndoManager)(recv.native))

	return
}

// EndNotUndoableAction is a wrapper around the C function gtk_source_undo_manager_end_not_undoable_action.
func (recv *UndoManager) EndNotUndoableAction() {
	C.gtk_source_undo_manager_end_not_undoable_action((*C.GtkSourceUndoManager)(recv.native))

	return
}

// Redo is a wrapper around the C function gtk_source_undo_manager_redo.
func (recv *UndoManager) Redo() {
	C.gtk_source_undo_manager_redo((*C.GtkSourceUndoManager)(recv.native))

	return
}

// Undo is a wrapper around the C function gtk_source_undo_manager_undo.
func (recv *UndoManager) Undo() {
	C.gtk_source_undo_manager_undo((*C.GtkSourceUndoManager)(recv.native))

	return
}

// BufferClass is a wrapper around the C record GtkSourceBufferClass.
type BufferClass struct {
	native *C.GtkSourceBufferClass
	// parent_class : record
	// no type for undo
	// no type for redo
	// no type for bracket_matched
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
	// no type for _gtk_source_reserved3
}

func BufferClassNewFromC(u unsafe.Pointer) *BufferClass {
	c := (*C.GtkSourceBufferClass)(u)
	if c == nil {
		return nil
	}

	g := &BufferClass{native: c}

	return g
}

func (recv *BufferClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferClass with another BufferClass, and returns true if they represent the same GObject.
func (recv *BufferClass) Equals(other *BufferClass) bool {
	return other.ToC() == recv.ToC()
}

// BufferPrivate is a wrapper around the C record GtkSourceBufferPrivate.
type BufferPrivate struct {
	native *C.GtkSourceBufferPrivate
}

func BufferPrivateNewFromC(u unsafe.Pointer) *BufferPrivate {
	c := (*C.GtkSourceBufferPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BufferPrivate{native: c}

	return g
}

func (recv *BufferPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BufferPrivate with another BufferPrivate, and returns true if they represent the same GObject.
func (recv *BufferPrivate) Equals(other *BufferPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CompletionClass is a wrapper around the C record GtkSourceCompletionClass.
type CompletionClass struct {
	native *C.GtkSourceCompletionClass
	// parent_class : record
	// no type for proposal_activated
	// no type for show
	// no type for hide
	// no type for populate_context
	// no type for move_cursor
	// no type for move_page
	// no type for activate_proposal
}

func CompletionClassNewFromC(u unsafe.Pointer) *CompletionClass {
	c := (*C.GtkSourceCompletionClass)(u)
	if c == nil {
		return nil
	}

	g := &CompletionClass{native: c}

	return g
}

func (recv *CompletionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionClass with another CompletionClass, and returns true if they represent the same GObject.
func (recv *CompletionClass) Equals(other *CompletionClass) bool {
	return other.ToC() == recv.ToC()
}

// CompletionContextClass is a wrapper around the C record GtkSourceCompletionContextClass.
type CompletionContextClass struct {
	native *C.GtkSourceCompletionContextClass
	// parent_class : record
	// no type for cancelled
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
	// no type for _gtk_source_reserved3
}

func CompletionContextClassNewFromC(u unsafe.Pointer) *CompletionContextClass {
	c := (*C.GtkSourceCompletionContextClass)(u)
	if c == nil {
		return nil
	}

	g := &CompletionContextClass{native: c}

	return g
}

func (recv *CompletionContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionContextClass with another CompletionContextClass, and returns true if they represent the same GObject.
func (recv *CompletionContextClass) Equals(other *CompletionContextClass) bool {
	return other.ToC() == recv.ToC()
}

// CompletionContextPrivate is a wrapper around the C record GtkSourceCompletionContextPrivate.
type CompletionContextPrivate struct {
	native *C.GtkSourceCompletionContextPrivate
}

func CompletionContextPrivateNewFromC(u unsafe.Pointer) *CompletionContextPrivate {
	c := (*C.GtkSourceCompletionContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CompletionContextPrivate{native: c}

	return g
}

func (recv *CompletionContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionContextPrivate with another CompletionContextPrivate, and returns true if they represent the same GObject.
func (recv *CompletionContextPrivate) Equals(other *CompletionContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CompletionInfoClass is a wrapper around the C record GtkSourceCompletionInfoClass.
type CompletionInfoClass struct {
	native *C.GtkSourceCompletionInfoClass
	// parent_class : record
	// no type for before_show
}

func CompletionInfoClassNewFromC(u unsafe.Pointer) *CompletionInfoClass {
	c := (*C.GtkSourceCompletionInfoClass)(u)
	if c == nil {
		return nil
	}

	g := &CompletionInfoClass{native: c}

	return g
}

func (recv *CompletionInfoClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionInfoClass with another CompletionInfoClass, and returns true if they represent the same GObject.
func (recv *CompletionInfoClass) Equals(other *CompletionInfoClass) bool {
	return other.ToC() == recv.ToC()
}

// CompletionInfoPrivate is a wrapper around the C record GtkSourceCompletionInfoPrivate.
type CompletionInfoPrivate struct {
	native *C.GtkSourceCompletionInfoPrivate
}

func CompletionInfoPrivateNewFromC(u unsafe.Pointer) *CompletionInfoPrivate {
	c := (*C.GtkSourceCompletionInfoPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CompletionInfoPrivate{native: c}

	return g
}

func (recv *CompletionInfoPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionInfoPrivate with another CompletionInfoPrivate, and returns true if they represent the same GObject.
func (recv *CompletionInfoPrivate) Equals(other *CompletionInfoPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CompletionItemClass is a wrapper around the C record GtkSourceCompletionItemClass.
type CompletionItemClass struct {
	native *C.GtkSourceCompletionItemClass
	// parent_class : record
}

func CompletionItemClassNewFromC(u unsafe.Pointer) *CompletionItemClass {
	c := (*C.GtkSourceCompletionItemClass)(u)
	if c == nil {
		return nil
	}

	g := &CompletionItemClass{native: c}

	return g
}

func (recv *CompletionItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionItemClass with another CompletionItemClass, and returns true if they represent the same GObject.
func (recv *CompletionItemClass) Equals(other *CompletionItemClass) bool {
	return other.ToC() == recv.ToC()
}

// CompletionItemPrivate is a wrapper around the C record GtkSourceCompletionItemPrivate.
type CompletionItemPrivate struct {
	native *C.GtkSourceCompletionItemPrivate
}

func CompletionItemPrivateNewFromC(u unsafe.Pointer) *CompletionItemPrivate {
	c := (*C.GtkSourceCompletionItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CompletionItemPrivate{native: c}

	return g
}

func (recv *CompletionItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionItemPrivate with another CompletionItemPrivate, and returns true if they represent the same GObject.
func (recv *CompletionItemPrivate) Equals(other *CompletionItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CompletionPrivate is a wrapper around the C record GtkSourceCompletionPrivate.
type CompletionPrivate struct {
	native *C.GtkSourceCompletionPrivate
}

func CompletionPrivateNewFromC(u unsafe.Pointer) *CompletionPrivate {
	c := (*C.GtkSourceCompletionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CompletionPrivate{native: c}

	return g
}

func (recv *CompletionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionPrivate with another CompletionPrivate, and returns true if they represent the same GObject.
func (recv *CompletionPrivate) Equals(other *CompletionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CompletionProposalIface is a wrapper around the C record GtkSourceCompletionProposalIface.
type CompletionProposalIface struct {
	native *C.GtkSourceCompletionProposalIface
	// parent : record
	// no type for get_label
	// no type for get_markup
	// no type for get_text
	// no type for get_icon
	// no type for get_icon_name
	// no type for get_gicon
	// no type for get_info
	// no type for hash
	// no type for equal
	// no type for changed
}

func CompletionProposalIfaceNewFromC(u unsafe.Pointer) *CompletionProposalIface {
	c := (*C.GtkSourceCompletionProposalIface)(u)
	if c == nil {
		return nil
	}

	g := &CompletionProposalIface{native: c}

	return g
}

func (recv *CompletionProposalIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionProposalIface with another CompletionProposalIface, and returns true if they represent the same GObject.
func (recv *CompletionProposalIface) Equals(other *CompletionProposalIface) bool {
	return other.ToC() == recv.ToC()
}

// CompletionProviderIface is a wrapper around the C record GtkSourceCompletionProviderIface.
type CompletionProviderIface struct {
	native *C.GtkSourceCompletionProviderIface
	// g_iface : record
	// no type for get_name
	// no type for get_icon
	// no type for get_icon_name
	// no type for get_gicon
	// no type for populate
	// no type for match
	// no type for get_activation
	// no type for get_info_widget
	// no type for update_info
	// no type for get_start_iter
	// no type for activate_proposal
	// no type for get_interactive_delay
	// no type for get_priority
}

func CompletionProviderIfaceNewFromC(u unsafe.Pointer) *CompletionProviderIface {
	c := (*C.GtkSourceCompletionProviderIface)(u)
	if c == nil {
		return nil
	}

	g := &CompletionProviderIface{native: c}

	return g
}

func (recv *CompletionProviderIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionProviderIface with another CompletionProviderIface, and returns true if they represent the same GObject.
func (recv *CompletionProviderIface) Equals(other *CompletionProviderIface) bool {
	return other.ToC() == recv.ToC()
}

// CompletionWordsClass is a wrapper around the C record GtkSourceCompletionWordsClass.
type CompletionWordsClass struct {
	native *C.GtkSourceCompletionWordsClass
	// parent_class : record
}

func CompletionWordsClassNewFromC(u unsafe.Pointer) *CompletionWordsClass {
	c := (*C.GtkSourceCompletionWordsClass)(u)
	if c == nil {
		return nil
	}

	g := &CompletionWordsClass{native: c}

	return g
}

func (recv *CompletionWordsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionWordsClass with another CompletionWordsClass, and returns true if they represent the same GObject.
func (recv *CompletionWordsClass) Equals(other *CompletionWordsClass) bool {
	return other.ToC() == recv.ToC()
}

// CompletionWordsPrivate is a wrapper around the C record GtkSourceCompletionWordsPrivate.
type CompletionWordsPrivate struct {
	native *C.GtkSourceCompletionWordsPrivate
}

func CompletionWordsPrivateNewFromC(u unsafe.Pointer) *CompletionWordsPrivate {
	c := (*C.GtkSourceCompletionWordsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CompletionWordsPrivate{native: c}

	return g
}

func (recv *CompletionWordsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CompletionWordsPrivate with another CompletionWordsPrivate, and returns true if they represent the same GObject.
func (recv *CompletionWordsPrivate) Equals(other *CompletionWordsPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Encoding is a wrapper around the C record GtkSourceEncoding.
type Encoding struct {
	native *C.GtkSourceEncoding
}

func EncodingNewFromC(u unsafe.Pointer) *Encoding {
	c := (*C.GtkSourceEncoding)(u)
	if c == nil {
		return nil
	}

	g := &Encoding{native: c}

	return g
}

func (recv *Encoding) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Encoding with another Encoding, and returns true if they represent the same GObject.
func (recv *Encoding) Equals(other *Encoding) bool {
	return other.ToC() == recv.ToC()
}

// EncodingGetAll is a wrapper around the C function gtk_source_encoding_get_all.
func EncodingGetAll() *glib.SList {
	retC := C.gtk_source_encoding_get_all()
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EncodingGetCurrent is a wrapper around the C function gtk_source_encoding_get_current.
func EncodingGetCurrent() *Encoding {
	retC := C.gtk_source_encoding_get_current()
	retGo := EncodingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EncodingGetDefaultCandidates is a wrapper around the C function gtk_source_encoding_get_default_candidates.
func EncodingGetDefaultCandidates() *glib.SList {
	retC := C.gtk_source_encoding_get_default_candidates()
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EncodingGetFromCharset is a wrapper around the C function gtk_source_encoding_get_from_charset.
func EncodingGetFromCharset(charset string) *Encoding {
	c_charset := C.CString(charset)
	defer C.free(unsafe.Pointer(c_charset))

	retC := C.gtk_source_encoding_get_from_charset(c_charset)
	var retGo (*Encoding)
	if retC == nil {
		retGo = nil
	} else {
		retGo = EncodingNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// EncodingGetUtf8 is a wrapper around the C function gtk_source_encoding_get_utf8.
func EncodingGetUtf8() *Encoding {
	retC := C.gtk_source_encoding_get_utf8()
	retGo := EncodingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_source_encoding_copy.
func (recv *Encoding) Copy() *Encoding {
	retC := C.gtk_source_encoding_copy((*C.GtkSourceEncoding)(recv.native))
	retGo := EncodingNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_source_encoding_free.
func (recv *Encoding) Free() {
	C.gtk_source_encoding_free((*C.GtkSourceEncoding)(recv.native))

	return
}

// GetCharset is a wrapper around the C function gtk_source_encoding_get_charset.
func (recv *Encoding) GetCharset() string {
	retC := C.gtk_source_encoding_get_charset((*C.GtkSourceEncoding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetName is a wrapper around the C function gtk_source_encoding_get_name.
func (recv *Encoding) GetName() string {
	retC := C.gtk_source_encoding_get_name((*C.GtkSourceEncoding)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// ToString is a wrapper around the C function gtk_source_encoding_to_string.
func (recv *Encoding) ToString() string {
	retC := C.gtk_source_encoding_to_string((*C.GtkSourceEncoding)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FileClass is a wrapper around the C record GtkSourceFileClass.
type FileClass struct {
	native *C.GtkSourceFileClass
	// parent_class : record
	// no type for padding
}

func FileClassNewFromC(u unsafe.Pointer) *FileClass {
	c := (*C.GtkSourceFileClass)(u)
	if c == nil {
		return nil
	}

	g := &FileClass{native: c}

	return g
}

func (recv *FileClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileClass with another FileClass, and returns true if they represent the same GObject.
func (recv *FileClass) Equals(other *FileClass) bool {
	return other.ToC() == recv.ToC()
}

// FileLoaderClass is a wrapper around the C record GtkSourceFileLoaderClass.
type FileLoaderClass struct {
	native *C.GtkSourceFileLoaderClass
	// parent_class : record
	// no type for padding
}

func FileLoaderClassNewFromC(u unsafe.Pointer) *FileLoaderClass {
	c := (*C.GtkSourceFileLoaderClass)(u)
	if c == nil {
		return nil
	}

	g := &FileLoaderClass{native: c}

	return g
}

func (recv *FileLoaderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileLoaderClass with another FileLoaderClass, and returns true if they represent the same GObject.
func (recv *FileLoaderClass) Equals(other *FileLoaderClass) bool {
	return other.ToC() == recv.ToC()
}

// FileLoaderPrivate is a wrapper around the C record GtkSourceFileLoaderPrivate.
type FileLoaderPrivate struct {
	native *C.GtkSourceFileLoaderPrivate
}

func FileLoaderPrivateNewFromC(u unsafe.Pointer) *FileLoaderPrivate {
	c := (*C.GtkSourceFileLoaderPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileLoaderPrivate{native: c}

	return g
}

func (recv *FileLoaderPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileLoaderPrivate with another FileLoaderPrivate, and returns true if they represent the same GObject.
func (recv *FileLoaderPrivate) Equals(other *FileLoaderPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FilePrivate is a wrapper around the C record GtkSourceFilePrivate.
type FilePrivate struct {
	native *C.GtkSourceFilePrivate
}

func FilePrivateNewFromC(u unsafe.Pointer) *FilePrivate {
	c := (*C.GtkSourceFilePrivate)(u)
	if c == nil {
		return nil
	}

	g := &FilePrivate{native: c}

	return g
}

func (recv *FilePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FilePrivate with another FilePrivate, and returns true if they represent the same GObject.
func (recv *FilePrivate) Equals(other *FilePrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileSaverClass is a wrapper around the C record GtkSourceFileSaverClass.
type FileSaverClass struct {
	native *C.GtkSourceFileSaverClass
	// parent_class : record
	// no type for padding
}

func FileSaverClassNewFromC(u unsafe.Pointer) *FileSaverClass {
	c := (*C.GtkSourceFileSaverClass)(u)
	if c == nil {
		return nil
	}

	g := &FileSaverClass{native: c}

	return g
}

func (recv *FileSaverClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileSaverClass with another FileSaverClass, and returns true if they represent the same GObject.
func (recv *FileSaverClass) Equals(other *FileSaverClass) bool {
	return other.ToC() == recv.ToC()
}

// FileSaverPrivate is a wrapper around the C record GtkSourceFileSaverPrivate.
type FileSaverPrivate struct {
	native *C.GtkSourceFileSaverPrivate
}

func FileSaverPrivateNewFromC(u unsafe.Pointer) *FileSaverPrivate {
	c := (*C.GtkSourceFileSaverPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileSaverPrivate{native: c}

	return g
}

func (recv *FileSaverPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileSaverPrivate with another FileSaverPrivate, and returns true if they represent the same GObject.
func (recv *FileSaverPrivate) Equals(other *FileSaverPrivate) bool {
	return other.ToC() == recv.ToC()
}

// GutterClass is a wrapper around the C record GtkSourceGutterClass.
type GutterClass struct {
	native *C.GtkSourceGutterClass
	// parent_class : record
}

func GutterClassNewFromC(u unsafe.Pointer) *GutterClass {
	c := (*C.GtkSourceGutterClass)(u)
	if c == nil {
		return nil
	}

	g := &GutterClass{native: c}

	return g
}

func (recv *GutterClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterClass with another GutterClass, and returns true if they represent the same GObject.
func (recv *GutterClass) Equals(other *GutterClass) bool {
	return other.ToC() == recv.ToC()
}

// GutterPrivate is a wrapper around the C record GtkSourceGutterPrivate.
type GutterPrivate struct {
	native *C.GtkSourceGutterPrivate
}

func GutterPrivateNewFromC(u unsafe.Pointer) *GutterPrivate {
	c := (*C.GtkSourceGutterPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GutterPrivate{native: c}

	return g
}

func (recv *GutterPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterPrivate with another GutterPrivate, and returns true if they represent the same GObject.
func (recv *GutterPrivate) Equals(other *GutterPrivate) bool {
	return other.ToC() == recv.ToC()
}

// GutterRendererClass is a wrapper around the C record GtkSourceGutterRendererClass.
type GutterRendererClass struct {
	native *C.GtkSourceGutterRendererClass
	// parent_class : record
	// no type for begin
	// no type for draw
	// no type for end
	// no type for change_view
	// no type for change_buffer
	// no type for query_activatable
	// no type for activate
	// no type for queue_draw
	// no type for query_tooltip
	// no type for query_data
}

func GutterRendererClassNewFromC(u unsafe.Pointer) *GutterRendererClass {
	c := (*C.GtkSourceGutterRendererClass)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererClass{native: c}

	return g
}

func (recv *GutterRendererClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererClass with another GutterRendererClass, and returns true if they represent the same GObject.
func (recv *GutterRendererClass) Equals(other *GutterRendererClass) bool {
	return other.ToC() == recv.ToC()
}

// GutterRendererPixbufClass is a wrapper around the C record GtkSourceGutterRendererPixbufClass.
type GutterRendererPixbufClass struct {
	native *C.GtkSourceGutterRendererPixbufClass
	// Private : parent_class
}

func GutterRendererPixbufClassNewFromC(u unsafe.Pointer) *GutterRendererPixbufClass {
	c := (*C.GtkSourceGutterRendererPixbufClass)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererPixbufClass{native: c}

	return g
}

func (recv *GutterRendererPixbufClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererPixbufClass with another GutterRendererPixbufClass, and returns true if they represent the same GObject.
func (recv *GutterRendererPixbufClass) Equals(other *GutterRendererPixbufClass) bool {
	return other.ToC() == recv.ToC()
}

// GutterRendererPixbufPrivate is a wrapper around the C record GtkSourceGutterRendererPixbufPrivate.
type GutterRendererPixbufPrivate struct {
	native *C.GtkSourceGutterRendererPixbufPrivate
}

func GutterRendererPixbufPrivateNewFromC(u unsafe.Pointer) *GutterRendererPixbufPrivate {
	c := (*C.GtkSourceGutterRendererPixbufPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererPixbufPrivate{native: c}

	return g
}

func (recv *GutterRendererPixbufPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererPixbufPrivate with another GutterRendererPixbufPrivate, and returns true if they represent the same GObject.
func (recv *GutterRendererPixbufPrivate) Equals(other *GutterRendererPixbufPrivate) bool {
	return other.ToC() == recv.ToC()
}

// GutterRendererPrivate is a wrapper around the C record GtkSourceGutterRendererPrivate.
type GutterRendererPrivate struct {
	native *C.GtkSourceGutterRendererPrivate
}

func GutterRendererPrivateNewFromC(u unsafe.Pointer) *GutterRendererPrivate {
	c := (*C.GtkSourceGutterRendererPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererPrivate{native: c}

	return g
}

func (recv *GutterRendererPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererPrivate with another GutterRendererPrivate, and returns true if they represent the same GObject.
func (recv *GutterRendererPrivate) Equals(other *GutterRendererPrivate) bool {
	return other.ToC() == recv.ToC()
}

// GutterRendererTextClass is a wrapper around the C record GtkSourceGutterRendererTextClass.
type GutterRendererTextClass struct {
	native *C.GtkSourceGutterRendererTextClass
	// Private : parent_class
}

func GutterRendererTextClassNewFromC(u unsafe.Pointer) *GutterRendererTextClass {
	c := (*C.GtkSourceGutterRendererTextClass)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererTextClass{native: c}

	return g
}

func (recv *GutterRendererTextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererTextClass with another GutterRendererTextClass, and returns true if they represent the same GObject.
func (recv *GutterRendererTextClass) Equals(other *GutterRendererTextClass) bool {
	return other.ToC() == recv.ToC()
}

// GutterRendererTextPrivate is a wrapper around the C record GtkSourceGutterRendererTextPrivate.
type GutterRendererTextPrivate struct {
	native *C.GtkSourceGutterRendererTextPrivate
}

func GutterRendererTextPrivateNewFromC(u unsafe.Pointer) *GutterRendererTextPrivate {
	c := (*C.GtkSourceGutterRendererTextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GutterRendererTextPrivate{native: c}

	return g
}

func (recv *GutterRendererTextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GutterRendererTextPrivate with another GutterRendererTextPrivate, and returns true if they represent the same GObject.
func (recv *GutterRendererTextPrivate) Equals(other *GutterRendererTextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// LanguageClass is a wrapper around the C record GtkSourceLanguageClass.
type LanguageClass struct {
	native *C.GtkSourceLanguageClass
	// parent_class : record
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
}

func LanguageClassNewFromC(u unsafe.Pointer) *LanguageClass {
	c := (*C.GtkSourceLanguageClass)(u)
	if c == nil {
		return nil
	}

	g := &LanguageClass{native: c}

	return g
}

func (recv *LanguageClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LanguageClass with another LanguageClass, and returns true if they represent the same GObject.
func (recv *LanguageClass) Equals(other *LanguageClass) bool {
	return other.ToC() == recv.ToC()
}

// LanguageManagerClass is a wrapper around the C record GtkSourceLanguageManagerClass.
type LanguageManagerClass struct {
	native *C.GtkSourceLanguageManagerClass
	// parent_class : record
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
	// no type for _gtk_source_reserved3
	// no type for _gtk_source_reserved4
}

func LanguageManagerClassNewFromC(u unsafe.Pointer) *LanguageManagerClass {
	c := (*C.GtkSourceLanguageManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &LanguageManagerClass{native: c}

	return g
}

func (recv *LanguageManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LanguageManagerClass with another LanguageManagerClass, and returns true if they represent the same GObject.
func (recv *LanguageManagerClass) Equals(other *LanguageManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// LanguageManagerPrivate is a wrapper around the C record GtkSourceLanguageManagerPrivate.
type LanguageManagerPrivate struct {
	native *C.GtkSourceLanguageManagerPrivate
}

func LanguageManagerPrivateNewFromC(u unsafe.Pointer) *LanguageManagerPrivate {
	c := (*C.GtkSourceLanguageManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &LanguageManagerPrivate{native: c}

	return g
}

func (recv *LanguageManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LanguageManagerPrivate with another LanguageManagerPrivate, and returns true if they represent the same GObject.
func (recv *LanguageManagerPrivate) Equals(other *LanguageManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// LanguagePrivate is a wrapper around the C record GtkSourceLanguagePrivate.
type LanguagePrivate struct {
	native *C.GtkSourceLanguagePrivate
}

func LanguagePrivateNewFromC(u unsafe.Pointer) *LanguagePrivate {
	c := (*C.GtkSourceLanguagePrivate)(u)
	if c == nil {
		return nil
	}

	g := &LanguagePrivate{native: c}

	return g
}

func (recv *LanguagePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LanguagePrivate with another LanguagePrivate, and returns true if they represent the same GObject.
func (recv *LanguagePrivate) Equals(other *LanguagePrivate) bool {
	return other.ToC() == recv.ToC()
}

// MapClass is a wrapper around the C record GtkSourceMapClass.
type MapClass struct {
	native *C.GtkSourceMapClass
	// parent_class : record
	// no type for padding
}

func MapClassNewFromC(u unsafe.Pointer) *MapClass {
	c := (*C.GtkSourceMapClass)(u)
	if c == nil {
		return nil
	}

	g := &MapClass{native: c}

	return g
}

func (recv *MapClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MapClass with another MapClass, and returns true if they represent the same GObject.
func (recv *MapClass) Equals(other *MapClass) bool {
	return other.ToC() == recv.ToC()
}

// MarkAttributesClass is a wrapper around the C record GtkSourceMarkAttributesClass.
type MarkAttributesClass struct {
	native *C.GtkSourceMarkAttributesClass
	// Private : parent_class
}

func MarkAttributesClassNewFromC(u unsafe.Pointer) *MarkAttributesClass {
	c := (*C.GtkSourceMarkAttributesClass)(u)
	if c == nil {
		return nil
	}

	g := &MarkAttributesClass{native: c}

	return g
}

func (recv *MarkAttributesClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkAttributesClass with another MarkAttributesClass, and returns true if they represent the same GObject.
func (recv *MarkAttributesClass) Equals(other *MarkAttributesClass) bool {
	return other.ToC() == recv.ToC()
}

// MarkAttributesPrivate is a wrapper around the C record GtkSourceMarkAttributesPrivate.
type MarkAttributesPrivate struct {
	native *C.GtkSourceMarkAttributesPrivate
}

func MarkAttributesPrivateNewFromC(u unsafe.Pointer) *MarkAttributesPrivate {
	c := (*C.GtkSourceMarkAttributesPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MarkAttributesPrivate{native: c}

	return g
}

func (recv *MarkAttributesPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkAttributesPrivate with another MarkAttributesPrivate, and returns true if they represent the same GObject.
func (recv *MarkAttributesPrivate) Equals(other *MarkAttributesPrivate) bool {
	return other.ToC() == recv.ToC()
}

// MarkClass is a wrapper around the C record GtkSourceMarkClass.
type MarkClass struct {
	native *C.GtkSourceMarkClass
	// parent_class : record
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
}

func MarkClassNewFromC(u unsafe.Pointer) *MarkClass {
	c := (*C.GtkSourceMarkClass)(u)
	if c == nil {
		return nil
	}

	g := &MarkClass{native: c}

	return g
}

func (recv *MarkClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkClass with another MarkClass, and returns true if they represent the same GObject.
func (recv *MarkClass) Equals(other *MarkClass) bool {
	return other.ToC() == recv.ToC()
}

// MarkPrivate is a wrapper around the C record GtkSourceMarkPrivate.
type MarkPrivate struct {
	native *C.GtkSourceMarkPrivate
}

func MarkPrivateNewFromC(u unsafe.Pointer) *MarkPrivate {
	c := (*C.GtkSourceMarkPrivate)(u)
	if c == nil {
		return nil
	}

	g := &MarkPrivate{native: c}

	return g
}

func (recv *MarkPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MarkPrivate with another MarkPrivate, and returns true if they represent the same GObject.
func (recv *MarkPrivate) Equals(other *MarkPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PrintCompositorClass is a wrapper around the C record GtkSourcePrintCompositorClass.
type PrintCompositorClass struct {
	native *C.GtkSourcePrintCompositorClass
	// parent_class : record
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
}

func PrintCompositorClassNewFromC(u unsafe.Pointer) *PrintCompositorClass {
	c := (*C.GtkSourcePrintCompositorClass)(u)
	if c == nil {
		return nil
	}

	g := &PrintCompositorClass{native: c}

	return g
}

func (recv *PrintCompositorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintCompositorClass with another PrintCompositorClass, and returns true if they represent the same GObject.
func (recv *PrintCompositorClass) Equals(other *PrintCompositorClass) bool {
	return other.ToC() == recv.ToC()
}

// PrintCompositorPrivate is a wrapper around the C record GtkSourcePrintCompositorPrivate.
type PrintCompositorPrivate struct {
	native *C.GtkSourcePrintCompositorPrivate
}

func PrintCompositorPrivateNewFromC(u unsafe.Pointer) *PrintCompositorPrivate {
	c := (*C.GtkSourcePrintCompositorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PrintCompositorPrivate{native: c}

	return g
}

func (recv *PrintCompositorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintCompositorPrivate with another PrintCompositorPrivate, and returns true if they represent the same GObject.
func (recv *PrintCompositorPrivate) Equals(other *PrintCompositorPrivate) bool {
	return other.ToC() == recv.ToC()
}

// RegionClass is a wrapper around the C record GtkSourceRegionClass.
type RegionClass struct {
	native *C.GtkSourceRegionClass
	// parent_class : record
	// no type for padding
}

func RegionClassNewFromC(u unsafe.Pointer) *RegionClass {
	c := (*C.GtkSourceRegionClass)(u)
	if c == nil {
		return nil
	}

	g := &RegionClass{native: c}

	return g
}

func (recv *RegionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RegionClass with another RegionClass, and returns true if they represent the same GObject.
func (recv *RegionClass) Equals(other *RegionClass) bool {
	return other.ToC() == recv.ToC()
}

// RegionIter is a wrapper around the C record GtkSourceRegionIter.
type RegionIter struct {
	native *C.GtkSourceRegionIter
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
}

func RegionIterNewFromC(u unsafe.Pointer) *RegionIter {
	c := (*C.GtkSourceRegionIter)(u)
	if c == nil {
		return nil
	}

	g := &RegionIter{native: c}

	return g
}

func (recv *RegionIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RegionIter with another RegionIter, and returns true if they represent the same GObject.
func (recv *RegionIter) Equals(other *RegionIter) bool {
	return other.ToC() == recv.ToC()
}

// GetSubregion is a wrapper around the C function gtk_source_region_iter_get_subregion.
func (recv *RegionIter) GetSubregion() (bool, *gtk.TextIter, *gtk.TextIter) {
	var c_start C.GtkTextIter

	var c_end C.GtkTextIter

	retC := C.gtk_source_region_iter_get_subregion((*C.GtkSourceRegionIter)(recv.native), &c_start, &c_end)
	retGo := retC == C.TRUE

	start := gtk.TextIterNewFromC(unsafe.Pointer(&c_start))

	end := gtk.TextIterNewFromC(unsafe.Pointer(&c_end))

	return retGo, start, end
}

// IsEnd is a wrapper around the C function gtk_source_region_iter_is_end.
func (recv *RegionIter) IsEnd() bool {
	retC := C.gtk_source_region_iter_is_end((*C.GtkSourceRegionIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Next is a wrapper around the C function gtk_source_region_iter_next.
func (recv *RegionIter) Next() bool {
	retC := C.gtk_source_region_iter_next((*C.GtkSourceRegionIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SearchContextClass is a wrapper around the C record GtkSourceSearchContextClass.
type SearchContextClass struct {
	native *C.GtkSourceSearchContextClass
	// parent_class : record
	// no type for padding
}

func SearchContextClassNewFromC(u unsafe.Pointer) *SearchContextClass {
	c := (*C.GtkSourceSearchContextClass)(u)
	if c == nil {
		return nil
	}

	g := &SearchContextClass{native: c}

	return g
}

func (recv *SearchContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SearchContextClass with another SearchContextClass, and returns true if they represent the same GObject.
func (recv *SearchContextClass) Equals(other *SearchContextClass) bool {
	return other.ToC() == recv.ToC()
}

// SearchContextPrivate is a wrapper around the C record GtkSourceSearchContextPrivate.
type SearchContextPrivate struct {
	native *C.GtkSourceSearchContextPrivate
}

func SearchContextPrivateNewFromC(u unsafe.Pointer) *SearchContextPrivate {
	c := (*C.GtkSourceSearchContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SearchContextPrivate{native: c}

	return g
}

func (recv *SearchContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SearchContextPrivate with another SearchContextPrivate, and returns true if they represent the same GObject.
func (recv *SearchContextPrivate) Equals(other *SearchContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SearchSettingsClass is a wrapper around the C record GtkSourceSearchSettingsClass.
type SearchSettingsClass struct {
	native *C.GtkSourceSearchSettingsClass
	// parent_class : record
	// no type for padding
}

func SearchSettingsClassNewFromC(u unsafe.Pointer) *SearchSettingsClass {
	c := (*C.GtkSourceSearchSettingsClass)(u)
	if c == nil {
		return nil
	}

	g := &SearchSettingsClass{native: c}

	return g
}

func (recv *SearchSettingsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SearchSettingsClass with another SearchSettingsClass, and returns true if they represent the same GObject.
func (recv *SearchSettingsClass) Equals(other *SearchSettingsClass) bool {
	return other.ToC() == recv.ToC()
}

// SearchSettingsPrivate is a wrapper around the C record GtkSourceSearchSettingsPrivate.
type SearchSettingsPrivate struct {
	native *C.GtkSourceSearchSettingsPrivate
}

func SearchSettingsPrivateNewFromC(u unsafe.Pointer) *SearchSettingsPrivate {
	c := (*C.GtkSourceSearchSettingsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SearchSettingsPrivate{native: c}

	return g
}

func (recv *SearchSettingsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SearchSettingsPrivate with another SearchSettingsPrivate, and returns true if they represent the same GObject.
func (recv *SearchSettingsPrivate) Equals(other *SearchSettingsPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GtkSourceSpaceDrawerPrivate

// StyleClass is a wrapper around the C record GtkSourceStyleClass.
type StyleClass struct {
	native *C.GtkSourceStyleClass
}

func StyleClassNewFromC(u unsafe.Pointer) *StyleClass {
	c := (*C.GtkSourceStyleClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleClass{native: c}

	return g
}

func (recv *StyleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleClass with another StyleClass, and returns true if they represent the same GObject.
func (recv *StyleClass) Equals(other *StyleClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemeChooserButtonClass is a wrapper around the C record GtkSourceStyleSchemeChooserButtonClass.
type StyleSchemeChooserButtonClass struct {
	native *C.GtkSourceStyleSchemeChooserButtonClass
	// parent : record
}

func StyleSchemeChooserButtonClassNewFromC(u unsafe.Pointer) *StyleSchemeChooserButtonClass {
	c := (*C.GtkSourceStyleSchemeChooserButtonClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeChooserButtonClass{native: c}

	return g
}

func (recv *StyleSchemeChooserButtonClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeChooserButtonClass with another StyleSchemeChooserButtonClass, and returns true if they represent the same GObject.
func (recv *StyleSchemeChooserButtonClass) Equals(other *StyleSchemeChooserButtonClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemeChooserInterface is a wrapper around the C record GtkSourceStyleSchemeChooserInterface.
type StyleSchemeChooserInterface struct {
	native *C.GtkSourceStyleSchemeChooserInterface
	// base_interface : record
	// no type for get_style_scheme
	// no type for set_style_scheme
	// no type for padding
}

func StyleSchemeChooserInterfaceNewFromC(u unsafe.Pointer) *StyleSchemeChooserInterface {
	c := (*C.GtkSourceStyleSchemeChooserInterface)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeChooserInterface{native: c}

	return g
}

func (recv *StyleSchemeChooserInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeChooserInterface with another StyleSchemeChooserInterface, and returns true if they represent the same GObject.
func (recv *StyleSchemeChooserInterface) Equals(other *StyleSchemeChooserInterface) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemeChooserWidgetClass is a wrapper around the C record GtkSourceStyleSchemeChooserWidgetClass.
type StyleSchemeChooserWidgetClass struct {
	native *C.GtkSourceStyleSchemeChooserWidgetClass
	// parent : record
}

func StyleSchemeChooserWidgetClassNewFromC(u unsafe.Pointer) *StyleSchemeChooserWidgetClass {
	c := (*C.GtkSourceStyleSchemeChooserWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeChooserWidgetClass{native: c}

	return g
}

func (recv *StyleSchemeChooserWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeChooserWidgetClass with another StyleSchemeChooserWidgetClass, and returns true if they represent the same GObject.
func (recv *StyleSchemeChooserWidgetClass) Equals(other *StyleSchemeChooserWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemeClass is a wrapper around the C record GtkSourceStyleSchemeClass.
type StyleSchemeClass struct {
	native *C.GtkSourceStyleSchemeClass
	// base_class : record
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
}

func StyleSchemeClassNewFromC(u unsafe.Pointer) *StyleSchemeClass {
	c := (*C.GtkSourceStyleSchemeClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeClass{native: c}

	return g
}

func (recv *StyleSchemeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeClass with another StyleSchemeClass, and returns true if they represent the same GObject.
func (recv *StyleSchemeClass) Equals(other *StyleSchemeClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemeManagerClass is a wrapper around the C record GtkSourceStyleSchemeManagerClass.
type StyleSchemeManagerClass struct {
	native *C.GtkSourceStyleSchemeManagerClass
	// parent_class : record
	// no type for _gtk_source_reserved1
	// no type for _gtk_source_reserved2
	// no type for _gtk_source_reserved3
	// no type for _gtk_source_reserved4
}

func StyleSchemeManagerClassNewFromC(u unsafe.Pointer) *StyleSchemeManagerClass {
	c := (*C.GtkSourceStyleSchemeManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeManagerClass{native: c}

	return g
}

func (recv *StyleSchemeManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeManagerClass with another StyleSchemeManagerClass, and returns true if they represent the same GObject.
func (recv *StyleSchemeManagerClass) Equals(other *StyleSchemeManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemeManagerPrivate is a wrapper around the C record GtkSourceStyleSchemeManagerPrivate.
type StyleSchemeManagerPrivate struct {
	native *C.GtkSourceStyleSchemeManagerPrivate
}

func StyleSchemeManagerPrivateNewFromC(u unsafe.Pointer) *StyleSchemeManagerPrivate {
	c := (*C.GtkSourceStyleSchemeManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemeManagerPrivate{native: c}

	return g
}

func (recv *StyleSchemeManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemeManagerPrivate with another StyleSchemeManagerPrivate, and returns true if they represent the same GObject.
func (recv *StyleSchemeManagerPrivate) Equals(other *StyleSchemeManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// StyleSchemePrivate is a wrapper around the C record GtkSourceStyleSchemePrivate.
type StyleSchemePrivate struct {
	native *C.GtkSourceStyleSchemePrivate
}

func StyleSchemePrivateNewFromC(u unsafe.Pointer) *StyleSchemePrivate {
	c := (*C.GtkSourceStyleSchemePrivate)(u)
	if c == nil {
		return nil
	}

	g := &StyleSchemePrivate{native: c}

	return g
}

func (recv *StyleSchemePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this StyleSchemePrivate with another StyleSchemePrivate, and returns true if they represent the same GObject.
func (recv *StyleSchemePrivate) Equals(other *StyleSchemePrivate) bool {
	return other.ToC() == recv.ToC()
}

// TagClass is a wrapper around the C record GtkSourceTagClass.
type TagClass struct {
	native *C.GtkSourceTagClass
	// parent_class : record
	// no type for padding
}

func TagClassNewFromC(u unsafe.Pointer) *TagClass {
	c := (*C.GtkSourceTagClass)(u)
	if c == nil {
		return nil
	}

	g := &TagClass{native: c}

	return g
}

func (recv *TagClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TagClass with another TagClass, and returns true if they represent the same GObject.
func (recv *TagClass) Equals(other *TagClass) bool {
	return other.ToC() == recv.ToC()
}

// UndoManagerIface is a wrapper around the C record GtkSourceUndoManagerIface.
type UndoManagerIface struct {
	native *C.GtkSourceUndoManagerIface
	// parent : record
	// no type for can_undo
	// no type for can_redo
	// no type for undo
	// no type for redo
	// no type for begin_not_undoable_action
	// no type for end_not_undoable_action
	// no type for can_undo_changed
	// no type for can_redo_changed
}

func UndoManagerIfaceNewFromC(u unsafe.Pointer) *UndoManagerIface {
	c := (*C.GtkSourceUndoManagerIface)(u)
	if c == nil {
		return nil
	}

	g := &UndoManagerIface{native: c}

	return g
}

func (recv *UndoManagerIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UndoManagerIface with another UndoManagerIface, and returns true if they represent the same GObject.
func (recv *UndoManagerIface) Equals(other *UndoManagerIface) bool {
	return other.ToC() == recv.ToC()
}

// ViewClass is a wrapper around the C record GtkSourceViewClass.
type ViewClass struct {
	native *C.GtkSourceViewClass
	// parent_class : record
	// no type for undo
	// no type for redo
	// no type for line_mark_activated
	// no type for show_completion
	// no type for move_lines
	// no type for move_words
}

func ViewClassNewFromC(u unsafe.Pointer) *ViewClass {
	c := (*C.GtkSourceViewClass)(u)
	if c == nil {
		return nil
	}

	g := &ViewClass{native: c}

	return g
}

func (recv *ViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ViewClass with another ViewClass, and returns true if they represent the same GObject.
func (recv *ViewClass) Equals(other *ViewClass) bool {
	return other.ToC() == recv.ToC()
}

// ViewPrivate is a wrapper around the C record GtkSourceViewPrivate.
type ViewPrivate struct {
	native *C.GtkSourceViewPrivate
}

func ViewPrivateNewFromC(u unsafe.Pointer) *ViewPrivate {
	c := (*C.GtkSourceViewPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ViewPrivate{native: c}

	return g
}

func (recv *ViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ViewPrivate with another ViewPrivate, and returns true if they represent the same GObject.
func (recv *ViewPrivate) Equals(other *ViewPrivate) bool {
	return other.ToC() == recv.ToC()
}
