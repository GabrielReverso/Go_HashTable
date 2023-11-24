package NameList

import (
	"HashTable/components/HashTable"
	"HashTable/components/MakeScreens"
	"fmt"
	"image/color"
	"math/rand"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type buttonLayout struct {
}

func (l *buttonLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 40) // set the size of the button
}

func (l *buttonLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	objects[0].Resize(fyne.NewSize(40, 40))  // circle
	objects[1].Resize(objects[1].MinSize())  // label
	objects[2].Resize(fyne.NewSize(700, 40)) // button

	objects[0].Move(fyne.NewPos(15, 0))
	objects[1].Move(fyne.NewPos(26, 3))
	objects[2].Move(fyne.NewPos(75, 0))
}

func (l *buttonLayout) DesiredSize(objects []fyne.CanvasObject, size fyne.Size) fyne.Size {
	return size
}

var colors = []color.RGBA{
	{R: 0, G: 0, B: 150, A: 255},
	{R: 0, G: 0, B: 200, A: 255},
	{R: 0, G: 120, B: 200, A: 255},
	{R: 30, G: 144, B: 255, A: 255},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomColor() color.RGBA {
	return colors[rand.Intn(len(colors))]
}

func CreateButtons(w fyne.Window, hash_table *HashTable.Hash, data []string) *fyne.Container {
	objects := make([]fyne.CanvasObject, len(data)*2-1)

	for i, str := range data {

		parts := strings.Split(str, ",")
		if i > 0 {
			parts[0] = fmt.Sprintf("%s (%d)", parts[0], i+1)
		}

		nome := parts[0]
		telefone := parts[1]
		endereco := parts[2]

		button := widget.NewButton(nome, func() { MakeScreens.MakeDataScreen(w, hash_table, nome, telefone, endereco, true) })

		circle := canvas.NewCircle(randomColor())
		circle.StrokeWidth = 0
		circle.StrokeColor = color.Black

		label := canvas.NewText(strings.ToUpper(string(str[0])), color.White)
		label.TextSize = 24
		label.TextStyle.Bold = true

		container := fyne.NewContainerWithLayout(&buttonLayout{}, circle, label, button)
		objects[i*2] = container

		if i < len(str)-1 {
			spacer := canvas.NewRectangle(color.Transparent)
			spacer.SetMinSize(fyne.NewSize(1, 3)) // 3 pixels of height
			objects[i*2+1] = spacer
		}
	}

	return container.NewVBox(objects...)
}
