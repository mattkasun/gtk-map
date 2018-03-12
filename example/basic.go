package main

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	gtkmap "github.com/mattkasun/gtk-map"
	"log"
	"time"
)

var (
	m      *gtkmap.Map
	entry  *gtk.Entry
	entry2 *gtk.Entry
)

func main() {
	//initalize gtk
	gtk.Init(nil)
	//create map widget
	m, err := gtkmap.MapNew()
	if err != nil {
		log.Fatal(err)
	}
	err = m.SetProperty("map-source", gtkmap.SourceVirtualEarthSatellite)
	if err != nil {
		fmt.Println("failed to set source", err)
	}
	m.SetCenterAndZoom(45.18, -75.93, 10)

	//radio buttons for changing map source
	radio1, _ := gtk.RadioButtonNewWithLabel(nil, "OpenStreetMap")
	radio2, _ := gtk.RadioButtonNewWithLabelFromWidget(radio1, "Google Hybrid")

	//entrys for display zoom, map center and instructions
	entry, _ = gtk.EntryNew()
	entry.SetText(fmt.Sprint("Zoom Level 10"))
	entry.SetProperty("editable", false)
	entry.SetProperty("max-width-chars", 13)
	entry2, _ = gtk.EntryNew()
	entry2.SetText(fmt.Sprint("45.18 -75.93"))
	entry2.SetProperty("editable", false)
	entry3, _ := gtk.EntryNew()
	entry3.SetText("use scroll wheel to zoom, drag map with left button")
	entry3.SetProperty("editable", false)

	//create and pack containters
	box1, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box2, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	box2.PackStart(radio1, false, false, 0)
	box2.PackStart(radio2, false, false, 0)
	box2.PackStart(entry, false, false, 0)
	box2.PackStart(entry2, true, false, 0)

	//connect signals
	radio1.Connect("clicked", func() {
		fmt.Println("changing source to OpenStreetMap")
		m.SetProperty("map-source", gtkmap.SourceOpenStreetMap1)
	})
	radio2.Connect("clicked", func() {
		fmt.Println("changing source to Google")
		m.SetProperty("map-source", gtkmap.SourceGoogleHybrid)
	})

	//pack top box
	box1.PackStart(box2, false, false, 0)
	box1.PackStart(m, true, true, 0)
	box1.PackStart(entry3, false, false, 0)

	//create main window
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//periodic function to update zoom level and map center indicators
	go func() {
		for {
			time.Sleep(time.Second)
			zoom, _ := m.GetProperty("zoom")
			s := fmt.Sprint("Zoom Level ", zoom)
			_, err := glib.IdleAdd(entry.SetText, s)
			if err != nil {
				log.Fatal("idle add failed: ", err)
			}
			lat, _ := m.GetProperty("latitude")
			lng, _ := m.GetProperty("longitude")
			s = fmt.Sprint(lat, lng)
			_, err = glib.IdleAdd(entry2.SetText, s)
			if err != nil {
				log.Fatal("idle add failed: ", err)
			}
		}
	}()

	//start
	window.Add(box1)
	window.SetDefaultSize(400, 400)
	window.ShowAll()

	gtk.Main()
}
