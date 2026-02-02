package main

import (
	"bytes"
	_ "embed"
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

	p := loadPet(a.Preferences())
	u := newUI(p)
	go p.runAging(u)

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
			container.New(&screenLayout{}, append([]fyne.CanvasObject{u.scr}, u.modes...)...),
			egg,
			container.NewThemeOverride(
				container.New(&buttonLayout{}, u.b1, u.b2, u.b3),
				lightThemeOverride{Theme: theme.DefaultTheme()}))))

	go func() {
		i := 0
		for {
			p.asleep = time.Now().Hour() < 9 || time.Now().Hour() > 21
			fyne.Do(func() {
				if p.asleep {
					if i == 1 {
						u.homePix = sleepLight1
					} else {
						u.homePix = sleepLight2
					}
				} else {
					u.homePix = ageFrames[int(p.age)][i]
				}

				if p.dirty {
					for id, row := range u.homePix {
						if i == 0 {
							u.homePix[id] = row | waste1[id]
						} else {
							u.homePix[id] = row | waste2[id]
						}
					}
				}
			})

			i++
			if i > 1 {
				i = 0
			}

			if u.hold {
				time.Sleep(time.Second)
				continue
			}
			fyne.Do(func() {
				if p.dark {
					if p.asleep {
						if i == 1 {
							u.scr.setPixels(sleepDark1)
						} else {
							u.scr.setPixels(sleepDark2)
						}
					} else {
						u.scr.setPixels(frameBlack)
					}
				} else {
					u.scr.setPixels(u.homePix)
				}
			})
			time.Sleep(time.Second)
		}
	}()

	w.SetPadded(false)
	w.Resize(fyne.NewSize(680/2, 880/2))
	w.ShowAndRun()
}
