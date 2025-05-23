package utils

import (
	"fmt"
	"strings"
)

func GenerateBaseFilename(title, subtitle, asin string) string {
	base := strings.TrimSpace(title)
	if subtitle != "" {
		base += ": " + strings.TrimSpace(subtitle)
	}
	base += fmt.Sprintf(" [%s]", asin)
	return base
}