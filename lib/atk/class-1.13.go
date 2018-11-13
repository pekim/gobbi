// This is a generated file - DO NOT EDIT
// +build atk_1.13 atk_1.20 atk_1.30 atk_2.8 atk_2.12

package atk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// Take the thread mutex for the GUI toolkit,
// if one exists.
// (This method is implemented by the toolkit ATK implementation layer;
// for instance, for GTK+, GAIL implements this via GDK_THREADS_ENTER).
/*

C function

atk_misc_threads_enter
*/
func (recv *Misc) ThreadsEnter() {
	C.atk_misc_threads_enter((*C.AtkMisc)(recv.native))

	return
}

// Release the thread mutex for the GUI toolkit,
// if one exists. This method, and atk_misc_threads_enter,
// are needed in some situations by threaded application code which
// services ATK requests, since fulfilling ATK requests often
// requires calling into the GUI toolkit.  If a long-running or
// potentially blocking call takes place inside such a block, it should
// be bracketed by atk_misc_threads_leave/atk_misc_threads_enter calls.
// (This method is implemented by the toolkit ATK implementation layer;
// for instance, for GTK+, GAIL implements this via GDK_THREADS_LEAVE).
/*

C function

atk_misc_threads_leave
*/
func (recv *Misc) ThreadsLeave() {
	C.atk_misc_threads_leave((*C.AtkMisc)(recv.native))

	return
}
