package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "selamat malam"

	words := strings.Split(kalimat, " ")

	wordCount := make(map[string]int)
	for _, word := range words {
		for _, char := range word {
			fmt.Println(string(char))
			wordCount[string(char)]++
		}
		fmt.Println()
	}

	fmt.Println(wordCount)
}
