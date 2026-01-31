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

type pet struct {
	mode petMode
	age  petAge

	pref fyne.Preferences

	alert, asleep, dark, dirty bool
}

func loadPet(pref fyne.Preferences) *pet {
	p := &pet{pref: pref}
	p.age = petAge(pref.Float(keyAge))

	return p
}
