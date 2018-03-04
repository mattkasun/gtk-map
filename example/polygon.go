package main

import (
	"fmt"
	"log"
	"github.com/gotk3/gotk3/gtk"
	gtkmap "github.com/mattkasun/gtk-map"
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
	m.SetCenterAndZoom(45.18, -75.93, 9)

	//create polygons
	poly := gtkmap.PolygonNew()
	t := poly.GetTrack()
	p1 := gtkmap.PointNewDegress(45.2, -75.99)
	p2 := gtkmap.PointNewDegress(45.3, -76)
	p3 := gtkmap.PointNewDegress(45.5, -76.2)
	t.AddPoint(p1)
	t.AddPoint(p2)
	t.AddPoint(p3)
	m.PolygonAdd(poly)

	poly2 := gtkmap.PolygonNew()
	t = poly2.GetTrack()
	p1 = gtkmap.PointNewDegress(45.2, -74.99)
	p2 = gtkmap.PointNewDegress(45.3, -75)
	p3 = gtkmap.PointNewDegress(45.5, -75.2)
	t.AddPoint(p1)
	t.AddPoint(p2)
	t.AddPoint(p3)
	m.PolygonAdd(poly2)




	poly.SetProperty ("editable", true)


	box1, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	but1, _ := gtk.ButtonNewWithLabel("Remove polygon")
	but2, _ := gtk.ButtonNewWithLabel("Restore polygon")
	but3, _ := gtk.ButtonNewWithLabel("Remove All")

	box1.PackStart(m, true, true, 0)
	box1.PackStart(but1, false, false, 0)
	box1.PackStart(but2, false, false, 0)
	box1.PackStart(but3, false, false, 0)

	but1.Connect ("clicked", func () {
		_ = m.PolygonRemove(poly2)
	})
	but2.Connect ("clicked", func () {
		m.PolygonAdd(poly2)
	})
	but3.Connect ("clicked", func() {
		m.PolygonRemoveAll()
	})

	//create main window
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})


	//start
	window.Add(box1)
	window.SetDefaultSize(800,400)
	window.ShowAll()

	gtk.Main()
}
