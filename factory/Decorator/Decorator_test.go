package Decorator

import (
	"fmt"
	"testing"
)

func Test_decorator(t *testing.T) {
	// 一碗干净的米饭
	rice := NewRice()
	res := rice.Eat()
	fmt.Println(res)

	// 一碗干净的面条
	noodle := NewNoodle()
	res = noodle.Eat()
	fmt.Println(res)

	// 米饭加个煎蛋
	rice = NewFriedEggDecorator(rice)
	res = rice.Eat()
	fmt.Println(res)

	// 面条加份火腿
	noodle = NewHamSausageDecorator(noodle)
	res = noodle.Eat()
	fmt.Println(res)

	// 米饭再分别加个煎蛋和一份老干妈
	rice = NewFriedEggDecorator(rice)
	rice = NewLaoGanMaDecorator(rice)
	res = rice.Eat()
	fmt.Println(res)
}
