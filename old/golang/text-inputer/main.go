package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("text-inputter")

	w.Icon()
	input := widget.NewMultiLineEntry()

	w.SetContent(input)
	w.Resize(fyne.NewSize(350, 350))
	w.ShowAndRun()
}
