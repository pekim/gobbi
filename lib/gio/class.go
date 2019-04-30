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

func AppLaunchContextNewFromC(u unsafe.Pointer) *AppLaunchContext {
	if u == nil {
		return nil
	}

	g := &AppLaunchContext{native: u}

	return g
}

func (recv *AppLaunchContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_app_launch_context_new : return type :

// ApplicationCommandLine is a wrapper around the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func ApplicationCommandLineNewFromC(u unsafe.Pointer) *ApplicationCommandLine {
	if u == nil {
		return nil
	}

	g := &ApplicationCommandLine{native: u}

	return g
}

func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BufferedInputStream is a wrapper around the C record GBufferedInputStream.
type BufferedInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func BufferedInputStreamNewFromC(u unsafe.Pointer) *BufferedInputStream {
	if u == nil {
		return nil
	}

	g := &BufferedInputStream{native: u}

	return g
}

func (recv *BufferedInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_buffered_input_stream_new : return type :

// Unsupported : g_buffered_input_stream_new_sized : return type :

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

func BufferedOutputStreamNewFromC(u unsafe.Pointer) *BufferedOutputStream {
	if u == nil {
		return nil
	}

	g := &BufferedOutputStream{native: u}

	return g
}

func (recv *BufferedOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_buffered_output_stream_new : return type :

// Unsupported : g_buffered_output_stream_new_sized : return type :

// Seekable returns the Seekable interface implemented by BufferedOutputStream
func (recv *BufferedOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// BytesIcon is a wrapper around the C record GBytesIcon.
type BytesIcon struct {
	native unsafe.Pointer
}

func BytesIconNewFromC(u unsafe.Pointer) *BytesIcon {
	if u == nil {
		return nil
	}

	g := &BytesIcon{native: u}

	return g
}

func (recv *BytesIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func CancellableNewFromC(u unsafe.Pointer) *Cancellable {
	if u == nil {
		return nil
	}

	g := &Cancellable{native: u}

	return g
}

func (recv *Cancellable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_cancellable_new : return type :

// CharsetConverter is a wrapper around the C record GCharsetConverter.
type CharsetConverter struct {
	native unsafe.Pointer
}

func CharsetConverterNewFromC(u unsafe.Pointer) *CharsetConverter {
	if u == nil {
		return nil
	}

	g := &CharsetConverter{native: u}

	return g
}

func (recv *CharsetConverter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ConverterInputStreamNewFromC(u unsafe.Pointer) *ConverterInputStream {
	if u == nil {
		return nil
	}

	g := &ConverterInputStream{native: u}

	return g
}

func (recv *ConverterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_converter_input_stream_new : return type :

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

func ConverterOutputStreamNewFromC(u unsafe.Pointer) *ConverterOutputStream {
	if u == nil {
		return nil
	}

	g := &ConverterOutputStream{native: u}

	return g
}

func (recv *ConverterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_converter_output_stream_new : return type :

// PollableOutputStream returns the PollableOutputStream interface implemented by ConverterOutputStream
func (recv *ConverterOutputStream) PollableOutputStream() *PollableOutputStream {
	return PollableOutputStreamNewFromC(recv.ToC())
}

// DBusActionGroup is a wrapper around the C record GDBusActionGroup.
type DBusActionGroup struct {
	native unsafe.Pointer
}

func DBusActionGroupNewFromC(u unsafe.Pointer) *DBusActionGroup {
	if u == nil {
		return nil
	}

	g := &DBusActionGroup{native: u}

	return g
}

func (recv *DBusActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func DBusMenuModelNewFromC(u unsafe.Pointer) *DBusMenuModel {
	if u == nil {
		return nil
	}

	g := &DBusMenuModel{native: u}

	return g
}

func (recv *DBusMenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DataInputStream is a wrapper around the C record GDataInputStream.
type DataInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func DataInputStreamNewFromC(u unsafe.Pointer) *DataInputStream {
	if u == nil {
		return nil
	}

	g := &DataInputStream{native: u}

	return g
}

func (recv *DataInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_data_input_stream_new : return type :

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

func DataOutputStreamNewFromC(u unsafe.Pointer) *DataOutputStream {
	if u == nil {
		return nil
	}

	g := &DataOutputStream{native: u}

	return g
}

func (recv *DataOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_data_output_stream_new : return type :

// Seekable returns the Seekable interface implemented by DataOutputStream
func (recv *DataOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// DesktopAppInfo is a wrapper around the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native unsafe.Pointer
}

func DesktopAppInfoNewFromC(u unsafe.Pointer) *DesktopAppInfo {
	if u == nil {
		return nil
	}

	g := &DesktopAppInfo{native: u}

	return g
}

func (recv *DesktopAppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_desktop_app_info_new : return type :

// Unsupported : g_desktop_app_info_new_from_filename : return type :

// AppInfo returns the AppInfo interface implemented by DesktopAppInfo
func (recv *DesktopAppInfo) AppInfo() *AppInfo {
	return AppInfoNewFromC(recv.ToC())
}

// Emblem is a wrapper around the C record GEmblem.
type Emblem struct {
	native unsafe.Pointer
}

func EmblemNewFromC(u unsafe.Pointer) *Emblem {
	if u == nil {
		return nil
	}

	g := &Emblem{native: u}

	return g
}

func (recv *Emblem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func EmblemedIconNewFromC(u unsafe.Pointer) *EmblemedIcon {
	if u == nil {
		return nil
	}

	g := &EmblemedIcon{native: u}

	return g
}

func (recv *EmblemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FileEnumeratorNewFromC(u unsafe.Pointer) *FileEnumerator {
	if u == nil {
		return nil
	}

	g := &FileEnumerator{native: u}

	return g
}

func (recv *FileEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileIOStream is a wrapper around the C record GFileIOStream.
type FileIOStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func FileIOStreamNewFromC(u unsafe.Pointer) *FileIOStream {
	if u == nil {
		return nil
	}

	g := &FileIOStream{native: u}

	return g
}

func (recv *FileIOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Seekable returns the Seekable interface implemented by FileIOStream
func (recv *FileIOStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FileIcon is a wrapper around the C record GFileIcon.
type FileIcon struct {
	native unsafe.Pointer
}

func FileIconNewFromC(u unsafe.Pointer) *FileIcon {
	if u == nil {
		return nil
	}

	g := &FileIcon{native: u}

	return g
}

func (recv *FileIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// FileInfo is a wrapper around the C record GFileInfo.
type FileInfo struct {
	native unsafe.Pointer
}

func FileInfoNewFromC(u unsafe.Pointer) *FileInfo {
	if u == nil {
		return nil
	}

	g := &FileInfo{native: u}

	return g
}

func (recv *FileInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_file_info_new : return type :

// FileInputStream is a wrapper around the C record GFileInputStream.
type FileInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func FileInputStreamNewFromC(u unsafe.Pointer) *FileInputStream {
	if u == nil {
		return nil
	}

	g := &FileInputStream{native: u}

	return g
}

func (recv *FileInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func FileMonitorNewFromC(u unsafe.Pointer) *FileMonitor {
	if u == nil {
		return nil
	}

	g := &FileMonitor{native: u}

	return g
}

func (recv *FileMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileOutputStream is a wrapper around the C record GFileOutputStream.
type FileOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func FileOutputStreamNewFromC(u unsafe.Pointer) *FileOutputStream {
	if u == nil {
		return nil
	}

	g := &FileOutputStream{native: u}

	return g
}

func (recv *FileOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Seekable returns the Seekable interface implemented by FileOutputStream
func (recv *FileOutputStream) Seekable() *Seekable {
	return SeekableNewFromC(recv.ToC())
}

// FilenameCompleter is a wrapper around the C record GFilenameCompleter.
type FilenameCompleter struct {
	native unsafe.Pointer
}

func FilenameCompleterNewFromC(u unsafe.Pointer) *FilenameCompleter {
	if u == nil {
		return nil
	}

	g := &FilenameCompleter{native: u}

	return g
}

func (recv *FilenameCompleter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_filename_completer_new : return type :

// FilterInputStream is a wrapper around the C record GFilterInputStream.
type FilterInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// base_stream : record
}

func FilterInputStreamNewFromC(u unsafe.Pointer) *FilterInputStream {
	if u == nil {
		return nil
	}

	g := &FilterInputStream{native: u}

	return g
}

func (recv *FilterInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FilterOutputStream is a wrapper around the C record GFilterOutputStream.
type FilterOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// base_stream : record
}

func FilterOutputStreamNewFromC(u unsafe.Pointer) *FilterOutputStream {
	if u == nil {
		return nil
	}

	g := &FilterOutputStream{native: u}

	return g
}

func (recv *FilterOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOModule is a wrapper around the C record GIOModule.
type IOModule struct {
	native unsafe.Pointer
}

func IOModuleNewFromC(u unsafe.Pointer) *IOModule {
	if u == nil {
		return nil
	}

	g := &IOModule{native: u}

	return g
}

func (recv *IOModule) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_io_module_new : return type :

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

func IOStreamNewFromC(u unsafe.Pointer) *IOStream {
	if u == nil {
		return nil
	}

	g := &IOStream{native: u}

	return g
}

func (recv *IOStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetAddress is a wrapper around the C record GInetAddress.
type InetAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func InetAddressNewFromC(u unsafe.Pointer) *InetAddress {
	if u == nil {
		return nil
	}

	g := &InetAddress{native: u}

	return g
}

func (recv *InetAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InetSocketAddress is a wrapper around the C record GInetSocketAddress.
type InetSocketAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func InetSocketAddressNewFromC(u unsafe.Pointer) *InetSocketAddress {
	if u == nil {
		return nil
	}

	g := &InetSocketAddress{native: u}

	return g
}

func (recv *InetSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func InputStreamNewFromC(u unsafe.Pointer) *InputStream {
	if u == nil {
		return nil
	}

	g := &InputStream{native: u}

	return g
}

func (recv *InputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListStore is a wrapper around the C record GListStore.
type ListStore struct {
	native unsafe.Pointer
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	if u == nil {
		return nil
	}

	g := &ListStore{native: u}

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MemoryInputStreamNewFromC(u unsafe.Pointer) *MemoryInputStream {
	if u == nil {
		return nil
	}

	g := &MemoryInputStream{native: u}

	return g
}

func (recv *MemoryInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// MemoryOutputStream is a wrapper around the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func MemoryOutputStreamNewFromC(u unsafe.Pointer) *MemoryOutputStream {
	if u == nil {
		return nil
	}

	g := &MemoryOutputStream{native: u}

	return g
}

func (recv *MemoryOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// MountOperation is a wrapper around the C record GMountOperation.
type MountOperation struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	if u == nil {
		return nil
	}

	g := &MountOperation{native: u}

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_mount_operation_new : return type :

// NativeVolumeMonitor is a wrapper around the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native unsafe.Pointer
	// parent_instance : record
}

func NativeVolumeMonitorNewFromC(u unsafe.Pointer) *NativeVolumeMonitor {
	if u == nil {
		return nil
	}

	g := &NativeVolumeMonitor{native: u}

	return g
}

func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NetworkAddress is a wrapper around the C record GNetworkAddress.
type NetworkAddress struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func NetworkAddressNewFromC(u unsafe.Pointer) *NetworkAddress {
	if u == nil {
		return nil
	}

	g := &NetworkAddress{native: u}

	return g
}

func (recv *NetworkAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func NetworkServiceNewFromC(u unsafe.Pointer) *NetworkService {
	if u == nil {
		return nil
	}

	g := &NetworkService{native: u}

	return g
}

func (recv *NetworkService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func OutputStreamNewFromC(u unsafe.Pointer) *OutputStream {
	if u == nil {
		return nil
	}

	g := &OutputStream{native: u}

	return g
}

func (recv *OutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Permission is a wrapper around the C record GPermission.
type Permission struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func PermissionNewFromC(u unsafe.Pointer) *Permission {
	if u == nil {
		return nil
	}

	g := &Permission{native: u}

	return g
}

func (recv *Permission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ProxyAddressEnumerator is a wrapper around the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func ProxyAddressEnumeratorNewFromC(u unsafe.Pointer) *ProxyAddressEnumerator {
	if u == nil {
		return nil
	}

	g := &ProxyAddressEnumerator{native: u}

	return g
}

func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Resolver is a wrapper around the C record GResolver.
type Resolver struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func ResolverNewFromC(u unsafe.Pointer) *Resolver {
	if u == nil {
		return nil
	}

	g := &Resolver{native: u}

	return g
}

func (recv *Resolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Settings is a wrapper around the C record GSettings.
type Settings struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	if u == nil {
		return nil
	}

	g := &Settings{native: u}

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SettingsBackend is a wrapper around the C record GSettingsBackend.
type SettingsBackend struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func SettingsBackendNewFromC(u unsafe.Pointer) *SettingsBackend {
	if u == nil {
		return nil
	}

	g := &SettingsBackend{native: u}

	return g
}

func (recv *SettingsBackend) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleAction is a wrapper around the C record GSimpleAction.
type SimpleAction struct {
	native unsafe.Pointer
}

func SimpleActionNewFromC(u unsafe.Pointer) *SimpleAction {
	if u == nil {
		return nil
	}

	g := &SimpleAction{native: u}

	return g
}

func (recv *SimpleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action returns the Action interface implemented by SimpleAction
func (recv *SimpleAction) Action() *Action {
	return ActionNewFromC(recv.ToC())
}

// SimpleAsyncResult is a wrapper around the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native unsafe.Pointer
}

func SimpleAsyncResultNewFromC(u unsafe.Pointer) *SimpleAsyncResult {
	if u == nil {
		return nil
	}

	g := &SimpleAsyncResult{native: u}

	return g
}

func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_simple_async_result_new : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_simple_async_result_new_from_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AsyncResult returns the AsyncResult interface implemented by SimpleAsyncResult
func (recv *SimpleAsyncResult) AsyncResult() *AsyncResult {
	return AsyncResultNewFromC(recv.ToC())
}

// SimplePermission is a wrapper around the C record GSimplePermission.
type SimplePermission struct {
	native unsafe.Pointer
}

func SimplePermissionNewFromC(u unsafe.Pointer) *SimplePermission {
	if u == nil {
		return nil
	}

	g := &SimplePermission{native: u}

	return g
}

func (recv *SimplePermission) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SimpleProxyResolver is a wrapper around the C record GSimpleProxyResolver.
type SimpleProxyResolver struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func SimpleProxyResolverNewFromC(u unsafe.Pointer) *SimpleProxyResolver {
	if u == nil {
		return nil
	}

	g := &SimpleProxyResolver{native: u}

	return g
}

func (recv *SimpleProxyResolver) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SocketAddressNewFromC(u unsafe.Pointer) *SocketAddress {
	if u == nil {
		return nil
	}

	g := &SocketAddress{native: u}

	return g
}

func (recv *SocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SocketAddressEnumeratorNewFromC(u unsafe.Pointer) *SocketAddressEnumerator {
	if u == nil {
		return nil
	}

	g := &SocketAddressEnumerator{native: u}

	return g
}

func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketControlMessage is a wrapper around the C record GSocketControlMessage.
type SocketControlMessage struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func SocketControlMessageNewFromC(u unsafe.Pointer) *SocketControlMessage {
	if u == nil {
		return nil
	}

	g := &SocketControlMessage{native: u}

	return g
}

func (recv *SocketControlMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Task is a wrapper around the C record GTask.
type Task struct {
	native unsafe.Pointer
}

func TaskNewFromC(u unsafe.Pointer) *Task {
	if u == nil {
		return nil
	}

	g := &Task{native: u}

	return g
}

func (recv *Task) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TcpWrapperConnectionNewFromC(u unsafe.Pointer) *TcpWrapperConnection {
	if u == nil {
		return nil
	}

	g := &TcpWrapperConnection{native: u}

	return g
}

func (recv *TcpWrapperConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThemedIcon is a wrapper around the C record GThemedIcon.
type ThemedIcon struct {
	native unsafe.Pointer
}

func ThemedIconNewFromC(u unsafe.Pointer) *ThemedIcon {
	if u == nil {
		return nil
	}

	g := &ThemedIcon{native: u}

	return g
}

func (recv *ThemedIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_themed_icon_new : return type :

// Unsupported : g_themed_icon_new_from_names : return type :

// Unsupported : g_themed_icon_new_with_default_fallbacks : return type :

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

func UnixConnectionNewFromC(u unsafe.Pointer) *UnixConnection {
	if u == nil {
		return nil
	}

	g := &UnixConnection{native: u}

	return g
}

func (recv *UnixConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDList is a wrapper around the C record GUnixFDList.
type UnixFDList struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func UnixFDListNewFromC(u unsafe.Pointer) *UnixFDList {
	if u == nil {
		return nil
	}

	g := &UnixFDList{native: u}

	return g
}

func (recv *UnixFDList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixFDMessage is a wrapper around the C record GUnixFDMessage.
type UnixFDMessage struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func UnixFDMessageNewFromC(u unsafe.Pointer) *UnixFDMessage {
	if u == nil {
		return nil
	}

	g := &UnixFDMessage{native: u}

	return g
}

func (recv *UnixFDMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// UnixInputStream is a wrapper around the C record GUnixInputStream.
type UnixInputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func UnixInputStreamNewFromC(u unsafe.Pointer) *UnixInputStream {
	if u == nil {
		return nil
	}

	g := &UnixInputStream{native: u}

	return g
}

func (recv *UnixInputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

// UnixMountMonitor is a wrapper around the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native unsafe.Pointer
}

func UnixMountMonitorNewFromC(u unsafe.Pointer) *UnixMountMonitor {
	if u == nil {
		return nil
	}

	g := &UnixMountMonitor{native: u}

	return g
}

func (recv *UnixMountMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_unix_mount_monitor_new : return type :

// UnixOutputStream is a wrapper around the C record GUnixOutputStream.
type UnixOutputStream struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func UnixOutputStreamNewFromC(u unsafe.Pointer) *UnixOutputStream {
	if u == nil {
		return nil
	}

	g := &UnixOutputStream{native: u}

	return g
}

func (recv *UnixOutputStream) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_unix_output_stream_new : return type :

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

func UnixSocketAddressNewFromC(u unsafe.Pointer) *UnixSocketAddress {
	if u == nil {
		return nil
	}

	g := &UnixSocketAddress{native: u}

	return g
}

func (recv *UnixSocketAddress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_unix_socket_address_new_abstract : return type :

// SocketConnectable returns the SocketConnectable interface implemented by UnixSocketAddress
func (recv *UnixSocketAddress) SocketConnectable() *SocketConnectable {
	return SocketConnectableNewFromC(recv.ToC())
}

// Vfs is a wrapper around the C record GVfs.
type Vfs struct {
	native unsafe.Pointer
	// parent_instance : record
}

func VfsNewFromC(u unsafe.Pointer) *Vfs {
	if u == nil {
		return nil
	}

	g := &Vfs{native: u}

	return g
}

func (recv *Vfs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VolumeMonitor is a wrapper around the C record GVolumeMonitor.
type VolumeMonitor struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func VolumeMonitorNewFromC(u unsafe.Pointer) *VolumeMonitor {
	if u == nil {
		return nil
	}

	g := &VolumeMonitor{native: u}

	return g
}

func (recv *VolumeMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ZlibCompressor is a wrapper around the C record GZlibCompressor.
type ZlibCompressor struct {
	native unsafe.Pointer
}

func ZlibCompressorNewFromC(u unsafe.Pointer) *ZlibCompressor {
	if u == nil {
		return nil
	}

	g := &ZlibCompressor{native: u}

	return g
}

func (recv *ZlibCompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Converter returns the Converter interface implemented by ZlibCompressor
func (recv *ZlibCompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}

// ZlibDecompressor is a wrapper around the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native unsafe.Pointer
}

func ZlibDecompressorNewFromC(u unsafe.Pointer) *ZlibDecompressor {
	if u == nil {
		return nil
	}

	g := &ZlibDecompressor{native: u}

	return g
}

func (recv *ZlibDecompressor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Converter returns the Converter interface implemented by ZlibDecompressor
func (recv *ZlibDecompressor) Converter() *Converter {
	return ConverterNewFromC(recv.ToC())
}
