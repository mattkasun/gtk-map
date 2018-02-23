package main

// #cgo pkg-config: osmgpsmap-1.0
//
// #include <gtk/gtk.h>
// #include <osm-gps-map.h>
import "C"
import (
	"errors"
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

type Map struct {
	gtk.Widget
}

func MapNew() (*Map, error) {
	c := C.osm_gps_map_new()
	if c == nil {
		return nil, errors.New("nil pointer")
	}

	p := unsafe.Pointer(c)
	fmt.Printf("c: %T %V\n", c, c)
	fmt.Printf("p: %T %V\n", p, p)
	obj := glib.Take(unsafe.Pointer(c))
	fmt.Printf("obj: %T %V\n", obj, obj)
	widget := &Map{gtk.Widget{glib.InitiallyUnowned{obj}}}
	fmt.Printf("widget: %T %V\n", widget, widget)

	return widget, nil
}

func main() {
	gtk.Init(nil)
	m, err := MapNew()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nwe got a map %T %v\n", m, m)
}
