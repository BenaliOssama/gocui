package main

import (
	"fmt"
)

func main() {
	//drawBox(5, 3, 20, 5) // Draw a box at position (5, 3) with width 20 and height 5
	fmt.Println()
	drawUnicodeBox(5, 3, 20 , 5 )
}

func drawBox(x, y, width, height int) {
	// Top border
	moveCursor(y, x)
	fmt.Print("+")
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	// Side borders
	for j := 0; j < height; j++ {
		moveCursor(y+j+1, x)
		fmt.Print("|")
		moveCursor(y+j+1, x+width+1)
		fmt.Println("|")
	}

	// Bottom border
	moveCursor(y+height+1, x)
	fmt.Print("+")
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col) // ANSI escape code to move the cursor
}
func drawUnicodeBox(x, y, width, height int) {
	// Top border
	moveCursor(y, x)
	fmt.Print("╔")
	for i := 0; i < width; i++ {
		fmt.Print("═")
	}
	fmt.Println("╗")

	// Side borders
	for j := 0; j < height; j++ {
		moveCursor(y+j+1, x)
		fmt.Print("║")
		moveCursor(y+j+1, x+width+1)
		fmt.Println("║")
	}

	// Bottom border
	moveCursor(y+height+1, x)
	fmt.Print("╚")
	for i := 0; i < width; i++ {
		fmt.Print("═")
	}
	fmt.Println("╝")
}

