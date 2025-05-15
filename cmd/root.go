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
	"yt2abs/internal/utils"
)

const toolName = "yt2abs"
const version = "1.1.0"

func printHelp() {
	fmt.Printf("Usage: %s -a <ASIN> [-i <InputAudio>] [-c <ChaptersFile>] [-o <OutputLocation>]\n", toolName)
	fmt.Println("\nOptions:")
	fmt.Println("  -a <ASIN>         Audible ASIN (e.g., B07KKMNZCH) - Required")
	fmt.Println("  -i <InputAudio>   Path to the MP3 file (default: audiobook.mp3)")
	fmt.Println("  -c <ChaptersFile> Path to the chapters file (default: chapters.txt)")
	fmt.Println("  -o <OutputPath>   Path to output location")
	fmt.Println("  -h                Show this help message")
	fmt.Println("  -v                Show version")
}

func Execute() {
	asin := flag.String("a", "", "Audible ASIN (e.g., B07KKMNZCH)")
	audioFile := flag.String("i", "audiobook.mp3", "Path to the MP3 file")
	chapterFile := flag.String("c", "chapters.txt", "Path to the chapters file")
	outputFolder := flag.String("o", "", "Path to output location")
	showHelp := flag.Bool("h", false, "Show this help message")
	showVersion := flag.Bool("v", false, "Show version")

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	if *showVersion {
		fmt.Println("yt2abs version", version)
		return
	}

	if *asin == "" {
		fmt.Println("Error: ASIN is required. Use -h for help.")
		os.Exit(1)
	}

	outputBase := "."
	if *outputFolder != "" {
		info, err := os.Stat(*outputFolder)
		if err != nil || !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Fehler: Ausgabeordner '%s' existiert nicht.\n", *outputFolder)
			os.Exit(1)
		}
		outputBase = *outputFolder
	}

	fmt.Println("Step 1: fetch Audible Metadata")
	product, err := audible.FetchMetadata(*asin)
	if err != nil {
		fmt.Println("Error while fetching metadata:", err)
		return
	}

	baseName := metadata.GenerateBaseFilename(product.Title, product.Subtitle, *asin)
	outputDir := utils.GenerateOutputDir(outputBase, product.Title, *asin)

	fmt.Println("Step 2: save cover image")
	err = cover.SaveImage(product.ProductImages.Image500)
	if err != nil {
		fmt.Println("Cover couldn't be saved:", err)
	}

	fmt.Println("Step 3: creating FFMETADATA.txt")
	metadata.CreateFFMETADATA(product, *chapterFile)

	fmt.Println("Step 4: creating .cue chapter file")
	cue.CreateCue(baseName, outputDir, *chapterFile)

	fmt.Println("Step 5: creating .m4b audiobook")
	ffmpeg.CreateAudiobook(baseName, outputDir, *audioFile)

	defer func() {
		err := utils.CleanTempDir()
		if err != nil {
			fmt.Println("Error cleaning up temp directory:", err)
		} else {
			fmt.Println("Final Step: Temporary files cleaned up.")
		}
	}()
}
