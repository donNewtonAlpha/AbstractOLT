package physicalModel

/*
Edgecore Implements the Edgecore linecard as an OLT
*/
type Edgecore struct {
	SimpleOLT
}

/*
CreateEdgecore takes simple olt struct and generates Edgecore OLT
*/
func CreateEdgecore(olt *SimpleOLT) *Edgecore {
	var newPorts [16]PONPort
	edge := Edgecore{SimpleOLT: *olt}
	edge.Ports = newPorts[:]
	return &edge
}
