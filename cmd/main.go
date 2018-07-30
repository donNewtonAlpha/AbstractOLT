package main

import (
	"fmt"

	"github.com/donNewtonAlpha/AbstractOLT/models"
)

func main() {
	fmt.Println("AbstractOLT")
	var chassis models.Chassis
	var slots [8]models.Slot
	for i := 0; i < 8; i++ {
		slot := models.Slot{Parent: &chassis}
		slots[i] = slot
		chassis.Slots = slots
	}

}
