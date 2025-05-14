package main

import (
	"fmt"
)

func main() {
	fmt.Println("Step 1: fetch Audible Metadata")
	product, err := fetchAudibleMetadata("B002V02KPU")
	if err != nil {
		fmt.Println("Fehler:", err)
		return
	}
	fmt.Println("Step 2: creating FFMETADATA.txt")
	FFMETADATA(product)
	fmt.Println("Step 3: creating .cue chapter file")
	CUE()
	fmt.Println("Step 4: creating .m4b audiobook")
	FFMPEG()
}
