package main

import "fmt"

func main() {
	var name string = "Jhon"
	var address string = "Jakarta"
	var age int = 60
	var weight float32 = 50.2

	//Typecasting
	var score int = int(weight)

	information := "Information"
	description := "Description"

	_, _ = description, information

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Weight:", weight)
	fmt.Println("Address:", address)
	fmt.Println(information, description)

	fmt.Println(score)
}
