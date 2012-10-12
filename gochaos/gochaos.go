package main

import (
	"os"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"github.com/mandel59/gochaos/mandelbrot"
)

func hsv(h, s, v float64) color.RGBA {
	hi, f := math.Modf(math.Mod(h / math.Pi * 3 + 6, 6))
	p := uint8(255 * v * (1-s))
	q := uint8(255 * v * (1 - f * s))
	t := uint8(255 * v * (1 - (1 - f) * s))
	V := uint8(255 * v)
	var r, g, b uint8
	switch int(hi) {
	case 0, 6:
		r, g, b = V, t, p
	case 1:
		r, g, b = q, V, p
	case 2:
		r, g, b = p, V, t
	case 3:
		r, g, b = p, q, V
	case 4:
		r, g, b = t, p, V
	case 5:
		r, g, b = V, p, q
	}
	return color.RGBA{r, g, b, 255}
}

func mapColor(ct, limit int, z complex128) color.RGBA {
	if ct == 0 {
		return hsv(cmplx.Phase(z), 1.0, 1.0)
	}
	return hsv(cmplx.Phase(z), 1.0, 0.2)
}

func Image(r image.Rectangle, min, max complex128, limit int) image.Image {
	dx, dy := r.Dx(), r.Dy()
	dc := max - min
	dr, di := real(dc) / float64(dx), imag(dc) / float64(dy)
	img := image.NewRGBA(r)
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			c := min + complex(float64(j) * dr, float64(i) * di)
			ct, z := mandelbrot.Calc(c, limit)
			img.SetRGBA(j, dy - i - 1, mapColor(ct, limit, z))
		}
	}
	return img
}

func main() {
	
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	min := -0.12065 - 0.980248i - (1e-4 + 1e-4i)
	max := -0.12065 - 0.980248i + (1e-4 + 1e-4i)
	img := Image(image.Rect(0, 0, 512, 512), min, max, 1000)
	png.Encode(f, img)
}

