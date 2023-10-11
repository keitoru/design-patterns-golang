package simpleFactory

import "fmt"

type Fruit interface {
	Eat()
}

// Orange
type Orange struct {
	name string
}

func NewOrange(name string) Fruit {
	return &Orange{
		name: name,
	}
}

func (o *Orange) Eat() {
	fmt.Printf("i am orange: %s, i am about to be eaten...\n", o.name)
}

// Strawberry
type Strawberry struct {
	name string
}

func NewStrawberry(name string) Fruit {
	return &Strawberry{
		name: name,
	}
}

func (s *Strawberry) Eat() {
	fmt.Printf("i am strawberry: %s, i am about to be eaten...\n", s.name)
}

// Cherry
type Cherry struct {
	name string
}

func NewCherry(name string) Fruit {
	return &Cherry{
		name: name,
	}
}

func (c *Cherry) Eat() {
	fmt.Printf("i am cherry: %s, i am about to be eaten...\n", c.name)
}
