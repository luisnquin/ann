package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/luisnquin/mocktail/internal/clipboard"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rivo/tview"
)

type tool struct {
	name        string
	description string
	task        func()
}

func main() {
	faker := gofakeit.New(time.Now().Unix())

	statusLeft, statusRight := tview.NewTextView(), tview.NewTextView()

	updateClipboardAndStatus := func(s string) func() {
		return func() {
			statusLeft.Lock()
			statusLeft.SetLabel(s)
			statusLeft.Unlock()

			if err := clipboard.Set(s); err != nil {
				panic(err) // TODO: improve error handling
			}
		}
	}

	tools := []tool{
		{
			name:        "UUID",
			description: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			task:        updateClipboardAndStatus(uuid.NewString()),
		},
		{
			name:        "Nano ID",
			description: "PPPPPPP-CCCCC",
			task:        updateClipboardAndStatus(gonanoid.Must()),
		},
		{
			name:        "Date time (UTC)",
			description: time.RFC3339,
			task:        updateClipboardAndStatus(faker.Date().UTC().Format(time.RFC3339)),
		},
		{
			name:        "Email",
			description: "example@mail.org",
			task:        updateClipboardAndStatus(faker.Email()),
		},
		{
			name:        "Full name",
			description: "John Doe",
			task:        updateClipboardAndStatus(fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName())),
		},
		{
			name:        "Username",
			description: "guest256",
			task:        updateClipboardAndStatus(faker.Username()),
		},
		{
			name:        "Phone number",
			description: "##########",
			task:        updateClipboardAndStatus(faker.Phone()),
		},
		{
			name:        "Credit card",
			description: "5370 1234 5678 9012",
			task: updateClipboardAndStatus(faker.CreditCardNumber(&gofakeit.CreditCardOptions{
				Types: []string{"visa", "mastercard"},
				Gaps:  true,
			})),
		},
		{
			name:        "Phrase",
			description: "How's it going?",
			task:        updateClipboardAndStatus(faker.Phrase()),
		},
	}

	app := tview.NewApplication().EnableMouse(true)

	list := tview.NewList()

	for i, t := range tools {
		shortcut := rune(int32(i + 1))

		localTask := t.task

		list.AddItem(t.name, t.description, shortcut, func() {
			localTask()
		})
	}

	statusLeft.SetTitle("Clipboard").SetBorder(true)
	statusLeft.SetLabelWidth(90)
	statusRight.SetBorder(true)

	clipboardText, err := clipboard.Get()
	if err != nil {
		panic(err)
	}

	grid := tview.NewGrid().
		AddItem(list, 0, 0, 7, 2, 0, 0, true).
		AddItem(statusLeft.SetLabel(clipboardText), 7, 0, 1, 1, 0, 0, false).
		AddItem(statusRight.SetLabel("Ninym Ralei the best girl"), 7, 1, 1, 1, 0, 0, false)

	grid.SetGap(1, 1).SetTitle("Main")

	go seekClipboardForChanges(statusLeft)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func seekClipboardForChanges(statusLeft *tview.TextView) {
	t := time.NewTicker(time.Second)

	statusLeft.Lock()
	lastClipText := statusLeft.GetLabel()
	statusLeft.Unlock()

	for range t.C {
		clipContent, err := clipboard.Get()
		if err != nil {
			panic(err)
		}

		if clipContent != lastClipText {
			lastClipText = clipContent

			statusLeft.Lock()
			statusLeft.SetLabel(lastClipText)
			statusLeft.Unlock()
		}
	}
}
