// Code generated - DO NOT EDIT.

package glib

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var arrayStruct *gi.Struct
var arrayStruct_Once sync.Once

func arrayStruct_Set() error {
	var err error
	arrayStruct_Once.Do(func() {
		arrayStruct, err = gi.StructNew("GLib", "Array")
	})
	return err
}

type Array struct {
	native uintptr
	Data   string
	Len    uint32
}

var asyncQueueStruct *gi.Struct
var asyncQueueStruct_Once sync.Once

func asyncQueueStruct_Set() error {
	var err error
	asyncQueueStruct_Once.Do(func() {
		asyncQueueStruct, err = gi.StructNew("GLib", "AsyncQueue")
	})
	return err
}

type AsyncQueue struct {
	native uintptr
}

var asyncQueueLengthFunction *gi.Function
var asyncQueueLengthFunction_Once sync.Once

func asyncQueueLengthFunction_Set() error {
	var err error
	asyncQueueLengthFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueLengthFunction, err = asyncQueueStruct.InvokerNew("length")
	})
	return err
}

// Length is a representation of the C type g_async_queue_length.
func (recv *AsyncQueue) Length() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := asyncQueueLengthFunction_Set()
	if err == nil {
		ret = asyncQueueLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var asyncQueueLengthUnlockedFunction *gi.Function
var asyncQueueLengthUnlockedFunction_Once sync.Once

func asyncQueueLengthUnlockedFunction_Set() error {
	var err error
	asyncQueueLengthUnlockedFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueLengthUnlockedFunction, err = asyncQueueStruct.InvokerNew("length_unlocked")
	})
	return err
}

// LengthUnlocked is a representation of the C type g_async_queue_length_unlocked.
func (recv *AsyncQueue) LengthUnlocked() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := asyncQueueLengthUnlockedFunction_Set()
	if err == nil {
		ret = asyncQueueLengthUnlockedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var asyncQueueLockFunction *gi.Function
var asyncQueueLockFunction_Once sync.Once

func asyncQueueLockFunction_Set() error {
	var err error
	asyncQueueLockFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueLockFunction, err = asyncQueueStruct.InvokerNew("lock")
	})
	return err
}

// Lock is a representation of the C type g_async_queue_lock.
func (recv *AsyncQueue) Lock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := asyncQueueLockFunction_Set()
	if err == nil {
		asyncQueueLockFunction.Invoke(inArgs[:], nil)
	}

	return err
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

func asyncQueueRefFunction_Set() error {
	var err error
	asyncQueueRefFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueRefFunction, err = asyncQueueStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_async_queue_ref.
func (recv *AsyncQueue) Ref() (*AsyncQueue, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := asyncQueueRefFunction_Set()
	if err == nil {
		ret = asyncQueueRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AsyncQueue{native: ret.Pointer()}

	return retGo, err
}

var asyncQueueRefUnlockedFunction *gi.Function
var asyncQueueRefUnlockedFunction_Once sync.Once

func asyncQueueRefUnlockedFunction_Set() error {
	var err error
	asyncQueueRefUnlockedFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueRefUnlockedFunction, err = asyncQueueStruct.InvokerNew("ref_unlocked")
	})
	return err
}

// RefUnlocked is a representation of the C type g_async_queue_ref_unlocked.
func (recv *AsyncQueue) RefUnlocked() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := asyncQueueRefUnlockedFunction_Set()
	if err == nil {
		asyncQueueRefUnlockedFunction.Invoke(inArgs[:], nil)
	}

	return err
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

func asyncQueueUnlockFunction_Set() error {
	var err error
	asyncQueueUnlockFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueUnlockFunction, err = asyncQueueStruct.InvokerNew("unlock")
	})
	return err
}

// Unlock is a representation of the C type g_async_queue_unlock.
func (recv *AsyncQueue) Unlock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := asyncQueueUnlockFunction_Set()
	if err == nil {
		asyncQueueUnlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var asyncQueueUnrefFunction *gi.Function
var asyncQueueUnrefFunction_Once sync.Once

func asyncQueueUnrefFunction_Set() error {
	var err error
	asyncQueueUnrefFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueUnrefFunction, err = asyncQueueStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_async_queue_unref.
func (recv *AsyncQueue) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := asyncQueueUnrefFunction_Set()
	if err == nil {
		asyncQueueUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var asyncQueueUnrefAndUnlockFunction *gi.Function
var asyncQueueUnrefAndUnlockFunction_Once sync.Once

func asyncQueueUnrefAndUnlockFunction_Set() error {
	var err error
	asyncQueueUnrefAndUnlockFunction_Once.Do(func() {
		err = asyncQueueStruct_Set()
		if err != nil {
			return
		}
		asyncQueueUnrefAndUnlockFunction, err = asyncQueueStruct.InvokerNew("unref_and_unlock")
	})
	return err
}

// UnrefAndUnlock is a representation of the C type g_async_queue_unref_and_unlock.
func (recv *AsyncQueue) UnrefAndUnlock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := asyncQueueUnrefAndUnlockFunction_Set()
	if err == nil {
		asyncQueueUnrefAndUnlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileStruct *gi.Struct
var bookmarkFileStruct_Once sync.Once

func bookmarkFileStruct_Set() error {
	var err error
	bookmarkFileStruct_Once.Do(func() {
		bookmarkFileStruct, err = gi.StructNew("GLib", "BookmarkFile")
	})
	return err
}

type BookmarkFile struct {
	native uintptr
}

var bookmarkFileAddApplicationFunction *gi.Function
var bookmarkFileAddApplicationFunction_Once sync.Once

func bookmarkFileAddApplicationFunction_Set() error {
	var err error
	bookmarkFileAddApplicationFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileAddApplicationFunction, err = bookmarkFileStruct.InvokerNew("add_application")
	})
	return err
}

// AddApplication is a representation of the C type g_bookmark_file_add_application.
func (recv *BookmarkFile) AddApplication(uri string, name string, exec string) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)
	inArgs[3].SetString(exec)

	err := bookmarkFileAddApplicationFunction_Set()
	if err == nil {
		bookmarkFileAddApplicationFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileAddGroupFunction *gi.Function
var bookmarkFileAddGroupFunction_Once sync.Once

func bookmarkFileAddGroupFunction_Set() error {
	var err error
	bookmarkFileAddGroupFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileAddGroupFunction, err = bookmarkFileStruct.InvokerNew("add_group")
	})
	return err
}

// AddGroup is a representation of the C type g_bookmark_file_add_group.
func (recv *BookmarkFile) AddGroup(uri string, group string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(group)

	err := bookmarkFileAddGroupFunction_Set()
	if err == nil {
		bookmarkFileAddGroupFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileFreeFunction *gi.Function
var bookmarkFileFreeFunction_Once sync.Once

func bookmarkFileFreeFunction_Set() error {
	var err error
	bookmarkFileFreeFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileFreeFunction, err = bookmarkFileStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_bookmark_file_free.
func (recv *BookmarkFile) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := bookmarkFileFreeFunction_Set()
	if err == nil {
		bookmarkFileFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileGetAddedFunction *gi.Function
var bookmarkFileGetAddedFunction_Once sync.Once

func bookmarkFileGetAddedFunction_Set() error {
	var err error
	bookmarkFileGetAddedFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetAddedFunction, err = bookmarkFileStruct.InvokerNew("get_added")
	})
	return err
}

// GetAdded is a representation of the C type g_bookmark_file_get_added.
func (recv *BookmarkFile) GetAdded(uri string) (int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetAddedFunction_Set()
	if err == nil {
		ret = bookmarkFileGetAddedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var bookmarkFileGetAppInfoFunction *gi.Function
var bookmarkFileGetAppInfoFunction_Once sync.Once

func bookmarkFileGetAppInfoFunction_Set() error {
	var err error
	bookmarkFileGetAppInfoFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetAppInfoFunction, err = bookmarkFileStruct.InvokerNew("get_app_info")
	})
	return err
}

// GetAppInfo is a representation of the C type g_bookmark_file_get_app_info.
func (recv *BookmarkFile) GetAppInfo(uri string, name string) (bool, string, uint32, int64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := bookmarkFileGetAppInfoFunction_Set()
	if err == nil {
		ret = bookmarkFileGetAppInfoFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)
	out1 := outArgs[1].Uint32()
	out2 := outArgs[2].Int64()

	return retGo, out0, out1, out2, err
}

// UNSUPPORTED : C value 'g_bookmark_file_get_applications' : parameter 'length' of type 'gsize' not supported

var bookmarkFileGetDescriptionFunction *gi.Function
var bookmarkFileGetDescriptionFunction_Once sync.Once

func bookmarkFileGetDescriptionFunction_Set() error {
	var err error
	bookmarkFileGetDescriptionFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetDescriptionFunction, err = bookmarkFileStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type g_bookmark_file_get_description.
func (recv *BookmarkFile) GetDescription(uri string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetDescriptionFunction_Set()
	if err == nil {
		ret = bookmarkFileGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_bookmark_file_get_groups' : parameter 'length' of type 'gsize' not supported

var bookmarkFileGetIconFunction *gi.Function
var bookmarkFileGetIconFunction_Once sync.Once

func bookmarkFileGetIconFunction_Set() error {
	var err error
	bookmarkFileGetIconFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetIconFunction, err = bookmarkFileStruct.InvokerNew("get_icon")
	})
	return err
}

// GetIcon is a representation of the C type g_bookmark_file_get_icon.
func (recv *BookmarkFile) GetIcon(uri string) (bool, string, string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := bookmarkFileGetIconFunction_Set()
	if err == nil {
		ret = bookmarkFileGetIconFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(true)
	out1 := outArgs[1].String(true)

	return retGo, out0, out1, err
}

var bookmarkFileGetIsPrivateFunction *gi.Function
var bookmarkFileGetIsPrivateFunction_Once sync.Once

func bookmarkFileGetIsPrivateFunction_Set() error {
	var err error
	bookmarkFileGetIsPrivateFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetIsPrivateFunction, err = bookmarkFileStruct.InvokerNew("get_is_private")
	})
	return err
}

// GetIsPrivate is a representation of the C type g_bookmark_file_get_is_private.
func (recv *BookmarkFile) GetIsPrivate(uri string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetIsPrivateFunction_Set()
	if err == nil {
		ret = bookmarkFileGetIsPrivateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileGetMimeTypeFunction *gi.Function
var bookmarkFileGetMimeTypeFunction_Once sync.Once

func bookmarkFileGetMimeTypeFunction_Set() error {
	var err error
	bookmarkFileGetMimeTypeFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetMimeTypeFunction, err = bookmarkFileStruct.InvokerNew("get_mime_type")
	})
	return err
}

// GetMimeType is a representation of the C type g_bookmark_file_get_mime_type.
func (recv *BookmarkFile) GetMimeType(uri string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetMimeTypeFunction_Set()
	if err == nil {
		ret = bookmarkFileGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var bookmarkFileGetModifiedFunction *gi.Function
var bookmarkFileGetModifiedFunction_Once sync.Once

func bookmarkFileGetModifiedFunction_Set() error {
	var err error
	bookmarkFileGetModifiedFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetModifiedFunction, err = bookmarkFileStruct.InvokerNew("get_modified")
	})
	return err
}

// GetModified is a representation of the C type g_bookmark_file_get_modified.
func (recv *BookmarkFile) GetModified(uri string) (int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetModifiedFunction_Set()
	if err == nil {
		ret = bookmarkFileGetModifiedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var bookmarkFileGetSizeFunction *gi.Function
var bookmarkFileGetSizeFunction_Once sync.Once

func bookmarkFileGetSizeFunction_Set() error {
	var err error
	bookmarkFileGetSizeFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetSizeFunction, err = bookmarkFileStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type g_bookmark_file_get_size.
func (recv *BookmarkFile) GetSize() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := bookmarkFileGetSizeFunction_Set()
	if err == nil {
		ret = bookmarkFileGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var bookmarkFileGetTitleFunction *gi.Function
var bookmarkFileGetTitleFunction_Once sync.Once

func bookmarkFileGetTitleFunction_Set() error {
	var err error
	bookmarkFileGetTitleFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetTitleFunction, err = bookmarkFileStruct.InvokerNew("get_title")
	})
	return err
}

// GetTitle is a representation of the C type g_bookmark_file_get_title.
func (recv *BookmarkFile) GetTitle(uri string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetTitleFunction_Set()
	if err == nil {
		ret = bookmarkFileGetTitleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_bookmark_file_get_uris' : parameter 'length' of type 'gsize' not supported

var bookmarkFileGetVisitedFunction *gi.Function
var bookmarkFileGetVisitedFunction_Once sync.Once

func bookmarkFileGetVisitedFunction_Set() error {
	var err error
	bookmarkFileGetVisitedFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileGetVisitedFunction, err = bookmarkFileStruct.InvokerNew("get_visited")
	})
	return err
}

// GetVisited is a representation of the C type g_bookmark_file_get_visited.
func (recv *BookmarkFile) GetVisited(uri string) (int64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileGetVisitedFunction_Set()
	if err == nil {
		ret = bookmarkFileGetVisitedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var bookmarkFileHasApplicationFunction *gi.Function
var bookmarkFileHasApplicationFunction_Once sync.Once

func bookmarkFileHasApplicationFunction_Set() error {
	var err error
	bookmarkFileHasApplicationFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileHasApplicationFunction, err = bookmarkFileStruct.InvokerNew("has_application")
	})
	return err
}

// HasApplication is a representation of the C type g_bookmark_file_has_application.
func (recv *BookmarkFile) HasApplication(uri string, name string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)

	var ret gi.Argument

	err := bookmarkFileHasApplicationFunction_Set()
	if err == nil {
		ret = bookmarkFileHasApplicationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileHasGroupFunction *gi.Function
var bookmarkFileHasGroupFunction_Once sync.Once

func bookmarkFileHasGroupFunction_Set() error {
	var err error
	bookmarkFileHasGroupFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileHasGroupFunction, err = bookmarkFileStruct.InvokerNew("has_group")
	})
	return err
}

// HasGroup is a representation of the C type g_bookmark_file_has_group.
func (recv *BookmarkFile) HasGroup(uri string, group string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(group)

	var ret gi.Argument

	err := bookmarkFileHasGroupFunction_Set()
	if err == nil {
		ret = bookmarkFileHasGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileHasItemFunction *gi.Function
var bookmarkFileHasItemFunction_Once sync.Once

func bookmarkFileHasItemFunction_Set() error {
	var err error
	bookmarkFileHasItemFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileHasItemFunction, err = bookmarkFileStruct.InvokerNew("has_item")
	})
	return err
}

// HasItem is a representation of the C type g_bookmark_file_has_item.
func (recv *BookmarkFile) HasItem(uri string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileHasItemFunction_Set()
	if err == nil {
		ret = bookmarkFileHasItemFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_bookmark_file_load_from_data' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bookmark_file_load_from_data_dirs' : parameter 'file' of type 'filename' not supported

// UNSUPPORTED : C value 'g_bookmark_file_load_from_file' : parameter 'filename' of type 'filename' not supported

var bookmarkFileMoveItemFunction *gi.Function
var bookmarkFileMoveItemFunction_Once sync.Once

func bookmarkFileMoveItemFunction_Set() error {
	var err error
	bookmarkFileMoveItemFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileMoveItemFunction, err = bookmarkFileStruct.InvokerNew("move_item")
	})
	return err
}

// MoveItem is a representation of the C type g_bookmark_file_move_item.
func (recv *BookmarkFile) MoveItem(oldUri string, newUri string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(oldUri)
	inArgs[2].SetString(newUri)

	var ret gi.Argument

	err := bookmarkFileMoveItemFunction_Set()
	if err == nil {
		ret = bookmarkFileMoveItemFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileRemoveApplicationFunction *gi.Function
var bookmarkFileRemoveApplicationFunction_Once sync.Once

func bookmarkFileRemoveApplicationFunction_Set() error {
	var err error
	bookmarkFileRemoveApplicationFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileRemoveApplicationFunction, err = bookmarkFileStruct.InvokerNew("remove_application")
	})
	return err
}

// RemoveApplication is a representation of the C type g_bookmark_file_remove_application.
func (recv *BookmarkFile) RemoveApplication(uri string, name string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)

	var ret gi.Argument

	err := bookmarkFileRemoveApplicationFunction_Set()
	if err == nil {
		ret = bookmarkFileRemoveApplicationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileRemoveGroupFunction *gi.Function
var bookmarkFileRemoveGroupFunction_Once sync.Once

func bookmarkFileRemoveGroupFunction_Set() error {
	var err error
	bookmarkFileRemoveGroupFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileRemoveGroupFunction, err = bookmarkFileStruct.InvokerNew("remove_group")
	})
	return err
}

// RemoveGroup is a representation of the C type g_bookmark_file_remove_group.
func (recv *BookmarkFile) RemoveGroup(uri string, group string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(group)

	var ret gi.Argument

	err := bookmarkFileRemoveGroupFunction_Set()
	if err == nil {
		ret = bookmarkFileRemoveGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileRemoveItemFunction *gi.Function
var bookmarkFileRemoveItemFunction_Once sync.Once

func bookmarkFileRemoveItemFunction_Set() error {
	var err error
	bookmarkFileRemoveItemFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileRemoveItemFunction, err = bookmarkFileStruct.InvokerNew("remove_item")
	})
	return err
}

// RemoveItem is a representation of the C type g_bookmark_file_remove_item.
func (recv *BookmarkFile) RemoveItem(uri string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := bookmarkFileRemoveItemFunction_Set()
	if err == nil {
		ret = bookmarkFileRemoveItemFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileSetAddedFunction *gi.Function
var bookmarkFileSetAddedFunction_Once sync.Once

func bookmarkFileSetAddedFunction_Set() error {
	var err error
	bookmarkFileSetAddedFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetAddedFunction, err = bookmarkFileStruct.InvokerNew("set_added")
	})
	return err
}

// SetAdded is a representation of the C type g_bookmark_file_set_added.
func (recv *BookmarkFile) SetAdded(uri string, added int64) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(added)

	err := bookmarkFileSetAddedFunction_Set()
	if err == nil {
		bookmarkFileSetAddedFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileSetAppInfoFunction *gi.Function
var bookmarkFileSetAppInfoFunction_Once sync.Once

func bookmarkFileSetAppInfoFunction_Set() error {
	var err error
	bookmarkFileSetAppInfoFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetAppInfoFunction, err = bookmarkFileStruct.InvokerNew("set_app_info")
	})
	return err
}

// SetAppInfo is a representation of the C type g_bookmark_file_set_app_info.
func (recv *BookmarkFile) SetAppInfo(uri string, name string, exec string, count int32, stamp int64) (bool, error) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(name)
	inArgs[3].SetString(exec)
	inArgs[4].SetInt32(count)
	inArgs[5].SetInt64(stamp)

	var ret gi.Argument

	err := bookmarkFileSetAppInfoFunction_Set()
	if err == nil {
		ret = bookmarkFileSetAppInfoFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var bookmarkFileSetDescriptionFunction *gi.Function
var bookmarkFileSetDescriptionFunction_Once sync.Once

func bookmarkFileSetDescriptionFunction_Set() error {
	var err error
	bookmarkFileSetDescriptionFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetDescriptionFunction, err = bookmarkFileStruct.InvokerNew("set_description")
	})
	return err
}

// SetDescription is a representation of the C type g_bookmark_file_set_description.
func (recv *BookmarkFile) SetDescription(uri string, description string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(description)

	err := bookmarkFileSetDescriptionFunction_Set()
	if err == nil {
		bookmarkFileSetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_bookmark_file_set_groups' : parameter 'groups' has no type

var bookmarkFileSetIconFunction *gi.Function
var bookmarkFileSetIconFunction_Once sync.Once

func bookmarkFileSetIconFunction_Set() error {
	var err error
	bookmarkFileSetIconFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetIconFunction, err = bookmarkFileStruct.InvokerNew("set_icon")
	})
	return err
}

// SetIcon is a representation of the C type g_bookmark_file_set_icon.
func (recv *BookmarkFile) SetIcon(uri string, href string, mimeType string) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(href)
	inArgs[3].SetString(mimeType)

	err := bookmarkFileSetIconFunction_Set()
	if err == nil {
		bookmarkFileSetIconFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileSetIsPrivateFunction *gi.Function
var bookmarkFileSetIsPrivateFunction_Once sync.Once

func bookmarkFileSetIsPrivateFunction_Set() error {
	var err error
	bookmarkFileSetIsPrivateFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetIsPrivateFunction, err = bookmarkFileStruct.InvokerNew("set_is_private")
	})
	return err
}

// SetIsPrivate is a representation of the C type g_bookmark_file_set_is_private.
func (recv *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetBoolean(isPrivate)

	err := bookmarkFileSetIsPrivateFunction_Set()
	if err == nil {
		bookmarkFileSetIsPrivateFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileSetMimeTypeFunction *gi.Function
var bookmarkFileSetMimeTypeFunction_Once sync.Once

func bookmarkFileSetMimeTypeFunction_Set() error {
	var err error
	bookmarkFileSetMimeTypeFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetMimeTypeFunction, err = bookmarkFileStruct.InvokerNew("set_mime_type")
	})
	return err
}

// SetMimeType is a representation of the C type g_bookmark_file_set_mime_type.
func (recv *BookmarkFile) SetMimeType(uri string, mimeType string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(mimeType)

	err := bookmarkFileSetMimeTypeFunction_Set()
	if err == nil {
		bookmarkFileSetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileSetModifiedFunction *gi.Function
var bookmarkFileSetModifiedFunction_Once sync.Once

func bookmarkFileSetModifiedFunction_Set() error {
	var err error
	bookmarkFileSetModifiedFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetModifiedFunction, err = bookmarkFileStruct.InvokerNew("set_modified")
	})
	return err
}

// SetModified is a representation of the C type g_bookmark_file_set_modified.
func (recv *BookmarkFile) SetModified(uri string, modified int64) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(modified)

	err := bookmarkFileSetModifiedFunction_Set()
	if err == nil {
		bookmarkFileSetModifiedFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileSetTitleFunction *gi.Function
var bookmarkFileSetTitleFunction_Once sync.Once

func bookmarkFileSetTitleFunction_Set() error {
	var err error
	bookmarkFileSetTitleFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetTitleFunction, err = bookmarkFileStruct.InvokerNew("set_title")
	})
	return err
}

// SetTitle is a representation of the C type g_bookmark_file_set_title.
func (recv *BookmarkFile) SetTitle(uri string, title string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetString(title)

	err := bookmarkFileSetTitleFunction_Set()
	if err == nil {
		bookmarkFileSetTitleFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bookmarkFileSetVisitedFunction *gi.Function
var bookmarkFileSetVisitedFunction_Once sync.Once

func bookmarkFileSetVisitedFunction_Set() error {
	var err error
	bookmarkFileSetVisitedFunction_Once.Do(func() {
		err = bookmarkFileStruct_Set()
		if err != nil {
			return
		}
		bookmarkFileSetVisitedFunction, err = bookmarkFileStruct.InvokerNew("set_visited")
	})
	return err
}

// SetVisited is a representation of the C type g_bookmark_file_set_visited.
func (recv *BookmarkFile) SetVisited(uri string, visited int64) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(uri)
	inArgs[2].SetInt64(visited)

	err := bookmarkFileSetVisitedFunction_Set()
	if err == nil {
		bookmarkFileSetVisitedFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_bookmark_file_to_data' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_bookmark_file_to_file' : parameter 'filename' of type 'filename' not supported

var byteArrayStruct *gi.Struct
var byteArrayStruct_Once sync.Once

func byteArrayStruct_Set() error {
	var err error
	byteArrayStruct_Once.Do(func() {
		byteArrayStruct, err = gi.StructNew("GLib", "ByteArray")
	})
	return err
}

type ByteArray struct {
	native uintptr
	Data   uint8
	Len    uint32
}

var bytesStruct *gi.Struct
var bytesStruct_Once sync.Once

func bytesStruct_Set() error {
	var err error
	bytesStruct_Once.Do(func() {
		bytesStruct, err = gi.StructNew("GLib", "Bytes")
	})
	return err
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

func bytesHashFunction_Set() error {
	var err error
	bytesHashFunction_Once.Do(func() {
		err = bytesStruct_Set()
		if err != nil {
			return
		}
		bytesHashFunction, err = bytesStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type g_bytes_hash.
func (recv *Bytes) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := bytesHashFunction_Set()
	if err == nil {
		ret = bytesHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_bytes_new_from_bytes' : parameter 'offset' of type 'gsize' not supported

var bytesRefFunction *gi.Function
var bytesRefFunction_Once sync.Once

func bytesRefFunction_Set() error {
	var err error
	bytesRefFunction_Once.Do(func() {
		err = bytesStruct_Set()
		if err != nil {
			return
		}
		bytesRefFunction, err = bytesStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_bytes_ref.
func (recv *Bytes) Ref() (*Bytes, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := bytesRefFunction_Set()
	if err == nil {
		ret = bytesRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Bytes{native: ret.Pointer()}

	return retGo, err
}

var bytesUnrefFunction *gi.Function
var bytesUnrefFunction_Once sync.Once

func bytesUnrefFunction_Set() error {
	var err error
	bytesUnrefFunction_Once.Do(func() {
		err = bytesStruct_Set()
		if err != nil {
			return
		}
		bytesUnrefFunction, err = bytesStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_bytes_unref.
func (recv *Bytes) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := bytesUnrefFunction_Set()
	if err == nil {
		bytesUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var bytesUnrefToArrayFunction *gi.Function
var bytesUnrefToArrayFunction_Once sync.Once

func bytesUnrefToArrayFunction_Set() error {
	var err error
	bytesUnrefToArrayFunction_Once.Do(func() {
		err = bytesStruct_Set()
		if err != nil {
			return
		}
		bytesUnrefToArrayFunction, err = bytesStruct.InvokerNew("unref_to_array")
	})
	return err
}

// UnrefToArray is a representation of the C type g_bytes_unref_to_array.
func (recv *Bytes) UnrefToArray() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := bytesUnrefToArrayFunction_Set()
	if err == nil {
		bytesUnrefToArrayFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_bytes_unref_to_data' : parameter 'size' of type 'gsize' not supported

var checksumStruct *gi.Struct
var checksumStruct_Once sync.Once

func checksumStruct_Set() error {
	var err error
	checksumStruct_Once.Do(func() {
		checksumStruct, err = gi.StructNew("GLib", "Checksum")
	})
	return err
}

type Checksum struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_checksum_new' : parameter 'checksum_type' of type 'ChecksumType' not supported

var checksumCopyFunction *gi.Function
var checksumCopyFunction_Once sync.Once

func checksumCopyFunction_Set() error {
	var err error
	checksumCopyFunction_Once.Do(func() {
		err = checksumStruct_Set()
		if err != nil {
			return
		}
		checksumCopyFunction, err = checksumStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_checksum_copy.
func (recv *Checksum) Copy() (*Checksum, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := checksumCopyFunction_Set()
	if err == nil {
		ret = checksumCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Checksum{native: ret.Pointer()}

	return retGo, err
}

var checksumFreeFunction *gi.Function
var checksumFreeFunction_Once sync.Once

func checksumFreeFunction_Set() error {
	var err error
	checksumFreeFunction_Once.Do(func() {
		err = checksumStruct_Set()
		if err != nil {
			return
		}
		checksumFreeFunction, err = checksumStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_checksum_free.
func (recv *Checksum) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := checksumFreeFunction_Set()
	if err == nil {
		checksumFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_checksum_get_digest' : parameter 'buffer' has no type

var checksumGetStringFunction *gi.Function
var checksumGetStringFunction_Once sync.Once

func checksumGetStringFunction_Set() error {
	var err error
	checksumGetStringFunction_Once.Do(func() {
		err = checksumStruct_Set()
		if err != nil {
			return
		}
		checksumGetStringFunction, err = checksumStruct.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_checksum_get_string.
func (recv *Checksum) GetString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := checksumGetStringFunction_Set()
	if err == nil {
		ret = checksumGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var checksumResetFunction *gi.Function
var checksumResetFunction_Once sync.Once

func checksumResetFunction_Set() error {
	var err error
	checksumResetFunction_Once.Do(func() {
		err = checksumStruct_Set()
		if err != nil {
			return
		}
		checksumResetFunction, err = checksumStruct.InvokerNew("reset")
	})
	return err
}

// Reset is a representation of the C type g_checksum_reset.
func (recv *Checksum) Reset() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := checksumResetFunction_Set()
	if err == nil {
		checksumResetFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_checksum_update' : parameter 'data' has no type

var condStruct *gi.Struct
var condStruct_Once sync.Once

func condStruct_Set() error {
	var err error
	condStruct_Once.Do(func() {
		condStruct, err = gi.StructNew("GLib", "Cond")
	})
	return err
}

type Cond struct {
	native uintptr
}

var condBroadcastFunction *gi.Function
var condBroadcastFunction_Once sync.Once

func condBroadcastFunction_Set() error {
	var err error
	condBroadcastFunction_Once.Do(func() {
		err = condStruct_Set()
		if err != nil {
			return
		}
		condBroadcastFunction, err = condStruct.InvokerNew("broadcast")
	})
	return err
}

// Broadcast is a representation of the C type g_cond_broadcast.
func (recv *Cond) Broadcast() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := condBroadcastFunction_Set()
	if err == nil {
		condBroadcastFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var condClearFunction *gi.Function
var condClearFunction_Once sync.Once

func condClearFunction_Set() error {
	var err error
	condClearFunction_Once.Do(func() {
		err = condStruct_Set()
		if err != nil {
			return
		}
		condClearFunction, err = condStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_cond_clear.
func (recv *Cond) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := condClearFunction_Set()
	if err == nil {
		condClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var condInitFunction *gi.Function
var condInitFunction_Once sync.Once

func condInitFunction_Set() error {
	var err error
	condInitFunction_Once.Do(func() {
		err = condStruct_Set()
		if err != nil {
			return
		}
		condInitFunction, err = condStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_cond_init.
func (recv *Cond) Init() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := condInitFunction_Set()
	if err == nil {
		condInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var condSignalFunction *gi.Function
var condSignalFunction_Once sync.Once

func condSignalFunction_Set() error {
	var err error
	condSignalFunction_Once.Do(func() {
		err = condStruct_Set()
		if err != nil {
			return
		}
		condSignalFunction, err = condStruct.InvokerNew("signal")
	})
	return err
}

// Signal is a representation of the C type g_cond_signal.
func (recv *Cond) Signal() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := condSignalFunction_Set()
	if err == nil {
		condSignalFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_cond_wait' : parameter 'mutex' of type 'Mutex' not supported

// UNSUPPORTED : C value 'g_cond_wait_until' : parameter 'mutex' of type 'Mutex' not supported

var dataStruct *gi.Struct
var dataStruct_Once sync.Once

func dataStruct_Set() error {
	var err error
	dataStruct_Once.Do(func() {
		dataStruct, err = gi.StructNew("GLib", "Data")
	})
	return err
}

type Data struct {
	native uintptr
}

var dateStruct *gi.Struct
var dateStruct_Once sync.Once

func dateStruct_Set() error {
	var err error
	dateStruct_Once.Do(func() {
		dateStruct, err = gi.StructNew("GLib", "Date")
	})
	return err
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

func dateNewFunction_Set() error {
	var err error
	dateNewFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateNewFunction, err = dateStruct.InvokerNew("new")
	})
	return err
}

// DateNew is a representation of the C type g_date_new.
func DateNew() (*Date, error) {

	var ret gi.Argument

	err := dateNewFunction_Set()
	if err == nil {
		ret = dateNewFunction.Invoke(nil, nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_new_dmy' : parameter 'day' of type 'DateDay' not supported

var dateNewJulianFunction *gi.Function
var dateNewJulianFunction_Once sync.Once

func dateNewJulianFunction_Set() error {
	var err error
	dateNewJulianFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateNewJulianFunction, err = dateStruct.InvokerNew("new_julian")
	})
	return err
}

// DateNewJulian is a representation of the C type g_date_new_julian.
func DateNewJulian(julianDay uint32) (*Date, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(julianDay)

	var ret gi.Argument

	err := dateNewJulianFunction_Set()
	if err == nil {
		ret = dateNewJulianFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo, err
}

var dateAddDaysFunction *gi.Function
var dateAddDaysFunction_Once sync.Once

func dateAddDaysFunction_Set() error {
	var err error
	dateAddDaysFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateAddDaysFunction, err = dateStruct.InvokerNew("add_days")
	})
	return err
}

// AddDays is a representation of the C type g_date_add_days.
func (recv *Date) AddDays(nDays uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDays)

	err := dateAddDaysFunction_Set()
	if err == nil {
		dateAddDaysFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dateAddMonthsFunction *gi.Function
var dateAddMonthsFunction_Once sync.Once

func dateAddMonthsFunction_Set() error {
	var err error
	dateAddMonthsFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateAddMonthsFunction, err = dateStruct.InvokerNew("add_months")
	})
	return err
}

// AddMonths is a representation of the C type g_date_add_months.
func (recv *Date) AddMonths(nMonths uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nMonths)

	err := dateAddMonthsFunction_Set()
	if err == nil {
		dateAddMonthsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dateAddYearsFunction *gi.Function
var dateAddYearsFunction_Once sync.Once

func dateAddYearsFunction_Set() error {
	var err error
	dateAddYearsFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateAddYearsFunction, err = dateStruct.InvokerNew("add_years")
	})
	return err
}

// AddYears is a representation of the C type g_date_add_years.
func (recv *Date) AddYears(nYears uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nYears)

	err := dateAddYearsFunction_Set()
	if err == nil {
		dateAddYearsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_clamp' : parameter 'min_date' of type 'Date' not supported

var dateClearFunction *gi.Function
var dateClearFunction_Once sync.Once

func dateClearFunction_Set() error {
	var err error
	dateClearFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateClearFunction, err = dateStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_date_clear.
func (recv *Date) Clear(nDates uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDates)

	err := dateClearFunction_Set()
	if err == nil {
		dateClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_compare' : parameter 'rhs' of type 'Date' not supported

var dateCopyFunction *gi.Function
var dateCopyFunction_Once sync.Once

func dateCopyFunction_Set() error {
	var err error
	dateCopyFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateCopyFunction, err = dateStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_date_copy.
func (recv *Date) Copy() (*Date, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateCopyFunction_Set()
	if err == nil {
		ret = dateCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Date{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_days_between' : parameter 'date2' of type 'Date' not supported

var dateFreeFunction *gi.Function
var dateFreeFunction_Once sync.Once

func dateFreeFunction_Set() error {
	var err error
	dateFreeFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateFreeFunction, err = dateStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_date_free.
func (recv *Date) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dateFreeFunction_Set()
	if err == nil {
		dateFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_get_day' : return type 'DateDay' not supported

var dateGetDayOfYearFunction *gi.Function
var dateGetDayOfYearFunction_Once sync.Once

func dateGetDayOfYearFunction_Set() error {
	var err error
	dateGetDayOfYearFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetDayOfYearFunction, err = dateStruct.InvokerNew("get_day_of_year")
	})
	return err
}

// GetDayOfYear is a representation of the C type g_date_get_day_of_year.
func (recv *Date) GetDayOfYear() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetDayOfYearFunction_Set()
	if err == nil {
		ret = dateGetDayOfYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var dateGetIso8601WeekOfYearFunction *gi.Function
var dateGetIso8601WeekOfYearFunction_Once sync.Once

func dateGetIso8601WeekOfYearFunction_Set() error {
	var err error
	dateGetIso8601WeekOfYearFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetIso8601WeekOfYearFunction, err = dateStruct.InvokerNew("get_iso8601_week_of_year")
	})
	return err
}

// GetIso8601WeekOfYear is a representation of the C type g_date_get_iso8601_week_of_year.
func (recv *Date) GetIso8601WeekOfYear() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetIso8601WeekOfYearFunction_Set()
	if err == nil {
		ret = dateGetIso8601WeekOfYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var dateGetJulianFunction *gi.Function
var dateGetJulianFunction_Once sync.Once

func dateGetJulianFunction_Set() error {
	var err error
	dateGetJulianFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetJulianFunction, err = dateStruct.InvokerNew("get_julian")
	})
	return err
}

// GetJulian is a representation of the C type g_date_get_julian.
func (recv *Date) GetJulian() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetJulianFunction_Set()
	if err == nil {
		ret = dateGetJulianFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var dateGetMondayWeekOfYearFunction *gi.Function
var dateGetMondayWeekOfYearFunction_Once sync.Once

func dateGetMondayWeekOfYearFunction_Set() error {
	var err error
	dateGetMondayWeekOfYearFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetMondayWeekOfYearFunction, err = dateStruct.InvokerNew("get_monday_week_of_year")
	})
	return err
}

// GetMondayWeekOfYear is a representation of the C type g_date_get_monday_week_of_year.
func (recv *Date) GetMondayWeekOfYear() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetMondayWeekOfYearFunction_Set()
	if err == nil {
		ret = dateGetMondayWeekOfYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_get_month' : return type 'DateMonth' not supported

var dateGetSundayWeekOfYearFunction *gi.Function
var dateGetSundayWeekOfYearFunction_Once sync.Once

func dateGetSundayWeekOfYearFunction_Set() error {
	var err error
	dateGetSundayWeekOfYearFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateGetSundayWeekOfYearFunction, err = dateStruct.InvokerNew("get_sunday_week_of_year")
	})
	return err
}

// GetSundayWeekOfYear is a representation of the C type g_date_get_sunday_week_of_year.
func (recv *Date) GetSundayWeekOfYear() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateGetSundayWeekOfYearFunction_Set()
	if err == nil {
		ret = dateGetSundayWeekOfYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_get_weekday' : return type 'DateWeekday' not supported

// UNSUPPORTED : C value 'g_date_get_year' : return type 'DateYear' not supported

var dateIsFirstOfMonthFunction *gi.Function
var dateIsFirstOfMonthFunction_Once sync.Once

func dateIsFirstOfMonthFunction_Set() error {
	var err error
	dateIsFirstOfMonthFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateIsFirstOfMonthFunction, err = dateStruct.InvokerNew("is_first_of_month")
	})
	return err
}

// IsFirstOfMonth is a representation of the C type g_date_is_first_of_month.
func (recv *Date) IsFirstOfMonth() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateIsFirstOfMonthFunction_Set()
	if err == nil {
		ret = dateIsFirstOfMonthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var dateIsLastOfMonthFunction *gi.Function
var dateIsLastOfMonthFunction_Once sync.Once

func dateIsLastOfMonthFunction_Set() error {
	var err error
	dateIsLastOfMonthFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateIsLastOfMonthFunction, err = dateStruct.InvokerNew("is_last_of_month")
	})
	return err
}

// IsLastOfMonth is a representation of the C type g_date_is_last_of_month.
func (recv *Date) IsLastOfMonth() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateIsLastOfMonthFunction_Set()
	if err == nil {
		ret = dateIsLastOfMonthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_order' : parameter 'date2' of type 'Date' not supported

// UNSUPPORTED : C value 'g_date_set_day' : parameter 'day' of type 'DateDay' not supported

// UNSUPPORTED : C value 'g_date_set_dmy' : parameter 'day' of type 'DateDay' not supported

var dateSetJulianFunction *gi.Function
var dateSetJulianFunction_Once sync.Once

func dateSetJulianFunction_Set() error {
	var err error
	dateSetJulianFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateSetJulianFunction, err = dateStruct.InvokerNew("set_julian")
	})
	return err
}

// SetJulian is a representation of the C type g_date_set_julian.
func (recv *Date) SetJulian(julianDate uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(julianDate)

	err := dateSetJulianFunction_Set()
	if err == nil {
		dateSetJulianFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_set_month' : parameter 'month' of type 'DateMonth' not supported

var dateSetParseFunction *gi.Function
var dateSetParseFunction_Once sync.Once

func dateSetParseFunction_Set() error {
	var err error
	dateSetParseFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateSetParseFunction, err = dateStruct.InvokerNew("set_parse")
	})
	return err
}

// SetParse is a representation of the C type g_date_set_parse.
func (recv *Date) SetParse(str string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(str)

	err := dateSetParseFunction_Set()
	if err == nil {
		dateSetParseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_set_time' : parameter 'time_' of type 'Time' not supported

var dateSetTimeTFunction *gi.Function
var dateSetTimeTFunction_Once sync.Once

func dateSetTimeTFunction_Set() error {
	var err error
	dateSetTimeTFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateSetTimeTFunction, err = dateStruct.InvokerNew("set_time_t")
	})
	return err
}

// SetTimeT is a representation of the C type g_date_set_time_t.
func (recv *Date) SetTimeT(timet int64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(timet)

	err := dateSetTimeTFunction_Set()
	if err == nil {
		dateSetTimeTFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_set_time_val' : parameter 'timeval' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_set_year' : parameter 'year' of type 'DateYear' not supported

var dateSubtractDaysFunction *gi.Function
var dateSubtractDaysFunction_Once sync.Once

func dateSubtractDaysFunction_Set() error {
	var err error
	dateSubtractDaysFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateSubtractDaysFunction, err = dateStruct.InvokerNew("subtract_days")
	})
	return err
}

// SubtractDays is a representation of the C type g_date_subtract_days.
func (recv *Date) SubtractDays(nDays uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nDays)

	err := dateSubtractDaysFunction_Set()
	if err == nil {
		dateSubtractDaysFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dateSubtractMonthsFunction *gi.Function
var dateSubtractMonthsFunction_Once sync.Once

func dateSubtractMonthsFunction_Set() error {
	var err error
	dateSubtractMonthsFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateSubtractMonthsFunction, err = dateStruct.InvokerNew("subtract_months")
	})
	return err
}

// SubtractMonths is a representation of the C type g_date_subtract_months.
func (recv *Date) SubtractMonths(nMonths uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nMonths)

	err := dateSubtractMonthsFunction_Set()
	if err == nil {
		dateSubtractMonthsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var dateSubtractYearsFunction *gi.Function
var dateSubtractYearsFunction_Once sync.Once

func dateSubtractYearsFunction_Set() error {
	var err error
	dateSubtractYearsFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateSubtractYearsFunction, err = dateStruct.InvokerNew("subtract_years")
	})
	return err
}

// SubtractYears is a representation of the C type g_date_subtract_years.
func (recv *Date) SubtractYears(nYears uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nYears)

	err := dateSubtractYearsFunction_Set()
	if err == nil {
		dateSubtractYearsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_date_to_struct_tm' : parameter 'tm' of type 'gpointer' not supported

var dateValidFunction *gi.Function
var dateValidFunction_Once sync.Once

func dateValidFunction_Set() error {
	var err error
	dateValidFunction_Once.Do(func() {
		err = dateStruct_Set()
		if err != nil {
			return
		}
		dateValidFunction, err = dateStruct.InvokerNew("valid")
	})
	return err
}

// Valid is a representation of the C type g_date_valid.
func (recv *Date) Valid() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateValidFunction_Set()
	if err == nil {
		ret = dateValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var dateTimeStruct *gi.Struct
var dateTimeStruct_Once sync.Once

func dateTimeStruct_Set() error {
	var err error
	dateTimeStruct_Once.Do(func() {
		dateTimeStruct, err = gi.StructNew("GLib", "DateTime")
	})
	return err
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

func dateTimeNewFromUnixLocalFunction_Set() error {
	var err error
	dateTimeNewFromUnixLocalFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeNewFromUnixLocalFunction, err = dateTimeStruct.InvokerNew("new_from_unix_local")
	})
	return err
}

// DateTimeNewFromUnixLocal is a representation of the C type g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) (*DateTime, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	var ret gi.Argument

	err := dateTimeNewFromUnixLocalFunction_Set()
	if err == nil {
		ret = dateTimeNewFromUnixLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeNewFromUnixUtcFunction *gi.Function
var dateTimeNewFromUnixUtcFunction_Once sync.Once

func dateTimeNewFromUnixUtcFunction_Set() error {
	var err error
	dateTimeNewFromUnixUtcFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeNewFromUnixUtcFunction, err = dateTimeStruct.InvokerNew("new_from_unix_utc")
	})
	return err
}

// DateTimeNewFromUnixUtc is a representation of the C type g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) (*DateTime, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	var ret gi.Argument

	err := dateTimeNewFromUnixUtcFunction_Set()
	if err == nil {
		ret = dateTimeNewFromUnixUtcFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeNewLocalFunction *gi.Function
var dateTimeNewLocalFunction_Once sync.Once

func dateTimeNewLocalFunction_Set() error {
	var err error
	dateTimeNewLocalFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeNewLocalFunction, err = dateTimeStruct.InvokerNew("new_local")
	})
	return err
}

// DateTimeNewLocal is a representation of the C type g_date_time_new_local.
func DateTimeNewLocal(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (*DateTime, error) {
	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(year)
	inArgs[1].SetInt32(month)
	inArgs[2].SetInt32(day)
	inArgs[3].SetInt32(hour)
	inArgs[4].SetInt32(minute)
	inArgs[5].SetFloat64(seconds)

	var ret gi.Argument

	err := dateTimeNewLocalFunction_Set()
	if err == nil {
		ret = dateTimeNewLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_time_new_now' : parameter 'tz' of type 'TimeZone' not supported

var dateTimeNewNowLocalFunction *gi.Function
var dateTimeNewNowLocalFunction_Once sync.Once

func dateTimeNewNowLocalFunction_Set() error {
	var err error
	dateTimeNewNowLocalFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeNewNowLocalFunction, err = dateTimeStruct.InvokerNew("new_now_local")
	})
	return err
}

// DateTimeNewNowLocal is a representation of the C type g_date_time_new_now_local.
func DateTimeNewNowLocal() (*DateTime, error) {

	var ret gi.Argument

	err := dateTimeNewNowLocalFunction_Set()
	if err == nil {
		ret = dateTimeNewNowLocalFunction.Invoke(nil, nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeNewNowUtcFunction *gi.Function
var dateTimeNewNowUtcFunction_Once sync.Once

func dateTimeNewNowUtcFunction_Set() error {
	var err error
	dateTimeNewNowUtcFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeNewNowUtcFunction, err = dateTimeStruct.InvokerNew("new_now_utc")
	})
	return err
}

// DateTimeNewNowUtc is a representation of the C type g_date_time_new_now_utc.
func DateTimeNewNowUtc() (*DateTime, error) {

	var ret gi.Argument

	err := dateTimeNewNowUtcFunction_Set()
	if err == nil {
		ret = dateTimeNewNowUtcFunction.Invoke(nil, nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeNewUtcFunction *gi.Function
var dateTimeNewUtcFunction_Once sync.Once

func dateTimeNewUtcFunction_Set() error {
	var err error
	dateTimeNewUtcFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeNewUtcFunction, err = dateTimeStruct.InvokerNew("new_utc")
	})
	return err
}

// DateTimeNewUtc is a representation of the C type g_date_time_new_utc.
func DateTimeNewUtc(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (*DateTime, error) {
	var inArgs [6]gi.Argument
	inArgs[0].SetInt32(year)
	inArgs[1].SetInt32(month)
	inArgs[2].SetInt32(day)
	inArgs[3].SetInt32(hour)
	inArgs[4].SetInt32(minute)
	inArgs[5].SetFloat64(seconds)

	var ret gi.Argument

	err := dateTimeNewUtcFunction_Set()
	if err == nil {
		ret = dateTimeNewUtcFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_time_add' : parameter 'timespan' of type 'TimeSpan' not supported

var dateTimeAddDaysFunction *gi.Function
var dateTimeAddDaysFunction_Once sync.Once

func dateTimeAddDaysFunction_Set() error {
	var err error
	dateTimeAddDaysFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddDaysFunction, err = dateTimeStruct.InvokerNew("add_days")
	})
	return err
}

// AddDays is a representation of the C type g_date_time_add_days.
func (recv *DateTime) AddDays(days int32) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(days)

	var ret gi.Argument

	err := dateTimeAddDaysFunction_Set()
	if err == nil {
		ret = dateTimeAddDaysFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddFullFunction *gi.Function
var dateTimeAddFullFunction_Once sync.Once

func dateTimeAddFullFunction_Set() error {
	var err error
	dateTimeAddFullFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddFullFunction, err = dateTimeStruct.InvokerNew("add_full")
	})
	return err
}

// AddFull is a representation of the C type g_date_time_add_full.
func (recv *DateTime) AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) (*DateTime, error) {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(years)
	inArgs[2].SetInt32(months)
	inArgs[3].SetInt32(days)
	inArgs[4].SetInt32(hours)
	inArgs[5].SetInt32(minutes)
	inArgs[6].SetFloat64(seconds)

	var ret gi.Argument

	err := dateTimeAddFullFunction_Set()
	if err == nil {
		ret = dateTimeAddFullFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddHoursFunction *gi.Function
var dateTimeAddHoursFunction_Once sync.Once

func dateTimeAddHoursFunction_Set() error {
	var err error
	dateTimeAddHoursFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddHoursFunction, err = dateTimeStruct.InvokerNew("add_hours")
	})
	return err
}

// AddHours is a representation of the C type g_date_time_add_hours.
func (recv *DateTime) AddHours(hours int32) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(hours)

	var ret gi.Argument

	err := dateTimeAddHoursFunction_Set()
	if err == nil {
		ret = dateTimeAddHoursFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddMinutesFunction *gi.Function
var dateTimeAddMinutesFunction_Once sync.Once

func dateTimeAddMinutesFunction_Set() error {
	var err error
	dateTimeAddMinutesFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddMinutesFunction, err = dateTimeStruct.InvokerNew("add_minutes")
	})
	return err
}

// AddMinutes is a representation of the C type g_date_time_add_minutes.
func (recv *DateTime) AddMinutes(minutes int32) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(minutes)

	var ret gi.Argument

	err := dateTimeAddMinutesFunction_Set()
	if err == nil {
		ret = dateTimeAddMinutesFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddMonthsFunction *gi.Function
var dateTimeAddMonthsFunction_Once sync.Once

func dateTimeAddMonthsFunction_Set() error {
	var err error
	dateTimeAddMonthsFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddMonthsFunction, err = dateTimeStruct.InvokerNew("add_months")
	})
	return err
}

// AddMonths is a representation of the C type g_date_time_add_months.
func (recv *DateTime) AddMonths(months int32) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(months)

	var ret gi.Argument

	err := dateTimeAddMonthsFunction_Set()
	if err == nil {
		ret = dateTimeAddMonthsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddSecondsFunction *gi.Function
var dateTimeAddSecondsFunction_Once sync.Once

func dateTimeAddSecondsFunction_Set() error {
	var err error
	dateTimeAddSecondsFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddSecondsFunction, err = dateTimeStruct.InvokerNew("add_seconds")
	})
	return err
}

// AddSeconds is a representation of the C type g_date_time_add_seconds.
func (recv *DateTime) AddSeconds(seconds float64) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(seconds)

	var ret gi.Argument

	err := dateTimeAddSecondsFunction_Set()
	if err == nil {
		ret = dateTimeAddSecondsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddWeeksFunction *gi.Function
var dateTimeAddWeeksFunction_Once sync.Once

func dateTimeAddWeeksFunction_Set() error {
	var err error
	dateTimeAddWeeksFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddWeeksFunction, err = dateTimeStruct.InvokerNew("add_weeks")
	})
	return err
}

// AddWeeks is a representation of the C type g_date_time_add_weeks.
func (recv *DateTime) AddWeeks(weeks int32) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(weeks)

	var ret gi.Argument

	err := dateTimeAddWeeksFunction_Set()
	if err == nil {
		ret = dateTimeAddWeeksFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeAddYearsFunction *gi.Function
var dateTimeAddYearsFunction_Once sync.Once

func dateTimeAddYearsFunction_Set() error {
	var err error
	dateTimeAddYearsFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeAddYearsFunction, err = dateTimeStruct.InvokerNew("add_years")
	})
	return err
}

// AddYears is a representation of the C type g_date_time_add_years.
func (recv *DateTime) AddYears(years int32) (*DateTime, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(years)

	var ret gi.Argument

	err := dateTimeAddYearsFunction_Set()
	if err == nil {
		ret = dateTimeAddYearsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_time_difference' : parameter 'begin' of type 'DateTime' not supported

var dateTimeFormatFunction *gi.Function
var dateTimeFormatFunction_Once sync.Once

func dateTimeFormatFunction_Set() error {
	var err error
	dateTimeFormatFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeFormatFunction, err = dateTimeStruct.InvokerNew("format")
	})
	return err
}

// Format is a representation of the C type g_date_time_format.
func (recv *DateTime) Format(format string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(format)

	var ret gi.Argument

	err := dateTimeFormatFunction_Set()
	if err == nil {
		ret = dateTimeFormatFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var dateTimeFormatIso8601Function *gi.Function
var dateTimeFormatIso8601Function_Once sync.Once

func dateTimeFormatIso8601Function_Set() error {
	var err error
	dateTimeFormatIso8601Function_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeFormatIso8601Function, err = dateTimeStruct.InvokerNew("format_iso8601")
	})
	return err
}

// FormatIso8601 is a representation of the C type g_date_time_format_iso8601.
func (recv *DateTime) FormatIso8601() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeFormatIso8601Function_Set()
	if err == nil {
		ret = dateTimeFormatIso8601Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var dateTimeGetDayOfMonthFunction *gi.Function
var dateTimeGetDayOfMonthFunction_Once sync.Once

func dateTimeGetDayOfMonthFunction_Set() error {
	var err error
	dateTimeGetDayOfMonthFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetDayOfMonthFunction, err = dateTimeStruct.InvokerNew("get_day_of_month")
	})
	return err
}

// GetDayOfMonth is a representation of the C type g_date_time_get_day_of_month.
func (recv *DateTime) GetDayOfMonth() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetDayOfMonthFunction_Set()
	if err == nil {
		ret = dateTimeGetDayOfMonthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetDayOfWeekFunction *gi.Function
var dateTimeGetDayOfWeekFunction_Once sync.Once

func dateTimeGetDayOfWeekFunction_Set() error {
	var err error
	dateTimeGetDayOfWeekFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetDayOfWeekFunction, err = dateTimeStruct.InvokerNew("get_day_of_week")
	})
	return err
}

// GetDayOfWeek is a representation of the C type g_date_time_get_day_of_week.
func (recv *DateTime) GetDayOfWeek() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetDayOfWeekFunction_Set()
	if err == nil {
		ret = dateTimeGetDayOfWeekFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetDayOfYearFunction *gi.Function
var dateTimeGetDayOfYearFunction_Once sync.Once

func dateTimeGetDayOfYearFunction_Set() error {
	var err error
	dateTimeGetDayOfYearFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetDayOfYearFunction, err = dateTimeStruct.InvokerNew("get_day_of_year")
	})
	return err
}

// GetDayOfYear is a representation of the C type g_date_time_get_day_of_year.
func (recv *DateTime) GetDayOfYear() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetDayOfYearFunction_Set()
	if err == nil {
		ret = dateTimeGetDayOfYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetHourFunction *gi.Function
var dateTimeGetHourFunction_Once sync.Once

func dateTimeGetHourFunction_Set() error {
	var err error
	dateTimeGetHourFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetHourFunction, err = dateTimeStruct.InvokerNew("get_hour")
	})
	return err
}

// GetHour is a representation of the C type g_date_time_get_hour.
func (recv *DateTime) GetHour() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetHourFunction_Set()
	if err == nil {
		ret = dateTimeGetHourFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetMicrosecondFunction *gi.Function
var dateTimeGetMicrosecondFunction_Once sync.Once

func dateTimeGetMicrosecondFunction_Set() error {
	var err error
	dateTimeGetMicrosecondFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetMicrosecondFunction, err = dateTimeStruct.InvokerNew("get_microsecond")
	})
	return err
}

// GetMicrosecond is a representation of the C type g_date_time_get_microsecond.
func (recv *DateTime) GetMicrosecond() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetMicrosecondFunction_Set()
	if err == nil {
		ret = dateTimeGetMicrosecondFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetMinuteFunction *gi.Function
var dateTimeGetMinuteFunction_Once sync.Once

func dateTimeGetMinuteFunction_Set() error {
	var err error
	dateTimeGetMinuteFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetMinuteFunction, err = dateTimeStruct.InvokerNew("get_minute")
	})
	return err
}

// GetMinute is a representation of the C type g_date_time_get_minute.
func (recv *DateTime) GetMinute() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetMinuteFunction_Set()
	if err == nil {
		ret = dateTimeGetMinuteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetMonthFunction *gi.Function
var dateTimeGetMonthFunction_Once sync.Once

func dateTimeGetMonthFunction_Set() error {
	var err error
	dateTimeGetMonthFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetMonthFunction, err = dateTimeStruct.InvokerNew("get_month")
	})
	return err
}

// GetMonth is a representation of the C type g_date_time_get_month.
func (recv *DateTime) GetMonth() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetMonthFunction_Set()
	if err == nil {
		ret = dateTimeGetMonthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetSecondFunction *gi.Function
var dateTimeGetSecondFunction_Once sync.Once

func dateTimeGetSecondFunction_Set() error {
	var err error
	dateTimeGetSecondFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetSecondFunction, err = dateTimeStruct.InvokerNew("get_second")
	})
	return err
}

// GetSecond is a representation of the C type g_date_time_get_second.
func (recv *DateTime) GetSecond() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetSecondFunction_Set()
	if err == nil {
		ret = dateTimeGetSecondFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetSecondsFunction *gi.Function
var dateTimeGetSecondsFunction_Once sync.Once

func dateTimeGetSecondsFunction_Set() error {
	var err error
	dateTimeGetSecondsFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetSecondsFunction, err = dateTimeStruct.InvokerNew("get_seconds")
	})
	return err
}

// GetSeconds is a representation of the C type g_date_time_get_seconds.
func (recv *DateTime) GetSeconds() (float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetSecondsFunction_Set()
	if err == nil {
		ret = dateTimeGetSecondsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var dateTimeGetTimezoneFunction *gi.Function
var dateTimeGetTimezoneFunction_Once sync.Once

func dateTimeGetTimezoneFunction_Set() error {
	var err error
	dateTimeGetTimezoneFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetTimezoneFunction, err = dateTimeStruct.InvokerNew("get_timezone")
	})
	return err
}

// GetTimezone is a representation of the C type g_date_time_get_timezone.
func (recv *DateTime) GetTimezone() (*TimeZone, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetTimezoneFunction_Set()
	if err == nil {
		ret = dateTimeGetTimezoneFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo, err
}

var dateTimeGetTimezoneAbbreviationFunction *gi.Function
var dateTimeGetTimezoneAbbreviationFunction_Once sync.Once

func dateTimeGetTimezoneAbbreviationFunction_Set() error {
	var err error
	dateTimeGetTimezoneAbbreviationFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetTimezoneAbbreviationFunction, err = dateTimeStruct.InvokerNew("get_timezone_abbreviation")
	})
	return err
}

// GetTimezoneAbbreviation is a representation of the C type g_date_time_get_timezone_abbreviation.
func (recv *DateTime) GetTimezoneAbbreviation() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetTimezoneAbbreviationFunction_Set()
	if err == nil {
		ret = dateTimeGetTimezoneAbbreviationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_time_get_utc_offset' : return type 'TimeSpan' not supported

var dateTimeGetWeekNumberingYearFunction *gi.Function
var dateTimeGetWeekNumberingYearFunction_Once sync.Once

func dateTimeGetWeekNumberingYearFunction_Set() error {
	var err error
	dateTimeGetWeekNumberingYearFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetWeekNumberingYearFunction, err = dateTimeStruct.InvokerNew("get_week_numbering_year")
	})
	return err
}

// GetWeekNumberingYear is a representation of the C type g_date_time_get_week_numbering_year.
func (recv *DateTime) GetWeekNumberingYear() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetWeekNumberingYearFunction_Set()
	if err == nil {
		ret = dateTimeGetWeekNumberingYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetWeekOfYearFunction *gi.Function
var dateTimeGetWeekOfYearFunction_Once sync.Once

func dateTimeGetWeekOfYearFunction_Set() error {
	var err error
	dateTimeGetWeekOfYearFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetWeekOfYearFunction, err = dateTimeStruct.InvokerNew("get_week_of_year")
	})
	return err
}

// GetWeekOfYear is a representation of the C type g_date_time_get_week_of_year.
func (recv *DateTime) GetWeekOfYear() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetWeekOfYearFunction_Set()
	if err == nil {
		ret = dateTimeGetWeekOfYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetYearFunction *gi.Function
var dateTimeGetYearFunction_Once sync.Once

func dateTimeGetYearFunction_Set() error {
	var err error
	dateTimeGetYearFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetYearFunction, err = dateTimeStruct.InvokerNew("get_year")
	})
	return err
}

// GetYear is a representation of the C type g_date_time_get_year.
func (recv *DateTime) GetYear() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeGetYearFunction_Set()
	if err == nil {
		ret = dateTimeGetYearFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var dateTimeGetYmdFunction *gi.Function
var dateTimeGetYmdFunction_Once sync.Once

func dateTimeGetYmdFunction_Set() error {
	var err error
	dateTimeGetYmdFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeGetYmdFunction, err = dateTimeStruct.InvokerNew("get_ymd")
	})
	return err
}

// GetYmd is a representation of the C type g_date_time_get_ymd.
func (recv *DateTime) GetYmd() (int32, int32, int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	err := dateTimeGetYmdFunction_Set()
	if err == nil {
		dateTimeGetYmdFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2, err
}

var dateTimeIsDaylightSavingsFunction *gi.Function
var dateTimeIsDaylightSavingsFunction_Once sync.Once

func dateTimeIsDaylightSavingsFunction_Set() error {
	var err error
	dateTimeIsDaylightSavingsFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeIsDaylightSavingsFunction, err = dateTimeStruct.InvokerNew("is_daylight_savings")
	})
	return err
}

// IsDaylightSavings is a representation of the C type g_date_time_is_daylight_savings.
func (recv *DateTime) IsDaylightSavings() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeIsDaylightSavingsFunction_Set()
	if err == nil {
		ret = dateTimeIsDaylightSavingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var dateTimeRefFunction *gi.Function
var dateTimeRefFunction_Once sync.Once

func dateTimeRefFunction_Set() error {
	var err error
	dateTimeRefFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeRefFunction, err = dateTimeStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_date_time_ref.
func (recv *DateTime) Ref() (*DateTime, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeRefFunction_Set()
	if err == nil {
		ret = dateTimeRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeToLocalFunction *gi.Function
var dateTimeToLocalFunction_Once sync.Once

func dateTimeToLocalFunction_Set() error {
	var err error
	dateTimeToLocalFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeToLocalFunction, err = dateTimeStruct.InvokerNew("to_local")
	})
	return err
}

// ToLocal is a representation of the C type g_date_time_to_local.
func (recv *DateTime) ToLocal() (*DateTime, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeToLocalFunction_Set()
	if err == nil {
		ret = dateTimeToLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_date_time_to_timeval' : parameter 'tv' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_time_to_timezone' : parameter 'tz' of type 'TimeZone' not supported

var dateTimeToUnixFunction *gi.Function
var dateTimeToUnixFunction_Once sync.Once

func dateTimeToUnixFunction_Set() error {
	var err error
	dateTimeToUnixFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeToUnixFunction, err = dateTimeStruct.InvokerNew("to_unix")
	})
	return err
}

// ToUnix is a representation of the C type g_date_time_to_unix.
func (recv *DateTime) ToUnix() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeToUnixFunction_Set()
	if err == nil {
		ret = dateTimeToUnixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var dateTimeToUtcFunction *gi.Function
var dateTimeToUtcFunction_Once sync.Once

func dateTimeToUtcFunction_Set() error {
	var err error
	dateTimeToUtcFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeToUtcFunction, err = dateTimeStruct.InvokerNew("to_utc")
	})
	return err
}

// ToUtc is a representation of the C type g_date_time_to_utc.
func (recv *DateTime) ToUtc() (*DateTime, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := dateTimeToUtcFunction_Set()
	if err == nil {
		ret = dateTimeToUtcFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DateTime{native: ret.Pointer()}

	return retGo, err
}

var dateTimeUnrefFunction *gi.Function
var dateTimeUnrefFunction_Once sync.Once

func dateTimeUnrefFunction_Set() error {
	var err error
	dateTimeUnrefFunction_Once.Do(func() {
		err = dateTimeStruct_Set()
		if err != nil {
			return
		}
		dateTimeUnrefFunction, err = dateTimeStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_date_time_unref.
func (recv *DateTime) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dateTimeUnrefFunction_Set()
	if err == nil {
		dateTimeUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var debugKeyStruct *gi.Struct
var debugKeyStruct_Once sync.Once

func debugKeyStruct_Set() error {
	var err error
	debugKeyStruct_Once.Do(func() {
		debugKeyStruct, err = gi.StructNew("GLib", "DebugKey")
	})
	return err
}

type DebugKey struct {
	native uintptr
	Key    string
	Value  uint32
}

var dirStruct *gi.Struct
var dirStruct_Once sync.Once

func dirStruct_Set() error {
	var err error
	dirStruct_Once.Do(func() {
		dirStruct, err = gi.StructNew("GLib", "Dir")
	})
	return err
}

type Dir struct {
	native uintptr
}

var dirCloseFunction *gi.Function
var dirCloseFunction_Once sync.Once

func dirCloseFunction_Set() error {
	var err error
	dirCloseFunction_Once.Do(func() {
		err = dirStruct_Set()
		if err != nil {
			return
		}
		dirCloseFunction, err = dirStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_dir_close.
func (recv *Dir) Close() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dirCloseFunction_Set()
	if err == nil {
		dirCloseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_dir_read_name' : return type 'filename' not supported

var dirRewindFunction *gi.Function
var dirRewindFunction_Once sync.Once

func dirRewindFunction_Set() error {
	var err error
	dirRewindFunction_Once.Do(func() {
		err = dirStruct_Set()
		if err != nil {
			return
		}
		dirRewindFunction, err = dirStruct.InvokerNew("rewind")
	})
	return err
}

// Rewind is a representation of the C type g_dir_rewind.
func (recv *Dir) Rewind() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := dirRewindFunction_Set()
	if err == nil {
		dirRewindFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var errorStruct *gi.Struct
var errorStruct_Once sync.Once

func errorStruct_Set() error {
	var err error
	errorStruct_Once.Do(func() {
		errorStruct, err = gi.StructNew("GLib", "Error")
	})
	return err
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

func errorCopyFunction_Set() error {
	var err error
	errorCopyFunction_Once.Do(func() {
		err = errorStruct_Set()
		if err != nil {
			return
		}
		errorCopyFunction, err = errorStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_error_copy.
func (recv *Error) Copy() (*Error, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := errorCopyFunction_Set()
	if err == nil {
		ret = errorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Error{native: ret.Pointer()}

	return retGo, err
}

var errorFreeFunction *gi.Function
var errorFreeFunction_Once sync.Once

func errorFreeFunction_Set() error {
	var err error
	errorFreeFunction_Once.Do(func() {
		err = errorStruct_Set()
		if err != nil {
			return
		}
		errorFreeFunction, err = errorStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_error_free.
func (recv *Error) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := errorFreeFunction_Set()
	if err == nil {
		errorFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_error_matches' : parameter 'domain' of type 'Quark' not supported

var hashTableStruct *gi.Struct
var hashTableStruct_Once sync.Once

func hashTableStruct_Set() error {
	var err error
	hashTableStruct_Once.Do(func() {
		hashTableStruct, err = gi.StructNew("GLib", "HashTable")
	})
	return err
}

type HashTable struct {
	native uintptr
}

var hashTableIterStruct *gi.Struct
var hashTableIterStruct_Once sync.Once

func hashTableIterStruct_Set() error {
	var err error
	hashTableIterStruct_Once.Do(func() {
		hashTableIterStruct, err = gi.StructNew("GLib", "HashTableIter")
	})
	return err
}

type HashTableIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_hash_table_iter_get_hash_table' : return type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_iter_init' : parameter 'hash_table' of type 'GLib.HashTable' not supported

// UNSUPPORTED : C value 'g_hash_table_iter_next' : parameter 'key' of type 'gpointer' not supported

var hashTableIterRemoveFunction *gi.Function
var hashTableIterRemoveFunction_Once sync.Once

func hashTableIterRemoveFunction_Set() error {
	var err error
	hashTableIterRemoveFunction_Once.Do(func() {
		err = hashTableIterStruct_Set()
		if err != nil {
			return
		}
		hashTableIterRemoveFunction, err = hashTableIterStruct.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type g_hash_table_iter_remove.
func (recv *HashTableIter) Remove() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := hashTableIterRemoveFunction_Set()
	if err == nil {
		hashTableIterRemoveFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_hash_table_iter_replace' : parameter 'value' of type 'gpointer' not supported

var hashTableIterStealFunction *gi.Function
var hashTableIterStealFunction_Once sync.Once

func hashTableIterStealFunction_Set() error {
	var err error
	hashTableIterStealFunction_Once.Do(func() {
		err = hashTableIterStruct_Set()
		if err != nil {
			return
		}
		hashTableIterStealFunction, err = hashTableIterStruct.InvokerNew("steal")
	})
	return err
}

// Steal is a representation of the C type g_hash_table_iter_steal.
func (recv *HashTableIter) Steal() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := hashTableIterStealFunction_Set()
	if err == nil {
		hashTableIterStealFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var hmacStruct *gi.Struct
var hmacStruct_Once sync.Once

func hmacStruct_Set() error {
	var err error
	hmacStruct_Once.Do(func() {
		hmacStruct, err = gi.StructNew("GLib", "Hmac")
	})
	return err
}

type Hmac struct {
	native uintptr
}

var hmacCopyFunction *gi.Function
var hmacCopyFunction_Once sync.Once

func hmacCopyFunction_Set() error {
	var err error
	hmacCopyFunction_Once.Do(func() {
		err = hmacStruct_Set()
		if err != nil {
			return
		}
		hmacCopyFunction, err = hmacStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_hmac_copy.
func (recv *Hmac) Copy() (*Hmac, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hmacCopyFunction_Set()
	if err == nil {
		ret = hmacCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Hmac{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_hmac_get_digest' : parameter 'buffer' has no type

var hmacGetStringFunction *gi.Function
var hmacGetStringFunction_Once sync.Once

func hmacGetStringFunction_Set() error {
	var err error
	hmacGetStringFunction_Once.Do(func() {
		err = hmacStruct_Set()
		if err != nil {
			return
		}
		hmacGetStringFunction, err = hmacStruct.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_hmac_get_string.
func (recv *Hmac) GetString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hmacGetStringFunction_Set()
	if err == nil {
		ret = hmacGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var hmacRefFunction *gi.Function
var hmacRefFunction_Once sync.Once

func hmacRefFunction_Set() error {
	var err error
	hmacRefFunction_Once.Do(func() {
		err = hmacStruct_Set()
		if err != nil {
			return
		}
		hmacRefFunction, err = hmacStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_hmac_ref.
func (recv *Hmac) Ref() (*Hmac, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hmacRefFunction_Set()
	if err == nil {
		ret = hmacRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Hmac{native: ret.Pointer()}

	return retGo, err
}

var hmacUnrefFunction *gi.Function
var hmacUnrefFunction_Once sync.Once

func hmacUnrefFunction_Set() error {
	var err error
	hmacUnrefFunction_Once.Do(func() {
		err = hmacStruct_Set()
		if err != nil {
			return
		}
		hmacUnrefFunction, err = hmacStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_hmac_unref.
func (recv *Hmac) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := hmacUnrefFunction_Set()
	if err == nil {
		hmacUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_hmac_update' : parameter 'data' has no type

var hookStruct *gi.Struct
var hookStruct_Once sync.Once

func hookStruct_Set() error {
	var err error
	hookStruct_Once.Do(func() {
		hookStruct, err = gi.StructNew("GLib", "Hook")
	})
	return err
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
var hookListStruct_Once sync.Once

func hookListStruct_Set() error {
	var err error
	hookListStruct_Once.Do(func() {
		hookListStruct, err = gi.StructNew("GLib", "HookList")
	})
	return err
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

func hookListClearFunction_Set() error {
	var err error
	hookListClearFunction_Once.Do(func() {
		err = hookListStruct_Set()
		if err != nil {
			return
		}
		hookListClearFunction, err = hookListStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_hook_list_clear.
func (recv *HookList) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := hookListClearFunction_Set()
	if err == nil {
		hookListClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var hookListInitFunction *gi.Function
var hookListInitFunction_Once sync.Once

func hookListInitFunction_Set() error {
	var err error
	hookListInitFunction_Once.Do(func() {
		err = hookListStruct_Set()
		if err != nil {
			return
		}
		hookListInitFunction, err = hookListStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_hook_list_init.
func (recv *HookList) Init(hookSize uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(hookSize)

	err := hookListInitFunction_Set()
	if err == nil {
		hookListInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var hookListInvokeFunction *gi.Function
var hookListInvokeFunction_Once sync.Once

func hookListInvokeFunction_Set() error {
	var err error
	hookListInvokeFunction_Once.Do(func() {
		err = hookListStruct_Set()
		if err != nil {
			return
		}
		hookListInvokeFunction, err = hookListStruct.InvokerNew("invoke")
	})
	return err
}

// Invoke is a representation of the C type g_hook_list_invoke.
func (recv *HookList) Invoke(mayRecurse bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(mayRecurse)

	err := hookListInvokeFunction_Set()
	if err == nil {
		hookListInvokeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var hookListInvokeCheckFunction *gi.Function
var hookListInvokeCheckFunction_Once sync.Once

func hookListInvokeCheckFunction_Set() error {
	var err error
	hookListInvokeCheckFunction_Once.Do(func() {
		err = hookListStruct_Set()
		if err != nil {
			return
		}
		hookListInvokeCheckFunction, err = hookListStruct.InvokerNew("invoke_check")
	})
	return err
}

// InvokeCheck is a representation of the C type g_hook_list_invoke_check.
func (recv *HookList) InvokeCheck(mayRecurse bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(mayRecurse)

	err := hookListInvokeCheckFunction_Set()
	if err == nil {
		hookListInvokeCheckFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_hook_list_marshal' : parameter 'marshaller' of type 'HookMarshaller' not supported

// UNSUPPORTED : C value 'g_hook_list_marshal_check' : parameter 'marshaller' of type 'HookCheckMarshaller' not supported

var iConvStruct *gi.Struct
var iConvStruct_Once sync.Once

func iConvStruct_Set() error {
	var err error
	iConvStruct_Once.Do(func() {
		iConvStruct, err = gi.StructNew("GLib", "IConv")
	})
	return err
}

type IConv struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_iconv' : parameter 'inbytes_left' of type 'gsize' not supported

var iConvCloseFunction *gi.Function
var iConvCloseFunction_Once sync.Once

func iConvCloseFunction_Set() error {
	var err error
	iConvCloseFunction_Once.Do(func() {
		err = iConvStruct_Set()
		if err != nil {
			return
		}
		iConvCloseFunction, err = iConvStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_iconv_close.
func (recv *IConv) Close() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iConvCloseFunction_Set()
	if err == nil {
		ret = iConvCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var iOChannelStruct *gi.Struct
var iOChannelStruct_Once sync.Once

func iOChannelStruct_Set() error {
	var err error
	iOChannelStruct_Once.Do(func() {
		iOChannelStruct, err = gi.StructNew("GLib", "IOChannel")
	})
	return err
}

type IOChannel struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_channel_new_file' : parameter 'filename' of type 'filename' not supported

var iOChannelUnixNewFunction *gi.Function
var iOChannelUnixNewFunction_Once sync.Once

func iOChannelUnixNewFunction_Set() error {
	var err error
	iOChannelUnixNewFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelUnixNewFunction, err = iOChannelStruct.InvokerNew("unix_new")
	})
	return err
}

// IOChannelUnixNew is a representation of the C type g_io_channel_unix_new.
func IOChannelUnixNew(fd int32) (*IOChannel, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	var ret gi.Argument

	err := iOChannelUnixNewFunction_Set()
	if err == nil {
		ret = iOChannelUnixNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOChannel{native: ret.Pointer()}

	return retGo, err
}

var iOChannelCloseFunction *gi.Function
var iOChannelCloseFunction_Once sync.Once

func iOChannelCloseFunction_Set() error {
	var err error
	iOChannelCloseFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelCloseFunction, err = iOChannelStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_io_channel_close.
func (recv *IOChannel) Close() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iOChannelCloseFunction_Set()
	if err == nil {
		iOChannelCloseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_io_channel_flush' : return type 'IOStatus' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffer_condition' : return type 'IOCondition' not supported

// UNSUPPORTED : C value 'g_io_channel_get_buffer_size' : return type 'gsize' not supported

var iOChannelGetBufferedFunction *gi.Function
var iOChannelGetBufferedFunction_Once sync.Once

func iOChannelGetBufferedFunction_Set() error {
	var err error
	iOChannelGetBufferedFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelGetBufferedFunction, err = iOChannelStruct.InvokerNew("get_buffered")
	})
	return err
}

// GetBuffered is a representation of the C type g_io_channel_get_buffered.
func (recv *IOChannel) GetBuffered() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOChannelGetBufferedFunction_Set()
	if err == nil {
		ret = iOChannelGetBufferedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var iOChannelGetCloseOnUnrefFunction *gi.Function
var iOChannelGetCloseOnUnrefFunction_Once sync.Once

func iOChannelGetCloseOnUnrefFunction_Set() error {
	var err error
	iOChannelGetCloseOnUnrefFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelGetCloseOnUnrefFunction, err = iOChannelStruct.InvokerNew("get_close_on_unref")
	})
	return err
}

// GetCloseOnUnref is a representation of the C type g_io_channel_get_close_on_unref.
func (recv *IOChannel) GetCloseOnUnref() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOChannelGetCloseOnUnrefFunction_Set()
	if err == nil {
		ret = iOChannelGetCloseOnUnrefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var iOChannelGetEncodingFunction *gi.Function
var iOChannelGetEncodingFunction_Once sync.Once

func iOChannelGetEncodingFunction_Set() error {
	var err error
	iOChannelGetEncodingFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelGetEncodingFunction, err = iOChannelStruct.InvokerNew("get_encoding")
	})
	return err
}

// GetEncoding is a representation of the C type g_io_channel_get_encoding.
func (recv *IOChannel) GetEncoding() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOChannelGetEncodingFunction_Set()
	if err == nil {
		ret = iOChannelGetEncodingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_io_channel_get_flags' : return type 'IOFlags' not supported

var iOChannelGetLineTermFunction *gi.Function
var iOChannelGetLineTermFunction_Once sync.Once

func iOChannelGetLineTermFunction_Set() error {
	var err error
	iOChannelGetLineTermFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelGetLineTermFunction, err = iOChannelStruct.InvokerNew("get_line_term")
	})
	return err
}

// GetLineTerm is a representation of the C type g_io_channel_get_line_term.
func (recv *IOChannel) GetLineTerm(length int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(length)

	var ret gi.Argument

	err := iOChannelGetLineTermFunction_Set()
	if err == nil {
		ret = iOChannelGetLineTermFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var iOChannelInitFunction *gi.Function
var iOChannelInitFunction_Once sync.Once

func iOChannelInitFunction_Set() error {
	var err error
	iOChannelInitFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelInitFunction, err = iOChannelStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_io_channel_init.
func (recv *IOChannel) Init() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iOChannelInitFunction_Set()
	if err == nil {
		iOChannelInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_io_channel_read' : parameter 'count' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_read_chars' : parameter 'buf' has no type

// UNSUPPORTED : C value 'g_io_channel_read_line' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_read_line_string' : parameter 'buffer' of type 'String' not supported

// UNSUPPORTED : C value 'g_io_channel_read_to_end' : parameter 'str_return' has no type

// UNSUPPORTED : C value 'g_io_channel_read_unichar' : parameter 'thechar' of type 'gunichar' not supported

var iOChannelRefFunction *gi.Function
var iOChannelRefFunction_Once sync.Once

func iOChannelRefFunction_Set() error {
	var err error
	iOChannelRefFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelRefFunction, err = iOChannelStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_io_channel_ref.
func (recv *IOChannel) Ref() (*IOChannel, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOChannelRefFunction_Set()
	if err == nil {
		ret = iOChannelRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IOChannel{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_io_channel_seek' : parameter 'type' of type 'SeekType' not supported

// UNSUPPORTED : C value 'g_io_channel_seek_position' : parameter 'type' of type 'SeekType' not supported

// UNSUPPORTED : C value 'g_io_channel_set_buffer_size' : parameter 'size' of type 'gsize' not supported

var iOChannelSetBufferedFunction *gi.Function
var iOChannelSetBufferedFunction_Once sync.Once

func iOChannelSetBufferedFunction_Set() error {
	var err error
	iOChannelSetBufferedFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelSetBufferedFunction, err = iOChannelStruct.InvokerNew("set_buffered")
	})
	return err
}

// SetBuffered is a representation of the C type g_io_channel_set_buffered.
func (recv *IOChannel) SetBuffered(buffered bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(buffered)

	err := iOChannelSetBufferedFunction_Set()
	if err == nil {
		iOChannelSetBufferedFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var iOChannelSetCloseOnUnrefFunction *gi.Function
var iOChannelSetCloseOnUnrefFunction_Once sync.Once

func iOChannelSetCloseOnUnrefFunction_Set() error {
	var err error
	iOChannelSetCloseOnUnrefFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelSetCloseOnUnrefFunction, err = iOChannelStruct.InvokerNew("set_close_on_unref")
	})
	return err
}

// SetCloseOnUnref is a representation of the C type g_io_channel_set_close_on_unref.
func (recv *IOChannel) SetCloseOnUnref(doClose bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(doClose)

	err := iOChannelSetCloseOnUnrefFunction_Set()
	if err == nil {
		iOChannelSetCloseOnUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_io_channel_set_encoding' : return type 'IOStatus' not supported

// UNSUPPORTED : C value 'g_io_channel_set_flags' : parameter 'flags' of type 'IOFlags' not supported

var iOChannelSetLineTermFunction *gi.Function
var iOChannelSetLineTermFunction_Once sync.Once

func iOChannelSetLineTermFunction_Set() error {
	var err error
	iOChannelSetLineTermFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelSetLineTermFunction, err = iOChannelStruct.InvokerNew("set_line_term")
	})
	return err
}

// SetLineTerm is a representation of the C type g_io_channel_set_line_term.
func (recv *IOChannel) SetLineTerm(lineTerm string, length int32) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(lineTerm)
	inArgs[2].SetInt32(length)

	err := iOChannelSetLineTermFunction_Set()
	if err == nil {
		iOChannelSetLineTermFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_io_channel_shutdown' : return type 'IOStatus' not supported

var iOChannelUnixGetFdFunction *gi.Function
var iOChannelUnixGetFdFunction_Once sync.Once

func iOChannelUnixGetFdFunction_Set() error {
	var err error
	iOChannelUnixGetFdFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelUnixGetFdFunction, err = iOChannelStruct.InvokerNew("unix_get_fd")
	})
	return err
}

// UnixGetFd is a representation of the C type g_io_channel_unix_get_fd.
func (recv *IOChannel) UnixGetFd() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iOChannelUnixGetFdFunction_Set()
	if err == nil {
		ret = iOChannelUnixGetFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var iOChannelUnrefFunction *gi.Function
var iOChannelUnrefFunction_Once sync.Once

func iOChannelUnrefFunction_Set() error {
	var err error
	iOChannelUnrefFunction_Once.Do(func() {
		err = iOChannelStruct_Set()
		if err != nil {
			return
		}
		iOChannelUnrefFunction, err = iOChannelStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_io_channel_unref.
func (recv *IOChannel) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iOChannelUnrefFunction_Set()
	if err == nil {
		iOChannelUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_io_channel_write' : parameter 'count' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_io_channel_write_chars' : parameter 'buf' has no type

// UNSUPPORTED : C value 'g_io_channel_write_unichar' : parameter 'thechar' of type 'gunichar' not supported

var iOFuncsStruct *gi.Struct
var iOFuncsStruct_Once sync.Once

func iOFuncsStruct_Set() error {
	var err error
	iOFuncsStruct_Once.Do(func() {
		iOFuncsStruct, err = gi.StructNew("GLib", "IOFuncs")
	})
	return err
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
var keyFileStruct_Once sync.Once

func keyFileStruct_Set() error {
	var err error
	keyFileStruct_Once.Do(func() {
		keyFileStruct, err = gi.StructNew("GLib", "KeyFile")
	})
	return err
}

type KeyFile struct {
	native uintptr
}

var keyFileNewFunction *gi.Function
var keyFileNewFunction_Once sync.Once

func keyFileNewFunction_Set() error {
	var err error
	keyFileNewFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileNewFunction, err = keyFileStruct.InvokerNew("new")
	})
	return err
}

// KeyFileNew is a representation of the C type g_key_file_new.
func KeyFileNew() (*KeyFile, error) {

	var ret gi.Argument

	err := keyFileNewFunction_Set()
	if err == nil {
		ret = keyFileNewFunction.Invoke(nil, nil)
	}

	retGo := &KeyFile{native: ret.Pointer()}

	return retGo, err
}

var keyFileFreeFunction *gi.Function
var keyFileFreeFunction_Once sync.Once

func keyFileFreeFunction_Set() error {
	var err error
	keyFileFreeFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileFreeFunction, err = keyFileStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_key_file_free.
func (recv *KeyFile) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := keyFileFreeFunction_Set()
	if err == nil {
		keyFileFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var keyFileGetBooleanFunction *gi.Function
var keyFileGetBooleanFunction_Once sync.Once

func keyFileGetBooleanFunction_Set() error {
	var err error
	keyFileGetBooleanFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetBooleanFunction, err = keyFileStruct.InvokerNew("get_boolean")
	})
	return err
}

// GetBoolean is a representation of the C type g_key_file_get_boolean.
func (recv *KeyFile) GetBoolean(groupName string, key string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetBooleanFunction_Set()
	if err == nil {
		ret = keyFileGetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_key_file_get_boolean_list' : parameter 'length' of type 'gsize' not supported

var keyFileGetCommentFunction *gi.Function
var keyFileGetCommentFunction_Once sync.Once

func keyFileGetCommentFunction_Set() error {
	var err error
	keyFileGetCommentFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetCommentFunction, err = keyFileStruct.InvokerNew("get_comment")
	})
	return err
}

// GetComment is a representation of the C type g_key_file_get_comment.
func (recv *KeyFile) GetComment(groupName string, key string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetCommentFunction_Set()
	if err == nil {
		ret = keyFileGetCommentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var keyFileGetDoubleFunction *gi.Function
var keyFileGetDoubleFunction_Once sync.Once

func keyFileGetDoubleFunction_Set() error {
	var err error
	keyFileGetDoubleFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetDoubleFunction, err = keyFileStruct.InvokerNew("get_double")
	})
	return err
}

// GetDouble is a representation of the C type g_key_file_get_double.
func (recv *KeyFile) GetDouble(groupName string, key string) (float64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetDoubleFunction_Set()
	if err == nil {
		ret = keyFileGetDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

// UNSUPPORTED : C value 'g_key_file_get_double_list' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_get_groups' : parameter 'length' of type 'gsize' not supported

var keyFileGetInt64Function *gi.Function
var keyFileGetInt64Function_Once sync.Once

func keyFileGetInt64Function_Set() error {
	var err error
	keyFileGetInt64Function_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetInt64Function, err = keyFileStruct.InvokerNew("get_int64")
	})
	return err
}

// GetInt64 is a representation of the C type g_key_file_get_int64.
func (recv *KeyFile) GetInt64(groupName string, key string) (int64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetInt64Function_Set()
	if err == nil {
		ret = keyFileGetInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var keyFileGetIntegerFunction *gi.Function
var keyFileGetIntegerFunction_Once sync.Once

func keyFileGetIntegerFunction_Set() error {
	var err error
	keyFileGetIntegerFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetIntegerFunction, err = keyFileStruct.InvokerNew("get_integer")
	})
	return err
}

// GetInteger is a representation of the C type g_key_file_get_integer.
func (recv *KeyFile) GetInteger(groupName string, key string) (int32, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetIntegerFunction_Set()
	if err == nil {
		ret = keyFileGetIntegerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_key_file_get_integer_list' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_get_keys' : parameter 'length' of type 'gsize' not supported

var keyFileGetLocaleForKeyFunction *gi.Function
var keyFileGetLocaleForKeyFunction_Once sync.Once

func keyFileGetLocaleForKeyFunction_Set() error {
	var err error
	keyFileGetLocaleForKeyFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetLocaleForKeyFunction, err = keyFileStruct.InvokerNew("get_locale_for_key")
	})
	return err
}

// GetLocaleForKey is a representation of the C type g_key_file_get_locale_for_key.
func (recv *KeyFile) GetLocaleForKey(groupName string, key string, locale string) (string, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)

	var ret gi.Argument

	err := keyFileGetLocaleForKeyFunction_Set()
	if err == nil {
		ret = keyFileGetLocaleForKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var keyFileGetLocaleStringFunction *gi.Function
var keyFileGetLocaleStringFunction_Once sync.Once

func keyFileGetLocaleStringFunction_Set() error {
	var err error
	keyFileGetLocaleStringFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetLocaleStringFunction, err = keyFileStruct.InvokerNew("get_locale_string")
	})
	return err
}

// GetLocaleString is a representation of the C type g_key_file_get_locale_string.
func (recv *KeyFile) GetLocaleString(groupName string, key string, locale string) (string, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)

	var ret gi.Argument

	err := keyFileGetLocaleStringFunction_Set()
	if err == nil {
		ret = keyFileGetLocaleStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_key_file_get_locale_string_list' : parameter 'length' of type 'gsize' not supported

var keyFileGetStartGroupFunction *gi.Function
var keyFileGetStartGroupFunction_Once sync.Once

func keyFileGetStartGroupFunction_Set() error {
	var err error
	keyFileGetStartGroupFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetStartGroupFunction, err = keyFileStruct.InvokerNew("get_start_group")
	})
	return err
}

// GetStartGroup is a representation of the C type g_key_file_get_start_group.
func (recv *KeyFile) GetStartGroup() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := keyFileGetStartGroupFunction_Set()
	if err == nil {
		ret = keyFileGetStartGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var keyFileGetStringFunction *gi.Function
var keyFileGetStringFunction_Once sync.Once

func keyFileGetStringFunction_Set() error {
	var err error
	keyFileGetStringFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetStringFunction, err = keyFileStruct.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_key_file_get_string.
func (recv *KeyFile) GetString(groupName string, key string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetStringFunction_Set()
	if err == nil {
		ret = keyFileGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_key_file_get_string_list' : parameter 'length' of type 'gsize' not supported

var keyFileGetUint64Function *gi.Function
var keyFileGetUint64Function_Once sync.Once

func keyFileGetUint64Function_Set() error {
	var err error
	keyFileGetUint64Function_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetUint64Function, err = keyFileStruct.InvokerNew("get_uint64")
	})
	return err
}

// GetUint64 is a representation of the C type g_key_file_get_uint64.
func (recv *KeyFile) GetUint64(groupName string, key string) (uint64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetUint64Function_Set()
	if err == nil {
		ret = keyFileGetUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo, err
}

var keyFileGetValueFunction *gi.Function
var keyFileGetValueFunction_Once sync.Once

func keyFileGetValueFunction_Set() error {
	var err error
	keyFileGetValueFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileGetValueFunction, err = keyFileStruct.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type g_key_file_get_value.
func (recv *KeyFile) GetValue(groupName string, key string) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileGetValueFunction_Set()
	if err == nil {
		ret = keyFileGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var keyFileHasGroupFunction *gi.Function
var keyFileHasGroupFunction_Once sync.Once

func keyFileHasGroupFunction_Set() error {
	var err error
	keyFileHasGroupFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileHasGroupFunction, err = keyFileStruct.InvokerNew("has_group")
	})
	return err
}

// HasGroup is a representation of the C type g_key_file_has_group.
func (recv *KeyFile) HasGroup(groupName string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)

	var ret gi.Argument

	err := keyFileHasGroupFunction_Set()
	if err == nil {
		ret = keyFileHasGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var keyFileHasKeyFunction *gi.Function
var keyFileHasKeyFunction_Once sync.Once

func keyFileHasKeyFunction_Set() error {
	var err error
	keyFileHasKeyFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileHasKeyFunction, err = keyFileStruct.InvokerNew("has_key")
	})
	return err
}

// HasKey is a representation of the C type g_key_file_has_key.
func (recv *KeyFile) HasKey(groupName string, key string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileHasKeyFunction_Set()
	if err == nil {
		ret = keyFileHasKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_key_file_load_from_bytes' : parameter 'bytes' of type 'Bytes' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_data' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_data_dirs' : parameter 'file' of type 'filename' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_dirs' : parameter 'file' of type 'filename' not supported

// UNSUPPORTED : C value 'g_key_file_load_from_file' : parameter 'file' of type 'filename' not supported

var keyFileRefFunction *gi.Function
var keyFileRefFunction_Once sync.Once

func keyFileRefFunction_Set() error {
	var err error
	keyFileRefFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileRefFunction, err = keyFileStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_key_file_ref.
func (recv *KeyFile) Ref() (*KeyFile, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := keyFileRefFunction_Set()
	if err == nil {
		ret = keyFileRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &KeyFile{native: ret.Pointer()}

	return retGo, err
}

var keyFileRemoveCommentFunction *gi.Function
var keyFileRemoveCommentFunction_Once sync.Once

func keyFileRemoveCommentFunction_Set() error {
	var err error
	keyFileRemoveCommentFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileRemoveCommentFunction, err = keyFileStruct.InvokerNew("remove_comment")
	})
	return err
}

// RemoveComment is a representation of the C type g_key_file_remove_comment.
func (recv *KeyFile) RemoveComment(groupName string, key string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileRemoveCommentFunction_Set()
	if err == nil {
		ret = keyFileRemoveCommentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var keyFileRemoveGroupFunction *gi.Function
var keyFileRemoveGroupFunction_Once sync.Once

func keyFileRemoveGroupFunction_Set() error {
	var err error
	keyFileRemoveGroupFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileRemoveGroupFunction, err = keyFileStruct.InvokerNew("remove_group")
	})
	return err
}

// RemoveGroup is a representation of the C type g_key_file_remove_group.
func (recv *KeyFile) RemoveGroup(groupName string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)

	var ret gi.Argument

	err := keyFileRemoveGroupFunction_Set()
	if err == nil {
		ret = keyFileRemoveGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var keyFileRemoveKeyFunction *gi.Function
var keyFileRemoveKeyFunction_Once sync.Once

func keyFileRemoveKeyFunction_Set() error {
	var err error
	keyFileRemoveKeyFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileRemoveKeyFunction, err = keyFileStruct.InvokerNew("remove_key")
	})
	return err
}

// RemoveKey is a representation of the C type g_key_file_remove_key.
func (recv *KeyFile) RemoveKey(groupName string, key string) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)

	var ret gi.Argument

	err := keyFileRemoveKeyFunction_Set()
	if err == nil {
		ret = keyFileRemoveKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var keyFileSaveToFileFunction *gi.Function
var keyFileSaveToFileFunction_Once sync.Once

func keyFileSaveToFileFunction_Set() error {
	var err error
	keyFileSaveToFileFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSaveToFileFunction, err = keyFileStruct.InvokerNew("save_to_file")
	})
	return err
}

// SaveToFile is a representation of the C type g_key_file_save_to_file.
func (recv *KeyFile) SaveToFile(filename string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(filename)

	var ret gi.Argument

	err := keyFileSaveToFileFunction_Set()
	if err == nil {
		ret = keyFileSaveToFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var keyFileSetBooleanFunction *gi.Function
var keyFileSetBooleanFunction_Once sync.Once

func keyFileSetBooleanFunction_Set() error {
	var err error
	keyFileSetBooleanFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetBooleanFunction, err = keyFileStruct.InvokerNew("set_boolean")
	})
	return err
}

// SetBoolean is a representation of the C type g_key_file_set_boolean.
func (recv *KeyFile) SetBoolean(groupName string, key string, value bool) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetBoolean(value)

	err := keyFileSetBooleanFunction_Set()
	if err == nil {
		keyFileSetBooleanFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_key_file_set_boolean_list' : parameter 'list' has no type

var keyFileSetCommentFunction *gi.Function
var keyFileSetCommentFunction_Once sync.Once

func keyFileSetCommentFunction_Set() error {
	var err error
	keyFileSetCommentFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetCommentFunction, err = keyFileStruct.InvokerNew("set_comment")
	})
	return err
}

// SetComment is a representation of the C type g_key_file_set_comment.
func (recv *KeyFile) SetComment(groupName string, key string, comment string) (bool, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(comment)

	var ret gi.Argument

	err := keyFileSetCommentFunction_Set()
	if err == nil {
		ret = keyFileSetCommentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var keyFileSetDoubleFunction *gi.Function
var keyFileSetDoubleFunction_Once sync.Once

func keyFileSetDoubleFunction_Set() error {
	var err error
	keyFileSetDoubleFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetDoubleFunction, err = keyFileStruct.InvokerNew("set_double")
	})
	return err
}

// SetDouble is a representation of the C type g_key_file_set_double.
func (recv *KeyFile) SetDouble(groupName string, key string, value float64) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetFloat64(value)

	err := keyFileSetDoubleFunction_Set()
	if err == nil {
		keyFileSetDoubleFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_key_file_set_double_list' : parameter 'list' has no type

var keyFileSetInt64Function *gi.Function
var keyFileSetInt64Function_Once sync.Once

func keyFileSetInt64Function_Set() error {
	var err error
	keyFileSetInt64Function_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetInt64Function, err = keyFileStruct.InvokerNew("set_int64")
	})
	return err
}

// SetInt64 is a representation of the C type g_key_file_set_int64.
func (recv *KeyFile) SetInt64(groupName string, key string, value int64) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetInt64(value)

	err := keyFileSetInt64Function_Set()
	if err == nil {
		keyFileSetInt64Function.Invoke(inArgs[:], nil)
	}

	return err
}

var keyFileSetIntegerFunction *gi.Function
var keyFileSetIntegerFunction_Once sync.Once

func keyFileSetIntegerFunction_Set() error {
	var err error
	keyFileSetIntegerFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetIntegerFunction, err = keyFileStruct.InvokerNew("set_integer")
	})
	return err
}

// SetInteger is a representation of the C type g_key_file_set_integer.
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetInt32(value)

	err := keyFileSetIntegerFunction_Set()
	if err == nil {
		keyFileSetIntegerFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_key_file_set_integer_list' : parameter 'list' has no type

var keyFileSetListSeparatorFunction *gi.Function
var keyFileSetListSeparatorFunction_Once sync.Once

func keyFileSetListSeparatorFunction_Set() error {
	var err error
	keyFileSetListSeparatorFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetListSeparatorFunction, err = keyFileStruct.InvokerNew("set_list_separator")
	})
	return err
}

// SetListSeparator is a representation of the C type g_key_file_set_list_separator.
func (recv *KeyFile) SetListSeparator(separator int8) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(separator)

	err := keyFileSetListSeparatorFunction_Set()
	if err == nil {
		keyFileSetListSeparatorFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var keyFileSetLocaleStringFunction *gi.Function
var keyFileSetLocaleStringFunction_Once sync.Once

func keyFileSetLocaleStringFunction_Set() error {
	var err error
	keyFileSetLocaleStringFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetLocaleStringFunction, err = keyFileStruct.InvokerNew("set_locale_string")
	})
	return err
}

// SetLocaleString is a representation of the C type g_key_file_set_locale_string.
func (recv *KeyFile) SetLocaleString(groupName string, key string, locale string, string_ string) error {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(locale)
	inArgs[4].SetString(string_)

	err := keyFileSetLocaleStringFunction_Set()
	if err == nil {
		keyFileSetLocaleStringFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_key_file_set_locale_string_list' : parameter 'list' has no type

var keyFileSetStringFunction *gi.Function
var keyFileSetStringFunction_Once sync.Once

func keyFileSetStringFunction_Set() error {
	var err error
	keyFileSetStringFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetStringFunction, err = keyFileStruct.InvokerNew("set_string")
	})
	return err
}

// SetString is a representation of the C type g_key_file_set_string.
func (recv *KeyFile) SetString(groupName string, key string, string_ string) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(string_)

	err := keyFileSetStringFunction_Set()
	if err == nil {
		keyFileSetStringFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_key_file_set_string_list' : parameter 'list' has no type

var keyFileSetUint64Function *gi.Function
var keyFileSetUint64Function_Once sync.Once

func keyFileSetUint64Function_Set() error {
	var err error
	keyFileSetUint64Function_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetUint64Function, err = keyFileStruct.InvokerNew("set_uint64")
	})
	return err
}

// SetUint64 is a representation of the C type g_key_file_set_uint64.
func (recv *KeyFile) SetUint64(groupName string, key string, value uint64) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetUint64(value)

	err := keyFileSetUint64Function_Set()
	if err == nil {
		keyFileSetUint64Function.Invoke(inArgs[:], nil)
	}

	return err
}

var keyFileSetValueFunction *gi.Function
var keyFileSetValueFunction_Once sync.Once

func keyFileSetValueFunction_Set() error {
	var err error
	keyFileSetValueFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileSetValueFunction, err = keyFileStruct.InvokerNew("set_value")
	})
	return err
}

// SetValue is a representation of the C type g_key_file_set_value.
func (recv *KeyFile) SetValue(groupName string, key string, value string) error {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)
	inArgs[2].SetString(key)
	inArgs[3].SetString(value)

	err := keyFileSetValueFunction_Set()
	if err == nil {
		keyFileSetValueFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_key_file_to_data' : parameter 'length' of type 'gsize' not supported

var keyFileUnrefFunction *gi.Function
var keyFileUnrefFunction_Once sync.Once

func keyFileUnrefFunction_Set() error {
	var err error
	keyFileUnrefFunction_Once.Do(func() {
		err = keyFileStruct_Set()
		if err != nil {
			return
		}
		keyFileUnrefFunction, err = keyFileStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_key_file_unref.
func (recv *KeyFile) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := keyFileUnrefFunction_Set()
	if err == nil {
		keyFileUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var listStruct *gi.Struct
var listStruct_Once sync.Once

func listStruct_Set() error {
	var err error
	listStruct_Once.Do(func() {
		listStruct, err = gi.StructNew("GLib", "List")
	})
	return err
}

type List struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.List'
	// UNSUPPORTED : C value 'prev' : no Go type for 'GLib.List'
}

var logFieldStruct *gi.Struct
var logFieldStruct_Once sync.Once

func logFieldStruct_Set() error {
	var err error
	logFieldStruct_Once.Do(func() {
		logFieldStruct, err = gi.StructNew("GLib", "LogField")
	})
	return err
}

type LogField struct {
	native uintptr
	Key    string
	// UNSUPPORTED : C value 'value' : no Go type for 'gpointer'
	Length int32
}

var mainContextStruct *gi.Struct
var mainContextStruct_Once sync.Once

func mainContextStruct_Set() error {
	var err error
	mainContextStruct_Once.Do(func() {
		mainContextStruct, err = gi.StructNew("GLib", "MainContext")
	})
	return err
}

type MainContext struct {
	native uintptr
}

var mainContextNewFunction *gi.Function
var mainContextNewFunction_Once sync.Once

func mainContextNewFunction_Set() error {
	var err error
	mainContextNewFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextNewFunction, err = mainContextStruct.InvokerNew("new")
	})
	return err
}

// MainContextNew is a representation of the C type g_main_context_new.
func MainContextNew() (*MainContext, error) {

	var ret gi.Argument

	err := mainContextNewFunction_Set()
	if err == nil {
		ret = mainContextNewFunction.Invoke(nil, nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

var mainContextAcquireFunction *gi.Function
var mainContextAcquireFunction_Once sync.Once

func mainContextAcquireFunction_Set() error {
	var err error
	mainContextAcquireFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextAcquireFunction, err = mainContextStruct.InvokerNew("acquire")
	})
	return err
}

// Acquire is a representation of the C type g_main_context_acquire.
func (recv *MainContext) Acquire() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainContextAcquireFunction_Set()
	if err == nil {
		ret = mainContextAcquireFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_main_context_add_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_main_context_check' : parameter 'fds' has no type

var mainContextDispatchFunction *gi.Function
var mainContextDispatchFunction_Once sync.Once

func mainContextDispatchFunction_Set() error {
	var err error
	mainContextDispatchFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextDispatchFunction, err = mainContextStruct.InvokerNew("dispatch")
	})
	return err
}

// Dispatch is a representation of the C type g_main_context_dispatch.
func (recv *MainContext) Dispatch() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainContextDispatchFunction_Set()
	if err == nil {
		mainContextDispatchFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_main_context_find_source_by_funcs_user_data' : parameter 'funcs' of type 'SourceFuncs' not supported

var mainContextFindSourceByIdFunction *gi.Function
var mainContextFindSourceByIdFunction_Once sync.Once

func mainContextFindSourceByIdFunction_Set() error {
	var err error
	mainContextFindSourceByIdFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextFindSourceByIdFunction, err = mainContextStruct.InvokerNew("find_source_by_id")
	})
	return err
}

// FindSourceById is a representation of the C type g_main_context_find_source_by_id.
func (recv *MainContext) FindSourceById(sourceId uint32) (*Source, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(sourceId)

	var ret gi.Argument

	err := mainContextFindSourceByIdFunction_Set()
	if err == nil {
		ret = mainContextFindSourceByIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_main_context_find_source_by_user_data' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_main_context_get_poll_func' : return type 'PollFunc' not supported

// UNSUPPORTED : C value 'g_main_context_invoke' : parameter 'function' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_main_context_invoke_full' : parameter 'function' of type 'SourceFunc' not supported

var mainContextIsOwnerFunction *gi.Function
var mainContextIsOwnerFunction_Once sync.Once

func mainContextIsOwnerFunction_Set() error {
	var err error
	mainContextIsOwnerFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextIsOwnerFunction, err = mainContextStruct.InvokerNew("is_owner")
	})
	return err
}

// IsOwner is a representation of the C type g_main_context_is_owner.
func (recv *MainContext) IsOwner() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainContextIsOwnerFunction_Set()
	if err == nil {
		ret = mainContextIsOwnerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var mainContextIterationFunction *gi.Function
var mainContextIterationFunction_Once sync.Once

func mainContextIterationFunction_Set() error {
	var err error
	mainContextIterationFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextIterationFunction, err = mainContextStruct.InvokerNew("iteration")
	})
	return err
}

// Iteration is a representation of the C type g_main_context_iteration.
func (recv *MainContext) Iteration(mayBlock bool) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(mayBlock)

	var ret gi.Argument

	err := mainContextIterationFunction_Set()
	if err == nil {
		ret = mainContextIterationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var mainContextPendingFunction *gi.Function
var mainContextPendingFunction_Once sync.Once

func mainContextPendingFunction_Set() error {
	var err error
	mainContextPendingFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextPendingFunction, err = mainContextStruct.InvokerNew("pending")
	})
	return err
}

// Pending is a representation of the C type g_main_context_pending.
func (recv *MainContext) Pending() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainContextPendingFunction_Set()
	if err == nil {
		ret = mainContextPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var mainContextPopThreadDefaultFunction *gi.Function
var mainContextPopThreadDefaultFunction_Once sync.Once

func mainContextPopThreadDefaultFunction_Set() error {
	var err error
	mainContextPopThreadDefaultFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextPopThreadDefaultFunction, err = mainContextStruct.InvokerNew("pop_thread_default")
	})
	return err
}

// PopThreadDefault is a representation of the C type g_main_context_pop_thread_default.
func (recv *MainContext) PopThreadDefault() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainContextPopThreadDefaultFunction_Set()
	if err == nil {
		mainContextPopThreadDefaultFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mainContextPrepareFunction *gi.Function
var mainContextPrepareFunction_Once sync.Once

func mainContextPrepareFunction_Set() error {
	var err error
	mainContextPrepareFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextPrepareFunction, err = mainContextStruct.InvokerNew("prepare")
	})
	return err
}

// Prepare is a representation of the C type g_main_context_prepare.
func (recv *MainContext) Prepare() (bool, int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := mainContextPrepareFunction_Set()
	if err == nil {
		ret = mainContextPrepareFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()

	return retGo, out0, err
}

var mainContextPushThreadDefaultFunction *gi.Function
var mainContextPushThreadDefaultFunction_Once sync.Once

func mainContextPushThreadDefaultFunction_Set() error {
	var err error
	mainContextPushThreadDefaultFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextPushThreadDefaultFunction, err = mainContextStruct.InvokerNew("push_thread_default")
	})
	return err
}

// PushThreadDefault is a representation of the C type g_main_context_push_thread_default.
func (recv *MainContext) PushThreadDefault() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainContextPushThreadDefaultFunction_Set()
	if err == nil {
		mainContextPushThreadDefaultFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_main_context_query' : parameter 'fds' has no type

var mainContextRefFunction *gi.Function
var mainContextRefFunction_Once sync.Once

func mainContextRefFunction_Set() error {
	var err error
	mainContextRefFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextRefFunction, err = mainContextStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_main_context_ref.
func (recv *MainContext) Ref() (*MainContext, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainContextRefFunction_Set()
	if err == nil {
		ret = mainContextRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

var mainContextReleaseFunction *gi.Function
var mainContextReleaseFunction_Once sync.Once

func mainContextReleaseFunction_Set() error {
	var err error
	mainContextReleaseFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextReleaseFunction, err = mainContextStruct.InvokerNew("release")
	})
	return err
}

// Release is a representation of the C type g_main_context_release.
func (recv *MainContext) Release() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainContextReleaseFunction_Set()
	if err == nil {
		mainContextReleaseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_main_context_remove_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_main_context_set_poll_func' : parameter 'func' of type 'PollFunc' not supported

var mainContextUnrefFunction *gi.Function
var mainContextUnrefFunction_Once sync.Once

func mainContextUnrefFunction_Set() error {
	var err error
	mainContextUnrefFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextUnrefFunction, err = mainContextStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_main_context_unref.
func (recv *MainContext) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainContextUnrefFunction_Set()
	if err == nil {
		mainContextUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_main_context_wait' : parameter 'cond' of type 'Cond' not supported

var mainContextWakeupFunction *gi.Function
var mainContextWakeupFunction_Once sync.Once

func mainContextWakeupFunction_Set() error {
	var err error
	mainContextWakeupFunction_Once.Do(func() {
		err = mainContextStruct_Set()
		if err != nil {
			return
		}
		mainContextWakeupFunction, err = mainContextStruct.InvokerNew("wakeup")
	})
	return err
}

// Wakeup is a representation of the C type g_main_context_wakeup.
func (recv *MainContext) Wakeup() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainContextWakeupFunction_Set()
	if err == nil {
		mainContextWakeupFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mainLoopStruct *gi.Struct
var mainLoopStruct_Once sync.Once

func mainLoopStruct_Set() error {
	var err error
	mainLoopStruct_Once.Do(func() {
		mainLoopStruct, err = gi.StructNew("GLib", "MainLoop")
	})
	return err
}

type MainLoop struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_main_loop_new' : parameter 'context' of type 'MainContext' not supported

var mainLoopGetContextFunction *gi.Function
var mainLoopGetContextFunction_Once sync.Once

func mainLoopGetContextFunction_Set() error {
	var err error
	mainLoopGetContextFunction_Once.Do(func() {
		err = mainLoopStruct_Set()
		if err != nil {
			return
		}
		mainLoopGetContextFunction, err = mainLoopStruct.InvokerNew("get_context")
	})
	return err
}

// GetContext is a representation of the C type g_main_loop_get_context.
func (recv *MainLoop) GetContext() (*MainContext, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainLoopGetContextFunction_Set()
	if err == nil {
		ret = mainLoopGetContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

var mainLoopIsRunningFunction *gi.Function
var mainLoopIsRunningFunction_Once sync.Once

func mainLoopIsRunningFunction_Set() error {
	var err error
	mainLoopIsRunningFunction_Once.Do(func() {
		err = mainLoopStruct_Set()
		if err != nil {
			return
		}
		mainLoopIsRunningFunction, err = mainLoopStruct.InvokerNew("is_running")
	})
	return err
}

// IsRunning is a representation of the C type g_main_loop_is_running.
func (recv *MainLoop) IsRunning() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainLoopIsRunningFunction_Set()
	if err == nil {
		ret = mainLoopIsRunningFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var mainLoopQuitFunction *gi.Function
var mainLoopQuitFunction_Once sync.Once

func mainLoopQuitFunction_Set() error {
	var err error
	mainLoopQuitFunction_Once.Do(func() {
		err = mainLoopStruct_Set()
		if err != nil {
			return
		}
		mainLoopQuitFunction, err = mainLoopStruct.InvokerNew("quit")
	})
	return err
}

// Quit is a representation of the C type g_main_loop_quit.
func (recv *MainLoop) Quit() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainLoopQuitFunction_Set()
	if err == nil {
		mainLoopQuitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mainLoopRefFunction *gi.Function
var mainLoopRefFunction_Once sync.Once

func mainLoopRefFunction_Set() error {
	var err error
	mainLoopRefFunction_Once.Do(func() {
		err = mainLoopStruct_Set()
		if err != nil {
			return
		}
		mainLoopRefFunction, err = mainLoopStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_main_loop_ref.
func (recv *MainLoop) Ref() (*MainLoop, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mainLoopRefFunction_Set()
	if err == nil {
		ret = mainLoopRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MainLoop{native: ret.Pointer()}

	return retGo, err
}

var mainLoopRunFunction *gi.Function
var mainLoopRunFunction_Once sync.Once

func mainLoopRunFunction_Set() error {
	var err error
	mainLoopRunFunction_Once.Do(func() {
		err = mainLoopStruct_Set()
		if err != nil {
			return
		}
		mainLoopRunFunction, err = mainLoopStruct.InvokerNew("run")
	})
	return err
}

// Run is a representation of the C type g_main_loop_run.
func (recv *MainLoop) Run() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainLoopRunFunction_Set()
	if err == nil {
		mainLoopRunFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mainLoopUnrefFunction *gi.Function
var mainLoopUnrefFunction_Once sync.Once

func mainLoopUnrefFunction_Set() error {
	var err error
	mainLoopUnrefFunction_Once.Do(func() {
		err = mainLoopStruct_Set()
		if err != nil {
			return
		}
		mainLoopUnrefFunction, err = mainLoopStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_main_loop_unref.
func (recv *MainLoop) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mainLoopUnrefFunction_Set()
	if err == nil {
		mainLoopUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mappedFileStruct *gi.Struct
var mappedFileStruct_Once sync.Once

func mappedFileStruct_Set() error {
	var err error
	mappedFileStruct_Once.Do(func() {
		mappedFileStruct, err = gi.StructNew("GLib", "MappedFile")
	})
	return err
}

type MappedFile struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_mapped_file_new' : parameter 'filename' of type 'filename' not supported

var mappedFileNewFromFdFunction *gi.Function
var mappedFileNewFromFdFunction_Once sync.Once

func mappedFileNewFromFdFunction_Set() error {
	var err error
	mappedFileNewFromFdFunction_Once.Do(func() {
		err = mappedFileStruct_Set()
		if err != nil {
			return
		}
		mappedFileNewFromFdFunction, err = mappedFileStruct.InvokerNew("new_from_fd")
	})
	return err
}

// MappedFileNewFromFd is a representation of the C type g_mapped_file_new_from_fd.
func MappedFileNewFromFd(fd int32, writable bool) (*MappedFile, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(fd)
	inArgs[1].SetBoolean(writable)

	var ret gi.Argument

	err := mappedFileNewFromFdFunction_Set()
	if err == nil {
		ret = mappedFileNewFromFdFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MappedFile{native: ret.Pointer()}

	return retGo, err
}

var mappedFileFreeFunction *gi.Function
var mappedFileFreeFunction_Once sync.Once

func mappedFileFreeFunction_Set() error {
	var err error
	mappedFileFreeFunction_Once.Do(func() {
		err = mappedFileStruct_Set()
		if err != nil {
			return
		}
		mappedFileFreeFunction, err = mappedFileStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_mapped_file_free.
func (recv *MappedFile) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mappedFileFreeFunction_Set()
	if err == nil {
		mappedFileFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var mappedFileGetBytesFunction *gi.Function
var mappedFileGetBytesFunction_Once sync.Once

func mappedFileGetBytesFunction_Set() error {
	var err error
	mappedFileGetBytesFunction_Once.Do(func() {
		err = mappedFileStruct_Set()
		if err != nil {
			return
		}
		mappedFileGetBytesFunction, err = mappedFileStruct.InvokerNew("get_bytes")
	})
	return err
}

// GetBytes is a representation of the C type g_mapped_file_get_bytes.
func (recv *MappedFile) GetBytes() (*Bytes, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mappedFileGetBytesFunction_Set()
	if err == nil {
		ret = mappedFileGetBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Bytes{native: ret.Pointer()}

	return retGo, err
}

var mappedFileGetContentsFunction *gi.Function
var mappedFileGetContentsFunction_Once sync.Once

func mappedFileGetContentsFunction_Set() error {
	var err error
	mappedFileGetContentsFunction_Once.Do(func() {
		err = mappedFileStruct_Set()
		if err != nil {
			return
		}
		mappedFileGetContentsFunction, err = mappedFileStruct.InvokerNew("get_contents")
	})
	return err
}

// GetContents is a representation of the C type g_mapped_file_get_contents.
func (recv *MappedFile) GetContents() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mappedFileGetContentsFunction_Set()
	if err == nil {
		ret = mappedFileGetContentsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_mapped_file_get_length' : return type 'gsize' not supported

var mappedFileRefFunction *gi.Function
var mappedFileRefFunction_Once sync.Once

func mappedFileRefFunction_Set() error {
	var err error
	mappedFileRefFunction_Once.Do(func() {
		err = mappedFileStruct_Set()
		if err != nil {
			return
		}
		mappedFileRefFunction, err = mappedFileStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_mapped_file_ref.
func (recv *MappedFile) Ref() (*MappedFile, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := mappedFileRefFunction_Set()
	if err == nil {
		ret = mappedFileRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MappedFile{native: ret.Pointer()}

	return retGo, err
}

var mappedFileUnrefFunction *gi.Function
var mappedFileUnrefFunction_Once sync.Once

func mappedFileUnrefFunction_Set() error {
	var err error
	mappedFileUnrefFunction_Once.Do(func() {
		err = mappedFileStruct_Set()
		if err != nil {
			return
		}
		mappedFileUnrefFunction, err = mappedFileStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_mapped_file_unref.
func (recv *MappedFile) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := mappedFileUnrefFunction_Set()
	if err == nil {
		mappedFileUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var markupParseContextStruct *gi.Struct
var markupParseContextStruct_Once sync.Once

func markupParseContextStruct_Set() error {
	var err error
	markupParseContextStruct_Once.Do(func() {
		markupParseContextStruct, err = gi.StructNew("GLib", "MarkupParseContext")
	})
	return err
}

type MarkupParseContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_markup_parse_context_new' : parameter 'parser' of type 'MarkupParser' not supported

var markupParseContextEndParseFunction *gi.Function
var markupParseContextEndParseFunction_Once sync.Once

func markupParseContextEndParseFunction_Set() error {
	var err error
	markupParseContextEndParseFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextEndParseFunction, err = markupParseContextStruct.InvokerNew("end_parse")
	})
	return err
}

// EndParse is a representation of the C type g_markup_parse_context_end_parse.
func (recv *MarkupParseContext) EndParse() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := markupParseContextEndParseFunction_Set()
	if err == nil {
		ret = markupParseContextEndParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var markupParseContextFreeFunction *gi.Function
var markupParseContextFreeFunction_Once sync.Once

func markupParseContextFreeFunction_Set() error {
	var err error
	markupParseContextFreeFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextFreeFunction, err = markupParseContextStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_markup_parse_context_free.
func (recv *MarkupParseContext) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := markupParseContextFreeFunction_Set()
	if err == nil {
		markupParseContextFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var markupParseContextGetElementFunction *gi.Function
var markupParseContextGetElementFunction_Once sync.Once

func markupParseContextGetElementFunction_Set() error {
	var err error
	markupParseContextGetElementFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextGetElementFunction, err = markupParseContextStruct.InvokerNew("get_element")
	})
	return err
}

// GetElement is a representation of the C type g_markup_parse_context_get_element.
func (recv *MarkupParseContext) GetElement() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := markupParseContextGetElementFunction_Set()
	if err == nil {
		ret = markupParseContextGetElementFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_markup_parse_context_get_element_stack' : return type 'GLib.SList' not supported

var markupParseContextGetPositionFunction *gi.Function
var markupParseContextGetPositionFunction_Once sync.Once

func markupParseContextGetPositionFunction_Set() error {
	var err error
	markupParseContextGetPositionFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextGetPositionFunction, err = markupParseContextStruct.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type g_markup_parse_context_get_position.
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(lineNumber)
	inArgs[2].SetInt32(charNumber)

	err := markupParseContextGetPositionFunction_Set()
	if err == nil {
		markupParseContextGetPositionFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_markup_parse_context_get_user_data' : return type 'gpointer' not supported

var markupParseContextParseFunction *gi.Function
var markupParseContextParseFunction_Once sync.Once

func markupParseContextParseFunction_Set() error {
	var err error
	markupParseContextParseFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextParseFunction, err = markupParseContextStruct.InvokerNew("parse")
	})
	return err
}

// Parse is a representation of the C type g_markup_parse_context_parse.
func (recv *MarkupParseContext) Parse(text string, textLen int32) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(textLen)

	var ret gi.Argument

	err := markupParseContextParseFunction_Set()
	if err == nil {
		ret = markupParseContextParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_markup_parse_context_pop' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_markup_parse_context_push' : parameter 'parser' of type 'MarkupParser' not supported

var markupParseContextRefFunction *gi.Function
var markupParseContextRefFunction_Once sync.Once

func markupParseContextRefFunction_Set() error {
	var err error
	markupParseContextRefFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextRefFunction, err = markupParseContextStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_markup_parse_context_ref.
func (recv *MarkupParseContext) Ref() (*MarkupParseContext, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := markupParseContextRefFunction_Set()
	if err == nil {
		ret = markupParseContextRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MarkupParseContext{native: ret.Pointer()}

	return retGo, err
}

var markupParseContextUnrefFunction *gi.Function
var markupParseContextUnrefFunction_Once sync.Once

func markupParseContextUnrefFunction_Set() error {
	var err error
	markupParseContextUnrefFunction_Once.Do(func() {
		err = markupParseContextStruct_Set()
		if err != nil {
			return
		}
		markupParseContextUnrefFunction, err = markupParseContextStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_markup_parse_context_unref.
func (recv *MarkupParseContext) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := markupParseContextUnrefFunction_Set()
	if err == nil {
		markupParseContextUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var markupParserStruct *gi.Struct
var markupParserStruct_Once sync.Once

func markupParserStruct_Set() error {
	var err error
	markupParserStruct_Once.Do(func() {
		markupParserStruct, err = gi.StructNew("GLib", "MarkupParser")
	})
	return err
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
var matchInfoStruct_Once sync.Once

func matchInfoStruct_Set() error {
	var err error
	matchInfoStruct_Once.Do(func() {
		matchInfoStruct, err = gi.StructNew("GLib", "MatchInfo")
	})
	return err
}

type MatchInfo struct {
	native uintptr
}

var matchInfoExpandReferencesFunction *gi.Function
var matchInfoExpandReferencesFunction_Once sync.Once

func matchInfoExpandReferencesFunction_Set() error {
	var err error
	matchInfoExpandReferencesFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoExpandReferencesFunction, err = matchInfoStruct.InvokerNew("expand_references")
	})
	return err
}

// ExpandReferences is a representation of the C type g_match_info_expand_references.
func (recv *MatchInfo) ExpandReferences(stringToExpand string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(stringToExpand)

	var ret gi.Argument

	err := matchInfoExpandReferencesFunction_Set()
	if err == nil {
		ret = matchInfoExpandReferencesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var matchInfoFetchFunction *gi.Function
var matchInfoFetchFunction_Once sync.Once

func matchInfoFetchFunction_Set() error {
	var err error
	matchInfoFetchFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoFetchFunction, err = matchInfoStruct.InvokerNew("fetch")
	})
	return err
}

// Fetch is a representation of the C type g_match_info_fetch.
func (recv *MatchInfo) Fetch(matchNum int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(matchNum)

	var ret gi.Argument

	err := matchInfoFetchFunction_Set()
	if err == nil {
		ret = matchInfoFetchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var matchInfoFetchAllFunction *gi.Function
var matchInfoFetchAllFunction_Once sync.Once

func matchInfoFetchAllFunction_Set() error {
	var err error
	matchInfoFetchAllFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoFetchAllFunction, err = matchInfoStruct.InvokerNew("fetch_all")
	})
	return err
}

// FetchAll is a representation of the C type g_match_info_fetch_all.
func (recv *MatchInfo) FetchAll() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := matchInfoFetchAllFunction_Set()
	if err == nil {
		matchInfoFetchAllFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var matchInfoFetchNamedFunction *gi.Function
var matchInfoFetchNamedFunction_Once sync.Once

func matchInfoFetchNamedFunction_Set() error {
	var err error
	matchInfoFetchNamedFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoFetchNamedFunction, err = matchInfoStruct.InvokerNew("fetch_named")
	})
	return err
}

// FetchNamed is a representation of the C type g_match_info_fetch_named.
func (recv *MatchInfo) FetchNamed(name string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := matchInfoFetchNamedFunction_Set()
	if err == nil {
		ret = matchInfoFetchNamedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var matchInfoFetchNamedPosFunction *gi.Function
var matchInfoFetchNamedPosFunction_Once sync.Once

func matchInfoFetchNamedPosFunction_Set() error {
	var err error
	matchInfoFetchNamedPosFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoFetchNamedPosFunction, err = matchInfoStruct.InvokerNew("fetch_named_pos")
	})
	return err
}

// FetchNamedPos is a representation of the C type g_match_info_fetch_named_pos.
func (recv *MatchInfo) FetchNamedPos(name string) (bool, int32, int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := matchInfoFetchNamedPosFunction_Set()
	if err == nil {
		ret = matchInfoFetchNamedPosFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1, err
}

var matchInfoFetchPosFunction *gi.Function
var matchInfoFetchPosFunction_Once sync.Once

func matchInfoFetchPosFunction_Set() error {
	var err error
	matchInfoFetchPosFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoFetchPosFunction, err = matchInfoStruct.InvokerNew("fetch_pos")
	})
	return err
}

// FetchPos is a representation of the C type g_match_info_fetch_pos.
func (recv *MatchInfo) FetchPos(matchNum int32) (bool, int32, int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(matchNum)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := matchInfoFetchPosFunction_Set()
	if err == nil {
		ret = matchInfoFetchPosFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1, err
}

var matchInfoFreeFunction *gi.Function
var matchInfoFreeFunction_Once sync.Once

func matchInfoFreeFunction_Set() error {
	var err error
	matchInfoFreeFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoFreeFunction, err = matchInfoStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_match_info_free.
func (recv *MatchInfo) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := matchInfoFreeFunction_Set()
	if err == nil {
		matchInfoFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var matchInfoGetMatchCountFunction *gi.Function
var matchInfoGetMatchCountFunction_Once sync.Once

func matchInfoGetMatchCountFunction_Set() error {
	var err error
	matchInfoGetMatchCountFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoGetMatchCountFunction, err = matchInfoStruct.InvokerNew("get_match_count")
	})
	return err
}

// GetMatchCount is a representation of the C type g_match_info_get_match_count.
func (recv *MatchInfo) GetMatchCount() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoGetMatchCountFunction_Set()
	if err == nil {
		ret = matchInfoGetMatchCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var matchInfoGetRegexFunction *gi.Function
var matchInfoGetRegexFunction_Once sync.Once

func matchInfoGetRegexFunction_Set() error {
	var err error
	matchInfoGetRegexFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoGetRegexFunction, err = matchInfoStruct.InvokerNew("get_regex")
	})
	return err
}

// GetRegex is a representation of the C type g_match_info_get_regex.
func (recv *MatchInfo) GetRegex() (*Regex, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoGetRegexFunction_Set()
	if err == nil {
		ret = matchInfoGetRegexFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Regex{native: ret.Pointer()}

	return retGo, err
}

var matchInfoGetStringFunction *gi.Function
var matchInfoGetStringFunction_Once sync.Once

func matchInfoGetStringFunction_Set() error {
	var err error
	matchInfoGetStringFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoGetStringFunction, err = matchInfoStruct.InvokerNew("get_string")
	})
	return err
}

// GetString is a representation of the C type g_match_info_get_string.
func (recv *MatchInfo) GetString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoGetStringFunction_Set()
	if err == nil {
		ret = matchInfoGetStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var matchInfoIsPartialMatchFunction *gi.Function
var matchInfoIsPartialMatchFunction_Once sync.Once

func matchInfoIsPartialMatchFunction_Set() error {
	var err error
	matchInfoIsPartialMatchFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoIsPartialMatchFunction, err = matchInfoStruct.InvokerNew("is_partial_match")
	})
	return err
}

// IsPartialMatch is a representation of the C type g_match_info_is_partial_match.
func (recv *MatchInfo) IsPartialMatch() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoIsPartialMatchFunction_Set()
	if err == nil {
		ret = matchInfoIsPartialMatchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var matchInfoMatchesFunction *gi.Function
var matchInfoMatchesFunction_Once sync.Once

func matchInfoMatchesFunction_Set() error {
	var err error
	matchInfoMatchesFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoMatchesFunction, err = matchInfoStruct.InvokerNew("matches")
	})
	return err
}

// Matches is a representation of the C type g_match_info_matches.
func (recv *MatchInfo) Matches() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoMatchesFunction_Set()
	if err == nil {
		ret = matchInfoMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var matchInfoNextFunction *gi.Function
var matchInfoNextFunction_Once sync.Once

func matchInfoNextFunction_Set() error {
	var err error
	matchInfoNextFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoNextFunction, err = matchInfoStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type g_match_info_next.
func (recv *MatchInfo) Next() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoNextFunction_Set()
	if err == nil {
		ret = matchInfoNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var matchInfoRefFunction *gi.Function
var matchInfoRefFunction_Once sync.Once

func matchInfoRefFunction_Set() error {
	var err error
	matchInfoRefFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoRefFunction, err = matchInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_match_info_ref.
func (recv *MatchInfo) Ref() (*MatchInfo, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matchInfoRefFunction_Set()
	if err == nil {
		ret = matchInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MatchInfo{native: ret.Pointer()}

	return retGo, err
}

var matchInfoUnrefFunction *gi.Function
var matchInfoUnrefFunction_Once sync.Once

func matchInfoUnrefFunction_Set() error {
	var err error
	matchInfoUnrefFunction_Once.Do(func() {
		err = matchInfoStruct_Set()
		if err != nil {
			return
		}
		matchInfoUnrefFunction, err = matchInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_match_info_unref.
func (recv *MatchInfo) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := matchInfoUnrefFunction_Set()
	if err == nil {
		matchInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var memVTableStruct *gi.Struct
var memVTableStruct_Once sync.Once

func memVTableStruct_Set() error {
	var err error
	memVTableStruct_Once.Do(func() {
		memVTableStruct, err = gi.StructNew("GLib", "MemVTable")
	})
	return err
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
var nodeStruct_Once sync.Once

func nodeStruct_Set() error {
	var err error
	nodeStruct_Once.Do(func() {
		nodeStruct, err = gi.StructNew("GLib", "Node")
	})
	return err
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

func nodeCopyFunction_Set() error {
	var err error
	nodeCopyFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeCopyFunction, err = nodeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_node_copy.
func (recv *Node) Copy() (*Node, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeCopyFunction_Set()
	if err == nil {
		ret = nodeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Node{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_node_copy_deep' : parameter 'copy_func' of type 'CopyFunc' not supported

var nodeDepthFunction *gi.Function
var nodeDepthFunction_Once sync.Once

func nodeDepthFunction_Set() error {
	var err error
	nodeDepthFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeDepthFunction, err = nodeStruct.InvokerNew("depth")
	})
	return err
}

// Depth is a representation of the C type g_node_depth.
func (recv *Node) Depth() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeDepthFunction_Set()
	if err == nil {
		ret = nodeDepthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var nodeDestroyFunction *gi.Function
var nodeDestroyFunction_Once sync.Once

func nodeDestroyFunction_Set() error {
	var err error
	nodeDestroyFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeDestroyFunction, err = nodeStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type g_node_destroy.
func (recv *Node) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := nodeDestroyFunction_Set()
	if err == nil {
		nodeDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_node_find' : parameter 'order' of type 'TraverseType' not supported

// UNSUPPORTED : C value 'g_node_find_child' : parameter 'flags' of type 'TraverseFlags' not supported

var nodeFirstSiblingFunction *gi.Function
var nodeFirstSiblingFunction_Once sync.Once

func nodeFirstSiblingFunction_Set() error {
	var err error
	nodeFirstSiblingFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeFirstSiblingFunction, err = nodeStruct.InvokerNew("first_sibling")
	})
	return err
}

// FirstSibling is a representation of the C type g_node_first_sibling.
func (recv *Node) FirstSibling() (*Node, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeFirstSiblingFunction_Set()
	if err == nil {
		ret = nodeFirstSiblingFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Node{native: ret.Pointer()}

	return retGo, err
}

var nodeGetRootFunction *gi.Function
var nodeGetRootFunction_Once sync.Once

func nodeGetRootFunction_Set() error {
	var err error
	nodeGetRootFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeGetRootFunction, err = nodeStruct.InvokerNew("get_root")
	})
	return err
}

// GetRoot is a representation of the C type g_node_get_root.
func (recv *Node) GetRoot() (*Node, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeGetRootFunction_Set()
	if err == nil {
		ret = nodeGetRootFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Node{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_node_insert' : parameter 'node' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_insert_after' : parameter 'sibling' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_insert_before' : parameter 'sibling' of type 'Node' not supported

// UNSUPPORTED : C value 'g_node_is_ancestor' : parameter 'descendant' of type 'Node' not supported

var nodeLastChildFunction *gi.Function
var nodeLastChildFunction_Once sync.Once

func nodeLastChildFunction_Set() error {
	var err error
	nodeLastChildFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeLastChildFunction, err = nodeStruct.InvokerNew("last_child")
	})
	return err
}

// LastChild is a representation of the C type g_node_last_child.
func (recv *Node) LastChild() (*Node, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeLastChildFunction_Set()
	if err == nil {
		ret = nodeLastChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Node{native: ret.Pointer()}

	return retGo, err
}

var nodeLastSiblingFunction *gi.Function
var nodeLastSiblingFunction_Once sync.Once

func nodeLastSiblingFunction_Set() error {
	var err error
	nodeLastSiblingFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeLastSiblingFunction, err = nodeStruct.InvokerNew("last_sibling")
	})
	return err
}

// LastSibling is a representation of the C type g_node_last_sibling.
func (recv *Node) LastSibling() (*Node, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeLastSiblingFunction_Set()
	if err == nil {
		ret = nodeLastSiblingFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Node{native: ret.Pointer()}

	return retGo, err
}

var nodeMaxHeightFunction *gi.Function
var nodeMaxHeightFunction_Once sync.Once

func nodeMaxHeightFunction_Set() error {
	var err error
	nodeMaxHeightFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeMaxHeightFunction, err = nodeStruct.InvokerNew("max_height")
	})
	return err
}

// MaxHeight is a representation of the C type g_node_max_height.
func (recv *Node) MaxHeight() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeMaxHeightFunction_Set()
	if err == nil {
		ret = nodeMaxHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var nodeNChildrenFunction *gi.Function
var nodeNChildrenFunction_Once sync.Once

func nodeNChildrenFunction_Set() error {
	var err error
	nodeNChildrenFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeNChildrenFunction, err = nodeStruct.InvokerNew("n_children")
	})
	return err
}

// NChildren is a representation of the C type g_node_n_children.
func (recv *Node) NChildren() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := nodeNChildrenFunction_Set()
	if err == nil {
		ret = nodeNChildrenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_node_n_nodes' : parameter 'flags' of type 'TraverseFlags' not supported

var nodeNthChildFunction *gi.Function
var nodeNthChildFunction_Once sync.Once

func nodeNthChildFunction_Set() error {
	var err error
	nodeNthChildFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeNthChildFunction, err = nodeStruct.InvokerNew("nth_child")
	})
	return err
}

// NthChild is a representation of the C type g_node_nth_child.
func (recv *Node) NthChild(n uint32) (*Node, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(n)

	var ret gi.Argument

	err := nodeNthChildFunction_Set()
	if err == nil {
		ret = nodeNthChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Node{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_node_prepend' : parameter 'node' of type 'Node' not supported

var nodeReverseChildrenFunction *gi.Function
var nodeReverseChildrenFunction_Once sync.Once

func nodeReverseChildrenFunction_Set() error {
	var err error
	nodeReverseChildrenFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeReverseChildrenFunction, err = nodeStruct.InvokerNew("reverse_children")
	})
	return err
}

// ReverseChildren is a representation of the C type g_node_reverse_children.
func (recv *Node) ReverseChildren() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := nodeReverseChildrenFunction_Set()
	if err == nil {
		nodeReverseChildrenFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_node_traverse' : parameter 'order' of type 'TraverseType' not supported

var nodeUnlinkFunction *gi.Function
var nodeUnlinkFunction_Once sync.Once

func nodeUnlinkFunction_Set() error {
	var err error
	nodeUnlinkFunction_Once.Do(func() {
		err = nodeStruct_Set()
		if err != nil {
			return
		}
		nodeUnlinkFunction, err = nodeStruct.InvokerNew("unlink")
	})
	return err
}

// Unlink is a representation of the C type g_node_unlink.
func (recv *Node) Unlink() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := nodeUnlinkFunction_Set()
	if err == nil {
		nodeUnlinkFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var onceStruct *gi.Struct
var onceStruct_Once sync.Once

func onceStruct_Set() error {
	var err error
	onceStruct_Once.Do(func() {
		onceStruct, err = gi.StructNew("GLib", "Once")
	})
	return err
}

type Once struct {
	native uintptr
	// UNSUPPORTED : C value 'status' : no Go type for 'OnceStatus'
	// UNSUPPORTED : C value 'retval' : no Go type for 'gpointer'
}

// UNSUPPORTED : C value 'g_once_impl' : parameter 'func' of type 'ThreadFunc' not supported

var optionContextStruct *gi.Struct
var optionContextStruct_Once sync.Once

func optionContextStruct_Set() error {
	var err error
	optionContextStruct_Once.Do(func() {
		optionContextStruct, err = gi.StructNew("GLib", "OptionContext")
	})
	return err
}

type OptionContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_context_add_group' : parameter 'group' of type 'OptionGroup' not supported

// UNSUPPORTED : C value 'g_option_context_add_main_entries' : parameter 'entries' of type 'OptionEntry' not supported

var optionContextFreeFunction *gi.Function
var optionContextFreeFunction_Once sync.Once

func optionContextFreeFunction_Set() error {
	var err error
	optionContextFreeFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextFreeFunction, err = optionContextStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_option_context_free.
func (recv *OptionContext) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := optionContextFreeFunction_Set()
	if err == nil {
		optionContextFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionContextGetDescriptionFunction *gi.Function
var optionContextGetDescriptionFunction_Once sync.Once

func optionContextGetDescriptionFunction_Set() error {
	var err error
	optionContextGetDescriptionFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextGetDescriptionFunction, err = optionContextStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type g_option_context_get_description.
func (recv *OptionContext) GetDescription() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionContextGetDescriptionFunction_Set()
	if err == nil {
		ret = optionContextGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_option_context_get_help' : parameter 'group' of type 'OptionGroup' not supported

var optionContextGetHelpEnabledFunction *gi.Function
var optionContextGetHelpEnabledFunction_Once sync.Once

func optionContextGetHelpEnabledFunction_Set() error {
	var err error
	optionContextGetHelpEnabledFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextGetHelpEnabledFunction, err = optionContextStruct.InvokerNew("get_help_enabled")
	})
	return err
}

// GetHelpEnabled is a representation of the C type g_option_context_get_help_enabled.
func (recv *OptionContext) GetHelpEnabled() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionContextGetHelpEnabledFunction_Set()
	if err == nil {
		ret = optionContextGetHelpEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionContextGetIgnoreUnknownOptionsFunction *gi.Function
var optionContextGetIgnoreUnknownOptionsFunction_Once sync.Once

func optionContextGetIgnoreUnknownOptionsFunction_Set() error {
	var err error
	optionContextGetIgnoreUnknownOptionsFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextGetIgnoreUnknownOptionsFunction, err = optionContextStruct.InvokerNew("get_ignore_unknown_options")
	})
	return err
}

// GetIgnoreUnknownOptions is a representation of the C type g_option_context_get_ignore_unknown_options.
func (recv *OptionContext) GetIgnoreUnknownOptions() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionContextGetIgnoreUnknownOptionsFunction_Set()
	if err == nil {
		ret = optionContextGetIgnoreUnknownOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionContextGetMainGroupFunction *gi.Function
var optionContextGetMainGroupFunction_Once sync.Once

func optionContextGetMainGroupFunction_Set() error {
	var err error
	optionContextGetMainGroupFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextGetMainGroupFunction, err = optionContextStruct.InvokerNew("get_main_group")
	})
	return err
}

// GetMainGroup is a representation of the C type g_option_context_get_main_group.
func (recv *OptionContext) GetMainGroup() (*OptionGroup, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionContextGetMainGroupFunction_Set()
	if err == nil {
		ret = optionContextGetMainGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &OptionGroup{native: ret.Pointer()}

	return retGo, err
}

var optionContextGetStrictPosixFunction *gi.Function
var optionContextGetStrictPosixFunction_Once sync.Once

func optionContextGetStrictPosixFunction_Set() error {
	var err error
	optionContextGetStrictPosixFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextGetStrictPosixFunction, err = optionContextStruct.InvokerNew("get_strict_posix")
	})
	return err
}

// GetStrictPosix is a representation of the C type g_option_context_get_strict_posix.
func (recv *OptionContext) GetStrictPosix() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionContextGetStrictPosixFunction_Set()
	if err == nil {
		ret = optionContextGetStrictPosixFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var optionContextGetSummaryFunction *gi.Function
var optionContextGetSummaryFunction_Once sync.Once

func optionContextGetSummaryFunction_Set() error {
	var err error
	optionContextGetSummaryFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextGetSummaryFunction, err = optionContextStruct.InvokerNew("get_summary")
	})
	return err
}

// GetSummary is a representation of the C type g_option_context_get_summary.
func (recv *OptionContext) GetSummary() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionContextGetSummaryFunction_Set()
	if err == nil {
		ret = optionContextGetSummaryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

// UNSUPPORTED : C value 'g_option_context_parse' : parameter 'argv' has no type

// UNSUPPORTED : C value 'g_option_context_parse_strv' : parameter 'arguments' has no type

var optionContextSetDescriptionFunction *gi.Function
var optionContextSetDescriptionFunction_Once sync.Once

func optionContextSetDescriptionFunction_Set() error {
	var err error
	optionContextSetDescriptionFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextSetDescriptionFunction, err = optionContextStruct.InvokerNew("set_description")
	})
	return err
}

// SetDescription is a representation of the C type g_option_context_set_description.
func (recv *OptionContext) SetDescription(description string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(description)

	err := optionContextSetDescriptionFunction_Set()
	if err == nil {
		optionContextSetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionContextSetHelpEnabledFunction *gi.Function
var optionContextSetHelpEnabledFunction_Once sync.Once

func optionContextSetHelpEnabledFunction_Set() error {
	var err error
	optionContextSetHelpEnabledFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextSetHelpEnabledFunction, err = optionContextStruct.InvokerNew("set_help_enabled")
	})
	return err
}

// SetHelpEnabled is a representation of the C type g_option_context_set_help_enabled.
func (recv *OptionContext) SetHelpEnabled(helpEnabled bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(helpEnabled)

	err := optionContextSetHelpEnabledFunction_Set()
	if err == nil {
		optionContextSetHelpEnabledFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionContextSetIgnoreUnknownOptionsFunction *gi.Function
var optionContextSetIgnoreUnknownOptionsFunction_Once sync.Once

func optionContextSetIgnoreUnknownOptionsFunction_Set() error {
	var err error
	optionContextSetIgnoreUnknownOptionsFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextSetIgnoreUnknownOptionsFunction, err = optionContextStruct.InvokerNew("set_ignore_unknown_options")
	})
	return err
}

// SetIgnoreUnknownOptions is a representation of the C type g_option_context_set_ignore_unknown_options.
func (recv *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(ignoreUnknown)

	err := optionContextSetIgnoreUnknownOptionsFunction_Set()
	if err == nil {
		optionContextSetIgnoreUnknownOptionsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_option_context_set_main_group' : parameter 'group' of type 'OptionGroup' not supported

var optionContextSetStrictPosixFunction *gi.Function
var optionContextSetStrictPosixFunction_Once sync.Once

func optionContextSetStrictPosixFunction_Set() error {
	var err error
	optionContextSetStrictPosixFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextSetStrictPosixFunction, err = optionContextStruct.InvokerNew("set_strict_posix")
	})
	return err
}

// SetStrictPosix is a representation of the C type g_option_context_set_strict_posix.
func (recv *OptionContext) SetStrictPosix(strictPosix bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(strictPosix)

	err := optionContextSetStrictPosixFunction_Set()
	if err == nil {
		optionContextSetStrictPosixFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionContextSetSummaryFunction *gi.Function
var optionContextSetSummaryFunction_Once sync.Once

func optionContextSetSummaryFunction_Set() error {
	var err error
	optionContextSetSummaryFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextSetSummaryFunction, err = optionContextStruct.InvokerNew("set_summary")
	})
	return err
}

// SetSummary is a representation of the C type g_option_context_set_summary.
func (recv *OptionContext) SetSummary(summary string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(summary)

	err := optionContextSetSummaryFunction_Set()
	if err == nil {
		optionContextSetSummaryFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_option_context_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

var optionContextSetTranslationDomainFunction *gi.Function
var optionContextSetTranslationDomainFunction_Once sync.Once

func optionContextSetTranslationDomainFunction_Set() error {
	var err error
	optionContextSetTranslationDomainFunction_Once.Do(func() {
		err = optionContextStruct_Set()
		if err != nil {
			return
		}
		optionContextSetTranslationDomainFunction, err = optionContextStruct.InvokerNew("set_translation_domain")
	})
	return err
}

// SetTranslationDomain is a representation of the C type g_option_context_set_translation_domain.
func (recv *OptionContext) SetTranslationDomain(domain string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	err := optionContextSetTranslationDomainFunction_Set()
	if err == nil {
		optionContextSetTranslationDomainFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionEntryStruct *gi.Struct
var optionEntryStruct_Once sync.Once

func optionEntryStruct_Set() error {
	var err error
	optionEntryStruct_Once.Do(func() {
		optionEntryStruct, err = gi.StructNew("GLib", "OptionEntry")
	})
	return err
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
var optionGroupStruct_Once sync.Once

func optionGroupStruct_Set() error {
	var err error
	optionGroupStruct_Once.Do(func() {
		optionGroupStruct, err = gi.StructNew("GLib", "OptionGroup")
	})
	return err
}

type OptionGroup struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_group_new' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_option_group_add_entries' : parameter 'entries' of type 'OptionEntry' not supported

var optionGroupFreeFunction *gi.Function
var optionGroupFreeFunction_Once sync.Once

func optionGroupFreeFunction_Set() error {
	var err error
	optionGroupFreeFunction_Once.Do(func() {
		err = optionGroupStruct_Set()
		if err != nil {
			return
		}
		optionGroupFreeFunction, err = optionGroupStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_option_group_free.
func (recv *OptionGroup) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := optionGroupFreeFunction_Set()
	if err == nil {
		optionGroupFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionGroupRefFunction *gi.Function
var optionGroupRefFunction_Once sync.Once

func optionGroupRefFunction_Set() error {
	var err error
	optionGroupRefFunction_Once.Do(func() {
		err = optionGroupStruct_Set()
		if err != nil {
			return
		}
		optionGroupRefFunction, err = optionGroupStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_option_group_ref.
func (recv *OptionGroup) Ref() (*OptionGroup, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := optionGroupRefFunction_Set()
	if err == nil {
		ret = optionGroupRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &OptionGroup{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_option_group_set_error_hook' : parameter 'error_func' of type 'OptionErrorFunc' not supported

// UNSUPPORTED : C value 'g_option_group_set_parse_hooks' : parameter 'pre_parse_func' of type 'OptionParseFunc' not supported

// UNSUPPORTED : C value 'g_option_group_set_translate_func' : parameter 'func' of type 'TranslateFunc' not supported

var optionGroupSetTranslationDomainFunction *gi.Function
var optionGroupSetTranslationDomainFunction_Once sync.Once

func optionGroupSetTranslationDomainFunction_Set() error {
	var err error
	optionGroupSetTranslationDomainFunction_Once.Do(func() {
		err = optionGroupStruct_Set()
		if err != nil {
			return
		}
		optionGroupSetTranslationDomainFunction, err = optionGroupStruct.InvokerNew("set_translation_domain")
	})
	return err
}

// SetTranslationDomain is a representation of the C type g_option_group_set_translation_domain.
func (recv *OptionGroup) SetTranslationDomain(domain string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(domain)

	err := optionGroupSetTranslationDomainFunction_Set()
	if err == nil {
		optionGroupSetTranslationDomainFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var optionGroupUnrefFunction *gi.Function
var optionGroupUnrefFunction_Once sync.Once

func optionGroupUnrefFunction_Set() error {
	var err error
	optionGroupUnrefFunction_Once.Do(func() {
		err = optionGroupStruct_Set()
		if err != nil {
			return
		}
		optionGroupUnrefFunction, err = optionGroupStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_option_group_unref.
func (recv *OptionGroup) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := optionGroupUnrefFunction_Set()
	if err == nil {
		optionGroupUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var patternSpecStruct *gi.Struct
var patternSpecStruct_Once sync.Once

func patternSpecStruct_Set() error {
	var err error
	patternSpecStruct_Once.Do(func() {
		patternSpecStruct, err = gi.StructNew("GLib", "PatternSpec")
	})
	return err
}

type PatternSpec struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_pattern_spec_equal' : parameter 'pspec2' of type 'PatternSpec' not supported

var patternSpecFreeFunction *gi.Function
var patternSpecFreeFunction_Once sync.Once

func patternSpecFreeFunction_Set() error {
	var err error
	patternSpecFreeFunction_Once.Do(func() {
		err = patternSpecStruct_Set()
		if err != nil {
			return
		}
		patternSpecFreeFunction, err = patternSpecStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_pattern_spec_free.
func (recv *PatternSpec) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := patternSpecFreeFunction_Set()
	if err == nil {
		patternSpecFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var pollFDStruct *gi.Struct
var pollFDStruct_Once sync.Once

func pollFDStruct_Set() error {
	var err error
	pollFDStruct_Once.Do(func() {
		pollFDStruct, err = gi.StructNew("GLib", "PollFD")
	})
	return err
}

type PollFD struct {
	native  uintptr
	Fd      int32
	Events  uint16
	Revents uint16
}

var privateStruct *gi.Struct
var privateStruct_Once sync.Once

func privateStruct_Set() error {
	var err error
	privateStruct_Once.Do(func() {
		privateStruct, err = gi.StructNew("GLib", "Private")
	})
	return err
}

type Private struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_private_get' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_private_replace' : parameter 'value' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_private_set' : parameter 'value' of type 'gpointer' not supported

var ptrArrayStruct *gi.Struct
var ptrArrayStruct_Once sync.Once

func ptrArrayStruct_Set() error {
	var err error
	ptrArrayStruct_Once.Do(func() {
		ptrArrayStruct, err = gi.StructNew("GLib", "PtrArray")
	})
	return err
}

type PtrArray struct {
	native uintptr
	// UNSUPPORTED : C value 'pdata' : no Go type for 'gpointer'
	Len uint32
}

var queueStruct *gi.Struct
var queueStruct_Once sync.Once

func queueStruct_Set() error {
	var err error
	queueStruct_Once.Do(func() {
		queueStruct, err = gi.StructNew("GLib", "Queue")
	})
	return err
}

type Queue struct {
	native uintptr
	// UNSUPPORTED : C value 'head' : no Go type for 'GLib.List'
	// UNSUPPORTED : C value 'tail' : no Go type for 'GLib.List'
	Length uint32
}

var queueClearFunction *gi.Function
var queueClearFunction_Once sync.Once

func queueClearFunction_Set() error {
	var err error
	queueClearFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueClearFunction, err = queueStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_queue_clear.
func (recv *Queue) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := queueClearFunction_Set()
	if err == nil {
		queueClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_queue_clear_full' : parameter 'free_func' of type 'DestroyNotify' not supported

var queueCopyFunction *gi.Function
var queueCopyFunction_Once sync.Once

func queueCopyFunction_Set() error {
	var err error
	queueCopyFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueCopyFunction, err = queueStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_queue_copy.
func (recv *Queue) Copy() (*Queue, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := queueCopyFunction_Set()
	if err == nil {
		ret = queueCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Queue{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_queue_delete_link' : parameter 'link_' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_find' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_find_custom' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_queue_foreach' : parameter 'func' of type 'Func' not supported

var queueFreeFunction *gi.Function
var queueFreeFunction_Once sync.Once

func queueFreeFunction_Set() error {
	var err error
	queueFreeFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueFreeFunction, err = queueStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_queue_free.
func (recv *Queue) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := queueFreeFunction_Set()
	if err == nil {
		queueFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_queue_free_full' : parameter 'free_func' of type 'DestroyNotify' not supported

var queueGetLengthFunction *gi.Function
var queueGetLengthFunction_Once sync.Once

func queueGetLengthFunction_Set() error {
	var err error
	queueGetLengthFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueGetLengthFunction, err = queueStruct.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type g_queue_get_length.
func (recv *Queue) GetLength() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := queueGetLengthFunction_Set()
	if err == nil {
		ret = queueGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_queue_index' : parameter 'data' of type 'gpointer' not supported

var queueInitFunction *gi.Function
var queueInitFunction_Once sync.Once

func queueInitFunction_Set() error {
	var err error
	queueInitFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueInitFunction, err = queueStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_queue_init.
func (recv *Queue) Init() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := queueInitFunction_Set()
	if err == nil {
		queueInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_queue_insert_after' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_after_link' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_before' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_before_link' : parameter 'sibling' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'g_queue_insert_sorted' : parameter 'data' of type 'gpointer' not supported

var queueIsEmptyFunction *gi.Function
var queueIsEmptyFunction_Once sync.Once

func queueIsEmptyFunction_Set() error {
	var err error
	queueIsEmptyFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueIsEmptyFunction, err = queueStruct.InvokerNew("is_empty")
	})
	return err
}

// IsEmpty is a representation of the C type g_queue_is_empty.
func (recv *Queue) IsEmpty() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := queueIsEmptyFunction_Set()
	if err == nil {
		ret = queueIsEmptyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

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

func queueReverseFunction_Set() error {
	var err error
	queueReverseFunction_Once.Do(func() {
		err = queueStruct_Set()
		if err != nil {
			return
		}
		queueReverseFunction, err = queueStruct.InvokerNew("reverse")
	})
	return err
}

// Reverse is a representation of the C type g_queue_reverse.
func (recv *Queue) Reverse() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := queueReverseFunction_Set()
	if err == nil {
		queueReverseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_queue_sort' : parameter 'compare_func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_queue_unlink' : parameter 'link_' of type 'GLib.List' not supported

var rWLockStruct *gi.Struct
var rWLockStruct_Once sync.Once

func rWLockStruct_Set() error {
	var err error
	rWLockStruct_Once.Do(func() {
		rWLockStruct, err = gi.StructNew("GLib", "RWLock")
	})
	return err
}

type RWLock struct {
	native uintptr
}

var rWLockClearFunction *gi.Function
var rWLockClearFunction_Once sync.Once

func rWLockClearFunction_Set() error {
	var err error
	rWLockClearFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockClearFunction, err = rWLockStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_rw_lock_clear.
func (recv *RWLock) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rWLockClearFunction_Set()
	if err == nil {
		rWLockClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rWLockInitFunction *gi.Function
var rWLockInitFunction_Once sync.Once

func rWLockInitFunction_Set() error {
	var err error
	rWLockInitFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockInitFunction, err = rWLockStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_rw_lock_init.
func (recv *RWLock) Init() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rWLockInitFunction_Set()
	if err == nil {
		rWLockInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rWLockReaderLockFunction *gi.Function
var rWLockReaderLockFunction_Once sync.Once

func rWLockReaderLockFunction_Set() error {
	var err error
	rWLockReaderLockFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockReaderLockFunction, err = rWLockStruct.InvokerNew("reader_lock")
	})
	return err
}

// ReaderLock is a representation of the C type g_rw_lock_reader_lock.
func (recv *RWLock) ReaderLock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rWLockReaderLockFunction_Set()
	if err == nil {
		rWLockReaderLockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rWLockReaderTrylockFunction *gi.Function
var rWLockReaderTrylockFunction_Once sync.Once

func rWLockReaderTrylockFunction_Set() error {
	var err error
	rWLockReaderTrylockFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockReaderTrylockFunction, err = rWLockStruct.InvokerNew("reader_trylock")
	})
	return err
}

// ReaderTrylock is a representation of the C type g_rw_lock_reader_trylock.
func (recv *RWLock) ReaderTrylock() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rWLockReaderTrylockFunction_Set()
	if err == nil {
		ret = rWLockReaderTrylockFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var rWLockReaderUnlockFunction *gi.Function
var rWLockReaderUnlockFunction_Once sync.Once

func rWLockReaderUnlockFunction_Set() error {
	var err error
	rWLockReaderUnlockFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockReaderUnlockFunction, err = rWLockStruct.InvokerNew("reader_unlock")
	})
	return err
}

// ReaderUnlock is a representation of the C type g_rw_lock_reader_unlock.
func (recv *RWLock) ReaderUnlock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rWLockReaderUnlockFunction_Set()
	if err == nil {
		rWLockReaderUnlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rWLockWriterLockFunction *gi.Function
var rWLockWriterLockFunction_Once sync.Once

func rWLockWriterLockFunction_Set() error {
	var err error
	rWLockWriterLockFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockWriterLockFunction, err = rWLockStruct.InvokerNew("writer_lock")
	})
	return err
}

// WriterLock is a representation of the C type g_rw_lock_writer_lock.
func (recv *RWLock) WriterLock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rWLockWriterLockFunction_Set()
	if err == nil {
		rWLockWriterLockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var rWLockWriterTrylockFunction *gi.Function
var rWLockWriterTrylockFunction_Once sync.Once

func rWLockWriterTrylockFunction_Set() error {
	var err error
	rWLockWriterTrylockFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockWriterTrylockFunction, err = rWLockStruct.InvokerNew("writer_trylock")
	})
	return err
}

// WriterTrylock is a representation of the C type g_rw_lock_writer_trylock.
func (recv *RWLock) WriterTrylock() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rWLockWriterTrylockFunction_Set()
	if err == nil {
		ret = rWLockWriterTrylockFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var rWLockWriterUnlockFunction *gi.Function
var rWLockWriterUnlockFunction_Once sync.Once

func rWLockWriterUnlockFunction_Set() error {
	var err error
	rWLockWriterUnlockFunction_Once.Do(func() {
		err = rWLockStruct_Set()
		if err != nil {
			return
		}
		rWLockWriterUnlockFunction, err = rWLockStruct.InvokerNew("writer_unlock")
	})
	return err
}

// WriterUnlock is a representation of the C type g_rw_lock_writer_unlock.
func (recv *RWLock) WriterUnlock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := rWLockWriterUnlockFunction_Set()
	if err == nil {
		rWLockWriterUnlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var randStruct *gi.Struct
var randStruct_Once sync.Once

func randStruct_Set() error {
	var err error
	randStruct_Once.Do(func() {
		randStruct, err = gi.StructNew("GLib", "Rand")
	})
	return err
}

type Rand struct {
	native uintptr
}

var randCopyFunction *gi.Function
var randCopyFunction_Once sync.Once

func randCopyFunction_Set() error {
	var err error
	randCopyFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randCopyFunction, err = randStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_rand_copy.
func (recv *Rand) Copy() (*Rand, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := randCopyFunction_Set()
	if err == nil {
		ret = randCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Rand{native: ret.Pointer()}

	return retGo, err
}

var randDoubleFunction *gi.Function
var randDoubleFunction_Once sync.Once

func randDoubleFunction_Set() error {
	var err error
	randDoubleFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randDoubleFunction, err = randStruct.InvokerNew("double")
	})
	return err
}

// Double is a representation of the C type g_rand_double.
func (recv *Rand) Double() (float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := randDoubleFunction_Set()
	if err == nil {
		ret = randDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var randDoubleRangeFunction *gi.Function
var randDoubleRangeFunction_Once sync.Once

func randDoubleRangeFunction_Set() error {
	var err error
	randDoubleRangeFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randDoubleRangeFunction, err = randStruct.InvokerNew("double_range")
	})
	return err
}

// DoubleRange is a representation of the C type g_rand_double_range.
func (recv *Rand) DoubleRange(begin float64, end float64) (float64, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(begin)
	inArgs[2].SetFloat64(end)

	var ret gi.Argument

	err := randDoubleRangeFunction_Set()
	if err == nil {
		ret = randDoubleRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var randFreeFunction *gi.Function
var randFreeFunction_Once sync.Once

func randFreeFunction_Set() error {
	var err error
	randFreeFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randFreeFunction, err = randStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_rand_free.
func (recv *Rand) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := randFreeFunction_Set()
	if err == nil {
		randFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var randIntFunction *gi.Function
var randIntFunction_Once sync.Once

func randIntFunction_Set() error {
	var err error
	randIntFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randIntFunction, err = randStruct.InvokerNew("int")
	})
	return err
}

// Int is a representation of the C type g_rand_int.
func (recv *Rand) Int() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := randIntFunction_Set()
	if err == nil {
		ret = randIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var randIntRangeFunction *gi.Function
var randIntRangeFunction_Once sync.Once

func randIntRangeFunction_Set() error {
	var err error
	randIntRangeFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randIntRangeFunction, err = randStruct.InvokerNew("int_range")
	})
	return err
}

// IntRange is a representation of the C type g_rand_int_range.
func (recv *Rand) IntRange(begin int32, end int32) (int32, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(begin)
	inArgs[2].SetInt32(end)

	var ret gi.Argument

	err := randIntRangeFunction_Set()
	if err == nil {
		ret = randIntRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var randSetSeedFunction *gi.Function
var randSetSeedFunction_Once sync.Once

func randSetSeedFunction_Set() error {
	var err error
	randSetSeedFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randSetSeedFunction, err = randStruct.InvokerNew("set_seed")
	})
	return err
}

// SetSeed is a representation of the C type g_rand_set_seed.
func (recv *Rand) SetSeed(seed uint32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(seed)

	err := randSetSeedFunction_Set()
	if err == nil {
		randSetSeedFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var randSetSeedArrayFunction *gi.Function
var randSetSeedArrayFunction_Once sync.Once

func randSetSeedArrayFunction_Set() error {
	var err error
	randSetSeedArrayFunction_Once.Do(func() {
		err = randStruct_Set()
		if err != nil {
			return
		}
		randSetSeedArrayFunction, err = randStruct.InvokerNew("set_seed_array")
	})
	return err
}

// SetSeedArray is a representation of the C type g_rand_set_seed_array.
func (recv *Rand) SetSeedArray(seed uint32, seedLength uint32) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(seed)
	inArgs[2].SetUint32(seedLength)

	err := randSetSeedArrayFunction_Set()
	if err == nil {
		randSetSeedArrayFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var recMutexStruct *gi.Struct
var recMutexStruct_Once sync.Once

func recMutexStruct_Set() error {
	var err error
	recMutexStruct_Once.Do(func() {
		recMutexStruct, err = gi.StructNew("GLib", "RecMutex")
	})
	return err
}

type RecMutex struct {
	native uintptr
}

var recMutexClearFunction *gi.Function
var recMutexClearFunction_Once sync.Once

func recMutexClearFunction_Set() error {
	var err error
	recMutexClearFunction_Once.Do(func() {
		err = recMutexStruct_Set()
		if err != nil {
			return
		}
		recMutexClearFunction, err = recMutexStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_rec_mutex_clear.
func (recv *RecMutex) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := recMutexClearFunction_Set()
	if err == nil {
		recMutexClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var recMutexInitFunction *gi.Function
var recMutexInitFunction_Once sync.Once

func recMutexInitFunction_Set() error {
	var err error
	recMutexInitFunction_Once.Do(func() {
		err = recMutexStruct_Set()
		if err != nil {
			return
		}
		recMutexInitFunction, err = recMutexStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type g_rec_mutex_init.
func (recv *RecMutex) Init() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := recMutexInitFunction_Set()
	if err == nil {
		recMutexInitFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var recMutexLockFunction *gi.Function
var recMutexLockFunction_Once sync.Once

func recMutexLockFunction_Set() error {
	var err error
	recMutexLockFunction_Once.Do(func() {
		err = recMutexStruct_Set()
		if err != nil {
			return
		}
		recMutexLockFunction, err = recMutexStruct.InvokerNew("lock")
	})
	return err
}

// Lock is a representation of the C type g_rec_mutex_lock.
func (recv *RecMutex) Lock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := recMutexLockFunction_Set()
	if err == nil {
		recMutexLockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var recMutexTrylockFunction *gi.Function
var recMutexTrylockFunction_Once sync.Once

func recMutexTrylockFunction_Set() error {
	var err error
	recMutexTrylockFunction_Once.Do(func() {
		err = recMutexStruct_Set()
		if err != nil {
			return
		}
		recMutexTrylockFunction, err = recMutexStruct.InvokerNew("trylock")
	})
	return err
}

// Trylock is a representation of the C type g_rec_mutex_trylock.
func (recv *RecMutex) Trylock() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recMutexTrylockFunction_Set()
	if err == nil {
		ret = recMutexTrylockFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var recMutexUnlockFunction *gi.Function
var recMutexUnlockFunction_Once sync.Once

func recMutexUnlockFunction_Set() error {
	var err error
	recMutexUnlockFunction_Once.Do(func() {
		err = recMutexStruct_Set()
		if err != nil {
			return
		}
		recMutexUnlockFunction, err = recMutexStruct.InvokerNew("unlock")
	})
	return err
}

// Unlock is a representation of the C type g_rec_mutex_unlock.
func (recv *RecMutex) Unlock() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := recMutexUnlockFunction_Set()
	if err == nil {
		recMutexUnlockFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var regexStruct *gi.Struct
var regexStruct_Once sync.Once

func regexStruct_Set() error {
	var err error
	regexStruct_Once.Do(func() {
		regexStruct, err = gi.StructNew("GLib", "Regex")
	})
	return err
}

type Regex struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_regex_new' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

var regexGetCaptureCountFunction *gi.Function
var regexGetCaptureCountFunction_Once sync.Once

func regexGetCaptureCountFunction_Set() error {
	var err error
	regexGetCaptureCountFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexGetCaptureCountFunction, err = regexStruct.InvokerNew("get_capture_count")
	})
	return err
}

// GetCaptureCount is a representation of the C type g_regex_get_capture_count.
func (recv *Regex) GetCaptureCount() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regexGetCaptureCountFunction_Set()
	if err == nil {
		ret = regexGetCaptureCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_regex_get_compile_flags' : return type 'RegexCompileFlags' not supported

var regexGetHasCrOrLfFunction *gi.Function
var regexGetHasCrOrLfFunction_Once sync.Once

func regexGetHasCrOrLfFunction_Set() error {
	var err error
	regexGetHasCrOrLfFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexGetHasCrOrLfFunction, err = regexStruct.InvokerNew("get_has_cr_or_lf")
	})
	return err
}

// GetHasCrOrLf is a representation of the C type g_regex_get_has_cr_or_lf.
func (recv *Regex) GetHasCrOrLf() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regexGetHasCrOrLfFunction_Set()
	if err == nil {
		ret = regexGetHasCrOrLfFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_regex_get_match_flags' : return type 'RegexMatchFlags' not supported

var regexGetMaxBackrefFunction *gi.Function
var regexGetMaxBackrefFunction_Once sync.Once

func regexGetMaxBackrefFunction_Set() error {
	var err error
	regexGetMaxBackrefFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexGetMaxBackrefFunction, err = regexStruct.InvokerNew("get_max_backref")
	})
	return err
}

// GetMaxBackref is a representation of the C type g_regex_get_max_backref.
func (recv *Regex) GetMaxBackref() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regexGetMaxBackrefFunction_Set()
	if err == nil {
		ret = regexGetMaxBackrefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var regexGetMaxLookbehindFunction *gi.Function
var regexGetMaxLookbehindFunction_Once sync.Once

func regexGetMaxLookbehindFunction_Set() error {
	var err error
	regexGetMaxLookbehindFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexGetMaxLookbehindFunction, err = regexStruct.InvokerNew("get_max_lookbehind")
	})
	return err
}

// GetMaxLookbehind is a representation of the C type g_regex_get_max_lookbehind.
func (recv *Regex) GetMaxLookbehind() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regexGetMaxLookbehindFunction_Set()
	if err == nil {
		ret = regexGetMaxLookbehindFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var regexGetPatternFunction *gi.Function
var regexGetPatternFunction_Once sync.Once

func regexGetPatternFunction_Set() error {
	var err error
	regexGetPatternFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexGetPatternFunction, err = regexStruct.InvokerNew("get_pattern")
	})
	return err
}

// GetPattern is a representation of the C type g_regex_get_pattern.
func (recv *Regex) GetPattern() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regexGetPatternFunction_Set()
	if err == nil {
		ret = regexGetPatternFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var regexGetStringNumberFunction *gi.Function
var regexGetStringNumberFunction_Once sync.Once

func regexGetStringNumberFunction_Set() error {
	var err error
	regexGetStringNumberFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexGetStringNumberFunction, err = regexStruct.InvokerNew("get_string_number")
	})
	return err
}

// GetStringNumber is a representation of the C type g_regex_get_string_number.
func (recv *Regex) GetStringNumber(name string) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := regexGetStringNumberFunction_Set()
	if err == nil {
		ret = regexGetStringNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_regex_match' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_match_all' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_match_all_full' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_match_full' : parameter 'string' has no type

var regexRefFunction *gi.Function
var regexRefFunction_Once sync.Once

func regexRefFunction_Set() error {
	var err error
	regexRefFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexRefFunction, err = regexStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_regex_ref.
func (recv *Regex) Ref() (*Regex, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := regexRefFunction_Set()
	if err == nil {
		ret = regexRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Regex{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_regex_replace' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_replace_eval' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_replace_literal' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_regex_split' : parameter 'match_options' of type 'RegexMatchFlags' not supported

// UNSUPPORTED : C value 'g_regex_split_full' : parameter 'string' has no type

var regexUnrefFunction *gi.Function
var regexUnrefFunction_Once sync.Once

func regexUnrefFunction_Set() error {
	var err error
	regexUnrefFunction_Once.Do(func() {
		err = regexStruct_Set()
		if err != nil {
			return
		}
		regexUnrefFunction, err = regexStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_regex_unref.
func (recv *Regex) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := regexUnrefFunction_Set()
	if err == nil {
		regexUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sListStruct *gi.Struct
var sListStruct_Once sync.Once

func sListStruct_Set() error {
	var err error
	sListStruct_Once.Do(func() {
		sListStruct, err = gi.StructNew("GLib", "SList")
	})
	return err
}

type SList struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.SList'
}

var scannerStruct *gi.Struct
var scannerStruct_Once sync.Once

func scannerStruct_Set() error {
	var err error
	scannerStruct_Once.Do(func() {
		scannerStruct, err = gi.StructNew("GLib", "Scanner")
	})
	return err
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

func scannerCurLineFunction_Set() error {
	var err error
	scannerCurLineFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerCurLineFunction, err = scannerStruct.InvokerNew("cur_line")
	})
	return err
}

// CurLine is a representation of the C type g_scanner_cur_line.
func (recv *Scanner) CurLine() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scannerCurLineFunction_Set()
	if err == nil {
		ret = scannerCurLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var scannerCurPositionFunction *gi.Function
var scannerCurPositionFunction_Once sync.Once

func scannerCurPositionFunction_Set() error {
	var err error
	scannerCurPositionFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerCurPositionFunction, err = scannerStruct.InvokerNew("cur_position")
	})
	return err
}

// CurPosition is a representation of the C type g_scanner_cur_position.
func (recv *Scanner) CurPosition() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scannerCurPositionFunction_Set()
	if err == nil {
		ret = scannerCurPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_scanner_cur_token' : return type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_cur_value' : return type 'TokenValue' not supported

var scannerDestroyFunction *gi.Function
var scannerDestroyFunction_Once sync.Once

func scannerDestroyFunction_Set() error {
	var err error
	scannerDestroyFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerDestroyFunction, err = scannerStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type g_scanner_destroy.
func (recv *Scanner) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scannerDestroyFunction_Set()
	if err == nil {
		scannerDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var scannerEofFunction *gi.Function
var scannerEofFunction_Once sync.Once

func scannerEofFunction_Set() error {
	var err error
	scannerEofFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerEofFunction, err = scannerStruct.InvokerNew("eof")
	})
	return err
}

// Eof is a representation of the C type g_scanner_eof.
func (recv *Scanner) Eof() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scannerEofFunction_Set()
	if err == nil {
		ret = scannerEofFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_scanner_error' : parameter '...' has no type

// UNSUPPORTED : C value 'g_scanner_get_next_token' : return type 'TokenType' not supported

var scannerInputFileFunction *gi.Function
var scannerInputFileFunction_Once sync.Once

func scannerInputFileFunction_Set() error {
	var err error
	scannerInputFileFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerInputFileFunction, err = scannerStruct.InvokerNew("input_file")
	})
	return err
}

// InputFile is a representation of the C type g_scanner_input_file.
func (recv *Scanner) InputFile(inputFd int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(inputFd)

	err := scannerInputFileFunction_Set()
	if err == nil {
		scannerInputFileFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var scannerInputTextFunction *gi.Function
var scannerInputTextFunction_Once sync.Once

func scannerInputTextFunction_Set() error {
	var err error
	scannerInputTextFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerInputTextFunction, err = scannerStruct.InvokerNew("input_text")
	})
	return err
}

// InputText is a representation of the C type g_scanner_input_text.
func (recv *Scanner) InputText(text string, textLen uint32) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetUint32(textLen)

	err := scannerInputTextFunction_Set()
	if err == nil {
		scannerInputTextFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_scanner_lookup_symbol' : return type 'gpointer' not supported

// UNSUPPORTED : C value 'g_scanner_peek_next_token' : return type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_scope_add_symbol' : parameter 'value' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_scanner_scope_foreach_symbol' : parameter 'func' of type 'HFunc' not supported

// UNSUPPORTED : C value 'g_scanner_scope_lookup_symbol' : return type 'gpointer' not supported

var scannerScopeRemoveSymbolFunction *gi.Function
var scannerScopeRemoveSymbolFunction_Once sync.Once

func scannerScopeRemoveSymbolFunction_Set() error {
	var err error
	scannerScopeRemoveSymbolFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerScopeRemoveSymbolFunction, err = scannerStruct.InvokerNew("scope_remove_symbol")
	})
	return err
}

// ScopeRemoveSymbol is a representation of the C type g_scanner_scope_remove_symbol.
func (recv *Scanner) ScopeRemoveSymbol(scopeId uint32, symbol string) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(scopeId)
	inArgs[2].SetString(symbol)

	err := scannerScopeRemoveSymbolFunction_Set()
	if err == nil {
		scannerScopeRemoveSymbolFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var scannerSetScopeFunction *gi.Function
var scannerSetScopeFunction_Once sync.Once

func scannerSetScopeFunction_Set() error {
	var err error
	scannerSetScopeFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerSetScopeFunction, err = scannerStruct.InvokerNew("set_scope")
	})
	return err
}

// SetScope is a representation of the C type g_scanner_set_scope.
func (recv *Scanner) SetScope(scopeId uint32) (uint32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(scopeId)

	var ret gi.Argument

	err := scannerSetScopeFunction_Set()
	if err == nil {
		ret = scannerSetScopeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var scannerSyncFileOffsetFunction *gi.Function
var scannerSyncFileOffsetFunction_Once sync.Once

func scannerSyncFileOffsetFunction_Set() error {
	var err error
	scannerSyncFileOffsetFunction_Once.Do(func() {
		err = scannerStruct_Set()
		if err != nil {
			return
		}
		scannerSyncFileOffsetFunction, err = scannerStruct.InvokerNew("sync_file_offset")
	})
	return err
}

// SyncFileOffset is a representation of the C type g_scanner_sync_file_offset.
func (recv *Scanner) SyncFileOffset() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scannerSyncFileOffsetFunction_Set()
	if err == nil {
		scannerSyncFileOffsetFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_scanner_unexp_token' : parameter 'expected_token' of type 'TokenType' not supported

// UNSUPPORTED : C value 'g_scanner_warn' : parameter '...' has no type

var scannerConfigStruct *gi.Struct
var scannerConfigStruct_Once sync.Once

func scannerConfigStruct_Set() error {
	var err error
	scannerConfigStruct_Once.Do(func() {
		scannerConfigStruct, err = gi.StructNew("GLib", "ScannerConfig")
	})
	return err
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
var sequenceStruct_Once sync.Once

func sequenceStruct_Set() error {
	var err error
	sequenceStruct_Once.Do(func() {
		sequenceStruct, err = gi.StructNew("GLib", "Sequence")
	})
	return err
}

type Sequence struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_sequence_append' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_foreach' : parameter 'func' of type 'Func' not supported

var sequenceFreeFunction *gi.Function
var sequenceFreeFunction_Once sync.Once

func sequenceFreeFunction_Set() error {
	var err error
	sequenceFreeFunction_Once.Do(func() {
		err = sequenceStruct_Set()
		if err != nil {
			return
		}
		sequenceFreeFunction, err = sequenceStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_sequence_free.
func (recv *Sequence) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := sequenceFreeFunction_Set()
	if err == nil {
		sequenceFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sequenceGetBeginIterFunction *gi.Function
var sequenceGetBeginIterFunction_Once sync.Once

func sequenceGetBeginIterFunction_Set() error {
	var err error
	sequenceGetBeginIterFunction_Once.Do(func() {
		err = sequenceStruct_Set()
		if err != nil {
			return
		}
		sequenceGetBeginIterFunction, err = sequenceStruct.InvokerNew("get_begin_iter")
	})
	return err
}

// GetBeginIter is a representation of the C type g_sequence_get_begin_iter.
func (recv *Sequence) GetBeginIter() (*SequenceIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceGetBeginIterFunction_Set()
	if err == nil {
		ret = sequenceGetBeginIterFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo, err
}

var sequenceGetEndIterFunction *gi.Function
var sequenceGetEndIterFunction_Once sync.Once

func sequenceGetEndIterFunction_Set() error {
	var err error
	sequenceGetEndIterFunction_Once.Do(func() {
		err = sequenceStruct_Set()
		if err != nil {
			return
		}
		sequenceGetEndIterFunction, err = sequenceStruct.InvokerNew("get_end_iter")
	})
	return err
}

// GetEndIter is a representation of the C type g_sequence_get_end_iter.
func (recv *Sequence) GetEndIter() (*SequenceIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceGetEndIterFunction_Set()
	if err == nil {
		ret = sequenceGetEndIterFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo, err
}

var sequenceGetIterAtPosFunction *gi.Function
var sequenceGetIterAtPosFunction_Once sync.Once

func sequenceGetIterAtPosFunction_Set() error {
	var err error
	sequenceGetIterAtPosFunction_Once.Do(func() {
		err = sequenceStruct_Set()
		if err != nil {
			return
		}
		sequenceGetIterAtPosFunction, err = sequenceStruct.InvokerNew("get_iter_at_pos")
	})
	return err
}

// GetIterAtPos is a representation of the C type g_sequence_get_iter_at_pos.
func (recv *Sequence) GetIterAtPos(pos int32) (*SequenceIter, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	var ret gi.Argument

	err := sequenceGetIterAtPosFunction_Set()
	if err == nil {
		ret = sequenceGetIterAtPosFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo, err
}

var sequenceGetLengthFunction *gi.Function
var sequenceGetLengthFunction_Once sync.Once

func sequenceGetLengthFunction_Set() error {
	var err error
	sequenceGetLengthFunction_Once.Do(func() {
		err = sequenceStruct_Set()
		if err != nil {
			return
		}
		sequenceGetLengthFunction, err = sequenceStruct.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type g_sequence_get_length.
func (recv *Sequence) GetLength() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceGetLengthFunction_Set()
	if err == nil {
		ret = sequenceGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_sequence_insert_sorted' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_insert_sorted_iter' : parameter 'data' of type 'gpointer' not supported

var sequenceIsEmptyFunction *gi.Function
var sequenceIsEmptyFunction_Once sync.Once

func sequenceIsEmptyFunction_Set() error {
	var err error
	sequenceIsEmptyFunction_Once.Do(func() {
		err = sequenceStruct_Set()
		if err != nil {
			return
		}
		sequenceIsEmptyFunction, err = sequenceStruct.InvokerNew("is_empty")
	})
	return err
}

// IsEmpty is a representation of the C type g_sequence_is_empty.
func (recv *Sequence) IsEmpty() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIsEmptyFunction_Set()
	if err == nil {
		ret = sequenceIsEmptyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_sequence_lookup' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_lookup_iter' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_prepend' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_search' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_search_iter' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_sequence_sort' : parameter 'cmp_func' of type 'CompareDataFunc' not supported

// UNSUPPORTED : C value 'g_sequence_sort_iter' : parameter 'cmp_func' of type 'SequenceIterCompareFunc' not supported

var sequenceIterStruct *gi.Struct
var sequenceIterStruct_Once sync.Once

func sequenceIterStruct_Set() error {
	var err error
	sequenceIterStruct_Once.Do(func() {
		sequenceIterStruct, err = gi.StructNew("GLib", "SequenceIter")
	})
	return err
}

type SequenceIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_sequence_iter_compare' : parameter 'b' of type 'SequenceIter' not supported

var sequenceIterGetPositionFunction *gi.Function
var sequenceIterGetPositionFunction_Once sync.Once

func sequenceIterGetPositionFunction_Set() error {
	var err error
	sequenceIterGetPositionFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterGetPositionFunction, err = sequenceIterStruct.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type g_sequence_iter_get_position.
func (recv *SequenceIter) GetPosition() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIterGetPositionFunction_Set()
	if err == nil {
		ret = sequenceIterGetPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var sequenceIterGetSequenceFunction *gi.Function
var sequenceIterGetSequenceFunction_Once sync.Once

func sequenceIterGetSequenceFunction_Set() error {
	var err error
	sequenceIterGetSequenceFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterGetSequenceFunction, err = sequenceIterStruct.InvokerNew("get_sequence")
	})
	return err
}

// GetSequence is a representation of the C type g_sequence_iter_get_sequence.
func (recv *SequenceIter) GetSequence() (*Sequence, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIterGetSequenceFunction_Set()
	if err == nil {
		ret = sequenceIterGetSequenceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Sequence{native: ret.Pointer()}

	return retGo, err
}

var sequenceIterIsBeginFunction *gi.Function
var sequenceIterIsBeginFunction_Once sync.Once

func sequenceIterIsBeginFunction_Set() error {
	var err error
	sequenceIterIsBeginFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterIsBeginFunction, err = sequenceIterStruct.InvokerNew("is_begin")
	})
	return err
}

// IsBegin is a representation of the C type g_sequence_iter_is_begin.
func (recv *SequenceIter) IsBegin() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIterIsBeginFunction_Set()
	if err == nil {
		ret = sequenceIterIsBeginFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var sequenceIterIsEndFunction *gi.Function
var sequenceIterIsEndFunction_Once sync.Once

func sequenceIterIsEndFunction_Set() error {
	var err error
	sequenceIterIsEndFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterIsEndFunction, err = sequenceIterStruct.InvokerNew("is_end")
	})
	return err
}

// IsEnd is a representation of the C type g_sequence_iter_is_end.
func (recv *SequenceIter) IsEnd() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIterIsEndFunction_Set()
	if err == nil {
		ret = sequenceIterIsEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var sequenceIterMoveFunction *gi.Function
var sequenceIterMoveFunction_Once sync.Once

func sequenceIterMoveFunction_Set() error {
	var err error
	sequenceIterMoveFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterMoveFunction, err = sequenceIterStruct.InvokerNew("move")
	})
	return err
}

// Move is a representation of the C type g_sequence_iter_move.
func (recv *SequenceIter) Move(delta int32) (*SequenceIter, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(delta)

	var ret gi.Argument

	err := sequenceIterMoveFunction_Set()
	if err == nil {
		ret = sequenceIterMoveFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo, err
}

var sequenceIterNextFunction *gi.Function
var sequenceIterNextFunction_Once sync.Once

func sequenceIterNextFunction_Set() error {
	var err error
	sequenceIterNextFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterNextFunction, err = sequenceIterStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type g_sequence_iter_next.
func (recv *SequenceIter) Next() (*SequenceIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIterNextFunction_Set()
	if err == nil {
		ret = sequenceIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo, err
}

var sequenceIterPrevFunction *gi.Function
var sequenceIterPrevFunction_Once sync.Once

func sequenceIterPrevFunction_Set() error {
	var err error
	sequenceIterPrevFunction_Once.Do(func() {
		err = sequenceIterStruct_Set()
		if err != nil {
			return
		}
		sequenceIterPrevFunction, err = sequenceIterStruct.InvokerNew("prev")
	})
	return err
}

// Prev is a representation of the C type g_sequence_iter_prev.
func (recv *SequenceIter) Prev() (*SequenceIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sequenceIterPrevFunction_Set()
	if err == nil {
		ret = sequenceIterPrevFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SequenceIter{native: ret.Pointer()}

	return retGo, err
}

var sourceStruct *gi.Struct
var sourceStruct_Once sync.Once

func sourceStruct_Set() error {
	var err error
	sourceStruct_Once.Do(func() {
		sourceStruct, err = gi.StructNew("GLib", "Source")
	})
	return err
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

func sourceDestroyFunction_Set() error {
	var err error
	sourceDestroyFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceDestroyFunction, err = sourceStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type g_source_destroy.
func (recv *Source) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := sourceDestroyFunction_Set()
	if err == nil {
		sourceDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sourceGetCanRecurseFunction *gi.Function
var sourceGetCanRecurseFunction_Once sync.Once

func sourceGetCanRecurseFunction_Set() error {
	var err error
	sourceGetCanRecurseFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetCanRecurseFunction, err = sourceStruct.InvokerNew("get_can_recurse")
	})
	return err
}

// GetCanRecurse is a representation of the C type g_source_get_can_recurse.
func (recv *Source) GetCanRecurse() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetCanRecurseFunction_Set()
	if err == nil {
		ret = sourceGetCanRecurseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var sourceGetContextFunction *gi.Function
var sourceGetContextFunction_Once sync.Once

func sourceGetContextFunction_Set() error {
	var err error
	sourceGetContextFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetContextFunction, err = sourceStruct.InvokerNew("get_context")
	})
	return err
}

// GetContext is a representation of the C type g_source_get_context.
func (recv *Source) GetContext() (*MainContext, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetContextFunction_Set()
	if err == nil {
		ret = sourceGetContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &MainContext{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_source_get_current_time' : parameter 'timeval' of type 'TimeVal' not supported

var sourceGetIdFunction *gi.Function
var sourceGetIdFunction_Once sync.Once

func sourceGetIdFunction_Set() error {
	var err error
	sourceGetIdFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetIdFunction, err = sourceStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type g_source_get_id.
func (recv *Source) GetId() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetIdFunction_Set()
	if err == nil {
		ret = sourceGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var sourceGetNameFunction *gi.Function
var sourceGetNameFunction_Once sync.Once

func sourceGetNameFunction_Set() error {
	var err error
	sourceGetNameFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetNameFunction, err = sourceStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type g_source_get_name.
func (recv *Source) GetName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetNameFunction_Set()
	if err == nil {
		ret = sourceGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var sourceGetPriorityFunction *gi.Function
var sourceGetPriorityFunction_Once sync.Once

func sourceGetPriorityFunction_Set() error {
	var err error
	sourceGetPriorityFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetPriorityFunction, err = sourceStruct.InvokerNew("get_priority")
	})
	return err
}

// GetPriority is a representation of the C type g_source_get_priority.
func (recv *Source) GetPriority() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetPriorityFunction_Set()
	if err == nil {
		ret = sourceGetPriorityFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var sourceGetReadyTimeFunction *gi.Function
var sourceGetReadyTimeFunction_Once sync.Once

func sourceGetReadyTimeFunction_Set() error {
	var err error
	sourceGetReadyTimeFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetReadyTimeFunction, err = sourceStruct.InvokerNew("get_ready_time")
	})
	return err
}

// GetReadyTime is a representation of the C type g_source_get_ready_time.
func (recv *Source) GetReadyTime() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetReadyTimeFunction_Set()
	if err == nil {
		ret = sourceGetReadyTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var sourceGetTimeFunction *gi.Function
var sourceGetTimeFunction_Once sync.Once

func sourceGetTimeFunction_Set() error {
	var err error
	sourceGetTimeFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceGetTimeFunction, err = sourceStruct.InvokerNew("get_time")
	})
	return err
}

// GetTime is a representation of the C type g_source_get_time.
func (recv *Source) GetTime() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceGetTimeFunction_Set()
	if err == nil {
		ret = sourceGetTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var sourceIsDestroyedFunction *gi.Function
var sourceIsDestroyedFunction_Once sync.Once

func sourceIsDestroyedFunction_Set() error {
	var err error
	sourceIsDestroyedFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceIsDestroyedFunction, err = sourceStruct.InvokerNew("is_destroyed")
	})
	return err
}

// IsDestroyed is a representation of the C type g_source_is_destroyed.
func (recv *Source) IsDestroyed() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceIsDestroyedFunction_Set()
	if err == nil {
		ret = sourceIsDestroyedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_source_modify_unix_fd' : parameter 'tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_source_query_unix_fd' : parameter 'tag' of type 'gpointer' not supported

var sourceRefFunction *gi.Function
var sourceRefFunction_Once sync.Once

func sourceRefFunction_Set() error {
	var err error
	sourceRefFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceRefFunction, err = sourceStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_source_ref.
func (recv *Source) Ref() (*Source, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := sourceRefFunction_Set()
	if err == nil {
		ret = sourceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Source{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_source_remove_child_source' : parameter 'child_source' of type 'Source' not supported

// UNSUPPORTED : C value 'g_source_remove_poll' : parameter 'fd' of type 'PollFD' not supported

// UNSUPPORTED : C value 'g_source_remove_unix_fd' : parameter 'tag' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_source_set_callback' : parameter 'func' of type 'SourceFunc' not supported

// UNSUPPORTED : C value 'g_source_set_callback_indirect' : parameter 'callback_data' of type 'gpointer' not supported

var sourceSetCanRecurseFunction *gi.Function
var sourceSetCanRecurseFunction_Once sync.Once

func sourceSetCanRecurseFunction_Set() error {
	var err error
	sourceSetCanRecurseFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceSetCanRecurseFunction, err = sourceStruct.InvokerNew("set_can_recurse")
	})
	return err
}

// SetCanRecurse is a representation of the C type g_source_set_can_recurse.
func (recv *Source) SetCanRecurse(canRecurse bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(canRecurse)

	err := sourceSetCanRecurseFunction_Set()
	if err == nil {
		sourceSetCanRecurseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_source_set_funcs' : parameter 'funcs' of type 'SourceFuncs' not supported

var sourceSetNameFunction *gi.Function
var sourceSetNameFunction_Once sync.Once

func sourceSetNameFunction_Set() error {
	var err error
	sourceSetNameFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceSetNameFunction, err = sourceStruct.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type g_source_set_name.
func (recv *Source) SetName(name string) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	err := sourceSetNameFunction_Set()
	if err == nil {
		sourceSetNameFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sourceSetPriorityFunction *gi.Function
var sourceSetPriorityFunction_Once sync.Once

func sourceSetPriorityFunction_Set() error {
	var err error
	sourceSetPriorityFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceSetPriorityFunction, err = sourceStruct.InvokerNew("set_priority")
	})
	return err
}

// SetPriority is a representation of the C type g_source_set_priority.
func (recv *Source) SetPriority(priority int32) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(priority)

	err := sourceSetPriorityFunction_Set()
	if err == nil {
		sourceSetPriorityFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sourceSetReadyTimeFunction *gi.Function
var sourceSetReadyTimeFunction_Once sync.Once

func sourceSetReadyTimeFunction_Set() error {
	var err error
	sourceSetReadyTimeFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceSetReadyTimeFunction, err = sourceStruct.InvokerNew("set_ready_time")
	})
	return err
}

// SetReadyTime is a representation of the C type g_source_set_ready_time.
func (recv *Source) SetReadyTime(readyTime int64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(readyTime)

	err := sourceSetReadyTimeFunction_Set()
	if err == nil {
		sourceSetReadyTimeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sourceUnrefFunction *gi.Function
var sourceUnrefFunction_Once sync.Once

func sourceUnrefFunction_Set() error {
	var err error
	sourceUnrefFunction_Once.Do(func() {
		err = sourceStruct_Set()
		if err != nil {
			return
		}
		sourceUnrefFunction, err = sourceStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_source_unref.
func (recv *Source) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := sourceUnrefFunction_Set()
	if err == nil {
		sourceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var sourceCallbackFuncsStruct *gi.Struct
var sourceCallbackFuncsStruct_Once sync.Once

func sourceCallbackFuncsStruct_Set() error {
	var err error
	sourceCallbackFuncsStruct_Once.Do(func() {
		sourceCallbackFuncsStruct, err = gi.StructNew("GLib", "SourceCallbackFuncs")
	})
	return err
}

type SourceCallbackFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'ref' : missing Type
	// UNSUPPORTED : C value 'unref' : missing Type
	// UNSUPPORTED : C value 'get' : missing Type
}

var sourceFuncsStruct *gi.Struct
var sourceFuncsStruct_Once sync.Once

func sourceFuncsStruct_Set() error {
	var err error
	sourceFuncsStruct_Once.Do(func() {
		sourceFuncsStruct, err = gi.StructNew("GLib", "SourceFuncs")
	})
	return err
}

type SourceFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'prepare' : missing Type
	// UNSUPPORTED : C value 'check' : missing Type
	// UNSUPPORTED : C value 'dispatch' : missing Type
	// UNSUPPORTED : C value 'finalize' : missing Type
}

var sourcePrivateStruct *gi.Struct
var sourcePrivateStruct_Once sync.Once

func sourcePrivateStruct_Set() error {
	var err error
	sourcePrivateStruct_Once.Do(func() {
		sourcePrivateStruct, err = gi.StructNew("GLib", "SourcePrivate")
	})
	return err
}

type SourcePrivate struct {
	native uintptr
}

var statBufStruct *gi.Struct
var statBufStruct_Once sync.Once

func statBufStruct_Set() error {
	var err error
	statBufStruct_Once.Do(func() {
		statBufStruct, err = gi.StructNew("GLib", "StatBuf")
	})
	return err
}

type StatBuf struct {
	native uintptr
}

var string_Struct *gi.Struct
var string_Struct_Once sync.Once

func string_Struct_Set() error {
	var err error
	string_Struct_Once.Do(func() {
		string_Struct, err = gi.StructNew("GLib", "String")
	})
	return err
}

type String struct {
	native       uintptr
	Str          string
	Len          uintptr
	AllocatedLen uintptr
}

var stringAppendFunction *gi.Function
var stringAppendFunction_Once sync.Once

func stringAppendFunction_Set() error {
	var err error
	stringAppendFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAppendFunction, err = string_Struct.InvokerNew("append")
	})
	return err
}

// Append is a representation of the C type g_string_append.
func (recv *String) Append(val string) (*String, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)

	var ret gi.Argument

	err := stringAppendFunction_Set()
	if err == nil {
		ret = stringAppendFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringAppendCFunction *gi.Function
var stringAppendCFunction_Once sync.Once

func stringAppendCFunction_Set() error {
	var err error
	stringAppendCFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAppendCFunction, err = string_Struct.InvokerNew("append_c")
	})
	return err
}

// AppendC is a representation of the C type g_string_append_c.
func (recv *String) AppendC(c int8) (*String, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(c)

	var ret gi.Argument

	err := stringAppendCFunction_Set()
	if err == nil {
		ret = stringAppendCFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringAppendLenFunction *gi.Function
var stringAppendLenFunction_Once sync.Once

func stringAppendLenFunction_Set() error {
	var err error
	stringAppendLenFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAppendLenFunction, err = string_Struct.InvokerNew("append_len")
	})
	return err
}

// AppendLen is a representation of the C type g_string_append_len.
func (recv *String) AppendLen(val string, len int32) (*String, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)
	inArgs[2].SetInt32(len)

	var ret gi.Argument

	err := stringAppendLenFunction_Set()
	if err == nil {
		ret = stringAppendLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_append_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_string_append_unichar' : parameter 'wc' of type 'gunichar' not supported

var stringAppendUriEscapedFunction *gi.Function
var stringAppendUriEscapedFunction_Once sync.Once

func stringAppendUriEscapedFunction_Set() error {
	var err error
	stringAppendUriEscapedFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAppendUriEscapedFunction, err = string_Struct.InvokerNew("append_uri_escaped")
	})
	return err
}

// AppendUriEscaped is a representation of the C type g_string_append_uri_escaped.
func (recv *String) AppendUriEscaped(unescaped string, reservedCharsAllowed string, allowUtf8 bool) (*String, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(unescaped)
	inArgs[2].SetString(reservedCharsAllowed)
	inArgs[3].SetBoolean(allowUtf8)

	var ret gi.Argument

	err := stringAppendUriEscapedFunction_Set()
	if err == nil {
		ret = stringAppendUriEscapedFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_append_vprintf' : parameter 'args' of type 'va_list' not supported

var stringAsciiDownFunction *gi.Function
var stringAsciiDownFunction_Once sync.Once

func stringAsciiDownFunction_Set() error {
	var err error
	stringAsciiDownFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAsciiDownFunction, err = string_Struct.InvokerNew("ascii_down")
	})
	return err
}

// AsciiDown is a representation of the C type g_string_ascii_down.
func (recv *String) AsciiDown() (*String, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringAsciiDownFunction_Set()
	if err == nil {
		ret = stringAsciiDownFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringAsciiUpFunction *gi.Function
var stringAsciiUpFunction_Once sync.Once

func stringAsciiUpFunction_Set() error {
	var err error
	stringAsciiUpFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAsciiUpFunction, err = string_Struct.InvokerNew("ascii_up")
	})
	return err
}

// AsciiUp is a representation of the C type g_string_ascii_up.
func (recv *String) AsciiUp() (*String, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringAsciiUpFunction_Set()
	if err == nil {
		ret = stringAsciiUpFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringAssignFunction *gi.Function
var stringAssignFunction_Once sync.Once

func stringAssignFunction_Set() error {
	var err error
	stringAssignFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringAssignFunction, err = string_Struct.InvokerNew("assign")
	})
	return err
}

// Assign is a representation of the C type g_string_assign.
func (recv *String) Assign(rval string) (*String, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(rval)

	var ret gi.Argument

	err := stringAssignFunction_Set()
	if err == nil {
		ret = stringAssignFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringDownFunction *gi.Function
var stringDownFunction_Once sync.Once

func stringDownFunction_Set() error {
	var err error
	stringDownFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringDownFunction, err = string_Struct.InvokerNew("down")
	})
	return err
}

// Down is a representation of the C type g_string_down.
func (recv *String) Down() (*String, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringDownFunction_Set()
	if err == nil {
		ret = stringDownFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_equal' : parameter 'v2' of type 'String' not supported

var stringEraseFunction *gi.Function
var stringEraseFunction_Once sync.Once

func stringEraseFunction_Set() error {
	var err error
	stringEraseFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringEraseFunction, err = string_Struct.InvokerNew("erase")
	})
	return err
}

// Erase is a representation of the C type g_string_erase.
func (recv *String) Erase(pos int32, len int32) (*String, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetInt32(len)

	var ret gi.Argument

	err := stringEraseFunction_Set()
	if err == nil {
		ret = stringEraseFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringFreeFunction *gi.Function
var stringFreeFunction_Once sync.Once

func stringFreeFunction_Set() error {
	var err error
	stringFreeFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringFreeFunction, err = string_Struct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_string_free.
func (recv *String) Free(freeSegment bool) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(freeSegment)

	var ret gi.Argument

	err := stringFreeFunction_Set()
	if err == nil {
		ret = stringFreeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var stringFreeToBytesFunction *gi.Function
var stringFreeToBytesFunction_Once sync.Once

func stringFreeToBytesFunction_Set() error {
	var err error
	stringFreeToBytesFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringFreeToBytesFunction, err = string_Struct.InvokerNew("free_to_bytes")
	})
	return err
}

// FreeToBytes is a representation of the C type g_string_free_to_bytes.
func (recv *String) FreeToBytes() (*Bytes, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringFreeToBytesFunction_Set()
	if err == nil {
		ret = stringFreeToBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Bytes{native: ret.Pointer()}

	return retGo, err
}

var stringHashFunction *gi.Function
var stringHashFunction_Once sync.Once

func stringHashFunction_Set() error {
	var err error
	stringHashFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringHashFunction, err = string_Struct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type g_string_hash.
func (recv *String) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringHashFunction_Set()
	if err == nil {
		ret = stringHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var stringInsertFunction *gi.Function
var stringInsertFunction_Once sync.Once

func stringInsertFunction_Set() error {
	var err error
	stringInsertFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringInsertFunction, err = string_Struct.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type g_string_insert.
func (recv *String) Insert(pos int32, val string) (*String, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(val)

	var ret gi.Argument

	err := stringInsertFunction_Set()
	if err == nil {
		ret = stringInsertFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringInsertCFunction *gi.Function
var stringInsertCFunction_Once sync.Once

func stringInsertCFunction_Set() error {
	var err error
	stringInsertCFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringInsertCFunction, err = string_Struct.InvokerNew("insert_c")
	})
	return err
}

// InsertC is a representation of the C type g_string_insert_c.
func (recv *String) InsertC(pos int32, c int8) (*String, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetInt8(c)

	var ret gi.Argument

	err := stringInsertCFunction_Set()
	if err == nil {
		ret = stringInsertCFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringInsertLenFunction *gi.Function
var stringInsertLenFunction_Once sync.Once

func stringInsertLenFunction_Set() error {
	var err error
	stringInsertLenFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringInsertLenFunction, err = string_Struct.InvokerNew("insert_len")
	})
	return err
}

// InsertLen is a representation of the C type g_string_insert_len.
func (recv *String) InsertLen(pos int32, val string, len int32) (*String, error) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(val)
	inArgs[3].SetInt32(len)

	var ret gi.Argument

	err := stringInsertLenFunction_Set()
	if err == nil {
		ret = stringInsertLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_insert_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_overwrite' : parameter 'pos' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_string_overwrite_len' : parameter 'pos' of type 'gsize' not supported

var stringPrependFunction *gi.Function
var stringPrependFunction_Once sync.Once

func stringPrependFunction_Set() error {
	var err error
	stringPrependFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringPrependFunction, err = string_Struct.InvokerNew("prepend")
	})
	return err
}

// Prepend is a representation of the C type g_string_prepend.
func (recv *String) Prepend(val string) (*String, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)

	var ret gi.Argument

	err := stringPrependFunction_Set()
	if err == nil {
		ret = stringPrependFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringPrependCFunction *gi.Function
var stringPrependCFunction_Once sync.Once

func stringPrependCFunction_Set() error {
	var err error
	stringPrependCFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringPrependCFunction, err = string_Struct.InvokerNew("prepend_c")
	})
	return err
}

// PrependC is a representation of the C type g_string_prepend_c.
func (recv *String) PrependC(c int8) (*String, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt8(c)

	var ret gi.Argument

	err := stringPrependCFunction_Set()
	if err == nil {
		ret = stringPrependCFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

var stringPrependLenFunction *gi.Function
var stringPrependLenFunction_Once sync.Once

func stringPrependLenFunction_Set() error {
	var err error
	stringPrependLenFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringPrependLenFunction, err = string_Struct.InvokerNew("prepend_len")
	})
	return err
}

// PrependLen is a representation of the C type g_string_prepend_len.
func (recv *String) PrependLen(val string, len int32) (*String, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(val)
	inArgs[2].SetInt32(len)

	var ret gi.Argument

	err := stringPrependLenFunction_Set()
	if err == nil {
		ret = stringPrependLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_prepend_unichar' : parameter 'wc' of type 'gunichar' not supported

// UNSUPPORTED : C value 'g_string_printf' : parameter '...' has no type

// UNSUPPORTED : C value 'g_string_set_size' : parameter 'len' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_string_truncate' : parameter 'len' of type 'gsize' not supported

var stringUpFunction *gi.Function
var stringUpFunction_Once sync.Once

func stringUpFunction_Set() error {
	var err error
	stringUpFunction_Once.Do(func() {
		err = string_Struct_Set()
		if err != nil {
			return
		}
		stringUpFunction, err = string_Struct.InvokerNew("up")
	})
	return err
}

// Up is a representation of the C type g_string_up.
func (recv *String) Up() (*String, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stringUpFunction_Set()
	if err == nil {
		ret = stringUpFunction.Invoke(inArgs[:], nil)
	}

	retGo := &String{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_string_vprintf' : parameter 'args' of type 'va_list' not supported

var stringChunkStruct *gi.Struct
var stringChunkStruct_Once sync.Once

func stringChunkStruct_Set() error {
	var err error
	stringChunkStruct_Once.Do(func() {
		stringChunkStruct, err = gi.StructNew("GLib", "StringChunk")
	})
	return err
}

type StringChunk struct {
	native uintptr
}

var stringChunkClearFunction *gi.Function
var stringChunkClearFunction_Once sync.Once

func stringChunkClearFunction_Set() error {
	var err error
	stringChunkClearFunction_Once.Do(func() {
		err = stringChunkStruct_Set()
		if err != nil {
			return
		}
		stringChunkClearFunction, err = stringChunkStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_string_chunk_clear.
func (recv *StringChunk) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := stringChunkClearFunction_Set()
	if err == nil {
		stringChunkClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var stringChunkFreeFunction *gi.Function
var stringChunkFreeFunction_Once sync.Once

func stringChunkFreeFunction_Set() error {
	var err error
	stringChunkFreeFunction_Once.Do(func() {
		err = stringChunkStruct_Set()
		if err != nil {
			return
		}
		stringChunkFreeFunction, err = stringChunkStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_string_chunk_free.
func (recv *StringChunk) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := stringChunkFreeFunction_Set()
	if err == nil {
		stringChunkFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var stringChunkInsertFunction *gi.Function
var stringChunkInsertFunction_Once sync.Once

func stringChunkInsertFunction_Set() error {
	var err error
	stringChunkInsertFunction_Once.Do(func() {
		err = stringChunkStruct_Set()
		if err != nil {
			return
		}
		stringChunkInsertFunction, err = stringChunkStruct.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type g_string_chunk_insert.
func (recv *StringChunk) Insert(string_ string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)

	var ret gi.Argument

	err := stringChunkInsertFunction_Set()
	if err == nil {
		ret = stringChunkInsertFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var stringChunkInsertConstFunction *gi.Function
var stringChunkInsertConstFunction_Once sync.Once

func stringChunkInsertConstFunction_Set() error {
	var err error
	stringChunkInsertConstFunction_Once.Do(func() {
		err = stringChunkStruct_Set()
		if err != nil {
			return
		}
		stringChunkInsertConstFunction, err = stringChunkStruct.InvokerNew("insert_const")
	})
	return err
}

// InsertConst is a representation of the C type g_string_chunk_insert_const.
func (recv *StringChunk) InsertConst(string_ string) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)

	var ret gi.Argument

	err := stringChunkInsertConstFunction_Set()
	if err == nil {
		ret = stringChunkInsertConstFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var stringChunkInsertLenFunction *gi.Function
var stringChunkInsertLenFunction_Once sync.Once

func stringChunkInsertLenFunction_Set() error {
	var err error
	stringChunkInsertLenFunction_Once.Do(func() {
		err = stringChunkStruct_Set()
		if err != nil {
			return
		}
		stringChunkInsertLenFunction, err = stringChunkStruct.InvokerNew("insert_len")
	})
	return err
}

// InsertLen is a representation of the C type g_string_chunk_insert_len.
func (recv *StringChunk) InsertLen(string_ string, len int32) (string, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(string_)
	inArgs[2].SetInt32(len)

	var ret gi.Argument

	err := stringChunkInsertLenFunction_Set()
	if err == nil {
		ret = stringChunkInsertLenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var testCaseStruct *gi.Struct
var testCaseStruct_Once sync.Once

func testCaseStruct_Set() error {
	var err error
	testCaseStruct_Once.Do(func() {
		testCaseStruct, err = gi.StructNew("GLib", "TestCase")
	})
	return err
}

type TestCase struct {
	native uintptr
}

var testConfigStruct *gi.Struct
var testConfigStruct_Once sync.Once

func testConfigStruct_Set() error {
	var err error
	testConfigStruct_Once.Do(func() {
		testConfigStruct, err = gi.StructNew("GLib", "TestConfig")
	})
	return err
}

type TestConfig struct {
	native          uintptr
	TestInitialized bool
	TestQuick       bool
	TestPerf        bool
	TestVerbose     bool
	TestQuiet       bool
	TestUndefined   bool
}

var testLogBufferStruct *gi.Struct
var testLogBufferStruct_Once sync.Once

func testLogBufferStruct_Set() error {
	var err error
	testLogBufferStruct_Once.Do(func() {
		testLogBufferStruct, err = gi.StructNew("GLib", "TestLogBuffer")
	})
	return err
}

type TestLogBuffer struct {
	native uintptr
}

var testLogBufferFreeFunction *gi.Function
var testLogBufferFreeFunction_Once sync.Once

func testLogBufferFreeFunction_Set() error {
	var err error
	testLogBufferFreeFunction_Once.Do(func() {
		err = testLogBufferStruct_Set()
		if err != nil {
			return
		}
		testLogBufferFreeFunction, err = testLogBufferStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_test_log_buffer_free.
func (recv *TestLogBuffer) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := testLogBufferFreeFunction_Set()
	if err == nil {
		testLogBufferFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testLogBufferPopFunction *gi.Function
var testLogBufferPopFunction_Once sync.Once

func testLogBufferPopFunction_Set() error {
	var err error
	testLogBufferPopFunction_Once.Do(func() {
		err = testLogBufferStruct_Set()
		if err != nil {
			return
		}
		testLogBufferPopFunction, err = testLogBufferStruct.InvokerNew("pop")
	})
	return err
}

// Pop is a representation of the C type g_test_log_buffer_pop.
func (recv *TestLogBuffer) Pop() (*TestLogMsg, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := testLogBufferPopFunction_Set()
	if err == nil {
		ret = testLogBufferPopFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TestLogMsg{native: ret.Pointer()}

	return retGo, err
}

var testLogBufferPushFunction *gi.Function
var testLogBufferPushFunction_Once sync.Once

func testLogBufferPushFunction_Set() error {
	var err error
	testLogBufferPushFunction_Once.Do(func() {
		err = testLogBufferStruct_Set()
		if err != nil {
			return
		}
		testLogBufferPushFunction, err = testLogBufferStruct.InvokerNew("push")
	})
	return err
}

// Push is a representation of the C type g_test_log_buffer_push.
func (recv *TestLogBuffer) Push(nBytes uint32, bytes uint8) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(nBytes)
	inArgs[2].SetUint8(bytes)

	err := testLogBufferPushFunction_Set()
	if err == nil {
		testLogBufferPushFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testLogMsgStruct *gi.Struct
var testLogMsgStruct_Once sync.Once

func testLogMsgStruct_Set() error {
	var err error
	testLogMsgStruct_Once.Do(func() {
		testLogMsgStruct, err = gi.StructNew("GLib", "TestLogMsg")
	})
	return err
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

func testLogMsgFreeFunction_Set() error {
	var err error
	testLogMsgFreeFunction_Once.Do(func() {
		err = testLogMsgStruct_Set()
		if err != nil {
			return
		}
		testLogMsgFreeFunction, err = testLogMsgStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_test_log_msg_free.
func (recv *TestLogMsg) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := testLogMsgFreeFunction_Set()
	if err == nil {
		testLogMsgFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var testSuiteStruct *gi.Struct
var testSuiteStruct_Once sync.Once

func testSuiteStruct_Set() error {
	var err error
	testSuiteStruct_Once.Do(func() {
		testSuiteStruct, err = gi.StructNew("GLib", "TestSuite")
	})
	return err
}

type TestSuite struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_test_suite_add' : parameter 'test_case' of type 'TestCase' not supported

// UNSUPPORTED : C value 'g_test_suite_add_suite' : parameter 'nestedsuite' of type 'TestSuite' not supported

var threadStruct *gi.Struct
var threadStruct_Once sync.Once

func threadStruct_Set() error {
	var err error
	threadStruct_Once.Do(func() {
		threadStruct, err = gi.StructNew("GLib", "Thread")
	})
	return err
}

type Thread struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_thread_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_try_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_join' : return type 'gpointer' not supported

var threadRefFunction *gi.Function
var threadRefFunction_Once sync.Once

func threadRefFunction_Set() error {
	var err error
	threadRefFunction_Once.Do(func() {
		err = threadStruct_Set()
		if err != nil {
			return
		}
		threadRefFunction, err = threadStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_thread_ref.
func (recv *Thread) Ref() (*Thread, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := threadRefFunction_Set()
	if err == nil {
		ret = threadRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Thread{native: ret.Pointer()}

	return retGo, err
}

var threadUnrefFunction *gi.Function
var threadUnrefFunction_Once sync.Once

func threadUnrefFunction_Set() error {
	var err error
	threadUnrefFunction_Once.Do(func() {
		err = threadStruct_Set()
		if err != nil {
			return
		}
		threadUnrefFunction, err = threadStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_thread_unref.
func (recv *Thread) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := threadUnrefFunction_Set()
	if err == nil {
		threadUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var threadPoolStruct *gi.Struct
var threadPoolStruct_Once sync.Once

func threadPoolStruct_Set() error {
	var err error
	threadPoolStruct_Once.Do(func() {
		threadPoolStruct, err = gi.StructNew("GLib", "ThreadPool")
	})
	return err
}

type ThreadPool struct {
	native uintptr
	// UNSUPPORTED : C value 'func' : no Go type for 'Func'
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'
	Exclusive bool
}

var threadPoolFreeFunction *gi.Function
var threadPoolFreeFunction_Once sync.Once

func threadPoolFreeFunction_Set() error {
	var err error
	threadPoolFreeFunction_Once.Do(func() {
		err = threadPoolStruct_Set()
		if err != nil {
			return
		}
		threadPoolFreeFunction, err = threadPoolStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_thread_pool_free.
func (recv *ThreadPool) Free(immediate bool, wait bool) error {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(immediate)
	inArgs[2].SetBoolean(wait)

	err := threadPoolFreeFunction_Set()
	if err == nil {
		threadPoolFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var threadPoolGetMaxThreadsFunction *gi.Function
var threadPoolGetMaxThreadsFunction_Once sync.Once

func threadPoolGetMaxThreadsFunction_Set() error {
	var err error
	threadPoolGetMaxThreadsFunction_Once.Do(func() {
		err = threadPoolStruct_Set()
		if err != nil {
			return
		}
		threadPoolGetMaxThreadsFunction, err = threadPoolStruct.InvokerNew("get_max_threads")
	})
	return err
}

// GetMaxThreads is a representation of the C type g_thread_pool_get_max_threads.
func (recv *ThreadPool) GetMaxThreads() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := threadPoolGetMaxThreadsFunction_Set()
	if err == nil {
		ret = threadPoolGetMaxThreadsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var threadPoolGetNumThreadsFunction *gi.Function
var threadPoolGetNumThreadsFunction_Once sync.Once

func threadPoolGetNumThreadsFunction_Set() error {
	var err error
	threadPoolGetNumThreadsFunction_Once.Do(func() {
		err = threadPoolStruct_Set()
		if err != nil {
			return
		}
		threadPoolGetNumThreadsFunction, err = threadPoolStruct.InvokerNew("get_num_threads")
	})
	return err
}

// GetNumThreads is a representation of the C type g_thread_pool_get_num_threads.
func (recv *ThreadPool) GetNumThreads() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := threadPoolGetNumThreadsFunction_Set()
	if err == nil {
		ret = threadPoolGetNumThreadsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_thread_pool_move_to_front' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_thread_pool_push' : parameter 'data' of type 'gpointer' not supported

var threadPoolSetMaxThreadsFunction *gi.Function
var threadPoolSetMaxThreadsFunction_Once sync.Once

func threadPoolSetMaxThreadsFunction_Set() error {
	var err error
	threadPoolSetMaxThreadsFunction_Once.Do(func() {
		err = threadPoolStruct_Set()
		if err != nil {
			return
		}
		threadPoolSetMaxThreadsFunction, err = threadPoolStruct.InvokerNew("set_max_threads")
	})
	return err
}

// SetMaxThreads is a representation of the C type g_thread_pool_set_max_threads.
func (recv *ThreadPool) SetMaxThreads(maxThreads int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(maxThreads)

	var ret gi.Argument

	err := threadPoolSetMaxThreadsFunction_Set()
	if err == nil {
		ret = threadPoolSetMaxThreadsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_thread_pool_set_sort_function' : parameter 'func' of type 'CompareDataFunc' not supported

var threadPoolUnprocessedFunction *gi.Function
var threadPoolUnprocessedFunction_Once sync.Once

func threadPoolUnprocessedFunction_Set() error {
	var err error
	threadPoolUnprocessedFunction_Once.Do(func() {
		err = threadPoolStruct_Set()
		if err != nil {
			return
		}
		threadPoolUnprocessedFunction, err = threadPoolStruct.InvokerNew("unprocessed")
	})
	return err
}

// Unprocessed is a representation of the C type g_thread_pool_unprocessed.
func (recv *ThreadPool) Unprocessed() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := threadPoolUnprocessedFunction_Set()
	if err == nil {
		ret = threadPoolUnprocessedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var timeValStruct *gi.Struct
var timeValStruct_Once sync.Once

func timeValStruct_Set() error {
	var err error
	timeValStruct_Once.Do(func() {
		timeValStruct, err = gi.StructNew("GLib", "TimeVal")
	})
	return err
}

type TimeVal struct {
	native uintptr
	TvSec  int64
	TvUsec int64
}

var timeValAddFunction *gi.Function
var timeValAddFunction_Once sync.Once

func timeValAddFunction_Set() error {
	var err error
	timeValAddFunction_Once.Do(func() {
		err = timeValStruct_Set()
		if err != nil {
			return
		}
		timeValAddFunction, err = timeValStruct.InvokerNew("add")
	})
	return err
}

// Add is a representation of the C type g_time_val_add.
func (recv *TimeVal) Add(microseconds int64) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt64(microseconds)

	err := timeValAddFunction_Set()
	if err == nil {
		timeValAddFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var timeValToIso8601Function *gi.Function
var timeValToIso8601Function_Once sync.Once

func timeValToIso8601Function_Set() error {
	var err error
	timeValToIso8601Function_Once.Do(func() {
		err = timeValStruct_Set()
		if err != nil {
			return
		}
		timeValToIso8601Function, err = timeValStruct.InvokerNew("to_iso8601")
	})
	return err
}

// ToIso8601 is a representation of the C type g_time_val_to_iso8601.
func (recv *TimeVal) ToIso8601() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := timeValToIso8601Function_Set()
	if err == nil {
		ret = timeValToIso8601Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var timeZoneStruct *gi.Struct
var timeZoneStruct_Once sync.Once

func timeZoneStruct_Set() error {
	var err error
	timeZoneStruct_Once.Do(func() {
		timeZoneStruct, err = gi.StructNew("GLib", "TimeZone")
	})
	return err
}

type TimeZone struct {
	native uintptr
}

var timeZoneNewFunction *gi.Function
var timeZoneNewFunction_Once sync.Once

func timeZoneNewFunction_Set() error {
	var err error
	timeZoneNewFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneNewFunction, err = timeZoneStruct.InvokerNew("new")
	})
	return err
}

// TimeZoneNew is a representation of the C type g_time_zone_new.
func TimeZoneNew(identifier string) (*TimeZone, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(identifier)

	var ret gi.Argument

	err := timeZoneNewFunction_Set()
	if err == nil {
		ret = timeZoneNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo, err
}

var timeZoneNewLocalFunction *gi.Function
var timeZoneNewLocalFunction_Once sync.Once

func timeZoneNewLocalFunction_Set() error {
	var err error
	timeZoneNewLocalFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneNewLocalFunction, err = timeZoneStruct.InvokerNew("new_local")
	})
	return err
}

// TimeZoneNewLocal is a representation of the C type g_time_zone_new_local.
func TimeZoneNewLocal() (*TimeZone, error) {

	var ret gi.Argument

	err := timeZoneNewLocalFunction_Set()
	if err == nil {
		ret = timeZoneNewLocalFunction.Invoke(nil, nil)
	}

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo, err
}

var timeZoneNewOffsetFunction *gi.Function
var timeZoneNewOffsetFunction_Once sync.Once

func timeZoneNewOffsetFunction_Set() error {
	var err error
	timeZoneNewOffsetFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneNewOffsetFunction, err = timeZoneStruct.InvokerNew("new_offset")
	})
	return err
}

// TimeZoneNewOffset is a representation of the C type g_time_zone_new_offset.
func TimeZoneNewOffset(seconds int32) (*TimeZone, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(seconds)

	var ret gi.Argument

	err := timeZoneNewOffsetFunction_Set()
	if err == nil {
		ret = timeZoneNewOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo, err
}

var timeZoneNewUtcFunction *gi.Function
var timeZoneNewUtcFunction_Once sync.Once

func timeZoneNewUtcFunction_Set() error {
	var err error
	timeZoneNewUtcFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneNewUtcFunction, err = timeZoneStruct.InvokerNew("new_utc")
	})
	return err
}

// TimeZoneNewUtc is a representation of the C type g_time_zone_new_utc.
func TimeZoneNewUtc() (*TimeZone, error) {

	var ret gi.Argument

	err := timeZoneNewUtcFunction_Set()
	if err == nil {
		ret = timeZoneNewUtcFunction.Invoke(nil, nil)
	}

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_time_zone_adjust_time' : parameter 'type' of type 'TimeType' not supported

// UNSUPPORTED : C value 'g_time_zone_find_interval' : parameter 'type' of type 'TimeType' not supported

var timeZoneGetAbbreviationFunction *gi.Function
var timeZoneGetAbbreviationFunction_Once sync.Once

func timeZoneGetAbbreviationFunction_Set() error {
	var err error
	timeZoneGetAbbreviationFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneGetAbbreviationFunction, err = timeZoneStruct.InvokerNew("get_abbreviation")
	})
	return err
}

// GetAbbreviation is a representation of the C type g_time_zone_get_abbreviation.
func (recv *TimeZone) GetAbbreviation(interval int32) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	var ret gi.Argument

	err := timeZoneGetAbbreviationFunction_Set()
	if err == nil {
		ret = timeZoneGetAbbreviationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var timeZoneGetIdentifierFunction *gi.Function
var timeZoneGetIdentifierFunction_Once sync.Once

func timeZoneGetIdentifierFunction_Set() error {
	var err error
	timeZoneGetIdentifierFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneGetIdentifierFunction, err = timeZoneStruct.InvokerNew("get_identifier")
	})
	return err
}

// GetIdentifier is a representation of the C type g_time_zone_get_identifier.
func (recv *TimeZone) GetIdentifier() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := timeZoneGetIdentifierFunction_Set()
	if err == nil {
		ret = timeZoneGetIdentifierFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var timeZoneGetOffsetFunction *gi.Function
var timeZoneGetOffsetFunction_Once sync.Once

func timeZoneGetOffsetFunction_Set() error {
	var err error
	timeZoneGetOffsetFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneGetOffsetFunction, err = timeZoneStruct.InvokerNew("get_offset")
	})
	return err
}

// GetOffset is a representation of the C type g_time_zone_get_offset.
func (recv *TimeZone) GetOffset(interval int32) (int32, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	var ret gi.Argument

	err := timeZoneGetOffsetFunction_Set()
	if err == nil {
		ret = timeZoneGetOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var timeZoneIsDstFunction *gi.Function
var timeZoneIsDstFunction_Once sync.Once

func timeZoneIsDstFunction_Set() error {
	var err error
	timeZoneIsDstFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneIsDstFunction, err = timeZoneStruct.InvokerNew("is_dst")
	})
	return err
}

// IsDst is a representation of the C type g_time_zone_is_dst.
func (recv *TimeZone) IsDst(interval int32) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(interval)

	var ret gi.Argument

	err := timeZoneIsDstFunction_Set()
	if err == nil {
		ret = timeZoneIsDstFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var timeZoneRefFunction *gi.Function
var timeZoneRefFunction_Once sync.Once

func timeZoneRefFunction_Set() error {
	var err error
	timeZoneRefFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneRefFunction, err = timeZoneStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_time_zone_ref.
func (recv *TimeZone) Ref() (*TimeZone, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := timeZoneRefFunction_Set()
	if err == nil {
		ret = timeZoneRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TimeZone{native: ret.Pointer()}

	return retGo, err
}

var timeZoneUnrefFunction *gi.Function
var timeZoneUnrefFunction_Once sync.Once

func timeZoneUnrefFunction_Set() error {
	var err error
	timeZoneUnrefFunction_Once.Do(func() {
		err = timeZoneStruct_Set()
		if err != nil {
			return
		}
		timeZoneUnrefFunction, err = timeZoneStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_time_zone_unref.
func (recv *TimeZone) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := timeZoneUnrefFunction_Set()
	if err == nil {
		timeZoneUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var timerStruct *gi.Struct
var timerStruct_Once sync.Once

func timerStruct_Set() error {
	var err error
	timerStruct_Once.Do(func() {
		timerStruct, err = gi.StructNew("GLib", "Timer")
	})
	return err
}

type Timer struct {
	native uintptr
}

var timerContinueFunction *gi.Function
var timerContinueFunction_Once sync.Once

func timerContinueFunction_Set() error {
	var err error
	timerContinueFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerContinueFunction, err = timerStruct.InvokerNew("continue")
	})
	return err
}

// Continue is a representation of the C type g_timer_continue.
func (recv *Timer) Continue() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := timerContinueFunction_Set()
	if err == nil {
		timerContinueFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var timerDestroyFunction *gi.Function
var timerDestroyFunction_Once sync.Once

func timerDestroyFunction_Set() error {
	var err error
	timerDestroyFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerDestroyFunction, err = timerStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type g_timer_destroy.
func (recv *Timer) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := timerDestroyFunction_Set()
	if err == nil {
		timerDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var timerElapsedFunction *gi.Function
var timerElapsedFunction_Once sync.Once

func timerElapsedFunction_Set() error {
	var err error
	timerElapsedFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerElapsedFunction, err = timerStruct.InvokerNew("elapsed")
	})
	return err
}

// Elapsed is a representation of the C type g_timer_elapsed.
func (recv *Timer) Elapsed(microseconds uint64) (float64, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint64(microseconds)

	var ret gi.Argument

	err := timerElapsedFunction_Set()
	if err == nil {
		ret = timerElapsedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

var timerIsActiveFunction *gi.Function
var timerIsActiveFunction_Once sync.Once

func timerIsActiveFunction_Set() error {
	var err error
	timerIsActiveFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerIsActiveFunction, err = timerStruct.InvokerNew("is_active")
	})
	return err
}

// IsActive is a representation of the C type g_timer_is_active.
func (recv *Timer) IsActive() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := timerIsActiveFunction_Set()
	if err == nil {
		ret = timerIsActiveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var timerResetFunction *gi.Function
var timerResetFunction_Once sync.Once

func timerResetFunction_Set() error {
	var err error
	timerResetFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerResetFunction, err = timerStruct.InvokerNew("reset")
	})
	return err
}

// Reset is a representation of the C type g_timer_reset.
func (recv *Timer) Reset() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := timerResetFunction_Set()
	if err == nil {
		timerResetFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var timerStartFunction *gi.Function
var timerStartFunction_Once sync.Once

func timerStartFunction_Set() error {
	var err error
	timerStartFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerStartFunction, err = timerStruct.InvokerNew("start")
	})
	return err
}

// Start is a representation of the C type g_timer_start.
func (recv *Timer) Start() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := timerStartFunction_Set()
	if err == nil {
		timerStartFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var timerStopFunction *gi.Function
var timerStopFunction_Once sync.Once

func timerStopFunction_Set() error {
	var err error
	timerStopFunction_Once.Do(func() {
		err = timerStruct_Set()
		if err != nil {
			return
		}
		timerStopFunction, err = timerStruct.InvokerNew("stop")
	})
	return err
}

// Stop is a representation of the C type g_timer_stop.
func (recv *Timer) Stop() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := timerStopFunction_Set()
	if err == nil {
		timerStopFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var trashStackStruct *gi.Struct
var trashStackStruct_Once sync.Once

func trashStackStruct_Set() error {
	var err error
	trashStackStruct_Once.Do(func() {
		trashStackStruct, err = gi.StructNew("GLib", "TrashStack")
	})
	return err
}

type TrashStack struct {
	native uintptr
	Next   *TrashStack
}

var treeStruct *gi.Struct
var treeStruct_Once sync.Once

func treeStruct_Set() error {
	var err error
	treeStruct_Once.Do(func() {
		treeStruct, err = gi.StructNew("GLib", "Tree")
	})
	return err
}

type Tree struct {
	native uintptr
}

var treeDestroyFunction *gi.Function
var treeDestroyFunction_Once sync.Once

func treeDestroyFunction_Set() error {
	var err error
	treeDestroyFunction_Once.Do(func() {
		err = treeStruct_Set()
		if err != nil {
			return
		}
		treeDestroyFunction, err = treeStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type g_tree_destroy.
func (recv *Tree) Destroy() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treeDestroyFunction_Set()
	if err == nil {
		treeDestroyFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_tree_foreach' : parameter 'func' of type 'TraverseFunc' not supported

var treeHeightFunction *gi.Function
var treeHeightFunction_Once sync.Once

func treeHeightFunction_Set() error {
	var err error
	treeHeightFunction_Once.Do(func() {
		err = treeStruct_Set()
		if err != nil {
			return
		}
		treeHeightFunction, err = treeStruct.InvokerNew("height")
	})
	return err
}

// Height is a representation of the C type g_tree_height.
func (recv *Tree) Height() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeHeightFunction_Set()
	if err == nil {
		ret = treeHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

// UNSUPPORTED : C value 'g_tree_insert' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_lookup' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_lookup_extended' : parameter 'lookup_key' of type 'gpointer' not supported

var treeNnodesFunction *gi.Function
var treeNnodesFunction_Once sync.Once

func treeNnodesFunction_Set() error {
	var err error
	treeNnodesFunction_Once.Do(func() {
		err = treeStruct_Set()
		if err != nil {
			return
		}
		treeNnodesFunction, err = treeStruct.InvokerNew("nnodes")
	})
	return err
}

// Nnodes is a representation of the C type g_tree_nnodes.
func (recv *Tree) Nnodes() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeNnodesFunction_Set()
	if err == nil {
		ret = treeNnodesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var treeRefFunction *gi.Function
var treeRefFunction_Once sync.Once

func treeRefFunction_Set() error {
	var err error
	treeRefFunction_Once.Do(func() {
		err = treeStruct_Set()
		if err != nil {
			return
		}
		treeRefFunction, err = treeStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_tree_ref.
func (recv *Tree) Ref() (*Tree, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeRefFunction_Set()
	if err == nil {
		ret = treeRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Tree{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_tree_remove' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_replace' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_search' : parameter 'search_func' of type 'CompareFunc' not supported

// UNSUPPORTED : C value 'g_tree_steal' : parameter 'key' of type 'gpointer' not supported

// UNSUPPORTED : C value 'g_tree_traverse' : parameter 'traverse_func' of type 'TraverseFunc' not supported

var treeUnrefFunction *gi.Function
var treeUnrefFunction_Once sync.Once

func treeUnrefFunction_Set() error {
	var err error
	treeUnrefFunction_Once.Do(func() {
		err = treeStruct_Set()
		if err != nil {
			return
		}
		treeUnrefFunction, err = treeStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_tree_unref.
func (recv *Tree) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treeUnrefFunction_Set()
	if err == nil {
		treeUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantStruct *gi.Struct
var variantStruct_Once sync.Once

func variantStruct_Set() error {
	var err error
	variantStruct_Once.Do(func() {
		variantStruct, err = gi.StructNew("GLib", "Variant")
	})
	return err
}

type Variant struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_new' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_array' : parameter 'child_type' of type 'VariantType' not supported

var variantNewBooleanFunction *gi.Function
var variantNewBooleanFunction_Once sync.Once

func variantNewBooleanFunction_Set() error {
	var err error
	variantNewBooleanFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewBooleanFunction, err = variantStruct.InvokerNew("new_boolean")
	})
	return err
}

// VariantNewBoolean is a representation of the C type g_variant_new_boolean.
func VariantNewBoolean(value bool) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(value)

	var ret gi.Argument

	err := variantNewBooleanFunction_Set()
	if err == nil {
		ret = variantNewBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewByteFunction *gi.Function
var variantNewByteFunction_Once sync.Once

func variantNewByteFunction_Set() error {
	var err error
	variantNewByteFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewByteFunction, err = variantStruct.InvokerNew("new_byte")
	})
	return err
}

// VariantNewByte is a representation of the C type g_variant_new_byte.
func VariantNewByte(value uint8) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint8(value)

	var ret gi.Argument

	err := variantNewByteFunction_Set()
	if err == nil {
		ret = variantNewByteFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_bytestring' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_variant_new_bytestring_array' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_dict_entry' : parameter 'key' of type 'Variant' not supported

var variantNewDoubleFunction *gi.Function
var variantNewDoubleFunction_Once sync.Once

func variantNewDoubleFunction_Set() error {
	var err error
	variantNewDoubleFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewDoubleFunction, err = variantStruct.InvokerNew("new_double")
	})
	return err
}

// VariantNewDouble is a representation of the C type g_variant_new_double.
func VariantNewDouble(value float64) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetFloat64(value)

	var ret gi.Argument

	err := variantNewDoubleFunction_Set()
	if err == nil {
		ret = variantNewDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_fixed_array' : parameter 'element_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_from_bytes' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_from_data' : parameter 'type' of type 'VariantType' not supported

var variantNewHandleFunction *gi.Function
var variantNewHandleFunction_Once sync.Once

func variantNewHandleFunction_Set() error {
	var err error
	variantNewHandleFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewHandleFunction, err = variantStruct.InvokerNew("new_handle")
	})
	return err
}

// VariantNewHandle is a representation of the C type g_variant_new_handle.
func VariantNewHandle(value int32) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	var ret gi.Argument

	err := variantNewHandleFunction_Set()
	if err == nil {
		ret = variantNewHandleFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewInt16Function *gi.Function
var variantNewInt16Function_Once sync.Once

func variantNewInt16Function_Set() error {
	var err error
	variantNewInt16Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewInt16Function, err = variantStruct.InvokerNew("new_int16")
	})
	return err
}

// VariantNewInt16 is a representation of the C type g_variant_new_int16.
func VariantNewInt16(value int16) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt16(value)

	var ret gi.Argument

	err := variantNewInt16Function_Set()
	if err == nil {
		ret = variantNewInt16Function.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewInt32Function *gi.Function
var variantNewInt32Function_Once sync.Once

func variantNewInt32Function_Set() error {
	var err error
	variantNewInt32Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewInt32Function, err = variantStruct.InvokerNew("new_int32")
	})
	return err
}

// VariantNewInt32 is a representation of the C type g_variant_new_int32.
func VariantNewInt32(value int32) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	var ret gi.Argument

	err := variantNewInt32Function_Set()
	if err == nil {
		ret = variantNewInt32Function.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewInt64Function *gi.Function
var variantNewInt64Function_Once sync.Once

func variantNewInt64Function_Set() error {
	var err error
	variantNewInt64Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewInt64Function, err = variantStruct.InvokerNew("new_int64")
	})
	return err
}

// VariantNewInt64 is a representation of the C type g_variant_new_int64.
func VariantNewInt64(value int64) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(value)

	var ret gi.Argument

	err := variantNewInt64Function_Set()
	if err == nil {
		ret = variantNewInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_maybe' : parameter 'child_type' of type 'VariantType' not supported

var variantNewObjectPathFunction *gi.Function
var variantNewObjectPathFunction_Once sync.Once

func variantNewObjectPathFunction_Set() error {
	var err error
	variantNewObjectPathFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewObjectPathFunction, err = variantStruct.InvokerNew("new_object_path")
	})
	return err
}

// VariantNewObjectPath is a representation of the C type g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(objectPath)

	var ret gi.Argument

	err := variantNewObjectPathFunction_Set()
	if err == nil {
		ret = variantNewObjectPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_objv' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_printf' : parameter '...' has no type

var variantNewSignatureFunction *gi.Function
var variantNewSignatureFunction_Once sync.Once

func variantNewSignatureFunction_Set() error {
	var err error
	variantNewSignatureFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewSignatureFunction, err = variantStruct.InvokerNew("new_signature")
	})
	return err
}

// VariantNewSignature is a representation of the C type g_variant_new_signature.
func VariantNewSignature(signature string) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(signature)

	var ret gi.Argument

	err := variantNewSignatureFunction_Set()
	if err == nil {
		ret = variantNewSignatureFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewStringFunction *gi.Function
var variantNewStringFunction_Once sync.Once

func variantNewStringFunction_Set() error {
	var err error
	variantNewStringFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewStringFunction, err = variantStruct.InvokerNew("new_string")
	})
	return err
}

// VariantNewString is a representation of the C type g_variant_new_string.
func VariantNewString(string_ string) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := variantNewStringFunction_Set()
	if err == nil {
		ret = variantNewStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_strv' : parameter 'strv' has no type

var variantNewTakeStringFunction *gi.Function
var variantNewTakeStringFunction_Once sync.Once

func variantNewTakeStringFunction_Set() error {
	var err error
	variantNewTakeStringFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewTakeStringFunction, err = variantStruct.InvokerNew("new_take_string")
	})
	return err
}

// VariantNewTakeString is a representation of the C type g_variant_new_take_string.
func VariantNewTakeString(string_ string) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	var ret gi.Argument

	err := variantNewTakeStringFunction_Set()
	if err == nil {
		ret = variantNewTakeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_tuple' : parameter 'children' has no type

var variantNewUint16Function *gi.Function
var variantNewUint16Function_Once sync.Once

func variantNewUint16Function_Set() error {
	var err error
	variantNewUint16Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewUint16Function, err = variantStruct.InvokerNew("new_uint16")
	})
	return err
}

// VariantNewUint16 is a representation of the C type g_variant_new_uint16.
func VariantNewUint16(value uint16) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(value)

	var ret gi.Argument

	err := variantNewUint16Function_Set()
	if err == nil {
		ret = variantNewUint16Function.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewUint32Function *gi.Function
var variantNewUint32Function_Once sync.Once

func variantNewUint32Function_Set() error {
	var err error
	variantNewUint32Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewUint32Function, err = variantStruct.InvokerNew("new_uint32")
	})
	return err
}

// VariantNewUint32 is a representation of the C type g_variant_new_uint32.
func VariantNewUint32(value uint32) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(value)

	var ret gi.Argument

	err := variantNewUint32Function_Set()
	if err == nil {
		ret = variantNewUint32Function.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantNewUint64Function *gi.Function
var variantNewUint64Function_Once sync.Once

func variantNewUint64Function_Set() error {
	var err error
	variantNewUint64Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantNewUint64Function, err = variantStruct.InvokerNew("new_uint64")
	})
	return err
}

// VariantNewUint64 is a representation of the C type g_variant_new_uint64.
func VariantNewUint64(value uint64) (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(value)

	var ret gi.Argument

	err := variantNewUint64Function_Set()
	if err == nil {
		ret = variantNewUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_new_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_variant' : parameter 'value' of type 'Variant' not supported

var variantByteswapFunction *gi.Function
var variantByteswapFunction_Once sync.Once

func variantByteswapFunction_Set() error {
	var err error
	variantByteswapFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantByteswapFunction, err = variantStruct.InvokerNew("byteswap")
	})
	return err
}

// Byteswap is a representation of the C type g_variant_byteswap.
func (recv *Variant) Byteswap() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantByteswapFunction_Set()
	if err == nil {
		ret = variantByteswapFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantCheckFormatStringFunction *gi.Function
var variantCheckFormatStringFunction_Once sync.Once

func variantCheckFormatStringFunction_Set() error {
	var err error
	variantCheckFormatStringFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantCheckFormatStringFunction, err = variantStruct.InvokerNew("check_format_string")
	})
	return err
}

// CheckFormatString is a representation of the C type g_variant_check_format_string.
func (recv *Variant) CheckFormatString(formatString string, copyOnly bool) (bool, error) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(formatString)
	inArgs[2].SetBoolean(copyOnly)

	var ret gi.Argument

	err := variantCheckFormatStringFunction_Set()
	if err == nil {
		ret = variantCheckFormatStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_classify' : return type 'VariantClass' not supported

// UNSUPPORTED : C value 'g_variant_compare' : parameter 'two' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dup_bytestring' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_bytestring_array' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_objv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_string' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_dup_strv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_equal' : parameter 'two' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_get' : parameter '...' has no type

var variantGetBooleanFunction *gi.Function
var variantGetBooleanFunction_Once sync.Once

func variantGetBooleanFunction_Set() error {
	var err error
	variantGetBooleanFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetBooleanFunction, err = variantStruct.InvokerNew("get_boolean")
	})
	return err
}

// GetBoolean is a representation of the C type g_variant_get_boolean.
func (recv *Variant) GetBoolean() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetBooleanFunction_Set()
	if err == nil {
		ret = variantGetBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantGetByteFunction *gi.Function
var variantGetByteFunction_Once sync.Once

func variantGetByteFunction_Set() error {
	var err error
	variantGetByteFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetByteFunction, err = variantStruct.InvokerNew("get_byte")
	})
	return err
}

// GetByte is a representation of the C type g_variant_get_byte.
func (recv *Variant) GetByte() (uint8, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetByteFunction_Set()
	if err == nil {
		ret = variantGetByteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo, err
}

var variantGetBytestringFunction *gi.Function
var variantGetBytestringFunction_Once sync.Once

func variantGetBytestringFunction_Set() error {
	var err error
	variantGetBytestringFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetBytestringFunction, err = variantStruct.InvokerNew("get_bytestring")
	})
	return err
}

// GetBytestring is a representation of the C type g_variant_get_bytestring.
func (recv *Variant) GetBytestring() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantGetBytestringFunction_Set()
	if err == nil {
		variantGetBytestringFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_variant_get_bytestring_array' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_child' : parameter 'index_' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_child_value' : parameter 'index_' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_data' : return type 'gpointer' not supported

var variantGetDataAsBytesFunction *gi.Function
var variantGetDataAsBytesFunction_Once sync.Once

func variantGetDataAsBytesFunction_Set() error {
	var err error
	variantGetDataAsBytesFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetDataAsBytesFunction, err = variantStruct.InvokerNew("get_data_as_bytes")
	})
	return err
}

// GetDataAsBytes is a representation of the C type g_variant_get_data_as_bytes.
func (recv *Variant) GetDataAsBytes() (*Bytes, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetDataAsBytesFunction_Set()
	if err == nil {
		ret = variantGetDataAsBytesFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Bytes{native: ret.Pointer()}

	return retGo, err
}

var variantGetDoubleFunction *gi.Function
var variantGetDoubleFunction_Once sync.Once

func variantGetDoubleFunction_Set() error {
	var err error
	variantGetDoubleFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetDoubleFunction, err = variantStruct.InvokerNew("get_double")
	})
	return err
}

// GetDouble is a representation of the C type g_variant_get_double.
func (recv *Variant) GetDouble() (float64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetDoubleFunction_Set()
	if err == nil {
		ret = variantGetDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_get_fixed_array' : parameter 'n_elements' of type 'gsize' not supported

var variantGetHandleFunction *gi.Function
var variantGetHandleFunction_Once sync.Once

func variantGetHandleFunction_Set() error {
	var err error
	variantGetHandleFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetHandleFunction, err = variantStruct.InvokerNew("get_handle")
	})
	return err
}

// GetHandle is a representation of the C type g_variant_get_handle.
func (recv *Variant) GetHandle() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetHandleFunction_Set()
	if err == nil {
		ret = variantGetHandleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var variantGetInt16Function *gi.Function
var variantGetInt16Function_Once sync.Once

func variantGetInt16Function_Set() error {
	var err error
	variantGetInt16Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetInt16Function, err = variantStruct.InvokerNew("get_int16")
	})
	return err
}

// GetInt16 is a representation of the C type g_variant_get_int16.
func (recv *Variant) GetInt16() (int16, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetInt16Function_Set()
	if err == nil {
		ret = variantGetInt16Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int16()

	return retGo, err
}

var variantGetInt32Function *gi.Function
var variantGetInt32Function_Once sync.Once

func variantGetInt32Function_Set() error {
	var err error
	variantGetInt32Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetInt32Function, err = variantStruct.InvokerNew("get_int32")
	})
	return err
}

// GetInt32 is a representation of the C type g_variant_get_int32.
func (recv *Variant) GetInt32() (int32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetInt32Function_Set()
	if err == nil {
		ret = variantGetInt32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo, err
}

var variantGetInt64Function *gi.Function
var variantGetInt64Function_Once sync.Once

func variantGetInt64Function_Set() error {
	var err error
	variantGetInt64Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetInt64Function, err = variantStruct.InvokerNew("get_int64")
	})
	return err
}

// GetInt64 is a representation of the C type g_variant_get_int64.
func (recv *Variant) GetInt64() (int64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetInt64Function_Set()
	if err == nil {
		ret = variantGetInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo, err
}

var variantGetMaybeFunction *gi.Function
var variantGetMaybeFunction_Once sync.Once

func variantGetMaybeFunction_Set() error {
	var err error
	variantGetMaybeFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetMaybeFunction, err = variantStruct.InvokerNew("get_maybe")
	})
	return err
}

// GetMaybe is a representation of the C type g_variant_get_maybe.
func (recv *Variant) GetMaybe() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetMaybeFunction_Set()
	if err == nil {
		ret = variantGetMaybeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantGetNormalFormFunction *gi.Function
var variantGetNormalFormFunction_Once sync.Once

func variantGetNormalFormFunction_Set() error {
	var err error
	variantGetNormalFormFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetNormalFormFunction, err = variantStruct.InvokerNew("get_normal_form")
	})
	return err
}

// GetNormalForm is a representation of the C type g_variant_get_normal_form.
func (recv *Variant) GetNormalForm() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetNormalFormFunction_Set()
	if err == nil {
		ret = variantGetNormalFormFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_get_objv' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_size' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_string' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_get_strv' : parameter 'length' of type 'gsize' not supported

var variantGetTypeFunction *gi.Function
var variantGetTypeFunction_Once sync.Once

func variantGetTypeFunction_Set() error {
	var err error
	variantGetTypeFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetTypeFunction, err = variantStruct.InvokerNew("get_type")
	})
	return err
}

// GetType is a representation of the C type g_variant_get_type.
func (recv *Variant) GetType() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetTypeFunction_Set()
	if err == nil {
		ret = variantGetTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

var variantGetTypeStringFunction *gi.Function
var variantGetTypeStringFunction_Once sync.Once

func variantGetTypeStringFunction_Set() error {
	var err error
	variantGetTypeStringFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetTypeStringFunction, err = variantStruct.InvokerNew("get_type_string")
	})
	return err
}

// GetTypeString is a representation of the C type g_variant_get_type_string.
func (recv *Variant) GetTypeString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetTypeStringFunction_Set()
	if err == nil {
		ret = variantGetTypeStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var variantGetUint16Function *gi.Function
var variantGetUint16Function_Once sync.Once

func variantGetUint16Function_Set() error {
	var err error
	variantGetUint16Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetUint16Function, err = variantStruct.InvokerNew("get_uint16")
	})
	return err
}

// GetUint16 is a representation of the C type g_variant_get_uint16.
func (recv *Variant) GetUint16() (uint16, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetUint16Function_Set()
	if err == nil {
		ret = variantGetUint16Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo, err
}

var variantGetUint32Function *gi.Function
var variantGetUint32Function_Once sync.Once

func variantGetUint32Function_Set() error {
	var err error
	variantGetUint32Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetUint32Function, err = variantStruct.InvokerNew("get_uint32")
	})
	return err
}

// GetUint32 is a representation of the C type g_variant_get_uint32.
func (recv *Variant) GetUint32() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetUint32Function_Set()
	if err == nil {
		ret = variantGetUint32Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var variantGetUint64Function *gi.Function
var variantGetUint64Function_Once sync.Once

func variantGetUint64Function_Set() error {
	var err error
	variantGetUint64Function_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetUint64Function, err = variantStruct.InvokerNew("get_uint64")
	})
	return err
}

// GetUint64 is a representation of the C type g_variant_get_uint64.
func (recv *Variant) GetUint64() (uint64, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetUint64Function_Set()
	if err == nil {
		ret = variantGetUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_get_va' : parameter 'app' of type 'va_list' not supported

var variantGetVariantFunction *gi.Function
var variantGetVariantFunction_Once sync.Once

func variantGetVariantFunction_Set() error {
	var err error
	variantGetVariantFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantGetVariantFunction, err = variantStruct.InvokerNew("get_variant")
	})
	return err
}

// GetVariant is a representation of the C type g_variant_get_variant.
func (recv *Variant) GetVariant() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantGetVariantFunction_Set()
	if err == nil {
		ret = variantGetVariantFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantHashFunction *gi.Function
var variantHashFunction_Once sync.Once

func variantHashFunction_Set() error {
	var err error
	variantHashFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantHashFunction, err = variantStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type g_variant_hash.
func (recv *Variant) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantHashFunction_Set()
	if err == nil {
		ret = variantHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var variantIsContainerFunction *gi.Function
var variantIsContainerFunction_Once sync.Once

func variantIsContainerFunction_Set() error {
	var err error
	variantIsContainerFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantIsContainerFunction, err = variantStruct.InvokerNew("is_container")
	})
	return err
}

// IsContainer is a representation of the C type g_variant_is_container.
func (recv *Variant) IsContainer() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantIsContainerFunction_Set()
	if err == nil {
		ret = variantIsContainerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantIsFloatingFunction *gi.Function
var variantIsFloatingFunction_Once sync.Once

func variantIsFloatingFunction_Set() error {
	var err error
	variantIsFloatingFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantIsFloatingFunction, err = variantStruct.InvokerNew("is_floating")
	})
	return err
}

// IsFloating is a representation of the C type g_variant_is_floating.
func (recv *Variant) IsFloating() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantIsFloatingFunction_Set()
	if err == nil {
		ret = variantIsFloatingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantIsNormalFormFunction *gi.Function
var variantIsNormalFormFunction_Once sync.Once

func variantIsNormalFormFunction_Set() error {
	var err error
	variantIsNormalFormFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantIsNormalFormFunction, err = variantStruct.InvokerNew("is_normal_form")
	})
	return err
}

// IsNormalForm is a representation of the C type g_variant_is_normal_form.
func (recv *Variant) IsNormalForm() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantIsNormalFormFunction_Set()
	if err == nil {
		ret = variantIsNormalFormFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_is_of_type' : parameter 'type' of type 'VariantType' not supported

var variantIterNewFunction *gi.Function
var variantIterNewFunction_Once sync.Once

func variantIterNewFunction_Set() error {
	var err error
	variantIterNewFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantIterNewFunction, err = variantStruct.InvokerNew("iter_new")
	})
	return err
}

// IterNew is a representation of the C type g_variant_iter_new.
func (recv *Variant) IterNew() (*VariantIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantIterNewFunction_Set()
	if err == nil {
		ret = variantIterNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantIter{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_lookup' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_lookup_value' : parameter 'expected_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_n_children' : return type 'gsize' not supported

var variantPrintFunction *gi.Function
var variantPrintFunction_Once sync.Once

func variantPrintFunction_Set() error {
	var err error
	variantPrintFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantPrintFunction, err = variantStruct.InvokerNew("print")
	})
	return err
}

// Print is a representation of the C type g_variant_print.
func (recv *Variant) Print(typeAnnotate bool) (string, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(typeAnnotate)

	var ret gi.Argument

	err := variantPrintFunction_Set()
	if err == nil {
		ret = variantPrintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_print_string' : parameter 'string' of type 'String' not supported

var variantRefFunction *gi.Function
var variantRefFunction_Once sync.Once

func variantRefFunction_Set() error {
	var err error
	variantRefFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantRefFunction, err = variantStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_variant_ref.
func (recv *Variant) Ref() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantRefFunction_Set()
	if err == nil {
		ret = variantRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantRefSinkFunction *gi.Function
var variantRefSinkFunction_Once sync.Once

func variantRefSinkFunction_Set() error {
	var err error
	variantRefSinkFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantRefSinkFunction, err = variantStruct.InvokerNew("ref_sink")
	})
	return err
}

// RefSink is a representation of the C type g_variant_ref_sink.
func (recv *Variant) RefSink() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantRefSinkFunction_Set()
	if err == nil {
		ret = variantRefSinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_store' : parameter 'data' of type 'gpointer' not supported

var variantTakeRefFunction *gi.Function
var variantTakeRefFunction_Once sync.Once

func variantTakeRefFunction_Set() error {
	var err error
	variantTakeRefFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantTakeRefFunction, err = variantStruct.InvokerNew("take_ref")
	})
	return err
}

// TakeRef is a representation of the C type g_variant_take_ref.
func (recv *Variant) TakeRef() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTakeRefFunction_Set()
	if err == nil {
		ret = variantTakeRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantUnrefFunction *gi.Function
var variantUnrefFunction_Once sync.Once

func variantUnrefFunction_Set() error {
	var err error
	variantUnrefFunction_Once.Do(func() {
		err = variantStruct_Set()
		if err != nil {
			return
		}
		variantUnrefFunction, err = variantStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_variant_unref.
func (recv *Variant) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantUnrefFunction_Set()
	if err == nil {
		variantUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantBuilderStruct *gi.Struct
var variantBuilderStruct_Once sync.Once

func variantBuilderStruct_Set() error {
	var err error
	variantBuilderStruct_Once.Do(func() {
		variantBuilderStruct, err = gi.StructNew("GLib", "VariantBuilder")
	})
	return err
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

func variantBuilderClearFunction_Set() error {
	var err error
	variantBuilderClearFunction_Once.Do(func() {
		err = variantBuilderStruct_Set()
		if err != nil {
			return
		}
		variantBuilderClearFunction, err = variantBuilderStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_variant_builder_clear.
func (recv *VariantBuilder) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantBuilderClearFunction_Set()
	if err == nil {
		variantBuilderClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantBuilderCloseFunction *gi.Function
var variantBuilderCloseFunction_Once sync.Once

func variantBuilderCloseFunction_Set() error {
	var err error
	variantBuilderCloseFunction_Once.Do(func() {
		err = variantBuilderStruct_Set()
		if err != nil {
			return
		}
		variantBuilderCloseFunction, err = variantBuilderStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type g_variant_builder_close.
func (recv *VariantBuilder) Close() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantBuilderCloseFunction_Set()
	if err == nil {
		variantBuilderCloseFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantBuilderEndFunction *gi.Function
var variantBuilderEndFunction_Once sync.Once

func variantBuilderEndFunction_Set() error {
	var err error
	variantBuilderEndFunction_Once.Do(func() {
		err = variantBuilderStruct_Set()
		if err != nil {
			return
		}
		variantBuilderEndFunction, err = variantBuilderStruct.InvokerNew("end")
	})
	return err
}

// End is a representation of the C type g_variant_builder_end.
func (recv *VariantBuilder) End() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantBuilderEndFunction_Set()
	if err == nil {
		ret = variantBuilderEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_builder_init' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_builder_open' : parameter 'type' of type 'VariantType' not supported

var variantBuilderRefFunction *gi.Function
var variantBuilderRefFunction_Once sync.Once

func variantBuilderRefFunction_Set() error {
	var err error
	variantBuilderRefFunction_Once.Do(func() {
		err = variantBuilderStruct_Set()
		if err != nil {
			return
		}
		variantBuilderRefFunction, err = variantBuilderStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_variant_builder_ref.
func (recv *VariantBuilder) Ref() (*VariantBuilder, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantBuilderRefFunction_Set()
	if err == nil {
		ret = variantBuilderRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantBuilder{native: ret.Pointer()}

	return retGo, err
}

var variantBuilderUnrefFunction *gi.Function
var variantBuilderUnrefFunction_Once sync.Once

func variantBuilderUnrefFunction_Set() error {
	var err error
	variantBuilderUnrefFunction_Once.Do(func() {
		err = variantBuilderStruct_Set()
		if err != nil {
			return
		}
		variantBuilderUnrefFunction, err = variantBuilderStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_variant_builder_unref.
func (recv *VariantBuilder) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantBuilderUnrefFunction_Set()
	if err == nil {
		variantBuilderUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantDictStruct *gi.Struct
var variantDictStruct_Once sync.Once

func variantDictStruct_Set() error {
	var err error
	variantDictStruct_Once.Do(func() {
		variantDictStruct, err = gi.StructNew("GLib", "VariantDict")
	})
	return err
}

type VariantDict struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_dict_new' : parameter 'from_asv' of type 'Variant' not supported

var variantDictClearFunction *gi.Function
var variantDictClearFunction_Once sync.Once

func variantDictClearFunction_Set() error {
	var err error
	variantDictClearFunction_Once.Do(func() {
		err = variantDictStruct_Set()
		if err != nil {
			return
		}
		variantDictClearFunction, err = variantDictStruct.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type g_variant_dict_clear.
func (recv *VariantDict) Clear() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantDictClearFunction_Set()
	if err == nil {
		variantDictClearFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantDictContainsFunction *gi.Function
var variantDictContainsFunction_Once sync.Once

func variantDictContainsFunction_Set() error {
	var err error
	variantDictContainsFunction_Once.Do(func() {
		err = variantDictStruct_Set()
		if err != nil {
			return
		}
		variantDictContainsFunction, err = variantDictStruct.InvokerNew("contains")
	})
	return err
}

// Contains is a representation of the C type g_variant_dict_contains.
func (recv *VariantDict) Contains(key string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := variantDictContainsFunction_Set()
	if err == nil {
		ret = variantDictContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantDictEndFunction *gi.Function
var variantDictEndFunction_Once sync.Once

func variantDictEndFunction_Set() error {
	var err error
	variantDictEndFunction_Once.Do(func() {
		err = variantDictStruct_Set()
		if err != nil {
			return
		}
		variantDictEndFunction, err = variantDictStruct.InvokerNew("end")
	})
	return err
}

// End is a representation of the C type g_variant_dict_end.
func (recv *VariantDict) End() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantDictEndFunction_Set()
	if err == nil {
		ret = variantDictEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_dict_init' : parameter 'from_asv' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dict_insert' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_dict_insert_value' : parameter 'value' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_dict_lookup' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_dict_lookup_value' : parameter 'expected_type' of type 'VariantType' not supported

var variantDictRefFunction *gi.Function
var variantDictRefFunction_Once sync.Once

func variantDictRefFunction_Set() error {
	var err error
	variantDictRefFunction_Once.Do(func() {
		err = variantDictStruct_Set()
		if err != nil {
			return
		}
		variantDictRefFunction, err = variantDictStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type g_variant_dict_ref.
func (recv *VariantDict) Ref() (*VariantDict, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantDictRefFunction_Set()
	if err == nil {
		ret = variantDictRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantDict{native: ret.Pointer()}

	return retGo, err
}

var variantDictRemoveFunction *gi.Function
var variantDictRemoveFunction_Once sync.Once

func variantDictRemoveFunction_Set() error {
	var err error
	variantDictRemoveFunction_Once.Do(func() {
		err = variantDictStruct_Set()
		if err != nil {
			return
		}
		variantDictRemoveFunction, err = variantDictStruct.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type g_variant_dict_remove.
func (recv *VariantDict) Remove(key string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := variantDictRemoveFunction_Set()
	if err == nil {
		ret = variantDictRemoveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantDictUnrefFunction *gi.Function
var variantDictUnrefFunction_Once sync.Once

func variantDictUnrefFunction_Set() error {
	var err error
	variantDictUnrefFunction_Once.Do(func() {
		err = variantDictStruct_Set()
		if err != nil {
			return
		}
		variantDictUnrefFunction, err = variantDictStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type g_variant_dict_unref.
func (recv *VariantDict) Unref() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantDictUnrefFunction_Set()
	if err == nil {
		variantDictUnrefFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var variantIterStruct *gi.Struct
var variantIterStruct_Once sync.Once

func variantIterStruct_Set() error {
	var err error
	variantIterStruct_Once.Do(func() {
		variantIterStruct, err = gi.StructNew("GLib", "VariantIter")
	})
	return err
}

type VariantIter struct {
	native uintptr
}

var variantIterCopyFunction *gi.Function
var variantIterCopyFunction_Once sync.Once

func variantIterCopyFunction_Set() error {
	var err error
	variantIterCopyFunction_Once.Do(func() {
		err = variantIterStruct_Set()
		if err != nil {
			return
		}
		variantIterCopyFunction, err = variantIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_variant_iter_copy.
func (recv *VariantIter) Copy() (*VariantIter, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantIterCopyFunction_Set()
	if err == nil {
		ret = variantIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantIter{native: ret.Pointer()}

	return retGo, err
}

var variantIterFreeFunction *gi.Function
var variantIterFreeFunction_Once sync.Once

func variantIterFreeFunction_Set() error {
	var err error
	variantIterFreeFunction_Once.Do(func() {
		err = variantIterStruct_Set()
		if err != nil {
			return
		}
		variantIterFreeFunction, err = variantIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_variant_iter_free.
func (recv *VariantIter) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantIterFreeFunction_Set()
	if err == nil {
		variantIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_variant_iter_init' : parameter 'value' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_iter_loop' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_iter_n_children' : return type 'gsize' not supported

// UNSUPPORTED : C value 'g_variant_iter_next' : parameter '...' has no type

var variantIterNextValueFunction *gi.Function
var variantIterNextValueFunction_Once sync.Once

func variantIterNextValueFunction_Set() error {
	var err error
	variantIterNextValueFunction_Once.Do(func() {
		err = variantIterStruct_Set()
		if err != nil {
			return
		}
		variantIterNextValueFunction, err = variantIterStruct.InvokerNew("next_value")
	})
	return err
}

// NextValue is a representation of the C type g_variant_iter_next_value.
func (recv *VariantIter) NextValue() (*Variant, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantIterNextValueFunction_Set()
	if err == nil {
		ret = variantIterNextValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Variant{native: ret.Pointer()}

	return retGo, err
}

var variantTypeStruct *gi.Struct
var variantTypeStruct_Once sync.Once

func variantTypeStruct_Set() error {
	var err error
	variantTypeStruct_Once.Do(func() {
		variantTypeStruct, err = gi.StructNew("GLib", "VariantType")
	})
	return err
}

type VariantType struct {
	native uintptr
}

var variantTypeNewFunction *gi.Function
var variantTypeNewFunction_Once sync.Once

func variantTypeNewFunction_Set() error {
	var err error
	variantTypeNewFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeNewFunction, err = variantTypeStruct.InvokerNew("new")
	})
	return err
}

// VariantTypeNew is a representation of the C type g_variant_type_new.
func VariantTypeNew(typeString string) (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	var ret gi.Argument

	err := variantTypeNewFunction_Set()
	if err == nil {
		ret = variantTypeNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_type_new_array' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_dict_entry' : parameter 'key' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_maybe' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_tuple' : parameter 'items' has no type

var variantTypeCopyFunction *gi.Function
var variantTypeCopyFunction_Once sync.Once

func variantTypeCopyFunction_Set() error {
	var err error
	variantTypeCopyFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeCopyFunction, err = variantTypeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type g_variant_type_copy.
func (recv *VariantType) Copy() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeCopyFunction_Set()
	if err == nil {
		ret = variantTypeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

var variantTypeDupStringFunction *gi.Function
var variantTypeDupStringFunction_Once sync.Once

func variantTypeDupStringFunction_Set() error {
	var err error
	variantTypeDupStringFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeDupStringFunction, err = variantTypeStruct.InvokerNew("dup_string")
	})
	return err
}

// DupString is a representation of the C type g_variant_type_dup_string.
func (recv *VariantType) DupString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeDupStringFunction_Set()
	if err == nil {
		ret = variantTypeDupStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var variantTypeElementFunction *gi.Function
var variantTypeElementFunction_Once sync.Once

func variantTypeElementFunction_Set() error {
	var err error
	variantTypeElementFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeElementFunction, err = variantTypeStruct.InvokerNew("element")
	})
	return err
}

// Element is a representation of the C type g_variant_type_element.
func (recv *VariantType) Element() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeElementFunction_Set()
	if err == nil {
		ret = variantTypeElementFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_type_equal' : parameter 'type2' of type 'VariantType' not supported

var variantTypeFirstFunction *gi.Function
var variantTypeFirstFunction_Once sync.Once

func variantTypeFirstFunction_Set() error {
	var err error
	variantTypeFirstFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeFirstFunction, err = variantTypeStruct.InvokerNew("first")
	})
	return err
}

// First is a representation of the C type g_variant_type_first.
func (recv *VariantType) First() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeFirstFunction_Set()
	if err == nil {
		ret = variantTypeFirstFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

var variantTypeFreeFunction *gi.Function
var variantTypeFreeFunction_Once sync.Once

func variantTypeFreeFunction_Set() error {
	var err error
	variantTypeFreeFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeFreeFunction, err = variantTypeStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type g_variant_type_free.
func (recv *VariantType) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := variantTypeFreeFunction_Set()
	if err == nil {
		variantTypeFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

// UNSUPPORTED : C value 'g_variant_type_get_string_length' : return type 'gsize' not supported

var variantTypeHashFunction *gi.Function
var variantTypeHashFunction_Once sync.Once

func variantTypeHashFunction_Set() error {
	var err error
	variantTypeHashFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeHashFunction, err = variantTypeStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type g_variant_type_hash.
func (recv *VariantType) Hash() (uint32, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeHashFunction_Set()
	if err == nil {
		ret = variantTypeHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo, err
}

var variantTypeIsArrayFunction *gi.Function
var variantTypeIsArrayFunction_Once sync.Once

func variantTypeIsArrayFunction_Set() error {
	var err error
	variantTypeIsArrayFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsArrayFunction, err = variantTypeStruct.InvokerNew("is_array")
	})
	return err
}

// IsArray is a representation of the C type g_variant_type_is_array.
func (recv *VariantType) IsArray() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsArrayFunction_Set()
	if err == nil {
		ret = variantTypeIsArrayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeIsBasicFunction *gi.Function
var variantTypeIsBasicFunction_Once sync.Once

func variantTypeIsBasicFunction_Set() error {
	var err error
	variantTypeIsBasicFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsBasicFunction, err = variantTypeStruct.InvokerNew("is_basic")
	})
	return err
}

// IsBasic is a representation of the C type g_variant_type_is_basic.
func (recv *VariantType) IsBasic() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsBasicFunction_Set()
	if err == nil {
		ret = variantTypeIsBasicFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeIsContainerFunction *gi.Function
var variantTypeIsContainerFunction_Once sync.Once

func variantTypeIsContainerFunction_Set() error {
	var err error
	variantTypeIsContainerFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsContainerFunction, err = variantTypeStruct.InvokerNew("is_container")
	})
	return err
}

// IsContainer is a representation of the C type g_variant_type_is_container.
func (recv *VariantType) IsContainer() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsContainerFunction_Set()
	if err == nil {
		ret = variantTypeIsContainerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeIsDefiniteFunction *gi.Function
var variantTypeIsDefiniteFunction_Once sync.Once

func variantTypeIsDefiniteFunction_Set() error {
	var err error
	variantTypeIsDefiniteFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsDefiniteFunction, err = variantTypeStruct.InvokerNew("is_definite")
	})
	return err
}

// IsDefinite is a representation of the C type g_variant_type_is_definite.
func (recv *VariantType) IsDefinite() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsDefiniteFunction_Set()
	if err == nil {
		ret = variantTypeIsDefiniteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeIsDictEntryFunction *gi.Function
var variantTypeIsDictEntryFunction_Once sync.Once

func variantTypeIsDictEntryFunction_Set() error {
	var err error
	variantTypeIsDictEntryFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsDictEntryFunction, err = variantTypeStruct.InvokerNew("is_dict_entry")
	})
	return err
}

// IsDictEntry is a representation of the C type g_variant_type_is_dict_entry.
func (recv *VariantType) IsDictEntry() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsDictEntryFunction_Set()
	if err == nil {
		ret = variantTypeIsDictEntryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeIsMaybeFunction *gi.Function
var variantTypeIsMaybeFunction_Once sync.Once

func variantTypeIsMaybeFunction_Set() error {
	var err error
	variantTypeIsMaybeFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsMaybeFunction, err = variantTypeStruct.InvokerNew("is_maybe")
	})
	return err
}

// IsMaybe is a representation of the C type g_variant_type_is_maybe.
func (recv *VariantType) IsMaybe() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsMaybeFunction_Set()
	if err == nil {
		ret = variantTypeIsMaybeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_type_is_subtype_of' : parameter 'supertype' of type 'VariantType' not supported

var variantTypeIsTupleFunction *gi.Function
var variantTypeIsTupleFunction_Once sync.Once

func variantTypeIsTupleFunction_Set() error {
	var err error
	variantTypeIsTupleFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsTupleFunction, err = variantTypeStruct.InvokerNew("is_tuple")
	})
	return err
}

// IsTuple is a representation of the C type g_variant_type_is_tuple.
func (recv *VariantType) IsTuple() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsTupleFunction_Set()
	if err == nil {
		ret = variantTypeIsTupleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeIsVariantFunction *gi.Function
var variantTypeIsVariantFunction_Once sync.Once

func variantTypeIsVariantFunction_Set() error {
	var err error
	variantTypeIsVariantFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeIsVariantFunction, err = variantTypeStruct.InvokerNew("is_variant")
	})
	return err
}

// IsVariant is a representation of the C type g_variant_type_is_variant.
func (recv *VariantType) IsVariant() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeIsVariantFunction_Set()
	if err == nil {
		ret = variantTypeIsVariantFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var variantTypeKeyFunction *gi.Function
var variantTypeKeyFunction_Once sync.Once

func variantTypeKeyFunction_Set() error {
	var err error
	variantTypeKeyFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeKeyFunction, err = variantTypeStruct.InvokerNew("key")
	})
	return err
}

// Key is a representation of the C type g_variant_type_key.
func (recv *VariantType) Key() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeKeyFunction_Set()
	if err == nil {
		ret = variantTypeKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

// UNSUPPORTED : C value 'g_variant_type_n_items' : return type 'gsize' not supported

var variantTypeNextFunction *gi.Function
var variantTypeNextFunction_Once sync.Once

func variantTypeNextFunction_Set() error {
	var err error
	variantTypeNextFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeNextFunction, err = variantTypeStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type g_variant_type_next.
func (recv *VariantType) Next() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeNextFunction_Set()
	if err == nil {
		ret = variantTypeNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}

var variantTypePeekStringFunction *gi.Function
var variantTypePeekStringFunction_Once sync.Once

func variantTypePeekStringFunction_Set() error {
	var err error
	variantTypePeekStringFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypePeekStringFunction, err = variantTypeStruct.InvokerNew("peek_string")
	})
	return err
}

// PeekString is a representation of the C type g_variant_type_peek_string.
func (recv *VariantType) PeekString() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypePeekStringFunction_Set()
	if err == nil {
		ret = variantTypePeekStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo, err
}

var variantTypeValueFunction *gi.Function
var variantTypeValueFunction_Once sync.Once

func variantTypeValueFunction_Set() error {
	var err error
	variantTypeValueFunction_Once.Do(func() {
		err = variantTypeStruct_Set()
		if err != nil {
			return
		}
		variantTypeValueFunction, err = variantTypeStruct.InvokerNew("value")
	})
	return err
}

// Value is a representation of the C type g_variant_type_value.
func (recv *VariantType) Value() (*VariantType, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := variantTypeValueFunction_Set()
	if err == nil {
		ret = variantTypeValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := &VariantType{native: ret.Pointer()}

	return retGo, err
}
