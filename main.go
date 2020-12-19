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

func (p *player) inHandInit() {
	p.inHand = map[string]int{"r": 0, "g": 0, "y": 0, "b": 0}
}

func (p *player) getCard(c string) {
	p.inHand[c] += 1
}

func (p *player) judgeWin() bool {
	gameEnd := false
	if p.inHand["r"]*p.inHand["g"]*p.inHand["y"]*p.inHand["b"] > 0 {
		fmt.Println(p.name, " is WIN!")
		gameEnd = true
	}
	return gameEnd
}

func main() {
	var stock []string
	var turn string
	turn = "p1"

	stock = stockInit()
	fmt.Print(stock, "\n")

	p1 := player{name: "p1"}
	p1.inHandInit()

	p2 := player{name: "p2"}
	p2.inHandInit()

	fmt.Println(p1.inHand)
	fmt.Println(p2.inHand)

	var nowPlayer *player
	var otherPlayer *player
	var gameEnd bool

game:
	for {
		if turn == "p1" {
			nowPlayer = &p1
			otherPlayer = &p2
			turn = "p2"
		} else {
			nowPlayer = &p2
			otherPlayer = &p1
			turn = "p1"
		}

		card := stock[0]
		stock = stock[1:]
		if card != "l" && card != "w" {
			nowPlayer.getCard(card)
		} else if card == "l" {
			tranceferAllCard(nowPlayer, otherPlayer)
		} else {
			tranceferAllCard(otherPlayer, nowPlayer)
		}
		gameEnd = nowPlayer.judgeWin()

		if gameEnd {
			break game
		}
		fmt.Println("++++++++++")
		fmt.Println(stock)
		fmt.Println(p1.inHand)
		fmt.Println(p2.inHand)
	}
	fmt.Println("==========")
	fmt.Println(stock)
	fmt.Println(p1.inHand)
	fmt.Println(p2.inHand)
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

func tranceferAllCard(fromPlayer *player, toPlayer *player) {
	for key := range fromPlayer.inHand {
		toPlayer.inHand[key] += fromPlayer.inHand[key]
		fromPlayer.inHand[key] = 0
	}
}
