package main

type Person struct {
	Name    string
	Address string
	Age     uint
}

type Employee struct {
	Common Person
	Salary uint
}

/*
	{
		"Salary": 140000
		"Common": {
			"Name": "John",
			"Address": "Jalan Kemang",
			"Age": 23
		}
	}
*/

type Student struct {
	Grade      string
	SchoolName string
	Person
}

/*
	{
		"Grade": "2A",
		"SchoolName":  "Jaya Abadi",
		"Name": "John",
		"Address": "Jalan Kemang",
		"Age": 23
	}
*/

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

/*
	{
		"value": 10,
		"left": {
			value: 03,
			left: {},
			right: {}
		},
		"right": {}
	}
*/

func structFunction() {
	var e1 Employee = Employee{
		Common: Person{
			Name: "John",
		},
		Salary: 10000,
	}

	e1.Common.Address = "Jakarta"

	e1.Common.Age = 24

	var s1 Student = Student{
		Grade:      "2A",
		SchoolName: "Tadika Puri",
		Person: Person{
			Name:    "Waldo",
			Address: "Jakarta",
		},
	}

	s1.Age = 20

	// s1.Person.Age = 20

	// fmt.Printf("s1 => %+v\n", s1)

	tempValue := struct {
		RandomNumber int
		RandomString string
	}{
		RandomNumber: 10,
		RandomString: "asdf",
	}

	tempValue.RandomNumber = 23

	tempValue.RandomString = "fdsa"
}
