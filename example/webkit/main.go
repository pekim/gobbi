package main

import (
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/pekim/gobbi/lib/webkit2"
	"os"
	"runtime"
)

/*
Create a window with webkit browser instance, and load a page.
*/

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("webkit example")
	window.SetDefaultSize(800, 600)
	window.Widget().ConnectDestroy(func(_ *gtk.Widget) { gtk.MainQuit() })

	webview := webkit2.WebViewNew()
	webview.ConnectClose(func(*webkit2.WebView) {
		window.Widget().Destroy()
	})
	webview.LoadUri("https://www.webkitgtk.org/")
	webview.Widget().GrabFocus()

	window.Container().Add(webview.Widget())

	window.Widget().ShowAll()

	gtk.Main()
}
