package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var (
	colorUp   = color.Transparent
	colorDown = color.NRGBA{R: 0xe0, G: 0xe0, B: 0xe0, A: 0x99}
)

type button struct {
	widget.BaseWidget

	OnTapped func()
	c        *canvas.Circle
}

func newButton(fn func()) *button {
	b := &button{OnTapped: fn, c: canvas.NewCircle(colorUp)}
	b.ExtendBaseWidget(b)
	return b
}

func (b *button) Tapped(_ *fyne.PointEvent) {
	b.c.FillColor = colorDown
	b.c.Refresh()

	b.OnTapped()

	canvas.NewColorRGBAAnimation(colorDown, colorUp, canvas.DurationShort, func(c color.Color) {
		b.c.FillColor = c
		b.c.Refresh()
	}).Start()
}

func (b *button) MinSize() fyne.Size {
	return fyne.NewSquareSize(16)
}

func (b *button) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(b.c)
}
