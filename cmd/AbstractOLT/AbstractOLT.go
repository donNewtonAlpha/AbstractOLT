package main

import (
	"fmt"

	"github.com/donNewtonAlpha/AbstractOLT/models"
	"github.com/donNewtonAlpha/AbstractOLT/internal/pkg/chassisSerialize"
)

func main() {
	fmt.Println("AbstractOLT")
	var chassis models.Chassis
	var slots [16]models.Slot
	for i := 0; i < 16; i++ {
		slot := models.Slot{Parent: &chassis, DeviceID: "DEVICE", Hostname: "HOSTNAME"}
		fmt.Printf("slot %d\n", i)
		slots[i] = slot
	}
	chassis.Slots = slots
	chassis.CLLI = "CLLI STRING"

	fmt.Println("Hello!")

	bytes, err := chassisSerialize.Serialize(chassis)
	fmt.Println(err)
	fmt.Println(string(bytes))
}
