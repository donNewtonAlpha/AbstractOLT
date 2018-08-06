package models

import (
	"fmt"

	"github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"
)

var chassisMap map[string]*abstractOLTModel.Chassis

/*
GetChassisMap return the chassis map singleton
*/
func GetChassisMap() *map[string]*abstractOLTModel.Chassis {
	if chassisMap == nil {
		fmt.Println("chassisMap was nil")
		chassisMap = make(map[string]*abstractOLTModel.Chassis)

	}
	fmt.Printf("chassis map %v\n", chassisMap)
	return &chassisMap
}
