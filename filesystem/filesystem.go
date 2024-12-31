package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/crlspe/snp/settings"
)

func CreateFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(filePath string) (string, error) {
	var fileContent, err = os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

func FileNotExists(filename string) bool {
	var _, err = os.Stat(filename)
	return os.IsNotExist(err)
}

func CreateDirectoryIfNotExist(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return nil
}

func GenerateUniqueID() int64 {
	return time.Now().UnixNano()
}

func ListDataFiles() ([]string, error) {
	return listDirectoryFilesByExtension(settings.Config.DataFolder, settings.Config.DataFileExtension)
}

func listDirectoryFilesByExtension(directory string, extension string) ([]string, error) {
	var files []string
	var err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
