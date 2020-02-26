package benchmark

import (
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

func BenchmarkCgoCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Malloc(1)
	}
}
