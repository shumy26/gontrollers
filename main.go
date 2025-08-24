package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shumy26/gontrollers/control"
	"github.com/shumy26/gontrollers/plot"
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
		go control.Control(ch, vari)
		plot.Plotting(ch)
		img.Refresh()
	})

	w.SetContent(container.NewVBox(message, button, img))

	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			fyne.Do(func() {
				vari += 10
				ch := make(chan float64)
				go control.Control(ch, vari)
				plot.Plotting(ch)
				img.Refresh()
			})
		}
	}()

	w.ShowAndRun()

}
