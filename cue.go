package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CUE() {
	inputFile := "input/chapters.txt"
	outputFile := "output/output.cue"
	audioFileName := "audiofile.mp3" // Kann angepasst oder als Argument übergeben werden

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Eingabedatei:", err)
		return
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Ausgabedatei:", err)
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
		title := strings.TrimSpace(parts[1])

		if strings.ToLower(title) == "end" {
			continue
		}

		writer.WriteString(fmt.Sprintf("  TRACK %02d AUDIO\n", trackNum))
		writer.WriteString(fmt.Sprintf("    TITLE \"%s\"\n", title))
		writer.WriteString(fmt.Sprintf("    INDEX 01 %s\n", timestamp))

		trackNum++
	}

	writer.Flush()
	fmt.Println("CUE-Datei geschrieben:", outputFile)
}
