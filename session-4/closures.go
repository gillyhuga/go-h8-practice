package main

import (
	"fmt"
	"time"
)

type MyCallBack func(num int) bool

func closures() {

	var result int = func(num1 int, num2 int) int {
		return num1 + num2
	}(20, 30)

	fmt.Println(result)

	_ = result

	fmt.Println(23, "Hallo semuanya")
	fmt.Println(MultipleReturnValueFunction())
	fmt.Printf(MultipleReturnValueFunction())

	word, randomNumber, m := MultipleReturnValueFunction()

	_, _, _ = word, randomNumber, m

	// fmt.Printf("word => %s, randomNumber => %d\n", word, randomNumber)

	isFoundV1 := FindNumberInSliceV1(10, []int{20, 30, 10, 40, 70})

	isFoundV2 := FindNumberInSliceV2(20, 30, 40, 50, 60)

	// fmt.Println(isFoundV1, isFoundV2)

	_, _ = isFoundV1, isFoundV2

	strResult, totalRunes := CountRunesInString("Hello World")

	_, _ = strResult, totalRunes

	// fmt.Printf("strResult => %s, totalRunes => %d\n", strResult, totalRunes)

	var nums []int = []int{10, 20, 30, 40, 59}

	var validator func(num int, findEven bool) bool = func(num int, findEven bool) bool {
		if findEven {
			if num%2 == 0 {
				return true
			}
			return false
		}

		return num%2 != 0
	}

	_, _ = nums, validator

	IsOddOrEvenNumber(nums, true, validator)

	SleepAndPrint(2, func() {
		fmt.Println("aku sudah bangun setelah tidur selama 2 detik")
	})

	SleepAndPrint(5, func() {
		fmt.Println("Hello semuanya, gmna kabar?")
	})

}

func SleepAndPrint(t int, log func()) {
	time.Sleep(time.Second * time.Duration(t))

	log()
}

func IsOddOrEvenNumber(
	numbers []int,
	isEven bool,
	oddOrEvenValidator func(num int, findEven bool) bool,
) {

	for _, num := range numbers {
		if isEven == true {
			if oddOrEvenValidator(num, isEven) == true {
				fmt.Printf("%d is an even number\n", num)
			}
		} else {
			if oddOrEvenValidator(num, isEven) == false {
				fmt.Printf("%d is an odd number\n", num)
			}
		}
	}
}

func CountRunesInString(str string) (strResult string, totalRunes int) {

	result := []rune(str)

	strResult = str

	totalRunes = len(result)

	return
}

func FindNumberInSliceV2(searchedNumber int, numbers ...int) bool {

	for _, eachNumber := range numbers {
		if searchedNumber == eachNumber {
			return true
		}
	}

	return false
}

func FindNumberInSliceV1(searchedNumber int, numbers []int) bool {
	for _, eachNumber := range numbers {
		if searchedNumber == eachNumber {
			return true
		}
	}

	return false
}

func MultipleReturnValueFunction() (string, int, map[string]int) {

	var greetings string = "Hallo semuanya"

	var randomNumber int = 23

	var m map[string]int = map[string]int{}

	return greetings, randomNumber, m
}
