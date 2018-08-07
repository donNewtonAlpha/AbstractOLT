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
	chassisMap := models.GetPhyChassisMap()
	clli := in.GetCLLI()
	chassis := (*chassisMap)[clli]
	if chassis != nil {
		return &AddChassisReturn{DeviceID: chassis.CLLI}, nil
	}
	newChassis := physicalModel.Chassis{CLLI: clli}
	fmt.Printf("new chassis %v\n", newChassis)
	(*chassisMap)[clli] = &newChassis
	return &AddChassisReturn{DeviceID: newChassis.CLLI}, nil
}
