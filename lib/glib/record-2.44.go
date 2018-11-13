// This is a generated file - DO NOT EDIT
// +build glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Returns whether strict POSIX code is enabled.
//
// See g_option_context_set_strict_posix() for more information.
/*

C function : g_option_context_get_strict_posix
*/
func (recv *OptionContext) GetStrictPosix() bool {
	retC := C.g_option_context_get_strict_posix((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets strict POSIX mode.
//
// By default, this mode is disabled.
//
// In strict POSIX mode, the first non-argument parameter encountered
// (eg: filename) terminates argument processing.  Remaining arguments
// are treated as non-options and are not attempted to be parsed.
//
// If strict POSIX mode is disabled then parsing is done in the GNU way
// where option arguments can be freely mixed with non-options.
//
// As an example, consider "ls foo -l".  With GNU style parsing, this
// will list "foo" in long mode.  In strict POSIX style, this will list
// the files named "foo" and "-l".
//
// It may be useful to force strict POSIX mode when creating "verb
// style" command line tools.  For example, the "gsettings" command line
// tool supports the global option "--schemadir" as well as many
// subcommands ("get", "set", etc.) which each have their own set of
// arguments.  Using strict POSIX mode will allow parsing the global
// options up to the verb name while leaving the remaining options to be
// parsed by the relevant subcommand (which can be determined by
// examining the verb name, which should be present in argv[1] after
// parsing).
/*

C function : g_option_context_set_strict_posix
*/
func (recv *OptionContext) SetStrictPosix(strictPosix bool) {
	c_strict_posix :=
		boolToGboolean(strictPosix)

	C.g_option_context_set_strict_posix((*C.GOptionContext)(recv.native), c_strict_posix)

	return
}

// Increments the reference count of @group by one.
/*

C function : g_option_group_ref
*/
func (recv *OptionGroup) Ref() *OptionGroup {
	retC := C.g_option_group_ref((*C.GOptionGroup)(recv.native))
	retGo := OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count of @group by one.
// If the reference count drops to 0, the @group will be freed.
// and all memory allocated by the @group is released.
/*

C function : g_option_group_unref
*/
func (recv *OptionGroup) Unref() {
	C.g_option_group_unref((*C.GOptionGroup)(recv.native))

	return
}
