package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"yt2abs/internal/utils"
)

func CreateAudiobook(baseName, outputDir, audioFile string, includeMetadata bool) {
	tempDir := utils.TempDirPath()
	if tempDir == "" {
		fmt.Println("Error: Could not create temporary directory.")
		return
	}

	m4bFile := baseName + ".m4b"
	m4bPath := filepath.Join(outputDir, m4bFile)

	args := []string{
		"-i", audioFile,
	}

	if includeMetadata {
		coverPath := filepath.Join(tempDir, "cover.jpg")
		metadataPath := filepath.Join(tempDir, "FFMETADATA.txt")

		args = append(args,
			"-i", coverPath,
			"-i", metadataPath,
			"-map", "0:a",
			"-map", "1:v",
			"-map_metadata", "2",
			"-c:v", "mjpeg",
			"-disposition:v", "attached_pic",
		)
	}

	args = append(args,
		"-c:a", "aac",
		"-b:a", "64k",
		"-movflags", "+faststart",
		"-metadata", "encoded_by=",
		m4bPath,
	)

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "FFmpeg error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("FFmpeg conversion successful:", m4bPath)
}
