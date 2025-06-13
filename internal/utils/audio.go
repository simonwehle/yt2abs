package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GetSortedAudioFiles(folder string) ([]string, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, fmt.Errorf("failed to read folder %s: %w", folder, err)
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(strings.ToLower(name), ".mp3") {
			files = append(files, filepath.Join(folder, name))
		}
	}

	sort.Strings(files)

	if len(files) == 0 {
		return nil, fmt.Errorf("no .mp3 files found in folder %s", folder)
	}

	return files, nil
}
