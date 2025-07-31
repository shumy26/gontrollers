package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	vari := 0.0
	img := canvas.NewImageFromFile("pid_plot.png")
	img.Refresh()
	img.FillMode = canvas.ImageFillOriginal
	img.ScaleMode = canvas.ImageScaleSmooth

	a := app.New()
	w := a.NewWindow("Update Time")

	message := widget.NewLabel("Welcome")
	button := widget.NewButton("Update", func() {
		vari += 10
		ch := make(chan float64)
		go control(ch, vari)
		plotting(ch)
		img.Refresh()
	})

	w.SetContent(container.NewVBox(message, button, img))
	w.ShowAndRun()
}
