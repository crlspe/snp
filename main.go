package main

import (
	"github.com/crlspe/snp/cli"
	"github.com/crlspe/snp/config"
	"github.com/crlspe/snp/logging"
	"github.com/crlspe/snp/settings"
)

func init() {
	settings.Config.ReadFile(settings.Config.GetConfigurationPath())
	logging.SetLogLevel(settings.Config.LogLevel)
}

func main() {
	logging.Debug(config.GetApplicationVersion())
	var input = cli.NewInput()
	input.Command.Exec()
}
