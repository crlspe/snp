package ui

import (
	"github.com/rivo/tview"
)

func GitHubCredentials(user, repository, token string) (string, string, string) {
	var app = StartForm()

	form := tview.NewForm()
	form.AddInputField("User", user, 30, nil, nil)
	form.AddInputField("Repository", repository, 30, nil, nil)
	form.AddInputField("Token", token, 0, nil, nil)

	form.AddButton("Submit", func() {
		user = form.GetFormItemByLabel("User").(*tview.InputField).GetText()
		repository = form.GetFormItemByLabel("Repository").(*tview.InputField).GetText()
		token = form.GetFormItemByLabel("Token").(*tview.InputField).GetText()
		app.Stop()
	})

	form.SetBorder(false).SetTitle("GitHub Credentials").SetTitleAlign(tview.AlignCenter)
	RunForm(app, form)

	return user, repository, token
}
