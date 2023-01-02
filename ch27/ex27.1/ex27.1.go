package main

import "fmt"

// 연습문제 3

type Attacker interface {
	Name() string
}

type DamageTaker interface {
	DealDamage(attacker Attacker, damage int)
}

type Player struct {
	name string
}

type Monster struct {
	hp int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(damageTaker DamageTaker) {
	damageTaker.DealDamage(p, 100)
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다.")
	}
}

func main() {
	player := &Player{"mz"}
	monster := &Monster{50}

	player.Attack(monster)
}
