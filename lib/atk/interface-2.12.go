// This is a generated file - DO NOT EDIT
// +build atk_2.12

package atk

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Unsupported signal 'page-changed' for Document : unsupported parameter page_number : type gint :

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

// Unsupported : atk_table_cell_get_column_header_cells : no return type

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

// Unsupported : atk_table_cell_get_row_header_cells : no return type

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

// Unsupported signal 'value-changed' for Value : unsupported parameter value : type gdouble :

// GetIncrement is a wrapper around the C function atk_value_get_increment.
func (recv *Value) GetIncrement() float64 {
	retC := C.atk_value_get_increment((*C.AtkValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetRange is a wrapper around the C function atk_value_get_range.
func (recv *Value) GetRange() *Range {
	retC := C.atk_value_get_range((*C.AtkValue)(recv.native))
	retGo := RangeNewFromC(unsafe.Pointer(retC))

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
