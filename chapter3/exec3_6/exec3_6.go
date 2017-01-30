// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		subpixelSize           = 4
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b, a uint
			for subpixelx := 0; subpixelx < subpixelSize/2; subpixelx++ {
				for subpixely := 0; subpixely < subpixelSize/2; subpixely++ {
					x := (float64(px)+float64(subpixelx)/(subpixelSize/2))/width*(xmax-xmin) + xmin
					y := (float64(py)+float64(subpixely)/(subpixelSize/2))/height*(ymax-ymin) + ymin
					z := complex(x, y)

					currentR, currentG, currentB, currentA := mandelbrot(z).RGBA()
					r += uint(currentR)
					g += uint(currentG)
					b += uint(currentB)
					a += uint(currentA)
				}
			}

			// https://jimdoescode.github.io/2015/05/22/manipulating-colors-in-go.html
			// https://blog.golang.org/go-image-package
			// http://stackoverflow.com/questions/35374300/why-does-golang-rgba-rgba-method-use-and
			img.Set(px, py, color.RGBA{uint8((r / subpixelSize) / 0x101), uint8((g / subpixelSize) / 0x101), uint8((b / subpixelSize) / 0x101), uint8((a / subpixelSize) / 0x101)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

//!-
