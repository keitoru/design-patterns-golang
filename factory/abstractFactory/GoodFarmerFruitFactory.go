package abstractFactory

import (
	"fmt"
)

// GoodFarmer
type GoodFarmerFruitFactory struct{}

func (g *GoodFarmerFruitFactory) myAspect() {
	fmt.Println("good farmer aspect...")
}

func (g *GoodFarmerFruitFactory) CreateStrawberry() Strawberry {
	g.myAspect()
	defer g.myAspect()

	return &GoodFarmerStrawberry{
		brand: "good farmer",
	}
}

func (g *GoodFarmerFruitFactory) CreateLemon() Lemon {
	g.myAspect()
	defer g.myAspect()

	return &GoodFarmerLemon{
		brand: "good farmer",
	}

}
