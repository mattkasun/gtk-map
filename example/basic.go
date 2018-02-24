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
	source, err := m.SetSource(gtkmap.SourceOpenStreetMap1)
	fmt.Printf("\nsource %T %v\n", source, source)
	fmt.Printf("\nerr  %T %v\n", err, err)
	coord := gtkmap.Coord (45.18, -75.93)
	m.SetCenterAndZoom(coord, 10)
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	window.Add(m)
	window.ShowAll()
	m.SetZoom(5)
	source, err = m.SetSource(gtkmap.SourceGoogleHybrid)
	fmt.Printf("\nsource %T %v\n", source, source)
	fmt.Printf("\nerr  %T %v\n", err, err)

	gtk.Main()
}
