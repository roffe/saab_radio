package ui

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/roffe/saab_radio/radio"
	"github.com/roffe/saab_radio/vin"
	sdialog "github.com/sqweek/dialog"
)

type MainWindow struct {
	app fyne.App
	fyne.Window
	vin  *widget.Entry
	form *widget.Form
}

func NewMainWindow(app fyne.App) {
	mw := &MainWindow{
		app: app,
	}
	mw.vin = &widget.Entry{
		TextStyle:   fyne.TextStyle{Monospace: true},
		Wrapping:    fyne.TextWrapOff,
		PlaceHolder: strings.Repeat("", 17),
	}
	mw.vin.Validator = func(str string) error {
		if len(str) != 17 {
			return fmt.Errorf("VIN must be 17 characters")
		}
		ok, _ := vin.VinCheck(str)
		if !ok {
			return fmt.Errorf("invalid VIN")
		}
		return nil
	}
	mw.vin.OnChanged = func(s string) {
		if len(s) > 17 {
			mw.vin.SetText(s[:17])
		}
	}
	mw.form = &widget.Form{
		Items: []*widget.FormItem{
			widget.NewFormItem("VIN", mw.vin),
		},
		SubmitText: "Generate",
		OnSubmit: func() {
			filename, err := sdialog.File().Filter("Bin file", "bin").Title("Save eeprom binary").Save()
			if err != nil {
				if err.Error() == "Cancelled" {
					return
				}
				return
			}
			filename = addSuffix(filename, ".bin")
			codes, err := radio.GenerateCodes(mw.vin.Text[len(mw.vin.Text)-6:])
			if err != nil {
				dialog.ShowError(err, mw)
				return
			}
			b := radio.GenerateBin(codes)
			if err := os.WriteFile(filename, b, 0644); err != nil {
				dialog.ShowError(err, mw)
				return
			}
			sdialog.Message("EEPROM binary saved to " + filename).Title("Success").Info()
		},
	}

	mw.Window = app.NewWindow("Saab 9-5 Radio Tool")
	mw.Show()
	mw.SetContent(mw.layout())
	mw.Resize(fyne.NewSize(300, 80))
	mw.SetFixedSize(true)
	mw.ShowAndRun()
}

func (mw *MainWindow) layout() fyne.CanvasObject {
	return mw.form
}

func addSuffix(s, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}
	return s
}
