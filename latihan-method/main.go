package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

func (p Player) AddScore(score int) {
	if p.Name == "John" {
		p.Score = p.Score + score
		fmt.Println(p.Name)
		fmt.Println(p.Score)
	}
}

func (p Player) DisplayInfo() {
	fmt.Println("Nama pemain :", p.Name, ", Skor :", p.Score)
}

func main() {
	player := Player{
		Name:  "John",
		Score: 0,
	}
	player.AddScore(10)
	player.AddScore(5)
	player.DisplayInfo()
}
