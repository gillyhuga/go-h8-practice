package main

import "fmt"

func ChangeValue(str *string, newStr string) {
	fmt.Println("pointer dari str =>", str)
	*str = newStr
}

func pointer() {

	var a string = "Hallo"
	var b *string = &a
	*b = "Apa kabar?"

	var num1 int = 10
	var num2 *int = &num1
	*num2 = 11

	var name string = "John"
	fmt.Println("pointer dari name =>", &name)

	ChangeValue(&name, "Aladin")

	fmt.Println(name)

	// if b != nil {
	// 	var c string = *b

	// 	_ = c

	// 	// fmt.Println(&a, b)

	// 	fmt.Println(*b)
	// }

	_, _ = a, b

}
