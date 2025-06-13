package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func GenerateChaptersFile(folder, outputPath string) error {
	files, err := filepath.Glob(filepath.Join(folder, "*.mp3"))
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("no mp3 files found in folder: %s", folder)
	}

	sort.Strings(files)

	var totalSeconds float64
	var lines []string

	for _, file := range files {
		duration, err := getAudioDuration(file)
		if err != nil {
			return fmt.Errorf("failed to get duration for %s: %v", file, err)
		}

		title := extractTitleFromFilename(filepath.Base(file))
		timestamp := formatTimestamp(totalSeconds)
		lines = append(lines, fmt.Sprintf("%s %s", timestamp, title))

		totalSeconds += duration
	}

	endTimestamp := formatTimestamp(totalSeconds)
	lines = append(lines, fmt.Sprintf("%s End", endTimestamp))

	return os.WriteFile(outputPath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}


func getAudioDuration(file string) (float64, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		file,
	)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
}

func formatTimestamp(seconds float64) string {
	h := int(seconds) / 3600
	m := (int(seconds) % 3600) / 60
	s := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func extractTitleFromFilename(name string) string {
	re := regexp.MustCompile(`^\d+\s*-\s*`)
	title := re.ReplaceAllString(name, "")
	return strings.TrimSuffix(title, filepath.Ext(title))
}
