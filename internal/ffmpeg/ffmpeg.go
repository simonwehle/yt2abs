package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"yt2abs/internal/utils"
)

func CreateAudiobook(baseName, audioFile string) {
	tempDir := utils.TempDirPath()
	if tempDir == "" {
		fmt.Println("Error: Could not create temporary directory.")
		return
	}

	m4bPath := baseName+".m4b"
	coverPath := filepath.Join(tempDir, "cover.jpg")
	metadataPath := filepath.Join(tempDir, "FFMETADATA.txt")

	cmd := exec.Command("ffmpeg",
		"-i", audioFile,
		"-i", coverPath,
		"-i", metadataPath,
		"-map", "0:a",
		"-map", "1:v",
		"-map_metadata", "2",
		"-c:a", "aac",
		"-b:a", "64k",
		"-c:v", "mjpeg",
		"-disposition:v", "attached_pic",
		"-movflags", "+faststart",
		m4bPath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "FFmpeg error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("FFmpeg-conversion successful:", m4bPath)
}
