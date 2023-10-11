package abstractFactory

import "fmt"

// Dole
type DoleFruitFactory struct{}

func (d *DoleFruitFactory) myAspect() {
	fmt.Println("dole aspect...")
}

func (d *DoleFruitFactory) CreateStrawberry() Strawberry {
	d.myAspect()
	defer d.myAspect()

	return &GoodFarmerStrawberry{
		brand: "dole",
	}
}

func (d *DoleFruitFactory) CreateLemon() Lemon {
	d.myAspect()
	defer d.myAspect()

	return &GoodFarmerLemon{
		brand: "dole",
	}

}
