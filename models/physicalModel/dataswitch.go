package physicalModel

import "net"

/*
Chassis is a model that takes up to 16 discreet OLT chassis as if it is a 16 slot OLT chassis
*/

type DataSwitch struct {
	Ports        [32]*Edgecore
	Driver       string
    Ipv4Loopback net.TCPAddr
    Ipv4NodeSid  int
    isEdgeRouter bool
    Name         string
    OfId         string
    RouterMac    string
}
