package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	runeStr := []rune(str)

	for i := 0; i < len(runeStr)/2; i++ {
		fmt.Printf("Comparing: %v, %v \n", string(runeStr[i]), string(runeStr[len(runeStr)-i-1]))

		if runeStr[i] != runeStr[len(runeStr)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("racecar", isPalindrome("racecar"))
	fmt.Println("test", isPalindrome("test"))
	fmt.Println("dad", isPalindrome("dad"))
}
