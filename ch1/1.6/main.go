package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}, color.White, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 0}}

const (
	blackIndex = 0
	greenIndex = 1
	whiteIndex = 2
	redIndex   = 3
	blueIndex  = 4
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		c := parseQuery(r)
		lissajous(w, c)
	}
	http.HandleFunc("/cycles", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseQuery(r *http.Request) int {
	params := r.URL.Query()["num"]
	cycles, err := strconv.Atoi(params[0])

	if err != nil {
		log.Fatal(err)
	}

	return cycles
}

func lissajous(out io.Writer, c int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(c)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colourIndex := rand.Intn(5)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colourIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
