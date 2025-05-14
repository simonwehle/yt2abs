package main

import (
	"flag"
	"fmt"
	"os"
)

const toolName = "yt2abs"

func printHelp() {
	fmt.Printf("Usage: %s --asin <ASIN> [--audio <AudioFile>] [--chapters <ChaptersFile>]\n", toolName)
	fmt.Println("\nOptions:")
	fmt.Println("  --asin <ASIN>        Audible ASIN (e.g., B0BJ35FVXD) - Required")
	fmt.Println("  --audio <AudioFile>  Path to the MP3 file (default: audiobook.mp3)")
	fmt.Println("  --chapters <ChaptersFile> Path to the chapters text file (default: chapters.txt)")
	fmt.Println("  --help               Show this help message")
}

func main() {
asin := flag.String("asin", "", "Audible ASIN (e.g., B0BJ35FVXD)")
	audioFile := flag.String("audio", "audiobook.mp3", "Path to the MP3 file")
	chapterFile := flag.String("chapters", "chapters.txt", "Path to the chapters file")
	showHelp := flag.Bool("help", false, "Show this help message")

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	if *asin == "" {
		fmt.Printf("Error: --asin is required. Use '%s --help' for more information.\n", toolName)
		os.Exit(1)
	}

	fmt.Println("ASIN:", *asin)
	fmt.Println("Audio:", *audioFile)
	fmt.Println("Kapitel:", *chapterFile)

	fmt.Println("Step 1: fetch Audible Metadata")
	product, err := fetchAudibleMetadata(*asin)
	if err != nil {
		fmt.Println("Error while fetching metadata:", err)
		return
	}
	baseName := generateBaseFilename(product.Title, product.Subtitle, *asin)
	fmt.Println("Step 2: save cover image")
	err = SaveCoverImage(product.ProductImages.Image500, "input/cover.jpg")
	if err != nil {
		fmt.Println("Cover konnte nicht gespeichert werden:", err)
	}
	fmt.Println("Step 2: creating FFMETADATA.txt")
	FFMETADATA(product)
	fmt.Println("Step 3: creating .cue chapter file")
	CUE(baseName)
	fmt.Println("Step 4: creating .m4b audiobook")
	FFMPEG(baseName)
}
