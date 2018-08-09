package physicalModel

import "net"

/*
Represents an arbitrary OLT linecard
*/
type OLT interface {
	GetCLLI() string
	GetHostname() string
	GetAddress() net.TCPAddr
	GetNumber() int
	GetPorts() []PONPort
	GetParent() *Chassis
	GetDataSwitchPort() int
}

/*
A basic representation of an OLT which fulfills the above interface,
and can be used in other OLT implementations
*/
type SimpleOLT struct {
	CLLI           string
	Hostname       string
	Address        net.TCPAddr
	Number         int
	Ports          []PONPort
	Parent         *Chassis `json:"-"`
	DataSwitchPort int
}

func (s *SimpleOLT) GetCLLI() string {
	return s.CLLI
}

func (s *SimpleOLT) GetHostname() string {
	return s.Hostname
}

func (s *SimpleOLT) GetAddress() net.TCPAddr {
	return s.Address
}

func (s *SimpleOLT) GetNumber() int {
	return s.Number
}

func (s *SimpleOLT) GetPorts() []PONPort {
	return s.Ports
}

func (s *SimpleOLT) GetParent() *Chassis {
	return s.Parent
}

func (s *SimpleOLT) GetDataSwitchPort() int {
	return s.DataSwitchPort
}
