package ui

import (
	"log"

	"github.com/rivo/tview"
)

func StartForm() *tview.Application {
	return tview.NewApplication()
}

func RunForm(app *tview.Application, form *tview.Form) {
	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		log.Fatal(err)
	}
}
