package main

import "fmt"

func main() {
	var stock []string
	for i := 0; i < 3; i++ {
		stock = append(stock, "r")
		stock = append(stock, "g")
		stock = append(stock, "y")
		stock = append(stock, "b")
	}
	for i := 0; i < 2; i++ {
		stock = append(stock, "b")
		stock = append(stock, "w")
	}
	fmt.Print(stock)

	p1Hand := map[string]int{"r": 0, "g": 0, "y": 0, "b": 0}
	p2Hand := map[string]int{"r": 0, "g": 0, "y": 0, "b": 0}
	fmt.Println(p1Hand)
	fmt.Println(p2Hand)
}
