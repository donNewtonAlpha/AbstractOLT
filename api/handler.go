package api

import (
	"fmt"
	"log"
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
Adds an OLT card to an existing physical chassis, allocating ports
in the physical card to those in the abstract model
*/
func AddCard(physChassis *physicalModel.Chassis, olt physicalModel.Olt) (error) {
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