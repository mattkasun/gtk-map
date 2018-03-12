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
	"math"
)

type Point struct {
	rlat float64
	rlng float64
}

func PointNew(lat, lng float64) *Point {
	var p Point
	p.rlat = lat
	p.rlng = lng
	return &p
}

func PointNewDegress(lat, lng float64) *Point {
	var p Point
	p.rlat = (lat * math.Pi / 180.0)
	p.rlng = (lng * math.Pi / 180.0)
	return &p
}

func (p *Point) Native() *C.OsmGpsMapPoint {
	c := C.osm_gps_map_point_new_radians((C.float)(p.rlat), (C.float)(p.rlng))
	return c
}
