package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
	fyne.Theme
	test fyne.Resource
}

func NewTheme() *myTheme {
	file, _ := os.ReadFile("/usr/share/fonts/truetype/ubuntu/UbuntuMono-BI.ttf") //TODO: do usunięcia po bundle
	// if error != nil {
	test := fyne.NewStaticResource("testResource", file)
	newTheme := &myTheme{theme.DefaultTheme(), test}
	// }
	return newTheme
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace { //TODO: użyć bundle na jakimś foncie i obsłużyć wszystko: monospace, bold, italics
		return m.test
	}
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name) * 2
}
