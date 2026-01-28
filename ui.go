package main

import (
	_ "embed"
	"log"
	"time"

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
	hold       bool

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
		u.hold = false
		if menuChoice == 0 {
			u.p.mode = modeHome
		} else {
			menuChoice = 0
			pix = homePix
			u.scr.Refresh()
		}
		u.refresh()
	}
	u.b1 = newButton(func() {
		if menuChoice > 0 {
			switch u.p.mode {
			case modeFeed:
				if menuChoice == 1 {
					menuChoice = 2
					pix = feedMenu2
				} else {
					menuChoice = 1
					pix = feedMenu1
				}
			case modeLight:
				if menuChoice == 1 {
					menuChoice = 2
					pix = lightMenuOff
				} else {
					menuChoice = 1
					pix = lightMenuOn
				}
			case modeStats:
				if menuChoice == 1 {
					menuChoice = 2
					pix = statsMenuDiscipline
				} else if menuChoice == 2 {
					menuChoice = 3
					pix = statsMenuHungry
				} else if menuChoice == 3 {
					menuChoice = 4
					pix = statsMenuHappy
				} else {
					menuChoice = 1
					pix = statsMenuHome
				}

			}
			u.scr.Refresh()

			return
		}
		p.mode++
		if p.mode > modeDiscipline {
			p.mode = modeHome
		}

		u.refresh()
	})
	u.b2 = newButton(func() {
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

			pix = feedMenu1
			menuChoice = 1
			u.scr.Refresh()
		case modeLight:
			if menuChoice > 0 {
				u.p.dark = menuChoice != 1
				if !u.p.dark {
					pix = homePix
					u.scr.Refresh()
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
		case modeClean:
			u.hold = true
			go func() {
				u.cleanAnimation()
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

			pix = statsMenuHome
			menuChoice = 1
			u.scr.Refresh()
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
		cancel()
	})

	return u
}

func (u *ui) alert() {
	u.p.alert = true
	u.refresh()

	fyne.CurrentApp().SendNotification(fyne.NewNotification("Attention!", "Your pet needs you..."))
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

func (u *ui) cleanAnimation() {
	i := 0
	for i < 36 {
		time.Sleep(time.Millisecond * 50)

		if i < 32 { // last 4 frames are a small freeze on the device
			for id, row := range pix {
				pix[id] = row << 1

				switch i % 32 {
				case 0:
					if id%4 == 2 {
						pix[id] |= 0x1
					}
				case 1:
					if id%4 != 0 {
						pix[id] |= 0x1
					}
				case 2:
					if id%4 != 2 {
						pix[id] |= 0x1
					}
				case 3:
					if id%2 == 0 {
						pix[id] |= 0x1
					}
				case 4:
					if id%2 == 1 {
						pix[id] |= 0x1
					}
				case 5:
					if id%4 == 0 {
						pix[id] |= 0x1
					}

				default:
					// no more pixels
				}
			}
		}
		i++

		fyne.Do(u.scr.Refresh)
	}
}
