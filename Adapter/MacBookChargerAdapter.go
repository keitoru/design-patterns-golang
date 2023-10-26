package Adapter

import "fmt"

type MacBookChargerAdapter struct {
	core *MacbookCharger
}

func NewMacBookChargerAdapter(charger *MacbookCharger) *MacBookChargerAdapter {
	return &MacBookChargerAdapter{
		core: charger,
	}
}

func (m *MacBookChargerAdapter) OutPut5V() {
	m.core.OutPut28V()
	fmt.Println("适配器将输出电压调整为 5V...")
}
