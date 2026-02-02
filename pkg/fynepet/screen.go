package fynepet

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Pixels [16]int64

type Screen struct {
	widget.BaseWidget

	out *canvas.Raster
	img *image.NRGBA

	pix Pixels
}

func NewScreen() *Screen {
	s := &Screen{}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Screen) Export() string {
	return fmt.Sprintf("%#v", s.pix)
}

func (s *Screen) Invert() {
	for y := 0; y < 16; y++ {
		for x := 0; x < 32; x++ {

			pow := int64(math.Pow(2, float64(31-x)))
			s.pix[y] ^= pow
		}
	}
	s.out.Refresh()
}

func (s *Screen) SetPixels(newPix [16]int64) {
	s.pix = newPix
	s.out.Refresh()
}

func (s *Screen) TogglePixel(x, y int) {
	pow := int64(math.Pow(2, float64(31-x)))
	s.pix[y] ^= pow
	s.out.Refresh()
}

func (s *Screen) CreateRenderer() fyne.WidgetRenderer {
	if s.out == nil {
		s.out = canvas.NewRaster(s.draw(&s.pix))
		s.out.ScaleMode = canvas.ImageScalePixels
	}

	s.img = image.NewNRGBA(image.Rect(0, 0, 32, 16))
	return widget.NewSimpleRenderer(s.out)
}

func (s *Screen) ScrollLeft() {
	for id, row := range s.pix {
		s.pix[id] = row << 1
	}
	s.out.Refresh()
}

func (s *Screen) draw(pix *Pixels) func(x, y int) image.Image {
	return func(_, _ int) image.Image {
		var onOff color.Color

		for y := 0; y < 16; y++ {
			for x := 0; x < 32; x++ {
				onOff = color.Transparent

				pow := int64(math.Pow(2, float64(31-x)))
				if (*pix)[y]&pow == pow {
					onOff = color.NRGBA{R: 0, G: 0, B: 0, A: 0xcc}
				}

				s.img.Set(x, y, onOff)
			}
		}

		return s.img
	}
}
