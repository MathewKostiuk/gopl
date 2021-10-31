// Mandelbrot emits a PNG image of the Mandelbrot fractal.

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	routines               = 1024
)

func main() {
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}

	img := Supersample(1024, 1024)

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func Supersample(width, height int) image.Image {
	doubleImg := normalImg(width*2, height*2)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		dpy := py * 2
		for px := 0; px < width; px++ {
			dpx := px * 2
			c0 := doubleImg.At(dpx, dpy)
			c1 := doubleImg.At(dpx+1, dpy)
			c2 := doubleImg.At(dpx, dpy+1)
			c3 := doubleImg.At(dpx+1, dpy+1)
			avg := avgColor(c0, c1, c2, c3)
			img.Set(px, py, avg)
		}
	}

	return img
}

func normalImg(width, height int) image.Image {
	var wg sync.WaitGroup
	var done = make(chan struct{})
	var size = height / routines
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < routines; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			for py := size * i; py < size*(i+1); py++ {
				y := float64(py)/float64(height)*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/float64(width)*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represent complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			done <- struct{}{}
		}(i)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func avgColor(colors ...color.Color) color.Color {
	var rsum, gsum, bsum uint32
	for _, c := range colors {
		r, g, b, _ := c.RGBA()
		rsum += r
		bsum += b
		gsum += g
	}

	n := uint32(len(colors))
	r := rsum / n
	g := gsum / n
	b := bsum / n
	return color.RGBA{
		R: uint8(r / 0x101),
		G: uint8(g / 0x101),
		B: uint8(b / 0x101),
		A: 255,
	}
}
