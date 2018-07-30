package models

type Port struct {
	Number int
	Onts   [64]Ont
	Parent *Slot
}
