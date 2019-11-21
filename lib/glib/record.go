// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var arrayStruct *gi.Struct
var arrayStructOnce sync.Once

func arrayStructSet() {
	arrayStructOnce.Do(func() {
		arrayStruct = gi.StructNew("GLib", "Array")
	})
}

type Array struct {
	native uintptr
	Data   string
	Len    uint32
}

var asyncQueueStruct *gi.Struct
var asyncQueueStructOnce sync.Once

func asyncQueueStructSet() {
	asyncQueueStructOnce.Do(func() {
		asyncQueueStruct = gi.StructNew("GLib", "AsyncQueue")
	})
}

type AsyncQueue struct {
	native uintptr
}

var asyncQueueLengthFunction *gi.Function
var asyncQueueLengthFunction_Once sync.Once

func asyncQueueLengthFunction_Set() {
	asyncQueueLengthFunction_Once.Do(func() {
		asyncQueueLengthFunction = gi.FunctionInvokerNew("GLib", "length")
	})
}

var lengthAsyncQueueInvoker *gi.Function

// Length is a representation of the C type g_async_queue_length.
func (recv *AsyncQueue) Length() int32 {
	asyncQueueLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := asyncQueueLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var asyncQueueLengthUnlockedFunction *gi.Function
var asyncQueueLengthUnlockedFunction_Once sync.Once

func asyncQueueLengthUnlockedFunction_Set() {
	asyncQueueLengthUnlockedFunction_Once.Do(func() {
		asyncQueueLengthUnlockedFunction = gi.FunctionInvokerNew("GLib", "length_unlocked")
	})
}

var lengthUnlockedAsyncQueueInvoker *gi.Function

// LengthUnlocked is a representation of the C type g_async_queue_length_unlocked.
func (recv *AsyncQueue) LengthUnlocked() int32 {
	asyncQueueLengthUnlockedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := asyncQueueLengthUnlockedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var asyncQueueLockFunction *gi.Function
var asyncQueueLockFunction_Once sync.Once

func asyncQueueLockFunction_Set() {
	asyncQueueLockFunction_Once.Do(func() {
		asyncQueueLockFunction = gi.FunctionInvokerNew("GLib", "lock")
	})
}

var lockAsyncQueueInvoker *gi.Function

// Lock is a representation of the C type g_async_queue_lock.
func (recv *AsyncQueue) Lock() {
	asyncQueueLockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	asyncQueueLockFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_async_queue_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_pop_unlocked' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_front' : parameter 'item' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_front_unlocked' : parameter 'item' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_sorted' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_sorted_unlocked' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_async_queue_push_unlocked' : parameter 'data' of type 'gpointer' not supported

var asyncQueueRefFunction *gi.Function
var asyncQueueRefFunction_Once sync.Once

func asyncQueueRefFunction_Set() {
	asyncQueueRefFunction_Once.Do(func() {
		asyncQueueRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refAsyncQueueInvoker *gi.Function

// Ref is a representation of the C type g_async_queue_ref.
func (recv *AsyncQueue) Ref() *AsyncQueue {
	asyncQueueRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := asyncQueueRefFunction.Invoke(inArgs[:], nil)

	retGo := &AsyncQueue{native: ret.Pointer()}

	return retGo
}

var asyncQueueRefUnlockedFunction *gi.Function
var asyncQueueRefUnlockedFunction_Once sync.Once

func asyncQueueRefUnlockedFunction_Set() {
	asyncQueueRefUnlockedFunction_Once.Do(func() {
		asyncQueueRefUnlockedFunction = gi.FunctionInvokerNew("GLib", "ref_unlocked")
	})
}

var refUnlockedAsyncQueueInvoker *gi.Function

// RefUnlocked is a representation of the C type g_async_queue_ref_unlocked.
func (recv *AsyncQueue) RefUnlocked() {
	asyncQueueRefUnlockedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	asyncQueueRefUnlockedFunction.Invoke(inArgs[:], nil)

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

var asyncQueueUnlockFunction *gi.Function
var asyncQueueUnlockFunction_Once sync.Once

func asyncQueueUnlockFunction_Set() {
	asyncQueueUnlockFunction_Once.Do(func() {
		asyncQueueUnlockFunction = gi.FunctionInvokerNew("GLib", "unlock")
	})
}

var unlockAsyncQueueInvoker *gi.Function

// Unlock is a representation of the C type g_async_queue_unlock.
func (recv *AsyncQueue) Unlock() {
	asyncQueueUnlockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	asyncQueueUnlockFunction.Invoke(inArgs[:], nil)

}

var asyncQueueUnrefFunction *gi.Function
var asyncQueueUnrefFunction_Once sync.Once

func asyncQueueUnrefFunction_Set() {
	asyncQueueUnrefFunction_Once.Do(func() {
		asyncQueueUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefAsyncQueueInvoker *gi.Function

// Unref is a representation of the C type g_async_queue_unref.
func (recv *AsyncQueue) Unref() {
	asyncQueueUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	asyncQueueUnrefFunction.Invoke(inArgs[:], nil)

}

var asyncQueueUnrefAndUnlockFunction *gi.Function
var asyncQueueUnrefAndUnlockFunction_Once sync.Once

func asyncQueueUnrefAndUnlockFunction_Set() {
	asyncQueueUnrefAndUnlockFunction_Once.Do(func() {
		asyncQueueUnrefAndUnlockFunction = gi.FunctionInvokerNew("GLib", "unref_and_unlock")
	})
}

var unrefAndUnlockAsyncQueueInvoker *gi.Function

// UnrefAndUnlock is a representation of the C type g_async_queue_unref_and_unlock.
func (recv *AsyncQueue) UnrefAndUnlock() {
	asyncQueueUnrefAndUnlockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	asyncQueueUnrefAndUnlockFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileStruct *gi.Struct
var bookmarkFileStructOnce sync.Once

func bookmarkFileStructSet() {
	bookmarkFileStructOnce.Do(func() {
		bookmarkFileStruct = gi.StructNew("GLib", "BookmarkFile")
	})
}

type BookmarkFile struct {
	native uintptr
}

var bookmarkFileAddApplicationFunction *gi.Function
var bookmarkFileAddApplicationFunction_Once sync.Once

func bookmarkFileAddApplicationFunction_Set() {
	bookmarkFileAddApplicationFunction_Once.Do(func() {
		bookmarkFileAddApplicationFunction = gi.FunctionInvokerNew("GLib", "add_application")
	})
}

var addApplicationBookmarkFileInvoker *gi.Function

// AddApplication is a representation of the C type g_bookmark_file_add_application.
func (recv *BookmarkFile) AddApplication(uri string, name string, exec string) {
	bookmarkFileAddApplicationFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)
	inArgs[3].SetString(exec)

	bookmarkFileAddApplicationFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileAddGroupFunction *gi.Function
var bookmarkFileAddGroupFunction_Once sync.Once

func bookmarkFileAddGroupFunction_Set() {
	bookmarkFileAddGroupFunction_Once.Do(func() {
		bookmarkFileAddGroupFunction = gi.FunctionInvokerNew("GLib", "add_group")
	})
}

var addGroupBookmarkFileInvoker *gi.Function

// AddGroup is a representation of the C type g_bookmark_file_add_group.
func (recv *BookmarkFile) AddGroup(uri string, group string) {
	bookmarkFileAddGroupFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(group)

	bookmarkFileAddGroupFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileFreeFunction *gi.Function
var bookmarkFileFreeFunction_Once sync.Once

func bookmarkFileFreeFunction_Set() {
	bookmarkFileFreeFunction_Once.Do(func() {
		bookmarkFileFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeBookmarkFileInvoker *gi.Function

// Free is a representation of the C type g_bookmark_file_free.
func (recv *BookmarkFile) Free() {
	bookmarkFileFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	bookmarkFileFreeFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileGetAddedFunction *gi.Function
var bookmarkFileGetAddedFunction_Once sync.Once

func bookmarkFileGetAddedFunction_Set() {
	bookmarkFileGetAddedFunction_Once.Do(func() {
		bookmarkFileGetAddedFunction = gi.FunctionInvokerNew("GLib", "get_added")
	})
}

var getAddedBookmarkFileInvoker *gi.Function

// GetAdded is a representation of the C type g_bookmark_file_get_added.
func (recv *BookmarkFile) GetAdded(uri string) int64 {
	bookmarkFileGetAddedFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := bookmarkFileGetAddedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_get_app_info' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_get_applications' : parameter 'length' of type 'gsize' not supported

var bookmarkFileGetDescriptionFunction *gi.Function
var bookmarkFileGetDescriptionFunction_Once sync.Once

func bookmarkFileGetDescriptionFunction_Set() {
	bookmarkFileGetDescriptionFunction_Once.Do(func() {
		bookmarkFileGetDescriptionFunction = gi.FunctionInvokerNew("GLib", "get_description")
	})
}

var getDescriptionBookmarkFileInvoker *gi.Function

// GetDescription is a representation of the C type g_bookmark_file_get_description.
func (recv *BookmarkFile) GetDescription(uri string) string {
	bookmarkFileGetDescriptionFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := bookmarkFileGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_get_groups' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_bookmark_file_get_icon' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_bookmark_file_get_is_private' : return type 'gboolean' not supported

var bookmarkFileGetMimeTypeFunction *gi.Function
var bookmarkFileGetMimeTypeFunction_Once sync.Once

func bookmarkFileGetMimeTypeFunction_Set() {
	bookmarkFileGetMimeTypeFunction_Once.Do(func() {
		bookmarkFileGetMimeTypeFunction = gi.FunctionInvokerNew("GLib", "get_mime_type")
	})
}

var getMimeTypeBookmarkFileInvoker *gi.Function

// GetMimeType is a representation of the C type g_bookmark_file_get_mime_type.
func (recv *BookmarkFile) GetMimeType(uri string) string {
	bookmarkFileGetMimeTypeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := bookmarkFileGetMimeTypeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var bookmarkFileGetModifiedFunction *gi.Function
var bookmarkFileGetModifiedFunction_Once sync.Once

func bookmarkFileGetModifiedFunction_Set() {
	bookmarkFileGetModifiedFunction_Once.Do(func() {
		bookmarkFileGetModifiedFunction = gi.FunctionInvokerNew("GLib", "get_modified")
	})
}

var getModifiedBookmarkFileInvoker *gi.Function

// GetModified is a representation of the C type g_bookmark_file_get_modified.
func (recv *BookmarkFile) GetModified(uri string) int64 {
	bookmarkFileGetModifiedFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := bookmarkFileGetModifiedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var bookmarkFileGetSizeFunction *gi.Function
var bookmarkFileGetSizeFunction_Once sync.Once

func bookmarkFileGetSizeFunction_Set() {
	bookmarkFileGetSizeFunction_Once.Do(func() {
		bookmarkFileGetSizeFunction = gi.FunctionInvokerNew("GLib", "get_size")
	})
}

var getSizeBookmarkFileInvoker *gi.Function

// GetSize is a representation of the C type g_bookmark_file_get_size.
func (recv *BookmarkFile) GetSize() int32 {
	bookmarkFileGetSizeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := bookmarkFileGetSizeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var bookmarkFileGetTitleFunction *gi.Function
var bookmarkFileGetTitleFunction_Once sync.Once

func bookmarkFileGetTitleFunction_Set() {
	bookmarkFileGetTitleFunction_Once.Do(func() {
		bookmarkFileGetTitleFunction = gi.FunctionInvokerNew("GLib", "get_title")
	})
}

var getTitleBookmarkFileInvoker *gi.Function

// GetTitle is a representation of the C type g_bookmark_file_get_title.
func (recv *BookmarkFile) GetTitle(uri string) string {
	bookmarkFileGetTitleFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := bookmarkFileGetTitleFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_bookmark_file_get_uris' : parameter 'length' of type 'gsize' not supported

var bookmarkFileGetVisitedFunction *gi.Function
var bookmarkFileGetVisitedFunction_Once sync.Once

func bookmarkFileGetVisitedFunction_Set() {
	bookmarkFileGetVisitedFunction_Once.Do(func() {
		bookmarkFileGetVisitedFunction = gi.FunctionInvokerNew("GLib", "get_visited")
	})
}

var getVisitedBookmarkFileInvoker *gi.Function

// GetVisited is a representation of the C type g_bookmark_file_get_visited.
func (recv *BookmarkFile) GetVisited(uri string) int64 {
	bookmarkFileGetVisitedFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	ret := bookmarkFileGetVisitedFunction.Invoke(inArgs[:], nil)

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

var bookmarkFileSetAddedFunction *gi.Function
var bookmarkFileSetAddedFunction_Once sync.Once

func bookmarkFileSetAddedFunction_Set() {
	bookmarkFileSetAddedFunction_Once.Do(func() {
		bookmarkFileSetAddedFunction = gi.FunctionInvokerNew("GLib", "set_added")
	})
}

var setAddedBookmarkFileInvoker *gi.Function

// SetAdded is a representation of the C type g_bookmark_file_set_added.
func (recv *BookmarkFile) SetAdded(uri string, added int64) {
	bookmarkFileSetAddedFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(added)

	bookmarkFileSetAddedFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_set_app_info' : return type 'gboolean' not supported

var bookmarkFileSetDescriptionFunction *gi.Function
var bookmarkFileSetDescriptionFunction_Once sync.Once

func bookmarkFileSetDescriptionFunction_Set() {
	bookmarkFileSetDescriptionFunction_Once.Do(func() {
		bookmarkFileSetDescriptionFunction = gi.FunctionInvokerNew("GLib", "set_description")
	})
}

var setDescriptionBookmarkFileInvoker *gi.Function

// SetDescription is a representation of the C type g_bookmark_file_set_description.
func (recv *BookmarkFile) SetDescription(uri string, description string) {
	bookmarkFileSetDescriptionFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(description)

	bookmarkFileSetDescriptionFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_set_groups' : parameter 'groups' has no type

var bookmarkFileSetIconFunction *gi.Function
var bookmarkFileSetIconFunction_Once sync.Once

func bookmarkFileSetIconFunction_Set() {
	bookmarkFileSetIconFunction_Once.Do(func() {
		bookmarkFileSetIconFunction = gi.FunctionInvokerNew("GLib", "set_icon")
	})
}

var setIconBookmarkFileInvoker *gi.Function

// SetIcon is a representation of the C type g_bookmark_file_set_icon.
func (recv *BookmarkFile) SetIcon(uri string, href string, mimeType string) {
	bookmarkFileSetIconFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(href)
	inArgs[3].SetString(mimeType)

	bookmarkFileSetIconFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_set_is_private' : parameter 'is_private' of type 'gboolean' not supported

var bookmarkFileSetMimeTypeFunction *gi.Function
var bookmarkFileSetMimeTypeFunction_Once sync.Once

func bookmarkFileSetMimeTypeFunction_Set() {
	bookmarkFileSetMimeTypeFunction_Once.Do(func() {
		bookmarkFileSetMimeTypeFunction = gi.FunctionInvokerNew("GLib", "set_mime_type")
	})
}

var setMimeTypeBookmarkFileInvoker *gi.Function

// SetMimeType is a representation of the C type g_bookmark_file_set_mime_type.
func (recv *BookmarkFile) SetMimeType(uri string, mimeType string) {
	bookmarkFileSetMimeTypeFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(mimeType)

	bookmarkFileSetMimeTypeFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileSetModifiedFunction *gi.Function
var bookmarkFileSetModifiedFunction_Once sync.Once

func bookmarkFileSetModifiedFunction_Set() {
	bookmarkFileSetModifiedFunction_Once.Do(func() {
		bookmarkFileSetModifiedFunction = gi.FunctionInvokerNew("GLib", "set_modified")
	})
}

var setModifiedBookmarkFileInvoker *gi.Function

// SetModified is a representation of the C type g_bookmark_file_set_modified.
func (recv *BookmarkFile) SetModified(uri string, modified int64) {
	bookmarkFileSetModifiedFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(modified)

	bookmarkFileSetModifiedFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileSetTitleFunction *gi.Function
var bookmarkFileSetTitleFunction_Once sync.Once

func bookmarkFileSetTitleFunction_Set() {
	bookmarkFileSetTitleFunction_Once.Do(func() {
		bookmarkFileSetTitleFunction = gi.FunctionInvokerNew("GLib", "set_title")
	})
}

var setTitleBookmarkFileInvoker *gi.Function

// SetTitle is a representation of the C type g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) {
	bookmarkFileSetTitleFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(title)

	bookmarkFileSetTitleFunction.Invoke(inArgs[:], nil)

}

var bookmarkFileSetVisitedFunction *gi.Function
var bookmarkFileSetVisitedFunction_Once sync.Once

func bookmarkFileSetVisitedFunction_Set() {
	bookmarkFileSetVisitedFunction_Once.Do(func() {
		bookmarkFileSetVisitedFunction = gi.FunctionInvokerNew("GLib", "set_visited")
	})
}

var setVisitedBookmarkFileInvoker *gi.Function

// SetVisited is a representation of the C type g_bookmark_file_set_visited.
func (recv *BookmarkFile) SetVisited(uri string, visited int64) {
	bookmarkFileSetVisitedFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(visited)

	bookmarkFileSetVisitedFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bookmark_file_to_data' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_bookmark_file_to_file' : parameter 'filename' of type 'filename' not supported

var byteArrayStruct *gi.Struct
var byteArrayStructOnce sync.Once

func byteArrayStructSet() {
	byteArrayStructOnce.Do(func() {
		byteArrayStruct = gi.StructNew("GLib", "ByteArray")
	})
}

type ByteArray struct {
	native uintptr
	Data   uint8
	Len    uint32
}

var bytesStruct *gi.Struct
var bytesStructOnce sync.Once

func bytesStructSet() {
	bytesStructOnce.Do(func() {
		bytesStruct = gi.StructNew("GLib", "Bytes")
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

var bytesHashFunction *gi.Function
var bytesHashFunction_Once sync.Once

func bytesHashFunction_Set() {
	bytesHashFunction_Once.Do(func() {
		bytesHashFunction = gi.FunctionInvokerNew("GLib", "hash")
	})
}

var hashBytesInvoker *gi.Function

// Hash is a representation of the C type g_bytes_hash.
func (recv *Bytes) Hash() uint32 {
	bytesHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := bytesHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_bytes_new_from_bytes' : parameter 'offset' of type 'gsize' not supported

var bytesRefFunction *gi.Function
var bytesRefFunction_Once sync.Once

func bytesRefFunction_Set() {
	bytesRefFunction_Once.Do(func() {
		bytesRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refBytesInvoker *gi.Function

// Ref is a representation of the C type g_bytes_ref.
func (recv *Bytes) Ref() *Bytes {
	bytesRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := bytesRefFunction.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

var bytesUnrefFunction *gi.Function
var bytesUnrefFunction_Once sync.Once

func bytesUnrefFunction_Set() {
	bytesUnrefFunction_Once.Do(func() {
		bytesUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefBytesInvoker *gi.Function

// Unref is a representation of the C type g_bytes_unref.
func (recv *Bytes) Unref() {
	bytesUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	bytesUnrefFunction.Invoke(inArgs[:], nil)

}

var bytesUnrefToArrayFunction *gi.Function
var bytesUnrefToArrayFunction_Once sync.Once

func bytesUnrefToArrayFunction_Set() {
	bytesUnrefToArrayFunction_Once.Do(func() {
		bytesUnrefToArrayFunction = gi.FunctionInvokerNew("GLib", "unref_to_array")
	})
}

var unrefToArrayBytesInvoker *gi.Function

// UnrefToArray is a representation of the C type g_bytes_unref_to_array.
func (recv *Bytes) UnrefToArray() {
	bytesUnrefToArrayFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	bytesUnrefToArrayFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_bytes_unref_to_data' : parameter 'size' of type 'gsize' not supported

var checksumStruct *gi.Struct
var checksumStructOnce sync.Once

func checksumStructSet() {
	checksumStructOnce.Do(func() {
		checksumStruct = gi.StructNew("GLib", "Checksum")
	})
}

type Checksum struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_checksum_new' : parameter 'checksum_type' of type 'ChecksumType' not supported

var checksumCopyFunction *gi.Function
var checksumCopyFunction_Once sync.Once

func checksumCopyFunction_Set() {
	checksumCopyFunction_Once.Do(func() {
		checksumCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyChecksumInvoker *gi.Function

// Copy is a representation of the C type g_checksum_copy.
func (recv *Checksum) Copy() *Checksum {
	checksumCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := checksumCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Checksum{native: ret.Pointer()}

	return retGo
}

var checksumFreeFunction *gi.Function
var checksumFreeFunction_Once sync.Once

func checksumFreeFunction_Set() {
	checksumFreeFunction_Once.Do(func() {
		checksumFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeChecksumInvoker *gi.Function

// Free is a representation of the C type g_checksum_free.
func (recv *Checksum) Free() {
	checksumFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	checksumFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_checksum_get_digest' : parameter 'buffer' has no type

var checksumGetStringFunction *gi.Function
var checksumGetStringFunction_Once sync.Once

func checksumGetStringFunction_Set() {
	checksumGetStringFunction_Once.Do(func() {
		checksumGetStringFunction = gi.FunctionInvokerNew("GLib", "get_string")
	})
}

var getStringChecksumInvoker *gi.Function

// GetString is a representation of the C type g_checksum_get_string.
func (recv *Checksum) GetString() string {
	checksumGetStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := checksumGetStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var checksumResetFunction *gi.Function
var checksumResetFunction_Once sync.Once

func checksumResetFunction_Set() {
	checksumResetFunction_Once.Do(func() {
		checksumResetFunction = gi.FunctionInvokerNew("GLib", "reset")
	})
}

var resetChecksumInvoker *gi.Function

// Reset is a representation of the C type g_checksum_reset.
func (recv *Checksum) Reset() {
	checksumResetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	checksumResetFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_checksum_update' : parameter 'data' has no type

var condStruct *gi.Struct
var condStructOnce sync.Once

func condStructSet() {
	condStructOnce.Do(func() {
		condStruct = gi.StructNew("GLib", "Cond")
	})
}

type Cond struct {
	native uintptr
}

var condBroadcastFunction *gi.Function
var condBroadcastFunction_Once sync.Once

func condBroadcastFunction_Set() {
	condBroadcastFunction_Once.Do(func() {
		condBroadcastFunction = gi.FunctionInvokerNew("GLib", "broadcast")
	})
}

var broadcastCondInvoker *gi.Function

// Broadcast is a representation of the C type g_cond_broadcast.
func (recv *Cond) Broadcast() {
	condBroadcastFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	condBroadcastFunction.Invoke(inArgs[:], nil)

}

var condClearFunction *gi.Function
var condClearFunction_Once sync.Once

func condClearFunction_Set() {
	condClearFunction_Once.Do(func() {
		condClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearCondInvoker *gi.Function

// Clear is a representation of the C type g_cond_clear.
func (recv *Cond) Clear() {
	condClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	condClearFunction.Invoke(inArgs[:], nil)

}

var condInitFunction *gi.Function
var condInitFunction_Once sync.Once

func condInitFunction_Set() {
	condInitFunction_Once.Do(func() {
		condInitFunction = gi.FunctionInvokerNew("GLib", "init")
	})
}

var initCondInvoker *gi.Function

// Init is a representation of the C type g_cond_init.
func (recv *Cond) Init() {
	condInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	condInitFunction.Invoke(inArgs[:], nil)

}

var condSignalFunction *gi.Function
var condSignalFunction_Once sync.Once

func condSignalFunction_Set() {
	condSignalFunction_Once.Do(func() {
		condSignalFunction = gi.FunctionInvokerNew("GLib", "signal")
	})
}

var signalCondInvoker *gi.Function

// Signal is a representation of the C type g_cond_signal.
func (recv *Cond) Signal() {
	condSignalFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	condSignalFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_cond_wait' : parameter 'mutex' of type 'Mutex' not supported

// UNSUPPORTED : C value 'g_cond_wait_until' : parameter 'mutex' of type 'Mutex' not supported

var dataStruct *gi.Struct
var dataStructOnce sync.Once

func dataStructSet() {
	dataStructOnce.Do(func() {
		dataStruct = gi.StructNew("GLib", "Data")
	})
}

type Data struct {
	native uintptr
}

var dateStruct *gi.Struct
var dateStructOnce sync.Once

func dateStructSet() {
	dateStructOnce.Do(func() {
		dateStruct = gi.StructNew("GLib", "Date")
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

var dateNewFunction *gi.Function
var dateNewFunction_Once sync.Once

func dateNewFunction_Set() {
	dateNewFunction_Once.Do(func() {
		dateNewFunction = gi.FunctionInvokerNew("GLib", "new")
	})
}

var newDateInvoker *gi.Function

// DateNew is a representation of the C type g_date_new.
func DateNew() *Date {
	dateNewFunction_Set()

	ret := dateNewFunction.Invoke(nil, nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_new_dmy' : parameter 'day' of type 'DateDay' not supported

var dateNewJulianFunction *gi.Function
var dateNewJulianFunction_Once sync.Once

func dateNewJulianFunction_Set() {
	dateNewJulianFunction_Once.Do(func() {
		dateNewJulianFunction = gi.FunctionInvokerNew("GLib", "new_julian")
	})
}

var newJulianDateInvoker *gi.Function

// DateNewJulian is a representation of the C type g_date_new_julian.
func DateNewJulian(julianDay uint32) *Date {
	dateNewJulianFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(julianDay)

	ret := dateNewJulianFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

var dateAddDaysFunction *gi.Function
var dateAddDaysFunction_Once sync.Once

func dateAddDaysFunction_Set() {
	dateAddDaysFunction_Once.Do(func() {
		dateAddDaysFunction = gi.FunctionInvokerNew("GLib", "add_days")
	})
}

var addDaysDateInvoker *gi.Function

// AddDays is a representation of the C type g_date_add_days.
func (recv *Date) AddDays(nDays uint32) {
	dateAddDaysFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDays)

	dateAddDaysFunction.Invoke(inArgs[:], nil)

}

var dateAddMonthsFunction *gi.Function
var dateAddMonthsFunction_Once sync.Once

func dateAddMonthsFunction_Set() {
	dateAddMonthsFunction_Once.Do(func() {
		dateAddMonthsFunction = gi.FunctionInvokerNew("GLib", "add_months")
	})
}

var addMonthsDateInvoker *gi.Function

// AddMonths is a representation of the C type g_date_add_months.
func (recv *Date) AddMonths(nMonths uint32) {
	dateAddMonthsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nMonths)

	dateAddMonthsFunction.Invoke(inArgs[:], nil)

}

var dateAddYearsFunction *gi.Function
var dateAddYearsFunction_Once sync.Once

func dateAddYearsFunction_Set() {
	dateAddYearsFunction_Once.Do(func() {
		dateAddYearsFunction = gi.FunctionInvokerNew("GLib", "add_years")
	})
}

var addYearsDateInvoker *gi.Function

// AddYears is a representation of the C type g_date_add_years.
func (recv *Date) AddYears(nYears uint32) {
	dateAddYearsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nYears)

	dateAddYearsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_clamp' : parameter 'min_date' of type 'Date' not supported

var dateClearFunction *gi.Function
var dateClearFunction_Once sync.Once

func dateClearFunction_Set() {
	dateClearFunction_Once.Do(func() {
		dateClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearDateInvoker *gi.Function

// Clear is a representation of the C type g_date_clear.
func (recv *Date) Clear(nDates uint32) {
	dateClearFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDates)

	dateClearFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_compare' : parameter 'rhs' of type 'Date' not supported

var dateCopyFunction *gi.Function
var dateCopyFunction_Once sync.Once

func dateCopyFunction_Set() {
	dateCopyFunction_Once.Do(func() {
		dateCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyDateInvoker *gi.Function

// Copy is a representation of the C type g_date_copy.
func (recv *Date) Copy() *Date {
	dateCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Date{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_days_between' : parameter 'date2' of type 'Date' not supported

var dateFreeFunction *gi.Function
var dateFreeFunction_Once sync.Once

func dateFreeFunction_Set() {
	dateFreeFunction_Once.Do(func() {
		dateFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeDateInvoker *gi.Function

// Free is a representation of the C type g_date_free.
func (recv *Date) Free() {
	dateFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dateFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_get_day' : return type 'DateDay' not supported

var dateGetDayOfYearFunction *gi.Function
var dateGetDayOfYearFunction_Once sync.Once

func dateGetDayOfYearFunction_Set() {
	dateGetDayOfYearFunction_Once.Do(func() {
		dateGetDayOfYearFunction = gi.FunctionInvokerNew("GLib", "get_day_of_year")
	})
}

var getDayOfYearDateInvoker *gi.Function

// GetDayOfYear is a representation of the C type g_date_get_day_of_year.
func (recv *Date) GetDayOfYear() uint32 {
	dateGetDayOfYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetDayOfYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var dateGetIso8601WeekOfYearFunction *gi.Function
var dateGetIso8601WeekOfYearFunction_Once sync.Once

func dateGetIso8601WeekOfYearFunction_Set() {
	dateGetIso8601WeekOfYearFunction_Once.Do(func() {
		dateGetIso8601WeekOfYearFunction = gi.FunctionInvokerNew("GLib", "get_iso8601_week_of_year")
	})
}

var getIso8601WeekOfYearDateInvoker *gi.Function

// GetIso8601WeekOfYear is a representation of the C type g_date_get_iso8601_week_of_year.
func (recv *Date) GetIso8601WeekOfYear() uint32 {
	dateGetIso8601WeekOfYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetIso8601WeekOfYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var dateGetJulianFunction *gi.Function
var dateGetJulianFunction_Once sync.Once

func dateGetJulianFunction_Set() {
	dateGetJulianFunction_Once.Do(func() {
		dateGetJulianFunction = gi.FunctionInvokerNew("GLib", "get_julian")
	})
}

var getJulianDateInvoker *gi.Function

// GetJulian is a representation of the C type g_date_get_julian.
func (recv *Date) GetJulian() uint32 {
	dateGetJulianFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetJulianFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var dateGetMondayWeekOfYearFunction *gi.Function
var dateGetMondayWeekOfYearFunction_Once sync.Once

func dateGetMondayWeekOfYearFunction_Set() {
	dateGetMondayWeekOfYearFunction_Once.Do(func() {
		dateGetMondayWeekOfYearFunction = gi.FunctionInvokerNew("GLib", "get_monday_week_of_year")
	})
}

var getMondayWeekOfYearDateInvoker *gi.Function

// GetMondayWeekOfYear is a representation of the C type g_date_get_monday_week_of_year.
func (recv *Date) GetMondayWeekOfYear() uint32 {
	dateGetMondayWeekOfYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetMondayWeekOfYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_date_get_month' : return type 'DateMonth' not supported

var dateGetSundayWeekOfYearFunction *gi.Function
var dateGetSundayWeekOfYearFunction_Once sync.Once

func dateGetSundayWeekOfYearFunction_Set() {
	dateGetSundayWeekOfYearFunction_Once.Do(func() {
		dateGetSundayWeekOfYearFunction = gi.FunctionInvokerNew("GLib", "get_sunday_week_of_year")
	})
}

var getSundayWeekOfYearDateInvoker *gi.Function

// GetSundayWeekOfYear is a representation of the C type g_date_get_sunday_week_of_year.
func (recv *Date) GetSundayWeekOfYear() uint32 {
	dateGetSundayWeekOfYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateGetSundayWeekOfYearFunction.Invoke(inArgs[:], nil)

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

var dateSetJulianFunction *gi.Function
var dateSetJulianFunction_Once sync.Once

func dateSetJulianFunction_Set() {
	dateSetJulianFunction_Once.Do(func() {
		dateSetJulianFunction = gi.FunctionInvokerNew("GLib", "set_julian")
	})
}

var setJulianDateInvoker *gi.Function

// SetJulian is a representation of the C type g_date_set_julian.
func (recv *Date) SetJulian(julianDate uint32) {
	dateSetJulianFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(julianDate)

	dateSetJulianFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_set_month' : parameter 'month' of type 'DateMonth' not supported

var dateSetParseFunction *gi.Function
var dateSetParseFunction_Once sync.Once

func dateSetParseFunction_Set() {
	dateSetParseFunction_Once.Do(func() {
		dateSetParseFunction = gi.FunctionInvokerNew("GLib", "set_parse")
	})
}

var setParseDateInvoker *gi.Function

// SetParse is a representation of the C type g_date_set_parse.
func (recv *Date) SetParse(str string) {
	dateSetParseFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(str)

	dateSetParseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_set_time' : parameter 'time_' of type 'Time' not supported

var dateSetTimeTFunction *gi.Function
var dateSetTimeTFunction_Once sync.Once

func dateSetTimeTFunction_Set() {
	dateSetTimeTFunction_Once.Do(func() {
		dateSetTimeTFunction = gi.FunctionInvokerNew("GLib", "set_time_t")
	})
}

var setTimeTDateInvoker *gi.Function

// SetTimeT is a representation of the C type g_date_set_time_t.
func (recv *Date) SetTimeT(timet int64) {
	dateSetTimeTFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(timet)

	dateSetTimeTFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_set_time_val' : parameter 'timeval' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_set_year' : parameter 'year' of type 'DateYear' not supported

var dateSubtractDaysFunction *gi.Function
var dateSubtractDaysFunction_Once sync.Once

func dateSubtractDaysFunction_Set() {
	dateSubtractDaysFunction_Once.Do(func() {
		dateSubtractDaysFunction = gi.FunctionInvokerNew("GLib", "subtract_days")
	})
}

var subtractDaysDateInvoker *gi.Function

// SubtractDays is a representation of the C type g_date_subtract_days.
func (recv *Date) SubtractDays(nDays uint32) {
	dateSubtractDaysFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDays)

	dateSubtractDaysFunction.Invoke(inArgs[:], nil)

}

var dateSubtractMonthsFunction *gi.Function
var dateSubtractMonthsFunction_Once sync.Once

func dateSubtractMonthsFunction_Set() {
	dateSubtractMonthsFunction_Once.Do(func() {
		dateSubtractMonthsFunction = gi.FunctionInvokerNew("GLib", "subtract_months")
	})
}

var subtractMonthsDateInvoker *gi.Function

// SubtractMonths is a representation of the C type g_date_subtract_months.
func (recv *Date) SubtractMonths(nMonths uint32) {
	dateSubtractMonthsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nMonths)

	dateSubtractMonthsFunction.Invoke(inArgs[:], nil)

}

var dateSubtractYearsFunction *gi.Function
var dateSubtractYearsFunction_Once sync.Once

func dateSubtractYearsFunction_Set() {
	dateSubtractYearsFunction_Once.Do(func() {
		dateSubtractYearsFunction = gi.FunctionInvokerNew("GLib", "subtract_years")
	})
}

var subtractYearsDateInvoker *gi.Function

// SubtractYears is a representation of the C type g_date_subtract_years.
func (recv *Date) SubtractYears(nYears uint32) {
	dateSubtractYearsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nYears)

	dateSubtractYearsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_date_to_struct_tm' : parameter 'tm' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_date_valid' : return type 'gboolean' not supported

var dateTimeStruct *gi.Struct
var dateTimeStructOnce sync.Once

func dateTimeStructSet() {
	dateTimeStructOnce.Do(func() {
		dateTimeStruct = gi.StructNew("GLib", "DateTime")
	})
}

type DateTime struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_date_time_new' : parameter 'tz' of type 'TimeZone' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_iso8601' : parameter 'default_tz' of type 'TimeZone' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_timeval_local' : parameter 'tv' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_timeval_utc' : parameter 'tv' of type 'TimeVal' not supported

var dateTimeNewFromUnixLocalFunction *gi.Function
var dateTimeNewFromUnixLocalFunction_Once sync.Once

func dateTimeNewFromUnixLocalFunction_Set() {
	dateTimeNewFromUnixLocalFunction_Once.Do(func() {
		dateTimeNewFromUnixLocalFunction = gi.FunctionInvokerNew("GLib", "new_from_unix_local")
	})
}

var newFromUnixLocalDateTimeInvoker *gi.Function

// DateTimeNewFromUnixLocal is a representation of the C type g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	dateTimeNewFromUnixLocalFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	ret := dateTimeNewFromUnixLocalFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeNewFromUnixUtcFunction *gi.Function
var dateTimeNewFromUnixUtcFunction_Once sync.Once

func dateTimeNewFromUnixUtcFunction_Set() {
	dateTimeNewFromUnixUtcFunction_Once.Do(func() {
		dateTimeNewFromUnixUtcFunction = gi.FunctionInvokerNew("GLib", "new_from_unix_utc")
	})
}

var newFromUnixUtcDateTimeInvoker *gi.Function

// DateTimeNewFromUnixUtc is a representation of the C type g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	dateTimeNewFromUnixUtcFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	ret := dateTimeNewFromUnixUtcFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_new_local' : parameter 'seconds' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_date_time_new_now' : parameter 'tz' of type 'TimeZone' not supported

var dateTimeNewNowLocalFunction *gi.Function
var dateTimeNewNowLocalFunction_Once sync.Once

func dateTimeNewNowLocalFunction_Set() {
	dateTimeNewNowLocalFunction_Once.Do(func() {
		dateTimeNewNowLocalFunction = gi.FunctionInvokerNew("GLib", "new_now_local")
	})
}

var newNowLocalDateTimeInvoker *gi.Function

// DateTimeNewNowLocal is a representation of the C type g_date_time_new_now_local.
func DateTimeNewNowLocal() *DateTime {
	dateTimeNewNowLocalFunction_Set()

	ret := dateTimeNewNowLocalFunction.Invoke(nil, nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeNewNowUtcFunction *gi.Function
var dateTimeNewNowUtcFunction_Once sync.Once

func dateTimeNewNowUtcFunction_Set() {
	dateTimeNewNowUtcFunction_Once.Do(func() {
		dateTimeNewNowUtcFunction = gi.FunctionInvokerNew("GLib", "new_now_utc")
	})
}

var newNowUtcDateTimeInvoker *gi.Function

// DateTimeNewNowUtc is a representation of the C type g_date_time_new_now_utc.
func DateTimeNewNowUtc() *DateTime {
	dateTimeNewNowUtcFunction_Set()

	ret := dateTimeNewNowUtcFunction.Invoke(nil, nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_new_utc' : parameter 'seconds' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_date_time_add' : parameter 'timespan' of type 'TimeSpan' not supported

var dateTimeAddDaysFunction *gi.Function
var dateTimeAddDaysFunction_Once sync.Once

func dateTimeAddDaysFunction_Set() {
	dateTimeAddDaysFunction_Once.Do(func() {
		dateTimeAddDaysFunction = gi.FunctionInvokerNew("GLib", "add_days")
	})
}

var addDaysDateTimeInvoker *gi.Function

// AddDays is a representation of the C type g_date_time_add_days.
func (recv *DateTime) AddDays(days int32) *DateTime {
	dateTimeAddDaysFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(days)

	ret := dateTimeAddDaysFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_add_full' : parameter 'seconds' of type 'gdouble' not supported

var dateTimeAddHoursFunction *gi.Function
var dateTimeAddHoursFunction_Once sync.Once

func dateTimeAddHoursFunction_Set() {
	dateTimeAddHoursFunction_Once.Do(func() {
		dateTimeAddHoursFunction = gi.FunctionInvokerNew("GLib", "add_hours")
	})
}

var addHoursDateTimeInvoker *gi.Function

// AddHours is a representation of the C type g_date_time_add_hours.
func (recv *DateTime) AddHours(hours int32) *DateTime {
	dateTimeAddHoursFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(hours)

	ret := dateTimeAddHoursFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeAddMinutesFunction *gi.Function
var dateTimeAddMinutesFunction_Once sync.Once

func dateTimeAddMinutesFunction_Set() {
	dateTimeAddMinutesFunction_Once.Do(func() {
		dateTimeAddMinutesFunction = gi.FunctionInvokerNew("GLib", "add_minutes")
	})
}

var addMinutesDateTimeInvoker *gi.Function

// AddMinutes is a representation of the C type g_date_time_add_minutes.
func (recv *DateTime) AddMinutes(minutes int32) *DateTime {
	dateTimeAddMinutesFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(minutes)

	ret := dateTimeAddMinutesFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeAddMonthsFunction *gi.Function
var dateTimeAddMonthsFunction_Once sync.Once

func dateTimeAddMonthsFunction_Set() {
	dateTimeAddMonthsFunction_Once.Do(func() {
		dateTimeAddMonthsFunction = gi.FunctionInvokerNew("GLib", "add_months")
	})
}

var addMonthsDateTimeInvoker *gi.Function

// AddMonths is a representation of the C type g_date_time_add_months.
func (recv *DateTime) AddMonths(months int32) *DateTime {
	dateTimeAddMonthsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(months)

	ret := dateTimeAddMonthsFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_add_seconds' : parameter 'seconds' of type 'gdouble' not supported

var dateTimeAddWeeksFunction *gi.Function
var dateTimeAddWeeksFunction_Once sync.Once

func dateTimeAddWeeksFunction_Set() {
	dateTimeAddWeeksFunction_Once.Do(func() {
		dateTimeAddWeeksFunction = gi.FunctionInvokerNew("GLib", "add_weeks")
	})
}

var addWeeksDateTimeInvoker *gi.Function

// AddWeeks is a representation of the C type g_date_time_add_weeks.
func (recv *DateTime) AddWeeks(weeks int32) *DateTime {
	dateTimeAddWeeksFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(weeks)

	ret := dateTimeAddWeeksFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeAddYearsFunction *gi.Function
var dateTimeAddYearsFunction_Once sync.Once

func dateTimeAddYearsFunction_Set() {
	dateTimeAddYearsFunction_Once.Do(func() {
		dateTimeAddYearsFunction = gi.FunctionInvokerNew("GLib", "add_years")
	})
}

var addYearsDateTimeInvoker *gi.Function

// AddYears is a representation of the C type g_date_time_add_years.
func (recv *DateTime) AddYears(years int32) *DateTime {
	dateTimeAddYearsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(years)

	ret := dateTimeAddYearsFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_difference' : parameter 'begin' of type 'DateTime' not supported

var dateTimeFormatFunction *gi.Function
var dateTimeFormatFunction_Once sync.Once

func dateTimeFormatFunction_Set() {
	dateTimeFormatFunction_Once.Do(func() {
		dateTimeFormatFunction = gi.FunctionInvokerNew("GLib", "format")
	})
}

var formatDateTimeInvoker *gi.Function

// Format is a representation of the C type g_date_time_format.
func (recv *DateTime) Format(format string) string {
	dateTimeFormatFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(format)

	ret := dateTimeFormatFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var dateTimeFormatIso8601Function *gi.Function
var dateTimeFormatIso8601Function_Once sync.Once

func dateTimeFormatIso8601Function_Set() {
	dateTimeFormatIso8601Function_Once.Do(func() {
		dateTimeFormatIso8601Function = gi.FunctionInvokerNew("GLib", "format_iso8601")
	})
}

var formatIso8601DateTimeInvoker *gi.Function

// FormatIso8601 is a representation of the C type g_date_time_format_iso8601.
func (recv *DateTime) FormatIso8601() string {
	dateTimeFormatIso8601Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeFormatIso8601Function.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var dateTimeGetDayOfMonthFunction *gi.Function
var dateTimeGetDayOfMonthFunction_Once sync.Once

func dateTimeGetDayOfMonthFunction_Set() {
	dateTimeGetDayOfMonthFunction_Once.Do(func() {
		dateTimeGetDayOfMonthFunction = gi.FunctionInvokerNew("GLib", "get_day_of_month")
	})
}

var getDayOfMonthDateTimeInvoker *gi.Function

// GetDayOfMonth is a representation of the C type g_date_time_get_day_of_month.
func (recv *DateTime) GetDayOfMonth() int32 {
	dateTimeGetDayOfMonthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetDayOfMonthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetDayOfWeekFunction *gi.Function
var dateTimeGetDayOfWeekFunction_Once sync.Once

func dateTimeGetDayOfWeekFunction_Set() {
	dateTimeGetDayOfWeekFunction_Once.Do(func() {
		dateTimeGetDayOfWeekFunction = gi.FunctionInvokerNew("GLib", "get_day_of_week")
	})
}

var getDayOfWeekDateTimeInvoker *gi.Function

// GetDayOfWeek is a representation of the C type g_date_time_get_day_of_week.
func (recv *DateTime) GetDayOfWeek() int32 {
	dateTimeGetDayOfWeekFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetDayOfWeekFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetDayOfYearFunction *gi.Function
var dateTimeGetDayOfYearFunction_Once sync.Once

func dateTimeGetDayOfYearFunction_Set() {
	dateTimeGetDayOfYearFunction_Once.Do(func() {
		dateTimeGetDayOfYearFunction = gi.FunctionInvokerNew("GLib", "get_day_of_year")
	})
}

var getDayOfYearDateTimeInvoker *gi.Function

// GetDayOfYear is a representation of the C type g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() int32 {
	dateTimeGetDayOfYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetDayOfYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetHourFunction *gi.Function
var dateTimeGetHourFunction_Once sync.Once

func dateTimeGetHourFunction_Set() {
	dateTimeGetHourFunction_Once.Do(func() {
		dateTimeGetHourFunction = gi.FunctionInvokerNew("GLib", "get_hour")
	})
}

var getHourDateTimeInvoker *gi.Function

// GetHour is a representation of the C type g_date_time_get_hour.
func (recv *DateTime) GetHour() int32 {
	dateTimeGetHourFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetHourFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetMicrosecondFunction *gi.Function
var dateTimeGetMicrosecondFunction_Once sync.Once

func dateTimeGetMicrosecondFunction_Set() {
	dateTimeGetMicrosecondFunction_Once.Do(func() {
		dateTimeGetMicrosecondFunction = gi.FunctionInvokerNew("GLib", "get_microsecond")
	})
}

var getMicrosecondDateTimeInvoker *gi.Function

// GetMicrosecond is a representation of the C type g_date_time_get_microsecond.
func (recv *DateTime) GetMicrosecond() int32 {
	dateTimeGetMicrosecondFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetMicrosecondFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetMinuteFunction *gi.Function
var dateTimeGetMinuteFunction_Once sync.Once

func dateTimeGetMinuteFunction_Set() {
	dateTimeGetMinuteFunction_Once.Do(func() {
		dateTimeGetMinuteFunction = gi.FunctionInvokerNew("GLib", "get_minute")
	})
}

var getMinuteDateTimeInvoker *gi.Function

// GetMinute is a representation of the C type g_date_time_get_minute.
func (recv *DateTime) GetMinute() int32 {
	dateTimeGetMinuteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetMinuteFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetMonthFunction *gi.Function
var dateTimeGetMonthFunction_Once sync.Once

func dateTimeGetMonthFunction_Set() {
	dateTimeGetMonthFunction_Once.Do(func() {
		dateTimeGetMonthFunction = gi.FunctionInvokerNew("GLib", "get_month")
	})
}

var getMonthDateTimeInvoker *gi.Function

// GetMonth is a representation of the C type g_date_time_get_month.
func (recv *DateTime) GetMonth() int32 {
	dateTimeGetMonthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetMonthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetSecondFunction *gi.Function
var dateTimeGetSecondFunction_Once sync.Once

func dateTimeGetSecondFunction_Set() {
	dateTimeGetSecondFunction_Once.Do(func() {
		dateTimeGetSecondFunction = gi.FunctionInvokerNew("GLib", "get_second")
	})
}

var getSecondDateTimeInvoker *gi.Function

// GetSecond is a representation of the C type g_date_time_get_second.
func (recv *DateTime) GetSecond() int32 {
	dateTimeGetSecondFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetSecondFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_get_seconds' : return type 'gdouble' not supported

var dateTimeGetTimezoneFunction *gi.Function
var dateTimeGetTimezoneFunction_Once sync.Once

func dateTimeGetTimezoneFunction_Set() {
	dateTimeGetTimezoneFunction_Once.Do(func() {
		dateTimeGetTimezoneFunction = gi.FunctionInvokerNew("GLib", "get_timezone")
	})
}

var getTimezoneDateTimeInvoker *gi.Function

// GetTimezone is a representation of the C type g_date_time_get_timezone.
func (recv *DateTime) GetTimezone() *TimeZone {
	dateTimeGetTimezoneFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetTimezoneFunction.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var dateTimeGetTimezoneAbbreviationFunction *gi.Function
var dateTimeGetTimezoneAbbreviationFunction_Once sync.Once

func dateTimeGetTimezoneAbbreviationFunction_Set() {
	dateTimeGetTimezoneAbbreviationFunction_Once.Do(func() {
		dateTimeGetTimezoneAbbreviationFunction = gi.FunctionInvokerNew("GLib", "get_timezone_abbreviation")
	})
}

var getTimezoneAbbreviationDateTimeInvoker *gi.Function

// GetTimezoneAbbreviation is a representation of the C type g_date_time_get_timezone_abbreviation.
func (recv *DateTime) GetTimezoneAbbreviation() string {
	dateTimeGetTimezoneAbbreviationFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetTimezoneAbbreviationFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_get_utc_offset' : return type 'TimeSpan' not supported

var dateTimeGetWeekNumberingYearFunction *gi.Function
var dateTimeGetWeekNumberingYearFunction_Once sync.Once

func dateTimeGetWeekNumberingYearFunction_Set() {
	dateTimeGetWeekNumberingYearFunction_Once.Do(func() {
		dateTimeGetWeekNumberingYearFunction = gi.FunctionInvokerNew("GLib", "get_week_numbering_year")
	})
}

var getWeekNumberingYearDateTimeInvoker *gi.Function

// GetWeekNumberingYear is a representation of the C type g_date_time_get_week_numbering_year.
func (recv *DateTime) GetWeekNumberingYear() int32 {
	dateTimeGetWeekNumberingYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetWeekNumberingYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetWeekOfYearFunction *gi.Function
var dateTimeGetWeekOfYearFunction_Once sync.Once

func dateTimeGetWeekOfYearFunction_Set() {
	dateTimeGetWeekOfYearFunction_Once.Do(func() {
		dateTimeGetWeekOfYearFunction = gi.FunctionInvokerNew("GLib", "get_week_of_year")
	})
}

var getWeekOfYearDateTimeInvoker *gi.Function

// GetWeekOfYear is a representation of the C type g_date_time_get_week_of_year.
func (recv *DateTime) GetWeekOfYear() int32 {
	dateTimeGetWeekOfYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetWeekOfYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetYearFunction *gi.Function
var dateTimeGetYearFunction_Once sync.Once

func dateTimeGetYearFunction_Set() {
	dateTimeGetYearFunction_Once.Do(func() {
		dateTimeGetYearFunction = gi.FunctionInvokerNew("GLib", "get_year")
	})
}

var getYearDateTimeInvoker *gi.Function

// GetYear is a representation of the C type g_date_time_get_year.
func (recv *DateTime) GetYear() int32 {
	dateTimeGetYearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeGetYearFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var dateTimeGetYmdFunction *gi.Function
var dateTimeGetYmdFunction_Once sync.Once

func dateTimeGetYmdFunction_Set() {
	dateTimeGetYmdFunction_Once.Do(func() {
		dateTimeGetYmdFunction = gi.FunctionInvokerNew("GLib", "get_ymd")
	})
}

var getYmdDateTimeInvoker *gi.Function

// GetYmd is a representation of the C type g_date_time_get_ymd.
func (recv *DateTime) GetYmd() (int32, int32, int32) {
	dateTimeGetYmdFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	dateTimeGetYmdFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2
}

// UNSUPPORTED : C value 'g_date_time_is_daylight_savings' : return type 'gboolean' not supported

var dateTimeRefFunction *gi.Function
var dateTimeRefFunction_Once sync.Once

func dateTimeRefFunction_Set() {
	dateTimeRefFunction_Once.Do(func() {
		dateTimeRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refDateTimeInvoker *gi.Function

// Ref is a representation of the C type g_date_time_ref.
func (recv *DateTime) Ref() *DateTime {
	dateTimeRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeRefFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeToLocalFunction *gi.Function
var dateTimeToLocalFunction_Once sync.Once

func dateTimeToLocalFunction_Set() {
	dateTimeToLocalFunction_Once.Do(func() {
		dateTimeToLocalFunction = gi.FunctionInvokerNew("GLib", "to_local")
	})
}

var toLocalDateTimeInvoker *gi.Function

// ToLocal is a representation of the C type g_date_time_to_local.
func (recv *DateTime) ToLocal() *DateTime {
	dateTimeToLocalFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeToLocalFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_date_time_to_timeval' : parameter 'tv' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_time_to_timezone' : parameter 'tz' of type 'TimeZone' not supported

var dateTimeToUnixFunction *gi.Function
var dateTimeToUnixFunction_Once sync.Once

func dateTimeToUnixFunction_Set() {
	dateTimeToUnixFunction_Once.Do(func() {
		dateTimeToUnixFunction = gi.FunctionInvokerNew("GLib", "to_unix")
	})
}

var toUnixDateTimeInvoker *gi.Function

// ToUnix is a representation of the C type g_date_time_to_unix.
func (recv *DateTime) ToUnix() int64 {
	dateTimeToUnixFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeToUnixFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var dateTimeToUtcFunction *gi.Function
var dateTimeToUtcFunction_Once sync.Once

func dateTimeToUtcFunction_Set() {
	dateTimeToUtcFunction_Once.Do(func() {
		dateTimeToUtcFunction = gi.FunctionInvokerNew("GLib", "to_utc")
	})
}

var toUtcDateTimeInvoker *gi.Function

// ToUtc is a representation of the C type g_date_time_to_utc.
func (recv *DateTime) ToUtc() *DateTime {
	dateTimeToUtcFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := dateTimeToUtcFunction.Invoke(inArgs[:], nil)

	retGo := &DateTime{native: ret.Pointer()}

	return retGo
}

var dateTimeUnrefFunction *gi.Function
var dateTimeUnrefFunction_Once sync.Once

func dateTimeUnrefFunction_Set() {
	dateTimeUnrefFunction_Once.Do(func() {
		dateTimeUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefDateTimeInvoker *gi.Function

// Unref is a representation of the C type g_date_time_unref.
func (recv *DateTime) Unref() {
	dateTimeUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dateTimeUnrefFunction.Invoke(inArgs[:], nil)

}

var debugKeyStruct *gi.Struct
var debugKeyStructOnce sync.Once

func debugKeyStructSet() {
	debugKeyStructOnce.Do(func() {
		debugKeyStruct = gi.StructNew("GLib", "DebugKey")
	})
}

type DebugKey struct {
	native uintptr
	Key    string
	Value  uint32
}

var dirStruct *gi.Struct
var dirStructOnce sync.Once

func dirStructSet() {
	dirStructOnce.Do(func() {
		dirStruct = gi.StructNew("GLib", "Dir")
	})
}

type Dir struct {
	native uintptr
}

var dirCloseFunction *gi.Function
var dirCloseFunction_Once sync.Once

func dirCloseFunction_Set() {
	dirCloseFunction_Once.Do(func() {
		dirCloseFunction = gi.FunctionInvokerNew("GLib", "close")
	})
}

var closeDirInvoker *gi.Function

// Close is a representation of the C type g_dir_close.
func (recv *Dir) Close() {
	dirCloseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dirCloseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_dir_read_name' : return type 'filename' not supported

var dirRewindFunction *gi.Function
var dirRewindFunction_Once sync.Once

func dirRewindFunction_Set() {
	dirRewindFunction_Once.Do(func() {
		dirRewindFunction = gi.FunctionInvokerNew("GLib", "rewind")
	})
}

var rewindDirInvoker *gi.Function

// Rewind is a representation of the C type g_dir_rewind.
func (recv *Dir) Rewind() {
	dirRewindFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	dirRewindFunction.Invoke(inArgs[:], nil)

}

var errorStruct *gi.Struct
var errorStructOnce sync.Once

func errorStructSet() {
	errorStructOnce.Do(func() {
		errorStruct = gi.StructNew("GLib", "Error")
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

var errorCopyFunction *gi.Function
var errorCopyFunction_Once sync.Once

func errorCopyFunction_Set() {
	errorCopyFunction_Once.Do(func() {
		errorCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyErrorInvoker *gi.Function

// Copy is a representation of the C type g_error_copy.
func (recv *Error) Copy() *Error {
	errorCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := errorCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Error{native: ret.Pointer()}

	return retGo
}

var errorFreeFunction *gi.Function
var errorFreeFunction_Once sync.Once

func errorFreeFunction_Set() {
	errorFreeFunction_Once.Do(func() {
		errorFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeErrorInvoker *gi.Function

// Free is a representation of the C type g_error_free.
func (recv *Error) Free() {
	errorFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	errorFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_error_matches' : parameter 'domain' of type 'Quark' not supported

var hashTableStruct *gi.Struct
var hashTableStructOnce sync.Once

func hashTableStructSet() {
	hashTableStructOnce.Do(func() {
		hashTableStruct = gi.StructNew("GLib", "HashTable")
	})
}

type HashTable struct {
	native uintptr
}

var hashTableIterStruct *gi.Struct
var hashTableIterStructOnce sync.Once

func hashTableIterStructSet() {
	hashTableIterStructOnce.Do(func() {
		hashTableIterStruct = gi.StructNew("GLib", "HashTableIter")
	})
}

type HashTableIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_hash_table_iter_get_hash_table' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_iter_init' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_iter_next' : parameter 'key' of type 'gpointer' not supported

var hashTableIterRemoveFunction *gi.Function
var hashTableIterRemoveFunction_Once sync.Once

func hashTableIterRemoveFunction_Set() {
	hashTableIterRemoveFunction_Once.Do(func() {
		hashTableIterRemoveFunction = gi.FunctionInvokerNew("GLib", "remove")
	})
}

var removeHashTableIterInvoker *gi.Function

// Remove is a representation of the C type g_hash_table_iter_remove.
func (recv *HashTableIter) Remove() {
	hashTableIterRemoveFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	hashTableIterRemoveFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_hash_table_iter_replace' : parameter 'value' of type 'gpointer' not supported

var hashTableIterStealFunction *gi.Function
var hashTableIterStealFunction_Once sync.Once

func hashTableIterStealFunction_Set() {
	hashTableIterStealFunction_Once.Do(func() {
		hashTableIterStealFunction = gi.FunctionInvokerNew("GLib", "steal")
	})
}

var stealHashTableIterInvoker *gi.Function

// Steal is a representation of the C type g_hash_table_iter_steal.
func (recv *HashTableIter) Steal() {
	hashTableIterStealFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	hashTableIterStealFunction.Invoke(inArgs[:], nil)

}

var hmacStruct *gi.Struct
var hmacStructOnce sync.Once

func hmacStructSet() {
	hmacStructOnce.Do(func() {
		hmacStruct = gi.StructNew("GLib", "Hmac")
	})
}

type Hmac struct {
	native uintptr
}

var hmacCopyFunction *gi.Function
var hmacCopyFunction_Once sync.Once

func hmacCopyFunction_Set() {
	hmacCopyFunction_Once.Do(func() {
		hmacCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyHmacInvoker *gi.Function

// Copy is a representation of the C type g_hmac_copy.
func (recv *Hmac) Copy() *Hmac {
	hmacCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hmacCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Hmac{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_hmac_get_digest' : parameter 'buffer' has no type

var hmacGetStringFunction *gi.Function
var hmacGetStringFunction_Once sync.Once

func hmacGetStringFunction_Set() {
	hmacGetStringFunction_Once.Do(func() {
		hmacGetStringFunction = gi.FunctionInvokerNew("GLib", "get_string")
	})
}

var getStringHmacInvoker *gi.Function

// GetString is a representation of the C type g_hmac_get_string.
func (recv *Hmac) GetString() string {
	hmacGetStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hmacGetStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var hmacRefFunction *gi.Function
var hmacRefFunction_Once sync.Once

func hmacRefFunction_Set() {
	hmacRefFunction_Once.Do(func() {
		hmacRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refHmacInvoker *gi.Function

// Ref is a representation of the C type g_hmac_ref.
func (recv *Hmac) Ref() *Hmac {
	hmacRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := hmacRefFunction.Invoke(inArgs[:], nil)

	retGo := &Hmac{native: ret.Pointer()}

	return retGo
}

var hmacUnrefFunction *gi.Function
var hmacUnrefFunction_Once sync.Once

func hmacUnrefFunction_Set() {
	hmacUnrefFunction_Once.Do(func() {
		hmacUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefHmacInvoker *gi.Function

// Unref is a representation of the C type g_hmac_unref.
func (recv *Hmac) Unref() {
	hmacUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	hmacUnrefFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_hmac_update' : parameter 'data' has no type

var hookStruct *gi.Struct
var hookStructOnce sync.Once

func hookStructSet() {
	hookStructOnce.Do(func() {
		hookStruct = gi.StructNew("GLib", "Hook")
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

var hookListStruct *gi.Struct
var hookListStructOnce sync.Once

func hookListStructSet() {
	hookListStructOnce.Do(func() {
		hookListStruct = gi.StructNew("GLib", "HookList")
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

var hookListClearFunction *gi.Function
var hookListClearFunction_Once sync.Once

func hookListClearFunction_Set() {
	hookListClearFunction_Once.Do(func() {
		hookListClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearHookListInvoker *gi.Function

// Clear is a representation of the C type g_hook_list_clear.
func (recv *HookList) Clear() {
	hookListClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	hookListClearFunction.Invoke(inArgs[:], nil)

}

var hookListInitFunction *gi.Function
var hookListInitFunction_Once sync.Once

func hookListInitFunction_Set() {
	hookListInitFunction_Once.Do(func() {
		hookListInitFunction = gi.FunctionInvokerNew("GLib", "init")
	})
}

var initHookListInvoker *gi.Function

// Init is a representation of the C type g_hook_list_init.
func (recv *HookList) Init(hookSize uint32) {
	hookListInitFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(hookSize)

	hookListInitFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_hook_list_invoke' : parameter 'may_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hook_list_invoke_check' : parameter 'may_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hook_list_marshal' : parameter 'may_recurse' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_hook_list_marshal_check' : parameter 'may_recurse' of type 'gboolean' not supported

var iConvStruct *gi.Struct
var iConvStructOnce sync.Once

func iConvStructSet() {
	iConvStructOnce.Do(func() {
		iConvStruct = gi.StructNew("GLib", "IConv")
	})
}

type IConv struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'inbytes_left' of type 'gsize' not supported

var iConvCloseFunction *gi.Function
var iConvCloseFunction_Once sync.Once

func iConvCloseFunction_Set() {
	iConvCloseFunction_Once.Do(func() {
		iConvCloseFunction = gi.FunctionInvokerNew("GLib", "close")
	})
}

var closeIConvInvoker *gi.Function

// Close is a representation of the C type g_iconv_close.
func (recv *IConv) Close() int32 {
	iConvCloseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iConvCloseFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var iOChannelStruct *gi.Struct
var iOChannelStructOnce sync.Once

func iOChannelStructSet() {
	iOChannelStructOnce.Do(func() {
		iOChannelStruct = gi.StructNew("GLib", "IOChannel")
	})
}

type IOChannel struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_channel_new_file' : parameter 'filename' of type 'filename' not supported

var iOChannelUnixNewFunction *gi.Function
var iOChannelUnixNewFunction_Once sync.Once

func iOChannelUnixNewFunction_Set() {
	iOChannelUnixNewFunction_Once.Do(func() {
		iOChannelUnixNewFunction = gi.FunctionInvokerNew("GLib", "unix_new")
	})
}

var unixNewIOChannelInvoker *gi.Function

// IOChannelUnixNew is a representation of the C type g_io_channel_unix_new.
func IOChannelUnixNew(fd int32) *IOChannel {
	iOChannelUnixNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	ret := iOChannelUnixNewFunction.Invoke(inArgs[:], nil)

	retGo := &IOChannel{native: ret.Pointer()}

	return retGo
}

var iOChannelCloseFunction *gi.Function
var iOChannelCloseFunction_Once sync.Once

func iOChannelCloseFunction_Set() {
	iOChannelCloseFunction_Once.Do(func() {
		iOChannelCloseFunction = gi.FunctionInvokerNew("GLib", "close")
	})
}

var closeIOChannelInvoker *gi.Function

// Close is a representation of the C type g_io_channel_close.
func (recv *IOChannel) Close() {
	iOChannelCloseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iOChannelCloseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_flush' : return type 'IOStatus' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffer_condition' : return type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffer_size' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffered' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_io_channel_get_close_on_unref' : return type 'gboolean' not supported

var iOChannelGetEncodingFunction *gi.Function
var iOChannelGetEncodingFunction_Once sync.Once

func iOChannelGetEncodingFunction_Set() {
	iOChannelGetEncodingFunction_Once.Do(func() {
		iOChannelGetEncodingFunction = gi.FunctionInvokerNew("GLib", "get_encoding")
	})
}

var getEncodingIOChannelInvoker *gi.Function

// GetEncoding is a representation of the C type g_io_channel_get_encoding.
func (recv *IOChannel) GetEncoding() string {
	iOChannelGetEncodingFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iOChannelGetEncodingFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_io_channel_get_flags' : return type 'IOFlags' not supported

var iOChannelGetLineTermFunction *gi.Function
var iOChannelGetLineTermFunction_Once sync.Once

func iOChannelGetLineTermFunction_Set() {
	iOChannelGetLineTermFunction_Once.Do(func() {
		iOChannelGetLineTermFunction = gi.FunctionInvokerNew("GLib", "get_line_term")
	})
}

var getLineTermIOChannelInvoker *gi.Function

// GetLineTerm is a representation of the C type g_io_channel_get_line_term.
func (recv *IOChannel) GetLineTerm(length int32) string {
	iOChannelGetLineTermFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(length)

	ret := iOChannelGetLineTermFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var iOChannelInitFunction *gi.Function
var iOChannelInitFunction_Once sync.Once

func iOChannelInitFunction_Set() {
	iOChannelInitFunction_Once.Do(func() {
		iOChannelInitFunction = gi.FunctionInvokerNew("GLib", "init")
	})
}

var initIOChannelInvoker *gi.Function

// Init is a representation of the C type g_io_channel_init.
func (recv *IOChannel) Init() {
	iOChannelInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iOChannelInitFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_read' : parameter 'count' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_read_chars' : parameter 'buf' has no type

// UNSUPPORTED : C value 'g_io_channel_read_line' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_read_line_string' : parameter 'buffer' of type 'String' not supported

// UNSUPPORTED : C value 'g_io_channel_read_to_end' : parameter 'str_return' has no type

// UNSUPPORTED : C value 'g_io_channel_read_unichar' : parameter 'thechar' of type 'gunichar' not supported

var iOChannelRefFunction *gi.Function
var iOChannelRefFunction_Once sync.Once

func iOChannelRefFunction_Set() {
	iOChannelRefFunction_Once.Do(func() {
		iOChannelRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refIOChannelInvoker *gi.Function

// Ref is a representation of the C type g_io_channel_ref.
func (recv *IOChannel) Ref() *IOChannel {
	iOChannelRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iOChannelRefFunction.Invoke(inArgs[:], nil)

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

var iOChannelSetLineTermFunction *gi.Function
var iOChannelSetLineTermFunction_Once sync.Once

func iOChannelSetLineTermFunction_Set() {
	iOChannelSetLineTermFunction_Once.Do(func() {
		iOChannelSetLineTermFunction = gi.FunctionInvokerNew("GLib", "set_line_term")
	})
}

var setLineTermIOChannelInvoker *gi.Function

// SetLineTerm is a representation of the C type g_io_channel_set_line_term.
func (recv *IOChannel) SetLineTerm(lineTerm string, length int32) {
	iOChannelSetLineTermFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(lineTerm)
	inArgs[2].SetInt32(length)

	iOChannelSetLineTermFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_shutdown' : parameter 'flush' of type 'gboolean' not supported

var iOChannelUnixGetFdFunction *gi.Function
var iOChannelUnixGetFdFunction_Once sync.Once

func iOChannelUnixGetFdFunction_Set() {
	iOChannelUnixGetFdFunction_Once.Do(func() {
		iOChannelUnixGetFdFunction = gi.FunctionInvokerNew("GLib", "unix_get_fd")
	})
}

var unixGetFdIOChannelInvoker *gi.Function

// UnixGetFd is a representation of the C type g_io_channel_unix_get_fd.
func (recv *IOChannel) UnixGetFd() int32 {
	iOChannelUnixGetFdFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iOChannelUnixGetFdFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var iOChannelUnrefFunction *gi.Function
var iOChannelUnrefFunction_Once sync.Once

func iOChannelUnrefFunction_Set() {
	iOChannelUnrefFunction_Once.Do(func() {
		iOChannelUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefIOChannelInvoker *gi.Function

// Unref is a representation of the C type g_io_channel_unref.
func (recv *IOChannel) Unref() {
	iOChannelUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iOChannelUnrefFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_io_channel_write' : parameter 'count' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_write_chars' : parameter 'buf' has no type

// UNSUPPORTED : C value 'g_io_channel_write_unichar' : parameter 'thechar' of type 'gunichar' not supported

var iOFuncsStruct *gi.Struct
var iOFuncsStructOnce sync.Once

func iOFuncsStructSet() {
	iOFuncsStructOnce.Do(func() {
		iOFuncsStruct = gi.StructNew("GLib", "IOFuncs")
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

var keyFileStruct *gi.Struct
var keyFileStructOnce sync.Once

func keyFileStructSet() {
	keyFileStructOnce.Do(func() {
		keyFileStruct = gi.StructNew("GLib", "KeyFile")
	})
}

type KeyFile struct {
	native uintptr
}

var keyFileNewFunction *gi.Function
var keyFileNewFunction_Once sync.Once

func keyFileNewFunction_Set() {
	keyFileNewFunction_Once.Do(func() {
		keyFileNewFunction = gi.FunctionInvokerNew("GLib", "new")
	})
}

var newKeyFileInvoker *gi.Function

// KeyFileNew is a representation of the C type g_key_file_new.
func KeyFileNew() *KeyFile {
	keyFileNewFunction_Set()

	ret := keyFileNewFunction.Invoke(nil, nil)

	retGo := &KeyFile{native: ret.Pointer()}

	return retGo
}

var keyFileFreeFunction *gi.Function
var keyFileFreeFunction_Once sync.Once

func keyFileFreeFunction_Set() {
	keyFileFreeFunction_Once.Do(func() {
		keyFileFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeKeyFileInvoker *gi.Function

// Free is a representation of the C type g_key_file_free.
func (recv *KeyFile) Free() {
	keyFileFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	keyFileFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_get_boolean' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_key_file_get_boolean_list' : parameter 'length' of type 'gsize' not supported

var keyFileGetCommentFunction *gi.Function
var keyFileGetCommentFunction_Once sync.Once

func keyFileGetCommentFunction_Set() {
	keyFileGetCommentFunction_Once.Do(func() {
		keyFileGetCommentFunction = gi.FunctionInvokerNew("GLib", "get_comment")
	})
}

var getCommentKeyFileInvoker *gi.Function

// GetComment is a representation of the C type g_key_file_get_comment.
func (recv *KeyFile) GetComment(groupName string, key string) string {
	keyFileGetCommentFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := keyFileGetCommentFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_key_file_get_double_list' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_get_groups' : parameter 'length' of type 'gsize' not supported

var keyFileGetInt64Function *gi.Function
var keyFileGetInt64Function_Once sync.Once

func keyFileGetInt64Function_Set() {
	keyFileGetInt64Function_Once.Do(func() {
		keyFileGetInt64Function = gi.FunctionInvokerNew("GLib", "get_int64")
	})
}

var getInt64KeyFileInvoker *gi.Function

// GetInt64 is a representation of the C type g_key_file_get_int64.
func (recv *KeyFile) GetInt64(groupName string, key string) int64 {
	keyFileGetInt64Function_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := keyFileGetInt64Function.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var keyFileGetIntegerFunction *gi.Function
var keyFileGetIntegerFunction_Once sync.Once

func keyFileGetIntegerFunction_Set() {
	keyFileGetIntegerFunction_Once.Do(func() {
		keyFileGetIntegerFunction = gi.FunctionInvokerNew("GLib", "get_integer")
	})
}

var getIntegerKeyFileInvoker *gi.Function

// GetInteger is a representation of the C type g_key_file_get_integer.
func (recv *KeyFile) GetInteger(groupName string, key string) int32 {
	keyFileGetIntegerFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := keyFileGetIntegerFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_integer_list' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_get_keys' : parameter 'length' of type 'gsize' not supported

var keyFileGetLocaleForKeyFunction *gi.Function
var keyFileGetLocaleForKeyFunction_Once sync.Once

func keyFileGetLocaleForKeyFunction_Set() {
	keyFileGetLocaleForKeyFunction_Once.Do(func() {
		keyFileGetLocaleForKeyFunction = gi.FunctionInvokerNew("GLib", "get_locale_for_key")
	})
}

var getLocaleForKeyKeyFileInvoker *gi.Function

// GetLocaleForKey is a representation of the C type g_key_file_get_locale_for_key.
func (recv *KeyFile) GetLocaleForKey(groupName string, key string, locale string) string {
	keyFileGetLocaleForKeyFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)

	ret := keyFileGetLocaleForKeyFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var keyFileGetLocaleStringFunction *gi.Function
var keyFileGetLocaleStringFunction_Once sync.Once

func keyFileGetLocaleStringFunction_Set() {
	keyFileGetLocaleStringFunction_Once.Do(func() {
		keyFileGetLocaleStringFunction = gi.FunctionInvokerNew("GLib", "get_locale_string")
	})
}

var getLocaleStringKeyFileInvoker *gi.Function

// GetLocaleString is a representation of the C type g_key_file_get_locale_string.
func (recv *KeyFile) GetLocaleString(groupName string, key string, locale string) string {
	keyFileGetLocaleStringFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)

	ret := keyFileGetLocaleStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_locale_string_list' : parameter 'length' of type 'gsize' not supported

var keyFileGetStartGroupFunction *gi.Function
var keyFileGetStartGroupFunction_Once sync.Once

func keyFileGetStartGroupFunction_Set() {
	keyFileGetStartGroupFunction_Once.Do(func() {
		keyFileGetStartGroupFunction = gi.FunctionInvokerNew("GLib", "get_start_group")
	})
}

var getStartGroupKeyFileInvoker *gi.Function

// GetStartGroup is a representation of the C type g_key_file_get_start_group.
func (recv *KeyFile) GetStartGroup() string {
	keyFileGetStartGroupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := keyFileGetStartGroupFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var keyFileGetStringFunction *gi.Function
var keyFileGetStringFunction_Once sync.Once

func keyFileGetStringFunction_Set() {
	keyFileGetStringFunction_Once.Do(func() {
		keyFileGetStringFunction = gi.FunctionInvokerNew("GLib", "get_string")
	})
}

var getStringKeyFileInvoker *gi.Function

// GetString is a representation of the C type g_key_file_get_string.
func (recv *KeyFile) GetString(groupName string, key string) string {
	keyFileGetStringFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := keyFileGetStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_key_file_get_string_list' : parameter 'length' of type 'gsize' not supported

var keyFileGetUint64Function *gi.Function
var keyFileGetUint64Function_Once sync.Once

func keyFileGetUint64Function_Set() {
	keyFileGetUint64Function_Once.Do(func() {
		keyFileGetUint64Function = gi.FunctionInvokerNew("GLib", "get_uint64")
	})
}

var getUint64KeyFileInvoker *gi.Function

// GetUint64 is a representation of the C type g_key_file_get_uint64.
func (recv *KeyFile) GetUint64(groupName string, key string) uint64 {
	keyFileGetUint64Function_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := keyFileGetUint64Function.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

var keyFileGetValueFunction *gi.Function
var keyFileGetValueFunction_Once sync.Once

func keyFileGetValueFunction_Set() {
	keyFileGetValueFunction_Once.Do(func() {
		keyFileGetValueFunction = gi.FunctionInvokerNew("GLib", "get_value")
	})
}

var getValueKeyFileInvoker *gi.Function

// GetValue is a representation of the C type g_key_file_get_value.
func (recv *KeyFile) GetValue(groupName string, key string) string {
	keyFileGetValueFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	ret := keyFileGetValueFunction.Invoke(inArgs[:], nil)

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

var keyFileRefFunction *gi.Function
var keyFileRefFunction_Once sync.Once

func keyFileRefFunction_Set() {
	keyFileRefFunction_Once.Do(func() {
		keyFileRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refKeyFileInvoker *gi.Function

// Ref is a representation of the C type g_key_file_ref.
func (recv *KeyFile) Ref() *KeyFile {
	keyFileRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := keyFileRefFunction.Invoke(inArgs[:], nil)

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

var keyFileSetInt64Function *gi.Function
var keyFileSetInt64Function_Once sync.Once

func keyFileSetInt64Function_Set() {
	keyFileSetInt64Function_Once.Do(func() {
		keyFileSetInt64Function = gi.FunctionInvokerNew("GLib", "set_int64")
	})
}

var setInt64KeyFileInvoker *gi.Function

// SetInt64 is a representation of the C type g_key_file_set_int64.
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) {
	keyFileSetInt64Function_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetInt64(value)

	keyFileSetInt64Function.Invoke(inArgs[:], nil)

}

var keyFileSetIntegerFunction *gi.Function
var keyFileSetIntegerFunction_Once sync.Once

func keyFileSetIntegerFunction_Set() {
	keyFileSetIntegerFunction_Once.Do(func() {
		keyFileSetIntegerFunction = gi.FunctionInvokerNew("GLib", "set_integer")
	})
}

var setIntegerKeyFileInvoker *gi.Function

// SetInteger is a representation of the C type g_key_file_set_integer.
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) {
	keyFileSetIntegerFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetInt32(value)

	keyFileSetIntegerFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_set_integer_list' : parameter 'list' has no type

var keyFileSetListSeparatorFunction *gi.Function
var keyFileSetListSeparatorFunction_Once sync.Once

func keyFileSetListSeparatorFunction_Set() {
	keyFileSetListSeparatorFunction_Once.Do(func() {
		keyFileSetListSeparatorFunction = gi.FunctionInvokerNew("GLib", "set_list_separator")
	})
}

var setListSeparatorKeyFileInvoker *gi.Function

// SetListSeparator is a representation of the C type g_key_file_set_list_separator.
func (recv *KeyFile) SetListSeparator(separator int8) {
	keyFileSetListSeparatorFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(separator)

	keyFileSetListSeparatorFunction.Invoke(inArgs[:], nil)

}

var keyFileSetLocaleStringFunction *gi.Function
var keyFileSetLocaleStringFunction_Once sync.Once

func keyFileSetLocaleStringFunction_Set() {
	keyFileSetLocaleStringFunction_Once.Do(func() {
		keyFileSetLocaleStringFunction = gi.FunctionInvokerNew("GLib", "set_locale_string")
	})
}

var setLocaleStringKeyFileInvoker *gi.Function

// SetLocaleString is a representation of the C type g_key_file_set_locale_string.
func (recv *KeyFile) SetLocaleString(groupName string, key string, locale string, string_ string) {
	keyFileSetLocaleStringFunction_Set()

	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)
	inArgs[4].SetString(string_)

	keyFileSetLocaleStringFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_set_locale_string_list' : parameter 'list' has no type

var keyFileSetStringFunction *gi.Function
var keyFileSetStringFunction_Once sync.Once

func keyFileSetStringFunction_Set() {
	keyFileSetStringFunction_Once.Do(func() {
		keyFileSetStringFunction = gi.FunctionInvokerNew("GLib", "set_string")
	})
}

var setStringKeyFileInvoker *gi.Function

// SetString is a representation of the C type g_key_file_set_string.
func (recv *KeyFile) SetString(groupName string, key string, string_ string) {
	keyFileSetStringFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(string_)

	keyFileSetStringFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_set_string_list' : parameter 'list' has no type

var keyFileSetUint64Function *gi.Function
var keyFileSetUint64Function_Once sync.Once

func keyFileSetUint64Function_Set() {
	keyFileSetUint64Function_Once.Do(func() {
		keyFileSetUint64Function = gi.FunctionInvokerNew("GLib", "set_uint64")
	})
}

var setUint64KeyFileInvoker *gi.Function

// SetUint64 is a representation of the C type g_key_file_set_uint64.
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) {
	keyFileSetUint64Function_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetUint64(value)

	keyFileSetUint64Function.Invoke(inArgs[:], nil)

}

var keyFileSetValueFunction *gi.Function
var keyFileSetValueFunction_Once sync.Once

func keyFileSetValueFunction_Set() {
	keyFileSetValueFunction_Once.Do(func() {
		keyFileSetValueFunction = gi.FunctionInvokerNew("GLib", "set_value")
	})
}

var setValueKeyFileInvoker *gi.Function

// SetValue is a representation of the C type g_key_file_set_value.
func (recv *KeyFile) SetValue(groupName string, key string, value string) {
	keyFileSetValueFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(value)

	keyFileSetValueFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_key_file_to_data' : parameter 'length' of type 'gsize' not supported

var keyFileUnrefFunction *gi.Function
var keyFileUnrefFunction_Once sync.Once

func keyFileUnrefFunction_Set() {
	keyFileUnrefFunction_Once.Do(func() {
		keyFileUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefKeyFileInvoker *gi.Function

// Unref is a representation of the C type g_key_file_unref.
func (recv *KeyFile) Unref() {
	keyFileUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	keyFileUnrefFunction.Invoke(inArgs[:], nil)

}

var listStruct *gi.Struct
var listStructOnce sync.Once

func listStructSet() {
	listStructOnce.Do(func() {
		listStruct = gi.StructNew("GLib", "List")
	})
}

type List struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.List'
	// UNSUPPORTED : C value 'prev' : no Go type for 'GLib.List'
}

var logFieldStruct *gi.Struct
var logFieldStructOnce sync.Once

func logFieldStructSet() {
	logFieldStructOnce.Do(func() {
		logFieldStruct = gi.StructNew("GLib", "LogField")
	})
}

type LogField struct {
	native uintptr
	Key    string
	// UNSUPPORTED : C value 'value' : no Go type for 'gpointer'
	Length int32
}

var mainContextStruct *gi.Struct
var mainContextStructOnce sync.Once

func mainContextStructSet() {
	mainContextStructOnce.Do(func() {
		mainContextStruct = gi.StructNew("GLib", "MainContext")
	})
}

type MainContext struct {
	native uintptr
}

var mainContextNewFunction *gi.Function
var mainContextNewFunction_Once sync.Once

func mainContextNewFunction_Set() {
	mainContextNewFunction_Once.Do(func() {
		mainContextNewFunction = gi.FunctionInvokerNew("GLib", "new")
	})
}

var newMainContextInvoker *gi.Function

// MainContextNew is a representation of the C type g_main_context_new.
func MainContextNew() *MainContext {
	mainContextNewFunction_Set()

	ret := mainContextNewFunction.Invoke(nil, nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_main_context_acquire' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_main_context_add_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_main_context_check' : parameter 'fds' has no type

var mainContextDispatchFunction *gi.Function
var mainContextDispatchFunction_Once sync.Once

func mainContextDispatchFunction_Set() {
	mainContextDispatchFunction_Once.Do(func() {
		mainContextDispatchFunction = gi.FunctionInvokerNew("GLib", "dispatch")
	})
}

var dispatchMainContextInvoker *gi.Function

// Dispatch is a representation of the C type g_main_context_dispatch.
func (recv *MainContext) Dispatch() {
	mainContextDispatchFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainContextDispatchFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_find_source_by_funcs_user_data' : parameter 'funcs' of type 'SourceFuncs' not supported

var mainContextFindSourceByIdFunction *gi.Function
var mainContextFindSourceByIdFunction_Once sync.Once

func mainContextFindSourceByIdFunction_Set() {
	mainContextFindSourceByIdFunction_Once.Do(func() {
		mainContextFindSourceByIdFunction = gi.FunctionInvokerNew("GLib", "find_source_by_id")
	})
}

var findSourceByIdMainContextInvoker *gi.Function

// FindSourceById is a representation of the C type g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	mainContextFindSourceByIdFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(sourceId)

	ret := mainContextFindSourceByIdFunction.Invoke(inArgs[:], nil)

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

var mainContextPopThreadDefaultFunction *gi.Function
var mainContextPopThreadDefaultFunction_Once sync.Once

func mainContextPopThreadDefaultFunction_Set() {
	mainContextPopThreadDefaultFunction_Once.Do(func() {
		mainContextPopThreadDefaultFunction = gi.FunctionInvokerNew("GLib", "pop_thread_default")
	})
}

var popThreadDefaultMainContextInvoker *gi.Function

// PopThreadDefault is a representation of the C type g_main_context_pop_thread_default.
func (recv *MainContext) PopThreadDefault() {
	mainContextPopThreadDefaultFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainContextPopThreadDefaultFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_prepare' : return type 'gboolean' not supported

var mainContextPushThreadDefaultFunction *gi.Function
var mainContextPushThreadDefaultFunction_Once sync.Once

func mainContextPushThreadDefaultFunction_Set() {
	mainContextPushThreadDefaultFunction_Once.Do(func() {
		mainContextPushThreadDefaultFunction = gi.FunctionInvokerNew("GLib", "push_thread_default")
	})
}

var pushThreadDefaultMainContextInvoker *gi.Function

// PushThreadDefault is a representation of the C type g_main_context_push_thread_default.
func (recv *MainContext) PushThreadDefault() {
	mainContextPushThreadDefaultFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainContextPushThreadDefaultFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_query' : parameter 'fds' has no type

var mainContextRefFunction *gi.Function
var mainContextRefFunction_Once sync.Once

func mainContextRefFunction_Set() {
	mainContextRefFunction_Once.Do(func() {
		mainContextRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refMainContextInvoker *gi.Function

// Ref is a representation of the C type g_main_context_ref.
func (recv *MainContext) Ref() *MainContext {
	mainContextRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mainContextRefFunction.Invoke(inArgs[:], nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

var mainContextReleaseFunction *gi.Function
var mainContextReleaseFunction_Once sync.Once

func mainContextReleaseFunction_Set() {
	mainContextReleaseFunction_Once.Do(func() {
		mainContextReleaseFunction = gi.FunctionInvokerNew("GLib", "release")
	})
}

var releaseMainContextInvoker *gi.Function

// Release is a representation of the C type g_main_context_release.
func (recv *MainContext) Release() {
	mainContextReleaseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainContextReleaseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_remove_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_main_context_set_poll_func' : parameter 'func' of type 'PollFunc' not supported

var mainContextUnrefFunction *gi.Function
var mainContextUnrefFunction_Once sync.Once

func mainContextUnrefFunction_Set() {
	mainContextUnrefFunction_Once.Do(func() {
		mainContextUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefMainContextInvoker *gi.Function

// Unref is a representation of the C type g_main_context_unref.
func (recv *MainContext) Unref() {
	mainContextUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainContextUnrefFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_main_context_wait' : parameter 'cond' of type 'Cond' not supported

var mainContextWakeupFunction *gi.Function
var mainContextWakeupFunction_Once sync.Once

func mainContextWakeupFunction_Set() {
	mainContextWakeupFunction_Once.Do(func() {
		mainContextWakeupFunction = gi.FunctionInvokerNew("GLib", "wakeup")
	})
}

var wakeupMainContextInvoker *gi.Function

// Wakeup is a representation of the C type g_main_context_wakeup.
func (recv *MainContext) Wakeup() {
	mainContextWakeupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainContextWakeupFunction.Invoke(inArgs[:], nil)

}

var mainLoopStruct *gi.Struct
var mainLoopStructOnce sync.Once

func mainLoopStructSet() {
	mainLoopStructOnce.Do(func() {
		mainLoopStruct = gi.StructNew("GLib", "MainLoop")
	})
}

type MainLoop struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_main_loop_new' : parameter 'context' of type 'MainContext' not supported

var mainLoopGetContextFunction *gi.Function
var mainLoopGetContextFunction_Once sync.Once

func mainLoopGetContextFunction_Set() {
	mainLoopGetContextFunction_Once.Do(func() {
		mainLoopGetContextFunction = gi.FunctionInvokerNew("GLib", "get_context")
	})
}

var getContextMainLoopInvoker *gi.Function

// GetContext is a representation of the C type g_main_loop_get_context.
func (recv *MainLoop) GetContext() *MainContext {
	mainLoopGetContextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mainLoopGetContextFunction.Invoke(inArgs[:], nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_main_loop_is_running' : return type 'gboolean' not supported

var mainLoopQuitFunction *gi.Function
var mainLoopQuitFunction_Once sync.Once

func mainLoopQuitFunction_Set() {
	mainLoopQuitFunction_Once.Do(func() {
		mainLoopQuitFunction = gi.FunctionInvokerNew("GLib", "quit")
	})
}

var quitMainLoopInvoker *gi.Function

// Quit is a representation of the C type g_main_loop_quit.
func (recv *MainLoop) Quit() {
	mainLoopQuitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainLoopQuitFunction.Invoke(inArgs[:], nil)

}

var mainLoopRefFunction *gi.Function
var mainLoopRefFunction_Once sync.Once

func mainLoopRefFunction_Set() {
	mainLoopRefFunction_Once.Do(func() {
		mainLoopRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refMainLoopInvoker *gi.Function

// Ref is a representation of the C type g_main_loop_ref.
func (recv *MainLoop) Ref() *MainLoop {
	mainLoopRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mainLoopRefFunction.Invoke(inArgs[:], nil)

	retGo := &MainLoop{native: ret.Pointer()}

	return retGo
}

var mainLoopRunFunction *gi.Function
var mainLoopRunFunction_Once sync.Once

func mainLoopRunFunction_Set() {
	mainLoopRunFunction_Once.Do(func() {
		mainLoopRunFunction = gi.FunctionInvokerNew("GLib", "run")
	})
}

var runMainLoopInvoker *gi.Function

// Run is a representation of the C type g_main_loop_run.
func (recv *MainLoop) Run() {
	mainLoopRunFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainLoopRunFunction.Invoke(inArgs[:], nil)

}

var mainLoopUnrefFunction *gi.Function
var mainLoopUnrefFunction_Once sync.Once

func mainLoopUnrefFunction_Set() {
	mainLoopUnrefFunction_Once.Do(func() {
		mainLoopUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefMainLoopInvoker *gi.Function

// Unref is a representation of the C type g_main_loop_unref.
func (recv *MainLoop) Unref() {
	mainLoopUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mainLoopUnrefFunction.Invoke(inArgs[:], nil)

}

var mappedFileStruct *gi.Struct
var mappedFileStructOnce sync.Once

func mappedFileStructSet() {
	mappedFileStructOnce.Do(func() {
		mappedFileStruct = gi.StructNew("GLib", "MappedFile")
	})
}

type MappedFile struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_mapped_file_new' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mapped_file_new_from_fd' : parameter 'writable' of type 'gboolean' not supported

var mappedFileFreeFunction *gi.Function
var mappedFileFreeFunction_Once sync.Once

func mappedFileFreeFunction_Set() {
	mappedFileFreeFunction_Once.Do(func() {
		mappedFileFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeMappedFileInvoker *gi.Function

// Free is a representation of the C type g_mapped_file_free.
func (recv *MappedFile) Free() {
	mappedFileFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mappedFileFreeFunction.Invoke(inArgs[:], nil)

}

var mappedFileGetBytesFunction *gi.Function
var mappedFileGetBytesFunction_Once sync.Once

func mappedFileGetBytesFunction_Set() {
	mappedFileGetBytesFunction_Once.Do(func() {
		mappedFileGetBytesFunction = gi.FunctionInvokerNew("GLib", "get_bytes")
	})
}

var getBytesMappedFileInvoker *gi.Function

// GetBytes is a representation of the C type g_mapped_file_get_bytes.
func (recv *MappedFile) GetBytes() *Bytes {
	mappedFileGetBytesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mappedFileGetBytesFunction.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

var mappedFileGetContentsFunction *gi.Function
var mappedFileGetContentsFunction_Once sync.Once

func mappedFileGetContentsFunction_Set() {
	mappedFileGetContentsFunction_Once.Do(func() {
		mappedFileGetContentsFunction = gi.FunctionInvokerNew("GLib", "get_contents")
	})
}

var getContentsMappedFileInvoker *gi.Function

// GetContents is a representation of the C type g_mapped_file_get_contents.
func (recv *MappedFile) GetContents() string {
	mappedFileGetContentsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mappedFileGetContentsFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_mapped_file_get_length' : return type 'gsize' not supported

var mappedFileRefFunction *gi.Function
var mappedFileRefFunction_Once sync.Once

func mappedFileRefFunction_Set() {
	mappedFileRefFunction_Once.Do(func() {
		mappedFileRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refMappedFileInvoker *gi.Function

// Ref is a representation of the C type g_mapped_file_ref.
func (recv *MappedFile) Ref() *MappedFile {
	mappedFileRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := mappedFileRefFunction.Invoke(inArgs[:], nil)

	retGo := &MappedFile{native: ret.Pointer()}

	return retGo
}

var mappedFileUnrefFunction *gi.Function
var mappedFileUnrefFunction_Once sync.Once

func mappedFileUnrefFunction_Set() {
	mappedFileUnrefFunction_Once.Do(func() {
		mappedFileUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefMappedFileInvoker *gi.Function

// Unref is a representation of the C type g_mapped_file_unref.
func (recv *MappedFile) Unref() {
	mappedFileUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	mappedFileUnrefFunction.Invoke(inArgs[:], nil)

}

var markupParseContextStruct *gi.Struct
var markupParseContextStructOnce sync.Once

func markupParseContextStructSet() {
	markupParseContextStructOnce.Do(func() {
		markupParseContextStruct = gi.StructNew("GLib", "MarkupParseContext")
	})
}

type MarkupParseContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_markup_parse_context_new' : parameter 'parser' of type 'MarkupParser' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_end_parse' : return type 'gboolean' not supported

var markupParseContextFreeFunction *gi.Function
var markupParseContextFreeFunction_Once sync.Once

func markupParseContextFreeFunction_Set() {
	markupParseContextFreeFunction_Once.Do(func() {
		markupParseContextFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeMarkupParseContextInvoker *gi.Function

// Free is a representation of the C type g_markup_parse_context_free.
func (recv *MarkupParseContext) Free() {
	markupParseContextFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	markupParseContextFreeFunction.Invoke(inArgs[:], nil)

}

var markupParseContextGetElementFunction *gi.Function
var markupParseContextGetElementFunction_Once sync.Once

func markupParseContextGetElementFunction_Set() {
	markupParseContextGetElementFunction_Once.Do(func() {
		markupParseContextGetElementFunction = gi.FunctionInvokerNew("GLib", "get_element")
	})
}

var getElementMarkupParseContextInvoker *gi.Function

// GetElement is a representation of the C type g_markup_parse_context_get_element.
func (recv *MarkupParseContext) GetElement() string {
	markupParseContextGetElementFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := markupParseContextGetElementFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_markup_parse_context_get_element_stack' : return type 'GLib.SList' not supported

var markupParseContextGetPositionFunction *gi.Function
var markupParseContextGetPositionFunction_Once sync.Once

func markupParseContextGetPositionFunction_Set() {
	markupParseContextGetPositionFunction_Once.Do(func() {
		markupParseContextGetPositionFunction = gi.FunctionInvokerNew("GLib", "get_position")
	})
}

var getPositionMarkupParseContextInvoker *gi.Function

// GetPosition is a representation of the C type g_markup_parse_context_get_position.
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) {
	markupParseContextGetPositionFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(lineNumber)
	inArgs[2].SetInt32(charNumber)

	markupParseContextGetPositionFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_markup_parse_context_get_user_data' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_parse' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_push' : parameter 'parser' of type 'MarkupParser' not supported

var markupParseContextRefFunction *gi.Function
var markupParseContextRefFunction_Once sync.Once

func markupParseContextRefFunction_Set() {
	markupParseContextRefFunction_Once.Do(func() {
		markupParseContextRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refMarkupParseContextInvoker *gi.Function

// Ref is a representation of the C type g_markup_parse_context_ref.
func (recv *MarkupParseContext) Ref() *MarkupParseContext {
	markupParseContextRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := markupParseContextRefFunction.Invoke(inArgs[:], nil)

	retGo := &MarkupParseContext{native: ret.Pointer()}

	return retGo
}

var markupParseContextUnrefFunction *gi.Function
var markupParseContextUnrefFunction_Once sync.Once

func markupParseContextUnrefFunction_Set() {
	markupParseContextUnrefFunction_Once.Do(func() {
		markupParseContextUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefMarkupParseContextInvoker *gi.Function

// Unref is a representation of the C type g_markup_parse_context_unref.
func (recv *MarkupParseContext) Unref() {
	markupParseContextUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	markupParseContextUnrefFunction.Invoke(inArgs[:], nil)

}

var markupParserStruct *gi.Struct
var markupParserStructOnce sync.Once

func markupParserStructSet() {
	markupParserStructOnce.Do(func() {
		markupParserStruct = gi.StructNew("GLib", "MarkupParser")
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

var matchInfoStruct *gi.Struct
var matchInfoStructOnce sync.Once

func matchInfoStructSet() {
	matchInfoStructOnce.Do(func() {
		matchInfoStruct = gi.StructNew("GLib", "MatchInfo")
	})
}

type MatchInfo struct {
	native uintptr
}

var matchInfoExpandReferencesFunction *gi.Function
var matchInfoExpandReferencesFunction_Once sync.Once

func matchInfoExpandReferencesFunction_Set() {
	matchInfoExpandReferencesFunction_Once.Do(func() {
		matchInfoExpandReferencesFunction = gi.FunctionInvokerNew("GLib", "expand_references")
	})
}

var expandReferencesMatchInfoInvoker *gi.Function

// ExpandReferences is a representation of the C type g_match_info_expand_references.
func (recv *MatchInfo) ExpandReferences(stringToExpand string) string {
	matchInfoExpandReferencesFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(stringToExpand)

	ret := matchInfoExpandReferencesFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var matchInfoFetchFunction *gi.Function
var matchInfoFetchFunction_Once sync.Once

func matchInfoFetchFunction_Set() {
	matchInfoFetchFunction_Once.Do(func() {
		matchInfoFetchFunction = gi.FunctionInvokerNew("GLib", "fetch")
	})
}

var fetchMatchInfoInvoker *gi.Function

// Fetch is a representation of the C type g_match_info_fetch.
func (recv *MatchInfo) Fetch(matchNum int32) string {
	matchInfoFetchFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(matchNum)

	ret := matchInfoFetchFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var matchInfoFetchAllFunction *gi.Function
var matchInfoFetchAllFunction_Once sync.Once

func matchInfoFetchAllFunction_Set() {
	matchInfoFetchAllFunction_Once.Do(func() {
		matchInfoFetchAllFunction = gi.FunctionInvokerNew("GLib", "fetch_all")
	})
}

var fetchAllMatchInfoInvoker *gi.Function

// FetchAll is a representation of the C type g_match_info_fetch_all.
func (recv *MatchInfo) FetchAll() {
	matchInfoFetchAllFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	matchInfoFetchAllFunction.Invoke(inArgs[:], nil)

}

var matchInfoFetchNamedFunction *gi.Function
var matchInfoFetchNamedFunction_Once sync.Once

func matchInfoFetchNamedFunction_Set() {
	matchInfoFetchNamedFunction_Once.Do(func() {
		matchInfoFetchNamedFunction = gi.FunctionInvokerNew("GLib", "fetch_named")
	})
}

var fetchNamedMatchInfoInvoker *gi.Function

// FetchNamed is a representation of the C type g_match_info_fetch_named.
func (recv *MatchInfo) FetchNamed(name string) string {
	matchInfoFetchNamedFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := matchInfoFetchNamedFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_match_info_fetch_named_pos' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_match_info_fetch_pos' : return type 'gboolean' not supported

var matchInfoFreeFunction *gi.Function
var matchInfoFreeFunction_Once sync.Once

func matchInfoFreeFunction_Set() {
	matchInfoFreeFunction_Once.Do(func() {
		matchInfoFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeMatchInfoInvoker *gi.Function

// Free is a representation of the C type g_match_info_free.
func (recv *MatchInfo) Free() {
	matchInfoFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	matchInfoFreeFunction.Invoke(inArgs[:], nil)

}

var matchInfoGetMatchCountFunction *gi.Function
var matchInfoGetMatchCountFunction_Once sync.Once

func matchInfoGetMatchCountFunction_Set() {
	matchInfoGetMatchCountFunction_Once.Do(func() {
		matchInfoGetMatchCountFunction = gi.FunctionInvokerNew("GLib", "get_match_count")
	})
}

var getMatchCountMatchInfoInvoker *gi.Function

// GetMatchCount is a representation of the C type g_match_info_get_match_count.
func (recv *MatchInfo) GetMatchCount() int32 {
	matchInfoGetMatchCountFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := matchInfoGetMatchCountFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var matchInfoGetRegexFunction *gi.Function
var matchInfoGetRegexFunction_Once sync.Once

func matchInfoGetRegexFunction_Set() {
	matchInfoGetRegexFunction_Once.Do(func() {
		matchInfoGetRegexFunction = gi.FunctionInvokerNew("GLib", "get_regex")
	})
}

var getRegexMatchInfoInvoker *gi.Function

// GetRegex is a representation of the C type g_match_info_get_regex.
func (recv *MatchInfo) GetRegex() *Regex {
	matchInfoGetRegexFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := matchInfoGetRegexFunction.Invoke(inArgs[:], nil)

	retGo := &Regex{native: ret.Pointer()}

	return retGo
}

var matchInfoGetStringFunction *gi.Function
var matchInfoGetStringFunction_Once sync.Once

func matchInfoGetStringFunction_Set() {
	matchInfoGetStringFunction_Once.Do(func() {
		matchInfoGetStringFunction = gi.FunctionInvokerNew("GLib", "get_string")
	})
}

var getStringMatchInfoInvoker *gi.Function

// GetString is a representation of the C type g_match_info_get_string.
func (recv *MatchInfo) GetString() string {
	matchInfoGetStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := matchInfoGetStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_match_info_is_partial_match' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_match_info_matches' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_match_info_next' : return type 'gboolean' not supported

var matchInfoRefFunction *gi.Function
var matchInfoRefFunction_Once sync.Once

func matchInfoRefFunction_Set() {
	matchInfoRefFunction_Once.Do(func() {
		matchInfoRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refMatchInfoInvoker *gi.Function

// Ref is a representation of the C type g_match_info_ref.
func (recv *MatchInfo) Ref() *MatchInfo {
	matchInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := matchInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &MatchInfo{native: ret.Pointer()}

	return retGo
}

var matchInfoUnrefFunction *gi.Function
var matchInfoUnrefFunction_Once sync.Once

func matchInfoUnrefFunction_Set() {
	matchInfoUnrefFunction_Once.Do(func() {
		matchInfoUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefMatchInfoInvoker *gi.Function

// Unref is a representation of the C type g_match_info_unref.
func (recv *MatchInfo) Unref() {
	matchInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	matchInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var memVTableStruct *gi.Struct
var memVTableStructOnce sync.Once

func memVTableStructSet() {
	memVTableStructOnce.Do(func() {
		memVTableStruct = gi.StructNew("GLib", "MemVTable")
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

var nodeStruct *gi.Struct
var nodeStructOnce sync.Once

func nodeStructSet() {
	nodeStructOnce.Do(func() {
		nodeStruct = gi.StructNew("GLib", "Node")
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

var nodeCopyFunction *gi.Function
var nodeCopyFunction_Once sync.Once

func nodeCopyFunction_Set() {
	nodeCopyFunction_Once.Do(func() {
		nodeCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyNodeInvoker *gi.Function

// Copy is a representation of the C type g_node_copy.
func (recv *Node) Copy() *Node {
	nodeCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_node_copy_deep' : parameter 'copy_func' of type 'CopyFunc' not supported

var nodeDepthFunction *gi.Function
var nodeDepthFunction_Once sync.Once

func nodeDepthFunction_Set() {
	nodeDepthFunction_Once.Do(func() {
		nodeDepthFunction = gi.FunctionInvokerNew("GLib", "depth")
	})
}

var depthNodeInvoker *gi.Function

// Depth is a representation of the C type g_node_depth.
func (recv *Node) Depth() uint32 {
	nodeDepthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeDepthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var nodeDestroyFunction *gi.Function
var nodeDestroyFunction_Once sync.Once

func nodeDestroyFunction_Set() {
	nodeDestroyFunction_Once.Do(func() {
		nodeDestroyFunction = gi.FunctionInvokerNew("GLib", "destroy")
	})
}

var destroyNodeInvoker *gi.Function

// Destroy is a representation of the C type g_node_destroy.
func (recv *Node) Destroy() {
	nodeDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	nodeDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_node_find' : parameter 'order' of type 'TraverseType' not supported

// UNSUPPORTED : C value 'g_node_find_child' : parameter 'flags' of type 'TraverseFlags' not supported

var nodeFirstSiblingFunction *gi.Function
var nodeFirstSiblingFunction_Once sync.Once

func nodeFirstSiblingFunction_Set() {
	nodeFirstSiblingFunction_Once.Do(func() {
		nodeFirstSiblingFunction = gi.FunctionInvokerNew("GLib", "first_sibling")
	})
}

var firstSiblingNodeInvoker *gi.Function

// FirstSibling is a representation of the C type g_node_first_sibling.
func (recv *Node) FirstSibling() *Node {
	nodeFirstSiblingFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeFirstSiblingFunction.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

var nodeGetRootFunction *gi.Function
var nodeGetRootFunction_Once sync.Once

func nodeGetRootFunction_Set() {
	nodeGetRootFunction_Once.Do(func() {
		nodeGetRootFunction = gi.FunctionInvokerNew("GLib", "get_root")
	})
}

var getRootNodeInvoker *gi.Function

// GetRoot is a representation of the C type g_node_get_root.
func (recv *Node) GetRoot() *Node {
	nodeGetRootFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeGetRootFunction.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_node_insert' : parameter 'node' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_insert_after' : parameter 'sibling' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_insert_before' : parameter 'sibling' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_is_ancestor' : parameter 'descendant' of type 'Node' not supported

var nodeLastChildFunction *gi.Function
var nodeLastChildFunction_Once sync.Once

func nodeLastChildFunction_Set() {
	nodeLastChildFunction_Once.Do(func() {
		nodeLastChildFunction = gi.FunctionInvokerNew("GLib", "last_child")
	})
}

var lastChildNodeInvoker *gi.Function

// LastChild is a representation of the C type g_node_last_child.
func (recv *Node) LastChild() *Node {
	nodeLastChildFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeLastChildFunction.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

var nodeLastSiblingFunction *gi.Function
var nodeLastSiblingFunction_Once sync.Once

func nodeLastSiblingFunction_Set() {
	nodeLastSiblingFunction_Once.Do(func() {
		nodeLastSiblingFunction = gi.FunctionInvokerNew("GLib", "last_sibling")
	})
}

var lastSiblingNodeInvoker *gi.Function

// LastSibling is a representation of the C type g_node_last_sibling.
func (recv *Node) LastSibling() *Node {
	nodeLastSiblingFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeLastSiblingFunction.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

var nodeMaxHeightFunction *gi.Function
var nodeMaxHeightFunction_Once sync.Once

func nodeMaxHeightFunction_Set() {
	nodeMaxHeightFunction_Once.Do(func() {
		nodeMaxHeightFunction = gi.FunctionInvokerNew("GLib", "max_height")
	})
}

var maxHeightNodeInvoker *gi.Function

// MaxHeight is a representation of the C type g_node_max_height.
func (recv *Node) MaxHeight() uint32 {
	nodeMaxHeightFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeMaxHeightFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var nodeNChildrenFunction *gi.Function
var nodeNChildrenFunction_Once sync.Once

func nodeNChildrenFunction_Set() {
	nodeNChildrenFunction_Once.Do(func() {
		nodeNChildrenFunction = gi.FunctionInvokerNew("GLib", "n_children")
	})
}

var nChildrenNodeInvoker *gi.Function

// NChildren is a representation of the C type g_node_n_children.
func (recv *Node) NChildren() uint32 {
	nodeNChildrenFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := nodeNChildrenFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_node_n_nodes' : parameter 'flags' of type 'TraverseFlags' not supported

var nodeNthChildFunction *gi.Function
var nodeNthChildFunction_Once sync.Once

func nodeNthChildFunction_Set() {
	nodeNthChildFunction_Once.Do(func() {
		nodeNthChildFunction = gi.FunctionInvokerNew("GLib", "nth_child")
	})
}

var nthChildNodeInvoker *gi.Function

// NthChild is a representation of the C type g_node_nth_child.
func (recv *Node) NthChild(n uint32) *Node {
	nodeNthChildFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(n)

	ret := nodeNthChildFunction.Invoke(inArgs[:], nil)

	retGo := &Node{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_node_prepend' : parameter 'node' of type 'Node' not supported

var nodeReverseChildrenFunction *gi.Function
var nodeReverseChildrenFunction_Once sync.Once

func nodeReverseChildrenFunction_Set() {
	nodeReverseChildrenFunction_Once.Do(func() {
		nodeReverseChildrenFunction = gi.FunctionInvokerNew("GLib", "reverse_children")
	})
}

var reverseChildrenNodeInvoker *gi.Function

// ReverseChildren is a representation of the C type g_node_reverse_children.
func (recv *Node) ReverseChildren() {
	nodeReverseChildrenFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	nodeReverseChildrenFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_node_traverse' : parameter 'order' of type 'TraverseType' not supported

var nodeUnlinkFunction *gi.Function
var nodeUnlinkFunction_Once sync.Once

func nodeUnlinkFunction_Set() {
	nodeUnlinkFunction_Once.Do(func() {
		nodeUnlinkFunction = gi.FunctionInvokerNew("GLib", "unlink")
	})
}

var unlinkNodeInvoker *gi.Function

// Unlink is a representation of the C type g_node_unlink.
func (recv *Node) Unlink() {
	nodeUnlinkFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	nodeUnlinkFunction.Invoke(inArgs[:], nil)

}

var onceStruct *gi.Struct
var onceStructOnce sync.Once

func onceStructSet() {
	onceStructOnce.Do(func() {
		onceStruct = gi.StructNew("GLib", "Once")
	})
}

type Once struct {
	native uintptr
	// UNSUPPORTED : C value 'status' : no Go type for 'OnceStatus'
	// UNSUPPORTED : C value 'retval' : no Go type for 'gpointer'
}

// UNSUPPORTED : C value 'g_once_impl' : parameter 'func' of type 'ThreadFunc' not supported

var optionContextStruct *gi.Struct
var optionContextStructOnce sync.Once

func optionContextStructSet() {
	optionContextStructOnce.Do(func() {
		optionContextStruct = gi.StructNew("GLib", "OptionContext")
	})
}

type OptionContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_context_add_group' : parameter 'group' of type 'OptionGroup' not supported

// UNSUPPORTED : C value 'g_option_context_add_main_entries' : parameter 'entries' of type 'OptionEntry' not supported

var optionContextFreeFunction *gi.Function
var optionContextFreeFunction_Once sync.Once

func optionContextFreeFunction_Set() {
	optionContextFreeFunction_Once.Do(func() {
		optionContextFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeOptionContextInvoker *gi.Function

// Free is a representation of the C type g_option_context_free.
func (recv *OptionContext) Free() {
	optionContextFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	optionContextFreeFunction.Invoke(inArgs[:], nil)

}

var optionContextGetDescriptionFunction *gi.Function
var optionContextGetDescriptionFunction_Once sync.Once

func optionContextGetDescriptionFunction_Set() {
	optionContextGetDescriptionFunction_Once.Do(func() {
		optionContextGetDescriptionFunction = gi.FunctionInvokerNew("GLib", "get_description")
	})
}

var getDescriptionOptionContextInvoker *gi.Function

// GetDescription is a representation of the C type g_option_context_get_description.
func (recv *OptionContext) GetDescription() string {
	optionContextGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionContextGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_option_context_get_help' : parameter 'main_help' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_get_help_enabled' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_get_ignore_unknown_options' : return type 'gboolean' not supported

var optionContextGetMainGroupFunction *gi.Function
var optionContextGetMainGroupFunction_Once sync.Once

func optionContextGetMainGroupFunction_Set() {
	optionContextGetMainGroupFunction_Once.Do(func() {
		optionContextGetMainGroupFunction = gi.FunctionInvokerNew("GLib", "get_main_group")
	})
}

var getMainGroupOptionContextInvoker *gi.Function

// GetMainGroup is a representation of the C type g_option_context_get_main_group.
func (recv *OptionContext) GetMainGroup() *OptionGroup {
	optionContextGetMainGroupFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionContextGetMainGroupFunction.Invoke(inArgs[:], nil)

	retGo := &OptionGroup{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_option_context_get_strict_posix' : return type 'gboolean' not supported

var optionContextGetSummaryFunction *gi.Function
var optionContextGetSummaryFunction_Once sync.Once

func optionContextGetSummaryFunction_Set() {
	optionContextGetSummaryFunction_Once.Do(func() {
		optionContextGetSummaryFunction = gi.FunctionInvokerNew("GLib", "get_summary")
	})
}

var getSummaryOptionContextInvoker *gi.Function

// GetSummary is a representation of the C type g_option_context_get_summary.
func (recv *OptionContext) GetSummary() string {
	optionContextGetSummaryFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionContextGetSummaryFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_option_context_parse' : parameter 'argv' has no type

// UNSUPPORTED : C value 'g_option_context_parse_strv' : parameter 'arguments' has no type

var optionContextSetDescriptionFunction *gi.Function
var optionContextSetDescriptionFunction_Once sync.Once

func optionContextSetDescriptionFunction_Set() {
	optionContextSetDescriptionFunction_Once.Do(func() {
		optionContextSetDescriptionFunction = gi.FunctionInvokerNew("GLib", "set_description")
	})
}

var setDescriptionOptionContextInvoker *gi.Function

// SetDescription is a representation of the C type g_option_context_set_description.
func (recv *OptionContext) SetDescription(description string) {
	optionContextSetDescriptionFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(description)

	optionContextSetDescriptionFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_option_context_set_help_enabled' : parameter 'help_enabled' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_set_ignore_unknown_options' : parameter 'ignore_unknown' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_option_context_set_main_group' : parameter 'group' of type 'OptionGroup' not supported

// UNSUPPORTED : C value 'g_option_context_set_strict_posix' : parameter 'strict_posix' of type 'gboolean' not supported

var optionContextSetSummaryFunction *gi.Function
var optionContextSetSummaryFunction_Once sync.Once

func optionContextSetSummaryFunction_Set() {
	optionContextSetSummaryFunction_Once.Do(func() {
		optionContextSetSummaryFunction = gi.FunctionInvokerNew("GLib", "set_summary")
	})
}

var setSummaryOptionContextInvoker *gi.Function

// SetSummary is a representation of the C type g_option_context_set_summary.
func (recv *OptionContext) SetSummary(summary string) {
	optionContextSetSummaryFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(summary)

	optionContextSetSummaryFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_option_context_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

var optionContextSetTranslationDomainFunction *gi.Function
var optionContextSetTranslationDomainFunction_Once sync.Once

func optionContextSetTranslationDomainFunction_Set() {
	optionContextSetTranslationDomainFunction_Once.Do(func() {
		optionContextSetTranslationDomainFunction = gi.FunctionInvokerNew("GLib", "set_translation_domain")
	})
}

var setTranslationDomainOptionContextInvoker *gi.Function

// SetTranslationDomain is a representation of the C type g_option_context_set_translation_domain.
func (recv *OptionContext) SetTranslationDomain(domain string) {
	optionContextSetTranslationDomainFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	optionContextSetTranslationDomainFunction.Invoke(inArgs[:], nil)

}

var optionEntryStruct *gi.Struct
var optionEntryStructOnce sync.Once

func optionEntryStructSet() {
	optionEntryStructOnce.Do(func() {
		optionEntryStruct = gi.StructNew("GLib", "OptionEntry")
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

var optionGroupStruct *gi.Struct
var optionGroupStructOnce sync.Once

func optionGroupStructSet() {
	optionGroupStructOnce.Do(func() {
		optionGroupStruct = gi.StructNew("GLib", "OptionGroup")
	})
}

type OptionGroup struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_group_new' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_option_group_add_entries' : parameter 'entries' of type 'OptionEntry' not supported

var optionGroupFreeFunction *gi.Function
var optionGroupFreeFunction_Once sync.Once

func optionGroupFreeFunction_Set() {
	optionGroupFreeFunction_Once.Do(func() {
		optionGroupFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeOptionGroupInvoker *gi.Function

// Free is a representation of the C type g_option_group_free.
func (recv *OptionGroup) Free() {
	optionGroupFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	optionGroupFreeFunction.Invoke(inArgs[:], nil)

}

var optionGroupRefFunction *gi.Function
var optionGroupRefFunction_Once sync.Once

func optionGroupRefFunction_Set() {
	optionGroupRefFunction_Once.Do(func() {
		optionGroupRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refOptionGroupInvoker *gi.Function

// Ref is a representation of the C type g_option_group_ref.
func (recv *OptionGroup) Ref() *OptionGroup {
	optionGroupRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := optionGroupRefFunction.Invoke(inArgs[:], nil)

	retGo := &OptionGroup{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_option_group_set_error_hook' : parameter 'error_func' of type 'OptionErrorFunc' not supported

// UNSUPPORTED : C value 'g_option_group_set_parse_hooks' : parameter 'pre_parse_func' of type 'OptionParseFunc' not supported

// UNSUPPORTED : C value 'g_option_group_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

var optionGroupSetTranslationDomainFunction *gi.Function
var optionGroupSetTranslationDomainFunction_Once sync.Once

func optionGroupSetTranslationDomainFunction_Set() {
	optionGroupSetTranslationDomainFunction_Once.Do(func() {
		optionGroupSetTranslationDomainFunction = gi.FunctionInvokerNew("GLib", "set_translation_domain")
	})
}

var setTranslationDomainOptionGroupInvoker *gi.Function

// SetTranslationDomain is a representation of the C type g_option_group_set_translation_domain.
func (recv *OptionGroup) SetTranslationDomain(domain string) {
	optionGroupSetTranslationDomainFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	optionGroupSetTranslationDomainFunction.Invoke(inArgs[:], nil)

}

var optionGroupUnrefFunction *gi.Function
var optionGroupUnrefFunction_Once sync.Once

func optionGroupUnrefFunction_Set() {
	optionGroupUnrefFunction_Once.Do(func() {
		optionGroupUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefOptionGroupInvoker *gi.Function

// Unref is a representation of the C type g_option_group_unref.
func (recv *OptionGroup) Unref() {
	optionGroupUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	optionGroupUnrefFunction.Invoke(inArgs[:], nil)

}

var patternSpecStruct *gi.Struct
var patternSpecStructOnce sync.Once

func patternSpecStructSet() {
	patternSpecStructOnce.Do(func() {
		patternSpecStruct = gi.StructNew("GLib", "PatternSpec")
	})
}

type PatternSpec struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_pattern_spec_equal' : parameter 'pspec2' of type 'PatternSpec' not supported

var patternSpecFreeFunction *gi.Function
var patternSpecFreeFunction_Once sync.Once

func patternSpecFreeFunction_Set() {
	patternSpecFreeFunction_Once.Do(func() {
		patternSpecFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freePatternSpecInvoker *gi.Function

// Free is a representation of the C type g_pattern_spec_free.
func (recv *PatternSpec) Free() {
	patternSpecFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	patternSpecFreeFunction.Invoke(inArgs[:], nil)

}

var pollFDStruct *gi.Struct
var pollFDStructOnce sync.Once

func pollFDStructSet() {
	pollFDStructOnce.Do(func() {
		pollFDStruct = gi.StructNew("GLib", "PollFD")
	})
}

type PollFD struct {
	native  uintptr
	Fd      int32
	Events  uint16
	Revents uint16
}

var privateStruct *gi.Struct
var privateStructOnce sync.Once

func privateStructSet() {
	privateStructOnce.Do(func() {
		privateStruct = gi.StructNew("GLib", "Private")
	})
}

type Private struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_private_get' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_private_replace' : parameter 'value' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_private_set' : parameter 'value' of type 'gpointer' not supported

var ptrArrayStruct *gi.Struct
var ptrArrayStructOnce sync.Once

func ptrArrayStructSet() {
	ptrArrayStructOnce.Do(func() {
		ptrArrayStruct = gi.StructNew("GLib", "PtrArray")
	})
}

type PtrArray struct {
	native uintptr
	// UNSUPPORTED : C value 'pdata' : no Go type for 'gpointer'
	Len uint32
}

var queueStruct *gi.Struct
var queueStructOnce sync.Once

func queueStructSet() {
	queueStructOnce.Do(func() {
		queueStruct = gi.StructNew("GLib", "Queue")
	})
}

type Queue struct {
	native uintptr
	// UNSUPPORTED : C value 'head' : no Go type for 'GLib.List'
	// UNSUPPORTED : C value 'tail' : no Go type for 'GLib.List'
	Length uint32
}

var queueClearFunction *gi.Function
var queueClearFunction_Once sync.Once

func queueClearFunction_Set() {
	queueClearFunction_Once.Do(func() {
		queueClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearQueueInvoker *gi.Function

// Clear is a representation of the C type g_queue_clear.
func (recv *Queue) Clear() {
	queueClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	queueClearFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_clear_full' : parameter 'free_func' of type 'DestroyNotify' not supported

var queueCopyFunction *gi.Function
var queueCopyFunction_Once sync.Once

func queueCopyFunction_Set() {
	queueCopyFunction_Once.Do(func() {
		queueCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyQueueInvoker *gi.Function

// Copy is a representation of the C type g_queue_copy.
func (recv *Queue) Copy() *Queue {
	queueCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := queueCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Queue{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_queue_delete_link' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_find' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_find_custom' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_foreach' : parameter 'func' of type 'Func' not supported

var queueFreeFunction *gi.Function
var queueFreeFunction_Once sync.Once

func queueFreeFunction_Set() {
	queueFreeFunction_Once.Do(func() {
		queueFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeQueueInvoker *gi.Function

// Free is a representation of the C type g_queue_free.
func (recv *Queue) Free() {
	queueFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	queueFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_free_full' : parameter 'free_func' of type 'DestroyNotify' not supported

var queueGetLengthFunction *gi.Function
var queueGetLengthFunction_Once sync.Once

func queueGetLengthFunction_Set() {
	queueGetLengthFunction_Once.Do(func() {
		queueGetLengthFunction = gi.FunctionInvokerNew("GLib", "get_length")
	})
}

var getLengthQueueInvoker *gi.Function

// GetLength is a representation of the C type g_queue_get_length.
func (recv *Queue) GetLength() uint32 {
	queueGetLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := queueGetLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_queue_index' : parameter 'data' of type 'gpointer' not supported

var queueInitFunction *gi.Function
var queueInitFunction_Once sync.Once

func queueInitFunction_Set() {
	queueInitFunction_Once.Do(func() {
		queueInitFunction = gi.FunctionInvokerNew("GLib", "init")
	})
}

var initQueueInvoker *gi.Function

// Init is a representation of the C type g_queue_init.
func (recv *Queue) Init() {
	queueInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	queueInitFunction.Invoke(inArgs[:], nil)

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

var queueReverseFunction *gi.Function
var queueReverseFunction_Once sync.Once

func queueReverseFunction_Set() {
	queueReverseFunction_Once.Do(func() {
		queueReverseFunction = gi.FunctionInvokerNew("GLib", "reverse")
	})
}

var reverseQueueInvoker *gi.Function

// Reverse is a representation of the C type g_queue_reverse.
func (recv *Queue) Reverse() {
	queueReverseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	queueReverseFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_queue_sort' : parameter 'compare_func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_queue_unlink' : parameter 'link_' of type 'GLib.List' not supported

var rWLockStruct *gi.Struct
var rWLockStructOnce sync.Once

func rWLockStructSet() {
	rWLockStructOnce.Do(func() {
		rWLockStruct = gi.StructNew("GLib", "RWLock")
	})
}

type RWLock struct {
	native uintptr
}

var rWLockClearFunction *gi.Function
var rWLockClearFunction_Once sync.Once

func rWLockClearFunction_Set() {
	rWLockClearFunction_Once.Do(func() {
		rWLockClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearRWLockInvoker *gi.Function

// Clear is a representation of the C type g_rw_lock_clear.
func (recv *RWLock) Clear() {
	rWLockClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rWLockClearFunction.Invoke(inArgs[:], nil)

}

var rWLockInitFunction *gi.Function
var rWLockInitFunction_Once sync.Once

func rWLockInitFunction_Set() {
	rWLockInitFunction_Once.Do(func() {
		rWLockInitFunction = gi.FunctionInvokerNew("GLib", "init")
	})
}

var initRWLockInvoker *gi.Function

// Init is a representation of the C type g_rw_lock_init.
func (recv *RWLock) Init() {
	rWLockInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rWLockInitFunction.Invoke(inArgs[:], nil)

}

var rWLockReaderLockFunction *gi.Function
var rWLockReaderLockFunction_Once sync.Once

func rWLockReaderLockFunction_Set() {
	rWLockReaderLockFunction_Once.Do(func() {
		rWLockReaderLockFunction = gi.FunctionInvokerNew("GLib", "reader_lock")
	})
}

var readerLockRWLockInvoker *gi.Function

// ReaderLock is a representation of the C type g_rw_lock_reader_lock.
func (recv *RWLock) ReaderLock() {
	rWLockReaderLockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rWLockReaderLockFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_rw_lock_reader_trylock' : return type 'gboolean' not supported

var rWLockReaderUnlockFunction *gi.Function
var rWLockReaderUnlockFunction_Once sync.Once

func rWLockReaderUnlockFunction_Set() {
	rWLockReaderUnlockFunction_Once.Do(func() {
		rWLockReaderUnlockFunction = gi.FunctionInvokerNew("GLib", "reader_unlock")
	})
}

var readerUnlockRWLockInvoker *gi.Function

// ReaderUnlock is a representation of the C type g_rw_lock_reader_unlock.
func (recv *RWLock) ReaderUnlock() {
	rWLockReaderUnlockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rWLockReaderUnlockFunction.Invoke(inArgs[:], nil)

}

var rWLockWriterLockFunction *gi.Function
var rWLockWriterLockFunction_Once sync.Once

func rWLockWriterLockFunction_Set() {
	rWLockWriterLockFunction_Once.Do(func() {
		rWLockWriterLockFunction = gi.FunctionInvokerNew("GLib", "writer_lock")
	})
}

var writerLockRWLockInvoker *gi.Function

// WriterLock is a representation of the C type g_rw_lock_writer_lock.
func (recv *RWLock) WriterLock() {
	rWLockWriterLockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rWLockWriterLockFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_rw_lock_writer_trylock' : return type 'gboolean' not supported

var rWLockWriterUnlockFunction *gi.Function
var rWLockWriterUnlockFunction_Once sync.Once

func rWLockWriterUnlockFunction_Set() {
	rWLockWriterUnlockFunction_Once.Do(func() {
		rWLockWriterUnlockFunction = gi.FunctionInvokerNew("GLib", "writer_unlock")
	})
}

var writerUnlockRWLockInvoker *gi.Function

// WriterUnlock is a representation of the C type g_rw_lock_writer_unlock.
func (recv *RWLock) WriterUnlock() {
	rWLockWriterUnlockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rWLockWriterUnlockFunction.Invoke(inArgs[:], nil)

}

var randStruct *gi.Struct
var randStructOnce sync.Once

func randStructSet() {
	randStructOnce.Do(func() {
		randStruct = gi.StructNew("GLib", "Rand")
	})
}

type Rand struct {
	native uintptr
}

var randCopyFunction *gi.Function
var randCopyFunction_Once sync.Once

func randCopyFunction_Set() {
	randCopyFunction_Once.Do(func() {
		randCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyRandInvoker *gi.Function

// Copy is a representation of the C type g_rand_copy.
func (recv *Rand) Copy() *Rand {
	randCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := randCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Rand{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_rand_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_rand_double_range' : parameter 'begin' of type 'gdouble' not supported

var randFreeFunction *gi.Function
var randFreeFunction_Once sync.Once

func randFreeFunction_Set() {
	randFreeFunction_Once.Do(func() {
		randFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeRandInvoker *gi.Function

// Free is a representation of the C type g_rand_free.
func (recv *Rand) Free() {
	randFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	randFreeFunction.Invoke(inArgs[:], nil)

}

var randIntFunction *gi.Function
var randIntFunction_Once sync.Once

func randIntFunction_Set() {
	randIntFunction_Once.Do(func() {
		randIntFunction = gi.FunctionInvokerNew("GLib", "int")
	})
}

var intRandInvoker *gi.Function

// Int is a representation of the C type g_rand_int.
func (recv *Rand) Int() uint32 {
	randIntFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := randIntFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var randIntRangeFunction *gi.Function
var randIntRangeFunction_Once sync.Once

func randIntRangeFunction_Set() {
	randIntRangeFunction_Once.Do(func() {
		randIntRangeFunction = gi.FunctionInvokerNew("GLib", "int_range")
	})
}

var intRangeRandInvoker *gi.Function

// IntRange is a representation of the C type g_rand_int_range.
func (recv *Rand) IntRange(begin int32, end int32) int32 {
	randIntRangeFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(begin)
	inArgs[2].SetInt32(end)

	ret := randIntRangeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var randSetSeedFunction *gi.Function
var randSetSeedFunction_Once sync.Once

func randSetSeedFunction_Set() {
	randSetSeedFunction_Once.Do(func() {
		randSetSeedFunction = gi.FunctionInvokerNew("GLib", "set_seed")
	})
}

var setSeedRandInvoker *gi.Function

// SetSeed is a representation of the C type g_rand_set_seed.
func (recv *Rand) SetSeed(seed uint32) {
	randSetSeedFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(seed)

	randSetSeedFunction.Invoke(inArgs[:], nil)

}

var randSetSeedArrayFunction *gi.Function
var randSetSeedArrayFunction_Once sync.Once

func randSetSeedArrayFunction_Set() {
	randSetSeedArrayFunction_Once.Do(func() {
		randSetSeedArrayFunction = gi.FunctionInvokerNew("GLib", "set_seed_array")
	})
}

var setSeedArrayRandInvoker *gi.Function

// SetSeedArray is a representation of the C type g_rand_set_seed_array.
func (recv *Rand) SetSeedArray(seed uint32, seedLength uint32) {
	randSetSeedArrayFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(seed)
	inArgs[2].SetUint32(seedLength)

	randSetSeedArrayFunction.Invoke(inArgs[:], nil)

}

var recMutexStruct *gi.Struct
var recMutexStructOnce sync.Once

func recMutexStructSet() {
	recMutexStructOnce.Do(func() {
		recMutexStruct = gi.StructNew("GLib", "RecMutex")
	})
}

type RecMutex struct {
	native uintptr
}

var recMutexClearFunction *gi.Function
var recMutexClearFunction_Once sync.Once

func recMutexClearFunction_Set() {
	recMutexClearFunction_Once.Do(func() {
		recMutexClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearRecMutexInvoker *gi.Function

// Clear is a representation of the C type g_rec_mutex_clear.
func (recv *RecMutex) Clear() {
	recMutexClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	recMutexClearFunction.Invoke(inArgs[:], nil)

}

var recMutexInitFunction *gi.Function
var recMutexInitFunction_Once sync.Once

func recMutexInitFunction_Set() {
	recMutexInitFunction_Once.Do(func() {
		recMutexInitFunction = gi.FunctionInvokerNew("GLib", "init")
	})
}

var initRecMutexInvoker *gi.Function

// Init is a representation of the C type g_rec_mutex_init.
func (recv *RecMutex) Init() {
	recMutexInitFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	recMutexInitFunction.Invoke(inArgs[:], nil)

}

var recMutexLockFunction *gi.Function
var recMutexLockFunction_Once sync.Once

func recMutexLockFunction_Set() {
	recMutexLockFunction_Once.Do(func() {
		recMutexLockFunction = gi.FunctionInvokerNew("GLib", "lock")
	})
}

var lockRecMutexInvoker *gi.Function

// Lock is a representation of the C type g_rec_mutex_lock.
func (recv *RecMutex) Lock() {
	recMutexLockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	recMutexLockFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_rec_mutex_trylock' : return type 'gboolean' not supported

var recMutexUnlockFunction *gi.Function
var recMutexUnlockFunction_Once sync.Once

func recMutexUnlockFunction_Set() {
	recMutexUnlockFunction_Once.Do(func() {
		recMutexUnlockFunction = gi.FunctionInvokerNew("GLib", "unlock")
	})
}

var unlockRecMutexInvoker *gi.Function

// Unlock is a representation of the C type g_rec_mutex_unlock.
func (recv *RecMutex) Unlock() {
	recMutexUnlockFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	recMutexUnlockFunction.Invoke(inArgs[:], nil)

}

var regexStruct *gi.Struct
var regexStructOnce sync.Once

func regexStructSet() {
	regexStructOnce.Do(func() {
		regexStruct = gi.StructNew("GLib", "Regex")
	})
}

type Regex struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_regex_new' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var regexGetCaptureCountFunction *gi.Function
var regexGetCaptureCountFunction_Once sync.Once

func regexGetCaptureCountFunction_Set() {
	regexGetCaptureCountFunction_Once.Do(func() {
		regexGetCaptureCountFunction = gi.FunctionInvokerNew("GLib", "get_capture_count")
	})
}

var getCaptureCountRegexInvoker *gi.Function

// GetCaptureCount is a representation of the C type g_regex_get_capture_count.
func (recv *Regex) GetCaptureCount() int32 {
	regexGetCaptureCountFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := regexGetCaptureCountFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_regex_get_compile_flags' : return type 'RegexCompileFlags' not supported

// UNSUPPORTED : C value 'g_regex_get_has_cr_or_lf' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_regex_get_match_flags' : return type 'RegexMatchFlags' not supported

var regexGetMaxBackrefFunction *gi.Function
var regexGetMaxBackrefFunction_Once sync.Once

func regexGetMaxBackrefFunction_Set() {
	regexGetMaxBackrefFunction_Once.Do(func() {
		regexGetMaxBackrefFunction = gi.FunctionInvokerNew("GLib", "get_max_backref")
	})
}

var getMaxBackrefRegexInvoker *gi.Function

// GetMaxBackref is a representation of the C type g_regex_get_max_backref.
func (recv *Regex) GetMaxBackref() int32 {
	regexGetMaxBackrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := regexGetMaxBackrefFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var regexGetMaxLookbehindFunction *gi.Function
var regexGetMaxLookbehindFunction_Once sync.Once

func regexGetMaxLookbehindFunction_Set() {
	regexGetMaxLookbehindFunction_Once.Do(func() {
		regexGetMaxLookbehindFunction = gi.FunctionInvokerNew("GLib", "get_max_lookbehind")
	})
}

var getMaxLookbehindRegexInvoker *gi.Function

// GetMaxLookbehind is a representation of the C type g_regex_get_max_lookbehind.
func (recv *Regex) GetMaxLookbehind() int32 {
	regexGetMaxLookbehindFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := regexGetMaxLookbehindFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var regexGetPatternFunction *gi.Function
var regexGetPatternFunction_Once sync.Once

func regexGetPatternFunction_Set() {
	regexGetPatternFunction_Once.Do(func() {
		regexGetPatternFunction = gi.FunctionInvokerNew("GLib", "get_pattern")
	})
}

var getPatternRegexInvoker *gi.Function

// GetPattern is a representation of the C type g_regex_get_pattern.
func (recv *Regex) GetPattern() string {
	regexGetPatternFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := regexGetPatternFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var regexGetStringNumberFunction *gi.Function
var regexGetStringNumberFunction_Once sync.Once

func regexGetStringNumberFunction_Set() {
	regexGetStringNumberFunction_Once.Do(func() {
		regexGetStringNumberFunction = gi.FunctionInvokerNew("GLib", "get_string_number")
	})
}

var getStringNumberRegexInvoker *gi.Function

// GetStringNumber is a representation of the C type g_regex_get_string_number.
func (recv *Regex) GetStringNumber(name string) int32 {
	regexGetStringNumberFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	ret := regexGetStringNumberFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_regex_match' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_match_all' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_match_all_full' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_full' : parameter 'string' has no type

var regexRefFunction *gi.Function
var regexRefFunction_Once sync.Once

func regexRefFunction_Set() {
	regexRefFunction_Once.Do(func() {
		regexRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refRegexInvoker *gi.Function

// Ref is a representation of the C type g_regex_ref.
func (recv *Regex) Ref() *Regex {
	regexRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := regexRefFunction.Invoke(inArgs[:], nil)

	retGo := &Regex{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_regex_replace' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_replace_eval' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_replace_literal' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_split' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_full' : parameter 'string' has no type

var regexUnrefFunction *gi.Function
var regexUnrefFunction_Once sync.Once

func regexUnrefFunction_Set() {
	regexUnrefFunction_Once.Do(func() {
		regexUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefRegexInvoker *gi.Function

// Unref is a representation of the C type g_regex_unref.
func (recv *Regex) Unref() {
	regexUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	regexUnrefFunction.Invoke(inArgs[:], nil)

}

var sListStruct *gi.Struct
var sListStructOnce sync.Once

func sListStructSet() {
	sListStructOnce.Do(func() {
		sListStruct = gi.StructNew("GLib", "SList")
	})
}

type SList struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.SList'
}

var scannerStruct *gi.Struct
var scannerStructOnce sync.Once

func scannerStructSet() {
	scannerStructOnce.Do(func() {
		scannerStruct = gi.StructNew("GLib", "Scanner")
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

var scannerCurLineFunction *gi.Function
var scannerCurLineFunction_Once sync.Once

func scannerCurLineFunction_Set() {
	scannerCurLineFunction_Once.Do(func() {
		scannerCurLineFunction = gi.FunctionInvokerNew("GLib", "cur_line")
	})
}

var curLineScannerInvoker *gi.Function

// CurLine is a representation of the C type g_scanner_cur_line.
func (recv *Scanner) CurLine() uint32 {
	scannerCurLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := scannerCurLineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var scannerCurPositionFunction *gi.Function
var scannerCurPositionFunction_Once sync.Once

func scannerCurPositionFunction_Set() {
	scannerCurPositionFunction_Once.Do(func() {
		scannerCurPositionFunction = gi.FunctionInvokerNew("GLib", "cur_position")
	})
}

var curPositionScannerInvoker *gi.Function

// CurPosition is a representation of the C type g_scanner_cur_position.
func (recv *Scanner) CurPosition() uint32 {
	scannerCurPositionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := scannerCurPositionFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_scanner_cur_token' : return type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_cur_value' : return type 'TokenValue' not supported

var scannerDestroyFunction *gi.Function
var scannerDestroyFunction_Once sync.Once

func scannerDestroyFunction_Set() {
	scannerDestroyFunction_Once.Do(func() {
		scannerDestroyFunction = gi.FunctionInvokerNew("GLib", "destroy")
	})
}

var destroyScannerInvoker *gi.Function

// Destroy is a representation of the C type g_scanner_destroy.
func (recv *Scanner) Destroy() {
	scannerDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	scannerDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_scanner_eof' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_scanner_error' : parameter '...' has no type

// UNSUPPORTED : C value 'g_scanner_get_next_token' : return type 'TokenType' not supported

var scannerInputFileFunction *gi.Function
var scannerInputFileFunction_Once sync.Once

func scannerInputFileFunction_Set() {
	scannerInputFileFunction_Once.Do(func() {
		scannerInputFileFunction = gi.FunctionInvokerNew("GLib", "input_file")
	})
}

var inputFileScannerInvoker *gi.Function

// InputFile is a representation of the C type g_scanner_input_file.
func (recv *Scanner) InputFile(inputFd int32) {
	scannerInputFileFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(inputFd)

	scannerInputFileFunction.Invoke(inArgs[:], nil)

}

var scannerInputTextFunction *gi.Function
var scannerInputTextFunction_Once sync.Once

func scannerInputTextFunction_Set() {
	scannerInputTextFunction_Once.Do(func() {
		scannerInputTextFunction = gi.FunctionInvokerNew("GLib", "input_text")
	})
}

var inputTextScannerInvoker *gi.Function

// InputText is a representation of the C type g_scanner_input_text.
func (recv *Scanner) InputText(text string, textLen uint32) {
	scannerInputTextFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetUint32(textLen)

	scannerInputTextFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_scanner_lookup_symbol' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_scanner_peek_next_token' : return type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_scope_add_symbol' : parameter 'value' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_scanner_scope_foreach_symbol' : parameter 'func' of type 'HFunc' not supported

// UNSUPPORTED : C value 'g_scanner_scope_lookup_symbol' : return type 'gpointer' not supported

var scannerScopeRemoveSymbolFunction *gi.Function
var scannerScopeRemoveSymbolFunction_Once sync.Once

func scannerScopeRemoveSymbolFunction_Set() {
	scannerScopeRemoveSymbolFunction_Once.Do(func() {
		scannerScopeRemoveSymbolFunction = gi.FunctionInvokerNew("GLib", "scope_remove_symbol")
	})
}

var scopeRemoveSymbolScannerInvoker *gi.Function

// ScopeRemoveSymbol is a representation of the C type g_scanner_scope_remove_symbol.
func (recv *Scanner) ScopeRemoveSymbol(scopeId uint32, symbol string) {
	scannerScopeRemoveSymbolFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(scopeId)
	inArgs[2].SetString(symbol)

	scannerScopeRemoveSymbolFunction.Invoke(inArgs[:], nil)

}

var scannerSetScopeFunction *gi.Function
var scannerSetScopeFunction_Once sync.Once

func scannerSetScopeFunction_Set() {
	scannerSetScopeFunction_Once.Do(func() {
		scannerSetScopeFunction = gi.FunctionInvokerNew("GLib", "set_scope")
	})
}

var setScopeScannerInvoker *gi.Function

// SetScope is a representation of the C type g_scanner_set_scope.
func (recv *Scanner) SetScope(scopeId uint32) uint32 {
	scannerSetScopeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(scopeId)

	ret := scannerSetScopeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var scannerSyncFileOffsetFunction *gi.Function
var scannerSyncFileOffsetFunction_Once sync.Once

func scannerSyncFileOffsetFunction_Set() {
	scannerSyncFileOffsetFunction_Once.Do(func() {
		scannerSyncFileOffsetFunction = gi.FunctionInvokerNew("GLib", "sync_file_offset")
	})
}

var syncFileOffsetScannerInvoker *gi.Function

// SyncFileOffset is a representation of the C type g_scanner_sync_file_offset.
func (recv *Scanner) SyncFileOffset() {
	scannerSyncFileOffsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	scannerSyncFileOffsetFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_scanner_unexp_token' : parameter 'expected_token' of type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_warn' : parameter '...' has no type

var scannerConfigStruct *gi.Struct
var scannerConfigStructOnce sync.Once

func scannerConfigStructSet() {
	scannerConfigStructOnce.Do(func() {
		scannerConfigStruct = gi.StructNew("GLib", "ScannerConfig")
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

var sequenceStruct *gi.Struct
var sequenceStructOnce sync.Once

func sequenceStructSet() {
	sequenceStructOnce.Do(func() {
		sequenceStruct = gi.StructNew("GLib", "Sequence")
	})
}

type Sequence struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_sequence_append' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_foreach' : parameter 'func' of type 'Func' not supported

var sequenceFreeFunction *gi.Function
var sequenceFreeFunction_Once sync.Once

func sequenceFreeFunction_Set() {
	sequenceFreeFunction_Once.Do(func() {
		sequenceFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeSequenceInvoker *gi.Function

// Free is a representation of the C type g_sequence_free.
func (recv *Sequence) Free() {
	sequenceFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	sequenceFreeFunction.Invoke(inArgs[:], nil)

}

var sequenceGetBeginIterFunction *gi.Function
var sequenceGetBeginIterFunction_Once sync.Once

func sequenceGetBeginIterFunction_Set() {
	sequenceGetBeginIterFunction_Once.Do(func() {
		sequenceGetBeginIterFunction = gi.FunctionInvokerNew("GLib", "get_begin_iter")
	})
}

var getBeginIterSequenceInvoker *gi.Function

// GetBeginIter is a representation of the C type g_sequence_get_begin_iter.
func (recv *Sequence) GetBeginIter() *SequenceIter {
	sequenceGetBeginIterFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceGetBeginIterFunction.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sequenceGetEndIterFunction *gi.Function
var sequenceGetEndIterFunction_Once sync.Once

func sequenceGetEndIterFunction_Set() {
	sequenceGetEndIterFunction_Once.Do(func() {
		sequenceGetEndIterFunction = gi.FunctionInvokerNew("GLib", "get_end_iter")
	})
}

var getEndIterSequenceInvoker *gi.Function

// GetEndIter is a representation of the C type g_sequence_get_end_iter.
func (recv *Sequence) GetEndIter() *SequenceIter {
	sequenceGetEndIterFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceGetEndIterFunction.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sequenceGetIterAtPosFunction *gi.Function
var sequenceGetIterAtPosFunction_Once sync.Once

func sequenceGetIterAtPosFunction_Set() {
	sequenceGetIterAtPosFunction_Once.Do(func() {
		sequenceGetIterAtPosFunction = gi.FunctionInvokerNew("GLib", "get_iter_at_pos")
	})
}

var getIterAtPosSequenceInvoker *gi.Function

// GetIterAtPos is a representation of the C type g_sequence_get_iter_at_pos.
func (recv *Sequence) GetIterAtPos(pos int32) *SequenceIter {
	sequenceGetIterAtPosFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := sequenceGetIterAtPosFunction.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sequenceGetLengthFunction *gi.Function
var sequenceGetLengthFunction_Once sync.Once

func sequenceGetLengthFunction_Set() {
	sequenceGetLengthFunction_Once.Do(func() {
		sequenceGetLengthFunction = gi.FunctionInvokerNew("GLib", "get_length")
	})
}

var getLengthSequenceInvoker *gi.Function

// GetLength is a representation of the C type g_sequence_get_length.
func (recv *Sequence) GetLength() int32 {
	sequenceGetLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceGetLengthFunction.Invoke(inArgs[:], nil)

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

var sequenceIterStruct *gi.Struct
var sequenceIterStructOnce sync.Once

func sequenceIterStructSet() {
	sequenceIterStructOnce.Do(func() {
		sequenceIterStruct = gi.StructNew("GLib", "SequenceIter")
	})
}

type SequenceIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_sequence_iter_compare' : parameter 'b' of type 'SequenceIter' not supported

var sequenceIterGetPositionFunction *gi.Function
var sequenceIterGetPositionFunction_Once sync.Once

func sequenceIterGetPositionFunction_Set() {
	sequenceIterGetPositionFunction_Once.Do(func() {
		sequenceIterGetPositionFunction = gi.FunctionInvokerNew("GLib", "get_position")
	})
}

var getPositionSequenceIterInvoker *gi.Function

// GetPosition is a representation of the C type g_sequence_iter_get_position.
func (recv *SequenceIter) GetPosition() int32 {
	sequenceIterGetPositionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceIterGetPositionFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var sequenceIterGetSequenceFunction *gi.Function
var sequenceIterGetSequenceFunction_Once sync.Once

func sequenceIterGetSequenceFunction_Set() {
	sequenceIterGetSequenceFunction_Once.Do(func() {
		sequenceIterGetSequenceFunction = gi.FunctionInvokerNew("GLib", "get_sequence")
	})
}

var getSequenceSequenceIterInvoker *gi.Function

// GetSequence is a representation of the C type g_sequence_iter_get_sequence.
func (recv *SequenceIter) GetSequence() *Sequence {
	sequenceIterGetSequenceFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceIterGetSequenceFunction.Invoke(inArgs[:], nil)

	retGo := &Sequence{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_sequence_iter_is_begin' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_sequence_iter_is_end' : return type 'gboolean' not supported

var sequenceIterMoveFunction *gi.Function
var sequenceIterMoveFunction_Once sync.Once

func sequenceIterMoveFunction_Set() {
	sequenceIterMoveFunction_Once.Do(func() {
		sequenceIterMoveFunction = gi.FunctionInvokerNew("GLib", "move")
	})
}

var moveSequenceIterInvoker *gi.Function

// Move is a representation of the C type g_sequence_iter_move.
func (recv *SequenceIter) Move(delta int32) *SequenceIter {
	sequenceIterMoveFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(delta)

	ret := sequenceIterMoveFunction.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sequenceIterNextFunction *gi.Function
var sequenceIterNextFunction_Once sync.Once

func sequenceIterNextFunction_Set() {
	sequenceIterNextFunction_Once.Do(func() {
		sequenceIterNextFunction = gi.FunctionInvokerNew("GLib", "next")
	})
}

var nextSequenceIterInvoker *gi.Function

// Next is a representation of the C type g_sequence_iter_next.
func (recv *SequenceIter) Next() *SequenceIter {
	sequenceIterNextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceIterNextFunction.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sequenceIterPrevFunction *gi.Function
var sequenceIterPrevFunction_Once sync.Once

func sequenceIterPrevFunction_Set() {
	sequenceIterPrevFunction_Once.Do(func() {
		sequenceIterPrevFunction = gi.FunctionInvokerNew("GLib", "prev")
	})
}

var prevSequenceIterInvoker *gi.Function

// Prev is a representation of the C type g_sequence_iter_prev.
func (recv *SequenceIter) Prev() *SequenceIter {
	sequenceIterPrevFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sequenceIterPrevFunction.Invoke(inArgs[:], nil)

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo
}

var sourceStruct *gi.Struct
var sourceStructOnce sync.Once

func sourceStructSet() {
	sourceStructOnce.Do(func() {
		sourceStruct = gi.StructNew("GLib", "Source")
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

var sourceDestroyFunction *gi.Function
var sourceDestroyFunction_Once sync.Once

func sourceDestroyFunction_Set() {
	sourceDestroyFunction_Once.Do(func() {
		sourceDestroyFunction = gi.FunctionInvokerNew("GLib", "destroy")
	})
}

var destroySourceInvoker *gi.Function

// Destroy is a representation of the C type g_source_destroy.
func (recv *Source) Destroy() {
	sourceDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	sourceDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_source_get_can_recurse' : return type 'gboolean' not supported

var sourceGetContextFunction *gi.Function
var sourceGetContextFunction_Once sync.Once

func sourceGetContextFunction_Set() {
	sourceGetContextFunction_Once.Do(func() {
		sourceGetContextFunction = gi.FunctionInvokerNew("GLib", "get_context")
	})
}

var getContextSourceInvoker *gi.Function

// GetContext is a representation of the C type g_source_get_context.
func (recv *Source) GetContext() *MainContext {
	sourceGetContextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceGetContextFunction.Invoke(inArgs[:], nil)

	retGo := &MainContext{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_source_get_current_time' : parameter 'timeval' of type 'TimeVal' not supported

var sourceGetIdFunction *gi.Function
var sourceGetIdFunction_Once sync.Once

func sourceGetIdFunction_Set() {
	sourceGetIdFunction_Once.Do(func() {
		sourceGetIdFunction = gi.FunctionInvokerNew("GLib", "get_id")
	})
}

var getIdSourceInvoker *gi.Function

// GetId is a representation of the C type g_source_get_id.
func (recv *Source) GetId() uint32 {
	sourceGetIdFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceGetIdFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var sourceGetNameFunction *gi.Function
var sourceGetNameFunction_Once sync.Once

func sourceGetNameFunction_Set() {
	sourceGetNameFunction_Once.Do(func() {
		sourceGetNameFunction = gi.FunctionInvokerNew("GLib", "get_name")
	})
}

var getNameSourceInvoker *gi.Function

// GetName is a representation of the C type g_source_get_name.
func (recv *Source) GetName() string {
	sourceGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var sourceGetPriorityFunction *gi.Function
var sourceGetPriorityFunction_Once sync.Once

func sourceGetPriorityFunction_Set() {
	sourceGetPriorityFunction_Once.Do(func() {
		sourceGetPriorityFunction = gi.FunctionInvokerNew("GLib", "get_priority")
	})
}

var getPrioritySourceInvoker *gi.Function

// GetPriority is a representation of the C type g_source_get_priority.
func (recv *Source) GetPriority() int32 {
	sourceGetPriorityFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceGetPriorityFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var sourceGetReadyTimeFunction *gi.Function
var sourceGetReadyTimeFunction_Once sync.Once

func sourceGetReadyTimeFunction_Set() {
	sourceGetReadyTimeFunction_Once.Do(func() {
		sourceGetReadyTimeFunction = gi.FunctionInvokerNew("GLib", "get_ready_time")
	})
}

var getReadyTimeSourceInvoker *gi.Function

// GetReadyTime is a representation of the C type g_source_get_ready_time.
func (recv *Source) GetReadyTime() int64 {
	sourceGetReadyTimeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceGetReadyTimeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var sourceGetTimeFunction *gi.Function
var sourceGetTimeFunction_Once sync.Once

func sourceGetTimeFunction_Set() {
	sourceGetTimeFunction_Once.Do(func() {
		sourceGetTimeFunction = gi.FunctionInvokerNew("GLib", "get_time")
	})
}

var getTimeSourceInvoker *gi.Function

// GetTime is a representation of the C type g_source_get_time.
func (recv *Source) GetTime() int64 {
	sourceGetTimeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceGetTimeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_source_is_destroyed' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_source_modify_unix_fd' : parameter 'tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_source_query_unix_fd' : parameter 'tag' of type 'gpointer' not supported

var sourceRefFunction *gi.Function
var sourceRefFunction_Once sync.Once

func sourceRefFunction_Set() {
	sourceRefFunction_Once.Do(func() {
		sourceRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refSourceInvoker *gi.Function

// Ref is a representation of the C type g_source_ref.
func (recv *Source) Ref() *Source {
	sourceRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := sourceRefFunction.Invoke(inArgs[:], nil)

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

var sourceSetNameFunction *gi.Function
var sourceSetNameFunction_Once sync.Once

func sourceSetNameFunction_Set() {
	sourceSetNameFunction_Once.Do(func() {
		sourceSetNameFunction = gi.FunctionInvokerNew("GLib", "set_name")
	})
}

var setNameSourceInvoker *gi.Function

// SetName is a representation of the C type g_source_set_name.
func (recv *Source) SetName(name string) {
	sourceSetNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	sourceSetNameFunction.Invoke(inArgs[:], nil)

}

var sourceSetPriorityFunction *gi.Function
var sourceSetPriorityFunction_Once sync.Once

func sourceSetPriorityFunction_Set() {
	sourceSetPriorityFunction_Once.Do(func() {
		sourceSetPriorityFunction = gi.FunctionInvokerNew("GLib", "set_priority")
	})
}

var setPrioritySourceInvoker *gi.Function

// SetPriority is a representation of the C type g_source_set_priority.
func (recv *Source) SetPriority(priority int32) {
	sourceSetPriorityFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(priority)

	sourceSetPriorityFunction.Invoke(inArgs[:], nil)

}

var sourceSetReadyTimeFunction *gi.Function
var sourceSetReadyTimeFunction_Once sync.Once

func sourceSetReadyTimeFunction_Set() {
	sourceSetReadyTimeFunction_Once.Do(func() {
		sourceSetReadyTimeFunction = gi.FunctionInvokerNew("GLib", "set_ready_time")
	})
}

var setReadyTimeSourceInvoker *gi.Function

// SetReadyTime is a representation of the C type g_source_set_ready_time.
func (recv *Source) SetReadyTime(readyTime int64) {
	sourceSetReadyTimeFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(readyTime)

	sourceSetReadyTimeFunction.Invoke(inArgs[:], nil)

}

var sourceUnrefFunction *gi.Function
var sourceUnrefFunction_Once sync.Once

func sourceUnrefFunction_Set() {
	sourceUnrefFunction_Once.Do(func() {
		sourceUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefSourceInvoker *gi.Function

// Unref is a representation of the C type g_source_unref.
func (recv *Source) Unref() {
	sourceUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	sourceUnrefFunction.Invoke(inArgs[:], nil)

}

var sourceCallbackFuncsStruct *gi.Struct
var sourceCallbackFuncsStructOnce sync.Once

func sourceCallbackFuncsStructSet() {
	sourceCallbackFuncsStructOnce.Do(func() {
		sourceCallbackFuncsStruct = gi.StructNew("GLib", "SourceCallbackFuncs")
	})
}

type SourceCallbackFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'ref' : missing Type
	// UNSUPPORTED : C value 'unref' : missing Type
	// UNSUPPORTED : C value 'get' : missing Type
}

var sourceFuncsStruct *gi.Struct
var sourceFuncsStructOnce sync.Once

func sourceFuncsStructSet() {
	sourceFuncsStructOnce.Do(func() {
		sourceFuncsStruct = gi.StructNew("GLib", "SourceFuncs")
	})
}

type SourceFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'prepare' : missing Type
	// UNSUPPORTED : C value 'check' : missing Type
	// UNSUPPORTED : C value 'dispatch' : missing Type
	// UNSUPPORTED : C value 'finalize' : missing Type
}

var sourcePrivateStruct *gi.Struct
var sourcePrivateStructOnce sync.Once

func sourcePrivateStructSet() {
	sourcePrivateStructOnce.Do(func() {
		sourcePrivateStruct = gi.StructNew("GLib", "SourcePrivate")
	})
}

type SourcePrivate struct {
	native uintptr
}

var statBufStruct *gi.Struct
var statBufStructOnce sync.Once

func statBufStructSet() {
	statBufStructOnce.Do(func() {
		statBufStruct = gi.StructNew("GLib", "StatBuf")
	})
}

type StatBuf struct {
	native uintptr
}

var string_Struct *gi.Struct
var string_StructOnce sync.Once

func string_StructSet() {
	string_StructOnce.Do(func() {
		string_Struct = gi.StructNew("GLib", "String")
	})
}

type String struct {
	native       uintptr
	Str          string
	Len          uintptr
	AllocatedLen uintptr
}

var stringAppendFunction *gi.Function
var stringAppendFunction_Once sync.Once

func stringAppendFunction_Set() {
	stringAppendFunction_Once.Do(func() {
		stringAppendFunction = gi.FunctionInvokerNew("GLib", "append")
	})
}

var appendStringInvoker *gi.Function

// Append is a representation of the C type g_string_append.
func (recv *String) Append(val string) *String {
	stringAppendFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)

	ret := stringAppendFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringAppendCFunction *gi.Function
var stringAppendCFunction_Once sync.Once

func stringAppendCFunction_Set() {
	stringAppendCFunction_Once.Do(func() {
		stringAppendCFunction = gi.FunctionInvokerNew("GLib", "append_c")
	})
}

var appendCStringInvoker *gi.Function

// AppendC is a representation of the C type g_string_append_c.
func (recv *String) AppendC(c int8) *String {
	stringAppendCFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(c)

	ret := stringAppendCFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringAppendLenFunction *gi.Function
var stringAppendLenFunction_Once sync.Once

func stringAppendLenFunction_Set() {
	stringAppendLenFunction_Once.Do(func() {
		stringAppendLenFunction = gi.FunctionInvokerNew("GLib", "append_len")
	})
}

var appendLenStringInvoker *gi.Function

// AppendLen is a representation of the C type g_string_append_len.
func (recv *String) AppendLen(val string, len int32) *String {
	stringAppendLenFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)
	inArgs[2].SetInt32(len)

	ret := stringAppendLenFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_append_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_string_append_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_append_uri_escaped' : parameter 'allow_utf8' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_string_append_vprintf' : parameter 'args' of type 'va_list' not supported

var stringAsciiDownFunction *gi.Function
var stringAsciiDownFunction_Once sync.Once

func stringAsciiDownFunction_Set() {
	stringAsciiDownFunction_Once.Do(func() {
		stringAsciiDownFunction = gi.FunctionInvokerNew("GLib", "ascii_down")
	})
}

var asciiDownStringInvoker *gi.Function

// AsciiDown is a representation of the C type g_string_ascii_down.
func (recv *String) AsciiDown() *String {
	stringAsciiDownFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stringAsciiDownFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringAsciiUpFunction *gi.Function
var stringAsciiUpFunction_Once sync.Once

func stringAsciiUpFunction_Set() {
	stringAsciiUpFunction_Once.Do(func() {
		stringAsciiUpFunction = gi.FunctionInvokerNew("GLib", "ascii_up")
	})
}

var asciiUpStringInvoker *gi.Function

// AsciiUp is a representation of the C type g_string_ascii_up.
func (recv *String) AsciiUp() *String {
	stringAsciiUpFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stringAsciiUpFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringAssignFunction *gi.Function
var stringAssignFunction_Once sync.Once

func stringAssignFunction_Set() {
	stringAssignFunction_Once.Do(func() {
		stringAssignFunction = gi.FunctionInvokerNew("GLib", "assign")
	})
}

var assignStringInvoker *gi.Function

// Assign is a representation of the C type g_string_assign.
func (recv *String) Assign(rval string) *String {
	stringAssignFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(rval)

	ret := stringAssignFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringDownFunction *gi.Function
var stringDownFunction_Once sync.Once

func stringDownFunction_Set() {
	stringDownFunction_Once.Do(func() {
		stringDownFunction = gi.FunctionInvokerNew("GLib", "down")
	})
}

var downStringInvoker *gi.Function

// Down is a representation of the C type g_string_down.
func (recv *String) Down() *String {
	stringDownFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stringDownFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_equal' : parameter 'v2' of type 'String' not supported

var stringEraseFunction *gi.Function
var stringEraseFunction_Once sync.Once

func stringEraseFunction_Set() {
	stringEraseFunction_Once.Do(func() {
		stringEraseFunction = gi.FunctionInvokerNew("GLib", "erase")
	})
}

var eraseStringInvoker *gi.Function

// Erase is a representation of the C type g_string_erase.
func (recv *String) Erase(pos int32, len int32) *String {
	stringEraseFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetInt32(len)

	ret := stringEraseFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_free' : parameter 'free_segment' of type 'gboolean' not supported

var stringFreeToBytesFunction *gi.Function
var stringFreeToBytesFunction_Once sync.Once

func stringFreeToBytesFunction_Set() {
	stringFreeToBytesFunction_Once.Do(func() {
		stringFreeToBytesFunction = gi.FunctionInvokerNew("GLib", "free_to_bytes")
	})
}

var freeToBytesStringInvoker *gi.Function

// FreeToBytes is a representation of the C type g_string_free_to_bytes.
func (recv *String) FreeToBytes() *Bytes {
	stringFreeToBytesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stringFreeToBytesFunction.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

var stringHashFunction *gi.Function
var stringHashFunction_Once sync.Once

func stringHashFunction_Set() {
	stringHashFunction_Once.Do(func() {
		stringHashFunction = gi.FunctionInvokerNew("GLib", "hash")
	})
}

var hashStringInvoker *gi.Function

// Hash is a representation of the C type g_string_hash.
func (recv *String) Hash() uint32 {
	stringHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stringHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var stringInsertFunction *gi.Function
var stringInsertFunction_Once sync.Once

func stringInsertFunction_Set() {
	stringInsertFunction_Once.Do(func() {
		stringInsertFunction = gi.FunctionInvokerNew("GLib", "insert")
	})
}

var insertStringInvoker *gi.Function

// Insert is a representation of the C type g_string_insert.
func (recv *String) Insert(pos int32, val string) *String {
	stringInsertFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(val)

	ret := stringInsertFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringInsertCFunction *gi.Function
var stringInsertCFunction_Once sync.Once

func stringInsertCFunction_Set() {
	stringInsertCFunction_Once.Do(func() {
		stringInsertCFunction = gi.FunctionInvokerNew("GLib", "insert_c")
	})
}

var insertCStringInvoker *gi.Function

// InsertC is a representation of the C type g_string_insert_c.
func (recv *String) InsertC(pos int32, c int8) *String {
	stringInsertCFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetInt8(c)

	ret := stringInsertCFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringInsertLenFunction *gi.Function
var stringInsertLenFunction_Once sync.Once

func stringInsertLenFunction_Set() {
	stringInsertLenFunction_Once.Do(func() {
		stringInsertLenFunction = gi.FunctionInvokerNew("GLib", "insert_len")
	})
}

var insertLenStringInvoker *gi.Function

// InsertLen is a representation of the C type g_string_insert_len.
func (recv *String) InsertLen(pos int32, val string, len int32) *String {
	stringInsertLenFunction_Set()

	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(val)
	inArgs[3].SetInt32(len)

	ret := stringInsertLenFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_insert_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_overwrite' : parameter 'pos' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_string_overwrite_len' : parameter 'pos' of type 'gsize' not supported

var stringPrependFunction *gi.Function
var stringPrependFunction_Once sync.Once

func stringPrependFunction_Set() {
	stringPrependFunction_Once.Do(func() {
		stringPrependFunction = gi.FunctionInvokerNew("GLib", "prepend")
	})
}

var prependStringInvoker *gi.Function

// Prepend is a representation of the C type g_string_prepend.
func (recv *String) Prepend(val string) *String {
	stringPrependFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)

	ret := stringPrependFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringPrependCFunction *gi.Function
var stringPrependCFunction_Once sync.Once

func stringPrependCFunction_Set() {
	stringPrependCFunction_Once.Do(func() {
		stringPrependCFunction = gi.FunctionInvokerNew("GLib", "prepend_c")
	})
}

var prependCStringInvoker *gi.Function

// PrependC is a representation of the C type g_string_prepend_c.
func (recv *String) PrependC(c int8) *String {
	stringPrependCFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(c)

	ret := stringPrependCFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

var stringPrependLenFunction *gi.Function
var stringPrependLenFunction_Once sync.Once

func stringPrependLenFunction_Set() {
	stringPrependLenFunction_Once.Do(func() {
		stringPrependLenFunction = gi.FunctionInvokerNew("GLib", "prepend_len")
	})
}

var prependLenStringInvoker *gi.Function

// PrependLen is a representation of the C type g_string_prepend_len.
func (recv *String) PrependLen(val string, len int32) *String {
	stringPrependLenFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)
	inArgs[2].SetInt32(len)

	ret := stringPrependLenFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_prepend_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_string_set_size' : parameter 'len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_string_truncate' : parameter 'len' of type 'gsize' not supported

var stringUpFunction *gi.Function
var stringUpFunction_Once sync.Once

func stringUpFunction_Set() {
	stringUpFunction_Once.Do(func() {
		stringUpFunction = gi.FunctionInvokerNew("GLib", "up")
	})
}

var upStringInvoker *gi.Function

// Up is a representation of the C type g_string_up.
func (recv *String) Up() *String {
	stringUpFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stringUpFunction.Invoke(inArgs[:], nil)

	retGo := &String{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_string_vprintf' : parameter 'args' of type 'va_list' not supported

var stringChunkStruct *gi.Struct
var stringChunkStructOnce sync.Once

func stringChunkStructSet() {
	stringChunkStructOnce.Do(func() {
		stringChunkStruct = gi.StructNew("GLib", "StringChunk")
	})
}

type StringChunk struct {
	native uintptr
}

var stringChunkClearFunction *gi.Function
var stringChunkClearFunction_Once sync.Once

func stringChunkClearFunction_Set() {
	stringChunkClearFunction_Once.Do(func() {
		stringChunkClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearStringChunkInvoker *gi.Function

// Clear is a representation of the C type g_string_chunk_clear.
func (recv *StringChunk) Clear() {
	stringChunkClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stringChunkClearFunction.Invoke(inArgs[:], nil)

}

var stringChunkFreeFunction *gi.Function
var stringChunkFreeFunction_Once sync.Once

func stringChunkFreeFunction_Set() {
	stringChunkFreeFunction_Once.Do(func() {
		stringChunkFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeStringChunkInvoker *gi.Function

// Free is a representation of the C type g_string_chunk_free.
func (recv *StringChunk) Free() {
	stringChunkFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stringChunkFreeFunction.Invoke(inArgs[:], nil)

}

var stringChunkInsertFunction *gi.Function
var stringChunkInsertFunction_Once sync.Once

func stringChunkInsertFunction_Set() {
	stringChunkInsertFunction_Once.Do(func() {
		stringChunkInsertFunction = gi.FunctionInvokerNew("GLib", "insert")
	})
}

var insertStringChunkInvoker *gi.Function

// Insert is a representation of the C type g_string_chunk_insert.
func (recv *StringChunk) Insert(string_ string) string {
	stringChunkInsertFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)

	ret := stringChunkInsertFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var stringChunkInsertConstFunction *gi.Function
var stringChunkInsertConstFunction_Once sync.Once

func stringChunkInsertConstFunction_Set() {
	stringChunkInsertConstFunction_Once.Do(func() {
		stringChunkInsertConstFunction = gi.FunctionInvokerNew("GLib", "insert_const")
	})
}

var insertConstStringChunkInvoker *gi.Function

// InsertConst is a representation of the C type g_string_chunk_insert_const.
func (recv *StringChunk) InsertConst(string_ string) string {
	stringChunkInsertConstFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)

	ret := stringChunkInsertConstFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var stringChunkInsertLenFunction *gi.Function
var stringChunkInsertLenFunction_Once sync.Once

func stringChunkInsertLenFunction_Set() {
	stringChunkInsertLenFunction_Once.Do(func() {
		stringChunkInsertLenFunction = gi.FunctionInvokerNew("GLib", "insert_len")
	})
}

var insertLenStringChunkInvoker *gi.Function

// InsertLen is a representation of the C type g_string_chunk_insert_len.
func (recv *StringChunk) InsertLen(string_ string, len int32) string {
	stringChunkInsertLenFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)
	inArgs[2].SetInt32(len)

	ret := stringChunkInsertLenFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var testCaseStruct *gi.Struct
var testCaseStructOnce sync.Once

func testCaseStructSet() {
	testCaseStructOnce.Do(func() {
		testCaseStruct = gi.StructNew("GLib", "TestCase")
	})
}

type TestCase struct {
	native uintptr
}

var testConfigStruct *gi.Struct
var testConfigStructOnce sync.Once

func testConfigStructSet() {
	testConfigStructOnce.Do(func() {
		testConfigStruct = gi.StructNew("GLib", "TestConfig")
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

var testLogBufferStruct *gi.Struct
var testLogBufferStructOnce sync.Once

func testLogBufferStructSet() {
	testLogBufferStructOnce.Do(func() {
		testLogBufferStruct = gi.StructNew("GLib", "TestLogBuffer")
	})
}

type TestLogBuffer struct {
	native uintptr
}

var testLogBufferFreeFunction *gi.Function
var testLogBufferFreeFunction_Once sync.Once

func testLogBufferFreeFunction_Set() {
	testLogBufferFreeFunction_Once.Do(func() {
		testLogBufferFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeTestLogBufferInvoker *gi.Function

// Free is a representation of the C type g_test_log_buffer_free.
func (recv *TestLogBuffer) Free() {
	testLogBufferFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	testLogBufferFreeFunction.Invoke(inArgs[:], nil)

}

var testLogBufferPopFunction *gi.Function
var testLogBufferPopFunction_Once sync.Once

func testLogBufferPopFunction_Set() {
	testLogBufferPopFunction_Once.Do(func() {
		testLogBufferPopFunction = gi.FunctionInvokerNew("GLib", "pop")
	})
}

var popTestLogBufferInvoker *gi.Function

// Pop is a representation of the C type g_test_log_buffer_pop.
func (recv *TestLogBuffer) Pop() *TestLogMsg {
	testLogBufferPopFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := testLogBufferPopFunction.Invoke(inArgs[:], nil)

	retGo := &TestLogMsg{native: ret.Pointer()}

	return retGo
}

var testLogBufferPushFunction *gi.Function
var testLogBufferPushFunction_Once sync.Once

func testLogBufferPushFunction_Set() {
	testLogBufferPushFunction_Once.Do(func() {
		testLogBufferPushFunction = gi.FunctionInvokerNew("GLib", "push")
	})
}

var pushTestLogBufferInvoker *gi.Function

// Push is a representation of the C type g_test_log_buffer_push.
func (recv *TestLogBuffer) Push(nBytes uint32, bytes uint8) {
	testLogBufferPushFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nBytes)
	inArgs[2].SetUint8(bytes)

	testLogBufferPushFunction.Invoke(inArgs[:], nil)

}

var testLogMsgStruct *gi.Struct
var testLogMsgStructOnce sync.Once

func testLogMsgStructSet() {
	testLogMsgStructOnce.Do(func() {
		testLogMsgStruct = gi.StructNew("GLib", "TestLogMsg")
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

var testLogMsgFreeFunction *gi.Function
var testLogMsgFreeFunction_Once sync.Once

func testLogMsgFreeFunction_Set() {
	testLogMsgFreeFunction_Once.Do(func() {
		testLogMsgFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeTestLogMsgInvoker *gi.Function

// Free is a representation of the C type g_test_log_msg_free.
func (recv *TestLogMsg) Free() {
	testLogMsgFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	testLogMsgFreeFunction.Invoke(inArgs[:], nil)

}

var testSuiteStruct *gi.Struct
var testSuiteStructOnce sync.Once

func testSuiteStructSet() {
	testSuiteStructOnce.Do(func() {
		testSuiteStruct = gi.StructNew("GLib", "TestSuite")
	})
}

type TestSuite struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_test_suite_add' : parameter 'test_case' of type 'TestCase' not supported

// UNSUPPORTED : C value 'g_test_suite_add_suite' : parameter 'nestedsuite' of type 'TestSuite' not supported

var threadStruct *gi.Struct
var threadStructOnce sync.Once

func threadStructSet() {
	threadStructOnce.Do(func() {
		threadStruct = gi.StructNew("GLib", "Thread")
	})
}

type Thread struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_thread_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_try_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_join' : return type 'gpointer' not supported

var threadRefFunction *gi.Function
var threadRefFunction_Once sync.Once

func threadRefFunction_Set() {
	threadRefFunction_Once.Do(func() {
		threadRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refThreadInvoker *gi.Function

// Ref is a representation of the C type g_thread_ref.
func (recv *Thread) Ref() *Thread {
	threadRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := threadRefFunction.Invoke(inArgs[:], nil)

	retGo := &Thread{native: ret.Pointer()}

	return retGo
}

var threadUnrefFunction *gi.Function
var threadUnrefFunction_Once sync.Once

func threadUnrefFunction_Set() {
	threadUnrefFunction_Once.Do(func() {
		threadUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefThreadInvoker *gi.Function

// Unref is a representation of the C type g_thread_unref.
func (recv *Thread) Unref() {
	threadUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	threadUnrefFunction.Invoke(inArgs[:], nil)

}

var threadPoolStruct *gi.Struct
var threadPoolStructOnce sync.Once

func threadPoolStructSet() {
	threadPoolStructOnce.Do(func() {
		threadPoolStruct = gi.StructNew("GLib", "ThreadPool")
	})
}

type ThreadPool struct {
	native uintptr
	// UNSUPPORTED : C value 'func' : no Go type for 'Func'
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'exclusive' : no Go type for 'gboolean'
}

// UNSUPPORTED : C value 'g_thread_pool_free' : parameter 'immediate' of type 'gboolean' not supported

var threadPoolGetMaxThreadsFunction *gi.Function
var threadPoolGetMaxThreadsFunction_Once sync.Once

func threadPoolGetMaxThreadsFunction_Set() {
	threadPoolGetMaxThreadsFunction_Once.Do(func() {
		threadPoolGetMaxThreadsFunction = gi.FunctionInvokerNew("GLib", "get_max_threads")
	})
}

var getMaxThreadsThreadPoolInvoker *gi.Function

// GetMaxThreads is a representation of the C type g_thread_pool_get_max_threads.
func (recv *ThreadPool) GetMaxThreads() int32 {
	threadPoolGetMaxThreadsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := threadPoolGetMaxThreadsFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var threadPoolGetNumThreadsFunction *gi.Function
var threadPoolGetNumThreadsFunction_Once sync.Once

func threadPoolGetNumThreadsFunction_Set() {
	threadPoolGetNumThreadsFunction_Once.Do(func() {
		threadPoolGetNumThreadsFunction = gi.FunctionInvokerNew("GLib", "get_num_threads")
	})
}

var getNumThreadsThreadPoolInvoker *gi.Function

// GetNumThreads is a representation of the C type g_thread_pool_get_num_threads.
func (recv *ThreadPool) GetNumThreads() uint32 {
	threadPoolGetNumThreadsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := threadPoolGetNumThreadsFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_thread_pool_move_to_front' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_thread_pool_push' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_thread_pool_set_max_threads' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_thread_pool_set_sort_function' : parameter 'func' of type 'CompareDataFunc' not supported

var threadPoolUnprocessedFunction *gi.Function
var threadPoolUnprocessedFunction_Once sync.Once

func threadPoolUnprocessedFunction_Set() {
	threadPoolUnprocessedFunction_Once.Do(func() {
		threadPoolUnprocessedFunction = gi.FunctionInvokerNew("GLib", "unprocessed")
	})
}

var unprocessedThreadPoolInvoker *gi.Function

// Unprocessed is a representation of the C type g_thread_pool_unprocessed.
func (recv *ThreadPool) Unprocessed() uint32 {
	threadPoolUnprocessedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := threadPoolUnprocessedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var timeValStruct *gi.Struct
var timeValStructOnce sync.Once

func timeValStructSet() {
	timeValStructOnce.Do(func() {
		timeValStruct = gi.StructNew("GLib", "TimeVal")
	})
}

type TimeVal struct {
	native uintptr
	TvSec  int64
	TvUsec int64
}

var timeValAddFunction *gi.Function
var timeValAddFunction_Once sync.Once

func timeValAddFunction_Set() {
	timeValAddFunction_Once.Do(func() {
		timeValAddFunction = gi.FunctionInvokerNew("GLib", "add")
	})
}

var addTimeValInvoker *gi.Function

// Add is a representation of the C type g_time_val_add.
func (recv *TimeVal) Add(microseconds int64) {
	timeValAddFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(microseconds)

	timeValAddFunction.Invoke(inArgs[:], nil)

}

var timeValToIso8601Function *gi.Function
var timeValToIso8601Function_Once sync.Once

func timeValToIso8601Function_Set() {
	timeValToIso8601Function_Once.Do(func() {
		timeValToIso8601Function = gi.FunctionInvokerNew("GLib", "to_iso8601")
	})
}

var toIso8601TimeValInvoker *gi.Function

// ToIso8601 is a representation of the C type g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() string {
	timeValToIso8601Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := timeValToIso8601Function.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var timeZoneStruct *gi.Struct
var timeZoneStructOnce sync.Once

func timeZoneStructSet() {
	timeZoneStructOnce.Do(func() {
		timeZoneStruct = gi.StructNew("GLib", "TimeZone")
	})
}

type TimeZone struct {
	native uintptr
}

var timeZoneNewFunction *gi.Function
var timeZoneNewFunction_Once sync.Once

func timeZoneNewFunction_Set() {
	timeZoneNewFunction_Once.Do(func() {
		timeZoneNewFunction = gi.FunctionInvokerNew("GLib", "new")
	})
}

var newTimeZoneInvoker *gi.Function

// TimeZoneNew is a representation of the C type g_time_zone_new.
func TimeZoneNew(identifier string) *TimeZone {
	timeZoneNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(identifier)

	ret := timeZoneNewFunction.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var timeZoneNewLocalFunction *gi.Function
var timeZoneNewLocalFunction_Once sync.Once

func timeZoneNewLocalFunction_Set() {
	timeZoneNewLocalFunction_Once.Do(func() {
		timeZoneNewLocalFunction = gi.FunctionInvokerNew("GLib", "new_local")
	})
}

var newLocalTimeZoneInvoker *gi.Function

// TimeZoneNewLocal is a representation of the C type g_time_zone_new_local.
func TimeZoneNewLocal() *TimeZone {
	timeZoneNewLocalFunction_Set()

	ret := timeZoneNewLocalFunction.Invoke(nil, nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var timeZoneNewOffsetFunction *gi.Function
var timeZoneNewOffsetFunction_Once sync.Once

func timeZoneNewOffsetFunction_Set() {
	timeZoneNewOffsetFunction_Once.Do(func() {
		timeZoneNewOffsetFunction = gi.FunctionInvokerNew("GLib", "new_offset")
	})
}

var newOffsetTimeZoneInvoker *gi.Function

// TimeZoneNewOffset is a representation of the C type g_time_zone_new_offset.
func TimeZoneNewOffset(seconds int32) *TimeZone {
	timeZoneNewOffsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(seconds)

	ret := timeZoneNewOffsetFunction.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var timeZoneNewUtcFunction *gi.Function
var timeZoneNewUtcFunction_Once sync.Once

func timeZoneNewUtcFunction_Set() {
	timeZoneNewUtcFunction_Once.Do(func() {
		timeZoneNewUtcFunction = gi.FunctionInvokerNew("GLib", "new_utc")
	})
}

var newUtcTimeZoneInvoker *gi.Function

// TimeZoneNewUtc is a representation of the C type g_time_zone_new_utc.
func TimeZoneNewUtc() *TimeZone {
	timeZoneNewUtcFunction_Set()

	ret := timeZoneNewUtcFunction.Invoke(nil, nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_time_zone_adjust_time' : parameter 'type' of type 'TimeType' not supported

// UNSUPPORTED : C value 'g_time_zone_find_interval' : parameter 'type' of type 'TimeType' not supported

var timeZoneGetAbbreviationFunction *gi.Function
var timeZoneGetAbbreviationFunction_Once sync.Once

func timeZoneGetAbbreviationFunction_Set() {
	timeZoneGetAbbreviationFunction_Once.Do(func() {
		timeZoneGetAbbreviationFunction = gi.FunctionInvokerNew("GLib", "get_abbreviation")
	})
}

var getAbbreviationTimeZoneInvoker *gi.Function

// GetAbbreviation is a representation of the C type g_time_zone_get_abbreviation.
func (recv *TimeZone) GetAbbreviation(interval int32) string {
	timeZoneGetAbbreviationFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	ret := timeZoneGetAbbreviationFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var timeZoneGetIdentifierFunction *gi.Function
var timeZoneGetIdentifierFunction_Once sync.Once

func timeZoneGetIdentifierFunction_Set() {
	timeZoneGetIdentifierFunction_Once.Do(func() {
		timeZoneGetIdentifierFunction = gi.FunctionInvokerNew("GLib", "get_identifier")
	})
}

var getIdentifierTimeZoneInvoker *gi.Function

// GetIdentifier is a representation of the C type g_time_zone_get_identifier.
func (recv *TimeZone) GetIdentifier() string {
	timeZoneGetIdentifierFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := timeZoneGetIdentifierFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var timeZoneGetOffsetFunction *gi.Function
var timeZoneGetOffsetFunction_Once sync.Once

func timeZoneGetOffsetFunction_Set() {
	timeZoneGetOffsetFunction_Once.Do(func() {
		timeZoneGetOffsetFunction = gi.FunctionInvokerNew("GLib", "get_offset")
	})
}

var getOffsetTimeZoneInvoker *gi.Function

// GetOffset is a representation of the C type g_time_zone_get_offset.
func (recv *TimeZone) GetOffset(interval int32) int32 {
	timeZoneGetOffsetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	ret := timeZoneGetOffsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_time_zone_is_dst' : return type 'gboolean' not supported

var timeZoneRefFunction *gi.Function
var timeZoneRefFunction_Once sync.Once

func timeZoneRefFunction_Set() {
	timeZoneRefFunction_Once.Do(func() {
		timeZoneRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refTimeZoneInvoker *gi.Function

// Ref is a representation of the C type g_time_zone_ref.
func (recv *TimeZone) Ref() *TimeZone {
	timeZoneRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := timeZoneRefFunction.Invoke(inArgs[:], nil)

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo
}

var timeZoneUnrefFunction *gi.Function
var timeZoneUnrefFunction_Once sync.Once

func timeZoneUnrefFunction_Set() {
	timeZoneUnrefFunction_Once.Do(func() {
		timeZoneUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefTimeZoneInvoker *gi.Function

// Unref is a representation of the C type g_time_zone_unref.
func (recv *TimeZone) Unref() {
	timeZoneUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	timeZoneUnrefFunction.Invoke(inArgs[:], nil)

}

var timerStruct *gi.Struct
var timerStructOnce sync.Once

func timerStructSet() {
	timerStructOnce.Do(func() {
		timerStruct = gi.StructNew("GLib", "Timer")
	})
}

type Timer struct {
	native uintptr
}

var timerContinueFunction *gi.Function
var timerContinueFunction_Once sync.Once

func timerContinueFunction_Set() {
	timerContinueFunction_Once.Do(func() {
		timerContinueFunction = gi.FunctionInvokerNew("GLib", "continue")
	})
}

var continueTimerInvoker *gi.Function

// Continue is a representation of the C type g_timer_continue.
func (recv *Timer) Continue() {
	timerContinueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	timerContinueFunction.Invoke(inArgs[:], nil)

}

var timerDestroyFunction *gi.Function
var timerDestroyFunction_Once sync.Once

func timerDestroyFunction_Set() {
	timerDestroyFunction_Once.Do(func() {
		timerDestroyFunction = gi.FunctionInvokerNew("GLib", "destroy")
	})
}

var destroyTimerInvoker *gi.Function

// Destroy is a representation of the C type g_timer_destroy.
func (recv *Timer) Destroy() {
	timerDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	timerDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_timer_elapsed' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_timer_is_active' : return type 'gboolean' not supported

var timerResetFunction *gi.Function
var timerResetFunction_Once sync.Once

func timerResetFunction_Set() {
	timerResetFunction_Once.Do(func() {
		timerResetFunction = gi.FunctionInvokerNew("GLib", "reset")
	})
}

var resetTimerInvoker *gi.Function

// Reset is a representation of the C type g_timer_reset.
func (recv *Timer) Reset() {
	timerResetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	timerResetFunction.Invoke(inArgs[:], nil)

}

var timerStartFunction *gi.Function
var timerStartFunction_Once sync.Once

func timerStartFunction_Set() {
	timerStartFunction_Once.Do(func() {
		timerStartFunction = gi.FunctionInvokerNew("GLib", "start")
	})
}

var startTimerInvoker *gi.Function

// Start is a representation of the C type g_timer_start.
func (recv *Timer) Start() {
	timerStartFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	timerStartFunction.Invoke(inArgs[:], nil)

}

var timerStopFunction *gi.Function
var timerStopFunction_Once sync.Once

func timerStopFunction_Set() {
	timerStopFunction_Once.Do(func() {
		timerStopFunction = gi.FunctionInvokerNew("GLib", "stop")
	})
}

var stopTimerInvoker *gi.Function

// Stop is a representation of the C type g_timer_stop.
func (recv *Timer) Stop() {
	timerStopFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	timerStopFunction.Invoke(inArgs[:], nil)

}

var trashStackStruct *gi.Struct
var trashStackStructOnce sync.Once

func trashStackStructSet() {
	trashStackStructOnce.Do(func() {
		trashStackStruct = gi.StructNew("GLib", "TrashStack")
	})
}

type TrashStack struct {
	native uintptr
	Next   *TrashStack
}

var treeStruct *gi.Struct
var treeStructOnce sync.Once

func treeStructSet() {
	treeStructOnce.Do(func() {
		treeStruct = gi.StructNew("GLib", "Tree")
	})
}

type Tree struct {
	native uintptr
}

var treeDestroyFunction *gi.Function
var treeDestroyFunction_Once sync.Once

func treeDestroyFunction_Set() {
	treeDestroyFunction_Once.Do(func() {
		treeDestroyFunction = gi.FunctionInvokerNew("GLib", "destroy")
	})
}

var destroyTreeInvoker *gi.Function

// Destroy is a representation of the C type g_tree_destroy.
func (recv *Tree) Destroy() {
	treeDestroyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treeDestroyFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_tree_foreach' : parameter 'func' of type 'TraverseFunc' not supported

var treeHeightFunction *gi.Function
var treeHeightFunction_Once sync.Once

func treeHeightFunction_Set() {
	treeHeightFunction_Once.Do(func() {
		treeHeightFunction = gi.FunctionInvokerNew("GLib", "height")
	})
}

var heightTreeInvoker *gi.Function

// Height is a representation of the C type g_tree_height.
func (recv *Tree) Height() int32 {
	treeHeightFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treeHeightFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_tree_insert' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_lookup' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_lookup_extended' : parameter 'lookup_key' of type 'gpointer' not supported

var treeNnodesFunction *gi.Function
var treeNnodesFunction_Once sync.Once

func treeNnodesFunction_Set() {
	treeNnodesFunction_Once.Do(func() {
		treeNnodesFunction = gi.FunctionInvokerNew("GLib", "nnodes")
	})
}

var nnodesTreeInvoker *gi.Function

// Nnodes is a representation of the C type g_tree_nnodes.
func (recv *Tree) Nnodes() int32 {
	treeNnodesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treeNnodesFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var treeRefFunction *gi.Function
var treeRefFunction_Once sync.Once

func treeRefFunction_Set() {
	treeRefFunction_Once.Do(func() {
		treeRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refTreeInvoker *gi.Function

// Ref is a representation of the C type g_tree_ref.
func (recv *Tree) Ref() *Tree {
	treeRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treeRefFunction.Invoke(inArgs[:], nil)

	retGo := &Tree{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_tree_remove' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_replace' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_search' : parameter 'search_func' of type 'CompareFunc' not supported

// UNSUPPORTED : C value 'g_tree_steal' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_traverse' : parameter 'traverse_func' of type 'TraverseFunc' not supported

var treeUnrefFunction *gi.Function
var treeUnrefFunction_Once sync.Once

func treeUnrefFunction_Set() {
	treeUnrefFunction_Once.Do(func() {
		treeUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefTreeInvoker *gi.Function

// Unref is a representation of the C type g_tree_unref.
func (recv *Tree) Unref() {
	treeUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treeUnrefFunction.Invoke(inArgs[:], nil)

}

var variantStruct *gi.Struct
var variantStructOnce sync.Once

func variantStructSet() {
	variantStructOnce.Do(func() {
		variantStruct = gi.StructNew("GLib", "Variant")
	})
}

type Variant struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_new' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_array' : parameter 'child_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_boolean' : parameter 'value' of type 'gboolean' not supported

var variantNewByteFunction *gi.Function
var variantNewByteFunction_Once sync.Once

func variantNewByteFunction_Set() {
	variantNewByteFunction_Once.Do(func() {
		variantNewByteFunction = gi.FunctionInvokerNew("GLib", "new_byte")
	})
}

var newByteVariantInvoker *gi.Function

// VariantNewByte is a representation of the C type g_variant_new_byte.
func VariantNewByte(value uint8) *Variant {
	variantNewByteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint8(value)

	ret := variantNewByteFunction.Invoke(inArgs[:], nil)

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

var variantNewHandleFunction *gi.Function
var variantNewHandleFunction_Once sync.Once

func variantNewHandleFunction_Set() {
	variantNewHandleFunction_Once.Do(func() {
		variantNewHandleFunction = gi.FunctionInvokerNew("GLib", "new_handle")
	})
}

var newHandleVariantInvoker *gi.Function

// VariantNewHandle is a representation of the C type g_variant_new_handle.
func VariantNewHandle(value int32) *Variant {
	variantNewHandleFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	ret := variantNewHandleFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantNewInt16Function *gi.Function
var variantNewInt16Function_Once sync.Once

func variantNewInt16Function_Set() {
	variantNewInt16Function_Once.Do(func() {
		variantNewInt16Function = gi.FunctionInvokerNew("GLib", "new_int16")
	})
}

var newInt16VariantInvoker *gi.Function

// VariantNewInt16 is a representation of the C type g_variant_new_int16.
func VariantNewInt16(value int16) *Variant {
	variantNewInt16Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt16(value)

	ret := variantNewInt16Function.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantNewInt32Function *gi.Function
var variantNewInt32Function_Once sync.Once

func variantNewInt32Function_Set() {
	variantNewInt32Function_Once.Do(func() {
		variantNewInt32Function = gi.FunctionInvokerNew("GLib", "new_int32")
	})
}

var newInt32VariantInvoker *gi.Function

// VariantNewInt32 is a representation of the C type g_variant_new_int32.
func VariantNewInt32(value int32) *Variant {
	variantNewInt32Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	ret := variantNewInt32Function.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantNewInt64Function *gi.Function
var variantNewInt64Function_Once sync.Once

func variantNewInt64Function_Set() {
	variantNewInt64Function_Once.Do(func() {
		variantNewInt64Function = gi.FunctionInvokerNew("GLib", "new_int64")
	})
}

var newInt64VariantInvoker *gi.Function

// VariantNewInt64 is a representation of the C type g_variant_new_int64.
func VariantNewInt64(value int64) *Variant {
	variantNewInt64Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(value)

	ret := variantNewInt64Function.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_maybe' : parameter 'child_type' of type 'VariantType' not supported

var variantNewObjectPathFunction *gi.Function
var variantNewObjectPathFunction_Once sync.Once

func variantNewObjectPathFunction_Set() {
	variantNewObjectPathFunction_Once.Do(func() {
		variantNewObjectPathFunction = gi.FunctionInvokerNew("GLib", "new_object_path")
	})
}

var newObjectPathVariantInvoker *gi.Function

// VariantNewObjectPath is a representation of the C type g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) *Variant {
	variantNewObjectPathFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(objectPath)

	ret := variantNewObjectPathFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_objv' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_printf' : parameter '...' has no type

var variantNewSignatureFunction *gi.Function
var variantNewSignatureFunction_Once sync.Once

func variantNewSignatureFunction_Set() {
	variantNewSignatureFunction_Once.Do(func() {
		variantNewSignatureFunction = gi.FunctionInvokerNew("GLib", "new_signature")
	})
}

var newSignatureVariantInvoker *gi.Function

// VariantNewSignature is a representation of the C type g_variant_new_signature.
func VariantNewSignature(signature string) *Variant {
	variantNewSignatureFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(signature)

	ret := variantNewSignatureFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantNewStringFunction *gi.Function
var variantNewStringFunction_Once sync.Once

func variantNewStringFunction_Set() {
	variantNewStringFunction_Once.Do(func() {
		variantNewStringFunction = gi.FunctionInvokerNew("GLib", "new_string")
	})
}

var newStringVariantInvoker *gi.Function

// VariantNewString is a representation of the C type g_variant_new_string.
func VariantNewString(string_ string) *Variant {
	variantNewStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := variantNewStringFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_strv' : parameter 'strv' has no type

var variantNewTakeStringFunction *gi.Function
var variantNewTakeStringFunction_Once sync.Once

func variantNewTakeStringFunction_Set() {
	variantNewTakeStringFunction_Once.Do(func() {
		variantNewTakeStringFunction = gi.FunctionInvokerNew("GLib", "new_take_string")
	})
}

var newTakeStringVariantInvoker *gi.Function

// VariantNewTakeString is a representation of the C type g_variant_new_take_string.
func VariantNewTakeString(string_ string) *Variant {
	variantNewTakeStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := variantNewTakeStringFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_tuple' : parameter 'children' has no type

var variantNewUint16Function *gi.Function
var variantNewUint16Function_Once sync.Once

func variantNewUint16Function_Set() {
	variantNewUint16Function_Once.Do(func() {
		variantNewUint16Function = gi.FunctionInvokerNew("GLib", "new_uint16")
	})
}

var newUint16VariantInvoker *gi.Function

// VariantNewUint16 is a representation of the C type g_variant_new_uint16.
func VariantNewUint16(value uint16) *Variant {
	variantNewUint16Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(value)

	ret := variantNewUint16Function.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantNewUint32Function *gi.Function
var variantNewUint32Function_Once sync.Once

func variantNewUint32Function_Set() {
	variantNewUint32Function_Once.Do(func() {
		variantNewUint32Function = gi.FunctionInvokerNew("GLib", "new_uint32")
	})
}

var newUint32VariantInvoker *gi.Function

// VariantNewUint32 is a representation of the C type g_variant_new_uint32.
func VariantNewUint32(value uint32) *Variant {
	variantNewUint32Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(value)

	ret := variantNewUint32Function.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantNewUint64Function *gi.Function
var variantNewUint64Function_Once sync.Once

func variantNewUint64Function_Set() {
	variantNewUint64Function_Once.Do(func() {
		variantNewUint64Function = gi.FunctionInvokerNew("GLib", "new_uint64")
	})
}

var newUint64VariantInvoker *gi.Function

// VariantNewUint64 is a representation of the C type g_variant_new_uint64.
func VariantNewUint64(value uint64) *Variant {
	variantNewUint64Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(value)

	ret := variantNewUint64Function.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_new_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_variant' : parameter 'value' of type 'Variant' not supported

var variantByteswapFunction *gi.Function
var variantByteswapFunction_Once sync.Once

func variantByteswapFunction_Set() {
	variantByteswapFunction_Once.Do(func() {
		variantByteswapFunction = gi.FunctionInvokerNew("GLib", "byteswap")
	})
}

var byteswapVariantInvoker *gi.Function

// Byteswap is a representation of the C type g_variant_byteswap.
func (recv *Variant) Byteswap() *Variant {
	variantByteswapFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantByteswapFunction.Invoke(inArgs[:], nil)

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

var variantGetByteFunction *gi.Function
var variantGetByteFunction_Once sync.Once

func variantGetByteFunction_Set() {
	variantGetByteFunction_Once.Do(func() {
		variantGetByteFunction = gi.FunctionInvokerNew("GLib", "get_byte")
	})
}

var getByteVariantInvoker *gi.Function

// GetByte is a representation of the C type g_variant_get_byte.
func (recv *Variant) GetByte() uint8 {
	variantGetByteFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetByteFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint8()

	return retGo
}

var variantGetBytestringFunction *gi.Function
var variantGetBytestringFunction_Once sync.Once

func variantGetBytestringFunction_Set() {
	variantGetBytestringFunction_Once.Do(func() {
		variantGetBytestringFunction = gi.FunctionInvokerNew("GLib", "get_bytestring")
	})
}

var getBytestringVariantInvoker *gi.Function

// GetBytestring is a representation of the C type g_variant_get_bytestring.
func (recv *Variant) GetBytestring() {
	variantGetBytestringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantGetBytestringFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_get_bytestring_array' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_child' : parameter 'index_' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_child_value' : parameter 'index_' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_data' : return type 'gpointer' not supported

var variantGetDataAsBytesFunction *gi.Function
var variantGetDataAsBytesFunction_Once sync.Once

func variantGetDataAsBytesFunction_Set() {
	variantGetDataAsBytesFunction_Once.Do(func() {
		variantGetDataAsBytesFunction = gi.FunctionInvokerNew("GLib", "get_data_as_bytes")
	})
}

var getDataAsBytesVariantInvoker *gi.Function

// GetDataAsBytes is a representation of the C type g_variant_get_data_as_bytes.
func (recv *Variant) GetDataAsBytes() *Bytes {
	variantGetDataAsBytesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetDataAsBytesFunction.Invoke(inArgs[:], nil)

	retGo := &Bytes{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_double' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'g_variant_get_fixed_array' : parameter 'n_elements' of type 'gsize' not supported

var variantGetHandleFunction *gi.Function
var variantGetHandleFunction_Once sync.Once

func variantGetHandleFunction_Set() {
	variantGetHandleFunction_Once.Do(func() {
		variantGetHandleFunction = gi.FunctionInvokerNew("GLib", "get_handle")
	})
}

var getHandleVariantInvoker *gi.Function

// GetHandle is a representation of the C type g_variant_get_handle.
func (recv *Variant) GetHandle() int32 {
	variantGetHandleFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetHandleFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var variantGetInt16Function *gi.Function
var variantGetInt16Function_Once sync.Once

func variantGetInt16Function_Set() {
	variantGetInt16Function_Once.Do(func() {
		variantGetInt16Function = gi.FunctionInvokerNew("GLib", "get_int16")
	})
}

var getInt16VariantInvoker *gi.Function

// GetInt16 is a representation of the C type g_variant_get_int16.
func (recv *Variant) GetInt16() int16 {
	variantGetInt16Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetInt16Function.Invoke(inArgs[:], nil)

	retGo := ret.Int16()

	return retGo
}

var variantGetInt32Function *gi.Function
var variantGetInt32Function_Once sync.Once

func variantGetInt32Function_Set() {
	variantGetInt32Function_Once.Do(func() {
		variantGetInt32Function = gi.FunctionInvokerNew("GLib", "get_int32")
	})
}

var getInt32VariantInvoker *gi.Function

// GetInt32 is a representation of the C type g_variant_get_int32.
func (recv *Variant) GetInt32() int32 {
	variantGetInt32Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetInt32Function.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var variantGetInt64Function *gi.Function
var variantGetInt64Function_Once sync.Once

func variantGetInt64Function_Set() {
	variantGetInt64Function_Once.Do(func() {
		variantGetInt64Function = gi.FunctionInvokerNew("GLib", "get_int64")
	})
}

var getInt64VariantInvoker *gi.Function

// GetInt64 is a representation of the C type g_variant_get_int64.
func (recv *Variant) GetInt64() int64 {
	variantGetInt64Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetInt64Function.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var variantGetMaybeFunction *gi.Function
var variantGetMaybeFunction_Once sync.Once

func variantGetMaybeFunction_Set() {
	variantGetMaybeFunction_Once.Do(func() {
		variantGetMaybeFunction = gi.FunctionInvokerNew("GLib", "get_maybe")
	})
}

var getMaybeVariantInvoker *gi.Function

// GetMaybe is a representation of the C type g_variant_get_maybe.
func (recv *Variant) GetMaybe() *Variant {
	variantGetMaybeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetMaybeFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantGetNormalFormFunction *gi.Function
var variantGetNormalFormFunction_Once sync.Once

func variantGetNormalFormFunction_Set() {
	variantGetNormalFormFunction_Once.Do(func() {
		variantGetNormalFormFunction = gi.FunctionInvokerNew("GLib", "get_normal_form")
	})
}

var getNormalFormVariantInvoker *gi.Function

// GetNormalForm is a representation of the C type g_variant_get_normal_form.
func (recv *Variant) GetNormalForm() *Variant {
	variantGetNormalFormFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetNormalFormFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_objv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_size' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_string' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_strv' : parameter 'length' of type 'gsize' not supported

var variantGetTypeFunction *gi.Function
var variantGetTypeFunction_Once sync.Once

func variantGetTypeFunction_Set() {
	variantGetTypeFunction_Once.Do(func() {
		variantGetTypeFunction = gi.FunctionInvokerNew("GLib", "get_type")
	})
}

var getTypeVariantInvoker *gi.Function

// GetType is a representation of the C type g_variant_get_type.
func (recv *Variant) GetType() *VariantType {
	variantGetTypeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetTypeFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var variantGetTypeStringFunction *gi.Function
var variantGetTypeStringFunction_Once sync.Once

func variantGetTypeStringFunction_Set() {
	variantGetTypeStringFunction_Once.Do(func() {
		variantGetTypeStringFunction = gi.FunctionInvokerNew("GLib", "get_type_string")
	})
}

var getTypeStringVariantInvoker *gi.Function

// GetTypeString is a representation of the C type g_variant_get_type_string.
func (recv *Variant) GetTypeString() string {
	variantGetTypeStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetTypeStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var variantGetUint16Function *gi.Function
var variantGetUint16Function_Once sync.Once

func variantGetUint16Function_Set() {
	variantGetUint16Function_Once.Do(func() {
		variantGetUint16Function = gi.FunctionInvokerNew("GLib", "get_uint16")
	})
}

var getUint16VariantInvoker *gi.Function

// GetUint16 is a representation of the C type g_variant_get_uint16.
func (recv *Variant) GetUint16() uint16 {
	variantGetUint16Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetUint16Function.Invoke(inArgs[:], nil)

	retGo := ret.Uint16()

	return retGo
}

var variantGetUint32Function *gi.Function
var variantGetUint32Function_Once sync.Once

func variantGetUint32Function_Set() {
	variantGetUint32Function_Once.Do(func() {
		variantGetUint32Function = gi.FunctionInvokerNew("GLib", "get_uint32")
	})
}

var getUint32VariantInvoker *gi.Function

// GetUint32 is a representation of the C type g_variant_get_uint32.
func (recv *Variant) GetUint32() uint32 {
	variantGetUint32Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetUint32Function.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var variantGetUint64Function *gi.Function
var variantGetUint64Function_Once sync.Once

func variantGetUint64Function_Set() {
	variantGetUint64Function_Once.Do(func() {
		variantGetUint64Function = gi.FunctionInvokerNew("GLib", "get_uint64")
	})
}

var getUint64VariantInvoker *gi.Function

// GetUint64 is a representation of the C type g_variant_get_uint64.
func (recv *Variant) GetUint64() uint64 {
	variantGetUint64Function_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetUint64Function.Invoke(inArgs[:], nil)

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_variant_get_va' : parameter 'app' of type 'va_list' not supported

var variantGetVariantFunction *gi.Function
var variantGetVariantFunction_Once sync.Once

func variantGetVariantFunction_Set() {
	variantGetVariantFunction_Once.Do(func() {
		variantGetVariantFunction = gi.FunctionInvokerNew("GLib", "get_variant")
	})
}

var getVariantVariantInvoker *gi.Function

// GetVariant is a representation of the C type g_variant_get_variant.
func (recv *Variant) GetVariant() *Variant {
	variantGetVariantFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantGetVariantFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantHashFunction *gi.Function
var variantHashFunction_Once sync.Once

func variantHashFunction_Set() {
	variantHashFunction_Once.Do(func() {
		variantHashFunction = gi.FunctionInvokerNew("GLib", "hash")
	})
}

var hashVariantInvoker *gi.Function

// Hash is a representation of the C type g_variant_hash.
func (recv *Variant) Hash() uint32 {
	variantHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantHashFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'g_variant_is_container' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_floating' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_normal_form' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_is_of_type' : parameter 'type' of type 'VariantType' not supported

var variantIterNewFunction *gi.Function
var variantIterNewFunction_Once sync.Once

func variantIterNewFunction_Set() {
	variantIterNewFunction_Once.Do(func() {
		variantIterNewFunction = gi.FunctionInvokerNew("GLib", "iter_new")
	})
}

var iterNewVariantInvoker *gi.Function

// IterNew is a representation of the C type g_variant_iter_new.
func (recv *Variant) IterNew() *VariantIter {
	variantIterNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantIterNewFunction.Invoke(inArgs[:], nil)

	retGo := &VariantIter{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_lookup' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_lookup_value' : parameter 'expected_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_n_children' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_print' : parameter 'type_annotate' of type 'gboolean' not supported

// UNSUPPORTED : C value 'g_variant_print_string' : parameter 'string' of type 'String' not supported

var variantRefFunction *gi.Function
var variantRefFunction_Once sync.Once

func variantRefFunction_Set() {
	variantRefFunction_Once.Do(func() {
		variantRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refVariantInvoker *gi.Function

// Ref is a representation of the C type g_variant_ref.
func (recv *Variant) Ref() *Variant {
	variantRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantRefFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantRefSinkFunction *gi.Function
var variantRefSinkFunction_Once sync.Once

func variantRefSinkFunction_Set() {
	variantRefSinkFunction_Once.Do(func() {
		variantRefSinkFunction = gi.FunctionInvokerNew("GLib", "ref_sink")
	})
}

var refSinkVariantInvoker *gi.Function

// RefSink is a representation of the C type g_variant_ref_sink.
func (recv *Variant) RefSink() *Variant {
	variantRefSinkFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantRefSinkFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_store' : parameter 'data' of type 'gpointer' not supported

var variantTakeRefFunction *gi.Function
var variantTakeRefFunction_Once sync.Once

func variantTakeRefFunction_Set() {
	variantTakeRefFunction_Once.Do(func() {
		variantTakeRefFunction = gi.FunctionInvokerNew("GLib", "take_ref")
	})
}

var takeRefVariantInvoker *gi.Function

// TakeRef is a representation of the C type g_variant_take_ref.
func (recv *Variant) TakeRef() *Variant {
	variantTakeRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTakeRefFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantUnrefFunction *gi.Function
var variantUnrefFunction_Once sync.Once

func variantUnrefFunction_Set() {
	variantUnrefFunction_Once.Do(func() {
		variantUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefVariantInvoker *gi.Function

// Unref is a representation of the C type g_variant_unref.
func (recv *Variant) Unref() {
	variantUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantUnrefFunction.Invoke(inArgs[:], nil)

}

var variantBuilderStruct *gi.Struct
var variantBuilderStructOnce sync.Once

func variantBuilderStructSet() {
	variantBuilderStructOnce.Do(func() {
		variantBuilderStruct = gi.StructNew("GLib", "VariantBuilder")
	})
}

type VariantBuilder struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_builder_new' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_builder_add' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_builder_add_parsed' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_builder_add_value' : parameter 'value' of type 'Variant' not supported

var variantBuilderClearFunction *gi.Function
var variantBuilderClearFunction_Once sync.Once

func variantBuilderClearFunction_Set() {
	variantBuilderClearFunction_Once.Do(func() {
		variantBuilderClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearVariantBuilderInvoker *gi.Function

// Clear is a representation of the C type g_variant_builder_clear.
func (recv *VariantBuilder) Clear() {
	variantBuilderClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantBuilderClearFunction.Invoke(inArgs[:], nil)

}

var variantBuilderCloseFunction *gi.Function
var variantBuilderCloseFunction_Once sync.Once

func variantBuilderCloseFunction_Set() {
	variantBuilderCloseFunction_Once.Do(func() {
		variantBuilderCloseFunction = gi.FunctionInvokerNew("GLib", "close")
	})
}

var closeVariantBuilderInvoker *gi.Function

// Close is a representation of the C type g_variant_builder_close.
func (recv *VariantBuilder) Close() {
	variantBuilderCloseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantBuilderCloseFunction.Invoke(inArgs[:], nil)

}

var variantBuilderEndFunction *gi.Function
var variantBuilderEndFunction_Once sync.Once

func variantBuilderEndFunction_Set() {
	variantBuilderEndFunction_Once.Do(func() {
		variantBuilderEndFunction = gi.FunctionInvokerNew("GLib", "end")
	})
}

var endVariantBuilderInvoker *gi.Function

// End is a representation of the C type g_variant_builder_end.
func (recv *VariantBuilder) End() *Variant {
	variantBuilderEndFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantBuilderEndFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_builder_init' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_builder_open' : parameter 'type' of type 'VariantType' not supported

var variantBuilderRefFunction *gi.Function
var variantBuilderRefFunction_Once sync.Once

func variantBuilderRefFunction_Set() {
	variantBuilderRefFunction_Once.Do(func() {
		variantBuilderRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refVariantBuilderInvoker *gi.Function

// Ref is a representation of the C type g_variant_builder_ref.
func (recv *VariantBuilder) Ref() *VariantBuilder {
	variantBuilderRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantBuilderRefFunction.Invoke(inArgs[:], nil)

	retGo := &VariantBuilder{native: ret.Pointer()}

	return retGo
}

var variantBuilderUnrefFunction *gi.Function
var variantBuilderUnrefFunction_Once sync.Once

func variantBuilderUnrefFunction_Set() {
	variantBuilderUnrefFunction_Once.Do(func() {
		variantBuilderUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefVariantBuilderInvoker *gi.Function

// Unref is a representation of the C type g_variant_builder_unref.
func (recv *VariantBuilder) Unref() {
	variantBuilderUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantBuilderUnrefFunction.Invoke(inArgs[:], nil)

}

var variantDictStruct *gi.Struct
var variantDictStructOnce sync.Once

func variantDictStructSet() {
	variantDictStructOnce.Do(func() {
		variantDictStruct = gi.StructNew("GLib", "VariantDict")
	})
}

type VariantDict struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_dict_new' : parameter 'from_asv' of type 'Variant' not supported

var variantDictClearFunction *gi.Function
var variantDictClearFunction_Once sync.Once

func variantDictClearFunction_Set() {
	variantDictClearFunction_Once.Do(func() {
		variantDictClearFunction = gi.FunctionInvokerNew("GLib", "clear")
	})
}

var clearVariantDictInvoker *gi.Function

// Clear is a representation of the C type g_variant_dict_clear.
func (recv *VariantDict) Clear() {
	variantDictClearFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantDictClearFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_dict_contains' : return type 'gboolean' not supported

var variantDictEndFunction *gi.Function
var variantDictEndFunction_Once sync.Once

func variantDictEndFunction_Set() {
	variantDictEndFunction_Once.Do(func() {
		variantDictEndFunction = gi.FunctionInvokerNew("GLib", "end")
	})
}

var endVariantDictInvoker *gi.Function

// End is a representation of the C type g_variant_dict_end.
func (recv *VariantDict) End() *Variant {
	variantDictEndFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantDictEndFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_dict_init' : parameter 'from_asv' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dict_insert' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_dict_insert_value' : parameter 'value' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dict_lookup' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_dict_lookup_value' : parameter 'expected_type' of type 'VariantType' not supported

var variantDictRefFunction *gi.Function
var variantDictRefFunction_Once sync.Once

func variantDictRefFunction_Set() {
	variantDictRefFunction_Once.Do(func() {
		variantDictRefFunction = gi.FunctionInvokerNew("GLib", "ref")
	})
}

var refVariantDictInvoker *gi.Function

// Ref is a representation of the C type g_variant_dict_ref.
func (recv *VariantDict) Ref() *VariantDict {
	variantDictRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantDictRefFunction.Invoke(inArgs[:], nil)

	retGo := &VariantDict{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_dict_remove' : return type 'gboolean' not supported

var variantDictUnrefFunction *gi.Function
var variantDictUnrefFunction_Once sync.Once

func variantDictUnrefFunction_Set() {
	variantDictUnrefFunction_Once.Do(func() {
		variantDictUnrefFunction = gi.FunctionInvokerNew("GLib", "unref")
	})
}

var unrefVariantDictInvoker *gi.Function

// Unref is a representation of the C type g_variant_dict_unref.
func (recv *VariantDict) Unref() {
	variantDictUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantDictUnrefFunction.Invoke(inArgs[:], nil)

}

var variantIterStruct *gi.Struct
var variantIterStructOnce sync.Once

func variantIterStructSet() {
	variantIterStructOnce.Do(func() {
		variantIterStruct = gi.StructNew("GLib", "VariantIter")
	})
}

type VariantIter struct {
	native uintptr
}

var variantIterCopyFunction *gi.Function
var variantIterCopyFunction_Once sync.Once

func variantIterCopyFunction_Set() {
	variantIterCopyFunction_Once.Do(func() {
		variantIterCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyVariantIterInvoker *gi.Function

// Copy is a representation of the C type g_variant_iter_copy.
func (recv *VariantIter) Copy() *VariantIter {
	variantIterCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantIterCopyFunction.Invoke(inArgs[:], nil)

	retGo := &VariantIter{native: ret.Pointer()}

	return retGo
}

var variantIterFreeFunction *gi.Function
var variantIterFreeFunction_Once sync.Once

func variantIterFreeFunction_Set() {
	variantIterFreeFunction_Once.Do(func() {
		variantIterFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeVariantIterInvoker *gi.Function

// Free is a representation of the C type g_variant_iter_free.
func (recv *VariantIter) Free() {
	variantIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantIterFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_iter_init' : parameter 'value' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_iter_loop' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_iter_n_children' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_iter_next' : parameter '...' has no type

var variantIterNextValueFunction *gi.Function
var variantIterNextValueFunction_Once sync.Once

func variantIterNextValueFunction_Set() {
	variantIterNextValueFunction_Once.Do(func() {
		variantIterNextValueFunction = gi.FunctionInvokerNew("GLib", "next_value")
	})
}

var nextValueVariantIterInvoker *gi.Function

// NextValue is a representation of the C type g_variant_iter_next_value.
func (recv *VariantIter) NextValue() *Variant {
	variantIterNextValueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantIterNextValueFunction.Invoke(inArgs[:], nil)

	retGo := &Variant{native: ret.Pointer()}

	return retGo
}

var variantTypeStruct *gi.Struct
var variantTypeStructOnce sync.Once

func variantTypeStructSet() {
	variantTypeStructOnce.Do(func() {
		variantTypeStruct = gi.StructNew("GLib", "VariantType")
	})
}

type VariantType struct {
	native uintptr
}

var variantTypeNewFunction *gi.Function
var variantTypeNewFunction_Once sync.Once

func variantTypeNewFunction_Set() {
	variantTypeNewFunction_Once.Do(func() {
		variantTypeNewFunction = gi.FunctionInvokerNew("GLib", "new")
	})
}

var newVariantTypeInvoker *gi.Function

// VariantTypeNew is a representation of the C type g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	variantTypeNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	ret := variantTypeNewFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_new_array' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_dict_entry' : parameter 'key' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_maybe' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_tuple' : parameter 'items' has no type

var variantTypeCopyFunction *gi.Function
var variantTypeCopyFunction_Once sync.Once

func variantTypeCopyFunction_Set() {
	variantTypeCopyFunction_Once.Do(func() {
		variantTypeCopyFunction = gi.FunctionInvokerNew("GLib", "copy")
	})
}

var copyVariantTypeInvoker *gi.Function

// Copy is a representation of the C type g_variant_type_copy.
func (recv *VariantType) Copy() *VariantType {
	variantTypeCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeCopyFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var variantTypeDupStringFunction *gi.Function
var variantTypeDupStringFunction_Once sync.Once

func variantTypeDupStringFunction_Set() {
	variantTypeDupStringFunction_Once.Do(func() {
		variantTypeDupStringFunction = gi.FunctionInvokerNew("GLib", "dup_string")
	})
}

var dupStringVariantTypeInvoker *gi.Function

// DupString is a representation of the C type g_variant_type_dup_string.
func (recv *VariantType) DupString() string {
	variantTypeDupStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeDupStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var variantTypeElementFunction *gi.Function
var variantTypeElementFunction_Once sync.Once

func variantTypeElementFunction_Set() {
	variantTypeElementFunction_Once.Do(func() {
		variantTypeElementFunction = gi.FunctionInvokerNew("GLib", "element")
	})
}

var elementVariantTypeInvoker *gi.Function

// Element is a representation of the C type g_variant_type_element.
func (recv *VariantType) Element() *VariantType {
	variantTypeElementFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeElementFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_equal' : parameter 'type2' of type 'VariantType' not supported

var variantTypeFirstFunction *gi.Function
var variantTypeFirstFunction_Once sync.Once

func variantTypeFirstFunction_Set() {
	variantTypeFirstFunction_Once.Do(func() {
		variantTypeFirstFunction = gi.FunctionInvokerNew("GLib", "first")
	})
}

var firstVariantTypeInvoker *gi.Function

// First is a representation of the C type g_variant_type_first.
func (recv *VariantType) First() *VariantType {
	variantTypeFirstFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeFirstFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var variantTypeFreeFunction *gi.Function
var variantTypeFreeFunction_Once sync.Once

func variantTypeFreeFunction_Set() {
	variantTypeFreeFunction_Once.Do(func() {
		variantTypeFreeFunction = gi.FunctionInvokerNew("GLib", "free")
	})
}

var freeVariantTypeInvoker *gi.Function

// Free is a representation of the C type g_variant_type_free.
func (recv *VariantType) Free() {
	variantTypeFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	variantTypeFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'g_variant_type_get_string_length' : return type 'gsize' not supported

var variantTypeHashFunction *gi.Function
var variantTypeHashFunction_Once sync.Once

func variantTypeHashFunction_Set() {
	variantTypeHashFunction_Once.Do(func() {
		variantTypeHashFunction = gi.FunctionInvokerNew("GLib", "hash")
	})
}

var hashVariantTypeInvoker *gi.Function

// Hash is a representation of the C type g_variant_type_hash.
func (recv *VariantType) Hash() uint32 {
	variantTypeHashFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeHashFunction.Invoke(inArgs[:], nil)

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

var variantTypeKeyFunction *gi.Function
var variantTypeKeyFunction_Once sync.Once

func variantTypeKeyFunction_Set() {
	variantTypeKeyFunction_Once.Do(func() {
		variantTypeKeyFunction = gi.FunctionInvokerNew("GLib", "key")
	})
}

var keyVariantTypeInvoker *gi.Function

// Key is a representation of the C type g_variant_type_key.
func (recv *VariantType) Key() *VariantType {
	variantTypeKeyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeKeyFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'g_variant_type_n_items' : return type 'gsize' not supported

var variantTypeNextFunction *gi.Function
var variantTypeNextFunction_Once sync.Once

func variantTypeNextFunction_Set() {
	variantTypeNextFunction_Once.Do(func() {
		variantTypeNextFunction = gi.FunctionInvokerNew("GLib", "next")
	})
}

var nextVariantTypeInvoker *gi.Function

// Next is a representation of the C type g_variant_type_next.
func (recv *VariantType) Next() *VariantType {
	variantTypeNextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeNextFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}

var variantTypePeekStringFunction *gi.Function
var variantTypePeekStringFunction_Once sync.Once

func variantTypePeekStringFunction_Set() {
	variantTypePeekStringFunction_Once.Do(func() {
		variantTypePeekStringFunction = gi.FunctionInvokerNew("GLib", "peek_string")
	})
}

var peekStringVariantTypeInvoker *gi.Function

// PeekString is a representation of the C type g_variant_type_peek_string.
func (recv *VariantType) PeekString() string {
	variantTypePeekStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypePeekStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var variantTypeValueFunction *gi.Function
var variantTypeValueFunction_Once sync.Once

func variantTypeValueFunction_Set() {
	variantTypeValueFunction_Once.Do(func() {
		variantTypeValueFunction = gi.FunctionInvokerNew("GLib", "value")
	})
}

var valueVariantTypeInvoker *gi.Function

// Value is a representation of the C type g_variant_type_value.
func (recv *VariantType) Value() *VariantType {
	variantTypeValueFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := variantTypeValueFunction.Invoke(inArgs[:], nil)

	retGo := &VariantType{native: ret.Pointer()}

	return retGo
}
