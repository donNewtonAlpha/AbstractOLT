package physicalModel

import "net"

// Represents an Edgecore linecard

type Edgecore struct {
	DeviceID string
	Hostname string
	Address  net.TCPAddr
	Number   int
	Ports    [16]Port
	Parent   *Chassis    `json:"-"`
	DataSwitchPort  int
}
