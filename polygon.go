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
	"unsafe"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

type Polygon struct {
	gtk.Widget
}

//OsmGpsMapPolygon*		osm_gps_map_polygon_new           (void);
func PolygonNew() *Polygon {
	c := C.osm_gps_map_polygon_new()
	obj := glib.Take(unsafe.Pointer (c))
	widget := &Polygon{gtk.Widget{glib.InitiallyUnowned{obj}}}
	return widget
}

func (t *Polygon) Native () *C.OsmGpsMapPolygon {
	return( (*C.OsmGpsMapPolygon) (unsafe.Pointer(t.GObject)))
}


/**
 * osm_gps_map_polygon_get_track:
 * @poly: a #OsmGpsMapPolygon
 *
 * Returns: (transfer none): The #OsmGpsMapTrack of the polygon
 **/
//OsmGpsMapTrack*			osm_gps_map_polygon_get_track(OsmGpsMapPolygon* poly);
func (p *Polygon) GetTrack () (*Track) {
	t := C.osm_gps_map_polygon_get_track(p.Native())
	obj := glib.Take(unsafe.Pointer (t))
	widget := &Track{gtk.Widget{glib.InitiallyUnowned{obj}}}
	return widget
}
