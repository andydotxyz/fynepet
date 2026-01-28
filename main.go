package main

import (
	"bytes"
	_ "embed"
	"math/rand/v2"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var (
	//go:embed "assets/background.png"
	imageBackground []byte
	//go:embed "assets/egg.png"
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
	u := newUI(p, scr)

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
			container.New(&screenLayout{}, append([]fyne.CanvasObject{scr}, u.modes...)...),
			egg,
			container.New(&buttonLayout{}, u.b1, u.b2, u.b3))))

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			if p.asleep {
				if p.dark {
					if i == 1 {
						homePix = sleepDark1
					} else {
						homePix = sleepDark2
					}
				} else {
					if i == 1 {
						homePix = sleepLight1
					} else {
						homePix = sleepLight2
					}
				}
			} else {
				if i == 1 {
					homePix = frameEggDown
				} else {
					homePix = frameEggUp
				}
			}

			if p.dirty {
				for id, row := range homePix {
					if i == 0 {
						homePix[id] = row | waste1[id]
					} else {
						homePix[id] = row | waste2[id]
					}
				}
			}

			i++
			if i > 1 {
				i = 0
			}

			if p.mode != modeHome {
				continue
			}
			pix = homePix
			fyne.Do(scr.Refresh)
		}
	}()

	go func() {
		for { // TODO remove this for real lifecycle
			delay := rand.IntN(15 * 60)
			time.Sleep(time.Second * time.Duration(delay))

			fyne.Do(func() {
				if p.alert {
					return
				}

				u.alert()
			})
		}
	}()

	w.SetPadded(false)
	w.Resize(fyne.NewSize(680/2, 880/2))
	w.ShowAndRun()
}
