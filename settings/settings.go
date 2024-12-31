package settings

import (
	"log/slog"
	"time"

	"github.com/crlspe/snp/config"
	"github.com/crlspe/snp/github"
)

var Config = config.Configuration{
	LogLevel: int(slog.LevelInfo),

	ConfigurationFolder:   config.GetOSConfigurationFolder(),
	ConfigurationFileName: "config.json",

	DataFolder:              config.GetOSDataFolder(),
	DataFileNameTemplate:    "%v-%v-%v",
	DataFileHeaderSeparator: "---",
	DataFileExtension:       ".md",

	DataUseSqliteDb:  false,
	DataSqliteDbName: "snp.db",

	GitHubUploadAfterAdd: true,
	GitHubDefaultComments: func() string {
		return "Saved on:" + time.Now().Format("02/01/2000")
	},
}

var GitHubClient *github.GitHub
