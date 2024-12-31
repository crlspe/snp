package command

import (
	"github.com/crlspe/snp/filesystem"
	"github.com/crlspe/snp/model"
	"github.com/crlspe/snp/settings"
	forms "github.com/crlspe/snp/ui"

	"golang.design/x/clipboard"
)

type AddCommand struct {
	Flags   model.Flags
	Snippet model.Snippet
}

func (a *AddCommand) Init(flags model.Flags) {
	a.Flags = flags
}

func (a *AddCommand) Exec() {
	a.parseFlags()
	a.askContents()
	filesystem.CreateDirectoryIfNotExist(settings.Config.DataFolder)
	a.Snippet.Save()
}

func (a *AddCommand) parseFlags() {
	switch {
	case *a.Flags.Clipboard:
		var content = string(clipboard.Read(clipboard.FmtText))
		a.Snippet.Code = content
	}
}

func (a *AddCommand) askContents() {
	a.Snippet.Code, a.Snippet.Description, a.Snippet.Scopes, a.Snippet.Tags = forms.SnippetAdd(a.Snippet.Code, a.Snippet.Description, a.Snippet.Scopes, a.Snippet.Tags)
}
