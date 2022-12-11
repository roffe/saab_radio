package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/roffe/saab_radio/assets"
	"github.com/roffe/saab_radio/ui"
)

func main() {
	app := app.NewWithID("com.github.9-5radio")
	app.SetIcon(fyne.NewStaticResource("saab-logo.png", assets.LogoBytes))
	ui.NewMainWindow(app)
}
