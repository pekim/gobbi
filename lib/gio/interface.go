// This is a generated file - DO NOT EDIT

package gio

import "unsafe"

// Action is a wrapper around the C record GAction.
type Action struct {
	native unsafe.Pointer
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	if u == nil {
		return nil
	}

	g := &Action{native: u}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionGroup is a wrapper around the C record GActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	if u == nil {
		return nil
	}

	g := &ActionGroup{native: u}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionMap is a wrapper around the C record GActionMap.
type ActionMap struct {
	native unsafe.Pointer
}

func ActionMapNewFromC(u unsafe.Pointer) *ActionMap {
	if u == nil {
		return nil
	}

	g := &ActionMap{native: u}

	return g
}

func (recv *ActionMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AppInfo is a wrapper around the C record GAppInfo.
type AppInfo struct {
	native unsafe.Pointer
}

func AppInfoNewFromC(u unsafe.Pointer) *AppInfo {
	if u == nil {
		return nil
	}

	g := &AppInfo{native: u}

	return g
}

func (recv *AppInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AsyncResult is a wrapper around the C record GAsyncResult.
type AsyncResult struct {
	native unsafe.Pointer
}

func AsyncResultNewFromC(u unsafe.Pointer) *AsyncResult {
	if u == nil {
		return nil
	}

	g := &AsyncResult{native: u}

	return g
}

func (recv *AsyncResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObject is a wrapper around the C record GDBusObject.
type DBusObject struct {
	native unsafe.Pointer
}

func DBusObjectNewFromC(u unsafe.Pointer) *DBusObject {
	if u == nil {
		return nil
	}

	g := &DBusObject{native: u}

	return g
}

func (recv *DBusObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DBusObjectManager is a wrapper around the C record GDBusObjectManager.
type DBusObjectManager struct {
	native unsafe.Pointer
}

func DBusObjectManagerNewFromC(u unsafe.Pointer) *DBusObjectManager {
	if u == nil {
		return nil
	}

	g := &DBusObjectManager{native: u}

	return g
}

func (recv *DBusObjectManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DesktopAppInfoLookup is a wrapper around the C record GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native unsafe.Pointer
}

func DesktopAppInfoLookupNewFromC(u unsafe.Pointer) *DesktopAppInfoLookup {
	if u == nil {
		return nil
	}

	g := &DesktopAppInfoLookup{native: u}

	return g
}

func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Drive is a wrapper around the C record GDrive.
type Drive struct {
	native unsafe.Pointer
}

func DriveNewFromC(u unsafe.Pointer) *Drive {
	if u == nil {
		return nil
	}

	g := &Drive{native: u}

	return g
}

func (recv *Drive) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// File is a wrapper around the C record GFile.
type File struct {
	native unsafe.Pointer
}

func FileNewFromC(u unsafe.Pointer) *File {
	if u == nil {
		return nil
	}

	g := &File{native: u}

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FileDescriptorBased is a wrapper around the C record GFileDescriptorBased.
type FileDescriptorBased struct {
	native unsafe.Pointer
}

func FileDescriptorBasedNewFromC(u unsafe.Pointer) *FileDescriptorBased {
	if u == nil {
		return nil
	}

	g := &FileDescriptorBased{native: u}

	return g
}

func (recv *FileDescriptorBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Icon is a wrapper around the C record GIcon.
type Icon struct {
	native unsafe.Pointer
}

func IconNewFromC(u unsafe.Pointer) *Icon {
	if u == nil {
		return nil
	}

	g := &Icon{native: u}

	return g
}

func (recv *Icon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ListModel is a wrapper around the C record GListModel.
type ListModel struct {
	native unsafe.Pointer
}

func ListModelNewFromC(u unsafe.Pointer) *ListModel {
	if u == nil {
		return nil
	}

	g := &ListModel{native: u}

	return g
}

func (recv *ListModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// LoadableIcon is a wrapper around the C record GLoadableIcon.
type LoadableIcon struct {
	native unsafe.Pointer
}

func LoadableIconNewFromC(u unsafe.Pointer) *LoadableIcon {
	if u == nil {
		return nil
	}

	g := &LoadableIcon{native: u}

	return g
}

func (recv *LoadableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Mount is a wrapper around the C record GMount.
type Mount struct {
	native unsafe.Pointer
}

func MountNewFromC(u unsafe.Pointer) *Mount {
	if u == nil {
		return nil
	}

	g := &Mount{native: u}

	return g
}

func (recv *Mount) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RemoteActionGroup is a wrapper around the C record GRemoteActionGroup.
type RemoteActionGroup struct {
	native unsafe.Pointer
}

func RemoteActionGroupNewFromC(u unsafe.Pointer) *RemoteActionGroup {
	if u == nil {
		return nil
	}

	g := &RemoteActionGroup{native: u}

	return g
}

func (recv *RemoteActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Seekable is a wrapper around the C record GSeekable.
type Seekable struct {
	native unsafe.Pointer
}

func SeekableNewFromC(u unsafe.Pointer) *Seekable {
	if u == nil {
		return nil
	}

	g := &Seekable{native: u}

	return g
}

func (recv *Seekable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketConnectable is a wrapper around the C record GSocketConnectable.
type SocketConnectable struct {
	native unsafe.Pointer
}

func SocketConnectableNewFromC(u unsafe.Pointer) *SocketConnectable {
	if u == nil {
		return nil
	}

	g := &SocketConnectable{native: u}

	return g
}

func (recv *SocketConnectable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Volume is a wrapper around the C record GVolume.
type Volume struct {
	native unsafe.Pointer
}

func VolumeNewFromC(u unsafe.Pointer) *Volume {
	if u == nil {
		return nil
	}

	g := &Volume{native: u}

	return g
}

func (recv *Volume) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
