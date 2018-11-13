// This is a generated file - DO NOT EDIT
// +build gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Moves @iter to the start of the previous visible line. Returns %TRUE if
// @iter could be moved; i.e. if @iter was at character offset 0, this
// function returns %FALSE. Therefore if @iter was already on line 0,
// but not at the start of the line, @iter is snapped to the start of
// the line and the function returns %TRUE. (Note that this implies that
// in a loop calling this function, the line number may not change on
// every iteration, if your first iteration is on line 0.)
/*

C function

gtk_text_iter_backward_visible_line
*/
func (recv *TextIter) BackwardVisibleLine() bool {
	retC := C.gtk_text_iter_backward_visible_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves @count visible lines backward, if possible (if @count would move
// past the start or end of the buffer, moves to the start or end of
// the buffer).  The return value indicates whether the iterator moved
// onto a dereferenceable position; if the iterator didn’t move, or
// moved onto the end iterator, then %FALSE is returned. If @count is 0,
// the function does nothing and returns %FALSE. If @count is negative,
// moves forward by 0 - @count lines.
/*

C function

gtk_text_iter_backward_visible_lines
*/
func (recv *TextIter) BackwardVisibleLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_visible_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Moves @iter to the start of the next visible line. Returns %TRUE if there
// was a next line to move to, and %FALSE if @iter was simply moved to
// the end of the buffer and is now not dereferenceable, or if @iter was
// already at the end of the buffer.
/*

C function

gtk_text_iter_forward_visible_line
*/
func (recv *TextIter) ForwardVisibleLine() bool {
	retC := C.gtk_text_iter_forward_visible_line((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Moves @count visible lines forward, if possible (if @count would move
// past the start or end of the buffer, moves to the start or end of
// the buffer).  The return value indicates whether the iterator moved
// onto a dereferenceable position; if the iterator didn’t move, or
// moved onto the end iterator, then %FALSE is returned. If @count is 0,
// the function does nothing and returns %FALSE. If @count is negative,
// moves backward by 0 - @count lines.
/*

C function

gtk_text_iter_forward_visible_lines
*/
func (recv *TextIter) ForwardVisibleLines(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_visible_lines((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// Returns the model that the row reference is monitoring.
/*

C function

gtk_tree_row_reference_get_model
*/
func (recv *TreeRowReference) GetModel() *TreeModel {
	retC := C.gtk_tree_row_reference_get_model((*C.GtkTreeRowReference)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}
