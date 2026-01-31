package main

import (
	"math/rand/v2"
	"time"

	"fyne.io/fyne/v2"
)

const (
	keyAge = "age"
)

type petAge float64

const (
	ageEgg petAge = iota
	ageBaby

	ageIota petAge = 0.001
)

var (
	ageFrame1 = frameEggDown
	ageFrame2 = frameEggUp
)

func (p *pet) runAging(u *ui) {
	if p.age < 1 {
		p.tickUntil(ageIota * 6 * 5)
		p.age++
	}
	ageFrame1 = frameBaby1
	ageFrame2 = frameBaby2

	go func() {
		for { // TODO remove this for real lifecycle
			delay := rand.IntN(15 * 60)
			time.Sleep(time.Second * time.Duration(delay))

			fyne.Do(func() {
				p.dirty = true
				if p.alert {
					return
				}

				u.alert()
			})
		}
	}()

	// TODO end of life
}

func (p *pet) tickUntil(age petAge) {
	delta := age - p.age
	if delta <= 0 {
		return
	}

	durationDelta := time.Duration(delta/ageIota) * 10 * time.Second
	p.tickFor(durationDelta)
}

func (p *pet) tickFor(d time.Duration) {
	lapsed := time.Duration(0)
	for lapsed < d {
		time.Sleep(time.Second * 10)
		lapsed += time.Second * 10
		p.age += 0.001

		p.pref.SetFloat(keyAge, float64(p.age))
	}
}
