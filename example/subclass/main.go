package main

import "fmt"

/*
#cgo pkg-config: gobject-2.0
//#include <glib-object.h>
//
#cgo pkg-config: gtk+-3.0
//#include <gtk/gtk.h>

#include "./obj-sc.h"
*/
import "C"

func main() {
	fmt.Println("sc")
}
