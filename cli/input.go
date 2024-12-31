package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/crlspe/snp/command"
	"github.com/crlspe/snp/model"
)

type Input struct {
	Command command.Command
}

func NewInput() Input {
	CheckCommands()
	var input = Input{}
	input.Command = command.GetCommand()
	input.Command.Init(*model.ParseArgs())
	return input
}

func CheckCommands() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(`

Please enter a command.

search <terms>: Searchs terms in the database.
add:            Adds a snippet.
		-c | --clipboard : Uses the content of the clipboard.

		`))

		os.Exit(0)
	}
}
