package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"yt2abs/internal/utils"
)

func CreateAudiobookFromFiles(baseName, outputDir string, files []string) error {
	tempDir := utils.TempDirPath()
	if tempDir == "" {
		return fmt.Errorf("could not create temporary directory")
	}

	concatListPath := filepath.Join(tempDir, "concat.txt")
	concatFile, err := os.Create(concatListPath)
	if err != nil {
		return err
	}
	defer concatFile.Close()

	for _, f := range files {
		absPath, err := filepath.Abs(f)
		if err != nil {
			return fmt.Errorf("failed to get absolute path: %v", err)
		}
		line := fmt.Sprintf("file '%s'\n", strings.ReplaceAll(absPath, "'", "'\\''"))
		_, err = concatFile.WriteString(line)
		if err != nil {
			return err
		}
	}

	m4bFile := baseName + ".m4b"
	m4bPath := filepath.Join(outputDir, m4bFile)

	args := []string{
		"-f", "concat",
		"-safe", "0",
		"-i", concatListPath,
		"-c:a", "aac",
		"-b:a", "64k",
		"-movflags", "+faststart",
		"-metadata", "encoded_by=",
		m4bPath,
	}

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running FFmpeg to create:", m4bPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w", err)
	}

	fmt.Println("FFmpeg conversion successful:", m4bPath)
	return nil
}
