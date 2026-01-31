package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type lightThemeOverride struct {
	fyne.Theme
}

func (t lightThemeOverride) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}
