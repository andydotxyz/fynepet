package main

import (
	_ "embed"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var (
	//go:embed "assets/feed.svg"
	picFeed []byte

	//go:embed "assets/light.svg"
	picLight []byte

	//go:embed "assets/game.svg"
	picGame []byte

	//go:embed "assets/medicine.svg"
	picMedicine []byte

	//go:embed "assets/wash.svg"
	picWash []byte

	//go:embed "assets/stats.svg"
	picStats []byte

	//go:embed "assets/discipline.svg"
	picDiscipline []byte

	//go:embed "assets/alert.svg"
	picAlert []byte
)

type ui struct {
	b1, b2, b3 *button
	modes      []fyne.CanvasObject

	p   *pet
	scr *canvas.Raster
}

func newUI(p *pet, s *canvas.Raster) *ui {
	u := &ui{p: p, scr: s}
	u.modes = []fyne.CanvasObject{
		u.newIcon("feed", picFeed),
		u.newIcon("light", picLight),
		u.newIcon("game", picGame),
		u.newIcon("medicine", picMedicine),
		u.newIcon("wash", picWash),
		u.newIcon("stats", picStats),
		u.newIcon("discipline", picDiscipline),
		u.newIcon("alert", picAlert),
	}

	menuChoice := 0
	cancel := func() {
		menuChoice = 0
		u.p.mode = modeHome
		pix = homePix
		u.scr.Refresh()
		u.refresh()
	}
	u.b1 = newButton(func() {
		if menuChoice > 0 {
			if u.p.mode == modeLight {
				if menuChoice == 1 {
					menuChoice = 2
					pix = lightMenuOff
				} else {
					menuChoice = 1
					pix = lightMenuOn
				}

				u.scr.Refresh()
			}

			return
		}
		p.mode++
		if p.mode > modeDiscipline {
			p.mode = modeHome
		}

		u.refresh()
	})
	u.b2 = newButton(func() {
		if u.p.mode == modeLight {
			if menuChoice > 0 {
				u.p.dark = menuChoice != 1
				if u.p.dark {
					homePix = sleepDark1
				} else {
					homePix = sleepLight1 // TODO possibly not? kick home animation
				}

				cancel()
				return
			}

			if u.p.dark {
				pix = lightMenuOff
				menuChoice = 2
			} else {
				pix = lightMenuOn
				menuChoice = 1
			}
			u.scr.Refresh()
			return
		}
		log.Println("Tap B")
	})
	u.b3 = newButton(func() {
		cancel()
	})

	if p.dark {
		pix = sleepDark1
	} else {
		pix = sleepLight1
	}
	return u
}

func (u *ui) newIcon(name string, data []byte) fyne.CanvasObject {
	res := fyne.NewStaticResource("pix"+name+".svg", data)
	i := canvas.NewImageFromResource(theme.NewColoredResource(res, theme.ColorNameBackground))
	if name != "alert" || !u.p.alert {
		i.Hide()
	}
	return i
}

func (u *ui) refresh() {
	for i, m := range u.modes {
		if i >= int(modeDiscipline) {
			if u.p.alert {
				m.Show()
			} else {
				m.Hide()
			}
			continue
		}
		if i == int(u.p.mode)-1 {
			m.Show()
			continue
		}
		m.Hide()
	}
}
