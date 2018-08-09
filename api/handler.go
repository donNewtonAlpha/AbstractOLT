package api

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/donNewtonAlpha/AbstractOLT/internal/pkg/settings"
	"github.com/donNewtonAlpha/AbstractOLT/models"
	"github.com/donNewtonAlpha/AbstractOLT/models/abstractOLTModel"
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
	abstractChassis := abstractOLTModel.GenerateChassis(clli)
	phyChassis := physicalModel.Chassis{CLLI: clli}
	if settings.GetDebug() {
		output := fmt.Sprintf("%v", abstractChassis)
		formatted := strings.Replace(output, "{", "\n{", -1)
		log.Printf("new chassis %s\n", formatted)
	}
	(*phyChassisMap)[clli] = &phyChassis
	(*absChassisMap)[clli] = abstractChassis
	return &AddChassisReturn{DeviceID: clli}, nil
}

/*
AddOLTChassis adds an OLT chassis/line card to the Physical chassis
*/
func (s *Server) CreateOLTChassis(ctx context.Context, in *AddOLTChassisMessage) (*AddOLTChassisReturn, error) {
	fmt.Printf(" CreateOLTChassis %v \n", *in)
	phyChassisMap := models.GetPhyChassisMap()
	clli := in.GetCLLI()
	chassis := (*phyChassisMap)[clli]
	if chassis == nil {
	}
	oltType := in.GetType()
	address := net.TCPAddr{IP: net.ParseIP(in.GetSlotIP()), Port: int(in.GetSlotPort())}
	sOlt := physicalModel.SimpleOLT{CLLI: clli, Hostname: in.GetHostname(), Address: address}

	var olt physicalModel.OLT
	switch oltType {
	case AddOLTChassisMessage_edgecore:
		olt = physicalModel.CreateEdgecore(&sOlt)
	case AddOLTChassisMessage_adtran:
	case AddOLTChassisMessage_tibit:
	}

	err := AddCard(chassis, olt)
	if err != nil {
		//TODO do something
	}

	return &AddOLTChassisReturn{DeviceID: in.GetHostname(), ChassisDeviceID: clli}, nil

}

/*
AddCard Adds an OLT card to an existing physical chassis, allocating ports
in the physical card to those in the abstract model
*/
func AddCard(physChassis *physicalModel.Chassis, olt physicalModel.OLT) error {
	physChassis.Linecards = append(physChassis.Linecards, olt)

	ports := olt.GetPorts()
	absChassis := (*models.GetAbstractChassisMap())[physChassis.CLLI]

	for i := 0; i < len(ports); i++ {
		absPort, _ := absChassis.NextPort()
		absPort.PhysPort = &ports[i]
		//AssignTraits(&ports[i], absPort)
	}

	//should probably worry about error at some point
	return nil
}
