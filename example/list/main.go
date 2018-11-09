package main

import (
	"github.com/pekim/gobbi/lib/gobject"
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

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 10).Container()
	container.Add(gtk.LabelNew("A tree view").Widget())

	listStore := gtk.ListStoreNewv([]gobject.Type{gobject.TYPE_INT, gobject.TYPE_STRING, gobject.TYPE_STRING})
	iter := listStore.Append()

	value0 := gobject.ValueNew(gobject.TYPE_INT)
	value0.SetInt(42)
	listStore.SetValue(iter, 0, value0)

	value1 := gobject.ValueNew(gobject.TYPE_STRING)
	value1.SetString("qaz")
	listStore.SetValue(iter, 1, value1)

	value2 := gobject.ValueNew(gobject.TYPE_STRING)
	value2.SetString("Red")
	listStore.SetValue(iter, 2, value2)

	iter = listStore.Append()

	value01 := gobject.ValueNew(gobject.TYPE_INT)
	value01.SetInt(43)
	listStore.SetValue(iter, 0, value01)

	value11 := gobject.ValueNew(gobject.TYPE_STRING)
	value11.SetString("qwerty")
	listStore.SetValue(iter, 1, value11)

	value21 := gobject.ValueNew(gobject.TYPE_STRING)
	value21.SetString("green")
	listStore.SetValue(iter, 2, value21)

	tree := gtk.TreeViewNewWithModel(listStore.TreeModel())

	column0 := gtk.TreeViewColumnNew()
	column0.SetResizable(true)
	column0.SetTitle("integer")
	tree.AppendColumn(column0)

	renderer0 := gtk.CellRendererTextNew()
	column0.PackStart(renderer0.CellRenderer(), true)
	column0.AddAttribute(renderer0.CellRenderer(), "text", 0)

	column1 := gtk.TreeViewColumnNew()
	column1.SetResizable(true)
	column1.SetTitle("text")
	tree.AppendColumn(column1)

	renderer1 := gtk.CellRendererTextNew()
	column1.PackStart(renderer1.CellRenderer(), true)
	column1.AddAttribute(renderer1.CellRenderer(), "text", 1)
	column1.AddAttribute(renderer1.CellRenderer(), "foreground", 2)

	container.Add(tree.Widget())

	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
