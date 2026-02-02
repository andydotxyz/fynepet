package main

import (
	"log"
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
	ageDead petAge = 99

	ageIota petAge = 0.006
)

var ageFrames = map[int][2][16]int64{
	int(ageEgg):  {frameEggDown, frameEggUp},
	int(ageBaby): {frameBaby1, frameBaby2},
	int(ageDead): {frameDead1, frameDead2},
}

func (p *pet) runAging(u *ui) {
	if p.age < 1 {
		p.tickUntil(ageIota * 5)
		p.age = 1
	}

	go func() {
		for { // TODO remove this for real lifecycle
			delay := rand.IntN(15 * 60)
			time.Sleep(time.Second * time.Duration(delay))
			if p.age < ageBaby || p.age >= ageDead {
				continue
			}

			fyne.Do(func() {
				p.dirty = true
				if p.alert {
					return
				}

				u.alert()
			})
		}
	}()

	if p.age >= ageBaby && p.age < ageDead {
		p.tickUntil(1.0 + ageIota*30)
		p.age = ageDead // dead
		p.pref.SetFloat(keyAge, float64(p.age))
	}
	// TODO dead stats page
}

func (p *pet) tickUntil(age petAge) {
	delta := age - p.age
	if delta <= 0 {
		return
	}
	for delta > 1 {
		// more than a life stage difference
		log.Println("Unsure how long until next generation")
	}

	durationDelta := time.Duration(delta/ageIota) * 60 * time.Second
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
