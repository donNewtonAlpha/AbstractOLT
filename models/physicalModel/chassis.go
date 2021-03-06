package physicalModel

import "net"

/*
Chassis is a model that takes up to 16 discreet OLT chassis as if it is a 16 slot OLT chassis
*/

type Chassis struct {
	CLLI         string
	VCoreAddress net.TCPAddr
	Dataswitch   DataSwitch
	Linecards    []OLT
}
