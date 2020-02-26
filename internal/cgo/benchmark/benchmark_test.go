package benchmark

import (
	"github.com/pekim/gobbi/internal/ffi"
	"github.com/pekim/gobbi/lib/glib"
	"testing"
)

func BenchmarkGiGlibCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		glib.Utf8Strdown("AaBbCc-Zz", -1)
	}
}

func BenchmarkCgoGlibCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Utf8Strdown("AaBbCc-Zz", -1)
	}
}

func BenchmarkFfiCall(b *testing.B) {
	ffi.SetErrorHandler(func(err error) {
		panic(err)
	})

	lib := ffi.OpenLibrary("GLib", "libglib-2.0.so.0")
	fn := lib.Function("g_utf8_strdown", []ffi.Type{ffi.Type_pointer, ffi.Type_slong}, ffi.Type_void)

	for n := 0; n < b.N; n++ {
		fn.Call()
	}
}

func BenchmarkCgoCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Malloc(1)
	}
}
