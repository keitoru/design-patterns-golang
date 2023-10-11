package factoryFunc

import "testing"

func Test_factory(t *testing.T) {
	orangeFactory := NewOrangeFactory()
	orange := orangeFactory.CreateFruit()
	orange.Eat()

	// 来颗樱桃
	cherryFactory := NewCherryFactory()
	cherry := cherryFactory.CreateFruit()
	cherry.Eat()

}
