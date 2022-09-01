package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func NewProduct(ID string, T string, D string, P float64) *Product {
	return &Product{ID, T, D, P}
}

func (prod *Product) printData() {
	fmt.Printf(" ID: %v ", prod.id)
	fmt.Printf("Title: %v ", prod.title)
	fmt.Printf("Description: %v ", prod.description)
	fmt.Printf("Price: USD %f ", prod.price)

}

func readUserInput(reader *bufio.Reader, promtText string) string {
	fmt.Print(promtText)
	userInput, _ := reader.ReadString('\n')
	//userInput = strings.Replace(userInput, "\n", " ", -1)
	return userInput
}
func getProduct() *Product {
	fmt.Println("Please Enter the product data")

	reader := bufio.NewReader(os.Stdin)
	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Title: ")
	dscptnInput := readUserInput(reader, "Description: ")
	priceInput := readUserInput(reader, "Price: ")
	priceValue, _ := strconv.ParseFloat(priceInput, 64)
	product := NewProduct(idInput, titleInput, dscptnInput, priceValue)
	return product
}
func main() {
	createdPoduct := getProduct()
	createdPoduct.printData()
}

// firstProduct := Product{
// 	"first-product",
// 	"A book",
// 	"A nice Book ",
// 	10.99,
// }

// secondProduct := NewProduct(
// 	"2nd product",
// 	"A carpet", "nice carpet",
// 	11.99)
// fmt.Println(firstProduct)
// fmt.Println(*secondProduct)
// firstProduct.printData()
// secondProduct.printData()
