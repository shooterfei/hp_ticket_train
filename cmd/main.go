package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func white_button(text string) *fyne.Container {
	var btn *widget.Button
	var btn_color *canvas.Rectangle
	var container1 *fyne.Container
	btn = widget.NewButton(text, func() {
		fmt.Println("Here is Go Button")
		fmt.Println(btn_color.FillColor)
		btn_color.FillColor = color.Black
		container1.Refresh()
	})
	btn_color = canvas.NewRectangle(color.White)

	container1 = container.New(
		layout.NewMaxLayout(), btn_color, btn)
	return container1
}

func main() {
	font := "assets/fonts/STHeiti Light.ttc"

	os.Setenv("FYNE_FONT", font)

	a := app.New()
	window := a.NewWindow("操作票排序")

	window.Resize(fyne.NewSize(1200, 800))

	filesView := widget.NewList(
		func() int {
			return len(resource10kVTxt)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
			// return widget.NewButton("test", func() {
			// 	fmt.Println("test")
			// })
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(resource10kVTxt[id].StaticName)
			// object.(*widget.Button).SetText("tewwefew")

		})

	contentText := widget.NewLabel("")

	// contentText := widget.NewButton("test", nil)
	filesView.OnSelected = func(id widget.ListItemID) {

		ticket := strings.Split(string(resource10kVTxt[id].StaticContent), "\n")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(ticket), func(i, j int) {
			ticket[i], ticket[j] = ticket[j], ticket[i]
		})

		context := strings.Join(ticket, "\n")
		contentText.SetText(context)

	}

	filesView.Select(0)

	split := container.NewHSplit(
		filesView,
		// container.NewMax(contentText),
		container.NewVBox(
			white_button("123dvhihv"),
			white_button("123dvhihv"),
			white_button("123dvhihv"),
			white_button("123dvhihv"),
		),
	)

	split.Offset = 0.2

	// window.SetContent(container.NewVBox(
	// 	red_button(),
	// 	red_button(),
	// 	red_button(),
	// 	red_button(),
	// 	red_button(),
	// 	red_button(),
	// ))
	window.SetContent(split)
	window.ShowAndRun()

}