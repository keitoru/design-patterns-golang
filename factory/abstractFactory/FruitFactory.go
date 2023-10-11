package abstractFactory

type FruitFactory interface {
	CreateStrawberry() Strawberry
	CreateLemon() Lemon
}
