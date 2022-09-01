package main

import "fmt"

func main() {
	hobbies := []string{
		"Coding\n",
		"Cooking\n",
		"Currently nothing else\n"}

	fmt.Println("11", hobbies)
	fmt.Println("12", hobbies[0:1])

	slicesHobbies := hobbies[1:]
	fmt.Println("15", slicesHobbies)

	reslicedhobbies := slicesHobbies[0:2]
	fmt.Println("18", reslicedhobbies)

	products := []Product{
		{
			"first",
			"a product",
			12,
		},
		{
			"2nd",
			"2nd product",
			15,
		},
	}
	fmt.Println(products)

}

type Product struct {
	id    string
	title string
	price float64
}
