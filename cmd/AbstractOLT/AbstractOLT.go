package main

import (
	"fmt"

	"github.com/donNewtonAlpha/AbstractOLT/models"
)

func main() {
	fmt.Println("AbstractOLT")
	var chassis models.Chassis
	var slots [16]models.Slot
	for i := 0; i < 16; i++ {
		slot := models.Slot{Parent: &chassis}
		fmt.Printf("slot %d\n", i)
		slots[i] = slot
	}
	chassis.Slots = slots
}
