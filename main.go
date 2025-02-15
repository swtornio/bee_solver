package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1) Prompt for 6 space-delimited letters
	fmt.Print("Enter 6 letters separated by spaces: ")
	scanner.Scan()
	inputLetters := strings.Fields(strings.ToLower(scanner.Text()))
	if len(inputLetters) != 6 {
		log.Fatal("Expected exactly 6 letters, got %d", len(inputLetters))
	}

	// 2) Prompt for the special letter
	fmt.Print("Enter the special letter: ")
	scanner.Scan()
	specialInput := strings.ToLower(scanner.Text())
	if len(specialInput) != 1 {
		log.Fatal("Expected exactly 1 letter, got %d", len(specialInput))
	}
	specialLetter := rune(specialInput[0])

	// Combine the 6 letters with the special letter to form the full set
	allLetters := append(inputLetters, string(specialLetter))

	// Build a lookup set (map[rune]bool) for quick membership checks
	allowedLetters := make(map[rune]bool)
	for _, letter := range allLetters {
		r := []rune(letter)[0]
		allowedLetters[r] = true
	}

	// 3) Prompt for the word list
	fmt.Print("Enter the path to the word list: ")
	scanner.Scan()
	wordListPath := scanner.Text()

	file, err := os.Open(wordListPath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var validWords []string
	fileScanner := bufio.NewScanner(file)

	// 4) For each word in the file:
	//    - Check length >= 4
	//    - Check if it contains the special letter
	//    - Check every letter in the word is in the allowed set
	for fileScanner.Scan() {
		word := strings.ToLower(strings.TrimSpace(fileScanner.Text()))
		if len(word) < 4 {
			continue
		}
		if !strings.ContainsRune(word, specialLetter) {
			continue
		}

		// check if all letters are in the allowed set
		valid := true
		for _, r := range word {
			if !allowedLetters[r] {
				valid = false
				break
			}
		}

		if valid {
			validWords = append(validWords, word)
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 5) sort and print the valid words
	sort.Strings(validWords)
	fmt.Println("Valid words:")
	for _, w := range validWords {
		fmt.Println(w)
	}
}
