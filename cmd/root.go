package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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
	fmt.Printf("Usage: %s [-a <ASIN>] | [-i <InputAudio>] [-c <ChaptersFile>] | [--nc] [-o <OutputLocation>]\n", toolName)
	fmt.Println("\nOptions:")
	fmt.Println("  -a <ASIN>         Audible ASIN (e.g., B07KKMNZCH) - Required unless -i is used")
	fmt.Println("  -i <InputAudio>   Path to the MP3 file (default: audiobook.mp3)")
	fmt.Println("  -c <ChaptersFile> Path to the chapters file (default: chapters.txt)")
	fmt.Println("  -o <OutputPath>   Path to output location")
	fmt.Println("  -nc               No chapters: skip creating .cue and FFMETADATA.txt")
	fmt.Println("  -h                Show this help message")
	fmt.Println("  -v                Show version")
	fmt.Println("\nExamples:")
	fmt.Println("  yt2abs -a B07KKMNZCH                 Full auto ASIN mode")
	fmt.Println("  yt2abs -i \"My Book.mp3\" -nc         Only MP3 to M4B no metadata and chapters")
	fmt.Println("  yt2abs -a B07KKMNZCH -o /output/path Define output folder")
}

func Execute() {
	asin := flag.String("a", "", "Audible ASIN (e.g., B07KKMNZCH)")
	title := flag.String("t", "", "Title of the output audiobook")
	audioFile := flag.String("i", "audiobook.mp3", "Path to the MP3 file")
	inputFolder := flag.String("f", "", "Path to input Folder")
	chapterFile := flag.String("c", "chapters.txt", "Path to the chapters file")
	noChapters := flag.Bool("nc", false, "No chapters")
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

	outputBase := "."
	if *outputFolder != "" {
		info, err := os.Stat(*outputFolder)
		if err != nil || !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Error: output folder '%s' does not exist.\n", *outputFolder)
			os.Exit(1)
		}
		outputBase = *outputFolder
	}

	var (
		baseName       string
		outputDir      string
		includeMetadata = false
	)

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

	if *inputFolder == "" {
		defaultChaptersExist := false
		if _, err := os.Stat("chapters.txt"); err == nil {
			defaultChaptersExist = true
		}

		if !defaultChaptersExist && *chapterFile == "chapters.txt" && !*noChapters {
			fmt.Fprintln(os.Stderr, "Error: No chapters file found. Provide -c (chapters) or use -nc to disable chapters.")
			os.Exit(1)
		}
	} else {
		if *chapterFile != "chapters.txt" {
			fmt.Println("Note: -c is ignored when using -f (folder).")
		}
		if *noChapters {
			fmt.Println("Note: -nc is ignored when using -f (folder).")
		}
	}

	if *asin != "" {
		fmt.Println("Step 1: fetch Audible Metadata")
		product, err := audible.FetchMetadata(*asin)
		if err != nil {
			fmt.Println("Error while fetching metadata:", err)
			return
		}
		includeMetadata = !*noChapters
		baseName = utils.GenerateBaseFilename(product.Title, product.Subtitle, *asin)
		outputDir = utils.GenerateOutputDir(outputBase, product.Title, *asin)

		fmt.Println("Step 2: save cover image")
		if err := cover.SaveImage(product.ProductImages.Image500); err != nil {
			fmt.Println("Cover couldn't be saved:", err)
		}

		if !*noChapters {
			fmt.Println("Step 3: creating FFMETADATA.txt")
			metadata.CreateFFMETADATA(product, *chapterFile)

			fmt.Println("Step 4: creating .cue chapter file")
			cue.CreateCue(baseName, outputDir, *chapterFile)
		}
	} else {
		baseName = utils.StripExtension(filepath.Base(*audioFile))
		outputDir = utils.GenerateOutputDir(outputBase, *title, "")
		includeMetadata = false

		if !*noChapters {
			fmt.Println("Step 1: creating .cue chapter file")
			cue.CreateCue(*title, outputDir, *chapterFile)
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
