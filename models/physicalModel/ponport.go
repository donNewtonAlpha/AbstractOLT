package physicalModel

/*
Port represents a single PON port on the OLT chassis
*/
type PONPort struct {
	Number   int
	DeviceID string
	Onts     [64]Ont
	Parent   *Slot    `json:"-"`
}
