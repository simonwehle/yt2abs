package main

import (
	"fmt"
)

func main() {
	asin := "B002UZKUC6"
	fmt.Println("Step 1: fetch Audible Metadata")
	product, err := fetchAudibleMetadata(asin)
	if err != nil {
		fmt.Println("Error while fetching metadata:", err)
		return
	}
	baseName := generateBaseFilename(product.Title, product.Subtitle, asin)
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
