// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

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

// GetCurrentPageNumber is a wrapper around the C function atk_document_get_current_page_number.
func (recv *Document) GetCurrentPageNumber() int32 {
	retC := C.atk_document_get_current_page_number((*C.AtkDocument)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPageCount is a wrapper around the C function atk_document_get_page_count.
func (recv *Document) GetPageCount() int32 {
	retC := C.atk_document_get_page_count((*C.AtkDocument)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : atk_table_cell_get_column_header_cells : array return type :

// GetColumnSpan is a wrapper around the C function atk_table_cell_get_column_span.
func (recv *TableCell) GetColumnSpan() int32 {
	retC := C.atk_table_cell_get_column_span((*C.AtkTableCell)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPosition is a wrapper around the C function atk_table_cell_get_position.
func (recv *TableCell) GetPosition() (bool, int32, int32) {
	var c_row C.gint

	var c_column C.gint

	retC := C.atk_table_cell_get_position((*C.AtkTableCell)(recv.native), &c_row, &c_column)
	retGo := retC == C.TRUE

	row := (int32)(c_row)

	column := (int32)(c_column)

	return retGo, row, column
}

// GetRowColumnSpan is a wrapper around the C function atk_table_cell_get_row_column_span.
func (recv *TableCell) GetRowColumnSpan() (bool, int32, int32, int32, int32) {
	var c_row C.gint

	var c_column C.gint

	var c_row_span C.gint

	var c_column_span C.gint

	retC := C.atk_table_cell_get_row_column_span((*C.AtkTableCell)(recv.native), &c_row, &c_column, &c_row_span, &c_column_span)
	retGo := retC == C.TRUE

	row := (int32)(c_row)

	column := (int32)(c_column)

	rowSpan := (int32)(c_row_span)

	columnSpan := (int32)(c_column_span)

	return retGo, row, column, rowSpan, columnSpan
}

// Unsupported : atk_table_cell_get_row_header_cells : array return type :

// GetRowSpan is a wrapper around the C function atk_table_cell_get_row_span.
func (recv *TableCell) GetRowSpan() int32 {
	retC := C.atk_table_cell_get_row_span((*C.AtkTableCell)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTable is a wrapper around the C function atk_table_cell_get_table.
func (recv *TableCell) GetTable() *Object {
	retC := C.atk_table_cell_get_table((*C.AtkTableCell)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

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

// GetIncrement is a wrapper around the C function atk_value_get_increment.
func (recv *Value) GetIncrement() float64 {
	retC := C.atk_value_get_increment((*C.AtkValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetRange is a wrapper around the C function atk_value_get_range.
func (recv *Value) GetRange() *Range {
	retC := C.atk_value_get_range((*C.AtkValue)(recv.native))
	var retGo (*Range)
	if retC == nil {
		retGo = nil
	} else {
		retGo = RangeNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSubRanges is a wrapper around the C function atk_value_get_sub_ranges.
func (recv *Value) GetSubRanges() *glib.SList {
	retC := C.atk_value_get_sub_ranges((*C.AtkValue)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValueAndText is a wrapper around the C function atk_value_get_value_and_text.
func (recv *Value) GetValueAndText() (float64, string) {
	var c_value C.gdouble

	var c_text *C.gchar

	C.atk_value_get_value_and_text((*C.AtkValue)(recv.native), &c_value, &c_text)

	value := (float64)(c_value)

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	return value, text
}

// SetValue is a wrapper around the C function atk_value_set_value.
func (recv *Value) SetValue(newValue float64) {
	c_new_value := (C.gdouble)(newValue)

	C.atk_value_set_value((*C.AtkValue)(recv.native), c_new_value)

	return
}
