// This is a generated file - DO NOT EDIT

package gio

import (
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// AppLaunchContext is a wrapper around the C record GAppLaunchContext.
type AppLaunchContext struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ApplicationCommandLine is a wrapper around the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

// BufferedInputStream is a wrapper around the C record GBufferedInputStream.
type BufferedInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Seekable returns the Seekable interface implemented by BufferedInputStream
func (recv *BufferedInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// BufferedOutputStream is a wrapper around the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Seekable returns the Seekable interface implemented by BufferedOutputStream
func (recv *BufferedOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// BytesIcon is a wrapper around the C record GBytesIcon.
type BytesIcon struct {
	native unsafe.Pointer
}

// Icon returns the Icon interface implemented by BytesIcon
func (recv *BytesIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by BytesIcon
func (recv *BytesIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// Cancellable is a wrapper around the C record GCancellable.
type Cancellable struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// CharsetConverter is a wrapper around the C record GCharsetConverter.
type CharsetConverter struct {
	native unsafe.Pointer
}

// Converter returns the Converter interface implemented by CharsetConverter
func (recv *CharsetConverter) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// Initable returns the Initable interface implemented by CharsetConverter
func (recv *CharsetConverter) Initable() *Initable {
	return InitableNewFromC(recv.ToC())
}

// ConverterInputStream is a wrapper around the C record GConverterInputStream.
type ConverterInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// PollableInputStream returns the PollableInputStream interface implemented by ConverterInputStream
func (recv *ConverterInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// ConverterOutputStream is a wrapper around the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// PollableOutputStream returns the PollableOutputStream interface implemented by ConverterOutputStream
func (recv *ConverterOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// DBusActionGroup is a wrapper around the C record GDBusActionGroup.
type DBusActionGroup struct {
	native unsafe.Pointer
}

// ActionGroup returns the ActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) ActionGroup() *ActionGroup {
	return ActionGroupNewFromC(recv.ToC())
}

// RemoteActionGroup returns the RemoteActionGroup interface implemented by DBusActionGroup
func (recv *DBusActionGroup) RemoteActionGroup() *RemoteActionGroup {
	return RemoteActionGroupNewFromC(recv.ToC())
}

// DBusMenuModel is a wrapper around the C record GDBusMenuModel.
type DBusMenuModel struct {
	native unsafe.Pointer
}

// DataInputStream is a wrapper around the C record GDataInputStream.
type DataInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Seekable returns the Seekable interface implemented by DataInputStream
func (recv *DataInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// DataOutputStream is a wrapper around the C record GDataOutputStream.
type DataOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Seekable returns the Seekable interface implemented by DataOutputStream
func (recv *DataOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// DesktopAppInfo is a wrapper around the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native unsafe.Pointer
}

// AppInfo returns the AppInfo interface implemented by DesktopAppInfo
func (recv *DesktopAppInfo) AppInfo() *AppInfo {
	return AppInfoNewFromC(recv.ToC())
}

// Emblem is a wrapper around the C record GEmblem.
type Emblem struct {
	native unsafe.Pointer
}

// Icon returns the Icon interface implemented by Emblem
func (recv *Emblem) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// EmblemedIcon is a wrapper around the C record GEmblemedIcon.
type EmblemedIcon struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Icon returns the Icon interface implemented by EmblemedIcon
func (recv *EmblemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// FileEnumerator is a wrapper around the C record GFileEnumerator.
type FileEnumerator struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// FileIOStream is a wrapper around the C record GFileIOStream.
type FileIOStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Seekable returns the Seekable interface implemented by FileIOStream
func (recv *FileIOStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FileIcon is a wrapper around the C record GFileIcon.
type FileIcon struct {
	native unsafe.Pointer
}

// Icon returns the Icon interface implemented by FileIcon
func (recv *FileIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by FileIcon
func (recv *FileIcon) LoadableIcon() *LoadableIcon {
	return LoadableIconNewFromC(recv.ToC())
}

// FileInfo is a wrapper around the C record GFileInfo.
type FileInfo struct {
	native unsafe.Pointer
}

// FileInputStream is a wrapper around the C record GFileInputStream.
type FileInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Seekable returns the Seekable interface implemented by FileInputStream
func (recv *FileInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FileMonitor is a wrapper around the C record GFileMonitor.
type FileMonitor struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// FileOutputStream is a wrapper around the C record GFileOutputStream.
type FileOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Seekable returns the Seekable interface implemented by FileOutputStream
func (recv *FileOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FilenameCompleter is a wrapper around the C record GFilenameCompleter.
type FilenameCompleter struct {
	native unsafe.Pointer
}

// FilterInputStream is a wrapper around the C record GFilterInputStream.
type FilterInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// base_stream : record
}

// FilterOutputStream is a wrapper around the C record GFilterOutputStream.
type FilterOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// base_stream : record
}

// IOModule is a wrapper around the C record GIOModule.
type IOModule struct {
	native unsafe.Pointer
}

// TypePlugin returns the TypePlugin interface implemented by IOModule
func (recv *IOModule) TypePlugin() *gobject.TypePlugin {
	return gobject.TypePluginNewFromC(recv.ToC())
}

// IOStream is a wrapper around the C record GIOStream.
type IOStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// InetAddress is a wrapper around the C record GInetAddress.
type InetAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// InetSocketAddress is a wrapper around the C record GInetSocketAddress.
type InetSocketAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// SocketConnectable returns the SocketConnectable interface implemented by InetSocketAddress
func (recv *InetSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// InputStream is a wrapper around the C record GInputStream.
type InputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ListStore is a wrapper around the C record GListStore.
type ListStore struct {
	native unsafe.Pointer
}

// ListModel returns the ListModel interface implemented by ListStore
func (recv *ListStore) ListModel() *ListModel {
	return ListModelNewFromC(recv.ToC())
}

// MemoryInputStream is a wrapper around the C record GMemoryInputStream.
type MemoryInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// PollableInputStream returns the PollableInputStream interface implemented by MemoryInputStream
func (recv *MemoryInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryInputStream
func (recv *MemoryInputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// MemoryOutputStream is a wrapper around the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// PollableOutputStream returns the PollableOutputStream interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// Seekable returns the Seekable interface implemented by MemoryOutputStream
func (recv *MemoryOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// MountOperation is a wrapper around the C record GMountOperation.
type MountOperation struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// NativeVolumeMonitor is a wrapper around the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native unsafe.Pointer
	// parent_instance : record
}

// NetworkAddress is a wrapper around the C record GNetworkAddress.
type NetworkAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkAddress
func (recv *NetworkAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// NetworkService is a wrapper around the C record GNetworkService.
type NetworkService struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// SocketConnectable returns the SocketConnectable interface implemented by NetworkService
func (recv *NetworkService) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// OutputStream is a wrapper around the C record GOutputStream.
type OutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// Permission is a wrapper around the C record GPermission.
type Permission struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ProxyAddressEnumerator is a wrapper around the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Resolver is a wrapper around the C record GResolver.
type Resolver struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Settings is a wrapper around the C record GSettings.
type Settings struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// SettingsBackend is a wrapper around the C record GSettingsBackend.
type SettingsBackend struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// SimpleAction is a wrapper around the C record GSimpleAction.
type SimpleAction struct {
	native unsafe.Pointer
}

// Action returns the Action interface implemented by SimpleAction
func (recv *SimpleAction) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// SimpleAsyncResult is a wrapper around the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native unsafe.Pointer
}

// AsyncResult returns the AsyncResult interface implemented by SimpleAsyncResult
func (recv *SimpleAsyncResult) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// SimplePermission is a wrapper around the C record GSimplePermission.
type SimplePermission struct {
	native unsafe.Pointer
}

// SimpleProxyResolver is a wrapper around the C record GSimpleProxyResolver.
type SimpleProxyResolver struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ProxyResolver returns the ProxyResolver interface implemented by SimpleProxyResolver
func (recv *SimpleProxyResolver) ProxyResolver() *ProxyResolver {
	return ProxyResolverNewFromC(recv.ToC())
}

// SocketAddress is a wrapper around the C record GSocketAddress.
type SocketAddress struct {
	native unsafe.Pointer
	// parent_instance : record
}

// SocketConnectable returns the SocketConnectable interface implemented by SocketAddress
func (recv *SocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// SocketAddressEnumerator is a wrapper around the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native unsafe.Pointer
	// parent_instance : record
}

// SocketControlMessage is a wrapper around the C record GSocketControlMessage.
type SocketControlMessage struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// Task is a wrapper around the C record GTask.
type Task struct {
	native unsafe.Pointer
}

// AsyncResult returns the AsyncResult interface implemented by Task
func (recv *Task) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// TcpWrapperConnection is a wrapper around the C record GTcpWrapperConnection.
type TcpWrapperConnection struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// ThemedIcon is a wrapper around the C record GThemedIcon.
type ThemedIcon struct {
	native unsafe.Pointer
}

// Icon returns the Icon interface implemented by ThemedIcon
func (recv *ThemedIcon) Icon() *Icon {
	return IconNewFromC(recv.ToC())
}

// UnixConnection is a wrapper around the C record GUnixConnection.
type UnixConnection struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// UnixFDList is a wrapper around the C record GUnixFDList.
type UnixFDList struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// UnixFDMessage is a wrapper around the C record GUnixFDMessage.
type UnixFDMessage struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

// UnixInputStream is a wrapper around the C record GUnixInputStream.
type UnixInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixInputStream
func (recv *UnixInputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableInputStream returns the PollableInputStream interface implemented by UnixInputStream
func (recv *UnixInputStream) PollableInputStream() *PollableInputStream {
	return PollableInputStreamNewFromC(recv.ToC())
}

// UnixMountMonitor is a wrapper around the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native unsafe.Pointer
}

// UnixOutputStream is a wrapper around the C record GUnixOutputStream.
type UnixOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// FileDescriptorBased returns the FileDescriptorBased interface implemented by UnixOutputStream
func (recv *UnixOutputStream) FileDescriptorBased() *FileDescriptorBased {
	return FileDescriptorBasedNewFromC(recv.ToC())
}

// PollableOutputStream returns the PollableOutputStream interface implemented by UnixOutputStream
func (recv *UnixOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// UnixSocketAddress is a wrapper around the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// SocketConnectable returns the SocketConnectable interface implemented by UnixSocketAddress
func (recv *UnixSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Vfs is a wrapper around the C record GVfs.
type Vfs struct {
	native unsafe.Pointer
	// parent_instance : record
}

// VolumeMonitor is a wrapper around the C record GVolumeMonitor.
type VolumeMonitor struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

// ZlibCompressor is a wrapper around the C record GZlibCompressor.
type ZlibCompressor struct {
	native unsafe.Pointer
}

// Converter returns the Converter interface implemented by ZlibCompressor
func (recv *ZlibCompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// ZlibDecompressor is a wrapper around the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native unsafe.Pointer
}

// Converter returns the Converter interface implemented by ZlibDecompressor
func (recv *ZlibDecompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}
