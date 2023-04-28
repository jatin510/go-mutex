package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	mu     sync.RWMutex
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
		p.mu.RLock()

		fmt.Printf("player health: %d\r", p.health)

		p.mu.RUnlock()
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)

	for {
		p.mu.Lock()

		p.health -= rand.Intn(40)
		if p.health <= 0 {
			fmt.Println("GAME OVER")
			break
		}

		p.mu.Unlock()
		<-ticker.C
	}
}

func main() {
	player := NewPlayer()

	go startUILoop(player)
	startGameLoop(player)
}
