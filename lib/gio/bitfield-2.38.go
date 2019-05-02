// This is a generated file - DO NOT EDIT
// +build gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "C"

type FileMeasureFlags C.GFileMeasureFlags

const (
	FILE_MEASURE_NONE             FileMeasureFlags = 0
	FILE_MEASURE_REPORT_ANY_ERROR FileMeasureFlags = 2
	FILE_MEASURE_APPARENT_SIZE    FileMeasureFlags = 4
	FILE_MEASURE_NO_XDEV          FileMeasureFlags = 8
)
