// +build glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

type TestFileType C.GTestFileType

const (
	TEST_DIST  TestFileType = 0
	TEST_BUILT TestFileType = 1
)