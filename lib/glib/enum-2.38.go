// This is a generated file - DO NOT EDIT
// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// The type of file to return the filename for, when used with
// g_test_build_filename().
//
// These two options correspond rather directly to the 'dist' and
// 'built' terminology that automake uses and are explicitly used to
// distinguish between the 'srcdir' and 'builddir' being separate.  All
// files in your project should either be dist (in the
// `EXTRA_DIST` or `dist_schema_DATA`
// sense, in which case they will always be in the srcdir) or built (in
// the `BUILT_SOURCES` sense, in which case they will
// always be in the builddir).
//
// Note: as a general rule of automake, files that are generated only as
// part of the build-from-git process (but then are distributed with the
// tarball) always go in srcdir (even if doing a srcdir != builddir
// build from git) and are considered as distributed files.
type TestFileType C.GTestFileType

const (
	// a file that was included in the distribution tarball
	TEST_DIST TestFileType = 0

	// a file that was built on the compiling machine
	TEST_BUILT TestFileType = 1
)
