package Adapter

import "fmt"

type PhoneCharger interface {
	OutPut5V()
}

// huawei
type HuaweiCharger struct {
}

func NewHuaweiCharger() *HuaweiCharger {
	return &HuaweiCharger{}
}

func (h *HuaweiCharger) OutPut5V() {
	fmt.Println("华为手机充电器输出 5V 电压...")
}

// xiaomi
type XiaomiCharger struct {
}

func NewXiaomiCharger() *XiaomiCharger {
	return &XiaomiCharger{}
}

func (x *XiaomiCharger) OutPut5V() {
	fmt.Println("小米手机充电器输出 5V 电压...")
}

// macbook
type MacbookCharger struct {
}

func NewMacbookCharger() *MacbookCharger {
	return &MacbookCharger{}
}

func (m *MacbookCharger) OutPut28V() {
	fmt.Println("苹果笔记本充电器输出 28V 电压...")
}
