package imageprocessing

import (
	"image"
	"image/color"
)

func ColorReversal(source image.Image) *image.RGBA {
	bounds := source.Bounds()
	dest := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := source.At(x, y)
			r, g, b, a := pixel.RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8

			col := color.RGBA{R: uint8(255 - r), G: uint8(255 - g), B: uint8(255 - b), A: uint8(a)}
			dest.Set(x, y, col)
		}
	}

	return dest
}
