package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	// case 1
	//actor := Actor{name, hp, speed}
	//return &actor

	// case 2
	//return &Actor{name, hp, speed}

	// case 3
	a := new(Actor)
	(*a).Name = name
	a.HP = hp
	a.Speed = speed
	return a
}

func main() {
	actor := NewActor("금도끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
