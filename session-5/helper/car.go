package helper

import "fmt"

type Car struct {
	Price          uint
	Brand          string
	machineVersion string
}

func (c *Car) SetMachineVersion(role string, machineVersions string) {
	if role != "admin" {
		panic("you are not authorized to access this resource")
	}
	c.machineVersion = machineVersions
}

func (c Car) MachineVersion() {
	fmt.Printf("machine version => %s\n", c.machineVersion)
	fmt.Println(secret)
}
