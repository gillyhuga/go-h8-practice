package main

type DrinkVariant string

const (
	CHOCOLATE_TEA DrinkVariant = "CHOCOLATE_TEA"
	THAI_TEA      DrinkVariant = "THAI_TEA"
	MANGO_TEA     DrinkVariant = "MANGO_TEA"
	GUAVA_TEA     DrinkVariant = "GUAVA_TEA"
)

type Score int

func aliaseAndNewType() {
	var score1 int = 10
	var score2 Score = Score(score1)
	_ = score2

	var juice string = "orange juice"
	var specialDrink DrinkVariant = DrinkVariant(juice)
	_ = specialDrink

	var drinks map[string]DrinkVariant = map[string]DrinkVariant{
		"monday":    CHOCOLATE_TEA,
		"tuesday":   MANGO_TEA,
		"wednesday": THAI_TEA,
		"thursday":  GUAVA_TEA,
	}

	_ = drinks

}
