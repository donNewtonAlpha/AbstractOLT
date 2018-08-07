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
func (s *Server) CreatePhyChassis(ctx context.Context, in *AddPhyChassisMessage) (*AddPhyChassisReturn, error) {
	chassisMap := models.GetPhyChassisMap()
	clli := in.GetCLLI()
	chassis := (*chassisMap)[clli]
	if chassis != nil {
		return &AddPhyChassisReturn{DeviceID: chassis.CLLI}, nil
	}
	vCoreAddress := net.TCPAddr{IP: net.ParseIP(in.GetVCoreIP()), Port: int(in.GetVCorePort())}
	newChassis := abstractOLTModel.Chassis{CLLI: clli, VCoreAddress: vCoreAddress}
	fmt.Printf("new chassis %v\n", newChassis)
	(*chassisMap)[clli] = &newChassis
	return &AddPhyChassisReturn{DeviceID: newChassis.CLLI}, nil
}
func (s *Server) CreateAbstractChassis(ctx context.Context, in *AddAbstractChassisMessage) (*AddAbstractChassisReturn, error) {
	chassisMap := models.GetAbstractChassisMap()
	clli := in.GetCLLI()
	chassis := (*chassisMap)[clli]
	if chassis != nil {
		return &AddAbstractChassisReturn{DeviceID: chassis.CLLI}, nil
	}
	vCoreAddress := net.TCPAddr{IP: net.ParseIP(in.GetVCoreIP()), Port: int(in.GetVCorePort())}
	newChassis := abstractOLTModel.Chassis{CLLI: clli, VCoreAddress: vCoreAddress}
	fmt.Printf("new chassis %v\n", newChassis)
	(*chassisMap)[clli] = &newChassis
	return &AddChassisReturn{DeviceID: newChassis.CLLI}, nil
}
