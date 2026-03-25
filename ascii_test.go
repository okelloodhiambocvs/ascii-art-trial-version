package main

import (
	"strings"
	"testing"
)

func TestRenderASCII(t *testing.T) {
	banner := Banner{
		'H': {
			" _  _ ",
			"| || |",
			"| || |",
			"| || |",
			"|_  _|",
			"      ",
			"      ",
			"      ",
		},
		'e': {
			" ___ ",
			"| __|",
			"| _| ",
			"|_|  ",
			"     ",
			"     ",
			"     ",
			"     ",
		},
		'l': {
			" _ ",
			"| |",
			"| |",
			"| |",
			"|_|",
			"   ",
			"   ",
			"   ",
		},
		'o': {
			" ___ ",
			"| _ |",
			"|  _|",
			"|_|  ",
			"     ",
			"     ",
			"     ",
			"     ",
		},
	}

	output := RenderASCII("Hello", banner)
	if !strings.Contains(output, "_  _") {
		t.Error("Expected ASCII output to contain '_  _'")
	}
}
