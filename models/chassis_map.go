package models

import (
	"fmt"

	"github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"
	"github.com/donNewtonAlpha/AbstractOLT/models/physicalModel"
)

var chassisMap map[string]*abstractOLTModel.Chassis
var abstractChassisMap map[string]*abstractOLTModel.Chassis

/*
GetPhyChassisMap return the chassis map singleton
*/
func GetPhyChassisMap() *map[string]*abstractOLTModel.Chassis {
	if chassisMap == nil {
		fmt.Println("chassisMap was nil")
		chassisMap = make(map[string]*physicalModel.Chassis)

	}
	fmt.Printf("chassis map %v\n", chassisMap)
	return &chassisMap
}

/*
GetAbstractChassisMap return the chassis map singleton
*/
func GetAbstractChassisMap() *map[string]*abstractOLTModel.Chassis {
	if abstractChassisMap == nil {
		fmt.Println("chassisMap was nil")
		abstractChassisMap = make(map[string]*abstractOLTModel.Chassis)

	}
	fmt.Printf("chassis map %v\n", chassisMap)
	return &chassisMap
}
