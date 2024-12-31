package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/google/go-github/v62/github"
)

type GitHub struct {
	Username           string
	Repository         string
	Token              string
	client             *github.Client
	context            context.Context
	FailAuthentication func(*GitHub)
}

func NewGitHub(username string, repo string, token string, failAuthentication func(githubRef *GitHub)) *GitHub {
	var github = GitHub{
		Username:           username,
		Repository:         repo,
		Token:              token,
		FailAuthentication: failAuthentication,
	}

	if github.CheckCredentials() {
		github.Authenticate()
	}

	return &github
}

func (g *GitHub) CheckCredentials() bool {
	var valid = true
	if g.Username == strings.TrimSpace("") {
		valid = false
		if g.FailAuthentication != nil {
			// log.Println("Unauthorized: No Username provided")
			g.FailAuthentication(g)
		}
	}
	if g.Repository == strings.TrimSpace("") {
		valid = false
		if g.FailAuthentication != nil {
			// log.Println("Unauthorized: No Repository provided")
			g.FailAuthentication(g)
		}
	}
	if g.Token == strings.TrimSpace("") {
		valid = false
		if g.FailAuthentication != nil {
			// log.Println("UnAuthorized: No AuthToken provided")
			g.FailAuthentication(g)
		}
	}
	return valid
}

func (g *GitHub) Authenticate() {
	g.context = context.Background()
	g.client = github.NewClient(nil).WithAuthToken(g.Token)
}

func (g *GitHub) createFile(filename string, fileOptions github.RepositoryContentFileOptions) error {
	var _, _, err = g.client.Repositories.CreateFile(
		g.context,
		g.Username,
		g.Repository,
		filename,
		&fileOptions)
	return err
}

func (g *GitHub) updateFile(filename string, fileOptions github.RepositoryContentFileOptions) error {
	var _, _, err = g.client.Repositories.UpdateFile(
		g.context,
		g.Username,
		g.Repository,
		filename,
		&fileOptions)
	return err
}

func (g *GitHub) getContent(path string) (*github.RepositoryContent, []*github.RepositoryContent, error) {
	var file, directory, _, err = g.client.Repositories.GetContents(
		g.context,
		g.Username,
		g.Repository,
		path,
		&github.RepositoryContentGetOptions{},
	)
	return file, directory, err
}

func (g *GitHub) DownloadDirectoryFiles(path ...string) map[string]string {
	var downloadedFiles = make(map[string]string)
	var contents, _ = g.GetDirectoryContent(path...)
	for _, content := range contents {
		if *content.Type == "file" {
			var fileContent, _ = g.GetFileContent(*content.Name)
			downloadedFiles[*content.Name] = string(fileContent)
		}
	}
	return downloadedFiles
}

func (g *GitHub) GetFileContent(filename string) (string, error) {
	var content, _, err = g.getContent(filename)
	var stringContent, _ = base64.StdEncoding.DecodeString(*content.Content)
	return string(stringContent), err
}

func (g *GitHub) GetDirectoryContent(path ...string) ([]*github.RepositoryContent, error) {
	var _, content, err = g.getContent(strings.Join(path, ""))
	return content, err
}

func (g *GitHub) UploadTextFile(filename string, content string, commitComments ...string) error {
	if !g.CheckCredentials() {
		fmt.Println(filename, "cannot be uploaded to github check credentials on the configuration file.")
		return nil
	}

	var file, _, err = g.getContent(filename)
	var sha *string

	if err == nil && file != nil {
		sha = file.SHA
	}

	if len(commitComments) == 0 {
		commitComments = append(commitComments, "Update text file")
	}

	var fileOptions = github.RepositoryContentFileOptions{
		Message: github.String(strings.Join(commitComments, "\n")),
		Content: []byte(content),
		SHA:     sha,
	}

	if sha == nil {
		g.createFile(filename, fileOptions)
	} else {
		g.updateFile(filename, fileOptions)
	}

	return err
}

func (g *GitHub) UploadBinaryFile(filePath string, commitComments ...string) error {
	if !g.CheckCredentials() {
		fmt.Println(filePath, "cannot be uploaded to GitHub, check credentials in the configuration file.")
		return fmt.Errorf("invalid credentials")
	}

	var fileName = path.Base(filePath)

	var file, _, err = g.getContent(fileName)
	var sha *string

	if err == nil && file != nil {
		sha = file.SHA
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	var content = base64.StdEncoding.EncodeToString(fileContent)

	if len(commitComments) == 0 {
		commitComments = append(commitComments, "Update binary file")
	}

	var fileOptions = github.RepositoryContentFileOptions{
		Message: github.String(strings.Join(commitComments, "\n")),
		Content: []byte(content),
		SHA:     sha,
	}

	if sha == nil {
		return g.createFile(fileName, fileOptions)
	} else {
		return g.updateFile(fileName, fileOptions)
	}
}
