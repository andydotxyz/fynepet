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
)

type pet struct {
	mode  petMode
	alert bool
}
