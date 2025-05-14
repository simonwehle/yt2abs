package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func FFMETADATA(product *Product) {
	inputFile := "input/chapters.txt"
	outputFile := "input/FFMETADATA.txt"

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
	writer.WriteString(";FFMETADATA1\n")

	authors := extractNames(product.Authors)
	narrators := extractNames(product.Narrators)
	comment := strings.ReplaceAll(product.PublisherSummary, "\n", " ")
	comment = strings.ReplaceAll(comment, "\r", " ")

	writer.WriteString(fmt.Sprintf("title=%s\n", product.Title))
	writer.WriteString(fmt.Sprintf("album=%s\n", product.Title))
	writer.WriteString(fmt.Sprintf("artist=%s\n", authors))
	writer.WriteString(fmt.Sprintf("composer=%s\n", narrators))
	writer.WriteString(fmt.Sprintf("date=%s\n", product.ReleaseDate))
	writer.WriteString(fmt.Sprintf("publisher=%s\n", product.PublisherName))
	writer.WriteString(fmt.Sprintf("comment=%s\n", comment))
	writer.WriteString("\n")

	scanner := bufio.NewScanner(in)
	type Chapter struct {
		Start int
		Title string
	}
	var chapters []Chapter
	var finalEnd int
	var lastLine string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lastLine = line
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		startTimeStr := strings.TrimSpace(parts[0])
		title := strings.TrimSpace(parts[1])

		startSec, err := parseTimeToSeconds(startTimeStr)
		if err != nil {
			fmt.Println("Ungültiges Zeitformat:", startTimeStr)
			continue
		}

		if strings.ToLower(title) == "end" {
			finalEnd = startSec
		} else {
			chapters = append(chapters, Chapter{Start: startSec, Title: title})
		}
	}

	if !strings.HasSuffix(strings.ToLower(lastLine), "end") {
		fmt.Println("Fehler: Die letzte Zeile in der Datei muss ein gültiger 'End'-Eintrag sein.")
		return
	}

	for i, c := range chapters {
		var end int
		if i+1 < len(chapters) {
			end = chapters[i+1].Start
		} else {
			end = finalEnd
		}
		writer.WriteString("[CHAPTER]\n")
		writer.WriteString("TIMEBASE=1/1\n")
		writer.WriteString("START=" + strconv.Itoa(c.Start) + "\n")
		writer.WriteString("END=" + strconv.Itoa(end) + "\n")
		writer.WriteString("title=" + c.Title + "\n")
	}

	writer.Flush()
	fmt.Println("Konvertierung abgeschlossen:", outputFile)
}

func parseTimeToSeconds(timeStr string) (int, error) {
	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return 0, err
	}
	return t.Hour()*3600 + t.Minute()*60 + t.Second(), nil
}
