package models

type Port struct {
	Number   int
	DeviceID string
	Onts     [64]Ont
	Parent   *Slot
}
