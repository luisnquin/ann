package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/d-tsuji/clipboard"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/rivo/tview"
)

type tool struct {
	name        string
	description string
	task        func() error
}

func main() {
	faker := gofakeit.New(time.Now().Unix())

	statusLeft, statusRight := tview.NewTextView(), tview.NewTextView()

	updateClipboardAndStatus := func(s string) error {
		statusLeft.SetLabel(s)

		return clipboard.Set(s)
	}

	tools := []tool{
		{
			name:        "UUID",
			description: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			task: func() error {
				return updateClipboardAndStatus(uuid.NewString())
			},
		},
		{
			name:        "Nano ID",
			description: "PPPPPPP-CCCCC",
			task: func() error {
				return updateClipboardAndStatus(gonanoid.Must())
			},
		},
		{
			name:        "Date time (UTC)",
			description: time.RFC3339,
			task: func() error {
				return updateClipboardAndStatus(faker.Date().UTC().Format(time.RFC3339))
			},
		},
		{
			name:        "Email",
			description: "example@mail.org",
			task: func() error {
				return updateClipboardAndStatus(faker.Email())
			},
		},
		{
			name:        "Full name",
			description: "John Doe",
			task: func() error {
				p := faker.Person()

				fullName := fmt.Sprintf("%s %s", p.FirstName, p.LastName)

				return updateClipboardAndStatus(fullName)
			},
		},
		{
			name:        "Username",
			description: "guest256",
			task: func() error {
				return updateClipboardAndStatus(faker.Username())
			},
		},
		{
			name:        "Phone number",
			description: "##########",
			task: func() error {
				return updateClipboardAndStatus(faker.Phone())
			},
		},
		{
			name:        "Credit card",
			description: "5370 1234 5678 9012",
			task: func() error {
				creditCard := faker.CreditCardNumber(&gofakeit.CreditCardOptions{
					Types: []string{"visa", "mastercard"},
					Gaps:  true,
				})

				return updateClipboardAndStatus(creditCard)
			},
		},
		{
			name:        "Phrase",
			description: "How's it going?",
			task: func() error {
				return updateClipboardAndStatus(faker.Phrase())
			},
		},
	}

	app := tview.NewApplication().EnableMouse(true)

	list := tview.NewList()

	for i, t := range tools {
		shortcut := rune(int32(i + 1))

		localTask := t.task

		list.AddItem(t.name, t.description, shortcut, func() {
			if err := localTask(); err != nil {
				panic(err) // TODO: improve error handling
			}
		})
	}

	statusLeft.SetTitle("Clipboard").SetBorder(true)
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

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
