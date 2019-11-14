package gi

import "fmt"

// #cgo pkg-config: gobject-introspection-1.0
// #include <girepository.h>
import "C"

const Qaz = 42

var defaultRepo *C.GIRepository

func init() {
	defaultRepo = C.g_irepository_get_default()
	fmt.Println(defaultRepo)
}
