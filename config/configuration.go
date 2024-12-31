package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
	"reflect"

	"github.com/crlspe/snp/color"
	"github.com/crlspe/snp/logging"
	"github.com/crlspe/snp/ui"
)

const ApplicationName = "snp"
const Version = "1.0"
const UnixDefaultConfigFolderName = ".config"
const UnixDefaultDataFolderName = ".local/share"
const JSON = "json"

type FuncCommentCommit func() string

type Configuration struct {
	LogLevel                int                 `json:"LogLevel"`
	DataFolder              string              `json:"DataFolder"`
	DataFileExtension       string              `json:"DefaultFileExtension"`
	DataFileNameTemplate    string              `json:"FileNameTemplate"`
	DataFileHeaderSeparator string              `json:"DefaultHeaderSeparator"`
	ConfigurationFolder     string              `json:"ConfigurationFolder"`
	ConfigurationFileName   string              `json:"ConfigurationFileName"`
	GitHub                  GitHubConfiguration `json:"GitHub"`
	GitHubDefaultComments   FuncCommentCommit   `json:"-"`
	GitHubUploadAfterAdd    bool                `json:"GitHubUploadAfterAdd"`
	DataUseSqliteDb         bool                `json:"DataUseSqliteDb"`
	DataSqliteDbName        string              `json:"DataSqliteDbName"`
}

type WritableConfiguration struct {
	DataFolder           string              `json:"DataFolder"`
	GitHub               GitHubConfiguration `json:"GitHub"`
	GitHubUploadAfterAdd bool                `json:"-"`
}

type GitHubConfiguration struct {
	User       string `json:"User"`
	Repository string `json:"Repository"`
	Token      string `json:"Token"`
}

func New(configFilePath string, config Configuration) Configuration {
	config.SaveFileOnlyIfNotExists(configFilePath)
	config.ReadFile(configFilePath)
	return config
}

func (c *Configuration) ReadFile(filePath string) {
	if FileNotExists(filePath) {
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		logging.Debug("Error opening configuration file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		logging.Debug("Error decoding JSON:", err)
		return
	}
}

func (c *Configuration) SaveFile(outputFilePath string) error {
	return c.saveFile(outputFilePath)
}

func (c *Configuration) SaveFileOnlyIfNotExists(outputFilePath string) error {
	var err error = nil
	if FileNotExists(outputFilePath) {
		err = c.saveFile(outputFilePath)
	}
	return err
}

func (c *Configuration) saveFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("could not create configuration file: %v", err)
	}
	defer file.Close()
	defaultConfig, err := json.MarshalIndent(c.WritetableConfiguration(), "", "    ")
	if err != nil {
		return fmt.Errorf("could not marshal default configuration: %v", err)
	}
	_, err = file.Write(defaultConfig)
	if err != nil {
		return fmt.Errorf("could not write to configuration file: %v", err)
	}
	logging.Debug("Configuration File: '", fileName, "' created with default values.")
	return nil
}

func (c *Configuration) WritetableConfiguration() WritableConfiguration {
	var writetableConfig = WritableConfiguration{
		DataFolder:           c.DataFolder,
		GitHubUploadAfterAdd: c.GitHubUploadAfterAdd,
		GitHub:               c.GitHub,
	}
	return writetableConfig
}

func (c *Configuration) GetDefaultCommitComments() string {
	return c.GitHubDefaultComments()
}

func (c *Configuration) GetConfigurationPath() string {
	return path.Join(c.ConfigurationFolder, c.ConfigurationFileName)
}

func (c Configuration) GetDataFilePath(fileName string) string {
	return path.Join(c.DataFolder, fileName)
}

func GetApplicationVersion() string {
	var version = ApplicationName + " v" + Version
	fmt.Print(fmt.Sprintln(color.Green(ApplicationName), color.Yellow("v"+Version)))
	return version
}

func FileNotExists(filename string) bool {
	var _, err = os.Stat(filename)
	return os.IsNotExist(err)
}

func GetOSConfigurationFolder() string {
	return path.Join(GetOSHomeFolder(), UnixDefaultConfigFolderName, ApplicationName)
}

func GetOSDataFolder() string {
	return path.Join(GetOSHomeFolder(), UnixDefaultDataFolderName, ApplicationName)
}

func GetOSHomeFolder() string {
	var usr, _ = user.Current()
	var homeFolder = usr.HomeDir
	return homeFolder
}

func (c *Configuration) GetFileNameTemplate() string {
	return c.DataFileNameTemplate + c.DataFileExtension
}

func (c Configuration) String() string {
	var str string = ""
	var confVal = reflect.ValueOf(c)
	var conf = reflect.TypeOf(c)
	for i := 0; i < confVal.NumField(); i++ {
		if conf.Field(i).Tag.Get(JSON) != "-" {
			str += fmt.Sprintln(color.BrightBlue(conf.Field(i).Name), confVal.Field(i))
		}
	}
	return str
}

func (g *GitHubConfiguration) AskCredentials() {
	g.User, g.Repository, g.Token = ui.GitHubCredentials(g.User, g.Repository, g.Token)
}
