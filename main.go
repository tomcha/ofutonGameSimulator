package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var stock []string
	stock = stockInit()
	fmt.Print(stock, "\n")

	p1Hand := playerInit()
	p2Hand := playerInit()
	fmt.Println(p1Hand, "\n", p2Hand, "\n")
}

func stockInit() []string {
	var s []string
	for i := 0; i < 3; i++ {
		s = append(s, "r")
		s = append(s, "g")
		s = append(s, "y")
		s = append(s, "b")
	}
	for i := 0; i < 2; i++ {
		s = append(s, "l")
		s = append(s, "w")
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}

func playerInit() map[string]int {
	return map[string]int{"r": 0, "g": 0, "y": 0, "b": 0}
}
