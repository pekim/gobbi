// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Flags that can be used with g_file_measure_disk_usage().
type FileMeasureFlags C.GFileMeasureFlags

const (
	// No flags set.
	FILE_MEASURE_NONE FileMeasureFlags = 0
	/*
	   Report any error encountered
	     while traversing the directory tree.  Normally errors are only
	     reported for the toplevel file.
	*/
	FILE_MEASURE_REPORT_ANY_ERROR FileMeasureFlags = 2
	/*
	   Tally usage based on apparent file
	     sizes.  Normally, the block-size is used, if available, as this is a
	     more accurate representation of disk space used.
	     Compare with `du --apparent-size`.
	*/
	FILE_MEASURE_APPARENT_SIZE FileMeasureFlags = 4
	/*
	   Do not cross mount point boundaries.
	     Compare with `du -x`.
	*/
	FILE_MEASURE_NO_XDEV FileMeasureFlags = 8
)
