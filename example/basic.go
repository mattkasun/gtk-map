package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	gtkmap "github.com/mattkasun/gtk-map"
)



func main() {
	gtk.Init(nil)
	m, err := gtkmap.MapNew()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nwe got a map %T %v\n", m, m)
	coord := gtkmap.Coord (45.18, -75.93)
	m.SetCenterAndZoom(coord, 10)
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	window.Add(m)
	window.ShowAll()
	gtk.Main()
}
