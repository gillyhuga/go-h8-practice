package main

import (
	"fmt"
	"session-5/helper"
)

func main() {
	var c1 helper.Car = helper.Car{
		Price: 50000000,
		Brand: "Nissan",
	}
	// fmt.Println(helper.secret)
	c1.SetMachineVersion("admin", "V3")
	c1.MachineVersion()
	fmt.Printf("c1 => %+v\n", c1)

	methodFunction()
	structFunction()
}
