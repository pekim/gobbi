// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #include <glib.h>
// #include <stdlib.h>
import "C"

// Blacklisted function: g_atomic_int_add

// Blacklisted function: g_atomic_int_compare_and_exchange

// Blacklisted function: g_atomic_int_dec_and_test

// Blacklisted function: g_atomic_int_exchange_and_add

// Blacklisted function: g_atomic_int_get

// Blacklisted function: g_atomic_int_inc

// Blacklisted function: g_atomic_int_set

// Blacklisted function: g_atomic_pointer_compare_and_exchange

// Blacklisted function: g_atomic_pointer_get

// Blacklisted function: g_atomic_pointer_set

// Unsupported function: g_child_watch_add : unsupported parameter pid : type Pid, GPid

// Unsupported function: g_child_watch_add_full : unsupported parameter pid : type Pid, GPid

// Unsupported function: g_child_watch_source_new : unsupported parameter pid : type Pid, GPid

// Unsupported function: g_file_read_link : unsupported parameter filename : type filename, const gchar*

// Unsupported function: g_markup_printf_escaped : unsupported parameter ... : varargs

// Unsupported function: g_markup_vprintf_escaped : unsupported parameter args : type va_list, va_list

// Unsupported function: g_setenv : unsupported parameter variable : type filename, const gchar*

// StripContext is a wrapper around the C function g_strip_context.
func StripContext(msgid string, msgval string) {}

// StrsplitSet is a wrapper around the C function g_strsplit_set.
func StrsplitSet(string string, delimiters string, maxTokens int32) {}

// Blacklisted function: g_unichar_get_mirror_char

// Unsupported function: g_unsetenv : unsupported parameter variable : type filename, const gchar*

// Unsupported function: g_vasprintf : unsupported parameter args : type va_list, va_list
