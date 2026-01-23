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

	p *pet
}

func newUI(p *pet) *ui {
	u := &ui{p: p}
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

	u.b1 = newButton(func() {
		p.mode++
		if p.mode > modeDiscipline {
			p.mode = modeNone
		}

		u.refresh()
	})
	u.b2 = newButton(func() {
		log.Println("Tap B")
	})
	u.b3 = newButton(func() {
		u.p.mode = modeNone
		u.refresh()
	})
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
