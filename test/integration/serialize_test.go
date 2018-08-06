package integration

import (
	"testing"
	"net"

	"github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"
	"github.com/donNewtonAlpha/AbstractOLT/internal/pkg/chassisSerialize"
)

func TestSerialize(t *testing.T) {
	chassis1 := generateTestChassis()
	bytes1, err1 := chassisSerialize.Serialize(chassis1)
	chassis2, err2 := chassisSerialize.Deserialize(bytes1)
	bytes2, err3 := chassisSerialize.Serialize(chassis2)
	chassis3, err4 := chassisSerialize.Deserialize(bytes2)

	ok(t, err1)
	ok(t, err2)
	ok(t, err3)
	ok(t, err4)
	equals(t, chassis1, chassis3)
	equals(t, chassis3.Slots[2].Parent, chassis3)
	equals(t, chassis3.Slots[15].Ports[8].Parent, &chassis3.Slots[15])
	equals(t, chassis3.Slots[0].Ports[10].Onts[15].Parent, &chassis3.Slots[0].Ports[10])
}


func generateTestChassis() (*abstractOLTModel.Chassis) {
	addr := net.TCPAddr{IP: net.IPv4(1,2,3,4), Port: 500, Zone: "VCore ZONE"}
	chassis := abstractOLTModel.Chassis{VCoreAddress: addr, CLLI: "CLLI STRING"}

	var slots [16]abstractOLTModel.Slot
	for i := 0; i < 16; i++ {
		slots[i] = generateTestSlot(i, &chassis)
	}

	chassis.Slots = slots
	return &chassis
}

func generateTestSlot(n int, c *abstractOLTModel.Chassis) (abstractOLTModel.Slot) {
	addr := net.TCPAddr{IP: net.IPv4(1,2,3,byte(n)), Port: 400 + n, Zone: "Slot " + string(n) + "Zone"}
	slot := abstractOLTModel.Slot{DeviceID: "Device Slot " + string(n), Hostname: "Host " + string(n), 
				Address: addr, Number: n, Parent: c}

	var ports [16]abstractOLTModel.Port
	for i := 0; i < 16; i++ {
		ports[i] = generateTestPort(16*n + i, &slot)
	}

	slot.Ports = ports
	return slot
}

func generateTestPort(n int, s *abstractOLTModel.Slot) (abstractOLTModel.Port) {
	port := abstractOLTModel.Port{Number: n, DeviceID: "Device Port " + string(n), Parent: s}

	var onts [64]abstractOLTModel.Ont
	for i := 0; i < 64; i++ {
		j := n*64 + i
		onts[i] = abstractOLTModel.Ont{Number: j, Svlan: j*10, Cvlan: j*10 + 5, Parent: &port}
	}

	port.Onts = onts
	return port
}

