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

func (p *Player) getHealth() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.health
}

func (p *Player) takeDamage(damage int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.health -= damage
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)

	for {
		fmt.Printf("player health: %d\r", p.getHealth())

		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)

	for {
		p.takeDamage(rand.Intn(30))
		if p.getHealth() <= 0 {
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
