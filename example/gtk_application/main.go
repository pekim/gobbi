package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gio"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
)

/*
An example of using gtk.Application.
*/

func init() {
	// Ensure that the ui's main thread is locked to the main thread.
	runtime.LockOSThread()
}

func main() {
	app := gtk.ApplicationNew("pekim.gobbi.example.application", gio.APPLICATION_FLAGS_NONE)

	app.ConnectWindowAdded(func(window *gtk.Window) {
		fmt.Printf("window added : %p : %s\n", window, window.GetTitle())
	})

	app.ConnectWindowRemoved(func(window *gtk.Window) {
		fmt.Printf("window removed : %p : %s\n", window, window.GetTitle())
	})

	app.Application().ConnectActivate(func() {
		fmt.Println("activate")
		createWindow(app, 1)
	})

	app.Application().ConnectStartup(func() {
		fmt.Println("startup")
	})

	app.Application().ConnectShutdown(func() {
		fmt.Println("shutdown")
	})

	status := app.Application().Run(os.Args)
	os.Exit(int(status))
}

func createWindow(app *gtk.Application, windowNumber int) {
	window := gtk.ApplicationWindowNew(app).Window()
	window.SetTitle(fmt.Sprintf("window #%d", windowNumber))
	window.SetDefaultSize(300, 300)

	button := gtk.ButtonNewWithLabel("click, to create new window")
	button.ConnectClicked(func() {
		createWindow(app, windowNumber+1)
	})
	window.Container().Add(button.Widget())

	app.AddWindow(window)
	window.Widget().ShowAll()

}
