package abstractOLTModel

const MAX_SLOTS int = 16
const MAX_PORTS int = 16

/*
Chassis is a model that takes up to 16 discreet OLT chassis as if it is a 16 slot OLT chassis
*/
type Chassis struct {
	CLLI      string
	Slots     [16]Slot
	AllocInfo PortAllocationInfo
}

type PortAllocationInfo struct {
	// Current info on next port to be allocated
	slot       int
	port       int
	outOfPorts bool
}
