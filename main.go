package main

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var (
	//go:embed "background.png"
	imageBackground []byte
	//go:embed "egg.png"
	imageEgg []byte
)

func main() {
	a := app.NewWithID("xyz.andy.fynepet")
	w := a.NewWindow("Fyne Pet")

	bg := canvas.NewImageFromReader(bytes.NewReader(imageBackground), "background.png")
	egg := canvas.NewImageFromReader(bytes.NewReader(imageEgg), "egg.png")
	t := canvas.NewRaster(draw(&pix))

	b1 := newButton(func() {
		log.Println("Tap A")
	})
	b2 := newButton(func() {
		log.Println("Tap B")
	})
	b3 := newButton(func() {
		log.Println("Tap C")
	})

	t.ScaleMode = canvas.ImageScalePixels
	w.SetContent(container.New(&fullLayout{}, bg, container.New(&screenLayout{}, t), egg,
		container.New(&buttonLayout{}, b1, b2, b3)))

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)

			if i == 1 {
				pix = sleepLight1
			} else {
				pix = sleepLight2
			}
			fyne.Do(t.Refresh)

			i++
			if i > 1 {
				i = 0
			}
		}
	}()

	w.SetPadded(false)
	w.Resize(fyne.NewSize(680/2, 880/2))
	w.ShowAndRun()
}
