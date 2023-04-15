package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/luisnquin/mocktail/internal/clipboard"
	"github.com/luisnquin/mocktail/internal/faker"
	"github.com/rivo/tview"
)

type tool struct {
	name      string
	generator func() string
}

const (
	DefaultStatusRightTitle = "A truth"
	DefaultStatusRightLabel = "Ninym Ralei the best girl"
)

func main() {
	tools := []tool{
		{
			name:      "UUID",
			generator: faker.UUID,
		},
		{
			name:      "Nano ID",
			generator: faker.NanoID,
		},
		{
			name:      "Date time (UTC)",
			generator: faker.DateTime,
		},
		{
			name:      "Email",
			generator: faker.Email,
		},
		{
			name:      "Full name",
			generator: faker.FullName,
		},
		{
			name:      "Username",
			generator: faker.Username,
		},
		{
			name:      "Phone number",
			generator: faker.PhoneNumber,
		},
		{
			name:      "Credit card",
			generator: faker.CreditCardNumber,
		},
		{
			name:      "Lorem sentence",
			generator: faker.LoremSentence,
		},
		{
			name:      "Postal code",
			generator: faker.PostalCode,
		},
		{
			name:      "City",
			generator: faker.City,
		},
		{
			name:      "Address",
			generator: faker.Address,
		},
		{
			name:      "Hexadecimal color",
			generator: faker.HexColor,
		},
	}

	app := tview.NewApplication().EnableMouse(true)

	list := tview.NewList()
	statusLeft, statusRight := tview.NewTextView(), tview.NewTextView()

	for i, t := range tools {
		shortcut := rune(int32(i + 1))

		generator := t.generator

		list.AddItem(t.name, generator(), shortcut, func() {
			text := generator()

			statusLeft.Lock()
			statusLeft.SetLabel(text)
			statusLeft.Unlock()

			if err := clipboard.Set(text); err != nil {
				panic(err) // TODO: improve error handling
			}
		})
	}

	statusLeft.SetTitle("Clipboard").SetBorder(true)
	statusLeft.SetLabelWidth(90)
	// statusLeft.SetBackgroundColor(tcell.ColorTeal)

	statusRight.SetTitle(DefaultStatusRightTitle).SetBorder(true)
	statusRight.SetBorderPadding(0, 0, 1, 0)

	clipboardText, err := clipboard.Get()
	if err != nil {
		panic(err)
	}

	grid := tview.NewGrid().
		AddItem(list, 0, 0, 7, 2, 0, 0, true).
		AddItem(statusLeft.SetLabel(clipboardText), 7, 0, 1, 1, 0, 0, false).
		AddItem(statusRight.SetLabel(DefaultStatusRightLabel), 7, 1, 1, 1, 0, 0, false)

	grid.SetGap(1, 1).SetTitle("Main")

	go seekClipboardForChanges(statusLeft, statusRight)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func seekClipboardForChanges(statusLeft, statusRight *tview.TextView) {
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
			statusRight.SetTitle(DefaultStatusRightTitle)
			statusRight.SetLabel(DefaultStatusRightLabel)
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
