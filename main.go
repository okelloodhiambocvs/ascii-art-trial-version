package main

import (
	"fmt"
	"os"
)

// Entry point of the program
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . \"Your text here\"")
		return
	}

	input := os.Args[1]

	// Load banner file
	banner, err := LoadBanner("banner.txt")
	if err != nil {
		fmt.Println("Error loading banner file:", err)
		return
	}

	// Render ASCII
	result := RenderASCII(input, banner)
	fmt.Print(result)
}
