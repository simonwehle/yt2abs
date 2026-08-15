package cover

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"yt2abs/internal/utils"
)

func SaveImage(url string) error {
	return SaveImageSource(url)
}

func SaveImageSource(source string) error {
	tempDir := utils.TempDirPath()
	if tempDir == "" {
		return fmt.Errorf("could not create temporary folder")
	}

	fileName := "cover.jpg"
	tempFilePath := filepath.Join(tempDir, fileName)

	if strings.HasPrefix(strings.ToLower(source), "http://") || strings.HasPrefix(strings.ToLower(source), "https://") {
		return saveRemote(source, tempFilePath)
	}

	if err := copyFile(source, tempFilePath); err != nil {
		return fmt.Errorf("error while copying cover: %w", err)
	}
	fmt.Println("Cover image saved in temporary directory:", tempFilePath)
	return nil
}

func RemoveImage() error {
	path := filepath.Join(utils.TempDirPath(), "cover.jpg")
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func saveRemote(url, tempFilePath string) error {
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

func copyFile(source, destination string) error {
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
