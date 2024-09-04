package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	clearScreen()
	moveCursor(5, 10)
	printWithColor("Hello, Terminal GUI!", 31) // Red text

	moveCursor(7, 10)
	printWithColor("Press 'q' to quit...", 32) // Green text

	// Simple event loop to wait for 'q' to quit
	for {
		var b = make([]byte, 1)
		os.Stdin.Read(b)
		if b[0] == 'q' {
			break
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear") // 'clear' command to clear the terminal
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col) // ANSI escape code to move the cursor
}

func printWithColor(text string, color int) {
	fmt.Printf("\033[%dm%s\033[0m", color, text) // ANSI escape code for colored text
}

