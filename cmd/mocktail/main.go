package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/luisnquin/mocktail/internal/clipboard"
	"github.com/rivo/tview"
)

type config struct {
	StatusRight struct {
		Title string
		Label string
	}
}

func getConfig() config {
	var c config

	c.StatusRight.Title = "A truth"
	c.StatusRight.Label = getRandomQuote()

	return c
}

func main() {
	app := tview.NewApplication().EnableMouse(true)
	config := getConfig()

	list := tview.NewList()
	statusLeft, statusRight := tview.NewTextView(), tview.NewTextView()

	for i, g := range getGenerators() {
		shortcut := rune(int32(i + 1))

		generator := g.generator

		list.AddItem(g.name, generator(), shortcut, func() {
			text := generator()

			statusLeft.Lock()
			statusLeft.SetLabel(text)
			statusLeft.Unlock()

			if err := clipboard.Set(text); err != nil {
				panic(err)
			}
		})
	}

	statusLeft.SetTitle("Clipboard").SetBorder(true)
	statusLeft.SetLabelWidth(90)
	// statusLeft.SetBackgroundColor(tcell.ColorTeal)

	statusRight.SetTitle(config.StatusRight.Title).SetBorder(true)
	statusRight.SetBorderPadding(0, 0, 1, 0)

	clipboardText, err := clipboard.Get()
	if err != nil {
		panic(err)
	}

	grid := tview.NewGrid().
		AddItem(list, 0, 0, 7, 2, 0, 0, true).
		AddItem(statusLeft.SetLabel(clipboardText), 7, 0, 1, 1, 0, 0, false).
		AddItem(statusRight.SetLabel(config.StatusRight.Label), 7, 1, 1, 1, 0, 0, false)

	grid.SetGap(1, 1).SetTitle("Main")

	go seekClipboardForChanges(config, statusLeft, statusRight)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func seekClipboardForChanges(config config, statusLeft, statusRight *tview.TextView) {
	t := time.NewTicker(time.Second)

	statusLeft.Lock()
	lastClipText := statusLeft.GetLabel()
	statusLeft.Unlock()

	for range t.C {
		clipContent, err := clipboard.Get()
		if err != nil {
			statusRight.Lock()
			statusRight.SetBorderColor(tcell.Color196)
			statusRight.SetTitle("Error")
			statusRight.SetLabel(fmt.Sprintf("from refresh func: %s", err.Error()))
			statusRight.Unlock()

			time.Sleep(time.Second * 3)

			statusRight.Lock()
			statusRight.SetBorderColor(tcell.ColorWhite)
			statusRight.SetTitle(config.StatusRight.Title)
			statusRight.SetLabel(config.StatusRight.Label)
			statusRight.Unlock()

			time.Sleep(time.Second)

			continue
		}

		if clipContent != lastClipText {
			lastClipText = clipContent

			statusLeft.Lock()
			statusLeft.SetLabel(lastClipText)
			statusLeft.Unlock()
		}
	}
}
