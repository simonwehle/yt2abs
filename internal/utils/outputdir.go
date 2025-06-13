package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateOutputDir(basePath, title, asin string) string {
	var dirName string
	title = strings.TrimSpace(title)

	if asin != "" {
		dirName = fmt.Sprintf("%s [%s]", title, asin)
	} else {
		dirName = title
	}

	fullPath := filepath.Join(basePath, dirName)
	os.MkdirAll(fullPath, 0755)
	return fullPath
}