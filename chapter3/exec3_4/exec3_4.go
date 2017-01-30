package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var width, height int
var xyscale, zscale float64
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	} else {
		height = parse(r, "height", 320)
		width = parse(r, "width", 600)
		color := parse(r, "color", 0xFFFFFF)
		xyscale = float64(width) / 2 / xyrange
		zscale = float64(height) * 0.4

		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(w, "<polygon fill='#%06X' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	}
}

func parse(r *http.Request, key string, defaultValue int) int {
	var value int
	value64, err := strconv.ParseInt(r.FormValue(key), 0, 32)
	if err != nil {
		value = defaultValue
	} else {
		value = int(value64)
	}
	return value
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
