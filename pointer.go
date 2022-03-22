package main

import (
	"fmt"
)

// func main() {
// 	var fisrtNumber *int
// 	var secondNumber *int
// 	_, _ = fisrtNumber, secondNumber
// }

// func main() {
// 	var fisrtNumber int = 4
// 	var secondNumber *int = &fisrtNumber

// 	fmt.Println("fisrtNumber (value) : ", fisrtNumber)
// 	fmt.Println("fisrtNumber (memori address) : ", &fisrtNumber)

// 	fmt.Println("secondNumber (value) : ", *secondNumber)
// 	fmt.Println("secondNumber (memori address) : ", secondNumber)
// }

// func main() {
// 	var firsPerson string = "John"
// 	var secondPerson *string = &firsPerson

// 	fmt.Println("firsPerson (value) :", firsPerson)
// 	fmt.Println("firsPerson (memori address) :", &firsPerson)

// 	fmt.Println("secondPerson (value) :", *secondPerson)
// 	fmt.Println("secondPerson (memori address) :", secondPerson)

// 	fmt.Println("\n", strings.Repeat("#", 30), "\n")

// 	*secondPerson = "Doe"

// 	fmt.Println("firsPerson (value) :", firsPerson)
// 	fmt.Println("firsPerson (memori address) :", &firsPerson)

// 	fmt.Println("secondPerson (value) :", *secondPerson)
// 	fmt.Println("secondPerson (memori address) :", secondPerson)

// }

func main() {
	var a int = 10
	fmt.Println("Before : ", a)

	chageValue(&a)
	fmt.Println("After : ", a)
}

func chageValue(number *int) {
	*number = 20
}
