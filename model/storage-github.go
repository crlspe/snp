package model

import "github.com/crlspe/snp/settings"

type GitHubStorage struct {
	Storage
}

func NewGitHubStorage(snippet Snippet) GitHubStorage {
	var storage = GitHubStorage{}
	storage.New(snippet)
	return storage
}

func (gh GitHubStorage) Save() {
	settings.GitHubClient.UploadTextFile(
		gh.Snippet.GenerateSnippetFileName(),
		gh.Snippet.GetSnippetContent(),
		settings.Config.GetDefaultCommitComments())
}
