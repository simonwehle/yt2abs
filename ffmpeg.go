package main

import (
	"fmt"
	"os"
	"os/exec"
)

func FFMPEG() {
	cmd := exec.Command("ffmpeg",
		"-i", "input.mp3",
		"-i", "cover.jpg",
		"-i", "chapters.txt",
		"-map", "0:a",
		"-map", "1",
		"-map_metadata", "2",
		"-c:a", "aac",
		"-b:a", "64k",
		"-vn",
		"-movflags", "+faststart",
		"output.m4b",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "FFmpeg-Fehler: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("FFmpeg-Konvertierung abgeschlossen: output.m4b")
}
