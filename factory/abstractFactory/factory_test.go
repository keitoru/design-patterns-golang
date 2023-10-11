package abstractFactory

import "testing"

func Test_factory(t *testing.T) {
	// 尝尝佳农品牌的水果
	goodFarmerFactory := GoodFarmerFruitFactory{}
	goodFarmerStrawberry := goodFarmerFactory.CreateStrawberry()
	goodFarmerStrawberry.SweetAttack()

	goodFarmerLemon := goodFarmerFactory.CreateLemon()
	goodFarmerLemon.AcidAttack()

	// 尝尝都乐品牌的水果
	doleFactory := DoleFruitFactory{}
	doleStrawberry := doleFactory.CreateStrawberry()
	doleStrawberry.SweetAttack()

	doleLemon := doleFactory.CreateLemon()
	doleLemon.AcidAttack()
}
