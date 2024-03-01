package main

import "fmt"

func loop() {

	var num int
	_ = num

	for num < 10 {
		fmt.Println(num)
		num++

		if num == 4 {
			break
		}
	}

	for num < 10 {
		if num%2 == 0 {
			fmt.Printf("angka %d adalah genap\n", num)
		} else {

			if num == 5 {
				num++
				continue
			}
			fmt.Printf("angka %d adalah angka ganjil\n", num)
		}
		num++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	var description string = "detail information"
	_ = description

parentLoop:
	for i := 1; i <= 3; i++ {
		fmt.Printf("parentLoop => %d\n", i)
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break parentLoop
			}

			fmt.Printf("child loop => %d\n", j)

		}

	}

	for i := 0; i < len(description); i++ {
		fmt.Println(string(description[i]))
	}

	for _, char := range description {
		fmt.Println(string(char))
	}
}
