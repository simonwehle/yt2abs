package cover

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SaveImage(url string, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while fetching cover: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: HTTP-Status %d during cover fetch", resp.StatusCode)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error during file creation: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error while saving cover: %w", err)
	}

	fmt.Println("Coverbild gespeichert als:", outputPath)
	return nil
}
