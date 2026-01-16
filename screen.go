package main

import (
	"image"
	"image/color"
	"math"
)

var (
	frameEggDown = [16]int{0, 0, 0, 0, 516096, 946176, 1468416, 1677312, 3775488, 2515968, 2515968, 3775488, 1677312, 946176, 2095104, 0}
	frameEggUp   = [16]int{0, 0, 245760, 417792, 946176, 626688, 1677312, 1468416, 2515968, 3775488, 3775488, 1468416, 946176, 368640, 1044480, 0}

	pix = frameEggDown
	img = image.NewNRGBA(image.Rect(0, 0, 32, 16))
)

func draw(pix *[16]int) func(x, y int) image.Image {
	return func(_, _ int) image.Image {
		var onOff color.Color

		for y := 0; y < 16; y++ {
			for x := 0; x < 32; x++ {
				onOff = color.Transparent

				pow := int(math.Pow(2, float64(31-x)))
				if (*pix)[y]&pow == pow {
					onOff = color.NRGBA{R: 0, G: 0, B: 0, A: 0xcc}
				}

				img.Set(x, y, onOff)
			}
		}

		return img
	}
}
