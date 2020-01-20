package factoryMethod

import (
	"log"
)

type Creater interface {
	CreatePerson(action string, name string, age int) Person
}

type Person interface {
	GetAction() string
	GetName() string
	GetAge() int
}

type ConcreteCreator struct {
}

func NewCreater() *ConcreteCreator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreatePerson(action string, name string, age int) Person {
	var person Person

	switch action {
	case "Student":
		person = &Student{action, name, age}
	case "Businessman":
		person = &Businessman{action, name, age, 1000000}
	case "Programmer":
		person = &Programmer{action, name, age, 0}
	default:
		log.Fatalln("Unknown Action")
	}

	return person
}

type Student struct {
	action string
	name   string
	age    int
}

func (s *Student) GetAge() int {
	return s.age
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetAction() string {
	return s.action
}

type Businessman struct {
	action string
	name   string
	age    int
	money  int
}

func (b *Businessman) GetAction() string {
	return b.action
}

func (b *Businessman) GetName() string {
	return b.name
}

func (b *Businessman) GetAge() int {
	return b.age
}

func (b *Businessman) TakeMoney() int {
	b.money = b.money - 1000
	return b.money
}

type Programmer struct {
	action string
	name   string
	age    int
	wallet int
}

func (p *Programmer) GetAction() string {
	return p.action
}

func (p *Programmer) GetName() string {
	return p.name
}

func (p *Programmer) GetAge() int {
	return p.age
}

func (p *Programmer) GiveMoney() int {
	p.wallet++
	return p.wallet
}
