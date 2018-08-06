package chassisSerialize

import (
	"encoding/json"

	"github.com/donNewtonAlpha/AbstractOLT/models"
)


func Serialize(chassis *models.Chassis) ([]byte, error) {
	return json.Marshal(chassis)
}

func Deserialize(jsonData []byte) (*models.Chassis, error) {
	var chassis models.Chassis
	err := json.Unmarshal(jsonData, &chassis)

	for i := 0; i < len(chassis.Slots); i++ {
		var slot *models.Slot
		slot = &chassis.Slots[i]
		slot.Parent = &chassis
		for j := 0; j < len(slot.Ports); j++ {
			var port *models.Port
			port = &slot.Ports[j]
			port.Parent = slot
			for k := 0; k < len(port.Onts); k++ {
				port.Onts[k].Parent = port
			}
		}
	}

	return &chassis, err
}
