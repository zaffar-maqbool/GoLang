package main

import (
	"fmt"

	"github.com/zaffar-maqbool/Golang/under/greet"
)

var Firstname string = "Zaffar "

func main() {
	fmt.Println("greet", greet.Greeting)
	fmt.Println("constant ", pi)
	var Lastname = "Maqbool"
	fullname := Firstname + Lastname
	Fullname := fmt.Sprintf("%v %v", Firstname, Lastname)
	sex := "Male"
	age := 22
	dob := "03.03.2000"
	currentYear := 2022
	email := "aalimwani94@gmail.com"
	fmt.Println(Firstname, Lastname, sex, age, dob, currentYear, email, Fullname+" wani")

	// %v is used for varaibles
	fmt.Printf("hi my name is %v and i am %v years old", fullname, age)
}
