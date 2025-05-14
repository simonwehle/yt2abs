package cover

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"yt2abs/internal/utils"
)

func SaveImage(url string) error {
	tempDir := utils.TempDirPath()
	if tempDir == "" {
		return fmt.Errorf("could not create temporary folder")
	}

	fileName := "cover.jpg"
	tempFilePath := filepath.Join(tempDir, fileName)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while fetching cover: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: HTTP status %d during cover fetch", resp.StatusCode)
	}

	outFile, err := os.Create(tempFilePath)
	if err != nil {
		return fmt.Errorf("error during file creation: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error while saving cover: %w", err)
	}

	fmt.Println("Cover image saved in temporary directory:", tempFilePath)
	return nil
}
