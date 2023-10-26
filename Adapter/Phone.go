package Adapter

import "fmt"

type Phone interface {
	Charge(charger PhoneCharger)
}

type HuaweiPhone struct {
}

func NewHuaweiPhone() *HuaweiPhone {
	return &HuaweiPhone{}
}

func (h *HuaweiPhone) Charge(charger PhoneCharger) {
	fmt.Println("华为手机准备开始充电...")
	charger.OutPut5V()
}
