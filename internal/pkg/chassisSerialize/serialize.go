package chassisSerialize

import (
	"encoding/json"

	"github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"
)


func Serialize(chassis *abstractOLTModel.Chassis) ([]byte, error) {
	return json.Marshal(chassis)
}

func Deserialize(jsonData []byte) (*abstractOLTModel.Chassis, error) {
	var chassis abstractOLTModel.Chassis
	err := json.Unmarshal(jsonData, &chassis)

	for i := 0; i < len(chassis.Slots); i++ {
		var slot *abstractOLTModel.Slot
		slot = &chassis.Slots[i]
		slot.Parent = &chassis
		for j := 0; j < len(slot.Ports); j++ {
			var port *abstractOLTModel.Port
			port = &slot.Ports[j]
			port.Parent = slot
			for k := 0; k < len(port.Onts); k++ {
				port.Onts[k].Parent = port
			}
		}
	}

	return &chassis, err
}
