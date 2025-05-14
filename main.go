package main

import (
	"fmt"
)

func main() {
	fmt.Println("Step 1: fetch Audible Metadata")
	product, err := fetchAudibleMetadata("B002UZKUC6")
	if err != nil {
		fmt.Println("Error while fetching metadata:", err)
		return
	}
	fmt.Println("Step 2: save cover image")
	err = SaveCoverImage(product.ProductImages.Image500, "input/cover.jpg")
	if err != nil {
		fmt.Println("Cover konnte nicht gespeichert werden:", err)
	}
	fmt.Println("Step 2: creating FFMETADATA.txt")
	FFMETADATA(product)
	fmt.Println("Step 3: creating .cue chapter file")
	CUE()
	fmt.Println("Step 4: creating .m4b audiobook")
	FFMPEG()
}
