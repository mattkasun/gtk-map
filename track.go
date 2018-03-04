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
type Track struct {
	gtk.Widget
}

//OsmGpsMapTrack *    osm_gps_map_track_new           (void);
func TrackNew() *Track {
	c := C.osm_gps_map_track_new()
	obj := glib.Take(unsafe.Pointer (c))
	widget := &Track{gtk.Widget{glib.InitiallyUnowned{obj}}}
	return widget
}

func (t *Track) Native () *C.OsmGpsMapTrack {
	return( (*C.OsmGpsMapTrack) (unsafe.Pointer(t.GObject)))
}

/**
 * osm_gps_map_track_add_point:
 * @track: (inout): a #OsmGpsMapTrack
 * @point: (in): point to add
 *
 * Since: 0.7.0
 **/ 
//void                osm_gps_map_track_add_point     (OsmGpsMapTrack *track, const OsmGpsMapPoint *point);
func (t *Track) AddPoint (p *Point) {
	C.osm_gps_map_track_add_point (t.Native(), p.Native())
}

/**
 * osm_gps_map_track_get_points:
 * @track: (in): a #OsmGpsMapTrack
 *
 * Returns: (element-type OsmGpsMapPoint) (transfer none): list of #OsmGpsMapPoint
 *
 * Since: 0.7.0
 **/
//GSList *            osm_gps_map_track_get_points    (OsmGpsMapTrack *track);
//void                osm_gps_map_track_set_color     (OsmGpsMapTrack *track, GdkRGBA *color);
//void                osm_gps_map_track_get_color     (OsmGpsMapTrack *track, GdkRGBA *color);

/**
 * osm_gps_map_track_remove_point:
 * @track: (in): a #OsmGpsMapTrack
 * @pos: Position of the point to remove
 *
 **/
//void                osm_gps_map_track_remove_point(OsmGpsMapTrack* track, int pos);

/**
 * osm_gps_map_track_n_points:
 * @track: a #OsmGpsMapTrack
 *
 * Returns: the number of points of the track.
 **/
//int                 osm_gps_map_track_n_points(OsmGpsMapTrack* track);

/**
 * osm_gps_map_track_insert_point:
 * @track: a #OsmGpsMapTrack
 * @np: a #OsmGpsMapPoint
 * @pos: Position for the point
 *
 **/
//void                osm_gps_map_track_insert_point(OsmGpsMapTrack* track, OsmGpsMapPoint* np, int pos);

/**
 * osm_gps_map_track_get_point:
 * @track: a #OsmGpsMapTrack
 * @pos: Position of the point to get
 *
 * Returns: a #OsmGpsMapPoint
 **/
//OsmGpsMapPoint*     osm_gps_map_track_get_point(OsmGpsMapTrack* track, int pos);

