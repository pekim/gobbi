package main

import (
	"fmt"
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
	window.SetDefaultSize(400, 300)

	container := gtk.BoxNew(gtk.GTK_ORIENTATION_VERTICAL, 20).Container()
	container.Widget().SetMarginTop(20)
	container.Widget().SetMarginLeft(20)
	container.Widget().SetMarginRight(20)

	selectionLabel := gtk.LabelNew("A tree view")
	selectionLabel.Widget().SetHalign(gtk.GTK_ALIGN_START)

	tree := createTree(selectionLabel)
	scrolledWindow := gtk.ScrolledWindowNew(nil, nil)
	scrolledWindow.Widget().SetVexpand(true)
	scrolledWindow.Container().Add(tree.Widget())

	container.Add(scrolledWindow.Widget())
	container.Add(selectionLabel.Widget())

	window.Container().Add(container.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}

func createTree(selectionLabel *gtk.Label) *gtk.TreeView {
	listStore := gtk.ListStoreNewv([]gobject.Type{
		gobject.TYPE_INT, gobject.TYPE_STRING, gobject.TYPE_INT, gobject.TYPE_STRING})
	addRow(listStore, 42, "some text")
	addRow(listStore, 43, "some more text")
	addRow(listStore, 43, "yet more text, don't you know")
	addRow(listStore, 107, "qaz")
	addRow(listStore, 109, "qaz qwerty")
	for i := int32(200); i < 500; i++ {
		addRow(listStore, i, fmt.Sprintf("just a bunch more rows %d", i))
	}

	tree := gtk.TreeViewNewWithModel(listStore.TreeModel())

	addColumn(tree, "number", COL_NUMBER, -1)
	addColumn(tree, "text", COL_TEXT, COL_COLOUR)
	addColumn(tree, "length of text", COL_LENGTH, -1)

	selection := tree.GetSelection()
	selection.SetMode(gtk.GTK_SELECTION_SINGLE)
	selection.ConnectChanged(selectionChange(selection, selectionLabel))

	return tree
}

func addRow(store *gtk.ListStore, number int32, text string) {
	valueNumber := gobject.ValueNew(gobject.TYPE_INT)
	valueNumber.SetInt(number)

	valueText := gobject.ValueNew(gobject.TYPE_STRING)
	valueText.SetString(text)

	valueLength := gobject.ValueNew(gobject.TYPE_INT)
	valueLength.SetInt(int32(len(text)))

	valueColour := gobject.ValueNew(gobject.TYPE_STRING)
	valueColour.SetString("red")

	iter := store.Append()
	store.SetValue(iter, COL_NUMBER, valueNumber)
	store.SetValue(iter, COL_TEXT, valueText)
	store.SetValue(iter, COL_LENGTH, valueLength)
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
	column.SetMinWidth(50)
	column.PackStart(renderer.CellRenderer(), true)
	column.AddAttribute(renderer.CellRenderer(), "text", modelColumn)

	if foregroundColumn >= 0 {
		column.AddAttribute(renderer.CellRenderer(), "foreground", foregroundColumn)
	}

	view.AppendColumn(column)
}

func selectionChange(selection *gtk.TreeSelection, selectionLabel *gtk.Label) func() {
	return func() {
		isSelected, model, iter := selection.GetSelected()

		if isSelected {
			number := model.GetValue(iter, COL_NUMBER).GetInt()
			text := model.GetValue(iter, COL_TEXT).GetString()
			selectionLabel.SetText(fmt.Sprintf("selected : %d, %s\n", number, text))
		} else {
			selectionLabel.SetText("no selection")
		}
	}
}
