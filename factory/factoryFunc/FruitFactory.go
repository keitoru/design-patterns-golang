package factoryFunc

type FruitFactory interface {
	CreateFruit() Fruit
}

// OrangeFactory
type OrangeFactory struct {
}

func NewOrangeFactory() FruitFactory {
	return &OrangeFactory{}
}

func (o *OrangeFactory) CreateFruit() Fruit {
	return NewOrange("")
}

// CherryFactory
type CherryFactory struct {
}

func NewCherryFactory() FruitFactory {
	return &CherryFactory{}
}

func (c *CherryFactory) CreateFruit() Fruit {
	return NewCherry("")
}

// StrawberryFactory
type StrawberryFactory struct {
}

func NewStrawberryFactory() FruitFactory {
	return &StrawberryFactory{}
}

func (s *StrawberryFactory) CreateFruit() Fruit {
	return NewStrawberry("")
}
