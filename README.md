# ASCII Art Generator in Go

This is a simple Go program that takes a string from the terminal and prints it in ASCII art using a custom banner file.

It can handle letters, numbers, spaces, special characters, and line breaks (`\n`).  

---

## Features

- Converts text to ASCII art
- Supports multiple lines
- Uses a banner file to define character styles
- Written in pure Go using only standard packages
- Includes unit tests

---

## Project Structure


ascii-art/
│
├─ main.go # Entry point
├─ ascii.go # Core logic to render ASCII art
├─ ascii_test.go # Unit tests
├─ banner.txt # ASCII font file (height=8)
├─ go.mod # Go module file
└─ README.md # Project documentation


---

## How to Run

1. Open the project folder in VS Code or your terminal.
2. Make sure Go is installed. Check with:

```bash
go version
Run the program with your text. Example:
go run . "Hello\nThere"

Use \n for new lines in your text.

Example output:
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                               
                               
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 