package models

import "net"

/*
Chassis is a model that takes up to 16 discreet OLT chassis as if it is a 16 slot OLT chassis
*/

type Chassis struct {
	VCoreAddress net.TCPAddr
	CLLI         string
	Slots        [16]Slot
}
