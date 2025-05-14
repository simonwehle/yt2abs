package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// SaveCoverImage lädt ein Bild von der angegebenen URL herunter und speichert es unter dem Pfad "input/cover.jpg"
func SaveCoverImage(url string, outputPath string) error {
	// HTTP-GET für das Bild
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while fetching cover: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: HTTP-Status %d during cover fetch", resp.StatusCode)
	}

	// Datei erstellen
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error during file creation: %w", err)
	}
	defer outFile.Close()

	// Daten in die Datei schreiben
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error while saving cover: %w", err)
	}

	fmt.Println("Coverbild gespeichert als:", outputPath)
	return nil
}
