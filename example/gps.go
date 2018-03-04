package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
	gtkmap "github.com/mattkasun/gtk-map"
)

var (
	m	*gtkmap.Map
	entry	*gtk.Entry
	entry2	*gtk.Entry

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

	osd :=gtkmap.OsdNewFull(true, true, true, true, true, false, false, false, 30.5)
	m.LayerAdd(osd)
	// add gps point
	lat := 45.0
	lng := -76.0
	head := 45.0
	m.GpsAdd(lat, lng, head)

	//create and pack containters
	box1, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)


	//pack top box
	box1.PackStart(m, true, true, 0)

	//create main window
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//periodic function to clear and update positionss
	go func () {
		for {
			time.Sleep(time.Second)
			m.GpsClear()
			time.Sleep(time.Second)
			lat := lat + rand.Float64()
			lng := lng + rand.Float64()
			head:= head + rand.Float64()
			m.GpsAdd(lat,lng,head)
			}
	}()

	//start
	window.Add(box1)
	window.SetDefaultSize(400,400)
	window.ShowAll()

	gtk.Main()
}
