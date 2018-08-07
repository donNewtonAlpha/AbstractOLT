package chassisUtils

import "github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"


func generateChassis(CLLI string) (*abstractOLTModel.Chassis) {
	chassis := abstractOLTModel.Chassis{CLLI: CLLI}

	var slots [16]abstractOLTModel.Slot
	for i := 0; i < 16; i++ {
		slots[i] = generateTestSlot(i, &chassis)
	}

	chassis.Slots = slots
	return &chassis
}

func generateSlot(n, int, c *abstractOLTModel.Chassis) (abstractOLTModel.Slot) {
	slot := abstractOLTModel.Slot{Number: n, Parent: c}

	var ports [16]abstractOLTModel.Port
	for i := 0; i < 16; i++ {
		ports[i] = generateTestPort(i, &slot)
	}

	slot.Ports = ports
	return slot
}

func generatePort(n int, s *abstractOLTModel.Slot) (abstractOLTModel.Port) {
	port := abstractOLTModel.Port{Number: n, Parent: s}

	var onts [64]abstractOLTModel.Ont
	for i := 0; i < 64; i++ {
		onts[i] = abstractOLTModel.Ont{Number: i, Svlan: calculateSvlan(s.Number, n, i),
					Cvlan: calculateSvlan(s.Number, n, i), Parent: &port}
	}

	port.Onts = onts
	return port
}


func calculateCvlan(slot int, port int, ont int) int {
	ONT_PORT_OFFSET := 120 // Max(ONT_SLOT) * Max(ONT_PORT) = 10 * 12 = 120
	ONT_SLOT_OFFSET := 12 //= Max(ONT_PORT) = 12
	VLAN_OFFSET := 1 //(VID 1 is reserved)

	CVID := ((ont - 1) % 32) * ONT_PORT_OFFSET +
			 	(slot - 1) * ONT_SLOT_OFFSET +  port + VLAN_OFFSET

	return CVID
}

func calculateSvlan(slot int, port int, ont int) int {
	LT_SLOT_OFFSET := 16
	VLAN_GAP := 288 // Max(LT_SLOT) * Max(LT_SLOT_OFFSET) = 18 * 16 = 288
	VLAN_OFFSET := 1  //(VID 1 is reserved)

	SVID := ((slot - 1) * LT_SLOT_OFFSET + port) + ((ont – 1) / 32) * VLAN_GAP] + VLAN_OFFSET

	return SVID
}



func (chassis *Chassis) NextPort() (*Port, error) {
	info := &chassis.AllocInfo

	if info.outOfPorts {
		return nil, errors.New("Abstract chassis out of ports")
	}

	nextPort := &chassis.Slots[info.slot].Ports[info.port]

	info.port++
	if info.port == MAX_PORTS {
		info.port = 0
		info.slot++
		if info.slot == MAX_SLOTS {
			info.slot = 0
			info.outOfPorts = true
		}
	}

	return nextPort, nil
}