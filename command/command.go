package command

import (
	"os"

	"github.com/crlspe/snp/github"
	"github.com/crlspe/snp/model"
	"github.com/crlspe/snp/settings"
)

func init() {
	settings.GitHubClient = github.NewGitHub(
		settings.Config.GitHub.User,
		settings.Config.GitHub.Repository,
		settings.Config.GitHub.Token,
		nil,
	)
}

type Command interface {
	Init(flags model.Flags)
	Exec()
}

type CommandName string

const (
	Search CommandName = "search"
	Add    CommandName = "add"
	Update CommandName = "update"
	Config CommandName = "config"
)

func GetCommand() Command {
	switch os.Args[1] {
	case string(Search):
		return &SearchCommand{}
	case string(Add):
		return &AddCommand{}
	case string(Update):
		return &UpdateCommand{}
	case string(Config):
		return &ConfCommand{}
	default:
		return &SearchCommand{}
	}
}
