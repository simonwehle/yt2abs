package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func TempDirPath() string {
	tempDir := filepath.Join(os.Getenv("HOME"), ".yt2abs")

	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err := os.MkdirAll(tempDir, 0755)
		if err != nil {
			fmt.Println("Error while creating temp dir:", err)
			return ""
		}
		fmt.Println("Temp dir created:", tempDir)
	}

	return tempDir
}

func CleanTempDir() error {
	tempDir := TempDirPath()
	if tempDir == "" {
		return fmt.Errorf("no temporary directory found")
	}

	files, err := os.ReadDir(tempDir)
	if err != nil {
		return fmt.Errorf("error while reading the directory contents: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(tempDir, file.Name())
		err := os.Remove(filePath)
		if err != nil {
			return fmt.Errorf("error deleting file %s: %v", filePath, err)
		}
	}

	fmt.Println("Temporary files have been deleted.")
	return nil
}