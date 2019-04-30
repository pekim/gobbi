// This is a generated file - DO NOT EDIT

package gio

import gobject "github.com/pekim/gobbi/lib/gobject"

// Unsupported : g_app_launch_context_new : return type :

// Unsupported : g_buffered_input_stream_new : return type :

// Unsupported : g_buffered_input_stream_new_sized : return type :

// Seekable returns the Seekable interface implemented by BufferedInputStream
func (recv *BufferedInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_buffered_output_stream_new : return type :

// Unsupported : g_buffered_output_stream_new_sized : return type :

// Seekable returns the Seekable interface implemented by BufferedOutputStream
func (recv *BufferedOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by BytesIcon
func (recv *BytesIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by BytesIcon
func (recv *BytesIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// Unsupported : g_cancellable_new : return type :

// Converter returns the Converter interface implemented by CharsetConverter
func (recv *CharsetConverter) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Initable returns the Initable interface implemented by CharsetConverter
func (recv *CharsetConverter) Initable() *Initable {
	return InitableNewFromC(recv.ToC())
}

// Unsupported : g_converter_input_stream_new : return type :

// PollableInputStream returns the PollableInputStream interface implemented by ConverterInputStream
func (recv *ConverterInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Unsupported : g_converter_output_stream_new : return type :

// PollableOutputStream returns the PollableOutputStream interface implemented by ConverterOutputStream
func (recv *ConverterOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) ActionGroup() *ActionGroup {
	return ActionGroupNewFromC(recv.ToC())
}

// RemoteActionGroup returns the RemoteActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) RemoteActionGroup() *RemoteActionGroup {
	return RemoteActionGroupNewFromC(recv.ToC())
}

// Unsupported : g_data_input_stream_new : return type :

// Seekable returns the Seekable interface implemented by DataInputStream
func (recv *DataInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_data_output_stream_new : return type :

// Seekable returns the Seekable interface implemented by DataOutputStream
func (recv *DataOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_desktop_app_info_new : return type :

// Unsupported : g_desktop_app_info_new_from_filename : return type :

// AppInfo returns the AppInfo interface implemented by DesktopAppInfo
func (recv *DesktopAppInfo) AppInfo() *AppInfo {
	return AppInfoNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by Emblem
func (recv *Emblem) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by EmblemedIcon
func (recv *EmblemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by FileIOStream
func (recv *FileIOStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_file_icon_new : return type :

// Icon returns the Icon interface implemented by FileIcon
func (recv *FileIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by FileIcon
func (recv *FileIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// Unsupported : g_file_info_new : return type :

// Seekable returns the Seekable interface implemented by FileInputStream
func (recv *FileInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by FileOutputStream
func (recv *FileOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_filename_completer_new : return type :

// Unsupported : g_io_module_new : return type :

// TypePlugin returns the TypePlugin interface implemented by IOModule
func (recv *IOModule) TypePlugin() *gobject.TypePlugin {
	return gobject.TypePluginNewFromC(recv.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by InetSocketAddress
func (recv *InetSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// ListModel returns the ListModel interface implemented by ListStore
func (recv *ListStore) ListModel() *ListModel {
	return ListModelNewFromC(recv.ToC())
}

// Unsupported : g_memory_input_stream_new : return type :

// Unsupported : g_memory_input_stream_new_from_data : unsupported parameter destroy : no type generator for GLib.DestroyNotify (GDestroyNotify) for param destroy

// PollableInputStream returns the PollableInputStream interface implemented by MemoryInputStream
func (recv *MemoryInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryInputStream
func (recv *MemoryInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_memory_output_stream_new : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PollableOutputStream returns the PollableOutputStream interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// Unsupported : g_mount_operation_new : return type :

// SocketConnectable returns the SocketConnectable interface implemented by NetworkAddress
func (recv *NetworkAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkService
func (recv *NetworkService) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by SimpleAction
func (recv *SimpleAction) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AsyncResult returns the AsyncResult interface implemented by SimpleAsyncResult
func (recv *SimpleAsyncResult) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// ProxyResolver returns the ProxyResolver interface implemented by SimpleProxyResolver
func (recv *SimpleProxyResolver) ProxyResolver() *ProxyResolver {
	return ProxyResolverNewFromC(recv.ToC())
}

// SocketConnectable returns the SocketConnectable interface implemented by SocketAddress
func (recv *SocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// AsyncResult returns the AsyncResult interface implemented by Task
func (recv *Task) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// Unsupported : g_themed_icon_new : return type :

// Unsupported : g_themed_icon_new_from_names : return type :

// Unsupported : g_themed_icon_new_with_default_fallbacks : return type :

// Icon returns the Icon interface implemented by ThemedIcon
func (recv *ThemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// Unsupported : g_unix_input_stream_new : return type :

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixInputStream
func (recv *UnixInputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableInputStream returns the PollableInputStream interface implemented by UnixInputStream
func (recv *UnixInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Unsupported : g_unix_mount_monitor_new : return type :

// Unsupported : g_unix_output_stream_new : return type :

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixOutputStream
func (recv *UnixOutputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableOutputStream returns the PollableOutputStream interface implemented by UnixOutputStream
func (recv *UnixOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// Unsupported : g_unix_socket_address_new_abstract : return type :

// SocketConnectable returns the SocketConnectable interface implemented by UnixSocketAddress
func (recv *UnixSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Converter returns the Converter interface implemented by ZlibCompressor
func (recv *ZlibCompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Converter returns the Converter interface implemented by ZlibDecompressor
func (recv *ZlibDecompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}
