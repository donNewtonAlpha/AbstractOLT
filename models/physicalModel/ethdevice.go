package physicalModel

import "net"

/*
EthDevice represents a device with an ethernet port
*/
type EthDevice interface {
}

type EthPort struct {
	Connected *EthDevice
}
