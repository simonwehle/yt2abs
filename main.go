package main

import (
	"fmt"
	//"os"
)

func main() {
/*     if len(os.Args) < 2 {
        fmt.Println("Usage: <toolname> <inputfile>")
        return
    }
    inputFile := os.Args[1] */
    // Hier kommt deine Umwandlungslogik

	fmt.Println("audiobook-creator Version 0.1.0")
	//fmt.Println("checking input files...")
	fmt.Println("Step 1: creating FFMETADATA.txt")
	FFMETADATA()
	fmt.Println("Step 2: creating .cue chapter file")
	CUE()
	fmt.Println("Step 3: creating .m4b audiobook")
}
