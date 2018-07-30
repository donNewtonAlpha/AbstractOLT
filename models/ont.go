package models

/*
Ont represents a single ont/onu connect to a splitter on a Port
*/
type Ont struct {
	Number int
	Svlan  int
	Cvlan  int
	Parent *Port
}
