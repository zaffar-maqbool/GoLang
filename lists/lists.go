package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices for Dynamic Array")

	prices := []float64{103, 123, 312}
	fmt.Println(prices)

	appedndedPrices := append(prices, 5.99)
	fmt.Println(appedndedPrices)

	prices = prices[1:]
	fmt.Println(prices)
	discountedPrices := []float64{101, 80, 99, 20}
	prices = append(prices, discountedPrices...)
	fmt.Println(prices)

}

// func main() {
// 	var productNames [4]string = [4]string{"a book"}

// 	prices := [4]float64{10.99, 9.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)

// 	// accessing an element using index
// 	// fmt.Println(prices[1])
// 	// overiding value
// 	// prices[0] = 1999
// 	// fmt.Println(prices)

// 	// slice			// 0 1 2   excluding 3
// 	slicedPrice := prices[0:2]
// 	fmt.Println(slicedPrice)

// 	// highlightedPrices := slicedPrice[:1]
// 	// fmt.Println(highlightedPrices)
// 	// fmt.Println(len(slicedPrice), cap(slicedPrice))
// 	// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// type TemperatureData struct {
// 	d1 float64
// 	d2 float64
// 	d3 float64
// 	d4 float64
// 	d5 float64
// }
