package main

import (
	"github.com/pekim/gobbi/lib/gobject"
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/pekim/gobbi/lib/pango"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

const (
	COL_NUMBER = iota
	COL_TEXT
	COL_LENGTH
	COL_COLOUR
)

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("A window title")
	window.SetDefaultSize(300, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 20).Container()
	container.Widget().SetMarginStart(20)
	container.Widget().SetMarginEnd(20)
	container.Add(gtk.LabelNew("A tree view").Widget())
	container.Add(createTree().Widget())

	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}

func createTree() *gtk.TreeView {
	listStore := gtk.ListStoreNewv([]gobject.Type{
		gobject.TYPE_INT, gobject.TYPE_STRING, gobject.TYPE_INT, gobject.TYPE_STRING})
	addRow(listStore, 42, "some text")
	addRow(listStore, 43, "some more text")
	addRow(listStore, 43, "yet more text, don't you know")
	addRow(listStore, 107, "qaz")
	addRow(listStore, 109, "qaz qwerty")

	tree := gtk.TreeViewNewWithModel(listStore.TreeModel())

	addColumn(tree, "number", COL_NUMBER, -1)
	addColumn(tree, "text", COL_TEXT, COL_COLOUR)
	addColumn(tree, "length of text", COL_LENGTH, -1)

	return tree
}

func addRow(store *gtk.ListStore, number int32, text string) {
	iter := store.Append()

	valueNumber := gobject.ValueNew(gobject.TYPE_INT)
	valueNumber.SetInt(number)
	store.SetValue(iter, COL_NUMBER, valueNumber)

	valueText := gobject.ValueNew(gobject.TYPE_STRING)
	valueText.SetString(text)
	store.SetValue(iter, COL_TEXT, valueText)

	valueLength := gobject.ValueNew(gobject.TYPE_INT)
	valueLength.SetInt(int32(len(text)))
	store.SetValue(iter, COL_LENGTH, valueLength)

	valueColour := gobject.ValueNew(gobject.TYPE_STRING)
	valueColour.SetString("red")
	store.SetValue(iter, COL_COLOUR, valueColour)
}

func addColumn(view *gtk.TreeView, title string, modelColumn int32, foregroundColumn int32) {
	ellipsize := gobject.ValueNew(gobject.TYPE_INT)
	ellipsize.SetInt(int32(pango.PANGO_ELLIPSIZE_END))

	renderer := gtk.CellRendererTextNew()
	renderer.Object().SetProperty("ellipsize", ellipsize)

	column := gtk.TreeViewColumnNew()
	column.SetResizable(true)
	column.SetTitle(title)
	column.PackStart(renderer.CellRenderer(), true)
	column.AddAttribute(renderer.CellRenderer(), "text", modelColumn)

	if foregroundColumn >= 0 {
		column.AddAttribute(renderer.CellRenderer(), "foreground", foregroundColumn)
	}

	view.AppendColumn(column)
}
