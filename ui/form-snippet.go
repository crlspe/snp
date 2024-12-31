package ui

import "github.com/rivo/tview"

func SnippetAdd(code, description, scopes, tags string) (string, string, string, string) {
	var app = StartForm()

	form := tview.NewForm()
	form.AddTextArea("Content", code, 0, 0, 0, nil)
	form.AddInputField("Description", description, 0, nil, nil)
	form.AddInputField("Scopes", scopes, 0, nil, nil)
	form.AddInputField("Tags", tags, 0, nil, nil)

	form.AddButton("Submit", func() {
		description = form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
		scopes = form.GetFormItemByLabel("Scopes").(*tview.InputField).GetText()
		tags = form.GetFormItemByLabel("Tags").(*tview.InputField).GetText()
		code = form.GetFormItemByLabel("Content").(*tview.TextArea).GetText()
		app.Stop()
	})

	form.SetBorder(false).SetTitle("Add Snippet").SetTitleAlign(tview.AlignCenter)
	RunForm(app, form)

	return code, description, scopes, tags
}
