package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	name   string
	inHand map[string]int
}

func main() {
	var stock []string
	stock = stockInit()
	fmt.Print(stock, "\n")

	p1 := player{name: "p1"}
	p1.inHandInit()

	p2 := player{name: "p2"}
	p2.inHandInit()

	fmt.Println(p1.inHand, "\n", p2.inHand, "\n")
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

func (p *player) inHandInit() {
	p.inHand = map[string]int{"r": 0, "g": 0, "y": 0, "b": 0}
}

//func judgeWin(player) {}
