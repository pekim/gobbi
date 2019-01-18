// This is a generated file - DO NOT EDIT
// +build glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_async_queue_push_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_push_sorted_unlocked : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_async_queue_sort : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// Unsupported : g_async_queue_sort_unlocked : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func

// SetTimeT is a wrapper around the C function g_date_set_time_t.
func (recv *Date) SetTimeT(timet int64) {
	c_timet := (C.time_t)(timet)

	C.g_date_set_time_t((*C.GDate)(recv.native), c_timet)

	return
}

// SetTimeVal is a wrapper around the C function g_date_set_time_val.
func (recv *Date) SetTimeVal(timeval *TimeVal) {
	c_timeval := (*C.GTimeVal)(C.NULL)
	if timeval != nil {
		c_timeval = (*C.GTimeVal)(timeval.ToC())
	}

	C.g_date_set_time_val((*C.GDate)(recv.native), c_timeval)

	return
}

// HashTableRef is a wrapper around the C function g_hash_table_ref.
func HashTableRef(hashTable *HashTable) *HashTable {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	retC := C.g_hash_table_ref(c_hash_table)
	retGo := HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HashTableUnref is a wrapper around the C function g_hash_table_unref.
func HashTableUnref(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_unref(c_hash_table)

	return
}

// g_list_insert_sorted_with_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// IsOwner is a wrapper around the C function g_main_context_is_owner.
func (recv *MainContext) IsOwner() bool {
	retC := C.g_main_context_is_owner((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// g_slist_insert_sorted_with_data : unsupported parameter data : no type generator for gpointer (gpointer) for param data
// ThreadPoolGetMaxIdleTime is a wrapper around the C function g_thread_pool_get_max_idle_time.
func ThreadPoolGetMaxIdleTime() uint32 {
	retC := C.g_thread_pool_get_max_idle_time()
	retGo := (uint32)(retC)

	return retGo
}

// ThreadPoolSetMaxIdleTime is a wrapper around the C function g_thread_pool_set_max_idle_time.
func ThreadPoolSetMaxIdleTime(interval uint32) {
	c_interval := (C.guint)(interval)

	C.g_thread_pool_set_max_idle_time(c_interval)

	return
}

// Unsupported : g_thread_pool_set_sort_function : unsupported parameter func : no type generator for CompareDataFunc (GCompareDataFunc) for param func
