package models

/*
Slot models a collection of PON ports (likely a single chassis) as if it is a
line card within a chassis
*/
type Slot struct {
	DeviceID string
	Hostname string
	Address  TCPAddr
	Number   int
	Ports    [16]Port
	Parent   *Chassis
}
