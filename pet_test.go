package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestFeedBurger(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()

	p := loadPet(a.Preferences())
	p.age = ageBaby
	p.hungry = 3
	p.happy = 1

	p.feed(true) // burger: hunger -2, happy +1

	if p.hungry != 1 {
		t.Errorf("expected hungry=1 after burger, got %d", p.hungry)
	}
	if p.happy != 2 {
		t.Errorf("expected happy=2 after burger, got %d", p.happy)
	}
}

func TestFeedCake(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()

	p := loadPet(a.Preferences())
	p.age = ageBaby
	p.hungry = 2
	p.happy = 1

	p.feed(false) // cake: hunger -1, happy +2

	if p.hungry != 1 {
		t.Errorf("expected hungry=1 after cake, got %d", p.hungry)
	}
	if p.happy != 3 {
		t.Errorf("expected happy=3 after cake, got %d", p.happy)
	}
}

func TestFeedBurgerClampsHungryAtZero(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()

	p := loadPet(a.Preferences())
	p.age = ageBaby
	p.hungry = 1
	p.happy = 0

	p.feed(true) // burger: hunger -2, but clamped to 0

	if p.hungry != 0 {
		t.Errorf("expected hungry=0 (clamped), got %d", p.hungry)
	}
}

func TestFeedCakeClampsHappyAtMax(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()

	p := loadPet(a.Preferences())
	p.age = ageBaby
	p.hungry = 2
	p.happy = maxStat - 1

	p.feed(false) // cake: happy +2, but clamped to maxStat

	if p.happy != maxStat {
		t.Errorf("expected happy=%d (clamped), got %d", maxStat, p.happy)
	}
}

func TestFeedPersistsToPreferences(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()

	pref := a.Preferences()
	p := loadPet(pref)
	p.age = ageBaby
	p.hungry = 3
	p.happy = 1

	p.feed(true)

	if pref.Int(keyHungry) != 1 {
		t.Errorf("expected persisted hungry=1, got %d", pref.Int(keyHungry))
	}
	if pref.Int(keyHappy) != 2 {
		t.Errorf("expected persisted happy=2, got %d", pref.Int(keyHappy))
	}
}
