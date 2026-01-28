package main

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

	alert, asleep, dark, dirty bool
}
