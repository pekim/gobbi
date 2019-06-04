package gobject

/*
#cgo pkg-config: libffi

#include <stdio.h>
#include <ffi.h>

void puts_binding(ffi_cif *cif, void *ret, void* args[], void *stream)
{
	*(ffi_arg *)ret = fputs(*(char **)args[0], (FILE *)stream);
}

typedef int (*puts_t)(char *);

int test_ffi()
{
  ffi_cif cif;
  ffi_type *args[1];
  ffi_closure *closure;

  void *bound_puts;
  int rc;

	closure = ffi_closure_alloc(sizeof(ffi_closure), &bound_puts);

	if (closure)
	{
		args[0] = &ffi_type_pointer;

		if (ffi_prep_cif(&cif, FFI_DEFAULT_ABI, 1,
			&ffi_type_sint, args) == FFI_OK)
		{
			if (ffi_prep_closure_loc(closure, &cif, puts_binding,
				stdout, bound_puts) == FFI_OK)
			{
				rc = ((puts_t)bound_puts)("Hello World!\n");
			}
		}
	}

	ffi_closure_free(closure);

	return 0;
}
*/
import "C"

import "fmt"

func FfiClosure() {
	fmt.Println("ffic")
	C.test_ffi()
}
