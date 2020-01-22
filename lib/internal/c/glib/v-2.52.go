// Code generated - DO NOT EDIT.
// +build glib_2.52 glib_2.54 glib_2.56 glib_2.58 glib_2.60 glib_2.62

package glib

// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-object.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : TestLogMsg : blacklisted
// UNSUPPORTED : g_array_append_vals : blacklisted

// UNSUPPORTED : g_array_binary_search : blacklisted

// UNSUPPORTED : g_array_copy : blacklisted

// UNSUPPORTED : g_array_free : blacklisted

// UNSUPPORTED : g_array_get_element_size : blacklisted

// UNSUPPORTED : g_array_insert_vals : blacklisted

// UNSUPPORTED : g_array_new : blacklisted

// UNSUPPORTED : g_array_prepend_vals : blacklisted

// UNSUPPORTED : g_array_ref : blacklisted

// UNSUPPORTED : g_array_remove_index : blacklisted

// UNSUPPORTED : g_array_remove_index_fast : blacklisted

// UNSUPPORTED : g_array_remove_range : blacklisted

// UNSUPPORTED : g_array_set_clear_func : blacklisted

// UNSUPPORTED : g_array_set_size : blacklisted

// UNSUPPORTED : g_array_sized_new : blacklisted

// UNSUPPORTED : g_array_sort : blacklisted

// UNSUPPORTED : g_array_sort_with_data : blacklisted

// UNSUPPORTED : g_array_unref : blacklisted

// UNSUPPORTED : g_async_queue_length : blacklisted

// UNSUPPORTED : g_async_queue_length_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_lock : blacklisted

// UNSUPPORTED : g_async_queue_pop : blacklisted

// UNSUPPORTED : g_async_queue_pop_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_push : blacklisted

// UNSUPPORTED : g_async_queue_push_front : blacklisted

// UNSUPPORTED : g_async_queue_push_front_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_push_sorted : blacklisted

// UNSUPPORTED : g_async_queue_push_sorted_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_push_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_ref : blacklisted

// UNSUPPORTED : g_async_queue_ref_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_remove : blacklisted

// UNSUPPORTED : g_async_queue_remove_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_sort : blacklisted

// UNSUPPORTED : g_async_queue_sort_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_timed_pop : blacklisted

// UNSUPPORTED : g_async_queue_timed_pop_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_timeout_pop : blacklisted

// UNSUPPORTED : g_async_queue_timeout_pop_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_try_pop : blacklisted

// UNSUPPORTED : g_async_queue_try_pop_unlocked : blacklisted

// UNSUPPORTED : g_async_queue_unlock : blacklisted

// UNSUPPORTED : g_async_queue_unref : blacklisted

// UNSUPPORTED : g_async_queue_unref_and_unlock : blacklisted

// UNSUPPORTED : g_async_queue_new : blacklisted

// UNSUPPORTED : g_async_queue_new_full : blacklisted

// UNSUPPORTED : g_bookmark_file_add_application : blacklisted

// UNSUPPORTED : g_bookmark_file_add_group : blacklisted

// UNSUPPORTED : g_bookmark_file_free : blacklisted

// UNSUPPORTED : g_bookmark_file_get_added : blacklisted

// UNSUPPORTED : g_bookmark_file_get_app_info : blacklisted

// UNSUPPORTED : g_bookmark_file_get_applications : blacklisted

// UNSUPPORTED : g_bookmark_file_get_description : blacklisted

// UNSUPPORTED : g_bookmark_file_get_groups : blacklisted

// UNSUPPORTED : g_bookmark_file_get_icon : blacklisted

// UNSUPPORTED : g_bookmark_file_get_is_private : blacklisted

// UNSUPPORTED : g_bookmark_file_get_mime_type : blacklisted

// UNSUPPORTED : g_bookmark_file_get_modified : blacklisted

// UNSUPPORTED : g_bookmark_file_get_size : blacklisted

// UNSUPPORTED : g_bookmark_file_get_title : blacklisted

// UNSUPPORTED : g_bookmark_file_get_uris : blacklisted

// UNSUPPORTED : g_bookmark_file_get_visited : blacklisted

// UNSUPPORTED : g_bookmark_file_has_application : blacklisted

// UNSUPPORTED : g_bookmark_file_has_group : blacklisted

// UNSUPPORTED : g_bookmark_file_has_item : blacklisted

// UNSUPPORTED : g_bookmark_file_load_from_data : blacklisted

// UNSUPPORTED : g_bookmark_file_load_from_data_dirs : blacklisted

// UNSUPPORTED : g_bookmark_file_load_from_file : blacklisted

// UNSUPPORTED : g_bookmark_file_move_item : blacklisted

// UNSUPPORTED : g_bookmark_file_remove_application : blacklisted

// UNSUPPORTED : g_bookmark_file_remove_group : blacklisted

// UNSUPPORTED : g_bookmark_file_remove_item : blacklisted

// UNSUPPORTED : g_bookmark_file_set_added : blacklisted

// UNSUPPORTED : g_bookmark_file_set_app_info : blacklisted

// UNSUPPORTED : g_bookmark_file_set_description : blacklisted

// UNSUPPORTED : g_bookmark_file_set_groups : blacklisted

// UNSUPPORTED : g_bookmark_file_set_icon : blacklisted

// UNSUPPORTED : g_bookmark_file_set_is_private : blacklisted

// UNSUPPORTED : g_bookmark_file_set_mime_type : blacklisted

// UNSUPPORTED : g_bookmark_file_set_modified : blacklisted

// UNSUPPORTED : g_bookmark_file_set_title : blacklisted

// UNSUPPORTED : g_bookmark_file_set_visited : blacklisted

// UNSUPPORTED : g_bookmark_file_to_data : blacklisted

// UNSUPPORTED : g_bookmark_file_to_file : blacklisted

// UNSUPPORTED : g_bookmark_file_error_quark : blacklisted

// UNSUPPORTED : g_bookmark_file_new : blacklisted

// UNSUPPORTED : g_byte_array_append : blacklisted

// UNSUPPORTED : g_byte_array_free : blacklisted

// UNSUPPORTED : g_byte_array_free_to_bytes : blacklisted

// UNSUPPORTED : g_byte_array_new : blacklisted

// UNSUPPORTED : g_byte_array_new_take : blacklisted

// UNSUPPORTED : g_byte_array_prepend : blacklisted

// UNSUPPORTED : g_byte_array_ref : blacklisted

// UNSUPPORTED : g_byte_array_remove_index : blacklisted

// UNSUPPORTED : g_byte_array_remove_index_fast : blacklisted

// UNSUPPORTED : g_byte_array_remove_range : blacklisted

// UNSUPPORTED : g_byte_array_set_size : blacklisted

// UNSUPPORTED : g_byte_array_sized_new : blacklisted

// UNSUPPORTED : g_byte_array_sort : blacklisted

// UNSUPPORTED : g_byte_array_sort_with_data : blacklisted

// UNSUPPORTED : g_byte_array_unref : blacklisted

// UNSUPPORTED : g_bytes_new : blacklisted

// UNSUPPORTED : g_bytes_new_static : blacklisted

// UNSUPPORTED : g_bytes_new_take : blacklisted

// UNSUPPORTED : g_bytes_new_with_free_func : blacklisted

// UNSUPPORTED : g_bytes_compare : blacklisted

// UNSUPPORTED : g_bytes_equal : blacklisted

// UNSUPPORTED : g_bytes_get_data : blacklisted

// UNSUPPORTED : g_bytes_get_size : blacklisted

// UNSUPPORTED : g_bytes_hash : blacklisted

// UNSUPPORTED : g_bytes_new_from_bytes : blacklisted

// UNSUPPORTED : g_bytes_ref : blacklisted

// UNSUPPORTED : g_bytes_unref : blacklisted

// UNSUPPORTED : g_bytes_unref_to_array : blacklisted

// UNSUPPORTED : g_bytes_unref_to_data : blacklisted

// UNSUPPORTED : g_checksum_new : blacklisted

// UNSUPPORTED : g_checksum_copy : blacklisted

// UNSUPPORTED : g_checksum_free : blacklisted

// UNSUPPORTED : g_checksum_get_digest : blacklisted

// UNSUPPORTED : g_checksum_get_string : blacklisted

// UNSUPPORTED : g_checksum_reset : blacklisted

// UNSUPPORTED : g_checksum_update : blacklisted

// UNSUPPORTED : g_checksum_type_get_length : blacklisted

// UNSUPPORTED : g_cond_broadcast : blacklisted

// UNSUPPORTED : g_cond_clear : blacklisted

// UNSUPPORTED : g_cond_init : blacklisted

// UNSUPPORTED : g_cond_signal : blacklisted

// UNSUPPORTED : g_cond_wait : blacklisted

// UNSUPPORTED : g_cond_wait_until : blacklisted

// UNSUPPORTED : g_date_new : blacklisted

// UNSUPPORTED : g_date_new_dmy : blacklisted

// UNSUPPORTED : g_date_new_julian : blacklisted

// UNSUPPORTED : g_date_add_days : blacklisted

// UNSUPPORTED : g_date_add_months : blacklisted

// UNSUPPORTED : g_date_add_years : blacklisted

// UNSUPPORTED : g_date_clamp : blacklisted

// UNSUPPORTED : g_date_clear : blacklisted

// UNSUPPORTED : g_date_compare : blacklisted

// UNSUPPORTED : g_date_copy : blacklisted

// UNSUPPORTED : g_date_days_between : blacklisted

// UNSUPPORTED : g_date_free : blacklisted

// UNSUPPORTED : g_date_get_day : blacklisted

// UNSUPPORTED : g_date_get_day_of_year : blacklisted

// UNSUPPORTED : g_date_get_iso8601_week_of_year : blacklisted

// UNSUPPORTED : g_date_get_julian : blacklisted

// UNSUPPORTED : g_date_get_monday_week_of_year : blacklisted

// UNSUPPORTED : g_date_get_month : blacklisted

// UNSUPPORTED : g_date_get_sunday_week_of_year : blacklisted

// UNSUPPORTED : g_date_get_weekday : blacklisted

// UNSUPPORTED : g_date_get_year : blacklisted

// UNSUPPORTED : g_date_is_first_of_month : blacklisted

// UNSUPPORTED : g_date_is_last_of_month : blacklisted

// UNSUPPORTED : g_date_order : blacklisted

// UNSUPPORTED : g_date_set_day : blacklisted

// UNSUPPORTED : g_date_set_dmy : blacklisted

// UNSUPPORTED : g_date_set_julian : blacklisted

// UNSUPPORTED : g_date_set_month : blacklisted

// UNSUPPORTED : g_date_set_parse : blacklisted

// UNSUPPORTED : g_date_set_time : blacklisted

// UNSUPPORTED : g_date_set_time_t : blacklisted

// UNSUPPORTED : g_date_set_time_val : blacklisted

// UNSUPPORTED : g_date_set_year : blacklisted

// UNSUPPORTED : g_date_subtract_days : blacklisted

// UNSUPPORTED : g_date_subtract_months : blacklisted

// UNSUPPORTED : g_date_subtract_years : blacklisted

// UNSUPPORTED : g_date_to_struct_tm : blacklisted

// UNSUPPORTED : g_date_valid : blacklisted

// UNSUPPORTED : g_date_get_days_in_month : blacklisted

// UNSUPPORTED : g_date_get_monday_weeks_in_year : blacklisted

// UNSUPPORTED : g_date_get_sunday_weeks_in_year : blacklisted

// UNSUPPORTED : g_date_is_leap_year : blacklisted

// UNSUPPORTED : g_date_strftime : blacklisted

// UNSUPPORTED : g_date_valid_day : blacklisted

// UNSUPPORTED : g_date_valid_dmy : blacklisted

// UNSUPPORTED : g_date_valid_julian : blacklisted

// UNSUPPORTED : g_date_valid_month : blacklisted

// UNSUPPORTED : g_date_valid_weekday : blacklisted

// UNSUPPORTED : g_date_valid_year : blacklisted

// UNSUPPORTED : g_date_time_new : blacklisted

// UNSUPPORTED : g_date_time_new_from_iso8601 : blacklisted

// UNSUPPORTED : g_date_time_new_from_timeval_local : blacklisted

// UNSUPPORTED : g_date_time_new_from_timeval_utc : blacklisted

// UNSUPPORTED : g_date_time_new_from_unix_local : blacklisted

// UNSUPPORTED : g_date_time_new_from_unix_utc : blacklisted

// UNSUPPORTED : g_date_time_new_local : blacklisted

// UNSUPPORTED : g_date_time_new_now : blacklisted

// UNSUPPORTED : g_date_time_new_now_local : blacklisted

// UNSUPPORTED : g_date_time_new_now_utc : blacklisted

// UNSUPPORTED : g_date_time_new_utc : blacklisted

// UNSUPPORTED : g_date_time_add : blacklisted

// UNSUPPORTED : g_date_time_add_days : blacklisted

// UNSUPPORTED : g_date_time_add_full : blacklisted

// UNSUPPORTED : g_date_time_add_hours : blacklisted

// UNSUPPORTED : g_date_time_add_minutes : blacklisted

// UNSUPPORTED : g_date_time_add_months : blacklisted

// UNSUPPORTED : g_date_time_add_seconds : blacklisted

// UNSUPPORTED : g_date_time_add_weeks : blacklisted

// UNSUPPORTED : g_date_time_add_years : blacklisted

// UNSUPPORTED : g_date_time_difference : blacklisted

// UNSUPPORTED : g_date_time_format : blacklisted

// UNSUPPORTED : g_date_time_format_iso8601 : blacklisted

// UNSUPPORTED : g_date_time_get_day_of_month : blacklisted

// UNSUPPORTED : g_date_time_get_day_of_week : blacklisted

// UNSUPPORTED : g_date_time_get_day_of_year : blacklisted

// UNSUPPORTED : g_date_time_get_hour : blacklisted

// UNSUPPORTED : g_date_time_get_microsecond : blacklisted

// UNSUPPORTED : g_date_time_get_minute : blacklisted

// UNSUPPORTED : g_date_time_get_month : blacklisted

// UNSUPPORTED : g_date_time_get_second : blacklisted

// UNSUPPORTED : g_date_time_get_seconds : blacklisted

// UNSUPPORTED : g_date_time_get_timezone : blacklisted

// UNSUPPORTED : g_date_time_get_timezone_abbreviation : blacklisted

// UNSUPPORTED : g_date_time_get_utc_offset : blacklisted

// UNSUPPORTED : g_date_time_get_week_numbering_year : blacklisted

// UNSUPPORTED : g_date_time_get_week_of_year : blacklisted

// UNSUPPORTED : g_date_time_get_year : blacklisted

// UNSUPPORTED : g_date_time_get_ymd : blacklisted

// UNSUPPORTED : g_date_time_is_daylight_savings : blacklisted

// UNSUPPORTED : g_date_time_ref : blacklisted

// UNSUPPORTED : g_date_time_to_local : blacklisted

// UNSUPPORTED : g_date_time_to_timeval : blacklisted

// UNSUPPORTED : g_date_time_to_timezone : blacklisted

// UNSUPPORTED : g_date_time_to_unix : blacklisted

// UNSUPPORTED : g_date_time_to_utc : blacklisted

// UNSUPPORTED : g_date_time_unref : blacklisted

// UNSUPPORTED : g_date_time_compare : blacklisted

// UNSUPPORTED : g_date_time_equal : blacklisted

// UNSUPPORTED : g_date_time_hash : blacklisted

// UNSUPPORTED : g_dir_close : blacklisted

// UNSUPPORTED : g_dir_read_name : blacklisted

// UNSUPPORTED : g_dir_rewind : blacklisted

// UNSUPPORTED : g_dir_make_tmp : blacklisted

// UNSUPPORTED : g_dir_open : blacklisted

// UNSUPPORTED : g_error_new : blacklisted

// UNSUPPORTED : g_error_new_literal : blacklisted

// UNSUPPORTED : g_error_new_valist : blacklisted

// UNSUPPORTED : g_error_copy : blacklisted

// UNSUPPORTED : g_error_free : blacklisted

// UNSUPPORTED : g_error_matches : blacklisted

// UNSUPPORTED : g_hash_table_add : blacklisted

// UNSUPPORTED : g_hash_table_contains : blacklisted

// UNSUPPORTED : g_hash_table_destroy : blacklisted

// UNSUPPORTED : g_hash_table_find : blacklisted

// UNSUPPORTED : g_hash_table_foreach : blacklisted

// UNSUPPORTED : g_hash_table_foreach_remove : blacklisted

// UNSUPPORTED : g_hash_table_foreach_steal : blacklisted

// UNSUPPORTED : g_hash_table_get_keys : blacklisted

// UNSUPPORTED : g_hash_table_get_keys_as_array : blacklisted

// UNSUPPORTED : g_hash_table_get_values : blacklisted

// UNSUPPORTED : g_hash_table_insert : blacklisted

// UNSUPPORTED : g_hash_table_lookup : blacklisted

// UNSUPPORTED : g_hash_table_lookup_extended : blacklisted

// UNSUPPORTED : g_hash_table_new : blacklisted

// UNSUPPORTED : g_hash_table_new_full : blacklisted

// UNSUPPORTED : g_hash_table_ref : blacklisted

// UNSUPPORTED : g_hash_table_remove : blacklisted

// UNSUPPORTED : g_hash_table_remove_all : blacklisted

// UNSUPPORTED : g_hash_table_replace : blacklisted

// UNSUPPORTED : g_hash_table_size : blacklisted

// UNSUPPORTED : g_hash_table_steal : blacklisted

// UNSUPPORTED : g_hash_table_steal_all : blacklisted

// UNSUPPORTED : g_hash_table_steal_extended : blacklisted

// UNSUPPORTED : g_hash_table_unref : blacklisted

// UNSUPPORTED : g_hash_table_iter_get_hash_table : blacklisted

// UNSUPPORTED : g_hash_table_iter_init : blacklisted

// UNSUPPORTED : g_hash_table_iter_next : blacklisted

// UNSUPPORTED : g_hash_table_iter_remove : blacklisted

// UNSUPPORTED : g_hash_table_iter_replace : blacklisted

// UNSUPPORTED : g_hash_table_iter_steal : blacklisted

// UNSUPPORTED : g_hmac_copy : blacklisted

// UNSUPPORTED : g_hmac_get_digest : blacklisted

// UNSUPPORTED : g_hmac_get_string : blacklisted

// UNSUPPORTED : g_hmac_ref : blacklisted

// UNSUPPORTED : g_hmac_unref : blacklisted

// UNSUPPORTED : g_hmac_update : blacklisted

// UNSUPPORTED : g_hmac_new : blacklisted

// UNSUPPORTED : g_hook_compare_ids : blacklisted

// UNSUPPORTED : g_hook_alloc : blacklisted

// UNSUPPORTED : g_hook_destroy : blacklisted

// UNSUPPORTED : g_hook_destroy_link : blacklisted

// UNSUPPORTED : g_hook_find : blacklisted

// UNSUPPORTED : g_hook_find_data : blacklisted

// UNSUPPORTED : g_hook_find_func : blacklisted

// UNSUPPORTED : g_hook_find_func_data : blacklisted

// UNSUPPORTED : g_hook_first_valid : blacklisted

// UNSUPPORTED : g_hook_free : blacklisted

// UNSUPPORTED : g_hook_get : blacklisted

// UNSUPPORTED : g_hook_insert_before : blacklisted

// UNSUPPORTED : g_hook_insert_sorted : blacklisted

// UNSUPPORTED : g_hook_next_valid : blacklisted

// UNSUPPORTED : g_hook_prepend : blacklisted

// UNSUPPORTED : g_hook_ref : blacklisted

// UNSUPPORTED : g_hook_unref : blacklisted

// UNSUPPORTED : g_hook_list_clear : blacklisted

// UNSUPPORTED : g_hook_list_init : blacklisted

// UNSUPPORTED : g_hook_list_invoke : blacklisted

// UNSUPPORTED : g_hook_list_invoke_check : blacklisted

// UNSUPPORTED : g_hook_list_marshal : blacklisted

// UNSUPPORTED : g_hook_list_marshal_check : blacklisted

// UNSUPPORTED : g_iconv : blacklisted

// UNSUPPORTED : g_iconv_close : blacklisted

// UNSUPPORTED : g_iconv_open : blacklisted

// UNSUPPORTED : g_io_channel_new_file : blacklisted

// UNSUPPORTED : g_io_channel_unix_new : blacklisted

// UNSUPPORTED : g_io_channel_close : blacklisted

// UNSUPPORTED : g_io_channel_flush : blacklisted

// UNSUPPORTED : g_io_channel_get_buffer_condition : blacklisted

// UNSUPPORTED : g_io_channel_get_buffer_size : blacklisted

// UNSUPPORTED : g_io_channel_get_buffered : blacklisted

// UNSUPPORTED : g_io_channel_get_close_on_unref : blacklisted

// UNSUPPORTED : g_io_channel_get_encoding : blacklisted

// UNSUPPORTED : g_io_channel_get_flags : blacklisted

// UNSUPPORTED : g_io_channel_get_line_term : blacklisted

// UNSUPPORTED : g_io_channel_init : blacklisted

// UNSUPPORTED : g_io_channel_read : blacklisted

// UNSUPPORTED : g_io_channel_read_chars : blacklisted

// UNSUPPORTED : g_io_channel_read_line : blacklisted

// UNSUPPORTED : g_io_channel_read_line_string : blacklisted

// UNSUPPORTED : g_io_channel_read_to_end : blacklisted

// UNSUPPORTED : g_io_channel_read_unichar : blacklisted

// UNSUPPORTED : g_io_channel_ref : blacklisted

// UNSUPPORTED : g_io_channel_seek : blacklisted

// UNSUPPORTED : g_io_channel_seek_position : blacklisted

// UNSUPPORTED : g_io_channel_set_buffer_size : blacklisted

// UNSUPPORTED : g_io_channel_set_buffered : blacklisted

// UNSUPPORTED : g_io_channel_set_close_on_unref : blacklisted

// UNSUPPORTED : g_io_channel_set_encoding : blacklisted

// UNSUPPORTED : g_io_channel_set_flags : blacklisted

// UNSUPPORTED : g_io_channel_set_line_term : blacklisted

// UNSUPPORTED : g_io_channel_shutdown : blacklisted

// UNSUPPORTED : g_io_channel_unix_get_fd : blacklisted

// UNSUPPORTED : g_io_channel_unref : blacklisted

// UNSUPPORTED : g_io_channel_write : blacklisted

// UNSUPPORTED : g_io_channel_write_chars : blacklisted

// UNSUPPORTED : g_io_channel_write_unichar : blacklisted

// UNSUPPORTED : g_io_channel_error_from_errno : blacklisted

// UNSUPPORTED : g_io_channel_error_quark : blacklisted

// UNSUPPORTED : g_key_file_new : blacklisted

// UNSUPPORTED : g_key_file_free : blacklisted

// UNSUPPORTED : g_key_file_get_boolean : blacklisted

// UNSUPPORTED : g_key_file_get_boolean_list : blacklisted

// UNSUPPORTED : g_key_file_get_comment : blacklisted

// UNSUPPORTED : g_key_file_get_double : blacklisted

// UNSUPPORTED : g_key_file_get_double_list : blacklisted

// UNSUPPORTED : g_key_file_get_groups : blacklisted

// UNSUPPORTED : g_key_file_get_int64 : blacklisted

// UNSUPPORTED : g_key_file_get_integer : blacklisted

// UNSUPPORTED : g_key_file_get_integer_list : blacklisted

// UNSUPPORTED : g_key_file_get_keys : blacklisted

// UNSUPPORTED : g_key_file_get_locale_for_key : blacklisted

// UNSUPPORTED : g_key_file_get_locale_string : blacklisted

// UNSUPPORTED : g_key_file_get_locale_string_list : blacklisted

// UNSUPPORTED : g_key_file_get_start_group : blacklisted

// UNSUPPORTED : g_key_file_get_string : blacklisted

// UNSUPPORTED : g_key_file_get_string_list : blacklisted

// UNSUPPORTED : g_key_file_get_uint64 : blacklisted

// UNSUPPORTED : g_key_file_get_value : blacklisted

// UNSUPPORTED : g_key_file_has_group : blacklisted

// UNSUPPORTED : g_key_file_has_key : blacklisted

// UNSUPPORTED : g_key_file_load_from_bytes : blacklisted

// UNSUPPORTED : g_key_file_load_from_data : blacklisted

// UNSUPPORTED : g_key_file_load_from_data_dirs : blacklisted

// UNSUPPORTED : g_key_file_load_from_dirs : blacklisted

// UNSUPPORTED : g_key_file_load_from_file : blacklisted

// UNSUPPORTED : g_key_file_ref : blacklisted

// UNSUPPORTED : g_key_file_remove_comment : blacklisted

// UNSUPPORTED : g_key_file_remove_group : blacklisted

// UNSUPPORTED : g_key_file_remove_key : blacklisted

// UNSUPPORTED : g_key_file_save_to_file : blacklisted

// UNSUPPORTED : g_key_file_set_boolean : blacklisted

// UNSUPPORTED : g_key_file_set_boolean_list : blacklisted

// UNSUPPORTED : g_key_file_set_comment : blacklisted

// UNSUPPORTED : g_key_file_set_double : blacklisted

// UNSUPPORTED : g_key_file_set_double_list : blacklisted

// UNSUPPORTED : g_key_file_set_int64 : blacklisted

// UNSUPPORTED : g_key_file_set_integer : blacklisted

// UNSUPPORTED : g_key_file_set_integer_list : blacklisted

// UNSUPPORTED : g_key_file_set_list_separator : blacklisted

// UNSUPPORTED : g_key_file_set_locale_string : blacklisted

// UNSUPPORTED : g_key_file_set_locale_string_list : blacklisted

// UNSUPPORTED : g_key_file_set_string : blacklisted

// UNSUPPORTED : g_key_file_set_string_list : blacklisted

// UNSUPPORTED : g_key_file_set_uint64 : blacklisted

// UNSUPPORTED : g_key_file_set_value : blacklisted

// UNSUPPORTED : g_key_file_to_data : blacklisted

// UNSUPPORTED : g_key_file_unref : blacklisted

// UNSUPPORTED : g_key_file_error_quark : blacklisted

// UNSUPPORTED : g_list_alloc : blacklisted

// UNSUPPORTED : g_list_append : blacklisted

// UNSUPPORTED : g_list_concat : blacklisted

// UNSUPPORTED : g_list_copy : blacklisted

// UNSUPPORTED : g_list_copy_deep : blacklisted

// UNSUPPORTED : g_list_delete_link : blacklisted

// UNSUPPORTED : g_list_find : blacklisted

// UNSUPPORTED : g_list_find_custom : blacklisted

// UNSUPPORTED : g_list_first : blacklisted

// UNSUPPORTED : g_list_foreach : blacklisted

// UNSUPPORTED : g_list_free : blacklisted

// UNSUPPORTED : g_list_free_1 : blacklisted

// UNSUPPORTED : g_list_free_full : blacklisted

// UNSUPPORTED : g_list_index : blacklisted

// UNSUPPORTED : g_list_insert : blacklisted

// UNSUPPORTED : g_list_insert_before : blacklisted

// UNSUPPORTED : g_list_insert_before_link : blacklisted

// UNSUPPORTED : g_list_insert_sorted : blacklisted

// UNSUPPORTED : g_list_insert_sorted_with_data : blacklisted

// UNSUPPORTED : g_list_last : blacklisted

// UNSUPPORTED : g_list_length : blacklisted

// UNSUPPORTED : g_list_nth : blacklisted

// UNSUPPORTED : g_list_nth_data : blacklisted

// UNSUPPORTED : g_list_nth_prev : blacklisted

// UNSUPPORTED : g_list_position : blacklisted

// UNSUPPORTED : g_list_prepend : blacklisted

// UNSUPPORTED : g_list_remove : blacklisted

// UNSUPPORTED : g_list_remove_all : blacklisted

// UNSUPPORTED : g_list_remove_link : blacklisted

// UNSUPPORTED : g_list_reverse : blacklisted

// UNSUPPORTED : g_list_sort : blacklisted

// UNSUPPORTED : g_list_sort_with_data : blacklisted

// UNSUPPORTED : g_main_context_new : blacklisted

// UNSUPPORTED : g_main_context_acquire : blacklisted

// UNSUPPORTED : g_main_context_add_poll : blacklisted

// UNSUPPORTED : g_main_context_check : blacklisted

// UNSUPPORTED : g_main_context_dispatch : blacklisted

// UNSUPPORTED : g_main_context_find_source_by_funcs_user_data : blacklisted

// UNSUPPORTED : g_main_context_find_source_by_id : blacklisted

// UNSUPPORTED : g_main_context_find_source_by_user_data : blacklisted

// UNSUPPORTED : g_main_context_get_poll_func : blacklisted

// UNSUPPORTED : g_main_context_invoke : blacklisted

// UNSUPPORTED : g_main_context_invoke_full : blacklisted

// UNSUPPORTED : g_main_context_is_owner : blacklisted

// UNSUPPORTED : g_main_context_iteration : blacklisted

// UNSUPPORTED : g_main_context_pending : blacklisted

// UNSUPPORTED : g_main_context_pop_thread_default : blacklisted

// UNSUPPORTED : g_main_context_prepare : blacklisted

// UNSUPPORTED : g_main_context_push_thread_default : blacklisted

// UNSUPPORTED : g_main_context_query : blacklisted

// UNSUPPORTED : g_main_context_ref : blacklisted

// UNSUPPORTED : g_main_context_release : blacklisted

// UNSUPPORTED : g_main_context_remove_poll : blacklisted

// UNSUPPORTED : g_main_context_set_poll_func : blacklisted

// UNSUPPORTED : g_main_context_unref : blacklisted

// UNSUPPORTED : g_main_context_wait : blacklisted

// UNSUPPORTED : g_main_context_wakeup : blacklisted

// UNSUPPORTED : g_main_context_default : blacklisted

// UNSUPPORTED : g_main_context_get_thread_default : blacklisted

// UNSUPPORTED : g_main_context_ref_thread_default : blacklisted

// UNSUPPORTED : g_main_loop_new : blacklisted

// UNSUPPORTED : g_main_loop_get_context : blacklisted

// UNSUPPORTED : g_main_loop_is_running : blacklisted

// UNSUPPORTED : g_main_loop_quit : blacklisted

// UNSUPPORTED : g_main_loop_ref : blacklisted

// UNSUPPORTED : g_main_loop_run : blacklisted

// UNSUPPORTED : g_main_loop_unref : blacklisted

// UNSUPPORTED : g_mapped_file_new : blacklisted

// UNSUPPORTED : g_mapped_file_new_from_fd : blacklisted

// UNSUPPORTED : g_mapped_file_free : blacklisted

// UNSUPPORTED : g_mapped_file_get_bytes : blacklisted

// UNSUPPORTED : g_mapped_file_get_contents : blacklisted

// UNSUPPORTED : g_mapped_file_get_length : blacklisted

// UNSUPPORTED : g_mapped_file_ref : blacklisted

// UNSUPPORTED : g_mapped_file_unref : blacklisted

// UNSUPPORTED : g_markup_parse_context_new : blacklisted

// UNSUPPORTED : g_markup_parse_context_end_parse : blacklisted

// UNSUPPORTED : g_markup_parse_context_free : blacklisted

// UNSUPPORTED : g_markup_parse_context_get_element : blacklisted

// UNSUPPORTED : g_markup_parse_context_get_element_stack : blacklisted

// UNSUPPORTED : g_markup_parse_context_get_position : blacklisted

// UNSUPPORTED : g_markup_parse_context_get_user_data : blacklisted

// UNSUPPORTED : g_markup_parse_context_parse : blacklisted

// UNSUPPORTED : g_markup_parse_context_pop : blacklisted

// UNSUPPORTED : g_markup_parse_context_push : blacklisted

// UNSUPPORTED : g_markup_parse_context_ref : blacklisted

// UNSUPPORTED : g_markup_parse_context_unref : blacklisted

// UNSUPPORTED : g_match_info_expand_references : blacklisted

// UNSUPPORTED : g_match_info_fetch : blacklisted

// UNSUPPORTED : g_match_info_fetch_all : blacklisted

// UNSUPPORTED : g_match_info_fetch_named : blacklisted

// UNSUPPORTED : g_match_info_fetch_named_pos : blacklisted

// UNSUPPORTED : g_match_info_fetch_pos : blacklisted

// UNSUPPORTED : g_match_info_free : blacklisted

// UNSUPPORTED : g_match_info_get_match_count : blacklisted

// UNSUPPORTED : g_match_info_get_regex : blacklisted

// UNSUPPORTED : g_match_info_get_string : blacklisted

// UNSUPPORTED : g_match_info_is_partial_match : blacklisted

// UNSUPPORTED : g_match_info_matches : blacklisted

// UNSUPPORTED : g_match_info_next : blacklisted

// UNSUPPORTED : g_match_info_ref : blacklisted

// UNSUPPORTED : g_match_info_unref : blacklisted

// UNSUPPORTED : g_node_child_index : blacklisted

// UNSUPPORTED : g_node_child_position : blacklisted

// UNSUPPORTED : g_node_children_foreach : blacklisted

// UNSUPPORTED : g_node_copy : blacklisted

// UNSUPPORTED : g_node_copy_deep : blacklisted

// UNSUPPORTED : g_node_depth : blacklisted

// UNSUPPORTED : g_node_destroy : blacklisted

// UNSUPPORTED : g_node_find : blacklisted

// UNSUPPORTED : g_node_find_child : blacklisted

// UNSUPPORTED : g_node_first_sibling : blacklisted

// UNSUPPORTED : g_node_get_root : blacklisted

// UNSUPPORTED : g_node_insert : blacklisted

// UNSUPPORTED : g_node_insert_after : blacklisted

// UNSUPPORTED : g_node_insert_before : blacklisted

// UNSUPPORTED : g_node_is_ancestor : blacklisted

// UNSUPPORTED : g_node_last_child : blacklisted

// UNSUPPORTED : g_node_last_sibling : blacklisted

// UNSUPPORTED : g_node_max_height : blacklisted

// UNSUPPORTED : g_node_n_children : blacklisted

// UNSUPPORTED : g_node_n_nodes : blacklisted

// UNSUPPORTED : g_node_nth_child : blacklisted

// UNSUPPORTED : g_node_prepend : blacklisted

// UNSUPPORTED : g_node_reverse_children : blacklisted

// UNSUPPORTED : g_node_traverse : blacklisted

// UNSUPPORTED : g_node_unlink : blacklisted

// UNSUPPORTED : g_node_new : blacklisted

// UNSUPPORTED : g_once_impl : blacklisted

// UNSUPPORTED : g_once_init_enter : blacklisted

// UNSUPPORTED : g_once_init_leave : blacklisted

// UNSUPPORTED : g_option_context_add_group : blacklisted

// UNSUPPORTED : g_option_context_add_main_entries : blacklisted

// UNSUPPORTED : g_option_context_free : blacklisted

// UNSUPPORTED : g_option_context_get_description : blacklisted

// UNSUPPORTED : g_option_context_get_help : blacklisted

// UNSUPPORTED : g_option_context_get_help_enabled : blacklisted

// UNSUPPORTED : g_option_context_get_ignore_unknown_options : blacklisted

// UNSUPPORTED : g_option_context_get_main_group : blacklisted

// UNSUPPORTED : g_option_context_get_strict_posix : blacklisted

// UNSUPPORTED : g_option_context_get_summary : blacklisted

// UNSUPPORTED : g_option_context_parse : blacklisted

// UNSUPPORTED : g_option_context_parse_strv : blacklisted

// UNSUPPORTED : g_option_context_set_description : blacklisted

// UNSUPPORTED : g_option_context_set_help_enabled : blacklisted

// UNSUPPORTED : g_option_context_set_ignore_unknown_options : blacklisted

// UNSUPPORTED : g_option_context_set_main_group : blacklisted

// UNSUPPORTED : g_option_context_set_strict_posix : blacklisted

// UNSUPPORTED : g_option_context_set_summary : blacklisted

// UNSUPPORTED : g_option_context_set_translate_func : blacklisted

// UNSUPPORTED : g_option_context_set_translation_domain : blacklisted

// UNSUPPORTED : g_option_context_new : blacklisted

// UNSUPPORTED : g_option_group_new : blacklisted

// UNSUPPORTED : g_option_group_add_entries : blacklisted

// UNSUPPORTED : g_option_group_free : blacklisted

// UNSUPPORTED : g_option_group_ref : blacklisted

// UNSUPPORTED : g_option_group_set_error_hook : blacklisted

// UNSUPPORTED : g_option_group_set_parse_hooks : blacklisted

// UNSUPPORTED : g_option_group_set_translate_func : blacklisted

// UNSUPPORTED : g_option_group_set_translation_domain : blacklisted

// UNSUPPORTED : g_option_group_unref : blacklisted

// UNSUPPORTED : g_pattern_spec_equal : blacklisted

// UNSUPPORTED : g_pattern_spec_free : blacklisted

// UNSUPPORTED : g_pattern_spec_new : blacklisted

// UNSUPPORTED : g_private_get : blacklisted

// UNSUPPORTED : g_private_replace : blacklisted

// UNSUPPORTED : g_private_set : blacklisted

// UNSUPPORTED : g_ptr_array_add : blacklisted

// UNSUPPORTED : g_ptr_array_copy : blacklisted

// UNSUPPORTED : g_ptr_array_extend : blacklisted

// UNSUPPORTED : g_ptr_array_extend_and_steal : blacklisted

// UNSUPPORTED : g_ptr_array_find : blacklisted

// UNSUPPORTED : g_ptr_array_find_with_equal_func : blacklisted

// UNSUPPORTED : g_ptr_array_foreach : blacklisted

// UNSUPPORTED : g_ptr_array_free : blacklisted

// UNSUPPORTED : g_ptr_array_insert : blacklisted

// UNSUPPORTED : g_ptr_array_new : blacklisted

// UNSUPPORTED : g_ptr_array_new_full : blacklisted

// UNSUPPORTED : g_ptr_array_new_with_free_func : blacklisted

// UNSUPPORTED : g_ptr_array_ref : blacklisted

// UNSUPPORTED : g_ptr_array_remove : blacklisted

// UNSUPPORTED : g_ptr_array_remove_fast : blacklisted

// UNSUPPORTED : g_ptr_array_remove_index : blacklisted

// UNSUPPORTED : g_ptr_array_remove_index_fast : blacklisted

// UNSUPPORTED : g_ptr_array_remove_range : blacklisted

// UNSUPPORTED : g_ptr_array_set_free_func : blacklisted

// UNSUPPORTED : g_ptr_array_set_size : blacklisted

// UNSUPPORTED : g_ptr_array_sized_new : blacklisted

// UNSUPPORTED : g_ptr_array_sort : blacklisted

// UNSUPPORTED : g_ptr_array_sort_with_data : blacklisted

// UNSUPPORTED : g_ptr_array_steal_index : blacklisted

// UNSUPPORTED : g_ptr_array_steal_index_fast : blacklisted

// UNSUPPORTED : g_ptr_array_unref : blacklisted

// UNSUPPORTED : g_queue_clear : blacklisted

// UNSUPPORTED : g_queue_clear_full : blacklisted

// UNSUPPORTED : g_queue_copy : blacklisted

// UNSUPPORTED : g_queue_delete_link : blacklisted

// UNSUPPORTED : g_queue_find : blacklisted

// UNSUPPORTED : g_queue_find_custom : blacklisted

// UNSUPPORTED : g_queue_foreach : blacklisted

// UNSUPPORTED : g_queue_free : blacklisted

// UNSUPPORTED : g_queue_free_full : blacklisted

// UNSUPPORTED : g_queue_get_length : blacklisted

// UNSUPPORTED : g_queue_index : blacklisted

// UNSUPPORTED : g_queue_init : blacklisted

// UNSUPPORTED : g_queue_insert_after : blacklisted

// UNSUPPORTED : g_queue_insert_after_link : blacklisted

// UNSUPPORTED : g_queue_insert_before : blacklisted

// UNSUPPORTED : g_queue_insert_before_link : blacklisted

// UNSUPPORTED : g_queue_insert_sorted : blacklisted

// UNSUPPORTED : g_queue_is_empty : blacklisted

// UNSUPPORTED : g_queue_link_index : blacklisted

// UNSUPPORTED : g_queue_peek_head : blacklisted

// UNSUPPORTED : g_queue_peek_head_link : blacklisted

// UNSUPPORTED : g_queue_peek_nth : blacklisted

// UNSUPPORTED : g_queue_peek_nth_link : blacklisted

// UNSUPPORTED : g_queue_peek_tail : blacklisted

// UNSUPPORTED : g_queue_peek_tail_link : blacklisted

// UNSUPPORTED : g_queue_pop_head : blacklisted

// UNSUPPORTED : g_queue_pop_head_link : blacklisted

// UNSUPPORTED : g_queue_pop_nth : blacklisted

// UNSUPPORTED : g_queue_pop_nth_link : blacklisted

// UNSUPPORTED : g_queue_pop_tail : blacklisted

// UNSUPPORTED : g_queue_pop_tail_link : blacklisted

// UNSUPPORTED : g_queue_push_head : blacklisted

// UNSUPPORTED : g_queue_push_head_link : blacklisted

// UNSUPPORTED : g_queue_push_nth : blacklisted

// UNSUPPORTED : g_queue_push_nth_link : blacklisted

// UNSUPPORTED : g_queue_push_tail : blacklisted

// UNSUPPORTED : g_queue_push_tail_link : blacklisted

// UNSUPPORTED : g_queue_remove : blacklisted

// UNSUPPORTED : g_queue_remove_all : blacklisted

// UNSUPPORTED : g_queue_reverse : blacklisted

// UNSUPPORTED : g_queue_sort : blacklisted

// UNSUPPORTED : g_queue_unlink : blacklisted

// UNSUPPORTED : g_queue_new : blacklisted

// UNSUPPORTED : g_rw_lock_clear : blacklisted

// UNSUPPORTED : g_rw_lock_init : blacklisted

// UNSUPPORTED : g_rw_lock_reader_lock : blacklisted

// UNSUPPORTED : g_rw_lock_reader_trylock : blacklisted

// UNSUPPORTED : g_rw_lock_reader_unlock : blacklisted

// UNSUPPORTED : g_rw_lock_writer_lock : blacklisted

// UNSUPPORTED : g_rw_lock_writer_trylock : blacklisted

// UNSUPPORTED : g_rw_lock_writer_unlock : blacklisted

// UNSUPPORTED : g_rand_copy : blacklisted

// UNSUPPORTED : g_rand_double : blacklisted

// UNSUPPORTED : g_rand_double_range : blacklisted

// UNSUPPORTED : g_rand_free : blacklisted

// UNSUPPORTED : g_rand_int : blacklisted

// UNSUPPORTED : g_rand_int_range : blacklisted

// UNSUPPORTED : g_rand_set_seed : blacklisted

// UNSUPPORTED : g_rand_set_seed_array : blacklisted

// UNSUPPORTED : g_rand_new : blacklisted

// UNSUPPORTED : g_rand_new_with_seed : blacklisted

// UNSUPPORTED : g_rand_new_with_seed_array : blacklisted

// UNSUPPORTED : g_rec_mutex_clear : blacklisted

// UNSUPPORTED : g_rec_mutex_init : blacklisted

// UNSUPPORTED : g_rec_mutex_lock : blacklisted

// UNSUPPORTED : g_rec_mutex_trylock : blacklisted

// UNSUPPORTED : g_rec_mutex_unlock : blacklisted

// UNSUPPORTED : g_regex_new : blacklisted

// UNSUPPORTED : g_regex_get_capture_count : blacklisted

// UNSUPPORTED : g_regex_get_compile_flags : blacklisted

// UNSUPPORTED : g_regex_get_has_cr_or_lf : blacklisted

// UNSUPPORTED : g_regex_get_match_flags : blacklisted

// UNSUPPORTED : g_regex_get_max_backref : blacklisted

// UNSUPPORTED : g_regex_get_max_lookbehind : blacklisted

// UNSUPPORTED : g_regex_get_pattern : blacklisted

// UNSUPPORTED : g_regex_get_string_number : blacklisted

// UNSUPPORTED : g_regex_match : blacklisted

// UNSUPPORTED : g_regex_match_all : blacklisted

// UNSUPPORTED : g_regex_match_all_full : blacklisted

// UNSUPPORTED : g_regex_match_full : blacklisted

// UNSUPPORTED : g_regex_ref : blacklisted

// UNSUPPORTED : g_regex_replace : blacklisted

// UNSUPPORTED : g_regex_replace_eval : blacklisted

// UNSUPPORTED : g_regex_replace_literal : blacklisted

// UNSUPPORTED : g_regex_split : blacklisted

// UNSUPPORTED : g_regex_split_full : blacklisted

// UNSUPPORTED : g_regex_unref : blacklisted

// UNSUPPORTED : g_regex_check_replacement : blacklisted

// UNSUPPORTED : g_regex_error_quark : blacklisted

// UNSUPPORTED : g_regex_escape_nul : blacklisted

// UNSUPPORTED : g_regex_escape_string : blacklisted

// UNSUPPORTED : g_regex_match_simple : blacklisted

// UNSUPPORTED : g_regex_split_simple : blacklisted

// UNSUPPORTED : g_slist_alloc : blacklisted

// UNSUPPORTED : g_slist_append : blacklisted

// UNSUPPORTED : g_slist_concat : blacklisted

// UNSUPPORTED : g_slist_copy : blacklisted

// UNSUPPORTED : g_slist_copy_deep : blacklisted

// UNSUPPORTED : g_slist_delete_link : blacklisted

// UNSUPPORTED : g_slist_find : blacklisted

// UNSUPPORTED : g_slist_find_custom : blacklisted

// UNSUPPORTED : g_slist_foreach : blacklisted

// UNSUPPORTED : g_slist_free : blacklisted

// UNSUPPORTED : g_slist_free_1 : blacklisted

// UNSUPPORTED : g_slist_free_full : blacklisted

// UNSUPPORTED : g_slist_index : blacklisted

// UNSUPPORTED : g_slist_insert : blacklisted

// UNSUPPORTED : g_slist_insert_before : blacklisted

// UNSUPPORTED : g_slist_insert_sorted : blacklisted

// UNSUPPORTED : g_slist_insert_sorted_with_data : blacklisted

// UNSUPPORTED : g_slist_last : blacklisted

// UNSUPPORTED : g_slist_length : blacklisted

// UNSUPPORTED : g_slist_nth : blacklisted

// UNSUPPORTED : g_slist_nth_data : blacklisted

// UNSUPPORTED : g_slist_position : blacklisted

// UNSUPPORTED : g_slist_prepend : blacklisted

// UNSUPPORTED : g_slist_remove : blacklisted

// UNSUPPORTED : g_slist_remove_all : blacklisted

// UNSUPPORTED : g_slist_remove_link : blacklisted

// UNSUPPORTED : g_slist_reverse : blacklisted

// UNSUPPORTED : g_slist_sort : blacklisted

// UNSUPPORTED : g_slist_sort_with_data : blacklisted

// UNSUPPORTED : g_scanner_cur_line : blacklisted

// UNSUPPORTED : g_scanner_cur_position : blacklisted

// UNSUPPORTED : g_scanner_cur_token : blacklisted

// UNSUPPORTED : g_scanner_cur_value : blacklisted

// UNSUPPORTED : g_scanner_destroy : blacklisted

// UNSUPPORTED : g_scanner_eof : blacklisted

// UNSUPPORTED : g_scanner_error : blacklisted

// UNSUPPORTED : g_scanner_get_next_token : blacklisted

// UNSUPPORTED : g_scanner_input_file : blacklisted

// UNSUPPORTED : g_scanner_input_text : blacklisted

// UNSUPPORTED : g_scanner_lookup_symbol : blacklisted

// UNSUPPORTED : g_scanner_peek_next_token : blacklisted

// UNSUPPORTED : g_scanner_scope_add_symbol : blacklisted

// UNSUPPORTED : g_scanner_scope_foreach_symbol : blacklisted

// UNSUPPORTED : g_scanner_scope_lookup_symbol : blacklisted

// UNSUPPORTED : g_scanner_scope_remove_symbol : blacklisted

// UNSUPPORTED : g_scanner_set_scope : blacklisted

// UNSUPPORTED : g_scanner_sync_file_offset : blacklisted

// UNSUPPORTED : g_scanner_unexp_token : blacklisted

// UNSUPPORTED : g_scanner_warn : blacklisted

// UNSUPPORTED : g_scanner_new : blacklisted

// UNSUPPORTED : g_sequence_append : blacklisted

// UNSUPPORTED : g_sequence_foreach : blacklisted

// UNSUPPORTED : g_sequence_free : blacklisted

// UNSUPPORTED : g_sequence_get_begin_iter : blacklisted

// UNSUPPORTED : g_sequence_get_end_iter : blacklisted

// UNSUPPORTED : g_sequence_get_iter_at_pos : blacklisted

// UNSUPPORTED : g_sequence_get_length : blacklisted

// UNSUPPORTED : g_sequence_insert_sorted : blacklisted

// UNSUPPORTED : g_sequence_insert_sorted_iter : blacklisted

// UNSUPPORTED : g_sequence_is_empty : blacklisted

// UNSUPPORTED : g_sequence_lookup : blacklisted

// UNSUPPORTED : g_sequence_lookup_iter : blacklisted

// UNSUPPORTED : g_sequence_prepend : blacklisted

// UNSUPPORTED : g_sequence_search : blacklisted

// UNSUPPORTED : g_sequence_search_iter : blacklisted

// UNSUPPORTED : g_sequence_sort : blacklisted

// UNSUPPORTED : g_sequence_sort_iter : blacklisted

// UNSUPPORTED : g_sequence_foreach_range : blacklisted

// UNSUPPORTED : g_sequence_get : blacklisted

// UNSUPPORTED : g_sequence_insert_before : blacklisted

// UNSUPPORTED : g_sequence_move : blacklisted

// UNSUPPORTED : g_sequence_move_range : blacklisted

// UNSUPPORTED : g_sequence_new : blacklisted

// UNSUPPORTED : g_sequence_range_get_midpoint : blacklisted

// UNSUPPORTED : g_sequence_remove : blacklisted

// UNSUPPORTED : g_sequence_remove_range : blacklisted

// UNSUPPORTED : g_sequence_set : blacklisted

// UNSUPPORTED : g_sequence_sort_changed : blacklisted

// UNSUPPORTED : g_sequence_sort_changed_iter : blacklisted

// UNSUPPORTED : g_sequence_swap : blacklisted

// UNSUPPORTED : g_sequence_iter_compare : blacklisted

// UNSUPPORTED : g_sequence_iter_get_position : blacklisted

// UNSUPPORTED : g_sequence_iter_get_sequence : blacklisted

// UNSUPPORTED : g_sequence_iter_is_begin : blacklisted

// UNSUPPORTED : g_sequence_iter_is_end : blacklisted

// UNSUPPORTED : g_sequence_iter_move : blacklisted

// UNSUPPORTED : g_sequence_iter_next : blacklisted

// UNSUPPORTED : g_sequence_iter_prev : blacklisted

// UNSUPPORTED : g_source_new : blacklisted

// UNSUPPORTED : g_source_add_child_source : blacklisted

// UNSUPPORTED : g_source_add_poll : blacklisted

// UNSUPPORTED : g_source_add_unix_fd : blacklisted

// UNSUPPORTED : g_source_attach : blacklisted

// UNSUPPORTED : g_source_destroy : blacklisted

// UNSUPPORTED : g_source_get_can_recurse : blacklisted

// UNSUPPORTED : g_source_get_context : blacklisted

// UNSUPPORTED : g_source_get_current_time : blacklisted

// UNSUPPORTED : g_source_get_id : blacklisted

// UNSUPPORTED : g_source_get_name : blacklisted

// UNSUPPORTED : g_source_get_priority : blacklisted

// UNSUPPORTED : g_source_get_ready_time : blacklisted

// UNSUPPORTED : g_source_get_time : blacklisted

// UNSUPPORTED : g_source_is_destroyed : blacklisted

// UNSUPPORTED : g_source_modify_unix_fd : blacklisted

// UNSUPPORTED : g_source_query_unix_fd : blacklisted

// UNSUPPORTED : g_source_ref : blacklisted

// UNSUPPORTED : g_source_remove_child_source : blacklisted

// UNSUPPORTED : g_source_remove_poll : blacklisted

// UNSUPPORTED : g_source_remove_unix_fd : blacklisted

// UNSUPPORTED : g_source_set_callback : blacklisted

// UNSUPPORTED : g_source_set_callback_indirect : blacklisted

// UNSUPPORTED : g_source_set_can_recurse : blacklisted

// UNSUPPORTED : g_source_set_funcs : blacklisted

// UNSUPPORTED : g_source_set_name : blacklisted

// UNSUPPORTED : g_source_set_priority : blacklisted

// UNSUPPORTED : g_source_set_ready_time : blacklisted

// UNSUPPORTED : g_source_unref : blacklisted

// UNSUPPORTED : g_source_remove : blacklisted

// UNSUPPORTED : g_source_remove_by_funcs_user_data : blacklisted

// UNSUPPORTED : g_source_remove_by_user_data : blacklisted

// UNSUPPORTED : g_source_set_name_by_id : blacklisted

// UNSUPPORTED : g_string_append : blacklisted

// UNSUPPORTED : g_string_append_c : blacklisted

// UNSUPPORTED : g_string_append_len : blacklisted

// UNSUPPORTED : g_string_append_printf : blacklisted

// UNSUPPORTED : g_string_append_unichar : blacklisted

// UNSUPPORTED : g_string_append_uri_escaped : blacklisted

// UNSUPPORTED : g_string_append_vprintf : blacklisted

// UNSUPPORTED : g_string_ascii_down : blacklisted

// UNSUPPORTED : g_string_ascii_up : blacklisted

// UNSUPPORTED : g_string_assign : blacklisted

// UNSUPPORTED : g_string_down : blacklisted

// UNSUPPORTED : g_string_equal : blacklisted

// UNSUPPORTED : g_string_erase : blacklisted

// UNSUPPORTED : g_string_free : blacklisted

// UNSUPPORTED : g_string_free_to_bytes : blacklisted

// UNSUPPORTED : g_string_hash : blacklisted

// UNSUPPORTED : g_string_insert : blacklisted

// UNSUPPORTED : g_string_insert_c : blacklisted

// UNSUPPORTED : g_string_insert_len : blacklisted

// UNSUPPORTED : g_string_insert_unichar : blacklisted

// UNSUPPORTED : g_string_overwrite : blacklisted

// UNSUPPORTED : g_string_overwrite_len : blacklisted

// UNSUPPORTED : g_string_prepend : blacklisted

// UNSUPPORTED : g_string_prepend_c : blacklisted

// UNSUPPORTED : g_string_prepend_len : blacklisted

// UNSUPPORTED : g_string_prepend_unichar : blacklisted

// UNSUPPORTED : g_string_printf : blacklisted

// UNSUPPORTED : g_string_set_size : blacklisted

// UNSUPPORTED : g_string_truncate : blacklisted

// UNSUPPORTED : g_string_up : blacklisted

// UNSUPPORTED : g_string_vprintf : blacklisted

// UNSUPPORTED : g_string_chunk_clear : blacklisted

// UNSUPPORTED : g_string_chunk_free : blacklisted

// UNSUPPORTED : g_string_chunk_insert : blacklisted

// UNSUPPORTED : g_string_chunk_insert_const : blacklisted

// UNSUPPORTED : g_string_chunk_insert_len : blacklisted

// UNSUPPORTED : g_string_chunk_new : blacklisted

// UNSUPPORTED : g_test_log_buffer_free : blacklisted

// UNSUPPORTED : g_test_log_buffer_pop : blacklisted

// UNSUPPORTED : g_test_log_buffer_push : blacklisted

// UNSUPPORTED : g_test_log_buffer_new : blacklisted

// UNSUPPORTED : g_test_log_msg_free : blacklisted

// UNSUPPORTED : g_test_suite_add : blacklisted

// UNSUPPORTED : g_test_suite_add_suite : blacklisted

// UNSUPPORTED : g_thread_new : blacklisted

// UNSUPPORTED : g_thread_try_new : blacklisted

// UNSUPPORTED : g_thread_join : blacklisted

// UNSUPPORTED : g_thread_ref : blacklisted

// UNSUPPORTED : g_thread_unref : blacklisted

// UNSUPPORTED : g_thread_error_quark : blacklisted

// UNSUPPORTED : g_thread_exit : blacklisted

// UNSUPPORTED : g_thread_self : blacklisted

// UNSUPPORTED : g_thread_yield : blacklisted

// UNSUPPORTED : g_thread_pool_free : blacklisted

// UNSUPPORTED : g_thread_pool_get_max_threads : blacklisted

// UNSUPPORTED : g_thread_pool_get_num_threads : blacklisted

// UNSUPPORTED : g_thread_pool_move_to_front : blacklisted

// UNSUPPORTED : g_thread_pool_push : blacklisted

// UNSUPPORTED : g_thread_pool_set_max_threads : blacklisted

// UNSUPPORTED : g_thread_pool_set_sort_function : blacklisted

// UNSUPPORTED : g_thread_pool_unprocessed : blacklisted

// UNSUPPORTED : g_thread_pool_get_max_idle_time : blacklisted

// UNSUPPORTED : g_thread_pool_get_max_unused_threads : blacklisted

// UNSUPPORTED : g_thread_pool_get_num_unused_threads : blacklisted

// UNSUPPORTED : g_thread_pool_new : blacklisted

// UNSUPPORTED : g_thread_pool_set_max_idle_time : blacklisted

// UNSUPPORTED : g_thread_pool_set_max_unused_threads : blacklisted

// UNSUPPORTED : g_thread_pool_stop_unused_threads : blacklisted

// UNSUPPORTED : g_time_val_add : blacklisted

// UNSUPPORTED : g_time_val_to_iso8601 : blacklisted

// UNSUPPORTED : g_time_val_from_iso8601 : blacklisted

// UNSUPPORTED : g_time_zone_new : blacklisted

// UNSUPPORTED : g_time_zone_new_local : blacklisted

// UNSUPPORTED : g_time_zone_new_offset : blacklisted

// UNSUPPORTED : g_time_zone_new_utc : blacklisted

// UNSUPPORTED : g_time_zone_adjust_time : blacklisted

// UNSUPPORTED : g_time_zone_find_interval : blacklisted

// UNSUPPORTED : g_time_zone_get_abbreviation : blacklisted

// UNSUPPORTED : g_time_zone_get_identifier : blacklisted

// UNSUPPORTED : g_time_zone_get_offset : blacklisted

// UNSUPPORTED : g_time_zone_is_dst : blacklisted

// UNSUPPORTED : g_time_zone_ref : blacklisted

// UNSUPPORTED : g_time_zone_unref : blacklisted

// UNSUPPORTED : g_timer_continue : blacklisted

// UNSUPPORTED : g_timer_destroy : blacklisted

// UNSUPPORTED : g_timer_elapsed : blacklisted

// UNSUPPORTED : g_timer_is_active : blacklisted

// UNSUPPORTED : g_timer_reset : blacklisted

// UNSUPPORTED : g_timer_start : blacklisted

// UNSUPPORTED : g_timer_stop : blacklisted

// UNSUPPORTED : g_timer_new : blacklisted

// UNSUPPORTED : g_trash_stack_height : blacklisted

// UNSUPPORTED : g_trash_stack_peek : blacklisted

// UNSUPPORTED : g_trash_stack_pop : blacklisted

// UNSUPPORTED : g_trash_stack_push : blacklisted

// UNSUPPORTED : g_tree_destroy : blacklisted

// UNSUPPORTED : g_tree_foreach : blacklisted

// UNSUPPORTED : g_tree_height : blacklisted

// UNSUPPORTED : g_tree_insert : blacklisted

// UNSUPPORTED : g_tree_lookup : blacklisted

// UNSUPPORTED : g_tree_lookup_extended : blacklisted

// UNSUPPORTED : g_tree_nnodes : blacklisted

// UNSUPPORTED : g_tree_ref : blacklisted

// UNSUPPORTED : g_tree_remove : blacklisted

// UNSUPPORTED : g_tree_replace : blacklisted

// UNSUPPORTED : g_tree_search : blacklisted

// UNSUPPORTED : g_tree_steal : blacklisted

// UNSUPPORTED : g_tree_traverse : blacklisted

// UNSUPPORTED : g_tree_unref : blacklisted

// UNSUPPORTED : g_tree_new : blacklisted

// UNSUPPORTED : g_tree_new_full : blacklisted

// UNSUPPORTED : g_tree_new_with_data : blacklisted

// UNSUPPORTED : g_variant_new : blacklisted

// UNSUPPORTED : g_variant_new_array : blacklisted

// UNSUPPORTED : g_variant_new_boolean : blacklisted

// UNSUPPORTED : g_variant_new_byte : blacklisted

// UNSUPPORTED : g_variant_new_bytestring : blacklisted

// UNSUPPORTED : g_variant_new_bytestring_array : blacklisted

// UNSUPPORTED : g_variant_new_dict_entry : blacklisted

// UNSUPPORTED : g_variant_new_double : blacklisted

// UNSUPPORTED : g_variant_new_fixed_array : blacklisted

// UNSUPPORTED : g_variant_new_from_bytes : blacklisted

// UNSUPPORTED : g_variant_new_from_data : blacklisted

// UNSUPPORTED : g_variant_new_handle : blacklisted

// UNSUPPORTED : g_variant_new_int16 : blacklisted

// UNSUPPORTED : g_variant_new_int32 : blacklisted

// UNSUPPORTED : g_variant_new_int64 : blacklisted

// UNSUPPORTED : g_variant_new_maybe : blacklisted

// UNSUPPORTED : g_variant_new_object_path : blacklisted

// UNSUPPORTED : g_variant_new_objv : blacklisted

// UNSUPPORTED : g_variant_new_parsed : blacklisted

// UNSUPPORTED : g_variant_new_parsed_va : blacklisted

// UNSUPPORTED : g_variant_new_printf : blacklisted

// UNSUPPORTED : g_variant_new_signature : blacklisted

// UNSUPPORTED : g_variant_new_string : blacklisted

// UNSUPPORTED : g_variant_new_strv : blacklisted

// UNSUPPORTED : g_variant_new_take_string : blacklisted

// UNSUPPORTED : g_variant_new_tuple : blacklisted

// UNSUPPORTED : g_variant_new_uint16 : blacklisted

// UNSUPPORTED : g_variant_new_uint32 : blacklisted

// UNSUPPORTED : g_variant_new_uint64 : blacklisted

// UNSUPPORTED : g_variant_new_va : blacklisted

// UNSUPPORTED : g_variant_new_variant : blacklisted

// UNSUPPORTED : g_variant_byteswap : blacklisted

// UNSUPPORTED : g_variant_check_format_string : blacklisted

// UNSUPPORTED : g_variant_classify : blacklisted

// UNSUPPORTED : g_variant_compare : blacklisted

// UNSUPPORTED : g_variant_dup_bytestring : blacklisted

// UNSUPPORTED : g_variant_dup_bytestring_array : blacklisted

// UNSUPPORTED : g_variant_dup_objv : blacklisted

// UNSUPPORTED : g_variant_dup_string : blacklisted

// UNSUPPORTED : g_variant_dup_strv : blacklisted

// UNSUPPORTED : g_variant_equal : blacklisted

// UNSUPPORTED : g_variant_get : blacklisted

// UNSUPPORTED : g_variant_get_boolean : blacklisted

// UNSUPPORTED : g_variant_get_byte : blacklisted

// UNSUPPORTED : g_variant_get_bytestring : blacklisted

// UNSUPPORTED : g_variant_get_bytestring_array : blacklisted

// UNSUPPORTED : g_variant_get_child : blacklisted

// UNSUPPORTED : g_variant_get_child_value : blacklisted

// UNSUPPORTED : g_variant_get_data : blacklisted

// UNSUPPORTED : g_variant_get_data_as_bytes : blacklisted

// UNSUPPORTED : g_variant_get_double : blacklisted

// UNSUPPORTED : g_variant_get_fixed_array : blacklisted

// UNSUPPORTED : g_variant_get_handle : blacklisted

// UNSUPPORTED : g_variant_get_int16 : blacklisted

// UNSUPPORTED : g_variant_get_int32 : blacklisted

// UNSUPPORTED : g_variant_get_int64 : blacklisted

// UNSUPPORTED : g_variant_get_maybe : blacklisted

// UNSUPPORTED : g_variant_get_normal_form : blacklisted

// UNSUPPORTED : g_variant_get_objv : blacklisted

// UNSUPPORTED : g_variant_get_size : blacklisted

// UNSUPPORTED : g_variant_get_string : blacklisted

// UNSUPPORTED : g_variant_get_strv : blacklisted

// UNSUPPORTED : g_variant_get_type : blacklisted

// UNSUPPORTED : g_variant_get_type_string : blacklisted

// UNSUPPORTED : g_variant_get_uint16 : blacklisted

// UNSUPPORTED : g_variant_get_uint32 : blacklisted

// UNSUPPORTED : g_variant_get_uint64 : blacklisted

// UNSUPPORTED : g_variant_get_va : blacklisted

// UNSUPPORTED : g_variant_get_variant : blacklisted

// UNSUPPORTED : g_variant_hash : blacklisted

// UNSUPPORTED : g_variant_is_container : blacklisted

// UNSUPPORTED : g_variant_is_floating : blacklisted

// UNSUPPORTED : g_variant_is_normal_form : blacklisted

// UNSUPPORTED : g_variant_is_of_type : blacklisted

// UNSUPPORTED : g_variant_iter_new : blacklisted

// UNSUPPORTED : g_variant_lookup : blacklisted

// UNSUPPORTED : g_variant_lookup_value : blacklisted

// UNSUPPORTED : g_variant_n_children : blacklisted

// UNSUPPORTED : g_variant_print : blacklisted

// UNSUPPORTED : g_variant_print_string : blacklisted

// UNSUPPORTED : g_variant_ref : blacklisted

// UNSUPPORTED : g_variant_ref_sink : blacklisted

// UNSUPPORTED : g_variant_store : blacklisted

// UNSUPPORTED : g_variant_take_ref : blacklisted

// UNSUPPORTED : g_variant_unref : blacklisted

// UNSUPPORTED : g_variant_is_object_path : blacklisted

// UNSUPPORTED : g_variant_is_signature : blacklisted

// UNSUPPORTED : g_variant_parse : blacklisted

// UNSUPPORTED : g_variant_parse_error_print_context : blacklisted

// UNSUPPORTED : g_variant_parse_error_quark : blacklisted

// UNSUPPORTED : g_variant_parser_get_error_quark : blacklisted

// UNSUPPORTED : g_variant_builder_new : blacklisted

// UNSUPPORTED : g_variant_builder_add : blacklisted

// UNSUPPORTED : g_variant_builder_add_parsed : blacklisted

// UNSUPPORTED : g_variant_builder_add_value : blacklisted

// UNSUPPORTED : g_variant_builder_clear : blacklisted

// UNSUPPORTED : g_variant_builder_close : blacklisted

// UNSUPPORTED : g_variant_builder_end : blacklisted

// UNSUPPORTED : g_variant_builder_init : blacklisted

// UNSUPPORTED : g_variant_builder_open : blacklisted

// UNSUPPORTED : g_variant_builder_ref : blacklisted

// UNSUPPORTED : g_variant_builder_unref : blacklisted

// UNSUPPORTED : g_variant_dict_new : blacklisted

// UNSUPPORTED : g_variant_dict_clear : blacklisted

// UNSUPPORTED : g_variant_dict_contains : blacklisted

// UNSUPPORTED : g_variant_dict_end : blacklisted

// UNSUPPORTED : g_variant_dict_init : blacklisted

// UNSUPPORTED : g_variant_dict_insert : blacklisted

// UNSUPPORTED : g_variant_dict_insert_value : blacklisted

// UNSUPPORTED : g_variant_dict_lookup : blacklisted

// UNSUPPORTED : g_variant_dict_lookup_value : blacklisted

// UNSUPPORTED : g_variant_dict_ref : blacklisted

// UNSUPPORTED : g_variant_dict_remove : blacklisted

// UNSUPPORTED : g_variant_dict_unref : blacklisted

// UNSUPPORTED : g_variant_iter_copy : blacklisted

// UNSUPPORTED : g_variant_iter_free : blacklisted

// UNSUPPORTED : g_variant_iter_init : blacklisted

// UNSUPPORTED : g_variant_iter_loop : blacklisted

// UNSUPPORTED : g_variant_iter_n_children : blacklisted

// UNSUPPORTED : g_variant_iter_next : blacklisted

// UNSUPPORTED : g_variant_iter_next_value : blacklisted

// UNSUPPORTED : g_variant_type_new : blacklisted

// UNSUPPORTED : g_variant_type_new_array : blacklisted

// UNSUPPORTED : g_variant_type_new_dict_entry : blacklisted

// UNSUPPORTED : g_variant_type_new_maybe : blacklisted

// UNSUPPORTED : g_variant_type_new_tuple : blacklisted

// UNSUPPORTED : g_variant_type_copy : blacklisted

// UNSUPPORTED : g_variant_type_dup_string : blacklisted

// UNSUPPORTED : g_variant_type_element : blacklisted

// UNSUPPORTED : g_variant_type_equal : blacklisted

// UNSUPPORTED : g_variant_type_first : blacklisted

// UNSUPPORTED : g_variant_type_free : blacklisted

// UNSUPPORTED : g_variant_type_get_string_length : blacklisted

// UNSUPPORTED : g_variant_type_hash : blacklisted

// UNSUPPORTED : g_variant_type_is_array : blacklisted

// UNSUPPORTED : g_variant_type_is_basic : blacklisted

// UNSUPPORTED : g_variant_type_is_container : blacklisted

// UNSUPPORTED : g_variant_type_is_definite : blacklisted

// UNSUPPORTED : g_variant_type_is_dict_entry : blacklisted

// UNSUPPORTED : g_variant_type_is_maybe : blacklisted

// UNSUPPORTED : g_variant_type_is_subtype_of : blacklisted

// UNSUPPORTED : g_variant_type_is_tuple : blacklisted

// UNSUPPORTED : g_variant_type_is_variant : blacklisted

// UNSUPPORTED : g_variant_type_key : blacklisted

// UNSUPPORTED : g_variant_type_n_items : blacklisted

// UNSUPPORTED : g_variant_type_next : blacklisted

// UNSUPPORTED : g_variant_type_peek_string : blacklisted

// UNSUPPORTED : g_variant_type_value : blacklisted

// UNSUPPORTED : g_variant_type_checked_ : blacklisted

// UNSUPPORTED : g_variant_type_string_get_depth_ : blacklisted

// UNSUPPORTED : g_variant_type_string_is_valid : blacklisted

// UNSUPPORTED : g_variant_type_string_scan : blacklisted

// UNSUPPORTED : g_access : blacklisted

// UNSUPPORTED : g_ascii_digit_value : blacklisted

// UNSUPPORTED : g_ascii_dtostr : blacklisted

// UNSUPPORTED : g_ascii_formatd : blacklisted

// UNSUPPORTED : g_ascii_strcasecmp : blacklisted

// UNSUPPORTED : g_ascii_strdown : blacklisted

// UNSUPPORTED : g_ascii_string_to_signed : blacklisted

// UNSUPPORTED : g_ascii_string_to_unsigned : blacklisted

// UNSUPPORTED : g_ascii_strncasecmp : blacklisted

// UNSUPPORTED : g_ascii_strtod : blacklisted

// UNSUPPORTED : g_ascii_strtoll : blacklisted

// UNSUPPORTED : g_ascii_strtoull : blacklisted

// UNSUPPORTED : g_ascii_strup : blacklisted

// UNSUPPORTED : g_ascii_tolower : blacklisted

// UNSUPPORTED : g_ascii_toupper : blacklisted

// UNSUPPORTED : g_ascii_xdigit_value : blacklisted

// UNSUPPORTED : g_assert_warning : blacklisted

// UNSUPPORTED : g_assertion_message : blacklisted

// UNSUPPORTED : g_assertion_message_cmpnum : blacklisted

// UNSUPPORTED : g_assertion_message_cmpstr : blacklisted

// UNSUPPORTED : g_assertion_message_error : blacklisted

// UNSUPPORTED : g_assertion_message_expr : blacklisted

// UNSUPPORTED : g_atexit : blacklisted

// UNSUPPORTED : g_atomic_int_add : blacklisted

// UNSUPPORTED : g_atomic_int_and : blacklisted

// UNSUPPORTED : g_atomic_int_compare_and_exchange : blacklisted

// UNSUPPORTED : g_atomic_int_dec_and_test : blacklisted

// UNSUPPORTED : g_atomic_int_exchange_and_add : blacklisted

// UNSUPPORTED : g_atomic_int_get : blacklisted

// UNSUPPORTED : g_atomic_int_inc : blacklisted

// UNSUPPORTED : g_atomic_int_or : blacklisted

// UNSUPPORTED : g_atomic_int_set : blacklisted

// UNSUPPORTED : g_atomic_int_xor : blacklisted

// UNSUPPORTED : g_atomic_pointer_add : blacklisted

// UNSUPPORTED : g_atomic_pointer_and : blacklisted

// UNSUPPORTED : g_atomic_pointer_compare_and_exchange : blacklisted

// UNSUPPORTED : g_atomic_pointer_get : blacklisted

// UNSUPPORTED : g_atomic_pointer_or : blacklisted

// UNSUPPORTED : g_atomic_pointer_set : blacklisted

// UNSUPPORTED : g_atomic_pointer_xor : blacklisted

// UNSUPPORTED : g_atomic_rc_box_acquire : blacklisted

// UNSUPPORTED : g_atomic_rc_box_alloc : blacklisted

// UNSUPPORTED : g_atomic_rc_box_alloc0 : blacklisted

// UNSUPPORTED : g_atomic_rc_box_dup : blacklisted

// UNSUPPORTED : g_atomic_rc_box_get_size : blacklisted

// UNSUPPORTED : g_atomic_rc_box_release : blacklisted

// UNSUPPORTED : g_atomic_rc_box_release_full : blacklisted

// UNSUPPORTED : g_atomic_ref_count_compare : blacklisted

// UNSUPPORTED : g_atomic_ref_count_dec : blacklisted

// UNSUPPORTED : g_atomic_ref_count_inc : blacklisted

// UNSUPPORTED : g_atomic_ref_count_init : blacklisted

// UNSUPPORTED : g_base64_decode : blacklisted

// UNSUPPORTED : g_base64_decode_inplace : blacklisted

// UNSUPPORTED : g_base64_decode_step : blacklisted

// UNSUPPORTED : g_base64_encode : blacklisted

// UNSUPPORTED : g_base64_encode_close : blacklisted

// UNSUPPORTED : g_base64_encode_step : blacklisted

// UNSUPPORTED : g_basename : blacklisted

// UNSUPPORTED : g_bit_lock : blacklisted

// UNSUPPORTED : g_bit_nth_lsf : blacklisted

// UNSUPPORTED : g_bit_nth_msf : blacklisted

// UNSUPPORTED : g_bit_storage : blacklisted

// UNSUPPORTED : g_bit_trylock : blacklisted

// UNSUPPORTED : g_bit_unlock : blacklisted

// UNSUPPORTED : g_bookmark_file_error_quark : blacklisted

// UNSUPPORTED : g_build_filename : blacklisted

// UNSUPPORTED : g_build_filename_valist : blacklisted

// UNSUPPORTED : g_build_filenamev : blacklisted

// UNSUPPORTED : g_build_path : blacklisted

// UNSUPPORTED : g_build_pathv : blacklisted

// UNSUPPORTED : g_byte_array_free : blacklisted

// UNSUPPORTED : g_byte_array_free_to_bytes : blacklisted

// UNSUPPORTED : g_byte_array_new : blacklisted

// UNSUPPORTED : g_byte_array_new_take : blacklisted

// UNSUPPORTED : g_byte_array_unref : blacklisted

// UNSUPPORTED : g_canonicalize_filename : blacklisted

// UNSUPPORTED : g_chdir : blacklisted

// UNSUPPORTED : glib_check_version : blacklisted

// UNSUPPORTED : g_checksum_type_get_length : blacklisted

// UNSUPPORTED : g_child_watch_add : blacklisted

// UNSUPPORTED : g_child_watch_add_full : blacklisted

// UNSUPPORTED : g_child_watch_source_new : blacklisted

// UNSUPPORTED : g_clear_error : blacklisted

// UNSUPPORTED : g_clear_handle_id : blacklisted

// UNSUPPORTED : g_clear_pointer : blacklisted

// UNSUPPORTED : g_close : blacklisted

// UNSUPPORTED : g_compute_checksum_for_bytes : blacklisted

// UNSUPPORTED : g_compute_checksum_for_data : blacklisted

// UNSUPPORTED : g_compute_checksum_for_string : blacklisted

// UNSUPPORTED : g_compute_hmac_for_bytes : blacklisted

// UNSUPPORTED : g_compute_hmac_for_data : blacklisted

// UNSUPPORTED : g_compute_hmac_for_string : blacklisted

// UNSUPPORTED : g_convert : blacklisted

// UNSUPPORTED : g_convert_error_quark : blacklisted

// UNSUPPORTED : g_convert_with_fallback : blacklisted

// UNSUPPORTED : g_convert_with_iconv : blacklisted

// UNSUPPORTED : g_datalist_clear : blacklisted

// UNSUPPORTED : g_datalist_foreach : blacklisted

// UNSUPPORTED : g_datalist_get_data : blacklisted

// UNSUPPORTED : g_datalist_get_flags : blacklisted

// UNSUPPORTED : g_datalist_id_dup_data : blacklisted

// UNSUPPORTED : g_datalist_id_get_data : blacklisted

// UNSUPPORTED : g_datalist_id_remove_no_notify : blacklisted

// UNSUPPORTED : g_datalist_id_replace_data : blacklisted

// UNSUPPORTED : g_datalist_id_set_data_full : blacklisted

// UNSUPPORTED : g_datalist_init : blacklisted

// UNSUPPORTED : g_datalist_set_flags : blacklisted

// UNSUPPORTED : g_datalist_unset_flags : blacklisted

// UNSUPPORTED : g_dataset_destroy : blacklisted

// UNSUPPORTED : g_dataset_foreach : blacklisted

// UNSUPPORTED : g_dataset_id_get_data : blacklisted

// UNSUPPORTED : g_dataset_id_remove_no_notify : blacklisted

// UNSUPPORTED : g_dataset_id_set_data_full : blacklisted

// UNSUPPORTED : g_date_get_days_in_month : blacklisted

// UNSUPPORTED : g_date_get_monday_weeks_in_year : blacklisted

// UNSUPPORTED : g_date_get_sunday_weeks_in_year : blacklisted

// UNSUPPORTED : g_date_is_leap_year : blacklisted

// UNSUPPORTED : g_date_strftime : blacklisted

// UNSUPPORTED : g_date_time_compare : blacklisted

// UNSUPPORTED : g_date_time_equal : blacklisted

// UNSUPPORTED : g_date_time_hash : blacklisted

// UNSUPPORTED : g_date_valid_day : blacklisted

// UNSUPPORTED : g_date_valid_dmy : blacklisted

// UNSUPPORTED : g_date_valid_julian : blacklisted

// UNSUPPORTED : g_date_valid_month : blacklisted

// UNSUPPORTED : g_date_valid_weekday : blacklisted

// UNSUPPORTED : g_date_valid_year : blacklisted

// UNSUPPORTED : g_dcgettext : blacklisted

// UNSUPPORTED : g_dgettext : blacklisted

// UNSUPPORTED : g_dir_make_tmp : blacklisted

// UNSUPPORTED : g_direct_equal : blacklisted

// UNSUPPORTED : g_direct_hash : blacklisted

// UNSUPPORTED : g_dngettext : blacklisted

// UNSUPPORTED : g_double_equal : blacklisted

// UNSUPPORTED : g_double_hash : blacklisted

// UNSUPPORTED : g_dpgettext : blacklisted

// UNSUPPORTED : g_dpgettext2 : blacklisted

// UNSUPPORTED : g_environ_getenv : blacklisted

// UNSUPPORTED : g_environ_setenv : blacklisted

// UNSUPPORTED : g_environ_unsetenv : blacklisted

// UNSUPPORTED : g_file_error_from_errno : blacklisted

// UNSUPPORTED : g_file_error_quark : blacklisted

// UNSUPPORTED : g_file_get_contents : blacklisted

// UNSUPPORTED : g_file_open_tmp : blacklisted

// UNSUPPORTED : g_file_read_link : blacklisted

// UNSUPPORTED : g_file_set_contents : blacklisted

// UNSUPPORTED : g_file_test : blacklisted

// UNSUPPORTED : g_filename_display_basename : blacklisted

// UNSUPPORTED : g_filename_display_name : blacklisted

// UNSUPPORTED : g_filename_from_uri : blacklisted

// UNSUPPORTED : g_filename_from_utf8 : blacklisted

// UNSUPPORTED : g_filename_to_uri : blacklisted

// UNSUPPORTED : g_filename_to_utf8 : blacklisted

// UNSUPPORTED : g_find_program_in_path : blacklisted

// UNSUPPORTED : g_format_size : blacklisted

// UNSUPPORTED : g_format_size_for_display : blacklisted

// UNSUPPORTED : g_format_size_full : blacklisted

// UNSUPPORTED : g_fprintf : blacklisted

// UNSUPPORTED : g_free : blacklisted

// UNSUPPORTED : g_get_application_name : blacklisted

// UNSUPPORTED : g_get_charset : blacklisted

// UNSUPPORTED : g_get_codeset : blacklisted

// UNSUPPORTED : g_get_console_charset : blacklisted

// UNSUPPORTED : g_get_current_dir : blacklisted

// UNSUPPORTED : g_get_current_time : blacklisted

// UNSUPPORTED : g_get_environ : blacklisted

// UNSUPPORTED : g_get_filename_charsets : blacklisted

// UNSUPPORTED : g_get_home_dir : blacklisted

// UNSUPPORTED : g_get_host_name : blacklisted

// UNSUPPORTED : g_get_language_names : blacklisted

// UNSUPPORTED : g_get_language_names_with_category : blacklisted

// UNSUPPORTED : g_get_locale_variants : blacklisted

// UNSUPPORTED : g_get_monotonic_time : blacklisted

// UNSUPPORTED : g_get_num_processors : blacklisted

// UNSUPPORTED : g_get_prgname : blacklisted

// UNSUPPORTED : g_get_real_name : blacklisted

// UNSUPPORTED : g_get_real_time : blacklisted

// UNSUPPORTED : g_get_system_config_dirs : blacklisted

// UNSUPPORTED : g_get_system_data_dirs : blacklisted

// UNSUPPORTED : g_get_tmp_dir : blacklisted

// UNSUPPORTED : g_get_user_cache_dir : blacklisted

// UNSUPPORTED : g_get_user_config_dir : blacklisted

// UNSUPPORTED : g_get_user_data_dir : blacklisted

// UNSUPPORTED : g_get_user_name : blacklisted

// UNSUPPORTED : g_get_user_runtime_dir : blacklisted

// UNSUPPORTED : g_get_user_special_dir : blacklisted

// UNSUPPORTED : g_getenv : blacklisted

// UNSUPPORTED : g_hash_table_add : blacklisted

// UNSUPPORTED : g_hash_table_contains : blacklisted

// UNSUPPORTED : g_hash_table_destroy : blacklisted

// UNSUPPORTED : g_hash_table_insert : blacklisted

// UNSUPPORTED : g_hash_table_lookup : blacklisted

// UNSUPPORTED : g_hash_table_lookup_extended : blacklisted

// UNSUPPORTED : g_hash_table_remove : blacklisted

// UNSUPPORTED : g_hash_table_remove_all : blacklisted

// UNSUPPORTED : g_hash_table_replace : blacklisted

// UNSUPPORTED : g_hash_table_size : blacklisted

// UNSUPPORTED : g_hash_table_steal : blacklisted

// UNSUPPORTED : g_hash_table_steal_all : blacklisted

// UNSUPPORTED : g_hash_table_steal_extended : blacklisted

// UNSUPPORTED : g_hash_table_unref : blacklisted

// UNSUPPORTED : g_hook_destroy : blacklisted

// UNSUPPORTED : g_hook_destroy_link : blacklisted

// UNSUPPORTED : g_hook_free : blacklisted

// UNSUPPORTED : g_hook_insert_before : blacklisted

// UNSUPPORTED : g_hook_prepend : blacklisted

// UNSUPPORTED : g_hook_unref : blacklisted

// UNSUPPORTED : g_hostname_is_ascii_encoded : blacklisted

// UNSUPPORTED : g_hostname_is_ip_address : blacklisted

// UNSUPPORTED : g_hostname_is_non_ascii : blacklisted

// UNSUPPORTED : g_hostname_to_ascii : blacklisted

// UNSUPPORTED : g_hostname_to_unicode : blacklisted

// UNSUPPORTED : g_iconv : blacklisted

// UNSUPPORTED : g_iconv_open : blacklisted

// UNSUPPORTED : g_idle_add : blacklisted

// UNSUPPORTED : g_idle_add_full : blacklisted

// UNSUPPORTED : g_idle_remove_by_data : blacklisted

// UNSUPPORTED : g_idle_source_new : blacklisted

// UNSUPPORTED : g_int64_equal : blacklisted

// UNSUPPORTED : g_int64_hash : blacklisted

// UNSUPPORTED : g_int_equal : blacklisted

// UNSUPPORTED : g_int_hash : blacklisted

// UNSUPPORTED : g_intern_static_string : blacklisted

// UNSUPPORTED : g_intern_string : blacklisted

// UNSUPPORTED : g_io_add_watch : blacklisted

// UNSUPPORTED : g_io_add_watch_full : blacklisted

// UNSUPPORTED : g_io_channel_error_from_errno : blacklisted

// UNSUPPORTED : g_io_channel_error_quark : blacklisted

// UNSUPPORTED : g_io_create_watch : blacklisted

// UNSUPPORTED : g_key_file_error_quark : blacklisted

// UNSUPPORTED : g_listenv : blacklisted

// UNSUPPORTED : g_locale_from_utf8 : blacklisted

// UNSUPPORTED : g_locale_to_utf8 : blacklisted

// UNSUPPORTED : g_log : blacklisted

// UNSUPPORTED : g_log_default_handler : blacklisted

// UNSUPPORTED : g_log_remove_handler : blacklisted

// UNSUPPORTED : g_log_set_always_fatal : blacklisted

// UNSUPPORTED : g_log_set_default_handler : blacklisted

// UNSUPPORTED : g_log_set_fatal_mask : blacklisted

// UNSUPPORTED : g_log_set_handler : blacklisted

// UNSUPPORTED : g_log_set_handler_full : blacklisted

// UNSUPPORTED : g_log_set_writer_func : blacklisted

// UNSUPPORTED : g_log_structured : blacklisted

// UNSUPPORTED : g_log_structured_array : blacklisted

// UNSUPPORTED : g_log_structured_standard : blacklisted

// UNSUPPORTED : g_log_variant : blacklisted

// UNSUPPORTED : g_log_writer_default : blacklisted

// UNSUPPORTED : g_log_writer_format_fields : blacklisted

// UNSUPPORTED : g_log_writer_is_journald : blacklisted

// UNSUPPORTED : g_log_writer_journald : blacklisted

// UNSUPPORTED : g_log_writer_standard_streams : blacklisted

// UNSUPPORTED : g_log_writer_supports_color : blacklisted

// UNSUPPORTED : g_logv : blacklisted

// UNSUPPORTED : g_main_context_default : blacklisted

// UNSUPPORTED : g_main_context_get_thread_default : blacklisted

// UNSUPPORTED : g_main_context_ref_thread_default : blacklisted

// UNSUPPORTED : g_main_current_source : blacklisted

// UNSUPPORTED : g_main_depth : blacklisted

// UNSUPPORTED : g_malloc : blacklisted

// UNSUPPORTED : g_malloc0 : blacklisted

// UNSUPPORTED : g_malloc0_n : blacklisted

// UNSUPPORTED : g_malloc_n : blacklisted

// UNSUPPORTED : g_markup_collect_attributes : blacklisted

// UNSUPPORTED : g_markup_error_quark : blacklisted

// UNSUPPORTED : g_markup_escape_text : blacklisted

// UNSUPPORTED : g_markup_printf_escaped : blacklisted

// UNSUPPORTED : g_markup_vprintf_escaped : blacklisted

// UNSUPPORTED : g_mem_is_system_malloc : blacklisted

// UNSUPPORTED : g_mem_profile : blacklisted

// UNSUPPORTED : g_mem_set_vtable : blacklisted

// UNSUPPORTED : g_memdup : blacklisted

// UNSUPPORTED : g_mkdir_with_parents : blacklisted

// UNSUPPORTED : g_mkdtemp : blacklisted

// UNSUPPORTED : g_mkdtemp_full : blacklisted

// UNSUPPORTED : g_mkstemp : blacklisted

// UNSUPPORTED : g_mkstemp_full : blacklisted

// UNSUPPORTED : g_nullify_pointer : blacklisted

// UNSUPPORTED : g_number_parser_error_quark : blacklisted

// UNSUPPORTED : g_on_error_query : blacklisted

// UNSUPPORTED : g_on_error_stack_trace : blacklisted

// UNSUPPORTED : g_once_init_enter : blacklisted

// UNSUPPORTED : g_once_init_leave : blacklisted

// UNSUPPORTED : g_option_error_quark : blacklisted

// UNSUPPORTED : g_parse_debug_string : blacklisted

// UNSUPPORTED : g_path_get_basename : blacklisted

// UNSUPPORTED : g_path_get_dirname : blacklisted

// UNSUPPORTED : g_path_is_absolute : blacklisted

// UNSUPPORTED : g_path_skip_root : blacklisted

// UNSUPPORTED : g_pattern_match : blacklisted

// UNSUPPORTED : g_pattern_match_simple : blacklisted

// UNSUPPORTED : g_pattern_match_string : blacklisted

// UNSUPPORTED : g_pointer_bit_lock : blacklisted

// UNSUPPORTED : g_pointer_bit_trylock : blacklisted

// UNSUPPORTED : g_pointer_bit_unlock : blacklisted

// UNSUPPORTED : g_poll : blacklisted

// UNSUPPORTED : g_prefix_error : blacklisted

// UNSUPPORTED : g_print : blacklisted

// UNSUPPORTED : g_printerr : blacklisted

// UNSUPPORTED : g_printf : blacklisted

// UNSUPPORTED : g_printf_string_upper_bound : blacklisted

// UNSUPPORTED : g_propagate_error : blacklisted

// UNSUPPORTED : g_propagate_prefixed_error : blacklisted

// UNSUPPORTED : g_ptr_array_find : blacklisted

// UNSUPPORTED : g_ptr_array_find_with_equal_func : blacklisted

// UNSUPPORTED : g_qsort_with_data : blacklisted

// UNSUPPORTED : g_quark_from_static_string : blacklisted

// UNSUPPORTED : g_quark_from_string : blacklisted

// UNSUPPORTED : g_quark_to_string : blacklisted

// UNSUPPORTED : g_quark_try_string : blacklisted

// UNSUPPORTED : g_random_double : blacklisted

// UNSUPPORTED : g_random_double_range : blacklisted

// UNSUPPORTED : g_random_int : blacklisted

// UNSUPPORTED : g_random_int_range : blacklisted

// UNSUPPORTED : g_random_set_seed : blacklisted

// UNSUPPORTED : g_rc_box_acquire : blacklisted

// UNSUPPORTED : g_rc_box_alloc : blacklisted

// UNSUPPORTED : g_rc_box_alloc0 : blacklisted

// UNSUPPORTED : g_rc_box_dup : blacklisted

// UNSUPPORTED : g_rc_box_get_size : blacklisted

// UNSUPPORTED : g_rc_box_release : blacklisted

// UNSUPPORTED : g_rc_box_release_full : blacklisted

// UNSUPPORTED : g_realloc : blacklisted

// UNSUPPORTED : g_realloc_n : blacklisted

// UNSUPPORTED : g_ref_count_compare : blacklisted

// UNSUPPORTED : g_ref_count_dec : blacklisted

// UNSUPPORTED : g_ref_count_inc : blacklisted

// UNSUPPORTED : g_ref_count_init : blacklisted

// UNSUPPORTED : g_ref_string_acquire : blacklisted

// UNSUPPORTED : g_ref_string_length : blacklisted

// UNSUPPORTED : g_ref_string_new : blacklisted

// UNSUPPORTED : g_ref_string_new_intern : blacklisted

// UNSUPPORTED : g_ref_string_new_len : blacklisted

// UNSUPPORTED : g_ref_string_release : blacklisted

// UNSUPPORTED : g_regex_check_replacement : blacklisted

// UNSUPPORTED : g_regex_error_quark : blacklisted

// UNSUPPORTED : g_regex_escape_nul : blacklisted

// UNSUPPORTED : g_regex_escape_string : blacklisted

// UNSUPPORTED : g_regex_match_simple : blacklisted

// UNSUPPORTED : g_regex_split_simple : blacklisted

// UNSUPPORTED : g_reload_user_special_dirs_cache : blacklisted

// UNSUPPORTED : g_return_if_fail_warning : blacklisted

// UNSUPPORTED : g_rmdir : blacklisted

// UNSUPPORTED : g_sequence_get : blacklisted

// UNSUPPORTED : g_sequence_insert_before : blacklisted

// UNSUPPORTED : g_sequence_move : blacklisted

// UNSUPPORTED : g_sequence_move_range : blacklisted

// UNSUPPORTED : g_sequence_range_get_midpoint : blacklisted

// UNSUPPORTED : g_sequence_remove : blacklisted

// UNSUPPORTED : g_sequence_remove_range : blacklisted

// UNSUPPORTED : g_sequence_set : blacklisted

// UNSUPPORTED : g_sequence_swap : blacklisted

// UNSUPPORTED : g_set_application_name : blacklisted

// UNSUPPORTED : g_set_error : blacklisted

// UNSUPPORTED : g_set_error_literal : blacklisted

// UNSUPPORTED : g_set_prgname : blacklisted

// UNSUPPORTED : g_set_print_handler : blacklisted

// UNSUPPORTED : g_set_printerr_handler : blacklisted

// UNSUPPORTED : g_setenv : blacklisted

// UNSUPPORTED : g_shell_error_quark : blacklisted

// UNSUPPORTED : g_shell_parse_argv : blacklisted

// UNSUPPORTED : g_shell_quote : blacklisted

// UNSUPPORTED : g_shell_unquote : blacklisted

// UNSUPPORTED : g_slice_alloc : blacklisted

// UNSUPPORTED : g_slice_alloc0 : blacklisted

// UNSUPPORTED : g_slice_copy : blacklisted

// UNSUPPORTED : g_slice_free1 : blacklisted

// UNSUPPORTED : g_slice_free_chain_with_offset : blacklisted

// UNSUPPORTED : g_slice_get_config : blacklisted

// UNSUPPORTED : g_slice_get_config_state : blacklisted

// UNSUPPORTED : g_slice_set_config : blacklisted

// UNSUPPORTED : g_snprintf : blacklisted

// UNSUPPORTED : g_source_remove : blacklisted

// UNSUPPORTED : g_source_remove_by_funcs_user_data : blacklisted

// UNSUPPORTED : g_source_remove_by_user_data : blacklisted

// UNSUPPORTED : g_source_set_name_by_id : blacklisted

// UNSUPPORTED : g_spaced_primes_closest : blacklisted

// UNSUPPORTED : g_spawn_async : blacklisted

// UNSUPPORTED : g_spawn_async_with_fds : blacklisted

// UNSUPPORTED : g_spawn_async_with_pipes : blacklisted

// UNSUPPORTED : g_spawn_check_exit_status : blacklisted

// UNSUPPORTED : g_spawn_close_pid : blacklisted

// UNSUPPORTED : g_spawn_command_line_async : blacklisted

// UNSUPPORTED : g_spawn_command_line_sync : blacklisted

// UNSUPPORTED : g_spawn_error_quark : blacklisted

// UNSUPPORTED : g_spawn_exit_error_quark : blacklisted

// UNSUPPORTED : g_spawn_sync : blacklisted

// UNSUPPORTED : g_sprintf : blacklisted

// UNSUPPORTED : g_stpcpy : blacklisted

// UNSUPPORTED : g_str_equal : blacklisted

// UNSUPPORTED : g_str_has_prefix : blacklisted

// UNSUPPORTED : g_str_has_suffix : blacklisted

// UNSUPPORTED : g_str_hash : blacklisted

// UNSUPPORTED : g_str_is_ascii : blacklisted

// UNSUPPORTED : g_str_match_string : blacklisted

// UNSUPPORTED : g_str_to_ascii : blacklisted

// UNSUPPORTED : g_str_tokenize_and_fold : blacklisted

// UNSUPPORTED : g_strcanon : blacklisted

// UNSUPPORTED : g_strcasecmp : blacklisted

// UNSUPPORTED : g_strchomp : blacklisted

// UNSUPPORTED : g_strchug : blacklisted

// UNSUPPORTED : g_strcmp0 : blacklisted

// UNSUPPORTED : g_strcompress : blacklisted

// UNSUPPORTED : g_strconcat : blacklisted

// UNSUPPORTED : g_strdelimit : blacklisted

// UNSUPPORTED : g_strdown : blacklisted

// UNSUPPORTED : g_strdup : blacklisted

// UNSUPPORTED : g_strdup_printf : blacklisted

// UNSUPPORTED : g_strdup_vprintf : blacklisted

// UNSUPPORTED : g_strdupv : blacklisted

// UNSUPPORTED : g_strerror : blacklisted

// UNSUPPORTED : g_strescape : blacklisted

// UNSUPPORTED : g_strfreev : blacklisted

// UNSUPPORTED : g_string_new : blacklisted

// UNSUPPORTED : g_string_new_len : blacklisted

// UNSUPPORTED : g_string_sized_new : blacklisted

// UNSUPPORTED : g_strip_context : blacklisted

// UNSUPPORTED : g_strjoin : blacklisted

// UNSUPPORTED : g_strjoinv : blacklisted

// UNSUPPORTED : g_strlcat : blacklisted

// UNSUPPORTED : g_strlcpy : blacklisted

// UNSUPPORTED : g_strncasecmp : blacklisted

// UNSUPPORTED : g_strndup : blacklisted

// UNSUPPORTED : g_strnfill : blacklisted

// UNSUPPORTED : g_strreverse : blacklisted

// UNSUPPORTED : g_strrstr : blacklisted

// UNSUPPORTED : g_strrstr_len : blacklisted

// UNSUPPORTED : g_strsignal : blacklisted

// UNSUPPORTED : g_strsplit : blacklisted

// UNSUPPORTED : g_strsplit_set : blacklisted

// UNSUPPORTED : g_strstr_len : blacklisted

// UNSUPPORTED : g_strtod : blacklisted

// UNSUPPORTED : g_strup : blacklisted

// UNSUPPORTED : g_strv_contains : blacklisted

// UNSUPPORTED : g_strv_equal : blacklisted

// UNSUPPORTED : g_strv_get_type : blacklisted

// UNSUPPORTED : g_strv_length : blacklisted

// UNSUPPORTED : g_test_add_data_func : blacklisted

// UNSUPPORTED : g_test_add_data_func_full : blacklisted

// UNSUPPORTED : g_test_add_func : blacklisted

// UNSUPPORTED : g_test_add_vtable : blacklisted

// UNSUPPORTED : g_test_assert_expected_messages_internal : blacklisted

// UNSUPPORTED : g_test_bug : blacklisted

// UNSUPPORTED : g_test_bug_base : blacklisted

// UNSUPPORTED : g_test_build_filename : blacklisted

// UNSUPPORTED : g_test_create_case : blacklisted

// UNSUPPORTED : g_test_create_suite : blacklisted

// UNSUPPORTED : g_test_expect_message : blacklisted

// UNSUPPORTED : g_test_fail : blacklisted

// UNSUPPORTED : g_test_failed : blacklisted

// UNSUPPORTED : g_test_get_dir : blacklisted

// UNSUPPORTED : g_test_get_filename : blacklisted

// UNSUPPORTED : g_test_get_root : blacklisted

// UNSUPPORTED : g_test_incomplete : blacklisted

// UNSUPPORTED : g_test_init : blacklisted

// UNSUPPORTED : g_test_log_set_fatal_handler : blacklisted

// UNSUPPORTED : g_test_log_type_name : blacklisted

// UNSUPPORTED : g_test_maximized_result : blacklisted

// UNSUPPORTED : g_test_message : blacklisted

// UNSUPPORTED : g_test_minimized_result : blacklisted

// UNSUPPORTED : g_test_queue_destroy : blacklisted

// UNSUPPORTED : g_test_queue_free : blacklisted

// UNSUPPORTED : g_test_rand_double : blacklisted

// UNSUPPORTED : g_test_rand_double_range : blacklisted

// UNSUPPORTED : g_test_rand_int : blacklisted

// UNSUPPORTED : g_test_rand_int_range : blacklisted

// UNSUPPORTED : g_test_run : blacklisted

// UNSUPPORTED : g_test_run_suite : blacklisted

// UNSUPPORTED : g_test_set_nonfatal_assertions : blacklisted

// UNSUPPORTED : g_test_skip : blacklisted

// UNSUPPORTED : g_test_subprocess : blacklisted

// UNSUPPORTED : g_test_summary : blacklisted

// UNSUPPORTED : g_test_timer_elapsed : blacklisted

// UNSUPPORTED : g_test_timer_last : blacklisted

// UNSUPPORTED : g_test_timer_start : blacklisted

// UNSUPPORTED : g_test_trap_assertions : blacklisted

// UNSUPPORTED : g_test_trap_fork : blacklisted

// UNSUPPORTED : g_test_trap_has_passed : blacklisted

// UNSUPPORTED : g_test_trap_reached_timeout : blacklisted

// UNSUPPORTED : g_test_trap_subprocess : blacklisted

// UNSUPPORTED : g_thread_error_quark : blacklisted

// UNSUPPORTED : g_thread_exit : blacklisted

// UNSUPPORTED : g_thread_pool_get_max_idle_time : blacklisted

// UNSUPPORTED : g_thread_pool_get_max_unused_threads : blacklisted

// UNSUPPORTED : g_thread_pool_get_num_unused_threads : blacklisted

// UNSUPPORTED : g_thread_pool_set_max_idle_time : blacklisted

// UNSUPPORTED : g_thread_pool_set_max_unused_threads : blacklisted

// UNSUPPORTED : g_thread_pool_stop_unused_threads : blacklisted

// UNSUPPORTED : g_thread_self : blacklisted

// UNSUPPORTED : g_thread_yield : blacklisted

// UNSUPPORTED : g_time_val_from_iso8601 : blacklisted

// UNSUPPORTED : g_timeout_add : blacklisted

// UNSUPPORTED : g_timeout_add_full : blacklisted

// UNSUPPORTED : g_timeout_add_seconds : blacklisted

// UNSUPPORTED : g_timeout_add_seconds_full : blacklisted

// UNSUPPORTED : g_timeout_source_new : blacklisted

// UNSUPPORTED : g_timeout_source_new_seconds : blacklisted

// UNSUPPORTED : g_trash_stack_height : blacklisted

// UNSUPPORTED : g_trash_stack_peek : blacklisted

// UNSUPPORTED : g_trash_stack_pop : blacklisted

// UNSUPPORTED : g_trash_stack_push : blacklisted

// UNSUPPORTED : g_try_malloc : blacklisted

// UNSUPPORTED : g_try_malloc0 : blacklisted

// UNSUPPORTED : g_try_malloc0_n : blacklisted

// UNSUPPORTED : g_try_malloc_n : blacklisted

// UNSUPPORTED : g_try_realloc : blacklisted

// UNSUPPORTED : g_try_realloc_n : blacklisted

// UNSUPPORTED : g_ucs4_to_utf16 : blacklisted

// UNSUPPORTED : g_ucs4_to_utf8 : blacklisted

// UNSUPPORTED : g_unichar_break_type : blacklisted

// UNSUPPORTED : g_unichar_combining_class : blacklisted

// UNSUPPORTED : g_unichar_compose : blacklisted

// UNSUPPORTED : g_unichar_decompose : blacklisted

// UNSUPPORTED : g_unichar_digit_value : blacklisted

// UNSUPPORTED : g_unichar_fully_decompose : blacklisted

// UNSUPPORTED : g_unichar_get_mirror_char : blacklisted

// UNSUPPORTED : g_unichar_get_script : blacklisted

// UNSUPPORTED : g_unichar_isalnum : blacklisted

// UNSUPPORTED : g_unichar_isalpha : blacklisted

// UNSUPPORTED : g_unichar_iscntrl : blacklisted

// UNSUPPORTED : g_unichar_isdefined : blacklisted

// UNSUPPORTED : g_unichar_isdigit : blacklisted

// UNSUPPORTED : g_unichar_isgraph : blacklisted

// UNSUPPORTED : g_unichar_islower : blacklisted

// UNSUPPORTED : g_unichar_ismark : blacklisted

// UNSUPPORTED : g_unichar_isprint : blacklisted

// UNSUPPORTED : g_unichar_ispunct : blacklisted

// UNSUPPORTED : g_unichar_isspace : blacklisted

// UNSUPPORTED : g_unichar_istitle : blacklisted

// UNSUPPORTED : g_unichar_isupper : blacklisted

// UNSUPPORTED : g_unichar_iswide : blacklisted

// UNSUPPORTED : g_unichar_iswide_cjk : blacklisted

// UNSUPPORTED : g_unichar_isxdigit : blacklisted

// UNSUPPORTED : g_unichar_iszerowidth : blacklisted

// UNSUPPORTED : g_unichar_to_utf8 : blacklisted

// UNSUPPORTED : g_unichar_tolower : blacklisted

// UNSUPPORTED : g_unichar_totitle : blacklisted

// UNSUPPORTED : g_unichar_toupper : blacklisted

// UNSUPPORTED : g_unichar_type : blacklisted

// UNSUPPORTED : g_unichar_validate : blacklisted

// UNSUPPORTED : g_unichar_xdigit_value : blacklisted

// UNSUPPORTED : g_unicode_canonical_decomposition : blacklisted

// UNSUPPORTED : g_unicode_canonical_ordering : blacklisted

// UNSUPPORTED : g_unicode_script_from_iso15924 : blacklisted

// UNSUPPORTED : g_unicode_script_to_iso15924 : blacklisted

// UNSUPPORTED : g_unix_error_quark : blacklisted

// UNSUPPORTED : g_unix_fd_add : blacklisted

// UNSUPPORTED : g_unix_fd_add_full : blacklisted

// UNSUPPORTED : g_unix_fd_source_new : blacklisted

// UNSUPPORTED : g_unix_open_pipe : blacklisted

// UNSUPPORTED : g_unix_set_fd_nonblocking : blacklisted

// UNSUPPORTED : g_unix_signal_add : blacklisted

// UNSUPPORTED : g_unix_signal_add_full : blacklisted

// UNSUPPORTED : g_unix_signal_source_new : blacklisted

// UNSUPPORTED : g_unlink : blacklisted

// UNSUPPORTED : g_unsetenv : blacklisted

// UNSUPPORTED : g_uri_escape_string : blacklisted

// UNSUPPORTED : g_uri_list_extract_uris : blacklisted

// UNSUPPORTED : g_uri_parse_scheme : blacklisted

// UNSUPPORTED : g_uri_unescape_segment : blacklisted

// UNSUPPORTED : g_uri_unescape_string : blacklisted

// UNSUPPORTED : g_usleep : blacklisted

// UNSUPPORTED : g_utf16_to_ucs4 : blacklisted

// UNSUPPORTED : g_utf16_to_utf8 : blacklisted

// UNSUPPORTED : g_utf8_casefold : blacklisted

// UNSUPPORTED : g_utf8_collate : blacklisted

// UNSUPPORTED : g_utf8_collate_key : blacklisted

// UNSUPPORTED : g_utf8_collate_key_for_filename : blacklisted

// UNSUPPORTED : g_utf8_find_next_char : blacklisted

// UNSUPPORTED : g_utf8_find_prev_char : blacklisted

// UNSUPPORTED : g_utf8_get_char : blacklisted

// UNSUPPORTED : g_utf8_get_char_validated : blacklisted

// UNSUPPORTED : g_utf8_make_valid : blacklisted

// UNSUPPORTED : g_utf8_normalize : blacklisted

// UNSUPPORTED : g_utf8_offset_to_pointer : blacklisted

// UNSUPPORTED : g_utf8_pointer_to_offset : blacklisted

// UNSUPPORTED : g_utf8_prev_char : blacklisted

// UNSUPPORTED : g_utf8_strchr : blacklisted

// UNSUPPORTED : g_utf8_strdown : blacklisted

// UNSUPPORTED : g_utf8_strlen : blacklisted

// UNSUPPORTED : g_utf8_strncpy : blacklisted

// UNSUPPORTED : g_utf8_strrchr : blacklisted

// UNSUPPORTED : g_utf8_strreverse : blacklisted

// UNSUPPORTED : g_utf8_strup : blacklisted

// UNSUPPORTED : g_utf8_substring : blacklisted

// UNSUPPORTED : g_utf8_to_ucs4 : blacklisted

// UNSUPPORTED : g_utf8_to_ucs4_fast : blacklisted

// UNSUPPORTED : g_utf8_to_utf16 : blacklisted

// UNSUPPORTED : g_utf8_validate : blacklisted

// UNSUPPORTED : g_utf8_validate_len : blacklisted

// UNSUPPORTED : g_uuid_string_is_valid : blacklisted

// UNSUPPORTED : g_uuid_string_random : blacklisted

// UNSUPPORTED : g_variant_get_gtype : blacklisted

// UNSUPPORTED : g_variant_is_object_path : blacklisted

// UNSUPPORTED : g_variant_is_signature : blacklisted

// UNSUPPORTED : g_variant_parse : blacklisted

// UNSUPPORTED : g_variant_parse_error_print_context : blacklisted

// UNSUPPORTED : g_variant_parse_error_quark : blacklisted

// UNSUPPORTED : g_variant_parser_get_error_quark : blacklisted

// UNSUPPORTED : g_variant_type_checked_ : blacklisted

// UNSUPPORTED : g_variant_type_string_get_depth_ : blacklisted

// UNSUPPORTED : g_variant_type_string_is_valid : blacklisted

// UNSUPPORTED : g_variant_type_string_scan : blacklisted

// UNSUPPORTED : g_vasprintf : blacklisted

// UNSUPPORTED : g_vfprintf : blacklisted

// UNSUPPORTED : g_vprintf : blacklisted

// UNSUPPORTED : g_vsnprintf : blacklisted

// UNSUPPORTED : g_vsprintf : blacklisted

// UNSUPPORTED : g_warn_message : blacklisted
