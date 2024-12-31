package model

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	Clipboard *bool
	Github    *bool
	Local     *bool
	Run       *bool
	DryRun    *bool
}

func ParseArgs() *Flags {
	var flags = Flags{}
	flags.Clipboard = pflag.BoolP("clipboard", "c", false, "Get the content from clipboard.")
	flags.Github = pflag.Bool("github", false, "Refer to GitHub repository.")
	flags.Local = pflag.Bool("local", false, "Refer to Local files.")
	flags.Run = pflag.Bool("run", false, "Run first snippet that matches the search.")
	flags.DryRun = pflag.Bool("dry-run", false, "Dry run cofiguration.")
	pflag.Parse()
	return &flags
}
