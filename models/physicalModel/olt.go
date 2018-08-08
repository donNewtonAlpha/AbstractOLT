package physicalModel

import "net"

/*
Represents an arbitrary OLT linecard
*/
type Olt interface {
	GetDeviceID() string
	GetHostname() string
	GetAddress()  net.TCPAddr
	GetNumber()   int
	GetPorts()    []PONPort
	GetParent()   *Chassis
	GetDataSwitchPort()  int
}


/*
A basic representation of an OLT which fulfills the above interface,
and can be used in other Olt implementations
*/
type SimpleOlt struct {
	DeviceID string
	Hostname string
	Address  net.TCPAddr
	Number   int
	Ports    []PONPort
	Parent   *Chassis    `json:"-"`
	DataSwitchPort  int
}


func (s *SimpleOlt) GetDeviceID() string {
	return s.DeviceID
}

func (s *SimpleOlt) GetHostname() string {
	return s.Hostname
}

func (s *SimpleOlt) GetAddress() net.TCPAddr {
	return s.Address
}

func (s *SimpleOlt) GetNumber() int {
	return s.Number
}

func (s *SimpleOlt) GetPorts() []PONPort {
	return s.Ports
}

func (s *SimpleOlt) GetParent() *Chassis {
	return s.Parent
}

func (s *SimpleOlt) GetDataSwitchPort() int {
	return s.DataSwitchPort
}