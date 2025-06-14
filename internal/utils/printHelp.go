package utils

import (
	"fmt"
)

func PrintHelp(toolName string) {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s metadata input [chapters] [output]\n\n", toolName)

	fmt.Println("Sections and Options:")

	fmt.Println("\n  metadata (required when using default audiobook.mp3)")
	fmt.Println("    -a <ASIN>         Audible ASIN (e.g., B07KKMNZCH)")
	fmt.Println("    -t <Title>        Title of the output audiobook")

	fmt.Println("\n  input (required if no metadata is set)")
	fmt.Println("    -i <InputAudio>   Path to the MP3 file (default: audiobook.mp3)")
	fmt.Println("    -f <InputFolder>  Path to input folder (chapters by file length)")

	fmt.Println("\n  chapters (optional)")
	fmt.Println("    -c <ChaptersFile> Path to the chapters file (default: chapters.txt)")
	fmt.Println("                      Chapters are skipped if no file is provided.")

	fmt.Println("\n  output")
	fmt.Println("    -o <OutputPath>   Path to output folder (default: current directory)")

	fmt.Println("\n  misc")
	fmt.Println("    -h                Show this help message")
	fmt.Println("    -v                Show version")

	fmt.Println("\nExamples:")
	fmt.Printf("  %s -a B07KKMNZCH                    Full auto ASIN mode\n", toolName)
	fmt.Printf("  %s -t \"My Audiobook\"                Only MP3 to M4B\n", toolName)
	fmt.Printf("  %s -i \"./path/MyBook.mp3\"           Input file\n", toolName)
	fmt.Printf("  %s -a B07KKMNZCH -o \"/output/path\"  Define output folder\n", toolName)

	fmt.Printf("\n  %s -a B017V4IM1G -f .               Folder mode\n", toolName)
	fmt.Printf("  %s -t \"My Book\" -f \"/file/path\"     Set title\n", toolName)
	fmt.Printf("  %s -f .                             Use folder name as title\n", toolName)
}