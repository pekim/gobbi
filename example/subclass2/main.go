package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gobject"
	"github.com/pekim/gobbi/lib/gtk"
)

func main() {
	//sc.SubclassCreate()
	gobject.SubclassCreate()
	gobject.FfiClosure()

	derived := gtk.WidgetDerive("test_widget", gtk.WidgetVirtualFunctions{})
	fmt.Println(derived)
}
