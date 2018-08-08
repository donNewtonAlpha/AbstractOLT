package api

import (
	"log"

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
	log.Printf("new chassis %v\n", abstractChassis)
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
		//should probably worry about error at some point
		absPort.PhysPort = &ports[i]
		ports[i].AssignTraits(absPort)
	}
}

/* Assigns properties of the abstract port, such as svlan and cvlan,
to the physical port that it has been mapped to
*/ 
func AssignTraits(phys *physicalModel.PONPort, abs *abstractOLTModel.Port) {

}