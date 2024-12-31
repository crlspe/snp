package model

import (
	"github.com/crlspe/snp/filesystem"
	"github.com/crlspe/snp/settings"
)

type FileStorage struct {
	Storage
	FileName string
}

func NewFileStorage(snippet Snippet) FileStorage {
	var storage = FileStorage{}
	storage.New(snippet)
	return storage
}

func (fs FileStorage) Save() string {
	filesystem.CreateFile(
		settings.Config.GetDataFilePath(fs.Snippet.GenerateSnippetFileName()),
		fs.Snippet.GetSnippetContent(),
	)
	return fs.FileName
}
