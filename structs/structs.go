package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func getUserData(promtText string) string {
	fmt.Print(promtText)
	userInput, _ := reader.ReadString('\n')

	/* These Two lines are creating problems somethinf
	is wrong with the syntax same happened with the calci also*/

	// cleanedInput := strings.Replace(userInput, "\n", " ", -1)
	// return cleanedInput
	return userInput
}

type Deatils struct {
	firstName string
	lastName  string
	birthdate string
	created   time.Time
}

// funcoutputUserDetails(user *Deatils) {
// 	fmt.Printf(" My Name Is  %v %v And my DOB is : %v ", user.firstName, user.lastName, user.birthdate)
// 	// fmt.Printf(" My Name Is  %v %v And my DOB is : %v ", (*user).firstName, (*user).lastName, (*user).birthdate)
// 	// Correct way but go automaticlly derefrence it
// }

// Methods
func (user *Deatils) outputUserDetails() {
	fmt.Printf(" My Name Is  %v %v And my DOB is : %v ", user.firstName, user.lastName, user.birthdate)
	// fmt.Printf(" My Name Is  %v %v And my DOB is : %v ", (*user).firstName, (*user).lastName, (*user).birthdate)
	// Correct way but go automaticlly derefrence it
}

func NewUser(fName string, lName string, bDate string) *Deatils {
	created := time.Now()

	user := Deatils{
		fName,
		lName,
		bDate,
		created,
	}
	return &user
}
func main() {

	//value  type
	var newUser Deatils
	FirstName := getUserData("Please Enter Your First Name: ")
	LastName := getUserData("Please Enter Your Last Name: ")
	BirthDate := getUserData("Please Enter Your Birthdate in (MM/DD/YYYY: ")
	// Created := time.Now()

	// newUser = User{firstName: FirstName,
	// 	lastName:  LastName,
	// 	birthdate: Birthdate,
	// 	created:   Created,
	// }
	// newUser = User{
	// 	FirstName,
	// 	LastName,
	// 	BirthDate,
	// 	Created,
	// }
	newUser = *NewUser(FirstName, LastName, BirthDate)

	//fmt.Println(FirstName, LastName, Birthdate, Created)
	// fmt.Println(newUser)

	// outputUserDetails(&newUser)
	newUser.outputUserDetails()
}
