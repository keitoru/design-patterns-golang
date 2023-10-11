package simpleFactory

import "testing"

func Test_factory(t *testing.T) {
	fruitFactory := NewFruitFactory()

	orange, _ := fruitFactory.CreateFruit("orange")
	orange.Eat()

	cherry, _ := fruitFactory.CreateFruit("cherry")
	cherry.Eat()

	watermelon, err := fruitFactory.CreateFruit("watermelon")
	if err != nil {
		t.Error(err)
		return
	}
	watermelon.Eat()

}
