package command

import (
	"fmt"

	"github.com/crlspe/snp/config"
	"github.com/crlspe/snp/filesystem"
	"github.com/crlspe/snp/model"
	"github.com/crlspe/snp/settings"
)

type ConfCommand struct {
	Flags model.Flags
}

func (i *ConfCommand) Init(flags model.Flags) {
	i.Flags = flags
}

func (i *ConfCommand) Exec() {
	switch {
	case *i.Flags.Local:
		InitLocalConfiguration()
	case *i.Flags.Github:
		InitLocalConfiguration()
		InitGitHubConfiguration()
	case *i.Flags.DryRun:
		PrintConfigurations()
	}
}

func InitLocalConfiguration() {
	filesystem.CreateDirectoryIfNotExist(settings.Config.ConfigurationFolder)
	filesystem.CreateDirectoryIfNotExist(settings.Config.DataFolder)
	settings.Config = config.New(settings.Config.GetConfigurationPath(), settings.Config)
}

func InitGitHubConfiguration() {
	settings.Config.GitHub.AskCredentials()
	settings.Config.SaveFile(settings.Config.GetConfigurationPath())
}

func PrintConfigurations() {
	settings.Config.ReadFile(settings.Config.GetConfigurationPath())
	fmt.Println(settings.Config)
}
