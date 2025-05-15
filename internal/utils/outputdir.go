package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateOutputDir(basePath, title, asin string) string {
	dirName := fmt.Sprintf("%s [%s]", strings.TrimSpace(title), asin)
	fullPath := filepath.Join(basePath, dirName)
	os.MkdirAll(fullPath, 0755)
	return fullPath
}