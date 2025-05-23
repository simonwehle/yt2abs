package utils

import (
	"path/filepath"
	"strings"
)

func StripExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}