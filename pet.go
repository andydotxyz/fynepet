package main

type petMode int

const (
	modeNone petMode = iota
	modeFeed
	modeLight
	modeGame
	modeMeds
	modeClean
	modeStats
	modeDiscipline
	modeAlert
)

type pet struct {
	mode petMode
}
