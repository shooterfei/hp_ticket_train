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
var total int
var sources []string
var ticket []string

func valid_button(text string) *fyne.Container {
	index := 0
	var btn *widget.Button
	var btn_color *canvas.Rectangle
	var container1 *fyne.Container
	btn = widget.NewButton(text, func() {

		if index < baseIndex {
			if index == 0 {
				index = baseIndex
				baseIndex++
			}
		} else {
			baseIndex++
		}
		if text == sources[index-1] {
			btn_color.FillColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		} else {
			btn_color.FillColor = color.NRGBA{R: 123, G: 123, B: 123, A: 255}
		}

		btn.SetText(fmt.Sprintf("%d. %s", index, text))
		container1.Refresh()
	})
	btn.Alignment = widget.ButtonAlignLeading
	btn_color = canvas.NewRectangle(color.White)

	container1 = container.New(
		layout.NewMaxLayout(), btn_color, btn)
	return container1
}

func white_button(text string) *fyne.Container {
	var btn *widget.Button
	var btn_color *canvas.Rectangle
	var container1 *fyne.Container
	btn = widget.NewButton(text, func() {})
	btn.Alignment = widget.ButtonAlignLeading
	btn_color = canvas.NewRectangle(color.White)

	container1 = container.New(
		layout.NewMaxLayout(), btn_color, btn)
	return container1
}

func main() {
	var sorts *fyne.Container

	baseIndex = 0
	sorts = container.NewVBox()

	font := "assets/fonts/STHeiti Light.ttc"

	os.Setenv("FYNE_FONT", font)

	a := app.New()
	window := a.NewWindow("操作票排序")

	window.Resize(fyne.NewSize(1600, 900))

	files := widget.NewList(
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

	right := container.NewVBox()

	files.OnSelected = func(id widget.ListItemID) {
		baseIndex = 1
		sorts.RemoveAll()
		right.RemoveAll()

		sources = strings.Split(string(resource10kVTxt[id].StaticContent), "\n")
		ticket = strings.Split(string(resource10kVTxt[id].StaticContent), "\n")
		total = len(sources)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(ticket), func(i, j int) {
			ticket[i], ticket[j] = ticket[j], ticket[i]
		})

		for i, v := range ticket {
			sorts.Add(valid_button(v))
			right.Add(white_button(fmt.Sprintf("%d. %s", i+1, sources[i])))
		}
	}

	files.Select(0)

	check := container.NewHSplit(sorts, right)

	var left *fyne.Container
	var split *container.Split
	var btn1 *widget.Button
	var btn2 *widget.Button

	btn1 = widget.NewButton("填写", func() {
		split.Trailing = sorts
		split.Refresh()
		btn1.Disable()
		btn2.Enable()

	})
	btn1.Disable()

	btn2 = widget.NewButton("对比", func() {
		split.Trailing = check
		split.Refresh()
		btn2.Disable()
		btn1.Enable()

	})
	tools := container.NewGridWithColumns(2, btn1, btn2)

	left = container.NewBorder(nil, tools, nil, nil, files)
	split = container.NewHSplit(left, sorts)
	split.Offset = 0.2

	window.SetContent(split)
	window.ShowAndRun()

}
