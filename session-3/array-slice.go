package main

import (
	"fmt"
	"strings"
)

func arraySlice() {
	var fruits [4]string = [4]string{
		"Banana",
		"Apple",
		"Mango",
	}

	fruits[3] = "Durian"

	fmt.Println("println =>", fruits)
	fmt.Printf("printf => %#v\n", fruits)

	fruits[0] = "Pineapple"
	fmt.Printf("printf => %#v\n", fruits)

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index (%d), name (%s)\n", i, fruits[i])
	}

	for index, eachFruit := range fruits {
		fmt.Printf("index (%d), name (%s)\n", index, eachFruit)
	}

	var cars []string = []string{
		"Honda",
		"Toyota",
		"BMW",
		"Wuling",
		"Hyundai",
	}

	var newCars []string = cars[0:4]

	var raceCars []string = cars[1:3]

	fmt.Printf("new Cars => %#v\n", newCars)
	fmt.Printf("newCars => len (%d), cap (%d)\n\n", len(newCars), cap(newCars))
	fmt.Printf("race Cars => %#v\n", raceCars)
	fmt.Printf("raceCars => len (%d), cap (%d)\n", len(raceCars), cap(raceCars))

	newCars = append(newCars, "Daihatsu")

	fmt.Printf("alamat memory index ke 0 cars => %v\n", &cars[0])
	fmt.Printf("alamat memory index ke 0 newCars => %v\n", &newCars[0])
	fmt.Printf("cars => len (%d), cap (%d), %#v\n", len(cars), cap(cars), cars)
	fmt.Printf("newCars => len (%d), cap (%d), %#v\n", len(newCars), cap(newCars), newCars)

	cars = append(cars, "Wuling")

	fmt.Printf("len (%d), cap (%d)\n", len(cars), cap(cars))

	cars = append(cars, "Mitsubishi")
	fmt.Printf("len (%d), cap (%d)\n", len(cars), cap(cars))

	cars = append(cars, "Mercedes")
	fmt.Printf("len (%d), cap (%d)\n", len(cars), cap(cars))

	cars = append(cars, "Hyundai")
	fmt.Printf("len (%d), cap (%d)\n", len(cars), cap(cars))

	cars = append(cars, "Honda", "Toyota", "BMW")

	// var newCars []string = cars

	// fmt.Printf("alamat memory dari index ke 0 cars => %v\n", &cars[0])
	// fmt.Printf("alamat memory dari index ke 0 newCars => %v\n", &newCars[0])
	// fmt.Printf("cars => %#v\n", cars)
	// fmt.Printf("newCars => %#v\n", newCars)

	newCars[0] = "Range Rover"

	fmt.Println(strings.Repeat("#", 60))

	fmt.Printf("cars => %#v\n", cars)
	fmt.Printf("newCars => %#v\n", newCars)

}
