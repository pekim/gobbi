// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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
/*

	gboolean threadedsocketservice_runHandler(GObject *, GSocketConnection *, GObject *, gpointer);

	static gulong ThreadedSocketService_signal_connect_run(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "run", G_CALLBACK(threadedsocketservice_runHandler), data);
	}

*/
/*

	void volumemonitor_driveStopButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_stop_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-stop-button", G_CALLBACK(volumemonitor_driveStopButtonHandler), data);
	}

*/
import "C"

// Unsupported : g_cancellable_connect : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Disconnects a handler from a cancellable instance similar to
// g_signal_handler_disconnect().  Additionally, in the event that a
// signal handler is currently running, this call will block until the
// handler has finished.  Calling this function from a
// #GCancellable::cancelled signal handler will therefore result in a
// deadlock.
//
// This avoids a race condition where a thread cancels at the
// same time as the cancellable operation is finished and the
// signal handler is removed. See #GCancellable::cancelled for
// details on how to use this.
//
// If @cancellable is %NULL or @handler_id is `0` this function does
// nothing.
/*

C function : g_cancellable_disconnect
*/
func (recv *Cancellable) Disconnect(handlerId uint64) {
	c_handler_id := (C.gulong)(handlerId)

	C.g_cancellable_disconnect((*C.GCancellable)(recv.native), c_handler_id)

	return
}

// Creates a #GPollFD corresponding to @cancellable; this can be passed
// to g_poll() and used to poll for cancellation. This is useful both
// for unix systems without a native poll and for portability to
// windows.
//
// When this function returns %TRUE, you should use
// g_cancellable_release_fd() to free up resources allocated for the
// @pollfd. After a %FALSE return, do not call g_cancellable_release_fd().
//
// If this function returns %FALSE, either no @cancellable was given or
// resource limits prevent this function from allocating the necessary
// structures for polling. (On Linux, you will likely have reached
// the maximum number of file descriptors.) The suggested way to handle
// these cases is to ignore the @cancellable.
//
// You are not supposed to read from the fd yourself, just check for
// readable status. Reading to unset the readable status is done
// with g_cancellable_reset().
/*

C function : g_cancellable_make_pollfd
*/
func (recv *Cancellable) MakePollfd(pollfd *glib.PollFD) bool {
	c_pollfd := (*C.GPollFD)(C.NULL)
	if pollfd != nil {
		c_pollfd = (*C.GPollFD)(pollfd.ToC())
	}

	retC := C.g_cancellable_make_pollfd((*C.GCancellable)(recv.native), c_pollfd)
	retGo := retC == C.TRUE

	return retGo
}

// Releases a resources previously allocated by g_cancellable_get_fd()
// or g_cancellable_make_pollfd().
//
// For compatibility reasons with older releases, calling this function
// is not strictly required, the resources will be automatically freed
// when the @cancellable is finalized. However, the @cancellable will
// block scarce file descriptors until it is finalized if this function
// is not called. This can cause the application to run out of file
// descriptors when many #GCancellables are used at the same time.
/*

C function : g_cancellable_release_fd
*/
func (recv *Cancellable) ReleaseFd() {
	C.g_cancellable_release_fd((*C.GCancellable)(recv.native))

	return
}

// Gets the entity tag for the file when it has been written.
// This must be called after the stream has been written
// and closed, as the etag can change while writing.
/*

C function : g_file_io_stream_get_etag
*/
func (recv *FileIOStream) GetEtag() string {
	retC := C.g_file_io_stream_get_etag((*C.GFileIOStream)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Queries a file io stream for the given @attributes.
// This function blocks while querying the stream. For the asynchronous
// version of this function, see g_file_io_stream_query_info_async().
// While the stream is blocked, the stream will set the pending flag
// internally, and any other operations on the stream will fail with
// %G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with @error being set to
// %G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
// set to %G_IO_ERROR_PENDING), or if querying info is not supported for
// the stream's interface (with @error being set to %G_IO_ERROR_NOT_SUPPORTED). I
// all cases of failure, %NULL will be returned.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be set, and %NULL will
// be returned.
/*

C function : g_file_io_stream_query_info
*/
func (recv *FileIOStream) QueryInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_io_stream_query_info((*C.GFileIOStream)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_io_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finalizes the asynchronous query started
// by g_file_io_stream_query_info_async().
/*

C function : g_file_io_stream_query_info_finish
*/
func (recv *FileIOStream) QueryInfoFinish(result *AsyncResult) (*FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_io_stream_query_info_finish((*C.GFileIOStream)(recv.native), c_result, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_info_get_attribute_stringv : no return type

// Checks if a file info structure has an attribute in the
// specified @name_space.
/*

C function : g_file_info_has_namespace
*/
func (recv *FileInfo) HasNamespace(nameSpace string) bool {
	c_name_space := C.CString(nameSpace)
	defer C.free(unsafe.Pointer(c_name_space))

	retC := C.g_file_info_has_namespace((*C.GFileInfo)(recv.native), c_name_space)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the attribute status for an attribute key. This is only
// needed by external code that implement g_file_set_attributes_from_info()
// or similar functions.
//
// The attribute must exist in @info for this to work. Otherwise %FALSE
// is returned and @info is unchanged.
/*

C function : g_file_info_set_attribute_status
*/
func (recv *FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_status := (C.GFileAttributeStatus)(status)

	retC := C.g_file_info_set_attribute_status((*C.GFileInfo)(recv.native), c_attribute, c_status)
	retGo := retC == C.TRUE

	return retGo
}

// Clears the pending flag on @stream.
/*

C function : g_io_stream_clear_pending
*/
func (recv *IOStream) ClearPending() {
	C.g_io_stream_clear_pending((*C.GIOStream)(recv.native))

	return
}

// Closes the stream, releasing resources related to it. This will also
// close the individual input and output streams, if they are not already
// closed.
//
// Once the stream is closed, all other operations will return
// %G_IO_ERROR_CLOSED. Closing a stream multiple times will not
// return an error.
//
// Closing a stream will automatically flush any outstanding buffers
// in the stream.
//
// Streams will be automatically closed when the last reference
// is dropped, but you might want to call this function to make sure
// resources are released as early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file
// descriptor) open after the stream is closed. See the documentation for
// the individual stream for details.
//
// On failure the first error that happened will be reported, but the
// close operation will finish as much as possible. A stream that failed
// to close will still return %G_IO_ERROR_CLOSED for all operations.
// Still, it is important to check and report the error to the user,
// otherwise there might be a loss of data as all data might not be written.
//
// If @cancellable is not NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
// Cancelling a close will still leave the stream closed, but some streams
// can use a faster close that doesn't block to e.g. check errors.
//
// The default implementation of this method just calls close on the
// individual input/output streams.
/*

C function : g_io_stream_close
*/
func (recv *IOStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_io_stream_close((*C.GIOStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_io_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Closes a stream.
/*

C function : g_io_stream_close_finish
*/
func (recv *IOStream) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_io_stream_close_finish((*C.GIOStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the input stream for this object. This is used
// for reading.
/*

C function : g_io_stream_get_input_stream
*/
func (recv *IOStream) GetInputStream() *InputStream {
	retC := C.g_io_stream_get_input_stream((*C.GIOStream)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the output stream for this object. This is used for
// writing.
/*

C function : g_io_stream_get_output_stream
*/
func (recv *IOStream) GetOutputStream() *OutputStream {
	retC := C.g_io_stream_get_output_stream((*C.GIOStream)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if a stream has pending actions.
/*

C function : g_io_stream_has_pending
*/
func (recv *IOStream) HasPending() bool {
	retC := C.g_io_stream_has_pending((*C.GIOStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if a stream is closed.
/*

C function : g_io_stream_is_closed
*/
func (recv *IOStream) IsClosed() bool {
	retC := C.g_io_stream_is_closed((*C.GIOStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return %FALSE and set
// @error.
/*

C function : g_io_stream_set_pending
*/
func (recv *IOStream) SetPending() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_io_stream_set_pending((*C.GIOStream)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a #GInetAddress for the "any" address (unassigned/"don't
// care") for @family.
/*

C function : g_inet_address_new_any
*/
func InetAddressNewAny(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_any(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GInetAddress from the given @family and @bytes.
// @bytes should be 4 bytes for %G_SOCKET_FAMILY_IPV4 and 16 bytes for
// %G_SOCKET_FAMILY_IPV6.
/*

C function : g_inet_address_new_from_bytes
*/
func InetAddressNewFromBytes(bytes []uint8, family SocketFamily) *InetAddress {
	c_bytes := &bytes[0]

	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_from_bytes((*C.guint8)(unsafe.Pointer(c_bytes)), c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parses @string as an IP address and creates a new #GInetAddress.
/*

C function : g_inet_address_new_from_string
*/
func InetAddressNewFromString(string string) *InetAddress {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_inet_address_new_from_string(c_string)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GInetAddress for the loopback address for @family.
/*

C function : g_inet_address_new_loopback
*/
func InetAddressNewLoopback(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_loopback(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @address's family
/*

C function : g_inet_address_get_family
*/
func (recv *InetAddress) GetFamily() SocketFamily {
	retC := C.g_inet_address_get_family((*C.GInetAddress)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// Tests whether @address is the "any" address for its family.
/*

C function : g_inet_address_get_is_any
*/
func (recv *InetAddress) GetIsAny() bool {
	retC := C.g_inet_address_get_is_any((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a link-local address (that is, if it
// identifies a host on a local network that is not connected to the
// Internet).
/*

C function : g_inet_address_get_is_link_local
*/
func (recv *InetAddress) GetIsLinkLocal() bool {
	retC := C.g_inet_address_get_is_link_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is the loopback address for its family.
/*

C function : g_inet_address_get_is_loopback
*/
func (recv *InetAddress) GetIsLoopback() bool {
	retC := C.g_inet_address_get_is_loopback((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a global multicast address.
/*

C function : g_inet_address_get_is_mc_global
*/
func (recv *InetAddress) GetIsMcGlobal() bool {
	retC := C.g_inet_address_get_is_mc_global((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a link-local multicast address.
/*

C function : g_inet_address_get_is_mc_link_local
*/
func (recv *InetAddress) GetIsMcLinkLocal() bool {
	retC := C.g_inet_address_get_is_mc_link_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a node-local multicast address.
/*

C function : g_inet_address_get_is_mc_node_local
*/
func (recv *InetAddress) GetIsMcNodeLocal() bool {
	retC := C.g_inet_address_get_is_mc_node_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is an organization-local multicast address.
/*

C function : g_inet_address_get_is_mc_org_local
*/
func (recv *InetAddress) GetIsMcOrgLocal() bool {
	retC := C.g_inet_address_get_is_mc_org_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a site-local multicast address.
/*

C function : g_inet_address_get_is_mc_site_local
*/
func (recv *InetAddress) GetIsMcSiteLocal() bool {
	retC := C.g_inet_address_get_is_mc_site_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a multicast address.
/*

C function : g_inet_address_get_is_multicast
*/
func (recv *InetAddress) GetIsMulticast() bool {
	retC := C.g_inet_address_get_is_multicast((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Tests whether @address is a site-local address such as 10.0.0.1
// (that is, the address identifies a host on a local network that can
// not be reached directly from the Internet, but which may have
// outgoing Internet connectivity via a NAT or firewall).
/*

C function : g_inet_address_get_is_site_local
*/
func (recv *InetAddress) GetIsSiteLocal() bool {
	retC := C.g_inet_address_get_is_site_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the size of the native raw binary address for @address. This
// is the size of the data that you get from g_inet_address_to_bytes().
/*

C function : g_inet_address_get_native_size
*/
func (recv *InetAddress) GetNativeSize() uint64 {
	retC := C.g_inet_address_get_native_size((*C.GInetAddress)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Blacklisted : g_inet_address_to_bytes

// Converts @address to string form.
/*

C function : g_inet_address_to_string
*/
func (recv *InetAddress) ToString() string {
	retC := C.g_inet_address_to_string((*C.GInetAddress)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GInetSocketAddress for @address and @port.
/*

C function : g_inet_socket_address_new
*/
func InetSocketAddressNew(address *InetAddress, port uint16) *InetSocketAddress {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	c_port := (C.guint16)(port)

	retC := C.g_inet_socket_address_new(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @address's #GInetAddress.
/*

C function : g_inet_socket_address_get_address
*/
func (recv *InetSocketAddress) GetAddress() *InetAddress {
	retC := C.g_inet_socket_address_get_address((*C.GInetSocketAddress)(recv.native))
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @address's port.
/*

C function : g_inet_socket_address_get_port
*/
func (recv *InetSocketAddress) GetPort() uint16 {
	retC := C.g_inet_socket_address_get_port((*C.GInetSocketAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported signal 'show-processes' for MountOperation : unsupported parameter message : type utf8 :

// Creates a new #GSocketConnectable for connecting to the given
// @hostname and @port.
//
// Note that depending on the configuration of the machine, a
// @hostname of `localhost` may refer to the IPv4 loopback address
// only, or to both IPv4 and IPv6; use
// g_network_address_new_loopback() to create a #GNetworkAddress that
// is guaranteed to resolve to both addresses.
/*

C function : g_network_address_new
*/
func NetworkAddressNew(hostname string, port uint16) *NetworkAddress {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	retC := C.g_network_address_new(c_hostname, c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets @addr's hostname. This might be either UTF-8 or ASCII-encoded,
// depending on what @addr was created with.
/*

C function : g_network_address_get_hostname
*/
func (recv *NetworkAddress) GetHostname() string {
	retC := C.g_network_address_get_hostname((*C.GNetworkAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @addr's port number
/*

C function : g_network_address_get_port
*/
func (recv *NetworkAddress) GetPort() uint16 {
	retC := C.g_network_address_get_port((*C.GNetworkAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Creates a new #GNetworkService representing the given @service,
// @protocol, and @domain. This will initially be unresolved; use the
// #GSocketConnectable interface to resolve it.
/*

C function : g_network_service_new
*/
func NetworkServiceNew(service string, protocol string, domain string) *NetworkService {
	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	retC := C.g_network_service_new(c_service, c_protocol, c_domain)
	retGo := NetworkServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the domain that @srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what @srv was created with.
/*

C function : g_network_service_get_domain
*/
func (recv *NetworkService) GetDomain() string {
	retC := C.g_network_service_get_domain((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @srv's protocol name (eg, "tcp").
/*

C function : g_network_service_get_protocol
*/
func (recv *NetworkService) GetProtocol() string {
	retC := C.g_network_service_get_protocol((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets @srv's service name (eg, "ldap").
/*

C function : g_network_service_get_service
*/
func (recv *NetworkService) GetService() string {
	retC := C.g_network_service_get_service((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Synchronously reverse-resolves @address to determine its
// associated hostname.
//
// If the DNS resolution fails, @error (if non-%NULL) will be set to
// a value from #GResolverError.
//
// If @cancellable is non-%NULL, it can be used to cancel the
// operation, in which case @error (if non-%NULL) will be set to
// %G_IO_ERROR_CANCELLED.
/*

C function : g_resolver_lookup_by_address
*/
func (recv *Resolver) LookupByAddress(address *InetAddress, cancellable *Cancellable) (string, error) {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_address((*C.GResolver)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Retrieves the result of a previous call to
// g_resolver_lookup_by_address_async().
//
// If the DNS resolution failed, @error (if non-%NULL) will be set to
// a value from #GResolverError. If the operation was cancelled,
// @error will be set to %G_IO_ERROR_CANCELLED.
/*

C function : g_resolver_lookup_by_address_finish
*/
func (recv *Resolver) LookupByAddressFinish(result *AsyncResult) (string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_address_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously resolves @hostname to determine its associated IP
// address(es). @hostname may be an ASCII-only or UTF-8 hostname, or
// the textual form of an IP address (in which case this just becomes
// a wrapper around g_inet_address_new_from_string()).
//
// On success, g_resolver_lookup_by_name() will return a non-empty #GList of
// #GInetAddress, sorted in order of preference and guaranteed to not
// contain duplicates. That is, if using the result to connect to
// @hostname, you should attempt to connect to the first address
// first, then the second if the first fails, etc. If you are using
// the result to listen on a socket, it is appropriate to add each
// result using e.g. g_socket_listener_add_address().
//
// If the DNS resolution fails, @error (if non-%NULL) will be set to a
// value from #GResolverError and %NULL will be returned.
//
// If @cancellable is non-%NULL, it can be used to cancel the
// operation, in which case @error (if non-%NULL) will be set to
// %G_IO_ERROR_CANCELLED.
//
// If you are planning to connect to a socket on the resolved IP
// address, it may be easier to create a #GNetworkAddress and use its
// #GSocketConnectable interface.
/*

C function : g_resolver_lookup_by_name
*/
func (recv *Resolver) LookupByName(hostname string, cancellable *Cancellable) (*glib.List, error) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_name((*C.GResolver)(recv.native), c_hostname, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Retrieves the result of a call to
// g_resolver_lookup_by_name_async().
//
// If the DNS resolution failed, @error (if non-%NULL) will be set to
// a value from #GResolverError. If the operation was cancelled,
// @error will be set to %G_IO_ERROR_CANCELLED.
/*

C function : g_resolver_lookup_by_name_finish
*/
func (recv *Resolver) LookupByNameFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_name_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Synchronously performs a DNS SRV lookup for the given @service and
// @protocol in the given @domain and returns an array of #GSrvTarget.
// @domain may be an ASCII-only or UTF-8 hostname. Note also that the
// @service and @protocol arguments do not include the leading underscore
// that appears in the actual DNS entry.
//
// On success, g_resolver_lookup_service() will return a non-empty #GList of
// #GSrvTarget, sorted in order of preference. (That is, you should
// attempt to connect to the first target first, then the second if
// the first fails, etc.)
//
// If the DNS resolution fails, @error (if non-%NULL) will be set to
// a value from #GResolverError and %NULL will be returned.
//
// If @cancellable is non-%NULL, it can be used to cancel the
// operation, in which case @error (if non-%NULL) will be set to
// %G_IO_ERROR_CANCELLED.
//
// If you are planning to connect to the service, it is usually easier
// to create a #GNetworkService and use its #GSocketConnectable
// interface.
/*

C function : g_resolver_lookup_service
*/
func (recv *Resolver) LookupService(service string, protocol string, domain string, cancellable *Cancellable) (*glib.List, error) {
	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_service((*C.GResolver)(recv.native), c_service, c_protocol, c_domain, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Retrieves the result of a previous call to
// g_resolver_lookup_service_async().
//
// If the DNS resolution failed, @error (if non-%NULL) will be set to
// a value from #GResolverError. If the operation was cancelled,
// @error will be set to %G_IO_ERROR_CANCELLED.
/*

C function : g_resolver_lookup_service_finish
*/
func (recv *Resolver) LookupServiceFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_service_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets @resolver to be the application's default resolver (reffing
// @resolver, and unreffing the previous default resolver, if any).
// Future calls to g_resolver_get_default() will return this resolver.
//
// This can be used if an application wants to perform any sort of DNS
// caching or "pinning"; it can implement its own #GResolver that
// calls the original default resolver for DNS operations, and
// implements its own cache policies on top of that, and then set
// itself as the default resolver for all later code to use.
/*

C function : g_resolver_set_default
*/
func (recv *Resolver) SetDefault() {
	C.g_resolver_set_default((*C.GResolver)(recv.native))

	return
}

// A #GSocket is a low-level networking primitive. It is a more or less
// direct mapping of the BSD socket API in a portable GObject based API.
// It supports both the UNIX socket implementations and winsock2 on Windows.
//
// #GSocket is the platform independent base upon which the higher level
// network primitives are based. Applications are not typically meant to
// use it directly, but rather through classes like #GSocketClient,
// #GSocketService and #GSocketConnection. However there may be cases where
// direct use of #GSocket is useful.
//
// #GSocket implements the #GInitable interface, so if it is manually constructed
// by e.g. g_object_new() you must call g_initable_init() and check the
// results before using the object. This is done automatically in
// g_socket_new() and g_socket_new_from_fd(), so these functions can return
// %NULL.
//
// Sockets operate in two general modes, blocking or non-blocking. When
// in blocking mode all operations (which don’t take an explicit blocking
// parameter) block until the requested operation
// is finished or there is an error. In non-blocking mode all calls that
// would block return immediately with a %G_IO_ERROR_WOULD_BLOCK error.
// To know when a call would successfully run you can call g_socket_condition_check(),
// or g_socket_condition_wait(). You can also use g_socket_create_source() and
// attach it to a #GMainContext to get callbacks when I/O is possible.
// Note that all sockets are always set to non blocking mode in the system, and
// blocking mode is emulated in GSocket.
//
// When working in non-blocking mode applications should always be able to
// handle getting a %G_IO_ERROR_WOULD_BLOCK error even when some other
// function said that I/O was possible. This can easily happen in case
// of a race condition in the application, but it can also happen for other
// reasons. For instance, on Windows a socket is always seen as writable
// until a write returns %G_IO_ERROR_WOULD_BLOCK.
//
// #GSockets can be either connection oriented or datagram based.
// For connection oriented types you must first establish a connection by
// either connecting to an address or accepting a connection from another
// address. For connectionless socket types the target/source address is
// specified or received in each I/O operation.
//
// All socket file descriptors are set to be close-on-exec.
//
// Note that creating a #GSocket causes the signal %SIGPIPE to be
// ignored for the remainder of the program. If you are writing a
// command-line utility that uses #GSocket, you may need to take into
// account the fact that your program will not automatically be killed
// if it tries to write to %stdout after it has been closed.
//
// Like most other APIs in GLib, #GSocket is not inherently thread safe. To use
// a #GSocket concurrently from multiple threads, you must implement your own
// locking.
/*

C record/class : GSocket
*/
type Socket struct {
	native *C.GSocket
	// parent_instance : record
	// priv : record
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	c := (*C.GSocket)(u)
	if c == nil {
		return nil
	}

	g := &Socket{native: c}

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Socket.
// Exercise care, as this is a potentially dangerous function if the Object is not a Socket.
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromC(object.ToC())
}

// Creates a new #GSocket with the defined family, type and protocol.
// If @protocol is 0 (%G_SOCKET_PROTOCOL_DEFAULT) the default protocol type
// for the family and type is used.
//
// The @protocol is a family and type specific int that specifies what
// kind of protocol to use. #GSocketProtocol lists several common ones.
// Many families only support one protocol, and use 0 for this, others
// support several and using 0 means to use the default protocol for
// the family and type.
//
// The protocol id is passed directly to the operating
// system, so you can use protocols not listed in #GSocketProtocol if you
// know the protocol number used for it.
/*

C function : g_socket_new
*/
func SocketNew(family SocketFamily, type_ SocketType, protocol SocketProtocol) (*Socket, error) {
	c_family := (C.GSocketFamily)(family)

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	var cThrowableError *C.GError

	retC := C.g_socket_new(c_family, c_type, c_protocol, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new #GSocket from a native file descriptor
// or winsock SOCKET handle.
//
// This reads all the settings from the file descriptor so that
// all properties should work. Note that the file descriptor
// will be set to non-blocking mode, independent on the blocking
// mode of the #GSocket.
//
// On success, the returned #GSocket takes ownership of @fd. On failure, the
// caller must close @fd themselves.
//
// Since GLib 2.46, it is no longer a fatal error to call this on a non-socket
// descriptor.  Instead, a GError will be set with code %G_IO_ERROR_FAILED
/*

C function : g_socket_new_from_fd
*/
func SocketNewFromFd(fd int32) (*Socket, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_socket_new_from_fd(c_fd, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Accept incoming connections on a connection-based socket. This removes
// the first outstanding connection request from the listening socket and
// creates a #GSocket object for it.
//
// The @socket must be bound to a local address with g_socket_bind() and
// must be listening for incoming connections (g_socket_listen()).
//
// If there are no outstanding connections then the operation will block
// or return %G_IO_ERROR_WOULD_BLOCK if non-blocking I/O is enabled.
// To be notified of an incoming connection, wait for the %G_IO_IN condition.
/*

C function : g_socket_accept
*/
func (recv *Socket) Accept(cancellable *Cancellable) (*Socket, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_accept((*C.GSocket)(recv.native), c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// When a socket is created it is attached to an address family, but it
// doesn't have an address in this family. g_socket_bind() assigns the
// address (sometimes called name) of the socket.
//
// It is generally required to bind to a local address before you can
// receive connections. (See g_socket_listen() and g_socket_accept() ).
// In certain situations, you may also want to bind a socket that will be
// used to initiate connections, though this is not normally required.
//
// If @socket is a TCP socket, then @allow_reuse controls the setting
// of the `SO_REUSEADDR` socket option; normally it should be %TRUE for
// server sockets (sockets that you will eventually call
// g_socket_accept() on), and %FALSE for client sockets. (Failing to
// set this flag on a server socket may cause g_socket_bind() to return
// %G_IO_ERROR_ADDRESS_IN_USE if the server program is stopped and then
// immediately restarted.)
//
// If @socket is a UDP socket, then @allow_reuse determines whether or
// not other UDP sockets can be bound to the same address at the same
// time. In particular, you can have several UDP sockets bound to the
// same address, and they will all receive all of the multicast and
// broadcast packets sent to that address. (The behavior of unicast
// UDP packets to an address with multiple listeners is not defined.)
/*

C function : g_socket_bind
*/
func (recv *Socket) Bind(address *SocketAddress, allowReuse bool) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_allow_reuse :=
		boolToGboolean(allowReuse)

	var cThrowableError *C.GError

	retC := C.g_socket_bind((*C.GSocket)(recv.native), c_address, c_allow_reuse, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks and resets the pending connect error for the socket.
// This is used to check for errors when g_socket_connect() is
// used in non-blocking mode.
/*

C function : g_socket_check_connect_result
*/
func (recv *Socket) CheckConnectResult() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_check_connect_result((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Closes the socket, shutting down any active connection.
//
// Closing a socket does not wait for all outstanding I/O operations
// to finish, so the caller should not rely on them to be guaranteed
// to complete even if the close returns with no error.
//
// Once the socket is closed, all other operations will return
// %G_IO_ERROR_CLOSED. Closing a socket multiple times will not
// return an error.
//
// Sockets will be automatically closed when the last reference
// is dropped, but you might want to call this function to make sure
// resources are released as early as possible.
//
// Beware that due to the way that TCP works, it is possible for
// recently-sent data to be lost if either you close a socket while the
// %G_IO_IN condition is set, or else if the remote connection tries to
// send something to you after you close the socket but before it has
// finished reading all of the data you sent. There is no easy generic
// way to avoid this problem; the easiest fix is to design the network
// protocol such that the client will never send data "out of turn".
// Another solution is for the server to half-close the connection by
// calling g_socket_shutdown() with only the @shutdown_write flag set,
// and then wait for the client to notice this and close its side of the
// connection, after which the server can safely call g_socket_close().
// (This is what #GTcpConnection does if you call
// g_tcp_connection_set_graceful_disconnect(). But of course, this
// only works if the client will close its connection after the server
// does.)
/*

C function : g_socket_close
*/
func (recv *Socket) Close() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_close((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks on the readiness of @socket to perform operations.
// The operations specified in @condition are checked for and masked
// against the currently-satisfied conditions on @socket. The result
// is returned.
//
// Note that on Windows, it is possible for an operation to return
// %G_IO_ERROR_WOULD_BLOCK even immediately after
// g_socket_condition_check() has claimed that the socket is ready for
// writing. Rather than calling g_socket_condition_check() and then
// writing to the socket if it succeeds, it is generally better to
// simply try writing to the socket right away, and try again later if
// the initial attempt returns %G_IO_ERROR_WOULD_BLOCK.
//
// It is meaningless to specify %G_IO_ERR or %G_IO_HUP in condition;
// these conditions will always be set in the output if they are true.
//
// This call never blocks.
/*

C function : g_socket_condition_check
*/
func (recv *Socket) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_socket_condition_check((*C.GSocket)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// Waits for @condition to become true on @socket. When the condition
// is met, %TRUE is returned.
//
// If @cancellable is cancelled before the condition is met, or if the
// socket has a timeout set and it is reached before the condition is
// met, then %FALSE is returned and @error, if non-%NULL, is set to
// the appropriate value (%G_IO_ERROR_CANCELLED or
// %G_IO_ERROR_TIMED_OUT).
//
// See also g_socket_condition_timed_wait().
/*

C function : g_socket_condition_wait
*/
func (recv *Socket) ConditionWait(condition glib.IOCondition, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_condition_wait((*C.GSocket)(recv.native), c_condition, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Connect the socket to the specified remote address.
//
// For connection oriented socket this generally means we attempt to make
// a connection to the @address. For a connection-less socket it sets
// the default address for g_socket_send() and discards all incoming datagrams
// from other sources.
//
// Generally connection oriented sockets can only connect once, but
// connection-less sockets can connect multiple times to change the
// default address.
//
// If the connect call needs to do network I/O it will block, unless
// non-blocking I/O is enabled. Then %G_IO_ERROR_PENDING is returned
// and the user can be notified of the connection finishing by waiting
// for the G_IO_OUT condition. The result of the connection must then be
// checked with g_socket_check_connect_result().
/*

C function : g_socket_connect
*/
func (recv *Socket) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_connect((*C.GSocket)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a #GSocketConnection subclass of the right type for
// @socket.
/*

C function : g_socket_connection_factory_create_connection
*/
func (recv *Socket) ConnectionFactoryCreateConnection() *SocketConnection {
	retC := C.g_socket_connection_factory_create_connection((*C.GSocket)(recv.native))
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GSource that can be attached to a %GMainContext to monitor
// for the availability of the specified @condition on the socket. The #GSource
// keeps a reference to the @socket.
//
// The callback on the source is of the #GSocketSourceFunc type.
//
// It is meaningless to specify %G_IO_ERR or %G_IO_HUP in @condition;
// these conditions will always be reported output if they are true.
//
// @cancellable if not %NULL can be used to cancel the source, which will
// cause the source to trigger, reporting the current condition (which
// is likely 0 unless cancellation happened at the same time as a
// condition change). You can check for this in the callback using
// g_cancellable_is_cancelled().
//
// If @socket has a timeout set, and it is reached before @condition
// occurs, the source will then trigger anyway, reporting %G_IO_IN or
// %G_IO_OUT depending on @condition. However, @socket will have been
// marked as having had a timeout, and so the next #GSocket I/O method
// you call will then fail with a %G_IO_ERROR_TIMED_OUT.
/*

C function : g_socket_create_source
*/
func (recv *Socket) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_socket_create_source((*C.GSocket)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the blocking mode of the socket. For details on blocking I/O,
// see g_socket_set_blocking().
/*

C function : g_socket_get_blocking
*/
func (recv *Socket) GetBlocking() bool {
	retC := C.g_socket_get_blocking((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the socket family of the socket.
/*

C function : g_socket_get_family
*/
func (recv *Socket) GetFamily() SocketFamily {
	retC := C.g_socket_get_family((*C.GSocket)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// Returns the underlying OS socket object. On unix this
// is a socket file descriptor, and on Windows this is
// a Winsock2 SOCKET handle. This may be useful for
// doing platform specific or otherwise unusual operations
// on the socket.
/*

C function : g_socket_get_fd
*/
func (recv *Socket) GetFd() int32 {
	retC := C.g_socket_get_fd((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the keepalive mode of the socket. For details on this,
// see g_socket_set_keepalive().
/*

C function : g_socket_get_keepalive
*/
func (recv *Socket) GetKeepalive() bool {
	retC := C.g_socket_get_keepalive((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the listen backlog setting of the socket. For details on this,
// see g_socket_set_listen_backlog().
/*

C function : g_socket_get_listen_backlog
*/
func (recv *Socket) GetListenBacklog() int32 {
	retC := C.g_socket_get_listen_backlog((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_socket_get_local_address

// Gets the socket protocol id the socket was created with.
// In case the protocol is unknown, -1 is returned.
/*

C function : g_socket_get_protocol
*/
func (recv *Socket) GetProtocol() SocketProtocol {
	retC := C.g_socket_get_protocol((*C.GSocket)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// Try to get the remote address of a connected socket. This is only
// useful for connection oriented sockets that have been connected.
/*

C function : g_socket_get_remote_address
*/
func (recv *Socket) GetRemoteAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_remote_address((*C.GSocket)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the socket type of the socket.
/*

C function : g_socket_get_socket_type
*/
func (recv *Socket) GetSocketType() SocketType {
	retC := C.g_socket_get_socket_type((*C.GSocket)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// Checks whether a socket is closed.
/*

C function : g_socket_is_closed
*/
func (recv *Socket) IsClosed() bool {
	retC := C.g_socket_is_closed((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Check whether the socket is connected. This is only useful for
// connection-oriented sockets.
//
// If using g_socket_shutdown(), this function will return %TRUE until the
// socket has been shut down for reading and writing. If you do a non-blocking
// connect, this function will not return %TRUE until after you call
// g_socket_check_connect_result().
/*

C function : g_socket_is_connected
*/
func (recv *Socket) IsConnected() bool {
	retC := C.g_socket_is_connected((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Marks the socket as a server socket, i.e. a socket that is used
// to accept incoming requests using g_socket_accept().
//
// Before calling this the socket must be bound to a local address using
// g_socket_bind().
//
// To set the maximum amount of outstanding clients, use
// g_socket_set_listen_backlog().
/*

C function : g_socket_listen
*/
func (recv *Socket) Listen() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_listen((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Receive data (up to @size bytes) from a socket. This is mainly used by
// connection-oriented sockets; it is identical to g_socket_receive_from()
// with @address set to %NULL.
//
// For %G_SOCKET_TYPE_DATAGRAM and %G_SOCKET_TYPE_SEQPACKET sockets,
// g_socket_receive() will always read either 0 or 1 complete messages from
// the socket. If the received message is too large to fit in @buffer, then
// the data beyond @size bytes will be discarded, without any explicit
// indication that this has occurred.
//
// For %G_SOCKET_TYPE_STREAM sockets, g_socket_receive() can return any
// number of bytes, up to @size. If more than @size bytes have been
// received, the additional data will be returned in future calls to
// g_socket_receive().
//
// If the socket is in blocking mode the call will block until there
// is some data to receive, the connection is closed, or there is an
// error. If there is no data available and the socket is in
// non-blocking mode, a %G_IO_ERROR_WOULD_BLOCK error will be
// returned. To be notified when data is available, wait for the
// %G_IO_IN condition.
//
// On error -1 is returned and @error is set accordingly.
/*

C function : g_socket_receive
*/
func (recv *Socket) Receive(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_receive((*C.GSocket)(recv.native), (*C.gchar)(unsafe.Pointer(c_buffer)), c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Receive data (up to @size bytes) from a socket.
//
// If @address is non-%NULL then @address will be set equal to the
// source address of the received packet.
// @address is owned by the caller.
//
// See g_socket_receive() for additional information.
/*

C function : g_socket_receive_from
*/
func (recv *Socket) ReceiveFrom(buffer []uint8, cancellable *Cancellable) (int64, *SocketAddress, error) {
	var c_address *C.GSocketAddress

	c_buffer := &buffer[0]

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_receive_from((*C.GSocket)(recv.native), &c_address, (*C.gchar)(unsafe.Pointer(c_buffer)), c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	address := SocketAddressNewFromC(unsafe.Pointer(c_address))

	return retGo, address, goThrowableError
}

// Unsupported : g_socket_receive_message : unsupported parameter vectors :

// Tries to send @size bytes from @buffer on the socket. This is
// mainly used by connection-oriented sockets; it is identical to
// g_socket_send_to() with @address set to %NULL.
//
// If the socket is in blocking mode the call will block until there is
// space for the data in the socket queue. If there is no space available
// and the socket is in non-blocking mode a %G_IO_ERROR_WOULD_BLOCK error
// will be returned. To be notified when space is available, wait for the
// %G_IO_OUT condition. Note though that you may still receive
// %G_IO_ERROR_WOULD_BLOCK from g_socket_send() even if you were previously
// notified of a %G_IO_OUT condition. (On Windows in particular, this is
// very common due to the way the underlying APIs work.)
//
// On error -1 is returned and @error is set accordingly.
/*

C function : g_socket_send
*/
func (recv *Socket) Send(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer := &buffer[0]

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_send((*C.GSocket)(recv.native), (*C.gchar)(unsafe.Pointer(c_buffer)), c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_send_message : unsupported parameter vectors :

// Tries to send @size bytes from @buffer to @address. If @address is
// %NULL then the message is sent to the default receiver (set by
// g_socket_connect()).
//
// See g_socket_send() for additional information.
/*

C function : g_socket_send_to
*/
func (recv *Socket) SendTo(address *SocketAddress, buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_buffer := &buffer[0]

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_send_to((*C.GSocket)(recv.native), c_address, (*C.gchar)(unsafe.Pointer(c_buffer)), c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the blocking mode of the socket. In blocking mode
// all operations (which don’t take an explicit blocking parameter) block until
// they succeed or there is an error. In
// non-blocking mode all functions return results immediately or
// with a %G_IO_ERROR_WOULD_BLOCK error.
//
// All sockets are created in blocking mode. However, note that the
// platform level socket is always non-blocking, and blocking mode
// is a GSocket level feature.
/*

C function : g_socket_set_blocking
*/
func (recv *Socket) SetBlocking(blocking bool) {
	c_blocking :=
		boolToGboolean(blocking)

	C.g_socket_set_blocking((*C.GSocket)(recv.native), c_blocking)

	return
}

// Sets or unsets the %SO_KEEPALIVE flag on the underlying socket. When
// this flag is set on a socket, the system will attempt to verify that the
// remote socket endpoint is still present if a sufficiently long period of
// time passes with no data being exchanged. If the system is unable to
// verify the presence of the remote endpoint, it will automatically close
// the connection.
//
// This option is only functional on certain kinds of sockets. (Notably,
// %G_SOCKET_PROTOCOL_TCP sockets.)
//
// The exact time between pings is system- and protocol-dependent, but will
// normally be at least two hours. Most commonly, you would set this flag
// on a server socket if you want to allow clients to remain idle for long
// periods of time, but also want to ensure that connections are eventually
// garbage-collected if clients crash or become unreachable.
/*

C function : g_socket_set_keepalive
*/
func (recv *Socket) SetKeepalive(keepalive bool) {
	c_keepalive :=
		boolToGboolean(keepalive)

	C.g_socket_set_keepalive((*C.GSocket)(recv.native), c_keepalive)

	return
}

// Sets the maximum number of outstanding connections allowed
// when listening on this socket. If more clients than this are
// connecting to the socket and the application is not handling them
// on time then the new connections will be refused.
//
// Note that this must be called before g_socket_listen() and has no
// effect if called after that.
/*

C function : g_socket_set_listen_backlog
*/
func (recv *Socket) SetListenBacklog(backlog int32) {
	c_backlog := (C.gint)(backlog)

	C.g_socket_set_listen_backlog((*C.GSocket)(recv.native), c_backlog)

	return
}

// Shut down part or all of a full-duplex connection.
//
// If @shutdown_read is %TRUE then the receiving side of the connection
// is shut down, and further reading is disallowed.
//
// If @shutdown_write is %TRUE then the sending side of the connection
// is shut down, and further writing is disallowed.
//
// It is allowed for both @shutdown_read and @shutdown_write to be %TRUE.
//
// One example where it is useful to shut down only one side of a connection is
// graceful disconnect for TCP connections where you close the sending side,
// then wait for the other side to close the connection, thus ensuring that the
// other side saw all sent data.
/*

C function : g_socket_shutdown
*/
func (recv *Socket) Shutdown(shutdownRead bool, shutdownWrite bool) (bool, error) {
	c_shutdown_read :=
		boolToGboolean(shutdownRead)

	c_shutdown_write :=
		boolToGboolean(shutdownWrite)

	var cThrowableError *C.GError

	retC := C.g_socket_shutdown((*C.GSocket)(recv.native), c_shutdown_read, c_shutdown_write, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if a socket is capable of speaking IPv4.
//
// IPv4 sockets are capable of speaking IPv4.  On some operating systems
// and under some combinations of circumstances IPv6 sockets are also
// capable of speaking IPv4.  See RFC 3493 section 3.7 for more
// information.
//
// No other types of sockets are currently considered as being capable
// of speaking IPv4.
/*

C function : g_socket_speaks_ipv4
*/
func (recv *Socket) SpeaksIpv4() bool {
	retC := C.g_socket_speaks_ipv4((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Creates a #GSocketAddress subclass corresponding to the native
// struct sockaddr @native.
/*

C function : g_socket_address_new_from_native
*/
func SocketAddressNewFromNative(native uintptr, len uint64) *SocketAddress {
	c_native := (C.gpointer)(native)

	c_len := (C.gsize)(len)

	retC := C.g_socket_address_new_from_native(c_native, c_len)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the socket family type of @address.
/*

C function : g_socket_address_get_family
*/
func (recv *SocketAddress) GetFamily() SocketFamily {
	retC := C.g_socket_address_get_family((*C.GSocketAddress)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// Gets the size of @address's native struct sockaddr.
// You can use this to allocate memory to pass to
// g_socket_address_to_native().
/*

C function : g_socket_address_get_native_size
*/
func (recv *SocketAddress) GetNativeSize() int64 {
	retC := C.g_socket_address_get_native_size((*C.GSocketAddress)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Converts a #GSocketAddress to a native struct sockaddr, which can
// be passed to low-level functions like connect() or bind().
//
// If not enough space is available, a %G_IO_ERROR_NO_SPACE error
// is returned. If the address type is not known on the system
// then a %G_IO_ERROR_NOT_SUPPORTED error is returned.
/*

C function : g_socket_address_to_native
*/
func (recv *SocketAddress) ToNative(dest uintptr, destlen uint64) (bool, error) {
	c_dest := (C.gpointer)(dest)

	c_destlen := (C.gsize)(destlen)

	var cThrowableError *C.GError

	retC := C.g_socket_address_to_native((*C.GSocketAddress)(recv.native), c_dest, c_destlen, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// #GSocketClient is a lightweight high-level utility class for connecting to
// a network host using a connection oriented socket type.
//
// You create a #GSocketClient object, set any options you want, and then
// call a sync or async connect operation, which returns a #GSocketConnection
// subclass on success.
//
// The type of the #GSocketConnection object returned depends on the type of
// the underlying socket that is in use. For instance, for a TCP/IP connection
// it will be a #GTcpConnection.
//
// As #GSocketClient is a lightweight object, you don't need to cache it. You
// can just create a new one any time you need one.
/*

C record/class : GSocketClient
*/
type SocketClient struct {
	native *C.GSocketClient
	// parent_instance : record
	// priv : record
}

func SocketClientNewFromC(u unsafe.Pointer) *SocketClient {
	c := (*C.GSocketClient)(u)
	if c == nil {
		return nil
	}

	g := &SocketClient{native: c}

	return g
}

func (recv *SocketClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SocketClient) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SocketClient.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketClient.
func CastToSocketClient(object *gobject.Object) *SocketClient {
	return SocketClientNewFromC(object.ToC())
}

// Creates a new #GSocketClient with the default options.
/*

C function : g_socket_client_new
*/
func SocketClientNew() *SocketClient {
	retC := C.g_socket_client_new()
	retGo := SocketClientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Enable proxy protocols to be handled by the application. When the
// indicated proxy protocol is returned by the #GProxyResolver,
// #GSocketClient will consider this protocol as supported but will
// not try to find a #GProxy instance to handle handshaking. The
// application must check for this case by calling
// g_socket_connection_get_remote_address() on the returned
// #GSocketConnection, and seeing if it's a #GProxyAddress of the
// appropriate type, to determine whether or not it needs to handle
// the proxy handshaking itself.
//
// This should be used for proxy protocols that are dialects of
// another protocol such as HTTP proxy. It also allows cohabitation of
// proxy protocols that are reused between protocols. A good example
// is HTTP. It can be used to proxy HTTP, FTP and Gopher and can also
// be use as generic socket proxy through the HTTP CONNECT method.
//
// When the proxy is detected as being an application proxy, TLS handshake
// will be skipped. This is required to let the application do the proxy
// specific handshake.
/*

C function : g_socket_client_add_application_proxy
*/
func (recv *SocketClient) AddApplicationProxy(protocol string) {
	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	C.g_socket_client_add_application_proxy((*C.GSocketClient)(recv.native), c_protocol)

	return
}

// Tries to resolve the @connectable and make a network connection to it.
//
// Upon a successful connection, a new #GSocketConnection is constructed
// and returned.  The caller owns this new object and must drop their
// reference to it when finished with it.
//
// The type of the #GSocketConnection object returned depends on the type of
// the underlying socket that is used. For instance, for a TCP/IP connection
// it will be a #GTcpConnection.
//
// The socket created will be the same family as the address that the
// @connectable resolves to, unless family is set with g_socket_client_set_family()
// or indirectly via g_socket_client_set_local_address(). The socket type
// defaults to %G_SOCKET_TYPE_STREAM but can be set with
// g_socket_client_set_socket_type().
//
// If a local address is specified with g_socket_client_set_local_address() the
// socket will be bound to this address before connecting.
/*

C function : g_socket_client_connect
*/
func (recv *SocketClient) Connect(connectable *SocketConnectable, cancellable *Cancellable) (*SocketConnection, error) {
	c_connectable := (*C.GSocketConnectable)(connectable.ToC())

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect((*C.GSocketClient)(recv.native), c_connectable, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async connect operation. See g_socket_client_connect_async()
/*

C function : g_socket_client_connect_finish
*/
func (recv *SocketClient) ConnectFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This is a helper function for g_socket_client_connect().
//
// Attempts to create a TCP connection to the named host.
//
// @host_and_port may be in any of a number of recognized formats; an IPv6
// address, an IPv4 address, or a domain name (in which case a DNS
// lookup is performed).  Quoting with [] is supported for all address
// types.  A port override may be specified in the usual way with a
// colon.  Ports may be given as decimal numbers or symbolic names (in
// which case an /etc/services lookup is performed).
//
// If no port override is given in @host_and_port then @default_port will be
// used as the port number to connect to.
//
// In general, @host_and_port is expected to be provided by the user (allowing
// them to give the hostname, and a port override if necessary) and
// @default_port is expected to be provided by the application.
//
// In the case that an IP address is given, a single connection
// attempt is made.  In the case that a name is given, multiple
// connection attempts may be made, in turn and according to the
// number of address records in DNS, until a connection succeeds.
//
// Upon a successful connection, a new #GSocketConnection is constructed
// and returned.  The caller owns this new object and must drop their
// reference to it when finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) %NULL is returned and @error (if non-%NULL) is set
// accordingly.
/*

C function : g_socket_client_connect_to_host
*/
func (recv *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_host_and_port := C.CString(hostAndPort)
	defer C.free(unsafe.Pointer(c_host_and_port))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_host((*C.GSocketClient)(recv.native), c_host_and_port, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async connect operation. See g_socket_client_connect_to_host_async()
/*

C function : g_socket_client_connect_to_host_finish
*/
func (recv *SocketClient) ConnectToHostFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_host_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Attempts to create a TCP connection to a service.
//
// This call looks up the SRV record for @service at @domain for the
// "tcp" protocol.  It then attempts to connect, in turn, to each of
// the hosts providing the service until either a connection succeeds
// or there are no hosts remaining.
//
// Upon a successful connection, a new #GSocketConnection is constructed
// and returned.  The caller owns this new object and must drop their
// reference to it when finished with it.
//
// In the event of any failure (DNS error, service not found, no hosts
// connectable) %NULL is returned and @error (if non-%NULL) is set
// accordingly.
/*

C function : g_socket_client_connect_to_service
*/
func (recv *SocketClient) ConnectToService(domain string, service string, cancellable *Cancellable) (*SocketConnection, error) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_service((*C.GSocketClient)(recv.native), c_domain, c_service, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async connect operation. See g_socket_client_connect_to_service_async()
/*

C function : g_socket_client_connect_to_service_finish
*/
func (recv *SocketClient) ConnectToServiceFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_service_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the socket family of the socket client.
//
// See g_socket_client_set_family() for details.
/*

C function : g_socket_client_get_family
*/
func (recv *SocketClient) GetFamily() SocketFamily {
	retC := C.g_socket_client_get_family((*C.GSocketClient)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// Gets the local address of the socket client.
//
// See g_socket_client_set_local_address() for details.
/*

C function : g_socket_client_get_local_address
*/
func (recv *SocketClient) GetLocalAddress() *SocketAddress {
	retC := C.g_socket_client_get_local_address((*C.GSocketClient)(recv.native))
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the protocol name type of the socket client.
//
// See g_socket_client_set_protocol() for details.
/*

C function : g_socket_client_get_protocol
*/
func (recv *SocketClient) GetProtocol() SocketProtocol {
	retC := C.g_socket_client_get_protocol((*C.GSocketClient)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// Gets the socket type of the socket client.
//
// See g_socket_client_set_socket_type() for details.
/*

C function : g_socket_client_get_socket_type
*/
func (recv *SocketClient) GetSocketType() SocketType {
	retC := C.g_socket_client_get_socket_type((*C.GSocketClient)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// Sets the socket family of the socket client.
// If this is set to something other than %G_SOCKET_FAMILY_INVALID
// then the sockets created by this object will be of the specified
// family.
//
// This might be useful for instance if you want to force the local
// connection to be an ipv4 socket, even though the address might
// be an ipv6 mapped to ipv4 address.
/*

C function : g_socket_client_set_family
*/
func (recv *SocketClient) SetFamily(family SocketFamily) {
	c_family := (C.GSocketFamily)(family)

	C.g_socket_client_set_family((*C.GSocketClient)(recv.native), c_family)

	return
}

// Sets the local address of the socket client.
// The sockets created by this object will bound to the
// specified address (if not %NULL) before connecting.
//
// This is useful if you want to ensure that the local
// side of the connection is on a specific port, or on
// a specific interface.
/*

C function : g_socket_client_set_local_address
*/
func (recv *SocketClient) SetLocalAddress(address *SocketAddress) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	C.g_socket_client_set_local_address((*C.GSocketClient)(recv.native), c_address)

	return
}

// Sets the protocol of the socket client.
// The sockets created by this object will use of the specified
// protocol.
//
// If @protocol is %0 that means to use the default
// protocol for the socket family and type.
/*

C function : g_socket_client_set_protocol
*/
func (recv *SocketClient) SetProtocol(protocol SocketProtocol) {
	c_protocol := (C.GSocketProtocol)(protocol)

	C.g_socket_client_set_protocol((*C.GSocketClient)(recv.native), c_protocol)

	return
}

// Sets the socket type of the socket client.
// The sockets created by this object will be of the specified
// type.
//
// It doesn't make sense to specify a type of %G_SOCKET_TYPE_DATAGRAM,
// as GSocketClient is used for connection oriented services.
/*

C function : g_socket_client_set_socket_type
*/
func (recv *SocketClient) SetSocketType(type_ SocketType) {
	c_type := (C.GSocketType)(type_)

	C.g_socket_client_set_socket_type((*C.GSocketClient)(recv.native), c_type)

	return
}

// #GSocketConnection is a #GIOStream for a connected socket. They
// can be created either by #GSocketClient when connecting to a host,
// or by #GSocketListener when accepting a new client.
//
// The type of the #GSocketConnection object returned from these calls
// depends on the type of the underlying socket that is in use. For
// instance, for a TCP/IP connection it will be a #GTcpConnection.
//
// Choosing what type of object to construct is done with the socket
// connection factory, and it is possible for 3rd parties to register
// custom socket connection types for specific combination of socket
// family/type/protocol using g_socket_connection_factory_register_type().
//
// To close a #GSocketConnection, use g_io_stream_close(). Closing both
// substreams of the #GIOStream separately will not close the underlying
// #GSocket.
/*

C record/class : GSocketConnection
*/
type SocketConnection struct {
	native *C.GSocketConnection
	// parent_instance : record
	// priv : record
}

func SocketConnectionNewFromC(u unsafe.Pointer) *SocketConnection {
	c := (*C.GSocketConnection)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnection{native: c}

	return g
}

func (recv *SocketConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOStream upcasts to *IOStream
func (recv *SocketConnection) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SocketConnection) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitary Object to SocketConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketConnection.
func CastToSocketConnection(object *gobject.Object) *SocketConnection {
	return SocketConnectionNewFromC(object.ToC())
}

// Try to get the local address of a socket connection.
/*

C function : g_socket_connection_get_local_address
*/
func (recv *SocketConnection) GetLocalAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_connection_get_local_address((*C.GSocketConnection)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Try to get the remote address of a socket connection.
//
// Since GLib 2.40, when used with g_socket_client_connect() or
// g_socket_client_connect_async(), during emission of
// %G_SOCKET_CLIENT_CONNECTING, this function will return the remote
// address that will be used for the connection.  This allows
// applications to print e.g. "Connecting to example.com
// (10.42.77.3)...".
/*

C function : g_socket_connection_get_remote_address
*/
func (recv *SocketConnection) GetRemoteAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_connection_get_remote_address((*C.GSocketConnection)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Gets the underlying #GSocket object of the connection.
// This can be useful if you want to do something unusual on it
// not supported by the #GSocketConnection APIs.
/*

C function : g_socket_connection_get_socket
*/
func (recv *SocketConnection) GetSocket() *Socket {
	retC := C.g_socket_connection_get_socket((*C.GSocketConnection)(recv.native))
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the "level" (i.e. the originating protocol) of the control message.
// This is often SOL_SOCKET.
/*

C function : g_socket_control_message_get_level
*/
func (recv *SocketControlMessage) GetLevel() int32 {
	retC := C.g_socket_control_message_get_level((*C.GSocketControlMessage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the protocol specific type of the control message.
// For instance, for UNIX fd passing this would be SCM_RIGHTS.
/*

C function : g_socket_control_message_get_msg_type
*/
func (recv *SocketControlMessage) GetMsgType() int32 {
	retC := C.g_socket_control_message_get_msg_type((*C.GSocketControlMessage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the space required for the control message, not including
// headers or alignment.
/*

C function : g_socket_control_message_get_size
*/
func (recv *SocketControlMessage) GetSize() uint64 {
	retC := C.g_socket_control_message_get_size((*C.GSocketControlMessage)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Converts the data in the message to bytes placed in the
// message.
//
// @data is guaranteed to have enough space to fit the size
// returned by g_socket_control_message_get_size() on this
// object.
/*

C function : g_socket_control_message_serialize
*/
func (recv *SocketControlMessage) Serialize(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_socket_control_message_serialize((*C.GSocketControlMessage)(recv.native), c_data)

	return
}

// A #GSocketListener is an object that keeps track of a set
// of server sockets and helps you accept sockets from any of the
// socket, either sync or async.
//
// If you want to implement a network server, also look at #GSocketService
// and #GThreadedSocketService which are subclass of #GSocketListener
// that makes this even easier.
/*

C record/class : GSocketListener
*/
type SocketListener struct {
	native *C.GSocketListener
	// parent_instance : record
	// priv : record
}

func SocketListenerNewFromC(u unsafe.Pointer) *SocketListener {
	c := (*C.GSocketListener)(u)
	if c == nil {
		return nil
	}

	g := &SocketListener{native: c}

	return g
}

func (recv *SocketListener) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SocketListener) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SocketListener.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketListener.
func CastToSocketListener(object *gobject.Object) *SocketListener {
	return SocketListenerNewFromC(object.ToC())
}

// Creates a new #GSocketListener with no sockets to listen for.
// New listeners can be added with e.g. g_socket_listener_add_address()
// or g_socket_listener_add_inet_port().
/*

C function : g_socket_listener_new
*/
func SocketListenerNew() *SocketListener {
	retC := C.g_socket_listener_new()
	retGo := SocketListenerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blocks waiting for a client to connect to any of the sockets added
// to the listener. Returns a #GSocketConnection for the socket that was
// accepted.
//
// If @source_object is not %NULL it will be filled out with the source
// object specified when the corresponding socket or address was added
// to the listener.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_socket_listener_accept
*/
func (recv *SocketListener) Accept(cancellable *Cancellable) (*SocketConnection, *gobject.Object, error) {
	var c_source_object *C.GObject

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept((*C.GSocketListener)(recv.native), &c_source_object, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goThrowableError
}

// Unsupported : g_socket_listener_accept_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async accept operation. See g_socket_listener_accept_async()
/*

C function : g_socket_listener_accept_finish
*/
func (recv *SocketListener) AcceptFinish(result *AsyncResult) (*SocketConnection, *gobject.Object, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_source_object *C.GObject

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_finish((*C.GSocketListener)(recv.native), c_result, &c_source_object, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goThrowableError
}

// Blocks waiting for a client to connect to any of the sockets added
// to the listener. Returns the #GSocket that was accepted.
//
// If you want to accept the high-level #GSocketConnection, not a #GSocket,
// which is often the case, then you should use g_socket_listener_accept()
// instead.
//
// If @source_object is not %NULL it will be filled out with the source
// object specified when the corresponding socket or address was added
// to the listener.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
/*

C function : g_socket_listener_accept_socket
*/
func (recv *SocketListener) AcceptSocket(cancellable *Cancellable) (*Socket, *gobject.Object, error) {
	var c_source_object *C.GObject

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_socket((*C.GSocketListener)(recv.native), &c_source_object, c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goThrowableError
}

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an async accept operation. See g_socket_listener_accept_socket_async()
/*

C function : g_socket_listener_accept_socket_finish
*/
func (recv *SocketListener) AcceptSocketFinish(result *AsyncResult) (*Socket, *gobject.Object, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_source_object *C.GObject

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_socket_finish((*C.GSocketListener)(recv.native), c_result, &c_source_object, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goThrowableError
}

// Creates a socket of type @type and protocol @protocol, binds
// it to @address and adds it to the set of sockets we're accepting
// sockets from.
//
// Note that adding an IPv6 address, depending on the platform,
// may or may not result in a listener that also accepts IPv4
// connections.  For more deterministic behavior, see
// g_socket_listener_add_inet_port().
//
// @source_object will be passed out in the various calls
// to accept to identify this particular source, which is
// useful if you're listening on multiple addresses and do
// different things depending on what address is connected to.
//
// If successful and @effective_address is non-%NULL then it will
// be set to the address that the binding actually occurred at.  This
// is helpful for determining the port number that was used for when
// requesting a binding to port 0 (ie: "any port").  This address, if
// requested, belongs to the caller and must be freed.
/*

C function : g_socket_listener_add_address
*/
func (recv *SocketListener) AddAddress(address *SocketAddress, type_ SocketType, protocol SocketProtocol, sourceObject *gobject.Object) (bool, *SocketAddress, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var c_effective_address *C.GSocketAddress

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_address((*C.GSocketListener)(recv.native), c_address, c_type, c_protocol, c_source_object, &c_effective_address, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	effectiveAddress := SocketAddressNewFromC(unsafe.Pointer(c_effective_address))

	return retGo, effectiveAddress, goThrowableError
}

// Helper function for g_socket_listener_add_address() that
// creates a TCP/IP socket listening on IPv4 and IPv6 (if
// supported) on the specified port on all interfaces.
//
// @source_object will be passed out in the various calls
// to accept to identify this particular source, which is
// useful if you're listening on multiple addresses and do
// different things depending on what address is connected to.
/*

C function : g_socket_listener_add_inet_port
*/
func (recv *SocketListener) AddInetPort(port uint16, sourceObject *gobject.Object) (bool, error) {
	c_port := (C.guint16)(port)

	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_inet_port((*C.GSocketListener)(recv.native), c_port, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Adds @socket to the set of sockets that we try to accept
// new clients from. The socket must be bound to a local
// address and listened to.
//
// @source_object will be passed out in the various calls
// to accept to identify this particular source, which is
// useful if you're listening on multiple addresses and do
// different things depending on what address is connected to.
//
// The @socket will not be automatically closed when the @listener is finalized
// unless the listener held the final reference to the socket. Before GLib 2.42,
// the @socket was automatically closed on finalization of the @listener, even
// if references to it were held elsewhere.
/*

C function : g_socket_listener_add_socket
*/
func (recv *SocketListener) AddSocket(socket *Socket, sourceObject *gobject.Object) (bool, error) {
	c_socket := (*C.GSocket)(C.NULL)
	if socket != nil {
		c_socket = (*C.GSocket)(socket.ToC())
	}

	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_socket((*C.GSocketListener)(recv.native), c_socket, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Closes all the sockets in the listener.
/*

C function : g_socket_listener_close
*/
func (recv *SocketListener) Close() {
	C.g_socket_listener_close((*C.GSocketListener)(recv.native))

	return
}

// Sets the listen backlog on the sockets in the listener.
//
// See g_socket_set_listen_backlog() for details
/*

C function : g_socket_listener_set_backlog
*/
func (recv *SocketListener) SetBacklog(listenBacklog int32) {
	c_listen_backlog := (C.int)(listenBacklog)

	C.g_socket_listener_set_backlog((*C.GSocketListener)(recv.native), c_listen_backlog)

	return
}

// A #GSocketService is an object that represents a service that
// is provided to the network or over local sockets.  When a new
// connection is made to the service the #GSocketService::incoming
// signal is emitted.
//
// A #GSocketService is a subclass of #GSocketListener and you need
// to add the addresses you want to accept connections on with the
// #GSocketListener APIs.
//
// There are two options for implementing a network service based on
// #GSocketService. The first is to create the service using
// g_socket_service_new() and to connect to the #GSocketService::incoming
// signal. The second is to subclass #GSocketService and override the
// default signal handler implementation.
//
// In either case, the handler must immediately return, or else it
// will block additional incoming connections from being serviced.
// If you are interested in writing connection handlers that contain
// blocking code then see #GThreadedSocketService.
//
// The socket service runs on the main loop of the
// [thread-default context][g-main-context-push-thread-default-context]
// of the thread it is created in, and is not
// threadsafe in general. However, the calls to start and stop the
// service are thread-safe so these can be used from threads that
// handle incoming clients.
/*

C record/class : GSocketService
*/
type SocketService struct {
	native *C.GSocketService
	// parent_instance : record
	// priv : record
}

func SocketServiceNewFromC(u unsafe.Pointer) *SocketService {
	c := (*C.GSocketService)(u)
	if c == nil {
		return nil
	}

	g := &SocketService{native: c}

	return g
}

func (recv *SocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketListener upcasts to *SocketListener
func (recv *SocketService) SocketListener() *SocketListener {
	return SocketListenerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SocketService) Object() *gobject.Object {
	return recv.SocketListener().Object()
}

// CastToWidget down casts any arbitary Object to SocketService.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketService.
func CastToSocketService(object *gobject.Object) *SocketService {
	return SocketServiceNewFromC(object.ToC())
}

// Creates a new #GSocketService with no sockets to listen for.
// New listeners can be added with e.g. g_socket_listener_add_address()
// or g_socket_listener_add_inet_port().
//
// New services are created active, there is no need to call
// g_socket_service_start(), unless g_socket_service_stop() has been
// called before.
/*

C function : g_socket_service_new
*/
func SocketServiceNew() *SocketService {
	retC := C.g_socket_service_new()
	retGo := SocketServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Check whether the service is active or not. An active
// service will accept new clients that connect, while
// a non-active service will let connecting clients queue
// up until the service is started.
/*

C function : g_socket_service_is_active
*/
func (recv *SocketService) IsActive() bool {
	retC := C.g_socket_service_is_active((*C.GSocketService)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Restarts the service, i.e. start accepting connections
// from the added sockets when the mainloop runs. This only needs
// to be called after the service has been stopped from
// g_socket_service_stop().
//
// This call is thread-safe, so it may be called from a thread
// handling an incoming client request.
/*

C function : g_socket_service_start
*/
func (recv *SocketService) Start() {
	C.g_socket_service_start((*C.GSocketService)(recv.native))

	return
}

// Stops the service, i.e. stops accepting connections
// from the added sockets when the mainloop runs.
//
// This call is thread-safe, so it may be called from a thread
// handling an incoming client request.
//
// Note that this only stops accepting new connections; it does not
// close the listening sockets, and you can call
// g_socket_service_start() again later to begin listening again. To
// close the listening sockets, call g_socket_listener_close(). (This
// will happen automatically when the #GSocketService is finalized.)
//
// This must be called before calling g_socket_listener_close() as
// the socket service will start accepting connections immediately
// when a new socket is added.
/*

C function : g_socket_service_stop
*/
func (recv *SocketService) Stop() {
	C.g_socket_service_stop((*C.GSocketService)(recv.native))

	return
}

// This is the subclass of #GSocketConnection that is created
// for TCP/IP sockets.
/*

C record/class : GTcpConnection
*/
type TcpConnection struct {
	native *C.GTcpConnection
	// parent_instance : record
	// priv : record
}

func TcpConnectionNewFromC(u unsafe.Pointer) *TcpConnection {
	c := (*C.GTcpConnection)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnection{native: c}

	return g
}

func (recv *TcpConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketConnection upcasts to *SocketConnection
func (recv *TcpConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromC(unsafe.Pointer(recv.native))
}

// IOStream upcasts to *IOStream
func (recv *TcpConnection) IOStream() *IOStream {
	return recv.SocketConnection().IOStream()
}

// Object upcasts to *Object
func (recv *TcpConnection) Object() *gobject.Object {
	return recv.SocketConnection().Object()
}

// CastToWidget down casts any arbitary Object to TcpConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TcpConnection.
func CastToTcpConnection(object *gobject.Object) *TcpConnection {
	return TcpConnectionNewFromC(object.ToC())
}

// Checks if graceful disconnects are used. See
// g_tcp_connection_set_graceful_disconnect().
/*

C function : g_tcp_connection_get_graceful_disconnect
*/
func (recv *TcpConnection) GetGracefulDisconnect() bool {
	retC := C.g_tcp_connection_get_graceful_disconnect((*C.GTcpConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This enables graceful disconnects on close. A graceful disconnect
// means that we signal the receiving end that the connection is terminated
// and wait for it to close the connection before closing the connection.
//
// A graceful disconnect means that we can be sure that we successfully sent
// all the outstanding data to the other end, or get an error reported.
// However, it also means we have to wait for all the data to reach the
// other side and for it to acknowledge this by closing the socket, which may
// take a while. For this reason it is disabled by default.
/*

C function : g_tcp_connection_set_graceful_disconnect
*/
func (recv *TcpConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	c_graceful_disconnect :=
		boolToGboolean(gracefulDisconnect)

	C.g_tcp_connection_set_graceful_disconnect((*C.GTcpConnection)(recv.native), c_graceful_disconnect)

	return
}

// A #GThreadedSocketService is a simple subclass of #GSocketService
// that handles incoming connections by creating a worker thread and
// dispatching the connection to it by emitting the
// #GThreadedSocketService::run signal in the new thread.
//
// The signal handler may perform blocking IO and need not return
// until the connection is closed.
//
// The service is implemented using a thread pool, so there is a
// limited amount of threads available to serve incoming requests.
// The service automatically stops the #GSocketService from accepting
// new connections when all threads are busy.
//
// As with #GSocketService, you may connect to #GThreadedSocketService::run,
// or subclass and override the default handler.
/*

C record/class : GThreadedSocketService
*/
type ThreadedSocketService struct {
	native *C.GThreadedSocketService
	// parent_instance : record
	// priv : record
}

func ThreadedSocketServiceNewFromC(u unsafe.Pointer) *ThreadedSocketService {
	c := (*C.GThreadedSocketService)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketService{native: c}

	return g
}

func (recv *ThreadedSocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketService upcasts to *SocketService
func (recv *ThreadedSocketService) SocketService() *SocketService {
	return SocketServiceNewFromC(unsafe.Pointer(recv.native))
}

// SocketListener upcasts to *SocketListener
func (recv *ThreadedSocketService) SocketListener() *SocketListener {
	return recv.SocketService().SocketListener()
}

// Object upcasts to *Object
func (recv *ThreadedSocketService) Object() *gobject.Object {
	return recv.SocketService().Object()
}

// CastToWidget down casts any arbitary Object to ThreadedSocketService.
// Exercise care, as this is a potentially dangerous function if the Object is not a ThreadedSocketService.
func CastToThreadedSocketService(object *gobject.Object) *ThreadedSocketService {
	return ThreadedSocketServiceNewFromC(object.ToC())
}

type signalThreadedSocketServiceRunDetail struct {
	callback  ThreadedSocketServiceSignalRunCallback
	handlerID C.gulong
}

var signalThreadedSocketServiceRunId int
var signalThreadedSocketServiceRunMap = make(map[int]signalThreadedSocketServiceRunDetail)
var signalThreadedSocketServiceRunLock sync.Mutex

// ThreadedSocketServiceSignalRunCallback is a callback function for a 'run' signal emitted from a ThreadedSocketService.
type ThreadedSocketServiceSignalRunCallback func(connection *SocketConnection, sourceObject *gobject.Object) bool

/*
ConnectRun connects the callback to the 'run' signal for the ThreadedSocketService.

The returned value represents the connection, and may be passed to DisconnectRun to remove it.
*/
func (recv *ThreadedSocketService) ConnectRun(callback ThreadedSocketServiceSignalRunCallback) int {
	signalThreadedSocketServiceRunLock.Lock()
	defer signalThreadedSocketServiceRunLock.Unlock()

	signalThreadedSocketServiceRunId++
	instance := C.gpointer(recv.native)
	handlerID := C.ThreadedSocketService_signal_connect_run(instance, C.gpointer(uintptr(signalThreadedSocketServiceRunId)))

	detail := signalThreadedSocketServiceRunDetail{callback, handlerID}
	signalThreadedSocketServiceRunMap[signalThreadedSocketServiceRunId] = detail

	return signalThreadedSocketServiceRunId
}

/*
DisconnectRun disconnects a callback from the 'run' signal for the ThreadedSocketService.

The connectionID should be a value returned from a call to ConnectRun.
*/
func (recv *ThreadedSocketService) DisconnectRun(connectionID int) {
	signalThreadedSocketServiceRunLock.Lock()
	defer signalThreadedSocketServiceRunLock.Unlock()

	detail, exists := signalThreadedSocketServiceRunMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalThreadedSocketServiceRunMap, connectionID)
}

//export threadedsocketservice_runHandler
func threadedsocketservice_runHandler(_ *C.GObject, c_connection *C.GSocketConnection, c_source_object *C.GObject, data C.gpointer) C.gboolean {
	connection := SocketConnectionNewFromC(unsafe.Pointer(c_connection))

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	index := int(uintptr(data))
	callback := signalThreadedSocketServiceRunMap[index].callback
	retGo := callback(connection, sourceObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Creates a new #GThreadedSocketService with no listeners. Listeners
// must be added with one of the #GSocketListener "add" methods.
/*

C function : g_threaded_socket_service_new
*/
func ThreadedSocketServiceNew(maxThreads int32) *ThreadedSocketService {
	c_max_threads := (C.int)(maxThreads)

	retC := C.g_threaded_socket_service_new(c_max_threads)
	retGo := ThreadedSocketServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Receives a file descriptor from the sending end of the connection.
// The sending end has to call g_unix_connection_send_fd() for this
// to work.
//
// As well as reading the fd this also reads a single byte from the
// stream, as this is required for fd passing to work on some
// implementations.
/*

C function : g_unix_connection_receive_fd
*/
func (recv *UnixConnection) ReceiveFd(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_fd((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Passes a file descriptor to the receiving side of the
// connection. The receiving end has to call g_unix_connection_receive_fd()
// to accept the file descriptor.
//
// As well as sending the fd this also writes a single byte to the
// stream, as this is required for fd passing to work on some
// implementations.
/*

C function : g_unix_connection_send_fd
*/
func (recv *UnixConnection) SendFd(fd int32, cancellable *Cancellable) (bool, error) {
	c_fd := (C.gint)(fd)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_fd((*C.GUnixConnection)(recv.native), c_fd, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Creates a new #GUnixFDMessage containing an empty file descriptor
// list.
/*

C function : g_unix_fd_message_new
*/
func UnixFDMessageNew() *UnixFDMessage {
	retC := C.g_unix_fd_message_new()
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a file descriptor to @message.
//
// The file descriptor is duplicated using dup(). You keep your copy
// of the descriptor and the copy contained in @message will be closed
// when @message is finalized.
//
// A possible cause of failure is exceeding the per-process or
// system-wide file descriptor limit.
/*

C function : g_unix_fd_message_append_fd
*/
func (recv *UnixFDMessage) AppendFd(fd int32) (bool, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_unix_fd_message_append_fd((*C.GUnixFDMessage)(recv.native), c_fd, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_unix_fd_message_steal_fds : no return type

// Creates a new #GUnixSocketAddress for @path.
//
// To create abstract socket addresses, on systems that support that,
// use g_unix_socket_address_new_abstract().
/*

C function : g_unix_socket_address_new
*/
func UnixSocketAddressNew(path string) *UnixSocketAddress {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_unix_socket_address_new(c_path)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Tests if @address is abstract.
/*

C function : g_unix_socket_address_get_is_abstract
*/
func (recv *UnixSocketAddress) GetIsAbstract() bool {
	retC := C.g_unix_socket_address_get_is_abstract((*C.GUnixSocketAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets @address's path, or for abstract sockets the "name".
//
// Guaranteed to be zero-terminated, but an abstract socket
// may contain embedded zeros, and thus you should use
// g_unix_socket_address_get_path_len() to get the true length
// of this string.
/*

C function : g_unix_socket_address_get_path
*/
func (recv *UnixSocketAddress) GetPath() string {
	retC := C.g_unix_socket_address_get_path((*C.GUnixSocketAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the length of @address's path.
//
// For details, see g_unix_socket_address_get_path().
/*

C function : g_unix_socket_address_get_path_len
*/
func (recv *UnixSocketAddress) GetPathLen() uint64 {
	retC := C.g_unix_socket_address_get_path_len((*C.GUnixSocketAddress)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

type signalVolumeMonitorDriveStopButtonDetail struct {
	callback  VolumeMonitorSignalDriveStopButtonCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveStopButtonId int
var signalVolumeMonitorDriveStopButtonMap = make(map[int]signalVolumeMonitorDriveStopButtonDetail)
var signalVolumeMonitorDriveStopButtonLock sync.Mutex

// VolumeMonitorSignalDriveStopButtonCallback is a callback function for a 'drive-stop-button' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveStopButtonCallback func(drive *Drive)

/*
ConnectDriveStopButton connects the callback to the 'drive-stop-button' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveStopButton to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveStopButton(callback VolumeMonitorSignalDriveStopButtonCallback) int {
	signalVolumeMonitorDriveStopButtonLock.Lock()
	defer signalVolumeMonitorDriveStopButtonLock.Unlock()

	signalVolumeMonitorDriveStopButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_stop_button(instance, C.gpointer(uintptr(signalVolumeMonitorDriveStopButtonId)))

	detail := signalVolumeMonitorDriveStopButtonDetail{callback, handlerID}
	signalVolumeMonitorDriveStopButtonMap[signalVolumeMonitorDriveStopButtonId] = detail

	return signalVolumeMonitorDriveStopButtonId
}

/*
DisconnectDriveStopButton disconnects a callback from the 'drive-stop-button' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveStopButton.
*/
func (recv *VolumeMonitor) DisconnectDriveStopButton(connectionID int) {
	signalVolumeMonitorDriveStopButtonLock.Lock()
	defer signalVolumeMonitorDriveStopButtonLock.Unlock()

	detail, exists := signalVolumeMonitorDriveStopButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveStopButtonMap, connectionID)
}

//export volumemonitor_driveStopButtonHandler
func volumemonitor_driveStopButtonHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveStopButtonMap[index].callback
	callback(drive)
}
