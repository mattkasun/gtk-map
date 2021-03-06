package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	gtkmap "github.com/mattkasun/gtk-map"
	"log"
)

func main() {
	//initalize gtk
	gtk.Init(nil)
	//create map widget
	m, err := gtkmap.MapNew()
	if err != nil {
		log.Fatal(err)
	}
	err = m.SetProperty("map-source", gtkmap.SourceGoogleHybrid)
	if err != nil {
		fmt.Println("failed to set source", err)
	}
	m.SetCenterAndZoom(45.18, -75.93, 10)

	//create track and points
	t := gtkmap.TrackNew()
	p1 := gtkmap.PointNewDegress(45.2, -75.99)
	p2 := gtkmap.PointNewDegress(45.3, -76)
	p3 := gtkmap.PointNewDegress(45.5, -76.2)
	t.AddPoint(p1)
	t.AddPoint(p2)
	t.AddPoint(p3)
	m.TrackAdd(t)

	t.SetProperty("editable", true)
	//create and pack containters
	box1, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	//pack top box
	box1.PackStart(m, true, true, 0)

	//create main window
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//start
	window.Add(box1)
	window.SetDefaultSize(400, 400)
	window.ShowAll()

	gtk.Main()
}
