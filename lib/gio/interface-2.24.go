// This is a generated file - DO NOT EDIT
// +build gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

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

// Gets the display name of the application. The display name is often more
// descriptive to the user than the name itself.
/*

C function : g_app_info_get_display_name
*/
func (recv *AppInfo) GetDisplayName() string {
	retC := C.g_app_info_get_display_name((*C.GAppInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Converter is a wrapper around the C record GConverter.
type Converter struct {
	native *C.GConverter
}

func ConverterNewFromC(u unsafe.Pointer) *Converter {
	c := (*C.GConverter)(u)
	if c == nil {
		return nil
	}

	g := &Converter{native: c}

	return g
}

func (recv *Converter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// This is the main operation used when converting data. It is to be called
// multiple times in a loop, and each time it will do some work, i.e.
// producing some output (in @outbuf) or consuming some input (from @inbuf) or
// both. If its not possible to do any work an error is returned.
//
// Note that a single call may not consume all input (or any input at all).
// Also a call may produce output even if given no input, due to state stored
// in the converter producing output.
//
// If any data was either produced or consumed, and then an error happens, then
// only the successful conversion is reported and the error is returned on the
// next call.
//
// A full conversion loop involves calling this method repeatedly, each time
// giving it new input and space output space. When there is no more input
// data after the data in @inbuf, the flag %G_CONVERTER_INPUT_AT_END must be set.
// The loop will be (unless some error happens) returning %G_CONVERTER_CONVERTED
// each time until all data is consumed and all output is produced, then
// %G_CONVERTER_FINISHED is returned instead. Note, that %G_CONVERTER_FINISHED
// may be returned even if %G_CONVERTER_INPUT_AT_END is not set, for instance
// in a decompression converter where the end of data is detectable from the
// data (and there might even be other data after the end of the compressed data).
//
// When some data has successfully been converted @bytes_read and is set to
// the number of bytes read from @inbuf, and @bytes_written is set to indicate
// how many bytes was written to @outbuf. If there are more data to output
// or consume (i.e. unless the %G_CONVERTER_INPUT_AT_END is specified) then
// %G_CONVERTER_CONVERTED is returned, and if no more data is to be output
// then %G_CONVERTER_FINISHED is returned.
//
// On error %G_CONVERTER_ERROR is returned and @error is set accordingly.
// Some errors need special handling:
//
// %G_IO_ERROR_NO_SPACE is returned if there is not enough space
// to write the resulting converted data, the application should
// call the function again with a larger @outbuf to continue.
//
// %G_IO_ERROR_PARTIAL_INPUT is returned if there is not enough
// input to fully determine what the conversion should produce,
// and the %G_CONVERTER_INPUT_AT_END flag is not set. This happens for
// example with an incomplete multibyte sequence when converting text,
// or when a regexp matches up to the end of the input (and may match
// further input). It may also happen when @inbuf_size is zero and
// there is no more data to produce.
//
// When this happens the application should read more input and then
// call the function again. If further input shows that there is no
// more data call the function again with the same data but with
// the %G_CONVERTER_INPUT_AT_END flag set. This may cause the conversion
// to finish as e.g. in the regexp match case (or, to fail again with
// %G_IO_ERROR_PARTIAL_INPUT in e.g. a charset conversion where the
// input is actually partial).
//
// After g_converter_convert() has returned %G_CONVERTER_FINISHED the
// converter object is in an invalid state where its not allowed
// to call g_converter_convert() anymore. At this time you can only
// free the object or call g_converter_reset() to reset it to the
// initial state.
//
// If the flag %G_CONVERTER_FLUSH is set then conversion is modified
// to try to write out all internal state to the output. The application
// has to call the function multiple times with the flag set, and when
// the available input has been consumed and all internal state has
// been produced then %G_CONVERTER_FLUSHED (or %G_CONVERTER_FINISHED if
// really at the end) is returned instead of %G_CONVERTER_CONVERTED.
// This is somewhat similar to what happens at the end of the input stream,
// but done in the middle of the data.
//
// This has different meanings for different conversions. For instance
// in a compression converter it would mean that we flush all the
// compression state into output such that if you uncompress the
// compressed data you get back all the input data. Doing this may
// make the final file larger due to padding though. Another example
// is a regexp conversion, where if you at the end of the flushed data
// have a match, but there is also a potential longer match. In the
// non-flushed case we would ask for more input, but when flushing we
// treat this as the end of input and do the match.
//
// Flushing is not always possible (like if a charset converter flushes
// at a partial multibyte sequence). Converters are supposed to try
// to produce as much output as possible and then return an error
// (typically %G_IO_ERROR_PARTIAL_INPUT).
/*

C function : g_converter_convert
*/
func (recv *Converter) Convert(inbuf []uint8, outbuf []uint8, flags ConverterFlags) (ConverterResult, uint64, uint64, error) {
	c_inbuf := &inbuf[0]

	c_inbuf_size := (C.gsize)(len(inbuf))

	c_outbuf := &outbuf[0]

	c_outbuf_size := (C.gsize)(len(outbuf))

	c_flags := (C.GConverterFlags)(flags)

	var c_bytes_read C.gsize

	var c_bytes_written C.gsize

	var cThrowableError *C.GError

	retC := C.g_converter_convert((*C.GConverter)(recv.native), (unsafe.Pointer(c_inbuf)), c_inbuf_size, (unsafe.Pointer(c_outbuf)), c_outbuf_size, c_flags, &c_bytes_read, &c_bytes_written, &cThrowableError)
	retGo := (ConverterResult)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	bytesRead := (uint64)(c_bytes_read)

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesRead, bytesWritten, goThrowableError
}

// Resets all internal state in the converter, making it behave
// as if it was just created. If the converter has any internal
// state that would produce output then that output is lost.
/*

C function : g_converter_reset
*/
func (recv *Converter) Reset() {
	C.g_converter_reset((*C.GConverter)(recv.native))

	return
}

// Checks if @file has a parent, and optionally, if it is @parent.
//
// If @parent is %NULL then this function returns %TRUE if @file has any
// parent at all.  If @parent is non-%NULL then %TRUE is only returned
// if @file is an immediate child of @parent.
/*

C function : g_file_has_parent
*/
func (recv *File) HasParent(parent *File) bool {
	c_parent := (*C.GFile)(parent.ToC())

	retC := C.g_file_has_parent((*C.GFile)(recv.native), c_parent)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the underlying file descriptor.
/*

C function : g_file_descriptor_based_get_fd
*/
func (recv *FileDescriptorBased) GetFd() int32 {
	retC := C.g_file_descriptor_based_get_fd((*C.GFileDescriptorBased)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
