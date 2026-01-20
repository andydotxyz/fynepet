package main

import (
	"fyne.io/fyne/v2"
)

var petMinSize = fyne.NewSize(340, 440)

type buttonLayout struct{}

func (l buttonLayout) Layout(objects []fyne.CanvasObject, s fyne.Size) {
	midWidth := (s.Width * .99) / 2 // image is slightly left aligned

	halfBSize := float32(21)
	bSize := fyne.NewSquareSize(halfBSize * 2)
	bTop := s.Height * .758
	b1 := objects[0]
	b1.Resize(bSize)
	b1.Move(fyne.NewPos(s.Width*.315-halfBSize, bTop))
	b2 := objects[1]
	b2.Resize(bSize)
	b2.Move(fyne.NewPos(midWidth-halfBSize, bTop+s.Height*.02))
	b3 := objects[2]
	b3.Resize(bSize)
	b3.Move(fyne.NewPos(s.Width*.68-halfBSize, bTop))
}

func (l buttonLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return petMinSize
}

type fullLayout struct{}

func (l fullLayout) Layout(objects []fyne.CanvasObject, s fyne.Size) {
	ratio := petMinSize.Width / petMinSize.Height
	winRatio := s.Width / s.Height

	offset := fyne.Position{}
	inner := s
	if ratio <= winRatio {
		inner.Width = s.Height * ratio
		offset.X = (s.Width - inner.Width) / 2
	} else if ratio >= winRatio {
		inner.Height = s.Width / ratio
		offset.Y = (s.Height - inner.Height) / 2
	}

	for _, o := range objects {
		o.Resize(inner)
		o.Move(offset)
	}
}

func (l fullLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return petMinSize
}

type screenLayout struct{}

func (l screenLayout) Layout(objects []fyne.CanvasObject, s fyne.Size) {
	midWidth, midHeight := s.Width/2, s.Height/2
	screen := objects[0]
	screen.Resize(l.MinSize(objects))
	screen.Move(fyne.NewPos(midWidth-82, midHeight-44))
}

func (l screenLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(160, 80)
}
