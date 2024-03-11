package main

import "fmt"

type DrinkVariant string

func (dv DrinkVariant) MixAndMatch(condiment string) DrinkVariant {
	fmt.Printf("dv => %s\n", dv)

	var mixedDrink string = fmt.Sprintf("%s mixed with %s", dv, condiment)

	return DrinkVariant(mixedDrink)
}

type User struct {
	Name    string
	Address string
}

func (p *User) InsertName(newName string) {
	p.Name = newName
}

func methodFunction() {
	var a DrinkVariant = "Chocolate Tea"
	var b DrinkVariant = "Kopi Susu Keluarga"
	var aMixed DrinkVariant = a.MixAndMatch("oreo")
	var bMixed DrinkVariant = b.MixAndMatch("soy milk")
	_, _ = aMixed, bMixed
	var p1 Person = Person{Address: "Jakarta", Name: "John"}
	// p1.InsertName("John")
	fmt.Printf("p1 => %+v\n", p1)

}
