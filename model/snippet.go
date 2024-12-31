package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/crlspe/snp/filesystem"
	"github.com/crlspe/snp/settings"
)

type Snippet struct {
	Description string
	Scopes      string
	Tags        string
	Code        string
}

func (s *Snippet) Save() {
	NewFileStorage(*s).Save()
	if settings.Config.GitHubUploadAfterAdd {
		NewGitHubStorage(*s).Save()
	}
}

func (s *Snippet) GenerateSnippetFileName() string {
	return formatFileName(
		fmt.Sprintf(
			settings.Config.GetFileNameTemplate(),
			formatSection(s.Scopes, "-"),
			formatSection(s.Tags, "-"),
			filesystem.GenerateUniqueID(),
		))
}

func (s *Snippet) GetSnippetContent() string {
	var description = "description: " + strings.ReplaceAll(s.Description, "\n", "")
	var scopes = "scopes: " + formatSection(s.Scopes, ",")
	var tags = "tags: " + formatSection(s.Tags, ",")

	return fmt.Sprintln(settings.Config.DataFileHeaderSeparator) +
		fmt.Sprintln(description) +
		fmt.Sprintln(scopes) +
		fmt.Sprintln(tags) +
		fmt.Sprintln(settings.Config.DataFileHeaderSeparator) +
		fmt.Sprintln(s.Code)
}

func formatFileName(fileName string) string {
	fileName = strings.ReplaceAll(fileName, "--", "")
	fileName = strings.ReplaceAll(fileName, "-.", ".")
	return fileName
}

func formatSection(section, separator string) string {
	return strings.ReplaceAll(section, " ", separator)
}

func ReadSnippetHeader(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open file %s: %v", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesRead := 0
	var lines []string

	for scanner.Scan() {
		if linesRead < 5 {
			line := scanner.Text()
			if line != "" {
				lines = append(lines, line)
			}
			linesRead++
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	return strings.Join(lines, "\n"), nil
}

func ReadSnippetFile(filePath string) (Snippet, error) {
	var snippet Snippet
	var snippetFile, err = os.ReadFile(filePath)
	if err != nil {
		return snippet, err
	}

	var fileContent = strings.Split(string(snippetFile), settings.Config.DataFileHeaderSeparator)

	if len(fileContent) == 1 {
		snippet.Code = strings.TrimSpace(fileContent[0])
	}

	if len(fileContent) >= 2 {
		snippet.Description = strings.TrimSpace(fileContent[1])
	}

	if len(fileContent) >= 3 {
		snippet.Code = strings.TrimSpace(fileContent[2])
	}
	return snippet, nil
}
