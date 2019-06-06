package main

import (
	"fmt"
	"github.com/pekim/gobbi/lib/gtk"
)

type MyWidget struct {
}

func (w *MyWidget) Init() {
	fmt.Println("my widget init")
}

func main() {
	gtk.Init([]string{})

	//sc.SubclassCreate()
	//gobject.SubclassCreate()
	//gobject.FfiClosure()

	class := gtk.WidgetDerive("test_widget", func() gtk.WidgetDerivedInitializer {
		return &MyWidget{}
	})
	fmt.Println(class)

	class.New()
}
