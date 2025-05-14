package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateAudiobook(baseName string) {
	outputDir := "output"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error while creating directory '%s': %v\n", outputDir, err)
			os.Exit(1)
		}
	}
	m4bPath := "output/" + baseName + ".m4b"

	cmd := exec.Command("ffmpeg",
		"-i", "input/audiobook.mp3",
		"-i", "input/cover.jpg",
		"-i", "input/FFMETADATA.txt",
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

	fmt.Println("FFmpeg-conversion successful")
}
