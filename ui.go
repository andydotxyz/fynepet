package main

import (
	_ "embed"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
	hold       bool

	p   *pet
	scr *screen

	homePix [16]int64
}

func newUI(p *pet) *ui {
	u := &ui{p: p, scr: newScreen()}

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
		u.hold = false
		if menuChoice == 0 {
			u.p.mode = modeHome
		} else {
			menuChoice = 0
			u.scr.setPixels(u.homePix)
		}
		u.refresh()
	}
	u.b1 = newButton(func() {
		if p.age < ageBaby || p.age >= ageDead {
			return
		}

		if menuChoice > 0 {
			switch u.p.mode {
			case modeFeed:
				if menuChoice == 1 {
					menuChoice = 2
					u.scr.setPixels(feedMenu2)
				} else {
					menuChoice = 1
					u.scr.setPixels(feedMenu1)
				}
			case modeLight:
				if menuChoice == 1 {
					menuChoice = 2
					u.scr.setPixels(lightMenuOff)
				} else {
					menuChoice = 1
					u.scr.setPixels(lightMenuOn)
				}
			case modeStats:
				if menuChoice == 1 {
					menuChoice = 2
					u.scr.setPixels(statsMenuDiscipline)
				} else if menuChoice == 2 {
					menuChoice = 3
					u.scr.setPixels(statsMenuHungry)
				} else if menuChoice == 3 {
					menuChoice = 4
					u.scr.setPixels(statsMenuHappy)
				} else {
					menuChoice = 1
					u.scr.setPixels(statsMenuHome)
				}

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
		if p.age < ageBaby || p.age >= ageDead {
			return
		}

		switch u.p.mode {
		case modeFeed:
			if menuChoice > 0 {
				if menuChoice == 1 {
					log.Println("TODO feed burger")
				} else {
					log.Println("TODO feed cake")
				}

				return
			}

			u.scr.setPixels(feedMenu1)
			menuChoice = 1
		case modeLight:
			if menuChoice > 0 {
				u.p.dark = menuChoice != 1
				if !u.p.dark {
					u.scr.setPixels(u.homePix)
				}

				cancel()
				return
			}

			if u.p.dark {
				u.scr.setPixels(lightMenuOff)
				menuChoice = 2
			} else {
				u.scr.setPixels(lightMenuOn)
				menuChoice = 1
			}
		case modeClean:
			u.hold = true
			go func() {
				u.scr.cleanAnimation()
				u.p.dirty = false
				u.p.alert = false
				u.hold = false
				u.refresh()
			}()
		case modeStats:
			if menuChoice > 0 {
				u.b1.OnTapped()
				return
			}

			u.scr.setPixels(statsMenuHome)
			menuChoice = 1
		case modeDiscipline:
			if p.alert {
				p.alert = false
				cancel()
			}
			return
		}
		if menuChoice != 0 {
			u.hold = true
		}
	})
	u.b3 = newButton(func() {
		if p.age < ageBaby || p.age >= ageDead {
			return
		}

		cancel()
	})

	return u
}

func (u *ui) alert() {
	u.p.alert = true
	u.refresh()

	fyne.CurrentApp().SendNotification(fyne.NewNotification("Fyne Pet is Sad", "Your pet needs you..."))
}

func (u *ui) newIcon(name string, data []byte) fyne.CanvasObject {
	res := fyne.NewStaticResource("pix"+name+".svg", data)
	i := canvas.NewImageFromResource(res)

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
