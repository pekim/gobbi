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

/*

C function : atk_document_get_current_page_number
*/
func (recv *Document) GetCurrentPageNumber() int32 {
	retC := C.atk_document_get_current_page_number((*C.AtkDocument)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

/*

C function : atk_document_get_page_count
*/
func (recv *Document) GetPageCount() int32 {
	retC := C.atk_document_get_page_count((*C.AtkDocument)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : atk_table_cell_get_column_header_cells : no return type

// Returns the number of columns occupied by this cell accessible.
/*

C function : atk_table_cell_get_column_span
*/
func (recv *TableCell) GetColumnSpan() int32 {
	retC := C.atk_table_cell_get_column_span((*C.AtkTableCell)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Retrieves the tabular position of this cell.
/*

C function : atk_table_cell_get_position
*/
func (recv *TableCell) GetPosition() (bool, int32, int32) {
	var c_row C.gint

	var c_column C.gint

	retC := C.atk_table_cell_get_position((*C.AtkTableCell)(recv.native), &c_row, &c_column)
	retGo := retC == C.TRUE

	row := (int32)(c_row)

	column := (int32)(c_column)

	return retGo, row, column
}

// Gets the row and column indexes and span of this cell accessible.
//
// Note: If the object does not implement this function, then, by default, atk
// will implement this function by calling get_row_span and get_column_span
// on the object.
/*

C function : atk_table_cell_get_row_column_span
*/
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

// Returns the number of rows occupied by this cell accessible.
/*

C function : atk_table_cell_get_row_span
*/
func (recv *TableCell) GetRowSpan() int32 {
	retC := C.atk_table_cell_get_row_span((*C.AtkTableCell)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns a reference to the accessible of the containing table.
/*

C function : atk_table_cell_get_table
*/
func (recv *TableCell) GetTable() *Object {
	retC := C.atk_table_cell_get_table((*C.AtkTableCell)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'value-changed' for Value : unsupported parameter value : type gdouble :

// Gets the minimum increment by which the value of this object may be
// changed.  If zero, the minimum increment is undefined, which may
// mean that it is limited only by the floating point precision of the
// platform.
/*

C function : atk_value_get_increment
*/
func (recv *Value) GetIncrement() float64 {
	retC := C.atk_value_get_increment((*C.AtkValue)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Gets the range of this object.
/*

C function : atk_value_get_range
*/
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

// Gets the list of subranges defined for this object. See #AtkValue
// introduction for examples of subranges and when to expose them.
/*

C function : atk_value_get_sub_ranges
*/
func (recv *Value) GetSubRanges() *glib.SList {
	retC := C.atk_value_get_sub_ranges((*C.AtkValue)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the current value and the human readable text alternative of
// @obj. @text is a newly created string, that must be freed by the
// caller. Can be NULL if no descriptor is available.
/*

C function : atk_value_get_value_and_text
*/
func (recv *Value) GetValueAndText() (float64, string) {
	var c_value C.gdouble

	var c_text *C.gchar

	C.atk_value_get_value_and_text((*C.AtkValue)(recv.native), &c_value, &c_text)

	value := (float64)(c_value)

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	return value, text
}

// Sets the value of this object.
//
// This method is intended to provide a way to change the value of the
// object. In any case, it is possible that the value can't be
// modified (ie: a read-only component). If the value changes due this
// call, it is possible that the text could change, and will trigger
// an #AtkValue::value-changed signal emission.
//
// Note for implementors: the deprecated atk_value_set_current_value()
// method returned TRUE or FALSE depending if the value was assigned
// or not. In the practice several implementors were not able to
// decide it, and returned TRUE in any case. For that reason it is not
// required anymore to return if the value was properly assigned or
// not.
/*

C function : atk_value_set_value
*/
func (recv *Value) SetValue(newValue float64) {
	c_new_value := (C.gdouble)(newValue)

	C.atk_value_set_value((*C.AtkValue)(recv.native), c_new_value)

	return
}
