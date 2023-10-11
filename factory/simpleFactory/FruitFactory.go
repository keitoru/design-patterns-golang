package simpleFactory

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type fruitCreator func(name string) Fruit

type FruitFactory struct {
	creators map[string]fruitCreator
}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{
		creators: map[string]fruitCreator{
			"orange":     NewOrange,
			"strawberry": NewStrawberry,
			"cherry":     NewCherry,
		},
	}
}

func (f *FruitFactory) CreateFruit(typ string) (Fruit, error) {
	fruitCreator, ok := f.creators[typ]
	if !ok {
		return nil, fmt.Errorf("fruit typ: %s is not supported yet", typ)
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	name := strconv.Itoa(r.Int())

	return fruitCreator(name), nil

}
