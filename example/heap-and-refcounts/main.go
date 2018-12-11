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

var stats0 runtime.MemStats
var stats1 runtime.MemStats
var stats2 runtime.MemStats
var stats3 runtime.MemStats

var mi0 C.struct_mallinfo
var mi1 C.struct_mallinfo
var mi2 C.struct_mallinfo
var mi3 C.struct_mallinfo

func doSomeGtkStuff() {
	runtime.ReadMemStats(&stats0)
	mi0 = C.mallinfo()

	windows := []*gtk.Window{}

	for w := 0; w < 100; w++ {
		window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
		window.SetDefaultSize(100, 100)

		container := gtk.VBoxNew(true, 0).Container()

		for i := 0; i < 1000; i++ {
			container.Add(gtk.LabelNew(fmt.Sprintf("Label %d", i)).Widget())
		}

		window.Container().Add(container.Widget())

		//window.Widget().ConnectDestroy(gtk.MainQuit)
		//window.Widget().ShowAll()

		windows = append(windows, window)
	}

	runtime.ReadMemStats(&stats1)
	mi1 = C.mallinfo()

	time.AfterFunc(time.Second/2, func() {
		glib.IdleAdd(func() bool {
			for _, w := range windows {
				w.Widget().Destroy()
			}

			gtk.MainQuit()

			return false
		})
	})

	gtk.Main()

	runtime.ReadMemStats(&stats2)
	mi2 = C.mallinfo()
}

func main() {
	gtk.Init(os.Args)
	doSomeGtkStuff()

	for i := 0; i < 10; i++ {
		for g := 100; g < 1; g++ {
			runtime.GC()
		}
		time.Sleep(time.Second / 10)

		runtime.ReadMemStats(&stats3)
		mi3 = C.mallinfo()

		fmt.Printf("%v %v %v %v \n",
			bToMb(stats0.HeapAlloc),
			bToMb(stats1.HeapAlloc),
			bToMb(stats2.HeapAlloc),
			bToMb(stats3.HeapAlloc),
		)

		fmt.Printf("%v %v %v %v \n",
			bToMb(stats0.HeapObjects),
			bToMb(stats1.HeapObjects),
			bToMb(stats2.HeapObjects),
			bToMb(stats3.HeapObjects),
		)

		fmt.Printf("%v %v %v %v \n",
			bToMb(uint64(mi0.uordblks)),
			bToMb(uint64(mi1.uordblks)),
			bToMb(uint64(mi2.uordblks)),
			bToMb(uint64(mi3.uordblks)),
		)

		fmt.Println()
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024
}
