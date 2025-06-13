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
const version = "1.3.0"

func Execute() {
	asin := flag.String("a", "", "Audible ASIN (e.g., B07KKMNZCH)")
	title := flag.String("t", "", "Title of the output audiobook")
	audioFile := flag.String("i", "audiobook.mp3", "Path to the MP3 file")
	inputFolder := flag.String("f", "", "Path to input folder")
	chapterFile := flag.String("c", "chapters.txt", "Path to the chapters file")
	outputFolder := flag.String("o", "", "Path to output location")
	showHelp := flag.Bool("h", false, "Show this help message")
	showVersion := flag.Bool("v", false, "Show version")

	if len(os.Args) == 1 {
		utils.PrintHelp(toolName)
		os.Exit(0)
	}

	flag.Parse()

	if *showHelp {
		utils.PrintHelp(toolName)
		return
	}
	if *showVersion {
		fmt.Printf("%s version %s\n", toolName, version)
		return
	}

	if *asin == "" && *title == "" {
		fmt.Fprintln(os.Stderr, "Error: You must specify either -a (ASIN) or -t (Title).")
		os.Exit(1)
	}

	defaultAudioExists := false
	if _, err := os.Stat("audiobook.mp3"); err == nil {
		defaultAudioExists = true
	}

	if !defaultAudioExists && *audioFile == "audiobook.mp3" && *inputFolder == "" {
		fmt.Fprintln(os.Stderr, "Error: No input audio found. Specify -i (audio file) or -f (folder).")
		os.Exit(1)
	}

	outputBase := "."
	if *outputFolder != "" {
		info, err := os.Stat(*outputFolder)
		if err != nil || !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Error: output folder '%s' does not exist.\n", *outputFolder)
			os.Exit(1)
		}
		outputBase = *outputFolder
	}

	chaptersEnabled := false
	if *inputFolder == "" {
		if *chapterFile != "chapters.txt" {
			chaptersEnabled = true
		} else if _, err := os.Stat("chapters.txt"); err == nil {
			chaptersEnabled = true
		} else {
			fmt.Println("Note: No chapters provided; skipping chapter and metadata generation.")
		}
	} else {
		if *chapterFile != "chapters.txt" {
			fmt.Println("Note: -c is ignored when using -f (folder).")
		}
	}

	var (
		baseName       string
		outputDir      string
		includeMetadata = false
	)

	if *asin != "" {
		fmt.Println("Step 1: fetch Audible Metadata")
		product, err := audible.FetchMetadata(*asin)
		if err != nil {
			fmt.Println("Error while fetching metadata:", err)
			return
		}
		baseName = utils.GenerateBaseFilename(product.Title, product.Subtitle, *asin)
		outputDir = utils.GenerateOutputDir(outputBase, product.Title, *asin)

		fmt.Println("Step 2: save cover image")
		if err := cover.SaveImage(product.ProductImages.Image500); err != nil {
			fmt.Println("Cover couldn't be saved:", err)
		}

		if chaptersEnabled {
			fmt.Println("Step 3: creating FFMETADATA.txt")
			metadata.CreateFFMETADATA(product, *chapterFile)

			fmt.Println("Step 4: creating .cue chapter file")
			cue.CreateCue(baseName, outputDir, *chapterFile)
		}
	} else {
		outputDir = utils.GenerateOutputDir(outputBase, *title, "")
		includeMetadata = false

		if chaptersEnabled {
			fmt.Println("Step 1: creating .cue chapter file")
			cue.CreateCue(baseName, outputDir, *chapterFile)
		}
	}

	fmt.Println("Creating .m4b audiobook")
	ffmpeg.CreateAudiobook(*title, outputDir, *audioFile, includeMetadata)

	defer func() {
		if err := utils.CleanTempDir(); err != nil {
			fmt.Println("Error cleaning up temp directory:", err)
		} else {
			fmt.Println("Temporary files cleaned up.")
		}
	}()
}
