package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Q: What is the msot common word in sherlock.txt?

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func wordFrequency(r io.Reader) (map[string]int, error) {
	// You don't want to load the entire file in memory. For many files, you can go single line by line in memory.
	dictionary := make(map[string]int) // word -> count
	s := bufio.NewScanner(r)
	lnum := 0

	for s.Scan() {
		lnum++
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			dictionary[strings.ToLower(w)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return dictionary, nil
}

func maxWord(dictionary map[string]int) (string, error) {
	if len(dictionary) == 0 {
		return "", fmt.Errorf("empty dictionary")
	}

	maxWord, maxCount := "", 0

	for word, count := range dictionary {
		if count > maxCount {
			maxWord = word
			maxCount = count
		}
	}

	fmt.Printf("%s -> %d", maxWord, maxCount)
	return maxWord, nil
}

func mapDemo() {
	stocks := make(map[string]float64) // symbol -> price
	symbol := "TTWO"
	price := stocks[symbol]

	fmt.Printf("%s -> $%.2f\n", symbol, price)

	if price, ok := stocks[symbol]; ok {
		fmt.Printf("%s -> $%.2f\n", symbol, price)
	} else {
		fmt.Printf("%s not found\n", symbol)
	}

	stocks[symbol] = 136.73
	stocks["AAPL"] = 217.90

	for k := range stocks { // Only key
		fmt.Println(k)
	}

	for k, v := range stocks { // key and value
		fmt.Printf("%s -> $%.2f\n", k, v)
	}

	for _, v := range stocks { // Only value
		fmt.Println(v)
	}
}

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer file.Close()

	dictionary, err := wordFrequency(file)
	if err != nil {
		log.Fatalf("Failed to get word frequency: %v", err)
	}

	maxWord(dictionary)
}
