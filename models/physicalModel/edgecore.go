package physicalModel


//Implements the Edgecore linecard as an OLT
type Edgecore struct {
	SimpleOlt
}

func CreateEdgecore() *Edgecore {
	var newPorts [16]PONPort
	edge := Edgecore{}
	edge.Ports = newPorts[:]
	return &edge
}