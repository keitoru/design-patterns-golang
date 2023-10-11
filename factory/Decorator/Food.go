package Decorator

type Food interface {
	Eat() string
	Cost() float32
}

// rice
type Rice struct {
}

func NewRice() Food {
	return &Rice{}
}

func (r *Rice) Eat() string {
	return "一碗香喷喷的米饭..."
}

func (r *Rice) Cost() float32 {
	return 1
}

// noodle
type Noodle struct {
}

func NewNoodle() Food {
	return &Noodle{}
}

func (n *Noodle) Eat() string {
	return "嗦面ing...嗦~"
}

func (n *Noodle) Cost() float32 {
	return 1.5
}
