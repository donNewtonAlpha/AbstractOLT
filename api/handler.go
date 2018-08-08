package api

import (
	"fmt"
	// "net"

	"github.com/donNewtonAlpha/AbstractOLT/models"
	"github.com/donNewtonAlpha/AbstractOLT/models/physicalModel"
	context "golang.org/x/net/context"
)

/*
Server instance of the grpc server
*/
type Server struct {
}

/*
CreateChassis - allocates a new Chassis struct and stores it in chassisMap
*/
func (s *Server) CreateChassis(ctx context.Context, in *AddChassisMessage) (*AddChassisReturn, error) {
	phyChassisMap := models.GetPhyChassisMap()
	absChassisMap := models.GetAbstractChassisMap()
	clli := in.GetCLLI()

	chassis := (*phyChassisMap)[clli]
	if chassis != nil {
		return &AddChassisReturn{DeviceID: chassis.CLLI}, nil
	}
	abstractChassis := modelUtils.generateChassis(clli)
	phyChassis := physicalModel.Chassis{CLLI: clli}
	fmt.Printf("new chassis %v\n", phyChassis)
	(*phyChassisMap)[clli] = &phyChassis
	(*abstractChassis)[clli] = &abstractChassis
	return &AddChassisReturn{DeviceID: newChassis.CLLI}, nil
}
