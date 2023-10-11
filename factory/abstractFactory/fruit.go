package abstractFactory

import "fmt"

// Strawberry
type Strawberry interface {
	SweetAttack()
}

// Lemon
type Lemon interface {
	AcidAttack()
}

// GoodFarmer
type GoodFarmerStrawberry struct {
	brand string
	Strawberry
}

func (g *GoodFarmerStrawberry) SweetAttack() {
	fmt.Println("sweet attack from ", g.brand)
}

type GoodFarmerLemon struct {
	brand string
	Lemon
}

func (g *GoodFarmerLemon) AcidAttack() {
	fmt.Println("acid attack from ", g.brand)
}

// Dole
type DoleStrawberry struct {
	brand string
	Strawberry
}

func (d *DoleStrawberry) SweetAttack() {
	fmt.Println("sweet attack from ", d.brand)
}

type DoleLemon struct {
	brand string
	Lemon
}

func (d *DoleLemon) AcidAttack() {
	fmt.Println("acid attack from ", d.brand)
}

// Zespri
type ZespriStrawBerry struct {
	brand string
	Strawberry
}

func (z *ZespriStrawBerry) SweetAttack() {
	fmt.Println("sweet attack from ", z.brand)
}

type ZespriLemon struct {
	brand string
	Lemon
}

func (z *ZespriLemon) AcidAttack() {
	fmt.Println("acid attack from ", z.brand)
}
