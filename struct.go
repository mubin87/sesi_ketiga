package main

import "fmt"

// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// func main() {
// 	var employee Employee
// 	employee.name = "Airel"
// 	employee.age = 23
// 	employee.division = "Curicullum Developer"

// 	fmt.Println(employee.name)
// 	fmt.Println(employee.age)
// 	fmt.Println(employee.division)
// }

// func main() {
// 	var employee1 Employee
// 	employee1.name = "Airel"
// 	employee1.age = 23
// 	employee1.division = "Curicullum Developer"

// 	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

// 	fmt.Printf("Employee1 : %+v\n", employee1)
// 	fmt.Printf("Employee2 : %+v\n", employee2)
// }

// //pointer to struct
// func main() {
// 	var employee1 = Employee{name: "Airel", age: 23, division: "Curicullum Developer"}

// 	var employee2 *Employee = &employee1

// 	fmt.Println("Employee1 name : ", employee1.name)
// 	fmt.Println("Employee2 name : ", employee2.name)

// 	fmt.Println(strings.Repeat("#", 20))

// 	employee2.name = "ananda"

// 	fmt.Println("Employee1 name : ", employee1.name)
// 	fmt.Println("Employee2 name : ", employee2.name)
// }

// type Person struct {
// 	name string
// 	age  int
// }

// type Employee struct {
// 	division string
// 	person   Person
// }

// func main() {
// 	var employee1 = Employee{}

// 	employee1.person.name = "Airell"
// 	employee1.person.age = 23
// 	employee1.division = "Curriculum Developer"

// 	fmt.Printf("%+v", employee1)
// }

// type Employee struct {
// 	division string
// 	person   Person
// }

// //Anonymous Struct
// func main() {
// 	//anonymous struct tanpa pengisian field
// 	var employee1 = struct {
// 		person   Person
// 		division string
// 	}{}
// 	employee1.person = Person{name: "Airell", age: 23}
// 	employee1.division = "Curricullum Developer"

// 	//anonymous struct dengan pengisian field

// 	var employee2 = struct {
// 		person   Person
// 		division string
// 	}{
// 		person:   Person{name: "Ananda", age: 23},
// 		division: "Finance",
// 	}
// 	fmt.Printf("Employee1 : %+v\n", ememployee1)
// 	fmt.Printf("Employee1 : %+v\n", ememployee2)
// }

type Person struct {
	name string
	age  int
}

func main() {
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Ananda", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
