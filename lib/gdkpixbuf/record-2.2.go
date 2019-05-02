// This is a generated file - DO NOT EDIT
// +build gdkpixbuf_2.2 gdkpixbuf_2.4 gdkpixbuf_2.6 gdkpixbuf_2.8 gdkpixbuf_2.12 gdkpixbuf_2.14 gdkpixbuf_2.18 gdkpixbuf_2.22 gdkpixbuf_2.24 gdkpixbuf_2.26 gdkpixbuf_2.28 gdkpixbuf_2.30 gdkpixbuf_2.32 gdkpixbuf_2.36

package gdkpixbuf

import (
	"C"
	"unsafe"
)

// GetDescription is a wrapper around the C function gdk_pixbuf_format_get_description.
func (recv *PixbufFormat) GetDescription() string {
	retC := C.gdk_pixbuf_format_get_description((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetExtensions is a wrapper around the C function gdk_pixbuf_format_get_extensions.
func (recv *PixbufFormat) GetExtensions() []string {
	retC := C.gdk_pixbuf_format_get_extensions((*C.GdkPixbufFormat)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetMimeTypes is a wrapper around the C function gdk_pixbuf_format_get_mime_types.
func (recv *PixbufFormat) GetMimeTypes() []string {
	retC := C.gdk_pixbuf_format_get_mime_types((*C.GdkPixbufFormat)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetName is a wrapper around the C function gdk_pixbuf_format_get_name.
func (recv *PixbufFormat) GetName() string {
	retC := C.gdk_pixbuf_format_get_name((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsWritable is a wrapper around the C function gdk_pixbuf_format_is_writable.
func (recv *PixbufFormat) IsWritable() bool {
	retC := C.gdk_pixbuf_format_is_writable((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}
