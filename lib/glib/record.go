// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var ArrayStruct *gi.Struct
var ArrayStructOnce sync.Once

func ArrayStructSet() {
	ArrayStructOnce.Do(func() {
		ArrayStruct = gi.StructNew("GLib", "Array")
	})
}

type Array struct {
	native uintptr
	Data   string
	Len    uint32
}

var AsyncQueueStruct *gi.Struct
var AsyncQueueStructOnce sync.Once

func AsyncQueueStructSet() {
	AsyncQueueStructOnce.Do(func() {
		AsyncQueueStruct = gi.StructNew("GLib", "AsyncQueue")
	})
}

type AsyncQueue struct {
	native uintptr
}

var lengthAsyncQueueInvoker *gi.Function

// Length is a representation of the C type g_async_queue_length.
func (recv *AsyncQueue) Length() int32 {
	if lengthAsyncQueueInvoker == nil {
		lengthAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := lengthAsyncQueueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var lengthUnlockedAsyncQueueInvoker *gi.Function

// LengthUnlocked is a representation of the C type g_async_queue_length_unlocked.
func (recv *AsyncQueue) LengthUnlocked() int32 {
	if lengthUnlockedAsyncQueueInvoker == nil {
		lengthUnlockedAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "length_unlocked")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := lengthUnlockedAsyncQueueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var lockAsyncQueueInvoker *gi.Function

// Lock is a representation of the C type g_async_queue_lock.
func (recv *AsyncQueue) Lock() {
	if lockAsyncQueueInvoker == nil {
		lockAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "lock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	lockAsyncQueueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_async_queue_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_pop_unlocked' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_front' : parameter 'item' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_front_unlocked' : parameter 'item' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_sorted' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_sorted_unlocked' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_unlocked' : parameter 'data' of type 'gpointer' not supported

var refAsyncQueueInvoker *gi.Function

// Ref is a representation of the C type g_async_queue_ref.
func (recv *AsyncQueue) Ref() *AsyncQueue {
	if refAsyncQueueInvoker == nil {
		refAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refAsyncQueueInvoker.Invoke(inArgs[:], nil)

	retGo := &AsyncQueue{native: ret.Pointer()}

	return retGo
}

var refUnlockedAsyncQueueInvoker *gi.Function

// RefUnlocked is a representation of the C type g_async_queue_ref_unlocked.
func (recv *AsyncQueue) RefUnlocked() {
	if refUnlockedAsyncQueueInvoker == nil {
		refUnlockedAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "ref_unlocked")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	refUnlockedAsyncQueueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_async_queue_remove' : parameter 'item' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_remove_unlocked' : parameter 'item' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_sort' : parameter 'func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_async_queue_sort_unlocked' : parameter 'func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_async_queue_timed_pop' : parameter 'end_time' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_async_queue_timed_pop_unlocked' : parameter 'end_time' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_async_queue_timeout_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_timeout_pop_unlocked' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_try_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_try_pop_unlocked' : return type 'gpointer' not supported

var unlockAsyncQueueInvoker *gi.Function

// Unlock is a representation of the C type g_async_queue_unlock.
func (recv *AsyncQueue) Unlock() {
	if unlockAsyncQueueInvoker == nil {
		unlockAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "unlock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unlockAsyncQueueInvoker.Invoke(inArgs[:], nil)

}

var unrefAsyncQueueInvoker *gi.Function

// Unref is a representation of the C type g_async_queue_unref.
func (recv *AsyncQueue) Unref() {
	if unrefAsyncQueueInvoker == nil {
		unrefAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefAsyncQueueInvoker.Invoke(inArgs[:], nil)

}

var unrefAndUnlockAsyncQueueInvoker *gi.Function

// UnrefAndUnlock is a representation of the C type g_async_queue_unref_and_unlock.
func (recv *AsyncQueue) UnrefAndUnlock() {
	if unrefAndUnlockAsyncQueueInvoker == nil {
		unrefAndUnlockAsyncQueueInvoker = gi.StructFunctionInvokerNew("GLib", "AsyncQueue", "unref_and_unlock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefAndUnlockAsyncQueueInvoker.Invoke(inArgs[:], nil)

}

var BookmarkFileStruct *gi.Struct
var BookmarkFileStructOnce sync.Once

func BookmarkFileStructSet() {
	BookmarkFileStructOnce.Do(func() {
		BookmarkFileStruct = gi.StructNew("GLib", "BookmarkFile")
	})
}

type BookmarkFile struct {
	native uintptr
}

var addApplicationBookmarkFileInvoker *gi.Function

// AddApplication is a representation of the C type g_bookmark_file_add_application.
func (recv *BookmarkFile) AddApplication(uri string, name string, exec string) {
	if addApplicationBookmarkFileInvoker == nil {
		addApplicationBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "add_application")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)
	inArgs[3].SetString(exec)

	addApplicationBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

var addGroupBookmarkFileInvoker *gi.Function

// AddGroup is a representation of the C type g_bookmark_file_add_group.
func (recv *BookmarkFile) AddGroup(uri string, group string) {
	if addGroupBookmarkFileInvoker == nil {
		addGroupBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "add_group")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(group)

	addGroupBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

var freeBookmarkFileInvoker *gi.Function

// Free is a representation of the C type g_bookmark_file_free.
func (recv *BookmarkFile) Free() {
	if freeBookmarkFileInvoker == nil {
		freeBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

var getAddedBookmarkFileInvoker *gi.Function

// GetAdded is a representation of the C type g_bookmark_file_get_added.
func (recv *BookmarkFile) GetAdded(uri string) int64 {
	if getAddedBookmarkFileInvoker == nil {
		getAddedBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_added")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := getAddedBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_get_app_info' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_get_applications' : parameter 'length' of type 'gsize' not supported

var getDescriptionBookmarkFileInvoker *gi.Function

// GetDescription is a representation of the C type g_bookmark_file_get_description.
func (recv *BookmarkFile) GetDescription(uri string) string {
	if getDescriptionBookmarkFileInvoker == nil {
		getDescriptionBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_description")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := getDescriptionBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_get_groups' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_bookmark_file_get_icon' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_get_is_private' : return type 'gboolean' not supported

var getMimeTypeBookmarkFileInvoker *gi.Function

// GetMimeType is a representation of the C type g_bookmark_file_get_mime_type.
func (recv *BookmarkFile) GetMimeType(uri string) string {
	if getMimeTypeBookmarkFileInvoker == nil {
		getMimeTypeBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_mime_type")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := getMimeTypeBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getModifiedBookmarkFileInvoker *gi.Function

// GetModified is a representation of the C type g_bookmark_file_get_modified.
func (recv *BookmarkFile) GetModified(uri string) int64 {
	if getModifiedBookmarkFileInvoker == nil {
		getModifiedBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_modified")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := getModifiedBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getSizeBookmarkFileInvoker *gi.Function

// GetSize is a representation of the C type g_bookmark_file_get_size.
func (recv *BookmarkFile) GetSize() int32 {
	if getSizeBookmarkFileInvoker == nil {
		getSizeBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_size")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSizeBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getTitleBookmarkFileInvoker *gi.Function

// GetTitle is a representation of the C type g_bookmark_file_get_title.
func (recv *BookmarkFile) GetTitle(uri string) string {
	if getTitleBookmarkFileInvoker == nil {
		getTitleBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_title")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := getTitleBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_get_uris' : parameter 'length' of type 'gsize' not supported

var getVisitedBookmarkFileInvoker *gi.Function

// GetVisited is a representation of the C type g_bookmark_file_get_visited.
func (recv *BookmarkFile) GetVisited(uri string) int64 {
	if getVisitedBookmarkFileInvoker == nil {
		getVisitedBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "get_visited")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := getVisitedBookmarkFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_has_application' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_has_group' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_has_item' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_load_from_data' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bookmark_file_load_from_data_dirs' : parameter 'file' of type 'filename' not supported

// UNSUPPORTED : C value 'g_bookmark_file_load_from_file' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_bookmark_file_move_item' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_remove_application' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_remove_group' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_remove_item' : return type 'gboolean' not supported

var setAddedBookmarkFileInvoker *gi.Function

// SetAdded is a representation of the C type g_bookmark_file_set_added.
func (recv *BookmarkFile) SetAdded(uri string, added int64) {
	if setAddedBookmarkFileInvoker == nil {
		setAddedBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_added")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(added)

	setAddedBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_set_app_info' : return type 'gboolean' not supported

var setDescriptionBookmarkFileInvoker *gi.Function

// SetDescription is a representation of the C type g_bookmark_file_set_description.
func (recv *BookmarkFile) SetDescription(uri string, description string) {
	if setDescriptionBookmarkFileInvoker == nil {
		setDescriptionBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_description")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(description)

	setDescriptionBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_set_groups' : parameter 'groups' has no type

var setIconBookmarkFileInvoker *gi.Function

// SetIcon is a representation of the C type g_bookmark_file_set_icon.
func (recv *BookmarkFile) SetIcon(uri string, href string, mimeType string) {
	if setIconBookmarkFileInvoker == nil {
		setIconBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_icon")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(href)
	inArgs[3].SetString(mimeType)

	setIconBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_set_is_private' : parameter 'is_private' of type 'gboolean' not supported

var setMimeTypeBookmarkFileInvoker *gi.Function

// SetMimeType is a representation of the C type g_bookmark_file_set_mime_type.
func (recv *BookmarkFile) SetMimeType(uri string, mimeType string) {
	if setMimeTypeBookmarkFileInvoker == nil {
		setMimeTypeBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_mime_type")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(mimeType)

	setMimeTypeBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

var setModifiedBookmarkFileInvoker *gi.Function

// SetModified is a representation of the C type g_bookmark_file_set_modified.
func (recv *BookmarkFile) SetModified(uri string, modified int64) {
	if setModifiedBookmarkFileInvoker == nil {
		setModifiedBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_modified")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(modified)

	setModifiedBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

var setTitleBookmarkFileInvoker *gi.Function

// SetTitle is a representation of the C type g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	if setTitleBookmarkFileInvoker == nil {
		setTitleBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_title")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(title)

	setTitleBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

var setVisitedBookmarkFileInvoker *gi.Function

// SetVisited is a representation of the C type g_bookmark_file_set_visited.
func (recv *BookmarkFile) SetVisited(uri string, visited int64) {
	if setVisitedBookmarkFileInvoker == nil {
		setVisitedBookmarkFileInvoker = gi.StructFunctionInvokerNew("GLib", "BookmarkFile", "set_visited")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(visited)

	setVisitedBookmarkFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_to_data' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_bookmark_file_to_file' : parameter 'filename' of type 'filename' not supported

var ByteArrayStruct *gi.Struct
var ByteArrayStructOnce sync.Once

func ByteArrayStructSet() {
	ByteArrayStructOnce.Do(func() {
		ByteArrayStruct = gi.StructNew("GLib", "ByteArray")
	})
}

type ByteArray struct {
	native uintptr
	Data   uint8
	Len    uint32
}

var BytesStruct *gi.Struct
var BytesStructOnce sync.Once

func BytesStructSet() {
	BytesStructOnce.Do(func() {
		BytesStruct = gi.StructNew("GLib", "Bytes")
	})
}

type Bytes struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_bytes_new' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_new_static' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_new_with_free_func' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_compare' : parameter 'bytes2' of type 'Bytes' not supported

// UNSUPPORTED : C value 'g_bytes_equal' : parameter 'bytes2' of type 'Bytes' not supported

// UNSUPPORTED : C value 'g_bytes_get_data' : parameter 'size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_bytes_get_size' : return type 'gsize' not supported

var hashBytesInvoker *gi.Function

// Hash is a representation of the C type g_bytes_hash.
func (recv *Bytes) Hash() uint32 {
	if hashBytesInvoker == nil {
		hashBytesInvoker = gi.StructFunctionInvokerNew("GLib", "Bytes", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashBytesInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_bytes_new_from_bytes' : parameter 'offset' of type 'gsize' not supported

var refBytesInvoker *gi.Function

// Ref is a representation of the C type g_bytes_ref.
func (recv *Bytes) Ref() *Bytes {
	if refBytesInvoker == nil {
		refBytesInvoker = gi.StructFunctionInvokerNew("GLib", "Bytes", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refBytesInvoker.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

var unrefBytesInvoker *gi.Function

// Unref is a representation of the C type g_bytes_unref.
func (recv *Bytes) Unref() {
	if unrefBytesInvoker == nil {
		unrefBytesInvoker = gi.StructFunctionInvokerNew("GLib", "Bytes", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefBytesInvoker.Invoke(inArgs[:], nil)

}

var unrefToArrayBytesInvoker *gi.Function

// UnrefToArray is a representation of the C type g_bytes_unref_to_array.
func (recv *Bytes) UnrefToArray() {
	if unrefToArrayBytesInvoker == nil {
		unrefToArrayBytesInvoker = gi.StructFunctionInvokerNew("GLib", "Bytes", "unref_to_array")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefToArrayBytesInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bytes_unref_to_data' : parameter 'size' of type 'gsize' not supported

var ChecksumStruct *gi.Struct
var ChecksumStructOnce sync.Once

func ChecksumStructSet() {
	ChecksumStructOnce.Do(func() {
		ChecksumStruct = gi.StructNew("GLib", "Checksum")
	})
}

type Checksum struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_checksum_new' : parameter 'checksum_type' of type 'ChecksumType' not supported

var copyChecksumInvoker *gi.Function

// Copy is a representation of the C type g_checksum_copy.
func (recv *Checksum) Copy() *Checksum {
	if copyChecksumInvoker == nil {
		copyChecksumInvoker = gi.StructFunctionInvokerNew("GLib", "Checksum", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyChecksumInvoker.Invoke(inArgs[:], nil)

	retGo := &Checksum{native: ret.Pointer()}

	return retGo
}

var freeChecksumInvoker *gi.Function

// Free is a representation of the C type g_checksum_free.
func (recv *Checksum) Free() {
	if freeChecksumInvoker == nil {
		freeChecksumInvoker = gi.StructFunctionInvokerNew("GLib", "Checksum", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeChecksumInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_checksum_get_digest' : parameter 'buffer' has no type

var getStringChecksumInvoker *gi.Function

// GetString is a representation of the C type g_checksum_get_string.
func (recv *Checksum) GetString() string {
	if getStringChecksumInvoker == nil {
		getStringChecksumInvoker = gi.StructFunctionInvokerNew("GLib", "Checksum", "get_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStringChecksumInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var resetChecksumInvoker *gi.Function

// Reset is a representation of the C type g_checksum_reset.
func (recv *Checksum) Reset() {
	if resetChecksumInvoker == nil {
		resetChecksumInvoker = gi.StructFunctionInvokerNew("GLib", "Checksum", "reset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	resetChecksumInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_checksum_update' : parameter 'data' has no type

var CondStruct *gi.Struct
var CondStructOnce sync.Once

func CondStructSet() {
	CondStructOnce.Do(func() {
		CondStruct = gi.StructNew("GLib", "Cond")
	})
}

type Cond struct {
	native uintptr
}

var broadcastCondInvoker *gi.Function

// Broadcast is a representation of the C type g_cond_broadcast.
func (recv *Cond) Broadcast() {
	if broadcastCondInvoker == nil {
		broadcastCondInvoker = gi.StructFunctionInvokerNew("GLib", "Cond", "broadcast")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	broadcastCondInvoker.Invoke(inArgs[:], nil)

}

var clearCondInvoker *gi.Function

// Clear is a representation of the C type g_cond_clear.
func (recv *Cond) Clear() {
	if clearCondInvoker == nil {
		clearCondInvoker = gi.StructFunctionInvokerNew("GLib", "Cond", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearCondInvoker.Invoke(inArgs[:], nil)

}

var initCondInvoker *gi.Function

// Init is a representation of the C type g_cond_init.
func (recv *Cond) Init() {
	if initCondInvoker == nil {
		initCondInvoker = gi.StructFunctionInvokerNew("GLib", "Cond", "init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	initCondInvoker.Invoke(inArgs[:], nil)

}

var signalCondInvoker *gi.Function

// Signal is a representation of the C type g_cond_signal.
func (recv *Cond) Signal() {
	if signalCondInvoker == nil {
		signalCondInvoker = gi.StructFunctionInvokerNew("GLib", "Cond", "signal")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	signalCondInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_cond_wait' : parameter 'mutex' of type 'Mutex' not supported

// UNSUPPORTED : C value 'g_cond_wait_until' : parameter 'mutex' of type 'Mutex' not supported

var DataStruct *gi.Struct
var DataStructOnce sync.Once

func DataStructSet() {
	DataStructOnce.Do(func() {
		DataStruct = gi.StructNew("GLib", "Data")
	})
}

type Data struct {
	native uintptr
}

var DateStruct *gi.Struct
var DateStructOnce sync.Once

func DateStructSet() {
	DateStructOnce.Do(func() {
		DateStruct = gi.StructNew("GLib", "Date")
	})
}

type Date struct {
	native     uintptr
	JulianDays uint32
	Julian     uint32
	Dmy        uint32
	Day        uint32
	Month      uint32
	Year       uint32
}

var newDateInvoker *gi.Function

// DateNew is a representation of the C type g_date_new.
func DateNew() *Date {
	if newDateInvoker == nil {
		newDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "new")
	}

	ret := newDateInvoker.Invoke(nil, nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_new_dmy' : parameter 'day' of type 'DateDay' not supported

var newJulianDateInvoker *gi.Function

// DateNewJulian is a representation of the C type g_date_new_julian.
func DateNewJulian(julianDay uint32) *Date {
	if newJulianDateInvoker == nil {
		newJulianDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "new_julian")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(julianDay)

	ret := newJulianDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var addDaysDateInvoker *gi.Function

// AddDays is a representation of the C type g_date_add_days.
func (recv *Date) AddDays(nDays uint32) {
	if addDaysDateInvoker == nil {
		addDaysDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "add_days")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDays)

	addDaysDateInvoker.Invoke(inArgs[:], nil)

}

var addMonthsDateInvoker *gi.Function

// AddMonths is a representation of the C type g_date_add_months.
func (recv *Date) AddMonths(nMonths uint32) {
	if addMonthsDateInvoker == nil {
		addMonthsDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "add_months")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nMonths)

	addMonthsDateInvoker.Invoke(inArgs[:], nil)

}

var addYearsDateInvoker *gi.Function

// AddYears is a representation of the C type g_date_add_years.
func (recv *Date) AddYears(nYears uint32) {
	if addYearsDateInvoker == nil {
		addYearsDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "add_years")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nYears)

	addYearsDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_clamp' : parameter 'min_date' of type 'Date' not supported

var clearDateInvoker *gi.Function

// Clear is a representation of the C type g_date_clear.
func (recv *Date) Clear(nDates uint32) {
	if clearDateInvoker == nil {
		clearDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "clear")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDates)

	clearDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_compare' : parameter 'rhs' of type 'Date' not supported

var copyDateInvoker *gi.Function

// Copy is a representation of the C type g_date_copy.
func (recv *Date) Copy() *Date {
	if copyDateInvoker == nil {
		copyDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyDateInvoker.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_days_between' : parameter 'date2' of type 'Date' not supported

var freeDateInvoker *gi.Function

// Free is a representation of the C type g_date_free.
func (recv *Date) Free() {
	if freeDateInvoker == nil {
		freeDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_get_day' : return type 'DateDay' not supported

var getDayOfYearDateInvoker *gi.Function

// GetDayOfYear is a representation of the C type g_date_get_day_of_year.
func (recv *Date) GetDayOfYear() uint32 {
	if getDayOfYearDateInvoker == nil {
		getDayOfYearDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "get_day_of_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDayOfYearDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getIso8601WeekOfYearDateInvoker *gi.Function

// GetIso8601WeekOfYear is a representation of the C type g_date_get_iso8601_week_of_year.
func (recv *Date) GetIso8601WeekOfYear() uint32 {
	if getIso8601WeekOfYearDateInvoker == nil {
		getIso8601WeekOfYearDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "get_iso8601_week_of_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIso8601WeekOfYearDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getJulianDateInvoker *gi.Function

// GetJulian is a representation of the C type g_date_get_julian.
func (recv *Date) GetJulian() uint32 {
	if getJulianDateInvoker == nil {
		getJulianDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "get_julian")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getJulianDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getMondayWeekOfYearDateInvoker *gi.Function

// GetMondayWeekOfYear is a representation of the C type g_date_get_monday_week_of_year.
func (recv *Date) GetMondayWeekOfYear() uint32 {
	if getMondayWeekOfYearDateInvoker == nil {
		getMondayWeekOfYearDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "get_monday_week_of_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMondayWeekOfYearDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_date_get_month' : return type 'DateMonth' not supported

var getSundayWeekOfYearDateInvoker *gi.Function

// GetSundayWeekOfYear is a representation of the C type g_date_get_sunday_week_of_year.
func (recv *Date) GetSundayWeekOfYear() uint32 {
	if getSundayWeekOfYearDateInvoker == nil {
		getSundayWeekOfYearDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "get_sunday_week_of_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSundayWeekOfYearDateInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_date_get_weekday' : return type 'DateWeekday' not supported

// UNSUPPORTED : C value 'g_date_get_year' : return type 'DateYear' not supported

// UNSUPPORTED : C value 'g_date_is_first_of_month' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_date_is_last_of_month' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_date_order' : parameter 'date2' of type 'Date' not supported

// UNSUPPORTED : C value 'g_date_set_day' : parameter 'day' of type 'DateDay' not supported

// UNSUPPORTED : C value 'g_date_set_dmy' : parameter 'day' of type 'DateDay' not supported

var setJulianDateInvoker *gi.Function

// SetJulian is a representation of the C type g_date_set_julian.
func (recv *Date) SetJulian(julianDate uint32) {
	if setJulianDateInvoker == nil {
		setJulianDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "set_julian")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(julianDate)

	setJulianDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_set_month' : parameter 'month' of type 'DateMonth' not supported

var setParseDateInvoker *gi.Function

// SetParse is a representation of the C type g_date_set_parse.
func (recv *Date) SetParse(str string) {
	if setParseDateInvoker == nil {
		setParseDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "set_parse")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(str)

	setParseDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_set_time' : parameter 'time_' of type 'Time' not supported

var setTimeTDateInvoker *gi.Function

// SetTimeT is a representation of the C type g_date_set_time_t.
func (recv *Date) SetTimeT(timet int64) {
	if setTimeTDateInvoker == nil {
		setTimeTDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "set_time_t")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(timet)

	setTimeTDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_set_time_val' : parameter 'timeval' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_set_year' : parameter 'year' of type 'DateYear' not supported

var subtractDaysDateInvoker *gi.Function

// SubtractDays is a representation of the C type g_date_subtract_days.
func (recv *Date) SubtractDays(nDays uint32) {
	if subtractDaysDateInvoker == nil {
		subtractDaysDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "subtract_days")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDays)

	subtractDaysDateInvoker.Invoke(inArgs[:], nil)

}

var subtractMonthsDateInvoker *gi.Function

// SubtractMonths is a representation of the C type g_date_subtract_months.
func (recv *Date) SubtractMonths(nMonths uint32) {
	if subtractMonthsDateInvoker == nil {
		subtractMonthsDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "subtract_months")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nMonths)

	subtractMonthsDateInvoker.Invoke(inArgs[:], nil)

}

var subtractYearsDateInvoker *gi.Function

// SubtractYears is a representation of the C type g_date_subtract_years.
func (recv *Date) SubtractYears(nYears uint32) {
	if subtractYearsDateInvoker == nil {
		subtractYearsDateInvoker = gi.StructFunctionInvokerNew("GLib", "Date", "subtract_years")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nYears)

	subtractYearsDateInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_to_struct_tm' : parameter 'tm' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_valid' : return type 'gboolean' not supported

var DateTimeStruct *gi.Struct
var DateTimeStructOnce sync.Once

func DateTimeStructSet() {
	DateTimeStructOnce.Do(func() {
		DateTimeStruct = gi.StructNew("GLib", "DateTime")
	})
}

type DateTime struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_date_time_new' : parameter 'tz' of type 'TimeZone' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_iso8601' : parameter 'default_tz' of type 'TimeZone' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_timeval_local' : parameter 'tv' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_timeval_utc' : parameter 'tv' of type 'TimeVal' not supported

var newFromUnixLocalDateTimeInvoker *gi.Function

// DateTimeNewFromUnixLocal is a representation of the C type g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	if newFromUnixLocalDateTimeInvoker == nil {
		newFromUnixLocalDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "new_from_unix_local")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	ret := newFromUnixLocalDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var newFromUnixUtcDateTimeInvoker *gi.Function

// DateTimeNewFromUnixUtc is a representation of the C type g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	if newFromUnixUtcDateTimeInvoker == nil {
		newFromUnixUtcDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "new_from_unix_utc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	ret := newFromUnixUtcDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_new_local' : parameter 'seconds' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_date_time_new_now' : parameter 'tz' of type 'TimeZone' not supported

var newNowLocalDateTimeInvoker *gi.Function

// DateTimeNewNowLocal is a representation of the C type g_date_time_new_now_local.
func DateTimeNewNowLocal() *DateTime {
	if newNowLocalDateTimeInvoker == nil {
		newNowLocalDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "new_now_local")
	}

	ret := newNowLocalDateTimeInvoker.Invoke(nil, nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var newNowUtcDateTimeInvoker *gi.Function

// DateTimeNewNowUtc is a representation of the C type g_date_time_new_now_utc.
func DateTimeNewNowUtc() *DateTime {
	if newNowUtcDateTimeInvoker == nil {
		newNowUtcDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "new_now_utc")
	}

	ret := newNowUtcDateTimeInvoker.Invoke(nil, nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_new_utc' : parameter 'seconds' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_date_time_add' : parameter 'timespan' of type 'TimeSpan' not supported

var addDaysDateTimeInvoker *gi.Function

// AddDays is a representation of the C type g_date_time_add_days.
func (recv *DateTime) AddDays(days int32) *DateTime {
	if addDaysDateTimeInvoker == nil {
		addDaysDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "add_days")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(days)

	ret := addDaysDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_add_full' : parameter 'seconds' of type 'gdouble' not supported

var addHoursDateTimeInvoker *gi.Function

// AddHours is a representation of the C type g_date_time_add_hours.
func (recv *DateTime) AddHours(hours int32) *DateTime {
	if addHoursDateTimeInvoker == nil {
		addHoursDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "add_hours")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(hours)

	ret := addHoursDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var addMinutesDateTimeInvoker *gi.Function

// AddMinutes is a representation of the C type g_date_time_add_minutes.
func (recv *DateTime) AddMinutes(minutes int32) *DateTime {
	if addMinutesDateTimeInvoker == nil {
		addMinutesDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "add_minutes")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(minutes)

	ret := addMinutesDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var addMonthsDateTimeInvoker *gi.Function

// AddMonths is a representation of the C type g_date_time_add_months.
func (recv *DateTime) AddMonths(months int32) *DateTime {
	if addMonthsDateTimeInvoker == nil {
		addMonthsDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "add_months")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(months)

	ret := addMonthsDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_add_seconds' : parameter 'seconds' of type 'gdouble' not supported

var addWeeksDateTimeInvoker *gi.Function

// AddWeeks is a representation of the C type g_date_time_add_weeks.
func (recv *DateTime) AddWeeks(weeks int32) *DateTime {
	if addWeeksDateTimeInvoker == nil {
		addWeeksDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "add_weeks")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(weeks)

	ret := addWeeksDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var addYearsDateTimeInvoker *gi.Function

// AddYears is a representation of the C type g_date_time_add_years.
func (recv *DateTime) AddYears(years int32) *DateTime {
	if addYearsDateTimeInvoker == nil {
		addYearsDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "add_years")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(years)

	ret := addYearsDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_difference' : parameter 'begin' of type 'DateTime' not supported

var formatDateTimeInvoker *gi.Function

// Format is a representation of the C type g_date_time_format.
func (recv *DateTime) Format(format string) string {
	if formatDateTimeInvoker == nil {
		formatDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "format")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(format)

	ret := formatDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var formatIso8601DateTimeInvoker *gi.Function

// FormatIso8601 is a representation of the C type g_date_time_format_iso8601.
func (recv *DateTime) FormatIso8601() string {
	if formatIso8601DateTimeInvoker == nil {
		formatIso8601DateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "format_iso8601")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := formatIso8601DateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getDayOfMonthDateTimeInvoker *gi.Function

// GetDayOfMonth is a representation of the C type g_date_time_get_day_of_month.
func (recv *DateTime) GetDayOfMonth() int32 {
	if getDayOfMonthDateTimeInvoker == nil {
		getDayOfMonthDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_day_of_month")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDayOfMonthDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getDayOfWeekDateTimeInvoker *gi.Function

// GetDayOfWeek is a representation of the C type g_date_time_get_day_of_week.
func (recv *DateTime) GetDayOfWeek() int32 {
	if getDayOfWeekDateTimeInvoker == nil {
		getDayOfWeekDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_day_of_week")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDayOfWeekDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getDayOfYearDateTimeInvoker *gi.Function

// GetDayOfYear is a representation of the C type g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() int32 {
	if getDayOfYearDateTimeInvoker == nil {
		getDayOfYearDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_day_of_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDayOfYearDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getHourDateTimeInvoker *gi.Function

// GetHour is a representation of the C type g_date_time_get_hour.
func (recv *DateTime) GetHour() int32 {
	if getHourDateTimeInvoker == nil {
		getHourDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_hour")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHourDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getMicrosecondDateTimeInvoker *gi.Function

// GetMicrosecond is a representation of the C type g_date_time_get_microsecond.
func (recv *DateTime) GetMicrosecond() int32 {
	if getMicrosecondDateTimeInvoker == nil {
		getMicrosecondDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_microsecond")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMicrosecondDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getMinuteDateTimeInvoker *gi.Function

// GetMinute is a representation of the C type g_date_time_get_minute.
func (recv *DateTime) GetMinute() int32 {
	if getMinuteDateTimeInvoker == nil {
		getMinuteDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_minute")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMinuteDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getMonthDateTimeInvoker *gi.Function

// GetMonth is a representation of the C type g_date_time_get_month.
func (recv *DateTime) GetMonth() int32 {
	if getMonthDateTimeInvoker == nil {
		getMonthDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_month")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMonthDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getSecondDateTimeInvoker *gi.Function

// GetSecond is a representation of the C type g_date_time_get_second.
func (recv *DateTime) GetSecond() int32 {
	if getSecondDateTimeInvoker == nil {
		getSecondDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_second")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSecondDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_get_seconds' : return type 'gdouble' not supported

var getTimezoneDateTimeInvoker *gi.Function

// GetTimezone is a representation of the C type g_date_time_get_timezone.
func (recv *DateTime) GetTimezone() *TimeZone {
	if getTimezoneDateTimeInvoker == nil {
		getTimezoneDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_timezone")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTimezoneDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var getTimezoneAbbreviationDateTimeInvoker *gi.Function

// GetTimezoneAbbreviation is a representation of the C type g_date_time_get_timezone_abbreviation.
func (recv *DateTime) GetTimezoneAbbreviation() string {
	if getTimezoneAbbreviationDateTimeInvoker == nil {
		getTimezoneAbbreviationDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_timezone_abbreviation")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTimezoneAbbreviationDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_get_utc_offset' : return type 'TimeSpan' not supported

var getWeekNumberingYearDateTimeInvoker *gi.Function

// GetWeekNumberingYear is a representation of the C type g_date_time_get_week_numbering_year.
func (recv *DateTime) GetWeekNumberingYear() int32 {
	if getWeekNumberingYearDateTimeInvoker == nil {
		getWeekNumberingYearDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_week_numbering_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getWeekNumberingYearDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getWeekOfYearDateTimeInvoker *gi.Function

// GetWeekOfYear is a representation of the C type g_date_time_get_week_of_year.
func (recv *DateTime) GetWeekOfYear() int32 {
	if getWeekOfYearDateTimeInvoker == nil {
		getWeekOfYearDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_week_of_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getWeekOfYearDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getYearDateTimeInvoker *gi.Function

// GetYear is a representation of the C type g_date_time_get_year.
func (recv *DateTime) GetYear() int32 {
	if getYearDateTimeInvoker == nil {
		getYearDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_year")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getYearDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getYmdDateTimeInvoker *gi.Function

// GetYmd is a representation of the C type g_date_time_get_ymd.
func (recv *DateTime) GetYmd() (int32, int32, int32) {
	if getYmdDateTimeInvoker == nil {
		getYmdDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "get_ymd")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	getYmdDateTimeInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2
}

// UNSUPPORTED : C value 'g_date_time_is_daylight_savings' : return type 'gboolean' not supported

var refDateTimeInvoker *gi.Function

// Ref is a representation of the C type g_date_time_ref.
func (recv *DateTime) Ref() *DateTime {
	if refDateTimeInvoker == nil {
		refDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var toLocalDateTimeInvoker *gi.Function

// ToLocal is a representation of the C type g_date_time_to_local.
func (recv *DateTime) ToLocal() *DateTime {
	if toLocalDateTimeInvoker == nil {
		toLocalDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "to_local")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toLocalDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_to_timeval' : parameter 'tv' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_time_to_timezone' : parameter 'tz' of type 'TimeZone' not supported

var toUnixDateTimeInvoker *gi.Function

// ToUnix is a representation of the C type g_date_time_to_unix.
func (recv *DateTime) ToUnix() int64 {
	if toUnixDateTimeInvoker == nil {
		toUnixDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "to_unix")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toUnixDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var toUtcDateTimeInvoker *gi.Function

// ToUtc is a representation of the C type g_date_time_to_utc.
func (recv *DateTime) ToUtc() *DateTime {
	if toUtcDateTimeInvoker == nil {
		toUtcDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "to_utc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toUtcDateTimeInvoker.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var unrefDateTimeInvoker *gi.Function

// Unref is a representation of the C type g_date_time_unref.
func (recv *DateTime) Unref() {
	if unrefDateTimeInvoker == nil {
		unrefDateTimeInvoker = gi.StructFunctionInvokerNew("GLib", "DateTime", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefDateTimeInvoker.Invoke(inArgs[:], nil)

}

var DebugKeyStruct *gi.Struct
var DebugKeyStructOnce sync.Once

func DebugKeyStructSet() {
	DebugKeyStructOnce.Do(func() {
		DebugKeyStruct = gi.StructNew("GLib", "DebugKey")
	})
}

type DebugKey struct {
	native uintptr
	Key    string
	Value  uint32
}

var DirStruct *gi.Struct
var DirStructOnce sync.Once

func DirStructSet() {
	DirStructOnce.Do(func() {
		DirStruct = gi.StructNew("GLib", "Dir")
	})
}

type Dir struct {
	native uintptr
}

var closeDirInvoker *gi.Function

// Close is a representation of the C type g_dir_close.
func (recv *Dir) Close() {
	if closeDirInvoker == nil {
		closeDirInvoker = gi.StructFunctionInvokerNew("GLib", "Dir", "close")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closeDirInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_dir_read_name' : return type 'filename' not supported

var rewindDirInvoker *gi.Function

// Rewind is a representation of the C type g_dir_rewind.
func (recv *Dir) Rewind() {
	if rewindDirInvoker == nil {
		rewindDirInvoker = gi.StructFunctionInvokerNew("GLib", "Dir", "rewind")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rewindDirInvoker.Invoke(inArgs[:], nil)

}

var ErrorStruct *gi.Struct
var ErrorStructOnce sync.Once

func ErrorStructSet() {
	ErrorStructOnce.Do(func() {
		ErrorStruct = gi.StructNew("GLib", "Error")
	})
}

type Error struct {
	native  uintptr
	Domain  Quark
	Code    int32
	Message string
}

// UNSUPPORTED : C value 'g_error_new' : parameter 'domain' of type 'Quark' not supported

// UNSUPPORTED : C value 'g_error_new_literal' : parameter 'domain' of type 'Quark' not supported

// UNSUPPORTED : C value 'g_error_new_valist' : parameter 'domain' of type 'Quark' not supported

var copyErrorInvoker *gi.Function

// Copy is a representation of the C type g_error_copy.
func (recv *Error) Copy() *Error {
	if copyErrorInvoker == nil {
		copyErrorInvoker = gi.StructFunctionInvokerNew("GLib", "Error", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyErrorInvoker.Invoke(inArgs[:], nil)

	retGo := &Error{native: ret.Pointer()}

	return retGo
}

var freeErrorInvoker *gi.Function

// Free is a representation of the C type g_error_free.
func (recv *Error) Free() {
	if freeErrorInvoker == nil {
		freeErrorInvoker = gi.StructFunctionInvokerNew("GLib", "Error", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeErrorInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_error_matches' : parameter 'domain' of type 'Quark' not supported

var HashTableStruct *gi.Struct
var HashTableStructOnce sync.Once

func HashTableStructSet() {
	HashTableStructOnce.Do(func() {
		HashTableStruct = gi.StructNew("GLib", "HashTable")
	})
}

type HashTable struct {
	native uintptr
}

var HashTableIterStruct *gi.Struct
var HashTableIterStructOnce sync.Once

func HashTableIterStructSet() {
	HashTableIterStructOnce.Do(func() {
		HashTableIterStruct = gi.StructNew("GLib", "HashTableIter")
	})
}

type HashTableIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_hash_table_iter_get_hash_table' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_iter_init' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_iter_next' : parameter 'key' of type 'gpointer' not supported

var removeHashTableIterInvoker *gi.Function

// Remove is a representation of the C type g_hash_table_iter_remove.
func (recv *HashTableIter) Remove() {
	if removeHashTableIterInvoker == nil {
		removeHashTableIterInvoker = gi.StructFunctionInvokerNew("GLib", "HashTableIter", "remove")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	removeHashTableIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_hash_table_iter_replace' : parameter 'value' of type 'gpointer' not supported

var stealHashTableIterInvoker *gi.Function

// Steal is a representation of the C type g_hash_table_iter_steal.
func (recv *HashTableIter) Steal() {
	if stealHashTableIterInvoker == nil {
		stealHashTableIterInvoker = gi.StructFunctionInvokerNew("GLib", "HashTableIter", "steal")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stealHashTableIterInvoker.Invoke(inArgs[:], nil)

}

var HmacStruct *gi.Struct
var HmacStructOnce sync.Once

func HmacStructSet() {
	HmacStructOnce.Do(func() {
		HmacStruct = gi.StructNew("GLib", "Hmac")
	})
}

type Hmac struct {
	native uintptr
}

var copyHmacInvoker *gi.Function

// Copy is a representation of the C type g_hmac_copy.
func (recv *Hmac) Copy() *Hmac {
	if copyHmacInvoker == nil {
		copyHmacInvoker = gi.StructFunctionInvokerNew("GLib", "Hmac", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyHmacInvoker.Invoke(inArgs[:], nil)

	retGo := &Hmac{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_hmac_get_digest' : parameter 'buffer' has no type

var getStringHmacInvoker *gi.Function

// GetString is a representation of the C type g_hmac_get_string.
func (recv *Hmac) GetString() string {
	if getStringHmacInvoker == nil {
		getStringHmacInvoker = gi.StructFunctionInvokerNew("GLib", "Hmac", "get_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStringHmacInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var refHmacInvoker *gi.Function

// Ref is a representation of the C type g_hmac_ref.
func (recv *Hmac) Ref() *Hmac {
	if refHmacInvoker == nil {
		refHmacInvoker = gi.StructFunctionInvokerNew("GLib", "Hmac", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refHmacInvoker.Invoke(inArgs[:], nil)

	retGo := &Hmac{native: ret.Pointer()}

	return retGo
}

var unrefHmacInvoker *gi.Function

// Unref is a representation of the C type g_hmac_unref.
func (recv *Hmac) Unref() {
	if unrefHmacInvoker == nil {
		unrefHmacInvoker = gi.StructFunctionInvokerNew("GLib", "Hmac", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefHmacInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_hmac_update' : parameter 'data' has no type

var HookStruct *gi.Struct
var HookStructOnce sync.Once

func HookStructSet() {
	HookStructOnce.Do(func() {
		HookStruct = gi.StructNew("GLib", "Hook")
	})
}

type Hook struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	Next     *Hook
	Prev     *Hook
	RefCount uint32
	HookId   uint64
	Flags    uint32
	// UNSUPPORTED : C value 'func' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'destroy' : no Go type for 'DestroyNotify'
}

// UNSUPPORTED : C value 'g_hook_compare_ids' : parameter 'sibling' of type 'Hook' not supported

var HookListStruct *gi.Struct
var HookListStructOnce sync.Once

func HookListStructSet() {
	HookListStructOnce.Do(func() {
		HookListStruct = gi.StructNew("GLib", "HookList")
	})
}

type HookList struct {
	native   uintptr
	SeqId    uint64
	HookSize uint32
	IsSetup  uint32
	Hooks    *Hook
	// UNSUPPORTED : C value 'dummy3' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'finalize_hook' : no Go type for 'HookFinalizeFunc'
	// UNSUPPORTED : C value 'dummy' : missing Type
}

var clearHookListInvoker *gi.Function

// Clear is a representation of the C type g_hook_list_clear.
func (recv *HookList) Clear() {
	if clearHookListInvoker == nil {
		clearHookListInvoker = gi.StructFunctionInvokerNew("GLib", "HookList", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearHookListInvoker.Invoke(inArgs[:], nil)

}

var initHookListInvoker *gi.Function

// Init is a representation of the C type g_hook_list_init.
func (recv *HookList) Init(hookSize uint32) {
	if initHookListInvoker == nil {
		initHookListInvoker = gi.StructFunctionInvokerNew("GLib", "HookList", "init")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(hookSize)

	initHookListInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_hook_list_invoke' : parameter 'may_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hook_list_invoke_check' : parameter 'may_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hook_list_marshal' : parameter 'may_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hook_list_marshal_check' : parameter 'may_recurse' of type 'gboolean' not supported

var IConvStruct *gi.Struct
var IConvStructOnce sync.Once

func IConvStructSet() {
	IConvStructOnce.Do(func() {
		IConvStruct = gi.StructNew("GLib", "IConv")
	})
}

type IConv struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'inbytes_left' of type 'gsize' not supported

var closeIConvInvoker *gi.Function

// Close is a representation of the C type g_iconv_close.
func (recv *IConv) Close() int32 {
	if closeIConvInvoker == nil {
		closeIConvInvoker = gi.StructFunctionInvokerNew("GLib", "IConv", "close")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := closeIConvInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var IOChannelStruct *gi.Struct
var IOChannelStructOnce sync.Once

func IOChannelStructSet() {
	IOChannelStructOnce.Do(func() {
		IOChannelStruct = gi.StructNew("GLib", "IOChannel")
	})
}

type IOChannel struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_channel_new_file' : parameter 'filename' of type 'filename' not supported

var unixNewIOChannelInvoker *gi.Function

// IOChannelUnixNew is a representation of the C type g_io_channel_unix_new.
func IOChannelUnixNew(fd int32) *IOChannel {
	if unixNewIOChannelInvoker == nil {
		unixNewIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "unix_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	ret := unixNewIOChannelInvoker.Invoke(inArgs[:], nil)

	retGo := &IOChannel{native: ret.Pointer()}

	return retGo
}

var closeIOChannelInvoker *gi.Function

// Close is a representation of the C type g_io_channel_close.
func (recv *IOChannel) Close() {
	if closeIOChannelInvoker == nil {
		closeIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "close")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closeIOChannelInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_flush' : return type 'IOStatus' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffer_condition' : return type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffer_size' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffered' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_io_channel_get_close_on_unref' : return type 'gboolean' not supported

var getEncodingIOChannelInvoker *gi.Function

// GetEncoding is a representation of the C type g_io_channel_get_encoding.
func (recv *IOChannel) GetEncoding() string {
	if getEncodingIOChannelInvoker == nil {
		getEncodingIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "get_encoding")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getEncodingIOChannelInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_io_channel_get_flags' : return type 'IOFlags' not supported

var getLineTermIOChannelInvoker *gi.Function

// GetLineTerm is a representation of the C type g_io_channel_get_line_term.
func (recv *IOChannel) GetLineTerm(length int32) string {
	if getLineTermIOChannelInvoker == nil {
		getLineTermIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "get_line_term")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(length)

	ret := getLineTermIOChannelInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var initIOChannelInvoker *gi.Function

// Init is a representation of the C type g_io_channel_init.
func (recv *IOChannel) Init() {
	if initIOChannelInvoker == nil {
		initIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	initIOChannelInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_read' : parameter 'count' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_read_chars' : parameter 'buf' has no type

// UNSUPPORTED : C value 'g_io_channel_read_line' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_read_line_string' : parameter 'buffer' of type 'String' not supported

// UNSUPPORTED : C value 'g_io_channel_read_to_end' : parameter 'str_return' has no type

// UNSUPPORTED : C value 'g_io_channel_read_unichar' : parameter 'thechar' of type 'gunichar' not supported

var refIOChannelInvoker *gi.Function

// Ref is a representation of the C type g_io_channel_ref.
func (recv *IOChannel) Ref() *IOChannel {
	if refIOChannelInvoker == nil {
		refIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refIOChannelInvoker.Invoke(inArgs[:], nil)

	retGo := &IOChannel{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_io_channel_seek' : parameter 'type' of type 'SeekType' not supported

// UNSUPPORTED : C value 'g_io_channel_seek_position' : parameter 'type' of type 'SeekType' not supported

// UNSUPPORTED : C value 'g_io_channel_set_buffer_size' : parameter 'size' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_set_buffered' : parameter 'buffered' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_io_channel_set_close_on_unref' : parameter 'do_close' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_io_channel_set_encoding' : return type 'IOStatus' not supported

// UNSUPPORTED : C value 'g_io_channel_set_flags' : parameter 'flags' of type 'IOFlags' not supported

var setLineTermIOChannelInvoker *gi.Function

// SetLineTerm is a representation of the C type g_io_channel_set_line_term.
func (recv *IOChannel) SetLineTerm(lineTerm string, length int32) {
	if setLineTermIOChannelInvoker == nil {
		setLineTermIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "set_line_term")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(lineTerm)
	inArgs[2].SetInt32(length)

	setLineTermIOChannelInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_shutdown' : parameter 'flush' of type 'gboolean' not supported

var unixGetFdIOChannelInvoker *gi.Function

// UnixGetFd is a representation of the C type g_io_channel_unix_get_fd.
func (recv *IOChannel) UnixGetFd() int32 {
	if unixGetFdIOChannelInvoker == nil {
		unixGetFdIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "unix_get_fd")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := unixGetFdIOChannelInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var unrefIOChannelInvoker *gi.Function

// Unref is a representation of the C type g_io_channel_unref.
func (recv *IOChannel) Unref() {
	if unrefIOChannelInvoker == nil {
		unrefIOChannelInvoker = gi.StructFunctionInvokerNew("GLib", "IOChannel", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefIOChannelInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_write' : parameter 'count' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_write_chars' : parameter 'buf' has no type

// UNSUPPORTED : C value 'g_io_channel_write_unichar' : parameter 'thechar' of type 'gunichar' not supported

var IOFuncsStruct *gi.Struct
var IOFuncsStructOnce sync.Once

func IOFuncsStructSet() {
	IOFuncsStructOnce.Do(func() {
		IOFuncsStruct = gi.StructNew("GLib", "IOFuncs")
	})
}

type IOFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'io_read' : missing Type
	// UNSUPPORTED : C value 'io_write' : missing Type
	// UNSUPPORTED : C value 'io_seek' : missing Type
	// UNSUPPORTED : C value 'io_close' : missing Type
	// UNSUPPORTED : C value 'io_create_watch' : missing Type
	// UNSUPPORTED : C value 'io_free' : missing Type
	// UNSUPPORTED : C value 'io_set_flags' : missing Type
	// UNSUPPORTED : C value 'io_get_flags' : missing Type
}

var KeyFileStruct *gi.Struct
var KeyFileStructOnce sync.Once

func KeyFileStructSet() {
	KeyFileStructOnce.Do(func() {
		KeyFileStruct = gi.StructNew("GLib", "KeyFile")
	})
}

type KeyFile struct {
	native uintptr
}

var newKeyFileInvoker *gi.Function

// KeyFileNew is a representation of the C type g_key_file_new.
func KeyFileNew() *KeyFile {
	if newKeyFileInvoker == nil {
		newKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "new")
	}

	ret := newKeyFileInvoker.Invoke(nil, nil)

	retGo := &KeyFile{native: ret.Pointer()}

	return retGo
}

var freeKeyFileInvoker *gi.Function

// Free is a representation of the C type g_key_file_free.
func (recv *KeyFile) Free() {
	if freeKeyFileInvoker == nil {
		freeKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeKeyFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_get_boolean' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_get_boolean_list' : parameter 'length' of type 'gsize' not supported

var getCommentKeyFileInvoker *gi.Function

// GetComment is a representation of the C type g_key_file_get_comment.
func (recv *KeyFile) GetComment(groupName string, key string) string {
	if getCommentKeyFileInvoker == nil {
		getCommentKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_comment")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := getCommentKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_key_file_get_double_list' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_get_groups' : parameter 'length' of type 'gsize' not supported

var getInt64KeyFileInvoker *gi.Function

// GetInt64 is a representation of the C type g_key_file_get_int64.
func (recv *KeyFile) GetInt64(groupName string, key string) int64 {
	if getInt64KeyFileInvoker == nil {
		getInt64KeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_int64")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := getInt64KeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getIntegerKeyFileInvoker *gi.Function

// GetInteger is a representation of the C type g_key_file_get_integer.
func (recv *KeyFile) GetInteger(groupName string, key string) int32 {
	if getIntegerKeyFileInvoker == nil {
		getIntegerKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_integer")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := getIntegerKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_integer_list' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_get_keys' : parameter 'length' of type 'gsize' not supported

var getLocaleForKeyKeyFileInvoker *gi.Function

// GetLocaleForKey is a representation of the C type g_key_file_get_locale_for_key.
func (recv *KeyFile) GetLocaleForKey(groupName string, key string, locale string) string {
	if getLocaleForKeyKeyFileInvoker == nil {
		getLocaleForKeyKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_locale_for_key")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)

	ret := getLocaleForKeyKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getLocaleStringKeyFileInvoker *gi.Function

// GetLocaleString is a representation of the C type g_key_file_get_locale_string.
func (recv *KeyFile) GetLocaleString(groupName string, key string, locale string) string {
	if getLocaleStringKeyFileInvoker == nil {
		getLocaleStringKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_locale_string")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)

	ret := getLocaleStringKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_locale_string_list' : parameter 'length' of type 'gsize' not supported

var getStartGroupKeyFileInvoker *gi.Function

// GetStartGroup is a representation of the C type g_key_file_get_start_group.
func (recv *KeyFile) GetStartGroup() string {
	if getStartGroupKeyFileInvoker == nil {
		getStartGroupKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_start_group")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStartGroupKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getStringKeyFileInvoker *gi.Function

// GetString is a representation of the C type g_key_file_get_string.
func (recv *KeyFile) GetString(groupName string, key string) string {
	if getStringKeyFileInvoker == nil {
		getStringKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_string")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := getStringKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_string_list' : parameter 'length' of type 'gsize' not supported

var getUint64KeyFileInvoker *gi.Function

// GetUint64 is a representation of the C type g_key_file_get_uint64.
func (recv *KeyFile) GetUint64(groupName string, key string) uint64 {
	if getUint64KeyFileInvoker == nil {
		getUint64KeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_uint64")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := getUint64KeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

var getValueKeyFileInvoker *gi.Function

// GetValue is a representation of the C type g_key_file_get_value.
func (recv *KeyFile) GetValue(groupName string, key string) string {
	if getValueKeyFileInvoker == nil {
		getValueKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "get_value")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := getValueKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_has_group' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_has_key' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_bytes' : parameter 'bytes' of type 'Bytes' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_data' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_data_dirs' : parameter 'file' of type 'filename' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_dirs' : parameter 'file' of type 'filename' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_file' : parameter 'file' of type 'filename' not supported

var refKeyFileInvoker *gi.Function

// Ref is a representation of the C type g_key_file_ref.
func (recv *KeyFile) Ref() *KeyFile {
	if refKeyFileInvoker == nil {
		refKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refKeyFileInvoker.Invoke(inArgs[:], nil)

	retGo := &KeyFile{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_remove_comment' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_remove_group' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_remove_key' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_save_to_file' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_set_boolean' : parameter 'value' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_set_boolean_list' : parameter 'list' has no type

// UNSUPPORTED : C value 'g_key_file_set_comment' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_set_double' : parameter 'value' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_key_file_set_double_list' : parameter 'list' has no type

var setInt64KeyFileInvoker *gi.Function

// SetInt64 is a representation of the C type g_key_file_set_int64.
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) {
	if setInt64KeyFileInvoker == nil {
		setInt64KeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_int64")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetInt64(value)

	setInt64KeyFileInvoker.Invoke(inArgs[:], nil)

}

var setIntegerKeyFileInvoker *gi.Function

// SetInteger is a representation of the C type g_key_file_set_integer.
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) {
	if setIntegerKeyFileInvoker == nil {
		setIntegerKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_integer")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetInt32(value)

	setIntegerKeyFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_set_integer_list' : parameter 'list' has no type

var setListSeparatorKeyFileInvoker *gi.Function

// SetListSeparator is a representation of the C type g_key_file_set_list_separator.
func (recv *KeyFile) SetListSeparator(separator int8) {
	if setListSeparatorKeyFileInvoker == nil {
		setListSeparatorKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_list_separator")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(separator)

	setListSeparatorKeyFileInvoker.Invoke(inArgs[:], nil)

}

var setLocaleStringKeyFileInvoker *gi.Function

// SetLocaleString is a representation of the C type g_key_file_set_locale_string.
func (recv *KeyFile) SetLocaleString(groupName string, key string, locale string, string_ string) {
	if setLocaleStringKeyFileInvoker == nil {
		setLocaleStringKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_locale_string")
	}

	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)
	inArgs[4].SetString(string_)

	setLocaleStringKeyFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_set_locale_string_list' : parameter 'list' has no type

var setStringKeyFileInvoker *gi.Function

// SetString is a representation of the C type g_key_file_set_string.
func (recv *KeyFile) SetString(groupName string, key string, string_ string) {
	if setStringKeyFileInvoker == nil {
		setStringKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_string")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(string_)

	setStringKeyFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_set_string_list' : parameter 'list' has no type

var setUint64KeyFileInvoker *gi.Function

// SetUint64 is a representation of the C type g_key_file_set_uint64.
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) {
	if setUint64KeyFileInvoker == nil {
		setUint64KeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_uint64")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetUint64(value)

	setUint64KeyFileInvoker.Invoke(inArgs[:], nil)

}

var setValueKeyFileInvoker *gi.Function

// SetValue is a representation of the C type g_key_file_set_value.
func (recv *KeyFile) SetValue(groupName string, key string, value string) {
	if setValueKeyFileInvoker == nil {
		setValueKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "set_value")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(value)

	setValueKeyFileInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_to_data' : parameter 'length' of type 'gsize' not supported

var unrefKeyFileInvoker *gi.Function

// Unref is a representation of the C type g_key_file_unref.
func (recv *KeyFile) Unref() {
	if unrefKeyFileInvoker == nil {
		unrefKeyFileInvoker = gi.StructFunctionInvokerNew("GLib", "KeyFile", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefKeyFileInvoker.Invoke(inArgs[:], nil)

}

var ListStruct *gi.Struct
var ListStructOnce sync.Once

func ListStructSet() {
	ListStructOnce.Do(func() {
		ListStruct = gi.StructNew("GLib", "List")
	})
}

type List struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.List'
	// UNSUPPORTED : C value 'prev' : no Go type for 'GLib.List'
}

var LogFieldStruct *gi.Struct
var LogFieldStructOnce sync.Once

func LogFieldStructSet() {
	LogFieldStructOnce.Do(func() {
		LogFieldStruct = gi.StructNew("GLib", "LogField")
	})
}

type LogField struct {
	native uintptr
	Key    string
	// UNSUPPORTED : C value 'value' : no Go type for 'gpointer'
	Length int32
}

var MainContextStruct *gi.Struct
var MainContextStructOnce sync.Once

func MainContextStructSet() {
	MainContextStructOnce.Do(func() {
		MainContextStruct = gi.StructNew("GLib", "MainContext")
	})
}

type MainContext struct {
	native uintptr
}

var newMainContextInvoker *gi.Function

// MainContextNew is a representation of the C type g_main_context_new.
func MainContextNew() *MainContext {
	if newMainContextInvoker == nil {
		newMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "new")
	}

	ret := newMainContextInvoker.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_main_context_acquire' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_main_context_add_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_main_context_check' : parameter 'fds' has no type

var dispatchMainContextInvoker *gi.Function

// Dispatch is a representation of the C type g_main_context_dispatch.
func (recv *MainContext) Dispatch() {
	if dispatchMainContextInvoker == nil {
		dispatchMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "dispatch")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dispatchMainContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_find_source_by_funcs_user_data' : parameter 'funcs' of type 'SourceFuncs' not supported

var findSourceByIdMainContextInvoker *gi.Function

// FindSourceById is a representation of the C type g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	if findSourceByIdMainContextInvoker == nil {
		findSourceByIdMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "find_source_by_id")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(sourceId)

	ret := findSourceByIdMainContextInvoker.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_main_context_find_source_by_user_data' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_main_context_get_poll_func' : return type 'PollFunc' not supported

// UNSUPPORTED : C value 'g_main_context_invoke' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_main_context_invoke_full' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_main_context_is_owner' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_main_context_iteration' : parameter 'may_block' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_main_context_pending' : return type 'gboolean' not supported

var popThreadDefaultMainContextInvoker *gi.Function

// PopThreadDefault is a representation of the C type g_main_context_pop_thread_default.
func (recv *MainContext) PopThreadDefault() {
	if popThreadDefaultMainContextInvoker == nil {
		popThreadDefaultMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "pop_thread_default")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	popThreadDefaultMainContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_prepare' : return type 'gboolean' not supported

var pushThreadDefaultMainContextInvoker *gi.Function

// PushThreadDefault is a representation of the C type g_main_context_push_thread_default.
func (recv *MainContext) PushThreadDefault() {
	if pushThreadDefaultMainContextInvoker == nil {
		pushThreadDefaultMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "push_thread_default")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pushThreadDefaultMainContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_query' : parameter 'fds' has no type

var refMainContextInvoker *gi.Function

// Ref is a representation of the C type g_main_context_ref.
func (recv *MainContext) Ref() *MainContext {
	if refMainContextInvoker == nil {
		refMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refMainContextInvoker.Invoke(inArgs[:], nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var releaseMainContextInvoker *gi.Function

// Release is a representation of the C type g_main_context_release.
func (recv *MainContext) Release() {
	if releaseMainContextInvoker == nil {
		releaseMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "release")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	releaseMainContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_remove_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_main_context_set_poll_func' : parameter 'func' of type 'PollFunc' not supported

var unrefMainContextInvoker *gi.Function

// Unref is a representation of the C type g_main_context_unref.
func (recv *MainContext) Unref() {
	if unrefMainContextInvoker == nil {
		unrefMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefMainContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_wait' : parameter 'cond' of type 'Cond' not supported

var wakeupMainContextInvoker *gi.Function

// Wakeup is a representation of the C type g_main_context_wakeup.
func (recv *MainContext) Wakeup() {
	if wakeupMainContextInvoker == nil {
		wakeupMainContextInvoker = gi.StructFunctionInvokerNew("GLib", "MainContext", "wakeup")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	wakeupMainContextInvoker.Invoke(inArgs[:], nil)

}

var MainLoopStruct *gi.Struct
var MainLoopStructOnce sync.Once

func MainLoopStructSet() {
	MainLoopStructOnce.Do(func() {
		MainLoopStruct = gi.StructNew("GLib", "MainLoop")
	})
}

type MainLoop struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_main_loop_new' : parameter 'context' of type 'MainContext' not supported

var getContextMainLoopInvoker *gi.Function

// GetContext is a representation of the C type g_main_loop_get_context.
func (recv *MainLoop) GetContext() *MainContext {
	if getContextMainLoopInvoker == nil {
		getContextMainLoopInvoker = gi.StructFunctionInvokerNew("GLib", "MainLoop", "get_context")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getContextMainLoopInvoker.Invoke(inArgs[:], nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_main_loop_is_running' : return type 'gboolean' not supported

var quitMainLoopInvoker *gi.Function

// Quit is a representation of the C type g_main_loop_quit.
func (recv *MainLoop) Quit() {
	if quitMainLoopInvoker == nil {
		quitMainLoopInvoker = gi.StructFunctionInvokerNew("GLib", "MainLoop", "quit")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	quitMainLoopInvoker.Invoke(inArgs[:], nil)

}

var refMainLoopInvoker *gi.Function

// Ref is a representation of the C type g_main_loop_ref.
func (recv *MainLoop) Ref() *MainLoop {
	if refMainLoopInvoker == nil {
		refMainLoopInvoker = gi.StructFunctionInvokerNew("GLib", "MainLoop", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refMainLoopInvoker.Invoke(inArgs[:], nil)

	retGo := &MainLoop{native: ret.Pointer()}

	return retGo
}

var runMainLoopInvoker *gi.Function

// Run is a representation of the C type g_main_loop_run.
func (recv *MainLoop) Run() {
	if runMainLoopInvoker == nil {
		runMainLoopInvoker = gi.StructFunctionInvokerNew("GLib", "MainLoop", "run")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	runMainLoopInvoker.Invoke(inArgs[:], nil)

}

var unrefMainLoopInvoker *gi.Function

// Unref is a representation of the C type g_main_loop_unref.
func (recv *MainLoop) Unref() {
	if unrefMainLoopInvoker == nil {
		unrefMainLoopInvoker = gi.StructFunctionInvokerNew("GLib", "MainLoop", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefMainLoopInvoker.Invoke(inArgs[:], nil)

}

var MappedFileStruct *gi.Struct
var MappedFileStructOnce sync.Once

func MappedFileStructSet() {
	MappedFileStructOnce.Do(func() {
		MappedFileStruct = gi.StructNew("GLib", "MappedFile")
	})
}

type MappedFile struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_mapped_file_new' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mapped_file_new_from_fd' : parameter 'writable' of type 'gboolean' not supported

var freeMappedFileInvoker *gi.Function

// Free is a representation of the C type g_mapped_file_free.
func (recv *MappedFile) Free() {
	if freeMappedFileInvoker == nil {
		freeMappedFileInvoker = gi.StructFunctionInvokerNew("GLib", "MappedFile", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMappedFileInvoker.Invoke(inArgs[:], nil)

}

var getBytesMappedFileInvoker *gi.Function

// GetBytes is a representation of the C type g_mapped_file_get_bytes.
func (recv *MappedFile) GetBytes() *Bytes {
	if getBytesMappedFileInvoker == nil {
		getBytesMappedFileInvoker = gi.StructFunctionInvokerNew("GLib", "MappedFile", "get_bytes")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getBytesMappedFileInvoker.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

var getContentsMappedFileInvoker *gi.Function

// GetContents is a representation of the C type g_mapped_file_get_contents.
func (recv *MappedFile) GetContents() string {
	if getContentsMappedFileInvoker == nil {
		getContentsMappedFileInvoker = gi.StructFunctionInvokerNew("GLib", "MappedFile", "get_contents")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getContentsMappedFileInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_mapped_file_get_length' : return type 'gsize' not supported

var refMappedFileInvoker *gi.Function

// Ref is a representation of the C type g_mapped_file_ref.
func (recv *MappedFile) Ref() *MappedFile {
	if refMappedFileInvoker == nil {
		refMappedFileInvoker = gi.StructFunctionInvokerNew("GLib", "MappedFile", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refMappedFileInvoker.Invoke(inArgs[:], nil)

	retGo := &MappedFile{native: ret.Pointer()}

	return retGo
}

var unrefMappedFileInvoker *gi.Function

// Unref is a representation of the C type g_mapped_file_unref.
func (recv *MappedFile) Unref() {
	if unrefMappedFileInvoker == nil {
		unrefMappedFileInvoker = gi.StructFunctionInvokerNew("GLib", "MappedFile", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefMappedFileInvoker.Invoke(inArgs[:], nil)

}

var MarkupParseContextStruct *gi.Struct
var MarkupParseContextStructOnce sync.Once

func MarkupParseContextStructSet() {
	MarkupParseContextStructOnce.Do(func() {
		MarkupParseContextStruct = gi.StructNew("GLib", "MarkupParseContext")
	})
}

type MarkupParseContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_markup_parse_context_new' : parameter 'parser' of type 'MarkupParser' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_end_parse' : return type 'gboolean' not supported

var freeMarkupParseContextInvoker *gi.Function

// Free is a representation of the C type g_markup_parse_context_free.
func (recv *MarkupParseContext) Free() {
	if freeMarkupParseContextInvoker == nil {
		freeMarkupParseContextInvoker = gi.StructFunctionInvokerNew("GLib", "MarkupParseContext", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMarkupParseContextInvoker.Invoke(inArgs[:], nil)

}

var getElementMarkupParseContextInvoker *gi.Function

// GetElement is a representation of the C type g_markup_parse_context_get_element.
func (recv *MarkupParseContext) GetElement() string {
	if getElementMarkupParseContextInvoker == nil {
		getElementMarkupParseContextInvoker = gi.StructFunctionInvokerNew("GLib", "MarkupParseContext", "get_element")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getElementMarkupParseContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_markup_parse_context_get_element_stack' : return type 'GLib.SList' not supported

var getPositionMarkupParseContextInvoker *gi.Function

// GetPosition is a representation of the C type g_markup_parse_context_get_position.
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) {
	if getPositionMarkupParseContextInvoker == nil {
		getPositionMarkupParseContextInvoker = gi.StructFunctionInvokerNew("GLib", "MarkupParseContext", "get_position")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(lineNumber)
	inArgs[2].SetInt32(charNumber)

	getPositionMarkupParseContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_markup_parse_context_get_user_data' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_parse' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_push' : parameter 'parser' of type 'MarkupParser' not supported

var refMarkupParseContextInvoker *gi.Function

// Ref is a representation of the C type g_markup_parse_context_ref.
func (recv *MarkupParseContext) Ref() *MarkupParseContext {
	if refMarkupParseContextInvoker == nil {
		refMarkupParseContextInvoker = gi.StructFunctionInvokerNew("GLib", "MarkupParseContext", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refMarkupParseContextInvoker.Invoke(inArgs[:], nil)

	retGo := &MarkupParseContext{native: ret.Pointer()}

	return retGo
}

var unrefMarkupParseContextInvoker *gi.Function

// Unref is a representation of the C type g_markup_parse_context_unref.
func (recv *MarkupParseContext) Unref() {
	if unrefMarkupParseContextInvoker == nil {
		unrefMarkupParseContextInvoker = gi.StructFunctionInvokerNew("GLib", "MarkupParseContext", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefMarkupParseContextInvoker.Invoke(inArgs[:], nil)

}

var MarkupParserStruct *gi.Struct
var MarkupParserStructOnce sync.Once

func MarkupParserStructSet() {
	MarkupParserStructOnce.Do(func() {
		MarkupParserStruct = gi.StructNew("GLib", "MarkupParser")
	})
}

type MarkupParser struct {
	native uintptr
	// UNSUPPORTED : C value 'start_element' : missing Type
	// UNSUPPORTED : C value 'end_element' : missing Type
	// UNSUPPORTED : C value 'text' : missing Type
	// UNSUPPORTED : C value 'passthrough' : missing Type
	// UNSUPPORTED : C value 'error' : missing Type
}

var MatchInfoStruct *gi.Struct
var MatchInfoStructOnce sync.Once

func MatchInfoStructSet() {
	MatchInfoStructOnce.Do(func() {
		MatchInfoStruct = gi.StructNew("GLib", "MatchInfo")
	})
}

type MatchInfo struct {
	native uintptr
}

var expandReferencesMatchInfoInvoker *gi.Function

// ExpandReferences is a representation of the C type g_match_info_expand_references.
func (recv *MatchInfo) ExpandReferences(stringToExpand string) string {
	if expandReferencesMatchInfoInvoker == nil {
		expandReferencesMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "expand_references")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(stringToExpand)

	ret := expandReferencesMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var fetchMatchInfoInvoker *gi.Function

// Fetch is a representation of the C type g_match_info_fetch.
func (recv *MatchInfo) Fetch(matchNum int32) string {
	if fetchMatchInfoInvoker == nil {
		fetchMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "fetch")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(matchNum)

	ret := fetchMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var fetchAllMatchInfoInvoker *gi.Function

// FetchAll is a representation of the C type g_match_info_fetch_all.
func (recv *MatchInfo) FetchAll() {
	if fetchAllMatchInfoInvoker == nil {
		fetchAllMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "fetch_all")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	fetchAllMatchInfoInvoker.Invoke(inArgs[:], nil)

}

var fetchNamedMatchInfoInvoker *gi.Function

// FetchNamed is a representation of the C type g_match_info_fetch_named.
func (recv *MatchInfo) FetchNamed(name string) string {
	if fetchNamedMatchInfoInvoker == nil {
		fetchNamedMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "fetch_named")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := fetchNamedMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_match_info_fetch_named_pos' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_match_info_fetch_pos' : return type 'gboolean' not supported

var freeMatchInfoInvoker *gi.Function

// Free is a representation of the C type g_match_info_free.
func (recv *MatchInfo) Free() {
	if freeMatchInfoInvoker == nil {
		freeMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeMatchInfoInvoker.Invoke(inArgs[:], nil)

}

var getMatchCountMatchInfoInvoker *gi.Function

// GetMatchCount is a representation of the C type g_match_info_get_match_count.
func (recv *MatchInfo) GetMatchCount() int32 {
	if getMatchCountMatchInfoInvoker == nil {
		getMatchCountMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "get_match_count")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMatchCountMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getRegexMatchInfoInvoker *gi.Function

// GetRegex is a representation of the C type g_match_info_get_regex.
func (recv *MatchInfo) GetRegex() *Regex {
	if getRegexMatchInfoInvoker == nil {
		getRegexMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "get_regex")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getRegexMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &Regex{native: ret.Pointer()}

	return retGo
}

var getStringMatchInfoInvoker *gi.Function

// GetString is a representation of the C type g_match_info_get_string.
func (recv *MatchInfo) GetString() string {
	if getStringMatchInfoInvoker == nil {
		getStringMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "get_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStringMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_match_info_is_partial_match' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_match_info_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_match_info_next' : return type 'gboolean' not supported

var refMatchInfoInvoker *gi.Function

// Ref is a representation of the C type g_match_info_ref.
func (recv *MatchInfo) Ref() *MatchInfo {
	if refMatchInfoInvoker == nil {
		refMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refMatchInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &MatchInfo{native: ret.Pointer()}

	return retGo
}

var unrefMatchInfoInvoker *gi.Function

// Unref is a representation of the C type g_match_info_unref.
func (recv *MatchInfo) Unref() {
	if unrefMatchInfoInvoker == nil {
		unrefMatchInfoInvoker = gi.StructFunctionInvokerNew("GLib", "MatchInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefMatchInfoInvoker.Invoke(inArgs[:], nil)

}

var MemVTableStruct *gi.Struct
var MemVTableStructOnce sync.Once

func MemVTableStructSet() {
	MemVTableStructOnce.Do(func() {
		MemVTableStruct = gi.StructNew("GLib", "MemVTable")
	})
}

type MemVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'malloc' : missing Type
	// UNSUPPORTED : C value 'realloc' : missing Type
	// UNSUPPORTED : C value 'free' : missing Type
	// UNSUPPORTED : C value 'calloc' : missing Type
	// UNSUPPORTED : C value 'try_malloc' : missing Type
	// UNSUPPORTED : C value 'try_realloc' : missing Type
}

var NodeStruct *gi.Struct
var NodeStructOnce sync.Once

func NodeStructSet() {
	NodeStructOnce.Do(func() {
		NodeStruct = gi.StructNew("GLib", "Node")
	})
}

type Node struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	Next     *Node
	Prev     *Node
	Parent   *Node
	Children *Node
}

// UNSUPPORTED : C value 'g_node_child_index' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_node_child_position' : parameter 'child' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_children_foreach' : parameter 'flags' of type 'TraverseFlags' not supported

var copyNodeInvoker *gi.Function

// Copy is a representation of the C type g_node_copy.
func (recv *Node) Copy() *Node {
	if copyNodeInvoker == nil {
		copyNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyNodeInvoker.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_node_copy_deep' : parameter 'copy_func' of type 'CopyFunc' not supported

var depthNodeInvoker *gi.Function

// Depth is a representation of the C type g_node_depth.
func (recv *Node) Depth() uint32 {
	if depthNodeInvoker == nil {
		depthNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "depth")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := depthNodeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var destroyNodeInvoker *gi.Function

// Destroy is a representation of the C type g_node_destroy.
func (recv *Node) Destroy() {
	if destroyNodeInvoker == nil {
		destroyNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyNodeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_node_find' : parameter 'order' of type 'TraverseType' not supported

// UNSUPPORTED : C value 'g_node_find_child' : parameter 'flags' of type 'TraverseFlags' not supported

var firstSiblingNodeInvoker *gi.Function

// FirstSibling is a representation of the C type g_node_first_sibling.
func (recv *Node) FirstSibling() *Node {
	if firstSiblingNodeInvoker == nil {
		firstSiblingNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "first_sibling")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := firstSiblingNodeInvoker.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

var getRootNodeInvoker *gi.Function

// GetRoot is a representation of the C type g_node_get_root.
func (recv *Node) GetRoot() *Node {
	if getRootNodeInvoker == nil {
		getRootNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "get_root")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getRootNodeInvoker.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_node_insert' : parameter 'node' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_insert_after' : parameter 'sibling' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_insert_before' : parameter 'sibling' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_is_ancestor' : parameter 'descendant' of type 'Node' not supported

var lastChildNodeInvoker *gi.Function

// LastChild is a representation of the C type g_node_last_child.
func (recv *Node) LastChild() *Node {
	if lastChildNodeInvoker == nil {
		lastChildNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "last_child")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := lastChildNodeInvoker.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

var lastSiblingNodeInvoker *gi.Function

// LastSibling is a representation of the C type g_node_last_sibling.
func (recv *Node) LastSibling() *Node {
	if lastSiblingNodeInvoker == nil {
		lastSiblingNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "last_sibling")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := lastSiblingNodeInvoker.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

var maxHeightNodeInvoker *gi.Function

// MaxHeight is a representation of the C type g_node_max_height.
func (recv *Node) MaxHeight() uint32 {
	if maxHeightNodeInvoker == nil {
		maxHeightNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "max_height")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := maxHeightNodeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var nChildrenNodeInvoker *gi.Function

// NChildren is a representation of the C type g_node_n_children.
func (recv *Node) NChildren() uint32 {
	if nChildrenNodeInvoker == nil {
		nChildrenNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "n_children")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nChildrenNodeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_node_n_nodes' : parameter 'flags' of type 'TraverseFlags' not supported

var nthChildNodeInvoker *gi.Function

// NthChild is a representation of the C type g_node_nth_child.
func (recv *Node) NthChild(n uint32) *Node {
	if nthChildNodeInvoker == nil {
		nthChildNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "nth_child")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(n)

	ret := nthChildNodeInvoker.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_node_prepend' : parameter 'node' of type 'Node' not supported

var reverseChildrenNodeInvoker *gi.Function

// ReverseChildren is a representation of the C type g_node_reverse_children.
func (recv *Node) ReverseChildren() {
	if reverseChildrenNodeInvoker == nil {
		reverseChildrenNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "reverse_children")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	reverseChildrenNodeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_node_traverse' : parameter 'order' of type 'TraverseType' not supported

var unlinkNodeInvoker *gi.Function

// Unlink is a representation of the C type g_node_unlink.
func (recv *Node) Unlink() {
	if unlinkNodeInvoker == nil {
		unlinkNodeInvoker = gi.StructFunctionInvokerNew("GLib", "Node", "unlink")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unlinkNodeInvoker.Invoke(inArgs[:], nil)

}

var OnceStruct *gi.Struct
var OnceStructOnce sync.Once

func OnceStructSet() {
	OnceStructOnce.Do(func() {
		OnceStruct = gi.StructNew("GLib", "Once")
	})
}

type Once struct {
	native uintptr
	// UNSUPPORTED : C value 'status' : no Go type for 'OnceStatus'
	// UNSUPPORTED : C value 'retval' : no Go type for 'gpointer'
}

// UNSUPPORTED : C value 'g_once_impl' : parameter 'func' of type 'ThreadFunc' not supported

var OptionContextStruct *gi.Struct
var OptionContextStructOnce sync.Once

func OptionContextStructSet() {
	OptionContextStructOnce.Do(func() {
		OptionContextStruct = gi.StructNew("GLib", "OptionContext")
	})
}

type OptionContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_context_add_group' : parameter 'group' of type 'OptionGroup' not supported

// UNSUPPORTED : C value 'g_option_context_add_main_entries' : parameter 'entries' of type 'OptionEntry' not supported

var freeOptionContextInvoker *gi.Function

// Free is a representation of the C type g_option_context_free.
func (recv *OptionContext) Free() {
	if freeOptionContextInvoker == nil {
		freeOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeOptionContextInvoker.Invoke(inArgs[:], nil)

}

var getDescriptionOptionContextInvoker *gi.Function

// GetDescription is a representation of the C type g_option_context_get_description.
func (recv *OptionContext) GetDescription() string {
	if getDescriptionOptionContextInvoker == nil {
		getDescriptionOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "get_description")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescriptionOptionContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_option_context_get_help' : parameter 'main_help' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_get_help_enabled' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_get_ignore_unknown_options' : return type 'gboolean' not supported

var getMainGroupOptionContextInvoker *gi.Function

// GetMainGroup is a representation of the C type g_option_context_get_main_group.
func (recv *OptionContext) GetMainGroup() *OptionGroup {
	if getMainGroupOptionContextInvoker == nil {
		getMainGroupOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "get_main_group")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMainGroupOptionContextInvoker.Invoke(inArgs[:], nil)

	retGo := &OptionGroup{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_option_context_get_strict_posix' : return type 'gboolean' not supported

var getSummaryOptionContextInvoker *gi.Function

// GetSummary is a representation of the C type g_option_context_get_summary.
func (recv *OptionContext) GetSummary() string {
	if getSummaryOptionContextInvoker == nil {
		getSummaryOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "get_summary")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSummaryOptionContextInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_option_context_parse' : parameter 'argv' has no type

// UNSUPPORTED : C value 'g_option_context_parse_strv' : parameter 'arguments' has no type

var setDescriptionOptionContextInvoker *gi.Function

// SetDescription is a representation of the C type g_option_context_set_description.
func (recv *OptionContext) SetDescription(description string) {
	if setDescriptionOptionContextInvoker == nil {
		setDescriptionOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "set_description")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(description)

	setDescriptionOptionContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_option_context_set_help_enabled' : parameter 'help_enabled' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_set_ignore_unknown_options' : parameter 'ignore_unknown' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_set_main_group' : parameter 'group' of type 'OptionGroup' not supported

// UNSUPPORTED : C value 'g_option_context_set_strict_posix' : parameter 'strict_posix' of type 'gboolean' not supported

var setSummaryOptionContextInvoker *gi.Function

// SetSummary is a representation of the C type g_option_context_set_summary.
func (recv *OptionContext) SetSummary(summary string) {
	if setSummaryOptionContextInvoker == nil {
		setSummaryOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "set_summary")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(summary)

	setSummaryOptionContextInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_option_context_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

var setTranslationDomainOptionContextInvoker *gi.Function

// SetTranslationDomain is a representation of the C type g_option_context_set_translation_domain.
func (recv *OptionContext) SetTranslationDomain(domain string) {
	if setTranslationDomainOptionContextInvoker == nil {
		setTranslationDomainOptionContextInvoker = gi.StructFunctionInvokerNew("GLib", "OptionContext", "set_translation_domain")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	setTranslationDomainOptionContextInvoker.Invoke(inArgs[:], nil)

}

var OptionEntryStruct *gi.Struct
var OptionEntryStructOnce sync.Once

func OptionEntryStructSet() {
	OptionEntryStructOnce.Do(func() {
		OptionEntryStruct = gi.StructNew("GLib", "OptionEntry")
	})
}

type OptionEntry struct {
	native    uintptr
	LongName  string
	ShortName int8
	Flags     int32
	// UNSUPPORTED : C value 'arg' : no Go type for 'OptionArg'
	// UNSUPPORTED : C value 'arg_data' : no Go type for 'gpointer'
	Description    string
	ArgDescription string
}

var OptionGroupStruct *gi.Struct
var OptionGroupStructOnce sync.Once

func OptionGroupStructSet() {
	OptionGroupStructOnce.Do(func() {
		OptionGroupStruct = gi.StructNew("GLib", "OptionGroup")
	})
}

type OptionGroup struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_group_new' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_option_group_add_entries' : parameter 'entries' of type 'OptionEntry' not supported

var freeOptionGroupInvoker *gi.Function

// Free is a representation of the C type g_option_group_free.
func (recv *OptionGroup) Free() {
	if freeOptionGroupInvoker == nil {
		freeOptionGroupInvoker = gi.StructFunctionInvokerNew("GLib", "OptionGroup", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeOptionGroupInvoker.Invoke(inArgs[:], nil)

}

var refOptionGroupInvoker *gi.Function

// Ref is a representation of the C type g_option_group_ref.
func (recv *OptionGroup) Ref() *OptionGroup {
	if refOptionGroupInvoker == nil {
		refOptionGroupInvoker = gi.StructFunctionInvokerNew("GLib", "OptionGroup", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refOptionGroupInvoker.Invoke(inArgs[:], nil)

	retGo := &OptionGroup{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_option_group_set_error_hook' : parameter 'error_func' of type 'OptionErrorFunc' not supported

// UNSUPPORTED : C value 'g_option_group_set_parse_hooks' : parameter 'pre_parse_func' of type 'OptionParseFunc' not supported

// UNSUPPORTED : C value 'g_option_group_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

var setTranslationDomainOptionGroupInvoker *gi.Function

// SetTranslationDomain is a representation of the C type g_option_group_set_translation_domain.
func (recv *OptionGroup) SetTranslationDomain(domain string) {
	if setTranslationDomainOptionGroupInvoker == nil {
		setTranslationDomainOptionGroupInvoker = gi.StructFunctionInvokerNew("GLib", "OptionGroup", "set_translation_domain")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	setTranslationDomainOptionGroupInvoker.Invoke(inArgs[:], nil)

}

var unrefOptionGroupInvoker *gi.Function

// Unref is a representation of the C type g_option_group_unref.
func (recv *OptionGroup) Unref() {
	if unrefOptionGroupInvoker == nil {
		unrefOptionGroupInvoker = gi.StructFunctionInvokerNew("GLib", "OptionGroup", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefOptionGroupInvoker.Invoke(inArgs[:], nil)

}

var PatternSpecStruct *gi.Struct
var PatternSpecStructOnce sync.Once

func PatternSpecStructSet() {
	PatternSpecStructOnce.Do(func() {
		PatternSpecStruct = gi.StructNew("GLib", "PatternSpec")
	})
}

type PatternSpec struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_pattern_spec_equal' : parameter 'pspec2' of type 'PatternSpec' not supported

var freePatternSpecInvoker *gi.Function

// Free is a representation of the C type g_pattern_spec_free.
func (recv *PatternSpec) Free() {
	if freePatternSpecInvoker == nil {
		freePatternSpecInvoker = gi.StructFunctionInvokerNew("GLib", "PatternSpec", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freePatternSpecInvoker.Invoke(inArgs[:], nil)

}

var PollFDStruct *gi.Struct
var PollFDStructOnce sync.Once

func PollFDStructSet() {
	PollFDStructOnce.Do(func() {
		PollFDStruct = gi.StructNew("GLib", "PollFD")
	})
}

type PollFD struct {
	native  uintptr
	Fd      int32
	Events  uint16
	Revents uint16
}

var PrivateStruct *gi.Struct
var PrivateStructOnce sync.Once

func PrivateStructSet() {
	PrivateStructOnce.Do(func() {
		PrivateStruct = gi.StructNew("GLib", "Private")
	})
}

type Private struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_private_get' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_private_replace' : parameter 'value' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_private_set' : parameter 'value' of type 'gpointer' not supported

var PtrArrayStruct *gi.Struct
var PtrArrayStructOnce sync.Once

func PtrArrayStructSet() {
	PtrArrayStructOnce.Do(func() {
		PtrArrayStruct = gi.StructNew("GLib", "PtrArray")
	})
}

type PtrArray struct {
	native uintptr
	// UNSUPPORTED : C value 'pdata' : no Go type for 'gpointer'
	Len uint32
}

var QueueStruct *gi.Struct
var QueueStructOnce sync.Once

func QueueStructSet() {
	QueueStructOnce.Do(func() {
		QueueStruct = gi.StructNew("GLib", "Queue")
	})
}

type Queue struct {
	native uintptr
	// UNSUPPORTED : C value 'head' : no Go type for 'GLib.List'
	// UNSUPPORTED : C value 'tail' : no Go type for 'GLib.List'
	Length uint32
}

var clearQueueInvoker *gi.Function

// Clear is a representation of the C type g_queue_clear.
func (recv *Queue) Clear() {
	if clearQueueInvoker == nil {
		clearQueueInvoker = gi.StructFunctionInvokerNew("GLib", "Queue", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearQueueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_clear_full' : parameter 'free_func' of type 'DestroyNotify' not supported

var copyQueueInvoker *gi.Function

// Copy is a representation of the C type g_queue_copy.
func (recv *Queue) Copy() *Queue {
	if copyQueueInvoker == nil {
		copyQueueInvoker = gi.StructFunctionInvokerNew("GLib", "Queue", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyQueueInvoker.Invoke(inArgs[:], nil)

	retGo := &Queue{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_queue_delete_link' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_find' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_find_custom' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_foreach' : parameter 'func' of type 'Func' not supported

var freeQueueInvoker *gi.Function

// Free is a representation of the C type g_queue_free.
func (recv *Queue) Free() {
	if freeQueueInvoker == nil {
		freeQueueInvoker = gi.StructFunctionInvokerNew("GLib", "Queue", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeQueueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_free_full' : parameter 'free_func' of type 'DestroyNotify' not supported

var getLengthQueueInvoker *gi.Function

// GetLength is a representation of the C type g_queue_get_length.
func (recv *Queue) GetLength() uint32 {
	if getLengthQueueInvoker == nil {
		getLengthQueueInvoker = gi.StructFunctionInvokerNew("GLib", "Queue", "get_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLengthQueueInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_queue_index' : parameter 'data' of type 'gpointer' not supported

var initQueueInvoker *gi.Function

// Init is a representation of the C type g_queue_init.
func (recv *Queue) Init() {
	if initQueueInvoker == nil {
		initQueueInvoker = gi.StructFunctionInvokerNew("GLib", "Queue", "init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	initQueueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_insert_after' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_after_link' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_before' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_before_link' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_sorted' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_is_empty' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_queue_link_index' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_peek_head' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_peek_head_link' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_peek_nth' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_peek_nth_link' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_peek_tail' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_peek_tail_link' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_pop_head' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_pop_head_link' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_pop_nth' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_pop_nth_link' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_pop_tail' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_pop_tail_link' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_push_head' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_push_head_link' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_push_nth' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_push_nth_link' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_push_tail' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_push_tail_link' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_remove' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_remove_all' : parameter 'data' of type 'gpointer' not supported

var reverseQueueInvoker *gi.Function

// Reverse is a representation of the C type g_queue_reverse.
func (recv *Queue) Reverse() {
	if reverseQueueInvoker == nil {
		reverseQueueInvoker = gi.StructFunctionInvokerNew("GLib", "Queue", "reverse")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	reverseQueueInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_sort' : parameter 'compare_func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_queue_unlink' : parameter 'link_' of type 'GLib.List' not supported

var RWLockStruct *gi.Struct
var RWLockStructOnce sync.Once

func RWLockStructSet() {
	RWLockStructOnce.Do(func() {
		RWLockStruct = gi.StructNew("GLib", "RWLock")
	})
}

type RWLock struct {
	native uintptr
}

var clearRWLockInvoker *gi.Function

// Clear is a representation of the C type g_rw_lock_clear.
func (recv *RWLock) Clear() {
	if clearRWLockInvoker == nil {
		clearRWLockInvoker = gi.StructFunctionInvokerNew("GLib", "RWLock", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearRWLockInvoker.Invoke(inArgs[:], nil)

}

var initRWLockInvoker *gi.Function

// Init is a representation of the C type g_rw_lock_init.
func (recv *RWLock) Init() {
	if initRWLockInvoker == nil {
		initRWLockInvoker = gi.StructFunctionInvokerNew("GLib", "RWLock", "init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	initRWLockInvoker.Invoke(inArgs[:], nil)

}

var readerLockRWLockInvoker *gi.Function

// ReaderLock is a representation of the C type g_rw_lock_reader_lock.
func (recv *RWLock) ReaderLock() {
	if readerLockRWLockInvoker == nil {
		readerLockRWLockInvoker = gi.StructFunctionInvokerNew("GLib", "RWLock", "reader_lock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	readerLockRWLockInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_rw_lock_reader_trylock' : return type 'gboolean' not supported

var readerUnlockRWLockInvoker *gi.Function

// ReaderUnlock is a representation of the C type g_rw_lock_reader_unlock.
func (recv *RWLock) ReaderUnlock() {
	if readerUnlockRWLockInvoker == nil {
		readerUnlockRWLockInvoker = gi.StructFunctionInvokerNew("GLib", "RWLock", "reader_unlock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	readerUnlockRWLockInvoker.Invoke(inArgs[:], nil)

}

var writerLockRWLockInvoker *gi.Function

// WriterLock is a representation of the C type g_rw_lock_writer_lock.
func (recv *RWLock) WriterLock() {
	if writerLockRWLockInvoker == nil {
		writerLockRWLockInvoker = gi.StructFunctionInvokerNew("GLib", "RWLock", "writer_lock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	writerLockRWLockInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_rw_lock_writer_trylock' : return type 'gboolean' not supported

var writerUnlockRWLockInvoker *gi.Function

// WriterUnlock is a representation of the C type g_rw_lock_writer_unlock.
func (recv *RWLock) WriterUnlock() {
	if writerUnlockRWLockInvoker == nil {
		writerUnlockRWLockInvoker = gi.StructFunctionInvokerNew("GLib", "RWLock", "writer_unlock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	writerUnlockRWLockInvoker.Invoke(inArgs[:], nil)

}

var RandStruct *gi.Struct
var RandStructOnce sync.Once

func RandStructSet() {
	RandStructOnce.Do(func() {
		RandStruct = gi.StructNew("GLib", "Rand")
	})
}

type Rand struct {
	native uintptr
}

var copyRandInvoker *gi.Function

// Copy is a representation of the C type g_rand_copy.
func (recv *Rand) Copy() *Rand {
	if copyRandInvoker == nil {
		copyRandInvoker = gi.StructFunctionInvokerNew("GLib", "Rand", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyRandInvoker.Invoke(inArgs[:], nil)

	retGo := &Rand{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_rand_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_rand_double_range' : parameter 'begin' of type 'gdouble' not supported

var freeRandInvoker *gi.Function

// Free is a representation of the C type g_rand_free.
func (recv *Rand) Free() {
	if freeRandInvoker == nil {
		freeRandInvoker = gi.StructFunctionInvokerNew("GLib", "Rand", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeRandInvoker.Invoke(inArgs[:], nil)

}

var intRandInvoker *gi.Function

// Int is a representation of the C type g_rand_int.
func (recv *Rand) Int() uint32 {
	if intRandInvoker == nil {
		intRandInvoker = gi.StructFunctionInvokerNew("GLib", "Rand", "int")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := intRandInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var intRangeRandInvoker *gi.Function

// IntRange is a representation of the C type g_rand_int_range.
func (recv *Rand) IntRange(begin int32, end int32) int32 {
	if intRangeRandInvoker == nil {
		intRangeRandInvoker = gi.StructFunctionInvokerNew("GLib", "Rand", "int_range")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(begin)
	inArgs[2].SetInt32(end)

	ret := intRangeRandInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var setSeedRandInvoker *gi.Function

// SetSeed is a representation of the C type g_rand_set_seed.
func (recv *Rand) SetSeed(seed uint32) {
	if setSeedRandInvoker == nil {
		setSeedRandInvoker = gi.StructFunctionInvokerNew("GLib", "Rand", "set_seed")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(seed)

	setSeedRandInvoker.Invoke(inArgs[:], nil)

}

var setSeedArrayRandInvoker *gi.Function

// SetSeedArray is a representation of the C type g_rand_set_seed_array.
func (recv *Rand) SetSeedArray(seed uint32, seedLength uint32) {
	if setSeedArrayRandInvoker == nil {
		setSeedArrayRandInvoker = gi.StructFunctionInvokerNew("GLib", "Rand", "set_seed_array")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(seed)
	inArgs[2].SetUint32(seedLength)

	setSeedArrayRandInvoker.Invoke(inArgs[:], nil)

}

var RecMutexStruct *gi.Struct
var RecMutexStructOnce sync.Once

func RecMutexStructSet() {
	RecMutexStructOnce.Do(func() {
		RecMutexStruct = gi.StructNew("GLib", "RecMutex")
	})
}

type RecMutex struct {
	native uintptr
}

var clearRecMutexInvoker *gi.Function

// Clear is a representation of the C type g_rec_mutex_clear.
func (recv *RecMutex) Clear() {
	if clearRecMutexInvoker == nil {
		clearRecMutexInvoker = gi.StructFunctionInvokerNew("GLib", "RecMutex", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearRecMutexInvoker.Invoke(inArgs[:], nil)

}

var initRecMutexInvoker *gi.Function

// Init is a representation of the C type g_rec_mutex_init.
func (recv *RecMutex) Init() {
	if initRecMutexInvoker == nil {
		initRecMutexInvoker = gi.StructFunctionInvokerNew("GLib", "RecMutex", "init")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	initRecMutexInvoker.Invoke(inArgs[:], nil)

}

var lockRecMutexInvoker *gi.Function

// Lock is a representation of the C type g_rec_mutex_lock.
func (recv *RecMutex) Lock() {
	if lockRecMutexInvoker == nil {
		lockRecMutexInvoker = gi.StructFunctionInvokerNew("GLib", "RecMutex", "lock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	lockRecMutexInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_rec_mutex_trylock' : return type 'gboolean' not supported

var unlockRecMutexInvoker *gi.Function

// Unlock is a representation of the C type g_rec_mutex_unlock.
func (recv *RecMutex) Unlock() {
	if unlockRecMutexInvoker == nil {
		unlockRecMutexInvoker = gi.StructFunctionInvokerNew("GLib", "RecMutex", "unlock")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unlockRecMutexInvoker.Invoke(inArgs[:], nil)

}

var RegexStruct *gi.Struct
var RegexStructOnce sync.Once

func RegexStructSet() {
	RegexStructOnce.Do(func() {
		RegexStruct = gi.StructNew("GLib", "Regex")
	})
}

type Regex struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_regex_new' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var getCaptureCountRegexInvoker *gi.Function

// GetCaptureCount is a representation of the C type g_regex_get_capture_count.
func (recv *Regex) GetCaptureCount() int32 {
	if getCaptureCountRegexInvoker == nil {
		getCaptureCountRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "get_capture_count")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getCaptureCountRegexInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_regex_get_compile_flags' : return type 'RegexCompileFlags' not supported

// UNSUPPORTED : C value 'g_regex_get_has_cr_or_lf' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_regex_get_match_flags' : return type 'RegexMatchFlags' not supported

var getMaxBackrefRegexInvoker *gi.Function

// GetMaxBackref is a representation of the C type g_regex_get_max_backref.
func (recv *Regex) GetMaxBackref() int32 {
	if getMaxBackrefRegexInvoker == nil {
		getMaxBackrefRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "get_max_backref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMaxBackrefRegexInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getMaxLookbehindRegexInvoker *gi.Function

// GetMaxLookbehind is a representation of the C type g_regex_get_max_lookbehind.
func (recv *Regex) GetMaxLookbehind() int32 {
	if getMaxLookbehindRegexInvoker == nil {
		getMaxLookbehindRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "get_max_lookbehind")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMaxLookbehindRegexInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getPatternRegexInvoker *gi.Function

// GetPattern is a representation of the C type g_regex_get_pattern.
func (recv *Regex) GetPattern() string {
	if getPatternRegexInvoker == nil {
		getPatternRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "get_pattern")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPatternRegexInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getStringNumberRegexInvoker *gi.Function

// GetStringNumber is a representation of the C type g_regex_get_string_number.
func (recv *Regex) GetStringNumber(name string) int32 {
	if getStringNumberRegexInvoker == nil {
		getStringNumberRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "get_string_number")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := getStringNumberRegexInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_regex_match' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_match_all' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_match_all_full' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_full' : parameter 'string' has no type

var refRegexInvoker *gi.Function

// Ref is a representation of the C type g_regex_ref.
func (recv *Regex) Ref() *Regex {
	if refRegexInvoker == nil {
		refRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refRegexInvoker.Invoke(inArgs[:], nil)

	retGo := &Regex{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_regex_replace' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_replace_eval' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_replace_literal' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_split' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_full' : parameter 'string' has no type

var unrefRegexInvoker *gi.Function

// Unref is a representation of the C type g_regex_unref.
func (recv *Regex) Unref() {
	if unrefRegexInvoker == nil {
		unrefRegexInvoker = gi.StructFunctionInvokerNew("GLib", "Regex", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefRegexInvoker.Invoke(inArgs[:], nil)

}

var SListStruct *gi.Struct
var SListStructOnce sync.Once

func SListStructSet() {
	SListStructOnce.Do(func() {
		SListStruct = gi.StructNew("GLib", "SList")
	})
}

type SList struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.SList'
}

var ScannerStruct *gi.Struct
var ScannerStructOnce sync.Once

func ScannerStructSet() {
	ScannerStructOnce.Do(func() {
		ScannerStruct = gi.StructNew("GLib", "Scanner")
	})
}

type Scanner struct {
	native uintptr
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	Qdata          *Data
	Config         *ScannerConfig
	// UNSUPPORTED : C value 'token' : no Go type for 'TokenType'
	// UNSUPPORTED : C value 'value' : no Go type for 'TokenValue'
	Line     uint32
	Position uint32
	// UNSUPPORTED : C value 'next_token' : no Go type for 'TokenType'
	// UNSUPPORTED : C value 'next_value' : no Go type for 'TokenValue'
	NextLine     uint32
	NextPosition uint32
	// UNSUPPORTED : C value 'msg_handler' : no Go type for 'ScannerMsgFunc'
}

var curLineScannerInvoker *gi.Function

// CurLine is a representation of the C type g_scanner_cur_line.
func (recv *Scanner) CurLine() uint32 {
	if curLineScannerInvoker == nil {
		curLineScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "cur_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := curLineScannerInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var curPositionScannerInvoker *gi.Function

// CurPosition is a representation of the C type g_scanner_cur_position.
func (recv *Scanner) CurPosition() uint32 {
	if curPositionScannerInvoker == nil {
		curPositionScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "cur_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := curPositionScannerInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_scanner_cur_token' : return type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_cur_value' : return type 'TokenValue' not supported

var destroyScannerInvoker *gi.Function

// Destroy is a representation of the C type g_scanner_destroy.
func (recv *Scanner) Destroy() {
	if destroyScannerInvoker == nil {
		destroyScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyScannerInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_scanner_eof' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_scanner_error' : parameter '...' has no type

// UNSUPPORTED : C value 'g_scanner_get_next_token' : return type 'TokenType' not supported

var inputFileScannerInvoker *gi.Function

// InputFile is a representation of the C type g_scanner_input_file.
func (recv *Scanner) InputFile(inputFd int32) {
	if inputFileScannerInvoker == nil {
		inputFileScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "input_file")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(inputFd)

	inputFileScannerInvoker.Invoke(inArgs[:], nil)

}

var inputTextScannerInvoker *gi.Function

// InputText is a representation of the C type g_scanner_input_text.
func (recv *Scanner) InputText(text string, textLen uint32) {
	if inputTextScannerInvoker == nil {
		inputTextScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "input_text")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetUint32(textLen)

	inputTextScannerInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_scanner_lookup_symbol' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_scanner_peek_next_token' : return type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_scope_add_symbol' : parameter 'value' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_scanner_scope_foreach_symbol' : parameter 'func' of type 'HFunc' not supported

// UNSUPPORTED : C value 'g_scanner_scope_lookup_symbol' : return type 'gpointer' not supported

var scopeRemoveSymbolScannerInvoker *gi.Function

// ScopeRemoveSymbol is a representation of the C type g_scanner_scope_remove_symbol.
func (recv *Scanner) ScopeRemoveSymbol(scopeId uint32, symbol string) {
	if scopeRemoveSymbolScannerInvoker == nil {
		scopeRemoveSymbolScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "scope_remove_symbol")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(scopeId)
	inArgs[2].SetString(symbol)

	scopeRemoveSymbolScannerInvoker.Invoke(inArgs[:], nil)

}

var setScopeScannerInvoker *gi.Function

// SetScope is a representation of the C type g_scanner_set_scope.
func (recv *Scanner) SetScope(scopeId uint32) uint32 {
	if setScopeScannerInvoker == nil {
		setScopeScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "set_scope")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(scopeId)

	ret := setScopeScannerInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var syncFileOffsetScannerInvoker *gi.Function

// SyncFileOffset is a representation of the C type g_scanner_sync_file_offset.
func (recv *Scanner) SyncFileOffset() {
	if syncFileOffsetScannerInvoker == nil {
		syncFileOffsetScannerInvoker = gi.StructFunctionInvokerNew("GLib", "Scanner", "sync_file_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	syncFileOffsetScannerInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_scanner_unexp_token' : parameter 'expected_token' of type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_warn' : parameter '...' has no type

var ScannerConfigStruct *gi.Struct
var ScannerConfigStructOnce sync.Once

func ScannerConfigStructSet() {
	ScannerConfigStructOnce.Do(func() {
		ScannerConfigStruct = gi.StructNew("GLib", "ScannerConfig")
	})
}

type ScannerConfig struct {
	native              uintptr
	CsetSkipCharacters  string
	CsetIdentifierFirst string
	CsetIdentifierNth   string
	CpairCommentSingle  string
	CaseSensitive       uint32
	SkipCommentMulti    uint32
	SkipCommentSingle   uint32
	ScanCommentMulti    uint32
	ScanIdentifier      uint32
	ScanIdentifier1char uint32
	ScanIdentifierNULL  uint32
	ScanSymbols         uint32
	ScanBinary          uint32
	ScanOctal           uint32
	ScanFloat           uint32
	ScanHex             uint32
	ScanHexDollar       uint32
	ScanStringSq        uint32
	ScanStringDq        uint32
	Numbers2Int         uint32
	Int2Float           uint32
	Identifier2String   uint32
	Char2Token          uint32
	Symbol2Token        uint32
	Scope0Fallback      uint32
	StoreInt64          uint32
}

var SequenceStruct *gi.Struct
var SequenceStructOnce sync.Once

func SequenceStructSet() {
	SequenceStructOnce.Do(func() {
		SequenceStruct = gi.StructNew("GLib", "Sequence")
	})
}

type Sequence struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_sequence_append' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_foreach' : parameter 'func' of type 'Func' not supported

var freeSequenceInvoker *gi.Function

// Free is a representation of the C type g_sequence_free.
func (recv *Sequence) Free() {
	if freeSequenceInvoker == nil {
		freeSequenceInvoker = gi.StructFunctionInvokerNew("GLib", "Sequence", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeSequenceInvoker.Invoke(inArgs[:], nil)

}

var getBeginIterSequenceInvoker *gi.Function

// GetBeginIter is a representation of the C type g_sequence_get_begin_iter.
func (recv *Sequence) GetBeginIter() *SequenceIter {
	if getBeginIterSequenceInvoker == nil {
		getBeginIterSequenceInvoker = gi.StructFunctionInvokerNew("GLib", "Sequence", "get_begin_iter")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getBeginIterSequenceInvoker.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var getEndIterSequenceInvoker *gi.Function

// GetEndIter is a representation of the C type g_sequence_get_end_iter.
func (recv *Sequence) GetEndIter() *SequenceIter {
	if getEndIterSequenceInvoker == nil {
		getEndIterSequenceInvoker = gi.StructFunctionInvokerNew("GLib", "Sequence", "get_end_iter")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getEndIterSequenceInvoker.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var getIterAtPosSequenceInvoker *gi.Function

// GetIterAtPos is a representation of the C type g_sequence_get_iter_at_pos.
func (recv *Sequence) GetIterAtPos(pos int32) *SequenceIter {
	if getIterAtPosSequenceInvoker == nil {
		getIterAtPosSequenceInvoker = gi.StructFunctionInvokerNew("GLib", "Sequence", "get_iter_at_pos")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := getIterAtPosSequenceInvoker.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var getLengthSequenceInvoker *gi.Function

// GetLength is a representation of the C type g_sequence_get_length.
func (recv *Sequence) GetLength() int32 {
	if getLengthSequenceInvoker == nil {
		getLengthSequenceInvoker = gi.StructFunctionInvokerNew("GLib", "Sequence", "get_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLengthSequenceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_sequence_insert_sorted' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_insert_sorted_iter' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_is_empty' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_sequence_lookup' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_lookup_iter' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_prepend' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_search' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_search_iter' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_sort' : parameter 'cmp_func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_sequence_sort_iter' : parameter 'cmp_func' of type 'SequenceIterCompareFunc' not supported

var SequenceIterStruct *gi.Struct
var SequenceIterStructOnce sync.Once

func SequenceIterStructSet() {
	SequenceIterStructOnce.Do(func() {
		SequenceIterStruct = gi.StructNew("GLib", "SequenceIter")
	})
}

type SequenceIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_sequence_iter_compare' : parameter 'b' of type 'SequenceIter' not supported

var getPositionSequenceIterInvoker *gi.Function

// GetPosition is a representation of the C type g_sequence_iter_get_position.
func (recv *SequenceIter) GetPosition() int32 {
	if getPositionSequenceIterInvoker == nil {
		getPositionSequenceIterInvoker = gi.StructFunctionInvokerNew("GLib", "SequenceIter", "get_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPositionSequenceIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getSequenceSequenceIterInvoker *gi.Function

// GetSequence is a representation of the C type g_sequence_iter_get_sequence.
func (recv *SequenceIter) GetSequence() *Sequence {
	if getSequenceSequenceIterInvoker == nil {
		getSequenceSequenceIterInvoker = gi.StructFunctionInvokerNew("GLib", "SequenceIter", "get_sequence")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getSequenceSequenceIterInvoker.Invoke(inArgs[:], nil)

	retGo := &Sequence{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_sequence_iter_is_begin' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_sequence_iter_is_end' : return type 'gboolean' not supported

var moveSequenceIterInvoker *gi.Function

// Move is a representation of the C type g_sequence_iter_move.
func (recv *SequenceIter) Move(delta int32) *SequenceIter {
	if moveSequenceIterInvoker == nil {
		moveSequenceIterInvoker = gi.StructFunctionInvokerNew("GLib", "SequenceIter", "move")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(delta)

	ret := moveSequenceIterInvoker.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var nextSequenceIterInvoker *gi.Function

// Next is a representation of the C type g_sequence_iter_next.
func (recv *SequenceIter) Next() *SequenceIter {
	if nextSequenceIterInvoker == nil {
		nextSequenceIterInvoker = gi.StructFunctionInvokerNew("GLib", "SequenceIter", "next")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nextSequenceIterInvoker.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var prevSequenceIterInvoker *gi.Function

// Prev is a representation of the C type g_sequence_iter_prev.
func (recv *SequenceIter) Prev() *SequenceIter {
	if prevSequenceIterInvoker == nil {
		prevSequenceIterInvoker = gi.StructFunctionInvokerNew("GLib", "SequenceIter", "prev")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := prevSequenceIterInvoker.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var SourceStruct *gi.Struct
var SourceStructOnce sync.Once

func SourceStructSet() {
	SourceStructOnce.Do(func() {
		SourceStruct = gi.StructNew("GLib", "Source")
	})
}

type Source struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_source_new' : parameter 'source_funcs' of type 'SourceFuncs' not supported

// UNSUPPORTED : C value 'g_source_add_child_source' : parameter 'child_source' of type 'Source' not supported

// UNSUPPORTED : C value 'g_source_add_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_source_add_unix_fd' : parameter 'events' of type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_source_attach' : parameter 'context' of type 'MainContext' not supported

var destroySourceInvoker *gi.Function

// Destroy is a representation of the C type g_source_destroy.
func (recv *Source) Destroy() {
	if destroySourceInvoker == nil {
		destroySourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroySourceInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_source_get_can_recurse' : return type 'gboolean' not supported

var getContextSourceInvoker *gi.Function

// GetContext is a representation of the C type g_source_get_context.
func (recv *Source) GetContext() *MainContext {
	if getContextSourceInvoker == nil {
		getContextSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "get_context")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getContextSourceInvoker.Invoke(inArgs[:], nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_source_get_current_time' : parameter 'timeval' of type 'TimeVal' not supported

var getIdSourceInvoker *gi.Function

// GetId is a representation of the C type g_source_get_id.
func (recv *Source) GetId() uint32 {
	if getIdSourceInvoker == nil {
		getIdSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "get_id")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIdSourceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getNameSourceInvoker *gi.Function

// GetName is a representation of the C type g_source_get_name.
func (recv *Source) GetName() string {
	if getNameSourceInvoker == nil {
		getNameSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNameSourceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPrioritySourceInvoker *gi.Function

// GetPriority is a representation of the C type g_source_get_priority.
func (recv *Source) GetPriority() int32 {
	if getPrioritySourceInvoker == nil {
		getPrioritySourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "get_priority")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPrioritySourceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getReadyTimeSourceInvoker *gi.Function

// GetReadyTime is a representation of the C type g_source_get_ready_time.
func (recv *Source) GetReadyTime() int64 {
	if getReadyTimeSourceInvoker == nil {
		getReadyTimeSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "get_ready_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getReadyTimeSourceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getTimeSourceInvoker *gi.Function

// GetTime is a representation of the C type g_source_get_time.
func (recv *Source) GetTime() int64 {
	if getTimeSourceInvoker == nil {
		getTimeSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "get_time")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTimeSourceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_source_is_destroyed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_source_modify_unix_fd' : parameter 'tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_source_query_unix_fd' : parameter 'tag' of type 'gpointer' not supported

var refSourceInvoker *gi.Function

// Ref is a representation of the C type g_source_ref.
func (recv *Source) Ref() *Source {
	if refSourceInvoker == nil {
		refSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSourceInvoker.Invoke(inArgs[:], nil)

	retGo := &Source{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_source_remove_child_source' : parameter 'child_source' of type 'Source' not supported

// UNSUPPORTED : C value 'g_source_remove_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_source_remove_unix_fd' : parameter 'tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_source_set_callback' : parameter 'func' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_source_set_callback_indirect' : parameter 'callback_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_source_set_can_recurse' : parameter 'can_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_source_set_funcs' : parameter 'funcs' of type 'SourceFuncs' not supported

var setNameSourceInvoker *gi.Function

// SetName is a representation of the C type g_source_set_name.
func (recv *Source) SetName(name string) {
	if setNameSourceInvoker == nil {
		setNameSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "set_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	setNameSourceInvoker.Invoke(inArgs[:], nil)

}

var setPrioritySourceInvoker *gi.Function

// SetPriority is a representation of the C type g_source_set_priority.
func (recv *Source) SetPriority(priority int32) {
	if setPrioritySourceInvoker == nil {
		setPrioritySourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "set_priority")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(priority)

	setPrioritySourceInvoker.Invoke(inArgs[:], nil)

}

var setReadyTimeSourceInvoker *gi.Function

// SetReadyTime is a representation of the C type g_source_set_ready_time.
func (recv *Source) SetReadyTime(readyTime int64) {
	if setReadyTimeSourceInvoker == nil {
		setReadyTimeSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "set_ready_time")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(readyTime)

	setReadyTimeSourceInvoker.Invoke(inArgs[:], nil)

}

var unrefSourceInvoker *gi.Function

// Unref is a representation of the C type g_source_unref.
func (recv *Source) Unref() {
	if unrefSourceInvoker == nil {
		unrefSourceInvoker = gi.StructFunctionInvokerNew("GLib", "Source", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefSourceInvoker.Invoke(inArgs[:], nil)

}

var SourceCallbackFuncsStruct *gi.Struct
var SourceCallbackFuncsStructOnce sync.Once

func SourceCallbackFuncsStructSet() {
	SourceCallbackFuncsStructOnce.Do(func() {
		SourceCallbackFuncsStruct = gi.StructNew("GLib", "SourceCallbackFuncs")
	})
}

type SourceCallbackFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'ref' : missing Type
	// UNSUPPORTED : C value 'unref' : missing Type
	// UNSUPPORTED : C value 'get' : missing Type
}

var SourceFuncsStruct *gi.Struct
var SourceFuncsStructOnce sync.Once

func SourceFuncsStructSet() {
	SourceFuncsStructOnce.Do(func() {
		SourceFuncsStruct = gi.StructNew("GLib", "SourceFuncs")
	})
}

type SourceFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'prepare' : missing Type
	// UNSUPPORTED : C value 'check' : missing Type
	// UNSUPPORTED : C value 'dispatch' : missing Type
	// UNSUPPORTED : C value 'finalize' : missing Type
}

var SourcePrivateStruct *gi.Struct
var SourcePrivateStructOnce sync.Once

func SourcePrivateStructSet() {
	SourcePrivateStructOnce.Do(func() {
		SourcePrivateStruct = gi.StructNew("GLib", "SourcePrivate")
	})
}

type SourcePrivate struct {
	native uintptr
}

var StatBufStruct *gi.Struct
var StatBufStructOnce sync.Once

func StatBufStructSet() {
	StatBufStructOnce.Do(func() {
		StatBufStruct = gi.StructNew("GLib", "StatBuf")
	})
}

type StatBuf struct {
	native uintptr
}

var StringStruct *gi.Struct
var StringStructOnce sync.Once

func StringStructSet() {
	StringStructOnce.Do(func() {
		StringStruct = gi.StructNew("GLib", "String")
	})
}

type String struct {
	native       uintptr
	Str          string
	Len          uintptr
	AllocatedLen uintptr
}

var appendStringInvoker *gi.Function

// Append is a representation of the C type g_string_append.
func (recv *String) Append(val string) *String {
	if appendStringInvoker == nil {
		appendStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "append")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)

	ret := appendStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var appendCStringInvoker *gi.Function

// AppendC is a representation of the C type g_string_append_c.
func (recv *String) AppendC(c int8) *String {
	if appendCStringInvoker == nil {
		appendCStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "append_c")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(c)

	ret := appendCStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var appendLenStringInvoker *gi.Function

// AppendLen is a representation of the C type g_string_append_len.
func (recv *String) AppendLen(val string, len int32) *String {
	if appendLenStringInvoker == nil {
		appendLenStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "append_len")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)
	inArgs[2].SetInt32(len)

	ret := appendLenStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_append_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_string_append_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_append_uri_escaped' : parameter 'allow_utf8' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_string_append_vprintf' : parameter 'args' of type 'va_list' not supported

var asciiDownStringInvoker *gi.Function

// AsciiDown is a representation of the C type g_string_ascii_down.
func (recv *String) AsciiDown() *String {
	if asciiDownStringInvoker == nil {
		asciiDownStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "ascii_down")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := asciiDownStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var asciiUpStringInvoker *gi.Function

// AsciiUp is a representation of the C type g_string_ascii_up.
func (recv *String) AsciiUp() *String {
	if asciiUpStringInvoker == nil {
		asciiUpStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "ascii_up")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := asciiUpStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var assignStringInvoker *gi.Function

// Assign is a representation of the C type g_string_assign.
func (recv *String) Assign(rval string) *String {
	if assignStringInvoker == nil {
		assignStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "assign")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(rval)

	ret := assignStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var downStringInvoker *gi.Function

// Down is a representation of the C type g_string_down.
func (recv *String) Down() *String {
	if downStringInvoker == nil {
		downStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "down")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := downStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_equal' : parameter 'v2' of type 'String' not supported

var eraseStringInvoker *gi.Function

// Erase is a representation of the C type g_string_erase.
func (recv *String) Erase(pos int32, len int32) *String {
	if eraseStringInvoker == nil {
		eraseStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "erase")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetInt32(len)

	ret := eraseStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_free' : parameter 'free_segment' of type 'gboolean' not supported

var freeToBytesStringInvoker *gi.Function

// FreeToBytes is a representation of the C type g_string_free_to_bytes.
func (recv *String) FreeToBytes() *Bytes {
	if freeToBytesStringInvoker == nil {
		freeToBytesStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "free_to_bytes")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := freeToBytesStringInvoker.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

var hashStringInvoker *gi.Function

// Hash is a representation of the C type g_string_hash.
func (recv *String) Hash() uint32 {
	if hashStringInvoker == nil {
		hashStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashStringInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var insertStringInvoker *gi.Function

// Insert is a representation of the C type g_string_insert.
func (recv *String) Insert(pos int32, val string) *String {
	if insertStringInvoker == nil {
		insertStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "insert")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(val)

	ret := insertStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var insertCStringInvoker *gi.Function

// InsertC is a representation of the C type g_string_insert_c.
func (recv *String) InsertC(pos int32, c int8) *String {
	if insertCStringInvoker == nil {
		insertCStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "insert_c")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetInt8(c)

	ret := insertCStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var insertLenStringInvoker *gi.Function

// InsertLen is a representation of the C type g_string_insert_len.
func (recv *String) InsertLen(pos int32, val string, len int32) *String {
	if insertLenStringInvoker == nil {
		insertLenStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "insert_len")
	}

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(val)
	inArgs[3].SetInt32(len)

	ret := insertLenStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_insert_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_overwrite' : parameter 'pos' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_string_overwrite_len' : parameter 'pos' of type 'gsize' not supported

var prependStringInvoker *gi.Function

// Prepend is a representation of the C type g_string_prepend.
func (recv *String) Prepend(val string) *String {
	if prependStringInvoker == nil {
		prependStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "prepend")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)

	ret := prependStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var prependCStringInvoker *gi.Function

// PrependC is a representation of the C type g_string_prepend_c.
func (recv *String) PrependC(c int8) *String {
	if prependCStringInvoker == nil {
		prependCStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "prepend_c")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(c)

	ret := prependCStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var prependLenStringInvoker *gi.Function

// PrependLen is a representation of the C type g_string_prepend_len.
func (recv *String) PrependLen(val string, len int32) *String {
	if prependLenStringInvoker == nil {
		prependLenStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "prepend_len")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)
	inArgs[2].SetInt32(len)

	ret := prependLenStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_prepend_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_string_set_size' : parameter 'len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_string_truncate' : parameter 'len' of type 'gsize' not supported

var upStringInvoker *gi.Function

// Up is a representation of the C type g_string_up.
func (recv *String) Up() *String {
	if upStringInvoker == nil {
		upStringInvoker = gi.StructFunctionInvokerNew("GLib", "String", "up")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := upStringInvoker.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_vprintf' : parameter 'args' of type 'va_list' not supported

var StringChunkStruct *gi.Struct
var StringChunkStructOnce sync.Once

func StringChunkStructSet() {
	StringChunkStructOnce.Do(func() {
		StringChunkStruct = gi.StructNew("GLib", "StringChunk")
	})
}

type StringChunk struct {
	native uintptr
}

var clearStringChunkInvoker *gi.Function

// Clear is a representation of the C type g_string_chunk_clear.
func (recv *StringChunk) Clear() {
	if clearStringChunkInvoker == nil {
		clearStringChunkInvoker = gi.StructFunctionInvokerNew("GLib", "StringChunk", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearStringChunkInvoker.Invoke(inArgs[:], nil)

}

var freeStringChunkInvoker *gi.Function

// Free is a representation of the C type g_string_chunk_free.
func (recv *StringChunk) Free() {
	if freeStringChunkInvoker == nil {
		freeStringChunkInvoker = gi.StructFunctionInvokerNew("GLib", "StringChunk", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeStringChunkInvoker.Invoke(inArgs[:], nil)

}

var insertStringChunkInvoker *gi.Function

// Insert is a representation of the C type g_string_chunk_insert.
func (recv *StringChunk) Insert(string_ string) string {
	if insertStringChunkInvoker == nil {
		insertStringChunkInvoker = gi.StructFunctionInvokerNew("GLib", "StringChunk", "insert")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)

	ret := insertStringChunkInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var insertConstStringChunkInvoker *gi.Function

// InsertConst is a representation of the C type g_string_chunk_insert_const.
func (recv *StringChunk) InsertConst(string_ string) string {
	if insertConstStringChunkInvoker == nil {
		insertConstStringChunkInvoker = gi.StructFunctionInvokerNew("GLib", "StringChunk", "insert_const")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)

	ret := insertConstStringChunkInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var insertLenStringChunkInvoker *gi.Function

// InsertLen is a representation of the C type g_string_chunk_insert_len.
func (recv *StringChunk) InsertLen(string_ string, len int32) string {
	if insertLenStringChunkInvoker == nil {
		insertLenStringChunkInvoker = gi.StructFunctionInvokerNew("GLib", "StringChunk", "insert_len")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)
	inArgs[2].SetInt32(len)

	ret := insertLenStringChunkInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var TestCaseStruct *gi.Struct
var TestCaseStructOnce sync.Once

func TestCaseStructSet() {
	TestCaseStructOnce.Do(func() {
		TestCaseStruct = gi.StructNew("GLib", "TestCase")
	})
}

type TestCase struct {
	native uintptr
}

var TestConfigStruct *gi.Struct
var TestConfigStructOnce sync.Once

func TestConfigStructSet() {
	TestConfigStructOnce.Do(func() {
		TestConfigStruct = gi.StructNew("GLib", "TestConfig")
	})
}

type TestConfig struct {
	native uintptr
	// UNSUPPORTED : C value 'test_initialized' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'test_quick' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'test_perf' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'test_verbose' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'test_quiet' : no Go type for 'gboolean'
	// UNSUPPORTED : C value 'test_undefined' : no Go type for 'gboolean'
}

var TestLogBufferStruct *gi.Struct
var TestLogBufferStructOnce sync.Once

func TestLogBufferStructSet() {
	TestLogBufferStructOnce.Do(func() {
		TestLogBufferStruct = gi.StructNew("GLib", "TestLogBuffer")
	})
}

type TestLogBuffer struct {
	native uintptr
}

var freeTestLogBufferInvoker *gi.Function

// Free is a representation of the C type g_test_log_buffer_free.
func (recv *TestLogBuffer) Free() {
	if freeTestLogBufferInvoker == nil {
		freeTestLogBufferInvoker = gi.StructFunctionInvokerNew("GLib", "TestLogBuffer", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTestLogBufferInvoker.Invoke(inArgs[:], nil)

}

var popTestLogBufferInvoker *gi.Function

// Pop is a representation of the C type g_test_log_buffer_pop.
func (recv *TestLogBuffer) Pop() *TestLogMsg {
	if popTestLogBufferInvoker == nil {
		popTestLogBufferInvoker = gi.StructFunctionInvokerNew("GLib", "TestLogBuffer", "pop")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := popTestLogBufferInvoker.Invoke(inArgs[:], nil)

	retGo := &TestLogMsg{native: ret.Pointer()}

	return retGo
}

var pushTestLogBufferInvoker *gi.Function

// Push is a representation of the C type g_test_log_buffer_push.
func (recv *TestLogBuffer) Push(nBytes uint32, bytes uint8) {
	if pushTestLogBufferInvoker == nil {
		pushTestLogBufferInvoker = gi.StructFunctionInvokerNew("GLib", "TestLogBuffer", "push")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nBytes)
	inArgs[2].SetUint8(bytes)

	pushTestLogBufferInvoker.Invoke(inArgs[:], nil)

}

var TestLogMsgStruct *gi.Struct
var TestLogMsgStructOnce sync.Once

func TestLogMsgStructSet() {
	TestLogMsgStructOnce.Do(func() {
		TestLogMsgStruct = gi.StructNew("GLib", "TestLogMsg")
	})
}

type TestLogMsg struct {
	native uintptr
	// UNSUPPORTED : C value 'log_type' : no Go type for 'TestLogType'
	NStrings uint32
	Strings  string
	NNums    uint32
	// UNSUPPORTED : C value 'nums' : no Go type for 'long double'
}

var freeTestLogMsgInvoker *gi.Function

// Free is a representation of the C type g_test_log_msg_free.
func (recv *TestLogMsg) Free() {
	if freeTestLogMsgInvoker == nil {
		freeTestLogMsgInvoker = gi.StructFunctionInvokerNew("GLib", "TestLogMsg", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTestLogMsgInvoker.Invoke(inArgs[:], nil)

}

var TestSuiteStruct *gi.Struct
var TestSuiteStructOnce sync.Once

func TestSuiteStructSet() {
	TestSuiteStructOnce.Do(func() {
		TestSuiteStruct = gi.StructNew("GLib", "TestSuite")
	})
}

type TestSuite struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_test_suite_add' : parameter 'test_case' of type 'TestCase' not supported

// UNSUPPORTED : C value 'g_test_suite_add_suite' : parameter 'nestedsuite' of type 'TestSuite' not supported

var ThreadStruct *gi.Struct
var ThreadStructOnce sync.Once

func ThreadStructSet() {
	ThreadStructOnce.Do(func() {
		ThreadStruct = gi.StructNew("GLib", "Thread")
	})
}

type Thread struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_thread_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_try_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_join' : return type 'gpointer' not supported

var refThreadInvoker *gi.Function

// Ref is a representation of the C type g_thread_ref.
func (recv *Thread) Ref() *Thread {
	if refThreadInvoker == nil {
		refThreadInvoker = gi.StructFunctionInvokerNew("GLib", "Thread", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refThreadInvoker.Invoke(inArgs[:], nil)

	retGo := &Thread{native: ret.Pointer()}

	return retGo
}

var unrefThreadInvoker *gi.Function

// Unref is a representation of the C type g_thread_unref.
func (recv *Thread) Unref() {
	if unrefThreadInvoker == nil {
		unrefThreadInvoker = gi.StructFunctionInvokerNew("GLib", "Thread", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefThreadInvoker.Invoke(inArgs[:], nil)

}

var ThreadPoolStruct *gi.Struct
var ThreadPoolStructOnce sync.Once

func ThreadPoolStructSet() {
	ThreadPoolStructOnce.Do(func() {
		ThreadPoolStruct = gi.StructNew("GLib", "ThreadPool")
	})
}

type ThreadPool struct {
	native uintptr
	// UNSUPPORTED : C value 'func' : no Go type for 'Func'
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'exclusive' : no Go type for 'gboolean'
}

// UNSUPPORTED : C value 'g_thread_pool_free' : parameter 'immediate' of type 'gboolean' not supported

var getMaxThreadsThreadPoolInvoker *gi.Function

// GetMaxThreads is a representation of the C type g_thread_pool_get_max_threads.
func (recv *ThreadPool) GetMaxThreads() int32 {
	if getMaxThreadsThreadPoolInvoker == nil {
		getMaxThreadsThreadPoolInvoker = gi.StructFunctionInvokerNew("GLib", "ThreadPool", "get_max_threads")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMaxThreadsThreadPoolInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getNumThreadsThreadPoolInvoker *gi.Function

// GetNumThreads is a representation of the C type g_thread_pool_get_num_threads.
func (recv *ThreadPool) GetNumThreads() uint32 {
	if getNumThreadsThreadPoolInvoker == nil {
		getNumThreadsThreadPoolInvoker = gi.StructFunctionInvokerNew("GLib", "ThreadPool", "get_num_threads")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNumThreadsThreadPoolInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_thread_pool_move_to_front' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_thread_pool_push' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_thread_pool_set_max_threads' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_thread_pool_set_sort_function' : parameter 'func' of type 'CompareDataFunc' not supported

var unprocessedThreadPoolInvoker *gi.Function

// Unprocessed is a representation of the C type g_thread_pool_unprocessed.
func (recv *ThreadPool) Unprocessed() uint32 {
	if unprocessedThreadPoolInvoker == nil {
		unprocessedThreadPoolInvoker = gi.StructFunctionInvokerNew("GLib", "ThreadPool", "unprocessed")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := unprocessedThreadPoolInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var TimeValStruct *gi.Struct
var TimeValStructOnce sync.Once

func TimeValStructSet() {
	TimeValStructOnce.Do(func() {
		TimeValStruct = gi.StructNew("GLib", "TimeVal")
	})
}

type TimeVal struct {
	native uintptr
	TvSec  int64
	TvUsec int64
}

var addTimeValInvoker *gi.Function

// Add is a representation of the C type g_time_val_add.
func (recv *TimeVal) Add(microseconds int64) {
	if addTimeValInvoker == nil {
		addTimeValInvoker = gi.StructFunctionInvokerNew("GLib", "TimeVal", "add")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(microseconds)

	addTimeValInvoker.Invoke(inArgs[:], nil)

}

var toIso8601TimeValInvoker *gi.Function

// ToIso8601 is a representation of the C type g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() string {
	if toIso8601TimeValInvoker == nil {
		toIso8601TimeValInvoker = gi.StructFunctionInvokerNew("GLib", "TimeVal", "to_iso8601")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toIso8601TimeValInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var TimeZoneStruct *gi.Struct
var TimeZoneStructOnce sync.Once

func TimeZoneStructSet() {
	TimeZoneStructOnce.Do(func() {
		TimeZoneStruct = gi.StructNew("GLib", "TimeZone")
	})
}

type TimeZone struct {
	native uintptr
}

var newTimeZoneInvoker *gi.Function

// TimeZoneNew is a representation of the C type g_time_zone_new.
func TimeZoneNew(identifier string) *TimeZone {
	if newTimeZoneInvoker == nil {
		newTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(identifier)

	ret := newTimeZoneInvoker.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var newLocalTimeZoneInvoker *gi.Function

// TimeZoneNewLocal is a representation of the C type g_time_zone_new_local.
func TimeZoneNewLocal() *TimeZone {
	if newLocalTimeZoneInvoker == nil {
		newLocalTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "new_local")
	}

	ret := newLocalTimeZoneInvoker.Invoke(nil, nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var newOffsetTimeZoneInvoker *gi.Function

// TimeZoneNewOffset is a representation of the C type g_time_zone_new_offset.
func TimeZoneNewOffset(seconds int32) *TimeZone {
	if newOffsetTimeZoneInvoker == nil {
		newOffsetTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "new_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(seconds)

	ret := newOffsetTimeZoneInvoker.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var newUtcTimeZoneInvoker *gi.Function

// TimeZoneNewUtc is a representation of the C type g_time_zone_new_utc.
func TimeZoneNewUtc() *TimeZone {
	if newUtcTimeZoneInvoker == nil {
		newUtcTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "new_utc")
	}

	ret := newUtcTimeZoneInvoker.Invoke(nil, nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_time_zone_adjust_time' : parameter 'type' of type 'TimeType' not supported

// UNSUPPORTED : C value 'g_time_zone_find_interval' : parameter 'type' of type 'TimeType' not supported

var getAbbreviationTimeZoneInvoker *gi.Function

// GetAbbreviation is a representation of the C type g_time_zone_get_abbreviation.
func (recv *TimeZone) GetAbbreviation(interval int32) string {
	if getAbbreviationTimeZoneInvoker == nil {
		getAbbreviationTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "get_abbreviation")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	ret := getAbbreviationTimeZoneInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getIdentifierTimeZoneInvoker *gi.Function

// GetIdentifier is a representation of the C type g_time_zone_get_identifier.
func (recv *TimeZone) GetIdentifier() string {
	if getIdentifierTimeZoneInvoker == nil {
		getIdentifierTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "get_identifier")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIdentifierTimeZoneInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getOffsetTimeZoneInvoker *gi.Function

// GetOffset is a representation of the C type g_time_zone_get_offset.
func (recv *TimeZone) GetOffset(interval int32) int32 {
	if getOffsetTimeZoneInvoker == nil {
		getOffsetTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "get_offset")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	ret := getOffsetTimeZoneInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_time_zone_is_dst' : return type 'gboolean' not supported

var refTimeZoneInvoker *gi.Function

// Ref is a representation of the C type g_time_zone_ref.
func (recv *TimeZone) Ref() *TimeZone {
	if refTimeZoneInvoker == nil {
		refTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refTimeZoneInvoker.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var unrefTimeZoneInvoker *gi.Function

// Unref is a representation of the C type g_time_zone_unref.
func (recv *TimeZone) Unref() {
	if unrefTimeZoneInvoker == nil {
		unrefTimeZoneInvoker = gi.StructFunctionInvokerNew("GLib", "TimeZone", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefTimeZoneInvoker.Invoke(inArgs[:], nil)

}

var TimerStruct *gi.Struct
var TimerStructOnce sync.Once

func TimerStructSet() {
	TimerStructOnce.Do(func() {
		TimerStruct = gi.StructNew("GLib", "Timer")
	})
}

type Timer struct {
	native uintptr
}

var continueTimerInvoker *gi.Function

// Continue is a representation of the C type g_timer_continue.
func (recv *Timer) Continue() {
	if continueTimerInvoker == nil {
		continueTimerInvoker = gi.StructFunctionInvokerNew("GLib", "Timer", "continue")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	continueTimerInvoker.Invoke(inArgs[:], nil)

}

var destroyTimerInvoker *gi.Function

// Destroy is a representation of the C type g_timer_destroy.
func (recv *Timer) Destroy() {
	if destroyTimerInvoker == nil {
		destroyTimerInvoker = gi.StructFunctionInvokerNew("GLib", "Timer", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyTimerInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_timer_elapsed' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_timer_is_active' : return type 'gboolean' not supported

var resetTimerInvoker *gi.Function

// Reset is a representation of the C type g_timer_reset.
func (recv *Timer) Reset() {
	if resetTimerInvoker == nil {
		resetTimerInvoker = gi.StructFunctionInvokerNew("GLib", "Timer", "reset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	resetTimerInvoker.Invoke(inArgs[:], nil)

}

var startTimerInvoker *gi.Function

// Start is a representation of the C type g_timer_start.
func (recv *Timer) Start() {
	if startTimerInvoker == nil {
		startTimerInvoker = gi.StructFunctionInvokerNew("GLib", "Timer", "start")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	startTimerInvoker.Invoke(inArgs[:], nil)

}

var stopTimerInvoker *gi.Function

// Stop is a representation of the C type g_timer_stop.
func (recv *Timer) Stop() {
	if stopTimerInvoker == nil {
		stopTimerInvoker = gi.StructFunctionInvokerNew("GLib", "Timer", "stop")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stopTimerInvoker.Invoke(inArgs[:], nil)

}

var TrashStackStruct *gi.Struct
var TrashStackStructOnce sync.Once

func TrashStackStructSet() {
	TrashStackStructOnce.Do(func() {
		TrashStackStruct = gi.StructNew("GLib", "TrashStack")
	})
}

type TrashStack struct {
	native uintptr
	Next   *TrashStack
}

var TreeStruct *gi.Struct
var TreeStructOnce sync.Once

func TreeStructSet() {
	TreeStructOnce.Do(func() {
		TreeStruct = gi.StructNew("GLib", "Tree")
	})
}

type Tree struct {
	native uintptr
}

var destroyTreeInvoker *gi.Function

// Destroy is a representation of the C type g_tree_destroy.
func (recv *Tree) Destroy() {
	if destroyTreeInvoker == nil {
		destroyTreeInvoker = gi.StructFunctionInvokerNew("GLib", "Tree", "destroy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	destroyTreeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_tree_foreach' : parameter 'func' of type 'TraverseFunc' not supported

var heightTreeInvoker *gi.Function

// Height is a representation of the C type g_tree_height.
func (recv *Tree) Height() int32 {
	if heightTreeInvoker == nil {
		heightTreeInvoker = gi.StructFunctionInvokerNew("GLib", "Tree", "height")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := heightTreeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_tree_insert' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_lookup' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_lookup_extended' : parameter 'lookup_key' of type 'gpointer' not supported

var nnodesTreeInvoker *gi.Function

// Nnodes is a representation of the C type g_tree_nnodes.
func (recv *Tree) Nnodes() int32 {
	if nnodesTreeInvoker == nil {
		nnodesTreeInvoker = gi.StructFunctionInvokerNew("GLib", "Tree", "nnodes")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nnodesTreeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var refTreeInvoker *gi.Function

// Ref is a representation of the C type g_tree_ref.
func (recv *Tree) Ref() *Tree {
	if refTreeInvoker == nil {
		refTreeInvoker = gi.StructFunctionInvokerNew("GLib", "Tree", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refTreeInvoker.Invoke(inArgs[:], nil)

	retGo := &Tree{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_tree_remove' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_replace' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_search' : parameter 'search_func' of type 'CompareFunc' not supported

// UNSUPPORTED : C value 'g_tree_steal' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_traverse' : parameter 'traverse_func' of type 'TraverseFunc' not supported

var unrefTreeInvoker *gi.Function

// Unref is a representation of the C type g_tree_unref.
func (recv *Tree) Unref() {
	if unrefTreeInvoker == nil {
		unrefTreeInvoker = gi.StructFunctionInvokerNew("GLib", "Tree", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefTreeInvoker.Invoke(inArgs[:], nil)

}

var VariantStruct *gi.Struct
var VariantStructOnce sync.Once

func VariantStructSet() {
	VariantStructOnce.Do(func() {
		VariantStruct = gi.StructNew("GLib", "Variant")
	})
}

type Variant struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_new' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_array' : parameter 'child_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_boolean' : parameter 'value' of type 'gboolean' not supported

var newByteVariantInvoker *gi.Function

// VariantNewByte is a representation of the C type g_variant_new_byte.
func VariantNewByte(value uint8) *Variant {
	if newByteVariantInvoker == nil {
		newByteVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_byte")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint8(value)

	ret := newByteVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_bytestring' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_variant_new_bytestring_array' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_dict_entry' : parameter 'key' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_new_double' : parameter 'value' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_variant_new_fixed_array' : parameter 'element_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_from_bytes' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_from_data' : parameter 'type' of type 'VariantType' not supported

var newHandleVariantInvoker *gi.Function

// VariantNewHandle is a representation of the C type g_variant_new_handle.
func VariantNewHandle(value int32) *Variant {
	if newHandleVariantInvoker == nil {
		newHandleVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_handle")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	ret := newHandleVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var newInt16VariantInvoker *gi.Function

// VariantNewInt16 is a representation of the C type g_variant_new_int16.
func VariantNewInt16(value int16) *Variant {
	if newInt16VariantInvoker == nil {
		newInt16VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_int16")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt16(value)

	ret := newInt16VariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var newInt32VariantInvoker *gi.Function

// VariantNewInt32 is a representation of the C type g_variant_new_int32.
func VariantNewInt32(value int32) *Variant {
	if newInt32VariantInvoker == nil {
		newInt32VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_int32")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	ret := newInt32VariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var newInt64VariantInvoker *gi.Function

// VariantNewInt64 is a representation of the C type g_variant_new_int64.
func VariantNewInt64(value int64) *Variant {
	if newInt64VariantInvoker == nil {
		newInt64VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_int64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(value)

	ret := newInt64VariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_maybe' : parameter 'child_type' of type 'VariantType' not supported

var newObjectPathVariantInvoker *gi.Function

// VariantNewObjectPath is a representation of the C type g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) *Variant {
	if newObjectPathVariantInvoker == nil {
		newObjectPathVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_object_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(objectPath)

	ret := newObjectPathVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_objv' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_printf' : parameter '...' has no type

var newSignatureVariantInvoker *gi.Function

// VariantNewSignature is a representation of the C type g_variant_new_signature.
func VariantNewSignature(signature string) *Variant {
	if newSignatureVariantInvoker == nil {
		newSignatureVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_signature")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(signature)

	ret := newSignatureVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var newStringVariantInvoker *gi.Function

// VariantNewString is a representation of the C type g_variant_new_string.
func VariantNewString(string_ string) *Variant {
	if newStringVariantInvoker == nil {
		newStringVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := newStringVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_strv' : parameter 'strv' has no type

var newTakeStringVariantInvoker *gi.Function

// VariantNewTakeString is a representation of the C type g_variant_new_take_string.
func VariantNewTakeString(string_ string) *Variant {
	if newTakeStringVariantInvoker == nil {
		newTakeStringVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_take_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := newTakeStringVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_tuple' : parameter 'children' has no type

var newUint16VariantInvoker *gi.Function

// VariantNewUint16 is a representation of the C type g_variant_new_uint16.
func VariantNewUint16(value uint16) *Variant {
	if newUint16VariantInvoker == nil {
		newUint16VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_uint16")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(value)

	ret := newUint16VariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var newUint32VariantInvoker *gi.Function

// VariantNewUint32 is a representation of the C type g_variant_new_uint32.
func VariantNewUint32(value uint32) *Variant {
	if newUint32VariantInvoker == nil {
		newUint32VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_uint32")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(value)

	ret := newUint32VariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var newUint64VariantInvoker *gi.Function

// VariantNewUint64 is a representation of the C type g_variant_new_uint64.
func VariantNewUint64(value uint64) *Variant {
	if newUint64VariantInvoker == nil {
		newUint64VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "new_uint64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(value)

	ret := newUint64VariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_variant' : parameter 'value' of type 'Variant' not supported

var byteswapVariantInvoker *gi.Function

// Byteswap is a representation of the C type g_variant_byteswap.
func (recv *Variant) Byteswap() *Variant {
	if byteswapVariantInvoker == nil {
		byteswapVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "byteswap")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := byteswapVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_check_format_string' : parameter 'copy_only' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_classify' : return type 'VariantClass' not supported

// UNSUPPORTED : C value 'g_variant_compare' : parameter 'two' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dup_bytestring' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_bytestring_array' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_objv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_string' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_strv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_equal' : parameter 'two' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_get' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_get_boolean' : return type 'gboolean' not supported

var getByteVariantInvoker *gi.Function

// GetByte is a representation of the C type g_variant_get_byte.
func (recv *Variant) GetByte() uint8 {
	if getByteVariantInvoker == nil {
		getByteVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_byte")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getByteVariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint8()

	return retGo
}

var getBytestringVariantInvoker *gi.Function

// GetBytestring is a representation of the C type g_variant_get_bytestring.
func (recv *Variant) GetBytestring() {
	if getBytestringVariantInvoker == nil {
		getBytestringVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_bytestring")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	getBytestringVariantInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_get_bytestring_array' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_child' : parameter 'index_' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_child_value' : parameter 'index_' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_data' : return type 'gpointer' not supported

var getDataAsBytesVariantInvoker *gi.Function

// GetDataAsBytes is a representation of the C type g_variant_get_data_as_bytes.
func (recv *Variant) GetDataAsBytes() *Bytes {
	if getDataAsBytesVariantInvoker == nil {
		getDataAsBytesVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_data_as_bytes")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDataAsBytesVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_variant_get_fixed_array' : parameter 'n_elements' of type 'gsize' not supported

var getHandleVariantInvoker *gi.Function

// GetHandle is a representation of the C type g_variant_get_handle.
func (recv *Variant) GetHandle() int32 {
	if getHandleVariantInvoker == nil {
		getHandleVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_handle")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getHandleVariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getInt16VariantInvoker *gi.Function

// GetInt16 is a representation of the C type g_variant_get_int16.
func (recv *Variant) GetInt16() int16 {
	if getInt16VariantInvoker == nil {
		getInt16VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_int16")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getInt16VariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int16()

	return retGo
}

var getInt32VariantInvoker *gi.Function

// GetInt32 is a representation of the C type g_variant_get_int32.
func (recv *Variant) GetInt32() int32 {
	if getInt32VariantInvoker == nil {
		getInt32VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_int32")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getInt32VariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getInt64VariantInvoker *gi.Function

// GetInt64 is a representation of the C type g_variant_get_int64.
func (recv *Variant) GetInt64() int64 {
	if getInt64VariantInvoker == nil {
		getInt64VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_int64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getInt64VariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getMaybeVariantInvoker *gi.Function

// GetMaybe is a representation of the C type g_variant_get_maybe.
func (recv *Variant) GetMaybe() *Variant {
	if getMaybeVariantInvoker == nil {
		getMaybeVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_maybe")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMaybeVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var getNormalFormVariantInvoker *gi.Function

// GetNormalForm is a representation of the C type g_variant_get_normal_form.
func (recv *Variant) GetNormalForm() *Variant {
	if getNormalFormVariantInvoker == nil {
		getNormalFormVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_normal_form")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNormalFormVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_objv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_size' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_string' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_strv' : parameter 'length' of type 'gsize' not supported

var getTypeVariantInvoker *gi.Function

// GetType is a representation of the C type g_variant_get_type.
func (recv *Variant) GetType() *VariantType {
	if getTypeVariantInvoker == nil {
		getTypeVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_type")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTypeVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var getTypeStringVariantInvoker *gi.Function

// GetTypeString is a representation of the C type g_variant_get_type_string.
func (recv *Variant) GetTypeString() string {
	if getTypeStringVariantInvoker == nil {
		getTypeStringVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_type_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTypeStringVariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getUint16VariantInvoker *gi.Function

// GetUint16 is a representation of the C type g_variant_get_uint16.
func (recv *Variant) GetUint16() uint16 {
	if getUint16VariantInvoker == nil {
		getUint16VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_uint16")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUint16VariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var getUint32VariantInvoker *gi.Function

// GetUint32 is a representation of the C type g_variant_get_uint32.
func (recv *Variant) GetUint32() uint32 {
	if getUint32VariantInvoker == nil {
		getUint32VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_uint32")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUint32VariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getUint64VariantInvoker *gi.Function

// GetUint64 is a representation of the C type g_variant_get_uint64.
func (recv *Variant) GetUint64() uint64 {
	if getUint64VariantInvoker == nil {
		getUint64VariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_uint64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUint64VariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_va' : parameter 'app' of type 'va_list' not supported

var getVariantVariantInvoker *gi.Function

// GetVariant is a representation of the C type g_variant_get_variant.
func (recv *Variant) GetVariant() *Variant {
	if getVariantVariantInvoker == nil {
		getVariantVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "get_variant")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getVariantVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var hashVariantInvoker *gi.Function

// Hash is a representation of the C type g_variant_hash.
func (recv *Variant) Hash() uint32 {
	if hashVariantInvoker == nil {
		hashVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashVariantInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_variant_is_container' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_floating' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_normal_form' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_of_type' : parameter 'type' of type 'VariantType' not supported

var iterNewVariantInvoker *gi.Function

// IterNew is a representation of the C type g_variant_iter_new.
func (recv *Variant) IterNew() *VariantIter {
	if iterNewVariantInvoker == nil {
		iterNewVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "iter_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iterNewVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantIter{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_lookup' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_lookup_value' : parameter 'expected_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_n_children' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_print' : parameter 'type_annotate' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_print_string' : parameter 'string' of type 'String' not supported

var refVariantInvoker *gi.Function

// Ref is a representation of the C type g_variant_ref.
func (recv *Variant) Ref() *Variant {
	if refVariantInvoker == nil {
		refVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var refSinkVariantInvoker *gi.Function

// RefSink is a representation of the C type g_variant_ref_sink.
func (recv *Variant) RefSink() *Variant {
	if refSinkVariantInvoker == nil {
		refSinkVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "ref_sink")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSinkVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_store' : parameter 'data' of type 'gpointer' not supported

var takeRefVariantInvoker *gi.Function

// TakeRef is a representation of the C type g_variant_take_ref.
func (recv *Variant) TakeRef() *Variant {
	if takeRefVariantInvoker == nil {
		takeRefVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "take_ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := takeRefVariantInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var unrefVariantInvoker *gi.Function

// Unref is a representation of the C type g_variant_unref.
func (recv *Variant) Unref() {
	if unrefVariantInvoker == nil {
		unrefVariantInvoker = gi.StructFunctionInvokerNew("GLib", "Variant", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefVariantInvoker.Invoke(inArgs[:], nil)

}

var VariantBuilderStruct *gi.Struct
var VariantBuilderStructOnce sync.Once

func VariantBuilderStructSet() {
	VariantBuilderStructOnce.Do(func() {
		VariantBuilderStruct = gi.StructNew("GLib", "VariantBuilder")
	})
}

type VariantBuilder struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_builder_new' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_builder_add' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_builder_add_parsed' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_builder_add_value' : parameter 'value' of type 'Variant' not supported

var clearVariantBuilderInvoker *gi.Function

// Clear is a representation of the C type g_variant_builder_clear.
func (recv *VariantBuilder) Clear() {
	if clearVariantBuilderInvoker == nil {
		clearVariantBuilderInvoker = gi.StructFunctionInvokerNew("GLib", "VariantBuilder", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearVariantBuilderInvoker.Invoke(inArgs[:], nil)

}

var closeVariantBuilderInvoker *gi.Function

// Close is a representation of the C type g_variant_builder_close.
func (recv *VariantBuilder) Close() {
	if closeVariantBuilderInvoker == nil {
		closeVariantBuilderInvoker = gi.StructFunctionInvokerNew("GLib", "VariantBuilder", "close")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	closeVariantBuilderInvoker.Invoke(inArgs[:], nil)

}

var endVariantBuilderInvoker *gi.Function

// End is a representation of the C type g_variant_builder_end.
func (recv *VariantBuilder) End() *Variant {
	if endVariantBuilderInvoker == nil {
		endVariantBuilderInvoker = gi.StructFunctionInvokerNew("GLib", "VariantBuilder", "end")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := endVariantBuilderInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_builder_init' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_builder_open' : parameter 'type' of type 'VariantType' not supported

var refVariantBuilderInvoker *gi.Function

// Ref is a representation of the C type g_variant_builder_ref.
func (recv *VariantBuilder) Ref() *VariantBuilder {
	if refVariantBuilderInvoker == nil {
		refVariantBuilderInvoker = gi.StructFunctionInvokerNew("GLib", "VariantBuilder", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refVariantBuilderInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantBuilder{native: ret.Pointer()}

	return retGo
}

var unrefVariantBuilderInvoker *gi.Function

// Unref is a representation of the C type g_variant_builder_unref.
func (recv *VariantBuilder) Unref() {
	if unrefVariantBuilderInvoker == nil {
		unrefVariantBuilderInvoker = gi.StructFunctionInvokerNew("GLib", "VariantBuilder", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefVariantBuilderInvoker.Invoke(inArgs[:], nil)

}

var VariantDictStruct *gi.Struct
var VariantDictStructOnce sync.Once

func VariantDictStructSet() {
	VariantDictStructOnce.Do(func() {
		VariantDictStruct = gi.StructNew("GLib", "VariantDict")
	})
}

type VariantDict struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_dict_new' : parameter 'from_asv' of type 'Variant' not supported

var clearVariantDictInvoker *gi.Function

// Clear is a representation of the C type g_variant_dict_clear.
func (recv *VariantDict) Clear() {
	if clearVariantDictInvoker == nil {
		clearVariantDictInvoker = gi.StructFunctionInvokerNew("GLib", "VariantDict", "clear")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	clearVariantDictInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_dict_contains' : return type 'gboolean' not supported

var endVariantDictInvoker *gi.Function

// End is a representation of the C type g_variant_dict_end.
func (recv *VariantDict) End() *Variant {
	if endVariantDictInvoker == nil {
		endVariantDictInvoker = gi.StructFunctionInvokerNew("GLib", "VariantDict", "end")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := endVariantDictInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_dict_init' : parameter 'from_asv' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dict_insert' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_dict_insert_value' : parameter 'value' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dict_lookup' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_dict_lookup_value' : parameter 'expected_type' of type 'VariantType' not supported

var refVariantDictInvoker *gi.Function

// Ref is a representation of the C type g_variant_dict_ref.
func (recv *VariantDict) Ref() *VariantDict {
	if refVariantDictInvoker == nil {
		refVariantDictInvoker = gi.StructFunctionInvokerNew("GLib", "VariantDict", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refVariantDictInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantDict{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_dict_remove' : return type 'gboolean' not supported

var unrefVariantDictInvoker *gi.Function

// Unref is a representation of the C type g_variant_dict_unref.
func (recv *VariantDict) Unref() {
	if unrefVariantDictInvoker == nil {
		unrefVariantDictInvoker = gi.StructFunctionInvokerNew("GLib", "VariantDict", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefVariantDictInvoker.Invoke(inArgs[:], nil)

}

var VariantIterStruct *gi.Struct
var VariantIterStructOnce sync.Once

func VariantIterStructSet() {
	VariantIterStructOnce.Do(func() {
		VariantIterStruct = gi.StructNew("GLib", "VariantIter")
	})
}

type VariantIter struct {
	native uintptr
}

var copyVariantIterInvoker *gi.Function

// Copy is a representation of the C type g_variant_iter_copy.
func (recv *VariantIter) Copy() *VariantIter {
	if copyVariantIterInvoker == nil {
		copyVariantIterInvoker = gi.StructFunctionInvokerNew("GLib", "VariantIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyVariantIterInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantIter{native: ret.Pointer()}

	return retGo
}

var freeVariantIterInvoker *gi.Function

// Free is a representation of the C type g_variant_iter_free.
func (recv *VariantIter) Free() {
	if freeVariantIterInvoker == nil {
		freeVariantIterInvoker = gi.StructFunctionInvokerNew("GLib", "VariantIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeVariantIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_iter_init' : parameter 'value' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_iter_loop' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_iter_n_children' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_iter_next' : parameter '...' has no type

var nextValueVariantIterInvoker *gi.Function

// NextValue is a representation of the C type g_variant_iter_next_value.
func (recv *VariantIter) NextValue() *Variant {
	if nextValueVariantIterInvoker == nil {
		nextValueVariantIterInvoker = gi.StructFunctionInvokerNew("GLib", "VariantIter", "next_value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nextValueVariantIterInvoker.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var VariantTypeStruct *gi.Struct
var VariantTypeStructOnce sync.Once

func VariantTypeStructSet() {
	VariantTypeStructOnce.Do(func() {
		VariantTypeStruct = gi.StructNew("GLib", "VariantType")
	})
}

type VariantType struct {
	native uintptr
}

var newVariantTypeInvoker *gi.Function

// VariantTypeNew is a representation of the C type g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	if newVariantTypeInvoker == nil {
		newVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	ret := newVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_new_array' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_dict_entry' : parameter 'key' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_maybe' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_tuple' : parameter 'items' has no type

var copyVariantTypeInvoker *gi.Function

// Copy is a representation of the C type g_variant_type_copy.
func (recv *VariantType) Copy() *VariantType {
	if copyVariantTypeInvoker == nil {
		copyVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var dupStringVariantTypeInvoker *gi.Function

// DupString is a representation of the C type g_variant_type_dup_string.
func (recv *VariantType) DupString() string {
	if dupStringVariantTypeInvoker == nil {
		dupStringVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "dup_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dupStringVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var elementVariantTypeInvoker *gi.Function

// Element is a representation of the C type g_variant_type_element.
func (recv *VariantType) Element() *VariantType {
	if elementVariantTypeInvoker == nil {
		elementVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "element")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := elementVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_equal' : parameter 'type2' of type 'VariantType' not supported

var firstVariantTypeInvoker *gi.Function

// First is a representation of the C type g_variant_type_first.
func (recv *VariantType) First() *VariantType {
	if firstVariantTypeInvoker == nil {
		firstVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "first")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := firstVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var freeVariantTypeInvoker *gi.Function

// Free is a representation of the C type g_variant_type_free.
func (recv *VariantType) Free() {
	if freeVariantTypeInvoker == nil {
		freeVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeVariantTypeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_type_get_string_length' : return type 'gsize' not supported

var hashVariantTypeInvoker *gi.Function

// Hash is a representation of the C type g_variant_type_hash.
func (recv *VariantType) Hash() uint32 {
	if hashVariantTypeInvoker == nil {
		hashVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "hash")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hashVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_is_array' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_basic' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_container' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_definite' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_dict_entry' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_maybe' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_subtype_of' : parameter 'supertype' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_is_tuple' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_type_is_variant' : return type 'gboolean' not supported

var keyVariantTypeInvoker *gi.Function

// Key is a representation of the C type g_variant_type_key.
func (recv *VariantType) Key() *VariantType {
	if keyVariantTypeInvoker == nil {
		keyVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "key")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := keyVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_n_items' : return type 'gsize' not supported

var nextVariantTypeInvoker *gi.Function

// Next is a representation of the C type g_variant_type_next.
func (recv *VariantType) Next() *VariantType {
	if nextVariantTypeInvoker == nil {
		nextVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "next")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nextVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var peekStringVariantTypeInvoker *gi.Function

// PeekString is a representation of the C type g_variant_type_peek_string.
func (recv *VariantType) PeekString() string {
	if peekStringVariantTypeInvoker == nil {
		peekStringVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "peek_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := peekStringVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var valueVariantTypeInvoker *gi.Function

// Value is a representation of the C type g_variant_type_value.
func (recv *VariantType) Value() *VariantType {
	if valueVariantTypeInvoker == nil {
		valueVariantTypeInvoker = gi.StructFunctionInvokerNew("GLib", "VariantType", "value")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := valueVariantTypeInvoker.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}
