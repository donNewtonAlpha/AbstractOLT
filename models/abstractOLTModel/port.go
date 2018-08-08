package abstractOLTModel

import "github.com/donNewtonAlpha/AbstractOLT/models/physicalModel"

/*
Port represents a single PON port on the OLT chassis
*/
type Port struct {
	Number   int
	// DeviceID string
	Onts     [64]Ont
	PhysPort *physicalModel.PONPort
	Parent   *Slot    `json:"-"`
}
