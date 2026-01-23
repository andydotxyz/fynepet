package main

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var (
	//go:embed "background.png"
	imageBackground []byte
	//go:embed "egg.png"
	imageEgg []byte

	//go:embed "Icon.png"
	iconDark []byte
	//go:embed "Icon-light.png"
	iconLight []byte
)

func main() {
	a := app.NewWithID("xyz.andy.fynepet")
	w := a.NewWindow("Fyne Pet")

	bg := canvas.NewImageFromReader(bytes.NewReader(imageBackground), "background.png")
	egg := canvas.NewImageFromReader(bytes.NewReader(imageEgg), "egg.png")

	p := &pet{}
	scr := canvas.NewRaster(draw(&pix))
	modes := []fyne.CanvasObject{
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
		canvas.NewImageFromResource(theme.NewColoredResource(theme.ViewFullScreenIcon(), theme.ColorNameBackground)),
	}
	for _, m := range modes {
		m.Hide()
	}

	b1 := newButton(func() {
		p.mode++
		if p.mode == modeAlert {
			p.mode = modeNone
		}

		for i, m := range modes {
			if i == int(p.mode)-1 {
				m.Show()
				continue
			}
			m.Hide()
		}
	})
	b2 := newButton(func() {
		log.Println("Tap B")
	})
	b3 := newButton(func() {
		log.Println("Tap C")
	})

	scr.ScaleMode = canvas.ImageScalePixels
	under := canvas.NewImageFromReader(bytes.NewReader(iconDark), "Icon.png")
	under.FillMode = canvas.ImageFillCover
	under.Translucency = .8
	if a.Settings().ThemeVariant() == theme.VariantLight {
		under.Resource = fyne.NewStaticResource("IconLight.png", iconLight)
		under.Translucency = .5
	}

	a.Settings().AddListener(func(s fyne.Settings) {
		switch s.ThemeVariant() {
		case theme.VariantLight:
			under.Resource = fyne.NewStaticResource("IconLight.png", iconLight)
			under.Translucency = .5
		default:
			under.Resource = fyne.NewStaticResource("Icon.png", iconDark)
			under.Translucency = .8
		}
		under.Refresh()
	})
	w.SetContent(container.NewStack(under,
		container.New(&fullLayout{}, bg,
			container.New(&screenLayout{}, append([]fyne.CanvasObject{scr}, modes...)...),
			egg,
			container.New(&buttonLayout{}, b1, b2, b3))))

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)

			if i == 1 {
				pix = sleepLight1
			} else {
				pix = sleepLight2
			}
			fyne.Do(scr.Refresh)

			i++
			if i > 1 {
				i = 0
			}
		}
	}()

	w.SetPadded(false)
	w.Resize(fyne.NewSize(680/2, 880/2))
	w.ShowAndRun()
}
