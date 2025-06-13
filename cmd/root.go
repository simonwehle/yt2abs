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

func printHelp() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s metadata [input] [chapters] [output]\n\n", toolName)

	fmt.Println("Sections and Options:")

	fmt.Println("\n  metadata (required)")
	fmt.Println("    -a <ASIN>         Audible ASIN (e.g., B07KKMNZCH)")
	fmt.Println("    -t <Title>        Title of the output audiobook")

	fmt.Println("\n  input")
	fmt.Println("    -i <InputAudio>   Path to the MP3 file (default: audiobook.mp3)")
	fmt.Println("    -f <InputFolder>  Path to input folder (disables chapter processing)")

	fmt.Println("\n  chapters (optional)")
	fmt.Println("    -c <ChaptersFile> Path to the chapters file (default: chapters.txt)")
	fmt.Println("                      Chapters are skipped if no file is provided.")

	fmt.Println("\n  output")
	fmt.Println("    -o <OutputPath>   Path to output folder (default: current directory)")

	fmt.Println("\n  misc")
	fmt.Println("    -h                Show this help message")
	fmt.Println("    -v                Show version")

	fmt.Println("\nExamples:")
	fmt.Println("  yt2abs -a B07KKMNZCH                    Full auto ASIN mode")
	fmt.Println("  yt2abs -t \"My Audiobook\"                Only MP3 to M4B")
	fmt.Println("  yt2abs -a B07KKMNZCH -o \"/output/path\"  Define output folder")
}

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
		printHelp()
		os.Exit(0)
	}

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}
	if *showVersion {
		fmt.Println("yt2abs version", version)
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
