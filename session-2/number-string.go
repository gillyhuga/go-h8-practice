package main

import "fmt"

type DrinksVariant string

const (
	CHOCOLATE_CAKE DrinksVariant = "CHOCOLATE_CAKE"
	BUBBLE_GUM     DrinksVariant = "DrinksVariant"
	THAI_TEA       DrinksVariant = "THAI_TEA"
)

func numberString() {
	var bgNumber uint = 256
	var smallNumber uint8 = uint8(bgNumber)
	var score float32 = 3.65
	var age int8 = int8(score)
	var totalScore int = int(score)
	_, _, _ = age, totalScore, smallNumber

	fmt.Println("age =>", age)
	fmt.Printf("score => %.2f\n", score)

	var isBool bool
	_ = isBool
	var name string
	var randomNumber int
	var randomFloat float32
	_, _ = randomNumber, randomFloat

	fmt.Printf("%s\n", name)
	fmt.Printf("%t\n", isBool)
	fmt.Printf("%d\n", randomNumber)
	fmt.Printf("%.1f\n", randomFloat)
	_ = name

	const fixedValue string = "DONT_CHANGE"
	var dynamicName string = "Flexible"
	// const fixedName string = dynamicName
	_ = dynamicName
	var incrementalNumber int = 10
	_ = incrementalNumber

	incrementalNumber = incrementalNumber + 1
	incrementalNumber += 1
	incrementalNumber++
	incrementalNumber += 5
	incrementalNumber = incrementalNumber + 5
	fmt.Println("Incremental Number :",incrementalNumber)
	_ = fixedValue

	var score1 int8 = 97
	var score2 uint8 = 97
	var result int = int(uint8(score1) + score2)
	_ = result
	fmt.Println(uint8(score1) + score2)

	var isTheSameNumber bool = score1 == int8(score2)
	fmt.Println("isTheSameNumber =>", isTheSameNumber)
}
