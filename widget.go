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
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)
type Source int

// Tile source repositories.
const (
	SourceNone    Source = C.OSM_GPS_MAP_SOURCE_NULL
	SourceDefault Source = SourceOpenStreetMap1
	// Sources.
	SourceOpenStreetMap1        Source = C.OSM_GPS_MAP_SOURCE_OPENSTREETMAP
	SourceOpenStreetMap2        Source = C.OSM_GPS_MAP_SOURCE_OPENSTREETMAP_RENDERER
	SourceOpenAerialMap         Source = C.OSM_GPS_MAP_SOURCE_OPENAERIALMAP
	SourceMapsForFree           Source = C.OSM_GPS_MAP_SOURCE_MAPS_FOR_FREE
	SourceOpenCycleMap          Source = C.OSM_GPS_MAP_SOURCE_OPENCYCLEMAP
	SourcePublicTransport       Source = C.OSM_GPS_MAP_SOURCE_OSM_PUBLIC_TRANSPORT
	SourceGoogleMaps            Source = C.OSM_GPS_MAP_SOURCE_GOOGLE_STREET
	SourceGoogleSatellite       Source = C.OSM_GPS_MAP_SOURCE_GOOGLE_SATELLITE
	SourceGoogleHybrid          Source = C.OSM_GPS_MAP_SOURCE_GOOGLE_HYBRID
	SourceVirtualEarth          Source = C.OSM_GPS_MAP_SOURCE_VIRTUAL_EARTH_STREET
	SourceVirtualEarthSatellite Source = C.OSM_GPS_MAP_SOURCE_VIRTUAL_EARTH_SATELLITE
	SourceVirtualEarthHybrid    Source = C.OSM_GPS_MAP_SOURCE_VIRTUAL_EARTH_HYBRID
	SourceOSMCTrails            Source = C.OSM_GPS_MAP_SOURCE_OSMC_TRAILS
)

type Map struct {
	gtk.Widget 
}

//Wraper function for
//GtkWidget* osm_gps_map_new (void);
func MapNew() (*Map, error) {
	c := C.osm_gps_map_new()
	if c == nil {
		return nil, errors.New("nil pointer")
	}
	obj := glib.Take(unsafe.Pointer(c))
	widget := &Map{gtk.Widget{glib.InitiallyUnowned{obj}}}

	return widget, nil
}

func (m *Map) Native() *C.OsmGpsMap {
	return (*C.OsmGpsMap)(unsafe.Pointer(m.GObject))
}

//use m.SetProperty("map-source", Source) to set map source


//gchar*          osm_gps_map_get_default_cache_directory (void);

//void            osm_gps_map_download_maps               (OsmGpsMap *map, OsmGpsMapPoint *pt1, OsmGpsMapPoint *pt2, int zoom_start, int zoom_end);
//void            osm_gps_map_download_cancel_all         (OsmGpsMap *map);
//void            osm_gps_map_get_bbox                    (OsmGpsMap *map, OsmGpsMapPoint *pt1, OsmGpsMapPoint *pt2);
//void            osm_gps_map_zoom_fit_bbox               (OsmGpsMap *map, float latitude1, float latitude2, float longitude1, float longitude2);

//Wrapper function for
//void            osm_gps_map_set_center_and_zoom         (OsmGpsMap *map, float latitude, float longitude, int zoom);
func (m *Map) SetCenterAndZoom(lat float64, lng float64, zoom int) {
	C.osm_gps_map_set_center_and_zoom(m.Native(), C.float(lat), C.float(lng), C.int(zoom))
}

//Wrapper function for
//void            osm_gps_map_set_center                  (OsmGpsMap *map, float latitude, float longitude);
func (m *Map) SetCenter(latitude, longitude float64) {
	C.osm_gps_map_set_center(m.Native(), C.float(latitude), C.float(longitude))
}
// SetZoom sets the zoom level of the map. It returns the new zoom level, which
// may differ if zoom was below the min or above the max zoom level of the
// current source.
//int             osm_gps_map_set_zoom                    (OsmGpsMap *map, int zoom);
func (m *Map) SetZoom (zoom int) int {
	return int(C.osm_gps_map_set_zoom(m.Native(), C.int(zoom)))
}

//void            osm_gps_map_set_zoom_offset             (OsmGpsMap *map, int zoom_offset);


// ZoomIn increases the zoom level by one. It returns the new zoom level.
//int             osm_gps_map_zoom_in                     (OsmGpsMap *map);
func (m *Map) ZoomIn() int {
	return int(C.osm_gps_map_zoom_in(m.Native()))
}

// ZoomOut decreases the zoom level by one. It returns the new zoom level.
//int             osm_gps_map_zoom_out                    (OsmGpsMap *map);
func (m *Map) ZoomOut() int {
	return int(C.osm_gps_map_zoom_out(m.Native()))
}

//void            osm_gps_map_scroll                      (OsmGpsMap *map, gint dx, gint dy);

// GetScale returns the scale at the center of the map, in meters/pixel.
//float           osm_gps_map_get_scale                   (OsmGpsMap *map);
func (m *Map) Scale() float64 {
	return float64(C.osm_gps_map_get_scale(m.Native()))
}

//void            osm_gps_map_set_keyboard_shortcut       (OsmGpsMap *map, OsmGpsMapKey_t key, guint keyval);

//void            osm_gps_map_track_add                   (OsmGpsMap *map, OsmGpsMapTrack *track);
func (m *Map) TrackAdd (t *Track) {
	C.osm_gps_map_track_add(m.Native(), t.Native())
}

//void            osm_gps_map_track_remove_all            (OsmGpsMap *map);
//gboolean        osm_gps_map_track_remove                (OsmGpsMap *map, OsmGpsMapTrack *track);

//void            osm_gps_map_polygon_add                 (OsmGpsMap *map, OsmGpsMapPolygon *poly);
func (m *Map) PolygonAdd (p *Polygon) {
	C.osm_gps_map_polygon_add(m.Native(), p.Native())
}

//void            osm_gps_map_polygon_remove_all          (OsmGpsMap *map);
func (m *Map) PolygonRemoveAll () {
	C.osm_gps_map_polygon_remove_all(m.Native())
}

//gboolean        osm_gps_map_polygon_remove                (OsmGpsMap *map, OsmGpsMapPolygon *poly);
func (m *Map) PolygonRemove (p *Polygon) bool {
	result := (C.osm_gps_map_polygon_remove(m.Native(), p.Native()))
	if result == 0 {
		return false
	}
	return true
}

//void            osm_gps_map_gps_add                     (OsmGpsMap *map, float latitude, float longitude, float heading);
func (m *Map) GpsAdd (latitude, longitude, heading float64) {
	C.osm_gps_map_gps_add(m.Native(), (C.float)(latitude), (C.float)(longitude), (C.float)(heading))
}

//void            osm_gps_map_gps_clear                   (OsmGpsMap *map);
func (m *Map) GpsClear () {
	C.osm_gps_map_gps_clear(m.Native())
}

//OsmGpsMapTrack *osm_gps_map_gps_get_track               (OsmGpsMap *map);
//OsmGpsMapImage *osm_gps_map_image_add                   (OsmGpsMap *map, float latitude, float longitude, GdkPixbuf *image);
//OsmGpsMapImage *osm_gps_map_image_add_z                 (OsmGpsMap *map, float latitude, float longitude, GdkPixbuf *image, gint zorder);
//OsmGpsMapImage *osm_gps_map_image_add_with_alignment    (OsmGpsMap *map, float latitude, float longitude, GdkPixbuf *image, float xalign, float yalign);
//OsmGpsMapImage *osm_gps_map_image_add_with_alignment_z  (OsmGpsMap *map, float latitude, float longitude, GdkPixbuf *image, float xalign, float yalign, gint zorder);
//gboolean        osm_gps_map_image_remove                (OsmGpsMap *map, OsmGpsMapImage *image);
//void            osm_gps_map_image_remove_all            (OsmGpsMap *map);
//void            osm_gps_map_layer_add                   (OsmGpsMap *map, OsmGpsMapLayer *layer);
func (m *Map) LayerAdd (o *Osd) () {
	
	C.osm_gps_map_layer_add(m.Native(),o.ToLayer())
}
//gboolean        osm_gps_map_layer_remove                (OsmGpsMap *map, OsmGpsMapLayer *layer);
//void            osm_gps_map_layer_remove_all            (OsmGpsMap *map);

//void            osm_gps_map_convert_screen_to_geographic(OsmGpsMap *map, gint pixel_x, gint pixel_y, OsmGpsMapPoint *pt);
//void            osm_gps_map_convert_geographic_to_screen(OsmGpsMap *map, OsmGpsMapPoint *pt, gint *pixel_x, gint *pixel_y);
//OsmGpsMapPoint *osm_gps_map_get_event_location          (OsmGpsMap *map, GdkEventButton *event);
//gboolean        osm_gps_map_map_redraw                  (OsmGpsMap *map);
//void            osm_gps_map_map_redraw_idle             (OsmGpsMap *map);

