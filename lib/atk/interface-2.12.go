// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import "sync"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
/*

	void document_pageChangedHandler(GObject *, gint, gpointer);

	static gulong Document_signal_connect_page_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "page-changed", G_CALLBACK(document_pageChangedHandler), data);
	}

*/
/*

	void value_valueChangedHandler(GObject *, gdouble, gchar*, gpointer);

	static gulong Value_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", G_CALLBACK(value_valueChangedHandler), data);
	}

*/
import "C"

type signalDocumentPageChangedDetail struct {
	callback  DocumentSignalPageChangedCallback
	handlerID C.gulong
}

var signalDocumentPageChangedId int
var signalDocumentPageChangedMap = make(map[int]signalDocumentPageChangedDetail)
var signalDocumentPageChangedLock sync.RWMutex

// DocumentSignalPageChangedCallback is a callback function for a 'page-changed' signal emitted from a Document.
type DocumentSignalPageChangedCallback func(pageNumber int32)

/*
ConnectPageChanged connects the callback to the 'page-changed' signal for the Document.

The returned value represents the connection, and may be passed to DisconnectPageChanged to remove it.
*/
func (recv *Document) ConnectPageChanged(callback DocumentSignalPageChangedCallback) int {
	signalDocumentPageChangedLock.Lock()
	defer signalDocumentPageChangedLock.Unlock()

	signalDocumentPageChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Document_signal_connect_page_changed(instance, C.gpointer(uintptr(signalDocumentPageChangedId)))

	detail := signalDocumentPageChangedDetail{callback, handlerID}
	signalDocumentPageChangedMap[signalDocumentPageChangedId] = detail

	return signalDocumentPageChangedId
}

/*
DisconnectPageChanged disconnects a callback from the 'page-changed' signal for the Document.

The connectionID should be a value returned from a call to ConnectPageChanged.
*/
func (recv *Document) DisconnectPageChanged(connectionID int) {
	signalDocumentPageChangedLock.Lock()
	defer signalDocumentPageChangedLock.Unlock()

	detail, exists := signalDocumentPageChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDocumentPageChangedMap, connectionID)
}

//export document_pageChangedHandler
func document_pageChangedHandler(_ *C.GObject, c_page_number C.gint, data C.gpointer) {
	signalDocumentPageChangedLock.RLock()
	defer signalDocumentPageChangedLock.RUnlock()

	pageNumber := int32(c_page_number)

	index := int(uintptr(data))
	callback := signalDocumentPageChangedMap[index].callback
	callback(pageNumber)
}

// Blacklisted : atk_document_get_current_page_number

// Blacklisted : atk_document_get_page_count

// Unsupported : atk_table_cell_get_column_header_cells : array return type :

// Blacklisted : atk_table_cell_get_column_span

// Blacklisted : atk_table_cell_get_position

// Blacklisted : atk_table_cell_get_row_column_span

// Unsupported : atk_table_cell_get_row_header_cells : array return type :

// Blacklisted : atk_table_cell_get_row_span

// Blacklisted : atk_table_cell_get_table

type signalValueValueChangedDetail struct {
	callback  ValueSignalValueChangedCallback
	handlerID C.gulong
}

var signalValueValueChangedId int
var signalValueValueChangedMap = make(map[int]signalValueValueChangedDetail)
var signalValueValueChangedLock sync.RWMutex

// ValueSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a Value.
type ValueSignalValueChangedCallback func(value float64, text string)

/*
ConnectValueChanged connects the callback to the 'value-changed' signal for the Value.

The returned value represents the connection, and may be passed to DisconnectValueChanged to remove it.
*/
func (recv *Value) ConnectValueChanged(callback ValueSignalValueChangedCallback) int {
	signalValueValueChangedLock.Lock()
	defer signalValueValueChangedLock.Unlock()

	signalValueValueChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Value_signal_connect_value_changed(instance, C.gpointer(uintptr(signalValueValueChangedId)))

	detail := signalValueValueChangedDetail{callback, handlerID}
	signalValueValueChangedMap[signalValueValueChangedId] = detail

	return signalValueValueChangedId
}

/*
DisconnectValueChanged disconnects a callback from the 'value-changed' signal for the Value.

The connectionID should be a value returned from a call to ConnectValueChanged.
*/
func (recv *Value) DisconnectValueChanged(connectionID int) {
	signalValueValueChangedLock.Lock()
	defer signalValueValueChangedLock.Unlock()

	detail, exists := signalValueValueChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalValueValueChangedMap, connectionID)
}

//export value_valueChangedHandler
func value_valueChangedHandler(_ *C.GObject, c_value C.gdouble, c_text *C.gchar, data C.gpointer) {
	signalValueValueChangedLock.RLock()
	defer signalValueValueChangedLock.RUnlock()

	value := float64(c_value)

	text := C.GoString(c_text)

	index := int(uintptr(data))
	callback := signalValueValueChangedMap[index].callback
	callback(value, text)
}

// Blacklisted : atk_value_get_increment

// Blacklisted : atk_value_get_range

// Blacklisted : atk_value_get_sub_ranges

// Blacklisted : atk_value_get_value_and_text

// Blacklisted : atk_value_set_value
