package models

type Slot struct {
	Number int
	Ports  [8]Port
	Parent *Chassis
}
