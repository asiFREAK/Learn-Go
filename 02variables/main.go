package main

import "fmt"

const Anothervarible string = "bdbx" // public

func main() {
	var username string = "Rishav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("Variable is of type: %T \n", smallint)

	var smallfloat float64 = 255.327532437875397
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	// default values

	var newvariable int
	fmt.Println(newvariable)
	fmt.Printf("Variable is of type: %T \n", newvariable)

	// implicit type

	var hello = 253
	fmt.Println(hello)
	fmt.Printf("Variable is of type: %T \n", hello)

	// no var

	countuser := "30"
	fmt.Println(countuser)
	fmt.Printf("Variable is of type: %T \n", countuser)

	// const var

	fmt.Println(Anothervarible)
	fmt.Printf("Variable is of type: %T \n", Anothervarible)

}