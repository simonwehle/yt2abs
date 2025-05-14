package cmd

import (
	"flag"
	"fmt"
	"os"

	"yt2abs/internal/audible"
	"yt2abs/internal/cover"
	"yt2abs/internal/cue"
	"yt2abs/internal/ffmpeg"
	"yt2abs/internal/metadata"
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

func Execute() {
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
	product, err := audible.FetchMetadata(*asin)
	if err != nil {
		fmt.Println("Error while fetching metadata:", err)
		return
	}

	baseName := metadata.GenerateBaseFilename(product.Title, product.Subtitle, *asin)

	fmt.Println("Step 2: save cover image")
	err = cover.SaveImage(product.ProductImages.Image500, "input/cover.jpg")
	if err != nil {
		fmt.Println("Cover couldn't be saved:", err)
	}

	fmt.Println("Step 3: creating FFMETADATA.txt")
	metadata.CreateFFMETADATA(product)

	fmt.Println("Step 4: creating .cue chapter file")
	cue.CreateCue(baseName, *chapterFile)

	fmt.Println("Step 5: creating .m4b audiobook")
	ffmpeg.CreateAudiobook(baseName)
}
