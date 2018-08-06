package api

import (
	"fmt"
	"net"

	"github.com/donNewtonAlpha/AbstractOLT/models"
	"github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"
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
	chassisMap := models.GetChassisMap()
	clli := in.GetCLLI()
	chassis := (*chassisMap)[clli]
	if chassis != nil {
		return &AddChassisReturn{DeviceID: chassis.CLLI}, nil
	}
	vCoreAddress := net.TCPAddr{IP: net.ParseIP(in.GetVCoreIP()), Port: int(in.GetVCorePort())}
	newChassis := abstractOLTModel.Chassis{CLLI: clli, VCoreAddress: vCoreAddress}
	fmt.Printf("new chassis %v\n", newChassis)
	(*chassisMap)[clli] = &newChassis
	return &AddChassisReturn{DeviceID: newChassis.CLLI}, nil
}
