package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	padding := (width - len(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func size(text string) {
	// fmt.Println("Size:", len(text)) // BUG: Doesn't work for values with higher than ASCII. Can't use len(). Always be careful about using len() for strings.
	fmt.Println("Size:", utf8.RuneCountInString(text)) // FIX: Use []rune() to get the correct length for strings with non-ASCII characters.

	for index, value := range text {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func main() {
	banner("Go!", 8)
	size("Go!")

	banner("G😊!", 8)
	size("G😊!")
}
