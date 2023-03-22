package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/d-tsuji/clipboard"
	"github.com/google/uuid"
	"github.com/marcusolsson/tui-go"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type tool struct {
	name string
	task func() (string, error)
}

func main() {
	faker := gofakeit.New(time.Now().Unix())

	tools := []tool{
		{
			name: "UUID",
			task: func() (string, error) {
				uuid := uuid.NewString()

				return uuid, clipboard.Set(uuid)
			},
		},
		{
			name: "Nano ID",
			task: func() (string, error) {
				nanoid := gonanoid.Must()

				return nanoid, clipboard.Set(nanoid)
			},
		},
		{
			name: "Date time (UTC)",
			task: func() (string, error) {
				dt := faker.Date().UTC().Format(time.RFC3339)

				return dt, clipboard.Set(dt)
			},
		},
		{
			name: "Email",
			task: func() (string, error) {
				email := faker.Email()

				return email, clipboard.Set(email)
			},
		},
		{
			name: "Full name",
			task: func() (string, error) {
				p := faker.Person()

				fullName := fmt.Sprintf("%s %s", p.FirstName, p.LastName)

				return fullName, clipboard.Set(fullName)
			},
		},
		{
			name: "Username",
			task: func() (string, error) {
				username := faker.Username()

				return username, clipboard.Set(username)
			},
		},
		{
			name: "Phone number",
			task: func() (string, error) {
				phoneNumber := faker.Phone()

				return phoneNumber, clipboard.Set(phoneNumber)
			},
		},
		{
			name: "Credit card",
			task: func() (string, error) {
				creditCard := faker.CreditCardNumber(&gofakeit.CreditCardOptions{
					Types: []string{"visa", "mastercard"},
					Gaps:  true,
				})

				return creditCard, clipboard.Set(creditCard)
			},
		},
		{
			name: "Company",
			task: func() (string, error) {
				company := faker.Company()

				return company, clipboard.Set(company)
			},
		},
		{
			name: "Phrase",
			task: func() (string, error) {
				phrase := faker.Phrase()

				return phrase, clipboard.Set(phrase)
			},
		},
	}

	generators := tui.NewTable(0, 0)
	// library.SetColumnStretch(0, 1)
	// library.SetColumnStretch(1, 1)
	// library.SetColumnStretch(2, 4)

	generators.SetFocused(true)

	generators.AppendRow(
		tui.NewLabel("What do you want to generate?"),
	)

	for _, tool := range tools {
		label := tui.NewLabel(tool.name)

		generators.AppendRow(
			label,
		)
	}

	currentClipboardText, err := clipboard.Get()
	if err != nil {
		panic(err)
	}

	statusLeft, statusRight := tui.NewStatusBar(""), tui.NewStatusBar("")

	// This implies a weird behavior that pushes the text to the right side of the HBox
	statusRight.SetPermanentText("Ninym Ralei the best possible girl")

	setClipboardStatus := func(s string) { statusLeft.SetText(fmt.Sprintf("Clipboard: %s", s)) }
	setClipboardStatus(currentClipboardText)

	statuses := tui.NewHBox(statusLeft, statusRight)
	statuses.SetBorder(true)

	root := tui.NewVBox(
		generators,
		tui.NewSpacer(),
		statuses,
	)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	generators.OnItemActivated(func(t *tui.Table) {
		if t.Selected() == 0 {
			return
		}

		out, err := tools[t.Selected()-1].task()
		if err != nil {
			panic(err)
		}

		setClipboardStatus(out)
	})

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("q", func() { ui.Quit() })

	theme := tui.NewTheme()
	theme.SetStyle("table.cell.selected", tui.Style{
		Bold: tui.DecorationOn, Reverse: tui.DecorationOn,
		Fg: tui.ColorBlue, Bg: tui.ColorWhite,
	})

	ui.SetTheme(theme)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
