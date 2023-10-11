package Decorator

type Decorator Food

func NewDecorator(f Food) Decorator {
	return f
}

// 老干妈 LaoGanMaDecorator
type LaoGanMaDecorator struct {
	Decorator
}

func NewLaoGanMaDecorator(d Decorator) Decorator {
	return &LaoGanMaDecorator{
		Decorator: d,
	}
}

func (l *LaoGanMaDecorator) Eat() string {
	// 加入老干妈配料
	return "加入一份老干妈~..." + l.Decorator.Eat()
}

func (l *LaoGanMaDecorator) Cost() float32 {
	// 价格增加 0.5 元
	return 0.5 + l.Decorator.Cost()
}

// 火腿肠 HamSausageDecorator
type HamSausageDecorator struct {
	Decorator
}

func NewHamSausageDecorator(d Decorator) Decorator {
	return &HamSausageDecorator{
		Decorator: d,
	}
}

func (h *HamSausageDecorator) Eat() string {
	// 加入火腿肠配料
	return "加入一份火腿~..." + h.Decorator.Eat()
}

func (h *HamSausageDecorator) Cost() float32 {
	// 价格增加 1.5 元
	return 1.5 + h.Decorator.Cost()
}

// 煎蛋 FriedEggDecorator
type FriedEggDecorator struct {
	Decorator
}

func NewFriedEggDecorator(d Decorator) Decorator {
	return &FriedEggDecorator{
		Decorator: d,
	}
}

func (f *FriedEggDecorator) Eat() string {
	// 加入煎蛋配料
	return "加入一份煎蛋~..." + f.Decorator.Eat()
}

func (f *FriedEggDecorator) Cost() float32 {
	// 价格增加 1 元
	return 1 + f.Decorator.Cost()
}
