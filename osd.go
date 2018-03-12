// Copyright (c) 2018 Matthew R Kasun <mkasun@gmail.com>
//
// This package provides a GTK map widget. It uses osm-gps-map as a backend.
// It was inspired by the gtkmap package at github.com/mewmew/gtkmap

// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package gtkmap

// #cgo pkg-config: osmgpsmap-1.0
//
// #include <gtk/gtk.h>
// #include <osm-gps-map.h>
import "C"
import (
	//"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

type Osd struct {
	gtk.Widget
}

//GType               osm_gps_map_osd_get_type (void);
//OsmGpsMapOsd*       osm_gps_map_osd_new      (void_);
//func OsdNew() *OsmGpsMapOsd {
func OsdNew() *Osd {
	c := C.osm_gps_map_osd_new()
	obj := glib.Take(unsafe.Pointer(c))
	widget := &Osd{gtk.Widget{glib.InitiallyUnowned{obj}}}
	return widget
}

func OsdNewFull(scale, coords, crosshair, dpad, zoom, gpsindpad, gpsinzoom, copywright bool, radius float64) *Osd {
	c := C.osm_gps_map_osd_new()
	obj := glib.Take(unsafe.Pointer(c))
	widget := &Osd{gtk.Widget{glib.InitiallyUnowned{obj}}}
	if scale {
		widget.SetProperty("show-scale", true)
	}
	if coords {
		widget.SetProperty("show-coordinates", true)
	}
	if crosshair {
		widget.SetProperty("show-crosshair", true)
	}
	if dpad {
		widget.SetProperty("show-dpad", true)
	}
	if zoom {
		widget.SetProperty("show-zoom", true)
	}
	if gpsindpad {
		widget.SetProperty("show-gps-in-dpad", true)
	}
	if gpsinzoom {
		widget.SetProperty("show-gps-in-zoom", true)
	}
	if copywright {
		widget.SetProperty("show-copywright", true)
	}
	widget.SetProperty("dpad-radius", radius)

	return widget
}

func (o *Osd) Native() *C.OsmGpsMapOsd {
	osd := (*C.OsmGpsMapOsd)(unsafe.Pointer(o.GObject))
	return osd
}

func (o *Osd) ToLayer() *C.OsmGpsMapLayer {
	osd := (*C.OsmGpsMapLayer)(unsafe.Pointer(o.GObject))
	return osd
}
