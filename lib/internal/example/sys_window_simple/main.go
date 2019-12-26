package main

import (
	"github.com/pekim/gobbi/lib/internal/c/gtk"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	argv := os.Args
	var argc int = len(argv)
	gtk.Fn_gtk_init(&argc, &argv)

	window := gtk.Fn_gtk_window_new(0)
	gtk.Fn_gtk_window_set_title(window, "some title")
	gtk.Fn_gtk_window_set_default_size(window, 400, 300)
	gtk.Fn_gtk_widget_show_all(window)

	gtk.Fn_gtk_main()
}
