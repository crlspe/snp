package command

import (
	"fmt"
	"path"
	"strings"

	"github.com/crlspe/snp/filesystem"
	"github.com/crlspe/snp/model"
	"github.com/crlspe/snp/settings"
)

type UpdateCommand struct {
	Flags model.Flags
}

func (i *UpdateCommand) Init(flags model.Flags) {
	i.Flags = flags
}

func (i *UpdateCommand) Exec() {
	filesystem.CreateDirectoryIfNotExist(settings.Config.DataFolder)
	settings.Config.ReadFile(settings.Config.GetConfigurationPath())
	switch {
	case *i.Flags.Local:
		DownloadGithubDataFiles()
	case *i.Flags.Github:
		UploadDataFilesToGithub()
	}
}

func UploadDataFilesToGithub() {
	var files, _ = filesystem.ListDataFiles()
	for _, fileName := range files {
		var fileContent, _ = filesystem.ReadFile(fileName)
		fileName = strings.ReplaceAll(fileName, settings.Config.DataFolder, "")
		fmt.Println(fileName, " saved to GitHub.")
		fileName = path.Base(fileName)
		settings.GitHubClient.UploadTextFile(fileName, fileContent, settings.Config.GetDefaultCommitComments())
	}
}

func DownloadGithubDataFiles() {
	var snippetFiles = settings.GitHubClient.DownloadDirectoryFiles()
	for fileName, fileContent := range snippetFiles {
		filesystem.CreateFile(settings.Config.DataFolder+fileName, fileContent)
		fmt.Println(settings.Config.DataFolder+fileName, " file created.")
	}
}
