package command

import (
	"os"

	"github.com/crlspe/snp/model"
)

type Command interface {
	Init(flags model.Flags)
	Exec()
}

type CommandName string

const (
	Search CommandName = "search"
	Add    CommandName = "add"
	Update CommandName = "update"
	Config   CommandName = "config"
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
