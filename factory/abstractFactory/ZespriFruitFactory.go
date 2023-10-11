package abstractFactory

import "fmt"

type ZespriFruitFactory struct{}

func (z *ZespriFruitFactory) myAspect() {
	fmt.Println("zespri aspect...")
}

func (z *ZespriFruitFactory) CreateStrawberry() Strawberry {
	z.myAspect()
	defer z.myAspect()

	return &ZespriStrawBerry{
		brand: "zespri",
	}
}

func (z *ZespriFruitFactory) CreateLemon() Lemon {
	z.myAspect()
	defer z.myAspect()

	return &ZespriLemon{
		brand: "zespri",
	}
}
