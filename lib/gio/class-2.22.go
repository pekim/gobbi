// This is a generated file - DO NOT EDIT
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
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

// Unsupported : g_cancellable_connect : unsupported parameter callback : no type generator for GObject.Callback, GCallback

// Disconnect is a wrapper around the C function g_cancellable_disconnect.
func (recv *Cancellable) Disconnect(handlerId uint64) {
	c_handler_id := (C.gulong)(handlerId)

	C.g_cancellable_disconnect((*C.GCancellable)(recv.native), c_handler_id)

	return
}

// MakePollfd is a wrapper around the C function g_cancellable_make_pollfd.
func (recv *Cancellable) MakePollfd(pollfd *glib.PollFD) bool {
	c_pollfd := (*C.GPollFD)(pollfd.ToC())

	retC := C.g_cancellable_make_pollfd((*C.GCancellable)(recv.native), c_pollfd)
	retGo := retC == C.TRUE

	return retGo
}

// ReleaseFd is a wrapper around the C function g_cancellable_release_fd.
func (recv *Cancellable) ReleaseFd() {
	C.g_cancellable_release_fd((*C.GCancellable)(recv.native))

	return
}

// Unsupported : g_converter_input_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_converter_output_stream_new : unsupported parameter converter : no type generator for Converter, GConverter*

// Unsupported : g_emblem_new : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblem_new_with_origin : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_emblemed_icon_new : unsupported parameter icon : no type generator for Icon, GIcon*

// GetEtag is a wrapper around the C function g_file_io_stream_get_etag.
func (recv *FileIOStream) GetEtag() string {
	retC := C.g_file_io_stream_get_etag((*C.GFileIOStream)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// QueryInfo is a wrapper around the C function g_file_io_stream_query_info.
func (recv *FileIOStream) QueryInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_io_stream_query_info((*C.GFileIOStream)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_file_io_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_file_io_stream_query_info_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_file_icon_new : unsupported parameter file : no type generator for File, GFile*

// Unsupported : g_file_info_get_attribute_stringv : no return type

// HasNamespace is a wrapper around the C function g_file_info_has_namespace.
func (recv *FileInfo) HasNamespace(nameSpace string) bool {
	c_name_space := C.CString(nameSpace)
	defer C.free(unsafe.Pointer(c_name_space))

	retC := C.g_file_info_has_namespace((*C.GFileInfo)(recv.native), c_name_space)
	retGo := retC == C.TRUE

	return retGo
}

// SetAttributeStatus is a wrapper around the C function g_file_info_set_attribute_status.
func (recv *FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_status := (C.GFileAttributeStatus)(status)

	retC := C.g_file_info_set_attribute_status((*C.GFileInfo)(recv.native), c_attribute, c_status)
	retGo := retC == C.TRUE

	return retGo
}

// ClearPending is a wrapper around the C function g_io_stream_clear_pending.
func (recv *IOStream) ClearPending() {
	C.g_io_stream_clear_pending((*C.GIOStream)(recv.native))

	return
}

// Close is a wrapper around the C function g_io_stream_close.
func (recv *IOStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_io_stream_close((*C.GIOStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_io_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_io_stream_close_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetInputStream is a wrapper around the C function g_io_stream_get_input_stream.
func (recv *IOStream) GetInputStream() *InputStream {
	retC := C.g_io_stream_get_input_stream((*C.GIOStream)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOutputStream is a wrapper around the C function g_io_stream_get_output_stream.
func (recv *IOStream) GetOutputStream() *OutputStream {
	retC := C.g_io_stream_get_output_stream((*C.GIOStream)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasPending is a wrapper around the C function g_io_stream_has_pending.
func (recv *IOStream) HasPending() bool {
	retC := C.g_io_stream_has_pending((*C.GIOStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsClosed is a wrapper around the C function g_io_stream_is_closed.
func (recv *IOStream) IsClosed() bool {
	retC := C.g_io_stream_is_closed((*C.GIOStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPending is a wrapper around the C function g_io_stream_set_pending.
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

// InetAddressNewAny is a wrapper around the C function g_inet_address_new_any.
func InetAddressNewAny(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_any(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_inet_address_new_from_bytes : unsupported parameter bytes : no param type

// InetAddressNewFromString is a wrapper around the C function g_inet_address_new_from_string.
func InetAddressNewFromString(string string) *InetAddress {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_inet_address_new_from_string(c_string)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InetAddressNewLoopback is a wrapper around the C function g_inet_address_new_loopback.
func InetAddressNewLoopback(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_loopback(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamily is a wrapper around the C function g_inet_address_get_family.
func (recv *InetAddress) GetFamily() SocketFamily {
	retC := C.g_inet_address_get_family((*C.GInetAddress)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetIsAny is a wrapper around the C function g_inet_address_get_is_any.
func (recv *InetAddress) GetIsAny() bool {
	retC := C.g_inet_address_get_is_any((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsLinkLocal is a wrapper around the C function g_inet_address_get_is_link_local.
func (recv *InetAddress) GetIsLinkLocal() bool {
	retC := C.g_inet_address_get_is_link_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsLoopback is a wrapper around the C function g_inet_address_get_is_loopback.
func (recv *InetAddress) GetIsLoopback() bool {
	retC := C.g_inet_address_get_is_loopback((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcGlobal is a wrapper around the C function g_inet_address_get_is_mc_global.
func (recv *InetAddress) GetIsMcGlobal() bool {
	retC := C.g_inet_address_get_is_mc_global((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcLinkLocal is a wrapper around the C function g_inet_address_get_is_mc_link_local.
func (recv *InetAddress) GetIsMcLinkLocal() bool {
	retC := C.g_inet_address_get_is_mc_link_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcNodeLocal is a wrapper around the C function g_inet_address_get_is_mc_node_local.
func (recv *InetAddress) GetIsMcNodeLocal() bool {
	retC := C.g_inet_address_get_is_mc_node_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcOrgLocal is a wrapper around the C function g_inet_address_get_is_mc_org_local.
func (recv *InetAddress) GetIsMcOrgLocal() bool {
	retC := C.g_inet_address_get_is_mc_org_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcSiteLocal is a wrapper around the C function g_inet_address_get_is_mc_site_local.
func (recv *InetAddress) GetIsMcSiteLocal() bool {
	retC := C.g_inet_address_get_is_mc_site_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMulticast is a wrapper around the C function g_inet_address_get_is_multicast.
func (recv *InetAddress) GetIsMulticast() bool {
	retC := C.g_inet_address_get_is_multicast((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsSiteLocal is a wrapper around the C function g_inet_address_get_is_site_local.
func (recv *InetAddress) GetIsSiteLocal() bool {
	retC := C.g_inet_address_get_is_site_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetNativeSize is a wrapper around the C function g_inet_address_get_native_size.
func (recv *InetAddress) GetNativeSize() uint64 {
	retC := C.g_inet_address_get_native_size((*C.GInetAddress)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Blacklisted : g_inet_address_to_bytes

// ToString is a wrapper around the C function g_inet_address_to_string.
func (recv *InetAddress) ToString() string {
	retC := C.g_inet_address_to_string((*C.GInetAddress)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InetSocketAddressNew is a wrapper around the C function g_inet_socket_address_new.
func InetSocketAddressNew(address *InetAddress, port uint16) *InetSocketAddress {
	c_address := (*C.GInetAddress)(address.ToC())

	c_port := (C.guint16)(port)

	retC := C.g_inet_socket_address_new(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAddress is a wrapper around the C function g_inet_socket_address_get_address.
func (recv *InetSocketAddress) GetAddress() *InetAddress {
	retC := C.g_inet_socket_address_get_address((*C.GInetSocketAddress)(recv.native))
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPort is a wrapper around the C function g_inet_socket_address_get_port.
func (recv *InetSocketAddress) GetPort() uint16 {
	retC := C.g_inet_socket_address_get_port((*C.GInetSocketAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported : g_list_store_new : unsupported parameter item_type : no type generator for GType, GType

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter data : no param type

// Unsupported : g_memory_output_stream_new : unsupported parameter realloc_function : no type generator for ReallocFunc, GReallocFunc

// Unsupported signal 'show-processes' for MountOperation : unsupported parameter message : type utf8 :

// NetworkAddressNew is a wrapper around the C function g_network_address_new.
func NetworkAddressNew(hostname string, port uint16) *NetworkAddress {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	retC := C.g_network_address_new(c_hostname, c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHostname is a wrapper around the C function g_network_address_get_hostname.
func (recv *NetworkAddress) GetHostname() string {
	retC := C.g_network_address_get_hostname((*C.GNetworkAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPort is a wrapper around the C function g_network_address_get_port.
func (recv *NetworkAddress) GetPort() uint16 {
	retC := C.g_network_address_get_port((*C.GNetworkAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// NetworkServiceNew is a wrapper around the C function g_network_service_new.
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

// GetDomain is a wrapper around the C function g_network_service_get_domain.
func (recv *NetworkService) GetDomain() string {
	retC := C.g_network_service_get_domain((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProtocol is a wrapper around the C function g_network_service_get_protocol.
func (recv *NetworkService) GetProtocol() string {
	retC := C.g_network_service_get_protocol((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetService is a wrapper around the C function g_network_service_get_service.
func (recv *NetworkService) GetService() string {
	retC := C.g_network_service_get_service((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// LookupByAddress is a wrapper around the C function g_resolver_lookup_by_address.
func (recv *Resolver) LookupByAddress(address *InetAddress, cancellable *Cancellable) (string, error) {
	c_address := (*C.GInetAddress)(address.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

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

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_address_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// LookupByName is a wrapper around the C function g_resolver_lookup_by_name.
func (recv *Resolver) LookupByName(hostname string, cancellable *Cancellable) (*glib.List, error) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_name((*C.GResolver)(recv.native), c_hostname, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_by_name_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// LookupService is a wrapper around the C function g_resolver_lookup_service.
func (recv *Resolver) LookupService(service string, protocol string, domain string, cancellable *Cancellable) (*glib.List, error) {
	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_service((*C.GResolver)(recv.native), c_service, c_protocol, c_domain, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_resolver_lookup_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_resolver_lookup_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// SetDefault is a wrapper around the C function g_resolver_set_default.
func (recv *Resolver) SetDefault() {
	C.g_resolver_set_default((*C.GResolver)(recv.native))

	return
}

// Unsupported : g_simple_action_new : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_action_new_stateful : unsupported parameter parameter_type : Blacklisted record : GVariantType

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Socket is a wrapper around the C record GSocket.
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

// SocketNew is a wrapper around the C function g_socket_new.
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

// SocketNewFromFd is a wrapper around the C function g_socket_new_from_fd.
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

// Accept is a wrapper around the C function g_socket_accept.
func (recv *Socket) Accept(cancellable *Cancellable) (*Socket, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_accept((*C.GSocket)(recv.native), c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Bind is a wrapper around the C function g_socket_bind.
func (recv *Socket) Bind(address *SocketAddress, allowReuse bool) (bool, error) {
	c_address := (*C.GSocketAddress)(address.ToC())

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

// CheckConnectResult is a wrapper around the C function g_socket_check_connect_result.
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

// Close is a wrapper around the C function g_socket_close.
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

// ConditionCheck is a wrapper around the C function g_socket_condition_check.
func (recv *Socket) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_socket_condition_check((*C.GSocket)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// ConditionWait is a wrapper around the C function g_socket_condition_wait.
func (recv *Socket) ConditionWait(condition glib.IOCondition, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_condition_wait((*C.GSocket)(recv.native), c_condition, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Connect is a wrapper around the C function g_socket_connect.
func (recv *Socket) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(address.ToC())

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_connect((*C.GSocket)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// ConnectionFactoryCreateConnection is a wrapper around the C function g_socket_connection_factory_create_connection.
func (recv *Socket) ConnectionFactoryCreateConnection() *SocketConnection {
	retC := C.g_socket_connection_factory_create_connection((*C.GSocket)(recv.native))
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateSource is a wrapper around the C function g_socket_create_source.
func (recv *Socket) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	retC := C.g_socket_create_source((*C.GSocket)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBlocking is a wrapper around the C function g_socket_get_blocking.
func (recv *Socket) GetBlocking() bool {
	retC := C.g_socket_get_blocking((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFamily is a wrapper around the C function g_socket_get_family.
func (recv *Socket) GetFamily() SocketFamily {
	retC := C.g_socket_get_family((*C.GSocket)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetFd is a wrapper around the C function g_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	retC := C.g_socket_get_fd((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetKeepalive is a wrapper around the C function g_socket_get_keepalive.
func (recv *Socket) GetKeepalive() bool {
	retC := C.g_socket_get_keepalive((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetListenBacklog is a wrapper around the C function g_socket_get_listen_backlog.
func (recv *Socket) GetListenBacklog() int32 {
	retC := C.g_socket_get_listen_backlog((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_socket_get_local_address

// GetProtocol is a wrapper around the C function g_socket_get_protocol.
func (recv *Socket) GetProtocol() SocketProtocol {
	retC := C.g_socket_get_protocol((*C.GSocket)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// GetRemoteAddress is a wrapper around the C function g_socket_get_remote_address.
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

// GetSocketType is a wrapper around the C function g_socket_get_socket_type.
func (recv *Socket) GetSocketType() SocketType {
	retC := C.g_socket_get_socket_type((*C.GSocket)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// IsClosed is a wrapper around the C function g_socket_is_closed.
func (recv *Socket) IsClosed() bool {
	retC := C.g_socket_is_closed((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsConnected is a wrapper around the C function g_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	retC := C.g_socket_is_connected((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Listen is a wrapper around the C function g_socket_listen.
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

// Unsupported : g_socket_receive : unsupported parameter buffer : no param type

// Unsupported : g_socket_receive_from : unsupported parameter address : record with indirection level of 2

// Unsupported : g_socket_receive_message : unsupported parameter address : record with indirection level of 2

// Unsupported : g_socket_send : unsupported parameter buffer : no param type

// Unsupported : g_socket_send_message : unsupported parameter vectors : no param type

// Unsupported : g_socket_send_to : unsupported parameter buffer : no param type

// SetBlocking is a wrapper around the C function g_socket_set_blocking.
func (recv *Socket) SetBlocking(blocking bool) {
	c_blocking :=
		boolToGboolean(blocking)

	C.g_socket_set_blocking((*C.GSocket)(recv.native), c_blocking)

	return
}

// SetKeepalive is a wrapper around the C function g_socket_set_keepalive.
func (recv *Socket) SetKeepalive(keepalive bool) {
	c_keepalive :=
		boolToGboolean(keepalive)

	C.g_socket_set_keepalive((*C.GSocket)(recv.native), c_keepalive)

	return
}

// SetListenBacklog is a wrapper around the C function g_socket_set_listen_backlog.
func (recv *Socket) SetListenBacklog(backlog int32) {
	c_backlog := (C.gint)(backlog)

	C.g_socket_set_listen_backlog((*C.GSocket)(recv.native), c_backlog)

	return
}

// Shutdown is a wrapper around the C function g_socket_shutdown.
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

// SpeaksIpv4 is a wrapper around the C function g_socket_speaks_ipv4.
func (recv *Socket) SpeaksIpv4() bool {
	retC := C.g_socket_speaks_ipv4((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SocketAddressNewFromNative is a wrapper around the C function g_socket_address_new_from_native.
func SocketAddressNewFromNative(native uintptr, len uint64) *SocketAddress {
	c_native := (C.gpointer)(native)

	c_len := (C.gsize)(len)

	retC := C.g_socket_address_new_from_native(c_native, c_len)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamily is a wrapper around the C function g_socket_address_get_family.
func (recv *SocketAddress) GetFamily() SocketFamily {
	retC := C.g_socket_address_get_family((*C.GSocketAddress)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetNativeSize is a wrapper around the C function g_socket_address_get_native_size.
func (recv *SocketAddress) GetNativeSize() int64 {
	retC := C.g_socket_address_get_native_size((*C.GSocketAddress)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// ToNative is a wrapper around the C function g_socket_address_to_native.
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

// SocketClient is a wrapper around the C record GSocketClient.
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

// SocketClientNew is a wrapper around the C function g_socket_client_new.
func SocketClientNew() *SocketClient {
	retC := C.g_socket_client_new()
	retGo := SocketClientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddApplicationProxy is a wrapper around the C function g_socket_client_add_application_proxy.
func (recv *SocketClient) AddApplicationProxy(protocol string) {
	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	C.g_socket_client_add_application_proxy((*C.GSocketClient)(recv.native), c_protocol)

	return
}

// Unsupported : g_socket_client_connect : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_async : unsupported parameter connectable : no type generator for SocketConnectable, GSocketConnectable*

// Unsupported : g_socket_client_connect_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ConnectToHost is a wrapper around the C function g_socket_client_connect_to_host.
func (recv *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_host_and_port := C.CString(hostAndPort)
	defer C.free(unsafe.Pointer(c_host_and_port))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_host((*C.GSocketClient)(recv.native), c_host_and_port, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_host_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// ConnectToService is a wrapper around the C function g_socket_client_connect_to_service.
func (recv *SocketClient) ConnectToService(domain string, service string, cancellable *Cancellable) (*SocketConnection, error) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_service((*C.GSocketClient)(recv.native), c_domain, c_service, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_client_connect_to_service_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// GetFamily is a wrapper around the C function g_socket_client_get_family.
func (recv *SocketClient) GetFamily() SocketFamily {
	retC := C.g_socket_client_get_family((*C.GSocketClient)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetLocalAddress is a wrapper around the C function g_socket_client_get_local_address.
func (recv *SocketClient) GetLocalAddress() *SocketAddress {
	retC := C.g_socket_client_get_local_address((*C.GSocketClient)(recv.native))
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProtocol is a wrapper around the C function g_socket_client_get_protocol.
func (recv *SocketClient) GetProtocol() SocketProtocol {
	retC := C.g_socket_client_get_protocol((*C.GSocketClient)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// GetSocketType is a wrapper around the C function g_socket_client_get_socket_type.
func (recv *SocketClient) GetSocketType() SocketType {
	retC := C.g_socket_client_get_socket_type((*C.GSocketClient)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// SetFamily is a wrapper around the C function g_socket_client_set_family.
func (recv *SocketClient) SetFamily(family SocketFamily) {
	c_family := (C.GSocketFamily)(family)

	C.g_socket_client_set_family((*C.GSocketClient)(recv.native), c_family)

	return
}

// SetLocalAddress is a wrapper around the C function g_socket_client_set_local_address.
func (recv *SocketClient) SetLocalAddress(address *SocketAddress) {
	c_address := (*C.GSocketAddress)(address.ToC())

	C.g_socket_client_set_local_address((*C.GSocketClient)(recv.native), c_address)

	return
}

// SetProtocol is a wrapper around the C function g_socket_client_set_protocol.
func (recv *SocketClient) SetProtocol(protocol SocketProtocol) {
	c_protocol := (C.GSocketProtocol)(protocol)

	C.g_socket_client_set_protocol((*C.GSocketClient)(recv.native), c_protocol)

	return
}

// SetSocketType is a wrapper around the C function g_socket_client_set_socket_type.
func (recv *SocketClient) SetSocketType(type_ SocketType) {
	c_type := (C.GSocketType)(type_)

	C.g_socket_client_set_socket_type((*C.GSocketClient)(recv.native), c_type)

	return
}

// SocketConnection is a wrapper around the C record GSocketConnection.
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

// GetLocalAddress is a wrapper around the C function g_socket_connection_get_local_address.
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

// GetRemoteAddress is a wrapper around the C function g_socket_connection_get_remote_address.
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

// GetSocket is a wrapper around the C function g_socket_connection_get_socket.
func (recv *SocketConnection) GetSocket() *Socket {
	retC := C.g_socket_connection_get_socket((*C.GSocketConnection)(recv.native))
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLevel is a wrapper around the C function g_socket_control_message_get_level.
func (recv *SocketControlMessage) GetLevel() int32 {
	retC := C.g_socket_control_message_get_level((*C.GSocketControlMessage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMsgType is a wrapper around the C function g_socket_control_message_get_msg_type.
func (recv *SocketControlMessage) GetMsgType() int32 {
	retC := C.g_socket_control_message_get_msg_type((*C.GSocketControlMessage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSize is a wrapper around the C function g_socket_control_message_get_size.
func (recv *SocketControlMessage) GetSize() uint64 {
	retC := C.g_socket_control_message_get_size((*C.GSocketControlMessage)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Serialize is a wrapper around the C function g_socket_control_message_serialize.
func (recv *SocketControlMessage) Serialize(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_socket_control_message_serialize((*C.GSocketControlMessage)(recv.native), c_data)

	return
}

// SocketListener is a wrapper around the C record GSocketListener.
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

// SocketListenerNew is a wrapper around the C function g_socket_listener_new.
func SocketListenerNew() *SocketListener {
	retC := C.g_socket_listener_new()
	retGo := SocketListenerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_socket_listener_accept : unsupported parameter source_object : record with indirection level of 2

// Unsupported : g_socket_listener_accept_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_listener_accept_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_listener_accept_socket : unsupported parameter source_object : record with indirection level of 2

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : g_socket_listener_accept_socket_finish : unsupported parameter result : no type generator for AsyncResult, GAsyncResult*

// Unsupported : g_socket_listener_add_address : unsupported parameter effective_address : record with indirection level of 2

// AddInetPort is a wrapper around the C function g_socket_listener_add_inet_port.
func (recv *SocketListener) AddInetPort(port uint16, sourceObject *gobject.Object) (bool, error) {
	c_port := (C.guint16)(port)

	c_source_object := (*C.GObject)(sourceObject.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_inet_port((*C.GSocketListener)(recv.native), c_port, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// AddSocket is a wrapper around the C function g_socket_listener_add_socket.
func (recv *SocketListener) AddSocket(socket *Socket, sourceObject *gobject.Object) (bool, error) {
	c_socket := (*C.GSocket)(socket.ToC())

	c_source_object := (*C.GObject)(sourceObject.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_socket((*C.GSocketListener)(recv.native), c_socket, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Close is a wrapper around the C function g_socket_listener_close.
func (recv *SocketListener) Close() {
	C.g_socket_listener_close((*C.GSocketListener)(recv.native))

	return
}

// SetBacklog is a wrapper around the C function g_socket_listener_set_backlog.
func (recv *SocketListener) SetBacklog(listenBacklog int32) {
	c_listen_backlog := (C.int)(listenBacklog)

	C.g_socket_listener_set_backlog((*C.GSocketListener)(recv.native), c_listen_backlog)

	return
}

// SocketService is a wrapper around the C record GSocketService.
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

// SocketServiceNew is a wrapper around the C function g_socket_service_new.
func SocketServiceNew() *SocketService {
	retC := C.g_socket_service_new()
	retGo := SocketServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsActive is a wrapper around the C function g_socket_service_is_active.
func (recv *SocketService) IsActive() bool {
	retC := C.g_socket_service_is_active((*C.GSocketService)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Start is a wrapper around the C function g_socket_service_start.
func (recv *SocketService) Start() {
	C.g_socket_service_start((*C.GSocketService)(recv.native))

	return
}

// Stop is a wrapper around the C function g_socket_service_stop.
func (recv *SocketService) Stop() {
	C.g_socket_service_stop((*C.GSocketService)(recv.native))

	return
}

// Unsupported : g_task_new : unsupported parameter callback : no type generator for AsyncReadyCallback, GAsyncReadyCallback

// TcpConnection is a wrapper around the C record GTcpConnection.
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

// GetGracefulDisconnect is a wrapper around the C function g_tcp_connection_get_graceful_disconnect.
func (recv *TcpConnection) GetGracefulDisconnect() bool {
	retC := C.g_tcp_connection_get_graceful_disconnect((*C.GTcpConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetGracefulDisconnect is a wrapper around the C function g_tcp_connection_set_graceful_disconnect.
func (recv *TcpConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	c_graceful_disconnect :=
		boolToGboolean(gracefulDisconnect)

	C.g_tcp_connection_set_graceful_disconnect((*C.GTcpConnection)(recv.native), c_graceful_disconnect)

	return
}

// Unsupported : g_themed_icon_new_from_names : unsupported parameter iconnames : no param type

// ThreadedSocketService is a wrapper around the C record GThreadedSocketService.
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

// Unsupported signal 'run' for ThreadedSocketService : unsupported parameter connection : type SocketConnection :

// ThreadedSocketServiceNew is a wrapper around the C function g_threaded_socket_service_new.
func ThreadedSocketServiceNew(maxThreads int32) *ThreadedSocketService {
	c_max_threads := (C.int)(maxThreads)

	retC := C.g_threaded_socket_service_new(c_max_threads)
	retGo := ThreadedSocketServiceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ReceiveFd is a wrapper around the C function g_unix_connection_receive_fd.
func (recv *UnixConnection) ReceiveFd(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_fd((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// SendFd is a wrapper around the C function g_unix_connection_send_fd.
func (recv *UnixConnection) SendFd(fd int32, cancellable *Cancellable) (bool, error) {
	c_fd := (C.gint)(fd)

	c_cancellable := (*C.GCancellable)(cancellable.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_fd((*C.GUnixConnection)(recv.native), c_fd, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_unix_fd_list_new_from_array : unsupported parameter fds : no param type

// UnixFDMessageNew is a wrapper around the C function g_unix_fd_message_new.
func UnixFDMessageNew() *UnixFDMessage {
	retC := C.g_unix_fd_message_new()
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendFd is a wrapper around the C function g_unix_fd_message_append_fd.
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

// UnixSocketAddressNew is a wrapper around the C function g_unix_socket_address_new.
func UnixSocketAddressNew(path string) *UnixSocketAddress {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_unix_socket_address_new(c_path)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_unix_socket_address_new_abstract : unsupported parameter path : no param type

// Unsupported : g_unix_socket_address_new_with_type : unsupported parameter path : no param type

// GetIsAbstract is a wrapper around the C function g_unix_socket_address_get_is_abstract.
func (recv *UnixSocketAddress) GetIsAbstract() bool {
	retC := C.g_unix_socket_address_get_is_abstract((*C.GUnixSocketAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPath is a wrapper around the C function g_unix_socket_address_get_path.
func (recv *UnixSocketAddress) GetPath() string {
	retC := C.g_unix_socket_address_get_path((*C.GUnixSocketAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPathLen is a wrapper around the C function g_unix_socket_address_get_path_len.
func (recv *UnixSocketAddress) GetPathLen() uint64 {
	retC := C.g_unix_socket_address_get_path_len((*C.GUnixSocketAddress)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported signal 'drive-stop-button' for VolumeMonitor : unsupported parameter drive : no type generator for Drive,
