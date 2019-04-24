#include <avcall.h>
#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>
#include "call.h"

extern int library_count;
extern void *library_handles[];
extern char *library_names[];

extern int function_count;
extern void *functions[];
extern char *function_names[];

char* open() {
    char *error;

	int i;
	for ( i = 0; i < library_count; i++ ) {
		library_handles[i] = dlopen(library_names[i], RTLD_LAZY);
		if (!library_handles[i])
		{
			return dlerror();
		}
	}

    return NULL;
}

void close() {
	int i;
	for ( i = 0; i < library_count; i++ ) {
	    dlclose(library_handles[i]);
	}
}

void call_function(int function_index, CallData* data) {
	void* fn = get_function(function_index);
	if (!fn) {
		fprintf(stderr, "failed to load function %d, %s\n", function_index, function_names[function_index]);
		exit(1);
	}

	av_alist alist;

    switch (data->return_type) {
    case rt_int:
       	av_start_int (alist, fn, &data->value_int);
        break;
    case rt_uint:
       	av_start_uint (alist, fn, &data->value_uint);
        break;
    case rt_long:
       	av_start_long (alist, fn, &data->value_long);
        break;
    case rt_double:
       	av_start_double (alist, fn, &data->value_double);
        break;
    case rt_void:
    default:
        av_start_void (alist, fn);
    }

	av_ptr(alist, char*, "qaz %d\n");
	av_int(alist, 42);
	av_call(alist);
}

void *get_function(int function_index) {
	char *function_name = function_names[function_index];
	void *fn = functions[function_index];
	if (fn) {
		return fn;
	}

	int i;
	for ( i = 0; i < library_count; i++ ) {
		fn = dlsym(library_handles[i], function_name);
		if (fn)
		{
			functions[function_index] = fn;
			return fn;
		}
	}

	return NULL;
}
