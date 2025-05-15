package cue

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// cue/cue.go
func CreateCue(baseName, outputDir, chapterFilePath string) {
	os.MkdirAll(outputDir, 0755)

	outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.cue", baseName))
	audioFileName := baseName + ".m4b"

	in, err := os.Open(chapterFilePath)
	if err != nil {
		fmt.Println("Error opening chapter file:", err)
		return
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating CUE file:", err)
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	writer.WriteString(fmt.Sprintf(`FILE "%s" MP3`, audioFileName) + "\n")

	scanner := bufio.NewScanner(in)
	trackNum := 1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		timestamp := strings.TrimSpace(parts[0])
		baseName := strings.TrimSpace(parts[1])

		if strings.ToLower(baseName) == "end" {
			continue
		}

		writer.WriteString(fmt.Sprintf("  TRACK %02d AUDIO\n", trackNum))
		writer.WriteString(fmt.Sprintf("    baseName \"%s\"\n", baseName))
		writer.WriteString(fmt.Sprintf("    INDEX 01 %s\n", timestamp))
		trackNum++
	}

	writer.Flush()
	fmt.Println("CUE file written to:", outputFile)
}
