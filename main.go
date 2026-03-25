package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// LoadBanner reads the banner file and maps each ASCII character to its 8-line ASCII art
func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)

	asciiStart := 32 // space
	asciiEnd := 126  // tilde (~)
	charHeight := 8  // each character is 8 lines

	expectedLines := (asciiEnd - asciiStart + 1) * charHeight
	if len(lines) < expectedLines {
		return nil, fmt.Errorf("banner file too short: expected %d lines, got %d", expectedLines, len(lines))
	}

	for i := asciiStart; i <= asciiEnd; i++ {
		startLine := (i - asciiStart) * charHeight
		banner[rune(i)] = lines[startLine : startLine+charHeight]
	}

	return banner, nil
}

// PrintASCII prints the input string in ASCII using the banner map
func PrintASCII(input string, banner map[rune][]string) {
	charHeight := 8
	lines := make([]string, charHeight)

	for _, ch := range input {
		if ch == '\n' {
			for _, line := range lines {
				fmt.Println(line)
			}
			lines = make([]string, charHeight)
			continue
		}

		if val, ok := banner[ch]; ok {
			for i := 0; i < charHeight; i++ {
				lines[i] += val[i]
			}
		} else {
			// Fallback: print 8 spaces for unknown characters
			for i := 0; i < charHeight; i++ {
				lines[i] += "        "
			}
		}
	}

	// print remaining lines
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string to convert")
		return
	}

	// Read input argument
	input := os.Args[1]

	// Convert literal "\n" to real newlines
	input = strings.ReplaceAll(input, `\n`, "\n")

	// Load the standard banner
	banner, err := LoadBanner("banners/standard.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	// Print ASCII art
	PrintASCII(input, banner)
}
