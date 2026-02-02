package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/andydotxyz/fynepet/pkg/fynepet"
)

//go:embed "invert-colors.svg"
var invert []byte

func main() {
	a := app.NewWithID("xyz.andy.fynepet.editor")
	w := a.NewWindow("Fyne Pet Editor")

	bg := canvas.NewRectangle(color.Gray{Y: 60})
	t := &tapper{}
	t.ExtendBaseWidget(t)

	// outer frame fynepet.Pixels{4294967295, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 2147483649, 4294967295}
	// crosshair fynepet.Pixels{98304, 0, 0, 0, 0, 0, 0, 2147581953, 2147581953, 0, 0, 0, 0, 0, 0, 98304}
	t.SetPixels(fynepet.Pixels{98304, 0, 0, 0, 0, 0, 0, 2147581953, 2147581953, 0, 0, 0, 0, 0, 0, 98304})

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			a.Clipboard().SetContent(t.Export())
		}),
		widget.NewToolbarAction(theme.NewThemedResource(fyne.NewStaticResource("invert-colors.svg", invert)), t.Invert),
	)

	w.SetContent(container.NewBorder(bar, nil, nil, nil,
		container.NewStack(bg, t)))
	w.SetPadded(false)
	w.Resize(fyne.NewSize(320, 160+bar.MinSize().Height))
	w.ShowAndRun()
}

type tapper struct {
	fynepet.Screen
}

func (t *tapper) Tapped(ev *fyne.PointEvent) {
	cellW := t.Size().Width / 32
	cellH := t.Size().Height / 16
	x := int(ev.Position.X / cellW)
	y := int(ev.Position.Y / cellH)

	t.TogglePixel(x, y)
}
