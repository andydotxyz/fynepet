package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// outer frame [16]int64{4294967295, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 4294967295}
// crosshair [16]int64{98304, 0, 0, 0, 0, 0, 0, 2147581953, 2147581953, 0, 0, 0, 0, 0, 0, 98304}
var pix = [16]int64{98304, 0, 0, 0, 0, 0, 0, 2147581953, 2147581953, 0, 0, 0, 0, 0, 0, 98304}

func main() {
	a := app.NewWithID("xyz.andy.fynepet.editor")
	w := a.NewWindow("Fyne Pet Editor")

	bg := canvas.NewRectangle(color.Gray{Y: 60})
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			a.Clipboard().SetContent(fmt.Sprintf("%#v", pix))
		}))

	t := &tapper{}
	t.ExtendBaseWidget(t)
	w.SetContent(container.NewBorder(bar, nil, nil, nil,
		container.NewStack(bg, t)))
	w.SetPadded(false)
	w.Resize(fyne.NewSize(320, 160))
	w.ShowAndRun()
}

var img = image.NewNRGBA(image.Rect(0, 0, 32, 16))

func draw(pix *[16]int64) func(x, y int) image.Image {
	return func(_, _ int) image.Image {
		var onOff color.Color

		for y := 0; y < 16; y++ {
			for x := 0; x < 32; x++ {
				onOff = color.Transparent

				pow := int64(math.Pow(2, float64(31-x)))
				if (*pix)[y]&pow == pow {
					onOff = color.Black
				}

				img.Set(x, y, onOff)
			}
		}

		return img
	}
}

type tapper struct {
	widget.BaseWidget
}

func (t *tapper) Tapped(ev *fyne.PointEvent) {
	cellW := t.Size().Width / 32
	cellH := t.Size().Height / 16
	x := int(ev.Position.X / cellW)
	y := int(ev.Position.Y / cellH)

	pow := int64(math.Pow(2, float64(31-x)))
	pix[y] ^= pow
	t.Refresh()
}

func (t *tapper) CreateRenderer() fyne.WidgetRenderer {
	r := canvas.NewRaster(draw(&pix))
	r.ScaleMode = canvas.ImageScalePixels

	return widget.NewSimpleRenderer(r)
}
