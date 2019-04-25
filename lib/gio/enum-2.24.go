// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

type ConverterResult int

const (
	CONVERTER_ERROR     ConverterResult = 0
	CONVERTER_CONVERTED ConverterResult = 1
	CONVERTER_FINISHED  ConverterResult = 2
	CONVERTER_FLUSHED   ConverterResult = 3
)

type ZlibCompressorFormat int

const (
	ZLIB_COMPRESSOR_FORMAT_ZLIB ZlibCompressorFormat = 0
	ZLIB_COMPRESSOR_FORMAT_GZIP ZlibCompressorFormat = 1
	ZLIB_COMPRESSOR_FORMAT_RAW  ZlibCompressorFormat = 2
)
