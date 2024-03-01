package main

import (
	"fmt"
	"strings"
)

func stringFunction() {
	var word string = "hello"

	_ = word

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i], string(word[i]))
	}

	var romanian string = "Mânca"
	fmt.Println(len(romanian))

	for i := 0; i < len(romanian); i++ {
		fmt.Println(romanian[i], string(romanian[i]))
	}

	fmt.Println(strings.Repeat("#", 60))

	r := []rune(romanian)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i], string(r[i]))
	}

	// char := word[0]
	var char string = string(word[0])
	fmt.Println(char)
	_ = char

	fmt.Println(word[0])

	var oldStudents []string = []string{"John", "Bayu"}
	var newStudents []string = []string{"Andre", "Toni"}
	var results []string = []string{}

	results = append(results, oldStudents...)
	results = append(results, newStudents...)
	_ = newStudents
	fmt.Printf("%#v\n", results)
}
