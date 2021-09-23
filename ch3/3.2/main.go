// Surface computes an SVG rendering of a 3D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrang..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var width, height int
var xyscale, zscale = 0.0, 0.0

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		width, height, xyscale, zscale, p, v := parseQuery(r)
		writeSVG(w, width, height, xyscale, zscale, p, v)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func parseQuery(r *http.Request) (float64, float64, float64, float64, string, string) {
	w := r.URL.Query()["width"]
	h := r.URL.Query()["height"]
	p := r.URL.Query()["peaks"]
	v := r.URL.Query()["valleys"]
	width, err := strconv.Atoi(w[0])

	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(h[0])

	if err != nil {
		log.Fatal(err)
	}

	xyscale := width / 2 / xyrange
	zscale := float64(height) * 0.4

	return float64(width), float64(height), float64(xyscale), zscale, p[0], v[0]
}

func writeSVG(out io.Writer, width, height, xyscale, zscale float64, p, v string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%v' height='%v'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, neg := corner(i+1, j, width, height, xyscale, zscale)
			if math.IsNaN(ax) {
				continue
			}
			bx, by, neg := corner(i, j, width, height, xyscale, zscale)
			if math.IsNaN(bx) {
				continue
			}
			cx, cy, neg := corner(i, j+1, width, height, xyscale, zscale)
			if math.IsNaN(cx) {
				continue
			}
			dx, dy, neg := corner(i+1, j+1, width, height, xyscale, zscale)
			if math.IsNaN(dx) {
				continue
			}

			var fill string

			if neg {
				fill = v
			} else {
				fill = p
			}

			fmt.Fprintf(out, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				fill, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, width, height, xyscale, zscale float64) (sx, sy float64, neg bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	neg = math.Signbit(z)

	if math.IsInf(z, 0) {
		return math.NaN(), math.NaN(), false
	}

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
