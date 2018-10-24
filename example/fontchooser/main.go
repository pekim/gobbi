package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	fontChooserButton := gtk.FontButtonNew()
	fontChooserButton.Widget().SetHalign(gtk.GTK_ALIGN_CENTER)

	// Use the FontChooser interface implemented by the button.
	fontChooser := fontChooserButton.FontChooser()

	printFont := func() {
		fmt.Println(fontChooserButton.GetFontName())
		fmt.Println("  ", fontChooser.GetFontDesc().GetFamily())

		fontChooserButton.Widget().ModifyFont(fontChooser.GetFontDesc())
	}

	printFont()
	fontChooserButton.ConnectFontSet(printFont)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Add(fontChooserButton.Widget())
	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
