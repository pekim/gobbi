// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

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

// Results returned from g_converter_convert().
type ConverterResult C.GConverterResult

const (
	// There was an error during conversion.
	CONVERTER_ERROR ConverterResult = 0
	// Some data was consumed or produced
	CONVERTER_CONVERTED ConverterResult = 1
	// The conversion is finished
	CONVERTER_FINISHED ConverterResult = 2
	// Flushing is finished
	CONVERTER_FLUSHED ConverterResult = 3
)

/*
Used to select the type of data format to use for #GZlibDecompressor
and #GZlibCompressor.
*/
type ZlibCompressorFormat C.GZlibCompressorFormat

const (
	// deflate compression with zlib header
	ZLIB_COMPRESSOR_FORMAT_ZLIB ZlibCompressorFormat = 0
	// gzip file format
	ZLIB_COMPRESSOR_FORMAT_GZIP ZlibCompressorFormat = 1
	// deflate compression with no header
	ZLIB_COMPRESSOR_FORMAT_RAW ZlibCompressorFormat = 2
)
