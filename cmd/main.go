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

var baseIndex int
func white_button(text string) *fyne.Container {
	index := 0
	var btn *widget.Button
	var btn_color *canvas.Rectangle
	var container1 *fyne.Container
	btn = widget.NewButton(text, func() {
		fmt.Println("---------------")
		fmt.Println(index, baseIndex)
		if index < baseIndex {
			if index == 0 {
				index = baseIndex
				baseIndex++
			}
		} else {
			baseIndex++
		}
		fmt.Println(index, baseIndex)
		fmt.Println("---------------")
		if btn_color.FillColor == color.White {
			btn_color.FillColor = color.NRGBA{R: 123, G: 123, B: 123, A: 123}
		} else {
			btn_color.FillColor = color.White
		}
		container1.Refresh()
	})
	btn.Alignment = widget.ButtonAlignLeading
	btn_color = canvas.NewRectangle(color.White)

	container1 = container.New(
		layout.NewMaxLayout(), btn_color, btn)
	return container1
}

func main() {
	var boxs *fyne.Container
	baseIndex = 0
	boxs = container.NewVBox()

	font := "assets/fonts/STHeiti Light.ttc"

	os.Setenv("FYNE_FONT", font)

	a := app.New()
	window := a.NewWindow("操作票排序")

	window.Resize(fyne.NewSize(1600, 900))

	filesView := widget.NewList(
		func() int {
			return len(resource10kVTxt)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(resource10kVTxt[id].StaticName)
			// object.(*widget.Button).SetText("tewwefew")

		})

	// contentText := widget.NewLabel("")

	// var ticket []string
	// contentText := widget.NewButton("test", nil)
	//

	right := container.NewVBox()
	filesView.OnSelected = func(id widget.ListItemID) {
		baseIndex = 0
		boxs.RemoveAll()
		right.RemoveAll()

		// sources.SetText(string(resource10kVTxt[id].StaticContent))
		sources := strings.Split(string(resource10kVTxt[id].StaticContent), "\n")
		ticket := strings.Split(string(resource10kVTxt[id].StaticContent), "\n")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(ticket), func(i, j int) {
			ticket[i], ticket[j] = ticket[j], ticket[i]
		})

		for i, v := range ticket {
			boxs.Add(white_button(v))
			right.Add(white_button(fmt.Sprintf("%d. %s", i + 1, sources[i])))
		}

		// context := strings.Join(ticket, "\n")
		// contentText.SetText(context)

	}

	filesView.Select(0)



	check := container.NewHSplit(boxs, right)

	split := container.NewHSplit(
		filesView,
		check,
	)

	split.Offset = 0.2


	window.SetContent(split)
	window.ShowAndRun()

}
