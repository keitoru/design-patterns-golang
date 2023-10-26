package Adapter

import "testing"

func Test_adapter(t *testing.T) {
	phone := NewHuaweiPhone()
	phone.Charge(NewHuaweiCharger())
	phone.Charge(NewXiaomiCharger())

	macbookCharge := NewMacbookCharger()
	adapter := NewMacBookChargerAdapter(macbookCharge)
	phone.Charge(adapter)
}
