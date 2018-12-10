package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/glib"
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
	"time"
)

// #include <malloc.h>
import "C"

func init() {
	runtime.LockOSThread()
}

func main() {
	gtk.Init(os.Args)

	var stats0 runtime.MemStats
	runtime.ReadMemStats(&stats0)
	mi0 := C.mallinfo()

	windows := []*gtk.Window{}

	for w := 0; w < 100; w++ {
		window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
		window.SetDefaultSize(100, 100)

		container := gtk.VBoxNew(true, 0).Container()

		for i := 0; i < 1000; i++ {
			container.Add(gtk.LabelNew(fmt.Sprintf("Label %d", i)).Widget())
		}

		window.Container().Add(container.Widget())

		window.Widget().ConnectDestroy(gtk.MainQuit)
		//window.Widget().ShowAll()

		windows = append(windows, window)
	}

	var stats2 runtime.MemStats
	runtime.ReadMemStats(&stats2)
	mi2 := C.mallinfo()

	glib.IdleAdd(func() bool {
		for _, w := range windows {
			w.Widget().Destroy()
		}
		return false
	})

	gtk.Main()

	windows = []*gtk.Window{}

	for g := 100; g < 1; g++ {
		runtime.GC()
	}
	time.Sleep(time.Second / 10)

	var stats3 runtime.MemStats
	runtime.ReadMemStats(&stats3)
	mi3 := C.mallinfo()

	fmt.Println(
		stats0.HeapAlloc,
		stats2.HeapAlloc,
		stats3.HeapAlloc,
	)

	//fmt.Println(mi0.ordblks, mi2.ordblks, mi3.ordblks)
	fmt.Println(mi0.uordblks, mi2.uordblks, mi3.uordblks)
}

//fmt.Println(mi.ordblks, mi.uordblks)
