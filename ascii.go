package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Banner maps characters to their ASCII art
type Banner map[rune][]string

// LoadBanner reads the banner file and returns a Banner
func LoadBanner(path string) (Banner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	banner := make(Banner)
	var lines []string
	var currentRune rune

	for scanner.Scan() {
		line := scanner.Text()

		// Empty line indicates new character
		if line == "" && len(lines) > 0 {
			banner[currentRune] = lines
			lines = nil
			currentRune++
			continue
		}

		// If first character, set rune
		if len(lines) == 0 {
			if currentRune == 0 {
				currentRune = 32 // start from space
			}
		}

		lines = append(lines, line)
	}

	// Add last character if file doesn't end with empty line
	if len(lines) > 0 {
		banner[currentRune] = lines
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banner, nil
}

// RenderASCII converts input string to ASCII art using a Banner
func RenderASCII(input string, banner Banner) string {
	lines := make([]string, 8) // height=8

	for _, r := range input {
		if r == '\n' {
			// Print current lines and reset
			fmt.Println(strings.Join(lines, "\n"))
			lines = make([]string, 8)
			continue
		}

		charArt, ok := banner[r]
		if !ok {
			// fallback to '?' if character not found
			charArt, ok = banner['?']
			if !ok {
				charArt = []string{"        ", "        ", "        ", "        ", "        ", "        ", "        ", "        "}
			}
		}

		for i := 0; i < 8; i++ {
			lines[i] += charArt[i]
		}
	}

	return strings.Join(lines, "\n") + "\n"
}
