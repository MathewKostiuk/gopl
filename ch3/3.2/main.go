// Surface computes an SVG rendering of a 3D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrang..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	peaks         = "#ff0000"           // Colour of the peaks (red)
	valleys       = "#0000ff"           // Colour of the valleys (blue)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		writeSVG(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func writeSVG(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, neg := corner(i+1, j)
			if math.IsNaN(ax) {
				continue
			}
			bx, by, neg := corner(i, j)
			if math.IsNaN(bx) {
				continue
			}
			cx, cy, neg := corner(i, j+1)
			if math.IsNaN(cx) {
				continue
			}
			dx, dy, neg := corner(i+1, j+1)
			if math.IsNaN(dx) {
				continue
			}

			var fill string

			if neg {
				fill = valleys
			} else {
				fill = peaks
			}

			fmt.Fprintf(out, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				fill, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	neg := math.Signbit(z)

	if math.IsInf(z, 0) {
		return math.NaN(), math.NaN(), false
	}

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, neg
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
