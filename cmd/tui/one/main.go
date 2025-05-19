package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts []Contact

var app = tview.NewApplication()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(q) to quit")

func main() {
	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
