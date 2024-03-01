package main

import (
	"fmt"
	"strings"
)

func stringFunction() {
	word := "mânca"
	var wordInRune []rune = []rune(word)
	var r rune = 'â'
	_, _ = wordInRune, r

	var b []byte = []byte{109, 120, 123, 104}
	fmt.Println(string(b))

	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]), word[i])
	}

	fmt.Println(strings.Repeat("#", 60))

	for i := 0; i < len(wordInRune); i++ {
		fmt.Println(string(wordInRune[i]), wordInRune[i])
	}

	for index, char := range word {
		fmt.Println(index, char, string(char))
	}
	fmt.Println(len(wordInRune))
}
