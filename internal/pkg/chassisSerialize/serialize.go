package chassisSerialize

import (
	"encoding/json"

	"github.com/donNewtonAlpha/AbstractOLT/models"
)


func Serialize(chassis models.Chassis) ([]byte, error) {
	return json.Marshal(chassis)
}

func Deserialize(jsonData []byte) (models.Chassis, error) {
	var chassis models.Chassis
	err := json.Unmarshal(jsonData, &chassis)

	for _, slot := range chassis.Slots {
		slot.Parent = &chassis
		for _, port := range slot.Ports {
			port.Parent = &slot
			for _, ont := range port.Onts {
				ont.Parent = &port
			}
		}
	}

	return chassis, err
}
