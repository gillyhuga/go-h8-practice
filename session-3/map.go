package main

import "fmt"

func mapFunction() {
	var person map[string]string = map[string]string{
		"name":    "John",
		"age":     "23",
		"address": "Jalan Kemang",
	}

	_ = person

	var schools map[string][]string = map[string][]string{
		"Jaya Abadi": []string{"Ahmad", "Handoko"},
		"Sentosa":    []string{"Charles", "James"},
	}

	_ = schools

	for key, value := range schools {
		for _, eachStudent := range value {
			fmt.Printf("school (%s), student (%s)\n", key, eachStudent)
		}
	}

	var footballPlayer map[string]map[string]string = map[string]map[string]string{
		"Arhan": {
			"height":      "1.70 cm",
			"total score": "120",
		},
		"Marcelino": {
			"height":      "1.74 cm",
			"total score": "156",
		},
	}

	_ = footballPlayer

	var studentScore map[string]int = map[string]int{
		"james":   10,
		"charles": 8,
	}

	_ = studentScore

	var footBallTeams map[string][]map[string]string = map[string][]map[string]string{
		"ac-milan": []map[string]string{
			{
				"name": "Pato",
				"age":  "23",
			},
			{
				"name": "Maldini",
				"age":  "32",
			},
		},
		"chelsea": []map[string]string{
			{
				"name": "Lampard",
				"age":  "31",
			},
		},
	}

	footBallTeams["liverpool"] = []map[string]string{
		{
			"name": "Dirk Kuyt",
			"age":  "33",
		},
	}

	liverpool, ok := footBallTeams["liverpool"]
	fmt.Printf("liverpool => (%#v), ok => (%t)\n", liverpool, ok)

	delete(footBallTeams, "chelsea")

	for key := range footBallTeams {
		fmt.Println(key)
	}

	fmt.Printf("%#v\n", person)
	for key, value := range person {
	fmt.Printf("key (%s)\n", key)
	fmt.Printf("value (%s)\n", value)
	fmt.Printf("key (%s), value (%s)\n", key, value)
	}

	person["name"] = "John"

}
