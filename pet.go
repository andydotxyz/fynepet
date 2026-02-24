package main

import (
	"fyne.io/fyne/v2"
)

type petMode int

const (
	modeHome petMode = iota
	modeFeed
	modeLight
	modeGame
	modeMeds
	modeClean
	modeStats
	modeDiscipline
)

const (
	keyHungry = "hungry"
	keyHappy  = "happy"

	maxStat = 4
)

type pet struct {
	mode petMode
	age  petAge

	pref fyne.Preferences

	alert, asleep, dark, dirty bool

	hungry int // 0 = full, maxStat = starving
	happy  int // 0 = sad, maxStat = happiest
}

func loadPet(pref fyne.Preferences) *pet {
	p := &pet{pref: pref}
	p.age = petAge(pref.Float(keyAge))
	p.hungry = pref.Int(keyHungry)
	p.happy = pref.Int(keyHappy)
	if p.happy == 0 && p.age >= ageBaby {
		p.happy = maxStat / 2 // default for existing pets
	}

	return p
}

// feed gives the pet food. Burger (meal=true) reduces hunger more;
// cake (meal=false) increases happiness more.
func (p *pet) feed(meal bool) {
	if meal {
		p.hungry -= 2
		p.happy++
	} else {
		p.hungry--
		p.happy += 2
	}

	if p.hungry < 0 {
		p.hungry = 0
	}
	if p.happy > maxStat {
		p.happy = maxStat
	}

	p.pref.SetInt(keyHungry, p.hungry)
	p.pref.SetInt(keyHappy, p.happy)
}
