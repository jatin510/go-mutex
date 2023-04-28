package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	health int
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)

	for {
		fmt.Printf("player health: %d\r", p.health)
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)

	for {
		p.health -= rand.Intn(40)
		if p.health <= 0 {
			fmt.Println("GAME OVER")
			break
		}

		<-ticker.C
	}
}

func main() {
	player := NewPlayer()

	go startUILoop(player)
	startGameLoop(player)
}
