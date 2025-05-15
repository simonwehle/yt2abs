package utils

import (
	"fmt"
	"strings"
)

func GenerateOutputDir(title, asin string) string {
	dirName := fmt.Sprintf("%s [%s]", strings.TrimSpace(title), asin)
	return dirName
}