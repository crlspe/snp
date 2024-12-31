package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/crlspe/snp/filesystem"
	"github.com/crlspe/snp/logging"
	"github.com/crlspe/snp/model"

	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

type SearchCommand struct {
	Flags model.Flags
}

type Metadata map[string]string

func (s *SearchCommand) Init(flags model.Flags) {
	s.Flags = flags
}

func (s *SearchCommand) Exec() {
	var metadata = getFileHeaderMetadata()
	var fileResults = searchFiles(metadata)

	switch {
	case *s.Flags.Run:
		if len(fileResults) >= 1 {
			var snippet, _ = model.ReadSnippetFile(fileResults[0])
			logging.Debug("Executing: ", snippet.Code)
			var commandParts = strings.Fields(snippet.Code)
			var cmd = exec.Command(commandParts[0], commandParts[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(output))
		}
	default:
		ShowResults(fileResults)
	}
}

func ShowResults(searchResults []string) {
	app := tview.NewApplication()
	list := tview.NewList()
	for index, fileName := range searchResults {
		var snippet, _ = model.ReadSnippetFile(fileName)
		list.AddItem(
			snippet.Code,
			strings.Join(strings.Split(snippet.Description, "\n"), " | ")+" | "+fileName,
			rune('0'+index),
			nil,
		)
	}
	list.AddItem("Quit", "Press to exit.", 'q', nil)

	list.SetSelectedFunc(func(i int, mainText string, secondaryText string, shortcut rune) {
		app.Stop()
		copyToClipboard(mainText)
	})

	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func copyToClipboard(mainText string) {
	if mainText != "Quit" {
		err := clipboard.Init()
		if err != nil {
			panic(err)
		}
		clipboard.Write(clipboard.FmtText, []byte(mainText))
		<-time.After(200 * time.Millisecond)
		logging.Debug("'%s' has been copied to clipboard.", mainText)
	}
}

func searchFiles(metadata Metadata) []string {
	var fileResults []string
	var searchTerms = os.Args[2:]

	if len(searchTerms) == 0 {
		searchTerms = []string{""}
	}

	for filename, header := range metadata {
		for _, term := range searchTerms {
			if contains(header, term) {
				fileResults = append(fileResults, filename)
			}
		}
	}

	return fileResults
}

func contains(text string, search string) bool {
	var contains = true
	text = strings.TrimSpace(text)
	search = strings.TrimSpace(search)

	if len(strings.Split(text, " ")) == 1 {
		return strings.Contains(text, search)
	} else {
		for _, word := range strings.Split(search, " ") {
			contains = contains && strings.Contains(text, word)
		}
	}
	return contains
}

func getFileHeaderMetadata() Metadata {
	var files, _ = filesystem.ListDataFiles()
	var metadata = make(Metadata)
	for _, snippet := range files {
		var header, _ = model.ReadSnippetHeader(snippet)
		metadata[snippet] = header
	}
	return metadata
}
